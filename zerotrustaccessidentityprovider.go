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

// ZeroTrustAccessIdentityProviderService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustAccessIdentityProviderService] method instead.
type ZeroTrustAccessIdentityProviderService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessIdentityProviderService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustAccessIdentityProviderService(opts ...option.RequestOption) (r *ZeroTrustAccessIdentityProviderService) {
	r = &ZeroTrustAccessIdentityProviderService{}
	r.Options = opts
	return
}

// Adds a new identity provider to Access.
func (r *ZeroTrustAccessIdentityProviderService) New(ctx context.Context, params ZeroTrustAccessIdentityProviderNewParams, opts ...option.RequestOption) (res *ZeroTrustAccessIdentityProviderNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessIdentityProviderNewResponseEnvelope
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
func (r *ZeroTrustAccessIdentityProviderService) Update(ctx context.Context, uuid string, params ZeroTrustAccessIdentityProviderUpdateParams, opts ...option.RequestOption) (res *ZeroTrustAccessIdentityProviderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessIdentityProviderUpdateResponseEnvelope
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
func (r *ZeroTrustAccessIdentityProviderService) List(ctx context.Context, query ZeroTrustAccessIdentityProviderListParams, opts ...option.RequestOption) (res *[]ZeroTrustAccessIdentityProviderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessIdentityProviderListResponseEnvelope
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
func (r *ZeroTrustAccessIdentityProviderService) Delete(ctx context.Context, uuid string, body ZeroTrustAccessIdentityProviderDeleteParams, opts ...option.RequestOption) (res *ZeroTrustAccessIdentityProviderDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessIdentityProviderDeleteResponseEnvelope
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
func (r *ZeroTrustAccessIdentityProviderService) Get(ctx context.Context, uuid string, query ZeroTrustAccessIdentityProviderGetParams, opts ...option.RequestOption) (res *ZeroTrustAccessIdentityProviderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessIdentityProviderGetResponseEnvelope
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

// Union satisfied by [ZeroTrustAccessIdentityProviderNewResponseAccessAzureAd],
// [ZeroTrustAccessIdentityProviderNewResponseAccessCentrify],
// [ZeroTrustAccessIdentityProviderNewResponseAccessFacebook],
// [ZeroTrustAccessIdentityProviderNewResponseAccessGitHub],
// [ZeroTrustAccessIdentityProviderNewResponseAccessGoogle],
// [ZeroTrustAccessIdentityProviderNewResponseAccessGoogleApps],
// [ZeroTrustAccessIdentityProviderNewResponseAccessLinkedin],
// [ZeroTrustAccessIdentityProviderNewResponseAccessOidc],
// [ZeroTrustAccessIdentityProviderNewResponseAccessOkta],
// [ZeroTrustAccessIdentityProviderNewResponseAccessOnelogin],
// [ZeroTrustAccessIdentityProviderNewResponseAccessPingone],
// [ZeroTrustAccessIdentityProviderNewResponseAccessSaml],
// [ZeroTrustAccessIdentityProviderNewResponseAccessYandex] or
// [ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepin].
type ZeroTrustAccessIdentityProviderNewResponse interface {
	implementsZeroTrustAccessIdentityProviderNewResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessIdentityProviderNewResponse)(nil)).Elem(), "")
}

type ZeroTrustAccessIdentityProviderNewResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessAzureAdJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessAzureAdJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessAzureAd]
type zeroTrustAccessIdentityProviderNewResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderNewResponseAccessAzureAd) implementsZeroTrustAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdConfig struct {
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
	JSON          zeroTrustAccessIdentityProviderNewResponseAccessAzureAdConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessAzureAdConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessAzureAdConfigJSON struct {
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

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType string

const (
	ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdTypeOnetimepin ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdTypeAzureAd    ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType = "azureAD"
	ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdTypeSaml       ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType = "saml"
	ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdTypeCentrify   ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType = "centrify"
	ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdTypeFacebook   ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType = "facebook"
	ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdTypeGitHub     ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType = "github"
	ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdTypeGoogleApps ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType = "google-apps"
	ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdTypeGoogle     ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType = "google"
	ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdTypeLinkedin   ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType = "linkedin"
	ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdTypeOidc       ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType = "oidc"
	ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdTypeOkta       ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType = "okta"
	ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdTypeOnelogin   ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType = "onelogin"
	ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdTypePingone    ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType = "pingone"
	ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdTypeYandex     ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessAzureAdScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdScimConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessCentrifyJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessCentrifyJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessCentrify]
type zeroTrustAccessIdentityProviderNewResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderNewResponseAccessCentrify) implementsZeroTrustAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyConfig struct {
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
	JSON           zeroTrustAccessIdentityProviderNewResponseAccessCentrifyConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessCentrifyConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType string

const (
	ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyTypeOnetimepin ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyTypeAzureAd    ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType = "azureAD"
	ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyTypeSaml       ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType = "saml"
	ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyTypeCentrify   ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType = "centrify"
	ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyTypeFacebook   ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType = "facebook"
	ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyTypeGitHub     ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType = "github"
	ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyTypeGoogleApps ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType = "google-apps"
	ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyTypeGoogle     ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType = "google"
	ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyTypeLinkedin   ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType = "linkedin"
	ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyTypeOidc       ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType = "oidc"
	ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyTypeOkta       ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType = "okta"
	ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyTypeOnelogin   ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType = "onelogin"
	ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyTypePingone    ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType = "pingone"
	ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyTypeYandex     ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessCentrifyScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyScimConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderNewResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderNewResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessFacebookJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessFacebookJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessFacebook]
type zeroTrustAccessIdentityProviderNewResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderNewResponseAccessFacebook) implementsZeroTrustAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                             `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderNewResponseAccessFacebookConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessFacebookConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessFacebookConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType string

const (
	ZeroTrustAccessIdentityProviderNewResponseAccessFacebookTypeOnetimepin ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewResponseAccessFacebookTypeAzureAd    ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType = "azureAD"
	ZeroTrustAccessIdentityProviderNewResponseAccessFacebookTypeSaml       ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType = "saml"
	ZeroTrustAccessIdentityProviderNewResponseAccessFacebookTypeCentrify   ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType = "centrify"
	ZeroTrustAccessIdentityProviderNewResponseAccessFacebookTypeFacebook   ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType = "facebook"
	ZeroTrustAccessIdentityProviderNewResponseAccessFacebookTypeGitHub     ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType = "github"
	ZeroTrustAccessIdentityProviderNewResponseAccessFacebookTypeGoogleApps ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType = "google-apps"
	ZeroTrustAccessIdentityProviderNewResponseAccessFacebookTypeGoogle     ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType = "google"
	ZeroTrustAccessIdentityProviderNewResponseAccessFacebookTypeLinkedin   ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType = "linkedin"
	ZeroTrustAccessIdentityProviderNewResponseAccessFacebookTypeOidc       ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType = "oidc"
	ZeroTrustAccessIdentityProviderNewResponseAccessFacebookTypeOkta       ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType = "okta"
	ZeroTrustAccessIdentityProviderNewResponseAccessFacebookTypeOnelogin   ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType = "onelogin"
	ZeroTrustAccessIdentityProviderNewResponseAccessFacebookTypePingone    ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType = "pingone"
	ZeroTrustAccessIdentityProviderNewResponseAccessFacebookTypeYandex     ZeroTrustAccessIdentityProviderNewResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderNewResponseAccessFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessFacebookScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessFacebookScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessFacebookScimConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderNewResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderNewResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessGitHubJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessGitHubJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderNewResponseAccessGitHub]
type zeroTrustAccessIdentityProviderNewResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderNewResponseAccessGitHub) implementsZeroTrustAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                           `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderNewResponseAccessGitHubConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessGitHubConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessGitHubConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType string

const (
	ZeroTrustAccessIdentityProviderNewResponseAccessGitHubTypeOnetimepin ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewResponseAccessGitHubTypeAzureAd    ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType = "azureAD"
	ZeroTrustAccessIdentityProviderNewResponseAccessGitHubTypeSaml       ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType = "saml"
	ZeroTrustAccessIdentityProviderNewResponseAccessGitHubTypeCentrify   ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType = "centrify"
	ZeroTrustAccessIdentityProviderNewResponseAccessGitHubTypeFacebook   ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType = "facebook"
	ZeroTrustAccessIdentityProviderNewResponseAccessGitHubTypeGitHub     ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType = "github"
	ZeroTrustAccessIdentityProviderNewResponseAccessGitHubTypeGoogleApps ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType = "google-apps"
	ZeroTrustAccessIdentityProviderNewResponseAccessGitHubTypeGoogle     ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType = "google"
	ZeroTrustAccessIdentityProviderNewResponseAccessGitHubTypeLinkedin   ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType = "linkedin"
	ZeroTrustAccessIdentityProviderNewResponseAccessGitHubTypeOidc       ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType = "oidc"
	ZeroTrustAccessIdentityProviderNewResponseAccessGitHubTypeOkta       ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType = "okta"
	ZeroTrustAccessIdentityProviderNewResponseAccessGitHubTypeOnelogin   ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType = "onelogin"
	ZeroTrustAccessIdentityProviderNewResponseAccessGitHubTypePingone    ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType = "pingone"
	ZeroTrustAccessIdentityProviderNewResponseAccessGitHubTypeYandex     ZeroTrustAccessIdentityProviderNewResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderNewResponseAccessGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessGitHubScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessGitHubScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessGitHubScimConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderNewResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderNewResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessGoogleJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessGoogleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderNewResponseAccessGoogle]
type zeroTrustAccessIdentityProviderNewResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderNewResponseAccessGoogle) implementsZeroTrustAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                           `json:"email_claim_name"`
	JSON           zeroTrustAccessIdentityProviderNewResponseAccessGoogleConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessGoogleConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessGoogleConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType string

const (
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleTypeOnetimepin ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleTypeAzureAd    ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType = "azureAD"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleTypeSaml       ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType = "saml"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleTypeCentrify   ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType = "centrify"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleTypeFacebook   ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType = "facebook"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleTypeGitHub     ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType = "github"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleTypeGoogleApps ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType = "google-apps"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleTypeGoogle     ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType = "google"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleTypeLinkedin   ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType = "linkedin"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleTypeOidc       ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType = "oidc"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleTypeOkta       ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType = "okta"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleTypeOnelogin   ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType = "onelogin"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleTypePingone    ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType = "pingone"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleTypeYandex     ZeroTrustAccessIdentityProviderNewResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderNewResponseAccessGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessGoogleScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessGoogleScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessGoogleScimConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessGoogleApps]
type zeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderNewResponseAccessGoogleApps) implementsZeroTrustAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsConfig struct {
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
	JSON           zeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType string

const (
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsTypeOnetimepin ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsTypeAzureAd    ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType = "azureAD"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsTypeSaml       ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType = "saml"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsTypeCentrify   ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType = "centrify"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsTypeFacebook   ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType = "facebook"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsTypeGitHub     ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType = "github"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsTypeGoogleApps ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType = "google-apps"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsTypeGoogle     ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType = "google"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsTypeLinkedin   ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType = "linkedin"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsTypeOidc       ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType = "oidc"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsTypeOkta       ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType = "okta"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsTypeOnelogin   ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType = "onelogin"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsTypePingone    ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType = "pingone"
	ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsTypeYandex     ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsScimConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessLinkedinJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessLinkedinJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessLinkedin]
type zeroTrustAccessIdentityProviderNewResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderNewResponseAccessLinkedin) implementsZeroTrustAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                             `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderNewResponseAccessLinkedinConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessLinkedinConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType string

const (
	ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinTypeOnetimepin ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinTypeAzureAd    ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType = "azureAD"
	ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinTypeSaml       ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType = "saml"
	ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinTypeCentrify   ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType = "centrify"
	ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinTypeFacebook   ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType = "facebook"
	ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinTypeGitHub     ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType = "github"
	ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinTypeGoogleApps ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType = "google-apps"
	ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinTypeGoogle     ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType = "google"
	ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinTypeLinkedin   ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType = "linkedin"
	ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinTypeOidc       ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType = "oidc"
	ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinTypeOkta       ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType = "okta"
	ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinTypeOnelogin   ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType = "onelogin"
	ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinTypePingone    ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType = "pingone"
	ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinTypeYandex     ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessLinkedinScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinScimConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderNewResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderNewResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderNewResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessOidcJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessOidcJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderNewResponseAccessOidc]
type zeroTrustAccessIdentityProviderNewResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderNewResponseAccessOidc) implementsZeroTrustAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessOidcConfig struct {
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
	JSON     zeroTrustAccessIdentityProviderNewResponseAccessOidcConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessOidcConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessOidcConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessOidcConfigJSON struct {
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

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessOidcType string

const (
	ZeroTrustAccessIdentityProviderNewResponseAccessOidcTypeOnetimepin ZeroTrustAccessIdentityProviderNewResponseAccessOidcType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewResponseAccessOidcTypeAzureAd    ZeroTrustAccessIdentityProviderNewResponseAccessOidcType = "azureAD"
	ZeroTrustAccessIdentityProviderNewResponseAccessOidcTypeSaml       ZeroTrustAccessIdentityProviderNewResponseAccessOidcType = "saml"
	ZeroTrustAccessIdentityProviderNewResponseAccessOidcTypeCentrify   ZeroTrustAccessIdentityProviderNewResponseAccessOidcType = "centrify"
	ZeroTrustAccessIdentityProviderNewResponseAccessOidcTypeFacebook   ZeroTrustAccessIdentityProviderNewResponseAccessOidcType = "facebook"
	ZeroTrustAccessIdentityProviderNewResponseAccessOidcTypeGitHub     ZeroTrustAccessIdentityProviderNewResponseAccessOidcType = "github"
	ZeroTrustAccessIdentityProviderNewResponseAccessOidcTypeGoogleApps ZeroTrustAccessIdentityProviderNewResponseAccessOidcType = "google-apps"
	ZeroTrustAccessIdentityProviderNewResponseAccessOidcTypeGoogle     ZeroTrustAccessIdentityProviderNewResponseAccessOidcType = "google"
	ZeroTrustAccessIdentityProviderNewResponseAccessOidcTypeLinkedin   ZeroTrustAccessIdentityProviderNewResponseAccessOidcType = "linkedin"
	ZeroTrustAccessIdentityProviderNewResponseAccessOidcTypeOidc       ZeroTrustAccessIdentityProviderNewResponseAccessOidcType = "oidc"
	ZeroTrustAccessIdentityProviderNewResponseAccessOidcTypeOkta       ZeroTrustAccessIdentityProviderNewResponseAccessOidcType = "okta"
	ZeroTrustAccessIdentityProviderNewResponseAccessOidcTypeOnelogin   ZeroTrustAccessIdentityProviderNewResponseAccessOidcType = "onelogin"
	ZeroTrustAccessIdentityProviderNewResponseAccessOidcTypePingone    ZeroTrustAccessIdentityProviderNewResponseAccessOidcType = "pingone"
	ZeroTrustAccessIdentityProviderNewResponseAccessOidcTypeYandex     ZeroTrustAccessIdentityProviderNewResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderNewResponseAccessOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessOidcScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessOidcScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessOidcScimConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderNewResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderNewResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderNewResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessOktaJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessOktaJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderNewResponseAccessOkta]
type zeroTrustAccessIdentityProviderNewResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderNewResponseAccessOkta) implementsZeroTrustAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessOktaConfig struct {
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
	JSON        zeroTrustAccessIdentityProviderNewResponseAccessOktaConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessOktaConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessOktaConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessOktaType string

const (
	ZeroTrustAccessIdentityProviderNewResponseAccessOktaTypeOnetimepin ZeroTrustAccessIdentityProviderNewResponseAccessOktaType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewResponseAccessOktaTypeAzureAd    ZeroTrustAccessIdentityProviderNewResponseAccessOktaType = "azureAD"
	ZeroTrustAccessIdentityProviderNewResponseAccessOktaTypeSaml       ZeroTrustAccessIdentityProviderNewResponseAccessOktaType = "saml"
	ZeroTrustAccessIdentityProviderNewResponseAccessOktaTypeCentrify   ZeroTrustAccessIdentityProviderNewResponseAccessOktaType = "centrify"
	ZeroTrustAccessIdentityProviderNewResponseAccessOktaTypeFacebook   ZeroTrustAccessIdentityProviderNewResponseAccessOktaType = "facebook"
	ZeroTrustAccessIdentityProviderNewResponseAccessOktaTypeGitHub     ZeroTrustAccessIdentityProviderNewResponseAccessOktaType = "github"
	ZeroTrustAccessIdentityProviderNewResponseAccessOktaTypeGoogleApps ZeroTrustAccessIdentityProviderNewResponseAccessOktaType = "google-apps"
	ZeroTrustAccessIdentityProviderNewResponseAccessOktaTypeGoogle     ZeroTrustAccessIdentityProviderNewResponseAccessOktaType = "google"
	ZeroTrustAccessIdentityProviderNewResponseAccessOktaTypeLinkedin   ZeroTrustAccessIdentityProviderNewResponseAccessOktaType = "linkedin"
	ZeroTrustAccessIdentityProviderNewResponseAccessOktaTypeOidc       ZeroTrustAccessIdentityProviderNewResponseAccessOktaType = "oidc"
	ZeroTrustAccessIdentityProviderNewResponseAccessOktaTypeOkta       ZeroTrustAccessIdentityProviderNewResponseAccessOktaType = "okta"
	ZeroTrustAccessIdentityProviderNewResponseAccessOktaTypeOnelogin   ZeroTrustAccessIdentityProviderNewResponseAccessOktaType = "onelogin"
	ZeroTrustAccessIdentityProviderNewResponseAccessOktaTypePingone    ZeroTrustAccessIdentityProviderNewResponseAccessOktaType = "pingone"
	ZeroTrustAccessIdentityProviderNewResponseAccessOktaTypeYandex     ZeroTrustAccessIdentityProviderNewResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderNewResponseAccessOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessOktaScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessOktaScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessOktaScimConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderNewResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderNewResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessOneloginJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessOneloginJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessOnelogin]
type zeroTrustAccessIdentityProviderNewResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderNewResponseAccessOnelogin) implementsZeroTrustAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessOneloginConfig struct {
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessOneloginConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessOneloginConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessOneloginConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType string

const (
	ZeroTrustAccessIdentityProviderNewResponseAccessOneloginTypeOnetimepin ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewResponseAccessOneloginTypeAzureAd    ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType = "azureAD"
	ZeroTrustAccessIdentityProviderNewResponseAccessOneloginTypeSaml       ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType = "saml"
	ZeroTrustAccessIdentityProviderNewResponseAccessOneloginTypeCentrify   ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType = "centrify"
	ZeroTrustAccessIdentityProviderNewResponseAccessOneloginTypeFacebook   ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType = "facebook"
	ZeroTrustAccessIdentityProviderNewResponseAccessOneloginTypeGitHub     ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType = "github"
	ZeroTrustAccessIdentityProviderNewResponseAccessOneloginTypeGoogleApps ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType = "google-apps"
	ZeroTrustAccessIdentityProviderNewResponseAccessOneloginTypeGoogle     ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType = "google"
	ZeroTrustAccessIdentityProviderNewResponseAccessOneloginTypeLinkedin   ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType = "linkedin"
	ZeroTrustAccessIdentityProviderNewResponseAccessOneloginTypeOidc       ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType = "oidc"
	ZeroTrustAccessIdentityProviderNewResponseAccessOneloginTypeOkta       ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType = "okta"
	ZeroTrustAccessIdentityProviderNewResponseAccessOneloginTypeOnelogin   ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType = "onelogin"
	ZeroTrustAccessIdentityProviderNewResponseAccessOneloginTypePingone    ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType = "pingone"
	ZeroTrustAccessIdentityProviderNewResponseAccessOneloginTypeYandex     ZeroTrustAccessIdentityProviderNewResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderNewResponseAccessOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessOneloginScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessOneloginScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessOneloginScimConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderNewResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderNewResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessPingoneJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessPingoneJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessPingone]
type zeroTrustAccessIdentityProviderNewResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderNewResponseAccessPingone) implementsZeroTrustAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessPingoneConfig struct {
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
	JSON      zeroTrustAccessIdentityProviderNewResponseAccessPingoneConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessPingoneConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessPingoneConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType string

const (
	ZeroTrustAccessIdentityProviderNewResponseAccessPingoneTypeOnetimepin ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewResponseAccessPingoneTypeAzureAd    ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType = "azureAD"
	ZeroTrustAccessIdentityProviderNewResponseAccessPingoneTypeSaml       ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType = "saml"
	ZeroTrustAccessIdentityProviderNewResponseAccessPingoneTypeCentrify   ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType = "centrify"
	ZeroTrustAccessIdentityProviderNewResponseAccessPingoneTypeFacebook   ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType = "facebook"
	ZeroTrustAccessIdentityProviderNewResponseAccessPingoneTypeGitHub     ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType = "github"
	ZeroTrustAccessIdentityProviderNewResponseAccessPingoneTypeGoogleApps ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType = "google-apps"
	ZeroTrustAccessIdentityProviderNewResponseAccessPingoneTypeGoogle     ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType = "google"
	ZeroTrustAccessIdentityProviderNewResponseAccessPingoneTypeLinkedin   ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType = "linkedin"
	ZeroTrustAccessIdentityProviderNewResponseAccessPingoneTypeOidc       ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType = "oidc"
	ZeroTrustAccessIdentityProviderNewResponseAccessPingoneTypeOkta       ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType = "okta"
	ZeroTrustAccessIdentityProviderNewResponseAccessPingoneTypeOnelogin   ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType = "onelogin"
	ZeroTrustAccessIdentityProviderNewResponseAccessPingoneTypePingone    ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType = "pingone"
	ZeroTrustAccessIdentityProviderNewResponseAccessPingoneTypeYandex     ZeroTrustAccessIdentityProviderNewResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderNewResponseAccessPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessPingoneScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessPingoneScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessPingoneScimConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderNewResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderNewResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderNewResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessSamlJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessSamlJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderNewResponseAccessSaml]
type zeroTrustAccessIdentityProviderNewResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderNewResponseAccessSaml) implementsZeroTrustAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []ZeroTrustAccessIdentityProviderNewResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                         `json:"sso_target_url"`
	JSON         zeroTrustAccessIdentityProviderNewResponseAccessSamlConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessSamlConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessSamlConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessSamlConfigJSON struct {
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

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                        `json:"header_name"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessSamlConfigHeaderAttribute]
type zeroTrustAccessIdentityProviderNewResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessSamlType string

const (
	ZeroTrustAccessIdentityProviderNewResponseAccessSamlTypeOnetimepin ZeroTrustAccessIdentityProviderNewResponseAccessSamlType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewResponseAccessSamlTypeAzureAd    ZeroTrustAccessIdentityProviderNewResponseAccessSamlType = "azureAD"
	ZeroTrustAccessIdentityProviderNewResponseAccessSamlTypeSaml       ZeroTrustAccessIdentityProviderNewResponseAccessSamlType = "saml"
	ZeroTrustAccessIdentityProviderNewResponseAccessSamlTypeCentrify   ZeroTrustAccessIdentityProviderNewResponseAccessSamlType = "centrify"
	ZeroTrustAccessIdentityProviderNewResponseAccessSamlTypeFacebook   ZeroTrustAccessIdentityProviderNewResponseAccessSamlType = "facebook"
	ZeroTrustAccessIdentityProviderNewResponseAccessSamlTypeGitHub     ZeroTrustAccessIdentityProviderNewResponseAccessSamlType = "github"
	ZeroTrustAccessIdentityProviderNewResponseAccessSamlTypeGoogleApps ZeroTrustAccessIdentityProviderNewResponseAccessSamlType = "google-apps"
	ZeroTrustAccessIdentityProviderNewResponseAccessSamlTypeGoogle     ZeroTrustAccessIdentityProviderNewResponseAccessSamlType = "google"
	ZeroTrustAccessIdentityProviderNewResponseAccessSamlTypeLinkedin   ZeroTrustAccessIdentityProviderNewResponseAccessSamlType = "linkedin"
	ZeroTrustAccessIdentityProviderNewResponseAccessSamlTypeOidc       ZeroTrustAccessIdentityProviderNewResponseAccessSamlType = "oidc"
	ZeroTrustAccessIdentityProviderNewResponseAccessSamlTypeOkta       ZeroTrustAccessIdentityProviderNewResponseAccessSamlType = "okta"
	ZeroTrustAccessIdentityProviderNewResponseAccessSamlTypeOnelogin   ZeroTrustAccessIdentityProviderNewResponseAccessSamlType = "onelogin"
	ZeroTrustAccessIdentityProviderNewResponseAccessSamlTypePingone    ZeroTrustAccessIdentityProviderNewResponseAccessSamlType = "pingone"
	ZeroTrustAccessIdentityProviderNewResponseAccessSamlTypeYandex     ZeroTrustAccessIdentityProviderNewResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderNewResponseAccessSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessSamlScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessSamlScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessSamlScimConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderNewResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderNewResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderNewResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessYandexJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessYandexJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderNewResponseAccessYandex]
type zeroTrustAccessIdentityProviderNewResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderNewResponseAccessYandex) implementsZeroTrustAccessIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                           `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderNewResponseAccessYandexConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessYandexConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessYandexConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessYandexType string

const (
	ZeroTrustAccessIdentityProviderNewResponseAccessYandexTypeOnetimepin ZeroTrustAccessIdentityProviderNewResponseAccessYandexType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewResponseAccessYandexTypeAzureAd    ZeroTrustAccessIdentityProviderNewResponseAccessYandexType = "azureAD"
	ZeroTrustAccessIdentityProviderNewResponseAccessYandexTypeSaml       ZeroTrustAccessIdentityProviderNewResponseAccessYandexType = "saml"
	ZeroTrustAccessIdentityProviderNewResponseAccessYandexTypeCentrify   ZeroTrustAccessIdentityProviderNewResponseAccessYandexType = "centrify"
	ZeroTrustAccessIdentityProviderNewResponseAccessYandexTypeFacebook   ZeroTrustAccessIdentityProviderNewResponseAccessYandexType = "facebook"
	ZeroTrustAccessIdentityProviderNewResponseAccessYandexTypeGitHub     ZeroTrustAccessIdentityProviderNewResponseAccessYandexType = "github"
	ZeroTrustAccessIdentityProviderNewResponseAccessYandexTypeGoogleApps ZeroTrustAccessIdentityProviderNewResponseAccessYandexType = "google-apps"
	ZeroTrustAccessIdentityProviderNewResponseAccessYandexTypeGoogle     ZeroTrustAccessIdentityProviderNewResponseAccessYandexType = "google"
	ZeroTrustAccessIdentityProviderNewResponseAccessYandexTypeLinkedin   ZeroTrustAccessIdentityProviderNewResponseAccessYandexType = "linkedin"
	ZeroTrustAccessIdentityProviderNewResponseAccessYandexTypeOidc       ZeroTrustAccessIdentityProviderNewResponseAccessYandexType = "oidc"
	ZeroTrustAccessIdentityProviderNewResponseAccessYandexTypeOkta       ZeroTrustAccessIdentityProviderNewResponseAccessYandexType = "okta"
	ZeroTrustAccessIdentityProviderNewResponseAccessYandexTypeOnelogin   ZeroTrustAccessIdentityProviderNewResponseAccessYandexType = "onelogin"
	ZeroTrustAccessIdentityProviderNewResponseAccessYandexTypePingone    ZeroTrustAccessIdentityProviderNewResponseAccessYandexType = "pingone"
	ZeroTrustAccessIdentityProviderNewResponseAccessYandexTypeYandex     ZeroTrustAccessIdentityProviderNewResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderNewResponseAccessYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessYandexScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessYandexScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessYandexScimConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderNewResponseAccessOnetimepinJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessOnetimepinJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepin]
type zeroTrustAccessIdentityProviderNewResponseAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepin) implementsZeroTrustAccessIdentityProviderNewResponse() {
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType string

const (
	ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinTypeOnetimepin ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinTypeAzureAd    ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType = "azureAD"
	ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinTypeSaml       ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType = "saml"
	ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinTypeCentrify   ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType = "centrify"
	ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinTypeFacebook   ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType = "facebook"
	ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinTypeGitHub     ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType = "github"
	ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinTypeGoogleApps ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType = "google-apps"
	ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinTypeGoogle     ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType = "google"
	ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinTypeLinkedin   ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType = "linkedin"
	ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinTypeOidc       ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType = "oidc"
	ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinTypeOkta       ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType = "okta"
	ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinTypeOnelogin   ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType = "onelogin"
	ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinTypePingone    ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType = "pingone"
	ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinTypeYandex     ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderNewResponseAccessOnetimepinScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseAccessOnetimepinScimConfigJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinScimConfig]
type zeroTrustAccessIdentityProviderNewResponseAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAd],
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrify],
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebook],
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHub],
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogle],
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleApps],
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedin],
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOidc],
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOkta],
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOnelogin],
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessPingone],
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessSaml],
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessYandex] or
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepin].
type ZeroTrustAccessIdentityProviderUpdateResponse interface {
	implementsZeroTrustAccessIdentityProviderUpdateResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessIdentityProviderUpdateResponse)(nil)).Elem(), "")
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAd]
type zeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAd) implementsZeroTrustAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdConfig struct {
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
	SupportGroups bool                                                                 `json:"support_groups"`
	JSON          zeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdConfigJSON struct {
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

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType string

const (
	ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdTypeSaml       ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType = "saml"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdTypeCentrify   ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdTypeFacebook   ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdTypeGitHub     ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType = "github"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdTypeGoogle     ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType = "google"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdTypeOidc       ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdTypeOkta       ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType = "okta"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdTypePingone    ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdTypeYandex     ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdScimConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrify]
type zeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrify) implementsZeroTrustAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyConfig struct {
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
	EmailClaimName string                                                                `json:"email_claim_name"`
	JSON           zeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType string

const (
	ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyTypeSaml       ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType = "saml"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyTypeCentrify   ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyTypeFacebook   ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyTypeGitHub     ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType = "github"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyTypeGoogle     ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType = "google"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyTypeOidc       ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyTypeOkta       ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType = "okta"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyTypePingone    ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyTypeYandex     ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                      `json:"user_deprovision"`
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyScimConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessFacebookJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessFacebookJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebook]
type zeroTrustAccessIdentityProviderUpdateResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebook) implementsZeroTrustAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderUpdateResponseAccessFacebookConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessFacebookConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType string

const (
	ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookTypeSaml       ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType = "saml"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookTypeCentrify   ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookTypeFacebook   ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookTypeGitHub     ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType = "github"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookTypeGoogle     ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType = "google"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookTypeOidc       ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookTypeOkta       ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType = "okta"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookTypePingone    ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookTypeYandex     ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                      `json:"user_deprovision"`
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessFacebookScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookScimConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessGitHubJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessGitHubJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHub]
type zeroTrustAccessIdentityProviderUpdateResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHub) implementsZeroTrustAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                              `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderUpdateResponseAccessGitHubConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessGitHubConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType string

const (
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubTypeSaml       ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType = "saml"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubTypeCentrify   ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubTypeFacebook   ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubTypeGitHub     ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType = "github"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubTypeGoogle     ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType = "google"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubTypeOidc       ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubTypeOkta       ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType = "okta"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubTypePingone    ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubTypeYandex     ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                    `json:"user_deprovision"`
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessGitHubScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessGitHubScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubScimConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogle]
type zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogle) implementsZeroTrustAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                              `json:"email_claim_name"`
	JSON           zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType string

const (
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleTypeSaml       ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType = "saml"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleTypeCentrify   ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleTypeFacebook   ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleTypeGitHub     ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType = "github"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleTypeGoogle     ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType = "google"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleTypeOidc       ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleTypeOkta       ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType = "okta"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleTypePingone    ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleTypeYandex     ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                    `json:"user_deprovision"`
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleScimConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleApps]
type zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleApps) implementsZeroTrustAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                  `json:"email_claim_name"`
	JSON           zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType string

const (
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsTypeSaml       ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType = "saml"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsTypeCentrify   ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsTypeFacebook   ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsTypeGitHub     ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType = "github"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsTypeGoogle     ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType = "google"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsTypeOidc       ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsTypeOkta       ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType = "okta"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsTypePingone    ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsTypeYandex     ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsScimConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedin]
type zeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedin) implementsZeroTrustAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType string

const (
	ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinTypeSaml       ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType = "saml"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinTypeCentrify   ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinTypeFacebook   ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinTypeGitHub     ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType = "github"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinTypeGoogle     ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType = "google"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinTypeOidc       ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinTypeOkta       ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType = "okta"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinTypePingone    ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinTypeYandex     ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                      `json:"user_deprovision"`
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinScimConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessOidcJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessOidcJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOidc]
type zeroTrustAccessIdentityProviderUpdateResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderUpdateResponseAccessOidc) implementsZeroTrustAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcConfig struct {
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
	TokenURL string                                                            `json:"token_url"`
	JSON     zeroTrustAccessIdentityProviderUpdateResponseAccessOidcConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessOidcConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessOidcConfigJSON struct {
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

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType string

const (
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcTypeSaml       ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType = "saml"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcTypeCentrify   ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcTypeFacebook   ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcTypeGitHub     ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType = "github"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcTypeGoogle     ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType = "google"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcTypeOidc       ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcTypeOkta       ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType = "okta"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcTypePingone    ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcTypeYandex     ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessOidcScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessOidcScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcScimConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessOktaJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessOktaJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOkta]
type zeroTrustAccessIdentityProviderUpdateResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderUpdateResponseAccessOkta) implementsZeroTrustAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaConfig struct {
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
	OktaAccount string                                                            `json:"okta_account"`
	JSON        zeroTrustAccessIdentityProviderUpdateResponseAccessOktaConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessOktaConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType string

const (
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaTypeSaml       ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType = "saml"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaTypeCentrify   ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaTypeFacebook   ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaTypeGitHub     ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType = "github"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaTypeGoogle     ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType = "google"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaTypeOidc       ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaTypeOkta       ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType = "okta"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaTypePingone    ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaTypeYandex     ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessOktaScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessOktaScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaScimConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessOneloginJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessOneloginJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOnelogin]
type zeroTrustAccessIdentityProviderUpdateResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderUpdateResponseAccessOnelogin) implementsZeroTrustAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                                `json:"onelogin_account"`
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessOneloginConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessOneloginConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType string

const (
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginTypeSaml       ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType = "saml"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginTypeCentrify   ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginTypeFacebook   ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginTypeGitHub     ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType = "github"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginTypeGoogle     ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType = "google"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginTypeOidc       ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginTypeOkta       ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType = "okta"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginTypePingone    ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginTypeYandex     ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                      `json:"user_deprovision"`
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessOneloginScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginScimConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessPingoneJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessPingoneJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessPingone]
type zeroTrustAccessIdentityProviderUpdateResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderUpdateResponseAccessPingone) implementsZeroTrustAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                               `json:"ping_env_id"`
	JSON      zeroTrustAccessIdentityProviderUpdateResponseAccessPingoneConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessPingoneConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType string

const (
	ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneTypeSaml       ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType = "saml"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneTypeCentrify   ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneTypeFacebook   ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneTypeGitHub     ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType = "github"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneTypeGoogle     ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType = "google"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneTypeOidc       ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneTypeOkta       ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType = "okta"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneTypePingone    ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneTypeYandex     ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessPingoneScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneScimConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessSamlJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessSamlJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessSaml]
type zeroTrustAccessIdentityProviderUpdateResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderUpdateResponseAccessSaml) implementsZeroTrustAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                            `json:"sso_target_url"`
	JSON         zeroTrustAccessIdentityProviderUpdateResponseAccessSamlConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessSamlConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessSamlConfigJSON struct {
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

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                           `json:"header_name"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute]
type zeroTrustAccessIdentityProviderUpdateResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType string

const (
	ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlTypeSaml       ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType = "saml"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlTypeCentrify   ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlTypeFacebook   ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlTypeGitHub     ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType = "github"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlTypeGoogle     ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType = "google"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlTypeOidc       ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlTypeOkta       ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType = "okta"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlTypePingone    ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlTypeYandex     ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessSamlScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessSamlScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlScimConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessYandexJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessYandexJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessYandex]
type zeroTrustAccessIdentityProviderUpdateResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderUpdateResponseAccessYandex) implementsZeroTrustAccessIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                              `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderUpdateResponseAccessYandexConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessYandexConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType string

const (
	ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexTypeSaml       ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType = "saml"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexTypeCentrify   ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexTypeFacebook   ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexTypeGitHub     ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType = "github"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexTypeGoogle     ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType = "google"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexTypeOidc       ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexTypeOkta       ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType = "okta"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexTypePingone    ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexTypeYandex     ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                    `json:"user_deprovision"`
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessYandexScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessYandexScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexScimConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepin]
type zeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepin) implementsZeroTrustAccessIdentityProviderUpdateResponse() {
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType string

const (
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinTypeSaml       ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType = "saml"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinTypeCentrify   ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinTypeFacebook   ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinTypeGitHub     ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType = "github"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinTypeGoogle     ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType = "google"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinTypeOidc       ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinTypeOkta       ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType = "okta"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinTypePingone    ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinTypeYandex     ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinScimConfigJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinScimConfig]
type zeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZeroTrustAccessIdentityProviderListResponseAccessAzureAd],
// [ZeroTrustAccessIdentityProviderListResponseAccessCentrify],
// [ZeroTrustAccessIdentityProviderListResponseAccessFacebook],
// [ZeroTrustAccessIdentityProviderListResponseAccessGitHub],
// [ZeroTrustAccessIdentityProviderListResponseAccessGoogle],
// [ZeroTrustAccessIdentityProviderListResponseAccessGoogleApps],
// [ZeroTrustAccessIdentityProviderListResponseAccessLinkedin],
// [ZeroTrustAccessIdentityProviderListResponseAccessOidc],
// [ZeroTrustAccessIdentityProviderListResponseAccessOkta],
// [ZeroTrustAccessIdentityProviderListResponseAccessOnelogin],
// [ZeroTrustAccessIdentityProviderListResponseAccessPingone],
// [ZeroTrustAccessIdentityProviderListResponseAccessSaml] or
// [ZeroTrustAccessIdentityProviderListResponseAccessYandex].
type ZeroTrustAccessIdentityProviderListResponse interface {
	implementsZeroTrustAccessIdentityProviderListResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessIdentityProviderListResponse)(nil)).Elem(), "")
}

type ZeroTrustAccessIdentityProviderListResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderListResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderListResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderListResponseAccessAzureAdJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessAzureAdJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessAzureAd]
type zeroTrustAccessIdentityProviderListResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderListResponseAccessAzureAd) implementsZeroTrustAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessAzureAdConfig struct {
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
	SupportGroups bool                                                               `json:"support_groups"`
	JSON          zeroTrustAccessIdentityProviderListResponseAccessAzureAdConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessAzureAdConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessAzureAdConfig]
type zeroTrustAccessIdentityProviderListResponseAccessAzureAdConfigJSON struct {
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

func (r *ZeroTrustAccessIdentityProviderListResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType string

const (
	ZeroTrustAccessIdentityProviderListResponseAccessAzureAdTypeOnetimepin ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType = "onetimepin"
	ZeroTrustAccessIdentityProviderListResponseAccessAzureAdTypeAzureAd    ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType = "azureAD"
	ZeroTrustAccessIdentityProviderListResponseAccessAzureAdTypeSaml       ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType = "saml"
	ZeroTrustAccessIdentityProviderListResponseAccessAzureAdTypeCentrify   ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType = "centrify"
	ZeroTrustAccessIdentityProviderListResponseAccessAzureAdTypeFacebook   ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType = "facebook"
	ZeroTrustAccessIdentityProviderListResponseAccessAzureAdTypeGitHub     ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType = "github"
	ZeroTrustAccessIdentityProviderListResponseAccessAzureAdTypeGoogleApps ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType = "google-apps"
	ZeroTrustAccessIdentityProviderListResponseAccessAzureAdTypeGoogle     ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType = "google"
	ZeroTrustAccessIdentityProviderListResponseAccessAzureAdTypeLinkedin   ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType = "linkedin"
	ZeroTrustAccessIdentityProviderListResponseAccessAzureAdTypeOidc       ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType = "oidc"
	ZeroTrustAccessIdentityProviderListResponseAccessAzureAdTypeOkta       ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType = "okta"
	ZeroTrustAccessIdentityProviderListResponseAccessAzureAdTypeOnelogin   ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType = "onelogin"
	ZeroTrustAccessIdentityProviderListResponseAccessAzureAdTypePingone    ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType = "pingone"
	ZeroTrustAccessIdentityProviderListResponseAccessAzureAdTypeYandex     ZeroTrustAccessIdentityProviderListResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderListResponseAccessAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderListResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessAzureAdScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessAzureAdScimConfig]
type zeroTrustAccessIdentityProviderListResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderListResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderListResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderListResponseAccessCentrifyJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessCentrifyJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessCentrify]
type zeroTrustAccessIdentityProviderListResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderListResponseAccessCentrify) implementsZeroTrustAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessCentrifyConfig struct {
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
	EmailClaimName string                                                              `json:"email_claim_name"`
	JSON           zeroTrustAccessIdentityProviderListResponseAccessCentrifyConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessCentrifyConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessCentrifyConfig]
type zeroTrustAccessIdentityProviderListResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType string

const (
	ZeroTrustAccessIdentityProviderListResponseAccessCentrifyTypeOnetimepin ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType = "onetimepin"
	ZeroTrustAccessIdentityProviderListResponseAccessCentrifyTypeAzureAd    ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType = "azureAD"
	ZeroTrustAccessIdentityProviderListResponseAccessCentrifyTypeSaml       ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType = "saml"
	ZeroTrustAccessIdentityProviderListResponseAccessCentrifyTypeCentrify   ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType = "centrify"
	ZeroTrustAccessIdentityProviderListResponseAccessCentrifyTypeFacebook   ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType = "facebook"
	ZeroTrustAccessIdentityProviderListResponseAccessCentrifyTypeGitHub     ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType = "github"
	ZeroTrustAccessIdentityProviderListResponseAccessCentrifyTypeGoogleApps ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType = "google-apps"
	ZeroTrustAccessIdentityProviderListResponseAccessCentrifyTypeGoogle     ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType = "google"
	ZeroTrustAccessIdentityProviderListResponseAccessCentrifyTypeLinkedin   ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType = "linkedin"
	ZeroTrustAccessIdentityProviderListResponseAccessCentrifyTypeOidc       ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType = "oidc"
	ZeroTrustAccessIdentityProviderListResponseAccessCentrifyTypeOkta       ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType = "okta"
	ZeroTrustAccessIdentityProviderListResponseAccessCentrifyTypeOnelogin   ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType = "onelogin"
	ZeroTrustAccessIdentityProviderListResponseAccessCentrifyTypePingone    ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType = "pingone"
	ZeroTrustAccessIdentityProviderListResponseAccessCentrifyTypeYandex     ZeroTrustAccessIdentityProviderListResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderListResponseAccessCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                    `json:"user_deprovision"`
	JSON            zeroTrustAccessIdentityProviderListResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessCentrifyScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessCentrifyScimConfig]
type zeroTrustAccessIdentityProviderListResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderListResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderListResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderListResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderListResponseAccessFacebookJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessFacebookJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessFacebook]
type zeroTrustAccessIdentityProviderListResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderListResponseAccessFacebook) implementsZeroTrustAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                              `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderListResponseAccessFacebookConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessFacebookConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessFacebookConfig]
type zeroTrustAccessIdentityProviderListResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessFacebookType string

const (
	ZeroTrustAccessIdentityProviderListResponseAccessFacebookTypeOnetimepin ZeroTrustAccessIdentityProviderListResponseAccessFacebookType = "onetimepin"
	ZeroTrustAccessIdentityProviderListResponseAccessFacebookTypeAzureAd    ZeroTrustAccessIdentityProviderListResponseAccessFacebookType = "azureAD"
	ZeroTrustAccessIdentityProviderListResponseAccessFacebookTypeSaml       ZeroTrustAccessIdentityProviderListResponseAccessFacebookType = "saml"
	ZeroTrustAccessIdentityProviderListResponseAccessFacebookTypeCentrify   ZeroTrustAccessIdentityProviderListResponseAccessFacebookType = "centrify"
	ZeroTrustAccessIdentityProviderListResponseAccessFacebookTypeFacebook   ZeroTrustAccessIdentityProviderListResponseAccessFacebookType = "facebook"
	ZeroTrustAccessIdentityProviderListResponseAccessFacebookTypeGitHub     ZeroTrustAccessIdentityProviderListResponseAccessFacebookType = "github"
	ZeroTrustAccessIdentityProviderListResponseAccessFacebookTypeGoogleApps ZeroTrustAccessIdentityProviderListResponseAccessFacebookType = "google-apps"
	ZeroTrustAccessIdentityProviderListResponseAccessFacebookTypeGoogle     ZeroTrustAccessIdentityProviderListResponseAccessFacebookType = "google"
	ZeroTrustAccessIdentityProviderListResponseAccessFacebookTypeLinkedin   ZeroTrustAccessIdentityProviderListResponseAccessFacebookType = "linkedin"
	ZeroTrustAccessIdentityProviderListResponseAccessFacebookTypeOidc       ZeroTrustAccessIdentityProviderListResponseAccessFacebookType = "oidc"
	ZeroTrustAccessIdentityProviderListResponseAccessFacebookTypeOkta       ZeroTrustAccessIdentityProviderListResponseAccessFacebookType = "okta"
	ZeroTrustAccessIdentityProviderListResponseAccessFacebookTypeOnelogin   ZeroTrustAccessIdentityProviderListResponseAccessFacebookType = "onelogin"
	ZeroTrustAccessIdentityProviderListResponseAccessFacebookTypePingone    ZeroTrustAccessIdentityProviderListResponseAccessFacebookType = "pingone"
	ZeroTrustAccessIdentityProviderListResponseAccessFacebookTypeYandex     ZeroTrustAccessIdentityProviderListResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderListResponseAccessFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                    `json:"user_deprovision"`
	JSON            zeroTrustAccessIdentityProviderListResponseAccessFacebookScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessFacebookScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessFacebookScimConfig]
type zeroTrustAccessIdentityProviderListResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderListResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderListResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderListResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderListResponseAccessGitHubJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessGitHubJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessGitHub]
type zeroTrustAccessIdentityProviderListResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderListResponseAccessGitHub) implementsZeroTrustAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                            `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderListResponseAccessGitHubConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessGitHubConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessGitHubConfig]
type zeroTrustAccessIdentityProviderListResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessGitHubType string

const (
	ZeroTrustAccessIdentityProviderListResponseAccessGitHubTypeOnetimepin ZeroTrustAccessIdentityProviderListResponseAccessGitHubType = "onetimepin"
	ZeroTrustAccessIdentityProviderListResponseAccessGitHubTypeAzureAd    ZeroTrustAccessIdentityProviderListResponseAccessGitHubType = "azureAD"
	ZeroTrustAccessIdentityProviderListResponseAccessGitHubTypeSaml       ZeroTrustAccessIdentityProviderListResponseAccessGitHubType = "saml"
	ZeroTrustAccessIdentityProviderListResponseAccessGitHubTypeCentrify   ZeroTrustAccessIdentityProviderListResponseAccessGitHubType = "centrify"
	ZeroTrustAccessIdentityProviderListResponseAccessGitHubTypeFacebook   ZeroTrustAccessIdentityProviderListResponseAccessGitHubType = "facebook"
	ZeroTrustAccessIdentityProviderListResponseAccessGitHubTypeGitHub     ZeroTrustAccessIdentityProviderListResponseAccessGitHubType = "github"
	ZeroTrustAccessIdentityProviderListResponseAccessGitHubTypeGoogleApps ZeroTrustAccessIdentityProviderListResponseAccessGitHubType = "google-apps"
	ZeroTrustAccessIdentityProviderListResponseAccessGitHubTypeGoogle     ZeroTrustAccessIdentityProviderListResponseAccessGitHubType = "google"
	ZeroTrustAccessIdentityProviderListResponseAccessGitHubTypeLinkedin   ZeroTrustAccessIdentityProviderListResponseAccessGitHubType = "linkedin"
	ZeroTrustAccessIdentityProviderListResponseAccessGitHubTypeOidc       ZeroTrustAccessIdentityProviderListResponseAccessGitHubType = "oidc"
	ZeroTrustAccessIdentityProviderListResponseAccessGitHubTypeOkta       ZeroTrustAccessIdentityProviderListResponseAccessGitHubType = "okta"
	ZeroTrustAccessIdentityProviderListResponseAccessGitHubTypeOnelogin   ZeroTrustAccessIdentityProviderListResponseAccessGitHubType = "onelogin"
	ZeroTrustAccessIdentityProviderListResponseAccessGitHubTypePingone    ZeroTrustAccessIdentityProviderListResponseAccessGitHubType = "pingone"
	ZeroTrustAccessIdentityProviderListResponseAccessGitHubTypeYandex     ZeroTrustAccessIdentityProviderListResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderListResponseAccessGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderListResponseAccessGitHubScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessGitHubScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessGitHubScimConfig]
type zeroTrustAccessIdentityProviderListResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderListResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderListResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderListResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderListResponseAccessGoogleJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessGoogleJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessGoogle]
type zeroTrustAccessIdentityProviderListResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderListResponseAccessGoogle) implementsZeroTrustAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                            `json:"email_claim_name"`
	JSON           zeroTrustAccessIdentityProviderListResponseAccessGoogleConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessGoogleConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessGoogleConfig]
type zeroTrustAccessIdentityProviderListResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessGoogleType string

const (
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleTypeOnetimepin ZeroTrustAccessIdentityProviderListResponseAccessGoogleType = "onetimepin"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleTypeAzureAd    ZeroTrustAccessIdentityProviderListResponseAccessGoogleType = "azureAD"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleTypeSaml       ZeroTrustAccessIdentityProviderListResponseAccessGoogleType = "saml"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleTypeCentrify   ZeroTrustAccessIdentityProviderListResponseAccessGoogleType = "centrify"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleTypeFacebook   ZeroTrustAccessIdentityProviderListResponseAccessGoogleType = "facebook"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleTypeGitHub     ZeroTrustAccessIdentityProviderListResponseAccessGoogleType = "github"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleTypeGoogleApps ZeroTrustAccessIdentityProviderListResponseAccessGoogleType = "google-apps"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleTypeGoogle     ZeroTrustAccessIdentityProviderListResponseAccessGoogleType = "google"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleTypeLinkedin   ZeroTrustAccessIdentityProviderListResponseAccessGoogleType = "linkedin"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleTypeOidc       ZeroTrustAccessIdentityProviderListResponseAccessGoogleType = "oidc"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleTypeOkta       ZeroTrustAccessIdentityProviderListResponseAccessGoogleType = "okta"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleTypeOnelogin   ZeroTrustAccessIdentityProviderListResponseAccessGoogleType = "onelogin"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleTypePingone    ZeroTrustAccessIdentityProviderListResponseAccessGoogleType = "pingone"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleTypeYandex     ZeroTrustAccessIdentityProviderListResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderListResponseAccessGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderListResponseAccessGoogleScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessGoogleScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessGoogleScimConfig]
type zeroTrustAccessIdentityProviderListResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderListResponseAccessGoogleAppsJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessGoogleAppsJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessGoogleApps]
type zeroTrustAccessIdentityProviderListResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderListResponseAccessGoogleApps) implementsZeroTrustAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                                `json:"email_claim_name"`
	JSON           zeroTrustAccessIdentityProviderListResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessGoogleAppsConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsConfig]
type zeroTrustAccessIdentityProviderListResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType string

const (
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsTypeOnetimepin ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType = "onetimepin"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsTypeAzureAd    ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType = "azureAD"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsTypeSaml       ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType = "saml"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsTypeCentrify   ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType = "centrify"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsTypeFacebook   ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType = "facebook"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsTypeGitHub     ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType = "github"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsTypeGoogleApps ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType = "google-apps"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsTypeGoogle     ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType = "google"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsTypeLinkedin   ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType = "linkedin"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsTypeOidc       ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType = "oidc"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsTypeOkta       ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType = "okta"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsTypeOnelogin   ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType = "onelogin"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsTypePingone    ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType = "pingone"
	ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsTypeYandex     ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                      `json:"user_deprovision"`
	JSON            zeroTrustAccessIdentityProviderListResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsScimConfig]
type zeroTrustAccessIdentityProviderListResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderListResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderListResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderListResponseAccessLinkedinJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessLinkedinJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessLinkedin]
type zeroTrustAccessIdentityProviderListResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderListResponseAccessLinkedin) implementsZeroTrustAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                              `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderListResponseAccessLinkedinConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessLinkedinConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessLinkedinConfig]
type zeroTrustAccessIdentityProviderListResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType string

const (
	ZeroTrustAccessIdentityProviderListResponseAccessLinkedinTypeOnetimepin ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType = "onetimepin"
	ZeroTrustAccessIdentityProviderListResponseAccessLinkedinTypeAzureAd    ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType = "azureAD"
	ZeroTrustAccessIdentityProviderListResponseAccessLinkedinTypeSaml       ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType = "saml"
	ZeroTrustAccessIdentityProviderListResponseAccessLinkedinTypeCentrify   ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType = "centrify"
	ZeroTrustAccessIdentityProviderListResponseAccessLinkedinTypeFacebook   ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType = "facebook"
	ZeroTrustAccessIdentityProviderListResponseAccessLinkedinTypeGitHub     ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType = "github"
	ZeroTrustAccessIdentityProviderListResponseAccessLinkedinTypeGoogleApps ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType = "google-apps"
	ZeroTrustAccessIdentityProviderListResponseAccessLinkedinTypeGoogle     ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType = "google"
	ZeroTrustAccessIdentityProviderListResponseAccessLinkedinTypeLinkedin   ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType = "linkedin"
	ZeroTrustAccessIdentityProviderListResponseAccessLinkedinTypeOidc       ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType = "oidc"
	ZeroTrustAccessIdentityProviderListResponseAccessLinkedinTypeOkta       ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType = "okta"
	ZeroTrustAccessIdentityProviderListResponseAccessLinkedinTypeOnelogin   ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType = "onelogin"
	ZeroTrustAccessIdentityProviderListResponseAccessLinkedinTypePingone    ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType = "pingone"
	ZeroTrustAccessIdentityProviderListResponseAccessLinkedinTypeYandex     ZeroTrustAccessIdentityProviderListResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderListResponseAccessLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                    `json:"user_deprovision"`
	JSON            zeroTrustAccessIdentityProviderListResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessLinkedinScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessLinkedinScimConfig]
type zeroTrustAccessIdentityProviderListResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderListResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderListResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderListResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderListResponseAccessOidcJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessOidcJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderListResponseAccessOidc]
type zeroTrustAccessIdentityProviderListResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderListResponseAccessOidc) implementsZeroTrustAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessOidcConfig struct {
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
	TokenURL string                                                          `json:"token_url"`
	JSON     zeroTrustAccessIdentityProviderListResponseAccessOidcConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessOidcConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessOidcConfig]
type zeroTrustAccessIdentityProviderListResponseAccessOidcConfigJSON struct {
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

func (r *ZeroTrustAccessIdentityProviderListResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessOidcType string

const (
	ZeroTrustAccessIdentityProviderListResponseAccessOidcTypeOnetimepin ZeroTrustAccessIdentityProviderListResponseAccessOidcType = "onetimepin"
	ZeroTrustAccessIdentityProviderListResponseAccessOidcTypeAzureAd    ZeroTrustAccessIdentityProviderListResponseAccessOidcType = "azureAD"
	ZeroTrustAccessIdentityProviderListResponseAccessOidcTypeSaml       ZeroTrustAccessIdentityProviderListResponseAccessOidcType = "saml"
	ZeroTrustAccessIdentityProviderListResponseAccessOidcTypeCentrify   ZeroTrustAccessIdentityProviderListResponseAccessOidcType = "centrify"
	ZeroTrustAccessIdentityProviderListResponseAccessOidcTypeFacebook   ZeroTrustAccessIdentityProviderListResponseAccessOidcType = "facebook"
	ZeroTrustAccessIdentityProviderListResponseAccessOidcTypeGitHub     ZeroTrustAccessIdentityProviderListResponseAccessOidcType = "github"
	ZeroTrustAccessIdentityProviderListResponseAccessOidcTypeGoogleApps ZeroTrustAccessIdentityProviderListResponseAccessOidcType = "google-apps"
	ZeroTrustAccessIdentityProviderListResponseAccessOidcTypeGoogle     ZeroTrustAccessIdentityProviderListResponseAccessOidcType = "google"
	ZeroTrustAccessIdentityProviderListResponseAccessOidcTypeLinkedin   ZeroTrustAccessIdentityProviderListResponseAccessOidcType = "linkedin"
	ZeroTrustAccessIdentityProviderListResponseAccessOidcTypeOidc       ZeroTrustAccessIdentityProviderListResponseAccessOidcType = "oidc"
	ZeroTrustAccessIdentityProviderListResponseAccessOidcTypeOkta       ZeroTrustAccessIdentityProviderListResponseAccessOidcType = "okta"
	ZeroTrustAccessIdentityProviderListResponseAccessOidcTypeOnelogin   ZeroTrustAccessIdentityProviderListResponseAccessOidcType = "onelogin"
	ZeroTrustAccessIdentityProviderListResponseAccessOidcTypePingone    ZeroTrustAccessIdentityProviderListResponseAccessOidcType = "pingone"
	ZeroTrustAccessIdentityProviderListResponseAccessOidcTypeYandex     ZeroTrustAccessIdentityProviderListResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderListResponseAccessOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderListResponseAccessOidcScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessOidcScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessOidcScimConfig]
type zeroTrustAccessIdentityProviderListResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderListResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderListResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderListResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderListResponseAccessOktaJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessOktaJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderListResponseAccessOkta]
type zeroTrustAccessIdentityProviderListResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderListResponseAccessOkta) implementsZeroTrustAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessOktaConfig struct {
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
	OktaAccount string                                                          `json:"okta_account"`
	JSON        zeroTrustAccessIdentityProviderListResponseAccessOktaConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessOktaConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessOktaConfig]
type zeroTrustAccessIdentityProviderListResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessOktaType string

const (
	ZeroTrustAccessIdentityProviderListResponseAccessOktaTypeOnetimepin ZeroTrustAccessIdentityProviderListResponseAccessOktaType = "onetimepin"
	ZeroTrustAccessIdentityProviderListResponseAccessOktaTypeAzureAd    ZeroTrustAccessIdentityProviderListResponseAccessOktaType = "azureAD"
	ZeroTrustAccessIdentityProviderListResponseAccessOktaTypeSaml       ZeroTrustAccessIdentityProviderListResponseAccessOktaType = "saml"
	ZeroTrustAccessIdentityProviderListResponseAccessOktaTypeCentrify   ZeroTrustAccessIdentityProviderListResponseAccessOktaType = "centrify"
	ZeroTrustAccessIdentityProviderListResponseAccessOktaTypeFacebook   ZeroTrustAccessIdentityProviderListResponseAccessOktaType = "facebook"
	ZeroTrustAccessIdentityProviderListResponseAccessOktaTypeGitHub     ZeroTrustAccessIdentityProviderListResponseAccessOktaType = "github"
	ZeroTrustAccessIdentityProviderListResponseAccessOktaTypeGoogleApps ZeroTrustAccessIdentityProviderListResponseAccessOktaType = "google-apps"
	ZeroTrustAccessIdentityProviderListResponseAccessOktaTypeGoogle     ZeroTrustAccessIdentityProviderListResponseAccessOktaType = "google"
	ZeroTrustAccessIdentityProviderListResponseAccessOktaTypeLinkedin   ZeroTrustAccessIdentityProviderListResponseAccessOktaType = "linkedin"
	ZeroTrustAccessIdentityProviderListResponseAccessOktaTypeOidc       ZeroTrustAccessIdentityProviderListResponseAccessOktaType = "oidc"
	ZeroTrustAccessIdentityProviderListResponseAccessOktaTypeOkta       ZeroTrustAccessIdentityProviderListResponseAccessOktaType = "okta"
	ZeroTrustAccessIdentityProviderListResponseAccessOktaTypeOnelogin   ZeroTrustAccessIdentityProviderListResponseAccessOktaType = "onelogin"
	ZeroTrustAccessIdentityProviderListResponseAccessOktaTypePingone    ZeroTrustAccessIdentityProviderListResponseAccessOktaType = "pingone"
	ZeroTrustAccessIdentityProviderListResponseAccessOktaTypeYandex     ZeroTrustAccessIdentityProviderListResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderListResponseAccessOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderListResponseAccessOktaScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessOktaScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessOktaScimConfig]
type zeroTrustAccessIdentityProviderListResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderListResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderListResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderListResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderListResponseAccessOneloginJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessOneloginJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessOnelogin]
type zeroTrustAccessIdentityProviderListResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderListResponseAccessOnelogin) implementsZeroTrustAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                              `json:"onelogin_account"`
	JSON            zeroTrustAccessIdentityProviderListResponseAccessOneloginConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessOneloginConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessOneloginConfig]
type zeroTrustAccessIdentityProviderListResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessOneloginType string

const (
	ZeroTrustAccessIdentityProviderListResponseAccessOneloginTypeOnetimepin ZeroTrustAccessIdentityProviderListResponseAccessOneloginType = "onetimepin"
	ZeroTrustAccessIdentityProviderListResponseAccessOneloginTypeAzureAd    ZeroTrustAccessIdentityProviderListResponseAccessOneloginType = "azureAD"
	ZeroTrustAccessIdentityProviderListResponseAccessOneloginTypeSaml       ZeroTrustAccessIdentityProviderListResponseAccessOneloginType = "saml"
	ZeroTrustAccessIdentityProviderListResponseAccessOneloginTypeCentrify   ZeroTrustAccessIdentityProviderListResponseAccessOneloginType = "centrify"
	ZeroTrustAccessIdentityProviderListResponseAccessOneloginTypeFacebook   ZeroTrustAccessIdentityProviderListResponseAccessOneloginType = "facebook"
	ZeroTrustAccessIdentityProviderListResponseAccessOneloginTypeGitHub     ZeroTrustAccessIdentityProviderListResponseAccessOneloginType = "github"
	ZeroTrustAccessIdentityProviderListResponseAccessOneloginTypeGoogleApps ZeroTrustAccessIdentityProviderListResponseAccessOneloginType = "google-apps"
	ZeroTrustAccessIdentityProviderListResponseAccessOneloginTypeGoogle     ZeroTrustAccessIdentityProviderListResponseAccessOneloginType = "google"
	ZeroTrustAccessIdentityProviderListResponseAccessOneloginTypeLinkedin   ZeroTrustAccessIdentityProviderListResponseAccessOneloginType = "linkedin"
	ZeroTrustAccessIdentityProviderListResponseAccessOneloginTypeOidc       ZeroTrustAccessIdentityProviderListResponseAccessOneloginType = "oidc"
	ZeroTrustAccessIdentityProviderListResponseAccessOneloginTypeOkta       ZeroTrustAccessIdentityProviderListResponseAccessOneloginType = "okta"
	ZeroTrustAccessIdentityProviderListResponseAccessOneloginTypeOnelogin   ZeroTrustAccessIdentityProviderListResponseAccessOneloginType = "onelogin"
	ZeroTrustAccessIdentityProviderListResponseAccessOneloginTypePingone    ZeroTrustAccessIdentityProviderListResponseAccessOneloginType = "pingone"
	ZeroTrustAccessIdentityProviderListResponseAccessOneloginTypeYandex     ZeroTrustAccessIdentityProviderListResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderListResponseAccessOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                    `json:"user_deprovision"`
	JSON            zeroTrustAccessIdentityProviderListResponseAccessOneloginScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessOneloginScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessOneloginScimConfig]
type zeroTrustAccessIdentityProviderListResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderListResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderListResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderListResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderListResponseAccessPingoneJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessPingoneJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessPingone]
type zeroTrustAccessIdentityProviderListResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderListResponseAccessPingone) implementsZeroTrustAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                             `json:"ping_env_id"`
	JSON      zeroTrustAccessIdentityProviderListResponseAccessPingoneConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessPingoneConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessPingoneConfig]
type zeroTrustAccessIdentityProviderListResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessPingoneType string

const (
	ZeroTrustAccessIdentityProviderListResponseAccessPingoneTypeOnetimepin ZeroTrustAccessIdentityProviderListResponseAccessPingoneType = "onetimepin"
	ZeroTrustAccessIdentityProviderListResponseAccessPingoneTypeAzureAd    ZeroTrustAccessIdentityProviderListResponseAccessPingoneType = "azureAD"
	ZeroTrustAccessIdentityProviderListResponseAccessPingoneTypeSaml       ZeroTrustAccessIdentityProviderListResponseAccessPingoneType = "saml"
	ZeroTrustAccessIdentityProviderListResponseAccessPingoneTypeCentrify   ZeroTrustAccessIdentityProviderListResponseAccessPingoneType = "centrify"
	ZeroTrustAccessIdentityProviderListResponseAccessPingoneTypeFacebook   ZeroTrustAccessIdentityProviderListResponseAccessPingoneType = "facebook"
	ZeroTrustAccessIdentityProviderListResponseAccessPingoneTypeGitHub     ZeroTrustAccessIdentityProviderListResponseAccessPingoneType = "github"
	ZeroTrustAccessIdentityProviderListResponseAccessPingoneTypeGoogleApps ZeroTrustAccessIdentityProviderListResponseAccessPingoneType = "google-apps"
	ZeroTrustAccessIdentityProviderListResponseAccessPingoneTypeGoogle     ZeroTrustAccessIdentityProviderListResponseAccessPingoneType = "google"
	ZeroTrustAccessIdentityProviderListResponseAccessPingoneTypeLinkedin   ZeroTrustAccessIdentityProviderListResponseAccessPingoneType = "linkedin"
	ZeroTrustAccessIdentityProviderListResponseAccessPingoneTypeOidc       ZeroTrustAccessIdentityProviderListResponseAccessPingoneType = "oidc"
	ZeroTrustAccessIdentityProviderListResponseAccessPingoneTypeOkta       ZeroTrustAccessIdentityProviderListResponseAccessPingoneType = "okta"
	ZeroTrustAccessIdentityProviderListResponseAccessPingoneTypeOnelogin   ZeroTrustAccessIdentityProviderListResponseAccessPingoneType = "onelogin"
	ZeroTrustAccessIdentityProviderListResponseAccessPingoneTypePingone    ZeroTrustAccessIdentityProviderListResponseAccessPingoneType = "pingone"
	ZeroTrustAccessIdentityProviderListResponseAccessPingoneTypeYandex     ZeroTrustAccessIdentityProviderListResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderListResponseAccessPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderListResponseAccessPingoneScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessPingoneScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessPingoneScimConfig]
type zeroTrustAccessIdentityProviderListResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderListResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderListResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderListResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderListResponseAccessSamlJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessSamlJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderListResponseAccessSaml]
type zeroTrustAccessIdentityProviderListResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderListResponseAccessSaml) implementsZeroTrustAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []ZeroTrustAccessIdentityProviderListResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                          `json:"sso_target_url"`
	JSON         zeroTrustAccessIdentityProviderListResponseAccessSamlConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessSamlConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessSamlConfig]
type zeroTrustAccessIdentityProviderListResponseAccessSamlConfigJSON struct {
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

func (r *ZeroTrustAccessIdentityProviderListResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                         `json:"header_name"`
	JSON       zeroTrustAccessIdentityProviderListResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessSamlConfigHeaderAttribute]
type zeroTrustAccessIdentityProviderListResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessSamlType string

const (
	ZeroTrustAccessIdentityProviderListResponseAccessSamlTypeOnetimepin ZeroTrustAccessIdentityProviderListResponseAccessSamlType = "onetimepin"
	ZeroTrustAccessIdentityProviderListResponseAccessSamlTypeAzureAd    ZeroTrustAccessIdentityProviderListResponseAccessSamlType = "azureAD"
	ZeroTrustAccessIdentityProviderListResponseAccessSamlTypeSaml       ZeroTrustAccessIdentityProviderListResponseAccessSamlType = "saml"
	ZeroTrustAccessIdentityProviderListResponseAccessSamlTypeCentrify   ZeroTrustAccessIdentityProviderListResponseAccessSamlType = "centrify"
	ZeroTrustAccessIdentityProviderListResponseAccessSamlTypeFacebook   ZeroTrustAccessIdentityProviderListResponseAccessSamlType = "facebook"
	ZeroTrustAccessIdentityProviderListResponseAccessSamlTypeGitHub     ZeroTrustAccessIdentityProviderListResponseAccessSamlType = "github"
	ZeroTrustAccessIdentityProviderListResponseAccessSamlTypeGoogleApps ZeroTrustAccessIdentityProviderListResponseAccessSamlType = "google-apps"
	ZeroTrustAccessIdentityProviderListResponseAccessSamlTypeGoogle     ZeroTrustAccessIdentityProviderListResponseAccessSamlType = "google"
	ZeroTrustAccessIdentityProviderListResponseAccessSamlTypeLinkedin   ZeroTrustAccessIdentityProviderListResponseAccessSamlType = "linkedin"
	ZeroTrustAccessIdentityProviderListResponseAccessSamlTypeOidc       ZeroTrustAccessIdentityProviderListResponseAccessSamlType = "oidc"
	ZeroTrustAccessIdentityProviderListResponseAccessSamlTypeOkta       ZeroTrustAccessIdentityProviderListResponseAccessSamlType = "okta"
	ZeroTrustAccessIdentityProviderListResponseAccessSamlTypeOnelogin   ZeroTrustAccessIdentityProviderListResponseAccessSamlType = "onelogin"
	ZeroTrustAccessIdentityProviderListResponseAccessSamlTypePingone    ZeroTrustAccessIdentityProviderListResponseAccessSamlType = "pingone"
	ZeroTrustAccessIdentityProviderListResponseAccessSamlTypeYandex     ZeroTrustAccessIdentityProviderListResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderListResponseAccessSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderListResponseAccessSamlScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessSamlScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessSamlScimConfig]
type zeroTrustAccessIdentityProviderListResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderListResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderListResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderListResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderListResponseAccessYandexJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessYandexJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessYandex]
type zeroTrustAccessIdentityProviderListResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderListResponseAccessYandex) implementsZeroTrustAccessIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                            `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderListResponseAccessYandexConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessYandexConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessYandexConfig]
type zeroTrustAccessIdentityProviderListResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderListResponseAccessYandexType string

const (
	ZeroTrustAccessIdentityProviderListResponseAccessYandexTypeOnetimepin ZeroTrustAccessIdentityProviderListResponseAccessYandexType = "onetimepin"
	ZeroTrustAccessIdentityProviderListResponseAccessYandexTypeAzureAd    ZeroTrustAccessIdentityProviderListResponseAccessYandexType = "azureAD"
	ZeroTrustAccessIdentityProviderListResponseAccessYandexTypeSaml       ZeroTrustAccessIdentityProviderListResponseAccessYandexType = "saml"
	ZeroTrustAccessIdentityProviderListResponseAccessYandexTypeCentrify   ZeroTrustAccessIdentityProviderListResponseAccessYandexType = "centrify"
	ZeroTrustAccessIdentityProviderListResponseAccessYandexTypeFacebook   ZeroTrustAccessIdentityProviderListResponseAccessYandexType = "facebook"
	ZeroTrustAccessIdentityProviderListResponseAccessYandexTypeGitHub     ZeroTrustAccessIdentityProviderListResponseAccessYandexType = "github"
	ZeroTrustAccessIdentityProviderListResponseAccessYandexTypeGoogleApps ZeroTrustAccessIdentityProviderListResponseAccessYandexType = "google-apps"
	ZeroTrustAccessIdentityProviderListResponseAccessYandexTypeGoogle     ZeroTrustAccessIdentityProviderListResponseAccessYandexType = "google"
	ZeroTrustAccessIdentityProviderListResponseAccessYandexTypeLinkedin   ZeroTrustAccessIdentityProviderListResponseAccessYandexType = "linkedin"
	ZeroTrustAccessIdentityProviderListResponseAccessYandexTypeOidc       ZeroTrustAccessIdentityProviderListResponseAccessYandexType = "oidc"
	ZeroTrustAccessIdentityProviderListResponseAccessYandexTypeOkta       ZeroTrustAccessIdentityProviderListResponseAccessYandexType = "okta"
	ZeroTrustAccessIdentityProviderListResponseAccessYandexTypeOnelogin   ZeroTrustAccessIdentityProviderListResponseAccessYandexType = "onelogin"
	ZeroTrustAccessIdentityProviderListResponseAccessYandexTypePingone    ZeroTrustAccessIdentityProviderListResponseAccessYandexType = "pingone"
	ZeroTrustAccessIdentityProviderListResponseAccessYandexTypeYandex     ZeroTrustAccessIdentityProviderListResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderListResponseAccessYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderListResponseAccessYandexScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseAccessYandexScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseAccessYandexScimConfig]
type zeroTrustAccessIdentityProviderListResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderDeleteResponse struct {
	// UUID
	ID   string                                            `json:"id"`
	JSON zeroTrustAccessIdentityProviderDeleteResponseJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderDeleteResponseJSON contains the JSON metadata for
// the struct [ZeroTrustAccessIdentityProviderDeleteResponse]
type zeroTrustAccessIdentityProviderDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZeroTrustAccessIdentityProviderGetResponseAccessAzureAd],
// [ZeroTrustAccessIdentityProviderGetResponseAccessCentrify],
// [ZeroTrustAccessIdentityProviderGetResponseAccessFacebook],
// [ZeroTrustAccessIdentityProviderGetResponseAccessGitHub],
// [ZeroTrustAccessIdentityProviderGetResponseAccessGoogle],
// [ZeroTrustAccessIdentityProviderGetResponseAccessGoogleApps],
// [ZeroTrustAccessIdentityProviderGetResponseAccessLinkedin],
// [ZeroTrustAccessIdentityProviderGetResponseAccessOidc],
// [ZeroTrustAccessIdentityProviderGetResponseAccessOkta],
// [ZeroTrustAccessIdentityProviderGetResponseAccessOnelogin],
// [ZeroTrustAccessIdentityProviderGetResponseAccessPingone],
// [ZeroTrustAccessIdentityProviderGetResponseAccessSaml],
// [ZeroTrustAccessIdentityProviderGetResponseAccessYandex] or
// [ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepin].
type ZeroTrustAccessIdentityProviderGetResponse interface {
	implementsZeroTrustAccessIdentityProviderGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustAccessIdentityProviderGetResponse)(nil)).Elem(), "")
}

type ZeroTrustAccessIdentityProviderGetResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessAzureAdJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessAzureAdJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessAzureAd]
type zeroTrustAccessIdentityProviderGetResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderGetResponseAccessAzureAd) implementsZeroTrustAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdConfig struct {
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
	JSON          zeroTrustAccessIdentityProviderGetResponseAccessAzureAdConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessAzureAdConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessAzureAdConfigJSON struct {
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

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType string

const (
	ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdTypeOnetimepin ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType = "onetimepin"
	ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdTypeAzureAd    ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType = "azureAD"
	ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdTypeSaml       ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType = "saml"
	ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdTypeCentrify   ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType = "centrify"
	ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdTypeFacebook   ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType = "facebook"
	ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdTypeGitHub     ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType = "github"
	ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdTypeGoogleApps ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType = "google-apps"
	ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdTypeGoogle     ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType = "google"
	ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdTypeLinkedin   ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType = "linkedin"
	ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdTypeOidc       ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType = "oidc"
	ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdTypeOkta       ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType = "okta"
	ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdTypeOnelogin   ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType = "onelogin"
	ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdTypePingone    ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType = "pingone"
	ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdTypeYandex     ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessAzureAdScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdScimConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessCentrifyJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessCentrifyJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessCentrify]
type zeroTrustAccessIdentityProviderGetResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderGetResponseAccessCentrify) implementsZeroTrustAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyConfig struct {
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
	JSON           zeroTrustAccessIdentityProviderGetResponseAccessCentrifyConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessCentrifyConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType string

const (
	ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyTypeOnetimepin ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType = "onetimepin"
	ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyTypeAzureAd    ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType = "azureAD"
	ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyTypeSaml       ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType = "saml"
	ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyTypeCentrify   ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType = "centrify"
	ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyTypeFacebook   ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType = "facebook"
	ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyTypeGitHub     ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType = "github"
	ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyTypeGoogleApps ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType = "google-apps"
	ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyTypeGoogle     ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType = "google"
	ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyTypeLinkedin   ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType = "linkedin"
	ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyTypeOidc       ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType = "oidc"
	ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyTypeOkta       ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType = "okta"
	ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyTypeOnelogin   ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType = "onelogin"
	ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyTypePingone    ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType = "pingone"
	ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyTypeYandex     ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessCentrifyScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyScimConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderGetResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderGetResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessFacebookJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessFacebookJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessFacebook]
type zeroTrustAccessIdentityProviderGetResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderGetResponseAccessFacebook) implementsZeroTrustAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                             `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderGetResponseAccessFacebookConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessFacebookConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessFacebookConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType string

const (
	ZeroTrustAccessIdentityProviderGetResponseAccessFacebookTypeOnetimepin ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType = "onetimepin"
	ZeroTrustAccessIdentityProviderGetResponseAccessFacebookTypeAzureAd    ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType = "azureAD"
	ZeroTrustAccessIdentityProviderGetResponseAccessFacebookTypeSaml       ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType = "saml"
	ZeroTrustAccessIdentityProviderGetResponseAccessFacebookTypeCentrify   ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType = "centrify"
	ZeroTrustAccessIdentityProviderGetResponseAccessFacebookTypeFacebook   ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType = "facebook"
	ZeroTrustAccessIdentityProviderGetResponseAccessFacebookTypeGitHub     ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType = "github"
	ZeroTrustAccessIdentityProviderGetResponseAccessFacebookTypeGoogleApps ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType = "google-apps"
	ZeroTrustAccessIdentityProviderGetResponseAccessFacebookTypeGoogle     ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType = "google"
	ZeroTrustAccessIdentityProviderGetResponseAccessFacebookTypeLinkedin   ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType = "linkedin"
	ZeroTrustAccessIdentityProviderGetResponseAccessFacebookTypeOidc       ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType = "oidc"
	ZeroTrustAccessIdentityProviderGetResponseAccessFacebookTypeOkta       ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType = "okta"
	ZeroTrustAccessIdentityProviderGetResponseAccessFacebookTypeOnelogin   ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType = "onelogin"
	ZeroTrustAccessIdentityProviderGetResponseAccessFacebookTypePingone    ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType = "pingone"
	ZeroTrustAccessIdentityProviderGetResponseAccessFacebookTypeYandex     ZeroTrustAccessIdentityProviderGetResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderGetResponseAccessFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessFacebookScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessFacebookScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessFacebookScimConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderGetResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderGetResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessGitHubJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessGitHubJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderGetResponseAccessGitHub]
type zeroTrustAccessIdentityProviderGetResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderGetResponseAccessGitHub) implementsZeroTrustAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                           `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderGetResponseAccessGitHubConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessGitHubConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessGitHubConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType string

const (
	ZeroTrustAccessIdentityProviderGetResponseAccessGitHubTypeOnetimepin ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType = "onetimepin"
	ZeroTrustAccessIdentityProviderGetResponseAccessGitHubTypeAzureAd    ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType = "azureAD"
	ZeroTrustAccessIdentityProviderGetResponseAccessGitHubTypeSaml       ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType = "saml"
	ZeroTrustAccessIdentityProviderGetResponseAccessGitHubTypeCentrify   ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType = "centrify"
	ZeroTrustAccessIdentityProviderGetResponseAccessGitHubTypeFacebook   ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType = "facebook"
	ZeroTrustAccessIdentityProviderGetResponseAccessGitHubTypeGitHub     ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType = "github"
	ZeroTrustAccessIdentityProviderGetResponseAccessGitHubTypeGoogleApps ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType = "google-apps"
	ZeroTrustAccessIdentityProviderGetResponseAccessGitHubTypeGoogle     ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType = "google"
	ZeroTrustAccessIdentityProviderGetResponseAccessGitHubTypeLinkedin   ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType = "linkedin"
	ZeroTrustAccessIdentityProviderGetResponseAccessGitHubTypeOidc       ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType = "oidc"
	ZeroTrustAccessIdentityProviderGetResponseAccessGitHubTypeOkta       ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType = "okta"
	ZeroTrustAccessIdentityProviderGetResponseAccessGitHubTypeOnelogin   ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType = "onelogin"
	ZeroTrustAccessIdentityProviderGetResponseAccessGitHubTypePingone    ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType = "pingone"
	ZeroTrustAccessIdentityProviderGetResponseAccessGitHubTypeYandex     ZeroTrustAccessIdentityProviderGetResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderGetResponseAccessGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessGitHubScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessGitHubScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessGitHubScimConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderGetResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderGetResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessGoogleJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessGoogleJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderGetResponseAccessGoogle]
type zeroTrustAccessIdentityProviderGetResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderGetResponseAccessGoogle) implementsZeroTrustAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                           `json:"email_claim_name"`
	JSON           zeroTrustAccessIdentityProviderGetResponseAccessGoogleConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessGoogleConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessGoogleConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType string

const (
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleTypeOnetimepin ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType = "onetimepin"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleTypeAzureAd    ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType = "azureAD"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleTypeSaml       ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType = "saml"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleTypeCentrify   ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType = "centrify"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleTypeFacebook   ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType = "facebook"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleTypeGitHub     ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType = "github"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleTypeGoogleApps ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType = "google-apps"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleTypeGoogle     ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType = "google"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleTypeLinkedin   ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType = "linkedin"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleTypeOidc       ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType = "oidc"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleTypeOkta       ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType = "okta"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleTypeOnelogin   ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType = "onelogin"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleTypePingone    ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType = "pingone"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleTypeYandex     ZeroTrustAccessIdentityProviderGetResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderGetResponseAccessGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessGoogleScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessGoogleScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessGoogleScimConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessGoogleApps]
type zeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderGetResponseAccessGoogleApps) implementsZeroTrustAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsConfig struct {
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
	JSON           zeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType string

const (
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsTypeOnetimepin ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType = "onetimepin"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsTypeAzureAd    ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType = "azureAD"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsTypeSaml       ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType = "saml"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsTypeCentrify   ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType = "centrify"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsTypeFacebook   ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType = "facebook"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsTypeGitHub     ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType = "github"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsTypeGoogleApps ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType = "google-apps"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsTypeGoogle     ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType = "google"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsTypeLinkedin   ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType = "linkedin"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsTypeOidc       ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType = "oidc"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsTypeOkta       ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType = "okta"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsTypeOnelogin   ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType = "onelogin"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsTypePingone    ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType = "pingone"
	ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsTypeYandex     ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsScimConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessLinkedinJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessLinkedinJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessLinkedin]
type zeroTrustAccessIdentityProviderGetResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderGetResponseAccessLinkedin) implementsZeroTrustAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                             `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderGetResponseAccessLinkedinConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessLinkedinConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType string

const (
	ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinTypeOnetimepin ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType = "onetimepin"
	ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinTypeAzureAd    ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType = "azureAD"
	ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinTypeSaml       ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType = "saml"
	ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinTypeCentrify   ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType = "centrify"
	ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinTypeFacebook   ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType = "facebook"
	ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinTypeGitHub     ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType = "github"
	ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinTypeGoogleApps ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType = "google-apps"
	ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinTypeGoogle     ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType = "google"
	ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinTypeLinkedin   ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType = "linkedin"
	ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinTypeOidc       ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType = "oidc"
	ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinTypeOkta       ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType = "okta"
	ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinTypeOnelogin   ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType = "onelogin"
	ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinTypePingone    ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType = "pingone"
	ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinTypeYandex     ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessLinkedinScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinScimConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderGetResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderGetResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderGetResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessOidcJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessOidcJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderGetResponseAccessOidc]
type zeroTrustAccessIdentityProviderGetResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderGetResponseAccessOidc) implementsZeroTrustAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessOidcConfig struct {
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
	JSON     zeroTrustAccessIdentityProviderGetResponseAccessOidcConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessOidcConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessOidcConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessOidcConfigJSON struct {
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

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessOidcType string

const (
	ZeroTrustAccessIdentityProviderGetResponseAccessOidcTypeOnetimepin ZeroTrustAccessIdentityProviderGetResponseAccessOidcType = "onetimepin"
	ZeroTrustAccessIdentityProviderGetResponseAccessOidcTypeAzureAd    ZeroTrustAccessIdentityProviderGetResponseAccessOidcType = "azureAD"
	ZeroTrustAccessIdentityProviderGetResponseAccessOidcTypeSaml       ZeroTrustAccessIdentityProviderGetResponseAccessOidcType = "saml"
	ZeroTrustAccessIdentityProviderGetResponseAccessOidcTypeCentrify   ZeroTrustAccessIdentityProviderGetResponseAccessOidcType = "centrify"
	ZeroTrustAccessIdentityProviderGetResponseAccessOidcTypeFacebook   ZeroTrustAccessIdentityProviderGetResponseAccessOidcType = "facebook"
	ZeroTrustAccessIdentityProviderGetResponseAccessOidcTypeGitHub     ZeroTrustAccessIdentityProviderGetResponseAccessOidcType = "github"
	ZeroTrustAccessIdentityProviderGetResponseAccessOidcTypeGoogleApps ZeroTrustAccessIdentityProviderGetResponseAccessOidcType = "google-apps"
	ZeroTrustAccessIdentityProviderGetResponseAccessOidcTypeGoogle     ZeroTrustAccessIdentityProviderGetResponseAccessOidcType = "google"
	ZeroTrustAccessIdentityProviderGetResponseAccessOidcTypeLinkedin   ZeroTrustAccessIdentityProviderGetResponseAccessOidcType = "linkedin"
	ZeroTrustAccessIdentityProviderGetResponseAccessOidcTypeOidc       ZeroTrustAccessIdentityProviderGetResponseAccessOidcType = "oidc"
	ZeroTrustAccessIdentityProviderGetResponseAccessOidcTypeOkta       ZeroTrustAccessIdentityProviderGetResponseAccessOidcType = "okta"
	ZeroTrustAccessIdentityProviderGetResponseAccessOidcTypeOnelogin   ZeroTrustAccessIdentityProviderGetResponseAccessOidcType = "onelogin"
	ZeroTrustAccessIdentityProviderGetResponseAccessOidcTypePingone    ZeroTrustAccessIdentityProviderGetResponseAccessOidcType = "pingone"
	ZeroTrustAccessIdentityProviderGetResponseAccessOidcTypeYandex     ZeroTrustAccessIdentityProviderGetResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderGetResponseAccessOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessOidcScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessOidcScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessOidcScimConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderGetResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderGetResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderGetResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessOktaJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessOktaJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderGetResponseAccessOkta]
type zeroTrustAccessIdentityProviderGetResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderGetResponseAccessOkta) implementsZeroTrustAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessOktaConfig struct {
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
	JSON        zeroTrustAccessIdentityProviderGetResponseAccessOktaConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessOktaConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessOktaConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessOktaType string

const (
	ZeroTrustAccessIdentityProviderGetResponseAccessOktaTypeOnetimepin ZeroTrustAccessIdentityProviderGetResponseAccessOktaType = "onetimepin"
	ZeroTrustAccessIdentityProviderGetResponseAccessOktaTypeAzureAd    ZeroTrustAccessIdentityProviderGetResponseAccessOktaType = "azureAD"
	ZeroTrustAccessIdentityProviderGetResponseAccessOktaTypeSaml       ZeroTrustAccessIdentityProviderGetResponseAccessOktaType = "saml"
	ZeroTrustAccessIdentityProviderGetResponseAccessOktaTypeCentrify   ZeroTrustAccessIdentityProviderGetResponseAccessOktaType = "centrify"
	ZeroTrustAccessIdentityProviderGetResponseAccessOktaTypeFacebook   ZeroTrustAccessIdentityProviderGetResponseAccessOktaType = "facebook"
	ZeroTrustAccessIdentityProviderGetResponseAccessOktaTypeGitHub     ZeroTrustAccessIdentityProviderGetResponseAccessOktaType = "github"
	ZeroTrustAccessIdentityProviderGetResponseAccessOktaTypeGoogleApps ZeroTrustAccessIdentityProviderGetResponseAccessOktaType = "google-apps"
	ZeroTrustAccessIdentityProviderGetResponseAccessOktaTypeGoogle     ZeroTrustAccessIdentityProviderGetResponseAccessOktaType = "google"
	ZeroTrustAccessIdentityProviderGetResponseAccessOktaTypeLinkedin   ZeroTrustAccessIdentityProviderGetResponseAccessOktaType = "linkedin"
	ZeroTrustAccessIdentityProviderGetResponseAccessOktaTypeOidc       ZeroTrustAccessIdentityProviderGetResponseAccessOktaType = "oidc"
	ZeroTrustAccessIdentityProviderGetResponseAccessOktaTypeOkta       ZeroTrustAccessIdentityProviderGetResponseAccessOktaType = "okta"
	ZeroTrustAccessIdentityProviderGetResponseAccessOktaTypeOnelogin   ZeroTrustAccessIdentityProviderGetResponseAccessOktaType = "onelogin"
	ZeroTrustAccessIdentityProviderGetResponseAccessOktaTypePingone    ZeroTrustAccessIdentityProviderGetResponseAccessOktaType = "pingone"
	ZeroTrustAccessIdentityProviderGetResponseAccessOktaTypeYandex     ZeroTrustAccessIdentityProviderGetResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderGetResponseAccessOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessOktaScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessOktaScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessOktaScimConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderGetResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderGetResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessOneloginJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessOneloginJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessOnelogin]
type zeroTrustAccessIdentityProviderGetResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderGetResponseAccessOnelogin) implementsZeroTrustAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessOneloginConfig struct {
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessOneloginConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessOneloginConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessOneloginConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType string

const (
	ZeroTrustAccessIdentityProviderGetResponseAccessOneloginTypeOnetimepin ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType = "onetimepin"
	ZeroTrustAccessIdentityProviderGetResponseAccessOneloginTypeAzureAd    ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType = "azureAD"
	ZeroTrustAccessIdentityProviderGetResponseAccessOneloginTypeSaml       ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType = "saml"
	ZeroTrustAccessIdentityProviderGetResponseAccessOneloginTypeCentrify   ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType = "centrify"
	ZeroTrustAccessIdentityProviderGetResponseAccessOneloginTypeFacebook   ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType = "facebook"
	ZeroTrustAccessIdentityProviderGetResponseAccessOneloginTypeGitHub     ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType = "github"
	ZeroTrustAccessIdentityProviderGetResponseAccessOneloginTypeGoogleApps ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType = "google-apps"
	ZeroTrustAccessIdentityProviderGetResponseAccessOneloginTypeGoogle     ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType = "google"
	ZeroTrustAccessIdentityProviderGetResponseAccessOneloginTypeLinkedin   ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType = "linkedin"
	ZeroTrustAccessIdentityProviderGetResponseAccessOneloginTypeOidc       ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType = "oidc"
	ZeroTrustAccessIdentityProviderGetResponseAccessOneloginTypeOkta       ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType = "okta"
	ZeroTrustAccessIdentityProviderGetResponseAccessOneloginTypeOnelogin   ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType = "onelogin"
	ZeroTrustAccessIdentityProviderGetResponseAccessOneloginTypePingone    ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType = "pingone"
	ZeroTrustAccessIdentityProviderGetResponseAccessOneloginTypeYandex     ZeroTrustAccessIdentityProviderGetResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderGetResponseAccessOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessOneloginScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessOneloginScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessOneloginScimConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderGetResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderGetResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessPingoneJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessPingoneJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessPingone]
type zeroTrustAccessIdentityProviderGetResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderGetResponseAccessPingone) implementsZeroTrustAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessPingoneConfig struct {
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
	JSON      zeroTrustAccessIdentityProviderGetResponseAccessPingoneConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessPingoneConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessPingoneConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType string

const (
	ZeroTrustAccessIdentityProviderGetResponseAccessPingoneTypeOnetimepin ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType = "onetimepin"
	ZeroTrustAccessIdentityProviderGetResponseAccessPingoneTypeAzureAd    ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType = "azureAD"
	ZeroTrustAccessIdentityProviderGetResponseAccessPingoneTypeSaml       ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType = "saml"
	ZeroTrustAccessIdentityProviderGetResponseAccessPingoneTypeCentrify   ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType = "centrify"
	ZeroTrustAccessIdentityProviderGetResponseAccessPingoneTypeFacebook   ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType = "facebook"
	ZeroTrustAccessIdentityProviderGetResponseAccessPingoneTypeGitHub     ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType = "github"
	ZeroTrustAccessIdentityProviderGetResponseAccessPingoneTypeGoogleApps ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType = "google-apps"
	ZeroTrustAccessIdentityProviderGetResponseAccessPingoneTypeGoogle     ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType = "google"
	ZeroTrustAccessIdentityProviderGetResponseAccessPingoneTypeLinkedin   ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType = "linkedin"
	ZeroTrustAccessIdentityProviderGetResponseAccessPingoneTypeOidc       ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType = "oidc"
	ZeroTrustAccessIdentityProviderGetResponseAccessPingoneTypeOkta       ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType = "okta"
	ZeroTrustAccessIdentityProviderGetResponseAccessPingoneTypeOnelogin   ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType = "onelogin"
	ZeroTrustAccessIdentityProviderGetResponseAccessPingoneTypePingone    ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType = "pingone"
	ZeroTrustAccessIdentityProviderGetResponseAccessPingoneTypeYandex     ZeroTrustAccessIdentityProviderGetResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderGetResponseAccessPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessPingoneScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessPingoneScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessPingoneScimConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderGetResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderGetResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderGetResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessSamlJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessSamlJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderGetResponseAccessSaml]
type zeroTrustAccessIdentityProviderGetResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderGetResponseAccessSaml) implementsZeroTrustAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []ZeroTrustAccessIdentityProviderGetResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                         `json:"sso_target_url"`
	JSON         zeroTrustAccessIdentityProviderGetResponseAccessSamlConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessSamlConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessSamlConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessSamlConfigJSON struct {
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

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                        `json:"header_name"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessSamlConfigHeaderAttribute]
type zeroTrustAccessIdentityProviderGetResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessSamlType string

const (
	ZeroTrustAccessIdentityProviderGetResponseAccessSamlTypeOnetimepin ZeroTrustAccessIdentityProviderGetResponseAccessSamlType = "onetimepin"
	ZeroTrustAccessIdentityProviderGetResponseAccessSamlTypeAzureAd    ZeroTrustAccessIdentityProviderGetResponseAccessSamlType = "azureAD"
	ZeroTrustAccessIdentityProviderGetResponseAccessSamlTypeSaml       ZeroTrustAccessIdentityProviderGetResponseAccessSamlType = "saml"
	ZeroTrustAccessIdentityProviderGetResponseAccessSamlTypeCentrify   ZeroTrustAccessIdentityProviderGetResponseAccessSamlType = "centrify"
	ZeroTrustAccessIdentityProviderGetResponseAccessSamlTypeFacebook   ZeroTrustAccessIdentityProviderGetResponseAccessSamlType = "facebook"
	ZeroTrustAccessIdentityProviderGetResponseAccessSamlTypeGitHub     ZeroTrustAccessIdentityProviderGetResponseAccessSamlType = "github"
	ZeroTrustAccessIdentityProviderGetResponseAccessSamlTypeGoogleApps ZeroTrustAccessIdentityProviderGetResponseAccessSamlType = "google-apps"
	ZeroTrustAccessIdentityProviderGetResponseAccessSamlTypeGoogle     ZeroTrustAccessIdentityProviderGetResponseAccessSamlType = "google"
	ZeroTrustAccessIdentityProviderGetResponseAccessSamlTypeLinkedin   ZeroTrustAccessIdentityProviderGetResponseAccessSamlType = "linkedin"
	ZeroTrustAccessIdentityProviderGetResponseAccessSamlTypeOidc       ZeroTrustAccessIdentityProviderGetResponseAccessSamlType = "oidc"
	ZeroTrustAccessIdentityProviderGetResponseAccessSamlTypeOkta       ZeroTrustAccessIdentityProviderGetResponseAccessSamlType = "okta"
	ZeroTrustAccessIdentityProviderGetResponseAccessSamlTypeOnelogin   ZeroTrustAccessIdentityProviderGetResponseAccessSamlType = "onelogin"
	ZeroTrustAccessIdentityProviderGetResponseAccessSamlTypePingone    ZeroTrustAccessIdentityProviderGetResponseAccessSamlType = "pingone"
	ZeroTrustAccessIdentityProviderGetResponseAccessSamlTypeYandex     ZeroTrustAccessIdentityProviderGetResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderGetResponseAccessSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessSamlScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessSamlScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessSamlScimConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustAccessIdentityProviderGetResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderGetResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderGetResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessYandexJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessYandexJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderGetResponseAccessYandex]
type zeroTrustAccessIdentityProviderGetResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderGetResponseAccessYandex) implementsZeroTrustAccessIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                           `json:"client_secret"`
	JSON         zeroTrustAccessIdentityProviderGetResponseAccessYandexConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessYandexConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessYandexConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessYandexType string

const (
	ZeroTrustAccessIdentityProviderGetResponseAccessYandexTypeOnetimepin ZeroTrustAccessIdentityProviderGetResponseAccessYandexType = "onetimepin"
	ZeroTrustAccessIdentityProviderGetResponseAccessYandexTypeAzureAd    ZeroTrustAccessIdentityProviderGetResponseAccessYandexType = "azureAD"
	ZeroTrustAccessIdentityProviderGetResponseAccessYandexTypeSaml       ZeroTrustAccessIdentityProviderGetResponseAccessYandexType = "saml"
	ZeroTrustAccessIdentityProviderGetResponseAccessYandexTypeCentrify   ZeroTrustAccessIdentityProviderGetResponseAccessYandexType = "centrify"
	ZeroTrustAccessIdentityProviderGetResponseAccessYandexTypeFacebook   ZeroTrustAccessIdentityProviderGetResponseAccessYandexType = "facebook"
	ZeroTrustAccessIdentityProviderGetResponseAccessYandexTypeGitHub     ZeroTrustAccessIdentityProviderGetResponseAccessYandexType = "github"
	ZeroTrustAccessIdentityProviderGetResponseAccessYandexTypeGoogleApps ZeroTrustAccessIdentityProviderGetResponseAccessYandexType = "google-apps"
	ZeroTrustAccessIdentityProviderGetResponseAccessYandexTypeGoogle     ZeroTrustAccessIdentityProviderGetResponseAccessYandexType = "google"
	ZeroTrustAccessIdentityProviderGetResponseAccessYandexTypeLinkedin   ZeroTrustAccessIdentityProviderGetResponseAccessYandexType = "linkedin"
	ZeroTrustAccessIdentityProviderGetResponseAccessYandexTypeOidc       ZeroTrustAccessIdentityProviderGetResponseAccessYandexType = "oidc"
	ZeroTrustAccessIdentityProviderGetResponseAccessYandexTypeOkta       ZeroTrustAccessIdentityProviderGetResponseAccessYandexType = "okta"
	ZeroTrustAccessIdentityProviderGetResponseAccessYandexTypeOnelogin   ZeroTrustAccessIdentityProviderGetResponseAccessYandexType = "onelogin"
	ZeroTrustAccessIdentityProviderGetResponseAccessYandexTypePingone    ZeroTrustAccessIdentityProviderGetResponseAccessYandexType = "pingone"
	ZeroTrustAccessIdentityProviderGetResponseAccessYandexTypeYandex     ZeroTrustAccessIdentityProviderGetResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderGetResponseAccessYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessYandexScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessYandexScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessYandexScimConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       zeroTrustAccessIdentityProviderGetResponseAccessOnetimepinJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessOnetimepinJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepin]
type zeroTrustAccessIdentityProviderGetResponseAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepin) implementsZeroTrustAccessIdentityProviderGetResponse() {
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType string

const (
	ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinTypeOnetimepin ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType = "onetimepin"
	ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinTypeAzureAd    ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType = "azureAD"
	ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinTypeSaml       ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType = "saml"
	ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinTypeCentrify   ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType = "centrify"
	ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinTypeFacebook   ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType = "facebook"
	ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinTypeGitHub     ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType = "github"
	ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinTypeGoogleApps ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType = "google-apps"
	ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinTypeGoogle     ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType = "google"
	ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinTypeLinkedin   ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType = "linkedin"
	ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinTypeOidc       ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType = "oidc"
	ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinTypeOkta       ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType = "okta"
	ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinTypeOnelogin   ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType = "onelogin"
	ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinTypePingone    ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType = "pingone"
	ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinTypeYandex     ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustAccessIdentityProviderGetResponseAccessOnetimepinScimConfigJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseAccessOnetimepinScimConfigJSON
// contains the JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinScimConfig]
type zeroTrustAccessIdentityProviderGetResponseAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewParams struct {
	Config param.Field[ZeroTrustAccessIdentityProviderNewParamsConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZeroTrustAccessIdentityProviderNewParamsType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID     param.Field[string]                                             `path:"zone_id"`
	ScimConfig param.Field[ZeroTrustAccessIdentityProviderNewParamsScimConfig] `json:"scim_config"`
}

func (r ZeroTrustAccessIdentityProviderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessIdentityProviderNewParamsConfig struct {
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
	HeaderAttributes param.Field[[]ZeroTrustAccessIdentityProviderNewParamsConfigHeaderAttribute] `json:"header_attributes"`
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

func (r ZeroTrustAccessIdentityProviderNewParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessIdentityProviderNewParamsConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r ZeroTrustAccessIdentityProviderNewParamsConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderNewParamsType string

const (
	ZeroTrustAccessIdentityProviderNewParamsTypeOnetimepin ZeroTrustAccessIdentityProviderNewParamsType = "onetimepin"
	ZeroTrustAccessIdentityProviderNewParamsTypeAzureAd    ZeroTrustAccessIdentityProviderNewParamsType = "azureAD"
	ZeroTrustAccessIdentityProviderNewParamsTypeSaml       ZeroTrustAccessIdentityProviderNewParamsType = "saml"
	ZeroTrustAccessIdentityProviderNewParamsTypeCentrify   ZeroTrustAccessIdentityProviderNewParamsType = "centrify"
	ZeroTrustAccessIdentityProviderNewParamsTypeFacebook   ZeroTrustAccessIdentityProviderNewParamsType = "facebook"
	ZeroTrustAccessIdentityProviderNewParamsTypeGitHub     ZeroTrustAccessIdentityProviderNewParamsType = "github"
	ZeroTrustAccessIdentityProviderNewParamsTypeGoogleApps ZeroTrustAccessIdentityProviderNewParamsType = "google-apps"
	ZeroTrustAccessIdentityProviderNewParamsTypeGoogle     ZeroTrustAccessIdentityProviderNewParamsType = "google"
	ZeroTrustAccessIdentityProviderNewParamsTypeLinkedin   ZeroTrustAccessIdentityProviderNewParamsType = "linkedin"
	ZeroTrustAccessIdentityProviderNewParamsTypeOidc       ZeroTrustAccessIdentityProviderNewParamsType = "oidc"
	ZeroTrustAccessIdentityProviderNewParamsTypeOkta       ZeroTrustAccessIdentityProviderNewParamsType = "okta"
	ZeroTrustAccessIdentityProviderNewParamsTypeOnelogin   ZeroTrustAccessIdentityProviderNewParamsType = "onelogin"
	ZeroTrustAccessIdentityProviderNewParamsTypePingone    ZeroTrustAccessIdentityProviderNewParamsType = "pingone"
	ZeroTrustAccessIdentityProviderNewParamsTypeYandex     ZeroTrustAccessIdentityProviderNewParamsType = "yandex"
)

type ZeroTrustAccessIdentityProviderNewParamsScimConfig struct {
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

func (r ZeroTrustAccessIdentityProviderNewParamsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessIdentityProviderNewResponseEnvelope struct {
	Errors   []ZeroTrustAccessIdentityProviderNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessIdentityProviderNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessIdentityProviderNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessIdentityProviderNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessIdentityProviderNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderNewResponseEnvelope]
type zeroTrustAccessIdentityProviderNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustAccessIdentityProviderNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseEnvelopeErrors]
type zeroTrustAccessIdentityProviderNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderNewResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustAccessIdentityProviderNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderNewResponseEnvelopeMessages]
type zeroTrustAccessIdentityProviderNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessIdentityProviderNewResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessIdentityProviderNewResponseEnvelopeSuccessTrue ZeroTrustAccessIdentityProviderNewResponseEnvelopeSuccess = true
)

type ZeroTrustAccessIdentityProviderUpdateParams struct {
	Config param.Field[ZeroTrustAccessIdentityProviderUpdateParamsConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZeroTrustAccessIdentityProviderUpdateParamsType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID     param.Field[string]                                                `path:"zone_id"`
	ScimConfig param.Field[ZeroTrustAccessIdentityProviderUpdateParamsScimConfig] `json:"scim_config"`
}

func (r ZeroTrustAccessIdentityProviderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessIdentityProviderUpdateParamsConfig struct {
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
	HeaderAttributes param.Field[[]ZeroTrustAccessIdentityProviderUpdateParamsConfigHeaderAttribute] `json:"header_attributes"`
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

func (r ZeroTrustAccessIdentityProviderUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessIdentityProviderUpdateParamsConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r ZeroTrustAccessIdentityProviderUpdateParamsConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustAccessIdentityProviderUpdateParamsType string

const (
	ZeroTrustAccessIdentityProviderUpdateParamsTypeOnetimepin ZeroTrustAccessIdentityProviderUpdateParamsType = "onetimepin"
	ZeroTrustAccessIdentityProviderUpdateParamsTypeAzureAd    ZeroTrustAccessIdentityProviderUpdateParamsType = "azureAD"
	ZeroTrustAccessIdentityProviderUpdateParamsTypeSaml       ZeroTrustAccessIdentityProviderUpdateParamsType = "saml"
	ZeroTrustAccessIdentityProviderUpdateParamsTypeCentrify   ZeroTrustAccessIdentityProviderUpdateParamsType = "centrify"
	ZeroTrustAccessIdentityProviderUpdateParamsTypeFacebook   ZeroTrustAccessIdentityProviderUpdateParamsType = "facebook"
	ZeroTrustAccessIdentityProviderUpdateParamsTypeGitHub     ZeroTrustAccessIdentityProviderUpdateParamsType = "github"
	ZeroTrustAccessIdentityProviderUpdateParamsTypeGoogleApps ZeroTrustAccessIdentityProviderUpdateParamsType = "google-apps"
	ZeroTrustAccessIdentityProviderUpdateParamsTypeGoogle     ZeroTrustAccessIdentityProviderUpdateParamsType = "google"
	ZeroTrustAccessIdentityProviderUpdateParamsTypeLinkedin   ZeroTrustAccessIdentityProviderUpdateParamsType = "linkedin"
	ZeroTrustAccessIdentityProviderUpdateParamsTypeOidc       ZeroTrustAccessIdentityProviderUpdateParamsType = "oidc"
	ZeroTrustAccessIdentityProviderUpdateParamsTypeOkta       ZeroTrustAccessIdentityProviderUpdateParamsType = "okta"
	ZeroTrustAccessIdentityProviderUpdateParamsTypeOnelogin   ZeroTrustAccessIdentityProviderUpdateParamsType = "onelogin"
	ZeroTrustAccessIdentityProviderUpdateParamsTypePingone    ZeroTrustAccessIdentityProviderUpdateParamsType = "pingone"
	ZeroTrustAccessIdentityProviderUpdateParamsTypeYandex     ZeroTrustAccessIdentityProviderUpdateParamsType = "yandex"
)

type ZeroTrustAccessIdentityProviderUpdateParamsScimConfig struct {
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

func (r ZeroTrustAccessIdentityProviderUpdateParamsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseEnvelope struct {
	Errors   []ZeroTrustAccessIdentityProviderUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessIdentityProviderUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessIdentityProviderUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessIdentityProviderUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessIdentityProviderUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderUpdateResponseEnvelope]
type zeroTrustAccessIdentityProviderUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustAccessIdentityProviderUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseEnvelopeErrors]
type zeroTrustAccessIdentityProviderUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderUpdateResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zeroTrustAccessIdentityProviderUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderUpdateResponseEnvelopeMessages]
type zeroTrustAccessIdentityProviderUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessIdentityProviderUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessIdentityProviderUpdateResponseEnvelopeSuccessTrue ZeroTrustAccessIdentityProviderUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustAccessIdentityProviderListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessIdentityProviderListResponseEnvelope struct {
	Errors   []ZeroTrustAccessIdentityProviderListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessIdentityProviderListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustAccessIdentityProviderListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessIdentityProviderListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessIdentityProviderListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessIdentityProviderListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderListResponseEnvelope]
type zeroTrustAccessIdentityProviderListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustAccessIdentityProviderListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseEnvelopeErrors]
type zeroTrustAccessIdentityProviderListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderListResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustAccessIdentityProviderListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseEnvelopeMessages]
type zeroTrustAccessIdentityProviderListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessIdentityProviderListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessIdentityProviderListResponseEnvelopeSuccessTrue ZeroTrustAccessIdentityProviderListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessIdentityProviderListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                           `json:"total_count"`
	JSON       zeroTrustAccessIdentityProviderListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderListResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderListResponseEnvelopeResultInfo]
type zeroTrustAccessIdentityProviderListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessIdentityProviderDeleteResponseEnvelope struct {
	Errors   []ZeroTrustAccessIdentityProviderDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessIdentityProviderDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessIdentityProviderDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessIdentityProviderDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessIdentityProviderDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessIdentityProviderDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderDeleteResponseEnvelope]
type zeroTrustAccessIdentityProviderDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderDeleteResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustAccessIdentityProviderDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderDeleteResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderDeleteResponseEnvelopeErrors]
type zeroTrustAccessIdentityProviderDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderDeleteResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zeroTrustAccessIdentityProviderDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderDeleteResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustAccessIdentityProviderDeleteResponseEnvelopeMessages]
type zeroTrustAccessIdentityProviderDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessIdentityProviderDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessIdentityProviderDeleteResponseEnvelopeSuccessTrue ZeroTrustAccessIdentityProviderDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustAccessIdentityProviderGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessIdentityProviderGetResponseEnvelope struct {
	Errors   []ZeroTrustAccessIdentityProviderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessIdentityProviderGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessIdentityProviderGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessIdentityProviderGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessIdentityProviderGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessIdentityProviderGetResponseEnvelope]
type zeroTrustAccessIdentityProviderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustAccessIdentityProviderGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseEnvelopeErrors]
type zeroTrustAccessIdentityProviderGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessIdentityProviderGetResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustAccessIdentityProviderGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessIdentityProviderGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessIdentityProviderGetResponseEnvelopeMessages]
type zeroTrustAccessIdentityProviderGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessIdentityProviderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessIdentityProviderGetResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessIdentityProviderGetResponseEnvelopeSuccessTrue ZeroTrustAccessIdentityProviderGetResponseEnvelopeSuccess = true
)

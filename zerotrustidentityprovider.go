// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
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
func (r *ZeroTrustIdentityProviderService) New(ctx context.Context, params ZeroTrustIdentityProviderNewParams, opts ...option.RequestOption) (res *ZeroTrustIdentityProviderNewResponse, err error) {
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
func (r *ZeroTrustIdentityProviderService) Update(ctx context.Context, uuid string, params ZeroTrustIdentityProviderUpdateParams, opts ...option.RequestOption) (res *ZeroTrustIdentityProviderUpdateResponse, err error) {
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
func (r *ZeroTrustIdentityProviderService) Get(ctx context.Context, uuid string, query ZeroTrustIdentityProviderGetParams, opts ...option.RequestOption) (res *ZeroTrustIdentityProviderGetResponse, err error) {
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

// Union satisfied by [ZeroTrustIdentityProviderNewResponseAccessAzureAd],
// [ZeroTrustIdentityProviderNewResponseAccessCentrify],
// [ZeroTrustIdentityProviderNewResponseAccessFacebook],
// [ZeroTrustIdentityProviderNewResponseAccessGitHub],
// [ZeroTrustIdentityProviderNewResponseAccessGoogle],
// [ZeroTrustIdentityProviderNewResponseAccessGoogleApps],
// [ZeroTrustIdentityProviderNewResponseAccessLinkedin],
// [ZeroTrustIdentityProviderNewResponseAccessOidc],
// [ZeroTrustIdentityProviderNewResponseAccessOkta],
// [ZeroTrustIdentityProviderNewResponseAccessOnelogin],
// [ZeroTrustIdentityProviderNewResponseAccessPingone],
// [ZeroTrustIdentityProviderNewResponseAccessSaml],
// [ZeroTrustIdentityProviderNewResponseAccessYandex] or
// [ZeroTrustIdentityProviderNewResponseAccessOnetimepin].
type ZeroTrustIdentityProviderNewResponse interface {
	implementsZeroTrustIdentityProviderNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustIdentityProviderNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderNewResponseAccessAzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderNewResponseAccessCentrify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderNewResponseAccessFacebook{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderNewResponseAccessGitHub{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderNewResponseAccessGoogle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderNewResponseAccessGoogleApps{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderNewResponseAccessLinkedin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderNewResponseAccessOidc{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderNewResponseAccessOkta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderNewResponseAccessOnelogin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderNewResponseAccessPingone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderNewResponseAccessSaml{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderNewResponseAccessYandex{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderNewResponseAccessOnetimepin{}),
		},
	)
}

type ZeroTrustIdentityProviderNewResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderNewResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderNewResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderNewResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderNewResponseAccessAzureAdJSON       `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessAzureAdJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderNewResponseAccessAzureAd]
type zeroTrustIdentityProviderNewResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessAzureAdJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderNewResponseAccessAzureAd) implementsZeroTrustIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessAzureAdConfig struct {
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
	JSON          zeroTrustIdentityProviderNewResponseAccessAzureAdConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessAzureAdConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessAzureAdConfig]
type zeroTrustIdentityProviderNewResponseAccessAzureAdConfigJSON struct {
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

func (r *ZeroTrustIdentityProviderNewResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessAzureAdConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessAzureAdType string

const (
	ZeroTrustIdentityProviderNewResponseAccessAzureAdTypeOnetimepin ZeroTrustIdentityProviderNewResponseAccessAzureAdType = "onetimepin"
	ZeroTrustIdentityProviderNewResponseAccessAzureAdTypeAzureAd    ZeroTrustIdentityProviderNewResponseAccessAzureAdType = "azureAD"
	ZeroTrustIdentityProviderNewResponseAccessAzureAdTypeSaml       ZeroTrustIdentityProviderNewResponseAccessAzureAdType = "saml"
	ZeroTrustIdentityProviderNewResponseAccessAzureAdTypeCentrify   ZeroTrustIdentityProviderNewResponseAccessAzureAdType = "centrify"
	ZeroTrustIdentityProviderNewResponseAccessAzureAdTypeFacebook   ZeroTrustIdentityProviderNewResponseAccessAzureAdType = "facebook"
	ZeroTrustIdentityProviderNewResponseAccessAzureAdTypeGitHub     ZeroTrustIdentityProviderNewResponseAccessAzureAdType = "github"
	ZeroTrustIdentityProviderNewResponseAccessAzureAdTypeGoogleApps ZeroTrustIdentityProviderNewResponseAccessAzureAdType = "google-apps"
	ZeroTrustIdentityProviderNewResponseAccessAzureAdTypeGoogle     ZeroTrustIdentityProviderNewResponseAccessAzureAdType = "google"
	ZeroTrustIdentityProviderNewResponseAccessAzureAdTypeLinkedin   ZeroTrustIdentityProviderNewResponseAccessAzureAdType = "linkedin"
	ZeroTrustIdentityProviderNewResponseAccessAzureAdTypeOidc       ZeroTrustIdentityProviderNewResponseAccessAzureAdType = "oidc"
	ZeroTrustIdentityProviderNewResponseAccessAzureAdTypeOkta       ZeroTrustIdentityProviderNewResponseAccessAzureAdType = "okta"
	ZeroTrustIdentityProviderNewResponseAccessAzureAdTypeOnelogin   ZeroTrustIdentityProviderNewResponseAccessAzureAdType = "onelogin"
	ZeroTrustIdentityProviderNewResponseAccessAzureAdTypePingone    ZeroTrustIdentityProviderNewResponseAccessAzureAdType = "pingone"
	ZeroTrustIdentityProviderNewResponseAccessAzureAdTypeYandex     ZeroTrustIdentityProviderNewResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderNewResponseAccessAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderNewResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessAzureAdScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessAzureAdScimConfig]
type zeroTrustIdentityProviderNewResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessAzureAdScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderNewResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderNewResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderNewResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderNewResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderNewResponseAccessCentrifyJSON       `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessCentrifyJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderNewResponseAccessCentrify]
type zeroTrustIdentityProviderNewResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessCentrifyJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderNewResponseAccessCentrify) implementsZeroTrustIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessCentrifyConfig struct {
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
	JSON           zeroTrustIdentityProviderNewResponseAccessCentrifyConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessCentrifyConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessCentrifyConfig]
type zeroTrustIdentityProviderNewResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessCentrifyConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessCentrifyType string

const (
	ZeroTrustIdentityProviderNewResponseAccessCentrifyTypeOnetimepin ZeroTrustIdentityProviderNewResponseAccessCentrifyType = "onetimepin"
	ZeroTrustIdentityProviderNewResponseAccessCentrifyTypeAzureAd    ZeroTrustIdentityProviderNewResponseAccessCentrifyType = "azureAD"
	ZeroTrustIdentityProviderNewResponseAccessCentrifyTypeSaml       ZeroTrustIdentityProviderNewResponseAccessCentrifyType = "saml"
	ZeroTrustIdentityProviderNewResponseAccessCentrifyTypeCentrify   ZeroTrustIdentityProviderNewResponseAccessCentrifyType = "centrify"
	ZeroTrustIdentityProviderNewResponseAccessCentrifyTypeFacebook   ZeroTrustIdentityProviderNewResponseAccessCentrifyType = "facebook"
	ZeroTrustIdentityProviderNewResponseAccessCentrifyTypeGitHub     ZeroTrustIdentityProviderNewResponseAccessCentrifyType = "github"
	ZeroTrustIdentityProviderNewResponseAccessCentrifyTypeGoogleApps ZeroTrustIdentityProviderNewResponseAccessCentrifyType = "google-apps"
	ZeroTrustIdentityProviderNewResponseAccessCentrifyTypeGoogle     ZeroTrustIdentityProviderNewResponseAccessCentrifyType = "google"
	ZeroTrustIdentityProviderNewResponseAccessCentrifyTypeLinkedin   ZeroTrustIdentityProviderNewResponseAccessCentrifyType = "linkedin"
	ZeroTrustIdentityProviderNewResponseAccessCentrifyTypeOidc       ZeroTrustIdentityProviderNewResponseAccessCentrifyType = "oidc"
	ZeroTrustIdentityProviderNewResponseAccessCentrifyTypeOkta       ZeroTrustIdentityProviderNewResponseAccessCentrifyType = "okta"
	ZeroTrustIdentityProviderNewResponseAccessCentrifyTypeOnelogin   ZeroTrustIdentityProviderNewResponseAccessCentrifyType = "onelogin"
	ZeroTrustIdentityProviderNewResponseAccessCentrifyTypePingone    ZeroTrustIdentityProviderNewResponseAccessCentrifyType = "pingone"
	ZeroTrustIdentityProviderNewResponseAccessCentrifyTypeYandex     ZeroTrustIdentityProviderNewResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderNewResponseAccessCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderNewResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessCentrifyScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessCentrifyScimConfig]
type zeroTrustIdentityProviderNewResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessCentrifyScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderNewResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderNewResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderNewResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderNewResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderNewResponseAccessFacebookJSON       `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessFacebookJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderNewResponseAccessFacebook]
type zeroTrustIdentityProviderNewResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessFacebookJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderNewResponseAccessFacebook) implementsZeroTrustIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                       `json:"client_secret"`
	JSON         zeroTrustIdentityProviderNewResponseAccessFacebookConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessFacebookConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessFacebookConfig]
type zeroTrustIdentityProviderNewResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessFacebookConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessFacebookType string

const (
	ZeroTrustIdentityProviderNewResponseAccessFacebookTypeOnetimepin ZeroTrustIdentityProviderNewResponseAccessFacebookType = "onetimepin"
	ZeroTrustIdentityProviderNewResponseAccessFacebookTypeAzureAd    ZeroTrustIdentityProviderNewResponseAccessFacebookType = "azureAD"
	ZeroTrustIdentityProviderNewResponseAccessFacebookTypeSaml       ZeroTrustIdentityProviderNewResponseAccessFacebookType = "saml"
	ZeroTrustIdentityProviderNewResponseAccessFacebookTypeCentrify   ZeroTrustIdentityProviderNewResponseAccessFacebookType = "centrify"
	ZeroTrustIdentityProviderNewResponseAccessFacebookTypeFacebook   ZeroTrustIdentityProviderNewResponseAccessFacebookType = "facebook"
	ZeroTrustIdentityProviderNewResponseAccessFacebookTypeGitHub     ZeroTrustIdentityProviderNewResponseAccessFacebookType = "github"
	ZeroTrustIdentityProviderNewResponseAccessFacebookTypeGoogleApps ZeroTrustIdentityProviderNewResponseAccessFacebookType = "google-apps"
	ZeroTrustIdentityProviderNewResponseAccessFacebookTypeGoogle     ZeroTrustIdentityProviderNewResponseAccessFacebookType = "google"
	ZeroTrustIdentityProviderNewResponseAccessFacebookTypeLinkedin   ZeroTrustIdentityProviderNewResponseAccessFacebookType = "linkedin"
	ZeroTrustIdentityProviderNewResponseAccessFacebookTypeOidc       ZeroTrustIdentityProviderNewResponseAccessFacebookType = "oidc"
	ZeroTrustIdentityProviderNewResponseAccessFacebookTypeOkta       ZeroTrustIdentityProviderNewResponseAccessFacebookType = "okta"
	ZeroTrustIdentityProviderNewResponseAccessFacebookTypeOnelogin   ZeroTrustIdentityProviderNewResponseAccessFacebookType = "onelogin"
	ZeroTrustIdentityProviderNewResponseAccessFacebookTypePingone    ZeroTrustIdentityProviderNewResponseAccessFacebookType = "pingone"
	ZeroTrustIdentityProviderNewResponseAccessFacebookTypeYandex     ZeroTrustIdentityProviderNewResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderNewResponseAccessFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderNewResponseAccessFacebookScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessFacebookScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessFacebookScimConfig]
type zeroTrustIdentityProviderNewResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessFacebookScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderNewResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderNewResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderNewResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderNewResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderNewResponseAccessGitHubJSON       `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessGitHubJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderNewResponseAccessGitHub]
type zeroTrustIdentityProviderNewResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessGitHubJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderNewResponseAccessGitHub) implementsZeroTrustIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                     `json:"client_secret"`
	JSON         zeroTrustIdentityProviderNewResponseAccessGitHubConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessGitHubConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderNewResponseAccessGitHubConfig]
type zeroTrustIdentityProviderNewResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessGitHubConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessGitHubType string

const (
	ZeroTrustIdentityProviderNewResponseAccessGitHubTypeOnetimepin ZeroTrustIdentityProviderNewResponseAccessGitHubType = "onetimepin"
	ZeroTrustIdentityProviderNewResponseAccessGitHubTypeAzureAd    ZeroTrustIdentityProviderNewResponseAccessGitHubType = "azureAD"
	ZeroTrustIdentityProviderNewResponseAccessGitHubTypeSaml       ZeroTrustIdentityProviderNewResponseAccessGitHubType = "saml"
	ZeroTrustIdentityProviderNewResponseAccessGitHubTypeCentrify   ZeroTrustIdentityProviderNewResponseAccessGitHubType = "centrify"
	ZeroTrustIdentityProviderNewResponseAccessGitHubTypeFacebook   ZeroTrustIdentityProviderNewResponseAccessGitHubType = "facebook"
	ZeroTrustIdentityProviderNewResponseAccessGitHubTypeGitHub     ZeroTrustIdentityProviderNewResponseAccessGitHubType = "github"
	ZeroTrustIdentityProviderNewResponseAccessGitHubTypeGoogleApps ZeroTrustIdentityProviderNewResponseAccessGitHubType = "google-apps"
	ZeroTrustIdentityProviderNewResponseAccessGitHubTypeGoogle     ZeroTrustIdentityProviderNewResponseAccessGitHubType = "google"
	ZeroTrustIdentityProviderNewResponseAccessGitHubTypeLinkedin   ZeroTrustIdentityProviderNewResponseAccessGitHubType = "linkedin"
	ZeroTrustIdentityProviderNewResponseAccessGitHubTypeOidc       ZeroTrustIdentityProviderNewResponseAccessGitHubType = "oidc"
	ZeroTrustIdentityProviderNewResponseAccessGitHubTypeOkta       ZeroTrustIdentityProviderNewResponseAccessGitHubType = "okta"
	ZeroTrustIdentityProviderNewResponseAccessGitHubTypeOnelogin   ZeroTrustIdentityProviderNewResponseAccessGitHubType = "onelogin"
	ZeroTrustIdentityProviderNewResponseAccessGitHubTypePingone    ZeroTrustIdentityProviderNewResponseAccessGitHubType = "pingone"
	ZeroTrustIdentityProviderNewResponseAccessGitHubTypeYandex     ZeroTrustIdentityProviderNewResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderNewResponseAccessGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderNewResponseAccessGitHubScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessGitHubScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessGitHubScimConfig]
type zeroTrustIdentityProviderNewResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessGitHubScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderNewResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderNewResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderNewResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderNewResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderNewResponseAccessGoogleJSON       `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessGoogleJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderNewResponseAccessGoogle]
type zeroTrustIdentityProviderNewResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessGoogleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderNewResponseAccessGoogle) implementsZeroTrustIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                     `json:"email_claim_name"`
	JSON           zeroTrustIdentityProviderNewResponseAccessGoogleConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessGoogleConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderNewResponseAccessGoogleConfig]
type zeroTrustIdentityProviderNewResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessGoogleConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessGoogleType string

const (
	ZeroTrustIdentityProviderNewResponseAccessGoogleTypeOnetimepin ZeroTrustIdentityProviderNewResponseAccessGoogleType = "onetimepin"
	ZeroTrustIdentityProviderNewResponseAccessGoogleTypeAzureAd    ZeroTrustIdentityProviderNewResponseAccessGoogleType = "azureAD"
	ZeroTrustIdentityProviderNewResponseAccessGoogleTypeSaml       ZeroTrustIdentityProviderNewResponseAccessGoogleType = "saml"
	ZeroTrustIdentityProviderNewResponseAccessGoogleTypeCentrify   ZeroTrustIdentityProviderNewResponseAccessGoogleType = "centrify"
	ZeroTrustIdentityProviderNewResponseAccessGoogleTypeFacebook   ZeroTrustIdentityProviderNewResponseAccessGoogleType = "facebook"
	ZeroTrustIdentityProviderNewResponseAccessGoogleTypeGitHub     ZeroTrustIdentityProviderNewResponseAccessGoogleType = "github"
	ZeroTrustIdentityProviderNewResponseAccessGoogleTypeGoogleApps ZeroTrustIdentityProviderNewResponseAccessGoogleType = "google-apps"
	ZeroTrustIdentityProviderNewResponseAccessGoogleTypeGoogle     ZeroTrustIdentityProviderNewResponseAccessGoogleType = "google"
	ZeroTrustIdentityProviderNewResponseAccessGoogleTypeLinkedin   ZeroTrustIdentityProviderNewResponseAccessGoogleType = "linkedin"
	ZeroTrustIdentityProviderNewResponseAccessGoogleTypeOidc       ZeroTrustIdentityProviderNewResponseAccessGoogleType = "oidc"
	ZeroTrustIdentityProviderNewResponseAccessGoogleTypeOkta       ZeroTrustIdentityProviderNewResponseAccessGoogleType = "okta"
	ZeroTrustIdentityProviderNewResponseAccessGoogleTypeOnelogin   ZeroTrustIdentityProviderNewResponseAccessGoogleType = "onelogin"
	ZeroTrustIdentityProviderNewResponseAccessGoogleTypePingone    ZeroTrustIdentityProviderNewResponseAccessGoogleType = "pingone"
	ZeroTrustIdentityProviderNewResponseAccessGoogleTypeYandex     ZeroTrustIdentityProviderNewResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderNewResponseAccessGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderNewResponseAccessGoogleScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessGoogleScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessGoogleScimConfig]
type zeroTrustIdentityProviderNewResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessGoogleScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderNewResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderNewResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderNewResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderNewResponseAccessGoogleAppsJSON       `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessGoogleAppsJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderNewResponseAccessGoogleApps]
type zeroTrustIdentityProviderNewResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessGoogleAppsJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderNewResponseAccessGoogleApps) implementsZeroTrustIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessGoogleAppsConfig struct {
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
	JSON           zeroTrustIdentityProviderNewResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessGoogleAppsConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessGoogleAppsConfig]
type zeroTrustIdentityProviderNewResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessGoogleAppsConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType string

const (
	ZeroTrustIdentityProviderNewResponseAccessGoogleAppsTypeOnetimepin ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType = "onetimepin"
	ZeroTrustIdentityProviderNewResponseAccessGoogleAppsTypeAzureAd    ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType = "azureAD"
	ZeroTrustIdentityProviderNewResponseAccessGoogleAppsTypeSaml       ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType = "saml"
	ZeroTrustIdentityProviderNewResponseAccessGoogleAppsTypeCentrify   ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType = "centrify"
	ZeroTrustIdentityProviderNewResponseAccessGoogleAppsTypeFacebook   ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType = "facebook"
	ZeroTrustIdentityProviderNewResponseAccessGoogleAppsTypeGitHub     ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType = "github"
	ZeroTrustIdentityProviderNewResponseAccessGoogleAppsTypeGoogleApps ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType = "google-apps"
	ZeroTrustIdentityProviderNewResponseAccessGoogleAppsTypeGoogle     ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType = "google"
	ZeroTrustIdentityProviderNewResponseAccessGoogleAppsTypeLinkedin   ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType = "linkedin"
	ZeroTrustIdentityProviderNewResponseAccessGoogleAppsTypeOidc       ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType = "oidc"
	ZeroTrustIdentityProviderNewResponseAccessGoogleAppsTypeOkta       ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType = "okta"
	ZeroTrustIdentityProviderNewResponseAccessGoogleAppsTypeOnelogin   ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType = "onelogin"
	ZeroTrustIdentityProviderNewResponseAccessGoogleAppsTypePingone    ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType = "pingone"
	ZeroTrustIdentityProviderNewResponseAccessGoogleAppsTypeYandex     ZeroTrustIdentityProviderNewResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderNewResponseAccessGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderNewResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessGoogleAppsScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessGoogleAppsScimConfig]
type zeroTrustIdentityProviderNewResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessGoogleAppsScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderNewResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderNewResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderNewResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderNewResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderNewResponseAccessLinkedinJSON       `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessLinkedinJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderNewResponseAccessLinkedin]
type zeroTrustIdentityProviderNewResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessLinkedinJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderNewResponseAccessLinkedin) implementsZeroTrustIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                       `json:"client_secret"`
	JSON         zeroTrustIdentityProviderNewResponseAccessLinkedinConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessLinkedinConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessLinkedinConfig]
type zeroTrustIdentityProviderNewResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessLinkedinConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessLinkedinType string

const (
	ZeroTrustIdentityProviderNewResponseAccessLinkedinTypeOnetimepin ZeroTrustIdentityProviderNewResponseAccessLinkedinType = "onetimepin"
	ZeroTrustIdentityProviderNewResponseAccessLinkedinTypeAzureAd    ZeroTrustIdentityProviderNewResponseAccessLinkedinType = "azureAD"
	ZeroTrustIdentityProviderNewResponseAccessLinkedinTypeSaml       ZeroTrustIdentityProviderNewResponseAccessLinkedinType = "saml"
	ZeroTrustIdentityProviderNewResponseAccessLinkedinTypeCentrify   ZeroTrustIdentityProviderNewResponseAccessLinkedinType = "centrify"
	ZeroTrustIdentityProviderNewResponseAccessLinkedinTypeFacebook   ZeroTrustIdentityProviderNewResponseAccessLinkedinType = "facebook"
	ZeroTrustIdentityProviderNewResponseAccessLinkedinTypeGitHub     ZeroTrustIdentityProviderNewResponseAccessLinkedinType = "github"
	ZeroTrustIdentityProviderNewResponseAccessLinkedinTypeGoogleApps ZeroTrustIdentityProviderNewResponseAccessLinkedinType = "google-apps"
	ZeroTrustIdentityProviderNewResponseAccessLinkedinTypeGoogle     ZeroTrustIdentityProviderNewResponseAccessLinkedinType = "google"
	ZeroTrustIdentityProviderNewResponseAccessLinkedinTypeLinkedin   ZeroTrustIdentityProviderNewResponseAccessLinkedinType = "linkedin"
	ZeroTrustIdentityProviderNewResponseAccessLinkedinTypeOidc       ZeroTrustIdentityProviderNewResponseAccessLinkedinType = "oidc"
	ZeroTrustIdentityProviderNewResponseAccessLinkedinTypeOkta       ZeroTrustIdentityProviderNewResponseAccessLinkedinType = "okta"
	ZeroTrustIdentityProviderNewResponseAccessLinkedinTypeOnelogin   ZeroTrustIdentityProviderNewResponseAccessLinkedinType = "onelogin"
	ZeroTrustIdentityProviderNewResponseAccessLinkedinTypePingone    ZeroTrustIdentityProviderNewResponseAccessLinkedinType = "pingone"
	ZeroTrustIdentityProviderNewResponseAccessLinkedinTypeYandex     ZeroTrustIdentityProviderNewResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderNewResponseAccessLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderNewResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessLinkedinScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessLinkedinScimConfig]
type zeroTrustIdentityProviderNewResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessLinkedinScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderNewResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderNewResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderNewResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderNewResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderNewResponseAccessOidcJSON       `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessOidcJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderNewResponseAccessOidc]
type zeroTrustIdentityProviderNewResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessOidcJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderNewResponseAccessOidc) implementsZeroTrustIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessOidcConfig struct {
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
	JSON     zeroTrustIdentityProviderNewResponseAccessOidcConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessOidcConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderNewResponseAccessOidcConfig]
type zeroTrustIdentityProviderNewResponseAccessOidcConfigJSON struct {
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

func (r *ZeroTrustIdentityProviderNewResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessOidcConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessOidcType string

const (
	ZeroTrustIdentityProviderNewResponseAccessOidcTypeOnetimepin ZeroTrustIdentityProviderNewResponseAccessOidcType = "onetimepin"
	ZeroTrustIdentityProviderNewResponseAccessOidcTypeAzureAd    ZeroTrustIdentityProviderNewResponseAccessOidcType = "azureAD"
	ZeroTrustIdentityProviderNewResponseAccessOidcTypeSaml       ZeroTrustIdentityProviderNewResponseAccessOidcType = "saml"
	ZeroTrustIdentityProviderNewResponseAccessOidcTypeCentrify   ZeroTrustIdentityProviderNewResponseAccessOidcType = "centrify"
	ZeroTrustIdentityProviderNewResponseAccessOidcTypeFacebook   ZeroTrustIdentityProviderNewResponseAccessOidcType = "facebook"
	ZeroTrustIdentityProviderNewResponseAccessOidcTypeGitHub     ZeroTrustIdentityProviderNewResponseAccessOidcType = "github"
	ZeroTrustIdentityProviderNewResponseAccessOidcTypeGoogleApps ZeroTrustIdentityProviderNewResponseAccessOidcType = "google-apps"
	ZeroTrustIdentityProviderNewResponseAccessOidcTypeGoogle     ZeroTrustIdentityProviderNewResponseAccessOidcType = "google"
	ZeroTrustIdentityProviderNewResponseAccessOidcTypeLinkedin   ZeroTrustIdentityProviderNewResponseAccessOidcType = "linkedin"
	ZeroTrustIdentityProviderNewResponseAccessOidcTypeOidc       ZeroTrustIdentityProviderNewResponseAccessOidcType = "oidc"
	ZeroTrustIdentityProviderNewResponseAccessOidcTypeOkta       ZeroTrustIdentityProviderNewResponseAccessOidcType = "okta"
	ZeroTrustIdentityProviderNewResponseAccessOidcTypeOnelogin   ZeroTrustIdentityProviderNewResponseAccessOidcType = "onelogin"
	ZeroTrustIdentityProviderNewResponseAccessOidcTypePingone    ZeroTrustIdentityProviderNewResponseAccessOidcType = "pingone"
	ZeroTrustIdentityProviderNewResponseAccessOidcTypeYandex     ZeroTrustIdentityProviderNewResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderNewResponseAccessOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderNewResponseAccessOidcScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessOidcScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessOidcScimConfig]
type zeroTrustIdentityProviderNewResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessOidcScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderNewResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderNewResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderNewResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderNewResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderNewResponseAccessOktaJSON       `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessOktaJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderNewResponseAccessOkta]
type zeroTrustIdentityProviderNewResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessOktaJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderNewResponseAccessOkta) implementsZeroTrustIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessOktaConfig struct {
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
	JSON        zeroTrustIdentityProviderNewResponseAccessOktaConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessOktaConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderNewResponseAccessOktaConfig]
type zeroTrustIdentityProviderNewResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessOktaConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessOktaType string

const (
	ZeroTrustIdentityProviderNewResponseAccessOktaTypeOnetimepin ZeroTrustIdentityProviderNewResponseAccessOktaType = "onetimepin"
	ZeroTrustIdentityProviderNewResponseAccessOktaTypeAzureAd    ZeroTrustIdentityProviderNewResponseAccessOktaType = "azureAD"
	ZeroTrustIdentityProviderNewResponseAccessOktaTypeSaml       ZeroTrustIdentityProviderNewResponseAccessOktaType = "saml"
	ZeroTrustIdentityProviderNewResponseAccessOktaTypeCentrify   ZeroTrustIdentityProviderNewResponseAccessOktaType = "centrify"
	ZeroTrustIdentityProviderNewResponseAccessOktaTypeFacebook   ZeroTrustIdentityProviderNewResponseAccessOktaType = "facebook"
	ZeroTrustIdentityProviderNewResponseAccessOktaTypeGitHub     ZeroTrustIdentityProviderNewResponseAccessOktaType = "github"
	ZeroTrustIdentityProviderNewResponseAccessOktaTypeGoogleApps ZeroTrustIdentityProviderNewResponseAccessOktaType = "google-apps"
	ZeroTrustIdentityProviderNewResponseAccessOktaTypeGoogle     ZeroTrustIdentityProviderNewResponseAccessOktaType = "google"
	ZeroTrustIdentityProviderNewResponseAccessOktaTypeLinkedin   ZeroTrustIdentityProviderNewResponseAccessOktaType = "linkedin"
	ZeroTrustIdentityProviderNewResponseAccessOktaTypeOidc       ZeroTrustIdentityProviderNewResponseAccessOktaType = "oidc"
	ZeroTrustIdentityProviderNewResponseAccessOktaTypeOkta       ZeroTrustIdentityProviderNewResponseAccessOktaType = "okta"
	ZeroTrustIdentityProviderNewResponseAccessOktaTypeOnelogin   ZeroTrustIdentityProviderNewResponseAccessOktaType = "onelogin"
	ZeroTrustIdentityProviderNewResponseAccessOktaTypePingone    ZeroTrustIdentityProviderNewResponseAccessOktaType = "pingone"
	ZeroTrustIdentityProviderNewResponseAccessOktaTypeYandex     ZeroTrustIdentityProviderNewResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderNewResponseAccessOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderNewResponseAccessOktaScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessOktaScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessOktaScimConfig]
type zeroTrustIdentityProviderNewResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessOktaScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderNewResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderNewResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderNewResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderNewResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderNewResponseAccessOneloginJSON       `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessOneloginJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderNewResponseAccessOnelogin]
type zeroTrustIdentityProviderNewResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessOneloginJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderNewResponseAccessOnelogin) implementsZeroTrustIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessOneloginConfig struct {
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
	JSON            zeroTrustIdentityProviderNewResponseAccessOneloginConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessOneloginConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessOneloginConfig]
type zeroTrustIdentityProviderNewResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessOneloginConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessOneloginType string

const (
	ZeroTrustIdentityProviderNewResponseAccessOneloginTypeOnetimepin ZeroTrustIdentityProviderNewResponseAccessOneloginType = "onetimepin"
	ZeroTrustIdentityProviderNewResponseAccessOneloginTypeAzureAd    ZeroTrustIdentityProviderNewResponseAccessOneloginType = "azureAD"
	ZeroTrustIdentityProviderNewResponseAccessOneloginTypeSaml       ZeroTrustIdentityProviderNewResponseAccessOneloginType = "saml"
	ZeroTrustIdentityProviderNewResponseAccessOneloginTypeCentrify   ZeroTrustIdentityProviderNewResponseAccessOneloginType = "centrify"
	ZeroTrustIdentityProviderNewResponseAccessOneloginTypeFacebook   ZeroTrustIdentityProviderNewResponseAccessOneloginType = "facebook"
	ZeroTrustIdentityProviderNewResponseAccessOneloginTypeGitHub     ZeroTrustIdentityProviderNewResponseAccessOneloginType = "github"
	ZeroTrustIdentityProviderNewResponseAccessOneloginTypeGoogleApps ZeroTrustIdentityProviderNewResponseAccessOneloginType = "google-apps"
	ZeroTrustIdentityProviderNewResponseAccessOneloginTypeGoogle     ZeroTrustIdentityProviderNewResponseAccessOneloginType = "google"
	ZeroTrustIdentityProviderNewResponseAccessOneloginTypeLinkedin   ZeroTrustIdentityProviderNewResponseAccessOneloginType = "linkedin"
	ZeroTrustIdentityProviderNewResponseAccessOneloginTypeOidc       ZeroTrustIdentityProviderNewResponseAccessOneloginType = "oidc"
	ZeroTrustIdentityProviderNewResponseAccessOneloginTypeOkta       ZeroTrustIdentityProviderNewResponseAccessOneloginType = "okta"
	ZeroTrustIdentityProviderNewResponseAccessOneloginTypeOnelogin   ZeroTrustIdentityProviderNewResponseAccessOneloginType = "onelogin"
	ZeroTrustIdentityProviderNewResponseAccessOneloginTypePingone    ZeroTrustIdentityProviderNewResponseAccessOneloginType = "pingone"
	ZeroTrustIdentityProviderNewResponseAccessOneloginTypeYandex     ZeroTrustIdentityProviderNewResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderNewResponseAccessOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderNewResponseAccessOneloginScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessOneloginScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessOneloginScimConfig]
type zeroTrustIdentityProviderNewResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessOneloginScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderNewResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderNewResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderNewResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderNewResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderNewResponseAccessPingoneJSON       `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessPingoneJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderNewResponseAccessPingone]
type zeroTrustIdentityProviderNewResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessPingoneJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderNewResponseAccessPingone) implementsZeroTrustIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessPingoneConfig struct {
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
	JSON      zeroTrustIdentityProviderNewResponseAccessPingoneConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessPingoneConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessPingoneConfig]
type zeroTrustIdentityProviderNewResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessPingoneConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessPingoneType string

const (
	ZeroTrustIdentityProviderNewResponseAccessPingoneTypeOnetimepin ZeroTrustIdentityProviderNewResponseAccessPingoneType = "onetimepin"
	ZeroTrustIdentityProviderNewResponseAccessPingoneTypeAzureAd    ZeroTrustIdentityProviderNewResponseAccessPingoneType = "azureAD"
	ZeroTrustIdentityProviderNewResponseAccessPingoneTypeSaml       ZeroTrustIdentityProviderNewResponseAccessPingoneType = "saml"
	ZeroTrustIdentityProviderNewResponseAccessPingoneTypeCentrify   ZeroTrustIdentityProviderNewResponseAccessPingoneType = "centrify"
	ZeroTrustIdentityProviderNewResponseAccessPingoneTypeFacebook   ZeroTrustIdentityProviderNewResponseAccessPingoneType = "facebook"
	ZeroTrustIdentityProviderNewResponseAccessPingoneTypeGitHub     ZeroTrustIdentityProviderNewResponseAccessPingoneType = "github"
	ZeroTrustIdentityProviderNewResponseAccessPingoneTypeGoogleApps ZeroTrustIdentityProviderNewResponseAccessPingoneType = "google-apps"
	ZeroTrustIdentityProviderNewResponseAccessPingoneTypeGoogle     ZeroTrustIdentityProviderNewResponseAccessPingoneType = "google"
	ZeroTrustIdentityProviderNewResponseAccessPingoneTypeLinkedin   ZeroTrustIdentityProviderNewResponseAccessPingoneType = "linkedin"
	ZeroTrustIdentityProviderNewResponseAccessPingoneTypeOidc       ZeroTrustIdentityProviderNewResponseAccessPingoneType = "oidc"
	ZeroTrustIdentityProviderNewResponseAccessPingoneTypeOkta       ZeroTrustIdentityProviderNewResponseAccessPingoneType = "okta"
	ZeroTrustIdentityProviderNewResponseAccessPingoneTypeOnelogin   ZeroTrustIdentityProviderNewResponseAccessPingoneType = "onelogin"
	ZeroTrustIdentityProviderNewResponseAccessPingoneTypePingone    ZeroTrustIdentityProviderNewResponseAccessPingoneType = "pingone"
	ZeroTrustIdentityProviderNewResponseAccessPingoneTypeYandex     ZeroTrustIdentityProviderNewResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderNewResponseAccessPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderNewResponseAccessPingoneScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessPingoneScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessPingoneScimConfig]
type zeroTrustIdentityProviderNewResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessPingoneScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderNewResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderNewResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderNewResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderNewResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderNewResponseAccessSamlJSON       `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessSamlJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderNewResponseAccessSaml]
type zeroTrustIdentityProviderNewResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessSamlJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderNewResponseAccessSaml) implementsZeroTrustIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []ZeroTrustIdentityProviderNewResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                   `json:"sso_target_url"`
	JSON         zeroTrustIdentityProviderNewResponseAccessSamlConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessSamlConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderNewResponseAccessSamlConfig]
type zeroTrustIdentityProviderNewResponseAccessSamlConfigJSON struct {
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

func (r *ZeroTrustIdentityProviderNewResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessSamlConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderNewResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                  `json:"header_name"`
	JSON       zeroTrustIdentityProviderNewResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessSamlConfigHeaderAttributeJSON contains
// the JSON metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessSamlConfigHeaderAttribute]
type zeroTrustIdentityProviderNewResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessSamlConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessSamlType string

const (
	ZeroTrustIdentityProviderNewResponseAccessSamlTypeOnetimepin ZeroTrustIdentityProviderNewResponseAccessSamlType = "onetimepin"
	ZeroTrustIdentityProviderNewResponseAccessSamlTypeAzureAd    ZeroTrustIdentityProviderNewResponseAccessSamlType = "azureAD"
	ZeroTrustIdentityProviderNewResponseAccessSamlTypeSaml       ZeroTrustIdentityProviderNewResponseAccessSamlType = "saml"
	ZeroTrustIdentityProviderNewResponseAccessSamlTypeCentrify   ZeroTrustIdentityProviderNewResponseAccessSamlType = "centrify"
	ZeroTrustIdentityProviderNewResponseAccessSamlTypeFacebook   ZeroTrustIdentityProviderNewResponseAccessSamlType = "facebook"
	ZeroTrustIdentityProviderNewResponseAccessSamlTypeGitHub     ZeroTrustIdentityProviderNewResponseAccessSamlType = "github"
	ZeroTrustIdentityProviderNewResponseAccessSamlTypeGoogleApps ZeroTrustIdentityProviderNewResponseAccessSamlType = "google-apps"
	ZeroTrustIdentityProviderNewResponseAccessSamlTypeGoogle     ZeroTrustIdentityProviderNewResponseAccessSamlType = "google"
	ZeroTrustIdentityProviderNewResponseAccessSamlTypeLinkedin   ZeroTrustIdentityProviderNewResponseAccessSamlType = "linkedin"
	ZeroTrustIdentityProviderNewResponseAccessSamlTypeOidc       ZeroTrustIdentityProviderNewResponseAccessSamlType = "oidc"
	ZeroTrustIdentityProviderNewResponseAccessSamlTypeOkta       ZeroTrustIdentityProviderNewResponseAccessSamlType = "okta"
	ZeroTrustIdentityProviderNewResponseAccessSamlTypeOnelogin   ZeroTrustIdentityProviderNewResponseAccessSamlType = "onelogin"
	ZeroTrustIdentityProviderNewResponseAccessSamlTypePingone    ZeroTrustIdentityProviderNewResponseAccessSamlType = "pingone"
	ZeroTrustIdentityProviderNewResponseAccessSamlTypeYandex     ZeroTrustIdentityProviderNewResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderNewResponseAccessSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderNewResponseAccessSamlScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessSamlScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessSamlScimConfig]
type zeroTrustIdentityProviderNewResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessSamlScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderNewResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderNewResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderNewResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderNewResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderNewResponseAccessYandexJSON       `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessYandexJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderNewResponseAccessYandex]
type zeroTrustIdentityProviderNewResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessYandexJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderNewResponseAccessYandex) implementsZeroTrustIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                     `json:"client_secret"`
	JSON         zeroTrustIdentityProviderNewResponseAccessYandexConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessYandexConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderNewResponseAccessYandexConfig]
type zeroTrustIdentityProviderNewResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessYandexConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessYandexType string

const (
	ZeroTrustIdentityProviderNewResponseAccessYandexTypeOnetimepin ZeroTrustIdentityProviderNewResponseAccessYandexType = "onetimepin"
	ZeroTrustIdentityProviderNewResponseAccessYandexTypeAzureAd    ZeroTrustIdentityProviderNewResponseAccessYandexType = "azureAD"
	ZeroTrustIdentityProviderNewResponseAccessYandexTypeSaml       ZeroTrustIdentityProviderNewResponseAccessYandexType = "saml"
	ZeroTrustIdentityProviderNewResponseAccessYandexTypeCentrify   ZeroTrustIdentityProviderNewResponseAccessYandexType = "centrify"
	ZeroTrustIdentityProviderNewResponseAccessYandexTypeFacebook   ZeroTrustIdentityProviderNewResponseAccessYandexType = "facebook"
	ZeroTrustIdentityProviderNewResponseAccessYandexTypeGitHub     ZeroTrustIdentityProviderNewResponseAccessYandexType = "github"
	ZeroTrustIdentityProviderNewResponseAccessYandexTypeGoogleApps ZeroTrustIdentityProviderNewResponseAccessYandexType = "google-apps"
	ZeroTrustIdentityProviderNewResponseAccessYandexTypeGoogle     ZeroTrustIdentityProviderNewResponseAccessYandexType = "google"
	ZeroTrustIdentityProviderNewResponseAccessYandexTypeLinkedin   ZeroTrustIdentityProviderNewResponseAccessYandexType = "linkedin"
	ZeroTrustIdentityProviderNewResponseAccessYandexTypeOidc       ZeroTrustIdentityProviderNewResponseAccessYandexType = "oidc"
	ZeroTrustIdentityProviderNewResponseAccessYandexTypeOkta       ZeroTrustIdentityProviderNewResponseAccessYandexType = "okta"
	ZeroTrustIdentityProviderNewResponseAccessYandexTypeOnelogin   ZeroTrustIdentityProviderNewResponseAccessYandexType = "onelogin"
	ZeroTrustIdentityProviderNewResponseAccessYandexTypePingone    ZeroTrustIdentityProviderNewResponseAccessYandexType = "pingone"
	ZeroTrustIdentityProviderNewResponseAccessYandexTypeYandex     ZeroTrustIdentityProviderNewResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderNewResponseAccessYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderNewResponseAccessYandexScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessYandexScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessYandexScimConfig]
type zeroTrustIdentityProviderNewResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessYandexScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderNewResponseAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderNewResponseAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderNewResponseAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderNewResponseAccessOnetimepinJSON       `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessOnetimepinJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderNewResponseAccessOnetimepin]
type zeroTrustIdentityProviderNewResponseAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessOnetimepinJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderNewResponseAccessOnetimepin) implementsZeroTrustIdentityProviderNewResponse() {
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderNewResponseAccessOnetimepinType string

const (
	ZeroTrustIdentityProviderNewResponseAccessOnetimepinTypeOnetimepin ZeroTrustIdentityProviderNewResponseAccessOnetimepinType = "onetimepin"
	ZeroTrustIdentityProviderNewResponseAccessOnetimepinTypeAzureAd    ZeroTrustIdentityProviderNewResponseAccessOnetimepinType = "azureAD"
	ZeroTrustIdentityProviderNewResponseAccessOnetimepinTypeSaml       ZeroTrustIdentityProviderNewResponseAccessOnetimepinType = "saml"
	ZeroTrustIdentityProviderNewResponseAccessOnetimepinTypeCentrify   ZeroTrustIdentityProviderNewResponseAccessOnetimepinType = "centrify"
	ZeroTrustIdentityProviderNewResponseAccessOnetimepinTypeFacebook   ZeroTrustIdentityProviderNewResponseAccessOnetimepinType = "facebook"
	ZeroTrustIdentityProviderNewResponseAccessOnetimepinTypeGitHub     ZeroTrustIdentityProviderNewResponseAccessOnetimepinType = "github"
	ZeroTrustIdentityProviderNewResponseAccessOnetimepinTypeGoogleApps ZeroTrustIdentityProviderNewResponseAccessOnetimepinType = "google-apps"
	ZeroTrustIdentityProviderNewResponseAccessOnetimepinTypeGoogle     ZeroTrustIdentityProviderNewResponseAccessOnetimepinType = "google"
	ZeroTrustIdentityProviderNewResponseAccessOnetimepinTypeLinkedin   ZeroTrustIdentityProviderNewResponseAccessOnetimepinType = "linkedin"
	ZeroTrustIdentityProviderNewResponseAccessOnetimepinTypeOidc       ZeroTrustIdentityProviderNewResponseAccessOnetimepinType = "oidc"
	ZeroTrustIdentityProviderNewResponseAccessOnetimepinTypeOkta       ZeroTrustIdentityProviderNewResponseAccessOnetimepinType = "okta"
	ZeroTrustIdentityProviderNewResponseAccessOnetimepinTypeOnelogin   ZeroTrustIdentityProviderNewResponseAccessOnetimepinType = "onelogin"
	ZeroTrustIdentityProviderNewResponseAccessOnetimepinTypePingone    ZeroTrustIdentityProviderNewResponseAccessOnetimepinType = "pingone"
	ZeroTrustIdentityProviderNewResponseAccessOnetimepinTypeYandex     ZeroTrustIdentityProviderNewResponseAccessOnetimepinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderNewResponseAccessOnetimepinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderNewResponseAccessOnetimepinScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderNewResponseAccessOnetimepinScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderNewResponseAccessOnetimepinScimConfig]
type zeroTrustIdentityProviderNewResponseAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderNewResponseAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderNewResponseAccessOnetimepinScimConfigJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ZeroTrustIdentityProviderUpdateResponseAccessAzureAd],
// [ZeroTrustIdentityProviderUpdateResponseAccessCentrify],
// [ZeroTrustIdentityProviderUpdateResponseAccessFacebook],
// [ZeroTrustIdentityProviderUpdateResponseAccessGitHub],
// [ZeroTrustIdentityProviderUpdateResponseAccessGoogle],
// [ZeroTrustIdentityProviderUpdateResponseAccessGoogleApps],
// [ZeroTrustIdentityProviderUpdateResponseAccessLinkedin],
// [ZeroTrustIdentityProviderUpdateResponseAccessOidc],
// [ZeroTrustIdentityProviderUpdateResponseAccessOkta],
// [ZeroTrustIdentityProviderUpdateResponseAccessOnelogin],
// [ZeroTrustIdentityProviderUpdateResponseAccessPingone],
// [ZeroTrustIdentityProviderUpdateResponseAccessSaml],
// [ZeroTrustIdentityProviderUpdateResponseAccessYandex] or
// [ZeroTrustIdentityProviderUpdateResponseAccessOnetimepin].
type ZeroTrustIdentityProviderUpdateResponse interface {
	implementsZeroTrustIdentityProviderUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustIdentityProviderUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderUpdateResponseAccessAzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderUpdateResponseAccessCentrify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderUpdateResponseAccessFacebook{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderUpdateResponseAccessGitHub{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderUpdateResponseAccessGoogle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderUpdateResponseAccessGoogleApps{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderUpdateResponseAccessLinkedin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderUpdateResponseAccessOidc{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderUpdateResponseAccessOkta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderUpdateResponseAccessOnelogin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderUpdateResponseAccessPingone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderUpdateResponseAccessSaml{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderUpdateResponseAccessYandex{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderUpdateResponseAccessOnetimepin{}),
		},
	)
}

type ZeroTrustIdentityProviderUpdateResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderUpdateResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderUpdateResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessAzureAdJSON       `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessAzureAdJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderUpdateResponseAccessAzureAd]
type zeroTrustIdentityProviderUpdateResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessAzureAdJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderUpdateResponseAccessAzureAd) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessAzureAdConfig struct {
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
	JSON          zeroTrustIdentityProviderUpdateResponseAccessAzureAdConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessAzureAdConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessAzureAdConfig]
type zeroTrustIdentityProviderUpdateResponseAccessAzureAdConfigJSON struct {
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

func (r *ZeroTrustIdentityProviderUpdateResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessAzureAdConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType string

const (
	ZeroTrustIdentityProviderUpdateResponseAccessAzureAdTypeOnetimepin ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType = "onetimepin"
	ZeroTrustIdentityProviderUpdateResponseAccessAzureAdTypeAzureAd    ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType = "azureAD"
	ZeroTrustIdentityProviderUpdateResponseAccessAzureAdTypeSaml       ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType = "saml"
	ZeroTrustIdentityProviderUpdateResponseAccessAzureAdTypeCentrify   ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType = "centrify"
	ZeroTrustIdentityProviderUpdateResponseAccessAzureAdTypeFacebook   ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType = "facebook"
	ZeroTrustIdentityProviderUpdateResponseAccessAzureAdTypeGitHub     ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType = "github"
	ZeroTrustIdentityProviderUpdateResponseAccessAzureAdTypeGoogleApps ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType = "google-apps"
	ZeroTrustIdentityProviderUpdateResponseAccessAzureAdTypeGoogle     ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType = "google"
	ZeroTrustIdentityProviderUpdateResponseAccessAzureAdTypeLinkedin   ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType = "linkedin"
	ZeroTrustIdentityProviderUpdateResponseAccessAzureAdTypeOidc       ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType = "oidc"
	ZeroTrustIdentityProviderUpdateResponseAccessAzureAdTypeOkta       ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType = "okta"
	ZeroTrustIdentityProviderUpdateResponseAccessAzureAdTypeOnelogin   ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType = "onelogin"
	ZeroTrustIdentityProviderUpdateResponseAccessAzureAdTypePingone    ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType = "pingone"
	ZeroTrustIdentityProviderUpdateResponseAccessAzureAdTypeYandex     ZeroTrustIdentityProviderUpdateResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderUpdateResponseAccessAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessAzureAdScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessAzureAdScimConfig]
type zeroTrustIdentityProviderUpdateResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessAzureAdScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderUpdateResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderUpdateResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderUpdateResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessCentrifyJSON       `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessCentrifyJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderUpdateResponseAccessCentrify]
type zeroTrustIdentityProviderUpdateResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessCentrifyJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderUpdateResponseAccessCentrify) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessCentrifyConfig struct {
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
	JSON           zeroTrustIdentityProviderUpdateResponseAccessCentrifyConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessCentrifyConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessCentrifyConfig]
type zeroTrustIdentityProviderUpdateResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessCentrifyConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType string

const (
	ZeroTrustIdentityProviderUpdateResponseAccessCentrifyTypeOnetimepin ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType = "onetimepin"
	ZeroTrustIdentityProviderUpdateResponseAccessCentrifyTypeAzureAd    ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType = "azureAD"
	ZeroTrustIdentityProviderUpdateResponseAccessCentrifyTypeSaml       ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType = "saml"
	ZeroTrustIdentityProviderUpdateResponseAccessCentrifyTypeCentrify   ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType = "centrify"
	ZeroTrustIdentityProviderUpdateResponseAccessCentrifyTypeFacebook   ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType = "facebook"
	ZeroTrustIdentityProviderUpdateResponseAccessCentrifyTypeGitHub     ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType = "github"
	ZeroTrustIdentityProviderUpdateResponseAccessCentrifyTypeGoogleApps ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType = "google-apps"
	ZeroTrustIdentityProviderUpdateResponseAccessCentrifyTypeGoogle     ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType = "google"
	ZeroTrustIdentityProviderUpdateResponseAccessCentrifyTypeLinkedin   ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType = "linkedin"
	ZeroTrustIdentityProviderUpdateResponseAccessCentrifyTypeOidc       ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType = "oidc"
	ZeroTrustIdentityProviderUpdateResponseAccessCentrifyTypeOkta       ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType = "okta"
	ZeroTrustIdentityProviderUpdateResponseAccessCentrifyTypeOnelogin   ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType = "onelogin"
	ZeroTrustIdentityProviderUpdateResponseAccessCentrifyTypePingone    ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType = "pingone"
	ZeroTrustIdentityProviderUpdateResponseAccessCentrifyTypeYandex     ZeroTrustIdentityProviderUpdateResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderUpdateResponseAccessCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessCentrifyScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessCentrifyScimConfig]
type zeroTrustIdentityProviderUpdateResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessCentrifyScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderUpdateResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderUpdateResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderUpdateResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderUpdateResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessFacebookJSON       `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessFacebookJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderUpdateResponseAccessFacebook]
type zeroTrustIdentityProviderUpdateResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessFacebookJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderUpdateResponseAccessFacebook) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                          `json:"client_secret"`
	JSON         zeroTrustIdentityProviderUpdateResponseAccessFacebookConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessFacebookConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessFacebookConfig]
type zeroTrustIdentityProviderUpdateResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessFacebookConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessFacebookType string

const (
	ZeroTrustIdentityProviderUpdateResponseAccessFacebookTypeOnetimepin ZeroTrustIdentityProviderUpdateResponseAccessFacebookType = "onetimepin"
	ZeroTrustIdentityProviderUpdateResponseAccessFacebookTypeAzureAd    ZeroTrustIdentityProviderUpdateResponseAccessFacebookType = "azureAD"
	ZeroTrustIdentityProviderUpdateResponseAccessFacebookTypeSaml       ZeroTrustIdentityProviderUpdateResponseAccessFacebookType = "saml"
	ZeroTrustIdentityProviderUpdateResponseAccessFacebookTypeCentrify   ZeroTrustIdentityProviderUpdateResponseAccessFacebookType = "centrify"
	ZeroTrustIdentityProviderUpdateResponseAccessFacebookTypeFacebook   ZeroTrustIdentityProviderUpdateResponseAccessFacebookType = "facebook"
	ZeroTrustIdentityProviderUpdateResponseAccessFacebookTypeGitHub     ZeroTrustIdentityProviderUpdateResponseAccessFacebookType = "github"
	ZeroTrustIdentityProviderUpdateResponseAccessFacebookTypeGoogleApps ZeroTrustIdentityProviderUpdateResponseAccessFacebookType = "google-apps"
	ZeroTrustIdentityProviderUpdateResponseAccessFacebookTypeGoogle     ZeroTrustIdentityProviderUpdateResponseAccessFacebookType = "google"
	ZeroTrustIdentityProviderUpdateResponseAccessFacebookTypeLinkedin   ZeroTrustIdentityProviderUpdateResponseAccessFacebookType = "linkedin"
	ZeroTrustIdentityProviderUpdateResponseAccessFacebookTypeOidc       ZeroTrustIdentityProviderUpdateResponseAccessFacebookType = "oidc"
	ZeroTrustIdentityProviderUpdateResponseAccessFacebookTypeOkta       ZeroTrustIdentityProviderUpdateResponseAccessFacebookType = "okta"
	ZeroTrustIdentityProviderUpdateResponseAccessFacebookTypeOnelogin   ZeroTrustIdentityProviderUpdateResponseAccessFacebookType = "onelogin"
	ZeroTrustIdentityProviderUpdateResponseAccessFacebookTypePingone    ZeroTrustIdentityProviderUpdateResponseAccessFacebookType = "pingone"
	ZeroTrustIdentityProviderUpdateResponseAccessFacebookTypeYandex     ZeroTrustIdentityProviderUpdateResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderUpdateResponseAccessFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessFacebookScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessFacebookScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessFacebookScimConfig]
type zeroTrustIdentityProviderUpdateResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessFacebookScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderUpdateResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderUpdateResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderUpdateResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderUpdateResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessGitHubJSON       `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessGitHubJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderUpdateResponseAccessGitHub]
type zeroTrustIdentityProviderUpdateResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessGitHubJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderUpdateResponseAccessGitHub) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                        `json:"client_secret"`
	JSON         zeroTrustIdentityProviderUpdateResponseAccessGitHubConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessGitHubConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessGitHubConfig]
type zeroTrustIdentityProviderUpdateResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessGitHubConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessGitHubType string

const (
	ZeroTrustIdentityProviderUpdateResponseAccessGitHubTypeOnetimepin ZeroTrustIdentityProviderUpdateResponseAccessGitHubType = "onetimepin"
	ZeroTrustIdentityProviderUpdateResponseAccessGitHubTypeAzureAd    ZeroTrustIdentityProviderUpdateResponseAccessGitHubType = "azureAD"
	ZeroTrustIdentityProviderUpdateResponseAccessGitHubTypeSaml       ZeroTrustIdentityProviderUpdateResponseAccessGitHubType = "saml"
	ZeroTrustIdentityProviderUpdateResponseAccessGitHubTypeCentrify   ZeroTrustIdentityProviderUpdateResponseAccessGitHubType = "centrify"
	ZeroTrustIdentityProviderUpdateResponseAccessGitHubTypeFacebook   ZeroTrustIdentityProviderUpdateResponseAccessGitHubType = "facebook"
	ZeroTrustIdentityProviderUpdateResponseAccessGitHubTypeGitHub     ZeroTrustIdentityProviderUpdateResponseAccessGitHubType = "github"
	ZeroTrustIdentityProviderUpdateResponseAccessGitHubTypeGoogleApps ZeroTrustIdentityProviderUpdateResponseAccessGitHubType = "google-apps"
	ZeroTrustIdentityProviderUpdateResponseAccessGitHubTypeGoogle     ZeroTrustIdentityProviderUpdateResponseAccessGitHubType = "google"
	ZeroTrustIdentityProviderUpdateResponseAccessGitHubTypeLinkedin   ZeroTrustIdentityProviderUpdateResponseAccessGitHubType = "linkedin"
	ZeroTrustIdentityProviderUpdateResponseAccessGitHubTypeOidc       ZeroTrustIdentityProviderUpdateResponseAccessGitHubType = "oidc"
	ZeroTrustIdentityProviderUpdateResponseAccessGitHubTypeOkta       ZeroTrustIdentityProviderUpdateResponseAccessGitHubType = "okta"
	ZeroTrustIdentityProviderUpdateResponseAccessGitHubTypeOnelogin   ZeroTrustIdentityProviderUpdateResponseAccessGitHubType = "onelogin"
	ZeroTrustIdentityProviderUpdateResponseAccessGitHubTypePingone    ZeroTrustIdentityProviderUpdateResponseAccessGitHubType = "pingone"
	ZeroTrustIdentityProviderUpdateResponseAccessGitHubTypeYandex     ZeroTrustIdentityProviderUpdateResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderUpdateResponseAccessGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessGitHubScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessGitHubScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessGitHubScimConfig]
type zeroTrustIdentityProviderUpdateResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessGitHubScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderUpdateResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderUpdateResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderUpdateResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderUpdateResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessGoogleJSON       `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessGoogleJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderUpdateResponseAccessGoogle]
type zeroTrustIdentityProviderUpdateResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessGoogleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderUpdateResponseAccessGoogle) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                        `json:"email_claim_name"`
	JSON           zeroTrustIdentityProviderUpdateResponseAccessGoogleConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessGoogleConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessGoogleConfig]
type zeroTrustIdentityProviderUpdateResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessGoogleConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessGoogleType string

const (
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleTypeOnetimepin ZeroTrustIdentityProviderUpdateResponseAccessGoogleType = "onetimepin"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleTypeAzureAd    ZeroTrustIdentityProviderUpdateResponseAccessGoogleType = "azureAD"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleTypeSaml       ZeroTrustIdentityProviderUpdateResponseAccessGoogleType = "saml"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleTypeCentrify   ZeroTrustIdentityProviderUpdateResponseAccessGoogleType = "centrify"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleTypeFacebook   ZeroTrustIdentityProviderUpdateResponseAccessGoogleType = "facebook"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleTypeGitHub     ZeroTrustIdentityProviderUpdateResponseAccessGoogleType = "github"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleTypeGoogleApps ZeroTrustIdentityProviderUpdateResponseAccessGoogleType = "google-apps"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleTypeGoogle     ZeroTrustIdentityProviderUpdateResponseAccessGoogleType = "google"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleTypeLinkedin   ZeroTrustIdentityProviderUpdateResponseAccessGoogleType = "linkedin"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleTypeOidc       ZeroTrustIdentityProviderUpdateResponseAccessGoogleType = "oidc"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleTypeOkta       ZeroTrustIdentityProviderUpdateResponseAccessGoogleType = "okta"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleTypeOnelogin   ZeroTrustIdentityProviderUpdateResponseAccessGoogleType = "onelogin"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleTypePingone    ZeroTrustIdentityProviderUpdateResponseAccessGoogleType = "pingone"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleTypeYandex     ZeroTrustIdentityProviderUpdateResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderUpdateResponseAccessGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessGoogleScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessGoogleScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessGoogleScimConfig]
type zeroTrustIdentityProviderUpdateResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessGoogleScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderUpdateResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessGoogleAppsJSON       `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessGoogleAppsJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessGoogleApps]
type zeroTrustIdentityProviderUpdateResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessGoogleAppsJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderUpdateResponseAccessGoogleApps) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsConfig struct {
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
	JSON           zeroTrustIdentityProviderUpdateResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessGoogleAppsConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsConfig]
type zeroTrustIdentityProviderUpdateResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessGoogleAppsConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType string

const (
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsTypeOnetimepin ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType = "onetimepin"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsTypeAzureAd    ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType = "azureAD"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsTypeSaml       ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType = "saml"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsTypeCentrify   ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType = "centrify"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsTypeFacebook   ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType = "facebook"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsTypeGitHub     ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType = "github"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsTypeGoogleApps ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType = "google-apps"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsTypeGoogle     ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType = "google"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsTypeLinkedin   ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType = "linkedin"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsTypeOidc       ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType = "oidc"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsTypeOkta       ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType = "okta"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsTypeOnelogin   ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType = "onelogin"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsTypePingone    ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType = "pingone"
	ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsTypeYandex     ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessGoogleAppsScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsScimConfig]
type zeroTrustIdentityProviderUpdateResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessGoogleAppsScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderUpdateResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderUpdateResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderUpdateResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessLinkedinJSON       `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessLinkedinJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderUpdateResponseAccessLinkedin]
type zeroTrustIdentityProviderUpdateResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessLinkedinJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderUpdateResponseAccessLinkedin) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                          `json:"client_secret"`
	JSON         zeroTrustIdentityProviderUpdateResponseAccessLinkedinConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessLinkedinConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessLinkedinConfig]
type zeroTrustIdentityProviderUpdateResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessLinkedinConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType string

const (
	ZeroTrustIdentityProviderUpdateResponseAccessLinkedinTypeOnetimepin ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType = "onetimepin"
	ZeroTrustIdentityProviderUpdateResponseAccessLinkedinTypeAzureAd    ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType = "azureAD"
	ZeroTrustIdentityProviderUpdateResponseAccessLinkedinTypeSaml       ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType = "saml"
	ZeroTrustIdentityProviderUpdateResponseAccessLinkedinTypeCentrify   ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType = "centrify"
	ZeroTrustIdentityProviderUpdateResponseAccessLinkedinTypeFacebook   ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType = "facebook"
	ZeroTrustIdentityProviderUpdateResponseAccessLinkedinTypeGitHub     ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType = "github"
	ZeroTrustIdentityProviderUpdateResponseAccessLinkedinTypeGoogleApps ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType = "google-apps"
	ZeroTrustIdentityProviderUpdateResponseAccessLinkedinTypeGoogle     ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType = "google"
	ZeroTrustIdentityProviderUpdateResponseAccessLinkedinTypeLinkedin   ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType = "linkedin"
	ZeroTrustIdentityProviderUpdateResponseAccessLinkedinTypeOidc       ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType = "oidc"
	ZeroTrustIdentityProviderUpdateResponseAccessLinkedinTypeOkta       ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType = "okta"
	ZeroTrustIdentityProviderUpdateResponseAccessLinkedinTypeOnelogin   ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType = "onelogin"
	ZeroTrustIdentityProviderUpdateResponseAccessLinkedinTypePingone    ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType = "pingone"
	ZeroTrustIdentityProviderUpdateResponseAccessLinkedinTypeYandex     ZeroTrustIdentityProviderUpdateResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderUpdateResponseAccessLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessLinkedinScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessLinkedinScimConfig]
type zeroTrustIdentityProviderUpdateResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessLinkedinScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderUpdateResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderUpdateResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderUpdateResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderUpdateResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessOidcJSON       `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessOidcJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderUpdateResponseAccessOidc]
type zeroTrustIdentityProviderUpdateResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessOidcJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderUpdateResponseAccessOidc) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessOidcConfig struct {
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
	JSON     zeroTrustIdentityProviderUpdateResponseAccessOidcConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessOidcConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessOidcConfig]
type zeroTrustIdentityProviderUpdateResponseAccessOidcConfigJSON struct {
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

func (r *ZeroTrustIdentityProviderUpdateResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessOidcConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessOidcType string

const (
	ZeroTrustIdentityProviderUpdateResponseAccessOidcTypeOnetimepin ZeroTrustIdentityProviderUpdateResponseAccessOidcType = "onetimepin"
	ZeroTrustIdentityProviderUpdateResponseAccessOidcTypeAzureAd    ZeroTrustIdentityProviderUpdateResponseAccessOidcType = "azureAD"
	ZeroTrustIdentityProviderUpdateResponseAccessOidcTypeSaml       ZeroTrustIdentityProviderUpdateResponseAccessOidcType = "saml"
	ZeroTrustIdentityProviderUpdateResponseAccessOidcTypeCentrify   ZeroTrustIdentityProviderUpdateResponseAccessOidcType = "centrify"
	ZeroTrustIdentityProviderUpdateResponseAccessOidcTypeFacebook   ZeroTrustIdentityProviderUpdateResponseAccessOidcType = "facebook"
	ZeroTrustIdentityProviderUpdateResponseAccessOidcTypeGitHub     ZeroTrustIdentityProviderUpdateResponseAccessOidcType = "github"
	ZeroTrustIdentityProviderUpdateResponseAccessOidcTypeGoogleApps ZeroTrustIdentityProviderUpdateResponseAccessOidcType = "google-apps"
	ZeroTrustIdentityProviderUpdateResponseAccessOidcTypeGoogle     ZeroTrustIdentityProviderUpdateResponseAccessOidcType = "google"
	ZeroTrustIdentityProviderUpdateResponseAccessOidcTypeLinkedin   ZeroTrustIdentityProviderUpdateResponseAccessOidcType = "linkedin"
	ZeroTrustIdentityProviderUpdateResponseAccessOidcTypeOidc       ZeroTrustIdentityProviderUpdateResponseAccessOidcType = "oidc"
	ZeroTrustIdentityProviderUpdateResponseAccessOidcTypeOkta       ZeroTrustIdentityProviderUpdateResponseAccessOidcType = "okta"
	ZeroTrustIdentityProviderUpdateResponseAccessOidcTypeOnelogin   ZeroTrustIdentityProviderUpdateResponseAccessOidcType = "onelogin"
	ZeroTrustIdentityProviderUpdateResponseAccessOidcTypePingone    ZeroTrustIdentityProviderUpdateResponseAccessOidcType = "pingone"
	ZeroTrustIdentityProviderUpdateResponseAccessOidcTypeYandex     ZeroTrustIdentityProviderUpdateResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderUpdateResponseAccessOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessOidcScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessOidcScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessOidcScimConfig]
type zeroTrustIdentityProviderUpdateResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessOidcScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderUpdateResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderUpdateResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderUpdateResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderUpdateResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessOktaJSON       `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessOktaJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderUpdateResponseAccessOkta]
type zeroTrustIdentityProviderUpdateResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessOktaJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderUpdateResponseAccessOkta) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessOktaConfig struct {
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
	JSON        zeroTrustIdentityProviderUpdateResponseAccessOktaConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessOktaConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessOktaConfig]
type zeroTrustIdentityProviderUpdateResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessOktaConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessOktaType string

const (
	ZeroTrustIdentityProviderUpdateResponseAccessOktaTypeOnetimepin ZeroTrustIdentityProviderUpdateResponseAccessOktaType = "onetimepin"
	ZeroTrustIdentityProviderUpdateResponseAccessOktaTypeAzureAd    ZeroTrustIdentityProviderUpdateResponseAccessOktaType = "azureAD"
	ZeroTrustIdentityProviderUpdateResponseAccessOktaTypeSaml       ZeroTrustIdentityProviderUpdateResponseAccessOktaType = "saml"
	ZeroTrustIdentityProviderUpdateResponseAccessOktaTypeCentrify   ZeroTrustIdentityProviderUpdateResponseAccessOktaType = "centrify"
	ZeroTrustIdentityProviderUpdateResponseAccessOktaTypeFacebook   ZeroTrustIdentityProviderUpdateResponseAccessOktaType = "facebook"
	ZeroTrustIdentityProviderUpdateResponseAccessOktaTypeGitHub     ZeroTrustIdentityProviderUpdateResponseAccessOktaType = "github"
	ZeroTrustIdentityProviderUpdateResponseAccessOktaTypeGoogleApps ZeroTrustIdentityProviderUpdateResponseAccessOktaType = "google-apps"
	ZeroTrustIdentityProviderUpdateResponseAccessOktaTypeGoogle     ZeroTrustIdentityProviderUpdateResponseAccessOktaType = "google"
	ZeroTrustIdentityProviderUpdateResponseAccessOktaTypeLinkedin   ZeroTrustIdentityProviderUpdateResponseAccessOktaType = "linkedin"
	ZeroTrustIdentityProviderUpdateResponseAccessOktaTypeOidc       ZeroTrustIdentityProviderUpdateResponseAccessOktaType = "oidc"
	ZeroTrustIdentityProviderUpdateResponseAccessOktaTypeOkta       ZeroTrustIdentityProviderUpdateResponseAccessOktaType = "okta"
	ZeroTrustIdentityProviderUpdateResponseAccessOktaTypeOnelogin   ZeroTrustIdentityProviderUpdateResponseAccessOktaType = "onelogin"
	ZeroTrustIdentityProviderUpdateResponseAccessOktaTypePingone    ZeroTrustIdentityProviderUpdateResponseAccessOktaType = "pingone"
	ZeroTrustIdentityProviderUpdateResponseAccessOktaTypeYandex     ZeroTrustIdentityProviderUpdateResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderUpdateResponseAccessOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessOktaScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessOktaScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessOktaScimConfig]
type zeroTrustIdentityProviderUpdateResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessOktaScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderUpdateResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderUpdateResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderUpdateResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderUpdateResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessOneloginJSON       `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessOneloginJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderUpdateResponseAccessOnelogin]
type zeroTrustIdentityProviderUpdateResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessOneloginJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderUpdateResponseAccessOnelogin) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessOneloginConfig struct {
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessOneloginConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessOneloginConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessOneloginConfig]
type zeroTrustIdentityProviderUpdateResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessOneloginConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessOneloginType string

const (
	ZeroTrustIdentityProviderUpdateResponseAccessOneloginTypeOnetimepin ZeroTrustIdentityProviderUpdateResponseAccessOneloginType = "onetimepin"
	ZeroTrustIdentityProviderUpdateResponseAccessOneloginTypeAzureAd    ZeroTrustIdentityProviderUpdateResponseAccessOneloginType = "azureAD"
	ZeroTrustIdentityProviderUpdateResponseAccessOneloginTypeSaml       ZeroTrustIdentityProviderUpdateResponseAccessOneloginType = "saml"
	ZeroTrustIdentityProviderUpdateResponseAccessOneloginTypeCentrify   ZeroTrustIdentityProviderUpdateResponseAccessOneloginType = "centrify"
	ZeroTrustIdentityProviderUpdateResponseAccessOneloginTypeFacebook   ZeroTrustIdentityProviderUpdateResponseAccessOneloginType = "facebook"
	ZeroTrustIdentityProviderUpdateResponseAccessOneloginTypeGitHub     ZeroTrustIdentityProviderUpdateResponseAccessOneloginType = "github"
	ZeroTrustIdentityProviderUpdateResponseAccessOneloginTypeGoogleApps ZeroTrustIdentityProviderUpdateResponseAccessOneloginType = "google-apps"
	ZeroTrustIdentityProviderUpdateResponseAccessOneloginTypeGoogle     ZeroTrustIdentityProviderUpdateResponseAccessOneloginType = "google"
	ZeroTrustIdentityProviderUpdateResponseAccessOneloginTypeLinkedin   ZeroTrustIdentityProviderUpdateResponseAccessOneloginType = "linkedin"
	ZeroTrustIdentityProviderUpdateResponseAccessOneloginTypeOidc       ZeroTrustIdentityProviderUpdateResponseAccessOneloginType = "oidc"
	ZeroTrustIdentityProviderUpdateResponseAccessOneloginTypeOkta       ZeroTrustIdentityProviderUpdateResponseAccessOneloginType = "okta"
	ZeroTrustIdentityProviderUpdateResponseAccessOneloginTypeOnelogin   ZeroTrustIdentityProviderUpdateResponseAccessOneloginType = "onelogin"
	ZeroTrustIdentityProviderUpdateResponseAccessOneloginTypePingone    ZeroTrustIdentityProviderUpdateResponseAccessOneloginType = "pingone"
	ZeroTrustIdentityProviderUpdateResponseAccessOneloginTypeYandex     ZeroTrustIdentityProviderUpdateResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderUpdateResponseAccessOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessOneloginScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessOneloginScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessOneloginScimConfig]
type zeroTrustIdentityProviderUpdateResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessOneloginScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderUpdateResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderUpdateResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderUpdateResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderUpdateResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessPingoneJSON       `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessPingoneJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderUpdateResponseAccessPingone]
type zeroTrustIdentityProviderUpdateResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessPingoneJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderUpdateResponseAccessPingone) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessPingoneConfig struct {
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
	JSON      zeroTrustIdentityProviderUpdateResponseAccessPingoneConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessPingoneConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessPingoneConfig]
type zeroTrustIdentityProviderUpdateResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessPingoneConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessPingoneType string

const (
	ZeroTrustIdentityProviderUpdateResponseAccessPingoneTypeOnetimepin ZeroTrustIdentityProviderUpdateResponseAccessPingoneType = "onetimepin"
	ZeroTrustIdentityProviderUpdateResponseAccessPingoneTypeAzureAd    ZeroTrustIdentityProviderUpdateResponseAccessPingoneType = "azureAD"
	ZeroTrustIdentityProviderUpdateResponseAccessPingoneTypeSaml       ZeroTrustIdentityProviderUpdateResponseAccessPingoneType = "saml"
	ZeroTrustIdentityProviderUpdateResponseAccessPingoneTypeCentrify   ZeroTrustIdentityProviderUpdateResponseAccessPingoneType = "centrify"
	ZeroTrustIdentityProviderUpdateResponseAccessPingoneTypeFacebook   ZeroTrustIdentityProviderUpdateResponseAccessPingoneType = "facebook"
	ZeroTrustIdentityProviderUpdateResponseAccessPingoneTypeGitHub     ZeroTrustIdentityProviderUpdateResponseAccessPingoneType = "github"
	ZeroTrustIdentityProviderUpdateResponseAccessPingoneTypeGoogleApps ZeroTrustIdentityProviderUpdateResponseAccessPingoneType = "google-apps"
	ZeroTrustIdentityProviderUpdateResponseAccessPingoneTypeGoogle     ZeroTrustIdentityProviderUpdateResponseAccessPingoneType = "google"
	ZeroTrustIdentityProviderUpdateResponseAccessPingoneTypeLinkedin   ZeroTrustIdentityProviderUpdateResponseAccessPingoneType = "linkedin"
	ZeroTrustIdentityProviderUpdateResponseAccessPingoneTypeOidc       ZeroTrustIdentityProviderUpdateResponseAccessPingoneType = "oidc"
	ZeroTrustIdentityProviderUpdateResponseAccessPingoneTypeOkta       ZeroTrustIdentityProviderUpdateResponseAccessPingoneType = "okta"
	ZeroTrustIdentityProviderUpdateResponseAccessPingoneTypeOnelogin   ZeroTrustIdentityProviderUpdateResponseAccessPingoneType = "onelogin"
	ZeroTrustIdentityProviderUpdateResponseAccessPingoneTypePingone    ZeroTrustIdentityProviderUpdateResponseAccessPingoneType = "pingone"
	ZeroTrustIdentityProviderUpdateResponseAccessPingoneTypeYandex     ZeroTrustIdentityProviderUpdateResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderUpdateResponseAccessPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessPingoneScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessPingoneScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessPingoneScimConfig]
type zeroTrustIdentityProviderUpdateResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessPingoneScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderUpdateResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderUpdateResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderUpdateResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderUpdateResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessSamlJSON       `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessSamlJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderUpdateResponseAccessSaml]
type zeroTrustIdentityProviderUpdateResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessSamlJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderUpdateResponseAccessSaml) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []ZeroTrustIdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                      `json:"sso_target_url"`
	JSON         zeroTrustIdentityProviderUpdateResponseAccessSamlConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessSamlConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessSamlConfig]
type zeroTrustIdentityProviderUpdateResponseAccessSamlConfigJSON struct {
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

func (r *ZeroTrustIdentityProviderUpdateResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessSamlConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                     `json:"header_name"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute]
type zeroTrustIdentityProviderUpdateResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessSamlConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessSamlType string

const (
	ZeroTrustIdentityProviderUpdateResponseAccessSamlTypeOnetimepin ZeroTrustIdentityProviderUpdateResponseAccessSamlType = "onetimepin"
	ZeroTrustIdentityProviderUpdateResponseAccessSamlTypeAzureAd    ZeroTrustIdentityProviderUpdateResponseAccessSamlType = "azureAD"
	ZeroTrustIdentityProviderUpdateResponseAccessSamlTypeSaml       ZeroTrustIdentityProviderUpdateResponseAccessSamlType = "saml"
	ZeroTrustIdentityProviderUpdateResponseAccessSamlTypeCentrify   ZeroTrustIdentityProviderUpdateResponseAccessSamlType = "centrify"
	ZeroTrustIdentityProviderUpdateResponseAccessSamlTypeFacebook   ZeroTrustIdentityProviderUpdateResponseAccessSamlType = "facebook"
	ZeroTrustIdentityProviderUpdateResponseAccessSamlTypeGitHub     ZeroTrustIdentityProviderUpdateResponseAccessSamlType = "github"
	ZeroTrustIdentityProviderUpdateResponseAccessSamlTypeGoogleApps ZeroTrustIdentityProviderUpdateResponseAccessSamlType = "google-apps"
	ZeroTrustIdentityProviderUpdateResponseAccessSamlTypeGoogle     ZeroTrustIdentityProviderUpdateResponseAccessSamlType = "google"
	ZeroTrustIdentityProviderUpdateResponseAccessSamlTypeLinkedin   ZeroTrustIdentityProviderUpdateResponseAccessSamlType = "linkedin"
	ZeroTrustIdentityProviderUpdateResponseAccessSamlTypeOidc       ZeroTrustIdentityProviderUpdateResponseAccessSamlType = "oidc"
	ZeroTrustIdentityProviderUpdateResponseAccessSamlTypeOkta       ZeroTrustIdentityProviderUpdateResponseAccessSamlType = "okta"
	ZeroTrustIdentityProviderUpdateResponseAccessSamlTypeOnelogin   ZeroTrustIdentityProviderUpdateResponseAccessSamlType = "onelogin"
	ZeroTrustIdentityProviderUpdateResponseAccessSamlTypePingone    ZeroTrustIdentityProviderUpdateResponseAccessSamlType = "pingone"
	ZeroTrustIdentityProviderUpdateResponseAccessSamlTypeYandex     ZeroTrustIdentityProviderUpdateResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderUpdateResponseAccessSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessSamlScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessSamlScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessSamlScimConfig]
type zeroTrustIdentityProviderUpdateResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessSamlScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderUpdateResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderUpdateResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderUpdateResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderUpdateResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessYandexJSON       `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessYandexJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderUpdateResponseAccessYandex]
type zeroTrustIdentityProviderUpdateResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessYandexJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderUpdateResponseAccessYandex) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                        `json:"client_secret"`
	JSON         zeroTrustIdentityProviderUpdateResponseAccessYandexConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessYandexConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessYandexConfig]
type zeroTrustIdentityProviderUpdateResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessYandexConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessYandexType string

const (
	ZeroTrustIdentityProviderUpdateResponseAccessYandexTypeOnetimepin ZeroTrustIdentityProviderUpdateResponseAccessYandexType = "onetimepin"
	ZeroTrustIdentityProviderUpdateResponseAccessYandexTypeAzureAd    ZeroTrustIdentityProviderUpdateResponseAccessYandexType = "azureAD"
	ZeroTrustIdentityProviderUpdateResponseAccessYandexTypeSaml       ZeroTrustIdentityProviderUpdateResponseAccessYandexType = "saml"
	ZeroTrustIdentityProviderUpdateResponseAccessYandexTypeCentrify   ZeroTrustIdentityProviderUpdateResponseAccessYandexType = "centrify"
	ZeroTrustIdentityProviderUpdateResponseAccessYandexTypeFacebook   ZeroTrustIdentityProviderUpdateResponseAccessYandexType = "facebook"
	ZeroTrustIdentityProviderUpdateResponseAccessYandexTypeGitHub     ZeroTrustIdentityProviderUpdateResponseAccessYandexType = "github"
	ZeroTrustIdentityProviderUpdateResponseAccessYandexTypeGoogleApps ZeroTrustIdentityProviderUpdateResponseAccessYandexType = "google-apps"
	ZeroTrustIdentityProviderUpdateResponseAccessYandexTypeGoogle     ZeroTrustIdentityProviderUpdateResponseAccessYandexType = "google"
	ZeroTrustIdentityProviderUpdateResponseAccessYandexTypeLinkedin   ZeroTrustIdentityProviderUpdateResponseAccessYandexType = "linkedin"
	ZeroTrustIdentityProviderUpdateResponseAccessYandexTypeOidc       ZeroTrustIdentityProviderUpdateResponseAccessYandexType = "oidc"
	ZeroTrustIdentityProviderUpdateResponseAccessYandexTypeOkta       ZeroTrustIdentityProviderUpdateResponseAccessYandexType = "okta"
	ZeroTrustIdentityProviderUpdateResponseAccessYandexTypeOnelogin   ZeroTrustIdentityProviderUpdateResponseAccessYandexType = "onelogin"
	ZeroTrustIdentityProviderUpdateResponseAccessYandexTypePingone    ZeroTrustIdentityProviderUpdateResponseAccessYandexType = "pingone"
	ZeroTrustIdentityProviderUpdateResponseAccessYandexTypeYandex     ZeroTrustIdentityProviderUpdateResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderUpdateResponseAccessYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessYandexScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessYandexScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessYandexScimConfig]
type zeroTrustIdentityProviderUpdateResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessYandexScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderUpdateResponseAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderUpdateResponseAccessOnetimepinJSON       `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessOnetimepinJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessOnetimepin]
type zeroTrustIdentityProviderUpdateResponseAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessOnetimepinJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderUpdateResponseAccessOnetimepin) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType string

const (
	ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinTypeOnetimepin ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType = "onetimepin"
	ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinTypeAzureAd    ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType = "azureAD"
	ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinTypeSaml       ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType = "saml"
	ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinTypeCentrify   ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType = "centrify"
	ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinTypeFacebook   ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType = "facebook"
	ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinTypeGitHub     ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType = "github"
	ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinTypeGoogleApps ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType = "google-apps"
	ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinTypeGoogle     ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType = "google"
	ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinTypeLinkedin   ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType = "linkedin"
	ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinTypeOidc       ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType = "oidc"
	ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinTypeOkta       ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType = "okta"
	ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinTypeOnelogin   ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType = "onelogin"
	ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinTypePingone    ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType = "pingone"
	ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinTypeYandex     ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderUpdateResponseAccessOnetimepinScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderUpdateResponseAccessOnetimepinScimConfigJSON contains
// the JSON metadata for the struct
// [ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinScimConfig]
type zeroTrustIdentityProviderUpdateResponseAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderUpdateResponseAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderUpdateResponseAccessOnetimepinScimConfigJSON) RawJSON() string {
	return r.raw
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
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustIdentityProviderListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderListResponseAccessAzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderListResponseAccessCentrify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderListResponseAccessFacebook{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderListResponseAccessGitHub{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderListResponseAccessGoogle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderListResponseAccessGoogleApps{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderListResponseAccessLinkedin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderListResponseAccessOidc{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderListResponseAccessOkta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderListResponseAccessOnelogin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderListResponseAccessPingone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderListResponseAccessSaml{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderListResponseAccessYandex{}),
		},
	)
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

func (r zeroTrustIdentityProviderListResponseAccessAzureAdJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessAzureAdConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessAzureAdScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessCentrifyJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessCentrifyConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessCentrifyScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessFacebookJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessFacebookConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessFacebookScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessGitHubJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessGitHubConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessGitHubScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessGoogleJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessGoogleConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessGoogleScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessGoogleAppsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessGoogleAppsConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessGoogleAppsScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessLinkedinJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessLinkedinConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessLinkedinScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessOidcJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessOidcConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessOidcScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessOktaJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessOktaConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessOktaScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessOneloginJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessOneloginConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessOneloginScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessPingoneJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessPingoneConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessPingoneScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessSamlJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessSamlConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessSamlConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessSamlScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessYandexJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessYandexConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseAccessYandexScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ZeroTrustIdentityProviderGetResponseAccessAzureAd],
// [ZeroTrustIdentityProviderGetResponseAccessCentrify],
// [ZeroTrustIdentityProviderGetResponseAccessFacebook],
// [ZeroTrustIdentityProviderGetResponseAccessGitHub],
// [ZeroTrustIdentityProviderGetResponseAccessGoogle],
// [ZeroTrustIdentityProviderGetResponseAccessGoogleApps],
// [ZeroTrustIdentityProviderGetResponseAccessLinkedin],
// [ZeroTrustIdentityProviderGetResponseAccessOidc],
// [ZeroTrustIdentityProviderGetResponseAccessOkta],
// [ZeroTrustIdentityProviderGetResponseAccessOnelogin],
// [ZeroTrustIdentityProviderGetResponseAccessPingone],
// [ZeroTrustIdentityProviderGetResponseAccessSaml],
// [ZeroTrustIdentityProviderGetResponseAccessYandex] or
// [ZeroTrustIdentityProviderGetResponseAccessOnetimepin].
type ZeroTrustIdentityProviderGetResponse interface {
	implementsZeroTrustIdentityProviderGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustIdentityProviderGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderGetResponseAccessAzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderGetResponseAccessCentrify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderGetResponseAccessFacebook{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderGetResponseAccessGitHub{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderGetResponseAccessGoogle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderGetResponseAccessGoogleApps{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderGetResponseAccessLinkedin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderGetResponseAccessOidc{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderGetResponseAccessOkta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderGetResponseAccessOnelogin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderGetResponseAccessPingone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderGetResponseAccessSaml{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderGetResponseAccessYandex{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProviderGetResponseAccessOnetimepin{}),
		},
	)
}

type ZeroTrustIdentityProviderGetResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderGetResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderGetResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderGetResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderGetResponseAccessAzureAdJSON       `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessAzureAdJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderGetResponseAccessAzureAd]
type zeroTrustIdentityProviderGetResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessAzureAdJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderGetResponseAccessAzureAd) implementsZeroTrustIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessAzureAdConfig struct {
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
	JSON          zeroTrustIdentityProviderGetResponseAccessAzureAdConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessAzureAdConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessAzureAdConfig]
type zeroTrustIdentityProviderGetResponseAccessAzureAdConfigJSON struct {
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

func (r *ZeroTrustIdentityProviderGetResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessAzureAdConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessAzureAdType string

const (
	ZeroTrustIdentityProviderGetResponseAccessAzureAdTypeOnetimepin ZeroTrustIdentityProviderGetResponseAccessAzureAdType = "onetimepin"
	ZeroTrustIdentityProviderGetResponseAccessAzureAdTypeAzureAd    ZeroTrustIdentityProviderGetResponseAccessAzureAdType = "azureAD"
	ZeroTrustIdentityProviderGetResponseAccessAzureAdTypeSaml       ZeroTrustIdentityProviderGetResponseAccessAzureAdType = "saml"
	ZeroTrustIdentityProviderGetResponseAccessAzureAdTypeCentrify   ZeroTrustIdentityProviderGetResponseAccessAzureAdType = "centrify"
	ZeroTrustIdentityProviderGetResponseAccessAzureAdTypeFacebook   ZeroTrustIdentityProviderGetResponseAccessAzureAdType = "facebook"
	ZeroTrustIdentityProviderGetResponseAccessAzureAdTypeGitHub     ZeroTrustIdentityProviderGetResponseAccessAzureAdType = "github"
	ZeroTrustIdentityProviderGetResponseAccessAzureAdTypeGoogleApps ZeroTrustIdentityProviderGetResponseAccessAzureAdType = "google-apps"
	ZeroTrustIdentityProviderGetResponseAccessAzureAdTypeGoogle     ZeroTrustIdentityProviderGetResponseAccessAzureAdType = "google"
	ZeroTrustIdentityProviderGetResponseAccessAzureAdTypeLinkedin   ZeroTrustIdentityProviderGetResponseAccessAzureAdType = "linkedin"
	ZeroTrustIdentityProviderGetResponseAccessAzureAdTypeOidc       ZeroTrustIdentityProviderGetResponseAccessAzureAdType = "oidc"
	ZeroTrustIdentityProviderGetResponseAccessAzureAdTypeOkta       ZeroTrustIdentityProviderGetResponseAccessAzureAdType = "okta"
	ZeroTrustIdentityProviderGetResponseAccessAzureAdTypeOnelogin   ZeroTrustIdentityProviderGetResponseAccessAzureAdType = "onelogin"
	ZeroTrustIdentityProviderGetResponseAccessAzureAdTypePingone    ZeroTrustIdentityProviderGetResponseAccessAzureAdType = "pingone"
	ZeroTrustIdentityProviderGetResponseAccessAzureAdTypeYandex     ZeroTrustIdentityProviderGetResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderGetResponseAccessAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderGetResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessAzureAdScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessAzureAdScimConfig]
type zeroTrustIdentityProviderGetResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessAzureAdScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderGetResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderGetResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderGetResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderGetResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderGetResponseAccessCentrifyJSON       `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessCentrifyJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderGetResponseAccessCentrify]
type zeroTrustIdentityProviderGetResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessCentrifyJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderGetResponseAccessCentrify) implementsZeroTrustIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessCentrifyConfig struct {
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
	JSON           zeroTrustIdentityProviderGetResponseAccessCentrifyConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessCentrifyConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessCentrifyConfig]
type zeroTrustIdentityProviderGetResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessCentrifyConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessCentrifyType string

const (
	ZeroTrustIdentityProviderGetResponseAccessCentrifyTypeOnetimepin ZeroTrustIdentityProviderGetResponseAccessCentrifyType = "onetimepin"
	ZeroTrustIdentityProviderGetResponseAccessCentrifyTypeAzureAd    ZeroTrustIdentityProviderGetResponseAccessCentrifyType = "azureAD"
	ZeroTrustIdentityProviderGetResponseAccessCentrifyTypeSaml       ZeroTrustIdentityProviderGetResponseAccessCentrifyType = "saml"
	ZeroTrustIdentityProviderGetResponseAccessCentrifyTypeCentrify   ZeroTrustIdentityProviderGetResponseAccessCentrifyType = "centrify"
	ZeroTrustIdentityProviderGetResponseAccessCentrifyTypeFacebook   ZeroTrustIdentityProviderGetResponseAccessCentrifyType = "facebook"
	ZeroTrustIdentityProviderGetResponseAccessCentrifyTypeGitHub     ZeroTrustIdentityProviderGetResponseAccessCentrifyType = "github"
	ZeroTrustIdentityProviderGetResponseAccessCentrifyTypeGoogleApps ZeroTrustIdentityProviderGetResponseAccessCentrifyType = "google-apps"
	ZeroTrustIdentityProviderGetResponseAccessCentrifyTypeGoogle     ZeroTrustIdentityProviderGetResponseAccessCentrifyType = "google"
	ZeroTrustIdentityProviderGetResponseAccessCentrifyTypeLinkedin   ZeroTrustIdentityProviderGetResponseAccessCentrifyType = "linkedin"
	ZeroTrustIdentityProviderGetResponseAccessCentrifyTypeOidc       ZeroTrustIdentityProviderGetResponseAccessCentrifyType = "oidc"
	ZeroTrustIdentityProviderGetResponseAccessCentrifyTypeOkta       ZeroTrustIdentityProviderGetResponseAccessCentrifyType = "okta"
	ZeroTrustIdentityProviderGetResponseAccessCentrifyTypeOnelogin   ZeroTrustIdentityProviderGetResponseAccessCentrifyType = "onelogin"
	ZeroTrustIdentityProviderGetResponseAccessCentrifyTypePingone    ZeroTrustIdentityProviderGetResponseAccessCentrifyType = "pingone"
	ZeroTrustIdentityProviderGetResponseAccessCentrifyTypeYandex     ZeroTrustIdentityProviderGetResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderGetResponseAccessCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderGetResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessCentrifyScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessCentrifyScimConfig]
type zeroTrustIdentityProviderGetResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessCentrifyScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderGetResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderGetResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderGetResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderGetResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderGetResponseAccessFacebookJSON       `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessFacebookJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderGetResponseAccessFacebook]
type zeroTrustIdentityProviderGetResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessFacebookJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderGetResponseAccessFacebook) implementsZeroTrustIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                       `json:"client_secret"`
	JSON         zeroTrustIdentityProviderGetResponseAccessFacebookConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessFacebookConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessFacebookConfig]
type zeroTrustIdentityProviderGetResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessFacebookConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessFacebookType string

const (
	ZeroTrustIdentityProviderGetResponseAccessFacebookTypeOnetimepin ZeroTrustIdentityProviderGetResponseAccessFacebookType = "onetimepin"
	ZeroTrustIdentityProviderGetResponseAccessFacebookTypeAzureAd    ZeroTrustIdentityProviderGetResponseAccessFacebookType = "azureAD"
	ZeroTrustIdentityProviderGetResponseAccessFacebookTypeSaml       ZeroTrustIdentityProviderGetResponseAccessFacebookType = "saml"
	ZeroTrustIdentityProviderGetResponseAccessFacebookTypeCentrify   ZeroTrustIdentityProviderGetResponseAccessFacebookType = "centrify"
	ZeroTrustIdentityProviderGetResponseAccessFacebookTypeFacebook   ZeroTrustIdentityProviderGetResponseAccessFacebookType = "facebook"
	ZeroTrustIdentityProviderGetResponseAccessFacebookTypeGitHub     ZeroTrustIdentityProviderGetResponseAccessFacebookType = "github"
	ZeroTrustIdentityProviderGetResponseAccessFacebookTypeGoogleApps ZeroTrustIdentityProviderGetResponseAccessFacebookType = "google-apps"
	ZeroTrustIdentityProviderGetResponseAccessFacebookTypeGoogle     ZeroTrustIdentityProviderGetResponseAccessFacebookType = "google"
	ZeroTrustIdentityProviderGetResponseAccessFacebookTypeLinkedin   ZeroTrustIdentityProviderGetResponseAccessFacebookType = "linkedin"
	ZeroTrustIdentityProviderGetResponseAccessFacebookTypeOidc       ZeroTrustIdentityProviderGetResponseAccessFacebookType = "oidc"
	ZeroTrustIdentityProviderGetResponseAccessFacebookTypeOkta       ZeroTrustIdentityProviderGetResponseAccessFacebookType = "okta"
	ZeroTrustIdentityProviderGetResponseAccessFacebookTypeOnelogin   ZeroTrustIdentityProviderGetResponseAccessFacebookType = "onelogin"
	ZeroTrustIdentityProviderGetResponseAccessFacebookTypePingone    ZeroTrustIdentityProviderGetResponseAccessFacebookType = "pingone"
	ZeroTrustIdentityProviderGetResponseAccessFacebookTypeYandex     ZeroTrustIdentityProviderGetResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderGetResponseAccessFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderGetResponseAccessFacebookScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessFacebookScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessFacebookScimConfig]
type zeroTrustIdentityProviderGetResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessFacebookScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderGetResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderGetResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderGetResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderGetResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderGetResponseAccessGitHubJSON       `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessGitHubJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderGetResponseAccessGitHub]
type zeroTrustIdentityProviderGetResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessGitHubJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderGetResponseAccessGitHub) implementsZeroTrustIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                     `json:"client_secret"`
	JSON         zeroTrustIdentityProviderGetResponseAccessGitHubConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessGitHubConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderGetResponseAccessGitHubConfig]
type zeroTrustIdentityProviderGetResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessGitHubConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessGitHubType string

const (
	ZeroTrustIdentityProviderGetResponseAccessGitHubTypeOnetimepin ZeroTrustIdentityProviderGetResponseAccessGitHubType = "onetimepin"
	ZeroTrustIdentityProviderGetResponseAccessGitHubTypeAzureAd    ZeroTrustIdentityProviderGetResponseAccessGitHubType = "azureAD"
	ZeroTrustIdentityProviderGetResponseAccessGitHubTypeSaml       ZeroTrustIdentityProviderGetResponseAccessGitHubType = "saml"
	ZeroTrustIdentityProviderGetResponseAccessGitHubTypeCentrify   ZeroTrustIdentityProviderGetResponseAccessGitHubType = "centrify"
	ZeroTrustIdentityProviderGetResponseAccessGitHubTypeFacebook   ZeroTrustIdentityProviderGetResponseAccessGitHubType = "facebook"
	ZeroTrustIdentityProviderGetResponseAccessGitHubTypeGitHub     ZeroTrustIdentityProviderGetResponseAccessGitHubType = "github"
	ZeroTrustIdentityProviderGetResponseAccessGitHubTypeGoogleApps ZeroTrustIdentityProviderGetResponseAccessGitHubType = "google-apps"
	ZeroTrustIdentityProviderGetResponseAccessGitHubTypeGoogle     ZeroTrustIdentityProviderGetResponseAccessGitHubType = "google"
	ZeroTrustIdentityProviderGetResponseAccessGitHubTypeLinkedin   ZeroTrustIdentityProviderGetResponseAccessGitHubType = "linkedin"
	ZeroTrustIdentityProviderGetResponseAccessGitHubTypeOidc       ZeroTrustIdentityProviderGetResponseAccessGitHubType = "oidc"
	ZeroTrustIdentityProviderGetResponseAccessGitHubTypeOkta       ZeroTrustIdentityProviderGetResponseAccessGitHubType = "okta"
	ZeroTrustIdentityProviderGetResponseAccessGitHubTypeOnelogin   ZeroTrustIdentityProviderGetResponseAccessGitHubType = "onelogin"
	ZeroTrustIdentityProviderGetResponseAccessGitHubTypePingone    ZeroTrustIdentityProviderGetResponseAccessGitHubType = "pingone"
	ZeroTrustIdentityProviderGetResponseAccessGitHubTypeYandex     ZeroTrustIdentityProviderGetResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderGetResponseAccessGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderGetResponseAccessGitHubScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessGitHubScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessGitHubScimConfig]
type zeroTrustIdentityProviderGetResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessGitHubScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderGetResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderGetResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderGetResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderGetResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderGetResponseAccessGoogleJSON       `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessGoogleJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderGetResponseAccessGoogle]
type zeroTrustIdentityProviderGetResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessGoogleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderGetResponseAccessGoogle) implementsZeroTrustIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                     `json:"email_claim_name"`
	JSON           zeroTrustIdentityProviderGetResponseAccessGoogleConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessGoogleConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderGetResponseAccessGoogleConfig]
type zeroTrustIdentityProviderGetResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessGoogleConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessGoogleType string

const (
	ZeroTrustIdentityProviderGetResponseAccessGoogleTypeOnetimepin ZeroTrustIdentityProviderGetResponseAccessGoogleType = "onetimepin"
	ZeroTrustIdentityProviderGetResponseAccessGoogleTypeAzureAd    ZeroTrustIdentityProviderGetResponseAccessGoogleType = "azureAD"
	ZeroTrustIdentityProviderGetResponseAccessGoogleTypeSaml       ZeroTrustIdentityProviderGetResponseAccessGoogleType = "saml"
	ZeroTrustIdentityProviderGetResponseAccessGoogleTypeCentrify   ZeroTrustIdentityProviderGetResponseAccessGoogleType = "centrify"
	ZeroTrustIdentityProviderGetResponseAccessGoogleTypeFacebook   ZeroTrustIdentityProviderGetResponseAccessGoogleType = "facebook"
	ZeroTrustIdentityProviderGetResponseAccessGoogleTypeGitHub     ZeroTrustIdentityProviderGetResponseAccessGoogleType = "github"
	ZeroTrustIdentityProviderGetResponseAccessGoogleTypeGoogleApps ZeroTrustIdentityProviderGetResponseAccessGoogleType = "google-apps"
	ZeroTrustIdentityProviderGetResponseAccessGoogleTypeGoogle     ZeroTrustIdentityProviderGetResponseAccessGoogleType = "google"
	ZeroTrustIdentityProviderGetResponseAccessGoogleTypeLinkedin   ZeroTrustIdentityProviderGetResponseAccessGoogleType = "linkedin"
	ZeroTrustIdentityProviderGetResponseAccessGoogleTypeOidc       ZeroTrustIdentityProviderGetResponseAccessGoogleType = "oidc"
	ZeroTrustIdentityProviderGetResponseAccessGoogleTypeOkta       ZeroTrustIdentityProviderGetResponseAccessGoogleType = "okta"
	ZeroTrustIdentityProviderGetResponseAccessGoogleTypeOnelogin   ZeroTrustIdentityProviderGetResponseAccessGoogleType = "onelogin"
	ZeroTrustIdentityProviderGetResponseAccessGoogleTypePingone    ZeroTrustIdentityProviderGetResponseAccessGoogleType = "pingone"
	ZeroTrustIdentityProviderGetResponseAccessGoogleTypeYandex     ZeroTrustIdentityProviderGetResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderGetResponseAccessGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderGetResponseAccessGoogleScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessGoogleScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessGoogleScimConfig]
type zeroTrustIdentityProviderGetResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessGoogleScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderGetResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderGetResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderGetResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderGetResponseAccessGoogleAppsJSON       `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessGoogleAppsJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderGetResponseAccessGoogleApps]
type zeroTrustIdentityProviderGetResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessGoogleAppsJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderGetResponseAccessGoogleApps) implementsZeroTrustIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessGoogleAppsConfig struct {
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
	JSON           zeroTrustIdentityProviderGetResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessGoogleAppsConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessGoogleAppsConfig]
type zeroTrustIdentityProviderGetResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessGoogleAppsConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType string

const (
	ZeroTrustIdentityProviderGetResponseAccessGoogleAppsTypeOnetimepin ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType = "onetimepin"
	ZeroTrustIdentityProviderGetResponseAccessGoogleAppsTypeAzureAd    ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType = "azureAD"
	ZeroTrustIdentityProviderGetResponseAccessGoogleAppsTypeSaml       ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType = "saml"
	ZeroTrustIdentityProviderGetResponseAccessGoogleAppsTypeCentrify   ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType = "centrify"
	ZeroTrustIdentityProviderGetResponseAccessGoogleAppsTypeFacebook   ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType = "facebook"
	ZeroTrustIdentityProviderGetResponseAccessGoogleAppsTypeGitHub     ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType = "github"
	ZeroTrustIdentityProviderGetResponseAccessGoogleAppsTypeGoogleApps ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType = "google-apps"
	ZeroTrustIdentityProviderGetResponseAccessGoogleAppsTypeGoogle     ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType = "google"
	ZeroTrustIdentityProviderGetResponseAccessGoogleAppsTypeLinkedin   ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType = "linkedin"
	ZeroTrustIdentityProviderGetResponseAccessGoogleAppsTypeOidc       ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType = "oidc"
	ZeroTrustIdentityProviderGetResponseAccessGoogleAppsTypeOkta       ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType = "okta"
	ZeroTrustIdentityProviderGetResponseAccessGoogleAppsTypeOnelogin   ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType = "onelogin"
	ZeroTrustIdentityProviderGetResponseAccessGoogleAppsTypePingone    ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType = "pingone"
	ZeroTrustIdentityProviderGetResponseAccessGoogleAppsTypeYandex     ZeroTrustIdentityProviderGetResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderGetResponseAccessGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderGetResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessGoogleAppsScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessGoogleAppsScimConfig]
type zeroTrustIdentityProviderGetResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessGoogleAppsScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderGetResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderGetResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderGetResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderGetResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderGetResponseAccessLinkedinJSON       `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessLinkedinJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderGetResponseAccessLinkedin]
type zeroTrustIdentityProviderGetResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessLinkedinJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderGetResponseAccessLinkedin) implementsZeroTrustIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                       `json:"client_secret"`
	JSON         zeroTrustIdentityProviderGetResponseAccessLinkedinConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessLinkedinConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessLinkedinConfig]
type zeroTrustIdentityProviderGetResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessLinkedinConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessLinkedinType string

const (
	ZeroTrustIdentityProviderGetResponseAccessLinkedinTypeOnetimepin ZeroTrustIdentityProviderGetResponseAccessLinkedinType = "onetimepin"
	ZeroTrustIdentityProviderGetResponseAccessLinkedinTypeAzureAd    ZeroTrustIdentityProviderGetResponseAccessLinkedinType = "azureAD"
	ZeroTrustIdentityProviderGetResponseAccessLinkedinTypeSaml       ZeroTrustIdentityProviderGetResponseAccessLinkedinType = "saml"
	ZeroTrustIdentityProviderGetResponseAccessLinkedinTypeCentrify   ZeroTrustIdentityProviderGetResponseAccessLinkedinType = "centrify"
	ZeroTrustIdentityProviderGetResponseAccessLinkedinTypeFacebook   ZeroTrustIdentityProviderGetResponseAccessLinkedinType = "facebook"
	ZeroTrustIdentityProviderGetResponseAccessLinkedinTypeGitHub     ZeroTrustIdentityProviderGetResponseAccessLinkedinType = "github"
	ZeroTrustIdentityProviderGetResponseAccessLinkedinTypeGoogleApps ZeroTrustIdentityProviderGetResponseAccessLinkedinType = "google-apps"
	ZeroTrustIdentityProviderGetResponseAccessLinkedinTypeGoogle     ZeroTrustIdentityProviderGetResponseAccessLinkedinType = "google"
	ZeroTrustIdentityProviderGetResponseAccessLinkedinTypeLinkedin   ZeroTrustIdentityProviderGetResponseAccessLinkedinType = "linkedin"
	ZeroTrustIdentityProviderGetResponseAccessLinkedinTypeOidc       ZeroTrustIdentityProviderGetResponseAccessLinkedinType = "oidc"
	ZeroTrustIdentityProviderGetResponseAccessLinkedinTypeOkta       ZeroTrustIdentityProviderGetResponseAccessLinkedinType = "okta"
	ZeroTrustIdentityProviderGetResponseAccessLinkedinTypeOnelogin   ZeroTrustIdentityProviderGetResponseAccessLinkedinType = "onelogin"
	ZeroTrustIdentityProviderGetResponseAccessLinkedinTypePingone    ZeroTrustIdentityProviderGetResponseAccessLinkedinType = "pingone"
	ZeroTrustIdentityProviderGetResponseAccessLinkedinTypeYandex     ZeroTrustIdentityProviderGetResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderGetResponseAccessLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderGetResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessLinkedinScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessLinkedinScimConfig]
type zeroTrustIdentityProviderGetResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessLinkedinScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderGetResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderGetResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderGetResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderGetResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderGetResponseAccessOidcJSON       `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessOidcJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderGetResponseAccessOidc]
type zeroTrustIdentityProviderGetResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessOidcJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderGetResponseAccessOidc) implementsZeroTrustIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessOidcConfig struct {
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
	JSON     zeroTrustIdentityProviderGetResponseAccessOidcConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessOidcConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderGetResponseAccessOidcConfig]
type zeroTrustIdentityProviderGetResponseAccessOidcConfigJSON struct {
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

func (r *ZeroTrustIdentityProviderGetResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessOidcConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessOidcType string

const (
	ZeroTrustIdentityProviderGetResponseAccessOidcTypeOnetimepin ZeroTrustIdentityProviderGetResponseAccessOidcType = "onetimepin"
	ZeroTrustIdentityProviderGetResponseAccessOidcTypeAzureAd    ZeroTrustIdentityProviderGetResponseAccessOidcType = "azureAD"
	ZeroTrustIdentityProviderGetResponseAccessOidcTypeSaml       ZeroTrustIdentityProviderGetResponseAccessOidcType = "saml"
	ZeroTrustIdentityProviderGetResponseAccessOidcTypeCentrify   ZeroTrustIdentityProviderGetResponseAccessOidcType = "centrify"
	ZeroTrustIdentityProviderGetResponseAccessOidcTypeFacebook   ZeroTrustIdentityProviderGetResponseAccessOidcType = "facebook"
	ZeroTrustIdentityProviderGetResponseAccessOidcTypeGitHub     ZeroTrustIdentityProviderGetResponseAccessOidcType = "github"
	ZeroTrustIdentityProviderGetResponseAccessOidcTypeGoogleApps ZeroTrustIdentityProviderGetResponseAccessOidcType = "google-apps"
	ZeroTrustIdentityProviderGetResponseAccessOidcTypeGoogle     ZeroTrustIdentityProviderGetResponseAccessOidcType = "google"
	ZeroTrustIdentityProviderGetResponseAccessOidcTypeLinkedin   ZeroTrustIdentityProviderGetResponseAccessOidcType = "linkedin"
	ZeroTrustIdentityProviderGetResponseAccessOidcTypeOidc       ZeroTrustIdentityProviderGetResponseAccessOidcType = "oidc"
	ZeroTrustIdentityProviderGetResponseAccessOidcTypeOkta       ZeroTrustIdentityProviderGetResponseAccessOidcType = "okta"
	ZeroTrustIdentityProviderGetResponseAccessOidcTypeOnelogin   ZeroTrustIdentityProviderGetResponseAccessOidcType = "onelogin"
	ZeroTrustIdentityProviderGetResponseAccessOidcTypePingone    ZeroTrustIdentityProviderGetResponseAccessOidcType = "pingone"
	ZeroTrustIdentityProviderGetResponseAccessOidcTypeYandex     ZeroTrustIdentityProviderGetResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderGetResponseAccessOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderGetResponseAccessOidcScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessOidcScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessOidcScimConfig]
type zeroTrustIdentityProviderGetResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessOidcScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderGetResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderGetResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderGetResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderGetResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderGetResponseAccessOktaJSON       `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessOktaJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderGetResponseAccessOkta]
type zeroTrustIdentityProviderGetResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessOktaJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderGetResponseAccessOkta) implementsZeroTrustIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessOktaConfig struct {
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
	JSON        zeroTrustIdentityProviderGetResponseAccessOktaConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessOktaConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderGetResponseAccessOktaConfig]
type zeroTrustIdentityProviderGetResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessOktaConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessOktaType string

const (
	ZeroTrustIdentityProviderGetResponseAccessOktaTypeOnetimepin ZeroTrustIdentityProviderGetResponseAccessOktaType = "onetimepin"
	ZeroTrustIdentityProviderGetResponseAccessOktaTypeAzureAd    ZeroTrustIdentityProviderGetResponseAccessOktaType = "azureAD"
	ZeroTrustIdentityProviderGetResponseAccessOktaTypeSaml       ZeroTrustIdentityProviderGetResponseAccessOktaType = "saml"
	ZeroTrustIdentityProviderGetResponseAccessOktaTypeCentrify   ZeroTrustIdentityProviderGetResponseAccessOktaType = "centrify"
	ZeroTrustIdentityProviderGetResponseAccessOktaTypeFacebook   ZeroTrustIdentityProviderGetResponseAccessOktaType = "facebook"
	ZeroTrustIdentityProviderGetResponseAccessOktaTypeGitHub     ZeroTrustIdentityProviderGetResponseAccessOktaType = "github"
	ZeroTrustIdentityProviderGetResponseAccessOktaTypeGoogleApps ZeroTrustIdentityProviderGetResponseAccessOktaType = "google-apps"
	ZeroTrustIdentityProviderGetResponseAccessOktaTypeGoogle     ZeroTrustIdentityProviderGetResponseAccessOktaType = "google"
	ZeroTrustIdentityProviderGetResponseAccessOktaTypeLinkedin   ZeroTrustIdentityProviderGetResponseAccessOktaType = "linkedin"
	ZeroTrustIdentityProviderGetResponseAccessOktaTypeOidc       ZeroTrustIdentityProviderGetResponseAccessOktaType = "oidc"
	ZeroTrustIdentityProviderGetResponseAccessOktaTypeOkta       ZeroTrustIdentityProviderGetResponseAccessOktaType = "okta"
	ZeroTrustIdentityProviderGetResponseAccessOktaTypeOnelogin   ZeroTrustIdentityProviderGetResponseAccessOktaType = "onelogin"
	ZeroTrustIdentityProviderGetResponseAccessOktaTypePingone    ZeroTrustIdentityProviderGetResponseAccessOktaType = "pingone"
	ZeroTrustIdentityProviderGetResponseAccessOktaTypeYandex     ZeroTrustIdentityProviderGetResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderGetResponseAccessOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderGetResponseAccessOktaScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessOktaScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessOktaScimConfig]
type zeroTrustIdentityProviderGetResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessOktaScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderGetResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderGetResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderGetResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderGetResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderGetResponseAccessOneloginJSON       `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessOneloginJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderGetResponseAccessOnelogin]
type zeroTrustIdentityProviderGetResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessOneloginJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderGetResponseAccessOnelogin) implementsZeroTrustIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessOneloginConfig struct {
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
	JSON            zeroTrustIdentityProviderGetResponseAccessOneloginConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessOneloginConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessOneloginConfig]
type zeroTrustIdentityProviderGetResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessOneloginConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessOneloginType string

const (
	ZeroTrustIdentityProviderGetResponseAccessOneloginTypeOnetimepin ZeroTrustIdentityProviderGetResponseAccessOneloginType = "onetimepin"
	ZeroTrustIdentityProviderGetResponseAccessOneloginTypeAzureAd    ZeroTrustIdentityProviderGetResponseAccessOneloginType = "azureAD"
	ZeroTrustIdentityProviderGetResponseAccessOneloginTypeSaml       ZeroTrustIdentityProviderGetResponseAccessOneloginType = "saml"
	ZeroTrustIdentityProviderGetResponseAccessOneloginTypeCentrify   ZeroTrustIdentityProviderGetResponseAccessOneloginType = "centrify"
	ZeroTrustIdentityProviderGetResponseAccessOneloginTypeFacebook   ZeroTrustIdentityProviderGetResponseAccessOneloginType = "facebook"
	ZeroTrustIdentityProviderGetResponseAccessOneloginTypeGitHub     ZeroTrustIdentityProviderGetResponseAccessOneloginType = "github"
	ZeroTrustIdentityProviderGetResponseAccessOneloginTypeGoogleApps ZeroTrustIdentityProviderGetResponseAccessOneloginType = "google-apps"
	ZeroTrustIdentityProviderGetResponseAccessOneloginTypeGoogle     ZeroTrustIdentityProviderGetResponseAccessOneloginType = "google"
	ZeroTrustIdentityProviderGetResponseAccessOneloginTypeLinkedin   ZeroTrustIdentityProviderGetResponseAccessOneloginType = "linkedin"
	ZeroTrustIdentityProviderGetResponseAccessOneloginTypeOidc       ZeroTrustIdentityProviderGetResponseAccessOneloginType = "oidc"
	ZeroTrustIdentityProviderGetResponseAccessOneloginTypeOkta       ZeroTrustIdentityProviderGetResponseAccessOneloginType = "okta"
	ZeroTrustIdentityProviderGetResponseAccessOneloginTypeOnelogin   ZeroTrustIdentityProviderGetResponseAccessOneloginType = "onelogin"
	ZeroTrustIdentityProviderGetResponseAccessOneloginTypePingone    ZeroTrustIdentityProviderGetResponseAccessOneloginType = "pingone"
	ZeroTrustIdentityProviderGetResponseAccessOneloginTypeYandex     ZeroTrustIdentityProviderGetResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderGetResponseAccessOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderGetResponseAccessOneloginScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessOneloginScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessOneloginScimConfig]
type zeroTrustIdentityProviderGetResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessOneloginScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderGetResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderGetResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderGetResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderGetResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderGetResponseAccessPingoneJSON       `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessPingoneJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderGetResponseAccessPingone]
type zeroTrustIdentityProviderGetResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessPingoneJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderGetResponseAccessPingone) implementsZeroTrustIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessPingoneConfig struct {
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
	JSON      zeroTrustIdentityProviderGetResponseAccessPingoneConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessPingoneConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessPingoneConfig]
type zeroTrustIdentityProviderGetResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessPingoneConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessPingoneType string

const (
	ZeroTrustIdentityProviderGetResponseAccessPingoneTypeOnetimepin ZeroTrustIdentityProviderGetResponseAccessPingoneType = "onetimepin"
	ZeroTrustIdentityProviderGetResponseAccessPingoneTypeAzureAd    ZeroTrustIdentityProviderGetResponseAccessPingoneType = "azureAD"
	ZeroTrustIdentityProviderGetResponseAccessPingoneTypeSaml       ZeroTrustIdentityProviderGetResponseAccessPingoneType = "saml"
	ZeroTrustIdentityProviderGetResponseAccessPingoneTypeCentrify   ZeroTrustIdentityProviderGetResponseAccessPingoneType = "centrify"
	ZeroTrustIdentityProviderGetResponseAccessPingoneTypeFacebook   ZeroTrustIdentityProviderGetResponseAccessPingoneType = "facebook"
	ZeroTrustIdentityProviderGetResponseAccessPingoneTypeGitHub     ZeroTrustIdentityProviderGetResponseAccessPingoneType = "github"
	ZeroTrustIdentityProviderGetResponseAccessPingoneTypeGoogleApps ZeroTrustIdentityProviderGetResponseAccessPingoneType = "google-apps"
	ZeroTrustIdentityProviderGetResponseAccessPingoneTypeGoogle     ZeroTrustIdentityProviderGetResponseAccessPingoneType = "google"
	ZeroTrustIdentityProviderGetResponseAccessPingoneTypeLinkedin   ZeroTrustIdentityProviderGetResponseAccessPingoneType = "linkedin"
	ZeroTrustIdentityProviderGetResponseAccessPingoneTypeOidc       ZeroTrustIdentityProviderGetResponseAccessPingoneType = "oidc"
	ZeroTrustIdentityProviderGetResponseAccessPingoneTypeOkta       ZeroTrustIdentityProviderGetResponseAccessPingoneType = "okta"
	ZeroTrustIdentityProviderGetResponseAccessPingoneTypeOnelogin   ZeroTrustIdentityProviderGetResponseAccessPingoneType = "onelogin"
	ZeroTrustIdentityProviderGetResponseAccessPingoneTypePingone    ZeroTrustIdentityProviderGetResponseAccessPingoneType = "pingone"
	ZeroTrustIdentityProviderGetResponseAccessPingoneTypeYandex     ZeroTrustIdentityProviderGetResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderGetResponseAccessPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderGetResponseAccessPingoneScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessPingoneScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessPingoneScimConfig]
type zeroTrustIdentityProviderGetResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessPingoneScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderGetResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderGetResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderGetResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderGetResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderGetResponseAccessSamlJSON       `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessSamlJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderGetResponseAccessSaml]
type zeroTrustIdentityProviderGetResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessSamlJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderGetResponseAccessSaml) implementsZeroTrustIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []ZeroTrustIdentityProviderGetResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                   `json:"sso_target_url"`
	JSON         zeroTrustIdentityProviderGetResponseAccessSamlConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessSamlConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderGetResponseAccessSamlConfig]
type zeroTrustIdentityProviderGetResponseAccessSamlConfigJSON struct {
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

func (r *ZeroTrustIdentityProviderGetResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessSamlConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderGetResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                  `json:"header_name"`
	JSON       zeroTrustIdentityProviderGetResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessSamlConfigHeaderAttributeJSON contains
// the JSON metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessSamlConfigHeaderAttribute]
type zeroTrustIdentityProviderGetResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessSamlConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessSamlType string

const (
	ZeroTrustIdentityProviderGetResponseAccessSamlTypeOnetimepin ZeroTrustIdentityProviderGetResponseAccessSamlType = "onetimepin"
	ZeroTrustIdentityProviderGetResponseAccessSamlTypeAzureAd    ZeroTrustIdentityProviderGetResponseAccessSamlType = "azureAD"
	ZeroTrustIdentityProviderGetResponseAccessSamlTypeSaml       ZeroTrustIdentityProviderGetResponseAccessSamlType = "saml"
	ZeroTrustIdentityProviderGetResponseAccessSamlTypeCentrify   ZeroTrustIdentityProviderGetResponseAccessSamlType = "centrify"
	ZeroTrustIdentityProviderGetResponseAccessSamlTypeFacebook   ZeroTrustIdentityProviderGetResponseAccessSamlType = "facebook"
	ZeroTrustIdentityProviderGetResponseAccessSamlTypeGitHub     ZeroTrustIdentityProviderGetResponseAccessSamlType = "github"
	ZeroTrustIdentityProviderGetResponseAccessSamlTypeGoogleApps ZeroTrustIdentityProviderGetResponseAccessSamlType = "google-apps"
	ZeroTrustIdentityProviderGetResponseAccessSamlTypeGoogle     ZeroTrustIdentityProviderGetResponseAccessSamlType = "google"
	ZeroTrustIdentityProviderGetResponseAccessSamlTypeLinkedin   ZeroTrustIdentityProviderGetResponseAccessSamlType = "linkedin"
	ZeroTrustIdentityProviderGetResponseAccessSamlTypeOidc       ZeroTrustIdentityProviderGetResponseAccessSamlType = "oidc"
	ZeroTrustIdentityProviderGetResponseAccessSamlTypeOkta       ZeroTrustIdentityProviderGetResponseAccessSamlType = "okta"
	ZeroTrustIdentityProviderGetResponseAccessSamlTypeOnelogin   ZeroTrustIdentityProviderGetResponseAccessSamlType = "onelogin"
	ZeroTrustIdentityProviderGetResponseAccessSamlTypePingone    ZeroTrustIdentityProviderGetResponseAccessSamlType = "pingone"
	ZeroTrustIdentityProviderGetResponseAccessSamlTypeYandex     ZeroTrustIdentityProviderGetResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderGetResponseAccessSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderGetResponseAccessSamlScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessSamlScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessSamlScimConfig]
type zeroTrustIdentityProviderGetResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessSamlScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderGetResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProviderGetResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderGetResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderGetResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderGetResponseAccessYandexJSON       `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessYandexJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProviderGetResponseAccessYandex]
type zeroTrustIdentityProviderGetResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessYandexJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderGetResponseAccessYandex) implementsZeroTrustIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                     `json:"client_secret"`
	JSON         zeroTrustIdentityProviderGetResponseAccessYandexConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessYandexConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderGetResponseAccessYandexConfig]
type zeroTrustIdentityProviderGetResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessYandexConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessYandexType string

const (
	ZeroTrustIdentityProviderGetResponseAccessYandexTypeOnetimepin ZeroTrustIdentityProviderGetResponseAccessYandexType = "onetimepin"
	ZeroTrustIdentityProviderGetResponseAccessYandexTypeAzureAd    ZeroTrustIdentityProviderGetResponseAccessYandexType = "azureAD"
	ZeroTrustIdentityProviderGetResponseAccessYandexTypeSaml       ZeroTrustIdentityProviderGetResponseAccessYandexType = "saml"
	ZeroTrustIdentityProviderGetResponseAccessYandexTypeCentrify   ZeroTrustIdentityProviderGetResponseAccessYandexType = "centrify"
	ZeroTrustIdentityProviderGetResponseAccessYandexTypeFacebook   ZeroTrustIdentityProviderGetResponseAccessYandexType = "facebook"
	ZeroTrustIdentityProviderGetResponseAccessYandexTypeGitHub     ZeroTrustIdentityProviderGetResponseAccessYandexType = "github"
	ZeroTrustIdentityProviderGetResponseAccessYandexTypeGoogleApps ZeroTrustIdentityProviderGetResponseAccessYandexType = "google-apps"
	ZeroTrustIdentityProviderGetResponseAccessYandexTypeGoogle     ZeroTrustIdentityProviderGetResponseAccessYandexType = "google"
	ZeroTrustIdentityProviderGetResponseAccessYandexTypeLinkedin   ZeroTrustIdentityProviderGetResponseAccessYandexType = "linkedin"
	ZeroTrustIdentityProviderGetResponseAccessYandexTypeOidc       ZeroTrustIdentityProviderGetResponseAccessYandexType = "oidc"
	ZeroTrustIdentityProviderGetResponseAccessYandexTypeOkta       ZeroTrustIdentityProviderGetResponseAccessYandexType = "okta"
	ZeroTrustIdentityProviderGetResponseAccessYandexTypeOnelogin   ZeroTrustIdentityProviderGetResponseAccessYandexType = "onelogin"
	ZeroTrustIdentityProviderGetResponseAccessYandexTypePingone    ZeroTrustIdentityProviderGetResponseAccessYandexType = "pingone"
	ZeroTrustIdentityProviderGetResponseAccessYandexTypeYandex     ZeroTrustIdentityProviderGetResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderGetResponseAccessYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderGetResponseAccessYandexScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessYandexScimConfigJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessYandexScimConfig]
type zeroTrustIdentityProviderGetResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessYandexScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProviderGetResponseAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProviderGetResponseAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProviderGetResponseAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProviderGetResponseAccessOnetimepinJSON       `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessOnetimepinJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProviderGetResponseAccessOnetimepin]
type zeroTrustIdentityProviderGetResponseAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessOnetimepinJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProviderGetResponseAccessOnetimepin) implementsZeroTrustIdentityProviderGetResponse() {
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProviderGetResponseAccessOnetimepinType string

const (
	ZeroTrustIdentityProviderGetResponseAccessOnetimepinTypeOnetimepin ZeroTrustIdentityProviderGetResponseAccessOnetimepinType = "onetimepin"
	ZeroTrustIdentityProviderGetResponseAccessOnetimepinTypeAzureAd    ZeroTrustIdentityProviderGetResponseAccessOnetimepinType = "azureAD"
	ZeroTrustIdentityProviderGetResponseAccessOnetimepinTypeSaml       ZeroTrustIdentityProviderGetResponseAccessOnetimepinType = "saml"
	ZeroTrustIdentityProviderGetResponseAccessOnetimepinTypeCentrify   ZeroTrustIdentityProviderGetResponseAccessOnetimepinType = "centrify"
	ZeroTrustIdentityProviderGetResponseAccessOnetimepinTypeFacebook   ZeroTrustIdentityProviderGetResponseAccessOnetimepinType = "facebook"
	ZeroTrustIdentityProviderGetResponseAccessOnetimepinTypeGitHub     ZeroTrustIdentityProviderGetResponseAccessOnetimepinType = "github"
	ZeroTrustIdentityProviderGetResponseAccessOnetimepinTypeGoogleApps ZeroTrustIdentityProviderGetResponseAccessOnetimepinType = "google-apps"
	ZeroTrustIdentityProviderGetResponseAccessOnetimepinTypeGoogle     ZeroTrustIdentityProviderGetResponseAccessOnetimepinType = "google"
	ZeroTrustIdentityProviderGetResponseAccessOnetimepinTypeLinkedin   ZeroTrustIdentityProviderGetResponseAccessOnetimepinType = "linkedin"
	ZeroTrustIdentityProviderGetResponseAccessOnetimepinTypeOidc       ZeroTrustIdentityProviderGetResponseAccessOnetimepinType = "oidc"
	ZeroTrustIdentityProviderGetResponseAccessOnetimepinTypeOkta       ZeroTrustIdentityProviderGetResponseAccessOnetimepinType = "okta"
	ZeroTrustIdentityProviderGetResponseAccessOnetimepinTypeOnelogin   ZeroTrustIdentityProviderGetResponseAccessOnetimepinType = "onelogin"
	ZeroTrustIdentityProviderGetResponseAccessOnetimepinTypePingone    ZeroTrustIdentityProviderGetResponseAccessOnetimepinType = "pingone"
	ZeroTrustIdentityProviderGetResponseAccessOnetimepinTypeYandex     ZeroTrustIdentityProviderGetResponseAccessOnetimepinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProviderGetResponseAccessOnetimepinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProviderGetResponseAccessOnetimepinScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProviderGetResponseAccessOnetimepinScimConfigJSON contains the
// JSON metadata for the struct
// [ZeroTrustIdentityProviderGetResponseAccessOnetimepinScimConfig]
type zeroTrustIdentityProviderGetResponseAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProviderGetResponseAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProviderGetResponseAccessOnetimepinScimConfigJSON) RawJSON() string {
	return r.raw
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
	Result   ZeroTrustIdentityProviderNewResponse                   `json:"result,required"`
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

func (r zeroTrustIdentityProviderNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   ZeroTrustIdentityProviderUpdateResponse                   `json:"result,required"`
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

func (r zeroTrustIdentityProviderUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
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
	Result   ZeroTrustIdentityProviderGetResponse                   `json:"result,required"`
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

func (r zeroTrustIdentityProviderGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r zeroTrustIdentityProviderGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustIdentityProviderGetResponseEnvelopeSuccess bool

const (
	ZeroTrustIdentityProviderGetResponseEnvelopeSuccessTrue ZeroTrustIdentityProviderGetResponseEnvelopeSuccess = true
)

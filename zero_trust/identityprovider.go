// File generated from our OpenAPI spec by Stainless.

package zero_trust

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

// IdentityProviderService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewIdentityProviderService] method
// instead.
type IdentityProviderService struct {
	Options []option.RequestOption
}

// NewIdentityProviderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIdentityProviderService(opts ...option.RequestOption) (r *IdentityProviderService) {
	r = &IdentityProviderService{}
	r.Options = opts
	return
}

// Adds a new identity provider to Access.
func (r *IdentityProviderService) New(ctx context.Context, params IdentityProviderNewParams, opts ...option.RequestOption) (res *IdentityProviderNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IdentityProviderNewResponseEnvelope
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
func (r *IdentityProviderService) Update(ctx context.Context, uuid string, params IdentityProviderUpdateParams, opts ...option.RequestOption) (res *IdentityProviderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IdentityProviderUpdateResponseEnvelope
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
func (r *IdentityProviderService) List(ctx context.Context, query IdentityProviderListParams, opts ...option.RequestOption) (res *[]IdentityProviderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IdentityProviderListResponseEnvelope
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
func (r *IdentityProviderService) Delete(ctx context.Context, uuid string, body IdentityProviderDeleteParams, opts ...option.RequestOption) (res *IdentityProviderDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IdentityProviderDeleteResponseEnvelope
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
func (r *IdentityProviderService) Get(ctx context.Context, uuid string, query IdentityProviderGetParams, opts ...option.RequestOption) (res *IdentityProviderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IdentityProviderGetResponseEnvelope
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

// Union satisfied by [zero_trust.IdentityProviderNewResponseAccessAzureAd],
// [zero_trust.IdentityProviderNewResponseAccessCentrify],
// [zero_trust.IdentityProviderNewResponseAccessFacebook],
// [zero_trust.IdentityProviderNewResponseAccessGitHub],
// [zero_trust.IdentityProviderNewResponseAccessGoogle],
// [zero_trust.IdentityProviderNewResponseAccessGoogleApps],
// [zero_trust.IdentityProviderNewResponseAccessLinkedin],
// [zero_trust.IdentityProviderNewResponseAccessOidc],
// [zero_trust.IdentityProviderNewResponseAccessOkta],
// [zero_trust.IdentityProviderNewResponseAccessOnelogin],
// [zero_trust.IdentityProviderNewResponseAccessPingone],
// [zero_trust.IdentityProviderNewResponseAccessSaml],
// [zero_trust.IdentityProviderNewResponseAccessYandex] or
// [zero_trust.IdentityProviderNewResponseAccessOnetimepin].
type IdentityProviderNewResponse interface {
	implementsZeroTrustIdentityProviderNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IdentityProviderNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessAzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessCentrify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessFacebook{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessGitHub{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessGoogle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessGoogleApps{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessLinkedin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessOidc{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessOkta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessOnelogin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessPingone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessSaml{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessYandex{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessOnetimepin{}),
		},
	)
}

type IdentityProviderNewResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderNewResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderNewResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       identityProviderNewResponseAccessAzureAdJSON       `json:"-"`
}

// identityProviderNewResponseAccessAzureAdJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseAccessAzureAd]
type identityProviderNewResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessAzureAdJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessAzureAd) implementsZeroTrustIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessAzureAdConfig struct {
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
	SupportGroups bool                                               `json:"support_groups"`
	JSON          identityProviderNewResponseAccessAzureAdConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessAzureAdConfigJSON contains the JSON metadata
// for the struct [IdentityProviderNewResponseAccessAzureAdConfig]
type identityProviderNewResponseAccessAzureAdConfigJSON struct {
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

func (r *IdentityProviderNewResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessAzureAdConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessAzureAdType string

const (
	IdentityProviderNewResponseAccessAzureAdTypeOnetimepin IdentityProviderNewResponseAccessAzureAdType = "onetimepin"
	IdentityProviderNewResponseAccessAzureAdTypeAzureAd    IdentityProviderNewResponseAccessAzureAdType = "azureAD"
	IdentityProviderNewResponseAccessAzureAdTypeSaml       IdentityProviderNewResponseAccessAzureAdType = "saml"
	IdentityProviderNewResponseAccessAzureAdTypeCentrify   IdentityProviderNewResponseAccessAzureAdType = "centrify"
	IdentityProviderNewResponseAccessAzureAdTypeFacebook   IdentityProviderNewResponseAccessAzureAdType = "facebook"
	IdentityProviderNewResponseAccessAzureAdTypeGitHub     IdentityProviderNewResponseAccessAzureAdType = "github"
	IdentityProviderNewResponseAccessAzureAdTypeGoogleApps IdentityProviderNewResponseAccessAzureAdType = "google-apps"
	IdentityProviderNewResponseAccessAzureAdTypeGoogle     IdentityProviderNewResponseAccessAzureAdType = "google"
	IdentityProviderNewResponseAccessAzureAdTypeLinkedin   IdentityProviderNewResponseAccessAzureAdType = "linkedin"
	IdentityProviderNewResponseAccessAzureAdTypeOidc       IdentityProviderNewResponseAccessAzureAdType = "oidc"
	IdentityProviderNewResponseAccessAzureAdTypeOkta       IdentityProviderNewResponseAccessAzureAdType = "okta"
	IdentityProviderNewResponseAccessAzureAdTypeOnelogin   IdentityProviderNewResponseAccessAzureAdType = "onelogin"
	IdentityProviderNewResponseAccessAzureAdTypePingone    IdentityProviderNewResponseAccessAzureAdType = "pingone"
	IdentityProviderNewResponseAccessAzureAdTypeYandex     IdentityProviderNewResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewResponseAccessAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderNewResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessAzureAdScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderNewResponseAccessAzureAdScimConfig]
type identityProviderNewResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessAzureAdScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderNewResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderNewResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       identityProviderNewResponseAccessCentrifyJSON       `json:"-"`
}

// identityProviderNewResponseAccessCentrifyJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseAccessCentrify]
type identityProviderNewResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessCentrifyJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessCentrify) implementsZeroTrustIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessCentrifyConfig struct {
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
	EmailClaimName string                                              `json:"email_claim_name"`
	JSON           identityProviderNewResponseAccessCentrifyConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessCentrifyConfigJSON contains the JSON metadata
// for the struct [IdentityProviderNewResponseAccessCentrifyConfig]
type identityProviderNewResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessCentrifyConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessCentrifyType string

const (
	IdentityProviderNewResponseAccessCentrifyTypeOnetimepin IdentityProviderNewResponseAccessCentrifyType = "onetimepin"
	IdentityProviderNewResponseAccessCentrifyTypeAzureAd    IdentityProviderNewResponseAccessCentrifyType = "azureAD"
	IdentityProviderNewResponseAccessCentrifyTypeSaml       IdentityProviderNewResponseAccessCentrifyType = "saml"
	IdentityProviderNewResponseAccessCentrifyTypeCentrify   IdentityProviderNewResponseAccessCentrifyType = "centrify"
	IdentityProviderNewResponseAccessCentrifyTypeFacebook   IdentityProviderNewResponseAccessCentrifyType = "facebook"
	IdentityProviderNewResponseAccessCentrifyTypeGitHub     IdentityProviderNewResponseAccessCentrifyType = "github"
	IdentityProviderNewResponseAccessCentrifyTypeGoogleApps IdentityProviderNewResponseAccessCentrifyType = "google-apps"
	IdentityProviderNewResponseAccessCentrifyTypeGoogle     IdentityProviderNewResponseAccessCentrifyType = "google"
	IdentityProviderNewResponseAccessCentrifyTypeLinkedin   IdentityProviderNewResponseAccessCentrifyType = "linkedin"
	IdentityProviderNewResponseAccessCentrifyTypeOidc       IdentityProviderNewResponseAccessCentrifyType = "oidc"
	IdentityProviderNewResponseAccessCentrifyTypeOkta       IdentityProviderNewResponseAccessCentrifyType = "okta"
	IdentityProviderNewResponseAccessCentrifyTypeOnelogin   IdentityProviderNewResponseAccessCentrifyType = "onelogin"
	IdentityProviderNewResponseAccessCentrifyTypePingone    IdentityProviderNewResponseAccessCentrifyType = "pingone"
	IdentityProviderNewResponseAccessCentrifyTypeYandex     IdentityProviderNewResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewResponseAccessCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderNewResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessCentrifyScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderNewResponseAccessCentrifyScimConfig]
type identityProviderNewResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessCentrifyScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderNewResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderNewResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       identityProviderNewResponseAccessFacebookJSON       `json:"-"`
}

// identityProviderNewResponseAccessFacebookJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseAccessFacebook]
type identityProviderNewResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessFacebookJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessFacebook) implementsZeroTrustIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                              `json:"client_secret"`
	JSON         identityProviderNewResponseAccessFacebookConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessFacebookConfigJSON contains the JSON metadata
// for the struct [IdentityProviderNewResponseAccessFacebookConfig]
type identityProviderNewResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessFacebookConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessFacebookType string

const (
	IdentityProviderNewResponseAccessFacebookTypeOnetimepin IdentityProviderNewResponseAccessFacebookType = "onetimepin"
	IdentityProviderNewResponseAccessFacebookTypeAzureAd    IdentityProviderNewResponseAccessFacebookType = "azureAD"
	IdentityProviderNewResponseAccessFacebookTypeSaml       IdentityProviderNewResponseAccessFacebookType = "saml"
	IdentityProviderNewResponseAccessFacebookTypeCentrify   IdentityProviderNewResponseAccessFacebookType = "centrify"
	IdentityProviderNewResponseAccessFacebookTypeFacebook   IdentityProviderNewResponseAccessFacebookType = "facebook"
	IdentityProviderNewResponseAccessFacebookTypeGitHub     IdentityProviderNewResponseAccessFacebookType = "github"
	IdentityProviderNewResponseAccessFacebookTypeGoogleApps IdentityProviderNewResponseAccessFacebookType = "google-apps"
	IdentityProviderNewResponseAccessFacebookTypeGoogle     IdentityProviderNewResponseAccessFacebookType = "google"
	IdentityProviderNewResponseAccessFacebookTypeLinkedin   IdentityProviderNewResponseAccessFacebookType = "linkedin"
	IdentityProviderNewResponseAccessFacebookTypeOidc       IdentityProviderNewResponseAccessFacebookType = "oidc"
	IdentityProviderNewResponseAccessFacebookTypeOkta       IdentityProviderNewResponseAccessFacebookType = "okta"
	IdentityProviderNewResponseAccessFacebookTypeOnelogin   IdentityProviderNewResponseAccessFacebookType = "onelogin"
	IdentityProviderNewResponseAccessFacebookTypePingone    IdentityProviderNewResponseAccessFacebookType = "pingone"
	IdentityProviderNewResponseAccessFacebookTypeYandex     IdentityProviderNewResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewResponseAccessFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderNewResponseAccessFacebookScimConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessFacebookScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderNewResponseAccessFacebookScimConfig]
type identityProviderNewResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessFacebookScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderNewResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderNewResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       identityProviderNewResponseAccessGitHubJSON       `json:"-"`
}

// identityProviderNewResponseAccessGitHubJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseAccessGitHub]
type identityProviderNewResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessGitHubJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessGitHub) implementsZeroTrustIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                            `json:"client_secret"`
	JSON         identityProviderNewResponseAccessGitHubConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessGitHubConfigJSON contains the JSON metadata for
// the struct [IdentityProviderNewResponseAccessGitHubConfig]
type identityProviderNewResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessGitHubConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessGitHubType string

const (
	IdentityProviderNewResponseAccessGitHubTypeOnetimepin IdentityProviderNewResponseAccessGitHubType = "onetimepin"
	IdentityProviderNewResponseAccessGitHubTypeAzureAd    IdentityProviderNewResponseAccessGitHubType = "azureAD"
	IdentityProviderNewResponseAccessGitHubTypeSaml       IdentityProviderNewResponseAccessGitHubType = "saml"
	IdentityProviderNewResponseAccessGitHubTypeCentrify   IdentityProviderNewResponseAccessGitHubType = "centrify"
	IdentityProviderNewResponseAccessGitHubTypeFacebook   IdentityProviderNewResponseAccessGitHubType = "facebook"
	IdentityProviderNewResponseAccessGitHubTypeGitHub     IdentityProviderNewResponseAccessGitHubType = "github"
	IdentityProviderNewResponseAccessGitHubTypeGoogleApps IdentityProviderNewResponseAccessGitHubType = "google-apps"
	IdentityProviderNewResponseAccessGitHubTypeGoogle     IdentityProviderNewResponseAccessGitHubType = "google"
	IdentityProviderNewResponseAccessGitHubTypeLinkedin   IdentityProviderNewResponseAccessGitHubType = "linkedin"
	IdentityProviderNewResponseAccessGitHubTypeOidc       IdentityProviderNewResponseAccessGitHubType = "oidc"
	IdentityProviderNewResponseAccessGitHubTypeOkta       IdentityProviderNewResponseAccessGitHubType = "okta"
	IdentityProviderNewResponseAccessGitHubTypeOnelogin   IdentityProviderNewResponseAccessGitHubType = "onelogin"
	IdentityProviderNewResponseAccessGitHubTypePingone    IdentityProviderNewResponseAccessGitHubType = "pingone"
	IdentityProviderNewResponseAccessGitHubTypeYandex     IdentityProviderNewResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewResponseAccessGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderNewResponseAccessGitHubScimConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessGitHubScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderNewResponseAccessGitHubScimConfig]
type identityProviderNewResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessGitHubScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderNewResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderNewResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       identityProviderNewResponseAccessGoogleJSON       `json:"-"`
}

// identityProviderNewResponseAccessGoogleJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseAccessGoogle]
type identityProviderNewResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessGoogleJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessGoogle) implementsZeroTrustIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                            `json:"email_claim_name"`
	JSON           identityProviderNewResponseAccessGoogleConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessGoogleConfigJSON contains the JSON metadata for
// the struct [IdentityProviderNewResponseAccessGoogleConfig]
type identityProviderNewResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessGoogleConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessGoogleType string

const (
	IdentityProviderNewResponseAccessGoogleTypeOnetimepin IdentityProviderNewResponseAccessGoogleType = "onetimepin"
	IdentityProviderNewResponseAccessGoogleTypeAzureAd    IdentityProviderNewResponseAccessGoogleType = "azureAD"
	IdentityProviderNewResponseAccessGoogleTypeSaml       IdentityProviderNewResponseAccessGoogleType = "saml"
	IdentityProviderNewResponseAccessGoogleTypeCentrify   IdentityProviderNewResponseAccessGoogleType = "centrify"
	IdentityProviderNewResponseAccessGoogleTypeFacebook   IdentityProviderNewResponseAccessGoogleType = "facebook"
	IdentityProviderNewResponseAccessGoogleTypeGitHub     IdentityProviderNewResponseAccessGoogleType = "github"
	IdentityProviderNewResponseAccessGoogleTypeGoogleApps IdentityProviderNewResponseAccessGoogleType = "google-apps"
	IdentityProviderNewResponseAccessGoogleTypeGoogle     IdentityProviderNewResponseAccessGoogleType = "google"
	IdentityProviderNewResponseAccessGoogleTypeLinkedin   IdentityProviderNewResponseAccessGoogleType = "linkedin"
	IdentityProviderNewResponseAccessGoogleTypeOidc       IdentityProviderNewResponseAccessGoogleType = "oidc"
	IdentityProviderNewResponseAccessGoogleTypeOkta       IdentityProviderNewResponseAccessGoogleType = "okta"
	IdentityProviderNewResponseAccessGoogleTypeOnelogin   IdentityProviderNewResponseAccessGoogleType = "onelogin"
	IdentityProviderNewResponseAccessGoogleTypePingone    IdentityProviderNewResponseAccessGoogleType = "pingone"
	IdentityProviderNewResponseAccessGoogleTypeYandex     IdentityProviderNewResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewResponseAccessGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderNewResponseAccessGoogleScimConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessGoogleScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderNewResponseAccessGoogleScimConfig]
type identityProviderNewResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessGoogleScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderNewResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderNewResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       identityProviderNewResponseAccessGoogleAppsJSON       `json:"-"`
}

// identityProviderNewResponseAccessGoogleAppsJSON contains the JSON metadata for
// the struct [IdentityProviderNewResponseAccessGoogleApps]
type identityProviderNewResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessGoogleAppsJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessGoogleApps) implementsZeroTrustIdentityProviderNewResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                `json:"email_claim_name"`
	JSON           identityProviderNewResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessGoogleAppsConfigJSON contains the JSON metadata
// for the struct [IdentityProviderNewResponseAccessGoogleAppsConfig]
type identityProviderNewResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessGoogleAppsConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessGoogleAppsType string

const (
	IdentityProviderNewResponseAccessGoogleAppsTypeOnetimepin IdentityProviderNewResponseAccessGoogleAppsType = "onetimepin"
	IdentityProviderNewResponseAccessGoogleAppsTypeAzureAd    IdentityProviderNewResponseAccessGoogleAppsType = "azureAD"
	IdentityProviderNewResponseAccessGoogleAppsTypeSaml       IdentityProviderNewResponseAccessGoogleAppsType = "saml"
	IdentityProviderNewResponseAccessGoogleAppsTypeCentrify   IdentityProviderNewResponseAccessGoogleAppsType = "centrify"
	IdentityProviderNewResponseAccessGoogleAppsTypeFacebook   IdentityProviderNewResponseAccessGoogleAppsType = "facebook"
	IdentityProviderNewResponseAccessGoogleAppsTypeGitHub     IdentityProviderNewResponseAccessGoogleAppsType = "github"
	IdentityProviderNewResponseAccessGoogleAppsTypeGoogleApps IdentityProviderNewResponseAccessGoogleAppsType = "google-apps"
	IdentityProviderNewResponseAccessGoogleAppsTypeGoogle     IdentityProviderNewResponseAccessGoogleAppsType = "google"
	IdentityProviderNewResponseAccessGoogleAppsTypeLinkedin   IdentityProviderNewResponseAccessGoogleAppsType = "linkedin"
	IdentityProviderNewResponseAccessGoogleAppsTypeOidc       IdentityProviderNewResponseAccessGoogleAppsType = "oidc"
	IdentityProviderNewResponseAccessGoogleAppsTypeOkta       IdentityProviderNewResponseAccessGoogleAppsType = "okta"
	IdentityProviderNewResponseAccessGoogleAppsTypeOnelogin   IdentityProviderNewResponseAccessGoogleAppsType = "onelogin"
	IdentityProviderNewResponseAccessGoogleAppsTypePingone    IdentityProviderNewResponseAccessGoogleAppsType = "pingone"
	IdentityProviderNewResponseAccessGoogleAppsTypeYandex     IdentityProviderNewResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewResponseAccessGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderNewResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessGoogleAppsScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderNewResponseAccessGoogleAppsScimConfig]
type identityProviderNewResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessGoogleAppsScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderNewResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderNewResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       identityProviderNewResponseAccessLinkedinJSON       `json:"-"`
}

// identityProviderNewResponseAccessLinkedinJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseAccessLinkedin]
type identityProviderNewResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessLinkedinJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessLinkedin) implementsZeroTrustIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                              `json:"client_secret"`
	JSON         identityProviderNewResponseAccessLinkedinConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessLinkedinConfigJSON contains the JSON metadata
// for the struct [IdentityProviderNewResponseAccessLinkedinConfig]
type identityProviderNewResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessLinkedinConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessLinkedinType string

const (
	IdentityProviderNewResponseAccessLinkedinTypeOnetimepin IdentityProviderNewResponseAccessLinkedinType = "onetimepin"
	IdentityProviderNewResponseAccessLinkedinTypeAzureAd    IdentityProviderNewResponseAccessLinkedinType = "azureAD"
	IdentityProviderNewResponseAccessLinkedinTypeSaml       IdentityProviderNewResponseAccessLinkedinType = "saml"
	IdentityProviderNewResponseAccessLinkedinTypeCentrify   IdentityProviderNewResponseAccessLinkedinType = "centrify"
	IdentityProviderNewResponseAccessLinkedinTypeFacebook   IdentityProviderNewResponseAccessLinkedinType = "facebook"
	IdentityProviderNewResponseAccessLinkedinTypeGitHub     IdentityProviderNewResponseAccessLinkedinType = "github"
	IdentityProviderNewResponseAccessLinkedinTypeGoogleApps IdentityProviderNewResponseAccessLinkedinType = "google-apps"
	IdentityProviderNewResponseAccessLinkedinTypeGoogle     IdentityProviderNewResponseAccessLinkedinType = "google"
	IdentityProviderNewResponseAccessLinkedinTypeLinkedin   IdentityProviderNewResponseAccessLinkedinType = "linkedin"
	IdentityProviderNewResponseAccessLinkedinTypeOidc       IdentityProviderNewResponseAccessLinkedinType = "oidc"
	IdentityProviderNewResponseAccessLinkedinTypeOkta       IdentityProviderNewResponseAccessLinkedinType = "okta"
	IdentityProviderNewResponseAccessLinkedinTypeOnelogin   IdentityProviderNewResponseAccessLinkedinType = "onelogin"
	IdentityProviderNewResponseAccessLinkedinTypePingone    IdentityProviderNewResponseAccessLinkedinType = "pingone"
	IdentityProviderNewResponseAccessLinkedinTypeYandex     IdentityProviderNewResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewResponseAccessLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderNewResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessLinkedinScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderNewResponseAccessLinkedinScimConfig]
type identityProviderNewResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessLinkedinScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderNewResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderNewResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       identityProviderNewResponseAccessOidcJSON       `json:"-"`
}

// identityProviderNewResponseAccessOidcJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseAccessOidc]
type identityProviderNewResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessOidcJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessOidc) implementsZeroTrustIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessOidcConfig struct {
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
	TokenURL string                                          `json:"token_url"`
	JSON     identityProviderNewResponseAccessOidcConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessOidcConfigJSON contains the JSON metadata for
// the struct [IdentityProviderNewResponseAccessOidcConfig]
type identityProviderNewResponseAccessOidcConfigJSON struct {
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

func (r *IdentityProviderNewResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessOidcConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessOidcType string

const (
	IdentityProviderNewResponseAccessOidcTypeOnetimepin IdentityProviderNewResponseAccessOidcType = "onetimepin"
	IdentityProviderNewResponseAccessOidcTypeAzureAd    IdentityProviderNewResponseAccessOidcType = "azureAD"
	IdentityProviderNewResponseAccessOidcTypeSaml       IdentityProviderNewResponseAccessOidcType = "saml"
	IdentityProviderNewResponseAccessOidcTypeCentrify   IdentityProviderNewResponseAccessOidcType = "centrify"
	IdentityProviderNewResponseAccessOidcTypeFacebook   IdentityProviderNewResponseAccessOidcType = "facebook"
	IdentityProviderNewResponseAccessOidcTypeGitHub     IdentityProviderNewResponseAccessOidcType = "github"
	IdentityProviderNewResponseAccessOidcTypeGoogleApps IdentityProviderNewResponseAccessOidcType = "google-apps"
	IdentityProviderNewResponseAccessOidcTypeGoogle     IdentityProviderNewResponseAccessOidcType = "google"
	IdentityProviderNewResponseAccessOidcTypeLinkedin   IdentityProviderNewResponseAccessOidcType = "linkedin"
	IdentityProviderNewResponseAccessOidcTypeOidc       IdentityProviderNewResponseAccessOidcType = "oidc"
	IdentityProviderNewResponseAccessOidcTypeOkta       IdentityProviderNewResponseAccessOidcType = "okta"
	IdentityProviderNewResponseAccessOidcTypeOnelogin   IdentityProviderNewResponseAccessOidcType = "onelogin"
	IdentityProviderNewResponseAccessOidcTypePingone    IdentityProviderNewResponseAccessOidcType = "pingone"
	IdentityProviderNewResponseAccessOidcTypeYandex     IdentityProviderNewResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewResponseAccessOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderNewResponseAccessOidcScimConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessOidcScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderNewResponseAccessOidcScimConfig]
type identityProviderNewResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessOidcScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderNewResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderNewResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       identityProviderNewResponseAccessOktaJSON       `json:"-"`
}

// identityProviderNewResponseAccessOktaJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseAccessOkta]
type identityProviderNewResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessOktaJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessOkta) implementsZeroTrustIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessOktaConfig struct {
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
	OktaAccount string                                          `json:"okta_account"`
	JSON        identityProviderNewResponseAccessOktaConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessOktaConfigJSON contains the JSON metadata for
// the struct [IdentityProviderNewResponseAccessOktaConfig]
type identityProviderNewResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessOktaConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessOktaType string

const (
	IdentityProviderNewResponseAccessOktaTypeOnetimepin IdentityProviderNewResponseAccessOktaType = "onetimepin"
	IdentityProviderNewResponseAccessOktaTypeAzureAd    IdentityProviderNewResponseAccessOktaType = "azureAD"
	IdentityProviderNewResponseAccessOktaTypeSaml       IdentityProviderNewResponseAccessOktaType = "saml"
	IdentityProviderNewResponseAccessOktaTypeCentrify   IdentityProviderNewResponseAccessOktaType = "centrify"
	IdentityProviderNewResponseAccessOktaTypeFacebook   IdentityProviderNewResponseAccessOktaType = "facebook"
	IdentityProviderNewResponseAccessOktaTypeGitHub     IdentityProviderNewResponseAccessOktaType = "github"
	IdentityProviderNewResponseAccessOktaTypeGoogleApps IdentityProviderNewResponseAccessOktaType = "google-apps"
	IdentityProviderNewResponseAccessOktaTypeGoogle     IdentityProviderNewResponseAccessOktaType = "google"
	IdentityProviderNewResponseAccessOktaTypeLinkedin   IdentityProviderNewResponseAccessOktaType = "linkedin"
	IdentityProviderNewResponseAccessOktaTypeOidc       IdentityProviderNewResponseAccessOktaType = "oidc"
	IdentityProviderNewResponseAccessOktaTypeOkta       IdentityProviderNewResponseAccessOktaType = "okta"
	IdentityProviderNewResponseAccessOktaTypeOnelogin   IdentityProviderNewResponseAccessOktaType = "onelogin"
	IdentityProviderNewResponseAccessOktaTypePingone    IdentityProviderNewResponseAccessOktaType = "pingone"
	IdentityProviderNewResponseAccessOktaTypeYandex     IdentityProviderNewResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewResponseAccessOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderNewResponseAccessOktaScimConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessOktaScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderNewResponseAccessOktaScimConfig]
type identityProviderNewResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessOktaScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderNewResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderNewResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       identityProviderNewResponseAccessOneloginJSON       `json:"-"`
}

// identityProviderNewResponseAccessOneloginJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseAccessOnelogin]
type identityProviderNewResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessOneloginJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessOnelogin) implementsZeroTrustIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                              `json:"onelogin_account"`
	JSON            identityProviderNewResponseAccessOneloginConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessOneloginConfigJSON contains the JSON metadata
// for the struct [IdentityProviderNewResponseAccessOneloginConfig]
type identityProviderNewResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessOneloginConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessOneloginType string

const (
	IdentityProviderNewResponseAccessOneloginTypeOnetimepin IdentityProviderNewResponseAccessOneloginType = "onetimepin"
	IdentityProviderNewResponseAccessOneloginTypeAzureAd    IdentityProviderNewResponseAccessOneloginType = "azureAD"
	IdentityProviderNewResponseAccessOneloginTypeSaml       IdentityProviderNewResponseAccessOneloginType = "saml"
	IdentityProviderNewResponseAccessOneloginTypeCentrify   IdentityProviderNewResponseAccessOneloginType = "centrify"
	IdentityProviderNewResponseAccessOneloginTypeFacebook   IdentityProviderNewResponseAccessOneloginType = "facebook"
	IdentityProviderNewResponseAccessOneloginTypeGitHub     IdentityProviderNewResponseAccessOneloginType = "github"
	IdentityProviderNewResponseAccessOneloginTypeGoogleApps IdentityProviderNewResponseAccessOneloginType = "google-apps"
	IdentityProviderNewResponseAccessOneloginTypeGoogle     IdentityProviderNewResponseAccessOneloginType = "google"
	IdentityProviderNewResponseAccessOneloginTypeLinkedin   IdentityProviderNewResponseAccessOneloginType = "linkedin"
	IdentityProviderNewResponseAccessOneloginTypeOidc       IdentityProviderNewResponseAccessOneloginType = "oidc"
	IdentityProviderNewResponseAccessOneloginTypeOkta       IdentityProviderNewResponseAccessOneloginType = "okta"
	IdentityProviderNewResponseAccessOneloginTypeOnelogin   IdentityProviderNewResponseAccessOneloginType = "onelogin"
	IdentityProviderNewResponseAccessOneloginTypePingone    IdentityProviderNewResponseAccessOneloginType = "pingone"
	IdentityProviderNewResponseAccessOneloginTypeYandex     IdentityProviderNewResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewResponseAccessOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderNewResponseAccessOneloginScimConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessOneloginScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderNewResponseAccessOneloginScimConfig]
type identityProviderNewResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessOneloginScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderNewResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderNewResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       identityProviderNewResponseAccessPingoneJSON       `json:"-"`
}

// identityProviderNewResponseAccessPingoneJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseAccessPingone]
type identityProviderNewResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessPingoneJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessPingone) implementsZeroTrustIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                             `json:"ping_env_id"`
	JSON      identityProviderNewResponseAccessPingoneConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessPingoneConfigJSON contains the JSON metadata
// for the struct [IdentityProviderNewResponseAccessPingoneConfig]
type identityProviderNewResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessPingoneConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessPingoneType string

const (
	IdentityProviderNewResponseAccessPingoneTypeOnetimepin IdentityProviderNewResponseAccessPingoneType = "onetimepin"
	IdentityProviderNewResponseAccessPingoneTypeAzureAd    IdentityProviderNewResponseAccessPingoneType = "azureAD"
	IdentityProviderNewResponseAccessPingoneTypeSaml       IdentityProviderNewResponseAccessPingoneType = "saml"
	IdentityProviderNewResponseAccessPingoneTypeCentrify   IdentityProviderNewResponseAccessPingoneType = "centrify"
	IdentityProviderNewResponseAccessPingoneTypeFacebook   IdentityProviderNewResponseAccessPingoneType = "facebook"
	IdentityProviderNewResponseAccessPingoneTypeGitHub     IdentityProviderNewResponseAccessPingoneType = "github"
	IdentityProviderNewResponseAccessPingoneTypeGoogleApps IdentityProviderNewResponseAccessPingoneType = "google-apps"
	IdentityProviderNewResponseAccessPingoneTypeGoogle     IdentityProviderNewResponseAccessPingoneType = "google"
	IdentityProviderNewResponseAccessPingoneTypeLinkedin   IdentityProviderNewResponseAccessPingoneType = "linkedin"
	IdentityProviderNewResponseAccessPingoneTypeOidc       IdentityProviderNewResponseAccessPingoneType = "oidc"
	IdentityProviderNewResponseAccessPingoneTypeOkta       IdentityProviderNewResponseAccessPingoneType = "okta"
	IdentityProviderNewResponseAccessPingoneTypeOnelogin   IdentityProviderNewResponseAccessPingoneType = "onelogin"
	IdentityProviderNewResponseAccessPingoneTypePingone    IdentityProviderNewResponseAccessPingoneType = "pingone"
	IdentityProviderNewResponseAccessPingoneTypeYandex     IdentityProviderNewResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewResponseAccessPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderNewResponseAccessPingoneScimConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessPingoneScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderNewResponseAccessPingoneScimConfig]
type identityProviderNewResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessPingoneScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderNewResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderNewResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       identityProviderNewResponseAccessSamlJSON       `json:"-"`
}

// identityProviderNewResponseAccessSamlJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseAccessSaml]
type identityProviderNewResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessSamlJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessSaml) implementsZeroTrustIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []IdentityProviderNewResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                          `json:"sso_target_url"`
	JSON         identityProviderNewResponseAccessSamlConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessSamlConfigJSON contains the JSON metadata for
// the struct [IdentityProviderNewResponseAccessSamlConfig]
type identityProviderNewResponseAccessSamlConfigJSON struct {
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

func (r *IdentityProviderNewResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessSamlConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                         `json:"header_name"`
	JSON       identityProviderNewResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// identityProviderNewResponseAccessSamlConfigHeaderAttributeJSON contains the JSON
// metadata for the struct
// [IdentityProviderNewResponseAccessSamlConfigHeaderAttribute]
type identityProviderNewResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessSamlConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessSamlType string

const (
	IdentityProviderNewResponseAccessSamlTypeOnetimepin IdentityProviderNewResponseAccessSamlType = "onetimepin"
	IdentityProviderNewResponseAccessSamlTypeAzureAd    IdentityProviderNewResponseAccessSamlType = "azureAD"
	IdentityProviderNewResponseAccessSamlTypeSaml       IdentityProviderNewResponseAccessSamlType = "saml"
	IdentityProviderNewResponseAccessSamlTypeCentrify   IdentityProviderNewResponseAccessSamlType = "centrify"
	IdentityProviderNewResponseAccessSamlTypeFacebook   IdentityProviderNewResponseAccessSamlType = "facebook"
	IdentityProviderNewResponseAccessSamlTypeGitHub     IdentityProviderNewResponseAccessSamlType = "github"
	IdentityProviderNewResponseAccessSamlTypeGoogleApps IdentityProviderNewResponseAccessSamlType = "google-apps"
	IdentityProviderNewResponseAccessSamlTypeGoogle     IdentityProviderNewResponseAccessSamlType = "google"
	IdentityProviderNewResponseAccessSamlTypeLinkedin   IdentityProviderNewResponseAccessSamlType = "linkedin"
	IdentityProviderNewResponseAccessSamlTypeOidc       IdentityProviderNewResponseAccessSamlType = "oidc"
	IdentityProviderNewResponseAccessSamlTypeOkta       IdentityProviderNewResponseAccessSamlType = "okta"
	IdentityProviderNewResponseAccessSamlTypeOnelogin   IdentityProviderNewResponseAccessSamlType = "onelogin"
	IdentityProviderNewResponseAccessSamlTypePingone    IdentityProviderNewResponseAccessSamlType = "pingone"
	IdentityProviderNewResponseAccessSamlTypeYandex     IdentityProviderNewResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewResponseAccessSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderNewResponseAccessSamlScimConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessSamlScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderNewResponseAccessSamlScimConfig]
type identityProviderNewResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessSamlScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderNewResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderNewResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       identityProviderNewResponseAccessYandexJSON       `json:"-"`
}

// identityProviderNewResponseAccessYandexJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseAccessYandex]
type identityProviderNewResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessYandexJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessYandex) implementsZeroTrustIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                            `json:"client_secret"`
	JSON         identityProviderNewResponseAccessYandexConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessYandexConfigJSON contains the JSON metadata for
// the struct [IdentityProviderNewResponseAccessYandexConfig]
type identityProviderNewResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessYandexConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessYandexType string

const (
	IdentityProviderNewResponseAccessYandexTypeOnetimepin IdentityProviderNewResponseAccessYandexType = "onetimepin"
	IdentityProviderNewResponseAccessYandexTypeAzureAd    IdentityProviderNewResponseAccessYandexType = "azureAD"
	IdentityProviderNewResponseAccessYandexTypeSaml       IdentityProviderNewResponseAccessYandexType = "saml"
	IdentityProviderNewResponseAccessYandexTypeCentrify   IdentityProviderNewResponseAccessYandexType = "centrify"
	IdentityProviderNewResponseAccessYandexTypeFacebook   IdentityProviderNewResponseAccessYandexType = "facebook"
	IdentityProviderNewResponseAccessYandexTypeGitHub     IdentityProviderNewResponseAccessYandexType = "github"
	IdentityProviderNewResponseAccessYandexTypeGoogleApps IdentityProviderNewResponseAccessYandexType = "google-apps"
	IdentityProviderNewResponseAccessYandexTypeGoogle     IdentityProviderNewResponseAccessYandexType = "google"
	IdentityProviderNewResponseAccessYandexTypeLinkedin   IdentityProviderNewResponseAccessYandexType = "linkedin"
	IdentityProviderNewResponseAccessYandexTypeOidc       IdentityProviderNewResponseAccessYandexType = "oidc"
	IdentityProviderNewResponseAccessYandexTypeOkta       IdentityProviderNewResponseAccessYandexType = "okta"
	IdentityProviderNewResponseAccessYandexTypeOnelogin   IdentityProviderNewResponseAccessYandexType = "onelogin"
	IdentityProviderNewResponseAccessYandexTypePingone    IdentityProviderNewResponseAccessYandexType = "pingone"
	IdentityProviderNewResponseAccessYandexTypeYandex     IdentityProviderNewResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewResponseAccessYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderNewResponseAccessYandexScimConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessYandexScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderNewResponseAccessYandexScimConfig]
type identityProviderNewResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessYandexScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderNewResponseAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderNewResponseAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       identityProviderNewResponseAccessOnetimepinJSON       `json:"-"`
}

// identityProviderNewResponseAccessOnetimepinJSON contains the JSON metadata for
// the struct [IdentityProviderNewResponseAccessOnetimepin]
type identityProviderNewResponseAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessOnetimepinJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessOnetimepin) implementsZeroTrustIdentityProviderNewResponse() {
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessOnetimepinType string

const (
	IdentityProviderNewResponseAccessOnetimepinTypeOnetimepin IdentityProviderNewResponseAccessOnetimepinType = "onetimepin"
	IdentityProviderNewResponseAccessOnetimepinTypeAzureAd    IdentityProviderNewResponseAccessOnetimepinType = "azureAD"
	IdentityProviderNewResponseAccessOnetimepinTypeSaml       IdentityProviderNewResponseAccessOnetimepinType = "saml"
	IdentityProviderNewResponseAccessOnetimepinTypeCentrify   IdentityProviderNewResponseAccessOnetimepinType = "centrify"
	IdentityProviderNewResponseAccessOnetimepinTypeFacebook   IdentityProviderNewResponseAccessOnetimepinType = "facebook"
	IdentityProviderNewResponseAccessOnetimepinTypeGitHub     IdentityProviderNewResponseAccessOnetimepinType = "github"
	IdentityProviderNewResponseAccessOnetimepinTypeGoogleApps IdentityProviderNewResponseAccessOnetimepinType = "google-apps"
	IdentityProviderNewResponseAccessOnetimepinTypeGoogle     IdentityProviderNewResponseAccessOnetimepinType = "google"
	IdentityProviderNewResponseAccessOnetimepinTypeLinkedin   IdentityProviderNewResponseAccessOnetimepinType = "linkedin"
	IdentityProviderNewResponseAccessOnetimepinTypeOidc       IdentityProviderNewResponseAccessOnetimepinType = "oidc"
	IdentityProviderNewResponseAccessOnetimepinTypeOkta       IdentityProviderNewResponseAccessOnetimepinType = "okta"
	IdentityProviderNewResponseAccessOnetimepinTypeOnelogin   IdentityProviderNewResponseAccessOnetimepinType = "onelogin"
	IdentityProviderNewResponseAccessOnetimepinTypePingone    IdentityProviderNewResponseAccessOnetimepinType = "pingone"
	IdentityProviderNewResponseAccessOnetimepinTypeYandex     IdentityProviderNewResponseAccessOnetimepinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewResponseAccessOnetimepinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderNewResponseAccessOnetimepinScimConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessOnetimepinScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderNewResponseAccessOnetimepinScimConfig]
type identityProviderNewResponseAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessOnetimepinScimConfigJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.IdentityProviderUpdateResponseAccessAzureAd],
// [zero_trust.IdentityProviderUpdateResponseAccessCentrify],
// [zero_trust.IdentityProviderUpdateResponseAccessFacebook],
// [zero_trust.IdentityProviderUpdateResponseAccessGitHub],
// [zero_trust.IdentityProviderUpdateResponseAccessGoogle],
// [zero_trust.IdentityProviderUpdateResponseAccessGoogleApps],
// [zero_trust.IdentityProviderUpdateResponseAccessLinkedin],
// [zero_trust.IdentityProviderUpdateResponseAccessOidc],
// [zero_trust.IdentityProviderUpdateResponseAccessOkta],
// [zero_trust.IdentityProviderUpdateResponseAccessOnelogin],
// [zero_trust.IdentityProviderUpdateResponseAccessPingone],
// [zero_trust.IdentityProviderUpdateResponseAccessSaml],
// [zero_trust.IdentityProviderUpdateResponseAccessYandex] or
// [zero_trust.IdentityProviderUpdateResponseAccessOnetimepin].
type IdentityProviderUpdateResponse interface {
	implementsZeroTrustIdentityProviderUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IdentityProviderUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessAzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessCentrify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessFacebook{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessGitHub{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessGoogle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessGoogleApps{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessLinkedin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessOidc{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessOkta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessOnelogin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessPingone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessSaml{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessYandex{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessOnetimepin{}),
		},
	)
}

type IdentityProviderUpdateResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderUpdateResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderUpdateResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessAzureAdJSON       `json:"-"`
}

// identityProviderUpdateResponseAccessAzureAdJSON contains the JSON metadata for
// the struct [IdentityProviderUpdateResponseAccessAzureAd]
type identityProviderUpdateResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessAzureAdJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessAzureAd) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessAzureAdConfig struct {
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
	SupportGroups bool                                                  `json:"support_groups"`
	JSON          identityProviderUpdateResponseAccessAzureAdConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessAzureAdConfigJSON contains the JSON metadata
// for the struct [IdentityProviderUpdateResponseAccessAzureAdConfig]
type identityProviderUpdateResponseAccessAzureAdConfigJSON struct {
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

func (r *IdentityProviderUpdateResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessAzureAdConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessAzureAdType string

const (
	IdentityProviderUpdateResponseAccessAzureAdTypeOnetimepin IdentityProviderUpdateResponseAccessAzureAdType = "onetimepin"
	IdentityProviderUpdateResponseAccessAzureAdTypeAzureAd    IdentityProviderUpdateResponseAccessAzureAdType = "azureAD"
	IdentityProviderUpdateResponseAccessAzureAdTypeSaml       IdentityProviderUpdateResponseAccessAzureAdType = "saml"
	IdentityProviderUpdateResponseAccessAzureAdTypeCentrify   IdentityProviderUpdateResponseAccessAzureAdType = "centrify"
	IdentityProviderUpdateResponseAccessAzureAdTypeFacebook   IdentityProviderUpdateResponseAccessAzureAdType = "facebook"
	IdentityProviderUpdateResponseAccessAzureAdTypeGitHub     IdentityProviderUpdateResponseAccessAzureAdType = "github"
	IdentityProviderUpdateResponseAccessAzureAdTypeGoogleApps IdentityProviderUpdateResponseAccessAzureAdType = "google-apps"
	IdentityProviderUpdateResponseAccessAzureAdTypeGoogle     IdentityProviderUpdateResponseAccessAzureAdType = "google"
	IdentityProviderUpdateResponseAccessAzureAdTypeLinkedin   IdentityProviderUpdateResponseAccessAzureAdType = "linkedin"
	IdentityProviderUpdateResponseAccessAzureAdTypeOidc       IdentityProviderUpdateResponseAccessAzureAdType = "oidc"
	IdentityProviderUpdateResponseAccessAzureAdTypeOkta       IdentityProviderUpdateResponseAccessAzureAdType = "okta"
	IdentityProviderUpdateResponseAccessAzureAdTypeOnelogin   IdentityProviderUpdateResponseAccessAzureAdType = "onelogin"
	IdentityProviderUpdateResponseAccessAzureAdTypePingone    IdentityProviderUpdateResponseAccessAzureAdType = "pingone"
	IdentityProviderUpdateResponseAccessAzureAdTypeYandex     IdentityProviderUpdateResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateResponseAccessAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderUpdateResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessAzureAdScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessAzureAdScimConfig]
type identityProviderUpdateResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessAzureAdScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderUpdateResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderUpdateResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessCentrifyJSON       `json:"-"`
}

// identityProviderUpdateResponseAccessCentrifyJSON contains the JSON metadata for
// the struct [IdentityProviderUpdateResponseAccessCentrify]
type identityProviderUpdateResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessCentrifyJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessCentrify) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessCentrifyConfig struct {
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
	EmailClaimName string                                                 `json:"email_claim_name"`
	JSON           identityProviderUpdateResponseAccessCentrifyConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessCentrifyConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessCentrifyConfig]
type identityProviderUpdateResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessCentrifyConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessCentrifyType string

const (
	IdentityProviderUpdateResponseAccessCentrifyTypeOnetimepin IdentityProviderUpdateResponseAccessCentrifyType = "onetimepin"
	IdentityProviderUpdateResponseAccessCentrifyTypeAzureAd    IdentityProviderUpdateResponseAccessCentrifyType = "azureAD"
	IdentityProviderUpdateResponseAccessCentrifyTypeSaml       IdentityProviderUpdateResponseAccessCentrifyType = "saml"
	IdentityProviderUpdateResponseAccessCentrifyTypeCentrify   IdentityProviderUpdateResponseAccessCentrifyType = "centrify"
	IdentityProviderUpdateResponseAccessCentrifyTypeFacebook   IdentityProviderUpdateResponseAccessCentrifyType = "facebook"
	IdentityProviderUpdateResponseAccessCentrifyTypeGitHub     IdentityProviderUpdateResponseAccessCentrifyType = "github"
	IdentityProviderUpdateResponseAccessCentrifyTypeGoogleApps IdentityProviderUpdateResponseAccessCentrifyType = "google-apps"
	IdentityProviderUpdateResponseAccessCentrifyTypeGoogle     IdentityProviderUpdateResponseAccessCentrifyType = "google"
	IdentityProviderUpdateResponseAccessCentrifyTypeLinkedin   IdentityProviderUpdateResponseAccessCentrifyType = "linkedin"
	IdentityProviderUpdateResponseAccessCentrifyTypeOidc       IdentityProviderUpdateResponseAccessCentrifyType = "oidc"
	IdentityProviderUpdateResponseAccessCentrifyTypeOkta       IdentityProviderUpdateResponseAccessCentrifyType = "okta"
	IdentityProviderUpdateResponseAccessCentrifyTypeOnelogin   IdentityProviderUpdateResponseAccessCentrifyType = "onelogin"
	IdentityProviderUpdateResponseAccessCentrifyTypePingone    IdentityProviderUpdateResponseAccessCentrifyType = "pingone"
	IdentityProviderUpdateResponseAccessCentrifyTypeYandex     IdentityProviderUpdateResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateResponseAccessCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderUpdateResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessCentrifyScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessCentrifyScimConfig]
type identityProviderUpdateResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessCentrifyScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderUpdateResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderUpdateResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessFacebookJSON       `json:"-"`
}

// identityProviderUpdateResponseAccessFacebookJSON contains the JSON metadata for
// the struct [IdentityProviderUpdateResponseAccessFacebook]
type identityProviderUpdateResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessFacebookJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessFacebook) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                 `json:"client_secret"`
	JSON         identityProviderUpdateResponseAccessFacebookConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessFacebookConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessFacebookConfig]
type identityProviderUpdateResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessFacebookConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessFacebookType string

const (
	IdentityProviderUpdateResponseAccessFacebookTypeOnetimepin IdentityProviderUpdateResponseAccessFacebookType = "onetimepin"
	IdentityProviderUpdateResponseAccessFacebookTypeAzureAd    IdentityProviderUpdateResponseAccessFacebookType = "azureAD"
	IdentityProviderUpdateResponseAccessFacebookTypeSaml       IdentityProviderUpdateResponseAccessFacebookType = "saml"
	IdentityProviderUpdateResponseAccessFacebookTypeCentrify   IdentityProviderUpdateResponseAccessFacebookType = "centrify"
	IdentityProviderUpdateResponseAccessFacebookTypeFacebook   IdentityProviderUpdateResponseAccessFacebookType = "facebook"
	IdentityProviderUpdateResponseAccessFacebookTypeGitHub     IdentityProviderUpdateResponseAccessFacebookType = "github"
	IdentityProviderUpdateResponseAccessFacebookTypeGoogleApps IdentityProviderUpdateResponseAccessFacebookType = "google-apps"
	IdentityProviderUpdateResponseAccessFacebookTypeGoogle     IdentityProviderUpdateResponseAccessFacebookType = "google"
	IdentityProviderUpdateResponseAccessFacebookTypeLinkedin   IdentityProviderUpdateResponseAccessFacebookType = "linkedin"
	IdentityProviderUpdateResponseAccessFacebookTypeOidc       IdentityProviderUpdateResponseAccessFacebookType = "oidc"
	IdentityProviderUpdateResponseAccessFacebookTypeOkta       IdentityProviderUpdateResponseAccessFacebookType = "okta"
	IdentityProviderUpdateResponseAccessFacebookTypeOnelogin   IdentityProviderUpdateResponseAccessFacebookType = "onelogin"
	IdentityProviderUpdateResponseAccessFacebookTypePingone    IdentityProviderUpdateResponseAccessFacebookType = "pingone"
	IdentityProviderUpdateResponseAccessFacebookTypeYandex     IdentityProviderUpdateResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateResponseAccessFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderUpdateResponseAccessFacebookScimConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessFacebookScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessFacebookScimConfig]
type identityProviderUpdateResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessFacebookScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderUpdateResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderUpdateResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessGitHubJSON       `json:"-"`
}

// identityProviderUpdateResponseAccessGitHubJSON contains the JSON metadata for
// the struct [IdentityProviderUpdateResponseAccessGitHub]
type identityProviderUpdateResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessGitHubJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessGitHub) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                               `json:"client_secret"`
	JSON         identityProviderUpdateResponseAccessGitHubConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessGitHubConfigJSON contains the JSON metadata
// for the struct [IdentityProviderUpdateResponseAccessGitHubConfig]
type identityProviderUpdateResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessGitHubConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessGitHubType string

const (
	IdentityProviderUpdateResponseAccessGitHubTypeOnetimepin IdentityProviderUpdateResponseAccessGitHubType = "onetimepin"
	IdentityProviderUpdateResponseAccessGitHubTypeAzureAd    IdentityProviderUpdateResponseAccessGitHubType = "azureAD"
	IdentityProviderUpdateResponseAccessGitHubTypeSaml       IdentityProviderUpdateResponseAccessGitHubType = "saml"
	IdentityProviderUpdateResponseAccessGitHubTypeCentrify   IdentityProviderUpdateResponseAccessGitHubType = "centrify"
	IdentityProviderUpdateResponseAccessGitHubTypeFacebook   IdentityProviderUpdateResponseAccessGitHubType = "facebook"
	IdentityProviderUpdateResponseAccessGitHubTypeGitHub     IdentityProviderUpdateResponseAccessGitHubType = "github"
	IdentityProviderUpdateResponseAccessGitHubTypeGoogleApps IdentityProviderUpdateResponseAccessGitHubType = "google-apps"
	IdentityProviderUpdateResponseAccessGitHubTypeGoogle     IdentityProviderUpdateResponseAccessGitHubType = "google"
	IdentityProviderUpdateResponseAccessGitHubTypeLinkedin   IdentityProviderUpdateResponseAccessGitHubType = "linkedin"
	IdentityProviderUpdateResponseAccessGitHubTypeOidc       IdentityProviderUpdateResponseAccessGitHubType = "oidc"
	IdentityProviderUpdateResponseAccessGitHubTypeOkta       IdentityProviderUpdateResponseAccessGitHubType = "okta"
	IdentityProviderUpdateResponseAccessGitHubTypeOnelogin   IdentityProviderUpdateResponseAccessGitHubType = "onelogin"
	IdentityProviderUpdateResponseAccessGitHubTypePingone    IdentityProviderUpdateResponseAccessGitHubType = "pingone"
	IdentityProviderUpdateResponseAccessGitHubTypeYandex     IdentityProviderUpdateResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateResponseAccessGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderUpdateResponseAccessGitHubScimConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessGitHubScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessGitHubScimConfig]
type identityProviderUpdateResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessGitHubScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderUpdateResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderUpdateResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessGoogleJSON       `json:"-"`
}

// identityProviderUpdateResponseAccessGoogleJSON contains the JSON metadata for
// the struct [IdentityProviderUpdateResponseAccessGoogle]
type identityProviderUpdateResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessGoogleJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessGoogle) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                               `json:"email_claim_name"`
	JSON           identityProviderUpdateResponseAccessGoogleConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessGoogleConfigJSON contains the JSON metadata
// for the struct [IdentityProviderUpdateResponseAccessGoogleConfig]
type identityProviderUpdateResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessGoogleConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessGoogleType string

const (
	IdentityProviderUpdateResponseAccessGoogleTypeOnetimepin IdentityProviderUpdateResponseAccessGoogleType = "onetimepin"
	IdentityProviderUpdateResponseAccessGoogleTypeAzureAd    IdentityProviderUpdateResponseAccessGoogleType = "azureAD"
	IdentityProviderUpdateResponseAccessGoogleTypeSaml       IdentityProviderUpdateResponseAccessGoogleType = "saml"
	IdentityProviderUpdateResponseAccessGoogleTypeCentrify   IdentityProviderUpdateResponseAccessGoogleType = "centrify"
	IdentityProviderUpdateResponseAccessGoogleTypeFacebook   IdentityProviderUpdateResponseAccessGoogleType = "facebook"
	IdentityProviderUpdateResponseAccessGoogleTypeGitHub     IdentityProviderUpdateResponseAccessGoogleType = "github"
	IdentityProviderUpdateResponseAccessGoogleTypeGoogleApps IdentityProviderUpdateResponseAccessGoogleType = "google-apps"
	IdentityProviderUpdateResponseAccessGoogleTypeGoogle     IdentityProviderUpdateResponseAccessGoogleType = "google"
	IdentityProviderUpdateResponseAccessGoogleTypeLinkedin   IdentityProviderUpdateResponseAccessGoogleType = "linkedin"
	IdentityProviderUpdateResponseAccessGoogleTypeOidc       IdentityProviderUpdateResponseAccessGoogleType = "oidc"
	IdentityProviderUpdateResponseAccessGoogleTypeOkta       IdentityProviderUpdateResponseAccessGoogleType = "okta"
	IdentityProviderUpdateResponseAccessGoogleTypeOnelogin   IdentityProviderUpdateResponseAccessGoogleType = "onelogin"
	IdentityProviderUpdateResponseAccessGoogleTypePingone    IdentityProviderUpdateResponseAccessGoogleType = "pingone"
	IdentityProviderUpdateResponseAccessGoogleTypeYandex     IdentityProviderUpdateResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateResponseAccessGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderUpdateResponseAccessGoogleScimConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessGoogleScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessGoogleScimConfig]
type identityProviderUpdateResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessGoogleScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderUpdateResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderUpdateResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessGoogleAppsJSON       `json:"-"`
}

// identityProviderUpdateResponseAccessGoogleAppsJSON contains the JSON metadata
// for the struct [IdentityProviderUpdateResponseAccessGoogleApps]
type identityProviderUpdateResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessGoogleAppsJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessGoogleApps) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                   `json:"email_claim_name"`
	JSON           identityProviderUpdateResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessGoogleAppsConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessGoogleAppsConfig]
type identityProviderUpdateResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessGoogleAppsConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessGoogleAppsType string

const (
	IdentityProviderUpdateResponseAccessGoogleAppsTypeOnetimepin IdentityProviderUpdateResponseAccessGoogleAppsType = "onetimepin"
	IdentityProviderUpdateResponseAccessGoogleAppsTypeAzureAd    IdentityProviderUpdateResponseAccessGoogleAppsType = "azureAD"
	IdentityProviderUpdateResponseAccessGoogleAppsTypeSaml       IdentityProviderUpdateResponseAccessGoogleAppsType = "saml"
	IdentityProviderUpdateResponseAccessGoogleAppsTypeCentrify   IdentityProviderUpdateResponseAccessGoogleAppsType = "centrify"
	IdentityProviderUpdateResponseAccessGoogleAppsTypeFacebook   IdentityProviderUpdateResponseAccessGoogleAppsType = "facebook"
	IdentityProviderUpdateResponseAccessGoogleAppsTypeGitHub     IdentityProviderUpdateResponseAccessGoogleAppsType = "github"
	IdentityProviderUpdateResponseAccessGoogleAppsTypeGoogleApps IdentityProviderUpdateResponseAccessGoogleAppsType = "google-apps"
	IdentityProviderUpdateResponseAccessGoogleAppsTypeGoogle     IdentityProviderUpdateResponseAccessGoogleAppsType = "google"
	IdentityProviderUpdateResponseAccessGoogleAppsTypeLinkedin   IdentityProviderUpdateResponseAccessGoogleAppsType = "linkedin"
	IdentityProviderUpdateResponseAccessGoogleAppsTypeOidc       IdentityProviderUpdateResponseAccessGoogleAppsType = "oidc"
	IdentityProviderUpdateResponseAccessGoogleAppsTypeOkta       IdentityProviderUpdateResponseAccessGoogleAppsType = "okta"
	IdentityProviderUpdateResponseAccessGoogleAppsTypeOnelogin   IdentityProviderUpdateResponseAccessGoogleAppsType = "onelogin"
	IdentityProviderUpdateResponseAccessGoogleAppsTypePingone    IdentityProviderUpdateResponseAccessGoogleAppsType = "pingone"
	IdentityProviderUpdateResponseAccessGoogleAppsTypeYandex     IdentityProviderUpdateResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateResponseAccessGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderUpdateResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessGoogleAppsScimConfigJSON contains the JSON
// metadata for the struct
// [IdentityProviderUpdateResponseAccessGoogleAppsScimConfig]
type identityProviderUpdateResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessGoogleAppsScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderUpdateResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderUpdateResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessLinkedinJSON       `json:"-"`
}

// identityProviderUpdateResponseAccessLinkedinJSON contains the JSON metadata for
// the struct [IdentityProviderUpdateResponseAccessLinkedin]
type identityProviderUpdateResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessLinkedinJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessLinkedin) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                 `json:"client_secret"`
	JSON         identityProviderUpdateResponseAccessLinkedinConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessLinkedinConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessLinkedinConfig]
type identityProviderUpdateResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessLinkedinConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessLinkedinType string

const (
	IdentityProviderUpdateResponseAccessLinkedinTypeOnetimepin IdentityProviderUpdateResponseAccessLinkedinType = "onetimepin"
	IdentityProviderUpdateResponseAccessLinkedinTypeAzureAd    IdentityProviderUpdateResponseAccessLinkedinType = "azureAD"
	IdentityProviderUpdateResponseAccessLinkedinTypeSaml       IdentityProviderUpdateResponseAccessLinkedinType = "saml"
	IdentityProviderUpdateResponseAccessLinkedinTypeCentrify   IdentityProviderUpdateResponseAccessLinkedinType = "centrify"
	IdentityProviderUpdateResponseAccessLinkedinTypeFacebook   IdentityProviderUpdateResponseAccessLinkedinType = "facebook"
	IdentityProviderUpdateResponseAccessLinkedinTypeGitHub     IdentityProviderUpdateResponseAccessLinkedinType = "github"
	IdentityProviderUpdateResponseAccessLinkedinTypeGoogleApps IdentityProviderUpdateResponseAccessLinkedinType = "google-apps"
	IdentityProviderUpdateResponseAccessLinkedinTypeGoogle     IdentityProviderUpdateResponseAccessLinkedinType = "google"
	IdentityProviderUpdateResponseAccessLinkedinTypeLinkedin   IdentityProviderUpdateResponseAccessLinkedinType = "linkedin"
	IdentityProviderUpdateResponseAccessLinkedinTypeOidc       IdentityProviderUpdateResponseAccessLinkedinType = "oidc"
	IdentityProviderUpdateResponseAccessLinkedinTypeOkta       IdentityProviderUpdateResponseAccessLinkedinType = "okta"
	IdentityProviderUpdateResponseAccessLinkedinTypeOnelogin   IdentityProviderUpdateResponseAccessLinkedinType = "onelogin"
	IdentityProviderUpdateResponseAccessLinkedinTypePingone    IdentityProviderUpdateResponseAccessLinkedinType = "pingone"
	IdentityProviderUpdateResponseAccessLinkedinTypeYandex     IdentityProviderUpdateResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateResponseAccessLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderUpdateResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessLinkedinScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessLinkedinScimConfig]
type identityProviderUpdateResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessLinkedinScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderUpdateResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderUpdateResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessOidcJSON       `json:"-"`
}

// identityProviderUpdateResponseAccessOidcJSON contains the JSON metadata for the
// struct [IdentityProviderUpdateResponseAccessOidc]
type identityProviderUpdateResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessOidcJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessOidc) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessOidcConfig struct {
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
	TokenURL string                                             `json:"token_url"`
	JSON     identityProviderUpdateResponseAccessOidcConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessOidcConfigJSON contains the JSON metadata
// for the struct [IdentityProviderUpdateResponseAccessOidcConfig]
type identityProviderUpdateResponseAccessOidcConfigJSON struct {
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

func (r *IdentityProviderUpdateResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessOidcConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessOidcType string

const (
	IdentityProviderUpdateResponseAccessOidcTypeOnetimepin IdentityProviderUpdateResponseAccessOidcType = "onetimepin"
	IdentityProviderUpdateResponseAccessOidcTypeAzureAd    IdentityProviderUpdateResponseAccessOidcType = "azureAD"
	IdentityProviderUpdateResponseAccessOidcTypeSaml       IdentityProviderUpdateResponseAccessOidcType = "saml"
	IdentityProviderUpdateResponseAccessOidcTypeCentrify   IdentityProviderUpdateResponseAccessOidcType = "centrify"
	IdentityProviderUpdateResponseAccessOidcTypeFacebook   IdentityProviderUpdateResponseAccessOidcType = "facebook"
	IdentityProviderUpdateResponseAccessOidcTypeGitHub     IdentityProviderUpdateResponseAccessOidcType = "github"
	IdentityProviderUpdateResponseAccessOidcTypeGoogleApps IdentityProviderUpdateResponseAccessOidcType = "google-apps"
	IdentityProviderUpdateResponseAccessOidcTypeGoogle     IdentityProviderUpdateResponseAccessOidcType = "google"
	IdentityProviderUpdateResponseAccessOidcTypeLinkedin   IdentityProviderUpdateResponseAccessOidcType = "linkedin"
	IdentityProviderUpdateResponseAccessOidcTypeOidc       IdentityProviderUpdateResponseAccessOidcType = "oidc"
	IdentityProviderUpdateResponseAccessOidcTypeOkta       IdentityProviderUpdateResponseAccessOidcType = "okta"
	IdentityProviderUpdateResponseAccessOidcTypeOnelogin   IdentityProviderUpdateResponseAccessOidcType = "onelogin"
	IdentityProviderUpdateResponseAccessOidcTypePingone    IdentityProviderUpdateResponseAccessOidcType = "pingone"
	IdentityProviderUpdateResponseAccessOidcTypeYandex     IdentityProviderUpdateResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateResponseAccessOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderUpdateResponseAccessOidcScimConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessOidcScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessOidcScimConfig]
type identityProviderUpdateResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessOidcScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderUpdateResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderUpdateResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessOktaJSON       `json:"-"`
}

// identityProviderUpdateResponseAccessOktaJSON contains the JSON metadata for the
// struct [IdentityProviderUpdateResponseAccessOkta]
type identityProviderUpdateResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessOktaJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessOkta) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessOktaConfig struct {
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
	OktaAccount string                                             `json:"okta_account"`
	JSON        identityProviderUpdateResponseAccessOktaConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessOktaConfigJSON contains the JSON metadata
// for the struct [IdentityProviderUpdateResponseAccessOktaConfig]
type identityProviderUpdateResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessOktaConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessOktaType string

const (
	IdentityProviderUpdateResponseAccessOktaTypeOnetimepin IdentityProviderUpdateResponseAccessOktaType = "onetimepin"
	IdentityProviderUpdateResponseAccessOktaTypeAzureAd    IdentityProviderUpdateResponseAccessOktaType = "azureAD"
	IdentityProviderUpdateResponseAccessOktaTypeSaml       IdentityProviderUpdateResponseAccessOktaType = "saml"
	IdentityProviderUpdateResponseAccessOktaTypeCentrify   IdentityProviderUpdateResponseAccessOktaType = "centrify"
	IdentityProviderUpdateResponseAccessOktaTypeFacebook   IdentityProviderUpdateResponseAccessOktaType = "facebook"
	IdentityProviderUpdateResponseAccessOktaTypeGitHub     IdentityProviderUpdateResponseAccessOktaType = "github"
	IdentityProviderUpdateResponseAccessOktaTypeGoogleApps IdentityProviderUpdateResponseAccessOktaType = "google-apps"
	IdentityProviderUpdateResponseAccessOktaTypeGoogle     IdentityProviderUpdateResponseAccessOktaType = "google"
	IdentityProviderUpdateResponseAccessOktaTypeLinkedin   IdentityProviderUpdateResponseAccessOktaType = "linkedin"
	IdentityProviderUpdateResponseAccessOktaTypeOidc       IdentityProviderUpdateResponseAccessOktaType = "oidc"
	IdentityProviderUpdateResponseAccessOktaTypeOkta       IdentityProviderUpdateResponseAccessOktaType = "okta"
	IdentityProviderUpdateResponseAccessOktaTypeOnelogin   IdentityProviderUpdateResponseAccessOktaType = "onelogin"
	IdentityProviderUpdateResponseAccessOktaTypePingone    IdentityProviderUpdateResponseAccessOktaType = "pingone"
	IdentityProviderUpdateResponseAccessOktaTypeYandex     IdentityProviderUpdateResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateResponseAccessOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderUpdateResponseAccessOktaScimConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessOktaScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessOktaScimConfig]
type identityProviderUpdateResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessOktaScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderUpdateResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderUpdateResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessOneloginJSON       `json:"-"`
}

// identityProviderUpdateResponseAccessOneloginJSON contains the JSON metadata for
// the struct [IdentityProviderUpdateResponseAccessOnelogin]
type identityProviderUpdateResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessOneloginJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessOnelogin) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                                 `json:"onelogin_account"`
	JSON            identityProviderUpdateResponseAccessOneloginConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessOneloginConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessOneloginConfig]
type identityProviderUpdateResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessOneloginConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessOneloginType string

const (
	IdentityProviderUpdateResponseAccessOneloginTypeOnetimepin IdentityProviderUpdateResponseAccessOneloginType = "onetimepin"
	IdentityProviderUpdateResponseAccessOneloginTypeAzureAd    IdentityProviderUpdateResponseAccessOneloginType = "azureAD"
	IdentityProviderUpdateResponseAccessOneloginTypeSaml       IdentityProviderUpdateResponseAccessOneloginType = "saml"
	IdentityProviderUpdateResponseAccessOneloginTypeCentrify   IdentityProviderUpdateResponseAccessOneloginType = "centrify"
	IdentityProviderUpdateResponseAccessOneloginTypeFacebook   IdentityProviderUpdateResponseAccessOneloginType = "facebook"
	IdentityProviderUpdateResponseAccessOneloginTypeGitHub     IdentityProviderUpdateResponseAccessOneloginType = "github"
	IdentityProviderUpdateResponseAccessOneloginTypeGoogleApps IdentityProviderUpdateResponseAccessOneloginType = "google-apps"
	IdentityProviderUpdateResponseAccessOneloginTypeGoogle     IdentityProviderUpdateResponseAccessOneloginType = "google"
	IdentityProviderUpdateResponseAccessOneloginTypeLinkedin   IdentityProviderUpdateResponseAccessOneloginType = "linkedin"
	IdentityProviderUpdateResponseAccessOneloginTypeOidc       IdentityProviderUpdateResponseAccessOneloginType = "oidc"
	IdentityProviderUpdateResponseAccessOneloginTypeOkta       IdentityProviderUpdateResponseAccessOneloginType = "okta"
	IdentityProviderUpdateResponseAccessOneloginTypeOnelogin   IdentityProviderUpdateResponseAccessOneloginType = "onelogin"
	IdentityProviderUpdateResponseAccessOneloginTypePingone    IdentityProviderUpdateResponseAccessOneloginType = "pingone"
	IdentityProviderUpdateResponseAccessOneloginTypeYandex     IdentityProviderUpdateResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateResponseAccessOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderUpdateResponseAccessOneloginScimConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessOneloginScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessOneloginScimConfig]
type identityProviderUpdateResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessOneloginScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderUpdateResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderUpdateResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessPingoneJSON       `json:"-"`
}

// identityProviderUpdateResponseAccessPingoneJSON contains the JSON metadata for
// the struct [IdentityProviderUpdateResponseAccessPingone]
type identityProviderUpdateResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessPingoneJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessPingone) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                                `json:"ping_env_id"`
	JSON      identityProviderUpdateResponseAccessPingoneConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessPingoneConfigJSON contains the JSON metadata
// for the struct [IdentityProviderUpdateResponseAccessPingoneConfig]
type identityProviderUpdateResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessPingoneConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessPingoneType string

const (
	IdentityProviderUpdateResponseAccessPingoneTypeOnetimepin IdentityProviderUpdateResponseAccessPingoneType = "onetimepin"
	IdentityProviderUpdateResponseAccessPingoneTypeAzureAd    IdentityProviderUpdateResponseAccessPingoneType = "azureAD"
	IdentityProviderUpdateResponseAccessPingoneTypeSaml       IdentityProviderUpdateResponseAccessPingoneType = "saml"
	IdentityProviderUpdateResponseAccessPingoneTypeCentrify   IdentityProviderUpdateResponseAccessPingoneType = "centrify"
	IdentityProviderUpdateResponseAccessPingoneTypeFacebook   IdentityProviderUpdateResponseAccessPingoneType = "facebook"
	IdentityProviderUpdateResponseAccessPingoneTypeGitHub     IdentityProviderUpdateResponseAccessPingoneType = "github"
	IdentityProviderUpdateResponseAccessPingoneTypeGoogleApps IdentityProviderUpdateResponseAccessPingoneType = "google-apps"
	IdentityProviderUpdateResponseAccessPingoneTypeGoogle     IdentityProviderUpdateResponseAccessPingoneType = "google"
	IdentityProviderUpdateResponseAccessPingoneTypeLinkedin   IdentityProviderUpdateResponseAccessPingoneType = "linkedin"
	IdentityProviderUpdateResponseAccessPingoneTypeOidc       IdentityProviderUpdateResponseAccessPingoneType = "oidc"
	IdentityProviderUpdateResponseAccessPingoneTypeOkta       IdentityProviderUpdateResponseAccessPingoneType = "okta"
	IdentityProviderUpdateResponseAccessPingoneTypeOnelogin   IdentityProviderUpdateResponseAccessPingoneType = "onelogin"
	IdentityProviderUpdateResponseAccessPingoneTypePingone    IdentityProviderUpdateResponseAccessPingoneType = "pingone"
	IdentityProviderUpdateResponseAccessPingoneTypeYandex     IdentityProviderUpdateResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateResponseAccessPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderUpdateResponseAccessPingoneScimConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessPingoneScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessPingoneScimConfig]
type identityProviderUpdateResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessPingoneScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderUpdateResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderUpdateResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessSamlJSON       `json:"-"`
}

// identityProviderUpdateResponseAccessSamlJSON contains the JSON metadata for the
// struct [IdentityProviderUpdateResponseAccessSaml]
type identityProviderUpdateResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessSamlJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessSaml) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []IdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                             `json:"sso_target_url"`
	JSON         identityProviderUpdateResponseAccessSamlConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessSamlConfigJSON contains the JSON metadata
// for the struct [IdentityProviderUpdateResponseAccessSamlConfig]
type identityProviderUpdateResponseAccessSamlConfigJSON struct {
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

func (r *IdentityProviderUpdateResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessSamlConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                            `json:"header_name"`
	JSON       identityProviderUpdateResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// identityProviderUpdateResponseAccessSamlConfigHeaderAttributeJSON contains the
// JSON metadata for the struct
// [IdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute]
type identityProviderUpdateResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessSamlConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessSamlType string

const (
	IdentityProviderUpdateResponseAccessSamlTypeOnetimepin IdentityProviderUpdateResponseAccessSamlType = "onetimepin"
	IdentityProviderUpdateResponseAccessSamlTypeAzureAd    IdentityProviderUpdateResponseAccessSamlType = "azureAD"
	IdentityProviderUpdateResponseAccessSamlTypeSaml       IdentityProviderUpdateResponseAccessSamlType = "saml"
	IdentityProviderUpdateResponseAccessSamlTypeCentrify   IdentityProviderUpdateResponseAccessSamlType = "centrify"
	IdentityProviderUpdateResponseAccessSamlTypeFacebook   IdentityProviderUpdateResponseAccessSamlType = "facebook"
	IdentityProviderUpdateResponseAccessSamlTypeGitHub     IdentityProviderUpdateResponseAccessSamlType = "github"
	IdentityProviderUpdateResponseAccessSamlTypeGoogleApps IdentityProviderUpdateResponseAccessSamlType = "google-apps"
	IdentityProviderUpdateResponseAccessSamlTypeGoogle     IdentityProviderUpdateResponseAccessSamlType = "google"
	IdentityProviderUpdateResponseAccessSamlTypeLinkedin   IdentityProviderUpdateResponseAccessSamlType = "linkedin"
	IdentityProviderUpdateResponseAccessSamlTypeOidc       IdentityProviderUpdateResponseAccessSamlType = "oidc"
	IdentityProviderUpdateResponseAccessSamlTypeOkta       IdentityProviderUpdateResponseAccessSamlType = "okta"
	IdentityProviderUpdateResponseAccessSamlTypeOnelogin   IdentityProviderUpdateResponseAccessSamlType = "onelogin"
	IdentityProviderUpdateResponseAccessSamlTypePingone    IdentityProviderUpdateResponseAccessSamlType = "pingone"
	IdentityProviderUpdateResponseAccessSamlTypeYandex     IdentityProviderUpdateResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateResponseAccessSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderUpdateResponseAccessSamlScimConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessSamlScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessSamlScimConfig]
type identityProviderUpdateResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessSamlScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderUpdateResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderUpdateResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessYandexJSON       `json:"-"`
}

// identityProviderUpdateResponseAccessYandexJSON contains the JSON metadata for
// the struct [IdentityProviderUpdateResponseAccessYandex]
type identityProviderUpdateResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessYandexJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessYandex) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                               `json:"client_secret"`
	JSON         identityProviderUpdateResponseAccessYandexConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessYandexConfigJSON contains the JSON metadata
// for the struct [IdentityProviderUpdateResponseAccessYandexConfig]
type identityProviderUpdateResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessYandexConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessYandexType string

const (
	IdentityProviderUpdateResponseAccessYandexTypeOnetimepin IdentityProviderUpdateResponseAccessYandexType = "onetimepin"
	IdentityProviderUpdateResponseAccessYandexTypeAzureAd    IdentityProviderUpdateResponseAccessYandexType = "azureAD"
	IdentityProviderUpdateResponseAccessYandexTypeSaml       IdentityProviderUpdateResponseAccessYandexType = "saml"
	IdentityProviderUpdateResponseAccessYandexTypeCentrify   IdentityProviderUpdateResponseAccessYandexType = "centrify"
	IdentityProviderUpdateResponseAccessYandexTypeFacebook   IdentityProviderUpdateResponseAccessYandexType = "facebook"
	IdentityProviderUpdateResponseAccessYandexTypeGitHub     IdentityProviderUpdateResponseAccessYandexType = "github"
	IdentityProviderUpdateResponseAccessYandexTypeGoogleApps IdentityProviderUpdateResponseAccessYandexType = "google-apps"
	IdentityProviderUpdateResponseAccessYandexTypeGoogle     IdentityProviderUpdateResponseAccessYandexType = "google"
	IdentityProviderUpdateResponseAccessYandexTypeLinkedin   IdentityProviderUpdateResponseAccessYandexType = "linkedin"
	IdentityProviderUpdateResponseAccessYandexTypeOidc       IdentityProviderUpdateResponseAccessYandexType = "oidc"
	IdentityProviderUpdateResponseAccessYandexTypeOkta       IdentityProviderUpdateResponseAccessYandexType = "okta"
	IdentityProviderUpdateResponseAccessYandexTypeOnelogin   IdentityProviderUpdateResponseAccessYandexType = "onelogin"
	IdentityProviderUpdateResponseAccessYandexTypePingone    IdentityProviderUpdateResponseAccessYandexType = "pingone"
	IdentityProviderUpdateResponseAccessYandexTypeYandex     IdentityProviderUpdateResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateResponseAccessYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderUpdateResponseAccessYandexScimConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessYandexScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderUpdateResponseAccessYandexScimConfig]
type identityProviderUpdateResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessYandexScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderUpdateResponseAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderUpdateResponseAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessOnetimepinJSON       `json:"-"`
}

// identityProviderUpdateResponseAccessOnetimepinJSON contains the JSON metadata
// for the struct [IdentityProviderUpdateResponseAccessOnetimepin]
type identityProviderUpdateResponseAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessOnetimepinJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessOnetimepin) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessOnetimepinType string

const (
	IdentityProviderUpdateResponseAccessOnetimepinTypeOnetimepin IdentityProviderUpdateResponseAccessOnetimepinType = "onetimepin"
	IdentityProviderUpdateResponseAccessOnetimepinTypeAzureAd    IdentityProviderUpdateResponseAccessOnetimepinType = "azureAD"
	IdentityProviderUpdateResponseAccessOnetimepinTypeSaml       IdentityProviderUpdateResponseAccessOnetimepinType = "saml"
	IdentityProviderUpdateResponseAccessOnetimepinTypeCentrify   IdentityProviderUpdateResponseAccessOnetimepinType = "centrify"
	IdentityProviderUpdateResponseAccessOnetimepinTypeFacebook   IdentityProviderUpdateResponseAccessOnetimepinType = "facebook"
	IdentityProviderUpdateResponseAccessOnetimepinTypeGitHub     IdentityProviderUpdateResponseAccessOnetimepinType = "github"
	IdentityProviderUpdateResponseAccessOnetimepinTypeGoogleApps IdentityProviderUpdateResponseAccessOnetimepinType = "google-apps"
	IdentityProviderUpdateResponseAccessOnetimepinTypeGoogle     IdentityProviderUpdateResponseAccessOnetimepinType = "google"
	IdentityProviderUpdateResponseAccessOnetimepinTypeLinkedin   IdentityProviderUpdateResponseAccessOnetimepinType = "linkedin"
	IdentityProviderUpdateResponseAccessOnetimepinTypeOidc       IdentityProviderUpdateResponseAccessOnetimepinType = "oidc"
	IdentityProviderUpdateResponseAccessOnetimepinTypeOkta       IdentityProviderUpdateResponseAccessOnetimepinType = "okta"
	IdentityProviderUpdateResponseAccessOnetimepinTypeOnelogin   IdentityProviderUpdateResponseAccessOnetimepinType = "onelogin"
	IdentityProviderUpdateResponseAccessOnetimepinTypePingone    IdentityProviderUpdateResponseAccessOnetimepinType = "pingone"
	IdentityProviderUpdateResponseAccessOnetimepinTypeYandex     IdentityProviderUpdateResponseAccessOnetimepinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateResponseAccessOnetimepinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderUpdateResponseAccessOnetimepinScimConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessOnetimepinScimConfigJSON contains the JSON
// metadata for the struct
// [IdentityProviderUpdateResponseAccessOnetimepinScimConfig]
type identityProviderUpdateResponseAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessOnetimepinScimConfigJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.IdentityProviderListResponseAccessAzureAd],
// [zero_trust.IdentityProviderListResponseAccessCentrify],
// [zero_trust.IdentityProviderListResponseAccessFacebook],
// [zero_trust.IdentityProviderListResponseAccessGitHub],
// [zero_trust.IdentityProviderListResponseAccessGoogle],
// [zero_trust.IdentityProviderListResponseAccessGoogleApps],
// [zero_trust.IdentityProviderListResponseAccessLinkedin],
// [zero_trust.IdentityProviderListResponseAccessOidc],
// [zero_trust.IdentityProviderListResponseAccessOkta],
// [zero_trust.IdentityProviderListResponseAccessOnelogin],
// [zero_trust.IdentityProviderListResponseAccessPingone],
// [zero_trust.IdentityProviderListResponseAccessSaml] or
// [zero_trust.IdentityProviderListResponseAccessYandex].
type IdentityProviderListResponse interface {
	implementsZeroTrustIdentityProviderListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IdentityProviderListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessAzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessCentrify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessFacebook{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessGitHub{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessGoogle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessGoogleApps{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessLinkedin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessOidc{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessOkta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessOnelogin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessPingone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessSaml{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessYandex{}),
		},
	)
}

type IdentityProviderListResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderListResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderListResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       identityProviderListResponseAccessAzureAdJSON       `json:"-"`
}

// identityProviderListResponseAccessAzureAdJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessAzureAd]
type identityProviderListResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessAzureAdJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessAzureAd) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessAzureAdConfig struct {
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
	SupportGroups bool                                                `json:"support_groups"`
	JSON          identityProviderListResponseAccessAzureAdConfigJSON `json:"-"`
}

// identityProviderListResponseAccessAzureAdConfigJSON contains the JSON metadata
// for the struct [IdentityProviderListResponseAccessAzureAdConfig]
type identityProviderListResponseAccessAzureAdConfigJSON struct {
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

func (r *IdentityProviderListResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessAzureAdConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessAzureAdType string

const (
	IdentityProviderListResponseAccessAzureAdTypeOnetimepin IdentityProviderListResponseAccessAzureAdType = "onetimepin"
	IdentityProviderListResponseAccessAzureAdTypeAzureAd    IdentityProviderListResponseAccessAzureAdType = "azureAD"
	IdentityProviderListResponseAccessAzureAdTypeSaml       IdentityProviderListResponseAccessAzureAdType = "saml"
	IdentityProviderListResponseAccessAzureAdTypeCentrify   IdentityProviderListResponseAccessAzureAdType = "centrify"
	IdentityProviderListResponseAccessAzureAdTypeFacebook   IdentityProviderListResponseAccessAzureAdType = "facebook"
	IdentityProviderListResponseAccessAzureAdTypeGitHub     IdentityProviderListResponseAccessAzureAdType = "github"
	IdentityProviderListResponseAccessAzureAdTypeGoogleApps IdentityProviderListResponseAccessAzureAdType = "google-apps"
	IdentityProviderListResponseAccessAzureAdTypeGoogle     IdentityProviderListResponseAccessAzureAdType = "google"
	IdentityProviderListResponseAccessAzureAdTypeLinkedin   IdentityProviderListResponseAccessAzureAdType = "linkedin"
	IdentityProviderListResponseAccessAzureAdTypeOidc       IdentityProviderListResponseAccessAzureAdType = "oidc"
	IdentityProviderListResponseAccessAzureAdTypeOkta       IdentityProviderListResponseAccessAzureAdType = "okta"
	IdentityProviderListResponseAccessAzureAdTypeOnelogin   IdentityProviderListResponseAccessAzureAdType = "onelogin"
	IdentityProviderListResponseAccessAzureAdTypePingone    IdentityProviderListResponseAccessAzureAdType = "pingone"
	IdentityProviderListResponseAccessAzureAdTypeYandex     IdentityProviderListResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderListResponseAccessAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderListResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// identityProviderListResponseAccessAzureAdScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderListResponseAccessAzureAdScimConfig]
type identityProviderListResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessAzureAdScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderListResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderListResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       identityProviderListResponseAccessCentrifyJSON       `json:"-"`
}

// identityProviderListResponseAccessCentrifyJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessCentrify]
type identityProviderListResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessCentrifyJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessCentrify) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessCentrifyConfig struct {
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
	EmailClaimName string                                               `json:"email_claim_name"`
	JSON           identityProviderListResponseAccessCentrifyConfigJSON `json:"-"`
}

// identityProviderListResponseAccessCentrifyConfigJSON contains the JSON metadata
// for the struct [IdentityProviderListResponseAccessCentrifyConfig]
type identityProviderListResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessCentrifyConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessCentrifyType string

const (
	IdentityProviderListResponseAccessCentrifyTypeOnetimepin IdentityProviderListResponseAccessCentrifyType = "onetimepin"
	IdentityProviderListResponseAccessCentrifyTypeAzureAd    IdentityProviderListResponseAccessCentrifyType = "azureAD"
	IdentityProviderListResponseAccessCentrifyTypeSaml       IdentityProviderListResponseAccessCentrifyType = "saml"
	IdentityProviderListResponseAccessCentrifyTypeCentrify   IdentityProviderListResponseAccessCentrifyType = "centrify"
	IdentityProviderListResponseAccessCentrifyTypeFacebook   IdentityProviderListResponseAccessCentrifyType = "facebook"
	IdentityProviderListResponseAccessCentrifyTypeGitHub     IdentityProviderListResponseAccessCentrifyType = "github"
	IdentityProviderListResponseAccessCentrifyTypeGoogleApps IdentityProviderListResponseAccessCentrifyType = "google-apps"
	IdentityProviderListResponseAccessCentrifyTypeGoogle     IdentityProviderListResponseAccessCentrifyType = "google"
	IdentityProviderListResponseAccessCentrifyTypeLinkedin   IdentityProviderListResponseAccessCentrifyType = "linkedin"
	IdentityProviderListResponseAccessCentrifyTypeOidc       IdentityProviderListResponseAccessCentrifyType = "oidc"
	IdentityProviderListResponseAccessCentrifyTypeOkta       IdentityProviderListResponseAccessCentrifyType = "okta"
	IdentityProviderListResponseAccessCentrifyTypeOnelogin   IdentityProviderListResponseAccessCentrifyType = "onelogin"
	IdentityProviderListResponseAccessCentrifyTypePingone    IdentityProviderListResponseAccessCentrifyType = "pingone"
	IdentityProviderListResponseAccessCentrifyTypeYandex     IdentityProviderListResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderListResponseAccessCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderListResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// identityProviderListResponseAccessCentrifyScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderListResponseAccessCentrifyScimConfig]
type identityProviderListResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessCentrifyScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderListResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderListResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       identityProviderListResponseAccessFacebookJSON       `json:"-"`
}

// identityProviderListResponseAccessFacebookJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessFacebook]
type identityProviderListResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessFacebookJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessFacebook) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                               `json:"client_secret"`
	JSON         identityProviderListResponseAccessFacebookConfigJSON `json:"-"`
}

// identityProviderListResponseAccessFacebookConfigJSON contains the JSON metadata
// for the struct [IdentityProviderListResponseAccessFacebookConfig]
type identityProviderListResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessFacebookConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessFacebookType string

const (
	IdentityProviderListResponseAccessFacebookTypeOnetimepin IdentityProviderListResponseAccessFacebookType = "onetimepin"
	IdentityProviderListResponseAccessFacebookTypeAzureAd    IdentityProviderListResponseAccessFacebookType = "azureAD"
	IdentityProviderListResponseAccessFacebookTypeSaml       IdentityProviderListResponseAccessFacebookType = "saml"
	IdentityProviderListResponseAccessFacebookTypeCentrify   IdentityProviderListResponseAccessFacebookType = "centrify"
	IdentityProviderListResponseAccessFacebookTypeFacebook   IdentityProviderListResponseAccessFacebookType = "facebook"
	IdentityProviderListResponseAccessFacebookTypeGitHub     IdentityProviderListResponseAccessFacebookType = "github"
	IdentityProviderListResponseAccessFacebookTypeGoogleApps IdentityProviderListResponseAccessFacebookType = "google-apps"
	IdentityProviderListResponseAccessFacebookTypeGoogle     IdentityProviderListResponseAccessFacebookType = "google"
	IdentityProviderListResponseAccessFacebookTypeLinkedin   IdentityProviderListResponseAccessFacebookType = "linkedin"
	IdentityProviderListResponseAccessFacebookTypeOidc       IdentityProviderListResponseAccessFacebookType = "oidc"
	IdentityProviderListResponseAccessFacebookTypeOkta       IdentityProviderListResponseAccessFacebookType = "okta"
	IdentityProviderListResponseAccessFacebookTypeOnelogin   IdentityProviderListResponseAccessFacebookType = "onelogin"
	IdentityProviderListResponseAccessFacebookTypePingone    IdentityProviderListResponseAccessFacebookType = "pingone"
	IdentityProviderListResponseAccessFacebookTypeYandex     IdentityProviderListResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderListResponseAccessFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderListResponseAccessFacebookScimConfigJSON `json:"-"`
}

// identityProviderListResponseAccessFacebookScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderListResponseAccessFacebookScimConfig]
type identityProviderListResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessFacebookScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderListResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderListResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       identityProviderListResponseAccessGitHubJSON       `json:"-"`
}

// identityProviderListResponseAccessGitHubJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessGitHub]
type identityProviderListResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessGitHubJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessGitHub) implementsZeroTrustIdentityProviderListResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                             `json:"client_secret"`
	JSON         identityProviderListResponseAccessGitHubConfigJSON `json:"-"`
}

// identityProviderListResponseAccessGitHubConfigJSON contains the JSON metadata
// for the struct [IdentityProviderListResponseAccessGitHubConfig]
type identityProviderListResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessGitHubConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessGitHubType string

const (
	IdentityProviderListResponseAccessGitHubTypeOnetimepin IdentityProviderListResponseAccessGitHubType = "onetimepin"
	IdentityProviderListResponseAccessGitHubTypeAzureAd    IdentityProviderListResponseAccessGitHubType = "azureAD"
	IdentityProviderListResponseAccessGitHubTypeSaml       IdentityProviderListResponseAccessGitHubType = "saml"
	IdentityProviderListResponseAccessGitHubTypeCentrify   IdentityProviderListResponseAccessGitHubType = "centrify"
	IdentityProviderListResponseAccessGitHubTypeFacebook   IdentityProviderListResponseAccessGitHubType = "facebook"
	IdentityProviderListResponseAccessGitHubTypeGitHub     IdentityProviderListResponseAccessGitHubType = "github"
	IdentityProviderListResponseAccessGitHubTypeGoogleApps IdentityProviderListResponseAccessGitHubType = "google-apps"
	IdentityProviderListResponseAccessGitHubTypeGoogle     IdentityProviderListResponseAccessGitHubType = "google"
	IdentityProviderListResponseAccessGitHubTypeLinkedin   IdentityProviderListResponseAccessGitHubType = "linkedin"
	IdentityProviderListResponseAccessGitHubTypeOidc       IdentityProviderListResponseAccessGitHubType = "oidc"
	IdentityProviderListResponseAccessGitHubTypeOkta       IdentityProviderListResponseAccessGitHubType = "okta"
	IdentityProviderListResponseAccessGitHubTypeOnelogin   IdentityProviderListResponseAccessGitHubType = "onelogin"
	IdentityProviderListResponseAccessGitHubTypePingone    IdentityProviderListResponseAccessGitHubType = "pingone"
	IdentityProviderListResponseAccessGitHubTypeYandex     IdentityProviderListResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderListResponseAccessGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderListResponseAccessGitHubScimConfigJSON `json:"-"`
}

// identityProviderListResponseAccessGitHubScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderListResponseAccessGitHubScimConfig]
type identityProviderListResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessGitHubScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderListResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderListResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       identityProviderListResponseAccessGoogleJSON       `json:"-"`
}

// identityProviderListResponseAccessGoogleJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessGoogle]
type identityProviderListResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessGoogleJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessGoogle) implementsZeroTrustIdentityProviderListResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                             `json:"email_claim_name"`
	JSON           identityProviderListResponseAccessGoogleConfigJSON `json:"-"`
}

// identityProviderListResponseAccessGoogleConfigJSON contains the JSON metadata
// for the struct [IdentityProviderListResponseAccessGoogleConfig]
type identityProviderListResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessGoogleConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessGoogleType string

const (
	IdentityProviderListResponseAccessGoogleTypeOnetimepin IdentityProviderListResponseAccessGoogleType = "onetimepin"
	IdentityProviderListResponseAccessGoogleTypeAzureAd    IdentityProviderListResponseAccessGoogleType = "azureAD"
	IdentityProviderListResponseAccessGoogleTypeSaml       IdentityProviderListResponseAccessGoogleType = "saml"
	IdentityProviderListResponseAccessGoogleTypeCentrify   IdentityProviderListResponseAccessGoogleType = "centrify"
	IdentityProviderListResponseAccessGoogleTypeFacebook   IdentityProviderListResponseAccessGoogleType = "facebook"
	IdentityProviderListResponseAccessGoogleTypeGitHub     IdentityProviderListResponseAccessGoogleType = "github"
	IdentityProviderListResponseAccessGoogleTypeGoogleApps IdentityProviderListResponseAccessGoogleType = "google-apps"
	IdentityProviderListResponseAccessGoogleTypeGoogle     IdentityProviderListResponseAccessGoogleType = "google"
	IdentityProviderListResponseAccessGoogleTypeLinkedin   IdentityProviderListResponseAccessGoogleType = "linkedin"
	IdentityProviderListResponseAccessGoogleTypeOidc       IdentityProviderListResponseAccessGoogleType = "oidc"
	IdentityProviderListResponseAccessGoogleTypeOkta       IdentityProviderListResponseAccessGoogleType = "okta"
	IdentityProviderListResponseAccessGoogleTypeOnelogin   IdentityProviderListResponseAccessGoogleType = "onelogin"
	IdentityProviderListResponseAccessGoogleTypePingone    IdentityProviderListResponseAccessGoogleType = "pingone"
	IdentityProviderListResponseAccessGoogleTypeYandex     IdentityProviderListResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderListResponseAccessGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderListResponseAccessGoogleScimConfigJSON `json:"-"`
}

// identityProviderListResponseAccessGoogleScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderListResponseAccessGoogleScimConfig]
type identityProviderListResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessGoogleScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderListResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderListResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       identityProviderListResponseAccessGoogleAppsJSON       `json:"-"`
}

// identityProviderListResponseAccessGoogleAppsJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessGoogleApps]
type identityProviderListResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessGoogleAppsJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessGoogleApps) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                 `json:"email_claim_name"`
	JSON           identityProviderListResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// identityProviderListResponseAccessGoogleAppsConfigJSON contains the JSON
// metadata for the struct [IdentityProviderListResponseAccessGoogleAppsConfig]
type identityProviderListResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessGoogleAppsConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessGoogleAppsType string

const (
	IdentityProviderListResponseAccessGoogleAppsTypeOnetimepin IdentityProviderListResponseAccessGoogleAppsType = "onetimepin"
	IdentityProviderListResponseAccessGoogleAppsTypeAzureAd    IdentityProviderListResponseAccessGoogleAppsType = "azureAD"
	IdentityProviderListResponseAccessGoogleAppsTypeSaml       IdentityProviderListResponseAccessGoogleAppsType = "saml"
	IdentityProviderListResponseAccessGoogleAppsTypeCentrify   IdentityProviderListResponseAccessGoogleAppsType = "centrify"
	IdentityProviderListResponseAccessGoogleAppsTypeFacebook   IdentityProviderListResponseAccessGoogleAppsType = "facebook"
	IdentityProviderListResponseAccessGoogleAppsTypeGitHub     IdentityProviderListResponseAccessGoogleAppsType = "github"
	IdentityProviderListResponseAccessGoogleAppsTypeGoogleApps IdentityProviderListResponseAccessGoogleAppsType = "google-apps"
	IdentityProviderListResponseAccessGoogleAppsTypeGoogle     IdentityProviderListResponseAccessGoogleAppsType = "google"
	IdentityProviderListResponseAccessGoogleAppsTypeLinkedin   IdentityProviderListResponseAccessGoogleAppsType = "linkedin"
	IdentityProviderListResponseAccessGoogleAppsTypeOidc       IdentityProviderListResponseAccessGoogleAppsType = "oidc"
	IdentityProviderListResponseAccessGoogleAppsTypeOkta       IdentityProviderListResponseAccessGoogleAppsType = "okta"
	IdentityProviderListResponseAccessGoogleAppsTypeOnelogin   IdentityProviderListResponseAccessGoogleAppsType = "onelogin"
	IdentityProviderListResponseAccessGoogleAppsTypePingone    IdentityProviderListResponseAccessGoogleAppsType = "pingone"
	IdentityProviderListResponseAccessGoogleAppsTypeYandex     IdentityProviderListResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderListResponseAccessGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderListResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// identityProviderListResponseAccessGoogleAppsScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderListResponseAccessGoogleAppsScimConfig]
type identityProviderListResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessGoogleAppsScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderListResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderListResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       identityProviderListResponseAccessLinkedinJSON       `json:"-"`
}

// identityProviderListResponseAccessLinkedinJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessLinkedin]
type identityProviderListResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessLinkedinJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessLinkedin) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                               `json:"client_secret"`
	JSON         identityProviderListResponseAccessLinkedinConfigJSON `json:"-"`
}

// identityProviderListResponseAccessLinkedinConfigJSON contains the JSON metadata
// for the struct [IdentityProviderListResponseAccessLinkedinConfig]
type identityProviderListResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessLinkedinConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessLinkedinType string

const (
	IdentityProviderListResponseAccessLinkedinTypeOnetimepin IdentityProviderListResponseAccessLinkedinType = "onetimepin"
	IdentityProviderListResponseAccessLinkedinTypeAzureAd    IdentityProviderListResponseAccessLinkedinType = "azureAD"
	IdentityProviderListResponseAccessLinkedinTypeSaml       IdentityProviderListResponseAccessLinkedinType = "saml"
	IdentityProviderListResponseAccessLinkedinTypeCentrify   IdentityProviderListResponseAccessLinkedinType = "centrify"
	IdentityProviderListResponseAccessLinkedinTypeFacebook   IdentityProviderListResponseAccessLinkedinType = "facebook"
	IdentityProviderListResponseAccessLinkedinTypeGitHub     IdentityProviderListResponseAccessLinkedinType = "github"
	IdentityProviderListResponseAccessLinkedinTypeGoogleApps IdentityProviderListResponseAccessLinkedinType = "google-apps"
	IdentityProviderListResponseAccessLinkedinTypeGoogle     IdentityProviderListResponseAccessLinkedinType = "google"
	IdentityProviderListResponseAccessLinkedinTypeLinkedin   IdentityProviderListResponseAccessLinkedinType = "linkedin"
	IdentityProviderListResponseAccessLinkedinTypeOidc       IdentityProviderListResponseAccessLinkedinType = "oidc"
	IdentityProviderListResponseAccessLinkedinTypeOkta       IdentityProviderListResponseAccessLinkedinType = "okta"
	IdentityProviderListResponseAccessLinkedinTypeOnelogin   IdentityProviderListResponseAccessLinkedinType = "onelogin"
	IdentityProviderListResponseAccessLinkedinTypePingone    IdentityProviderListResponseAccessLinkedinType = "pingone"
	IdentityProviderListResponseAccessLinkedinTypeYandex     IdentityProviderListResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderListResponseAccessLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderListResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// identityProviderListResponseAccessLinkedinScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderListResponseAccessLinkedinScimConfig]
type identityProviderListResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessLinkedinScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderListResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderListResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       identityProviderListResponseAccessOidcJSON       `json:"-"`
}

// identityProviderListResponseAccessOidcJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessOidc]
type identityProviderListResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessOidcJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessOidc) implementsZeroTrustIdentityProviderListResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessOidcConfig struct {
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
	TokenURL string                                           `json:"token_url"`
	JSON     identityProviderListResponseAccessOidcConfigJSON `json:"-"`
}

// identityProviderListResponseAccessOidcConfigJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessOidcConfig]
type identityProviderListResponseAccessOidcConfigJSON struct {
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

func (r *IdentityProviderListResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessOidcConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessOidcType string

const (
	IdentityProviderListResponseAccessOidcTypeOnetimepin IdentityProviderListResponseAccessOidcType = "onetimepin"
	IdentityProviderListResponseAccessOidcTypeAzureAd    IdentityProviderListResponseAccessOidcType = "azureAD"
	IdentityProviderListResponseAccessOidcTypeSaml       IdentityProviderListResponseAccessOidcType = "saml"
	IdentityProviderListResponseAccessOidcTypeCentrify   IdentityProviderListResponseAccessOidcType = "centrify"
	IdentityProviderListResponseAccessOidcTypeFacebook   IdentityProviderListResponseAccessOidcType = "facebook"
	IdentityProviderListResponseAccessOidcTypeGitHub     IdentityProviderListResponseAccessOidcType = "github"
	IdentityProviderListResponseAccessOidcTypeGoogleApps IdentityProviderListResponseAccessOidcType = "google-apps"
	IdentityProviderListResponseAccessOidcTypeGoogle     IdentityProviderListResponseAccessOidcType = "google"
	IdentityProviderListResponseAccessOidcTypeLinkedin   IdentityProviderListResponseAccessOidcType = "linkedin"
	IdentityProviderListResponseAccessOidcTypeOidc       IdentityProviderListResponseAccessOidcType = "oidc"
	IdentityProviderListResponseAccessOidcTypeOkta       IdentityProviderListResponseAccessOidcType = "okta"
	IdentityProviderListResponseAccessOidcTypeOnelogin   IdentityProviderListResponseAccessOidcType = "onelogin"
	IdentityProviderListResponseAccessOidcTypePingone    IdentityProviderListResponseAccessOidcType = "pingone"
	IdentityProviderListResponseAccessOidcTypeYandex     IdentityProviderListResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderListResponseAccessOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderListResponseAccessOidcScimConfigJSON `json:"-"`
}

// identityProviderListResponseAccessOidcScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderListResponseAccessOidcScimConfig]
type identityProviderListResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessOidcScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderListResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderListResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       identityProviderListResponseAccessOktaJSON       `json:"-"`
}

// identityProviderListResponseAccessOktaJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessOkta]
type identityProviderListResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessOktaJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessOkta) implementsZeroTrustIdentityProviderListResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessOktaConfig struct {
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
	OktaAccount string                                           `json:"okta_account"`
	JSON        identityProviderListResponseAccessOktaConfigJSON `json:"-"`
}

// identityProviderListResponseAccessOktaConfigJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessOktaConfig]
type identityProviderListResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessOktaConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessOktaType string

const (
	IdentityProviderListResponseAccessOktaTypeOnetimepin IdentityProviderListResponseAccessOktaType = "onetimepin"
	IdentityProviderListResponseAccessOktaTypeAzureAd    IdentityProviderListResponseAccessOktaType = "azureAD"
	IdentityProviderListResponseAccessOktaTypeSaml       IdentityProviderListResponseAccessOktaType = "saml"
	IdentityProviderListResponseAccessOktaTypeCentrify   IdentityProviderListResponseAccessOktaType = "centrify"
	IdentityProviderListResponseAccessOktaTypeFacebook   IdentityProviderListResponseAccessOktaType = "facebook"
	IdentityProviderListResponseAccessOktaTypeGitHub     IdentityProviderListResponseAccessOktaType = "github"
	IdentityProviderListResponseAccessOktaTypeGoogleApps IdentityProviderListResponseAccessOktaType = "google-apps"
	IdentityProviderListResponseAccessOktaTypeGoogle     IdentityProviderListResponseAccessOktaType = "google"
	IdentityProviderListResponseAccessOktaTypeLinkedin   IdentityProviderListResponseAccessOktaType = "linkedin"
	IdentityProviderListResponseAccessOktaTypeOidc       IdentityProviderListResponseAccessOktaType = "oidc"
	IdentityProviderListResponseAccessOktaTypeOkta       IdentityProviderListResponseAccessOktaType = "okta"
	IdentityProviderListResponseAccessOktaTypeOnelogin   IdentityProviderListResponseAccessOktaType = "onelogin"
	IdentityProviderListResponseAccessOktaTypePingone    IdentityProviderListResponseAccessOktaType = "pingone"
	IdentityProviderListResponseAccessOktaTypeYandex     IdentityProviderListResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderListResponseAccessOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderListResponseAccessOktaScimConfigJSON `json:"-"`
}

// identityProviderListResponseAccessOktaScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderListResponseAccessOktaScimConfig]
type identityProviderListResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessOktaScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderListResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderListResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       identityProviderListResponseAccessOneloginJSON       `json:"-"`
}

// identityProviderListResponseAccessOneloginJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessOnelogin]
type identityProviderListResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessOneloginJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessOnelogin) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                               `json:"onelogin_account"`
	JSON            identityProviderListResponseAccessOneloginConfigJSON `json:"-"`
}

// identityProviderListResponseAccessOneloginConfigJSON contains the JSON metadata
// for the struct [IdentityProviderListResponseAccessOneloginConfig]
type identityProviderListResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessOneloginConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessOneloginType string

const (
	IdentityProviderListResponseAccessOneloginTypeOnetimepin IdentityProviderListResponseAccessOneloginType = "onetimepin"
	IdentityProviderListResponseAccessOneloginTypeAzureAd    IdentityProviderListResponseAccessOneloginType = "azureAD"
	IdentityProviderListResponseAccessOneloginTypeSaml       IdentityProviderListResponseAccessOneloginType = "saml"
	IdentityProviderListResponseAccessOneloginTypeCentrify   IdentityProviderListResponseAccessOneloginType = "centrify"
	IdentityProviderListResponseAccessOneloginTypeFacebook   IdentityProviderListResponseAccessOneloginType = "facebook"
	IdentityProviderListResponseAccessOneloginTypeGitHub     IdentityProviderListResponseAccessOneloginType = "github"
	IdentityProviderListResponseAccessOneloginTypeGoogleApps IdentityProviderListResponseAccessOneloginType = "google-apps"
	IdentityProviderListResponseAccessOneloginTypeGoogle     IdentityProviderListResponseAccessOneloginType = "google"
	IdentityProviderListResponseAccessOneloginTypeLinkedin   IdentityProviderListResponseAccessOneloginType = "linkedin"
	IdentityProviderListResponseAccessOneloginTypeOidc       IdentityProviderListResponseAccessOneloginType = "oidc"
	IdentityProviderListResponseAccessOneloginTypeOkta       IdentityProviderListResponseAccessOneloginType = "okta"
	IdentityProviderListResponseAccessOneloginTypeOnelogin   IdentityProviderListResponseAccessOneloginType = "onelogin"
	IdentityProviderListResponseAccessOneloginTypePingone    IdentityProviderListResponseAccessOneloginType = "pingone"
	IdentityProviderListResponseAccessOneloginTypeYandex     IdentityProviderListResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderListResponseAccessOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderListResponseAccessOneloginScimConfigJSON `json:"-"`
}

// identityProviderListResponseAccessOneloginScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderListResponseAccessOneloginScimConfig]
type identityProviderListResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessOneloginScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderListResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderListResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       identityProviderListResponseAccessPingoneJSON       `json:"-"`
}

// identityProviderListResponseAccessPingoneJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessPingone]
type identityProviderListResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessPingoneJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessPingone) implementsZeroTrustIdentityProviderListResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                              `json:"ping_env_id"`
	JSON      identityProviderListResponseAccessPingoneConfigJSON `json:"-"`
}

// identityProviderListResponseAccessPingoneConfigJSON contains the JSON metadata
// for the struct [IdentityProviderListResponseAccessPingoneConfig]
type identityProviderListResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessPingoneConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessPingoneType string

const (
	IdentityProviderListResponseAccessPingoneTypeOnetimepin IdentityProviderListResponseAccessPingoneType = "onetimepin"
	IdentityProviderListResponseAccessPingoneTypeAzureAd    IdentityProviderListResponseAccessPingoneType = "azureAD"
	IdentityProviderListResponseAccessPingoneTypeSaml       IdentityProviderListResponseAccessPingoneType = "saml"
	IdentityProviderListResponseAccessPingoneTypeCentrify   IdentityProviderListResponseAccessPingoneType = "centrify"
	IdentityProviderListResponseAccessPingoneTypeFacebook   IdentityProviderListResponseAccessPingoneType = "facebook"
	IdentityProviderListResponseAccessPingoneTypeGitHub     IdentityProviderListResponseAccessPingoneType = "github"
	IdentityProviderListResponseAccessPingoneTypeGoogleApps IdentityProviderListResponseAccessPingoneType = "google-apps"
	IdentityProviderListResponseAccessPingoneTypeGoogle     IdentityProviderListResponseAccessPingoneType = "google"
	IdentityProviderListResponseAccessPingoneTypeLinkedin   IdentityProviderListResponseAccessPingoneType = "linkedin"
	IdentityProviderListResponseAccessPingoneTypeOidc       IdentityProviderListResponseAccessPingoneType = "oidc"
	IdentityProviderListResponseAccessPingoneTypeOkta       IdentityProviderListResponseAccessPingoneType = "okta"
	IdentityProviderListResponseAccessPingoneTypeOnelogin   IdentityProviderListResponseAccessPingoneType = "onelogin"
	IdentityProviderListResponseAccessPingoneTypePingone    IdentityProviderListResponseAccessPingoneType = "pingone"
	IdentityProviderListResponseAccessPingoneTypeYandex     IdentityProviderListResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderListResponseAccessPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderListResponseAccessPingoneScimConfigJSON `json:"-"`
}

// identityProviderListResponseAccessPingoneScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderListResponseAccessPingoneScimConfig]
type identityProviderListResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessPingoneScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderListResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderListResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       identityProviderListResponseAccessSamlJSON       `json:"-"`
}

// identityProviderListResponseAccessSamlJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessSaml]
type identityProviderListResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessSamlJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessSaml) implementsZeroTrustIdentityProviderListResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []IdentityProviderListResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                           `json:"sso_target_url"`
	JSON         identityProviderListResponseAccessSamlConfigJSON `json:"-"`
}

// identityProviderListResponseAccessSamlConfigJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessSamlConfig]
type identityProviderListResponseAccessSamlConfigJSON struct {
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

func (r *IdentityProviderListResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessSamlConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                          `json:"header_name"`
	JSON       identityProviderListResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// identityProviderListResponseAccessSamlConfigHeaderAttributeJSON contains the
// JSON metadata for the struct
// [IdentityProviderListResponseAccessSamlConfigHeaderAttribute]
type identityProviderListResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessSamlConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessSamlType string

const (
	IdentityProviderListResponseAccessSamlTypeOnetimepin IdentityProviderListResponseAccessSamlType = "onetimepin"
	IdentityProviderListResponseAccessSamlTypeAzureAd    IdentityProviderListResponseAccessSamlType = "azureAD"
	IdentityProviderListResponseAccessSamlTypeSaml       IdentityProviderListResponseAccessSamlType = "saml"
	IdentityProviderListResponseAccessSamlTypeCentrify   IdentityProviderListResponseAccessSamlType = "centrify"
	IdentityProviderListResponseAccessSamlTypeFacebook   IdentityProviderListResponseAccessSamlType = "facebook"
	IdentityProviderListResponseAccessSamlTypeGitHub     IdentityProviderListResponseAccessSamlType = "github"
	IdentityProviderListResponseAccessSamlTypeGoogleApps IdentityProviderListResponseAccessSamlType = "google-apps"
	IdentityProviderListResponseAccessSamlTypeGoogle     IdentityProviderListResponseAccessSamlType = "google"
	IdentityProviderListResponseAccessSamlTypeLinkedin   IdentityProviderListResponseAccessSamlType = "linkedin"
	IdentityProviderListResponseAccessSamlTypeOidc       IdentityProviderListResponseAccessSamlType = "oidc"
	IdentityProviderListResponseAccessSamlTypeOkta       IdentityProviderListResponseAccessSamlType = "okta"
	IdentityProviderListResponseAccessSamlTypeOnelogin   IdentityProviderListResponseAccessSamlType = "onelogin"
	IdentityProviderListResponseAccessSamlTypePingone    IdentityProviderListResponseAccessSamlType = "pingone"
	IdentityProviderListResponseAccessSamlTypeYandex     IdentityProviderListResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderListResponseAccessSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderListResponseAccessSamlScimConfigJSON `json:"-"`
}

// identityProviderListResponseAccessSamlScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderListResponseAccessSamlScimConfig]
type identityProviderListResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessSamlScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderListResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderListResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       identityProviderListResponseAccessYandexJSON       `json:"-"`
}

// identityProviderListResponseAccessYandexJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessYandex]
type identityProviderListResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessYandexJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessYandex) implementsZeroTrustIdentityProviderListResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                             `json:"client_secret"`
	JSON         identityProviderListResponseAccessYandexConfigJSON `json:"-"`
}

// identityProviderListResponseAccessYandexConfigJSON contains the JSON metadata
// for the struct [IdentityProviderListResponseAccessYandexConfig]
type identityProviderListResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessYandexConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessYandexType string

const (
	IdentityProviderListResponseAccessYandexTypeOnetimepin IdentityProviderListResponseAccessYandexType = "onetimepin"
	IdentityProviderListResponseAccessYandexTypeAzureAd    IdentityProviderListResponseAccessYandexType = "azureAD"
	IdentityProviderListResponseAccessYandexTypeSaml       IdentityProviderListResponseAccessYandexType = "saml"
	IdentityProviderListResponseAccessYandexTypeCentrify   IdentityProviderListResponseAccessYandexType = "centrify"
	IdentityProviderListResponseAccessYandexTypeFacebook   IdentityProviderListResponseAccessYandexType = "facebook"
	IdentityProviderListResponseAccessYandexTypeGitHub     IdentityProviderListResponseAccessYandexType = "github"
	IdentityProviderListResponseAccessYandexTypeGoogleApps IdentityProviderListResponseAccessYandexType = "google-apps"
	IdentityProviderListResponseAccessYandexTypeGoogle     IdentityProviderListResponseAccessYandexType = "google"
	IdentityProviderListResponseAccessYandexTypeLinkedin   IdentityProviderListResponseAccessYandexType = "linkedin"
	IdentityProviderListResponseAccessYandexTypeOidc       IdentityProviderListResponseAccessYandexType = "oidc"
	IdentityProviderListResponseAccessYandexTypeOkta       IdentityProviderListResponseAccessYandexType = "okta"
	IdentityProviderListResponseAccessYandexTypeOnelogin   IdentityProviderListResponseAccessYandexType = "onelogin"
	IdentityProviderListResponseAccessYandexTypePingone    IdentityProviderListResponseAccessYandexType = "pingone"
	IdentityProviderListResponseAccessYandexTypeYandex     IdentityProviderListResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderListResponseAccessYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderListResponseAccessYandexScimConfigJSON `json:"-"`
}

// identityProviderListResponseAccessYandexScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderListResponseAccessYandexScimConfig]
type identityProviderListResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessYandexScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderDeleteResponse struct {
	// UUID
	ID   string                             `json:"id"`
	JSON identityProviderDeleteResponseJSON `json:"-"`
}

// identityProviderDeleteResponseJSON contains the JSON metadata for the struct
// [IdentityProviderDeleteResponse]
type identityProviderDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.IdentityProviderGetResponseAccessAzureAd],
// [zero_trust.IdentityProviderGetResponseAccessCentrify],
// [zero_trust.IdentityProviderGetResponseAccessFacebook],
// [zero_trust.IdentityProviderGetResponseAccessGitHub],
// [zero_trust.IdentityProviderGetResponseAccessGoogle],
// [zero_trust.IdentityProviderGetResponseAccessGoogleApps],
// [zero_trust.IdentityProviderGetResponseAccessLinkedin],
// [zero_trust.IdentityProviderGetResponseAccessOidc],
// [zero_trust.IdentityProviderGetResponseAccessOkta],
// [zero_trust.IdentityProviderGetResponseAccessOnelogin],
// [zero_trust.IdentityProviderGetResponseAccessPingone],
// [zero_trust.IdentityProviderGetResponseAccessSaml],
// [zero_trust.IdentityProviderGetResponseAccessYandex] or
// [zero_trust.IdentityProviderGetResponseAccessOnetimepin].
type IdentityProviderGetResponse interface {
	implementsZeroTrustIdentityProviderGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IdentityProviderGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessAzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessCentrify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessFacebook{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessGitHub{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessGoogle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessGoogleApps{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessLinkedin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessOidc{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessOkta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessOnelogin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessPingone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessSaml{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessYandex{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessOnetimepin{}),
		},
	)
}

type IdentityProviderGetResponseAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderGetResponseAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderGetResponseAccessAzureAdScimConfig `json:"scim_config"`
	JSON       identityProviderGetResponseAccessAzureAdJSON       `json:"-"`
}

// identityProviderGetResponseAccessAzureAdJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseAccessAzureAd]
type identityProviderGetResponseAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessAzureAdJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessAzureAd) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessAzureAdConfig struct {
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
	SupportGroups bool                                               `json:"support_groups"`
	JSON          identityProviderGetResponseAccessAzureAdConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessAzureAdConfigJSON contains the JSON metadata
// for the struct [IdentityProviderGetResponseAccessAzureAdConfig]
type identityProviderGetResponseAccessAzureAdConfigJSON struct {
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

func (r *IdentityProviderGetResponseAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessAzureAdConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessAzureAdType string

const (
	IdentityProviderGetResponseAccessAzureAdTypeOnetimepin IdentityProviderGetResponseAccessAzureAdType = "onetimepin"
	IdentityProviderGetResponseAccessAzureAdTypeAzureAd    IdentityProviderGetResponseAccessAzureAdType = "azureAD"
	IdentityProviderGetResponseAccessAzureAdTypeSaml       IdentityProviderGetResponseAccessAzureAdType = "saml"
	IdentityProviderGetResponseAccessAzureAdTypeCentrify   IdentityProviderGetResponseAccessAzureAdType = "centrify"
	IdentityProviderGetResponseAccessAzureAdTypeFacebook   IdentityProviderGetResponseAccessAzureAdType = "facebook"
	IdentityProviderGetResponseAccessAzureAdTypeGitHub     IdentityProviderGetResponseAccessAzureAdType = "github"
	IdentityProviderGetResponseAccessAzureAdTypeGoogleApps IdentityProviderGetResponseAccessAzureAdType = "google-apps"
	IdentityProviderGetResponseAccessAzureAdTypeGoogle     IdentityProviderGetResponseAccessAzureAdType = "google"
	IdentityProviderGetResponseAccessAzureAdTypeLinkedin   IdentityProviderGetResponseAccessAzureAdType = "linkedin"
	IdentityProviderGetResponseAccessAzureAdTypeOidc       IdentityProviderGetResponseAccessAzureAdType = "oidc"
	IdentityProviderGetResponseAccessAzureAdTypeOkta       IdentityProviderGetResponseAccessAzureAdType = "okta"
	IdentityProviderGetResponseAccessAzureAdTypeOnelogin   IdentityProviderGetResponseAccessAzureAdType = "onelogin"
	IdentityProviderGetResponseAccessAzureAdTypePingone    IdentityProviderGetResponseAccessAzureAdType = "pingone"
	IdentityProviderGetResponseAccessAzureAdTypeYandex     IdentityProviderGetResponseAccessAzureAdType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderGetResponseAccessAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderGetResponseAccessAzureAdScimConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessAzureAdScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderGetResponseAccessAzureAdScimConfig]
type identityProviderGetResponseAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessAzureAdScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderGetResponseAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderGetResponseAccessCentrifyScimConfig `json:"scim_config"`
	JSON       identityProviderGetResponseAccessCentrifyJSON       `json:"-"`
}

// identityProviderGetResponseAccessCentrifyJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseAccessCentrify]
type identityProviderGetResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessCentrifyJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessCentrify) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessCentrifyConfig struct {
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
	EmailClaimName string                                              `json:"email_claim_name"`
	JSON           identityProviderGetResponseAccessCentrifyConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessCentrifyConfigJSON contains the JSON metadata
// for the struct [IdentityProviderGetResponseAccessCentrifyConfig]
type identityProviderGetResponseAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessCentrifyConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessCentrifyType string

const (
	IdentityProviderGetResponseAccessCentrifyTypeOnetimepin IdentityProviderGetResponseAccessCentrifyType = "onetimepin"
	IdentityProviderGetResponseAccessCentrifyTypeAzureAd    IdentityProviderGetResponseAccessCentrifyType = "azureAD"
	IdentityProviderGetResponseAccessCentrifyTypeSaml       IdentityProviderGetResponseAccessCentrifyType = "saml"
	IdentityProviderGetResponseAccessCentrifyTypeCentrify   IdentityProviderGetResponseAccessCentrifyType = "centrify"
	IdentityProviderGetResponseAccessCentrifyTypeFacebook   IdentityProviderGetResponseAccessCentrifyType = "facebook"
	IdentityProviderGetResponseAccessCentrifyTypeGitHub     IdentityProviderGetResponseAccessCentrifyType = "github"
	IdentityProviderGetResponseAccessCentrifyTypeGoogleApps IdentityProviderGetResponseAccessCentrifyType = "google-apps"
	IdentityProviderGetResponseAccessCentrifyTypeGoogle     IdentityProviderGetResponseAccessCentrifyType = "google"
	IdentityProviderGetResponseAccessCentrifyTypeLinkedin   IdentityProviderGetResponseAccessCentrifyType = "linkedin"
	IdentityProviderGetResponseAccessCentrifyTypeOidc       IdentityProviderGetResponseAccessCentrifyType = "oidc"
	IdentityProviderGetResponseAccessCentrifyTypeOkta       IdentityProviderGetResponseAccessCentrifyType = "okta"
	IdentityProviderGetResponseAccessCentrifyTypeOnelogin   IdentityProviderGetResponseAccessCentrifyType = "onelogin"
	IdentityProviderGetResponseAccessCentrifyTypePingone    IdentityProviderGetResponseAccessCentrifyType = "pingone"
	IdentityProviderGetResponseAccessCentrifyTypeYandex     IdentityProviderGetResponseAccessCentrifyType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderGetResponseAccessCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderGetResponseAccessCentrifyScimConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessCentrifyScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderGetResponseAccessCentrifyScimConfig]
type identityProviderGetResponseAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessCentrifyScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderGetResponseAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderGetResponseAccessFacebookScimConfig `json:"scim_config"`
	JSON       identityProviderGetResponseAccessFacebookJSON       `json:"-"`
}

// identityProviderGetResponseAccessFacebookJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseAccessFacebook]
type identityProviderGetResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessFacebookJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessFacebook) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                              `json:"client_secret"`
	JSON         identityProviderGetResponseAccessFacebookConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessFacebookConfigJSON contains the JSON metadata
// for the struct [IdentityProviderGetResponseAccessFacebookConfig]
type identityProviderGetResponseAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessFacebookConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessFacebookType string

const (
	IdentityProviderGetResponseAccessFacebookTypeOnetimepin IdentityProviderGetResponseAccessFacebookType = "onetimepin"
	IdentityProviderGetResponseAccessFacebookTypeAzureAd    IdentityProviderGetResponseAccessFacebookType = "azureAD"
	IdentityProviderGetResponseAccessFacebookTypeSaml       IdentityProviderGetResponseAccessFacebookType = "saml"
	IdentityProviderGetResponseAccessFacebookTypeCentrify   IdentityProviderGetResponseAccessFacebookType = "centrify"
	IdentityProviderGetResponseAccessFacebookTypeFacebook   IdentityProviderGetResponseAccessFacebookType = "facebook"
	IdentityProviderGetResponseAccessFacebookTypeGitHub     IdentityProviderGetResponseAccessFacebookType = "github"
	IdentityProviderGetResponseAccessFacebookTypeGoogleApps IdentityProviderGetResponseAccessFacebookType = "google-apps"
	IdentityProviderGetResponseAccessFacebookTypeGoogle     IdentityProviderGetResponseAccessFacebookType = "google"
	IdentityProviderGetResponseAccessFacebookTypeLinkedin   IdentityProviderGetResponseAccessFacebookType = "linkedin"
	IdentityProviderGetResponseAccessFacebookTypeOidc       IdentityProviderGetResponseAccessFacebookType = "oidc"
	IdentityProviderGetResponseAccessFacebookTypeOkta       IdentityProviderGetResponseAccessFacebookType = "okta"
	IdentityProviderGetResponseAccessFacebookTypeOnelogin   IdentityProviderGetResponseAccessFacebookType = "onelogin"
	IdentityProviderGetResponseAccessFacebookTypePingone    IdentityProviderGetResponseAccessFacebookType = "pingone"
	IdentityProviderGetResponseAccessFacebookTypeYandex     IdentityProviderGetResponseAccessFacebookType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderGetResponseAccessFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderGetResponseAccessFacebookScimConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessFacebookScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderGetResponseAccessFacebookScimConfig]
type identityProviderGetResponseAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessFacebookScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderGetResponseAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderGetResponseAccessGitHubScimConfig `json:"scim_config"`
	JSON       identityProviderGetResponseAccessGitHubJSON       `json:"-"`
}

// identityProviderGetResponseAccessGitHubJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseAccessGitHub]
type identityProviderGetResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessGitHubJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessGitHub) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                            `json:"client_secret"`
	JSON         identityProviderGetResponseAccessGitHubConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessGitHubConfigJSON contains the JSON metadata for
// the struct [IdentityProviderGetResponseAccessGitHubConfig]
type identityProviderGetResponseAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessGitHubConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessGitHubType string

const (
	IdentityProviderGetResponseAccessGitHubTypeOnetimepin IdentityProviderGetResponseAccessGitHubType = "onetimepin"
	IdentityProviderGetResponseAccessGitHubTypeAzureAd    IdentityProviderGetResponseAccessGitHubType = "azureAD"
	IdentityProviderGetResponseAccessGitHubTypeSaml       IdentityProviderGetResponseAccessGitHubType = "saml"
	IdentityProviderGetResponseAccessGitHubTypeCentrify   IdentityProviderGetResponseAccessGitHubType = "centrify"
	IdentityProviderGetResponseAccessGitHubTypeFacebook   IdentityProviderGetResponseAccessGitHubType = "facebook"
	IdentityProviderGetResponseAccessGitHubTypeGitHub     IdentityProviderGetResponseAccessGitHubType = "github"
	IdentityProviderGetResponseAccessGitHubTypeGoogleApps IdentityProviderGetResponseAccessGitHubType = "google-apps"
	IdentityProviderGetResponseAccessGitHubTypeGoogle     IdentityProviderGetResponseAccessGitHubType = "google"
	IdentityProviderGetResponseAccessGitHubTypeLinkedin   IdentityProviderGetResponseAccessGitHubType = "linkedin"
	IdentityProviderGetResponseAccessGitHubTypeOidc       IdentityProviderGetResponseAccessGitHubType = "oidc"
	IdentityProviderGetResponseAccessGitHubTypeOkta       IdentityProviderGetResponseAccessGitHubType = "okta"
	IdentityProviderGetResponseAccessGitHubTypeOnelogin   IdentityProviderGetResponseAccessGitHubType = "onelogin"
	IdentityProviderGetResponseAccessGitHubTypePingone    IdentityProviderGetResponseAccessGitHubType = "pingone"
	IdentityProviderGetResponseAccessGitHubTypeYandex     IdentityProviderGetResponseAccessGitHubType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderGetResponseAccessGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderGetResponseAccessGitHubScimConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessGitHubScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderGetResponseAccessGitHubScimConfig]
type identityProviderGetResponseAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessGitHubScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderGetResponseAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderGetResponseAccessGoogleScimConfig `json:"scim_config"`
	JSON       identityProviderGetResponseAccessGoogleJSON       `json:"-"`
}

// identityProviderGetResponseAccessGoogleJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseAccessGoogle]
type identityProviderGetResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessGoogleJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessGoogle) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                            `json:"email_claim_name"`
	JSON           identityProviderGetResponseAccessGoogleConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessGoogleConfigJSON contains the JSON metadata for
// the struct [IdentityProviderGetResponseAccessGoogleConfig]
type identityProviderGetResponseAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessGoogleConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessGoogleType string

const (
	IdentityProviderGetResponseAccessGoogleTypeOnetimepin IdentityProviderGetResponseAccessGoogleType = "onetimepin"
	IdentityProviderGetResponseAccessGoogleTypeAzureAd    IdentityProviderGetResponseAccessGoogleType = "azureAD"
	IdentityProviderGetResponseAccessGoogleTypeSaml       IdentityProviderGetResponseAccessGoogleType = "saml"
	IdentityProviderGetResponseAccessGoogleTypeCentrify   IdentityProviderGetResponseAccessGoogleType = "centrify"
	IdentityProviderGetResponseAccessGoogleTypeFacebook   IdentityProviderGetResponseAccessGoogleType = "facebook"
	IdentityProviderGetResponseAccessGoogleTypeGitHub     IdentityProviderGetResponseAccessGoogleType = "github"
	IdentityProviderGetResponseAccessGoogleTypeGoogleApps IdentityProviderGetResponseAccessGoogleType = "google-apps"
	IdentityProviderGetResponseAccessGoogleTypeGoogle     IdentityProviderGetResponseAccessGoogleType = "google"
	IdentityProviderGetResponseAccessGoogleTypeLinkedin   IdentityProviderGetResponseAccessGoogleType = "linkedin"
	IdentityProviderGetResponseAccessGoogleTypeOidc       IdentityProviderGetResponseAccessGoogleType = "oidc"
	IdentityProviderGetResponseAccessGoogleTypeOkta       IdentityProviderGetResponseAccessGoogleType = "okta"
	IdentityProviderGetResponseAccessGoogleTypeOnelogin   IdentityProviderGetResponseAccessGoogleType = "onelogin"
	IdentityProviderGetResponseAccessGoogleTypePingone    IdentityProviderGetResponseAccessGoogleType = "pingone"
	IdentityProviderGetResponseAccessGoogleTypeYandex     IdentityProviderGetResponseAccessGoogleType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderGetResponseAccessGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderGetResponseAccessGoogleScimConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessGoogleScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderGetResponseAccessGoogleScimConfig]
type identityProviderGetResponseAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessGoogleScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderGetResponseAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderGetResponseAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       identityProviderGetResponseAccessGoogleAppsJSON       `json:"-"`
}

// identityProviderGetResponseAccessGoogleAppsJSON contains the JSON metadata for
// the struct [IdentityProviderGetResponseAccessGoogleApps]
type identityProviderGetResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessGoogleAppsJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessGoogleApps) implementsZeroTrustIdentityProviderGetResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                                `json:"email_claim_name"`
	JSON           identityProviderGetResponseAccessGoogleAppsConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessGoogleAppsConfigJSON contains the JSON metadata
// for the struct [IdentityProviderGetResponseAccessGoogleAppsConfig]
type identityProviderGetResponseAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessGoogleAppsConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessGoogleAppsType string

const (
	IdentityProviderGetResponseAccessGoogleAppsTypeOnetimepin IdentityProviderGetResponseAccessGoogleAppsType = "onetimepin"
	IdentityProviderGetResponseAccessGoogleAppsTypeAzureAd    IdentityProviderGetResponseAccessGoogleAppsType = "azureAD"
	IdentityProviderGetResponseAccessGoogleAppsTypeSaml       IdentityProviderGetResponseAccessGoogleAppsType = "saml"
	IdentityProviderGetResponseAccessGoogleAppsTypeCentrify   IdentityProviderGetResponseAccessGoogleAppsType = "centrify"
	IdentityProviderGetResponseAccessGoogleAppsTypeFacebook   IdentityProviderGetResponseAccessGoogleAppsType = "facebook"
	IdentityProviderGetResponseAccessGoogleAppsTypeGitHub     IdentityProviderGetResponseAccessGoogleAppsType = "github"
	IdentityProviderGetResponseAccessGoogleAppsTypeGoogleApps IdentityProviderGetResponseAccessGoogleAppsType = "google-apps"
	IdentityProviderGetResponseAccessGoogleAppsTypeGoogle     IdentityProviderGetResponseAccessGoogleAppsType = "google"
	IdentityProviderGetResponseAccessGoogleAppsTypeLinkedin   IdentityProviderGetResponseAccessGoogleAppsType = "linkedin"
	IdentityProviderGetResponseAccessGoogleAppsTypeOidc       IdentityProviderGetResponseAccessGoogleAppsType = "oidc"
	IdentityProviderGetResponseAccessGoogleAppsTypeOkta       IdentityProviderGetResponseAccessGoogleAppsType = "okta"
	IdentityProviderGetResponseAccessGoogleAppsTypeOnelogin   IdentityProviderGetResponseAccessGoogleAppsType = "onelogin"
	IdentityProviderGetResponseAccessGoogleAppsTypePingone    IdentityProviderGetResponseAccessGoogleAppsType = "pingone"
	IdentityProviderGetResponseAccessGoogleAppsTypeYandex     IdentityProviderGetResponseAccessGoogleAppsType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderGetResponseAccessGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderGetResponseAccessGoogleAppsScimConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessGoogleAppsScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderGetResponseAccessGoogleAppsScimConfig]
type identityProviderGetResponseAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessGoogleAppsScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderGetResponseAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderGetResponseAccessLinkedinScimConfig `json:"scim_config"`
	JSON       identityProviderGetResponseAccessLinkedinJSON       `json:"-"`
}

// identityProviderGetResponseAccessLinkedinJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseAccessLinkedin]
type identityProviderGetResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessLinkedinJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessLinkedin) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                              `json:"client_secret"`
	JSON         identityProviderGetResponseAccessLinkedinConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessLinkedinConfigJSON contains the JSON metadata
// for the struct [IdentityProviderGetResponseAccessLinkedinConfig]
type identityProviderGetResponseAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessLinkedinConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessLinkedinType string

const (
	IdentityProviderGetResponseAccessLinkedinTypeOnetimepin IdentityProviderGetResponseAccessLinkedinType = "onetimepin"
	IdentityProviderGetResponseAccessLinkedinTypeAzureAd    IdentityProviderGetResponseAccessLinkedinType = "azureAD"
	IdentityProviderGetResponseAccessLinkedinTypeSaml       IdentityProviderGetResponseAccessLinkedinType = "saml"
	IdentityProviderGetResponseAccessLinkedinTypeCentrify   IdentityProviderGetResponseAccessLinkedinType = "centrify"
	IdentityProviderGetResponseAccessLinkedinTypeFacebook   IdentityProviderGetResponseAccessLinkedinType = "facebook"
	IdentityProviderGetResponseAccessLinkedinTypeGitHub     IdentityProviderGetResponseAccessLinkedinType = "github"
	IdentityProviderGetResponseAccessLinkedinTypeGoogleApps IdentityProviderGetResponseAccessLinkedinType = "google-apps"
	IdentityProviderGetResponseAccessLinkedinTypeGoogle     IdentityProviderGetResponseAccessLinkedinType = "google"
	IdentityProviderGetResponseAccessLinkedinTypeLinkedin   IdentityProviderGetResponseAccessLinkedinType = "linkedin"
	IdentityProviderGetResponseAccessLinkedinTypeOidc       IdentityProviderGetResponseAccessLinkedinType = "oidc"
	IdentityProviderGetResponseAccessLinkedinTypeOkta       IdentityProviderGetResponseAccessLinkedinType = "okta"
	IdentityProviderGetResponseAccessLinkedinTypeOnelogin   IdentityProviderGetResponseAccessLinkedinType = "onelogin"
	IdentityProviderGetResponseAccessLinkedinTypePingone    IdentityProviderGetResponseAccessLinkedinType = "pingone"
	IdentityProviderGetResponseAccessLinkedinTypeYandex     IdentityProviderGetResponseAccessLinkedinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderGetResponseAccessLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderGetResponseAccessLinkedinScimConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessLinkedinScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderGetResponseAccessLinkedinScimConfig]
type identityProviderGetResponseAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessLinkedinScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderGetResponseAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderGetResponseAccessOidcScimConfig `json:"scim_config"`
	JSON       identityProviderGetResponseAccessOidcJSON       `json:"-"`
}

// identityProviderGetResponseAccessOidcJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseAccessOidc]
type identityProviderGetResponseAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessOidcJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessOidc) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessOidcConfig struct {
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
	TokenURL string                                          `json:"token_url"`
	JSON     identityProviderGetResponseAccessOidcConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessOidcConfigJSON contains the JSON metadata for
// the struct [IdentityProviderGetResponseAccessOidcConfig]
type identityProviderGetResponseAccessOidcConfigJSON struct {
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

func (r *IdentityProviderGetResponseAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessOidcConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessOidcType string

const (
	IdentityProviderGetResponseAccessOidcTypeOnetimepin IdentityProviderGetResponseAccessOidcType = "onetimepin"
	IdentityProviderGetResponseAccessOidcTypeAzureAd    IdentityProviderGetResponseAccessOidcType = "azureAD"
	IdentityProviderGetResponseAccessOidcTypeSaml       IdentityProviderGetResponseAccessOidcType = "saml"
	IdentityProviderGetResponseAccessOidcTypeCentrify   IdentityProviderGetResponseAccessOidcType = "centrify"
	IdentityProviderGetResponseAccessOidcTypeFacebook   IdentityProviderGetResponseAccessOidcType = "facebook"
	IdentityProviderGetResponseAccessOidcTypeGitHub     IdentityProviderGetResponseAccessOidcType = "github"
	IdentityProviderGetResponseAccessOidcTypeGoogleApps IdentityProviderGetResponseAccessOidcType = "google-apps"
	IdentityProviderGetResponseAccessOidcTypeGoogle     IdentityProviderGetResponseAccessOidcType = "google"
	IdentityProviderGetResponseAccessOidcTypeLinkedin   IdentityProviderGetResponseAccessOidcType = "linkedin"
	IdentityProviderGetResponseAccessOidcTypeOidc       IdentityProviderGetResponseAccessOidcType = "oidc"
	IdentityProviderGetResponseAccessOidcTypeOkta       IdentityProviderGetResponseAccessOidcType = "okta"
	IdentityProviderGetResponseAccessOidcTypeOnelogin   IdentityProviderGetResponseAccessOidcType = "onelogin"
	IdentityProviderGetResponseAccessOidcTypePingone    IdentityProviderGetResponseAccessOidcType = "pingone"
	IdentityProviderGetResponseAccessOidcTypeYandex     IdentityProviderGetResponseAccessOidcType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderGetResponseAccessOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderGetResponseAccessOidcScimConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessOidcScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderGetResponseAccessOidcScimConfig]
type identityProviderGetResponseAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessOidcScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderGetResponseAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderGetResponseAccessOktaScimConfig `json:"scim_config"`
	JSON       identityProviderGetResponseAccessOktaJSON       `json:"-"`
}

// identityProviderGetResponseAccessOktaJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseAccessOkta]
type identityProviderGetResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessOktaJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessOkta) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessOktaConfig struct {
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
	OktaAccount string                                          `json:"okta_account"`
	JSON        identityProviderGetResponseAccessOktaConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessOktaConfigJSON contains the JSON metadata for
// the struct [IdentityProviderGetResponseAccessOktaConfig]
type identityProviderGetResponseAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessOktaConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessOktaType string

const (
	IdentityProviderGetResponseAccessOktaTypeOnetimepin IdentityProviderGetResponseAccessOktaType = "onetimepin"
	IdentityProviderGetResponseAccessOktaTypeAzureAd    IdentityProviderGetResponseAccessOktaType = "azureAD"
	IdentityProviderGetResponseAccessOktaTypeSaml       IdentityProviderGetResponseAccessOktaType = "saml"
	IdentityProviderGetResponseAccessOktaTypeCentrify   IdentityProviderGetResponseAccessOktaType = "centrify"
	IdentityProviderGetResponseAccessOktaTypeFacebook   IdentityProviderGetResponseAccessOktaType = "facebook"
	IdentityProviderGetResponseAccessOktaTypeGitHub     IdentityProviderGetResponseAccessOktaType = "github"
	IdentityProviderGetResponseAccessOktaTypeGoogleApps IdentityProviderGetResponseAccessOktaType = "google-apps"
	IdentityProviderGetResponseAccessOktaTypeGoogle     IdentityProviderGetResponseAccessOktaType = "google"
	IdentityProviderGetResponseAccessOktaTypeLinkedin   IdentityProviderGetResponseAccessOktaType = "linkedin"
	IdentityProviderGetResponseAccessOktaTypeOidc       IdentityProviderGetResponseAccessOktaType = "oidc"
	IdentityProviderGetResponseAccessOktaTypeOkta       IdentityProviderGetResponseAccessOktaType = "okta"
	IdentityProviderGetResponseAccessOktaTypeOnelogin   IdentityProviderGetResponseAccessOktaType = "onelogin"
	IdentityProviderGetResponseAccessOktaTypePingone    IdentityProviderGetResponseAccessOktaType = "pingone"
	IdentityProviderGetResponseAccessOktaTypeYandex     IdentityProviderGetResponseAccessOktaType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderGetResponseAccessOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderGetResponseAccessOktaScimConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessOktaScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderGetResponseAccessOktaScimConfig]
type identityProviderGetResponseAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessOktaScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderGetResponseAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderGetResponseAccessOneloginScimConfig `json:"scim_config"`
	JSON       identityProviderGetResponseAccessOneloginJSON       `json:"-"`
}

// identityProviderGetResponseAccessOneloginJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseAccessOnelogin]
type identityProviderGetResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessOneloginJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessOnelogin) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                              `json:"onelogin_account"`
	JSON            identityProviderGetResponseAccessOneloginConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessOneloginConfigJSON contains the JSON metadata
// for the struct [IdentityProviderGetResponseAccessOneloginConfig]
type identityProviderGetResponseAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessOneloginConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessOneloginType string

const (
	IdentityProviderGetResponseAccessOneloginTypeOnetimepin IdentityProviderGetResponseAccessOneloginType = "onetimepin"
	IdentityProviderGetResponseAccessOneloginTypeAzureAd    IdentityProviderGetResponseAccessOneloginType = "azureAD"
	IdentityProviderGetResponseAccessOneloginTypeSaml       IdentityProviderGetResponseAccessOneloginType = "saml"
	IdentityProviderGetResponseAccessOneloginTypeCentrify   IdentityProviderGetResponseAccessOneloginType = "centrify"
	IdentityProviderGetResponseAccessOneloginTypeFacebook   IdentityProviderGetResponseAccessOneloginType = "facebook"
	IdentityProviderGetResponseAccessOneloginTypeGitHub     IdentityProviderGetResponseAccessOneloginType = "github"
	IdentityProviderGetResponseAccessOneloginTypeGoogleApps IdentityProviderGetResponseAccessOneloginType = "google-apps"
	IdentityProviderGetResponseAccessOneloginTypeGoogle     IdentityProviderGetResponseAccessOneloginType = "google"
	IdentityProviderGetResponseAccessOneloginTypeLinkedin   IdentityProviderGetResponseAccessOneloginType = "linkedin"
	IdentityProviderGetResponseAccessOneloginTypeOidc       IdentityProviderGetResponseAccessOneloginType = "oidc"
	IdentityProviderGetResponseAccessOneloginTypeOkta       IdentityProviderGetResponseAccessOneloginType = "okta"
	IdentityProviderGetResponseAccessOneloginTypeOnelogin   IdentityProviderGetResponseAccessOneloginType = "onelogin"
	IdentityProviderGetResponseAccessOneloginTypePingone    IdentityProviderGetResponseAccessOneloginType = "pingone"
	IdentityProviderGetResponseAccessOneloginTypeYandex     IdentityProviderGetResponseAccessOneloginType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderGetResponseAccessOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderGetResponseAccessOneloginScimConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessOneloginScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderGetResponseAccessOneloginScimConfig]
type identityProviderGetResponseAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessOneloginScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderGetResponseAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderGetResponseAccessPingoneScimConfig `json:"scim_config"`
	JSON       identityProviderGetResponseAccessPingoneJSON       `json:"-"`
}

// identityProviderGetResponseAccessPingoneJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseAccessPingone]
type identityProviderGetResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessPingoneJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessPingone) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                             `json:"ping_env_id"`
	JSON      identityProviderGetResponseAccessPingoneConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessPingoneConfigJSON contains the JSON metadata
// for the struct [IdentityProviderGetResponseAccessPingoneConfig]
type identityProviderGetResponseAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessPingoneConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessPingoneType string

const (
	IdentityProviderGetResponseAccessPingoneTypeOnetimepin IdentityProviderGetResponseAccessPingoneType = "onetimepin"
	IdentityProviderGetResponseAccessPingoneTypeAzureAd    IdentityProviderGetResponseAccessPingoneType = "azureAD"
	IdentityProviderGetResponseAccessPingoneTypeSaml       IdentityProviderGetResponseAccessPingoneType = "saml"
	IdentityProviderGetResponseAccessPingoneTypeCentrify   IdentityProviderGetResponseAccessPingoneType = "centrify"
	IdentityProviderGetResponseAccessPingoneTypeFacebook   IdentityProviderGetResponseAccessPingoneType = "facebook"
	IdentityProviderGetResponseAccessPingoneTypeGitHub     IdentityProviderGetResponseAccessPingoneType = "github"
	IdentityProviderGetResponseAccessPingoneTypeGoogleApps IdentityProviderGetResponseAccessPingoneType = "google-apps"
	IdentityProviderGetResponseAccessPingoneTypeGoogle     IdentityProviderGetResponseAccessPingoneType = "google"
	IdentityProviderGetResponseAccessPingoneTypeLinkedin   IdentityProviderGetResponseAccessPingoneType = "linkedin"
	IdentityProviderGetResponseAccessPingoneTypeOidc       IdentityProviderGetResponseAccessPingoneType = "oidc"
	IdentityProviderGetResponseAccessPingoneTypeOkta       IdentityProviderGetResponseAccessPingoneType = "okta"
	IdentityProviderGetResponseAccessPingoneTypeOnelogin   IdentityProviderGetResponseAccessPingoneType = "onelogin"
	IdentityProviderGetResponseAccessPingoneTypePingone    IdentityProviderGetResponseAccessPingoneType = "pingone"
	IdentityProviderGetResponseAccessPingoneTypeYandex     IdentityProviderGetResponseAccessPingoneType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderGetResponseAccessPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderGetResponseAccessPingoneScimConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessPingoneScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderGetResponseAccessPingoneScimConfig]
type identityProviderGetResponseAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessPingoneScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderGetResponseAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderGetResponseAccessSamlScimConfig `json:"scim_config"`
	JSON       identityProviderGetResponseAccessSamlJSON       `json:"-"`
}

// identityProviderGetResponseAccessSamlJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseAccessSaml]
type identityProviderGetResponseAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessSamlJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessSaml) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []IdentityProviderGetResponseAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                          `json:"sso_target_url"`
	JSON         identityProviderGetResponseAccessSamlConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessSamlConfigJSON contains the JSON metadata for
// the struct [IdentityProviderGetResponseAccessSamlConfig]
type identityProviderGetResponseAccessSamlConfigJSON struct {
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

func (r *IdentityProviderGetResponseAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessSamlConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                         `json:"header_name"`
	JSON       identityProviderGetResponseAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// identityProviderGetResponseAccessSamlConfigHeaderAttributeJSON contains the JSON
// metadata for the struct
// [IdentityProviderGetResponseAccessSamlConfigHeaderAttribute]
type identityProviderGetResponseAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessSamlConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessSamlType string

const (
	IdentityProviderGetResponseAccessSamlTypeOnetimepin IdentityProviderGetResponseAccessSamlType = "onetimepin"
	IdentityProviderGetResponseAccessSamlTypeAzureAd    IdentityProviderGetResponseAccessSamlType = "azureAD"
	IdentityProviderGetResponseAccessSamlTypeSaml       IdentityProviderGetResponseAccessSamlType = "saml"
	IdentityProviderGetResponseAccessSamlTypeCentrify   IdentityProviderGetResponseAccessSamlType = "centrify"
	IdentityProviderGetResponseAccessSamlTypeFacebook   IdentityProviderGetResponseAccessSamlType = "facebook"
	IdentityProviderGetResponseAccessSamlTypeGitHub     IdentityProviderGetResponseAccessSamlType = "github"
	IdentityProviderGetResponseAccessSamlTypeGoogleApps IdentityProviderGetResponseAccessSamlType = "google-apps"
	IdentityProviderGetResponseAccessSamlTypeGoogle     IdentityProviderGetResponseAccessSamlType = "google"
	IdentityProviderGetResponseAccessSamlTypeLinkedin   IdentityProviderGetResponseAccessSamlType = "linkedin"
	IdentityProviderGetResponseAccessSamlTypeOidc       IdentityProviderGetResponseAccessSamlType = "oidc"
	IdentityProviderGetResponseAccessSamlTypeOkta       IdentityProviderGetResponseAccessSamlType = "okta"
	IdentityProviderGetResponseAccessSamlTypeOnelogin   IdentityProviderGetResponseAccessSamlType = "onelogin"
	IdentityProviderGetResponseAccessSamlTypePingone    IdentityProviderGetResponseAccessSamlType = "pingone"
	IdentityProviderGetResponseAccessSamlTypeYandex     IdentityProviderGetResponseAccessSamlType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderGetResponseAccessSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderGetResponseAccessSamlScimConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessSamlScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderGetResponseAccessSamlScimConfig]
type identityProviderGetResponseAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessSamlScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderGetResponseAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderGetResponseAccessYandexScimConfig `json:"scim_config"`
	JSON       identityProviderGetResponseAccessYandexJSON       `json:"-"`
}

// identityProviderGetResponseAccessYandexJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseAccessYandex]
type identityProviderGetResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessYandexJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessYandex) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                            `json:"client_secret"`
	JSON         identityProviderGetResponseAccessYandexConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessYandexConfigJSON contains the JSON metadata for
// the struct [IdentityProviderGetResponseAccessYandexConfig]
type identityProviderGetResponseAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessYandexConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessYandexType string

const (
	IdentityProviderGetResponseAccessYandexTypeOnetimepin IdentityProviderGetResponseAccessYandexType = "onetimepin"
	IdentityProviderGetResponseAccessYandexTypeAzureAd    IdentityProviderGetResponseAccessYandexType = "azureAD"
	IdentityProviderGetResponseAccessYandexTypeSaml       IdentityProviderGetResponseAccessYandexType = "saml"
	IdentityProviderGetResponseAccessYandexTypeCentrify   IdentityProviderGetResponseAccessYandexType = "centrify"
	IdentityProviderGetResponseAccessYandexTypeFacebook   IdentityProviderGetResponseAccessYandexType = "facebook"
	IdentityProviderGetResponseAccessYandexTypeGitHub     IdentityProviderGetResponseAccessYandexType = "github"
	IdentityProviderGetResponseAccessYandexTypeGoogleApps IdentityProviderGetResponseAccessYandexType = "google-apps"
	IdentityProviderGetResponseAccessYandexTypeGoogle     IdentityProviderGetResponseAccessYandexType = "google"
	IdentityProviderGetResponseAccessYandexTypeLinkedin   IdentityProviderGetResponseAccessYandexType = "linkedin"
	IdentityProviderGetResponseAccessYandexTypeOidc       IdentityProviderGetResponseAccessYandexType = "oidc"
	IdentityProviderGetResponseAccessYandexTypeOkta       IdentityProviderGetResponseAccessYandexType = "okta"
	IdentityProviderGetResponseAccessYandexTypeOnelogin   IdentityProviderGetResponseAccessYandexType = "onelogin"
	IdentityProviderGetResponseAccessYandexTypePingone    IdentityProviderGetResponseAccessYandexType = "pingone"
	IdentityProviderGetResponseAccessYandexTypeYandex     IdentityProviderGetResponseAccessYandexType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderGetResponseAccessYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderGetResponseAccessYandexScimConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessYandexScimConfigJSON contains the JSON metadata
// for the struct [IdentityProviderGetResponseAccessYandexScimConfig]
type identityProviderGetResponseAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessYandexScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderGetResponseAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProviderGetResponseAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       identityProviderGetResponseAccessOnetimepinJSON       `json:"-"`
}

// identityProviderGetResponseAccessOnetimepinJSON contains the JSON metadata for
// the struct [IdentityProviderGetResponseAccessOnetimepin]
type identityProviderGetResponseAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessOnetimepinJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessOnetimepin) implementsZeroTrustIdentityProviderGetResponse() {
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessOnetimepinType string

const (
	IdentityProviderGetResponseAccessOnetimepinTypeOnetimepin IdentityProviderGetResponseAccessOnetimepinType = "onetimepin"
	IdentityProviderGetResponseAccessOnetimepinTypeAzureAd    IdentityProviderGetResponseAccessOnetimepinType = "azureAD"
	IdentityProviderGetResponseAccessOnetimepinTypeSaml       IdentityProviderGetResponseAccessOnetimepinType = "saml"
	IdentityProviderGetResponseAccessOnetimepinTypeCentrify   IdentityProviderGetResponseAccessOnetimepinType = "centrify"
	IdentityProviderGetResponseAccessOnetimepinTypeFacebook   IdentityProviderGetResponseAccessOnetimepinType = "facebook"
	IdentityProviderGetResponseAccessOnetimepinTypeGitHub     IdentityProviderGetResponseAccessOnetimepinType = "github"
	IdentityProviderGetResponseAccessOnetimepinTypeGoogleApps IdentityProviderGetResponseAccessOnetimepinType = "google-apps"
	IdentityProviderGetResponseAccessOnetimepinTypeGoogle     IdentityProviderGetResponseAccessOnetimepinType = "google"
	IdentityProviderGetResponseAccessOnetimepinTypeLinkedin   IdentityProviderGetResponseAccessOnetimepinType = "linkedin"
	IdentityProviderGetResponseAccessOnetimepinTypeOidc       IdentityProviderGetResponseAccessOnetimepinType = "oidc"
	IdentityProviderGetResponseAccessOnetimepinTypeOkta       IdentityProviderGetResponseAccessOnetimepinType = "okta"
	IdentityProviderGetResponseAccessOnetimepinTypeOnelogin   IdentityProviderGetResponseAccessOnetimepinType = "onelogin"
	IdentityProviderGetResponseAccessOnetimepinTypePingone    IdentityProviderGetResponseAccessOnetimepinType = "pingone"
	IdentityProviderGetResponseAccessOnetimepinTypeYandex     IdentityProviderGetResponseAccessOnetimepinType = "yandex"
)

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderGetResponseAccessOnetimepinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProviderGetResponseAccessOnetimepinScimConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessOnetimepinScimConfigJSON contains the JSON
// metadata for the struct [IdentityProviderGetResponseAccessOnetimepinScimConfig]
type identityProviderGetResponseAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessOnetimepinScimConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewParams struct {
	Config param.Field[IdentityProviderNewParamsConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID     param.Field[string]                              `path:"zone_id"`
	ScimConfig param.Field[IdentityProviderNewParamsScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsConfig struct {
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
	HeaderAttributes param.Field[[]IdentityProviderNewParamsConfigHeaderAttribute] `json:"header_attributes"`
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

func (r IdentityProviderNewParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r IdentityProviderNewParamsConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsType string

const (
	IdentityProviderNewParamsTypeOnetimepin IdentityProviderNewParamsType = "onetimepin"
	IdentityProviderNewParamsTypeAzureAd    IdentityProviderNewParamsType = "azureAD"
	IdentityProviderNewParamsTypeSaml       IdentityProviderNewParamsType = "saml"
	IdentityProviderNewParamsTypeCentrify   IdentityProviderNewParamsType = "centrify"
	IdentityProviderNewParamsTypeFacebook   IdentityProviderNewParamsType = "facebook"
	IdentityProviderNewParamsTypeGitHub     IdentityProviderNewParamsType = "github"
	IdentityProviderNewParamsTypeGoogleApps IdentityProviderNewParamsType = "google-apps"
	IdentityProviderNewParamsTypeGoogle     IdentityProviderNewParamsType = "google"
	IdentityProviderNewParamsTypeLinkedin   IdentityProviderNewParamsType = "linkedin"
	IdentityProviderNewParamsTypeOidc       IdentityProviderNewParamsType = "oidc"
	IdentityProviderNewParamsTypeOkta       IdentityProviderNewParamsType = "okta"
	IdentityProviderNewParamsTypeOnelogin   IdentityProviderNewParamsType = "onelogin"
	IdentityProviderNewParamsTypePingone    IdentityProviderNewParamsType = "pingone"
	IdentityProviderNewParamsTypeYandex     IdentityProviderNewParamsType = "yandex"
)

type IdentityProviderNewParamsScimConfig struct {
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

func (r IdentityProviderNewParamsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewResponseEnvelope struct {
	Errors   []IdentityProviderNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IdentityProviderNewResponseEnvelopeMessages `json:"messages,required"`
	Result   IdentityProviderNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IdentityProviderNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    identityProviderNewResponseEnvelopeJSON    `json:"-"`
}

// identityProviderNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseEnvelope]
type identityProviderNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    identityProviderNewResponseEnvelopeErrorsJSON `json:"-"`
}

// identityProviderNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseEnvelopeErrors]
type identityProviderNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    identityProviderNewResponseEnvelopeMessagesJSON `json:"-"`
}

// identityProviderNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [IdentityProviderNewResponseEnvelopeMessages]
type identityProviderNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IdentityProviderNewResponseEnvelopeSuccess bool

const (
	IdentityProviderNewResponseEnvelopeSuccessTrue IdentityProviderNewResponseEnvelopeSuccess = true
)

type IdentityProviderUpdateParams struct {
	Config param.Field[IdentityProviderUpdateParamsConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID     param.Field[string]                                 `path:"zone_id"`
	ScimConfig param.Field[IdentityProviderUpdateParamsScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsConfig struct {
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
	HeaderAttributes param.Field[[]IdentityProviderUpdateParamsConfigHeaderAttribute] `json:"header_attributes"`
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

func (r IdentityProviderUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r IdentityProviderUpdateParamsConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsType string

const (
	IdentityProviderUpdateParamsTypeOnetimepin IdentityProviderUpdateParamsType = "onetimepin"
	IdentityProviderUpdateParamsTypeAzureAd    IdentityProviderUpdateParamsType = "azureAD"
	IdentityProviderUpdateParamsTypeSaml       IdentityProviderUpdateParamsType = "saml"
	IdentityProviderUpdateParamsTypeCentrify   IdentityProviderUpdateParamsType = "centrify"
	IdentityProviderUpdateParamsTypeFacebook   IdentityProviderUpdateParamsType = "facebook"
	IdentityProviderUpdateParamsTypeGitHub     IdentityProviderUpdateParamsType = "github"
	IdentityProviderUpdateParamsTypeGoogleApps IdentityProviderUpdateParamsType = "google-apps"
	IdentityProviderUpdateParamsTypeGoogle     IdentityProviderUpdateParamsType = "google"
	IdentityProviderUpdateParamsTypeLinkedin   IdentityProviderUpdateParamsType = "linkedin"
	IdentityProviderUpdateParamsTypeOidc       IdentityProviderUpdateParamsType = "oidc"
	IdentityProviderUpdateParamsTypeOkta       IdentityProviderUpdateParamsType = "okta"
	IdentityProviderUpdateParamsTypeOnelogin   IdentityProviderUpdateParamsType = "onelogin"
	IdentityProviderUpdateParamsTypePingone    IdentityProviderUpdateParamsType = "pingone"
	IdentityProviderUpdateParamsTypeYandex     IdentityProviderUpdateParamsType = "yandex"
)

type IdentityProviderUpdateParamsScimConfig struct {
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

func (r IdentityProviderUpdateParamsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateResponseEnvelope struct {
	Errors   []IdentityProviderUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IdentityProviderUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   IdentityProviderUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IdentityProviderUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    identityProviderUpdateResponseEnvelopeJSON    `json:"-"`
}

// identityProviderUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [IdentityProviderUpdateResponseEnvelope]
type identityProviderUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    identityProviderUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// identityProviderUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [IdentityProviderUpdateResponseEnvelopeErrors]
type identityProviderUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    identityProviderUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// identityProviderUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [IdentityProviderUpdateResponseEnvelopeMessages]
type identityProviderUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IdentityProviderUpdateResponseEnvelopeSuccess bool

const (
	IdentityProviderUpdateResponseEnvelopeSuccessTrue IdentityProviderUpdateResponseEnvelopeSuccess = true
)

type IdentityProviderListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type IdentityProviderListResponseEnvelope struct {
	Errors   []IdentityProviderListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IdentityProviderListResponseEnvelopeMessages `json:"messages,required"`
	Result   []IdentityProviderListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    IdentityProviderListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo IdentityProviderListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       identityProviderListResponseEnvelopeJSON       `json:"-"`
}

// identityProviderListResponseEnvelopeJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseEnvelope]
type identityProviderListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    identityProviderListResponseEnvelopeErrorsJSON `json:"-"`
}

// identityProviderListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseEnvelopeErrors]
type identityProviderListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    identityProviderListResponseEnvelopeMessagesJSON `json:"-"`
}

// identityProviderListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseEnvelopeMessages]
type identityProviderListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IdentityProviderListResponseEnvelopeSuccess bool

const (
	IdentityProviderListResponseEnvelopeSuccessTrue IdentityProviderListResponseEnvelopeSuccess = true
)

type IdentityProviderListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       identityProviderListResponseEnvelopeResultInfoJSON `json:"-"`
}

// identityProviderListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [IdentityProviderListResponseEnvelopeResultInfo]
type identityProviderListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type IdentityProviderDeleteResponseEnvelope struct {
	Errors   []IdentityProviderDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IdentityProviderDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   IdentityProviderDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IdentityProviderDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    identityProviderDeleteResponseEnvelopeJSON    `json:"-"`
}

// identityProviderDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [IdentityProviderDeleteResponseEnvelope]
type identityProviderDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderDeleteResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    identityProviderDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// identityProviderDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [IdentityProviderDeleteResponseEnvelopeErrors]
type identityProviderDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderDeleteResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    identityProviderDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// identityProviderDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [IdentityProviderDeleteResponseEnvelopeMessages]
type identityProviderDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IdentityProviderDeleteResponseEnvelopeSuccess bool

const (
	IdentityProviderDeleteResponseEnvelopeSuccessTrue IdentityProviderDeleteResponseEnvelopeSuccess = true
)

type IdentityProviderGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type IdentityProviderGetResponseEnvelope struct {
	Errors   []IdentityProviderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IdentityProviderGetResponseEnvelopeMessages `json:"messages,required"`
	Result   IdentityProviderGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IdentityProviderGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    identityProviderGetResponseEnvelopeJSON    `json:"-"`
}

// identityProviderGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseEnvelope]
type identityProviderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    identityProviderGetResponseEnvelopeErrorsJSON `json:"-"`
}

// identityProviderGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseEnvelopeErrors]
type identityProviderGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    identityProviderGetResponseEnvelopeMessagesJSON `json:"-"`
}

// identityProviderGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [IdentityProviderGetResponseEnvelopeMessages]
type identityProviderGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IdentityProviderGetResponseEnvelopeSuccess bool

const (
	IdentityProviderGetResponseEnvelopeSuccessTrue IdentityProviderGetResponseEnvelopeSuccess = true
)

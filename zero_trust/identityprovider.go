// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *IdentityProviderService) New(ctx context.Context, params IdentityProviderNewParams, opts ...option.RequestOption) (res *ZeroTrustIdentityProviders, err error) {
	opts = append(r.Options[:], opts...)
	var env IdentityProviderNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.getAccountID().Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.getAccountID()
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.getZoneID()
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
func (r *IdentityProviderService) Update(ctx context.Context, uuid string, params IdentityProviderUpdateParams, opts ...option.RequestOption) (res *ZeroTrustIdentityProviders, err error) {
	opts = append(r.Options[:], opts...)
	var env IdentityProviderUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.getAccountID().Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.getAccountID()
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.getZoneID()
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a configured identity provider.
func (r *IdentityProviderService) Get(ctx context.Context, uuid string, query IdentityProviderGetParams, opts ...option.RequestOption) (res *ZeroTrustIdentityProviders, err error) {
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [zero_trust.ZeroTrustIdentityProvidersAccessAzureAd],
// [zero_trust.ZeroTrustIdentityProvidersAccessCentrify],
// [zero_trust.ZeroTrustIdentityProvidersAccessFacebook],
// [zero_trust.ZeroTrustIdentityProvidersAccessGitHub],
// [zero_trust.ZeroTrustIdentityProvidersAccessGoogle],
// [zero_trust.ZeroTrustIdentityProvidersAccessGoogleApps],
// [zero_trust.ZeroTrustIdentityProvidersAccessLinkedin],
// [zero_trust.ZeroTrustIdentityProvidersAccessOidc],
// [zero_trust.ZeroTrustIdentityProvidersAccessOkta],
// [zero_trust.ZeroTrustIdentityProvidersAccessOnelogin],
// [zero_trust.ZeroTrustIdentityProvidersAccessPingone],
// [zero_trust.ZeroTrustIdentityProvidersAccessSaml],
// [zero_trust.ZeroTrustIdentityProvidersAccessYandex] or
// [zero_trust.ZeroTrustIdentityProvidersAccessOnetimepin].
type ZeroTrustIdentityProviders interface {
	implementsZeroTrustZeroTrustIdentityProviders()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustIdentityProviders)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProvidersAccessAzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProvidersAccessCentrify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProvidersAccessFacebook{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProvidersAccessGitHub{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProvidersAccessGoogle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProvidersAccessGoogleApps{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProvidersAccessLinkedin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProvidersAccessOidc{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProvidersAccessOkta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProvidersAccessOnelogin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProvidersAccessPingone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProvidersAccessSaml{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProvidersAccessYandex{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ZeroTrustIdentityProvidersAccessOnetimepin{}),
		},
	)
}

type ZeroTrustIdentityProvidersAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProvidersAccessAzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProvidersAccessAzureAdType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProvidersAccessAzureAdScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProvidersAccessAzureAdJSON       `json:"-"`
}

// zeroTrustIdentityProvidersAccessAzureAdJSON contains the JSON metadata for the
// struct [ZeroTrustIdentityProvidersAccessAzureAd]
type zeroTrustIdentityProvidersAccessAzureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessAzureAdJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProvidersAccessAzureAd) implementsZeroTrustZeroTrustIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessAzureAdConfig struct {
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
	// Indicates the type of user interaction that is required. prompt=login forces the
	// user to enter their credentials on that request, negating single-sign on.
	// prompt=none is the opposite. It ensures that the user isn't presented with any
	// interactive prompt. If the request can't be completed silently by using
	// single-sign on, the Microsoft identity platform returns an interaction_required
	// error. prompt=select_account interrupts single sign-on providing account
	// selection experience listing all the accounts either in session or any
	// remembered account or an option to choose to use a different account altogether.
	Prompt ZeroTrustIdentityProvidersAccessAzureAdConfigPrompt `json:"prompt"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool                                              `json:"support_groups"`
	JSON          zeroTrustIdentityProvidersAccessAzureAdConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessAzureAdConfigJSON contains the JSON metadata for
// the struct [ZeroTrustIdentityProvidersAccessAzureAdConfig]
type zeroTrustIdentityProvidersAccessAzureAdConfigJSON struct {
	Claims                   apijson.Field
	ClientID                 apijson.Field
	ClientSecret             apijson.Field
	ConditionalAccessEnabled apijson.Field
	DirectoryID              apijson.Field
	EmailClaimName           apijson.Field
	Prompt                   apijson.Field
	SupportGroups            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessAzureAdConfigJSON) RawJSON() string {
	return r.raw
}

// Indicates the type of user interaction that is required. prompt=login forces the
// user to enter their credentials on that request, negating single-sign on.
// prompt=none is the opposite. It ensures that the user isn't presented with any
// interactive prompt. If the request can't be completed silently by using
// single-sign on, the Microsoft identity platform returns an interaction_required
// error. prompt=select_account interrupts single sign-on providing account
// selection experience listing all the accounts either in session or any
// remembered account or an option to choose to use a different account altogether.
type ZeroTrustIdentityProvidersAccessAzureAdConfigPrompt string

const (
	ZeroTrustIdentityProvidersAccessAzureAdConfigPromptLogin         ZeroTrustIdentityProvidersAccessAzureAdConfigPrompt = "login"
	ZeroTrustIdentityProvidersAccessAzureAdConfigPromptSelectAccount ZeroTrustIdentityProvidersAccessAzureAdConfigPrompt = "select_account"
	ZeroTrustIdentityProvidersAccessAzureAdConfigPromptNone          ZeroTrustIdentityProvidersAccessAzureAdConfigPrompt = "none"
)

func (r ZeroTrustIdentityProvidersAccessAzureAdConfigPrompt) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessAzureAdConfigPromptLogin, ZeroTrustIdentityProvidersAccessAzureAdConfigPromptSelectAccount, ZeroTrustIdentityProvidersAccessAzureAdConfigPromptNone:
		return true
	}
	return false
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessAzureAdType string

const (
	ZeroTrustIdentityProvidersAccessAzureAdTypeOnetimepin ZeroTrustIdentityProvidersAccessAzureAdType = "onetimepin"
	ZeroTrustIdentityProvidersAccessAzureAdTypeAzureAd    ZeroTrustIdentityProvidersAccessAzureAdType = "azureAD"
	ZeroTrustIdentityProvidersAccessAzureAdTypeSaml       ZeroTrustIdentityProvidersAccessAzureAdType = "saml"
	ZeroTrustIdentityProvidersAccessAzureAdTypeCentrify   ZeroTrustIdentityProvidersAccessAzureAdType = "centrify"
	ZeroTrustIdentityProvidersAccessAzureAdTypeFacebook   ZeroTrustIdentityProvidersAccessAzureAdType = "facebook"
	ZeroTrustIdentityProvidersAccessAzureAdTypeGitHub     ZeroTrustIdentityProvidersAccessAzureAdType = "github"
	ZeroTrustIdentityProvidersAccessAzureAdTypeGoogleApps ZeroTrustIdentityProvidersAccessAzureAdType = "google-apps"
	ZeroTrustIdentityProvidersAccessAzureAdTypeGoogle     ZeroTrustIdentityProvidersAccessAzureAdType = "google"
	ZeroTrustIdentityProvidersAccessAzureAdTypeLinkedin   ZeroTrustIdentityProvidersAccessAzureAdType = "linkedin"
	ZeroTrustIdentityProvidersAccessAzureAdTypeOidc       ZeroTrustIdentityProvidersAccessAzureAdType = "oidc"
	ZeroTrustIdentityProvidersAccessAzureAdTypeOkta       ZeroTrustIdentityProvidersAccessAzureAdType = "okta"
	ZeroTrustIdentityProvidersAccessAzureAdTypeOnelogin   ZeroTrustIdentityProvidersAccessAzureAdType = "onelogin"
	ZeroTrustIdentityProvidersAccessAzureAdTypePingone    ZeroTrustIdentityProvidersAccessAzureAdType = "pingone"
	ZeroTrustIdentityProvidersAccessAzureAdTypeYandex     ZeroTrustIdentityProvidersAccessAzureAdType = "yandex"
)

func (r ZeroTrustIdentityProvidersAccessAzureAdType) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessAzureAdTypeOnetimepin, ZeroTrustIdentityProvidersAccessAzureAdTypeAzureAd, ZeroTrustIdentityProvidersAccessAzureAdTypeSaml, ZeroTrustIdentityProvidersAccessAzureAdTypeCentrify, ZeroTrustIdentityProvidersAccessAzureAdTypeFacebook, ZeroTrustIdentityProvidersAccessAzureAdTypeGitHub, ZeroTrustIdentityProvidersAccessAzureAdTypeGoogleApps, ZeroTrustIdentityProvidersAccessAzureAdTypeGoogle, ZeroTrustIdentityProvidersAccessAzureAdTypeLinkedin, ZeroTrustIdentityProvidersAccessAzureAdTypeOidc, ZeroTrustIdentityProvidersAccessAzureAdTypeOkta, ZeroTrustIdentityProvidersAccessAzureAdTypeOnelogin, ZeroTrustIdentityProvidersAccessAzureAdTypePingone, ZeroTrustIdentityProvidersAccessAzureAdTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProvidersAccessAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProvidersAccessAzureAdScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessAzureAdScimConfigJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProvidersAccessAzureAdScimConfig]
type zeroTrustIdentityProvidersAccessAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessAzureAdScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProvidersAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProvidersAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProvidersAccessCentrifyType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProvidersAccessCentrifyScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProvidersAccessCentrifyJSON       `json:"-"`
}

// zeroTrustIdentityProvidersAccessCentrifyJSON contains the JSON metadata for the
// struct [ZeroTrustIdentityProvidersAccessCentrify]
type zeroTrustIdentityProvidersAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessCentrifyJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProvidersAccessCentrify) implementsZeroTrustZeroTrustIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessCentrifyConfig struct {
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
	EmailClaimName string                                             `json:"email_claim_name"`
	JSON           zeroTrustIdentityProvidersAccessCentrifyConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessCentrifyConfigJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProvidersAccessCentrifyConfig]
type zeroTrustIdentityProvidersAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessCentrifyConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessCentrifyType string

const (
	ZeroTrustIdentityProvidersAccessCentrifyTypeOnetimepin ZeroTrustIdentityProvidersAccessCentrifyType = "onetimepin"
	ZeroTrustIdentityProvidersAccessCentrifyTypeAzureAd    ZeroTrustIdentityProvidersAccessCentrifyType = "azureAD"
	ZeroTrustIdentityProvidersAccessCentrifyTypeSaml       ZeroTrustIdentityProvidersAccessCentrifyType = "saml"
	ZeroTrustIdentityProvidersAccessCentrifyTypeCentrify   ZeroTrustIdentityProvidersAccessCentrifyType = "centrify"
	ZeroTrustIdentityProvidersAccessCentrifyTypeFacebook   ZeroTrustIdentityProvidersAccessCentrifyType = "facebook"
	ZeroTrustIdentityProvidersAccessCentrifyTypeGitHub     ZeroTrustIdentityProvidersAccessCentrifyType = "github"
	ZeroTrustIdentityProvidersAccessCentrifyTypeGoogleApps ZeroTrustIdentityProvidersAccessCentrifyType = "google-apps"
	ZeroTrustIdentityProvidersAccessCentrifyTypeGoogle     ZeroTrustIdentityProvidersAccessCentrifyType = "google"
	ZeroTrustIdentityProvidersAccessCentrifyTypeLinkedin   ZeroTrustIdentityProvidersAccessCentrifyType = "linkedin"
	ZeroTrustIdentityProvidersAccessCentrifyTypeOidc       ZeroTrustIdentityProvidersAccessCentrifyType = "oidc"
	ZeroTrustIdentityProvidersAccessCentrifyTypeOkta       ZeroTrustIdentityProvidersAccessCentrifyType = "okta"
	ZeroTrustIdentityProvidersAccessCentrifyTypeOnelogin   ZeroTrustIdentityProvidersAccessCentrifyType = "onelogin"
	ZeroTrustIdentityProvidersAccessCentrifyTypePingone    ZeroTrustIdentityProvidersAccessCentrifyType = "pingone"
	ZeroTrustIdentityProvidersAccessCentrifyTypeYandex     ZeroTrustIdentityProvidersAccessCentrifyType = "yandex"
)

func (r ZeroTrustIdentityProvidersAccessCentrifyType) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessCentrifyTypeOnetimepin, ZeroTrustIdentityProvidersAccessCentrifyTypeAzureAd, ZeroTrustIdentityProvidersAccessCentrifyTypeSaml, ZeroTrustIdentityProvidersAccessCentrifyTypeCentrify, ZeroTrustIdentityProvidersAccessCentrifyTypeFacebook, ZeroTrustIdentityProvidersAccessCentrifyTypeGitHub, ZeroTrustIdentityProvidersAccessCentrifyTypeGoogleApps, ZeroTrustIdentityProvidersAccessCentrifyTypeGoogle, ZeroTrustIdentityProvidersAccessCentrifyTypeLinkedin, ZeroTrustIdentityProvidersAccessCentrifyTypeOidc, ZeroTrustIdentityProvidersAccessCentrifyTypeOkta, ZeroTrustIdentityProvidersAccessCentrifyTypeOnelogin, ZeroTrustIdentityProvidersAccessCentrifyTypePingone, ZeroTrustIdentityProvidersAccessCentrifyTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProvidersAccessCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProvidersAccessCentrifyScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessCentrifyScimConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProvidersAccessCentrifyScimConfig]
type zeroTrustIdentityProvidersAccessCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessCentrifyScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProvidersAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProvidersAccessFacebookConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProvidersAccessFacebookType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProvidersAccessFacebookScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProvidersAccessFacebookJSON       `json:"-"`
}

// zeroTrustIdentityProvidersAccessFacebookJSON contains the JSON metadata for the
// struct [ZeroTrustIdentityProvidersAccessFacebook]
type zeroTrustIdentityProvidersAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessFacebookJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProvidersAccessFacebook) implementsZeroTrustZeroTrustIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                             `json:"client_secret"`
	JSON         zeroTrustIdentityProvidersAccessFacebookConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessFacebookConfigJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProvidersAccessFacebookConfig]
type zeroTrustIdentityProvidersAccessFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessFacebookConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessFacebookType string

const (
	ZeroTrustIdentityProvidersAccessFacebookTypeOnetimepin ZeroTrustIdentityProvidersAccessFacebookType = "onetimepin"
	ZeroTrustIdentityProvidersAccessFacebookTypeAzureAd    ZeroTrustIdentityProvidersAccessFacebookType = "azureAD"
	ZeroTrustIdentityProvidersAccessFacebookTypeSaml       ZeroTrustIdentityProvidersAccessFacebookType = "saml"
	ZeroTrustIdentityProvidersAccessFacebookTypeCentrify   ZeroTrustIdentityProvidersAccessFacebookType = "centrify"
	ZeroTrustIdentityProvidersAccessFacebookTypeFacebook   ZeroTrustIdentityProvidersAccessFacebookType = "facebook"
	ZeroTrustIdentityProvidersAccessFacebookTypeGitHub     ZeroTrustIdentityProvidersAccessFacebookType = "github"
	ZeroTrustIdentityProvidersAccessFacebookTypeGoogleApps ZeroTrustIdentityProvidersAccessFacebookType = "google-apps"
	ZeroTrustIdentityProvidersAccessFacebookTypeGoogle     ZeroTrustIdentityProvidersAccessFacebookType = "google"
	ZeroTrustIdentityProvidersAccessFacebookTypeLinkedin   ZeroTrustIdentityProvidersAccessFacebookType = "linkedin"
	ZeroTrustIdentityProvidersAccessFacebookTypeOidc       ZeroTrustIdentityProvidersAccessFacebookType = "oidc"
	ZeroTrustIdentityProvidersAccessFacebookTypeOkta       ZeroTrustIdentityProvidersAccessFacebookType = "okta"
	ZeroTrustIdentityProvidersAccessFacebookTypeOnelogin   ZeroTrustIdentityProvidersAccessFacebookType = "onelogin"
	ZeroTrustIdentityProvidersAccessFacebookTypePingone    ZeroTrustIdentityProvidersAccessFacebookType = "pingone"
	ZeroTrustIdentityProvidersAccessFacebookTypeYandex     ZeroTrustIdentityProvidersAccessFacebookType = "yandex"
)

func (r ZeroTrustIdentityProvidersAccessFacebookType) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessFacebookTypeOnetimepin, ZeroTrustIdentityProvidersAccessFacebookTypeAzureAd, ZeroTrustIdentityProvidersAccessFacebookTypeSaml, ZeroTrustIdentityProvidersAccessFacebookTypeCentrify, ZeroTrustIdentityProvidersAccessFacebookTypeFacebook, ZeroTrustIdentityProvidersAccessFacebookTypeGitHub, ZeroTrustIdentityProvidersAccessFacebookTypeGoogleApps, ZeroTrustIdentityProvidersAccessFacebookTypeGoogle, ZeroTrustIdentityProvidersAccessFacebookTypeLinkedin, ZeroTrustIdentityProvidersAccessFacebookTypeOidc, ZeroTrustIdentityProvidersAccessFacebookTypeOkta, ZeroTrustIdentityProvidersAccessFacebookTypeOnelogin, ZeroTrustIdentityProvidersAccessFacebookTypePingone, ZeroTrustIdentityProvidersAccessFacebookTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProvidersAccessFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProvidersAccessFacebookScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessFacebookScimConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProvidersAccessFacebookScimConfig]
type zeroTrustIdentityProvidersAccessFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessFacebookScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProvidersAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProvidersAccessGitHubConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProvidersAccessGitHubType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProvidersAccessGitHubScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProvidersAccessGitHubJSON       `json:"-"`
}

// zeroTrustIdentityProvidersAccessGitHubJSON contains the JSON metadata for the
// struct [ZeroTrustIdentityProvidersAccessGitHub]
type zeroTrustIdentityProvidersAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessGitHubJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProvidersAccessGitHub) implementsZeroTrustZeroTrustIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                           `json:"client_secret"`
	JSON         zeroTrustIdentityProvidersAccessGitHubConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessGitHubConfigJSON contains the JSON metadata for
// the struct [ZeroTrustIdentityProvidersAccessGitHubConfig]
type zeroTrustIdentityProvidersAccessGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessGitHubConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessGitHubType string

const (
	ZeroTrustIdentityProvidersAccessGitHubTypeOnetimepin ZeroTrustIdentityProvidersAccessGitHubType = "onetimepin"
	ZeroTrustIdentityProvidersAccessGitHubTypeAzureAd    ZeroTrustIdentityProvidersAccessGitHubType = "azureAD"
	ZeroTrustIdentityProvidersAccessGitHubTypeSaml       ZeroTrustIdentityProvidersAccessGitHubType = "saml"
	ZeroTrustIdentityProvidersAccessGitHubTypeCentrify   ZeroTrustIdentityProvidersAccessGitHubType = "centrify"
	ZeroTrustIdentityProvidersAccessGitHubTypeFacebook   ZeroTrustIdentityProvidersAccessGitHubType = "facebook"
	ZeroTrustIdentityProvidersAccessGitHubTypeGitHub     ZeroTrustIdentityProvidersAccessGitHubType = "github"
	ZeroTrustIdentityProvidersAccessGitHubTypeGoogleApps ZeroTrustIdentityProvidersAccessGitHubType = "google-apps"
	ZeroTrustIdentityProvidersAccessGitHubTypeGoogle     ZeroTrustIdentityProvidersAccessGitHubType = "google"
	ZeroTrustIdentityProvidersAccessGitHubTypeLinkedin   ZeroTrustIdentityProvidersAccessGitHubType = "linkedin"
	ZeroTrustIdentityProvidersAccessGitHubTypeOidc       ZeroTrustIdentityProvidersAccessGitHubType = "oidc"
	ZeroTrustIdentityProvidersAccessGitHubTypeOkta       ZeroTrustIdentityProvidersAccessGitHubType = "okta"
	ZeroTrustIdentityProvidersAccessGitHubTypeOnelogin   ZeroTrustIdentityProvidersAccessGitHubType = "onelogin"
	ZeroTrustIdentityProvidersAccessGitHubTypePingone    ZeroTrustIdentityProvidersAccessGitHubType = "pingone"
	ZeroTrustIdentityProvidersAccessGitHubTypeYandex     ZeroTrustIdentityProvidersAccessGitHubType = "yandex"
)

func (r ZeroTrustIdentityProvidersAccessGitHubType) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessGitHubTypeOnetimepin, ZeroTrustIdentityProvidersAccessGitHubTypeAzureAd, ZeroTrustIdentityProvidersAccessGitHubTypeSaml, ZeroTrustIdentityProvidersAccessGitHubTypeCentrify, ZeroTrustIdentityProvidersAccessGitHubTypeFacebook, ZeroTrustIdentityProvidersAccessGitHubTypeGitHub, ZeroTrustIdentityProvidersAccessGitHubTypeGoogleApps, ZeroTrustIdentityProvidersAccessGitHubTypeGoogle, ZeroTrustIdentityProvidersAccessGitHubTypeLinkedin, ZeroTrustIdentityProvidersAccessGitHubTypeOidc, ZeroTrustIdentityProvidersAccessGitHubTypeOkta, ZeroTrustIdentityProvidersAccessGitHubTypeOnelogin, ZeroTrustIdentityProvidersAccessGitHubTypePingone, ZeroTrustIdentityProvidersAccessGitHubTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProvidersAccessGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProvidersAccessGitHubScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessGitHubScimConfigJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProvidersAccessGitHubScimConfig]
type zeroTrustIdentityProvidersAccessGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessGitHubScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProvidersAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProvidersAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProvidersAccessGoogleType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProvidersAccessGoogleScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProvidersAccessGoogleJSON       `json:"-"`
}

// zeroTrustIdentityProvidersAccessGoogleJSON contains the JSON metadata for the
// struct [ZeroTrustIdentityProvidersAccessGoogle]
type zeroTrustIdentityProvidersAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessGoogleJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProvidersAccessGoogle) implementsZeroTrustZeroTrustIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                           `json:"email_claim_name"`
	JSON           zeroTrustIdentityProvidersAccessGoogleConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessGoogleConfigJSON contains the JSON metadata for
// the struct [ZeroTrustIdentityProvidersAccessGoogleConfig]
type zeroTrustIdentityProvidersAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessGoogleConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessGoogleType string

const (
	ZeroTrustIdentityProvidersAccessGoogleTypeOnetimepin ZeroTrustIdentityProvidersAccessGoogleType = "onetimepin"
	ZeroTrustIdentityProvidersAccessGoogleTypeAzureAd    ZeroTrustIdentityProvidersAccessGoogleType = "azureAD"
	ZeroTrustIdentityProvidersAccessGoogleTypeSaml       ZeroTrustIdentityProvidersAccessGoogleType = "saml"
	ZeroTrustIdentityProvidersAccessGoogleTypeCentrify   ZeroTrustIdentityProvidersAccessGoogleType = "centrify"
	ZeroTrustIdentityProvidersAccessGoogleTypeFacebook   ZeroTrustIdentityProvidersAccessGoogleType = "facebook"
	ZeroTrustIdentityProvidersAccessGoogleTypeGitHub     ZeroTrustIdentityProvidersAccessGoogleType = "github"
	ZeroTrustIdentityProvidersAccessGoogleTypeGoogleApps ZeroTrustIdentityProvidersAccessGoogleType = "google-apps"
	ZeroTrustIdentityProvidersAccessGoogleTypeGoogle     ZeroTrustIdentityProvidersAccessGoogleType = "google"
	ZeroTrustIdentityProvidersAccessGoogleTypeLinkedin   ZeroTrustIdentityProvidersAccessGoogleType = "linkedin"
	ZeroTrustIdentityProvidersAccessGoogleTypeOidc       ZeroTrustIdentityProvidersAccessGoogleType = "oidc"
	ZeroTrustIdentityProvidersAccessGoogleTypeOkta       ZeroTrustIdentityProvidersAccessGoogleType = "okta"
	ZeroTrustIdentityProvidersAccessGoogleTypeOnelogin   ZeroTrustIdentityProvidersAccessGoogleType = "onelogin"
	ZeroTrustIdentityProvidersAccessGoogleTypePingone    ZeroTrustIdentityProvidersAccessGoogleType = "pingone"
	ZeroTrustIdentityProvidersAccessGoogleTypeYandex     ZeroTrustIdentityProvidersAccessGoogleType = "yandex"
)

func (r ZeroTrustIdentityProvidersAccessGoogleType) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessGoogleTypeOnetimepin, ZeroTrustIdentityProvidersAccessGoogleTypeAzureAd, ZeroTrustIdentityProvidersAccessGoogleTypeSaml, ZeroTrustIdentityProvidersAccessGoogleTypeCentrify, ZeroTrustIdentityProvidersAccessGoogleTypeFacebook, ZeroTrustIdentityProvidersAccessGoogleTypeGitHub, ZeroTrustIdentityProvidersAccessGoogleTypeGoogleApps, ZeroTrustIdentityProvidersAccessGoogleTypeGoogle, ZeroTrustIdentityProvidersAccessGoogleTypeLinkedin, ZeroTrustIdentityProvidersAccessGoogleTypeOidc, ZeroTrustIdentityProvidersAccessGoogleTypeOkta, ZeroTrustIdentityProvidersAccessGoogleTypeOnelogin, ZeroTrustIdentityProvidersAccessGoogleTypePingone, ZeroTrustIdentityProvidersAccessGoogleTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProvidersAccessGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProvidersAccessGoogleScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessGoogleScimConfigJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProvidersAccessGoogleScimConfig]
type zeroTrustIdentityProvidersAccessGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessGoogleScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProvidersAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProvidersAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProvidersAccessGoogleAppsType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProvidersAccessGoogleAppsScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProvidersAccessGoogleAppsJSON       `json:"-"`
}

// zeroTrustIdentityProvidersAccessGoogleAppsJSON contains the JSON metadata for
// the struct [ZeroTrustIdentityProvidersAccessGoogleApps]
type zeroTrustIdentityProvidersAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessGoogleAppsJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProvidersAccessGoogleApps) implementsZeroTrustZeroTrustIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                               `json:"email_claim_name"`
	JSON           zeroTrustIdentityProvidersAccessGoogleAppsConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessGoogleAppsConfigJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProvidersAccessGoogleAppsConfig]
type zeroTrustIdentityProvidersAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessGoogleAppsConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessGoogleAppsType string

const (
	ZeroTrustIdentityProvidersAccessGoogleAppsTypeOnetimepin ZeroTrustIdentityProvidersAccessGoogleAppsType = "onetimepin"
	ZeroTrustIdentityProvidersAccessGoogleAppsTypeAzureAd    ZeroTrustIdentityProvidersAccessGoogleAppsType = "azureAD"
	ZeroTrustIdentityProvidersAccessGoogleAppsTypeSaml       ZeroTrustIdentityProvidersAccessGoogleAppsType = "saml"
	ZeroTrustIdentityProvidersAccessGoogleAppsTypeCentrify   ZeroTrustIdentityProvidersAccessGoogleAppsType = "centrify"
	ZeroTrustIdentityProvidersAccessGoogleAppsTypeFacebook   ZeroTrustIdentityProvidersAccessGoogleAppsType = "facebook"
	ZeroTrustIdentityProvidersAccessGoogleAppsTypeGitHub     ZeroTrustIdentityProvidersAccessGoogleAppsType = "github"
	ZeroTrustIdentityProvidersAccessGoogleAppsTypeGoogleApps ZeroTrustIdentityProvidersAccessGoogleAppsType = "google-apps"
	ZeroTrustIdentityProvidersAccessGoogleAppsTypeGoogle     ZeroTrustIdentityProvidersAccessGoogleAppsType = "google"
	ZeroTrustIdentityProvidersAccessGoogleAppsTypeLinkedin   ZeroTrustIdentityProvidersAccessGoogleAppsType = "linkedin"
	ZeroTrustIdentityProvidersAccessGoogleAppsTypeOidc       ZeroTrustIdentityProvidersAccessGoogleAppsType = "oidc"
	ZeroTrustIdentityProvidersAccessGoogleAppsTypeOkta       ZeroTrustIdentityProvidersAccessGoogleAppsType = "okta"
	ZeroTrustIdentityProvidersAccessGoogleAppsTypeOnelogin   ZeroTrustIdentityProvidersAccessGoogleAppsType = "onelogin"
	ZeroTrustIdentityProvidersAccessGoogleAppsTypePingone    ZeroTrustIdentityProvidersAccessGoogleAppsType = "pingone"
	ZeroTrustIdentityProvidersAccessGoogleAppsTypeYandex     ZeroTrustIdentityProvidersAccessGoogleAppsType = "yandex"
)

func (r ZeroTrustIdentityProvidersAccessGoogleAppsType) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessGoogleAppsTypeOnetimepin, ZeroTrustIdentityProvidersAccessGoogleAppsTypeAzureAd, ZeroTrustIdentityProvidersAccessGoogleAppsTypeSaml, ZeroTrustIdentityProvidersAccessGoogleAppsTypeCentrify, ZeroTrustIdentityProvidersAccessGoogleAppsTypeFacebook, ZeroTrustIdentityProvidersAccessGoogleAppsTypeGitHub, ZeroTrustIdentityProvidersAccessGoogleAppsTypeGoogleApps, ZeroTrustIdentityProvidersAccessGoogleAppsTypeGoogle, ZeroTrustIdentityProvidersAccessGoogleAppsTypeLinkedin, ZeroTrustIdentityProvidersAccessGoogleAppsTypeOidc, ZeroTrustIdentityProvidersAccessGoogleAppsTypeOkta, ZeroTrustIdentityProvidersAccessGoogleAppsTypeOnelogin, ZeroTrustIdentityProvidersAccessGoogleAppsTypePingone, ZeroTrustIdentityProvidersAccessGoogleAppsTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProvidersAccessGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProvidersAccessGoogleAppsScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessGoogleAppsScimConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProvidersAccessGoogleAppsScimConfig]
type zeroTrustIdentityProvidersAccessGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessGoogleAppsScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProvidersAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProvidersAccessLinkedinConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProvidersAccessLinkedinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProvidersAccessLinkedinScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProvidersAccessLinkedinJSON       `json:"-"`
}

// zeroTrustIdentityProvidersAccessLinkedinJSON contains the JSON metadata for the
// struct [ZeroTrustIdentityProvidersAccessLinkedin]
type zeroTrustIdentityProvidersAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessLinkedinJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProvidersAccessLinkedin) implementsZeroTrustZeroTrustIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                             `json:"client_secret"`
	JSON         zeroTrustIdentityProvidersAccessLinkedinConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessLinkedinConfigJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProvidersAccessLinkedinConfig]
type zeroTrustIdentityProvidersAccessLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessLinkedinConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessLinkedinType string

const (
	ZeroTrustIdentityProvidersAccessLinkedinTypeOnetimepin ZeroTrustIdentityProvidersAccessLinkedinType = "onetimepin"
	ZeroTrustIdentityProvidersAccessLinkedinTypeAzureAd    ZeroTrustIdentityProvidersAccessLinkedinType = "azureAD"
	ZeroTrustIdentityProvidersAccessLinkedinTypeSaml       ZeroTrustIdentityProvidersAccessLinkedinType = "saml"
	ZeroTrustIdentityProvidersAccessLinkedinTypeCentrify   ZeroTrustIdentityProvidersAccessLinkedinType = "centrify"
	ZeroTrustIdentityProvidersAccessLinkedinTypeFacebook   ZeroTrustIdentityProvidersAccessLinkedinType = "facebook"
	ZeroTrustIdentityProvidersAccessLinkedinTypeGitHub     ZeroTrustIdentityProvidersAccessLinkedinType = "github"
	ZeroTrustIdentityProvidersAccessLinkedinTypeGoogleApps ZeroTrustIdentityProvidersAccessLinkedinType = "google-apps"
	ZeroTrustIdentityProvidersAccessLinkedinTypeGoogle     ZeroTrustIdentityProvidersAccessLinkedinType = "google"
	ZeroTrustIdentityProvidersAccessLinkedinTypeLinkedin   ZeroTrustIdentityProvidersAccessLinkedinType = "linkedin"
	ZeroTrustIdentityProvidersAccessLinkedinTypeOidc       ZeroTrustIdentityProvidersAccessLinkedinType = "oidc"
	ZeroTrustIdentityProvidersAccessLinkedinTypeOkta       ZeroTrustIdentityProvidersAccessLinkedinType = "okta"
	ZeroTrustIdentityProvidersAccessLinkedinTypeOnelogin   ZeroTrustIdentityProvidersAccessLinkedinType = "onelogin"
	ZeroTrustIdentityProvidersAccessLinkedinTypePingone    ZeroTrustIdentityProvidersAccessLinkedinType = "pingone"
	ZeroTrustIdentityProvidersAccessLinkedinTypeYandex     ZeroTrustIdentityProvidersAccessLinkedinType = "yandex"
)

func (r ZeroTrustIdentityProvidersAccessLinkedinType) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessLinkedinTypeOnetimepin, ZeroTrustIdentityProvidersAccessLinkedinTypeAzureAd, ZeroTrustIdentityProvidersAccessLinkedinTypeSaml, ZeroTrustIdentityProvidersAccessLinkedinTypeCentrify, ZeroTrustIdentityProvidersAccessLinkedinTypeFacebook, ZeroTrustIdentityProvidersAccessLinkedinTypeGitHub, ZeroTrustIdentityProvidersAccessLinkedinTypeGoogleApps, ZeroTrustIdentityProvidersAccessLinkedinTypeGoogle, ZeroTrustIdentityProvidersAccessLinkedinTypeLinkedin, ZeroTrustIdentityProvidersAccessLinkedinTypeOidc, ZeroTrustIdentityProvidersAccessLinkedinTypeOkta, ZeroTrustIdentityProvidersAccessLinkedinTypeOnelogin, ZeroTrustIdentityProvidersAccessLinkedinTypePingone, ZeroTrustIdentityProvidersAccessLinkedinTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProvidersAccessLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProvidersAccessLinkedinScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessLinkedinScimConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProvidersAccessLinkedinScimConfig]
type zeroTrustIdentityProvidersAccessLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessLinkedinScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProvidersAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProvidersAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProvidersAccessOidcType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProvidersAccessOidcScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProvidersAccessOidcJSON       `json:"-"`
}

// zeroTrustIdentityProvidersAccessOidcJSON contains the JSON metadata for the
// struct [ZeroTrustIdentityProvidersAccessOidc]
type zeroTrustIdentityProvidersAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessOidcJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProvidersAccessOidc) implementsZeroTrustZeroTrustIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL string `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CERTsURL string `json:"certs_url"`
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
	TokenURL string                                         `json:"token_url"`
	JSON     zeroTrustIdentityProvidersAccessOidcConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessOidcConfigJSON contains the JSON metadata for
// the struct [ZeroTrustIdentityProvidersAccessOidcConfig]
type zeroTrustIdentityProvidersAccessOidcConfigJSON struct {
	AuthURL        apijson.Field
	CERTsURL       apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	Scopes         apijson.Field
	TokenURL       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessOidcConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessOidcType string

const (
	ZeroTrustIdentityProvidersAccessOidcTypeOnetimepin ZeroTrustIdentityProvidersAccessOidcType = "onetimepin"
	ZeroTrustIdentityProvidersAccessOidcTypeAzureAd    ZeroTrustIdentityProvidersAccessOidcType = "azureAD"
	ZeroTrustIdentityProvidersAccessOidcTypeSaml       ZeroTrustIdentityProvidersAccessOidcType = "saml"
	ZeroTrustIdentityProvidersAccessOidcTypeCentrify   ZeroTrustIdentityProvidersAccessOidcType = "centrify"
	ZeroTrustIdentityProvidersAccessOidcTypeFacebook   ZeroTrustIdentityProvidersAccessOidcType = "facebook"
	ZeroTrustIdentityProvidersAccessOidcTypeGitHub     ZeroTrustIdentityProvidersAccessOidcType = "github"
	ZeroTrustIdentityProvidersAccessOidcTypeGoogleApps ZeroTrustIdentityProvidersAccessOidcType = "google-apps"
	ZeroTrustIdentityProvidersAccessOidcTypeGoogle     ZeroTrustIdentityProvidersAccessOidcType = "google"
	ZeroTrustIdentityProvidersAccessOidcTypeLinkedin   ZeroTrustIdentityProvidersAccessOidcType = "linkedin"
	ZeroTrustIdentityProvidersAccessOidcTypeOidc       ZeroTrustIdentityProvidersAccessOidcType = "oidc"
	ZeroTrustIdentityProvidersAccessOidcTypeOkta       ZeroTrustIdentityProvidersAccessOidcType = "okta"
	ZeroTrustIdentityProvidersAccessOidcTypeOnelogin   ZeroTrustIdentityProvidersAccessOidcType = "onelogin"
	ZeroTrustIdentityProvidersAccessOidcTypePingone    ZeroTrustIdentityProvidersAccessOidcType = "pingone"
	ZeroTrustIdentityProvidersAccessOidcTypeYandex     ZeroTrustIdentityProvidersAccessOidcType = "yandex"
)

func (r ZeroTrustIdentityProvidersAccessOidcType) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessOidcTypeOnetimepin, ZeroTrustIdentityProvidersAccessOidcTypeAzureAd, ZeroTrustIdentityProvidersAccessOidcTypeSaml, ZeroTrustIdentityProvidersAccessOidcTypeCentrify, ZeroTrustIdentityProvidersAccessOidcTypeFacebook, ZeroTrustIdentityProvidersAccessOidcTypeGitHub, ZeroTrustIdentityProvidersAccessOidcTypeGoogleApps, ZeroTrustIdentityProvidersAccessOidcTypeGoogle, ZeroTrustIdentityProvidersAccessOidcTypeLinkedin, ZeroTrustIdentityProvidersAccessOidcTypeOidc, ZeroTrustIdentityProvidersAccessOidcTypeOkta, ZeroTrustIdentityProvidersAccessOidcTypeOnelogin, ZeroTrustIdentityProvidersAccessOidcTypePingone, ZeroTrustIdentityProvidersAccessOidcTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProvidersAccessOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProvidersAccessOidcScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessOidcScimConfigJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProvidersAccessOidcScimConfig]
type zeroTrustIdentityProvidersAccessOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessOidcScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProvidersAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProvidersAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProvidersAccessOktaType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProvidersAccessOktaScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProvidersAccessOktaJSON       `json:"-"`
}

// zeroTrustIdentityProvidersAccessOktaJSON contains the JSON metadata for the
// struct [ZeroTrustIdentityProvidersAccessOkta]
type zeroTrustIdentityProvidersAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessOktaJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProvidersAccessOkta) implementsZeroTrustZeroTrustIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessOktaConfig struct {
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
	OktaAccount string                                         `json:"okta_account"`
	JSON        zeroTrustIdentityProvidersAccessOktaConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessOktaConfigJSON contains the JSON metadata for
// the struct [ZeroTrustIdentityProvidersAccessOktaConfig]
type zeroTrustIdentityProvidersAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessOktaConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessOktaType string

const (
	ZeroTrustIdentityProvidersAccessOktaTypeOnetimepin ZeroTrustIdentityProvidersAccessOktaType = "onetimepin"
	ZeroTrustIdentityProvidersAccessOktaTypeAzureAd    ZeroTrustIdentityProvidersAccessOktaType = "azureAD"
	ZeroTrustIdentityProvidersAccessOktaTypeSaml       ZeroTrustIdentityProvidersAccessOktaType = "saml"
	ZeroTrustIdentityProvidersAccessOktaTypeCentrify   ZeroTrustIdentityProvidersAccessOktaType = "centrify"
	ZeroTrustIdentityProvidersAccessOktaTypeFacebook   ZeroTrustIdentityProvidersAccessOktaType = "facebook"
	ZeroTrustIdentityProvidersAccessOktaTypeGitHub     ZeroTrustIdentityProvidersAccessOktaType = "github"
	ZeroTrustIdentityProvidersAccessOktaTypeGoogleApps ZeroTrustIdentityProvidersAccessOktaType = "google-apps"
	ZeroTrustIdentityProvidersAccessOktaTypeGoogle     ZeroTrustIdentityProvidersAccessOktaType = "google"
	ZeroTrustIdentityProvidersAccessOktaTypeLinkedin   ZeroTrustIdentityProvidersAccessOktaType = "linkedin"
	ZeroTrustIdentityProvidersAccessOktaTypeOidc       ZeroTrustIdentityProvidersAccessOktaType = "oidc"
	ZeroTrustIdentityProvidersAccessOktaTypeOkta       ZeroTrustIdentityProvidersAccessOktaType = "okta"
	ZeroTrustIdentityProvidersAccessOktaTypeOnelogin   ZeroTrustIdentityProvidersAccessOktaType = "onelogin"
	ZeroTrustIdentityProvidersAccessOktaTypePingone    ZeroTrustIdentityProvidersAccessOktaType = "pingone"
	ZeroTrustIdentityProvidersAccessOktaTypeYandex     ZeroTrustIdentityProvidersAccessOktaType = "yandex"
)

func (r ZeroTrustIdentityProvidersAccessOktaType) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessOktaTypeOnetimepin, ZeroTrustIdentityProvidersAccessOktaTypeAzureAd, ZeroTrustIdentityProvidersAccessOktaTypeSaml, ZeroTrustIdentityProvidersAccessOktaTypeCentrify, ZeroTrustIdentityProvidersAccessOktaTypeFacebook, ZeroTrustIdentityProvidersAccessOktaTypeGitHub, ZeroTrustIdentityProvidersAccessOktaTypeGoogleApps, ZeroTrustIdentityProvidersAccessOktaTypeGoogle, ZeroTrustIdentityProvidersAccessOktaTypeLinkedin, ZeroTrustIdentityProvidersAccessOktaTypeOidc, ZeroTrustIdentityProvidersAccessOktaTypeOkta, ZeroTrustIdentityProvidersAccessOktaTypeOnelogin, ZeroTrustIdentityProvidersAccessOktaTypePingone, ZeroTrustIdentityProvidersAccessOktaTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProvidersAccessOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProvidersAccessOktaScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessOktaScimConfigJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProvidersAccessOktaScimConfig]
type zeroTrustIdentityProvidersAccessOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessOktaScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProvidersAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProvidersAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProvidersAccessOneloginType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProvidersAccessOneloginScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProvidersAccessOneloginJSON       `json:"-"`
}

// zeroTrustIdentityProvidersAccessOneloginJSON contains the JSON metadata for the
// struct [ZeroTrustIdentityProvidersAccessOnelogin]
type zeroTrustIdentityProvidersAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessOneloginJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProvidersAccessOnelogin) implementsZeroTrustZeroTrustIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                             `json:"onelogin_account"`
	JSON            zeroTrustIdentityProvidersAccessOneloginConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessOneloginConfigJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProvidersAccessOneloginConfig]
type zeroTrustIdentityProvidersAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessOneloginConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessOneloginType string

const (
	ZeroTrustIdentityProvidersAccessOneloginTypeOnetimepin ZeroTrustIdentityProvidersAccessOneloginType = "onetimepin"
	ZeroTrustIdentityProvidersAccessOneloginTypeAzureAd    ZeroTrustIdentityProvidersAccessOneloginType = "azureAD"
	ZeroTrustIdentityProvidersAccessOneloginTypeSaml       ZeroTrustIdentityProvidersAccessOneloginType = "saml"
	ZeroTrustIdentityProvidersAccessOneloginTypeCentrify   ZeroTrustIdentityProvidersAccessOneloginType = "centrify"
	ZeroTrustIdentityProvidersAccessOneloginTypeFacebook   ZeroTrustIdentityProvidersAccessOneloginType = "facebook"
	ZeroTrustIdentityProvidersAccessOneloginTypeGitHub     ZeroTrustIdentityProvidersAccessOneloginType = "github"
	ZeroTrustIdentityProvidersAccessOneloginTypeGoogleApps ZeroTrustIdentityProvidersAccessOneloginType = "google-apps"
	ZeroTrustIdentityProvidersAccessOneloginTypeGoogle     ZeroTrustIdentityProvidersAccessOneloginType = "google"
	ZeroTrustIdentityProvidersAccessOneloginTypeLinkedin   ZeroTrustIdentityProvidersAccessOneloginType = "linkedin"
	ZeroTrustIdentityProvidersAccessOneloginTypeOidc       ZeroTrustIdentityProvidersAccessOneloginType = "oidc"
	ZeroTrustIdentityProvidersAccessOneloginTypeOkta       ZeroTrustIdentityProvidersAccessOneloginType = "okta"
	ZeroTrustIdentityProvidersAccessOneloginTypeOnelogin   ZeroTrustIdentityProvidersAccessOneloginType = "onelogin"
	ZeroTrustIdentityProvidersAccessOneloginTypePingone    ZeroTrustIdentityProvidersAccessOneloginType = "pingone"
	ZeroTrustIdentityProvidersAccessOneloginTypeYandex     ZeroTrustIdentityProvidersAccessOneloginType = "yandex"
)

func (r ZeroTrustIdentityProvidersAccessOneloginType) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessOneloginTypeOnetimepin, ZeroTrustIdentityProvidersAccessOneloginTypeAzureAd, ZeroTrustIdentityProvidersAccessOneloginTypeSaml, ZeroTrustIdentityProvidersAccessOneloginTypeCentrify, ZeroTrustIdentityProvidersAccessOneloginTypeFacebook, ZeroTrustIdentityProvidersAccessOneloginTypeGitHub, ZeroTrustIdentityProvidersAccessOneloginTypeGoogleApps, ZeroTrustIdentityProvidersAccessOneloginTypeGoogle, ZeroTrustIdentityProvidersAccessOneloginTypeLinkedin, ZeroTrustIdentityProvidersAccessOneloginTypeOidc, ZeroTrustIdentityProvidersAccessOneloginTypeOkta, ZeroTrustIdentityProvidersAccessOneloginTypeOnelogin, ZeroTrustIdentityProvidersAccessOneloginTypePingone, ZeroTrustIdentityProvidersAccessOneloginTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProvidersAccessOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProvidersAccessOneloginScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessOneloginScimConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProvidersAccessOneloginScimConfig]
type zeroTrustIdentityProvidersAccessOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessOneloginScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProvidersAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProvidersAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProvidersAccessPingoneType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProvidersAccessPingoneScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProvidersAccessPingoneJSON       `json:"-"`
}

// zeroTrustIdentityProvidersAccessPingoneJSON contains the JSON metadata for the
// struct [ZeroTrustIdentityProvidersAccessPingone]
type zeroTrustIdentityProvidersAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessPingoneJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProvidersAccessPingone) implementsZeroTrustZeroTrustIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                            `json:"ping_env_id"`
	JSON      zeroTrustIdentityProvidersAccessPingoneConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessPingoneConfigJSON contains the JSON metadata for
// the struct [ZeroTrustIdentityProvidersAccessPingoneConfig]
type zeroTrustIdentityProvidersAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessPingoneConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessPingoneType string

const (
	ZeroTrustIdentityProvidersAccessPingoneTypeOnetimepin ZeroTrustIdentityProvidersAccessPingoneType = "onetimepin"
	ZeroTrustIdentityProvidersAccessPingoneTypeAzureAd    ZeroTrustIdentityProvidersAccessPingoneType = "azureAD"
	ZeroTrustIdentityProvidersAccessPingoneTypeSaml       ZeroTrustIdentityProvidersAccessPingoneType = "saml"
	ZeroTrustIdentityProvidersAccessPingoneTypeCentrify   ZeroTrustIdentityProvidersAccessPingoneType = "centrify"
	ZeroTrustIdentityProvidersAccessPingoneTypeFacebook   ZeroTrustIdentityProvidersAccessPingoneType = "facebook"
	ZeroTrustIdentityProvidersAccessPingoneTypeGitHub     ZeroTrustIdentityProvidersAccessPingoneType = "github"
	ZeroTrustIdentityProvidersAccessPingoneTypeGoogleApps ZeroTrustIdentityProvidersAccessPingoneType = "google-apps"
	ZeroTrustIdentityProvidersAccessPingoneTypeGoogle     ZeroTrustIdentityProvidersAccessPingoneType = "google"
	ZeroTrustIdentityProvidersAccessPingoneTypeLinkedin   ZeroTrustIdentityProvidersAccessPingoneType = "linkedin"
	ZeroTrustIdentityProvidersAccessPingoneTypeOidc       ZeroTrustIdentityProvidersAccessPingoneType = "oidc"
	ZeroTrustIdentityProvidersAccessPingoneTypeOkta       ZeroTrustIdentityProvidersAccessPingoneType = "okta"
	ZeroTrustIdentityProvidersAccessPingoneTypeOnelogin   ZeroTrustIdentityProvidersAccessPingoneType = "onelogin"
	ZeroTrustIdentityProvidersAccessPingoneTypePingone    ZeroTrustIdentityProvidersAccessPingoneType = "pingone"
	ZeroTrustIdentityProvidersAccessPingoneTypeYandex     ZeroTrustIdentityProvidersAccessPingoneType = "yandex"
)

func (r ZeroTrustIdentityProvidersAccessPingoneType) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessPingoneTypeOnetimepin, ZeroTrustIdentityProvidersAccessPingoneTypeAzureAd, ZeroTrustIdentityProvidersAccessPingoneTypeSaml, ZeroTrustIdentityProvidersAccessPingoneTypeCentrify, ZeroTrustIdentityProvidersAccessPingoneTypeFacebook, ZeroTrustIdentityProvidersAccessPingoneTypeGitHub, ZeroTrustIdentityProvidersAccessPingoneTypeGoogleApps, ZeroTrustIdentityProvidersAccessPingoneTypeGoogle, ZeroTrustIdentityProvidersAccessPingoneTypeLinkedin, ZeroTrustIdentityProvidersAccessPingoneTypeOidc, ZeroTrustIdentityProvidersAccessPingoneTypeOkta, ZeroTrustIdentityProvidersAccessPingoneTypeOnelogin, ZeroTrustIdentityProvidersAccessPingoneTypePingone, ZeroTrustIdentityProvidersAccessPingoneTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProvidersAccessPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProvidersAccessPingoneScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessPingoneScimConfigJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProvidersAccessPingoneScimConfig]
type zeroTrustIdentityProvidersAccessPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessPingoneScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProvidersAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProvidersAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProvidersAccessSamlType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProvidersAccessSamlScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProvidersAccessSamlJSON       `json:"-"`
}

// zeroTrustIdentityProvidersAccessSamlJSON contains the JSON metadata for the
// struct [ZeroTrustIdentityProvidersAccessSaml]
type zeroTrustIdentityProvidersAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessSamlJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProvidersAccessSaml) implementsZeroTrustZeroTrustIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []ZeroTrustIdentityProvidersAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IDPPublicCERTs []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                         `json:"sso_target_url"`
	JSON         zeroTrustIdentityProvidersAccessSamlConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessSamlConfigJSON contains the JSON metadata for
// the struct [ZeroTrustIdentityProvidersAccessSamlConfig]
type zeroTrustIdentityProvidersAccessSamlConfigJSON struct {
	Attributes         apijson.Field
	EmailAttributeName apijson.Field
	HeaderAttributes   apijson.Field
	IDPPublicCERTs     apijson.Field
	IssuerURL          apijson.Field
	SignRequest        apijson.Field
	SSOTargetURL       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessSamlConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProvidersAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                        `json:"header_name"`
	JSON       zeroTrustIdentityProvidersAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessSamlConfigHeaderAttributeJSON contains the JSON
// metadata for the struct
// [ZeroTrustIdentityProvidersAccessSamlConfigHeaderAttribute]
type zeroTrustIdentityProvidersAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessSamlConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessSamlType string

const (
	ZeroTrustIdentityProvidersAccessSamlTypeOnetimepin ZeroTrustIdentityProvidersAccessSamlType = "onetimepin"
	ZeroTrustIdentityProvidersAccessSamlTypeAzureAd    ZeroTrustIdentityProvidersAccessSamlType = "azureAD"
	ZeroTrustIdentityProvidersAccessSamlTypeSaml       ZeroTrustIdentityProvidersAccessSamlType = "saml"
	ZeroTrustIdentityProvidersAccessSamlTypeCentrify   ZeroTrustIdentityProvidersAccessSamlType = "centrify"
	ZeroTrustIdentityProvidersAccessSamlTypeFacebook   ZeroTrustIdentityProvidersAccessSamlType = "facebook"
	ZeroTrustIdentityProvidersAccessSamlTypeGitHub     ZeroTrustIdentityProvidersAccessSamlType = "github"
	ZeroTrustIdentityProvidersAccessSamlTypeGoogleApps ZeroTrustIdentityProvidersAccessSamlType = "google-apps"
	ZeroTrustIdentityProvidersAccessSamlTypeGoogle     ZeroTrustIdentityProvidersAccessSamlType = "google"
	ZeroTrustIdentityProvidersAccessSamlTypeLinkedin   ZeroTrustIdentityProvidersAccessSamlType = "linkedin"
	ZeroTrustIdentityProvidersAccessSamlTypeOidc       ZeroTrustIdentityProvidersAccessSamlType = "oidc"
	ZeroTrustIdentityProvidersAccessSamlTypeOkta       ZeroTrustIdentityProvidersAccessSamlType = "okta"
	ZeroTrustIdentityProvidersAccessSamlTypeOnelogin   ZeroTrustIdentityProvidersAccessSamlType = "onelogin"
	ZeroTrustIdentityProvidersAccessSamlTypePingone    ZeroTrustIdentityProvidersAccessSamlType = "pingone"
	ZeroTrustIdentityProvidersAccessSamlTypeYandex     ZeroTrustIdentityProvidersAccessSamlType = "yandex"
)

func (r ZeroTrustIdentityProvidersAccessSamlType) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessSamlTypeOnetimepin, ZeroTrustIdentityProvidersAccessSamlTypeAzureAd, ZeroTrustIdentityProvidersAccessSamlTypeSaml, ZeroTrustIdentityProvidersAccessSamlTypeCentrify, ZeroTrustIdentityProvidersAccessSamlTypeFacebook, ZeroTrustIdentityProvidersAccessSamlTypeGitHub, ZeroTrustIdentityProvidersAccessSamlTypeGoogleApps, ZeroTrustIdentityProvidersAccessSamlTypeGoogle, ZeroTrustIdentityProvidersAccessSamlTypeLinkedin, ZeroTrustIdentityProvidersAccessSamlTypeOidc, ZeroTrustIdentityProvidersAccessSamlTypeOkta, ZeroTrustIdentityProvidersAccessSamlTypeOnelogin, ZeroTrustIdentityProvidersAccessSamlTypePingone, ZeroTrustIdentityProvidersAccessSamlTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProvidersAccessSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProvidersAccessSamlScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessSamlScimConfigJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProvidersAccessSamlScimConfig]
type zeroTrustIdentityProvidersAccessSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessSamlScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProvidersAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config ZeroTrustIdentityProvidersAccessYandexConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProvidersAccessYandexType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProvidersAccessYandexScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProvidersAccessYandexJSON       `json:"-"`
}

// zeroTrustIdentityProvidersAccessYandexJSON contains the JSON metadata for the
// struct [ZeroTrustIdentityProvidersAccessYandex]
type zeroTrustIdentityProvidersAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessYandexJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProvidersAccessYandex) implementsZeroTrustZeroTrustIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                           `json:"client_secret"`
	JSON         zeroTrustIdentityProvidersAccessYandexConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessYandexConfigJSON contains the JSON metadata for
// the struct [ZeroTrustIdentityProvidersAccessYandexConfig]
type zeroTrustIdentityProvidersAccessYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessYandexConfigJSON) RawJSON() string {
	return r.raw
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessYandexType string

const (
	ZeroTrustIdentityProvidersAccessYandexTypeOnetimepin ZeroTrustIdentityProvidersAccessYandexType = "onetimepin"
	ZeroTrustIdentityProvidersAccessYandexTypeAzureAd    ZeroTrustIdentityProvidersAccessYandexType = "azureAD"
	ZeroTrustIdentityProvidersAccessYandexTypeSaml       ZeroTrustIdentityProvidersAccessYandexType = "saml"
	ZeroTrustIdentityProvidersAccessYandexTypeCentrify   ZeroTrustIdentityProvidersAccessYandexType = "centrify"
	ZeroTrustIdentityProvidersAccessYandexTypeFacebook   ZeroTrustIdentityProvidersAccessYandexType = "facebook"
	ZeroTrustIdentityProvidersAccessYandexTypeGitHub     ZeroTrustIdentityProvidersAccessYandexType = "github"
	ZeroTrustIdentityProvidersAccessYandexTypeGoogleApps ZeroTrustIdentityProvidersAccessYandexType = "google-apps"
	ZeroTrustIdentityProvidersAccessYandexTypeGoogle     ZeroTrustIdentityProvidersAccessYandexType = "google"
	ZeroTrustIdentityProvidersAccessYandexTypeLinkedin   ZeroTrustIdentityProvidersAccessYandexType = "linkedin"
	ZeroTrustIdentityProvidersAccessYandexTypeOidc       ZeroTrustIdentityProvidersAccessYandexType = "oidc"
	ZeroTrustIdentityProvidersAccessYandexTypeOkta       ZeroTrustIdentityProvidersAccessYandexType = "okta"
	ZeroTrustIdentityProvidersAccessYandexTypeOnelogin   ZeroTrustIdentityProvidersAccessYandexType = "onelogin"
	ZeroTrustIdentityProvidersAccessYandexTypePingone    ZeroTrustIdentityProvidersAccessYandexType = "pingone"
	ZeroTrustIdentityProvidersAccessYandexTypeYandex     ZeroTrustIdentityProvidersAccessYandexType = "yandex"
)

func (r ZeroTrustIdentityProvidersAccessYandexType) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessYandexTypeOnetimepin, ZeroTrustIdentityProvidersAccessYandexTypeAzureAd, ZeroTrustIdentityProvidersAccessYandexTypeSaml, ZeroTrustIdentityProvidersAccessYandexTypeCentrify, ZeroTrustIdentityProvidersAccessYandexTypeFacebook, ZeroTrustIdentityProvidersAccessYandexTypeGitHub, ZeroTrustIdentityProvidersAccessYandexTypeGoogleApps, ZeroTrustIdentityProvidersAccessYandexTypeGoogle, ZeroTrustIdentityProvidersAccessYandexTypeLinkedin, ZeroTrustIdentityProvidersAccessYandexTypeOidc, ZeroTrustIdentityProvidersAccessYandexTypeOkta, ZeroTrustIdentityProvidersAccessYandexTypeOnelogin, ZeroTrustIdentityProvidersAccessYandexTypePingone, ZeroTrustIdentityProvidersAccessYandexTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProvidersAccessYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProvidersAccessYandexScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessYandexScimConfigJSON contains the JSON metadata
// for the struct [ZeroTrustIdentityProvidersAccessYandexScimConfig]
type zeroTrustIdentityProvidersAccessYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessYandexScimConfigJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustIdentityProvidersAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZeroTrustIdentityProvidersAccessOnetimepinType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZeroTrustIdentityProvidersAccessOnetimepinScimConfig `json:"scim_config"`
	JSON       zeroTrustIdentityProvidersAccessOnetimepinJSON       `json:"-"`
}

// zeroTrustIdentityProvidersAccessOnetimepinJSON contains the JSON metadata for
// the struct [ZeroTrustIdentityProvidersAccessOnetimepin]
type zeroTrustIdentityProvidersAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessOnetimepinJSON) RawJSON() string {
	return r.raw
}

func (r ZeroTrustIdentityProvidersAccessOnetimepin) implementsZeroTrustZeroTrustIdentityProviders() {}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZeroTrustIdentityProvidersAccessOnetimepinType string

const (
	ZeroTrustIdentityProvidersAccessOnetimepinTypeOnetimepin ZeroTrustIdentityProvidersAccessOnetimepinType = "onetimepin"
	ZeroTrustIdentityProvidersAccessOnetimepinTypeAzureAd    ZeroTrustIdentityProvidersAccessOnetimepinType = "azureAD"
	ZeroTrustIdentityProvidersAccessOnetimepinTypeSaml       ZeroTrustIdentityProvidersAccessOnetimepinType = "saml"
	ZeroTrustIdentityProvidersAccessOnetimepinTypeCentrify   ZeroTrustIdentityProvidersAccessOnetimepinType = "centrify"
	ZeroTrustIdentityProvidersAccessOnetimepinTypeFacebook   ZeroTrustIdentityProvidersAccessOnetimepinType = "facebook"
	ZeroTrustIdentityProvidersAccessOnetimepinTypeGitHub     ZeroTrustIdentityProvidersAccessOnetimepinType = "github"
	ZeroTrustIdentityProvidersAccessOnetimepinTypeGoogleApps ZeroTrustIdentityProvidersAccessOnetimepinType = "google-apps"
	ZeroTrustIdentityProvidersAccessOnetimepinTypeGoogle     ZeroTrustIdentityProvidersAccessOnetimepinType = "google"
	ZeroTrustIdentityProvidersAccessOnetimepinTypeLinkedin   ZeroTrustIdentityProvidersAccessOnetimepinType = "linkedin"
	ZeroTrustIdentityProvidersAccessOnetimepinTypeOidc       ZeroTrustIdentityProvidersAccessOnetimepinType = "oidc"
	ZeroTrustIdentityProvidersAccessOnetimepinTypeOkta       ZeroTrustIdentityProvidersAccessOnetimepinType = "okta"
	ZeroTrustIdentityProvidersAccessOnetimepinTypeOnelogin   ZeroTrustIdentityProvidersAccessOnetimepinType = "onelogin"
	ZeroTrustIdentityProvidersAccessOnetimepinTypePingone    ZeroTrustIdentityProvidersAccessOnetimepinType = "pingone"
	ZeroTrustIdentityProvidersAccessOnetimepinTypeYandex     ZeroTrustIdentityProvidersAccessOnetimepinType = "yandex"
)

func (r ZeroTrustIdentityProvidersAccessOnetimepinType) IsKnown() bool {
	switch r {
	case ZeroTrustIdentityProvidersAccessOnetimepinTypeOnetimepin, ZeroTrustIdentityProvidersAccessOnetimepinTypeAzureAd, ZeroTrustIdentityProvidersAccessOnetimepinTypeSaml, ZeroTrustIdentityProvidersAccessOnetimepinTypeCentrify, ZeroTrustIdentityProvidersAccessOnetimepinTypeFacebook, ZeroTrustIdentityProvidersAccessOnetimepinTypeGitHub, ZeroTrustIdentityProvidersAccessOnetimepinTypeGoogleApps, ZeroTrustIdentityProvidersAccessOnetimepinTypeGoogle, ZeroTrustIdentityProvidersAccessOnetimepinTypeLinkedin, ZeroTrustIdentityProvidersAccessOnetimepinTypeOidc, ZeroTrustIdentityProvidersAccessOnetimepinTypeOkta, ZeroTrustIdentityProvidersAccessOnetimepinTypeOnelogin, ZeroTrustIdentityProvidersAccessOnetimepinTypePingone, ZeroTrustIdentityProvidersAccessOnetimepinTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZeroTrustIdentityProvidersAccessOnetimepinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zeroTrustIdentityProvidersAccessOnetimepinScimConfigJSON `json:"-"`
}

// zeroTrustIdentityProvidersAccessOnetimepinScimConfigJSON contains the JSON
// metadata for the struct [ZeroTrustIdentityProvidersAccessOnetimepinScimConfig]
type zeroTrustIdentityProvidersAccessOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZeroTrustIdentityProvidersAccessOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustIdentityProvidersAccessOnetimepinScimConfigJSON) RawJSON() string {
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
	// Indicates the type of user interaction that is required. prompt=login forces the
	// user to enter their credentials on that request, negating single-sign on.
	// prompt=none is the opposite. It ensures that the user isn't presented with any
	// interactive prompt. If the request can't be completed silently by using
	// single-sign on, the Microsoft identity platform returns an interaction_required
	// error. prompt=select_account interrupts single sign-on providing account
	// selection experience listing all the accounts either in session or any
	// remembered account or an option to choose to use a different account altogether.
	Prompt IdentityProviderListResponseAccessAzureAdConfigPrompt `json:"prompt"`
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
	Prompt                   apijson.Field
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

// Indicates the type of user interaction that is required. prompt=login forces the
// user to enter their credentials on that request, negating single-sign on.
// prompt=none is the opposite. It ensures that the user isn't presented with any
// interactive prompt. If the request can't be completed silently by using
// single-sign on, the Microsoft identity platform returns an interaction_required
// error. prompt=select_account interrupts single sign-on providing account
// selection experience listing all the accounts either in session or any
// remembered account or an option to choose to use a different account altogether.
type IdentityProviderListResponseAccessAzureAdConfigPrompt string

const (
	IdentityProviderListResponseAccessAzureAdConfigPromptLogin         IdentityProviderListResponseAccessAzureAdConfigPrompt = "login"
	IdentityProviderListResponseAccessAzureAdConfigPromptSelectAccount IdentityProviderListResponseAccessAzureAdConfigPrompt = "select_account"
	IdentityProviderListResponseAccessAzureAdConfigPromptNone          IdentityProviderListResponseAccessAzureAdConfigPrompt = "none"
)

func (r IdentityProviderListResponseAccessAzureAdConfigPrompt) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseAccessAzureAdConfigPromptLogin, IdentityProviderListResponseAccessAzureAdConfigPromptSelectAccount, IdentityProviderListResponseAccessAzureAdConfigPromptNone:
		return true
	}
	return false
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

func (r IdentityProviderListResponseAccessAzureAdType) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseAccessAzureAdTypeOnetimepin, IdentityProviderListResponseAccessAzureAdTypeAzureAd, IdentityProviderListResponseAccessAzureAdTypeSaml, IdentityProviderListResponseAccessAzureAdTypeCentrify, IdentityProviderListResponseAccessAzureAdTypeFacebook, IdentityProviderListResponseAccessAzureAdTypeGitHub, IdentityProviderListResponseAccessAzureAdTypeGoogleApps, IdentityProviderListResponseAccessAzureAdTypeGoogle, IdentityProviderListResponseAccessAzureAdTypeLinkedin, IdentityProviderListResponseAccessAzureAdTypeOidc, IdentityProviderListResponseAccessAzureAdTypeOkta, IdentityProviderListResponseAccessAzureAdTypeOnelogin, IdentityProviderListResponseAccessAzureAdTypePingone, IdentityProviderListResponseAccessAzureAdTypeYandex:
		return true
	}
	return false
}

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

func (r IdentityProviderListResponseAccessCentrifyType) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseAccessCentrifyTypeOnetimepin, IdentityProviderListResponseAccessCentrifyTypeAzureAd, IdentityProviderListResponseAccessCentrifyTypeSaml, IdentityProviderListResponseAccessCentrifyTypeCentrify, IdentityProviderListResponseAccessCentrifyTypeFacebook, IdentityProviderListResponseAccessCentrifyTypeGitHub, IdentityProviderListResponseAccessCentrifyTypeGoogleApps, IdentityProviderListResponseAccessCentrifyTypeGoogle, IdentityProviderListResponseAccessCentrifyTypeLinkedin, IdentityProviderListResponseAccessCentrifyTypeOidc, IdentityProviderListResponseAccessCentrifyTypeOkta, IdentityProviderListResponseAccessCentrifyTypeOnelogin, IdentityProviderListResponseAccessCentrifyTypePingone, IdentityProviderListResponseAccessCentrifyTypeYandex:
		return true
	}
	return false
}

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

func (r IdentityProviderListResponseAccessFacebookType) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseAccessFacebookTypeOnetimepin, IdentityProviderListResponseAccessFacebookTypeAzureAd, IdentityProviderListResponseAccessFacebookTypeSaml, IdentityProviderListResponseAccessFacebookTypeCentrify, IdentityProviderListResponseAccessFacebookTypeFacebook, IdentityProviderListResponseAccessFacebookTypeGitHub, IdentityProviderListResponseAccessFacebookTypeGoogleApps, IdentityProviderListResponseAccessFacebookTypeGoogle, IdentityProviderListResponseAccessFacebookTypeLinkedin, IdentityProviderListResponseAccessFacebookTypeOidc, IdentityProviderListResponseAccessFacebookTypeOkta, IdentityProviderListResponseAccessFacebookTypeOnelogin, IdentityProviderListResponseAccessFacebookTypePingone, IdentityProviderListResponseAccessFacebookTypeYandex:
		return true
	}
	return false
}

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

func (r IdentityProviderListResponseAccessGitHubType) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseAccessGitHubTypeOnetimepin, IdentityProviderListResponseAccessGitHubTypeAzureAd, IdentityProviderListResponseAccessGitHubTypeSaml, IdentityProviderListResponseAccessGitHubTypeCentrify, IdentityProviderListResponseAccessGitHubTypeFacebook, IdentityProviderListResponseAccessGitHubTypeGitHub, IdentityProviderListResponseAccessGitHubTypeGoogleApps, IdentityProviderListResponseAccessGitHubTypeGoogle, IdentityProviderListResponseAccessGitHubTypeLinkedin, IdentityProviderListResponseAccessGitHubTypeOidc, IdentityProviderListResponseAccessGitHubTypeOkta, IdentityProviderListResponseAccessGitHubTypeOnelogin, IdentityProviderListResponseAccessGitHubTypePingone, IdentityProviderListResponseAccessGitHubTypeYandex:
		return true
	}
	return false
}

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

func (r IdentityProviderListResponseAccessGoogleType) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseAccessGoogleTypeOnetimepin, IdentityProviderListResponseAccessGoogleTypeAzureAd, IdentityProviderListResponseAccessGoogleTypeSaml, IdentityProviderListResponseAccessGoogleTypeCentrify, IdentityProviderListResponseAccessGoogleTypeFacebook, IdentityProviderListResponseAccessGoogleTypeGitHub, IdentityProviderListResponseAccessGoogleTypeGoogleApps, IdentityProviderListResponseAccessGoogleTypeGoogle, IdentityProviderListResponseAccessGoogleTypeLinkedin, IdentityProviderListResponseAccessGoogleTypeOidc, IdentityProviderListResponseAccessGoogleTypeOkta, IdentityProviderListResponseAccessGoogleTypeOnelogin, IdentityProviderListResponseAccessGoogleTypePingone, IdentityProviderListResponseAccessGoogleTypeYandex:
		return true
	}
	return false
}

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

func (r IdentityProviderListResponseAccessGoogleAppsType) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseAccessGoogleAppsTypeOnetimepin, IdentityProviderListResponseAccessGoogleAppsTypeAzureAd, IdentityProviderListResponseAccessGoogleAppsTypeSaml, IdentityProviderListResponseAccessGoogleAppsTypeCentrify, IdentityProviderListResponseAccessGoogleAppsTypeFacebook, IdentityProviderListResponseAccessGoogleAppsTypeGitHub, IdentityProviderListResponseAccessGoogleAppsTypeGoogleApps, IdentityProviderListResponseAccessGoogleAppsTypeGoogle, IdentityProviderListResponseAccessGoogleAppsTypeLinkedin, IdentityProviderListResponseAccessGoogleAppsTypeOidc, IdentityProviderListResponseAccessGoogleAppsTypeOkta, IdentityProviderListResponseAccessGoogleAppsTypeOnelogin, IdentityProviderListResponseAccessGoogleAppsTypePingone, IdentityProviderListResponseAccessGoogleAppsTypeYandex:
		return true
	}
	return false
}

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

func (r IdentityProviderListResponseAccessLinkedinType) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseAccessLinkedinTypeOnetimepin, IdentityProviderListResponseAccessLinkedinTypeAzureAd, IdentityProviderListResponseAccessLinkedinTypeSaml, IdentityProviderListResponseAccessLinkedinTypeCentrify, IdentityProviderListResponseAccessLinkedinTypeFacebook, IdentityProviderListResponseAccessLinkedinTypeGitHub, IdentityProviderListResponseAccessLinkedinTypeGoogleApps, IdentityProviderListResponseAccessLinkedinTypeGoogle, IdentityProviderListResponseAccessLinkedinTypeLinkedin, IdentityProviderListResponseAccessLinkedinTypeOidc, IdentityProviderListResponseAccessLinkedinTypeOkta, IdentityProviderListResponseAccessLinkedinTypeOnelogin, IdentityProviderListResponseAccessLinkedinTypePingone, IdentityProviderListResponseAccessLinkedinTypeYandex:
		return true
	}
	return false
}

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
	CERTsURL string `json:"certs_url"`
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
	CERTsURL       apijson.Field
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

func (r IdentityProviderListResponseAccessOidcType) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseAccessOidcTypeOnetimepin, IdentityProviderListResponseAccessOidcTypeAzureAd, IdentityProviderListResponseAccessOidcTypeSaml, IdentityProviderListResponseAccessOidcTypeCentrify, IdentityProviderListResponseAccessOidcTypeFacebook, IdentityProviderListResponseAccessOidcTypeGitHub, IdentityProviderListResponseAccessOidcTypeGoogleApps, IdentityProviderListResponseAccessOidcTypeGoogle, IdentityProviderListResponseAccessOidcTypeLinkedin, IdentityProviderListResponseAccessOidcTypeOidc, IdentityProviderListResponseAccessOidcTypeOkta, IdentityProviderListResponseAccessOidcTypeOnelogin, IdentityProviderListResponseAccessOidcTypePingone, IdentityProviderListResponseAccessOidcTypeYandex:
		return true
	}
	return false
}

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

func (r IdentityProviderListResponseAccessOktaType) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseAccessOktaTypeOnetimepin, IdentityProviderListResponseAccessOktaTypeAzureAd, IdentityProviderListResponseAccessOktaTypeSaml, IdentityProviderListResponseAccessOktaTypeCentrify, IdentityProviderListResponseAccessOktaTypeFacebook, IdentityProviderListResponseAccessOktaTypeGitHub, IdentityProviderListResponseAccessOktaTypeGoogleApps, IdentityProviderListResponseAccessOktaTypeGoogle, IdentityProviderListResponseAccessOktaTypeLinkedin, IdentityProviderListResponseAccessOktaTypeOidc, IdentityProviderListResponseAccessOktaTypeOkta, IdentityProviderListResponseAccessOktaTypeOnelogin, IdentityProviderListResponseAccessOktaTypePingone, IdentityProviderListResponseAccessOktaTypeYandex:
		return true
	}
	return false
}

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

func (r IdentityProviderListResponseAccessOneloginType) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseAccessOneloginTypeOnetimepin, IdentityProviderListResponseAccessOneloginTypeAzureAd, IdentityProviderListResponseAccessOneloginTypeSaml, IdentityProviderListResponseAccessOneloginTypeCentrify, IdentityProviderListResponseAccessOneloginTypeFacebook, IdentityProviderListResponseAccessOneloginTypeGitHub, IdentityProviderListResponseAccessOneloginTypeGoogleApps, IdentityProviderListResponseAccessOneloginTypeGoogle, IdentityProviderListResponseAccessOneloginTypeLinkedin, IdentityProviderListResponseAccessOneloginTypeOidc, IdentityProviderListResponseAccessOneloginTypeOkta, IdentityProviderListResponseAccessOneloginTypeOnelogin, IdentityProviderListResponseAccessOneloginTypePingone, IdentityProviderListResponseAccessOneloginTypeYandex:
		return true
	}
	return false
}

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

func (r IdentityProviderListResponseAccessPingoneType) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseAccessPingoneTypeOnetimepin, IdentityProviderListResponseAccessPingoneTypeAzureAd, IdentityProviderListResponseAccessPingoneTypeSaml, IdentityProviderListResponseAccessPingoneTypeCentrify, IdentityProviderListResponseAccessPingoneTypeFacebook, IdentityProviderListResponseAccessPingoneTypeGitHub, IdentityProviderListResponseAccessPingoneTypeGoogleApps, IdentityProviderListResponseAccessPingoneTypeGoogle, IdentityProviderListResponseAccessPingoneTypeLinkedin, IdentityProviderListResponseAccessPingoneTypeOidc, IdentityProviderListResponseAccessPingoneTypeOkta, IdentityProviderListResponseAccessPingoneTypeOnelogin, IdentityProviderListResponseAccessPingoneTypePingone, IdentityProviderListResponseAccessPingoneTypeYandex:
		return true
	}
	return false
}

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
	IDPPublicCERTs []string `json:"idp_public_certs"`
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
	IDPPublicCERTs     apijson.Field
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

func (r IdentityProviderListResponseAccessSamlType) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseAccessSamlTypeOnetimepin, IdentityProviderListResponseAccessSamlTypeAzureAd, IdentityProviderListResponseAccessSamlTypeSaml, IdentityProviderListResponseAccessSamlTypeCentrify, IdentityProviderListResponseAccessSamlTypeFacebook, IdentityProviderListResponseAccessSamlTypeGitHub, IdentityProviderListResponseAccessSamlTypeGoogleApps, IdentityProviderListResponseAccessSamlTypeGoogle, IdentityProviderListResponseAccessSamlTypeLinkedin, IdentityProviderListResponseAccessSamlTypeOidc, IdentityProviderListResponseAccessSamlTypeOkta, IdentityProviderListResponseAccessSamlTypeOnelogin, IdentityProviderListResponseAccessSamlTypePingone, IdentityProviderListResponseAccessSamlTypeYandex:
		return true
	}
	return false
}

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

func (r IdentityProviderListResponseAccessYandexType) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseAccessYandexTypeOnetimepin, IdentityProviderListResponseAccessYandexTypeAzureAd, IdentityProviderListResponseAccessYandexTypeSaml, IdentityProviderListResponseAccessYandexTypeCentrify, IdentityProviderListResponseAccessYandexTypeFacebook, IdentityProviderListResponseAccessYandexTypeGitHub, IdentityProviderListResponseAccessYandexTypeGoogleApps, IdentityProviderListResponseAccessYandexTypeGoogle, IdentityProviderListResponseAccessYandexTypeLinkedin, IdentityProviderListResponseAccessYandexTypeOidc, IdentityProviderListResponseAccessYandexTypeOkta, IdentityProviderListResponseAccessYandexTypeOnelogin, IdentityProviderListResponseAccessYandexTypePingone, IdentityProviderListResponseAccessYandexTypeYandex:
		return true
	}
	return false
}

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

// This interface is a union satisfied by one of the following:
// [IdentityProviderNewParamsAccessAzureAd],
// [IdentityProviderNewParamsAccessCentrify],
// [IdentityProviderNewParamsAccessFacebook],
// [IdentityProviderNewParamsAccessGitHub],
// [IdentityProviderNewParamsAccessGoogle],
// [IdentityProviderNewParamsAccessGoogleApps],
// [IdentityProviderNewParamsAccessLinkedin],
// [IdentityProviderNewParamsAccessOidc], [IdentityProviderNewParamsAccessOkta],
// [IdentityProviderNewParamsAccessOnelogin],
// [IdentityProviderNewParamsAccessPingone], [IdentityProviderNewParamsAccessSaml],
// [IdentityProviderNewParamsAccessYandex],
// [IdentityProviderNewParamsAccessOnetimepin].
type IdentityProviderNewParams interface {
	ImplementsIdentityProviderNewParams()

	getAccountID() param.Field[string]

	getZoneID() param.Field[string]
}

type IdentityProviderNewParamsAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsAccessAzureAdConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsAccessAzureAdType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderNewParamsAccessAzureAdScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAccessAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAccessAzureAd) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAccessAzureAd) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAccessAzureAd) ImplementsIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessAzureAdConfig struct {
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
	// Indicates the type of user interaction that is required. prompt=login forces the
	// user to enter their credentials on that request, negating single-sign on.
	// prompt=none is the opposite. It ensures that the user isn't presented with any
	// interactive prompt. If the request can't be completed silently by using
	// single-sign on, the Microsoft identity platform returns an interaction_required
	// error. prompt=select_account interrupts single sign-on providing account
	// selection experience listing all the accounts either in session or any
	// remembered account or an option to choose to use a different account altogether.
	Prompt param.Field[IdentityProviderNewParamsAccessAzureAdConfigPrompt] `json:"prompt"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
}

func (r IdentityProviderNewParamsAccessAzureAdConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Indicates the type of user interaction that is required. prompt=login forces the
// user to enter their credentials on that request, negating single-sign on.
// prompt=none is the opposite. It ensures that the user isn't presented with any
// interactive prompt. If the request can't be completed silently by using
// single-sign on, the Microsoft identity platform returns an interaction_required
// error. prompt=select_account interrupts single sign-on providing account
// selection experience listing all the accounts either in session or any
// remembered account or an option to choose to use a different account altogether.
type IdentityProviderNewParamsAccessAzureAdConfigPrompt string

const (
	IdentityProviderNewParamsAccessAzureAdConfigPromptLogin         IdentityProviderNewParamsAccessAzureAdConfigPrompt = "login"
	IdentityProviderNewParamsAccessAzureAdConfigPromptSelectAccount IdentityProviderNewParamsAccessAzureAdConfigPrompt = "select_account"
	IdentityProviderNewParamsAccessAzureAdConfigPromptNone          IdentityProviderNewParamsAccessAzureAdConfigPrompt = "none"
)

func (r IdentityProviderNewParamsAccessAzureAdConfigPrompt) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessAzureAdConfigPromptLogin, IdentityProviderNewParamsAccessAzureAdConfigPromptSelectAccount, IdentityProviderNewParamsAccessAzureAdConfigPromptNone:
		return true
	}
	return false
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessAzureAdType string

const (
	IdentityProviderNewParamsAccessAzureAdTypeOnetimepin IdentityProviderNewParamsAccessAzureAdType = "onetimepin"
	IdentityProviderNewParamsAccessAzureAdTypeAzureAd    IdentityProviderNewParamsAccessAzureAdType = "azureAD"
	IdentityProviderNewParamsAccessAzureAdTypeSaml       IdentityProviderNewParamsAccessAzureAdType = "saml"
	IdentityProviderNewParamsAccessAzureAdTypeCentrify   IdentityProviderNewParamsAccessAzureAdType = "centrify"
	IdentityProviderNewParamsAccessAzureAdTypeFacebook   IdentityProviderNewParamsAccessAzureAdType = "facebook"
	IdentityProviderNewParamsAccessAzureAdTypeGitHub     IdentityProviderNewParamsAccessAzureAdType = "github"
	IdentityProviderNewParamsAccessAzureAdTypeGoogleApps IdentityProviderNewParamsAccessAzureAdType = "google-apps"
	IdentityProviderNewParamsAccessAzureAdTypeGoogle     IdentityProviderNewParamsAccessAzureAdType = "google"
	IdentityProviderNewParamsAccessAzureAdTypeLinkedin   IdentityProviderNewParamsAccessAzureAdType = "linkedin"
	IdentityProviderNewParamsAccessAzureAdTypeOidc       IdentityProviderNewParamsAccessAzureAdType = "oidc"
	IdentityProviderNewParamsAccessAzureAdTypeOkta       IdentityProviderNewParamsAccessAzureAdType = "okta"
	IdentityProviderNewParamsAccessAzureAdTypeOnelogin   IdentityProviderNewParamsAccessAzureAdType = "onelogin"
	IdentityProviderNewParamsAccessAzureAdTypePingone    IdentityProviderNewParamsAccessAzureAdType = "pingone"
	IdentityProviderNewParamsAccessAzureAdTypeYandex     IdentityProviderNewParamsAccessAzureAdType = "yandex"
)

func (r IdentityProviderNewParamsAccessAzureAdType) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessAzureAdTypeOnetimepin, IdentityProviderNewParamsAccessAzureAdTypeAzureAd, IdentityProviderNewParamsAccessAzureAdTypeSaml, IdentityProviderNewParamsAccessAzureAdTypeCentrify, IdentityProviderNewParamsAccessAzureAdTypeFacebook, IdentityProviderNewParamsAccessAzureAdTypeGitHub, IdentityProviderNewParamsAccessAzureAdTypeGoogleApps, IdentityProviderNewParamsAccessAzureAdTypeGoogle, IdentityProviderNewParamsAccessAzureAdTypeLinkedin, IdentityProviderNewParamsAccessAzureAdTypeOidc, IdentityProviderNewParamsAccessAzureAdTypeOkta, IdentityProviderNewParamsAccessAzureAdTypeOnelogin, IdentityProviderNewParamsAccessAzureAdTypePingone, IdentityProviderNewParamsAccessAzureAdTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewParamsAccessAzureAdScimConfig struct {
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

func (r IdentityProviderNewParamsAccessAzureAdScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsAccessCentrifyConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsAccessCentrifyType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderNewParamsAccessCentrifyScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAccessCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAccessCentrify) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAccessCentrify) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAccessCentrify) ImplementsIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessCentrifyConfig struct {
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

func (r IdentityProviderNewParamsAccessCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessCentrifyType string

const (
	IdentityProviderNewParamsAccessCentrifyTypeOnetimepin IdentityProviderNewParamsAccessCentrifyType = "onetimepin"
	IdentityProviderNewParamsAccessCentrifyTypeAzureAd    IdentityProviderNewParamsAccessCentrifyType = "azureAD"
	IdentityProviderNewParamsAccessCentrifyTypeSaml       IdentityProviderNewParamsAccessCentrifyType = "saml"
	IdentityProviderNewParamsAccessCentrifyTypeCentrify   IdentityProviderNewParamsAccessCentrifyType = "centrify"
	IdentityProviderNewParamsAccessCentrifyTypeFacebook   IdentityProviderNewParamsAccessCentrifyType = "facebook"
	IdentityProviderNewParamsAccessCentrifyTypeGitHub     IdentityProviderNewParamsAccessCentrifyType = "github"
	IdentityProviderNewParamsAccessCentrifyTypeGoogleApps IdentityProviderNewParamsAccessCentrifyType = "google-apps"
	IdentityProviderNewParamsAccessCentrifyTypeGoogle     IdentityProviderNewParamsAccessCentrifyType = "google"
	IdentityProviderNewParamsAccessCentrifyTypeLinkedin   IdentityProviderNewParamsAccessCentrifyType = "linkedin"
	IdentityProviderNewParamsAccessCentrifyTypeOidc       IdentityProviderNewParamsAccessCentrifyType = "oidc"
	IdentityProviderNewParamsAccessCentrifyTypeOkta       IdentityProviderNewParamsAccessCentrifyType = "okta"
	IdentityProviderNewParamsAccessCentrifyTypeOnelogin   IdentityProviderNewParamsAccessCentrifyType = "onelogin"
	IdentityProviderNewParamsAccessCentrifyTypePingone    IdentityProviderNewParamsAccessCentrifyType = "pingone"
	IdentityProviderNewParamsAccessCentrifyTypeYandex     IdentityProviderNewParamsAccessCentrifyType = "yandex"
)

func (r IdentityProviderNewParamsAccessCentrifyType) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessCentrifyTypeOnetimepin, IdentityProviderNewParamsAccessCentrifyTypeAzureAd, IdentityProviderNewParamsAccessCentrifyTypeSaml, IdentityProviderNewParamsAccessCentrifyTypeCentrify, IdentityProviderNewParamsAccessCentrifyTypeFacebook, IdentityProviderNewParamsAccessCentrifyTypeGitHub, IdentityProviderNewParamsAccessCentrifyTypeGoogleApps, IdentityProviderNewParamsAccessCentrifyTypeGoogle, IdentityProviderNewParamsAccessCentrifyTypeLinkedin, IdentityProviderNewParamsAccessCentrifyTypeOidc, IdentityProviderNewParamsAccessCentrifyTypeOkta, IdentityProviderNewParamsAccessCentrifyTypeOnelogin, IdentityProviderNewParamsAccessCentrifyTypePingone, IdentityProviderNewParamsAccessCentrifyTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewParamsAccessCentrifyScimConfig struct {
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

func (r IdentityProviderNewParamsAccessCentrifyScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsAccessFacebookConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsAccessFacebookType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderNewParamsAccessFacebookScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAccessFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAccessFacebook) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAccessFacebook) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAccessFacebook) ImplementsIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r IdentityProviderNewParamsAccessFacebookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessFacebookType string

const (
	IdentityProviderNewParamsAccessFacebookTypeOnetimepin IdentityProviderNewParamsAccessFacebookType = "onetimepin"
	IdentityProviderNewParamsAccessFacebookTypeAzureAd    IdentityProviderNewParamsAccessFacebookType = "azureAD"
	IdentityProviderNewParamsAccessFacebookTypeSaml       IdentityProviderNewParamsAccessFacebookType = "saml"
	IdentityProviderNewParamsAccessFacebookTypeCentrify   IdentityProviderNewParamsAccessFacebookType = "centrify"
	IdentityProviderNewParamsAccessFacebookTypeFacebook   IdentityProviderNewParamsAccessFacebookType = "facebook"
	IdentityProviderNewParamsAccessFacebookTypeGitHub     IdentityProviderNewParamsAccessFacebookType = "github"
	IdentityProviderNewParamsAccessFacebookTypeGoogleApps IdentityProviderNewParamsAccessFacebookType = "google-apps"
	IdentityProviderNewParamsAccessFacebookTypeGoogle     IdentityProviderNewParamsAccessFacebookType = "google"
	IdentityProviderNewParamsAccessFacebookTypeLinkedin   IdentityProviderNewParamsAccessFacebookType = "linkedin"
	IdentityProviderNewParamsAccessFacebookTypeOidc       IdentityProviderNewParamsAccessFacebookType = "oidc"
	IdentityProviderNewParamsAccessFacebookTypeOkta       IdentityProviderNewParamsAccessFacebookType = "okta"
	IdentityProviderNewParamsAccessFacebookTypeOnelogin   IdentityProviderNewParamsAccessFacebookType = "onelogin"
	IdentityProviderNewParamsAccessFacebookTypePingone    IdentityProviderNewParamsAccessFacebookType = "pingone"
	IdentityProviderNewParamsAccessFacebookTypeYandex     IdentityProviderNewParamsAccessFacebookType = "yandex"
)

func (r IdentityProviderNewParamsAccessFacebookType) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessFacebookTypeOnetimepin, IdentityProviderNewParamsAccessFacebookTypeAzureAd, IdentityProviderNewParamsAccessFacebookTypeSaml, IdentityProviderNewParamsAccessFacebookTypeCentrify, IdentityProviderNewParamsAccessFacebookTypeFacebook, IdentityProviderNewParamsAccessFacebookTypeGitHub, IdentityProviderNewParamsAccessFacebookTypeGoogleApps, IdentityProviderNewParamsAccessFacebookTypeGoogle, IdentityProviderNewParamsAccessFacebookTypeLinkedin, IdentityProviderNewParamsAccessFacebookTypeOidc, IdentityProviderNewParamsAccessFacebookTypeOkta, IdentityProviderNewParamsAccessFacebookTypeOnelogin, IdentityProviderNewParamsAccessFacebookTypePingone, IdentityProviderNewParamsAccessFacebookTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewParamsAccessFacebookScimConfig struct {
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

func (r IdentityProviderNewParamsAccessFacebookScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsAccessGitHubConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsAccessGitHubType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderNewParamsAccessGitHubScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAccessGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAccessGitHub) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAccessGitHub) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAccessGitHub) ImplementsIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r IdentityProviderNewParamsAccessGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessGitHubType string

const (
	IdentityProviderNewParamsAccessGitHubTypeOnetimepin IdentityProviderNewParamsAccessGitHubType = "onetimepin"
	IdentityProviderNewParamsAccessGitHubTypeAzureAd    IdentityProviderNewParamsAccessGitHubType = "azureAD"
	IdentityProviderNewParamsAccessGitHubTypeSaml       IdentityProviderNewParamsAccessGitHubType = "saml"
	IdentityProviderNewParamsAccessGitHubTypeCentrify   IdentityProviderNewParamsAccessGitHubType = "centrify"
	IdentityProviderNewParamsAccessGitHubTypeFacebook   IdentityProviderNewParamsAccessGitHubType = "facebook"
	IdentityProviderNewParamsAccessGitHubTypeGitHub     IdentityProviderNewParamsAccessGitHubType = "github"
	IdentityProviderNewParamsAccessGitHubTypeGoogleApps IdentityProviderNewParamsAccessGitHubType = "google-apps"
	IdentityProviderNewParamsAccessGitHubTypeGoogle     IdentityProviderNewParamsAccessGitHubType = "google"
	IdentityProviderNewParamsAccessGitHubTypeLinkedin   IdentityProviderNewParamsAccessGitHubType = "linkedin"
	IdentityProviderNewParamsAccessGitHubTypeOidc       IdentityProviderNewParamsAccessGitHubType = "oidc"
	IdentityProviderNewParamsAccessGitHubTypeOkta       IdentityProviderNewParamsAccessGitHubType = "okta"
	IdentityProviderNewParamsAccessGitHubTypeOnelogin   IdentityProviderNewParamsAccessGitHubType = "onelogin"
	IdentityProviderNewParamsAccessGitHubTypePingone    IdentityProviderNewParamsAccessGitHubType = "pingone"
	IdentityProviderNewParamsAccessGitHubTypeYandex     IdentityProviderNewParamsAccessGitHubType = "yandex"
)

func (r IdentityProviderNewParamsAccessGitHubType) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessGitHubTypeOnetimepin, IdentityProviderNewParamsAccessGitHubTypeAzureAd, IdentityProviderNewParamsAccessGitHubTypeSaml, IdentityProviderNewParamsAccessGitHubTypeCentrify, IdentityProviderNewParamsAccessGitHubTypeFacebook, IdentityProviderNewParamsAccessGitHubTypeGitHub, IdentityProviderNewParamsAccessGitHubTypeGoogleApps, IdentityProviderNewParamsAccessGitHubTypeGoogle, IdentityProviderNewParamsAccessGitHubTypeLinkedin, IdentityProviderNewParamsAccessGitHubTypeOidc, IdentityProviderNewParamsAccessGitHubTypeOkta, IdentityProviderNewParamsAccessGitHubTypeOnelogin, IdentityProviderNewParamsAccessGitHubTypePingone, IdentityProviderNewParamsAccessGitHubTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewParamsAccessGitHubScimConfig struct {
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

func (r IdentityProviderNewParamsAccessGitHubScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsAccessGoogleConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsAccessGoogleType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderNewParamsAccessGoogleScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAccessGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAccessGoogle) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAccessGoogle) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAccessGoogle) ImplementsIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessGoogleConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r IdentityProviderNewParamsAccessGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessGoogleType string

const (
	IdentityProviderNewParamsAccessGoogleTypeOnetimepin IdentityProviderNewParamsAccessGoogleType = "onetimepin"
	IdentityProviderNewParamsAccessGoogleTypeAzureAd    IdentityProviderNewParamsAccessGoogleType = "azureAD"
	IdentityProviderNewParamsAccessGoogleTypeSaml       IdentityProviderNewParamsAccessGoogleType = "saml"
	IdentityProviderNewParamsAccessGoogleTypeCentrify   IdentityProviderNewParamsAccessGoogleType = "centrify"
	IdentityProviderNewParamsAccessGoogleTypeFacebook   IdentityProviderNewParamsAccessGoogleType = "facebook"
	IdentityProviderNewParamsAccessGoogleTypeGitHub     IdentityProviderNewParamsAccessGoogleType = "github"
	IdentityProviderNewParamsAccessGoogleTypeGoogleApps IdentityProviderNewParamsAccessGoogleType = "google-apps"
	IdentityProviderNewParamsAccessGoogleTypeGoogle     IdentityProviderNewParamsAccessGoogleType = "google"
	IdentityProviderNewParamsAccessGoogleTypeLinkedin   IdentityProviderNewParamsAccessGoogleType = "linkedin"
	IdentityProviderNewParamsAccessGoogleTypeOidc       IdentityProviderNewParamsAccessGoogleType = "oidc"
	IdentityProviderNewParamsAccessGoogleTypeOkta       IdentityProviderNewParamsAccessGoogleType = "okta"
	IdentityProviderNewParamsAccessGoogleTypeOnelogin   IdentityProviderNewParamsAccessGoogleType = "onelogin"
	IdentityProviderNewParamsAccessGoogleTypePingone    IdentityProviderNewParamsAccessGoogleType = "pingone"
	IdentityProviderNewParamsAccessGoogleTypeYandex     IdentityProviderNewParamsAccessGoogleType = "yandex"
)

func (r IdentityProviderNewParamsAccessGoogleType) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessGoogleTypeOnetimepin, IdentityProviderNewParamsAccessGoogleTypeAzureAd, IdentityProviderNewParamsAccessGoogleTypeSaml, IdentityProviderNewParamsAccessGoogleTypeCentrify, IdentityProviderNewParamsAccessGoogleTypeFacebook, IdentityProviderNewParamsAccessGoogleTypeGitHub, IdentityProviderNewParamsAccessGoogleTypeGoogleApps, IdentityProviderNewParamsAccessGoogleTypeGoogle, IdentityProviderNewParamsAccessGoogleTypeLinkedin, IdentityProviderNewParamsAccessGoogleTypeOidc, IdentityProviderNewParamsAccessGoogleTypeOkta, IdentityProviderNewParamsAccessGoogleTypeOnelogin, IdentityProviderNewParamsAccessGoogleTypePingone, IdentityProviderNewParamsAccessGoogleTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewParamsAccessGoogleScimConfig struct {
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

func (r IdentityProviderNewParamsAccessGoogleScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsAccessGoogleAppsConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsAccessGoogleAppsType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderNewParamsAccessGoogleAppsScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAccessGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAccessGoogleApps) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAccessGoogleApps) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAccessGoogleApps) ImplementsIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessGoogleAppsConfig struct {
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

func (r IdentityProviderNewParamsAccessGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessGoogleAppsType string

const (
	IdentityProviderNewParamsAccessGoogleAppsTypeOnetimepin IdentityProviderNewParamsAccessGoogleAppsType = "onetimepin"
	IdentityProviderNewParamsAccessGoogleAppsTypeAzureAd    IdentityProviderNewParamsAccessGoogleAppsType = "azureAD"
	IdentityProviderNewParamsAccessGoogleAppsTypeSaml       IdentityProviderNewParamsAccessGoogleAppsType = "saml"
	IdentityProviderNewParamsAccessGoogleAppsTypeCentrify   IdentityProviderNewParamsAccessGoogleAppsType = "centrify"
	IdentityProviderNewParamsAccessGoogleAppsTypeFacebook   IdentityProviderNewParamsAccessGoogleAppsType = "facebook"
	IdentityProviderNewParamsAccessGoogleAppsTypeGitHub     IdentityProviderNewParamsAccessGoogleAppsType = "github"
	IdentityProviderNewParamsAccessGoogleAppsTypeGoogleApps IdentityProviderNewParamsAccessGoogleAppsType = "google-apps"
	IdentityProviderNewParamsAccessGoogleAppsTypeGoogle     IdentityProviderNewParamsAccessGoogleAppsType = "google"
	IdentityProviderNewParamsAccessGoogleAppsTypeLinkedin   IdentityProviderNewParamsAccessGoogleAppsType = "linkedin"
	IdentityProviderNewParamsAccessGoogleAppsTypeOidc       IdentityProviderNewParamsAccessGoogleAppsType = "oidc"
	IdentityProviderNewParamsAccessGoogleAppsTypeOkta       IdentityProviderNewParamsAccessGoogleAppsType = "okta"
	IdentityProviderNewParamsAccessGoogleAppsTypeOnelogin   IdentityProviderNewParamsAccessGoogleAppsType = "onelogin"
	IdentityProviderNewParamsAccessGoogleAppsTypePingone    IdentityProviderNewParamsAccessGoogleAppsType = "pingone"
	IdentityProviderNewParamsAccessGoogleAppsTypeYandex     IdentityProviderNewParamsAccessGoogleAppsType = "yandex"
)

func (r IdentityProviderNewParamsAccessGoogleAppsType) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessGoogleAppsTypeOnetimepin, IdentityProviderNewParamsAccessGoogleAppsTypeAzureAd, IdentityProviderNewParamsAccessGoogleAppsTypeSaml, IdentityProviderNewParamsAccessGoogleAppsTypeCentrify, IdentityProviderNewParamsAccessGoogleAppsTypeFacebook, IdentityProviderNewParamsAccessGoogleAppsTypeGitHub, IdentityProviderNewParamsAccessGoogleAppsTypeGoogleApps, IdentityProviderNewParamsAccessGoogleAppsTypeGoogle, IdentityProviderNewParamsAccessGoogleAppsTypeLinkedin, IdentityProviderNewParamsAccessGoogleAppsTypeOidc, IdentityProviderNewParamsAccessGoogleAppsTypeOkta, IdentityProviderNewParamsAccessGoogleAppsTypeOnelogin, IdentityProviderNewParamsAccessGoogleAppsTypePingone, IdentityProviderNewParamsAccessGoogleAppsTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewParamsAccessGoogleAppsScimConfig struct {
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

func (r IdentityProviderNewParamsAccessGoogleAppsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsAccessLinkedinConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsAccessLinkedinType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderNewParamsAccessLinkedinScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAccessLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAccessLinkedin) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAccessLinkedin) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAccessLinkedin) ImplementsIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r IdentityProviderNewParamsAccessLinkedinConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessLinkedinType string

const (
	IdentityProviderNewParamsAccessLinkedinTypeOnetimepin IdentityProviderNewParamsAccessLinkedinType = "onetimepin"
	IdentityProviderNewParamsAccessLinkedinTypeAzureAd    IdentityProviderNewParamsAccessLinkedinType = "azureAD"
	IdentityProviderNewParamsAccessLinkedinTypeSaml       IdentityProviderNewParamsAccessLinkedinType = "saml"
	IdentityProviderNewParamsAccessLinkedinTypeCentrify   IdentityProviderNewParamsAccessLinkedinType = "centrify"
	IdentityProviderNewParamsAccessLinkedinTypeFacebook   IdentityProviderNewParamsAccessLinkedinType = "facebook"
	IdentityProviderNewParamsAccessLinkedinTypeGitHub     IdentityProviderNewParamsAccessLinkedinType = "github"
	IdentityProviderNewParamsAccessLinkedinTypeGoogleApps IdentityProviderNewParamsAccessLinkedinType = "google-apps"
	IdentityProviderNewParamsAccessLinkedinTypeGoogle     IdentityProviderNewParamsAccessLinkedinType = "google"
	IdentityProviderNewParamsAccessLinkedinTypeLinkedin   IdentityProviderNewParamsAccessLinkedinType = "linkedin"
	IdentityProviderNewParamsAccessLinkedinTypeOidc       IdentityProviderNewParamsAccessLinkedinType = "oidc"
	IdentityProviderNewParamsAccessLinkedinTypeOkta       IdentityProviderNewParamsAccessLinkedinType = "okta"
	IdentityProviderNewParamsAccessLinkedinTypeOnelogin   IdentityProviderNewParamsAccessLinkedinType = "onelogin"
	IdentityProviderNewParamsAccessLinkedinTypePingone    IdentityProviderNewParamsAccessLinkedinType = "pingone"
	IdentityProviderNewParamsAccessLinkedinTypeYandex     IdentityProviderNewParamsAccessLinkedinType = "yandex"
)

func (r IdentityProviderNewParamsAccessLinkedinType) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessLinkedinTypeOnetimepin, IdentityProviderNewParamsAccessLinkedinTypeAzureAd, IdentityProviderNewParamsAccessLinkedinTypeSaml, IdentityProviderNewParamsAccessLinkedinTypeCentrify, IdentityProviderNewParamsAccessLinkedinTypeFacebook, IdentityProviderNewParamsAccessLinkedinTypeGitHub, IdentityProviderNewParamsAccessLinkedinTypeGoogleApps, IdentityProviderNewParamsAccessLinkedinTypeGoogle, IdentityProviderNewParamsAccessLinkedinTypeLinkedin, IdentityProviderNewParamsAccessLinkedinTypeOidc, IdentityProviderNewParamsAccessLinkedinTypeOkta, IdentityProviderNewParamsAccessLinkedinTypeOnelogin, IdentityProviderNewParamsAccessLinkedinTypePingone, IdentityProviderNewParamsAccessLinkedinTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewParamsAccessLinkedinScimConfig struct {
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

func (r IdentityProviderNewParamsAccessLinkedinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsAccessOidcConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsAccessOidcType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderNewParamsAccessOidcScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAccessOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAccessOidc) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAccessOidc) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAccessOidc) ImplementsIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL param.Field[string] `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CERTsURL param.Field[string] `json:"certs_url"`
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

func (r IdentityProviderNewParamsAccessOidcConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessOidcType string

const (
	IdentityProviderNewParamsAccessOidcTypeOnetimepin IdentityProviderNewParamsAccessOidcType = "onetimepin"
	IdentityProviderNewParamsAccessOidcTypeAzureAd    IdentityProviderNewParamsAccessOidcType = "azureAD"
	IdentityProviderNewParamsAccessOidcTypeSaml       IdentityProviderNewParamsAccessOidcType = "saml"
	IdentityProviderNewParamsAccessOidcTypeCentrify   IdentityProviderNewParamsAccessOidcType = "centrify"
	IdentityProviderNewParamsAccessOidcTypeFacebook   IdentityProviderNewParamsAccessOidcType = "facebook"
	IdentityProviderNewParamsAccessOidcTypeGitHub     IdentityProviderNewParamsAccessOidcType = "github"
	IdentityProviderNewParamsAccessOidcTypeGoogleApps IdentityProviderNewParamsAccessOidcType = "google-apps"
	IdentityProviderNewParamsAccessOidcTypeGoogle     IdentityProviderNewParamsAccessOidcType = "google"
	IdentityProviderNewParamsAccessOidcTypeLinkedin   IdentityProviderNewParamsAccessOidcType = "linkedin"
	IdentityProviderNewParamsAccessOidcTypeOidc       IdentityProviderNewParamsAccessOidcType = "oidc"
	IdentityProviderNewParamsAccessOidcTypeOkta       IdentityProviderNewParamsAccessOidcType = "okta"
	IdentityProviderNewParamsAccessOidcTypeOnelogin   IdentityProviderNewParamsAccessOidcType = "onelogin"
	IdentityProviderNewParamsAccessOidcTypePingone    IdentityProviderNewParamsAccessOidcType = "pingone"
	IdentityProviderNewParamsAccessOidcTypeYandex     IdentityProviderNewParamsAccessOidcType = "yandex"
)

func (r IdentityProviderNewParamsAccessOidcType) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessOidcTypeOnetimepin, IdentityProviderNewParamsAccessOidcTypeAzureAd, IdentityProviderNewParamsAccessOidcTypeSaml, IdentityProviderNewParamsAccessOidcTypeCentrify, IdentityProviderNewParamsAccessOidcTypeFacebook, IdentityProviderNewParamsAccessOidcTypeGitHub, IdentityProviderNewParamsAccessOidcTypeGoogleApps, IdentityProviderNewParamsAccessOidcTypeGoogle, IdentityProviderNewParamsAccessOidcTypeLinkedin, IdentityProviderNewParamsAccessOidcTypeOidc, IdentityProviderNewParamsAccessOidcTypeOkta, IdentityProviderNewParamsAccessOidcTypeOnelogin, IdentityProviderNewParamsAccessOidcTypePingone, IdentityProviderNewParamsAccessOidcTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewParamsAccessOidcScimConfig struct {
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

func (r IdentityProviderNewParamsAccessOidcScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsAccessOktaConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsAccessOktaType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderNewParamsAccessOktaScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAccessOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAccessOkta) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAccessOkta) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAccessOkta) ImplementsIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessOktaConfig struct {
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

func (r IdentityProviderNewParamsAccessOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessOktaType string

const (
	IdentityProviderNewParamsAccessOktaTypeOnetimepin IdentityProviderNewParamsAccessOktaType = "onetimepin"
	IdentityProviderNewParamsAccessOktaTypeAzureAd    IdentityProviderNewParamsAccessOktaType = "azureAD"
	IdentityProviderNewParamsAccessOktaTypeSaml       IdentityProviderNewParamsAccessOktaType = "saml"
	IdentityProviderNewParamsAccessOktaTypeCentrify   IdentityProviderNewParamsAccessOktaType = "centrify"
	IdentityProviderNewParamsAccessOktaTypeFacebook   IdentityProviderNewParamsAccessOktaType = "facebook"
	IdentityProviderNewParamsAccessOktaTypeGitHub     IdentityProviderNewParamsAccessOktaType = "github"
	IdentityProviderNewParamsAccessOktaTypeGoogleApps IdentityProviderNewParamsAccessOktaType = "google-apps"
	IdentityProviderNewParamsAccessOktaTypeGoogle     IdentityProviderNewParamsAccessOktaType = "google"
	IdentityProviderNewParamsAccessOktaTypeLinkedin   IdentityProviderNewParamsAccessOktaType = "linkedin"
	IdentityProviderNewParamsAccessOktaTypeOidc       IdentityProviderNewParamsAccessOktaType = "oidc"
	IdentityProviderNewParamsAccessOktaTypeOkta       IdentityProviderNewParamsAccessOktaType = "okta"
	IdentityProviderNewParamsAccessOktaTypeOnelogin   IdentityProviderNewParamsAccessOktaType = "onelogin"
	IdentityProviderNewParamsAccessOktaTypePingone    IdentityProviderNewParamsAccessOktaType = "pingone"
	IdentityProviderNewParamsAccessOktaTypeYandex     IdentityProviderNewParamsAccessOktaType = "yandex"
)

func (r IdentityProviderNewParamsAccessOktaType) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessOktaTypeOnetimepin, IdentityProviderNewParamsAccessOktaTypeAzureAd, IdentityProviderNewParamsAccessOktaTypeSaml, IdentityProviderNewParamsAccessOktaTypeCentrify, IdentityProviderNewParamsAccessOktaTypeFacebook, IdentityProviderNewParamsAccessOktaTypeGitHub, IdentityProviderNewParamsAccessOktaTypeGoogleApps, IdentityProviderNewParamsAccessOktaTypeGoogle, IdentityProviderNewParamsAccessOktaTypeLinkedin, IdentityProviderNewParamsAccessOktaTypeOidc, IdentityProviderNewParamsAccessOktaTypeOkta, IdentityProviderNewParamsAccessOktaTypeOnelogin, IdentityProviderNewParamsAccessOktaTypePingone, IdentityProviderNewParamsAccessOktaTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewParamsAccessOktaScimConfig struct {
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

func (r IdentityProviderNewParamsAccessOktaScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsAccessOneloginConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsAccessOneloginType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderNewParamsAccessOneloginScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAccessOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAccessOnelogin) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAccessOnelogin) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAccessOnelogin) ImplementsIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessOneloginConfig struct {
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

func (r IdentityProviderNewParamsAccessOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessOneloginType string

const (
	IdentityProviderNewParamsAccessOneloginTypeOnetimepin IdentityProviderNewParamsAccessOneloginType = "onetimepin"
	IdentityProviderNewParamsAccessOneloginTypeAzureAd    IdentityProviderNewParamsAccessOneloginType = "azureAD"
	IdentityProviderNewParamsAccessOneloginTypeSaml       IdentityProviderNewParamsAccessOneloginType = "saml"
	IdentityProviderNewParamsAccessOneloginTypeCentrify   IdentityProviderNewParamsAccessOneloginType = "centrify"
	IdentityProviderNewParamsAccessOneloginTypeFacebook   IdentityProviderNewParamsAccessOneloginType = "facebook"
	IdentityProviderNewParamsAccessOneloginTypeGitHub     IdentityProviderNewParamsAccessOneloginType = "github"
	IdentityProviderNewParamsAccessOneloginTypeGoogleApps IdentityProviderNewParamsAccessOneloginType = "google-apps"
	IdentityProviderNewParamsAccessOneloginTypeGoogle     IdentityProviderNewParamsAccessOneloginType = "google"
	IdentityProviderNewParamsAccessOneloginTypeLinkedin   IdentityProviderNewParamsAccessOneloginType = "linkedin"
	IdentityProviderNewParamsAccessOneloginTypeOidc       IdentityProviderNewParamsAccessOneloginType = "oidc"
	IdentityProviderNewParamsAccessOneloginTypeOkta       IdentityProviderNewParamsAccessOneloginType = "okta"
	IdentityProviderNewParamsAccessOneloginTypeOnelogin   IdentityProviderNewParamsAccessOneloginType = "onelogin"
	IdentityProviderNewParamsAccessOneloginTypePingone    IdentityProviderNewParamsAccessOneloginType = "pingone"
	IdentityProviderNewParamsAccessOneloginTypeYandex     IdentityProviderNewParamsAccessOneloginType = "yandex"
)

func (r IdentityProviderNewParamsAccessOneloginType) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessOneloginTypeOnetimepin, IdentityProviderNewParamsAccessOneloginTypeAzureAd, IdentityProviderNewParamsAccessOneloginTypeSaml, IdentityProviderNewParamsAccessOneloginTypeCentrify, IdentityProviderNewParamsAccessOneloginTypeFacebook, IdentityProviderNewParamsAccessOneloginTypeGitHub, IdentityProviderNewParamsAccessOneloginTypeGoogleApps, IdentityProviderNewParamsAccessOneloginTypeGoogle, IdentityProviderNewParamsAccessOneloginTypeLinkedin, IdentityProviderNewParamsAccessOneloginTypeOidc, IdentityProviderNewParamsAccessOneloginTypeOkta, IdentityProviderNewParamsAccessOneloginTypeOnelogin, IdentityProviderNewParamsAccessOneloginTypePingone, IdentityProviderNewParamsAccessOneloginTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewParamsAccessOneloginScimConfig struct {
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

func (r IdentityProviderNewParamsAccessOneloginScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsAccessPingoneConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsAccessPingoneType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderNewParamsAccessPingoneScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAccessPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAccessPingone) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAccessPingone) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAccessPingone) ImplementsIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessPingoneConfig struct {
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

func (r IdentityProviderNewParamsAccessPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessPingoneType string

const (
	IdentityProviderNewParamsAccessPingoneTypeOnetimepin IdentityProviderNewParamsAccessPingoneType = "onetimepin"
	IdentityProviderNewParamsAccessPingoneTypeAzureAd    IdentityProviderNewParamsAccessPingoneType = "azureAD"
	IdentityProviderNewParamsAccessPingoneTypeSaml       IdentityProviderNewParamsAccessPingoneType = "saml"
	IdentityProviderNewParamsAccessPingoneTypeCentrify   IdentityProviderNewParamsAccessPingoneType = "centrify"
	IdentityProviderNewParamsAccessPingoneTypeFacebook   IdentityProviderNewParamsAccessPingoneType = "facebook"
	IdentityProviderNewParamsAccessPingoneTypeGitHub     IdentityProviderNewParamsAccessPingoneType = "github"
	IdentityProviderNewParamsAccessPingoneTypeGoogleApps IdentityProviderNewParamsAccessPingoneType = "google-apps"
	IdentityProviderNewParamsAccessPingoneTypeGoogle     IdentityProviderNewParamsAccessPingoneType = "google"
	IdentityProviderNewParamsAccessPingoneTypeLinkedin   IdentityProviderNewParamsAccessPingoneType = "linkedin"
	IdentityProviderNewParamsAccessPingoneTypeOidc       IdentityProviderNewParamsAccessPingoneType = "oidc"
	IdentityProviderNewParamsAccessPingoneTypeOkta       IdentityProviderNewParamsAccessPingoneType = "okta"
	IdentityProviderNewParamsAccessPingoneTypeOnelogin   IdentityProviderNewParamsAccessPingoneType = "onelogin"
	IdentityProviderNewParamsAccessPingoneTypePingone    IdentityProviderNewParamsAccessPingoneType = "pingone"
	IdentityProviderNewParamsAccessPingoneTypeYandex     IdentityProviderNewParamsAccessPingoneType = "yandex"
)

func (r IdentityProviderNewParamsAccessPingoneType) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessPingoneTypeOnetimepin, IdentityProviderNewParamsAccessPingoneTypeAzureAd, IdentityProviderNewParamsAccessPingoneTypeSaml, IdentityProviderNewParamsAccessPingoneTypeCentrify, IdentityProviderNewParamsAccessPingoneTypeFacebook, IdentityProviderNewParamsAccessPingoneTypeGitHub, IdentityProviderNewParamsAccessPingoneTypeGoogleApps, IdentityProviderNewParamsAccessPingoneTypeGoogle, IdentityProviderNewParamsAccessPingoneTypeLinkedin, IdentityProviderNewParamsAccessPingoneTypeOidc, IdentityProviderNewParamsAccessPingoneTypeOkta, IdentityProviderNewParamsAccessPingoneTypeOnelogin, IdentityProviderNewParamsAccessPingoneTypePingone, IdentityProviderNewParamsAccessPingoneTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewParamsAccessPingoneScimConfig struct {
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

func (r IdentityProviderNewParamsAccessPingoneScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsAccessSamlConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsAccessSamlType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderNewParamsAccessSamlScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAccessSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAccessSaml) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAccessSaml) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAccessSaml) ImplementsIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]IdentityProviderNewParamsAccessSamlConfigHeaderAttribute] `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IDPPublicCERTs param.Field[[]string] `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL param.Field[string] `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest param.Field[bool] `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL param.Field[string] `json:"sso_target_url"`
}

func (r IdentityProviderNewParamsAccessSamlConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r IdentityProviderNewParamsAccessSamlConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessSamlType string

const (
	IdentityProviderNewParamsAccessSamlTypeOnetimepin IdentityProviderNewParamsAccessSamlType = "onetimepin"
	IdentityProviderNewParamsAccessSamlTypeAzureAd    IdentityProviderNewParamsAccessSamlType = "azureAD"
	IdentityProviderNewParamsAccessSamlTypeSaml       IdentityProviderNewParamsAccessSamlType = "saml"
	IdentityProviderNewParamsAccessSamlTypeCentrify   IdentityProviderNewParamsAccessSamlType = "centrify"
	IdentityProviderNewParamsAccessSamlTypeFacebook   IdentityProviderNewParamsAccessSamlType = "facebook"
	IdentityProviderNewParamsAccessSamlTypeGitHub     IdentityProviderNewParamsAccessSamlType = "github"
	IdentityProviderNewParamsAccessSamlTypeGoogleApps IdentityProviderNewParamsAccessSamlType = "google-apps"
	IdentityProviderNewParamsAccessSamlTypeGoogle     IdentityProviderNewParamsAccessSamlType = "google"
	IdentityProviderNewParamsAccessSamlTypeLinkedin   IdentityProviderNewParamsAccessSamlType = "linkedin"
	IdentityProviderNewParamsAccessSamlTypeOidc       IdentityProviderNewParamsAccessSamlType = "oidc"
	IdentityProviderNewParamsAccessSamlTypeOkta       IdentityProviderNewParamsAccessSamlType = "okta"
	IdentityProviderNewParamsAccessSamlTypeOnelogin   IdentityProviderNewParamsAccessSamlType = "onelogin"
	IdentityProviderNewParamsAccessSamlTypePingone    IdentityProviderNewParamsAccessSamlType = "pingone"
	IdentityProviderNewParamsAccessSamlTypeYandex     IdentityProviderNewParamsAccessSamlType = "yandex"
)

func (r IdentityProviderNewParamsAccessSamlType) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessSamlTypeOnetimepin, IdentityProviderNewParamsAccessSamlTypeAzureAd, IdentityProviderNewParamsAccessSamlTypeSaml, IdentityProviderNewParamsAccessSamlTypeCentrify, IdentityProviderNewParamsAccessSamlTypeFacebook, IdentityProviderNewParamsAccessSamlTypeGitHub, IdentityProviderNewParamsAccessSamlTypeGoogleApps, IdentityProviderNewParamsAccessSamlTypeGoogle, IdentityProviderNewParamsAccessSamlTypeLinkedin, IdentityProviderNewParamsAccessSamlTypeOidc, IdentityProviderNewParamsAccessSamlTypeOkta, IdentityProviderNewParamsAccessSamlTypeOnelogin, IdentityProviderNewParamsAccessSamlTypePingone, IdentityProviderNewParamsAccessSamlTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewParamsAccessSamlScimConfig struct {
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

func (r IdentityProviderNewParamsAccessSamlScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsAccessYandexConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsAccessYandexType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderNewParamsAccessYandexScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAccessYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAccessYandex) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAccessYandex) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAccessYandex) ImplementsIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r IdentityProviderNewParamsAccessYandexConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessYandexType string

const (
	IdentityProviderNewParamsAccessYandexTypeOnetimepin IdentityProviderNewParamsAccessYandexType = "onetimepin"
	IdentityProviderNewParamsAccessYandexTypeAzureAd    IdentityProviderNewParamsAccessYandexType = "azureAD"
	IdentityProviderNewParamsAccessYandexTypeSaml       IdentityProviderNewParamsAccessYandexType = "saml"
	IdentityProviderNewParamsAccessYandexTypeCentrify   IdentityProviderNewParamsAccessYandexType = "centrify"
	IdentityProviderNewParamsAccessYandexTypeFacebook   IdentityProviderNewParamsAccessYandexType = "facebook"
	IdentityProviderNewParamsAccessYandexTypeGitHub     IdentityProviderNewParamsAccessYandexType = "github"
	IdentityProviderNewParamsAccessYandexTypeGoogleApps IdentityProviderNewParamsAccessYandexType = "google-apps"
	IdentityProviderNewParamsAccessYandexTypeGoogle     IdentityProviderNewParamsAccessYandexType = "google"
	IdentityProviderNewParamsAccessYandexTypeLinkedin   IdentityProviderNewParamsAccessYandexType = "linkedin"
	IdentityProviderNewParamsAccessYandexTypeOidc       IdentityProviderNewParamsAccessYandexType = "oidc"
	IdentityProviderNewParamsAccessYandexTypeOkta       IdentityProviderNewParamsAccessYandexType = "okta"
	IdentityProviderNewParamsAccessYandexTypeOnelogin   IdentityProviderNewParamsAccessYandexType = "onelogin"
	IdentityProviderNewParamsAccessYandexTypePingone    IdentityProviderNewParamsAccessYandexType = "pingone"
	IdentityProviderNewParamsAccessYandexTypeYandex     IdentityProviderNewParamsAccessYandexType = "yandex"
)

func (r IdentityProviderNewParamsAccessYandexType) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessYandexTypeOnetimepin, IdentityProviderNewParamsAccessYandexTypeAzureAd, IdentityProviderNewParamsAccessYandexTypeSaml, IdentityProviderNewParamsAccessYandexTypeCentrify, IdentityProviderNewParamsAccessYandexTypeFacebook, IdentityProviderNewParamsAccessYandexTypeGitHub, IdentityProviderNewParamsAccessYandexTypeGoogleApps, IdentityProviderNewParamsAccessYandexTypeGoogle, IdentityProviderNewParamsAccessYandexTypeLinkedin, IdentityProviderNewParamsAccessYandexTypeOidc, IdentityProviderNewParamsAccessYandexTypeOkta, IdentityProviderNewParamsAccessYandexTypeOnelogin, IdentityProviderNewParamsAccessYandexTypePingone, IdentityProviderNewParamsAccessYandexTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewParamsAccessYandexScimConfig struct {
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

func (r IdentityProviderNewParamsAccessYandexScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[interface{}] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderNewParamsAccessOnetimepinType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderNewParamsAccessOnetimepinScimConfig] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAccessOnetimepin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAccessOnetimepin) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAccessOnetimepin) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAccessOnetimepin) ImplementsIdentityProviderNewParams() {

}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAccessOnetimepinType string

const (
	IdentityProviderNewParamsAccessOnetimepinTypeOnetimepin IdentityProviderNewParamsAccessOnetimepinType = "onetimepin"
	IdentityProviderNewParamsAccessOnetimepinTypeAzureAd    IdentityProviderNewParamsAccessOnetimepinType = "azureAD"
	IdentityProviderNewParamsAccessOnetimepinTypeSaml       IdentityProviderNewParamsAccessOnetimepinType = "saml"
	IdentityProviderNewParamsAccessOnetimepinTypeCentrify   IdentityProviderNewParamsAccessOnetimepinType = "centrify"
	IdentityProviderNewParamsAccessOnetimepinTypeFacebook   IdentityProviderNewParamsAccessOnetimepinType = "facebook"
	IdentityProviderNewParamsAccessOnetimepinTypeGitHub     IdentityProviderNewParamsAccessOnetimepinType = "github"
	IdentityProviderNewParamsAccessOnetimepinTypeGoogleApps IdentityProviderNewParamsAccessOnetimepinType = "google-apps"
	IdentityProviderNewParamsAccessOnetimepinTypeGoogle     IdentityProviderNewParamsAccessOnetimepinType = "google"
	IdentityProviderNewParamsAccessOnetimepinTypeLinkedin   IdentityProviderNewParamsAccessOnetimepinType = "linkedin"
	IdentityProviderNewParamsAccessOnetimepinTypeOidc       IdentityProviderNewParamsAccessOnetimepinType = "oidc"
	IdentityProviderNewParamsAccessOnetimepinTypeOkta       IdentityProviderNewParamsAccessOnetimepinType = "okta"
	IdentityProviderNewParamsAccessOnetimepinTypeOnelogin   IdentityProviderNewParamsAccessOnetimepinType = "onelogin"
	IdentityProviderNewParamsAccessOnetimepinTypePingone    IdentityProviderNewParamsAccessOnetimepinType = "pingone"
	IdentityProviderNewParamsAccessOnetimepinTypeYandex     IdentityProviderNewParamsAccessOnetimepinType = "yandex"
)

func (r IdentityProviderNewParamsAccessOnetimepinType) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAccessOnetimepinTypeOnetimepin, IdentityProviderNewParamsAccessOnetimepinTypeAzureAd, IdentityProviderNewParamsAccessOnetimepinTypeSaml, IdentityProviderNewParamsAccessOnetimepinTypeCentrify, IdentityProviderNewParamsAccessOnetimepinTypeFacebook, IdentityProviderNewParamsAccessOnetimepinTypeGitHub, IdentityProviderNewParamsAccessOnetimepinTypeGoogleApps, IdentityProviderNewParamsAccessOnetimepinTypeGoogle, IdentityProviderNewParamsAccessOnetimepinTypeLinkedin, IdentityProviderNewParamsAccessOnetimepinTypeOidc, IdentityProviderNewParamsAccessOnetimepinTypeOkta, IdentityProviderNewParamsAccessOnetimepinTypeOnelogin, IdentityProviderNewParamsAccessOnetimepinTypePingone, IdentityProviderNewParamsAccessOnetimepinTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderNewParamsAccessOnetimepinScimConfig struct {
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

func (r IdentityProviderNewParamsAccessOnetimepinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewResponseEnvelope struct {
	Errors   []IdentityProviderNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IdentityProviderNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustIdentityProviders                    `json:"result,required"`
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

func (r IdentityProviderNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IdentityProviderNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

// This interface is a union satisfied by one of the following:
// [IdentityProviderUpdateParamsAccessAzureAd],
// [IdentityProviderUpdateParamsAccessCentrify],
// [IdentityProviderUpdateParamsAccessFacebook],
// [IdentityProviderUpdateParamsAccessGitHub],
// [IdentityProviderUpdateParamsAccessGoogle],
// [IdentityProviderUpdateParamsAccessGoogleApps],
// [IdentityProviderUpdateParamsAccessLinkedin],
// [IdentityProviderUpdateParamsAccessOidc],
// [IdentityProviderUpdateParamsAccessOkta],
// [IdentityProviderUpdateParamsAccessOnelogin],
// [IdentityProviderUpdateParamsAccessPingone],
// [IdentityProviderUpdateParamsAccessSaml],
// [IdentityProviderUpdateParamsAccessYandex],
// [IdentityProviderUpdateParamsAccessOnetimepin].
type IdentityProviderUpdateParams interface {
	ImplementsIdentityProviderUpdateParams()

	getAccountID() param.Field[string]

	getZoneID() param.Field[string]
}

type IdentityProviderUpdateParamsAccessAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsAccessAzureAdConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsAccessAzureAdType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderUpdateParamsAccessAzureAdScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAccessAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAccessAzureAd) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAccessAzureAd) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAccessAzureAd) ImplementsIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessAzureAdConfig struct {
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
	// Indicates the type of user interaction that is required. prompt=login forces the
	// user to enter their credentials on that request, negating single-sign on.
	// prompt=none is the opposite. It ensures that the user isn't presented with any
	// interactive prompt. If the request can't be completed silently by using
	// single-sign on, the Microsoft identity platform returns an interaction_required
	// error. prompt=select_account interrupts single sign-on providing account
	// selection experience listing all the accounts either in session or any
	// remembered account or an option to choose to use a different account altogether.
	Prompt param.Field[IdentityProviderUpdateParamsAccessAzureAdConfigPrompt] `json:"prompt"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
}

func (r IdentityProviderUpdateParamsAccessAzureAdConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Indicates the type of user interaction that is required. prompt=login forces the
// user to enter their credentials on that request, negating single-sign on.
// prompt=none is the opposite. It ensures that the user isn't presented with any
// interactive prompt. If the request can't be completed silently by using
// single-sign on, the Microsoft identity platform returns an interaction_required
// error. prompt=select_account interrupts single sign-on providing account
// selection experience listing all the accounts either in session or any
// remembered account or an option to choose to use a different account altogether.
type IdentityProviderUpdateParamsAccessAzureAdConfigPrompt string

const (
	IdentityProviderUpdateParamsAccessAzureAdConfigPromptLogin         IdentityProviderUpdateParamsAccessAzureAdConfigPrompt = "login"
	IdentityProviderUpdateParamsAccessAzureAdConfigPromptSelectAccount IdentityProviderUpdateParamsAccessAzureAdConfigPrompt = "select_account"
	IdentityProviderUpdateParamsAccessAzureAdConfigPromptNone          IdentityProviderUpdateParamsAccessAzureAdConfigPrompt = "none"
)

func (r IdentityProviderUpdateParamsAccessAzureAdConfigPrompt) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessAzureAdConfigPromptLogin, IdentityProviderUpdateParamsAccessAzureAdConfigPromptSelectAccount, IdentityProviderUpdateParamsAccessAzureAdConfigPromptNone:
		return true
	}
	return false
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessAzureAdType string

const (
	IdentityProviderUpdateParamsAccessAzureAdTypeOnetimepin IdentityProviderUpdateParamsAccessAzureAdType = "onetimepin"
	IdentityProviderUpdateParamsAccessAzureAdTypeAzureAd    IdentityProviderUpdateParamsAccessAzureAdType = "azureAD"
	IdentityProviderUpdateParamsAccessAzureAdTypeSaml       IdentityProviderUpdateParamsAccessAzureAdType = "saml"
	IdentityProviderUpdateParamsAccessAzureAdTypeCentrify   IdentityProviderUpdateParamsAccessAzureAdType = "centrify"
	IdentityProviderUpdateParamsAccessAzureAdTypeFacebook   IdentityProviderUpdateParamsAccessAzureAdType = "facebook"
	IdentityProviderUpdateParamsAccessAzureAdTypeGitHub     IdentityProviderUpdateParamsAccessAzureAdType = "github"
	IdentityProviderUpdateParamsAccessAzureAdTypeGoogleApps IdentityProviderUpdateParamsAccessAzureAdType = "google-apps"
	IdentityProviderUpdateParamsAccessAzureAdTypeGoogle     IdentityProviderUpdateParamsAccessAzureAdType = "google"
	IdentityProviderUpdateParamsAccessAzureAdTypeLinkedin   IdentityProviderUpdateParamsAccessAzureAdType = "linkedin"
	IdentityProviderUpdateParamsAccessAzureAdTypeOidc       IdentityProviderUpdateParamsAccessAzureAdType = "oidc"
	IdentityProviderUpdateParamsAccessAzureAdTypeOkta       IdentityProviderUpdateParamsAccessAzureAdType = "okta"
	IdentityProviderUpdateParamsAccessAzureAdTypeOnelogin   IdentityProviderUpdateParamsAccessAzureAdType = "onelogin"
	IdentityProviderUpdateParamsAccessAzureAdTypePingone    IdentityProviderUpdateParamsAccessAzureAdType = "pingone"
	IdentityProviderUpdateParamsAccessAzureAdTypeYandex     IdentityProviderUpdateParamsAccessAzureAdType = "yandex"
)

func (r IdentityProviderUpdateParamsAccessAzureAdType) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessAzureAdTypeOnetimepin, IdentityProviderUpdateParamsAccessAzureAdTypeAzureAd, IdentityProviderUpdateParamsAccessAzureAdTypeSaml, IdentityProviderUpdateParamsAccessAzureAdTypeCentrify, IdentityProviderUpdateParamsAccessAzureAdTypeFacebook, IdentityProviderUpdateParamsAccessAzureAdTypeGitHub, IdentityProviderUpdateParamsAccessAzureAdTypeGoogleApps, IdentityProviderUpdateParamsAccessAzureAdTypeGoogle, IdentityProviderUpdateParamsAccessAzureAdTypeLinkedin, IdentityProviderUpdateParamsAccessAzureAdTypeOidc, IdentityProviderUpdateParamsAccessAzureAdTypeOkta, IdentityProviderUpdateParamsAccessAzureAdTypeOnelogin, IdentityProviderUpdateParamsAccessAzureAdTypePingone, IdentityProviderUpdateParamsAccessAzureAdTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateParamsAccessAzureAdScimConfig struct {
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

func (r IdentityProviderUpdateParamsAccessAzureAdScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsAccessCentrifyConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsAccessCentrifyType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderUpdateParamsAccessCentrifyScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAccessCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAccessCentrify) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAccessCentrify) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAccessCentrify) ImplementsIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessCentrifyConfig struct {
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

func (r IdentityProviderUpdateParamsAccessCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessCentrifyType string

const (
	IdentityProviderUpdateParamsAccessCentrifyTypeOnetimepin IdentityProviderUpdateParamsAccessCentrifyType = "onetimepin"
	IdentityProviderUpdateParamsAccessCentrifyTypeAzureAd    IdentityProviderUpdateParamsAccessCentrifyType = "azureAD"
	IdentityProviderUpdateParamsAccessCentrifyTypeSaml       IdentityProviderUpdateParamsAccessCentrifyType = "saml"
	IdentityProviderUpdateParamsAccessCentrifyTypeCentrify   IdentityProviderUpdateParamsAccessCentrifyType = "centrify"
	IdentityProviderUpdateParamsAccessCentrifyTypeFacebook   IdentityProviderUpdateParamsAccessCentrifyType = "facebook"
	IdentityProviderUpdateParamsAccessCentrifyTypeGitHub     IdentityProviderUpdateParamsAccessCentrifyType = "github"
	IdentityProviderUpdateParamsAccessCentrifyTypeGoogleApps IdentityProviderUpdateParamsAccessCentrifyType = "google-apps"
	IdentityProviderUpdateParamsAccessCentrifyTypeGoogle     IdentityProviderUpdateParamsAccessCentrifyType = "google"
	IdentityProviderUpdateParamsAccessCentrifyTypeLinkedin   IdentityProviderUpdateParamsAccessCentrifyType = "linkedin"
	IdentityProviderUpdateParamsAccessCentrifyTypeOidc       IdentityProviderUpdateParamsAccessCentrifyType = "oidc"
	IdentityProviderUpdateParamsAccessCentrifyTypeOkta       IdentityProviderUpdateParamsAccessCentrifyType = "okta"
	IdentityProviderUpdateParamsAccessCentrifyTypeOnelogin   IdentityProviderUpdateParamsAccessCentrifyType = "onelogin"
	IdentityProviderUpdateParamsAccessCentrifyTypePingone    IdentityProviderUpdateParamsAccessCentrifyType = "pingone"
	IdentityProviderUpdateParamsAccessCentrifyTypeYandex     IdentityProviderUpdateParamsAccessCentrifyType = "yandex"
)

func (r IdentityProviderUpdateParamsAccessCentrifyType) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessCentrifyTypeOnetimepin, IdentityProviderUpdateParamsAccessCentrifyTypeAzureAd, IdentityProviderUpdateParamsAccessCentrifyTypeSaml, IdentityProviderUpdateParamsAccessCentrifyTypeCentrify, IdentityProviderUpdateParamsAccessCentrifyTypeFacebook, IdentityProviderUpdateParamsAccessCentrifyTypeGitHub, IdentityProviderUpdateParamsAccessCentrifyTypeGoogleApps, IdentityProviderUpdateParamsAccessCentrifyTypeGoogle, IdentityProviderUpdateParamsAccessCentrifyTypeLinkedin, IdentityProviderUpdateParamsAccessCentrifyTypeOidc, IdentityProviderUpdateParamsAccessCentrifyTypeOkta, IdentityProviderUpdateParamsAccessCentrifyTypeOnelogin, IdentityProviderUpdateParamsAccessCentrifyTypePingone, IdentityProviderUpdateParamsAccessCentrifyTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateParamsAccessCentrifyScimConfig struct {
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

func (r IdentityProviderUpdateParamsAccessCentrifyScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsAccessFacebookConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsAccessFacebookType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderUpdateParamsAccessFacebookScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAccessFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAccessFacebook) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAccessFacebook) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAccessFacebook) ImplementsIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r IdentityProviderUpdateParamsAccessFacebookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessFacebookType string

const (
	IdentityProviderUpdateParamsAccessFacebookTypeOnetimepin IdentityProviderUpdateParamsAccessFacebookType = "onetimepin"
	IdentityProviderUpdateParamsAccessFacebookTypeAzureAd    IdentityProviderUpdateParamsAccessFacebookType = "azureAD"
	IdentityProviderUpdateParamsAccessFacebookTypeSaml       IdentityProviderUpdateParamsAccessFacebookType = "saml"
	IdentityProviderUpdateParamsAccessFacebookTypeCentrify   IdentityProviderUpdateParamsAccessFacebookType = "centrify"
	IdentityProviderUpdateParamsAccessFacebookTypeFacebook   IdentityProviderUpdateParamsAccessFacebookType = "facebook"
	IdentityProviderUpdateParamsAccessFacebookTypeGitHub     IdentityProviderUpdateParamsAccessFacebookType = "github"
	IdentityProviderUpdateParamsAccessFacebookTypeGoogleApps IdentityProviderUpdateParamsAccessFacebookType = "google-apps"
	IdentityProviderUpdateParamsAccessFacebookTypeGoogle     IdentityProviderUpdateParamsAccessFacebookType = "google"
	IdentityProviderUpdateParamsAccessFacebookTypeLinkedin   IdentityProviderUpdateParamsAccessFacebookType = "linkedin"
	IdentityProviderUpdateParamsAccessFacebookTypeOidc       IdentityProviderUpdateParamsAccessFacebookType = "oidc"
	IdentityProviderUpdateParamsAccessFacebookTypeOkta       IdentityProviderUpdateParamsAccessFacebookType = "okta"
	IdentityProviderUpdateParamsAccessFacebookTypeOnelogin   IdentityProviderUpdateParamsAccessFacebookType = "onelogin"
	IdentityProviderUpdateParamsAccessFacebookTypePingone    IdentityProviderUpdateParamsAccessFacebookType = "pingone"
	IdentityProviderUpdateParamsAccessFacebookTypeYandex     IdentityProviderUpdateParamsAccessFacebookType = "yandex"
)

func (r IdentityProviderUpdateParamsAccessFacebookType) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessFacebookTypeOnetimepin, IdentityProviderUpdateParamsAccessFacebookTypeAzureAd, IdentityProviderUpdateParamsAccessFacebookTypeSaml, IdentityProviderUpdateParamsAccessFacebookTypeCentrify, IdentityProviderUpdateParamsAccessFacebookTypeFacebook, IdentityProviderUpdateParamsAccessFacebookTypeGitHub, IdentityProviderUpdateParamsAccessFacebookTypeGoogleApps, IdentityProviderUpdateParamsAccessFacebookTypeGoogle, IdentityProviderUpdateParamsAccessFacebookTypeLinkedin, IdentityProviderUpdateParamsAccessFacebookTypeOidc, IdentityProviderUpdateParamsAccessFacebookTypeOkta, IdentityProviderUpdateParamsAccessFacebookTypeOnelogin, IdentityProviderUpdateParamsAccessFacebookTypePingone, IdentityProviderUpdateParamsAccessFacebookTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateParamsAccessFacebookScimConfig struct {
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

func (r IdentityProviderUpdateParamsAccessFacebookScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsAccessGitHubConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsAccessGitHubType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderUpdateParamsAccessGitHubScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAccessGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAccessGitHub) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAccessGitHub) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAccessGitHub) ImplementsIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r IdentityProviderUpdateParamsAccessGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessGitHubType string

const (
	IdentityProviderUpdateParamsAccessGitHubTypeOnetimepin IdentityProviderUpdateParamsAccessGitHubType = "onetimepin"
	IdentityProviderUpdateParamsAccessGitHubTypeAzureAd    IdentityProviderUpdateParamsAccessGitHubType = "azureAD"
	IdentityProviderUpdateParamsAccessGitHubTypeSaml       IdentityProviderUpdateParamsAccessGitHubType = "saml"
	IdentityProviderUpdateParamsAccessGitHubTypeCentrify   IdentityProviderUpdateParamsAccessGitHubType = "centrify"
	IdentityProviderUpdateParamsAccessGitHubTypeFacebook   IdentityProviderUpdateParamsAccessGitHubType = "facebook"
	IdentityProviderUpdateParamsAccessGitHubTypeGitHub     IdentityProviderUpdateParamsAccessGitHubType = "github"
	IdentityProviderUpdateParamsAccessGitHubTypeGoogleApps IdentityProviderUpdateParamsAccessGitHubType = "google-apps"
	IdentityProviderUpdateParamsAccessGitHubTypeGoogle     IdentityProviderUpdateParamsAccessGitHubType = "google"
	IdentityProviderUpdateParamsAccessGitHubTypeLinkedin   IdentityProviderUpdateParamsAccessGitHubType = "linkedin"
	IdentityProviderUpdateParamsAccessGitHubTypeOidc       IdentityProviderUpdateParamsAccessGitHubType = "oidc"
	IdentityProviderUpdateParamsAccessGitHubTypeOkta       IdentityProviderUpdateParamsAccessGitHubType = "okta"
	IdentityProviderUpdateParamsAccessGitHubTypeOnelogin   IdentityProviderUpdateParamsAccessGitHubType = "onelogin"
	IdentityProviderUpdateParamsAccessGitHubTypePingone    IdentityProviderUpdateParamsAccessGitHubType = "pingone"
	IdentityProviderUpdateParamsAccessGitHubTypeYandex     IdentityProviderUpdateParamsAccessGitHubType = "yandex"
)

func (r IdentityProviderUpdateParamsAccessGitHubType) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessGitHubTypeOnetimepin, IdentityProviderUpdateParamsAccessGitHubTypeAzureAd, IdentityProviderUpdateParamsAccessGitHubTypeSaml, IdentityProviderUpdateParamsAccessGitHubTypeCentrify, IdentityProviderUpdateParamsAccessGitHubTypeFacebook, IdentityProviderUpdateParamsAccessGitHubTypeGitHub, IdentityProviderUpdateParamsAccessGitHubTypeGoogleApps, IdentityProviderUpdateParamsAccessGitHubTypeGoogle, IdentityProviderUpdateParamsAccessGitHubTypeLinkedin, IdentityProviderUpdateParamsAccessGitHubTypeOidc, IdentityProviderUpdateParamsAccessGitHubTypeOkta, IdentityProviderUpdateParamsAccessGitHubTypeOnelogin, IdentityProviderUpdateParamsAccessGitHubTypePingone, IdentityProviderUpdateParamsAccessGitHubTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateParamsAccessGitHubScimConfig struct {
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

func (r IdentityProviderUpdateParamsAccessGitHubScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsAccessGoogleConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsAccessGoogleType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderUpdateParamsAccessGoogleScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAccessGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAccessGoogle) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAccessGoogle) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAccessGoogle) ImplementsIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessGoogleConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r IdentityProviderUpdateParamsAccessGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessGoogleType string

const (
	IdentityProviderUpdateParamsAccessGoogleTypeOnetimepin IdentityProviderUpdateParamsAccessGoogleType = "onetimepin"
	IdentityProviderUpdateParamsAccessGoogleTypeAzureAd    IdentityProviderUpdateParamsAccessGoogleType = "azureAD"
	IdentityProviderUpdateParamsAccessGoogleTypeSaml       IdentityProviderUpdateParamsAccessGoogleType = "saml"
	IdentityProviderUpdateParamsAccessGoogleTypeCentrify   IdentityProviderUpdateParamsAccessGoogleType = "centrify"
	IdentityProviderUpdateParamsAccessGoogleTypeFacebook   IdentityProviderUpdateParamsAccessGoogleType = "facebook"
	IdentityProviderUpdateParamsAccessGoogleTypeGitHub     IdentityProviderUpdateParamsAccessGoogleType = "github"
	IdentityProviderUpdateParamsAccessGoogleTypeGoogleApps IdentityProviderUpdateParamsAccessGoogleType = "google-apps"
	IdentityProviderUpdateParamsAccessGoogleTypeGoogle     IdentityProviderUpdateParamsAccessGoogleType = "google"
	IdentityProviderUpdateParamsAccessGoogleTypeLinkedin   IdentityProviderUpdateParamsAccessGoogleType = "linkedin"
	IdentityProviderUpdateParamsAccessGoogleTypeOidc       IdentityProviderUpdateParamsAccessGoogleType = "oidc"
	IdentityProviderUpdateParamsAccessGoogleTypeOkta       IdentityProviderUpdateParamsAccessGoogleType = "okta"
	IdentityProviderUpdateParamsAccessGoogleTypeOnelogin   IdentityProviderUpdateParamsAccessGoogleType = "onelogin"
	IdentityProviderUpdateParamsAccessGoogleTypePingone    IdentityProviderUpdateParamsAccessGoogleType = "pingone"
	IdentityProviderUpdateParamsAccessGoogleTypeYandex     IdentityProviderUpdateParamsAccessGoogleType = "yandex"
)

func (r IdentityProviderUpdateParamsAccessGoogleType) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessGoogleTypeOnetimepin, IdentityProviderUpdateParamsAccessGoogleTypeAzureAd, IdentityProviderUpdateParamsAccessGoogleTypeSaml, IdentityProviderUpdateParamsAccessGoogleTypeCentrify, IdentityProviderUpdateParamsAccessGoogleTypeFacebook, IdentityProviderUpdateParamsAccessGoogleTypeGitHub, IdentityProviderUpdateParamsAccessGoogleTypeGoogleApps, IdentityProviderUpdateParamsAccessGoogleTypeGoogle, IdentityProviderUpdateParamsAccessGoogleTypeLinkedin, IdentityProviderUpdateParamsAccessGoogleTypeOidc, IdentityProviderUpdateParamsAccessGoogleTypeOkta, IdentityProviderUpdateParamsAccessGoogleTypeOnelogin, IdentityProviderUpdateParamsAccessGoogleTypePingone, IdentityProviderUpdateParamsAccessGoogleTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateParamsAccessGoogleScimConfig struct {
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

func (r IdentityProviderUpdateParamsAccessGoogleScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsAccessGoogleAppsConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsAccessGoogleAppsType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderUpdateParamsAccessGoogleAppsScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAccessGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAccessGoogleApps) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAccessGoogleApps) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAccessGoogleApps) ImplementsIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessGoogleAppsConfig struct {
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

func (r IdentityProviderUpdateParamsAccessGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessGoogleAppsType string

const (
	IdentityProviderUpdateParamsAccessGoogleAppsTypeOnetimepin IdentityProviderUpdateParamsAccessGoogleAppsType = "onetimepin"
	IdentityProviderUpdateParamsAccessGoogleAppsTypeAzureAd    IdentityProviderUpdateParamsAccessGoogleAppsType = "azureAD"
	IdentityProviderUpdateParamsAccessGoogleAppsTypeSaml       IdentityProviderUpdateParamsAccessGoogleAppsType = "saml"
	IdentityProviderUpdateParamsAccessGoogleAppsTypeCentrify   IdentityProviderUpdateParamsAccessGoogleAppsType = "centrify"
	IdentityProviderUpdateParamsAccessGoogleAppsTypeFacebook   IdentityProviderUpdateParamsAccessGoogleAppsType = "facebook"
	IdentityProviderUpdateParamsAccessGoogleAppsTypeGitHub     IdentityProviderUpdateParamsAccessGoogleAppsType = "github"
	IdentityProviderUpdateParamsAccessGoogleAppsTypeGoogleApps IdentityProviderUpdateParamsAccessGoogleAppsType = "google-apps"
	IdentityProviderUpdateParamsAccessGoogleAppsTypeGoogle     IdentityProviderUpdateParamsAccessGoogleAppsType = "google"
	IdentityProviderUpdateParamsAccessGoogleAppsTypeLinkedin   IdentityProviderUpdateParamsAccessGoogleAppsType = "linkedin"
	IdentityProviderUpdateParamsAccessGoogleAppsTypeOidc       IdentityProviderUpdateParamsAccessGoogleAppsType = "oidc"
	IdentityProviderUpdateParamsAccessGoogleAppsTypeOkta       IdentityProviderUpdateParamsAccessGoogleAppsType = "okta"
	IdentityProviderUpdateParamsAccessGoogleAppsTypeOnelogin   IdentityProviderUpdateParamsAccessGoogleAppsType = "onelogin"
	IdentityProviderUpdateParamsAccessGoogleAppsTypePingone    IdentityProviderUpdateParamsAccessGoogleAppsType = "pingone"
	IdentityProviderUpdateParamsAccessGoogleAppsTypeYandex     IdentityProviderUpdateParamsAccessGoogleAppsType = "yandex"
)

func (r IdentityProviderUpdateParamsAccessGoogleAppsType) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessGoogleAppsTypeOnetimepin, IdentityProviderUpdateParamsAccessGoogleAppsTypeAzureAd, IdentityProviderUpdateParamsAccessGoogleAppsTypeSaml, IdentityProviderUpdateParamsAccessGoogleAppsTypeCentrify, IdentityProviderUpdateParamsAccessGoogleAppsTypeFacebook, IdentityProviderUpdateParamsAccessGoogleAppsTypeGitHub, IdentityProviderUpdateParamsAccessGoogleAppsTypeGoogleApps, IdentityProviderUpdateParamsAccessGoogleAppsTypeGoogle, IdentityProviderUpdateParamsAccessGoogleAppsTypeLinkedin, IdentityProviderUpdateParamsAccessGoogleAppsTypeOidc, IdentityProviderUpdateParamsAccessGoogleAppsTypeOkta, IdentityProviderUpdateParamsAccessGoogleAppsTypeOnelogin, IdentityProviderUpdateParamsAccessGoogleAppsTypePingone, IdentityProviderUpdateParamsAccessGoogleAppsTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateParamsAccessGoogleAppsScimConfig struct {
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

func (r IdentityProviderUpdateParamsAccessGoogleAppsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsAccessLinkedinConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsAccessLinkedinType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderUpdateParamsAccessLinkedinScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAccessLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAccessLinkedin) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAccessLinkedin) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAccessLinkedin) ImplementsIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r IdentityProviderUpdateParamsAccessLinkedinConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessLinkedinType string

const (
	IdentityProviderUpdateParamsAccessLinkedinTypeOnetimepin IdentityProviderUpdateParamsAccessLinkedinType = "onetimepin"
	IdentityProviderUpdateParamsAccessLinkedinTypeAzureAd    IdentityProviderUpdateParamsAccessLinkedinType = "azureAD"
	IdentityProviderUpdateParamsAccessLinkedinTypeSaml       IdentityProviderUpdateParamsAccessLinkedinType = "saml"
	IdentityProviderUpdateParamsAccessLinkedinTypeCentrify   IdentityProviderUpdateParamsAccessLinkedinType = "centrify"
	IdentityProviderUpdateParamsAccessLinkedinTypeFacebook   IdentityProviderUpdateParamsAccessLinkedinType = "facebook"
	IdentityProviderUpdateParamsAccessLinkedinTypeGitHub     IdentityProviderUpdateParamsAccessLinkedinType = "github"
	IdentityProviderUpdateParamsAccessLinkedinTypeGoogleApps IdentityProviderUpdateParamsAccessLinkedinType = "google-apps"
	IdentityProviderUpdateParamsAccessLinkedinTypeGoogle     IdentityProviderUpdateParamsAccessLinkedinType = "google"
	IdentityProviderUpdateParamsAccessLinkedinTypeLinkedin   IdentityProviderUpdateParamsAccessLinkedinType = "linkedin"
	IdentityProviderUpdateParamsAccessLinkedinTypeOidc       IdentityProviderUpdateParamsAccessLinkedinType = "oidc"
	IdentityProviderUpdateParamsAccessLinkedinTypeOkta       IdentityProviderUpdateParamsAccessLinkedinType = "okta"
	IdentityProviderUpdateParamsAccessLinkedinTypeOnelogin   IdentityProviderUpdateParamsAccessLinkedinType = "onelogin"
	IdentityProviderUpdateParamsAccessLinkedinTypePingone    IdentityProviderUpdateParamsAccessLinkedinType = "pingone"
	IdentityProviderUpdateParamsAccessLinkedinTypeYandex     IdentityProviderUpdateParamsAccessLinkedinType = "yandex"
)

func (r IdentityProviderUpdateParamsAccessLinkedinType) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessLinkedinTypeOnetimepin, IdentityProviderUpdateParamsAccessLinkedinTypeAzureAd, IdentityProviderUpdateParamsAccessLinkedinTypeSaml, IdentityProviderUpdateParamsAccessLinkedinTypeCentrify, IdentityProviderUpdateParamsAccessLinkedinTypeFacebook, IdentityProviderUpdateParamsAccessLinkedinTypeGitHub, IdentityProviderUpdateParamsAccessLinkedinTypeGoogleApps, IdentityProviderUpdateParamsAccessLinkedinTypeGoogle, IdentityProviderUpdateParamsAccessLinkedinTypeLinkedin, IdentityProviderUpdateParamsAccessLinkedinTypeOidc, IdentityProviderUpdateParamsAccessLinkedinTypeOkta, IdentityProviderUpdateParamsAccessLinkedinTypeOnelogin, IdentityProviderUpdateParamsAccessLinkedinTypePingone, IdentityProviderUpdateParamsAccessLinkedinTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateParamsAccessLinkedinScimConfig struct {
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

func (r IdentityProviderUpdateParamsAccessLinkedinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsAccessOidcConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsAccessOidcType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderUpdateParamsAccessOidcScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAccessOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAccessOidc) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAccessOidc) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAccessOidc) ImplementsIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL param.Field[string] `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CERTsURL param.Field[string] `json:"certs_url"`
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

func (r IdentityProviderUpdateParamsAccessOidcConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessOidcType string

const (
	IdentityProviderUpdateParamsAccessOidcTypeOnetimepin IdentityProviderUpdateParamsAccessOidcType = "onetimepin"
	IdentityProviderUpdateParamsAccessOidcTypeAzureAd    IdentityProviderUpdateParamsAccessOidcType = "azureAD"
	IdentityProviderUpdateParamsAccessOidcTypeSaml       IdentityProviderUpdateParamsAccessOidcType = "saml"
	IdentityProviderUpdateParamsAccessOidcTypeCentrify   IdentityProviderUpdateParamsAccessOidcType = "centrify"
	IdentityProviderUpdateParamsAccessOidcTypeFacebook   IdentityProviderUpdateParamsAccessOidcType = "facebook"
	IdentityProviderUpdateParamsAccessOidcTypeGitHub     IdentityProviderUpdateParamsAccessOidcType = "github"
	IdentityProviderUpdateParamsAccessOidcTypeGoogleApps IdentityProviderUpdateParamsAccessOidcType = "google-apps"
	IdentityProviderUpdateParamsAccessOidcTypeGoogle     IdentityProviderUpdateParamsAccessOidcType = "google"
	IdentityProviderUpdateParamsAccessOidcTypeLinkedin   IdentityProviderUpdateParamsAccessOidcType = "linkedin"
	IdentityProviderUpdateParamsAccessOidcTypeOidc       IdentityProviderUpdateParamsAccessOidcType = "oidc"
	IdentityProviderUpdateParamsAccessOidcTypeOkta       IdentityProviderUpdateParamsAccessOidcType = "okta"
	IdentityProviderUpdateParamsAccessOidcTypeOnelogin   IdentityProviderUpdateParamsAccessOidcType = "onelogin"
	IdentityProviderUpdateParamsAccessOidcTypePingone    IdentityProviderUpdateParamsAccessOidcType = "pingone"
	IdentityProviderUpdateParamsAccessOidcTypeYandex     IdentityProviderUpdateParamsAccessOidcType = "yandex"
)

func (r IdentityProviderUpdateParamsAccessOidcType) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessOidcTypeOnetimepin, IdentityProviderUpdateParamsAccessOidcTypeAzureAd, IdentityProviderUpdateParamsAccessOidcTypeSaml, IdentityProviderUpdateParamsAccessOidcTypeCentrify, IdentityProviderUpdateParamsAccessOidcTypeFacebook, IdentityProviderUpdateParamsAccessOidcTypeGitHub, IdentityProviderUpdateParamsAccessOidcTypeGoogleApps, IdentityProviderUpdateParamsAccessOidcTypeGoogle, IdentityProviderUpdateParamsAccessOidcTypeLinkedin, IdentityProviderUpdateParamsAccessOidcTypeOidc, IdentityProviderUpdateParamsAccessOidcTypeOkta, IdentityProviderUpdateParamsAccessOidcTypeOnelogin, IdentityProviderUpdateParamsAccessOidcTypePingone, IdentityProviderUpdateParamsAccessOidcTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateParamsAccessOidcScimConfig struct {
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

func (r IdentityProviderUpdateParamsAccessOidcScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsAccessOktaConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsAccessOktaType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderUpdateParamsAccessOktaScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAccessOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAccessOkta) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAccessOkta) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAccessOkta) ImplementsIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessOktaConfig struct {
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

func (r IdentityProviderUpdateParamsAccessOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessOktaType string

const (
	IdentityProviderUpdateParamsAccessOktaTypeOnetimepin IdentityProviderUpdateParamsAccessOktaType = "onetimepin"
	IdentityProviderUpdateParamsAccessOktaTypeAzureAd    IdentityProviderUpdateParamsAccessOktaType = "azureAD"
	IdentityProviderUpdateParamsAccessOktaTypeSaml       IdentityProviderUpdateParamsAccessOktaType = "saml"
	IdentityProviderUpdateParamsAccessOktaTypeCentrify   IdentityProviderUpdateParamsAccessOktaType = "centrify"
	IdentityProviderUpdateParamsAccessOktaTypeFacebook   IdentityProviderUpdateParamsAccessOktaType = "facebook"
	IdentityProviderUpdateParamsAccessOktaTypeGitHub     IdentityProviderUpdateParamsAccessOktaType = "github"
	IdentityProviderUpdateParamsAccessOktaTypeGoogleApps IdentityProviderUpdateParamsAccessOktaType = "google-apps"
	IdentityProviderUpdateParamsAccessOktaTypeGoogle     IdentityProviderUpdateParamsAccessOktaType = "google"
	IdentityProviderUpdateParamsAccessOktaTypeLinkedin   IdentityProviderUpdateParamsAccessOktaType = "linkedin"
	IdentityProviderUpdateParamsAccessOktaTypeOidc       IdentityProviderUpdateParamsAccessOktaType = "oidc"
	IdentityProviderUpdateParamsAccessOktaTypeOkta       IdentityProviderUpdateParamsAccessOktaType = "okta"
	IdentityProviderUpdateParamsAccessOktaTypeOnelogin   IdentityProviderUpdateParamsAccessOktaType = "onelogin"
	IdentityProviderUpdateParamsAccessOktaTypePingone    IdentityProviderUpdateParamsAccessOktaType = "pingone"
	IdentityProviderUpdateParamsAccessOktaTypeYandex     IdentityProviderUpdateParamsAccessOktaType = "yandex"
)

func (r IdentityProviderUpdateParamsAccessOktaType) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessOktaTypeOnetimepin, IdentityProviderUpdateParamsAccessOktaTypeAzureAd, IdentityProviderUpdateParamsAccessOktaTypeSaml, IdentityProviderUpdateParamsAccessOktaTypeCentrify, IdentityProviderUpdateParamsAccessOktaTypeFacebook, IdentityProviderUpdateParamsAccessOktaTypeGitHub, IdentityProviderUpdateParamsAccessOktaTypeGoogleApps, IdentityProviderUpdateParamsAccessOktaTypeGoogle, IdentityProviderUpdateParamsAccessOktaTypeLinkedin, IdentityProviderUpdateParamsAccessOktaTypeOidc, IdentityProviderUpdateParamsAccessOktaTypeOkta, IdentityProviderUpdateParamsAccessOktaTypeOnelogin, IdentityProviderUpdateParamsAccessOktaTypePingone, IdentityProviderUpdateParamsAccessOktaTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateParamsAccessOktaScimConfig struct {
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

func (r IdentityProviderUpdateParamsAccessOktaScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsAccessOneloginConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsAccessOneloginType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderUpdateParamsAccessOneloginScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAccessOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAccessOnelogin) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAccessOnelogin) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAccessOnelogin) ImplementsIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessOneloginConfig struct {
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

func (r IdentityProviderUpdateParamsAccessOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessOneloginType string

const (
	IdentityProviderUpdateParamsAccessOneloginTypeOnetimepin IdentityProviderUpdateParamsAccessOneloginType = "onetimepin"
	IdentityProviderUpdateParamsAccessOneloginTypeAzureAd    IdentityProviderUpdateParamsAccessOneloginType = "azureAD"
	IdentityProviderUpdateParamsAccessOneloginTypeSaml       IdentityProviderUpdateParamsAccessOneloginType = "saml"
	IdentityProviderUpdateParamsAccessOneloginTypeCentrify   IdentityProviderUpdateParamsAccessOneloginType = "centrify"
	IdentityProviderUpdateParamsAccessOneloginTypeFacebook   IdentityProviderUpdateParamsAccessOneloginType = "facebook"
	IdentityProviderUpdateParamsAccessOneloginTypeGitHub     IdentityProviderUpdateParamsAccessOneloginType = "github"
	IdentityProviderUpdateParamsAccessOneloginTypeGoogleApps IdentityProviderUpdateParamsAccessOneloginType = "google-apps"
	IdentityProviderUpdateParamsAccessOneloginTypeGoogle     IdentityProviderUpdateParamsAccessOneloginType = "google"
	IdentityProviderUpdateParamsAccessOneloginTypeLinkedin   IdentityProviderUpdateParamsAccessOneloginType = "linkedin"
	IdentityProviderUpdateParamsAccessOneloginTypeOidc       IdentityProviderUpdateParamsAccessOneloginType = "oidc"
	IdentityProviderUpdateParamsAccessOneloginTypeOkta       IdentityProviderUpdateParamsAccessOneloginType = "okta"
	IdentityProviderUpdateParamsAccessOneloginTypeOnelogin   IdentityProviderUpdateParamsAccessOneloginType = "onelogin"
	IdentityProviderUpdateParamsAccessOneloginTypePingone    IdentityProviderUpdateParamsAccessOneloginType = "pingone"
	IdentityProviderUpdateParamsAccessOneloginTypeYandex     IdentityProviderUpdateParamsAccessOneloginType = "yandex"
)

func (r IdentityProviderUpdateParamsAccessOneloginType) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessOneloginTypeOnetimepin, IdentityProviderUpdateParamsAccessOneloginTypeAzureAd, IdentityProviderUpdateParamsAccessOneloginTypeSaml, IdentityProviderUpdateParamsAccessOneloginTypeCentrify, IdentityProviderUpdateParamsAccessOneloginTypeFacebook, IdentityProviderUpdateParamsAccessOneloginTypeGitHub, IdentityProviderUpdateParamsAccessOneloginTypeGoogleApps, IdentityProviderUpdateParamsAccessOneloginTypeGoogle, IdentityProviderUpdateParamsAccessOneloginTypeLinkedin, IdentityProviderUpdateParamsAccessOneloginTypeOidc, IdentityProviderUpdateParamsAccessOneloginTypeOkta, IdentityProviderUpdateParamsAccessOneloginTypeOnelogin, IdentityProviderUpdateParamsAccessOneloginTypePingone, IdentityProviderUpdateParamsAccessOneloginTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateParamsAccessOneloginScimConfig struct {
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

func (r IdentityProviderUpdateParamsAccessOneloginScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsAccessPingoneConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsAccessPingoneType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderUpdateParamsAccessPingoneScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAccessPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAccessPingone) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAccessPingone) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAccessPingone) ImplementsIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessPingoneConfig struct {
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

func (r IdentityProviderUpdateParamsAccessPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessPingoneType string

const (
	IdentityProviderUpdateParamsAccessPingoneTypeOnetimepin IdentityProviderUpdateParamsAccessPingoneType = "onetimepin"
	IdentityProviderUpdateParamsAccessPingoneTypeAzureAd    IdentityProviderUpdateParamsAccessPingoneType = "azureAD"
	IdentityProviderUpdateParamsAccessPingoneTypeSaml       IdentityProviderUpdateParamsAccessPingoneType = "saml"
	IdentityProviderUpdateParamsAccessPingoneTypeCentrify   IdentityProviderUpdateParamsAccessPingoneType = "centrify"
	IdentityProviderUpdateParamsAccessPingoneTypeFacebook   IdentityProviderUpdateParamsAccessPingoneType = "facebook"
	IdentityProviderUpdateParamsAccessPingoneTypeGitHub     IdentityProviderUpdateParamsAccessPingoneType = "github"
	IdentityProviderUpdateParamsAccessPingoneTypeGoogleApps IdentityProviderUpdateParamsAccessPingoneType = "google-apps"
	IdentityProviderUpdateParamsAccessPingoneTypeGoogle     IdentityProviderUpdateParamsAccessPingoneType = "google"
	IdentityProviderUpdateParamsAccessPingoneTypeLinkedin   IdentityProviderUpdateParamsAccessPingoneType = "linkedin"
	IdentityProviderUpdateParamsAccessPingoneTypeOidc       IdentityProviderUpdateParamsAccessPingoneType = "oidc"
	IdentityProviderUpdateParamsAccessPingoneTypeOkta       IdentityProviderUpdateParamsAccessPingoneType = "okta"
	IdentityProviderUpdateParamsAccessPingoneTypeOnelogin   IdentityProviderUpdateParamsAccessPingoneType = "onelogin"
	IdentityProviderUpdateParamsAccessPingoneTypePingone    IdentityProviderUpdateParamsAccessPingoneType = "pingone"
	IdentityProviderUpdateParamsAccessPingoneTypeYandex     IdentityProviderUpdateParamsAccessPingoneType = "yandex"
)

func (r IdentityProviderUpdateParamsAccessPingoneType) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessPingoneTypeOnetimepin, IdentityProviderUpdateParamsAccessPingoneTypeAzureAd, IdentityProviderUpdateParamsAccessPingoneTypeSaml, IdentityProviderUpdateParamsAccessPingoneTypeCentrify, IdentityProviderUpdateParamsAccessPingoneTypeFacebook, IdentityProviderUpdateParamsAccessPingoneTypeGitHub, IdentityProviderUpdateParamsAccessPingoneTypeGoogleApps, IdentityProviderUpdateParamsAccessPingoneTypeGoogle, IdentityProviderUpdateParamsAccessPingoneTypeLinkedin, IdentityProviderUpdateParamsAccessPingoneTypeOidc, IdentityProviderUpdateParamsAccessPingoneTypeOkta, IdentityProviderUpdateParamsAccessPingoneTypeOnelogin, IdentityProviderUpdateParamsAccessPingoneTypePingone, IdentityProviderUpdateParamsAccessPingoneTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateParamsAccessPingoneScimConfig struct {
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

func (r IdentityProviderUpdateParamsAccessPingoneScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsAccessSamlConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsAccessSamlType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderUpdateParamsAccessSamlScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAccessSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAccessSaml) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAccessSaml) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAccessSaml) ImplementsIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]IdentityProviderUpdateParamsAccessSamlConfigHeaderAttribute] `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IDPPublicCERTs param.Field[[]string] `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL param.Field[string] `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest param.Field[bool] `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL param.Field[string] `json:"sso_target_url"`
}

func (r IdentityProviderUpdateParamsAccessSamlConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r IdentityProviderUpdateParamsAccessSamlConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessSamlType string

const (
	IdentityProviderUpdateParamsAccessSamlTypeOnetimepin IdentityProviderUpdateParamsAccessSamlType = "onetimepin"
	IdentityProviderUpdateParamsAccessSamlTypeAzureAd    IdentityProviderUpdateParamsAccessSamlType = "azureAD"
	IdentityProviderUpdateParamsAccessSamlTypeSaml       IdentityProviderUpdateParamsAccessSamlType = "saml"
	IdentityProviderUpdateParamsAccessSamlTypeCentrify   IdentityProviderUpdateParamsAccessSamlType = "centrify"
	IdentityProviderUpdateParamsAccessSamlTypeFacebook   IdentityProviderUpdateParamsAccessSamlType = "facebook"
	IdentityProviderUpdateParamsAccessSamlTypeGitHub     IdentityProviderUpdateParamsAccessSamlType = "github"
	IdentityProviderUpdateParamsAccessSamlTypeGoogleApps IdentityProviderUpdateParamsAccessSamlType = "google-apps"
	IdentityProviderUpdateParamsAccessSamlTypeGoogle     IdentityProviderUpdateParamsAccessSamlType = "google"
	IdentityProviderUpdateParamsAccessSamlTypeLinkedin   IdentityProviderUpdateParamsAccessSamlType = "linkedin"
	IdentityProviderUpdateParamsAccessSamlTypeOidc       IdentityProviderUpdateParamsAccessSamlType = "oidc"
	IdentityProviderUpdateParamsAccessSamlTypeOkta       IdentityProviderUpdateParamsAccessSamlType = "okta"
	IdentityProviderUpdateParamsAccessSamlTypeOnelogin   IdentityProviderUpdateParamsAccessSamlType = "onelogin"
	IdentityProviderUpdateParamsAccessSamlTypePingone    IdentityProviderUpdateParamsAccessSamlType = "pingone"
	IdentityProviderUpdateParamsAccessSamlTypeYandex     IdentityProviderUpdateParamsAccessSamlType = "yandex"
)

func (r IdentityProviderUpdateParamsAccessSamlType) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessSamlTypeOnetimepin, IdentityProviderUpdateParamsAccessSamlTypeAzureAd, IdentityProviderUpdateParamsAccessSamlTypeSaml, IdentityProviderUpdateParamsAccessSamlTypeCentrify, IdentityProviderUpdateParamsAccessSamlTypeFacebook, IdentityProviderUpdateParamsAccessSamlTypeGitHub, IdentityProviderUpdateParamsAccessSamlTypeGoogleApps, IdentityProviderUpdateParamsAccessSamlTypeGoogle, IdentityProviderUpdateParamsAccessSamlTypeLinkedin, IdentityProviderUpdateParamsAccessSamlTypeOidc, IdentityProviderUpdateParamsAccessSamlTypeOkta, IdentityProviderUpdateParamsAccessSamlTypeOnelogin, IdentityProviderUpdateParamsAccessSamlTypePingone, IdentityProviderUpdateParamsAccessSamlTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateParamsAccessSamlScimConfig struct {
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

func (r IdentityProviderUpdateParamsAccessSamlScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsAccessYandexConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsAccessYandexType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderUpdateParamsAccessYandexScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAccessYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAccessYandex) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAccessYandex) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAccessYandex) ImplementsIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r IdentityProviderUpdateParamsAccessYandexConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessYandexType string

const (
	IdentityProviderUpdateParamsAccessYandexTypeOnetimepin IdentityProviderUpdateParamsAccessYandexType = "onetimepin"
	IdentityProviderUpdateParamsAccessYandexTypeAzureAd    IdentityProviderUpdateParamsAccessYandexType = "azureAD"
	IdentityProviderUpdateParamsAccessYandexTypeSaml       IdentityProviderUpdateParamsAccessYandexType = "saml"
	IdentityProviderUpdateParamsAccessYandexTypeCentrify   IdentityProviderUpdateParamsAccessYandexType = "centrify"
	IdentityProviderUpdateParamsAccessYandexTypeFacebook   IdentityProviderUpdateParamsAccessYandexType = "facebook"
	IdentityProviderUpdateParamsAccessYandexTypeGitHub     IdentityProviderUpdateParamsAccessYandexType = "github"
	IdentityProviderUpdateParamsAccessYandexTypeGoogleApps IdentityProviderUpdateParamsAccessYandexType = "google-apps"
	IdentityProviderUpdateParamsAccessYandexTypeGoogle     IdentityProviderUpdateParamsAccessYandexType = "google"
	IdentityProviderUpdateParamsAccessYandexTypeLinkedin   IdentityProviderUpdateParamsAccessYandexType = "linkedin"
	IdentityProviderUpdateParamsAccessYandexTypeOidc       IdentityProviderUpdateParamsAccessYandexType = "oidc"
	IdentityProviderUpdateParamsAccessYandexTypeOkta       IdentityProviderUpdateParamsAccessYandexType = "okta"
	IdentityProviderUpdateParamsAccessYandexTypeOnelogin   IdentityProviderUpdateParamsAccessYandexType = "onelogin"
	IdentityProviderUpdateParamsAccessYandexTypePingone    IdentityProviderUpdateParamsAccessYandexType = "pingone"
	IdentityProviderUpdateParamsAccessYandexTypeYandex     IdentityProviderUpdateParamsAccessYandexType = "yandex"
)

func (r IdentityProviderUpdateParamsAccessYandexType) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessYandexTypeOnetimepin, IdentityProviderUpdateParamsAccessYandexTypeAzureAd, IdentityProviderUpdateParamsAccessYandexTypeSaml, IdentityProviderUpdateParamsAccessYandexTypeCentrify, IdentityProviderUpdateParamsAccessYandexTypeFacebook, IdentityProviderUpdateParamsAccessYandexTypeGitHub, IdentityProviderUpdateParamsAccessYandexTypeGoogleApps, IdentityProviderUpdateParamsAccessYandexTypeGoogle, IdentityProviderUpdateParamsAccessYandexTypeLinkedin, IdentityProviderUpdateParamsAccessYandexTypeOidc, IdentityProviderUpdateParamsAccessYandexTypeOkta, IdentityProviderUpdateParamsAccessYandexTypeOnelogin, IdentityProviderUpdateParamsAccessYandexTypePingone, IdentityProviderUpdateParamsAccessYandexTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateParamsAccessYandexScimConfig struct {
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

func (r IdentityProviderUpdateParamsAccessYandexScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[interface{}] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderUpdateParamsAccessOnetimepinType] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[IdentityProviderUpdateParamsAccessOnetimepinScimConfig] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAccessOnetimepin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAccessOnetimepin) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAccessOnetimepin) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAccessOnetimepin) ImplementsIdentityProviderUpdateParams() {

}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAccessOnetimepinType string

const (
	IdentityProviderUpdateParamsAccessOnetimepinTypeOnetimepin IdentityProviderUpdateParamsAccessOnetimepinType = "onetimepin"
	IdentityProviderUpdateParamsAccessOnetimepinTypeAzureAd    IdentityProviderUpdateParamsAccessOnetimepinType = "azureAD"
	IdentityProviderUpdateParamsAccessOnetimepinTypeSaml       IdentityProviderUpdateParamsAccessOnetimepinType = "saml"
	IdentityProviderUpdateParamsAccessOnetimepinTypeCentrify   IdentityProviderUpdateParamsAccessOnetimepinType = "centrify"
	IdentityProviderUpdateParamsAccessOnetimepinTypeFacebook   IdentityProviderUpdateParamsAccessOnetimepinType = "facebook"
	IdentityProviderUpdateParamsAccessOnetimepinTypeGitHub     IdentityProviderUpdateParamsAccessOnetimepinType = "github"
	IdentityProviderUpdateParamsAccessOnetimepinTypeGoogleApps IdentityProviderUpdateParamsAccessOnetimepinType = "google-apps"
	IdentityProviderUpdateParamsAccessOnetimepinTypeGoogle     IdentityProviderUpdateParamsAccessOnetimepinType = "google"
	IdentityProviderUpdateParamsAccessOnetimepinTypeLinkedin   IdentityProviderUpdateParamsAccessOnetimepinType = "linkedin"
	IdentityProviderUpdateParamsAccessOnetimepinTypeOidc       IdentityProviderUpdateParamsAccessOnetimepinType = "oidc"
	IdentityProviderUpdateParamsAccessOnetimepinTypeOkta       IdentityProviderUpdateParamsAccessOnetimepinType = "okta"
	IdentityProviderUpdateParamsAccessOnetimepinTypeOnelogin   IdentityProviderUpdateParamsAccessOnetimepinType = "onelogin"
	IdentityProviderUpdateParamsAccessOnetimepinTypePingone    IdentityProviderUpdateParamsAccessOnetimepinType = "pingone"
	IdentityProviderUpdateParamsAccessOnetimepinTypeYandex     IdentityProviderUpdateParamsAccessOnetimepinType = "yandex"
)

func (r IdentityProviderUpdateParamsAccessOnetimepinType) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAccessOnetimepinTypeOnetimepin, IdentityProviderUpdateParamsAccessOnetimepinTypeAzureAd, IdentityProviderUpdateParamsAccessOnetimepinTypeSaml, IdentityProviderUpdateParamsAccessOnetimepinTypeCentrify, IdentityProviderUpdateParamsAccessOnetimepinTypeFacebook, IdentityProviderUpdateParamsAccessOnetimepinTypeGitHub, IdentityProviderUpdateParamsAccessOnetimepinTypeGoogleApps, IdentityProviderUpdateParamsAccessOnetimepinTypeGoogle, IdentityProviderUpdateParamsAccessOnetimepinTypeLinkedin, IdentityProviderUpdateParamsAccessOnetimepinTypeOidc, IdentityProviderUpdateParamsAccessOnetimepinTypeOkta, IdentityProviderUpdateParamsAccessOnetimepinTypeOnelogin, IdentityProviderUpdateParamsAccessOnetimepinTypePingone, IdentityProviderUpdateParamsAccessOnetimepinTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProviderUpdateParamsAccessOnetimepinScimConfig struct {
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

func (r IdentityProviderUpdateParamsAccessOnetimepinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateResponseEnvelope struct {
	Errors   []IdentityProviderUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IdentityProviderUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustIdentityProviders                       `json:"result,required"`
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

func (r IdentityProviderUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r IdentityProviderListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IdentityProviderListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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

func (r IdentityProviderDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IdentityProviderDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IdentityProviderGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type IdentityProviderGetResponseEnvelope struct {
	Errors   []IdentityProviderGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IdentityProviderGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustIdentityProviders                    `json:"result,required"`
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

func (r IdentityProviderGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IdentityProviderGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

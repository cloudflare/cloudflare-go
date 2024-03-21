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
func (r *IdentityProviderService) New(ctx context.Context, params IdentityProviderNewParams, opts ...option.RequestOption) (res *AccessIdentityProviders, err error) {
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
func (r *IdentityProviderService) Update(ctx context.Context, uuid string, params IdentityProviderUpdateParams, opts ...option.RequestOption) (res *AccessIdentityProviders, err error) {
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
func (r *IdentityProviderService) Get(ctx context.Context, uuid string, query IdentityProviderGetParams, opts ...option.RequestOption) (res *AccessIdentityProviders, err error) {
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

// Union satisfied by [zero_trust.AccessIdentityProvidersAccessAzureAd],
// [zero_trust.AccessIdentityProvidersAccessCentrify],
// [zero_trust.AccessIdentityProvidersAccessFacebook],
// [zero_trust.AccessIdentityProvidersAccessGitHub],
// [zero_trust.AccessIdentityProvidersAccessGoogle],
// [zero_trust.AccessIdentityProvidersAccessGoogleApps],
// [zero_trust.AccessIdentityProvidersAccessLinkedin],
// [zero_trust.AccessIdentityProvidersAccessOidc],
// [zero_trust.AccessIdentityProvidersAccessOkta],
// [zero_trust.AccessIdentityProvidersAccessOnelogin],
// [zero_trust.AccessIdentityProvidersAccessPingone],
// [zero_trust.AccessIdentityProvidersAccessSaml],
// [zero_trust.AccessIdentityProvidersAccessYandex] or
// [zero_trust.AccessIdentityProvidersAccessOnetimepin].
type AccessIdentityProviders interface {
	implementsZeroTrustAccessIdentityProviders()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessIdentityProviders)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessIdentityProvidersAccessAzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessIdentityProvidersAccessCentrify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessIdentityProvidersAccessFacebook{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessIdentityProvidersAccessGitHub{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessIdentityProvidersAccessGoogle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessIdentityProvidersAccessGoogleApps{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessIdentityProvidersAccessLinkedin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessIdentityProvidersAccessOidc{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessIdentityProvidersAccessOkta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessIdentityProvidersAccessOnelogin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessIdentityProvidersAccessPingone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessIdentityProvidersAccessSaml{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessIdentityProvidersAccessYandex{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessIdentityProvidersAccessOnetimepin{}),
		},
	)
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

func (r accessIdentityProvidersAccessAzureAdJSON) RawJSON() string {
	return r.raw
}

func (r AccessIdentityProvidersAccessAzureAd) implementsZeroTrustAccessIdentityProviders() {}

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
	// Indicates the type of user interaction that is required. prompt=login forces the
	// user to enter their credentials on that request, negating single-sign on.
	// prompt=none is the opposite. It ensures that the user isn't presented with any
	// interactive prompt. If the request can't be completed silently by using
	// single-sign on, the Microsoft identity platform returns an interaction_required
	// error. prompt=select_account interrupts single sign-on providing account
	// selection experience listing all the accounts either in session or any
	// remembered account or an option to choose to use a different account altogether.
	Prompt AccessIdentityProvidersAccessAzureAdConfigPrompt `json:"prompt"`
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
	Prompt                   apijson.Field
	SupportGroups            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdentityProvidersAccessAzureAdConfigJSON) RawJSON() string {
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
type AccessIdentityProvidersAccessAzureAdConfigPrompt string

const (
	AccessIdentityProvidersAccessAzureAdConfigPromptLogin         AccessIdentityProvidersAccessAzureAdConfigPrompt = "login"
	AccessIdentityProvidersAccessAzureAdConfigPromptSelectAccount AccessIdentityProvidersAccessAzureAdConfigPrompt = "select_account"
	AccessIdentityProvidersAccessAzureAdConfigPromptNone          AccessIdentityProvidersAccessAzureAdConfigPrompt = "none"
)

func (r AccessIdentityProvidersAccessAzureAdConfigPrompt) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessAzureAdConfigPromptLogin, AccessIdentityProvidersAccessAzureAdConfigPromptSelectAccount, AccessIdentityProvidersAccessAzureAdConfigPromptNone:
		return true
	}
	return false
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

func (r AccessIdentityProvidersAccessAzureAdType) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessAzureAdTypeOnetimepin, AccessIdentityProvidersAccessAzureAdTypeAzureAd, AccessIdentityProvidersAccessAzureAdTypeSaml, AccessIdentityProvidersAccessAzureAdTypeCentrify, AccessIdentityProvidersAccessAzureAdTypeFacebook, AccessIdentityProvidersAccessAzureAdTypeGitHub, AccessIdentityProvidersAccessAzureAdTypeGoogleApps, AccessIdentityProvidersAccessAzureAdTypeGoogle, AccessIdentityProvidersAccessAzureAdTypeLinkedin, AccessIdentityProvidersAccessAzureAdTypeOidc, AccessIdentityProvidersAccessAzureAdTypeOkta, AccessIdentityProvidersAccessAzureAdTypeOnelogin, AccessIdentityProvidersAccessAzureAdTypePingone, AccessIdentityProvidersAccessAzureAdTypeYandex:
		return true
	}
	return false
}

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

func (r accessIdentityProvidersAccessAzureAdScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r accessIdentityProvidersAccessCentrifyJSON) RawJSON() string {
	return r.raw
}

func (r AccessIdentityProvidersAccessCentrify) implementsZeroTrustAccessIdentityProviders() {}

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

func (r accessIdentityProvidersAccessCentrifyConfigJSON) RawJSON() string {
	return r.raw
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

func (r AccessIdentityProvidersAccessCentrifyType) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessCentrifyTypeOnetimepin, AccessIdentityProvidersAccessCentrifyTypeAzureAd, AccessIdentityProvidersAccessCentrifyTypeSaml, AccessIdentityProvidersAccessCentrifyTypeCentrify, AccessIdentityProvidersAccessCentrifyTypeFacebook, AccessIdentityProvidersAccessCentrifyTypeGitHub, AccessIdentityProvidersAccessCentrifyTypeGoogleApps, AccessIdentityProvidersAccessCentrifyTypeGoogle, AccessIdentityProvidersAccessCentrifyTypeLinkedin, AccessIdentityProvidersAccessCentrifyTypeOidc, AccessIdentityProvidersAccessCentrifyTypeOkta, AccessIdentityProvidersAccessCentrifyTypeOnelogin, AccessIdentityProvidersAccessCentrifyTypePingone, AccessIdentityProvidersAccessCentrifyTypeYandex:
		return true
	}
	return false
}

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

func (r accessIdentityProvidersAccessCentrifyScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r accessIdentityProvidersAccessFacebookJSON) RawJSON() string {
	return r.raw
}

func (r AccessIdentityProvidersAccessFacebook) implementsZeroTrustAccessIdentityProviders() {}

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

func (r accessIdentityProvidersAccessFacebookConfigJSON) RawJSON() string {
	return r.raw
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

func (r AccessIdentityProvidersAccessFacebookType) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessFacebookTypeOnetimepin, AccessIdentityProvidersAccessFacebookTypeAzureAd, AccessIdentityProvidersAccessFacebookTypeSaml, AccessIdentityProvidersAccessFacebookTypeCentrify, AccessIdentityProvidersAccessFacebookTypeFacebook, AccessIdentityProvidersAccessFacebookTypeGitHub, AccessIdentityProvidersAccessFacebookTypeGoogleApps, AccessIdentityProvidersAccessFacebookTypeGoogle, AccessIdentityProvidersAccessFacebookTypeLinkedin, AccessIdentityProvidersAccessFacebookTypeOidc, AccessIdentityProvidersAccessFacebookTypeOkta, AccessIdentityProvidersAccessFacebookTypeOnelogin, AccessIdentityProvidersAccessFacebookTypePingone, AccessIdentityProvidersAccessFacebookTypeYandex:
		return true
	}
	return false
}

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

func (r accessIdentityProvidersAccessFacebookScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r accessIdentityProvidersAccessGitHubJSON) RawJSON() string {
	return r.raw
}

func (r AccessIdentityProvidersAccessGitHub) implementsZeroTrustAccessIdentityProviders() {}

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

func (r accessIdentityProvidersAccessGitHubConfigJSON) RawJSON() string {
	return r.raw
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

func (r AccessIdentityProvidersAccessGitHubType) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessGitHubTypeOnetimepin, AccessIdentityProvidersAccessGitHubTypeAzureAd, AccessIdentityProvidersAccessGitHubTypeSaml, AccessIdentityProvidersAccessGitHubTypeCentrify, AccessIdentityProvidersAccessGitHubTypeFacebook, AccessIdentityProvidersAccessGitHubTypeGitHub, AccessIdentityProvidersAccessGitHubTypeGoogleApps, AccessIdentityProvidersAccessGitHubTypeGoogle, AccessIdentityProvidersAccessGitHubTypeLinkedin, AccessIdentityProvidersAccessGitHubTypeOidc, AccessIdentityProvidersAccessGitHubTypeOkta, AccessIdentityProvidersAccessGitHubTypeOnelogin, AccessIdentityProvidersAccessGitHubTypePingone, AccessIdentityProvidersAccessGitHubTypeYandex:
		return true
	}
	return false
}

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

func (r accessIdentityProvidersAccessGitHubScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r accessIdentityProvidersAccessGoogleJSON) RawJSON() string {
	return r.raw
}

func (r AccessIdentityProvidersAccessGoogle) implementsZeroTrustAccessIdentityProviders() {}

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

func (r accessIdentityProvidersAccessGoogleConfigJSON) RawJSON() string {
	return r.raw
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

func (r AccessIdentityProvidersAccessGoogleType) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessGoogleTypeOnetimepin, AccessIdentityProvidersAccessGoogleTypeAzureAd, AccessIdentityProvidersAccessGoogleTypeSaml, AccessIdentityProvidersAccessGoogleTypeCentrify, AccessIdentityProvidersAccessGoogleTypeFacebook, AccessIdentityProvidersAccessGoogleTypeGitHub, AccessIdentityProvidersAccessGoogleTypeGoogleApps, AccessIdentityProvidersAccessGoogleTypeGoogle, AccessIdentityProvidersAccessGoogleTypeLinkedin, AccessIdentityProvidersAccessGoogleTypeOidc, AccessIdentityProvidersAccessGoogleTypeOkta, AccessIdentityProvidersAccessGoogleTypeOnelogin, AccessIdentityProvidersAccessGoogleTypePingone, AccessIdentityProvidersAccessGoogleTypeYandex:
		return true
	}
	return false
}

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

func (r accessIdentityProvidersAccessGoogleScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r accessIdentityProvidersAccessGoogleAppsJSON) RawJSON() string {
	return r.raw
}

func (r AccessIdentityProvidersAccessGoogleApps) implementsZeroTrustAccessIdentityProviders() {}

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

func (r accessIdentityProvidersAccessGoogleAppsConfigJSON) RawJSON() string {
	return r.raw
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

func (r AccessIdentityProvidersAccessGoogleAppsType) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessGoogleAppsTypeOnetimepin, AccessIdentityProvidersAccessGoogleAppsTypeAzureAd, AccessIdentityProvidersAccessGoogleAppsTypeSaml, AccessIdentityProvidersAccessGoogleAppsTypeCentrify, AccessIdentityProvidersAccessGoogleAppsTypeFacebook, AccessIdentityProvidersAccessGoogleAppsTypeGitHub, AccessIdentityProvidersAccessGoogleAppsTypeGoogleApps, AccessIdentityProvidersAccessGoogleAppsTypeGoogle, AccessIdentityProvidersAccessGoogleAppsTypeLinkedin, AccessIdentityProvidersAccessGoogleAppsTypeOidc, AccessIdentityProvidersAccessGoogleAppsTypeOkta, AccessIdentityProvidersAccessGoogleAppsTypeOnelogin, AccessIdentityProvidersAccessGoogleAppsTypePingone, AccessIdentityProvidersAccessGoogleAppsTypeYandex:
		return true
	}
	return false
}

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

func (r accessIdentityProvidersAccessGoogleAppsScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r accessIdentityProvidersAccessLinkedinJSON) RawJSON() string {
	return r.raw
}

func (r AccessIdentityProvidersAccessLinkedin) implementsZeroTrustAccessIdentityProviders() {}

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

func (r accessIdentityProvidersAccessLinkedinConfigJSON) RawJSON() string {
	return r.raw
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

func (r AccessIdentityProvidersAccessLinkedinType) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessLinkedinTypeOnetimepin, AccessIdentityProvidersAccessLinkedinTypeAzureAd, AccessIdentityProvidersAccessLinkedinTypeSaml, AccessIdentityProvidersAccessLinkedinTypeCentrify, AccessIdentityProvidersAccessLinkedinTypeFacebook, AccessIdentityProvidersAccessLinkedinTypeGitHub, AccessIdentityProvidersAccessLinkedinTypeGoogleApps, AccessIdentityProvidersAccessLinkedinTypeGoogle, AccessIdentityProvidersAccessLinkedinTypeLinkedin, AccessIdentityProvidersAccessLinkedinTypeOidc, AccessIdentityProvidersAccessLinkedinTypeOkta, AccessIdentityProvidersAccessLinkedinTypeOnelogin, AccessIdentityProvidersAccessLinkedinTypePingone, AccessIdentityProvidersAccessLinkedinTypeYandex:
		return true
	}
	return false
}

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

func (r accessIdentityProvidersAccessLinkedinScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r accessIdentityProvidersAccessOidcJSON) RawJSON() string {
	return r.raw
}

func (r AccessIdentityProvidersAccessOidc) implementsZeroTrustAccessIdentityProviders() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AccessIdentityProvidersAccessOidcConfig struct {
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
	TokenURL string                                      `json:"token_url"`
	JSON     accessIdentityProvidersAccessOidcConfigJSON `json:"-"`
}

// accessIdentityProvidersAccessOidcConfigJSON contains the JSON metadata for the
// struct [AccessIdentityProvidersAccessOidcConfig]
type accessIdentityProvidersAccessOidcConfigJSON struct {
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

func (r *AccessIdentityProvidersAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdentityProvidersAccessOidcConfigJSON) RawJSON() string {
	return r.raw
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

func (r AccessIdentityProvidersAccessOidcType) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessOidcTypeOnetimepin, AccessIdentityProvidersAccessOidcTypeAzureAd, AccessIdentityProvidersAccessOidcTypeSaml, AccessIdentityProvidersAccessOidcTypeCentrify, AccessIdentityProvidersAccessOidcTypeFacebook, AccessIdentityProvidersAccessOidcTypeGitHub, AccessIdentityProvidersAccessOidcTypeGoogleApps, AccessIdentityProvidersAccessOidcTypeGoogle, AccessIdentityProvidersAccessOidcTypeLinkedin, AccessIdentityProvidersAccessOidcTypeOidc, AccessIdentityProvidersAccessOidcTypeOkta, AccessIdentityProvidersAccessOidcTypeOnelogin, AccessIdentityProvidersAccessOidcTypePingone, AccessIdentityProvidersAccessOidcTypeYandex:
		return true
	}
	return false
}

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

func (r accessIdentityProvidersAccessOidcScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r accessIdentityProvidersAccessOktaJSON) RawJSON() string {
	return r.raw
}

func (r AccessIdentityProvidersAccessOkta) implementsZeroTrustAccessIdentityProviders() {}

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

func (r accessIdentityProvidersAccessOktaConfigJSON) RawJSON() string {
	return r.raw
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

func (r AccessIdentityProvidersAccessOktaType) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessOktaTypeOnetimepin, AccessIdentityProvidersAccessOktaTypeAzureAd, AccessIdentityProvidersAccessOktaTypeSaml, AccessIdentityProvidersAccessOktaTypeCentrify, AccessIdentityProvidersAccessOktaTypeFacebook, AccessIdentityProvidersAccessOktaTypeGitHub, AccessIdentityProvidersAccessOktaTypeGoogleApps, AccessIdentityProvidersAccessOktaTypeGoogle, AccessIdentityProvidersAccessOktaTypeLinkedin, AccessIdentityProvidersAccessOktaTypeOidc, AccessIdentityProvidersAccessOktaTypeOkta, AccessIdentityProvidersAccessOktaTypeOnelogin, AccessIdentityProvidersAccessOktaTypePingone, AccessIdentityProvidersAccessOktaTypeYandex:
		return true
	}
	return false
}

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

func (r accessIdentityProvidersAccessOktaScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r accessIdentityProvidersAccessOneloginJSON) RawJSON() string {
	return r.raw
}

func (r AccessIdentityProvidersAccessOnelogin) implementsZeroTrustAccessIdentityProviders() {}

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

func (r accessIdentityProvidersAccessOneloginConfigJSON) RawJSON() string {
	return r.raw
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

func (r AccessIdentityProvidersAccessOneloginType) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessOneloginTypeOnetimepin, AccessIdentityProvidersAccessOneloginTypeAzureAd, AccessIdentityProvidersAccessOneloginTypeSaml, AccessIdentityProvidersAccessOneloginTypeCentrify, AccessIdentityProvidersAccessOneloginTypeFacebook, AccessIdentityProvidersAccessOneloginTypeGitHub, AccessIdentityProvidersAccessOneloginTypeGoogleApps, AccessIdentityProvidersAccessOneloginTypeGoogle, AccessIdentityProvidersAccessOneloginTypeLinkedin, AccessIdentityProvidersAccessOneloginTypeOidc, AccessIdentityProvidersAccessOneloginTypeOkta, AccessIdentityProvidersAccessOneloginTypeOnelogin, AccessIdentityProvidersAccessOneloginTypePingone, AccessIdentityProvidersAccessOneloginTypeYandex:
		return true
	}
	return false
}

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

func (r accessIdentityProvidersAccessOneloginScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r accessIdentityProvidersAccessPingoneJSON) RawJSON() string {
	return r.raw
}

func (r AccessIdentityProvidersAccessPingone) implementsZeroTrustAccessIdentityProviders() {}

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

func (r accessIdentityProvidersAccessPingoneConfigJSON) RawJSON() string {
	return r.raw
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

func (r AccessIdentityProvidersAccessPingoneType) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessPingoneTypeOnetimepin, AccessIdentityProvidersAccessPingoneTypeAzureAd, AccessIdentityProvidersAccessPingoneTypeSaml, AccessIdentityProvidersAccessPingoneTypeCentrify, AccessIdentityProvidersAccessPingoneTypeFacebook, AccessIdentityProvidersAccessPingoneTypeGitHub, AccessIdentityProvidersAccessPingoneTypeGoogleApps, AccessIdentityProvidersAccessPingoneTypeGoogle, AccessIdentityProvidersAccessPingoneTypeLinkedin, AccessIdentityProvidersAccessPingoneTypeOidc, AccessIdentityProvidersAccessPingoneTypeOkta, AccessIdentityProvidersAccessPingoneTypeOnelogin, AccessIdentityProvidersAccessPingoneTypePingone, AccessIdentityProvidersAccessPingoneTypeYandex:
		return true
	}
	return false
}

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

func (r accessIdentityProvidersAccessPingoneScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r accessIdentityProvidersAccessSamlJSON) RawJSON() string {
	return r.raw
}

func (r AccessIdentityProvidersAccessSaml) implementsZeroTrustAccessIdentityProviders() {}

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
	IDPPublicCERTs []string `json:"idp_public_certs"`
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
	IDPPublicCERTs     apijson.Field
	IssuerURL          apijson.Field
	SignRequest        apijson.Field
	SSOTargetURL       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessIdentityProvidersAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessIdentityProvidersAccessSamlConfigJSON) RawJSON() string {
	return r.raw
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

func (r accessIdentityProvidersAccessSamlConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
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

func (r AccessIdentityProvidersAccessSamlType) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessSamlTypeOnetimepin, AccessIdentityProvidersAccessSamlTypeAzureAd, AccessIdentityProvidersAccessSamlTypeSaml, AccessIdentityProvidersAccessSamlTypeCentrify, AccessIdentityProvidersAccessSamlTypeFacebook, AccessIdentityProvidersAccessSamlTypeGitHub, AccessIdentityProvidersAccessSamlTypeGoogleApps, AccessIdentityProvidersAccessSamlTypeGoogle, AccessIdentityProvidersAccessSamlTypeLinkedin, AccessIdentityProvidersAccessSamlTypeOidc, AccessIdentityProvidersAccessSamlTypeOkta, AccessIdentityProvidersAccessSamlTypeOnelogin, AccessIdentityProvidersAccessSamlTypePingone, AccessIdentityProvidersAccessSamlTypeYandex:
		return true
	}
	return false
}

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

func (r accessIdentityProvidersAccessSamlScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r accessIdentityProvidersAccessYandexJSON) RawJSON() string {
	return r.raw
}

func (r AccessIdentityProvidersAccessYandex) implementsZeroTrustAccessIdentityProviders() {}

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

func (r accessIdentityProvidersAccessYandexConfigJSON) RawJSON() string {
	return r.raw
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

func (r AccessIdentityProvidersAccessYandexType) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessYandexTypeOnetimepin, AccessIdentityProvidersAccessYandexTypeAzureAd, AccessIdentityProvidersAccessYandexTypeSaml, AccessIdentityProvidersAccessYandexTypeCentrify, AccessIdentityProvidersAccessYandexTypeFacebook, AccessIdentityProvidersAccessYandexTypeGitHub, AccessIdentityProvidersAccessYandexTypeGoogleApps, AccessIdentityProvidersAccessYandexTypeGoogle, AccessIdentityProvidersAccessYandexTypeLinkedin, AccessIdentityProvidersAccessYandexTypeOidc, AccessIdentityProvidersAccessYandexTypeOkta, AccessIdentityProvidersAccessYandexTypeOnelogin, AccessIdentityProvidersAccessYandexTypePingone, AccessIdentityProvidersAccessYandexTypeYandex:
		return true
	}
	return false
}

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

func (r accessIdentityProvidersAccessYandexScimConfigJSON) RawJSON() string {
	return r.raw
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

func (r accessIdentityProvidersAccessOnetimepinJSON) RawJSON() string {
	return r.raw
}

func (r AccessIdentityProvidersAccessOnetimepin) implementsZeroTrustAccessIdentityProviders() {}

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

func (r AccessIdentityProvidersAccessOnetimepinType) IsKnown() bool {
	switch r {
	case AccessIdentityProvidersAccessOnetimepinTypeOnetimepin, AccessIdentityProvidersAccessOnetimepinTypeAzureAd, AccessIdentityProvidersAccessOnetimepinTypeSaml, AccessIdentityProvidersAccessOnetimepinTypeCentrify, AccessIdentityProvidersAccessOnetimepinTypeFacebook, AccessIdentityProvidersAccessOnetimepinTypeGitHub, AccessIdentityProvidersAccessOnetimepinTypeGoogleApps, AccessIdentityProvidersAccessOnetimepinTypeGoogle, AccessIdentityProvidersAccessOnetimepinTypeLinkedin, AccessIdentityProvidersAccessOnetimepinTypeOidc, AccessIdentityProvidersAccessOnetimepinTypeOkta, AccessIdentityProvidersAccessOnetimepinTypeOnelogin, AccessIdentityProvidersAccessOnetimepinTypePingone, AccessIdentityProvidersAccessOnetimepinTypeYandex:
		return true
	}
	return false
}

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

func (r accessIdentityProvidersAccessOnetimepinScimConfigJSON) RawJSON() string {
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
	Result   AccessIdentityProviders                       `json:"result,required"`
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
	Result   AccessIdentityProviders                          `json:"result,required"`
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
	Result   AccessIdentityProviders                       `json:"result,required"`
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
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
func (r *IdentityProviderService) List(ctx context.Context, query IdentityProviderListParams, opts ...option.RequestOption) (res *pagination.SinglePage[IdentityProviderListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all configured identity providers.
func (r *IdentityProviderService) ListAutoPaging(ctx context.Context, query IdentityProviderListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[IdentityProviderListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AzureAD struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AzureADConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig  `json:"scim_config"`
	JSON       azureADJSON `json:"-"`
}

// azureADJSON contains the JSON metadata for the struct [AzureAD]
type azureADJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AzureAD) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r azureADJSON) RawJSON() string {
	return r.raw
}

func (r AzureAD) implementsZeroTrustIdentityProviderNewResponse() {}

func (r AzureAD) implementsZeroTrustIdentityProviderUpdateResponse() {}

func (r AzureAD) implementsZeroTrustIdentityProviderListResponse() {}

func (r AzureAD) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AzureADConfig struct {
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
	Prompt AzureADConfigPrompt `json:"prompt"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool              `json:"support_groups"`
	JSON          azureADConfigJSON `json:"-"`
}

// azureADConfigJSON contains the JSON metadata for the struct [AzureADConfig]
type azureADConfigJSON struct {
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

func (r *AzureADConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r azureADConfigJSON) RawJSON() string {
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
type AzureADConfigPrompt string

const (
	AzureADConfigPromptLogin         AzureADConfigPrompt = "login"
	AzureADConfigPromptSelectAccount AzureADConfigPrompt = "select_account"
	AzureADConfigPromptNone          AzureADConfigPrompt = "none"
)

func (r AzureADConfigPrompt) IsKnown() bool {
	switch r {
	case AzureADConfigPromptLogin, AzureADConfigPromptSelectAccount, AzureADConfigPromptNone:
		return true
	}
	return false
}

type AzureADParam struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[AzureADConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r AzureADParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AzureADParam) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {}

func (r AzureADParam) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AzureADConfigParam struct {
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
	Prompt param.Field[AzureADConfigPrompt] `json:"prompt"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
}

func (r AzureADConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GenericOAuthConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                 `json:"client_secret"`
	JSON         genericOAuthConfigJSON `json:"-"`
}

// genericOAuthConfigJSON contains the JSON metadata for the struct
// [GenericOAuthConfig]
type genericOAuthConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GenericOAuthConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r genericOAuthConfigJSON) RawJSON() string {
	return r.raw
}

type GenericOAuthConfigParam struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r GenericOAuthConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderType string

const (
	IdentityProviderTypeOnetimepin IdentityProviderType = "onetimepin"
	IdentityProviderTypeAzureAD    IdentityProviderType = "azureAD"
	IdentityProviderTypeSAML       IdentityProviderType = "saml"
	IdentityProviderTypeCentrify   IdentityProviderType = "centrify"
	IdentityProviderTypeFacebook   IdentityProviderType = "facebook"
	IdentityProviderTypeGitHub     IdentityProviderType = "github"
	IdentityProviderTypeGoogleApps IdentityProviderType = "google-apps"
	IdentityProviderTypeGoogle     IdentityProviderType = "google"
	IdentityProviderTypeLinkedin   IdentityProviderType = "linkedin"
	IdentityProviderTypeOIDC       IdentityProviderType = "oidc"
	IdentityProviderTypeOkta       IdentityProviderType = "okta"
	IdentityProviderTypeOnelogin   IdentityProviderType = "onelogin"
	IdentityProviderTypePingone    IdentityProviderType = "pingone"
	IdentityProviderTypeYandex     IdentityProviderType = "yandex"
)

func (r IdentityProviderType) IsKnown() bool {
	switch r {
	case IdentityProviderTypeOnetimepin, IdentityProviderTypeAzureAD, IdentityProviderTypeSAML, IdentityProviderTypeCentrify, IdentityProviderTypeFacebook, IdentityProviderTypeGitHub, IdentityProviderTypeGoogleApps, IdentityProviderTypeGoogle, IdentityProviderTypeLinkedin, IdentityProviderTypeOIDC, IdentityProviderTypeOkta, IdentityProviderTypeOnelogin, IdentityProviderTypePingone, IdentityProviderTypeYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool           `json:"user_deprovision"`
	JSON            scimConfigJSON `json:"-"`
}

// scimConfigJSON contains the JSON metadata for the struct [ScimConfig]
type scimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigJSON) RawJSON() string {
	return r.raw
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ScimConfigParam struct {
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

func (r ScimConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewResponse struct {
	Config interface{} `json:"config"`
	// UUID
	ID string `json:"id"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type  IdentityProviderType            `json:"type,required"`
	JSON  identityProviderNewResponseJSON `json:"-"`
	union IdentityProviderNewResponseUnion
}

// identityProviderNewResponseJSON contains the JSON metadata for the struct
// [IdentityProviderNewResponse]
type identityProviderNewResponseJSON struct {
	Config      apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r identityProviderNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *IdentityProviderNewResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r IdentityProviderNewResponse) AsUnion() IdentityProviderNewResponseUnion {
	return r.union
}

// Union satisfied by [zero_trust.AzureAD],
// [zero_trust.IdentityProviderNewResponseAccessCentrify],
// [zero_trust.IdentityProviderNewResponseAccessFacebook],
// [zero_trust.IdentityProviderNewResponseAccessGitHub],
// [zero_trust.IdentityProviderNewResponseAccessGoogle],
// [zero_trust.IdentityProviderNewResponseAccessGoogleApps],
// [zero_trust.IdentityProviderNewResponseAccessLinkedin],
// [zero_trust.IdentityProviderNewResponseAccessOIDC],
// [zero_trust.IdentityProviderNewResponseAccessOkta],
// [zero_trust.IdentityProviderNewResponseAccessOnelogin],
// [zero_trust.IdentityProviderNewResponseAccessPingone],
// [zero_trust.IdentityProviderNewResponseAccessSAML],
// [zero_trust.IdentityProviderNewResponseAccessYandex] or
// [zero_trust.IdentityProviderNewResponseAccessOnetimepin].
type IdentityProviderNewResponseUnion interface {
	implementsZeroTrustIdentityProviderNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IdentityProviderNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AzureAD{}),
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
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessOIDC{}),
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
			Type:       reflect.TypeOf(IdentityProviderNewResponseAccessSAML{}),
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                    `json:"scim_config"`
	JSON       identityProviderNewResponseAccessCentrifyJSON `json:"-"`
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

type IdentityProviderNewResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                    `json:"scim_config"`
	JSON       identityProviderNewResponseAccessFacebookJSON `json:"-"`
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

type IdentityProviderNewResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                  `json:"scim_config"`
	JSON       identityProviderNewResponseAccessGitHubJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                  `json:"scim_config"`
	JSON       identityProviderNewResponseAccessGoogleJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                      `json:"scim_config"`
	JSON       identityProviderNewResponseAccessGoogleAppsJSON `json:"-"`
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

type IdentityProviderNewResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                    `json:"scim_config"`
	JSON       identityProviderNewResponseAccessLinkedinJSON `json:"-"`
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

type IdentityProviderNewResponseAccessOIDC struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessOIDCConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                `json:"scim_config"`
	JSON       identityProviderNewResponseAccessOIDCJSON `json:"-"`
}

// identityProviderNewResponseAccessOIDCJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseAccessOIDC]
type identityProviderNewResponseAccessOIDCJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessOIDC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessOIDCJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessOIDC) implementsZeroTrustIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessOIDCConfig struct {
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
	TokenURL string                                          `json:"token_url"`
	JSON     identityProviderNewResponseAccessOIDCConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessOIDCConfigJSON contains the JSON metadata for
// the struct [IdentityProviderNewResponseAccessOIDCConfig]
type identityProviderNewResponseAccessOIDCConfigJSON struct {
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

func (r *IdentityProviderNewResponseAccessOIDCConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessOIDCConfigJSON) RawJSON() string {
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                `json:"scim_config"`
	JSON       identityProviderNewResponseAccessOktaJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                    `json:"scim_config"`
	JSON       identityProviderNewResponseAccessOneloginJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                   `json:"scim_config"`
	JSON       identityProviderNewResponseAccessPingoneJSON `json:"-"`
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

type IdentityProviderNewResponseAccessSAML struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderNewResponseAccessSAMLConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                `json:"scim_config"`
	JSON       identityProviderNewResponseAccessSAMLJSON `json:"-"`
}

// identityProviderNewResponseAccessSAMLJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseAccessSAML]
type identityProviderNewResponseAccessSAMLJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessSAML) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessSAMLJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderNewResponseAccessSAML) implementsZeroTrustIdentityProviderNewResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewResponseAccessSAMLConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []IdentityProviderNewResponseAccessSAMLConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdPPublicCERTs []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                          `json:"sso_target_url"`
	JSON         identityProviderNewResponseAccessSAMLConfigJSON `json:"-"`
}

// identityProviderNewResponseAccessSAMLConfigJSON contains the JSON metadata for
// the struct [IdentityProviderNewResponseAccessSAMLConfig]
type identityProviderNewResponseAccessSAMLConfigJSON struct {
	Attributes         apijson.Field
	EmailAttributeName apijson.Field
	HeaderAttributes   apijson.Field
	IdPPublicCERTs     apijson.Field
	IssuerURL          apijson.Field
	SignRequest        apijson.Field
	SSOTargetURL       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessSAMLConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessSAMLConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessSAMLConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                         `json:"header_name"`
	JSON       identityProviderNewResponseAccessSAMLConfigHeaderAttributeJSON `json:"-"`
}

// identityProviderNewResponseAccessSAMLConfigHeaderAttributeJSON contains the JSON
// metadata for the struct
// [IdentityProviderNewResponseAccessSAMLConfigHeaderAttribute]
type identityProviderNewResponseAccessSAMLConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityProviderNewResponseAccessSAMLConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseAccessSAMLConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderNewResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                  `json:"scim_config"`
	JSON       identityProviderNewResponseAccessYandexJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                      `json:"scim_config"`
	JSON       identityProviderNewResponseAccessOnetimepinJSON `json:"-"`
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

type IdentityProviderUpdateResponse struct {
	Config interface{} `json:"config"`
	// UUID
	ID string `json:"id"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type  IdentityProviderType               `json:"type,required"`
	JSON  identityProviderUpdateResponseJSON `json:"-"`
	union IdentityProviderUpdateResponseUnion
}

// identityProviderUpdateResponseJSON contains the JSON metadata for the struct
// [IdentityProviderUpdateResponse]
type identityProviderUpdateResponseJSON struct {
	Config      apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r identityProviderUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *IdentityProviderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r IdentityProviderUpdateResponse) AsUnion() IdentityProviderUpdateResponseUnion {
	return r.union
}

// Union satisfied by [zero_trust.AzureAD],
// [zero_trust.IdentityProviderUpdateResponseAccessCentrify],
// [zero_trust.IdentityProviderUpdateResponseAccessFacebook],
// [zero_trust.IdentityProviderUpdateResponseAccessGitHub],
// [zero_trust.IdentityProviderUpdateResponseAccessGoogle],
// [zero_trust.IdentityProviderUpdateResponseAccessGoogleApps],
// [zero_trust.IdentityProviderUpdateResponseAccessLinkedin],
// [zero_trust.IdentityProviderUpdateResponseAccessOIDC],
// [zero_trust.IdentityProviderUpdateResponseAccessOkta],
// [zero_trust.IdentityProviderUpdateResponseAccessOnelogin],
// [zero_trust.IdentityProviderUpdateResponseAccessPingone],
// [zero_trust.IdentityProviderUpdateResponseAccessSAML],
// [zero_trust.IdentityProviderUpdateResponseAccessYandex] or
// [zero_trust.IdentityProviderUpdateResponseAccessOnetimepin].
type IdentityProviderUpdateResponseUnion interface {
	implementsZeroTrustIdentityProviderUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IdentityProviderUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AzureAD{}),
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
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessOIDC{}),
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
			Type:       reflect.TypeOf(IdentityProviderUpdateResponseAccessSAML{}),
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                       `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessCentrifyJSON `json:"-"`
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

type IdentityProviderUpdateResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                       `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessFacebookJSON `json:"-"`
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

type IdentityProviderUpdateResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                     `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessGitHubJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                     `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessGoogleJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                         `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessGoogleAppsJSON `json:"-"`
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

type IdentityProviderUpdateResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                       `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessLinkedinJSON `json:"-"`
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

type IdentityProviderUpdateResponseAccessOIDC struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessOIDCConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                   `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessOIDCJSON `json:"-"`
}

// identityProviderUpdateResponseAccessOIDCJSON contains the JSON metadata for the
// struct [IdentityProviderUpdateResponseAccessOIDC]
type identityProviderUpdateResponseAccessOIDCJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessOIDC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessOIDCJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessOIDC) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessOIDCConfig struct {
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
	TokenURL string                                             `json:"token_url"`
	JSON     identityProviderUpdateResponseAccessOIDCConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessOIDCConfigJSON contains the JSON metadata
// for the struct [IdentityProviderUpdateResponseAccessOIDCConfig]
type identityProviderUpdateResponseAccessOIDCConfigJSON struct {
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

func (r *IdentityProviderUpdateResponseAccessOIDCConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessOIDCConfigJSON) RawJSON() string {
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                   `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessOktaJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                       `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessOneloginJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                      `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessPingoneJSON `json:"-"`
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

type IdentityProviderUpdateResponseAccessSAML struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderUpdateResponseAccessSAMLConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                   `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessSAMLJSON `json:"-"`
}

// identityProviderUpdateResponseAccessSAMLJSON contains the JSON metadata for the
// struct [IdentityProviderUpdateResponseAccessSAML]
type identityProviderUpdateResponseAccessSAMLJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessSAML) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessSAMLJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderUpdateResponseAccessSAML) implementsZeroTrustIdentityProviderUpdateResponse() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateResponseAccessSAMLConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []IdentityProviderUpdateResponseAccessSAMLConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdPPublicCERTs []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                             `json:"sso_target_url"`
	JSON         identityProviderUpdateResponseAccessSAMLConfigJSON `json:"-"`
}

// identityProviderUpdateResponseAccessSAMLConfigJSON contains the JSON metadata
// for the struct [IdentityProviderUpdateResponseAccessSAMLConfig]
type identityProviderUpdateResponseAccessSAMLConfigJSON struct {
	Attributes         apijson.Field
	EmailAttributeName apijson.Field
	HeaderAttributes   apijson.Field
	IdPPublicCERTs     apijson.Field
	IssuerURL          apijson.Field
	SignRequest        apijson.Field
	SSOTargetURL       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessSAMLConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessSAMLConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessSAMLConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                            `json:"header_name"`
	JSON       identityProviderUpdateResponseAccessSAMLConfigHeaderAttributeJSON `json:"-"`
}

// identityProviderUpdateResponseAccessSAMLConfigHeaderAttributeJSON contains the
// JSON metadata for the struct
// [IdentityProviderUpdateResponseAccessSAMLConfigHeaderAttribute]
type identityProviderUpdateResponseAccessSAMLConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseAccessSAMLConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseAccessSAMLConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderUpdateResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                     `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessYandexJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                         `json:"scim_config"`
	JSON       identityProviderUpdateResponseAccessOnetimepinJSON `json:"-"`
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

type IdentityProviderListResponse struct {
	Config interface{} `json:"config"`
	// UUID
	ID string `json:"id"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type  IdentityProviderType             `json:"type,required"`
	JSON  identityProviderListResponseJSON `json:"-"`
	union IdentityProviderListResponseUnion
}

// identityProviderListResponseJSON contains the JSON metadata for the struct
// [IdentityProviderListResponse]
type identityProviderListResponseJSON struct {
	Config      apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r identityProviderListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *IdentityProviderListResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r IdentityProviderListResponse) AsUnion() IdentityProviderListResponseUnion {
	return r.union
}

// Union satisfied by [zero_trust.AzureAD],
// [zero_trust.IdentityProviderListResponseAccessCentrify],
// [zero_trust.IdentityProviderListResponseAccessFacebook],
// [zero_trust.IdentityProviderListResponseAccessGitHub],
// [zero_trust.IdentityProviderListResponseAccessGoogle],
// [zero_trust.IdentityProviderListResponseAccessGoogleApps],
// [zero_trust.IdentityProviderListResponseAccessLinkedin],
// [zero_trust.IdentityProviderListResponseAccessOIDC],
// [zero_trust.IdentityProviderListResponseAccessOkta],
// [zero_trust.IdentityProviderListResponseAccessOnelogin],
// [zero_trust.IdentityProviderListResponseAccessPingone],
// [zero_trust.IdentityProviderListResponseAccessSAML] or
// [zero_trust.IdentityProviderListResponseAccessYandex].
type IdentityProviderListResponseUnion interface {
	implementsZeroTrustIdentityProviderListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IdentityProviderListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AzureAD{}),
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
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessOIDC{}),
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
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessSAML{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderListResponseAccessYandex{}),
		},
	)
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                     `json:"scim_config"`
	JSON       identityProviderListResponseAccessCentrifyJSON `json:"-"`
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

type IdentityProviderListResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                     `json:"scim_config"`
	JSON       identityProviderListResponseAccessFacebookJSON `json:"-"`
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

type IdentityProviderListResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                   `json:"scim_config"`
	JSON       identityProviderListResponseAccessGitHubJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                   `json:"scim_config"`
	JSON       identityProviderListResponseAccessGoogleJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                       `json:"scim_config"`
	JSON       identityProviderListResponseAccessGoogleAppsJSON `json:"-"`
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

type IdentityProviderListResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                     `json:"scim_config"`
	JSON       identityProviderListResponseAccessLinkedinJSON `json:"-"`
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

type IdentityProviderListResponseAccessOIDC struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessOIDCConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                 `json:"scim_config"`
	JSON       identityProviderListResponseAccessOIDCJSON `json:"-"`
}

// identityProviderListResponseAccessOIDCJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessOIDC]
type identityProviderListResponseAccessOIDCJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessOIDC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessOIDCJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessOIDC) implementsZeroTrustIdentityProviderListResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessOIDCConfig struct {
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
	JSON     identityProviderListResponseAccessOIDCConfigJSON `json:"-"`
}

// identityProviderListResponseAccessOIDCConfigJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessOIDCConfig]
type identityProviderListResponseAccessOIDCConfigJSON struct {
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

func (r *IdentityProviderListResponseAccessOIDCConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessOIDCConfigJSON) RawJSON() string {
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                 `json:"scim_config"`
	JSON       identityProviderListResponseAccessOktaJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                     `json:"scim_config"`
	JSON       identityProviderListResponseAccessOneloginJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                    `json:"scim_config"`
	JSON       identityProviderListResponseAccessPingoneJSON `json:"-"`
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

type IdentityProviderListResponseAccessSAML struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderListResponseAccessSAMLConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                 `json:"scim_config"`
	JSON       identityProviderListResponseAccessSAMLJSON `json:"-"`
}

// identityProviderListResponseAccessSAMLJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessSAML]
type identityProviderListResponseAccessSAMLJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessSAML) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessSAMLJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderListResponseAccessSAML) implementsZeroTrustIdentityProviderListResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderListResponseAccessSAMLConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []IdentityProviderListResponseAccessSAMLConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdPPublicCERTs []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                           `json:"sso_target_url"`
	JSON         identityProviderListResponseAccessSAMLConfigJSON `json:"-"`
}

// identityProviderListResponseAccessSAMLConfigJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessSAMLConfig]
type identityProviderListResponseAccessSAMLConfigJSON struct {
	Attributes         apijson.Field
	EmailAttributeName apijson.Field
	HeaderAttributes   apijson.Field
	IdPPublicCERTs     apijson.Field
	IssuerURL          apijson.Field
	SignRequest        apijson.Field
	SSOTargetURL       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessSAMLConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessSAMLConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessSAMLConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                          `json:"header_name"`
	JSON       identityProviderListResponseAccessSAMLConfigHeaderAttributeJSON `json:"-"`
}

// identityProviderListResponseAccessSAMLConfigHeaderAttributeJSON contains the
// JSON metadata for the struct
// [IdentityProviderListResponseAccessSAMLConfigHeaderAttribute]
type identityProviderListResponseAccessSAMLConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityProviderListResponseAccessSAMLConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderListResponseAccessSAMLConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderListResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                   `json:"scim_config"`
	JSON       identityProviderListResponseAccessYandexJSON `json:"-"`
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

type IdentityProviderGetResponse struct {
	Config interface{} `json:"config"`
	// UUID
	ID string `json:"id"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type  IdentityProviderType            `json:"type,required"`
	JSON  identityProviderGetResponseJSON `json:"-"`
	union IdentityProviderGetResponseUnion
}

// identityProviderGetResponseJSON contains the JSON metadata for the struct
// [IdentityProviderGetResponse]
type identityProviderGetResponseJSON struct {
	Config      apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r identityProviderGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *IdentityProviderGetResponse) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r IdentityProviderGetResponse) AsUnion() IdentityProviderGetResponseUnion {
	return r.union
}

// Union satisfied by [zero_trust.AzureAD],
// [zero_trust.IdentityProviderGetResponseAccessCentrify],
// [zero_trust.IdentityProviderGetResponseAccessFacebook],
// [zero_trust.IdentityProviderGetResponseAccessGitHub],
// [zero_trust.IdentityProviderGetResponseAccessGoogle],
// [zero_trust.IdentityProviderGetResponseAccessGoogleApps],
// [zero_trust.IdentityProviderGetResponseAccessLinkedin],
// [zero_trust.IdentityProviderGetResponseAccessOIDC],
// [zero_trust.IdentityProviderGetResponseAccessOkta],
// [zero_trust.IdentityProviderGetResponseAccessOnelogin],
// [zero_trust.IdentityProviderGetResponseAccessPingone],
// [zero_trust.IdentityProviderGetResponseAccessSAML],
// [zero_trust.IdentityProviderGetResponseAccessYandex] or
// [zero_trust.IdentityProviderGetResponseAccessOnetimepin].
type IdentityProviderGetResponseUnion interface {
	implementsZeroTrustIdentityProviderGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IdentityProviderGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AzureAD{}),
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
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessOIDC{}),
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
			Type:       reflect.TypeOf(IdentityProviderGetResponseAccessSAML{}),
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                    `json:"scim_config"`
	JSON       identityProviderGetResponseAccessCentrifyJSON `json:"-"`
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

type IdentityProviderGetResponseAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                    `json:"scim_config"`
	JSON       identityProviderGetResponseAccessFacebookJSON `json:"-"`
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

type IdentityProviderGetResponseAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                  `json:"scim_config"`
	JSON       identityProviderGetResponseAccessGitHubJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                  `json:"scim_config"`
	JSON       identityProviderGetResponseAccessGoogleJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                      `json:"scim_config"`
	JSON       identityProviderGetResponseAccessGoogleAppsJSON `json:"-"`
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

type IdentityProviderGetResponseAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                    `json:"scim_config"`
	JSON       identityProviderGetResponseAccessLinkedinJSON `json:"-"`
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

type IdentityProviderGetResponseAccessOIDC struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessOIDCConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                `json:"scim_config"`
	JSON       identityProviderGetResponseAccessOIDCJSON `json:"-"`
}

// identityProviderGetResponseAccessOIDCJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseAccessOIDC]
type identityProviderGetResponseAccessOIDCJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessOIDC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessOIDCJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessOIDC) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessOIDCConfig struct {
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
	TokenURL string                                          `json:"token_url"`
	JSON     identityProviderGetResponseAccessOIDCConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessOIDCConfigJSON contains the JSON metadata for
// the struct [IdentityProviderGetResponseAccessOIDCConfig]
type identityProviderGetResponseAccessOIDCConfigJSON struct {
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

func (r *IdentityProviderGetResponseAccessOIDCConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessOIDCConfigJSON) RawJSON() string {
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                `json:"scim_config"`
	JSON       identityProviderGetResponseAccessOktaJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                    `json:"scim_config"`
	JSON       identityProviderGetResponseAccessOneloginJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                   `json:"scim_config"`
	JSON       identityProviderGetResponseAccessPingoneJSON `json:"-"`
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

type IdentityProviderGetResponseAccessSAML struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderGetResponseAccessSAMLConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                `json:"scim_config"`
	JSON       identityProviderGetResponseAccessSAMLJSON `json:"-"`
}

// identityProviderGetResponseAccessSAMLJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseAccessSAML]
type identityProviderGetResponseAccessSAMLJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessSAML) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessSAMLJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderGetResponseAccessSAML) implementsZeroTrustIdentityProviderGetResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderGetResponseAccessSAMLConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []IdentityProviderGetResponseAccessSAMLConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdPPublicCERTs []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                          `json:"sso_target_url"`
	JSON         identityProviderGetResponseAccessSAMLConfigJSON `json:"-"`
}

// identityProviderGetResponseAccessSAMLConfigJSON contains the JSON metadata for
// the struct [IdentityProviderGetResponseAccessSAMLConfig]
type identityProviderGetResponseAccessSAMLConfigJSON struct {
	Attributes         apijson.Field
	EmailAttributeName apijson.Field
	HeaderAttributes   apijson.Field
	IdPPublicCERTs     apijson.Field
	IssuerURL          apijson.Field
	SignRequest        apijson.Field
	SSOTargetURL       apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessSAMLConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessSAMLConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessSAMLConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                         `json:"header_name"`
	JSON       identityProviderGetResponseAccessSAMLConfigHeaderAttributeJSON `json:"-"`
}

// identityProviderGetResponseAccessSAMLConfigHeaderAttributeJSON contains the JSON
// metadata for the struct
// [IdentityProviderGetResponseAccessSAMLConfigHeaderAttribute]
type identityProviderGetResponseAccessSAMLConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityProviderGetResponseAccessSAMLConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseAccessSAMLConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderGetResponseAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                  `json:"scim_config"`
	JSON       identityProviderGetResponseAccessYandexJSON `json:"-"`
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ScimConfig                                      `json:"scim_config"`
	JSON       identityProviderGetResponseAccessOnetimepinJSON `json:"-"`
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

type IdentityProviderNewParams struct {
	Body IdentityProviderNewParamsBodyUnion `json:"body,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r IdentityProviderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type IdentityProviderNewParamsBody struct {
	Config param.Field[interface{}] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
}

func (r IdentityProviderNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsBody) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {}

// Satisfied by [zero_trust.AzureADParam],
// [zero_trust.IdentityProviderNewParamsBodyAccessCentrify],
// [zero_trust.IdentityProviderNewParamsBodyAccessFacebook],
// [zero_trust.IdentityProviderNewParamsBodyAccessGitHub],
// [zero_trust.IdentityProviderNewParamsBodyAccessGoogle],
// [zero_trust.IdentityProviderNewParamsBodyAccessGoogleApps],
// [zero_trust.IdentityProviderNewParamsBodyAccessLinkedin],
// [zero_trust.IdentityProviderNewParamsBodyAccessOIDC],
// [zero_trust.IdentityProviderNewParamsBodyAccessOkta],
// [zero_trust.IdentityProviderNewParamsBodyAccessOnelogin],
// [zero_trust.IdentityProviderNewParamsBodyAccessPingone],
// [zero_trust.IdentityProviderNewParamsBodyAccessSAML],
// [zero_trust.IdentityProviderNewParamsBodyAccessYandex],
// [zero_trust.IdentityProviderNewParamsBodyAccessOnetimepin],
// [IdentityProviderNewParamsBody].
type IdentityProviderNewParamsBodyUnion interface {
	implementsZeroTrustIdentityProviderNewParamsBodyUnion()
}

type IdentityProviderNewParamsBodyAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsBodyAccessCentrifyConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderNewParamsBodyAccessCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsBodyAccessCentrify) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsBodyAccessCentrifyConfig struct {
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

func (r IdentityProviderNewParamsBodyAccessCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsBodyAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderNewParamsBodyAccessFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsBodyAccessFacebook) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {
}

type IdentityProviderNewParamsBodyAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderNewParamsBodyAccessGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsBodyAccessGitHub) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {
}

type IdentityProviderNewParamsBodyAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsBodyAccessGoogleConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderNewParamsBodyAccessGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsBodyAccessGoogle) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsBodyAccessGoogleConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r IdentityProviderNewParamsBodyAccessGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsBodyAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsBodyAccessGoogleAppsConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderNewParamsBodyAccessGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsBodyAccessGoogleApps) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsBodyAccessGoogleAppsConfig struct {
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

func (r IdentityProviderNewParamsBodyAccessGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsBodyAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderNewParamsBodyAccessLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsBodyAccessLinkedin) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {
}

type IdentityProviderNewParamsBodyAccessOIDC struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsBodyAccessOIDCConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderNewParamsBodyAccessOIDC) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsBodyAccessOIDC) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsBodyAccessOIDCConfig struct {
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

func (r IdentityProviderNewParamsBodyAccessOIDCConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsBodyAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsBodyAccessOktaConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderNewParamsBodyAccessOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsBodyAccessOkta) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsBodyAccessOktaConfig struct {
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

func (r IdentityProviderNewParamsBodyAccessOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsBodyAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsBodyAccessOneloginConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderNewParamsBodyAccessOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsBodyAccessOnelogin) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsBodyAccessOneloginConfig struct {
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

func (r IdentityProviderNewParamsBodyAccessOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsBodyAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsBodyAccessPingoneConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderNewParamsBodyAccessPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsBodyAccessPingone) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsBodyAccessPingoneConfig struct {
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

func (r IdentityProviderNewParamsBodyAccessPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsBodyAccessSAML struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsBodyAccessSAMLConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderNewParamsBodyAccessSAML) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsBodyAccessSAML) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsBodyAccessSAMLConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]IdentityProviderNewParamsBodyAccessSAMLConfigHeaderAttribute] `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdPPublicCERTs param.Field[[]string] `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL param.Field[string] `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest param.Field[bool] `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL param.Field[string] `json:"sso_target_url"`
}

func (r IdentityProviderNewParamsBodyAccessSAMLConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsBodyAccessSAMLConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r IdentityProviderNewParamsBodyAccessSAMLConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderNewParamsBodyAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderNewParamsBodyAccessYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsBodyAccessYandex) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {
}

type IdentityProviderNewParamsBodyAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[interface{}] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderNewParamsBodyAccessOnetimepin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsBodyAccessOnetimepin) implementsZeroTrustIdentityProviderNewParamsBodyUnion() {
}

type IdentityProviderNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IdentityProviderNewResponseEnvelopeSuccess `json:"success,required"`
	Result  IdentityProviderNewResponse                `json:"result"`
	JSON    identityProviderNewResponseEnvelopeJSON    `json:"-"`
}

// identityProviderNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [IdentityProviderNewResponseEnvelope]
type identityProviderNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderNewResponseEnvelopeJSON) RawJSON() string {
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

type IdentityProviderUpdateParams struct {
	Body IdentityProviderUpdateParamsBodyUnion `json:"body,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r IdentityProviderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type IdentityProviderUpdateParamsBody struct {
	Config param.Field[interface{}] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
}

func (r IdentityProviderUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsBody) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {
}

// Satisfied by [zero_trust.AzureADParam],
// [zero_trust.IdentityProviderUpdateParamsBodyAccessCentrify],
// [zero_trust.IdentityProviderUpdateParamsBodyAccessFacebook],
// [zero_trust.IdentityProviderUpdateParamsBodyAccessGitHub],
// [zero_trust.IdentityProviderUpdateParamsBodyAccessGoogle],
// [zero_trust.IdentityProviderUpdateParamsBodyAccessGoogleApps],
// [zero_trust.IdentityProviderUpdateParamsBodyAccessLinkedin],
// [zero_trust.IdentityProviderUpdateParamsBodyAccessOIDC],
// [zero_trust.IdentityProviderUpdateParamsBodyAccessOkta],
// [zero_trust.IdentityProviderUpdateParamsBodyAccessOnelogin],
// [zero_trust.IdentityProviderUpdateParamsBodyAccessPingone],
// [zero_trust.IdentityProviderUpdateParamsBodyAccessSAML],
// [zero_trust.IdentityProviderUpdateParamsBodyAccessYandex],
// [zero_trust.IdentityProviderUpdateParamsBodyAccessOnetimepin],
// [IdentityProviderUpdateParamsBody].
type IdentityProviderUpdateParamsBodyUnion interface {
	implementsZeroTrustIdentityProviderUpdateParamsBodyUnion()
}

type IdentityProviderUpdateParamsBodyAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsBodyAccessCentrifyConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsBodyAccessCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsBodyAccessCentrify) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsBodyAccessCentrifyConfig struct {
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

func (r IdentityProviderUpdateParamsBodyAccessCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsBodyAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsBodyAccessFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsBodyAccessFacebook) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {
}

type IdentityProviderUpdateParamsBodyAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsBodyAccessGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsBodyAccessGitHub) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {
}

type IdentityProviderUpdateParamsBodyAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsBodyAccessGoogleConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsBodyAccessGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsBodyAccessGoogle) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsBodyAccessGoogleConfig struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r IdentityProviderUpdateParamsBodyAccessGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsBodyAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsBodyAccessGoogleAppsConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsBodyAccessGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsBodyAccessGoogleApps) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsBodyAccessGoogleAppsConfig struct {
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

func (r IdentityProviderUpdateParamsBodyAccessGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsBodyAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsBodyAccessLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsBodyAccessLinkedin) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {
}

type IdentityProviderUpdateParamsBodyAccessOIDC struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsBodyAccessOIDCConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsBodyAccessOIDC) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsBodyAccessOIDC) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsBodyAccessOIDCConfig struct {
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

func (r IdentityProviderUpdateParamsBodyAccessOIDCConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsBodyAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsBodyAccessOktaConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsBodyAccessOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsBodyAccessOkta) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsBodyAccessOktaConfig struct {
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

func (r IdentityProviderUpdateParamsBodyAccessOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsBodyAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsBodyAccessOneloginConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsBodyAccessOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsBodyAccessOnelogin) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsBodyAccessOneloginConfig struct {
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

func (r IdentityProviderUpdateParamsBodyAccessOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsBodyAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsBodyAccessPingoneConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsBodyAccessPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsBodyAccessPingone) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsBodyAccessPingoneConfig struct {
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

func (r IdentityProviderUpdateParamsBodyAccessPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsBodyAccessSAML struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsBodyAccessSAMLConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsBodyAccessSAML) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsBodyAccessSAML) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {
}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsBodyAccessSAMLConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]IdentityProviderUpdateParamsBodyAccessSAMLConfigHeaderAttribute] `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdPPublicCERTs param.Field[[]string] `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL param.Field[string] `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest param.Field[bool] `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL param.Field[string] `json:"sso_target_url"`
}

func (r IdentityProviderUpdateParamsBodyAccessSAMLConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsBodyAccessSAMLConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r IdentityProviderUpdateParamsBodyAccessSAMLConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderUpdateParamsBodyAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsBodyAccessYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsBodyAccessYandex) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {
}

type IdentityProviderUpdateParamsBodyAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[interface{}] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ScimConfigParam] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsBodyAccessOnetimepin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsBodyAccessOnetimepin) implementsZeroTrustIdentityProviderUpdateParamsBodyUnion() {
}

type IdentityProviderUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IdentityProviderUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  IdentityProviderUpdateResponse                `json:"result"`
	JSON    identityProviderUpdateResponseEnvelopeJSON    `json:"-"`
}

// identityProviderUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [IdentityProviderUpdateResponseEnvelope]
type identityProviderUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderUpdateResponseEnvelopeJSON) RawJSON() string {
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

type IdentityProviderDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type IdentityProviderDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IdentityProviderDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  IdentityProviderDeleteResponse                `json:"result"`
	JSON    identityProviderDeleteResponseEnvelopeJSON    `json:"-"`
}

// identityProviderDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [IdentityProviderDeleteResponseEnvelope]
type identityProviderDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderDeleteResponseEnvelopeJSON) RawJSON() string {
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IdentityProviderGetResponseEnvelopeSuccess `json:"success,required"`
	Result  IdentityProviderGetResponse                `json:"result"`
	JSON    identityProviderGetResponseEnvelopeJSON    `json:"-"`
}

// identityProviderGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [IdentityProviderGetResponseEnvelope]
type identityProviderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderGetResponseEnvelopeJSON) RawJSON() string {
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

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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
func (r *IdentityProviderService) New(ctx context.Context, params IdentityProviderNewParams, opts ...option.RequestOption) (res *IdentityProvider, err error) {
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
func (r *IdentityProviderService) Update(ctx context.Context, uuid string, params IdentityProviderUpdateParams, opts ...option.RequestOption) (res *IdentityProvider, err error) {
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
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *IdentityProviderService) Get(ctx context.Context, uuid string, query IdentityProviderGetParams, opts ...option.RequestOption) (res *IdentityProvider, err error) {
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

type AzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config AzureAdConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       azureAdJSON                                      `json:"-"`
}

// azureAdJSON contains the JSON metadata for the struct [AzureAd]
type azureAdJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r azureAdJSON) RawJSON() string {
	return r.raw
}

func (r AzureAd) implementsZeroTrustIdentityProvider() {}

func (r AzureAd) implementsZeroTrustIdentityProviderListResponse() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type AzureAdConfig struct {
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
	Prompt AzureAdConfigPrompt `json:"prompt"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool              `json:"support_groups"`
	JSON          azureAdConfigJSON `json:"-"`
}

// azureAdConfigJSON contains the JSON metadata for the struct [AzureAdConfig]
type azureAdConfigJSON struct {
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

func (r *AzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r azureAdConfigJSON) RawJSON() string {
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
type AzureAdConfigPrompt string

const (
	AzureAdConfigPromptLogin         AzureAdConfigPrompt = "login"
	AzureAdConfigPromptSelectAccount AzureAdConfigPrompt = "select_account"
	AzureAdConfigPromptNone          AzureAdConfigPrompt = "none"
)

func (r AzureAdConfigPrompt) IsKnown() bool {
	switch r {
	case AzureAdConfigPromptLogin, AzureAdConfigPromptSelectAccount, AzureAdConfigPromptNone:
		return true
	}
	return false
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

type IdentityProvider struct {
	Config interface{} `json:"config"`
	// UUID
	ID string `json:"id"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type  UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	JSON  identityProviderJSON                             `json:"-"`
	union IdentityProviderUnion
}

// identityProviderJSON contains the JSON metadata for the struct
// [IdentityProvider]
type identityProviderJSON struct {
	Config      apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r identityProviderJSON) RawJSON() string {
	return r.raw
}

func (r *IdentityProvider) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r IdentityProvider) AsUnion() IdentityProviderUnion {
	return r.union
}

// Union satisfied by [zero_trust.AzureAd],
// [zero_trust.IdentityProviderAccessCentrify],
// [zero_trust.IdentityProviderAccessFacebook],
// [zero_trust.IdentityProviderAccessGitHub],
// [zero_trust.IdentityProviderAccessGoogle],
// [zero_trust.IdentityProviderAccessGoogleApps],
// [zero_trust.IdentityProviderAccessLinkedin],
// [zero_trust.IdentityProviderAccessOidc],
// [zero_trust.IdentityProviderAccessOkta],
// [zero_trust.IdentityProviderAccessOnelogin],
// [zero_trust.IdentityProviderAccessPingone],
// [zero_trust.IdentityProviderAccessSaml],
// [zero_trust.IdentityProviderAccessYandex] or
// [zero_trust.IdentityProviderAccessOnetimepin].
type IdentityProviderUnion interface {
	implementsZeroTrustIdentityProvider()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IdentityProviderUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AzureAd{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderAccessCentrify{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderAccessFacebook{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderAccessGitHub{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderAccessGoogle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderAccessGoogleApps{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderAccessLinkedin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderAccessOidc{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderAccessOkta{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderAccessOnelogin{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderAccessPingone{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderAccessSaml{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderAccessYandex{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IdentityProviderAccessOnetimepin{}),
		},
	)
}

type IdentityProviderAccessCentrify struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderAccessCentrifyConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderAccessCentrifyJSON               `json:"-"`
}

// identityProviderAccessCentrifyJSON contains the JSON metadata for the struct
// [IdentityProviderAccessCentrify]
type identityProviderAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessCentrifyJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessCentrify) implementsZeroTrustIdentityProvider() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessCentrifyConfig struct {
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
	EmailClaimName string                                   `json:"email_claim_name"`
	JSON           identityProviderAccessCentrifyConfigJSON `json:"-"`
}

// identityProviderAccessCentrifyConfigJSON contains the JSON metadata for the
// struct [IdentityProviderAccessCentrifyConfig]
type identityProviderAccessCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IdentityProviderAccessCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessCentrifyConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderAccessFacebookJSON               `json:"-"`
}

// identityProviderAccessFacebookJSON contains the JSON metadata for the struct
// [IdentityProviderAccessFacebook]
type identityProviderAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessFacebookJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessFacebook) implementsZeroTrustIdentityProvider() {}

type IdentityProviderAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderAccessGitHubJSON                 `json:"-"`
}

// identityProviderAccessGitHubJSON contains the JSON metadata for the struct
// [IdentityProviderAccessGitHub]
type identityProviderAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessGitHubJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessGitHub) implementsZeroTrustIdentityProvider() {}

type IdentityProviderAccessGoogle struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderAccessGoogleConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderAccessGoogleJSON                 `json:"-"`
}

// identityProviderAccessGoogleJSON contains the JSON metadata for the struct
// [IdentityProviderAccessGoogle]
type identityProviderAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessGoogleJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessGoogle) implementsZeroTrustIdentityProvider() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessGoogleConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                 `json:"email_claim_name"`
	JSON           identityProviderAccessGoogleConfigJSON `json:"-"`
}

// identityProviderAccessGoogleConfigJSON contains the JSON metadata for the struct
// [IdentityProviderAccessGoogleConfig]
type identityProviderAccessGoogleConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderAccessGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessGoogleConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderAccessGoogleApps struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderAccessGoogleAppsConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderAccessGoogleAppsJSON             `json:"-"`
}

// identityProviderAccessGoogleAppsJSON contains the JSON metadata for the struct
// [IdentityProviderAccessGoogleApps]
type identityProviderAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessGoogleAppsJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessGoogleApps) implementsZeroTrustIdentityProvider() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string                                     `json:"email_claim_name"`
	JSON           identityProviderAccessGoogleAppsConfigJSON `json:"-"`
}

// identityProviderAccessGoogleAppsConfigJSON contains the JSON metadata for the
// struct [IdentityProviderAccessGoogleAppsConfig]
type identityProviderAccessGoogleAppsConfigJSON struct {
	AppsDomain     apijson.Field
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderAccessGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessGoogleAppsConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderAccessLinkedinJSON               `json:"-"`
}

// identityProviderAccessLinkedinJSON contains the JSON metadata for the struct
// [IdentityProviderAccessLinkedin]
type identityProviderAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessLinkedinJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessLinkedin) implementsZeroTrustIdentityProvider() {}

type IdentityProviderAccessOidc struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderAccessOidcConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderAccessOidcJSON                   `json:"-"`
}

// identityProviderAccessOidcJSON contains the JSON metadata for the struct
// [IdentityProviderAccessOidc]
type identityProviderAccessOidcJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessOidcJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessOidc) implementsZeroTrustIdentityProvider() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessOidcConfig struct {
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
	TokenURL string                               `json:"token_url"`
	JSON     identityProviderAccessOidcConfigJSON `json:"-"`
}

// identityProviderAccessOidcConfigJSON contains the JSON metadata for the struct
// [IdentityProviderAccessOidcConfig]
type identityProviderAccessOidcConfigJSON struct {
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

func (r *IdentityProviderAccessOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessOidcConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderAccessOkta struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderAccessOktaConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderAccessOktaJSON                   `json:"-"`
}

// identityProviderAccessOktaJSON contains the JSON metadata for the struct
// [IdentityProviderAccessOkta]
type identityProviderAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessOktaJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessOkta) implementsZeroTrustIdentityProvider() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessOktaConfig struct {
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
	OktaAccount string                               `json:"okta_account"`
	JSON        identityProviderAccessOktaConfigJSON `json:"-"`
}

// identityProviderAccessOktaConfigJSON contains the JSON metadata for the struct
// [IdentityProviderAccessOktaConfig]
type identityProviderAccessOktaConfigJSON struct {
	AuthorizationServerID apijson.Field
	Claims                apijson.Field
	ClientID              apijson.Field
	ClientSecret          apijson.Field
	EmailClaimName        apijson.Field
	OktaAccount           apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *IdentityProviderAccessOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessOktaConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderAccessOnelogin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderAccessOneloginConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderAccessOneloginJSON               `json:"-"`
}

// identityProviderAccessOneloginJSON contains the JSON metadata for the struct
// [IdentityProviderAccessOnelogin]
type identityProviderAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessOneloginJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessOnelogin) implementsZeroTrustIdentityProvider() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessOneloginConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your OneLogin account url
	OneloginAccount string                                   `json:"onelogin_account"`
	JSON            identityProviderAccessOneloginConfigJSON `json:"-"`
}

// identityProviderAccessOneloginConfigJSON contains the JSON metadata for the
// struct [IdentityProviderAccessOneloginConfig]
type identityProviderAccessOneloginConfigJSON struct {
	Claims          apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	EmailClaimName  apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IdentityProviderAccessOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessOneloginConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderAccessPingone struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderAccessPingoneConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderAccessPingoneJSON                `json:"-"`
}

// identityProviderAccessPingoneJSON contains the JSON metadata for the struct
// [IdentityProviderAccessPingone]
type identityProviderAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessPingoneJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessPingone) implementsZeroTrustIdentityProvider() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessPingoneConfig struct {
	// Custom claims
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName string `json:"email_claim_name"`
	// Your PingOne environment identifier
	PingEnvID string                                  `json:"ping_env_id"`
	JSON      identityProviderAccessPingoneConfigJSON `json:"-"`
}

// identityProviderAccessPingoneConfigJSON contains the JSON metadata for the
// struct [IdentityProviderAccessPingoneConfig]
type identityProviderAccessPingoneConfigJSON struct {
	Claims         apijson.Field
	ClientID       apijson.Field
	ClientSecret   apijson.Field
	EmailClaimName apijson.Field
	PingEnvID      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *IdentityProviderAccessPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessPingoneConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderAccessSaml struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderAccessSamlConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderAccessSamlJSON                   `json:"-"`
}

// identityProviderAccessSamlJSON contains the JSON metadata for the struct
// [IdentityProviderAccessSaml]
type identityProviderAccessSamlJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessSamlJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessSaml) implementsZeroTrustIdentityProvider() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []IdentityProviderAccessSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IDPPublicCERTs []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                               `json:"sso_target_url"`
	JSON         identityProviderAccessSamlConfigJSON `json:"-"`
}

// identityProviderAccessSamlConfigJSON contains the JSON metadata for the struct
// [IdentityProviderAccessSamlConfig]
type identityProviderAccessSamlConfigJSON struct {
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

func (r *IdentityProviderAccessSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessSamlConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderAccessSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                              `json:"header_name"`
	JSON       identityProviderAccessSamlConfigHeaderAttributeJSON `json:"-"`
}

// identityProviderAccessSamlConfigHeaderAttributeJSON contains the JSON metadata
// for the struct [IdentityProviderAccessSamlConfigHeaderAttribute]
type identityProviderAccessSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityProviderAccessSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessSamlConfigHeaderAttributeJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config GenericOAuthConfig `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderAccessYandexJSON                 `json:"-"`
}

// identityProviderAccessYandexJSON contains the JSON metadata for the struct
// [IdentityProviderAccessYandex]
type identityProviderAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessYandexJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessYandex) implementsZeroTrustIdentityProvider() {}

type IdentityProviderAccessOnetimepin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config interface{} `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderAccessOnetimepinJSON             `json:"-"`
}

// identityProviderAccessOnetimepinJSON contains the JSON metadata for the struct
// [IdentityProviderAccessOnetimepin]
type identityProviderAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	ScimConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessOnetimepinJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessOnetimepin) implementsZeroTrustIdentityProvider() {}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb string

const (
	UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbOnetimepin UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb = "onetimepin"
	UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbAzureAd    UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb = "azureAD"
	UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbSaml       UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb = "saml"
	UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbCentrify   UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb = "centrify"
	UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbFacebook   UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb = "facebook"
	UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbGitHub     UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb = "github"
	UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbGoogleApps UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb = "google-apps"
	UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbGoogle     UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb = "google"
	UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbLinkedin   UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb = "linkedin"
	UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbOidc       UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb = "oidc"
	UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbOkta       UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb = "okta"
	UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbOnelogin   UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb = "onelogin"
	UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbPingone    UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb = "pingone"
	UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbYandex     UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb = "yandex"
)

func (r UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbOnetimepin, UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbAzureAd, UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbSaml, UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbCentrify, UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbFacebook, UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbGitHub, UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbGoogleApps, UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbGoogle, UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbLinkedin, UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbOidc, UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbOkta, UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbOnelogin, UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbPingone, UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcbYandex:
		return true
	}
	return false
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            unnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1JSON `json:"-"`
}

// unnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1JSON contains the JSON metadata
// for the struct [UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1]
type unnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1JSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1JSON) RawJSON() string {
	return r.raw
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param struct {
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

func (r UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderListResponse struct {
	Config interface{} `json:"config"`
	// UUID
	ID string `json:"id"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type  UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	JSON  identityProviderListResponseJSON                 `json:"-"`
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

// Union satisfied by [zero_trust.AzureAd],
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
type IdentityProviderListResponseUnion interface {
	implementsZeroTrustIdentityProviderListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IdentityProviderListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AzureAd{}),
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
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderListResponseAccessCentrifyJSON   `json:"-"`
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
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderListResponseAccessFacebookJSON   `json:"-"`
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
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderListResponseAccessGitHubJSON     `json:"-"`
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
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderListResponseAccessGoogleJSON     `json:"-"`
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
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
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
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderListResponseAccessLinkedinJSON   `json:"-"`
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
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
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
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
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
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderListResponseAccessOneloginJSON   `json:"-"`
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
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderListResponseAccessPingoneJSON    `json:"-"`
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
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
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
	Type UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1 `json:"scim_config"`
	JSON       identityProviderListResponseAccessYandexJSON     `json:"-"`
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

// This interface is a union satisfied by one of the following:
// [IdentityProviderNewParamsAzureAd], [IdentityProviderNewParamsAccessCentrify],
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

type IdentityProviderNewParamsAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderNewParamsAzureAdConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
}

func (r IdentityProviderNewParamsAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderNewParamsAzureAd) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderNewParamsAzureAd) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderNewParamsAzureAd) ImplementsIdentityProviderNewParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderNewParamsAzureAdConfig struct {
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
	Prompt param.Field[IdentityProviderNewParamsAzureAdConfigPrompt] `json:"prompt"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
}

func (r IdentityProviderNewParamsAzureAdConfig) MarshalJSON() (data []byte, err error) {
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
type IdentityProviderNewParamsAzureAdConfigPrompt string

const (
	IdentityProviderNewParamsAzureAdConfigPromptLogin         IdentityProviderNewParamsAzureAdConfigPrompt = "login"
	IdentityProviderNewParamsAzureAdConfigPromptSelectAccount IdentityProviderNewParamsAzureAdConfigPrompt = "select_account"
	IdentityProviderNewParamsAzureAdConfigPromptNone          IdentityProviderNewParamsAzureAdConfigPrompt = "none"
)

func (r IdentityProviderNewParamsAzureAdConfigPrompt) IsKnown() bool {
	switch r {
	case IdentityProviderNewParamsAzureAdConfigPromptLogin, IdentityProviderNewParamsAzureAdConfigPromptSelectAccount, IdentityProviderNewParamsAzureAdConfigPromptNone:
		return true
	}
	return false
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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

type IdentityProviderNewParamsAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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

type IdentityProviderNewParamsAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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

type IdentityProviderNewParamsAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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

type IdentityProviderNewParamsAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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

type IdentityProviderNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   IdentityProvider      `json:"result,required"`
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
// [IdentityProviderUpdateParamsAzureAd],
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

type IdentityProviderUpdateParamsAzureAd struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderUpdateParamsAzureAdConfig] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
}

func (r IdentityProviderUpdateParamsAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderUpdateParamsAzureAd) getAccountID() param.Field[string] {
	return r.AccountID
}

func (r IdentityProviderUpdateParamsAzureAd) getZoneID() param.Field[string] {
	return r.ZoneID
}

func (IdentityProviderUpdateParamsAzureAd) ImplementsIdentityProviderUpdateParams() {

}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderUpdateParamsAzureAdConfig struct {
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
	Prompt param.Field[IdentityProviderUpdateParamsAzureAdConfigPrompt] `json:"prompt"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
}

func (r IdentityProviderUpdateParamsAzureAdConfig) MarshalJSON() (data []byte, err error) {
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
type IdentityProviderUpdateParamsAzureAdConfigPrompt string

const (
	IdentityProviderUpdateParamsAzureAdConfigPromptLogin         IdentityProviderUpdateParamsAzureAdConfigPrompt = "login"
	IdentityProviderUpdateParamsAzureAdConfigPromptSelectAccount IdentityProviderUpdateParamsAzureAdConfigPrompt = "select_account"
	IdentityProviderUpdateParamsAzureAdConfigPromptNone          IdentityProviderUpdateParamsAzureAdConfigPrompt = "none"
)

func (r IdentityProviderUpdateParamsAzureAdConfigPrompt) IsKnown() bool {
	switch r {
	case IdentityProviderUpdateParamsAzureAdConfigPromptLogin, IdentityProviderUpdateParamsAzureAdConfigPromptSelectAccount, IdentityProviderUpdateParamsAzureAdConfigPromptNone:
		return true
	}
	return false
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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

type IdentityProviderUpdateParamsAccessFacebook struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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

type IdentityProviderUpdateParamsAccessGitHub struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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

type IdentityProviderUpdateParamsAccessLinkedin struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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

type IdentityProviderUpdateParamsAccessYandex struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[GenericOAuthConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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
	Type param.Field[UnnamedSchemaRef9ab84e842cdf571c8f3898648bcdabcb] `json:"type,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[UnnamedSchemaRefDd86d8b7ea73283da7b160ed3f86cae1Param] `json:"scim_config"`
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

type IdentityProviderUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   IdentityProvider      `json:"result,required"`
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
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   IdentityProviderDeleteResponse `json:"result,required"`
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
	Result   IdentityProvider      `json:"result,required"`
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

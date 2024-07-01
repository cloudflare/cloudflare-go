// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
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
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityProviderService] method instead.
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
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
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
func (r *IdentityProviderService) Update(ctx context.Context, identityProviderID string, params IdentityProviderUpdateParams, opts ...option.RequestOption) (res *IdentityProvider, err error) {
	opts = append(r.Options[:], opts...)
	var env IdentityProviderUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/identity_providers/%s", accountOrZone, accountOrZoneID, identityProviderID)
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
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
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
func (r *IdentityProviderService) Delete(ctx context.Context, identityProviderID string, body IdentityProviderDeleteParams, opts ...option.RequestOption) (res *IdentityProviderDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IdentityProviderDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/identity_providers/%s", accountOrZone, accountOrZoneID, identityProviderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a configured identity provider.
func (r *IdentityProviderService) Get(ctx context.Context, identityProviderID string, query IdentityProviderGetParams, opts ...option.RequestOption) (res *IdentityProvider, err error) {
	opts = append(r.Options[:], opts...)
	var env IdentityProviderGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	if identityProviderID == "" {
		err = errors.New("missing required identity_provider_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/identity_providers/%s", accountOrZone, accountOrZoneID, identityProviderID)
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
	SCIMConfig SCIMConfig  `json:"scim_config"`
	JSON       azureADJSON `json:"-"`
}

// azureADJSON contains the JSON metadata for the struct [AzureAD]
type azureADJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AzureAD) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r azureADJSON) RawJSON() string {
	return r.raw
}

func (r AzureAD) implementsZeroTrustIdentityProvider() {}

func (r AzureAD) implementsZeroTrustIdentityProviderListResponse() {}

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
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
}

func (r AzureADParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AzureADParam) implementsZeroTrustIdentityProviderUnionParam() {}

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

type IdentityProvider struct {
	// This field can have the runtime type of [AzureADConfig],
	// [IdentityProviderAccessCentrifyConfig], [GenericOAuthConfig],
	// [IdentityProviderAccessGoogleConfig], [IdentityProviderAccessGoogleAppsConfig],
	// [IdentityProviderAccessOIDCConfig], [IdentityProviderAccessOktaConfig],
	// [IdentityProviderAccessOneloginConfig], [IdentityProviderAccessPingoneConfig],
	// [IdentityProviderAccessSAMLConfig], [interface{}].
	Config interface{} `json:"config"`
	// UUID
	ID string `json:"id"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig SCIMConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type  IdentityProviderType `json:"type,required"`
	JSON  identityProviderJSON `json:"-"`
	union IdentityProviderUnion
}

// identityProviderJSON contains the JSON metadata for the struct
// [IdentityProvider]
type identityProviderJSON struct {
	Config      apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	SCIMConfig  apijson.Field
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

// AsUnion returns a [IdentityProviderUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [zero_trust.AzureAD],
// [zero_trust.IdentityProviderAccessCentrify],
// [zero_trust.IdentityProviderAccessFacebook],
// [zero_trust.IdentityProviderAccessGitHub],
// [zero_trust.IdentityProviderAccessGoogle],
// [zero_trust.IdentityProviderAccessGoogleApps],
// [zero_trust.IdentityProviderAccessLinkedin],
// [zero_trust.IdentityProviderAccessOIDC],
// [zero_trust.IdentityProviderAccessOkta],
// [zero_trust.IdentityProviderAccessOnelogin],
// [zero_trust.IdentityProviderAccessPingone],
// [zero_trust.IdentityProviderAccessSAML],
// [zero_trust.IdentityProviderAccessYandex],
// [zero_trust.IdentityProviderAccessOnetimepin].
func (r IdentityProvider) AsUnion() IdentityProviderUnion {
	return r.union
}

// Union satisfied by [zero_trust.AzureAD],
// [zero_trust.IdentityProviderAccessCentrify],
// [zero_trust.IdentityProviderAccessFacebook],
// [zero_trust.IdentityProviderAccessGitHub],
// [zero_trust.IdentityProviderAccessGoogle],
// [zero_trust.IdentityProviderAccessGoogleApps],
// [zero_trust.IdentityProviderAccessLinkedin],
// [zero_trust.IdentityProviderAccessOIDC],
// [zero_trust.IdentityProviderAccessOkta],
// [zero_trust.IdentityProviderAccessOnelogin],
// [zero_trust.IdentityProviderAccessPingone],
// [zero_trust.IdentityProviderAccessSAML],
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
			Type:       reflect.TypeOf(AzureAD{}),
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
			Type:       reflect.TypeOf(IdentityProviderAccessOIDC{}),
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
			Type:       reflect.TypeOf(IdentityProviderAccessSAML{}),
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig SCIMConfig                         `json:"scim_config"`
	JSON       identityProviderAccessCentrifyJSON `json:"-"`
}

// identityProviderAccessCentrifyJSON contains the JSON metadata for the struct
// [IdentityProviderAccessCentrify]
type identityProviderAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig SCIMConfig                         `json:"scim_config"`
	JSON       identityProviderAccessFacebookJSON `json:"-"`
}

// identityProviderAccessFacebookJSON contains the JSON metadata for the struct
// [IdentityProviderAccessFacebook]
type identityProviderAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig SCIMConfig                       `json:"scim_config"`
	JSON       identityProviderAccessGitHubJSON `json:"-"`
}

// identityProviderAccessGitHubJSON contains the JSON metadata for the struct
// [IdentityProviderAccessGitHub]
type identityProviderAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig SCIMConfig                       `json:"scim_config"`
	JSON       identityProviderAccessGoogleJSON `json:"-"`
}

// identityProviderAccessGoogleJSON contains the JSON metadata for the struct
// [IdentityProviderAccessGoogle]
type identityProviderAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig SCIMConfig                           `json:"scim_config"`
	JSON       identityProviderAccessGoogleAppsJSON `json:"-"`
}

// identityProviderAccessGoogleAppsJSON contains the JSON metadata for the struct
// [IdentityProviderAccessGoogleApps]
type identityProviderAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig SCIMConfig                         `json:"scim_config"`
	JSON       identityProviderAccessLinkedinJSON `json:"-"`
}

// identityProviderAccessLinkedinJSON contains the JSON metadata for the struct
// [IdentityProviderAccessLinkedin]
type identityProviderAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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

type IdentityProviderAccessOIDC struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderAccessOIDCConfig `json:"config,required"`
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
	SCIMConfig SCIMConfig                     `json:"scim_config"`
	JSON       identityProviderAccessOIDCJSON `json:"-"`
}

// identityProviderAccessOIDCJSON contains the JSON metadata for the struct
// [IdentityProviderAccessOIDC]
type identityProviderAccessOIDCJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessOIDC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessOIDCJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessOIDC) implementsZeroTrustIdentityProvider() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessOIDCConfig struct {
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
	JSON     identityProviderAccessOIDCConfigJSON `json:"-"`
}

// identityProviderAccessOIDCConfigJSON contains the JSON metadata for the struct
// [IdentityProviderAccessOIDCConfig]
type identityProviderAccessOIDCConfigJSON struct {
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

func (r *IdentityProviderAccessOIDCConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessOIDCConfigJSON) RawJSON() string {
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig SCIMConfig                     `json:"scim_config"`
	JSON       identityProviderAccessOktaJSON `json:"-"`
}

// identityProviderAccessOktaJSON contains the JSON metadata for the struct
// [IdentityProviderAccessOkta]
type identityProviderAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig SCIMConfig                         `json:"scim_config"`
	JSON       identityProviderAccessOneloginJSON `json:"-"`
}

// identityProviderAccessOneloginJSON contains the JSON metadata for the struct
// [IdentityProviderAccessOnelogin]
type identityProviderAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig SCIMConfig                        `json:"scim_config"`
	JSON       identityProviderAccessPingoneJSON `json:"-"`
}

// identityProviderAccessPingoneJSON contains the JSON metadata for the struct
// [IdentityProviderAccessPingone]
type identityProviderAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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

type IdentityProviderAccessSAML struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config IdentityProviderAccessSAMLConfig `json:"config,required"`
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
	SCIMConfig SCIMConfig                     `json:"scim_config"`
	JSON       identityProviderAccessSAMLJSON `json:"-"`
}

// identityProviderAccessSAMLJSON contains the JSON metadata for the struct
// [IdentityProviderAccessSAML]
type identityProviderAccessSAMLJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProviderAccessSAML) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessSAMLJSON) RawJSON() string {
	return r.raw
}

func (r IdentityProviderAccessSAML) implementsZeroTrustIdentityProvider() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessSAMLConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []IdentityProviderAccessSAMLConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdPPublicCERTs []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                               `json:"sso_target_url"`
	JSON         identityProviderAccessSAMLConfigJSON `json:"-"`
}

// identityProviderAccessSAMLConfigJSON contains the JSON metadata for the struct
// [IdentityProviderAccessSAMLConfig]
type identityProviderAccessSAMLConfigJSON struct {
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

func (r *IdentityProviderAccessSAMLConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessSAMLConfigJSON) RawJSON() string {
	return r.raw
}

type IdentityProviderAccessSAMLConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                              `json:"header_name"`
	JSON       identityProviderAccessSAMLConfigHeaderAttributeJSON `json:"-"`
}

// identityProviderAccessSAMLConfigHeaderAttributeJSON contains the JSON metadata
// for the struct [IdentityProviderAccessSAMLConfigHeaderAttribute]
type identityProviderAccessSAMLConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityProviderAccessSAMLConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityProviderAccessSAMLConfigHeaderAttributeJSON) RawJSON() string {
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig SCIMConfig                       `json:"scim_config"`
	JSON       identityProviderAccessYandexJSON `json:"-"`
}

// identityProviderAccessYandexJSON contains the JSON metadata for the struct
// [IdentityProviderAccessYandex]
type identityProviderAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	Type IdentityProviderType `json:"type,required"`
	// UUID
	ID string `json:"id"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig SCIMConfig                           `json:"scim_config"`
	JSON       identityProviderAccessOnetimepinJSON `json:"-"`
}

// identityProviderAccessOnetimepinJSON contains the JSON metadata for the struct
// [IdentityProviderAccessOnetimepin]
type identityProviderAccessOnetimepinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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

type IdentityProviderParam struct {
	Config param.Field[interface{}] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
}

func (r IdentityProviderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderParam) implementsZeroTrustIdentityProviderUnionParam() {}

// Satisfied by [zero_trust.AzureADParam],
// [zero_trust.IdentityProviderAccessCentrifyParam],
// [zero_trust.IdentityProviderAccessFacebookParam],
// [zero_trust.IdentityProviderAccessGitHubParam],
// [zero_trust.IdentityProviderAccessGoogleParam],
// [zero_trust.IdentityProviderAccessGoogleAppsParam],
// [zero_trust.IdentityProviderAccessLinkedinParam],
// [zero_trust.IdentityProviderAccessOIDCParam],
// [zero_trust.IdentityProviderAccessOktaParam],
// [zero_trust.IdentityProviderAccessOneloginParam],
// [zero_trust.IdentityProviderAccessPingoneParam],
// [zero_trust.IdentityProviderAccessSAMLParam],
// [zero_trust.IdentityProviderAccessYandexParam],
// [zero_trust.IdentityProviderAccessOnetimepinParam], [IdentityProviderParam].
type IdentityProviderUnionParam interface {
	implementsZeroTrustIdentityProviderUnionParam()
}

type IdentityProviderAccessCentrifyParam struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderAccessCentrifyConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
}

func (r IdentityProviderAccessCentrifyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderAccessCentrifyParam) implementsZeroTrustIdentityProviderUnionParam() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessCentrifyConfigParam struct {
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

func (r IdentityProviderAccessCentrifyConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderAccessFacebookParam struct {
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
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
}

func (r IdentityProviderAccessFacebookParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderAccessFacebookParam) implementsZeroTrustIdentityProviderUnionParam() {}

type IdentityProviderAccessGitHubParam struct {
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
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
}

func (r IdentityProviderAccessGitHubParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderAccessGitHubParam) implementsZeroTrustIdentityProviderUnionParam() {}

type IdentityProviderAccessGoogleParam struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderAccessGoogleConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
}

func (r IdentityProviderAccessGoogleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderAccessGoogleParam) implementsZeroTrustIdentityProviderUnionParam() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessGoogleConfigParam struct {
	// Custom claims
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// The claim name for email in the id_token response.
	EmailClaimName param.Field[string] `json:"email_claim_name"`
}

func (r IdentityProviderAccessGoogleConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderAccessGoogleAppsParam struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderAccessGoogleAppsConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
}

func (r IdentityProviderAccessGoogleAppsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderAccessGoogleAppsParam) implementsZeroTrustIdentityProviderUnionParam() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessGoogleAppsConfigParam struct {
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

func (r IdentityProviderAccessGoogleAppsConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderAccessLinkedinParam struct {
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
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
}

func (r IdentityProviderAccessLinkedinParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderAccessLinkedinParam) implementsZeroTrustIdentityProviderUnionParam() {}

type IdentityProviderAccessOIDCParam struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderAccessOIDCConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
}

func (r IdentityProviderAccessOIDCParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderAccessOIDCParam) implementsZeroTrustIdentityProviderUnionParam() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessOIDCConfigParam struct {
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

func (r IdentityProviderAccessOIDCConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderAccessOktaParam struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderAccessOktaConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
}

func (r IdentityProviderAccessOktaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderAccessOktaParam) implementsZeroTrustIdentityProviderUnionParam() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessOktaConfigParam struct {
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

func (r IdentityProviderAccessOktaConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderAccessOneloginParam struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderAccessOneloginConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
}

func (r IdentityProviderAccessOneloginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderAccessOneloginParam) implementsZeroTrustIdentityProviderUnionParam() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessOneloginConfigParam struct {
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

func (r IdentityProviderAccessOneloginConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderAccessPingoneParam struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderAccessPingoneConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
}

func (r IdentityProviderAccessPingoneParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderAccessPingoneParam) implementsZeroTrustIdentityProviderUnionParam() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessPingoneConfigParam struct {
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

func (r IdentityProviderAccessPingoneConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderAccessSAMLParam struct {
	// The configuration parameters for the identity provider. To view the required
	// parameters for a specific provider, refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Config param.Field[IdentityProviderAccessSAMLConfigParam] `json:"config,required"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name,required"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[IdentityProviderType] `json:"type,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
}

func (r IdentityProviderAccessSAMLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderAccessSAMLParam) implementsZeroTrustIdentityProviderUnionParam() {}

// The configuration parameters for the identity provider. To view the required
// parameters for a specific provider, refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type IdentityProviderAccessSAMLConfigParam struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]IdentityProviderAccessSAMLConfigHeaderAttributeParam] `json:"header_attributes"`
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

func (r IdentityProviderAccessSAMLConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderAccessSAMLConfigHeaderAttributeParam struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r IdentityProviderAccessSAMLConfigHeaderAttributeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderAccessYandexParam struct {
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
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
}

func (r IdentityProviderAccessYandexParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderAccessYandexParam) implementsZeroTrustIdentityProviderUnionParam() {}

type IdentityProviderAccessOnetimepinParam struct {
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
	SCIMConfig param.Field[SCIMConfigParam] `json:"scim_config"`
}

func (r IdentityProviderAccessOnetimepinParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IdentityProviderAccessOnetimepinParam) implementsZeroTrustIdentityProviderUnionParam() {}

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
type SCIMConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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

// scimConfigJSON contains the JSON metadata for the struct [SCIMConfig]
type scimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SCIMConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scimConfigJSON) RawJSON() string {
	return r.raw
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type SCIMConfigParam struct {
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

func (r SCIMConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IdentityProviderListResponse struct {
	// This field can have the runtime type of [AzureADConfig],
	// [IdentityProviderListResponseAccessCentrifyConfig], [GenericOAuthConfig],
	// [IdentityProviderListResponseAccessGoogleConfig],
	// [IdentityProviderListResponseAccessGoogleAppsConfig],
	// [IdentityProviderListResponseAccessOIDCConfig],
	// [IdentityProviderListResponseAccessOktaConfig],
	// [IdentityProviderListResponseAccessOneloginConfig],
	// [IdentityProviderListResponseAccessPingoneConfig],
	// [IdentityProviderListResponseAccessSAMLConfig].
	Config interface{} `json:"config"`
	// UUID
	ID string `json:"id"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name,required"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	SCIMConfig SCIMConfig `json:"scim_config"`
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
	SCIMConfig  apijson.Field
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

// AsUnion returns a [IdentityProviderListResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [zero_trust.AzureAD],
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
// [zero_trust.IdentityProviderListResponseAccessSAML],
// [zero_trust.IdentityProviderListResponseAccessYandex].
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
	SCIMConfig SCIMConfig                                     `json:"scim_config"`
	JSON       identityProviderListResponseAccessCentrifyJSON `json:"-"`
}

// identityProviderListResponseAccessCentrifyJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessCentrify]
type identityProviderListResponseAccessCentrifyJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	SCIMConfig SCIMConfig                                     `json:"scim_config"`
	JSON       identityProviderListResponseAccessFacebookJSON `json:"-"`
}

// identityProviderListResponseAccessFacebookJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessFacebook]
type identityProviderListResponseAccessFacebookJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	SCIMConfig SCIMConfig                                   `json:"scim_config"`
	JSON       identityProviderListResponseAccessGitHubJSON `json:"-"`
}

// identityProviderListResponseAccessGitHubJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessGitHub]
type identityProviderListResponseAccessGitHubJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	SCIMConfig SCIMConfig                                   `json:"scim_config"`
	JSON       identityProviderListResponseAccessGoogleJSON `json:"-"`
}

// identityProviderListResponseAccessGoogleJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessGoogle]
type identityProviderListResponseAccessGoogleJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	SCIMConfig SCIMConfig                                       `json:"scim_config"`
	JSON       identityProviderListResponseAccessGoogleAppsJSON `json:"-"`
}

// identityProviderListResponseAccessGoogleAppsJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessGoogleApps]
type identityProviderListResponseAccessGoogleAppsJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	SCIMConfig SCIMConfig                                     `json:"scim_config"`
	JSON       identityProviderListResponseAccessLinkedinJSON `json:"-"`
}

// identityProviderListResponseAccessLinkedinJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessLinkedin]
type identityProviderListResponseAccessLinkedinJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	SCIMConfig SCIMConfig                                 `json:"scim_config"`
	JSON       identityProviderListResponseAccessOIDCJSON `json:"-"`
}

// identityProviderListResponseAccessOIDCJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessOIDC]
type identityProviderListResponseAccessOIDCJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	SCIMConfig SCIMConfig                                 `json:"scim_config"`
	JSON       identityProviderListResponseAccessOktaJSON `json:"-"`
}

// identityProviderListResponseAccessOktaJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessOkta]
type identityProviderListResponseAccessOktaJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	SCIMConfig SCIMConfig                                     `json:"scim_config"`
	JSON       identityProviderListResponseAccessOneloginJSON `json:"-"`
}

// identityProviderListResponseAccessOneloginJSON contains the JSON metadata for
// the struct [IdentityProviderListResponseAccessOnelogin]
type identityProviderListResponseAccessOneloginJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	SCIMConfig SCIMConfig                                    `json:"scim_config"`
	JSON       identityProviderListResponseAccessPingoneJSON `json:"-"`
}

// identityProviderListResponseAccessPingoneJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessPingone]
type identityProviderListResponseAccessPingoneJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	SCIMConfig SCIMConfig                                 `json:"scim_config"`
	JSON       identityProviderListResponseAccessSAMLJSON `json:"-"`
}

// identityProviderListResponseAccessSAMLJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessSAML]
type identityProviderListResponseAccessSAMLJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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
	SCIMConfig SCIMConfig                                   `json:"scim_config"`
	JSON       identityProviderListResponseAccessYandexJSON `json:"-"`
}

// identityProviderListResponseAccessYandexJSON contains the JSON metadata for the
// struct [IdentityProviderListResponseAccessYandex]
type identityProviderListResponseAccessYandexJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	SCIMConfig  apijson.Field
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

type IdentityProviderNewParams struct {
	IdentityProvider IdentityProviderUnionParam `json:"identity_provider,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r IdentityProviderNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.IdentityProvider)
}

type IdentityProviderNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IdentityProviderNewResponseEnvelopeSuccess `json:"success,required"`
	Result  IdentityProvider                           `json:"result"`
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
	IdentityProvider IdentityProviderUnionParam `json:"identity_provider,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

func (r IdentityProviderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.IdentityProvider)
}

type IdentityProviderUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success IdentityProviderUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  IdentityProvider                              `json:"result"`
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
	Result  IdentityProvider                           `json:"result"`
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

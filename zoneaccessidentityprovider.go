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

// ZoneAccessIdentityProviderService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneAccessIdentityProviderService] method instead.
type ZoneAccessIdentityProviderService struct {
	Options []option.RequestOption
}

// NewZoneAccessIdentityProviderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneAccessIdentityProviderService(opts ...option.RequestOption) (r *ZoneAccessIdentityProviderService) {
	r = &ZoneAccessIdentityProviderService{}
	r.Options = opts
	return
}

// Fetches a configured identity provider.
func (r *ZoneAccessIdentityProviderService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *IdentityProvidersSingleResponse8KOtOl1H, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/identity_providers/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured identity provider.
func (r *ZoneAccessIdentityProviderService) Update(ctx context.Context, identifier interface{}, uuid string, body ZoneAccessIdentityProviderUpdateParams, opts ...option.RequestOption) (res *IdentityProvidersSingleResponse8KOtOl1H, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/identity_providers/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all configured identity providers.
func (r *ZoneAccessIdentityProviderService) List(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *IdentityProvidersResponseCollection4ARq9pOu, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/identity_providers", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an identity provider from Access.
func (r *ZoneAccessIdentityProviderService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *APIResponseSingleIDY7yIk9Qy, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/identity_providers/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new identity provider to Access.
func (r *ZoneAccessIdentityProviderService) Add(ctx context.Context, identifier interface{}, body ZoneAccessIdentityProviderAddParams, opts ...option.RequestOption) (res *IdentityProvidersSingleResponse8KOtOl1H, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/identity_providers", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type IdentityProvidersResponseCollection4ARq9pOu struct {
	Errors     []IdentityProvidersResponseCollection4ARq9pOuError    `json:"errors"`
	Messages   []IdentityProvidersResponseCollection4ARq9pOuMessage  `json:"messages"`
	Result     []IdentityProvidersResponseCollection4ARq9pOuResult   `json:"result"`
	ResultInfo IdentityProvidersResponseCollection4ARq9pOuResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success IdentityProvidersResponseCollection4ARq9pOuSuccess `json:"success"`
	JSON    identityProvidersResponseCollection4ARq9pOuJSON    `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuJSON contains the JSON metadata for
// the struct [IdentityProvidersResponseCollection4ARq9pOu]
type identityProvidersResponseCollection4ARq9pOuJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOu) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    identityProvidersResponseCollection4ARq9pOuErrorJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuErrorJSON contains the JSON metadata
// for the struct [IdentityProvidersResponseCollection4ARq9pOuError]
type identityProvidersResponseCollection4ARq9pOuErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    identityProvidersResponseCollection4ARq9pOuMessageJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuMessageJSON contains the JSON
// metadata for the struct [IdentityProvidersResponseCollection4ARq9pOuMessage]
type identityProvidersResponseCollection4ARq9pOuMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasAzureAd],
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasCentrify],
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasFacebook],
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasGitHub],
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogle],
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleApps],
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasLinkedin],
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasOidc],
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasOkta],
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasOnelogin],
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasPingone],
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasSaml] or
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasYandex].
type IdentityProvidersResponseCollection4ARq9pOuResult interface {
	implementsIdentityProvidersResponseCollection4ARq9pOuResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*IdentityProvidersResponseCollection4ARq9pOuResult)(nil)).Elem(), "")
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasAzureAd struct {
	// UUID
	ID     string                                                                `json:"id"`
	Config IdentityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                              `json:"type"`
	JSON identityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdJSON contains the
// JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasAzureAd]
type identityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersResponseCollection4ARq9pOuResultSchemasAzureAd) implementsIdentityProvidersResponseCollection4ARq9pOuResult() {
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your Azure directory uuid
	DirectoryID string `json:"directory_id"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool                                                                      `json:"support_groups"`
	JSON          identityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdConfigJSON struct {
	ClientID      apijson.Field
	ClientSecret  apijson.Field
	DirectoryID   apijson.Field
	SupportGroups apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                          `json:"user_deprovision"`
	JSON            identityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdScimConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdScimConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasCentrify struct {
	// UUID
	ID     string                                                                 `json:"id"`
	Config IdentityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                               `json:"type"`
	JSON identityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyJSON contains
// the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasCentrify]
type identityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersResponseCollection4ARq9pOuResultSchemasCentrify) implementsIdentityProvidersResponseCollection4ARq9pOuResult() {
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount string `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID string `json:"centrify_app_id"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                     `json:"client_secret"`
	JSON         identityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                           `json:"user_deprovision"`
	JSON            identityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyScimConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyScimConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasFacebook struct {
	// UUID
	ID     string                                                                 `json:"id"`
	Config IdentityProvidersResponseCollection4ARq9pOuResultSchemasFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersResponseCollection4ARq9pOuResultSchemasFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                               `json:"type"`
	JSON identityProvidersResponseCollection4ARq9pOuResultSchemasFacebookJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasFacebookJSON contains
// the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasFacebook]
type identityProvidersResponseCollection4ARq9pOuResultSchemasFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersResponseCollection4ARq9pOuResultSchemasFacebook) implementsIdentityProvidersResponseCollection4ARq9pOuResult() {
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                     `json:"client_secret"`
	JSON         identityProvidersResponseCollection4ARq9pOuResultSchemasFacebookConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasFacebookConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasFacebookConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersResponseCollection4ARq9pOuResultSchemasFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                           `json:"user_deprovision"`
	JSON            identityProvidersResponseCollection4ARq9pOuResultSchemasFacebookScimConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasFacebookScimConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasGitHub struct {
	// UUID
	ID     string                                                               `json:"id"`
	Config IdentityProvidersResponseCollection4ARq9pOuResultSchemasGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersResponseCollection4ARq9pOuResultSchemasGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                             `json:"type"`
	JSON identityProvidersResponseCollection4ARq9pOuResultSchemasGitHubJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasGitHubJSON contains the
// JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasGitHub]
type identityProvidersResponseCollection4ARq9pOuResultSchemasGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersResponseCollection4ARq9pOuResultSchemasGitHub) implementsIdentityProvidersResponseCollection4ARq9pOuResult() {
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                   `json:"client_secret"`
	JSON         identityProvidersResponseCollection4ARq9pOuResultSchemasGitHubConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasGitHubConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasGitHubConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersResponseCollection4ARq9pOuResultSchemasGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                         `json:"user_deprovision"`
	JSON            identityProvidersResponseCollection4ARq9pOuResultSchemasGitHubScimConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasGitHubScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasGitHubScimConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogle struct {
	// UUID
	ID     string                                                               `json:"id"`
	Config IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                             `json:"type"`
	JSON identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleJSON contains the
// JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogle]
type identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogle) implementsIdentityProvidersResponseCollection4ARq9pOuResult() {
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                   `json:"client_secret"`
	JSON         identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                         `json:"user_deprovision"`
	JSON            identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleScimConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleScimConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleApps struct {
	// UUID
	ID     string                                                                   `json:"id"`
	Config IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                                 `json:"type"`
	JSON identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsJSON contains
// the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleApps]
type identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleApps) implementsIdentityProvidersResponseCollection4ARq9pOuResult() {
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                       `json:"client_secret"`
	JSON         identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsConfigJSON struct {
	AppsDomain   apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                             `json:"user_deprovision"`
	JSON            identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsScimConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsScimConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasLinkedin struct {
	// UUID
	ID     string                                                                 `json:"id"`
	Config IdentityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                               `json:"type"`
	JSON identityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinJSON contains
// the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasLinkedin]
type identityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersResponseCollection4ARq9pOuResultSchemasLinkedin) implementsIdentityProvidersResponseCollection4ARq9pOuResult() {
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                     `json:"client_secret"`
	JSON         identityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                           `json:"user_deprovision"`
	JSON            identityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinScimConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinScimConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasOidc struct {
	// UUID
	ID     string                                                             `json:"id"`
	Config IdentityProvidersResponseCollection4ARq9pOuResultSchemasOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersResponseCollection4ARq9pOuResultSchemasOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                           `json:"type"`
	JSON identityProvidersResponseCollection4ARq9pOuResultSchemasOidcJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasOidcJSON contains the
// JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasOidc]
type identityProvidersResponseCollection4ARq9pOuResultSchemasOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersResponseCollection4ARq9pOuResultSchemasOidc) implementsIdentityProvidersResponseCollection4ARq9pOuResult() {
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL string `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL string `json:"certs_url"`
	// List of custom claims that will be pulled from your id_token and added to your
	// signed Access JWT token.
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// OAuth scopes
	Scopes []string `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL string                                                                 `json:"token_url"`
	JSON     identityProvidersResponseCollection4ARq9pOuResultSchemasOidcConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasOidcConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasOidcConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasOidcConfigJSON struct {
	AuthURL      apijson.Field
	CertsURL     apijson.Field
	Claims       apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	Scopes       apijson.Field
	TokenURL     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersResponseCollection4ARq9pOuResultSchemasOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                       `json:"user_deprovision"`
	JSON            identityProvidersResponseCollection4ARq9pOuResultSchemasOidcScimConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasOidcScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasOidcScimConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasOkta struct {
	// UUID
	ID     string                                                             `json:"id"`
	Config IdentityProvidersResponseCollection4ARq9pOuResultSchemasOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersResponseCollection4ARq9pOuResultSchemasOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                           `json:"type"`
	JSON identityProvidersResponseCollection4ARq9pOuResultSchemasOktaJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasOktaJSON contains the
// JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasOkta]
type identityProvidersResponseCollection4ARq9pOuResultSchemasOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersResponseCollection4ARq9pOuResultSchemasOkta) implementsIdentityProvidersResponseCollection4ARq9pOuResult() {
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasOktaConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your okta account url
	OktaAccount string                                                                 `json:"okta_account"`
	JSON        identityProvidersResponseCollection4ARq9pOuResultSchemasOktaConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasOktaConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasOktaConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasOktaConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	OktaAccount  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersResponseCollection4ARq9pOuResultSchemasOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                       `json:"user_deprovision"`
	JSON            identityProvidersResponseCollection4ARq9pOuResultSchemasOktaScimConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasOktaScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasOktaScimConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasOnelogin struct {
	// UUID
	ID     string                                                                 `json:"id"`
	Config IdentityProvidersResponseCollection4ARq9pOuResultSchemasOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersResponseCollection4ARq9pOuResultSchemasOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                               `json:"type"`
	JSON identityProvidersResponseCollection4ARq9pOuResultSchemasOneloginJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasOneloginJSON contains
// the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasOnelogin]
type identityProvidersResponseCollection4ARq9pOuResultSchemasOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersResponseCollection4ARq9pOuResultSchemasOnelogin) implementsIdentityProvidersResponseCollection4ARq9pOuResult() {
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasOneloginConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your OneLogin account url
	OneloginAccount string                                                                     `json:"onelogin_account"`
	JSON            identityProvidersResponseCollection4ARq9pOuResultSchemasOneloginConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasOneloginConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasOneloginConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasOneloginConfigJSON struct {
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersResponseCollection4ARq9pOuResultSchemasOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                           `json:"user_deprovision"`
	JSON            identityProvidersResponseCollection4ARq9pOuResultSchemasOneloginScimConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasOneloginScimConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasPingone struct {
	// UUID
	ID     string                                                                `json:"id"`
	Config IdentityProvidersResponseCollection4ARq9pOuResultSchemasPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersResponseCollection4ARq9pOuResultSchemasPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                              `json:"type"`
	JSON identityProvidersResponseCollection4ARq9pOuResultSchemasPingoneJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasPingoneJSON contains the
// JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasPingone]
type identityProvidersResponseCollection4ARq9pOuResultSchemasPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersResponseCollection4ARq9pOuResultSchemasPingone) implementsIdentityProvidersResponseCollection4ARq9pOuResult() {
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasPingoneConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your PingOne environment identifier
	PingEnvID string                                                                    `json:"ping_env_id"`
	JSON      identityProvidersResponseCollection4ARq9pOuResultSchemasPingoneConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasPingoneConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasPingoneConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasPingoneConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	PingEnvID    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersResponseCollection4ARq9pOuResultSchemasPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                          `json:"user_deprovision"`
	JSON            identityProvidersResponseCollection4ARq9pOuResultSchemasPingoneScimConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasPingoneScimConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasSaml struct {
	// UUID
	ID     string                                                             `json:"id"`
	Config IdentityProvidersResponseCollection4ARq9pOuResultSchemasSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersResponseCollection4ARq9pOuResultSchemasSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                           `json:"type"`
	JSON identityProvidersResponseCollection4ARq9pOuResultSchemasSamlJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasSamlJSON contains the
// JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasSaml]
type identityProvidersResponseCollection4ARq9pOuResultSchemasSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersResponseCollection4ARq9pOuResultSchemasSaml) implementsIdentityProvidersResponseCollection4ARq9pOuResult() {
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []IdentityProvidersResponseCollection4ARq9pOuResultSchemasSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                                 `json:"sso_target_url"`
	JSON         identityProvidersResponseCollection4ARq9pOuResultSchemasSamlConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasSamlConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasSamlConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasSamlConfigJSON struct {
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

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                                `json:"header_name"`
	JSON       identityProvidersResponseCollection4ARq9pOuResultSchemasSamlConfigHeaderAttributeJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasSamlConfigHeaderAttribute]
type identityProvidersResponseCollection4ARq9pOuResultSchemasSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersResponseCollection4ARq9pOuResultSchemasSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                       `json:"user_deprovision"`
	JSON            identityProvidersResponseCollection4ARq9pOuResultSchemasSamlScimConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasSamlScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasSamlScimConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasYandex struct {
	// UUID
	ID     string                                                               `json:"id"`
	Config IdentityProvidersResponseCollection4ARq9pOuResultSchemasYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersResponseCollection4ARq9pOuResultSchemasYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                             `json:"type"`
	JSON identityProvidersResponseCollection4ARq9pOuResultSchemasYandexJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasYandexJSON contains the
// JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasYandex]
type identityProvidersResponseCollection4ARq9pOuResultSchemasYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersResponseCollection4ARq9pOuResultSchemasYandex) implementsIdentityProvidersResponseCollection4ARq9pOuResult() {
}

type IdentityProvidersResponseCollection4ARq9pOuResultSchemasYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                   `json:"client_secret"`
	JSON         identityProvidersResponseCollection4ARq9pOuResultSchemasYandexConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasYandexConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasYandexConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersResponseCollection4ARq9pOuResultSchemasYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                         `json:"user_deprovision"`
	JSON            identityProvidersResponseCollection4ARq9pOuResultSchemasYandexScimConfigJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultSchemasYandexScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersResponseCollection4ARq9pOuResultSchemasYandexScimConfig]
type identityProvidersResponseCollection4ARq9pOuResultSchemasYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultSchemasYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersResponseCollection4ARq9pOuResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       identityProvidersResponseCollection4ARq9pOuResultInfoJSON `json:"-"`
}

// identityProvidersResponseCollection4ARq9pOuResultInfoJSON contains the JSON
// metadata for the struct [IdentityProvidersResponseCollection4ARq9pOuResultInfo]
type identityProvidersResponseCollection4ARq9pOuResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersResponseCollection4ARq9pOuResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IdentityProvidersResponseCollection4ARq9pOuSuccess bool

const (
	IdentityProvidersResponseCollection4ARq9pOuSuccessTrue IdentityProvidersResponseCollection4ARq9pOuSuccess = true
)

type IdentityProvidersSingleResponse8KOtOl1H struct {
	Errors   []IdentityProvidersSingleResponse8KOtOl1HError   `json:"errors"`
	Messages []IdentityProvidersSingleResponse8KOtOl1HMessage `json:"messages"`
	Result   IdentityProvidersSingleResponse8KOtOl1HResult    `json:"result"`
	// Whether the API call was successful
	Success IdentityProvidersSingleResponse8KOtOl1HSuccess `json:"success"`
	JSON    identityProvidersSingleResponse8KOtOl1HJSON    `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HJSON contains the JSON metadata for the
// struct [IdentityProvidersSingleResponse8KOtOl1H]
type identityProvidersSingleResponse8KOtOl1HJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1H) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    identityProvidersSingleResponse8KOtOl1HErrorJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HErrorJSON contains the JSON metadata for
// the struct [IdentityProvidersSingleResponse8KOtOl1HError]
type identityProvidersSingleResponse8KOtOl1HErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    identityProvidersSingleResponse8KOtOl1HMessageJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HMessageJSON contains the JSON metadata
// for the struct [IdentityProvidersSingleResponse8KOtOl1HMessage]
type identityProvidersSingleResponse8KOtOl1HMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasAzureAd],
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasCentrify],
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasFacebook],
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasGitHub],
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogle],
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleApps],
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasLinkedin],
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasOidc],
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasOkta],
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasOnelogin],
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasPingone],
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasSaml] or
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasYandex].
type IdentityProvidersSingleResponse8KOtOl1HResult interface {
	implementsIdentityProvidersSingleResponse8KOtOl1HResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*IdentityProvidersSingleResponse8KOtOl1HResult)(nil)).Elem(), "")
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasAzureAd struct {
	// UUID
	ID     string                                                            `json:"id"`
	Config IdentityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                          `json:"type"`
	JSON identityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdJSON contains the
// JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasAzureAd]
type identityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersSingleResponse8KOtOl1HResultSchemasAzureAd) implementsIdentityProvidersSingleResponse8KOtOl1HResult() {
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your Azure directory uuid
	DirectoryID string `json:"directory_id"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool                                                                  `json:"support_groups"`
	JSON          identityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdConfigJSON struct {
	ClientID      apijson.Field
	ClientSecret  apijson.Field
	DirectoryID   apijson.Field
	SupportGroups apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdScimConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdScimConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasCentrify struct {
	// UUID
	ID     string                                                             `json:"id"`
	Config IdentityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                           `json:"type"`
	JSON identityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyJSON contains the
// JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasCentrify]
type identityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersSingleResponse8KOtOl1HResultSchemasCentrify) implementsIdentityProvidersSingleResponse8KOtOl1HResult() {
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount string `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID string `json:"centrify_app_id"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                 `json:"client_secret"`
	JSON         identityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                       `json:"user_deprovision"`
	JSON            identityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyScimConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyScimConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasFacebook struct {
	// UUID
	ID     string                                                             `json:"id"`
	Config IdentityProvidersSingleResponse8KOtOl1HResultSchemasFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersSingleResponse8KOtOl1HResultSchemasFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                           `json:"type"`
	JSON identityProvidersSingleResponse8KOtOl1HResultSchemasFacebookJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasFacebookJSON contains the
// JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasFacebook]
type identityProvidersSingleResponse8KOtOl1HResultSchemasFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersSingleResponse8KOtOl1HResultSchemasFacebook) implementsIdentityProvidersSingleResponse8KOtOl1HResult() {
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                 `json:"client_secret"`
	JSON         identityProvidersSingleResponse8KOtOl1HResultSchemasFacebookConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasFacebookConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasFacebookConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersSingleResponse8KOtOl1HResultSchemasFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                       `json:"user_deprovision"`
	JSON            identityProvidersSingleResponse8KOtOl1HResultSchemasFacebookScimConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasFacebookScimConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasGitHub struct {
	// UUID
	ID     string                                                           `json:"id"`
	Config IdentityProvidersSingleResponse8KOtOl1HResultSchemasGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersSingleResponse8KOtOl1HResultSchemasGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                         `json:"type"`
	JSON identityProvidersSingleResponse8KOtOl1HResultSchemasGitHubJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasGitHubJSON contains the JSON
// metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasGitHub]
type identityProvidersSingleResponse8KOtOl1HResultSchemasGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersSingleResponse8KOtOl1HResultSchemasGitHub) implementsIdentityProvidersSingleResponse8KOtOl1HResult() {
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                               `json:"client_secret"`
	JSON         identityProvidersSingleResponse8KOtOl1HResultSchemasGitHubConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasGitHubConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasGitHubConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersSingleResponse8KOtOl1HResultSchemasGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProvidersSingleResponse8KOtOl1HResultSchemasGitHubScimConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasGitHubScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasGitHubScimConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogle struct {
	// UUID
	ID     string                                                           `json:"id"`
	Config IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                         `json:"type"`
	JSON identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleJSON contains the JSON
// metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogle]
type identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogle) implementsIdentityProvidersSingleResponse8KOtOl1HResult() {
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                               `json:"client_secret"`
	JSON         identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleScimConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleScimConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleApps struct {
	// UUID
	ID     string                                                               `json:"id"`
	Config IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                             `json:"type"`
	JSON identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsJSON contains the
// JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleApps]
type identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleApps) implementsIdentityProvidersSingleResponse8KOtOl1HResult() {
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                   `json:"client_secret"`
	JSON         identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsConfigJSON struct {
	AppsDomain   apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                         `json:"user_deprovision"`
	JSON            identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsScimConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsScimConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasLinkedin struct {
	// UUID
	ID     string                                                             `json:"id"`
	Config IdentityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                           `json:"type"`
	JSON identityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinJSON contains the
// JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasLinkedin]
type identityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersSingleResponse8KOtOl1HResultSchemasLinkedin) implementsIdentityProvidersSingleResponse8KOtOl1HResult() {
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                 `json:"client_secret"`
	JSON         identityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                       `json:"user_deprovision"`
	JSON            identityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinScimConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinScimConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasOidc struct {
	// UUID
	ID     string                                                         `json:"id"`
	Config IdentityProvidersSingleResponse8KOtOl1HResultSchemasOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersSingleResponse8KOtOl1HResultSchemasOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                       `json:"type"`
	JSON identityProvidersSingleResponse8KOtOl1HResultSchemasOidcJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasOidcJSON contains the JSON
// metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasOidc]
type identityProvidersSingleResponse8KOtOl1HResultSchemasOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersSingleResponse8KOtOl1HResultSchemasOidc) implementsIdentityProvidersSingleResponse8KOtOl1HResult() {
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL string `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL string `json:"certs_url"`
	// List of custom claims that will be pulled from your id_token and added to your
	// signed Access JWT token.
	Claims []string `json:"claims"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// OAuth scopes
	Scopes []string `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL string                                                             `json:"token_url"`
	JSON     identityProvidersSingleResponse8KOtOl1HResultSchemasOidcConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasOidcConfigJSON contains the
// JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasOidcConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasOidcConfigJSON struct {
	AuthURL      apijson.Field
	CertsURL     apijson.Field
	Claims       apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	Scopes       apijson.Field
	TokenURL     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersSingleResponse8KOtOl1HResultSchemasOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProvidersSingleResponse8KOtOl1HResultSchemasOidcScimConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasOidcScimConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasOidcScimConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasOkta struct {
	// UUID
	ID     string                                                         `json:"id"`
	Config IdentityProvidersSingleResponse8KOtOl1HResultSchemasOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersSingleResponse8KOtOl1HResultSchemasOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                       `json:"type"`
	JSON identityProvidersSingleResponse8KOtOl1HResultSchemasOktaJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasOktaJSON contains the JSON
// metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasOkta]
type identityProvidersSingleResponse8KOtOl1HResultSchemasOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersSingleResponse8KOtOl1HResultSchemasOkta) implementsIdentityProvidersSingleResponse8KOtOl1HResult() {
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasOktaConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your okta account url
	OktaAccount string                                                             `json:"okta_account"`
	JSON        identityProvidersSingleResponse8KOtOl1HResultSchemasOktaConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasOktaConfigJSON contains the
// JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasOktaConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasOktaConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	OktaAccount  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersSingleResponse8KOtOl1HResultSchemasOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProvidersSingleResponse8KOtOl1HResultSchemasOktaScimConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasOktaScimConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasOktaScimConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasOnelogin struct {
	// UUID
	ID     string                                                             `json:"id"`
	Config IdentityProvidersSingleResponse8KOtOl1HResultSchemasOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersSingleResponse8KOtOl1HResultSchemasOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                           `json:"type"`
	JSON identityProvidersSingleResponse8KOtOl1HResultSchemasOneloginJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasOneloginJSON contains the
// JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasOnelogin]
type identityProvidersSingleResponse8KOtOl1HResultSchemasOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersSingleResponse8KOtOl1HResultSchemasOnelogin) implementsIdentityProvidersSingleResponse8KOtOl1HResult() {
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasOneloginConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your OneLogin account url
	OneloginAccount string                                                                 `json:"onelogin_account"`
	JSON            identityProvidersSingleResponse8KOtOl1HResultSchemasOneloginConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasOneloginConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasOneloginConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasOneloginConfigJSON struct {
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersSingleResponse8KOtOl1HResultSchemasOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                       `json:"user_deprovision"`
	JSON            identityProvidersSingleResponse8KOtOl1HResultSchemasOneloginScimConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasOneloginScimConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasPingone struct {
	// UUID
	ID     string                                                            `json:"id"`
	Config IdentityProvidersSingleResponse8KOtOl1HResultSchemasPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersSingleResponse8KOtOl1HResultSchemasPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                          `json:"type"`
	JSON identityProvidersSingleResponse8KOtOl1HResultSchemasPingoneJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasPingoneJSON contains the
// JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasPingone]
type identityProvidersSingleResponse8KOtOl1HResultSchemasPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersSingleResponse8KOtOl1HResultSchemasPingone) implementsIdentityProvidersSingleResponse8KOtOl1HResult() {
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasPingoneConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your PingOne environment identifier
	PingEnvID string                                                                `json:"ping_env_id"`
	JSON      identityProvidersSingleResponse8KOtOl1HResultSchemasPingoneConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasPingoneConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasPingoneConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasPingoneConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	PingEnvID    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersSingleResponse8KOtOl1HResultSchemasPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProvidersSingleResponse8KOtOl1HResultSchemasPingoneScimConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasPingoneScimConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasSaml struct {
	// UUID
	ID     string                                                         `json:"id"`
	Config IdentityProvidersSingleResponse8KOtOl1HResultSchemasSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersSingleResponse8KOtOl1HResultSchemasSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                       `json:"type"`
	JSON identityProvidersSingleResponse8KOtOl1HResultSchemasSamlJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasSamlJSON contains the JSON
// metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasSaml]
type identityProvidersSingleResponse8KOtOl1HResultSchemasSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersSingleResponse8KOtOl1HResultSchemasSaml) implementsIdentityProvidersSingleResponse8KOtOl1HResult() {
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []IdentityProvidersSingleResponse8KOtOl1HResultSchemasSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                             `json:"sso_target_url"`
	JSON         identityProvidersSingleResponse8KOtOl1HResultSchemasSamlConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasSamlConfigJSON contains the
// JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasSamlConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasSamlConfigJSON struct {
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

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                            `json:"header_name"`
	JSON       identityProvidersSingleResponse8KOtOl1HResultSchemasSamlConfigHeaderAttributeJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasSamlConfigHeaderAttribute]
type identityProvidersSingleResponse8KOtOl1HResultSchemasSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersSingleResponse8KOtOl1HResultSchemasSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProvidersSingleResponse8KOtOl1HResultSchemasSamlScimConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasSamlScimConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasSamlScimConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasYandex struct {
	// UUID
	ID     string                                                           `json:"id"`
	Config IdentityProvidersSingleResponse8KOtOl1HResultSchemasYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig IdentityProvidersSingleResponse8KOtOl1HResultSchemasYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type string                                                         `json:"type"`
	JSON identityProvidersSingleResponse8KOtOl1HResultSchemasYandexJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasYandexJSON contains the JSON
// metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasYandex]
type identityProvidersSingleResponse8KOtOl1HResultSchemasYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r IdentityProvidersSingleResponse8KOtOl1HResultSchemasYandex) implementsIdentityProvidersSingleResponse8KOtOl1HResult() {
}

type IdentityProvidersSingleResponse8KOtOl1HResultSchemasYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                               `json:"client_secret"`
	JSON         identityProvidersSingleResponse8KOtOl1HResultSchemasYandexConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasYandexConfigJSON contains
// the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasYandexConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type IdentityProvidersSingleResponse8KOtOl1HResultSchemasYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            identityProvidersSingleResponse8KOtOl1HResultSchemasYandexScimConfigJSON `json:"-"`
}

// identityProvidersSingleResponse8KOtOl1HResultSchemasYandexScimConfigJSON
// contains the JSON metadata for the struct
// [IdentityProvidersSingleResponse8KOtOl1HResultSchemasYandexScimConfig]
type identityProvidersSingleResponse8KOtOl1HResultSchemasYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IdentityProvidersSingleResponse8KOtOl1HResultSchemasYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IdentityProvidersSingleResponse8KOtOl1HSuccess bool

const (
	IdentityProvidersSingleResponse8KOtOl1HSuccessTrue IdentityProvidersSingleResponse8KOtOl1HSuccess = true
)

// This interface is a union satisfied by one of the following:
// [ZoneAccessIdentityProviderUpdateParamsSchemasAzureAd],
// [ZoneAccessIdentityProviderUpdateParamsSchemasCentrify],
// [ZoneAccessIdentityProviderUpdateParamsSchemasFacebook],
// [ZoneAccessIdentityProviderUpdateParamsSchemasGitHub],
// [ZoneAccessIdentityProviderUpdateParamsSchemasGoogle],
// [ZoneAccessIdentityProviderUpdateParamsSchemasGoogleApps],
// [ZoneAccessIdentityProviderUpdateParamsSchemasLinkedin],
// [ZoneAccessIdentityProviderUpdateParamsSchemasOidc],
// [ZoneAccessIdentityProviderUpdateParamsSchemasOkta],
// [ZoneAccessIdentityProviderUpdateParamsSchemasOnelogin],
// [ZoneAccessIdentityProviderUpdateParamsSchemasPingone],
// [ZoneAccessIdentityProviderUpdateParamsSchemasSaml],
// [ZoneAccessIdentityProviderUpdateParamsSchemasYandex].
type ZoneAccessIdentityProviderUpdateParams interface {
	ImplementsZoneAccessIdentityProviderUpdateParams()
}

type ZoneAccessIdentityProviderUpdateParamsSchemasAzureAd struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasAzureAdConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasAzureAdScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsSchemasAzureAd) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsSchemasAzureAdConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your Azure directory uuid
	DirectoryID param.Field[string] `json:"directory_id"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasAzureAdConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsSchemasAzureAdScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsSchemasAzureAdScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderUpdateParamsSchemasCentrify struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasCentrifyConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasCentrifyScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsSchemasCentrify) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsSchemasCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount param.Field[string] `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID param.Field[string] `json:"centrify_app_id"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsSchemasCentrifyScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsSchemasCentrifyScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderUpdateParamsSchemasFacebook struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasFacebookConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasFacebookScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsSchemasFacebook) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsSchemasFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasFacebookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsSchemasFacebookScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsSchemasFacebookScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderUpdateParamsSchemasGitHub struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasGitHubConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasGitHubScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsSchemasGitHub) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsSchemasGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsSchemasGitHubScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsSchemasGitHubScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderUpdateParamsSchemasGoogle struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasGoogleConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasGoogleScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsSchemasGoogle) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsSchemasGoogleConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsSchemasGoogleScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsSchemasGoogleScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderUpdateParamsSchemasGoogleApps struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasGoogleAppsConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasGoogleAppsScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsSchemasGoogleApps) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsSchemasGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain param.Field[string] `json:"apps_domain"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsSchemasGoogleAppsScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsSchemasGoogleAppsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderUpdateParamsSchemasLinkedin struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasLinkedinConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasLinkedinScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsSchemasLinkedin) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsSchemasLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasLinkedinConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsSchemasLinkedinScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsSchemasLinkedinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderUpdateParamsSchemasOidc struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasOidcConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasOidcScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsSchemasOidc) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsSchemasOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL param.Field[string] `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL param.Field[string] `json:"certs_url"`
	// List of custom claims that will be pulled from your id_token and added to your
	// signed Access JWT token.
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// OAuth scopes
	Scopes param.Field[[]string] `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL param.Field[string] `json:"token_url"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasOidcConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsSchemasOidcScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsSchemasOidcScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderUpdateParamsSchemasOkta struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasOktaConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasOktaScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsSchemasOkta) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsSchemasOktaConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your okta account url
	OktaAccount param.Field[string] `json:"okta_account"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsSchemasOktaScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsSchemasOktaScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderUpdateParamsSchemasOnelogin struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasOneloginConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasOneloginScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsSchemasOnelogin) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsSchemasOneloginConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your OneLogin account url
	OneloginAccount param.Field[string] `json:"onelogin_account"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsSchemasOneloginScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsSchemasOneloginScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderUpdateParamsSchemasPingone struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasPingoneConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasPingoneScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsSchemasPingone) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsSchemasPingoneConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your PingOne environment identifier
	PingEnvID param.Field[string] `json:"ping_env_id"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsSchemasPingoneScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsSchemasPingoneScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderUpdateParamsSchemasSaml struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasSamlConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasSamlScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsSchemasSaml) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsSchemasSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]ZoneAccessIdentityProviderUpdateParamsSchemasSamlConfigHeaderAttribute] `json:"header_attributes"`
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

func (r ZoneAccessIdentityProviderUpdateParamsSchemasSamlConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderUpdateParamsSchemasSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasSamlConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsSchemasSamlScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsSchemasSamlScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderUpdateParamsSchemasYandex struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasYandexConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsSchemasYandexScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsSchemasYandex) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsSchemasYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderUpdateParamsSchemasYandexConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsSchemasYandexScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsSchemasYandexScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This interface is a union satisfied by one of the following:
// [ZoneAccessIdentityProviderAddParamsSchemasAzureAd],
// [ZoneAccessIdentityProviderAddParamsSchemasCentrify],
// [ZoneAccessIdentityProviderAddParamsSchemasFacebook],
// [ZoneAccessIdentityProviderAddParamsSchemasGitHub],
// [ZoneAccessIdentityProviderAddParamsSchemasGoogle],
// [ZoneAccessIdentityProviderAddParamsSchemasGoogleApps],
// [ZoneAccessIdentityProviderAddParamsSchemasLinkedin],
// [ZoneAccessIdentityProviderAddParamsSchemasOidc],
// [ZoneAccessIdentityProviderAddParamsSchemasOkta],
// [ZoneAccessIdentityProviderAddParamsSchemasOnelogin],
// [ZoneAccessIdentityProviderAddParamsSchemasPingone],
// [ZoneAccessIdentityProviderAddParamsSchemasSaml],
// [ZoneAccessIdentityProviderAddParamsSchemasYandex].
type ZoneAccessIdentityProviderAddParams interface {
	ImplementsZoneAccessIdentityProviderAddParams()
}

type ZoneAccessIdentityProviderAddParamsSchemasAzureAd struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsSchemasAzureAdConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsSchemasAzureAdScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsSchemasAzureAd) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsSchemasAzureAdConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your Azure directory uuid
	DirectoryID param.Field[string] `json:"directory_id"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasAzureAdConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsSchemasAzureAdScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsSchemasAzureAdScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderAddParamsSchemasCentrify struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsSchemasCentrifyConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsSchemasCentrifyScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsSchemasCentrify) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsSchemasCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount param.Field[string] `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID param.Field[string] `json:"centrify_app_id"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsSchemasCentrifyScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsSchemasCentrifyScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderAddParamsSchemasFacebook struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsSchemasFacebookConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsSchemasFacebookScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsSchemasFacebook) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsSchemasFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasFacebookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsSchemasFacebookScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsSchemasFacebookScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderAddParamsSchemasGitHub struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsSchemasGitHubConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsSchemasGitHubScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsSchemasGitHub) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsSchemasGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsSchemasGitHubScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsSchemasGitHubScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderAddParamsSchemasGoogle struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsSchemasGoogleConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsSchemasGoogleScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsSchemasGoogle) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsSchemasGoogleConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsSchemasGoogleScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsSchemasGoogleScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderAddParamsSchemasGoogleApps struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsSchemasGoogleAppsConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsSchemasGoogleAppsScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsSchemasGoogleApps) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsSchemasGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain param.Field[string] `json:"apps_domain"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsSchemasGoogleAppsScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsSchemasGoogleAppsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderAddParamsSchemasLinkedin struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsSchemasLinkedinConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsSchemasLinkedinScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsSchemasLinkedin) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsSchemasLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasLinkedinConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsSchemasLinkedinScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsSchemasLinkedinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderAddParamsSchemasOidc struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsSchemasOidcConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsSchemasOidcScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsSchemasOidc) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsSchemasOidcConfig struct {
	// The authorization_endpoint URL of your IdP
	AuthURL param.Field[string] `json:"auth_url"`
	// The jwks_uri endpoint of your IdP to allow the IdP keys to sign the tokens
	CertsURL param.Field[string] `json:"certs_url"`
	// List of custom claims that will be pulled from your id_token and added to your
	// signed Access JWT token.
	Claims param.Field[[]string] `json:"claims"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// OAuth scopes
	Scopes param.Field[[]string] `json:"scopes"`
	// The token_endpoint URL of your IdP
	TokenURL param.Field[string] `json:"token_url"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasOidcConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsSchemasOidcScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsSchemasOidcScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderAddParamsSchemasOkta struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsSchemasOktaConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsSchemasOktaScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsSchemasOkta) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsSchemasOktaConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your okta account url
	OktaAccount param.Field[string] `json:"okta_account"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsSchemasOktaScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsSchemasOktaScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderAddParamsSchemasOnelogin struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsSchemasOneloginConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsSchemasOneloginScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsSchemasOnelogin) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsSchemasOneloginConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your OneLogin account url
	OneloginAccount param.Field[string] `json:"onelogin_account"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsSchemasOneloginScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsSchemasOneloginScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderAddParamsSchemasPingone struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsSchemasPingoneConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsSchemasPingoneScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsSchemasPingone) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsSchemasPingoneConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your PingOne environment identifier
	PingEnvID param.Field[string] `json:"ping_env_id"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsSchemasPingoneScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsSchemasPingoneScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderAddParamsSchemasSaml struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsSchemasSamlConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsSchemasSamlScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsSchemasSaml) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsSchemasSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]ZoneAccessIdentityProviderAddParamsSchemasSamlConfigHeaderAttribute] `json:"header_attributes"`
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

func (r ZoneAccessIdentityProviderAddParamsSchemasSamlConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderAddParamsSchemasSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasSamlConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsSchemasSamlScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsSchemasSamlScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderAddParamsSchemasYandex struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsSchemasYandexConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsSchemasYandexScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[string] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsSchemasYandex) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsSchemasYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderAddParamsSchemasYandexConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsSchemasYandexScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsSchemasYandexScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

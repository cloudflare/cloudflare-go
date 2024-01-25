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
func (r *ZoneAccessIdentityProviderService) Get(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ZoneAccessIdentityProviderGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/identity_providers/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured identity provider.
func (r *ZoneAccessIdentityProviderService) Update(ctx context.Context, identifier string, uuid string, body ZoneAccessIdentityProviderUpdateParams, opts ...option.RequestOption) (res *ZoneAccessIdentityProviderUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/identity_providers/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all configured identity providers.
func (r *ZoneAccessIdentityProviderService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ZoneAccessIdentityProviderListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/identity_providers", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an identity provider from Access.
func (r *ZoneAccessIdentityProviderService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ZoneAccessIdentityProviderDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/identity_providers/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new identity provider to Access.
func (r *ZoneAccessIdentityProviderService) Add(ctx context.Context, identifier string, body ZoneAccessIdentityProviderAddParams, opts ...option.RequestOption) (res *ZoneAccessIdentityProviderAddResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/identity_providers", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneAccessIdentityProviderGetResponse struct {
	Errors   []ZoneAccessIdentityProviderGetResponseError   `json:"errors"`
	Messages []ZoneAccessIdentityProviderGetResponseMessage `json:"messages"`
	Result   ZoneAccessIdentityProviderGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessIdentityProviderGetResponseSuccess `json:"success"`
	JSON    zoneAccessIdentityProviderGetResponseJSON    `json:"-"`
}

// zoneAccessIdentityProviderGetResponseJSON contains the JSON metadata for the
// struct [ZoneAccessIdentityProviderGetResponse]
type zoneAccessIdentityProviderGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderGetResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneAccessIdentityProviderGetResponseErrorJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseErrorJSON contains the JSON metadata for
// the struct [ZoneAccessIdentityProviderGetResponseError]
type zoneAccessIdentityProviderGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderGetResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneAccessIdentityProviderGetResponseMessageJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseMessageJSON contains the JSON metadata for
// the struct [ZoneAccessIdentityProviderGetResponseMessage]
type zoneAccessIdentityProviderGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAd],
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrify],
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebook],
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHub],
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogle],
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleApps],
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedin],
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidc],
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOkta],
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOnelogin],
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingone],
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSaml] or
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandex].
type ZoneAccessIdentityProviderGetResponseResult interface {
	implementsZoneAccessIdentityProviderGetResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessIdentityProviderGetResponseResult)(nil)).Elem(), "")
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAd struct {
	// UUID
	ID     string                                                                  `json:"id"`
	Config ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType `json:"type"`
	JSON zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAd]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAd) implementsZoneAccessIdentityProviderGetResponseResult() {
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Should Cloudflare try to load authentication contexts from your account
	ConditionalAccessEnabled bool `json:"conditional_access_enabled"`
	// Your Azure directory uuid
	DirectoryID string `json:"directory_id"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool                                                                        `json:"support_groups"`
	JSON          zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdConfigJSON struct {
	ClientID                 apijson.Field
	ClientSecret             apijson.Field
	ConditionalAccessEnabled apijson.Field
	DirectoryID              apijson.Field
	SupportGroups            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                            `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdScimConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType string

const (
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdTypeOnetimepin ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType = "onetimepin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdTypeAzureAd    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType = "azureAD"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdTypeSaml       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType = "saml"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdTypeCentrify   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType = "centrify"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdTypeFacebook   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType = "facebook"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdTypeGitHub     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType = "github"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdTypeGoogleApps ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType = "google-apps"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdTypeGoogle     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType = "google"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdTypeLinkedin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType = "linkedin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdTypeOidc       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType = "oidc"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdTypeOkta       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType = "okta"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdTypeOnelogin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType = "onelogin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdTypePingone    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType = "pingone"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdTypeYandex     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasAzureAdType = "yandex"
)

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrify struct {
	// UUID
	ID     string                                                                   `json:"id"`
	Config ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType `json:"type"`
	JSON zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrify]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrify) implementsZoneAccessIdentityProviderGetResponseResult() {
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount string `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID string `json:"centrify_app_id"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                       `json:"client_secret"`
	JSON         zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyScimConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType string

const (
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyTypeOnetimepin ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType = "onetimepin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyTypeAzureAd    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType = "azureAD"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyTypeSaml       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType = "saml"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyTypeCentrify   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType = "centrify"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyTypeFacebook   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType = "facebook"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyTypeGitHub     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType = "github"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyTypeGoogleApps ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType = "google-apps"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyTypeGoogle     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType = "google"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyTypeLinkedin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType = "linkedin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyTypeOidc       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType = "oidc"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyTypeOkta       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType = "okta"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyTypeOnelogin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType = "onelogin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyTypePingone    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType = "pingone"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyTypeYandex     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasCentrifyType = "yandex"
)

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebook struct {
	// UUID
	ID     string                                                                   `json:"id"`
	Config ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType `json:"type"`
	JSON zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebook]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebook) implementsZoneAccessIdentityProviderGetResponseResult() {
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                       `json:"client_secret"`
	JSON         zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookScimConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType string

const (
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookTypeOnetimepin ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType = "onetimepin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookTypeAzureAd    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType = "azureAD"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookTypeSaml       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType = "saml"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookTypeCentrify   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType = "centrify"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookTypeFacebook   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType = "facebook"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookTypeGitHub     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType = "github"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookTypeGoogleApps ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType = "google-apps"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookTypeGoogle     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType = "google"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookTypeLinkedin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType = "linkedin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookTypeOidc       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType = "oidc"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookTypeOkta       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType = "okta"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookTypeOnelogin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType = "onelogin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookTypePingone    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType = "pingone"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookTypeYandex     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasFacebookType = "yandex"
)

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHub struct {
	// UUID
	ID     string                                                                 `json:"id"`
	Config ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType `json:"type"`
	JSON zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHub]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHub) implementsZoneAccessIdentityProviderGetResponseResult() {
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                     `json:"client_secret"`
	JSON         zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubScimConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType string

const (
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubTypeOnetimepin ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType = "onetimepin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubTypeAzureAd    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType = "azureAD"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubTypeSaml       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType = "saml"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubTypeCentrify   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType = "centrify"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubTypeFacebook   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType = "facebook"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubTypeGitHub     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType = "github"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubTypeGoogleApps ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType = "google-apps"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubTypeGoogle     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType = "google"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubTypeLinkedin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType = "linkedin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubTypeOidc       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType = "oidc"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubTypeOkta       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType = "okta"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubTypeOnelogin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType = "onelogin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubTypePingone    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType = "pingone"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubTypeYandex     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGitHubType = "yandex"
)

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogle struct {
	// UUID
	ID     string                                                                 `json:"id"`
	Config ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType `json:"type"`
	JSON zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogle]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogle) implementsZoneAccessIdentityProviderGetResponseResult() {
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                     `json:"client_secret"`
	JSON         zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleScimConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType string

const (
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleTypeOnetimepin ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType = "onetimepin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleTypeAzureAd    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType = "azureAD"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleTypeSaml       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType = "saml"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleTypeCentrify   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType = "centrify"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleTypeFacebook   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType = "facebook"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleTypeGitHub     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType = "github"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleTypeGoogleApps ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType = "google-apps"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleTypeGoogle     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType = "google"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleTypeLinkedin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType = "linkedin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleTypeOidc       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType = "oidc"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleTypeOkta       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType = "okta"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleTypeOnelogin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType = "onelogin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleTypePingone    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType = "pingone"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleTypeYandex     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleType = "yandex"
)

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleApps struct {
	// UUID
	ID     string                                                                     `json:"id"`
	Config ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType `json:"type"`
	JSON zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleApps]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleApps) implementsZoneAccessIdentityProviderGetResponseResult() {
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                         `json:"client_secret"`
	JSON         zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsConfigJSON struct {
	AppsDomain   apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                               `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsScimConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType string

const (
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsTypeOnetimepin ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType = "onetimepin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsTypeAzureAd    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType = "azureAD"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsTypeSaml       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType = "saml"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsTypeCentrify   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType = "centrify"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsTypeFacebook   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType = "facebook"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsTypeGitHub     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType = "github"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsTypeGoogleApps ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType = "google-apps"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsTypeGoogle     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType = "google"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsTypeLinkedin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType = "linkedin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsTypeOidc       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType = "oidc"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsTypeOkta       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType = "okta"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsTypeOnelogin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType = "onelogin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsTypePingone    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType = "pingone"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsTypeYandex     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasGoogleAppsType = "yandex"
)

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedin struct {
	// UUID
	ID     string                                                                   `json:"id"`
	Config ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType `json:"type"`
	JSON zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedin]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedin) implementsZoneAccessIdentityProviderGetResponseResult() {
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                       `json:"client_secret"`
	JSON         zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinScimConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType string

const (
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinTypeOnetimepin ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType = "onetimepin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinTypeAzureAd    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType = "azureAD"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinTypeSaml       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType = "saml"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinTypeCentrify   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType = "centrify"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinTypeFacebook   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType = "facebook"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinTypeGitHub     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType = "github"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinTypeGoogleApps ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType = "google-apps"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinTypeGoogle     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType = "google"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinTypeLinkedin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType = "linkedin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinTypeOidc       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType = "oidc"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinTypeOkta       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType = "okta"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinTypeOnelogin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType = "onelogin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinTypePingone    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType = "pingone"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinTypeYandex     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasLinkedinType = "yandex"
)

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidc struct {
	// UUID
	ID     string                                                               `json:"id"`
	Config ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType `json:"type"`
	JSON zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcJSON contains the
// JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidc]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidc) implementsZoneAccessIdentityProviderGetResponseResult() {
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcConfig struct {
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
	TokenURL string                                                                   `json:"token_url"`
	JSON     zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcConfigJSON struct {
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

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcScimConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType string

const (
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcTypeOnetimepin ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType = "onetimepin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcTypeAzureAd    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType = "azureAD"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcTypeSaml       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType = "saml"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcTypeCentrify   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType = "centrify"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcTypeFacebook   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType = "facebook"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcTypeGitHub     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType = "github"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcTypeGoogleApps ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType = "google-apps"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcTypeGoogle     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType = "google"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcTypeLinkedin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType = "linkedin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcTypeOidc       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType = "oidc"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcTypeOkta       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType = "okta"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcTypeOnelogin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType = "onelogin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcTypePingone    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType = "pingone"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcTypeYandex     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOidcType = "yandex"
)

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOkta struct {
	// UUID
	ID     string                                                               `json:"id"`
	Config ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType `json:"type"`
	JSON zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaJSON contains the
// JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOkta]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOkta) implementsZoneAccessIdentityProviderGetResponseResult() {
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your okta account url
	OktaAccount string                                                                   `json:"okta_account"`
	JSON        zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	OktaAccount  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaScimConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType string

const (
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaTypeOnetimepin ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType = "onetimepin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaTypeAzureAd    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType = "azureAD"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaTypeSaml       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType = "saml"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaTypeCentrify   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType = "centrify"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaTypeFacebook   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType = "facebook"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaTypeGitHub     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType = "github"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaTypeGoogleApps ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType = "google-apps"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaTypeGoogle     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType = "google"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaTypeLinkedin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType = "linkedin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaTypeOidc       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType = "oidc"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaTypeOkta       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType = "okta"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaTypeOnelogin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType = "onelogin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaTypePingone    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType = "pingone"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaTypeYandex     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOktaType = "yandex"
)

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOnelogin struct {
	// UUID
	ID     string                                                                   `json:"id"`
	Config ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType `json:"type"`
	JSON zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOnelogin]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOnelogin) implementsZoneAccessIdentityProviderGetResponseResult() {
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your OneLogin account url
	OneloginAccount string                                                                       `json:"onelogin_account"`
	JSON            zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginConfigJSON struct {
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginScimConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType string

const (
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginTypeOnetimepin ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType = "onetimepin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginTypeAzureAd    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType = "azureAD"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginTypeSaml       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType = "saml"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginTypeCentrify   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType = "centrify"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginTypeFacebook   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType = "facebook"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginTypeGitHub     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType = "github"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginTypeGoogleApps ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType = "google-apps"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginTypeGoogle     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType = "google"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginTypeLinkedin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType = "linkedin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginTypeOidc       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType = "oidc"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginTypeOkta       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType = "okta"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginTypeOnelogin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType = "onelogin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginTypePingone    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType = "pingone"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginTypeYandex     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasOneloginType = "yandex"
)

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingone struct {
	// UUID
	ID     string                                                                  `json:"id"`
	Config ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType `json:"type"`
	JSON zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingone]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingone) implementsZoneAccessIdentityProviderGetResponseResult() {
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your PingOne environment identifier
	PingEnvID string                                                                      `json:"ping_env_id"`
	JSON      zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	PingEnvID    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                            `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneScimConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType string

const (
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneTypeOnetimepin ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType = "onetimepin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneTypeAzureAd    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType = "azureAD"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneTypeSaml       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType = "saml"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneTypeCentrify   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType = "centrify"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneTypeFacebook   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType = "facebook"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneTypeGitHub     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType = "github"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneTypeGoogleApps ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType = "google-apps"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneTypeGoogle     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType = "google"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneTypeLinkedin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType = "linkedin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneTypeOidc       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType = "oidc"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneTypeOkta       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType = "okta"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneTypeOnelogin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType = "onelogin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneTypePingone    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType = "pingone"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneTypeYandex     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasPingoneType = "yandex"
)

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSaml struct {
	// UUID
	ID     string                                                               `json:"id"`
	Config ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType `json:"type"`
	JSON zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlJSON contains the
// JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSaml]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSaml) implementsZoneAccessIdentityProviderGetResponseResult() {
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                                   `json:"sso_target_url"`
	JSON         zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlConfigJSON struct {
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

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                                  `json:"header_name"`
	JSON       zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlConfigHeaderAttributeJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlConfigHeaderAttribute]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlScimConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType string

const (
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlTypeOnetimepin ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType = "onetimepin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlTypeAzureAd    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType = "azureAD"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlTypeSaml       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType = "saml"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlTypeCentrify   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType = "centrify"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlTypeFacebook   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType = "facebook"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlTypeGitHub     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType = "github"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlTypeGoogleApps ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType = "google-apps"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlTypeGoogle     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType = "google"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlTypeLinkedin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType = "linkedin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlTypeOidc       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType = "oidc"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlTypeOkta       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType = "okta"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlTypeOnelogin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType = "onelogin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlTypePingone    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType = "pingone"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlTypeYandex     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasSamlType = "yandex"
)

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandex struct {
	// UUID
	ID     string                                                                 `json:"id"`
	Config ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType `json:"type"`
	JSON zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandex]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandex) implementsZoneAccessIdentityProviderGetResponseResult() {
}

type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                     `json:"client_secret"`
	JSON         zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexScimConfig]
type zoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType string

const (
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexTypeOnetimepin ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType = "onetimepin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexTypeAzureAd    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType = "azureAD"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexTypeSaml       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType = "saml"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexTypeCentrify   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType = "centrify"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexTypeFacebook   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType = "facebook"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexTypeGitHub     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType = "github"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexTypeGoogleApps ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType = "google-apps"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexTypeGoogle     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType = "google"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexTypeLinkedin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType = "linkedin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexTypeOidc       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType = "oidc"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexTypeOkta       ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType = "okta"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexTypeOnelogin   ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType = "onelogin"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexTypePingone    ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType = "pingone"
	ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexTypeYandex     ZoneAccessIdentityProviderGetResponseResultPajwohLqSchemasYandexType = "yandex"
)

// Whether the API call was successful
type ZoneAccessIdentityProviderGetResponseSuccess bool

const (
	ZoneAccessIdentityProviderGetResponseSuccessTrue ZoneAccessIdentityProviderGetResponseSuccess = true
)

type ZoneAccessIdentityProviderUpdateResponse struct {
	Errors   []ZoneAccessIdentityProviderUpdateResponseError   `json:"errors"`
	Messages []ZoneAccessIdentityProviderUpdateResponseMessage `json:"messages"`
	Result   ZoneAccessIdentityProviderUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessIdentityProviderUpdateResponseSuccess `json:"success"`
	JSON    zoneAccessIdentityProviderUpdateResponseJSON    `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneAccessIdentityProviderUpdateResponse]
type zoneAccessIdentityProviderUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderUpdateResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneAccessIdentityProviderUpdateResponseErrorJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseErrorJSON contains the JSON metadata for
// the struct [ZoneAccessIdentityProviderUpdateResponseError]
type zoneAccessIdentityProviderUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderUpdateResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneAccessIdentityProviderUpdateResponseMessageJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseMessageJSON contains the JSON metadata
// for the struct [ZoneAccessIdentityProviderUpdateResponseMessage]
type zoneAccessIdentityProviderUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAd],
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrify],
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebook],
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHub],
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogle],
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleApps],
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedin],
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidc],
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOkta],
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOnelogin],
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingone],
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSaml] or
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandex].
type ZoneAccessIdentityProviderUpdateResponseResult interface {
	implementsZoneAccessIdentityProviderUpdateResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessIdentityProviderUpdateResponseResult)(nil)).Elem(), "")
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAd struct {
	// UUID
	ID     string                                                                     `json:"id"`
	Config ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType `json:"type"`
	JSON zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAd]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAd) implementsZoneAccessIdentityProviderUpdateResponseResult() {
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Should Cloudflare try to load authentication contexts from your account
	ConditionalAccessEnabled bool `json:"conditional_access_enabled"`
	// Your Azure directory uuid
	DirectoryID string `json:"directory_id"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool                                                                           `json:"support_groups"`
	JSON          zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdConfigJSON struct {
	ClientID                 apijson.Field
	ClientSecret             apijson.Field
	ConditionalAccessEnabled apijson.Field
	DirectoryID              apijson.Field
	SupportGroups            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                               `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdScimConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType string

const (
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdTypeOnetimepin ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType = "onetimepin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdTypeAzureAd    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType = "azureAD"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdTypeSaml       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType = "saml"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdTypeCentrify   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType = "centrify"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdTypeFacebook   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType = "facebook"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdTypeGitHub     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType = "github"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdTypeGoogleApps ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType = "google-apps"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdTypeGoogle     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType = "google"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdTypeLinkedin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType = "linkedin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdTypeOidc       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType = "oidc"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdTypeOkta       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType = "okta"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdTypeOnelogin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType = "onelogin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdTypePingone    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType = "pingone"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdTypeYandex     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasAzureAdType = "yandex"
)

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrify struct {
	// UUID
	ID     string                                                                      `json:"id"`
	Config ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType `json:"type"`
	JSON zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrify]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrify) implementsZoneAccessIdentityProviderUpdateResponseResult() {
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount string `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID string `json:"centrify_app_id"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                          `json:"client_secret"`
	JSON         zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                                `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyScimConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType string

const (
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyTypeOnetimepin ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType = "onetimepin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyTypeAzureAd    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType = "azureAD"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyTypeSaml       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType = "saml"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyTypeCentrify   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType = "centrify"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyTypeFacebook   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType = "facebook"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyTypeGitHub     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType = "github"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyTypeGoogleApps ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType = "google-apps"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyTypeGoogle     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType = "google"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyTypeLinkedin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType = "linkedin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyTypeOidc       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType = "oidc"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyTypeOkta       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType = "okta"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyTypeOnelogin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType = "onelogin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyTypePingone    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType = "pingone"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyTypeYandex     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasCentrifyType = "yandex"
)

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebook struct {
	// UUID
	ID     string                                                                      `json:"id"`
	Config ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType `json:"type"`
	JSON zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebook]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebook) implementsZoneAccessIdentityProviderUpdateResponseResult() {
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                          `json:"client_secret"`
	JSON         zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                                `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookScimConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType string

const (
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookTypeOnetimepin ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType = "onetimepin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookTypeAzureAd    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType = "azureAD"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookTypeSaml       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType = "saml"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookTypeCentrify   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType = "centrify"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookTypeFacebook   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType = "facebook"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookTypeGitHub     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType = "github"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookTypeGoogleApps ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType = "google-apps"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookTypeGoogle     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType = "google"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookTypeLinkedin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType = "linkedin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookTypeOidc       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType = "oidc"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookTypeOkta       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType = "okta"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookTypeOnelogin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType = "onelogin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookTypePingone    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType = "pingone"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookTypeYandex     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasFacebookType = "yandex"
)

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHub struct {
	// UUID
	ID     string                                                                    `json:"id"`
	Config ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType `json:"type"`
	JSON zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHub]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHub) implementsZoneAccessIdentityProviderUpdateResponseResult() {
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                        `json:"client_secret"`
	JSON         zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                              `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubScimConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType string

const (
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubTypeOnetimepin ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType = "onetimepin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubTypeAzureAd    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType = "azureAD"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubTypeSaml       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType = "saml"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubTypeCentrify   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType = "centrify"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubTypeFacebook   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType = "facebook"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubTypeGitHub     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType = "github"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubTypeGoogleApps ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType = "google-apps"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubTypeGoogle     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType = "google"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubTypeLinkedin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType = "linkedin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubTypeOidc       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType = "oidc"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubTypeOkta       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType = "okta"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubTypeOnelogin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType = "onelogin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubTypePingone    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType = "pingone"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubTypeYandex     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGitHubType = "yandex"
)

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogle struct {
	// UUID
	ID     string                                                                    `json:"id"`
	Config ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType `json:"type"`
	JSON zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogle]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogle) implementsZoneAccessIdentityProviderUpdateResponseResult() {
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                        `json:"client_secret"`
	JSON         zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                              `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleScimConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType string

const (
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleTypeOnetimepin ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType = "onetimepin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleTypeAzureAd    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType = "azureAD"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleTypeSaml       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType = "saml"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleTypeCentrify   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType = "centrify"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleTypeFacebook   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType = "facebook"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleTypeGitHub     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType = "github"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleTypeGoogleApps ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType = "google-apps"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleTypeGoogle     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType = "google"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleTypeLinkedin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType = "linkedin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleTypeOidc       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType = "oidc"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleTypeOkta       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType = "okta"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleTypeOnelogin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType = "onelogin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleTypePingone    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType = "pingone"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleTypeYandex     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleType = "yandex"
)

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleApps struct {
	// UUID
	ID     string                                                                        `json:"id"`
	Config ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType `json:"type"`
	JSON zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleApps]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleApps) implementsZoneAccessIdentityProviderUpdateResponseResult() {
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                            `json:"client_secret"`
	JSON         zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsConfigJSON struct {
	AppsDomain   apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                                  `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsScimConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType string

const (
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsTypeOnetimepin ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType = "onetimepin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsTypeAzureAd    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType = "azureAD"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsTypeSaml       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType = "saml"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsTypeCentrify   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType = "centrify"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsTypeFacebook   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType = "facebook"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsTypeGitHub     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType = "github"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsTypeGoogleApps ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType = "google-apps"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsTypeGoogle     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType = "google"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsTypeLinkedin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType = "linkedin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsTypeOidc       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType = "oidc"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsTypeOkta       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType = "okta"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsTypeOnelogin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType = "onelogin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsTypePingone    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType = "pingone"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsTypeYandex     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasGoogleAppsType = "yandex"
)

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedin struct {
	// UUID
	ID     string                                                                      `json:"id"`
	Config ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType `json:"type"`
	JSON zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedin]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedin) implementsZoneAccessIdentityProviderUpdateResponseResult() {
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                          `json:"client_secret"`
	JSON         zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                                `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinScimConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType string

const (
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinTypeOnetimepin ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType = "onetimepin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinTypeAzureAd    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType = "azureAD"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinTypeSaml       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType = "saml"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinTypeCentrify   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType = "centrify"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinTypeFacebook   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType = "facebook"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinTypeGitHub     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType = "github"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinTypeGoogleApps ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType = "google-apps"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinTypeGoogle     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType = "google"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinTypeLinkedin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType = "linkedin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinTypeOidc       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType = "oidc"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinTypeOkta       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType = "okta"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinTypeOnelogin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType = "onelogin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinTypePingone    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType = "pingone"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinTypeYandex     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasLinkedinType = "yandex"
)

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidc struct {
	// UUID
	ID     string                                                                  `json:"id"`
	Config ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType `json:"type"`
	JSON zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidc]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidc) implementsZoneAccessIdentityProviderUpdateResponseResult() {
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcConfig struct {
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
	TokenURL string                                                                      `json:"token_url"`
	JSON     zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcConfigJSON struct {
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

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                            `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcScimConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType string

const (
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcTypeOnetimepin ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType = "onetimepin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcTypeAzureAd    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType = "azureAD"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcTypeSaml       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType = "saml"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcTypeCentrify   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType = "centrify"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcTypeFacebook   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType = "facebook"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcTypeGitHub     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType = "github"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcTypeGoogleApps ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType = "google-apps"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcTypeGoogle     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType = "google"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcTypeLinkedin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType = "linkedin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcTypeOidc       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType = "oidc"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcTypeOkta       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType = "okta"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcTypeOnelogin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType = "onelogin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcTypePingone    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType = "pingone"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcTypeYandex     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOidcType = "yandex"
)

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOkta struct {
	// UUID
	ID     string                                                                  `json:"id"`
	Config ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType `json:"type"`
	JSON zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOkta]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOkta) implementsZoneAccessIdentityProviderUpdateResponseResult() {
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your okta account url
	OktaAccount string                                                                      `json:"okta_account"`
	JSON        zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	OktaAccount  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                            `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaScimConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType string

const (
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaTypeOnetimepin ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType = "onetimepin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaTypeAzureAd    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType = "azureAD"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaTypeSaml       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType = "saml"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaTypeCentrify   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType = "centrify"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaTypeFacebook   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType = "facebook"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaTypeGitHub     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType = "github"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaTypeGoogleApps ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType = "google-apps"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaTypeGoogle     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType = "google"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaTypeLinkedin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType = "linkedin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaTypeOidc       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType = "oidc"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaTypeOkta       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType = "okta"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaTypeOnelogin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType = "onelogin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaTypePingone    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType = "pingone"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaTypeYandex     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOktaType = "yandex"
)

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOnelogin struct {
	// UUID
	ID     string                                                                      `json:"id"`
	Config ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType `json:"type"`
	JSON zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOnelogin]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOnelogin) implementsZoneAccessIdentityProviderUpdateResponseResult() {
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your OneLogin account url
	OneloginAccount string                                                                          `json:"onelogin_account"`
	JSON            zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginConfigJSON struct {
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                                `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginScimConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType string

const (
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginTypeOnetimepin ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType = "onetimepin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginTypeAzureAd    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType = "azureAD"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginTypeSaml       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType = "saml"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginTypeCentrify   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType = "centrify"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginTypeFacebook   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType = "facebook"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginTypeGitHub     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType = "github"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginTypeGoogleApps ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType = "google-apps"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginTypeGoogle     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType = "google"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginTypeLinkedin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType = "linkedin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginTypeOidc       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType = "oidc"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginTypeOkta       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType = "okta"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginTypeOnelogin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType = "onelogin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginTypePingone    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType = "pingone"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginTypeYandex     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasOneloginType = "yandex"
)

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingone struct {
	// UUID
	ID     string                                                                     `json:"id"`
	Config ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType `json:"type"`
	JSON zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingone]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingone) implementsZoneAccessIdentityProviderUpdateResponseResult() {
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your PingOne environment identifier
	PingEnvID string                                                                         `json:"ping_env_id"`
	JSON      zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	PingEnvID    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                               `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneScimConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType string

const (
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneTypeOnetimepin ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType = "onetimepin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneTypeAzureAd    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType = "azureAD"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneTypeSaml       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType = "saml"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneTypeCentrify   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType = "centrify"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneTypeFacebook   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType = "facebook"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneTypeGitHub     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType = "github"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneTypeGoogleApps ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType = "google-apps"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneTypeGoogle     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType = "google"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneTypeLinkedin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType = "linkedin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneTypeOidc       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType = "oidc"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneTypeOkta       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType = "okta"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneTypeOnelogin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType = "onelogin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneTypePingone    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType = "pingone"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneTypeYandex     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasPingoneType = "yandex"
)

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSaml struct {
	// UUID
	ID     string                                                                  `json:"id"`
	Config ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType `json:"type"`
	JSON zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSaml]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSaml) implementsZoneAccessIdentityProviderUpdateResponseResult() {
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                                      `json:"sso_target_url"`
	JSON         zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlConfigJSON struct {
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

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                                     `json:"header_name"`
	JSON       zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlConfigHeaderAttributeJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlConfigHeaderAttribute]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                            `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlScimConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType string

const (
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlTypeOnetimepin ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType = "onetimepin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlTypeAzureAd    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType = "azureAD"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlTypeSaml       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType = "saml"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlTypeCentrify   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType = "centrify"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlTypeFacebook   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType = "facebook"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlTypeGitHub     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType = "github"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlTypeGoogleApps ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType = "google-apps"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlTypeGoogle     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType = "google"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlTypeLinkedin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType = "linkedin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlTypeOidc       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType = "oidc"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlTypeOkta       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType = "okta"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlTypeOnelogin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType = "onelogin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlTypePingone    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType = "pingone"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlTypeYandex     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasSamlType = "yandex"
)

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandex struct {
	// UUID
	ID     string                                                                    `json:"id"`
	Config ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType `json:"type"`
	JSON zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandex]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandex) implementsZoneAccessIdentityProviderUpdateResponseResult() {
}

type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                        `json:"client_secret"`
	JSON         zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                              `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexScimConfig]
type zoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType string

const (
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexTypeOnetimepin ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType = "onetimepin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexTypeAzureAd    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType = "azureAD"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexTypeSaml       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType = "saml"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexTypeCentrify   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType = "centrify"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexTypeFacebook   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType = "facebook"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexTypeGitHub     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType = "github"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexTypeGoogleApps ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType = "google-apps"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexTypeGoogle     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType = "google"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexTypeLinkedin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType = "linkedin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexTypeOidc       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType = "oidc"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexTypeOkta       ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType = "okta"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexTypeOnelogin   ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType = "onelogin"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexTypePingone    ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType = "pingone"
	ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexTypeYandex     ZoneAccessIdentityProviderUpdateResponseResultPajwohLqSchemasYandexType = "yandex"
)

// Whether the API call was successful
type ZoneAccessIdentityProviderUpdateResponseSuccess bool

const (
	ZoneAccessIdentityProviderUpdateResponseSuccessTrue ZoneAccessIdentityProviderUpdateResponseSuccess = true
)

type ZoneAccessIdentityProviderListResponse struct {
	Errors     []ZoneAccessIdentityProviderListResponseError    `json:"errors"`
	Messages   []ZoneAccessIdentityProviderListResponseMessage  `json:"messages"`
	Result     []ZoneAccessIdentityProviderListResponseResult   `json:"result"`
	ResultInfo ZoneAccessIdentityProviderListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneAccessIdentityProviderListResponseSuccess `json:"success"`
	JSON    zoneAccessIdentityProviderListResponseJSON    `json:"-"`
}

// zoneAccessIdentityProviderListResponseJSON contains the JSON metadata for the
// struct [ZoneAccessIdentityProviderListResponse]
type zoneAccessIdentityProviderListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderListResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneAccessIdentityProviderListResponseErrorJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneAccessIdentityProviderListResponseError]
type zoneAccessIdentityProviderListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderListResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneAccessIdentityProviderListResponseMessageJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseMessageJSON contains the JSON metadata for
// the struct [ZoneAccessIdentityProviderListResponseMessage]
type zoneAccessIdentityProviderListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAd],
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrify],
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebook],
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHub],
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogle],
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleApps],
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedin],
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidc],
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOkta],
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnelogin],
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingone],
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSaml],
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandex] or
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepin].
type ZoneAccessIdentityProviderListResponseResult interface {
	implementsZoneAccessIdentityProviderListResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessIdentityProviderListResponseResult)(nil)).Elem(), "")
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAd struct {
	// UUID
	ID     string                                                                   `json:"id"`
	Config ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType `json:"type"`
	JSON zoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAd]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAd) implementsZoneAccessIdentityProviderListResponseResult() {
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Should Cloudflare try to load authentication contexts from your account
	ConditionalAccessEnabled bool `json:"conditional_access_enabled"`
	// Your Azure directory uuid
	DirectoryID string `json:"directory_id"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool                                                                         `json:"support_groups"`
	JSON          zoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdConfigJSON struct {
	ClientID                 apijson.Field
	ClientSecret             apijson.Field
	ConditionalAccessEnabled apijson.Field
	DirectoryID              apijson.Field
	SupportGroups            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdScimConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType string

const (
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdTypeOnetimepin ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType = "onetimepin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdTypeAzureAd    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType = "azureAD"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdTypeSaml       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType = "saml"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdTypeCentrify   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType = "centrify"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdTypeFacebook   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType = "facebook"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdTypeGitHub     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType = "github"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdTypeGoogleApps ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType = "google-apps"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdTypeGoogle     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType = "google"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdTypeLinkedin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType = "linkedin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdTypeOidc       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType = "oidc"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdTypeOkta       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType = "okta"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdTypeOnelogin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType = "onelogin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdTypePingone    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType = "pingone"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdTypeYandex     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasAzureAdType = "yandex"
)

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrify struct {
	// UUID
	ID     string                                                                    `json:"id"`
	Config ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType `json:"type"`
	JSON zoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrify]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrify) implementsZoneAccessIdentityProviderListResponseResult() {
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount string `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID string `json:"centrify_app_id"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                        `json:"client_secret"`
	JSON         zoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                              `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyScimConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType string

const (
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyTypeOnetimepin ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType = "onetimepin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyTypeAzureAd    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType = "azureAD"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyTypeSaml       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType = "saml"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyTypeCentrify   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType = "centrify"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyTypeFacebook   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType = "facebook"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyTypeGitHub     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType = "github"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyTypeGoogleApps ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType = "google-apps"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyTypeGoogle     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType = "google"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyTypeLinkedin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType = "linkedin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyTypeOidc       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType = "oidc"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyTypeOkta       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType = "okta"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyTypeOnelogin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType = "onelogin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyTypePingone    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType = "pingone"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyTypeYandex     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasCentrifyType = "yandex"
)

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebook struct {
	// UUID
	ID     string                                                                    `json:"id"`
	Config ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType `json:"type"`
	JSON zoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebook]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebook) implementsZoneAccessIdentityProviderListResponseResult() {
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                        `json:"client_secret"`
	JSON         zoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                              `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookScimConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType string

const (
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookTypeOnetimepin ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType = "onetimepin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookTypeAzureAd    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType = "azureAD"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookTypeSaml       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType = "saml"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookTypeCentrify   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType = "centrify"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookTypeFacebook   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType = "facebook"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookTypeGitHub     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType = "github"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookTypeGoogleApps ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType = "google-apps"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookTypeGoogle     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType = "google"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookTypeLinkedin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType = "linkedin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookTypeOidc       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType = "oidc"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookTypeOkta       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType = "okta"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookTypeOnelogin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType = "onelogin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookTypePingone    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType = "pingone"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookTypeYandex     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasFacebookType = "yandex"
)

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHub struct {
	// UUID
	ID     string                                                                  `json:"id"`
	Config ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType `json:"type"`
	JSON zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHub]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHub) implementsZoneAccessIdentityProviderListResponseResult() {
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                      `json:"client_secret"`
	JSON         zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                            `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubScimConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType string

const (
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubTypeOnetimepin ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType = "onetimepin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubTypeAzureAd    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType = "azureAD"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubTypeSaml       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType = "saml"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubTypeCentrify   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType = "centrify"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubTypeFacebook   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType = "facebook"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubTypeGitHub     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType = "github"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubTypeGoogleApps ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType = "google-apps"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubTypeGoogle     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType = "google"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubTypeLinkedin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType = "linkedin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubTypeOidc       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType = "oidc"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubTypeOkta       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType = "okta"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubTypeOnelogin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType = "onelogin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubTypePingone    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType = "pingone"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubTypeYandex     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGitHubType = "yandex"
)

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogle struct {
	// UUID
	ID     string                                                                  `json:"id"`
	Config ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType `json:"type"`
	JSON zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogle]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogle) implementsZoneAccessIdentityProviderListResponseResult() {
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                      `json:"client_secret"`
	JSON         zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                            `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleScimConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType string

const (
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleTypeOnetimepin ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType = "onetimepin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleTypeAzureAd    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType = "azureAD"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleTypeSaml       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType = "saml"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleTypeCentrify   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType = "centrify"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleTypeFacebook   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType = "facebook"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleTypeGitHub     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType = "github"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleTypeGoogleApps ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType = "google-apps"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleTypeGoogle     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType = "google"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleTypeLinkedin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType = "linkedin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleTypeOidc       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType = "oidc"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleTypeOkta       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType = "okta"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleTypeOnelogin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType = "onelogin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleTypePingone    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType = "pingone"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleTypeYandex     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleType = "yandex"
)

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleApps struct {
	// UUID
	ID     string                                                                      `json:"id"`
	Config ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType `json:"type"`
	JSON zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleApps]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleApps) implementsZoneAccessIdentityProviderListResponseResult() {
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                          `json:"client_secret"`
	JSON         zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsConfigJSON struct {
	AppsDomain   apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                                `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsScimConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType string

const (
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsTypeOnetimepin ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType = "onetimepin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsTypeAzureAd    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType = "azureAD"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsTypeSaml       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType = "saml"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsTypeCentrify   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType = "centrify"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsTypeFacebook   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType = "facebook"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsTypeGitHub     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType = "github"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsTypeGoogleApps ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType = "google-apps"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsTypeGoogle     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType = "google"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsTypeLinkedin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType = "linkedin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsTypeOidc       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType = "oidc"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsTypeOkta       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType = "okta"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsTypeOnelogin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType = "onelogin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsTypePingone    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType = "pingone"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsTypeYandex     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasGoogleAppsType = "yandex"
)

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedin struct {
	// UUID
	ID     string                                                                    `json:"id"`
	Config ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType `json:"type"`
	JSON zoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedin]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedin) implementsZoneAccessIdentityProviderListResponseResult() {
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                        `json:"client_secret"`
	JSON         zoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                              `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinScimConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType string

const (
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinTypeOnetimepin ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType = "onetimepin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinTypeAzureAd    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType = "azureAD"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinTypeSaml       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType = "saml"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinTypeCentrify   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType = "centrify"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinTypeFacebook   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType = "facebook"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinTypeGitHub     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType = "github"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinTypeGoogleApps ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType = "google-apps"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinTypeGoogle     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType = "google"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinTypeLinkedin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType = "linkedin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinTypeOidc       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType = "oidc"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinTypeOkta       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType = "okta"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinTypeOnelogin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType = "onelogin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinTypePingone    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType = "pingone"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinTypeYandex     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasLinkedinType = "yandex"
)

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidc struct {
	// UUID
	ID     string                                                                `json:"id"`
	Config ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType `json:"type"`
	JSON zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcJSON contains the
// JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidc]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidc) implementsZoneAccessIdentityProviderListResponseResult() {
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcConfig struct {
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
	TokenURL string                                                                    `json:"token_url"`
	JSON     zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcConfigJSON struct {
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

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcScimConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType string

const (
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcTypeOnetimepin ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType = "onetimepin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcTypeAzureAd    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType = "azureAD"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcTypeSaml       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType = "saml"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcTypeCentrify   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType = "centrify"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcTypeFacebook   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType = "facebook"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcTypeGitHub     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType = "github"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcTypeGoogleApps ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType = "google-apps"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcTypeGoogle     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType = "google"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcTypeLinkedin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType = "linkedin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcTypeOidc       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType = "oidc"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcTypeOkta       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType = "okta"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcTypeOnelogin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType = "onelogin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcTypePingone    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType = "pingone"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcTypeYandex     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOidcType = "yandex"
)

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOkta struct {
	// UUID
	ID     string                                                                `json:"id"`
	Config ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType `json:"type"`
	JSON zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaJSON contains the
// JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOkta]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOkta) implementsZoneAccessIdentityProviderListResponseResult() {
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your okta account url
	OktaAccount string                                                                    `json:"okta_account"`
	JSON        zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	OktaAccount  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaScimConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType string

const (
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaTypeOnetimepin ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType = "onetimepin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaTypeAzureAd    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType = "azureAD"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaTypeSaml       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType = "saml"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaTypeCentrify   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType = "centrify"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaTypeFacebook   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType = "facebook"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaTypeGitHub     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType = "github"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaTypeGoogleApps ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType = "google-apps"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaTypeGoogle     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType = "google"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaTypeLinkedin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType = "linkedin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaTypeOidc       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType = "oidc"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaTypeOkta       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType = "okta"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaTypeOnelogin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType = "onelogin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaTypePingone    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType = "pingone"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaTypeYandex     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOktaType = "yandex"
)

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnelogin struct {
	// UUID
	ID     string                                                                    `json:"id"`
	Config ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType `json:"type"`
	JSON zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnelogin]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnelogin) implementsZoneAccessIdentityProviderListResponseResult() {
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your OneLogin account url
	OneloginAccount string                                                                        `json:"onelogin_account"`
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginConfigJSON struct {
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                              `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginScimConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType string

const (
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginTypeOnetimepin ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType = "onetimepin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginTypeAzureAd    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType = "azureAD"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginTypeSaml       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType = "saml"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginTypeCentrify   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType = "centrify"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginTypeFacebook   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType = "facebook"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginTypeGitHub     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType = "github"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginTypeGoogleApps ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType = "google-apps"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginTypeGoogle     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType = "google"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginTypeLinkedin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType = "linkedin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginTypeOidc       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType = "oidc"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginTypeOkta       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType = "okta"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginTypeOnelogin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType = "onelogin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginTypePingone    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType = "pingone"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginTypeYandex     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOneloginType = "yandex"
)

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingone struct {
	// UUID
	ID     string                                                                   `json:"id"`
	Config ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType `json:"type"`
	JSON zoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingone]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingone) implementsZoneAccessIdentityProviderListResponseResult() {
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your PingOne environment identifier
	PingEnvID string                                                                       `json:"ping_env_id"`
	JSON      zoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	PingEnvID    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneScimConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType string

const (
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneTypeOnetimepin ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType = "onetimepin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneTypeAzureAd    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType = "azureAD"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneTypeSaml       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType = "saml"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneTypeCentrify   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType = "centrify"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneTypeFacebook   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType = "facebook"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneTypeGitHub     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType = "github"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneTypeGoogleApps ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType = "google-apps"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneTypeGoogle     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType = "google"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneTypeLinkedin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType = "linkedin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneTypeOidc       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType = "oidc"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneTypeOkta       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType = "okta"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneTypeOnelogin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType = "onelogin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneTypePingone    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType = "pingone"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneTypeYandex     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasPingoneType = "yandex"
)

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSaml struct {
	// UUID
	ID     string                                                                `json:"id"`
	Config ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType `json:"type"`
	JSON zoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlJSON contains the
// JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSaml]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSaml) implementsZoneAccessIdentityProviderListResponseResult() {
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                                    `json:"sso_target_url"`
	JSON         zoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlConfigJSON struct {
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

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                                   `json:"header_name"`
	JSON       zoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlConfigHeaderAttributeJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlConfigHeaderAttribute]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlScimConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType string

const (
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlTypeOnetimepin ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType = "onetimepin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlTypeAzureAd    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType = "azureAD"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlTypeSaml       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType = "saml"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlTypeCentrify   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType = "centrify"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlTypeFacebook   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType = "facebook"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlTypeGitHub     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType = "github"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlTypeGoogleApps ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType = "google-apps"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlTypeGoogle     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType = "google"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlTypeLinkedin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType = "linkedin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlTypeOidc       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType = "oidc"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlTypeOkta       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType = "okta"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlTypeOnelogin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType = "onelogin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlTypePingone    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType = "pingone"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlTypeYandex     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasSamlType = "yandex"
)

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandex struct {
	// UUID
	ID     string                                                                  `json:"id"`
	Config ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType `json:"type"`
	JSON zoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandex]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandex) implementsZoneAccessIdentityProviderListResponseResult() {
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                      `json:"client_secret"`
	JSON         zoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                            `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexScimConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType string

const (
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexTypeOnetimepin ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType = "onetimepin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexTypeAzureAd    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType = "azureAD"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexTypeSaml       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType = "saml"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexTypeCentrify   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType = "centrify"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexTypeFacebook   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType = "facebook"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexTypeGitHub     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType = "github"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexTypeGoogleApps ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType = "google-apps"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexTypeGoogle     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType = "google"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexTypeLinkedin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType = "linkedin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexTypeOidc       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType = "oidc"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexTypeOkta       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType = "okta"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexTypeOnelogin   ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType = "onelogin"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexTypePingone    ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType = "pingone"
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexTypeYandex     ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasYandexType = "yandex"
)

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepin struct {
	// UUID
	ID     string      `json:"id"`
	Config interface{} `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepinScimConfig `json:"scim_config"`
	Type       ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepinType       `json:"type"`
	JSON       zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepinJSON       `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepinJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepin]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepin) implementsZoneAccessIdentityProviderListResponseResult() {
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                                `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepinScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepinScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepinScimConfig]
type zoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepinType string

const (
	ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepinTypeOnetimepin ZoneAccessIdentityProviderListResponseResultPajwohLqSchemasOnetimepinType = "onetimepin"
)

type ZoneAccessIdentityProviderListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       zoneAccessIdentityProviderListResponseResultInfoJSON `json:"-"`
}

// zoneAccessIdentityProviderListResponseResultInfoJSON contains the JSON metadata
// for the struct [ZoneAccessIdentityProviderListResponseResultInfo]
type zoneAccessIdentityProviderListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessIdentityProviderListResponseSuccess bool

const (
	ZoneAccessIdentityProviderListResponseSuccessTrue ZoneAccessIdentityProviderListResponseSuccess = true
)

type ZoneAccessIdentityProviderDeleteResponse struct {
	Errors   []ZoneAccessIdentityProviderDeleteResponseError   `json:"errors"`
	Messages []ZoneAccessIdentityProviderDeleteResponseMessage `json:"messages"`
	Result   ZoneAccessIdentityProviderDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessIdentityProviderDeleteResponseSuccess `json:"success"`
	JSON    zoneAccessIdentityProviderDeleteResponseJSON    `json:"-"`
}

// zoneAccessIdentityProviderDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneAccessIdentityProviderDeleteResponse]
type zoneAccessIdentityProviderDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderDeleteResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneAccessIdentityProviderDeleteResponseErrorJSON `json:"-"`
}

// zoneAccessIdentityProviderDeleteResponseErrorJSON contains the JSON metadata for
// the struct [ZoneAccessIdentityProviderDeleteResponseError]
type zoneAccessIdentityProviderDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderDeleteResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zoneAccessIdentityProviderDeleteResponseMessageJSON `json:"-"`
}

// zoneAccessIdentityProviderDeleteResponseMessageJSON contains the JSON metadata
// for the struct [ZoneAccessIdentityProviderDeleteResponseMessage]
type zoneAccessIdentityProviderDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderDeleteResponseResult struct {
	// UUID
	ID   string                                             `json:"id"`
	JSON zoneAccessIdentityProviderDeleteResponseResultJSON `json:"-"`
}

// zoneAccessIdentityProviderDeleteResponseResultJSON contains the JSON metadata
// for the struct [ZoneAccessIdentityProviderDeleteResponseResult]
type zoneAccessIdentityProviderDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessIdentityProviderDeleteResponseSuccess bool

const (
	ZoneAccessIdentityProviderDeleteResponseSuccessTrue ZoneAccessIdentityProviderDeleteResponseSuccess = true
)

type ZoneAccessIdentityProviderAddResponse struct {
	Errors   []ZoneAccessIdentityProviderAddResponseError   `json:"errors"`
	Messages []ZoneAccessIdentityProviderAddResponseMessage `json:"messages"`
	Result   ZoneAccessIdentityProviderAddResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessIdentityProviderAddResponseSuccess `json:"success"`
	JSON    zoneAccessIdentityProviderAddResponseJSON    `json:"-"`
}

// zoneAccessIdentityProviderAddResponseJSON contains the JSON metadata for the
// struct [ZoneAccessIdentityProviderAddResponse]
type zoneAccessIdentityProviderAddResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderAddResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneAccessIdentityProviderAddResponseErrorJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseErrorJSON contains the JSON metadata for
// the struct [ZoneAccessIdentityProviderAddResponseError]
type zoneAccessIdentityProviderAddResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderAddResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneAccessIdentityProviderAddResponseMessageJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseMessageJSON contains the JSON metadata for
// the struct [ZoneAccessIdentityProviderAddResponseMessage]
type zoneAccessIdentityProviderAddResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAd],
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrify],
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebook],
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHub],
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogle],
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleApps],
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedin],
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidc],
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOkta],
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOnelogin],
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingone],
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSaml] or
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandex].
type ZoneAccessIdentityProviderAddResponseResult interface {
	implementsZoneAccessIdentityProviderAddResponseResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneAccessIdentityProviderAddResponseResult)(nil)).Elem(), "")
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAd struct {
	// UUID
	ID     string                                                                  `json:"id"`
	Config ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType `json:"type"`
	JSON zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAd]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAd) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAd) implementsZoneAccessIdentityProviderAddResponseResult() {
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Should Cloudflare try to load authentication contexts from your account
	ConditionalAccessEnabled bool `json:"conditional_access_enabled"`
	// Your Azure directory uuid
	DirectoryID string `json:"directory_id"`
	// Should Cloudflare try to load groups from your account
	SupportGroups bool                                                                        `json:"support_groups"`
	JSON          zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdConfigJSON struct {
	ClientID                 apijson.Field
	ClientSecret             apijson.Field
	ConditionalAccessEnabled apijson.Field
	DirectoryID              apijson.Field
	SupportGroups            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                            `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdScimConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType string

const (
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdTypeOnetimepin ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType = "onetimepin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdTypeAzureAd    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType = "azureAD"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdTypeSaml       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType = "saml"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdTypeCentrify   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType = "centrify"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdTypeFacebook   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType = "facebook"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdTypeGitHub     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType = "github"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdTypeGoogleApps ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType = "google-apps"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdTypeGoogle     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType = "google"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdTypeLinkedin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType = "linkedin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdTypeOidc       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType = "oidc"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdTypeOkta       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType = "okta"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdTypeOnelogin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType = "onelogin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdTypePingone    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType = "pingone"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdTypeYandex     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasAzureAdType = "yandex"
)

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrify struct {
	// UUID
	ID     string                                                                   `json:"id"`
	Config ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType `json:"type"`
	JSON zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrify]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrify) implementsZoneAccessIdentityProviderAddResponseResult() {
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount string `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID string `json:"centrify_app_id"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                       `json:"client_secret"`
	JSON         zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyConfigJSON struct {
	CentrifyAccount apijson.Field
	CentrifyAppID   apijson.Field
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyScimConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType string

const (
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyTypeOnetimepin ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType = "onetimepin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyTypeAzureAd    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType = "azureAD"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyTypeSaml       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType = "saml"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyTypeCentrify   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType = "centrify"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyTypeFacebook   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType = "facebook"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyTypeGitHub     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType = "github"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyTypeGoogleApps ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType = "google-apps"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyTypeGoogle     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType = "google"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyTypeLinkedin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType = "linkedin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyTypeOidc       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType = "oidc"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyTypeOkta       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType = "okta"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyTypeOnelogin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType = "onelogin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyTypePingone    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType = "pingone"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyTypeYandex     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasCentrifyType = "yandex"
)

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebook struct {
	// UUID
	ID     string                                                                   `json:"id"`
	Config ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType `json:"type"`
	JSON zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebook]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebook) implementsZoneAccessIdentityProviderAddResponseResult() {
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                       `json:"client_secret"`
	JSON         zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookScimConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType string

const (
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookTypeOnetimepin ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType = "onetimepin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookTypeAzureAd    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType = "azureAD"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookTypeSaml       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType = "saml"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookTypeCentrify   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType = "centrify"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookTypeFacebook   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType = "facebook"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookTypeGitHub     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType = "github"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookTypeGoogleApps ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType = "google-apps"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookTypeGoogle     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType = "google"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookTypeLinkedin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType = "linkedin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookTypeOidc       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType = "oidc"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookTypeOkta       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType = "okta"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookTypeOnelogin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType = "onelogin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookTypePingone    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType = "pingone"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookTypeYandex     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasFacebookType = "yandex"
)

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHub struct {
	// UUID
	ID     string                                                                 `json:"id"`
	Config ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType `json:"type"`
	JSON zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHub]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHub) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHub) implementsZoneAccessIdentityProviderAddResponseResult() {
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                     `json:"client_secret"`
	JSON         zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubScimConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType string

const (
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubTypeOnetimepin ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType = "onetimepin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubTypeAzureAd    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType = "azureAD"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubTypeSaml       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType = "saml"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubTypeCentrify   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType = "centrify"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubTypeFacebook   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType = "facebook"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubTypeGitHub     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType = "github"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubTypeGoogleApps ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType = "google-apps"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubTypeGoogle     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType = "google"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubTypeLinkedin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType = "linkedin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubTypeOidc       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType = "oidc"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubTypeOkta       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType = "okta"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubTypeOnelogin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType = "onelogin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubTypePingone    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType = "pingone"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubTypeYandex     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGitHubType = "yandex"
)

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogle struct {
	// UUID
	ID     string                                                                 `json:"id"`
	Config ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType `json:"type"`
	JSON zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogle]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogle) implementsZoneAccessIdentityProviderAddResponseResult() {
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                     `json:"client_secret"`
	JSON         zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleScimConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType string

const (
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleTypeOnetimepin ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType = "onetimepin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleTypeAzureAd    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType = "azureAD"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleTypeSaml       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType = "saml"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleTypeCentrify   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType = "centrify"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleTypeFacebook   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType = "facebook"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleTypeGitHub     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType = "github"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleTypeGoogleApps ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType = "google-apps"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleTypeGoogle     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType = "google"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleTypeLinkedin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType = "linkedin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleTypeOidc       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType = "oidc"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleTypeOkta       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType = "okta"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleTypeOnelogin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType = "onelogin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleTypePingone    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType = "pingone"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleTypeYandex     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleType = "yandex"
)

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleApps struct {
	// UUID
	ID     string                                                                     `json:"id"`
	Config ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType `json:"type"`
	JSON zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleApps]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleApps) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleApps) implementsZoneAccessIdentityProviderAddResponseResult() {
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain string `json:"apps_domain"`
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                         `json:"client_secret"`
	JSON         zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsConfigJSON struct {
	AppsDomain   apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                               `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsScimConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType string

const (
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsTypeOnetimepin ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType = "onetimepin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsTypeAzureAd    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType = "azureAD"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsTypeSaml       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType = "saml"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsTypeCentrify   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType = "centrify"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsTypeFacebook   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType = "facebook"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsTypeGitHub     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType = "github"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsTypeGoogleApps ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType = "google-apps"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsTypeGoogle     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType = "google"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsTypeLinkedin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType = "linkedin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsTypeOidc       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType = "oidc"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsTypeOkta       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType = "okta"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsTypeOnelogin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType = "onelogin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsTypePingone    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType = "pingone"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsTypeYandex     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasGoogleAppsType = "yandex"
)

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedin struct {
	// UUID
	ID     string                                                                   `json:"id"`
	Config ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType `json:"type"`
	JSON zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedin]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedin) implementsZoneAccessIdentityProviderAddResponseResult() {
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                       `json:"client_secret"`
	JSON         zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinScimConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType string

const (
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinTypeOnetimepin ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType = "onetimepin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinTypeAzureAd    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType = "azureAD"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinTypeSaml       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType = "saml"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinTypeCentrify   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType = "centrify"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinTypeFacebook   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType = "facebook"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinTypeGitHub     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType = "github"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinTypeGoogleApps ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType = "google-apps"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinTypeGoogle     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType = "google"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinTypeLinkedin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType = "linkedin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinTypeOidc       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType = "oidc"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinTypeOkta       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType = "okta"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinTypeOnelogin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType = "onelogin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinTypePingone    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType = "pingone"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinTypeYandex     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasLinkedinType = "yandex"
)

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidc struct {
	// UUID
	ID     string                                                               `json:"id"`
	Config ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType `json:"type"`
	JSON zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcJSON contains the
// JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidc]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidc) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidc) implementsZoneAccessIdentityProviderAddResponseResult() {
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcConfig struct {
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
	TokenURL string                                                                   `json:"token_url"`
	JSON     zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcConfigJSON struct {
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

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcScimConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType string

const (
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcTypeOnetimepin ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType = "onetimepin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcTypeAzureAd    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType = "azureAD"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcTypeSaml       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType = "saml"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcTypeCentrify   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType = "centrify"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcTypeFacebook   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType = "facebook"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcTypeGitHub     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType = "github"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcTypeGoogleApps ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType = "google-apps"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcTypeGoogle     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType = "google"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcTypeLinkedin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType = "linkedin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcTypeOidc       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType = "oidc"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcTypeOkta       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType = "okta"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcTypeOnelogin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType = "onelogin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcTypePingone    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType = "pingone"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcTypeYandex     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOidcType = "yandex"
)

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOkta struct {
	// UUID
	ID     string                                                               `json:"id"`
	Config ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType `json:"type"`
	JSON zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaJSON contains the
// JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOkta]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOkta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOkta) implementsZoneAccessIdentityProviderAddResponseResult() {
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your okta account url
	OktaAccount string                                                                   `json:"okta_account"`
	JSON        zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	OktaAccount  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaScimConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType string

const (
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaTypeOnetimepin ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType = "onetimepin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaTypeAzureAd    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType = "azureAD"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaTypeSaml       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType = "saml"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaTypeCentrify   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType = "centrify"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaTypeFacebook   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType = "facebook"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaTypeGitHub     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType = "github"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaTypeGoogleApps ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType = "google-apps"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaTypeGoogle     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType = "google"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaTypeLinkedin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType = "linkedin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaTypeOidc       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType = "oidc"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaTypeOkta       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType = "okta"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaTypeOnelogin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType = "onelogin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaTypePingone    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType = "pingone"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaTypeYandex     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOktaType = "yandex"
)

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOnelogin struct {
	// UUID
	ID     string                                                                   `json:"id"`
	Config ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType `json:"type"`
	JSON zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOnelogin]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOnelogin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOnelogin) implementsZoneAccessIdentityProviderAddResponseResult() {
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your OneLogin account url
	OneloginAccount string                                                                       `json:"onelogin_account"`
	JSON            zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginConfigJSON struct {
	ClientID        apijson.Field
	ClientSecret    apijson.Field
	OneloginAccount apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginScimConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType string

const (
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginTypeOnetimepin ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType = "onetimepin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginTypeAzureAd    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType = "azureAD"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginTypeSaml       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType = "saml"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginTypeCentrify   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType = "centrify"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginTypeFacebook   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType = "facebook"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginTypeGitHub     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType = "github"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginTypeGoogleApps ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType = "google-apps"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginTypeGoogle     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType = "google"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginTypeLinkedin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType = "linkedin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginTypeOidc       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType = "oidc"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginTypeOkta       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType = "okta"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginTypeOnelogin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType = "onelogin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginTypePingone    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType = "pingone"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginTypeYandex     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasOneloginType = "yandex"
)

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingone struct {
	// UUID
	ID     string                                                                  `json:"id"`
	Config ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType `json:"type"`
	JSON zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingone]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingone) implementsZoneAccessIdentityProviderAddResponseResult() {
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string `json:"client_secret"`
	// Your PingOne environment identifier
	PingEnvID string                                                                      `json:"ping_env_id"`
	JSON      zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	PingEnvID    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
	// in the Identity Provider. This cannot be enabled unless user_deprovision is also
	// enabled.
	SeatDeprovision bool `json:"seat_deprovision"`
	// A read-only token generated when the SCIM integration is enabled for the first
	// time. It is redacted on subsequent requests. If you lose this you will need to
	// refresh it token at /access/identity_providers/:idpID/refresh_scim_secret.
	Secret string `json:"secret"`
	// A flag to enable revoking a user's session in Access and Gateway when they have
	// been deprovisioned in the Identity Provider.
	UserDeprovision bool                                                                            `json:"user_deprovision"`
	JSON            zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneScimConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType string

const (
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneTypeOnetimepin ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType = "onetimepin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneTypeAzureAd    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType = "azureAD"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneTypeSaml       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType = "saml"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneTypeCentrify   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType = "centrify"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneTypeFacebook   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType = "facebook"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneTypeGitHub     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType = "github"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneTypeGoogleApps ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType = "google-apps"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneTypeGoogle     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType = "google"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneTypeLinkedin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType = "linkedin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneTypeOidc       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType = "oidc"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneTypeOkta       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType = "okta"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneTypeOnelogin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType = "onelogin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneTypePingone    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType = "pingone"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneTypeYandex     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasPingoneType = "yandex"
)

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSaml struct {
	// UUID
	ID     string                                                               `json:"id"`
	Config ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType `json:"type"`
	JSON zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlJSON contains the
// JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSaml]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSaml) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSaml) implementsZoneAccessIdentityProviderAddResponseResult() {
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes []string `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName string `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes []ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlConfigHeaderAttribute `json:"header_attributes"`
	// X509 certificate to verify the signature in the SAML authentication response
	IdpPublicCerts []string `json:"idp_public_certs"`
	// IdP Entity ID or Issuer URL
	IssuerURL string `json:"issuer_url"`
	// Sign the SAML authentication request with Access credentials. To verify the
	// signature, use the public key from the Access certs endpoints.
	SignRequest bool `json:"sign_request"`
	// URL to send the SAML authentication requests to
	SSOTargetURL string                                                                   `json:"sso_target_url"`
	JSON         zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlConfigJSON struct {
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

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName string `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName string                                                                                  `json:"header_name"`
	JSON       zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlConfigHeaderAttributeJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlConfigHeaderAttributeJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlConfigHeaderAttribute]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlConfigHeaderAttributeJSON struct {
	AttributeName apijson.Field
	HeaderName    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlConfigHeaderAttribute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlScimConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType string

const (
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlTypeOnetimepin ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType = "onetimepin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlTypeAzureAd    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType = "azureAD"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlTypeSaml       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType = "saml"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlTypeCentrify   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType = "centrify"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlTypeFacebook   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType = "facebook"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlTypeGitHub     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType = "github"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlTypeGoogleApps ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType = "google-apps"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlTypeGoogle     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType = "google"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlTypeLinkedin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType = "linkedin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlTypeOidc       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType = "oidc"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlTypeOkta       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType = "okta"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlTypeOnelogin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType = "onelogin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlTypePingone    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType = "pingone"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlTypeYandex     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasSamlType = "yandex"
)

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandex struct {
	// UUID
	ID     string                                                                 `json:"id"`
	Config ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexConfig `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name string `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexScimConfig `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType `json:"type"`
	JSON zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexJSON contains
// the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandex]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexJSON struct {
	ID          apijson.Field
	Config      apijson.Field
	Name        apijson.Field
	ScimConfig  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandex) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandex) implementsZoneAccessIdentityProviderAddResponseResult() {
}

type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexConfig struct {
	// Your OAuth Client ID
	ClientID string `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret string                                                                     `json:"client_secret"`
	JSON         zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexConfigJSON struct {
	ClientID     apijson.Field
	ClientSecret apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexScimConfig struct {
	// A flag to enable or disable SCIM for the identity provider.
	Enabled bool `json:"enabled"`
	// A flag to revoke a user's session in Access and force a reauthentication on the
	// user's Gateway session when they have been added or removed from a group in the
	// Identity Provider.
	GroupMemberDeprovision bool `json:"group_member_deprovision"`
	// A flag to remove a user's seat in Zero Trust when they have been deprovisioned
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
	JSON            zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexScimConfigJSON `json:"-"`
}

// zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexScimConfigJSON
// contains the JSON metadata for the struct
// [ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexScimConfig]
type zoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexScimConfigJSON struct {
	Enabled                apijson.Field
	GroupMemberDeprovision apijson.Field
	SeatDeprovision        apijson.Field
	Secret                 apijson.Field
	UserDeprovision        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexScimConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType string

const (
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexTypeOnetimepin ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType = "onetimepin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexTypeAzureAd    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType = "azureAD"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexTypeSaml       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType = "saml"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexTypeCentrify   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType = "centrify"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexTypeFacebook   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType = "facebook"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexTypeGitHub     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType = "github"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexTypeGoogleApps ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType = "google-apps"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexTypeGoogle     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType = "google"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexTypeLinkedin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType = "linkedin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexTypeOidc       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType = "oidc"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexTypeOkta       ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType = "okta"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexTypeOnelogin   ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType = "onelogin"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexTypePingone    ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType = "pingone"
	ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexTypeYandex     ZoneAccessIdentityProviderAddResponseResultPajwohLqSchemasYandexType = "yandex"
)

// Whether the API call was successful
type ZoneAccessIdentityProviderAddResponseSuccess bool

const (
	ZoneAccessIdentityProviderAddResponseSuccessTrue ZoneAccessIdentityProviderAddResponseSuccess = true
)

// This interface is a union satisfied by one of the following:
// [ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAd],
// [ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrify],
// [ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebook],
// [ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHub],
// [ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogle],
// [ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleApps],
// [ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedin],
// [ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidc],
// [ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOkta],
// [ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOnelogin],
// [ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingone],
// [ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSaml],
// [ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandex].
type ZoneAccessIdentityProviderUpdateParams interface {
	ImplementsZoneAccessIdentityProviderUpdateParams()
}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAd struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAd) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Should Cloudflare try to load authentication contexts from your account
	ConditionalAccessEnabled param.Field[bool] `json:"conditional_access_enabled"`
	// Your Azure directory uuid
	DirectoryID param.Field[string] `json:"directory_id"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType string

const (
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdTypeOnetimepin ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType = "onetimepin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdTypeAzureAd    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType = "azureAD"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdTypeSaml       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType = "saml"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdTypeCentrify   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType = "centrify"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdTypeFacebook   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType = "facebook"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdTypeGitHub     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType = "github"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdTypeGoogleApps ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType = "google-apps"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdTypeGoogle     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType = "google"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdTypeLinkedin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType = "linkedin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdTypeOidc       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType = "oidc"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdTypeOkta       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType = "okta"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdTypeOnelogin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType = "onelogin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdTypePingone    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType = "pingone"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdTypeYandex     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasAzureAdType = "yandex"
)

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrify struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrify) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount param.Field[string] `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID param.Field[string] `json:"centrify_app_id"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType string

const (
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyTypeOnetimepin ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType = "onetimepin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyTypeAzureAd    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType = "azureAD"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyTypeSaml       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType = "saml"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyTypeCentrify   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType = "centrify"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyTypeFacebook   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType = "facebook"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyTypeGitHub     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType = "github"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyTypeGoogleApps ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType = "google-apps"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyTypeGoogle     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType = "google"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyTypeLinkedin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType = "linkedin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyTypeOidc       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType = "oidc"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyTypeOkta       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType = "okta"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyTypeOnelogin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType = "onelogin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyTypePingone    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType = "pingone"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyTypeYandex     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasCentrifyType = "yandex"
)

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebook struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebook) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType string

const (
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookTypeOnetimepin ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType = "onetimepin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookTypeAzureAd    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType = "azureAD"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookTypeSaml       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType = "saml"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookTypeCentrify   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType = "centrify"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookTypeFacebook   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType = "facebook"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookTypeGitHub     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType = "github"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookTypeGoogleApps ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType = "google-apps"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookTypeGoogle     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType = "google"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookTypeLinkedin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType = "linkedin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookTypeOidc       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType = "oidc"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookTypeOkta       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType = "okta"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookTypeOnelogin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType = "onelogin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookTypePingone    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType = "pingone"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookTypeYandex     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasFacebookType = "yandex"
)

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHub struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHub) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType string

const (
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubTypeOnetimepin ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType = "onetimepin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubTypeAzureAd    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType = "azureAD"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubTypeSaml       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType = "saml"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubTypeCentrify   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType = "centrify"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubTypeFacebook   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType = "facebook"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubTypeGitHub     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType = "github"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubTypeGoogleApps ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType = "google-apps"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubTypeGoogle     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType = "google"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubTypeLinkedin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType = "linkedin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubTypeOidc       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType = "oidc"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubTypeOkta       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType = "okta"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubTypeOnelogin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType = "onelogin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubTypePingone    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType = "pingone"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubTypeYandex     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGitHubType = "yandex"
)

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogle struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogle) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType string

const (
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleTypeOnetimepin ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType = "onetimepin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleTypeAzureAd    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType = "azureAD"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleTypeSaml       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType = "saml"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleTypeCentrify   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType = "centrify"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleTypeFacebook   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType = "facebook"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleTypeGitHub     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType = "github"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleTypeGoogleApps ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType = "google-apps"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleTypeGoogle     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType = "google"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleTypeLinkedin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType = "linkedin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleTypeOidc       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType = "oidc"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleTypeOkta       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType = "okta"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleTypeOnelogin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType = "onelogin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleTypePingone    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType = "pingone"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleTypeYandex     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleType = "yandex"
)

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleApps struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleApps) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain param.Field[string] `json:"apps_domain"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType string

const (
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsTypeOnetimepin ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType = "onetimepin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsTypeAzureAd    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType = "azureAD"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsTypeSaml       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType = "saml"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsTypeCentrify   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType = "centrify"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsTypeFacebook   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType = "facebook"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsTypeGitHub     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType = "github"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsTypeGoogleApps ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType = "google-apps"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsTypeGoogle     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType = "google"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsTypeLinkedin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType = "linkedin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsTypeOidc       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType = "oidc"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsTypeOkta       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType = "okta"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsTypeOnelogin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType = "onelogin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsTypePingone    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType = "pingone"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsTypeYandex     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasGoogleAppsType = "yandex"
)

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedin struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedin) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType string

const (
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinTypeOnetimepin ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType = "onetimepin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinTypeAzureAd    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType = "azureAD"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinTypeSaml       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType = "saml"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinTypeCentrify   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType = "centrify"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinTypeFacebook   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType = "facebook"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinTypeGitHub     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType = "github"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinTypeGoogleApps ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType = "google-apps"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinTypeGoogle     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType = "google"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinTypeLinkedin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType = "linkedin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinTypeOidc       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType = "oidc"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinTypeOkta       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType = "okta"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinTypeOnelogin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType = "onelogin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinTypePingone    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType = "pingone"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinTypeYandex     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasLinkedinType = "yandex"
)

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidc struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidc) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType string

const (
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcTypeOnetimepin ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType = "onetimepin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcTypeAzureAd    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType = "azureAD"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcTypeSaml       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType = "saml"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcTypeCentrify   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType = "centrify"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcTypeFacebook   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType = "facebook"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcTypeGitHub     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType = "github"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcTypeGoogleApps ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType = "google-apps"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcTypeGoogle     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType = "google"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcTypeLinkedin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType = "linkedin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcTypeOidc       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType = "oidc"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcTypeOkta       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType = "okta"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcTypeOnelogin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType = "onelogin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcTypePingone    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType = "pingone"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcTypeYandex     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOidcType = "yandex"
)

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOkta struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOkta) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your okta account url
	OktaAccount param.Field[string] `json:"okta_account"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType string

const (
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaTypeOnetimepin ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType = "onetimepin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaTypeAzureAd    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType = "azureAD"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaTypeSaml       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType = "saml"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaTypeCentrify   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType = "centrify"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaTypeFacebook   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType = "facebook"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaTypeGitHub     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType = "github"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaTypeGoogleApps ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType = "google-apps"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaTypeGoogle     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType = "google"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaTypeLinkedin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType = "linkedin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaTypeOidc       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType = "oidc"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaTypeOkta       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType = "okta"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaTypeOnelogin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType = "onelogin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaTypePingone    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType = "pingone"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaTypeYandex     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOktaType = "yandex"
)

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOnelogin struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOnelogin) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your OneLogin account url
	OneloginAccount param.Field[string] `json:"onelogin_account"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType string

const (
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginTypeOnetimepin ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType = "onetimepin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginTypeAzureAd    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType = "azureAD"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginTypeSaml       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType = "saml"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginTypeCentrify   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType = "centrify"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginTypeFacebook   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType = "facebook"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginTypeGitHub     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType = "github"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginTypeGoogleApps ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType = "google-apps"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginTypeGoogle     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType = "google"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginTypeLinkedin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType = "linkedin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginTypeOidc       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType = "oidc"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginTypeOkta       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType = "okta"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginTypeOnelogin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType = "onelogin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginTypePingone    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType = "pingone"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginTypeYandex     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasOneloginType = "yandex"
)

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingone struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingone) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your PingOne environment identifier
	PingEnvID param.Field[string] `json:"ping_env_id"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType string

const (
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneTypeOnetimepin ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType = "onetimepin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneTypeAzureAd    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType = "azureAD"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneTypeSaml       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType = "saml"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneTypeCentrify   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType = "centrify"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneTypeFacebook   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType = "facebook"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneTypeGitHub     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType = "github"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneTypeGoogleApps ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType = "google-apps"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneTypeGoogle     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType = "google"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneTypeLinkedin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType = "linkedin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneTypeOidc       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType = "oidc"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneTypeOkta       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType = "okta"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneTypeOnelogin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType = "onelogin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneTypePingone    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType = "pingone"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneTypeYandex     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasPingoneType = "yandex"
)

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSaml struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSaml) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlConfigHeaderAttribute] `json:"header_attributes"`
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType string

const (
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlTypeOnetimepin ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType = "onetimepin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlTypeAzureAd    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType = "azureAD"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlTypeSaml       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType = "saml"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlTypeCentrify   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType = "centrify"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlTypeFacebook   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType = "facebook"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlTypeGitHub     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType = "github"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlTypeGoogleApps ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType = "google-apps"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlTypeGoogle     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType = "google"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlTypeLinkedin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType = "linkedin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlTypeOidc       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType = "oidc"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlTypeOkta       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType = "okta"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlTypeOnelogin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType = "onelogin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlTypePingone    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType = "pingone"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlTypeYandex     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasSamlType = "yandex"
)

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandex struct {
	Config param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType] `json:"type"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandex) ImplementsZoneAccessIdentityProviderUpdateParams() {

}

type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexScimConfig struct {
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

func (r ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType string

const (
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexTypeOnetimepin ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType = "onetimepin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexTypeAzureAd    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType = "azureAD"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexTypeSaml       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType = "saml"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexTypeCentrify   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType = "centrify"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexTypeFacebook   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType = "facebook"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexTypeGitHub     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType = "github"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexTypeGoogleApps ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType = "google-apps"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexTypeGoogle     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType = "google"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexTypeLinkedin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType = "linkedin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexTypeOidc       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType = "oidc"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexTypeOkta       ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType = "okta"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexTypeOnelogin   ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType = "onelogin"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexTypePingone    ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType = "pingone"
	ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexTypeYandex     ZoneAccessIdentityProviderUpdateParamsPajwohLqSchemasYandexType = "yandex"
)

// This interface is a union satisfied by one of the following:
// [ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAd],
// [ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrify],
// [ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebook],
// [ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHub],
// [ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogle],
// [ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleApps],
// [ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedin],
// [ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidc],
// [ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOkta],
// [ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOnelogin],
// [ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingone],
// [ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSaml],
// [ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandex].
type ZoneAccessIdentityProviderAddParams interface {
	ImplementsZoneAccessIdentityProviderAddParams()
}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAd struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAd) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAd) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Should Cloudflare try to load authentication contexts from your account
	ConditionalAccessEnabled param.Field[bool] `json:"conditional_access_enabled"`
	// Your Azure directory uuid
	DirectoryID param.Field[string] `json:"directory_id"`
	// Should Cloudflare try to load groups from your account
	SupportGroups param.Field[bool] `json:"support_groups"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType string

const (
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdTypeOnetimepin ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType = "onetimepin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdTypeAzureAd    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType = "azureAD"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdTypeSaml       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType = "saml"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdTypeCentrify   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType = "centrify"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdTypeFacebook   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType = "facebook"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdTypeGitHub     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType = "github"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdTypeGoogleApps ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType = "google-apps"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdTypeGoogle     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType = "google"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdTypeLinkedin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType = "linkedin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdTypeOidc       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType = "oidc"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdTypeOkta       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType = "okta"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdTypeOnelogin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType = "onelogin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdTypePingone    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType = "pingone"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdTypeYandex     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasAzureAdType = "yandex"
)

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrify struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrify) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyConfig struct {
	// Your centrify account url
	CentrifyAccount param.Field[string] `json:"centrify_account"`
	// Your centrify app id
	CentrifyAppID param.Field[string] `json:"centrify_app_id"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType string

const (
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyTypeOnetimepin ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType = "onetimepin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyTypeAzureAd    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType = "azureAD"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyTypeSaml       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType = "saml"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyTypeCentrify   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType = "centrify"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyTypeFacebook   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType = "facebook"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyTypeGitHub     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType = "github"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyTypeGoogleApps ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType = "google-apps"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyTypeGoogle     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType = "google"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyTypeLinkedin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType = "linkedin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyTypeOidc       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType = "oidc"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyTypeOkta       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType = "okta"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyTypeOnelogin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType = "onelogin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyTypePingone    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType = "pingone"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyTypeYandex     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasCentrifyType = "yandex"
)

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebook struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebook) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebook) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType string

const (
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookTypeOnetimepin ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType = "onetimepin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookTypeAzureAd    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType = "azureAD"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookTypeSaml       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType = "saml"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookTypeCentrify   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType = "centrify"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookTypeFacebook   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType = "facebook"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookTypeGitHub     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType = "github"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookTypeGoogleApps ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType = "google-apps"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookTypeGoogle     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType = "google"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookTypeLinkedin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType = "linkedin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookTypeOidc       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType = "oidc"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookTypeOkta       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType = "okta"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookTypeOnelogin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType = "onelogin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookTypePingone    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType = "pingone"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookTypeYandex     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasFacebookType = "yandex"
)

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHub struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHub) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHub) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType string

const (
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubTypeOnetimepin ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType = "onetimepin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubTypeAzureAd    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType = "azureAD"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubTypeSaml       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType = "saml"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubTypeCentrify   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType = "centrify"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubTypeFacebook   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType = "facebook"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubTypeGitHub     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType = "github"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubTypeGoogleApps ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType = "google-apps"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubTypeGoogle     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType = "google"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubTypeLinkedin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType = "linkedin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubTypeOidc       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType = "oidc"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubTypeOkta       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType = "okta"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubTypeOnelogin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType = "onelogin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubTypePingone    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType = "pingone"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubTypeYandex     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGitHubType = "yandex"
)

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogle struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogle) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType string

const (
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleTypeOnetimepin ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType = "onetimepin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleTypeAzureAd    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType = "azureAD"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleTypeSaml       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType = "saml"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleTypeCentrify   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType = "centrify"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleTypeFacebook   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType = "facebook"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleTypeGitHub     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType = "github"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleTypeGoogleApps ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType = "google-apps"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleTypeGoogle     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType = "google"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleTypeLinkedin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType = "linkedin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleTypeOidc       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType = "oidc"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleTypeOkta       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType = "okta"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleTypeOnelogin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType = "onelogin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleTypePingone    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType = "pingone"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleTypeYandex     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleType = "yandex"
)

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleApps struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleApps) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleApps) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsConfig struct {
	// Your companies TLD
	AppsDomain param.Field[string] `json:"apps_domain"`
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType string

const (
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsTypeOnetimepin ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType = "onetimepin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsTypeAzureAd    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType = "azureAD"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsTypeSaml       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType = "saml"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsTypeCentrify   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType = "centrify"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsTypeFacebook   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType = "facebook"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsTypeGitHub     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType = "github"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsTypeGoogleApps ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType = "google-apps"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsTypeGoogle     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType = "google"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsTypeLinkedin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType = "linkedin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsTypeOidc       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType = "oidc"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsTypeOkta       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType = "okta"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsTypeOnelogin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType = "onelogin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsTypePingone    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType = "pingone"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsTypeYandex     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasGoogleAppsType = "yandex"
)

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedin struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedin) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType string

const (
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinTypeOnetimepin ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType = "onetimepin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinTypeAzureAd    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType = "azureAD"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinTypeSaml       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType = "saml"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinTypeCentrify   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType = "centrify"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinTypeFacebook   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType = "facebook"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinTypeGitHub     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType = "github"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinTypeGoogleApps ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType = "google-apps"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinTypeGoogle     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType = "google"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinTypeLinkedin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType = "linkedin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinTypeOidc       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType = "oidc"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinTypeOkta       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType = "okta"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinTypeOnelogin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType = "onelogin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinTypePingone    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType = "pingone"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinTypeYandex     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasLinkedinType = "yandex"
)

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidc struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidc) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidc) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType string

const (
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcTypeOnetimepin ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType = "onetimepin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcTypeAzureAd    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType = "azureAD"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcTypeSaml       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType = "saml"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcTypeCentrify   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType = "centrify"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcTypeFacebook   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType = "facebook"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcTypeGitHub     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType = "github"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcTypeGoogleApps ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType = "google-apps"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcTypeGoogle     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType = "google"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcTypeLinkedin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType = "linkedin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcTypeOidc       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType = "oidc"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcTypeOkta       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType = "okta"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcTypeOnelogin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType = "onelogin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcTypePingone    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType = "pingone"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcTypeYandex     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOidcType = "yandex"
)

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOkta struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOkta) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOkta) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your okta account url
	OktaAccount param.Field[string] `json:"okta_account"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType string

const (
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaTypeOnetimepin ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType = "onetimepin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaTypeAzureAd    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType = "azureAD"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaTypeSaml       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType = "saml"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaTypeCentrify   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType = "centrify"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaTypeFacebook   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType = "facebook"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaTypeGitHub     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType = "github"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaTypeGoogleApps ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType = "google-apps"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaTypeGoogle     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType = "google"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaTypeLinkedin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType = "linkedin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaTypeOidc       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType = "oidc"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaTypeOkta       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType = "okta"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaTypeOnelogin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType = "onelogin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaTypePingone    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType = "pingone"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaTypeYandex     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOktaType = "yandex"
)

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOnelogin struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOnelogin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOnelogin) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your OneLogin account url
	OneloginAccount param.Field[string] `json:"onelogin_account"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType string

const (
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginTypeOnetimepin ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType = "onetimepin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginTypeAzureAd    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType = "azureAD"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginTypeSaml       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType = "saml"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginTypeCentrify   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType = "centrify"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginTypeFacebook   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType = "facebook"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginTypeGitHub     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType = "github"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginTypeGoogleApps ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType = "google-apps"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginTypeGoogle     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType = "google"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginTypeLinkedin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType = "linkedin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginTypeOidc       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType = "oidc"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginTypeOkta       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType = "okta"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginTypeOnelogin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType = "onelogin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginTypePingone    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType = "pingone"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginTypeYandex     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasOneloginType = "yandex"
)

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingone struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingone) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingone) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
	// Your PingOne environment identifier
	PingEnvID param.Field[string] `json:"ping_env_id"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType string

const (
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneTypeOnetimepin ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType = "onetimepin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneTypeAzureAd    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType = "azureAD"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneTypeSaml       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType = "saml"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneTypeCentrify   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType = "centrify"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneTypeFacebook   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType = "facebook"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneTypeGitHub     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType = "github"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneTypeGoogleApps ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType = "google-apps"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneTypeGoogle     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType = "google"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneTypeLinkedin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType = "linkedin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneTypeOidc       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType = "oidc"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneTypeOkta       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType = "okta"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneTypeOnelogin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType = "onelogin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneTypePingone    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType = "pingone"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneTypeYandex     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasPingoneType = "yandex"
)

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSaml struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSaml) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSaml) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlConfig struct {
	// A list of SAML attribute names that will be added to your signed JWT token and
	// can be used in SAML policy rules.
	Attributes param.Field[[]string] `json:"attributes"`
	// The attribute name for email in the SAML response.
	EmailAttributeName param.Field[string] `json:"email_attribute_name"`
	// Add a list of attribute names that will be returned in the response header from
	// the Access callback.
	HeaderAttributes param.Field[[]ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlConfigHeaderAttribute] `json:"header_attributes"`
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlConfigHeaderAttribute struct {
	// attribute name from the IDP
	AttributeName param.Field[string] `json:"attribute_name"`
	// header that will be added on the request to the origin
	HeaderName param.Field[string] `json:"header_name"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlConfigHeaderAttribute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType string

const (
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlTypeOnetimepin ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType = "onetimepin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlTypeAzureAd    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType = "azureAD"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlTypeSaml       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType = "saml"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlTypeCentrify   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType = "centrify"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlTypeFacebook   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType = "facebook"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlTypeGitHub     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType = "github"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlTypeGoogleApps ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType = "google-apps"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlTypeGoogle     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType = "google"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlTypeLinkedin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType = "linkedin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlTypeOidc       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType = "oidc"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlTypeOkta       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType = "okta"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlTypeOnelogin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType = "onelogin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlTypePingone    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType = "pingone"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlTypeYandex     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasSamlType = "yandex"
)

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandex struct {
	Config param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexConfig] `json:"config"`
	// The name of the identity provider, shown to users on the login page.
	Name param.Field[string] `json:"name"`
	// The configuration settings for enabling a System for Cross-Domain Identity
	// Management (SCIM) with the identity provider.
	ScimConfig param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexScimConfig] `json:"scim_config"`
	// The type of identity provider. To determine the value for a specific provider,
	// refer to our
	// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
	Type param.Field[ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType] `json:"type"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandex) ImplementsZoneAccessIdentityProviderAddParams() {

}

type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexConfig struct {
	// Your OAuth Client ID
	ClientID param.Field[string] `json:"client_id"`
	// Your OAuth Client Secret
	ClientSecret param.Field[string] `json:"client_secret"`
}

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration settings for enabling a System for Cross-Domain Identity
// Management (SCIM) with the identity provider.
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexScimConfig struct {
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

func (r ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexScimConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of identity provider. To determine the value for a specific provider,
// refer to our
// [developer documentation](https://developers.cloudflare.com/cloudflare-one/identity/idp-integration/).
type ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType string

const (
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexTypeOnetimepin ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType = "onetimepin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexTypeAzureAd    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType = "azureAD"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexTypeSaml       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType = "saml"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexTypeCentrify   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType = "centrify"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexTypeFacebook   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType = "facebook"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexTypeGitHub     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType = "github"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexTypeGoogleApps ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType = "google-apps"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexTypeGoogle     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType = "google"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexTypeLinkedin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType = "linkedin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexTypeOidc       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType = "oidc"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexTypeOkta       ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType = "okta"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexTypeOnelogin   ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType = "onelogin"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexTypePingone    ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType = "pingone"
	ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexTypeYandex     ZoneAccessIdentityProviderAddParamsPajwohLqSchemasYandexType = "yandex"
)

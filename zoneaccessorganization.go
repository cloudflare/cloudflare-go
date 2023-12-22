// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneAccessOrganizationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAccessOrganizationService]
// method instead.
type ZoneAccessOrganizationService struct {
	Options []option.RequestOption
}

// NewZoneAccessOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAccessOrganizationService(opts ...option.RequestOption) (r *ZoneAccessOrganizationService) {
	r = &ZoneAccessOrganizationService{}
	r.Options = opts
	return
}

// Revokes a user's access across all applications.
func (r *ZoneAccessOrganizationService) RevokeUser(ctx context.Context, identifier interface{}, body ZoneAccessOrganizationRevokeUserParams, opts ...option.RequestOption) (res *ZoneAccessOrganizationRevokeUserResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/organizations/revoke_user", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Sets up a Zero Trust organization for your account.
func (r *ZoneAccessOrganizationService) ZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganization(ctx context.Context, identifier interface{}, body ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationParams, opts ...option.RequestOption) (res *SingleResponsePxUl6qf, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/organizations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the configuration for your Zero Trust organization.
func (r *ZoneAccessOrganizationService) ZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganization(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *SingleResponsePxUl6qf, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/organizations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the configuration for your Zero Trust organization.
func (r *ZoneAccessOrganizationService) ZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganization(ctx context.Context, identifier interface{}, body ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams, opts ...option.RequestOption) (res *SingleResponsePxUl6qf, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/organizations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneAccessOrganizationRevokeUserResponse struct {
	Result  ZoneAccessOrganizationRevokeUserResponseResult  `json:"result"`
	Success ZoneAccessOrganizationRevokeUserResponseSuccess `json:"success"`
	JSON    zoneAccessOrganizationRevokeUserResponseJSON    `json:"-"`
}

// zoneAccessOrganizationRevokeUserResponseJSON contains the JSON metadata for the
// struct [ZoneAccessOrganizationRevokeUserResponse]
type zoneAccessOrganizationRevokeUserResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessOrganizationRevokeUserResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessOrganizationRevokeUserResponseResult bool

const (
	ZoneAccessOrganizationRevokeUserResponseResultTrue  ZoneAccessOrganizationRevokeUserResponseResult = true
	ZoneAccessOrganizationRevokeUserResponseResultFalse ZoneAccessOrganizationRevokeUserResponseResult = false
)

type ZoneAccessOrganizationRevokeUserResponseSuccess bool

const (
	ZoneAccessOrganizationRevokeUserResponseSuccessTrue  ZoneAccessOrganizationRevokeUserResponseSuccess = true
	ZoneAccessOrganizationRevokeUserResponseSuccessFalse ZoneAccessOrganizationRevokeUserResponseSuccess = false
)

type ZoneAccessOrganizationRevokeUserParams struct {
	// The email of the user to revoke.
	Email param.Field[string] `json:"email,required"`
}

func (r ZoneAccessOrganizationRevokeUserParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationParams struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain,required"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name,required"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]                                                                                              `json:"is_ui_read_only"`
	LoginDesign  param.Field[ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign] `json:"login_design"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
}

func (r ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign struct {
	// The background color on your login page.
	BackgroundColor param.Field[string] `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText param.Field[string] `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText param.Field[string] `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath param.Field[string] `json:"logo_path"`
	// The text color on your login page.
	TextColor param.Field[string] `json:"text_color"`
}

func (r ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]                                                                                                 `json:"is_ui_read_only"`
	LoginDesign  param.Field[ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign] `json:"login_design"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
}

func (r ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign struct {
	// The background color on your login page.
	BackgroundColor param.Field[string] `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText param.Field[string] `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText param.Field[string] `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath param.Field[string] `json:"logo_path"`
	// The text color on your login page.
	TextColor param.Field[string] `json:"text_color"`
}

func (r ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

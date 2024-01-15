// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
func (r *ZoneAccessOrganizationService) ZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganization(ctx context.Context, identifier interface{}, body ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationParams, opts ...option.RequestOption) (res *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/organizations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the configuration for your Zero Trust organization.
func (r *ZoneAccessOrganizationService) ZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganization(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/access/organizations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the configuration for your Zero Trust organization.
func (r *ZoneAccessOrganizationService) ZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganization(ctx context.Context, identifier interface{}, body ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams, opts ...option.RequestOption) (res *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse, err error) {
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

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponse struct {
	Errors   []ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseError   `json:"errors"`
	Messages []ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessage `json:"messages"`
	Result   ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseSuccess `json:"success"`
	JSON    zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseJSON    `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponse]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseError struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseErrorJSON `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseError]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessage struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessageJSON `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessage]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResult struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string    `json:"auth_domain"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                                                                                      `json:"is_ui_read_only"`
	LoginDesign  ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string                                                                                             `json:"user_seat_expiration_inactive_time"`
	JSON                           zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultJSON `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResult]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultJSON struct {
	AuthDomain                     apijson.Field
	CreatedAt                      apijson.Field
	IsUiReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	UiReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                                                                        `json:"text_color"`
	JSON      zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesignJSON `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesignJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesign]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseSuccess bool

const (
	ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseSuccessTrue ZoneAccessOrganizationZoneLevelZeroTrustOrganizationNewYourZeroTrustOrganizationResponseSuccess = true
)

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponse struct {
	Errors   []ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseError   `json:"errors"`
	Messages []ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessage `json:"messages"`
	Result   ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseSuccess `json:"success"`
	JSON    zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseJSON    `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponse]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseError struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseErrorJSON `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseError]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessage struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessageJSON `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessage]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResult struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string    `json:"auth_domain"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                                                                                      `json:"is_ui_read_only"`
	LoginDesign  ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string                                                                                             `json:"user_seat_expiration_inactive_time"`
	JSON                           zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultJSON `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResult]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultJSON struct {
	AuthDomain                     apijson.Field
	CreatedAt                      apijson.Field
	IsUiReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	UiReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                                                                        `json:"text_color"`
	JSON      zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesignJSON `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesignJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesign]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseSuccess bool

const (
	ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseSuccessTrue ZoneAccessOrganizationZoneLevelZeroTrustOrganizationGetYourZeroTrustOrganizationResponseSuccess = true
)

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse struct {
	Errors   []ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseError   `json:"errors"`
	Messages []ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessage `json:"messages"`
	Result   ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseSuccess `json:"success"`
	JSON    zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseJSON    `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseError struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseErrorJSON `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseError]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessage struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessageJSON `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessage]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResult struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string    `json:"auth_domain"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                                                                                         `json:"is_ui_read_only"`
	LoginDesign  ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string                                                                                                `json:"user_seat_expiration_inactive_time"`
	JSON                           zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultJSON `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResult]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultJSON struct {
	AuthDomain                     apijson.Field
	CreatedAt                      apijson.Field
	IsUiReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	UiReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                                                                           `json:"text_color"`
	JSON      zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesignJSON `json:"-"`
}

// zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesignJSON
// contains the JSON metadata for the struct
// [ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesign]
type zoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseSuccess bool

const (
	ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseSuccessTrue ZoneAccessOrganizationZoneLevelZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseSuccess = true
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

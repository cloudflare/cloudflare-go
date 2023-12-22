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

// AccountAccessOrganizationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountAccessOrganizationService] method instead.
type AccountAccessOrganizationService struct {
	Options     []option.RequestOption
	RevokeUsers *AccountAccessOrganizationRevokeUserService
}

// NewAccountAccessOrganizationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessOrganizationService(opts ...option.RequestOption) (r *AccountAccessOrganizationService) {
	r = &AccountAccessOrganizationService{}
	r.Options = opts
	r.RevokeUsers = NewAccountAccessOrganizationRevokeUserService(opts...)
	return
}

// Sets up a Zero Trust organization for your account.
func (r *AccountAccessOrganizationService) ZeroTrustOrganizationNewYourZeroTrustOrganization(ctx context.Context, identifier interface{}, body AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParams, opts ...option.RequestOption) (res *SingleResponsePxUl6qf, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/organizations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the configuration for your Zero Trust organization.
func (r *AccountAccessOrganizationService) ZeroTrustOrganizationGetYourZeroTrustOrganization(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *SingleResponsePxUl6qf, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/organizations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the configuration for your Zero Trust organization.
func (r *AccountAccessOrganizationService) ZeroTrustOrganizationUpdateYourZeroTrustOrganization(ctx context.Context, identifier interface{}, body AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams, opts ...option.RequestOption) (res *SingleResponsePxUl6qf, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/access/organizations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SingleResponsePxUl6qf struct {
	Errors   []SingleResponsePxUl6qfError   `json:"errors"`
	Messages []SingleResponsePxUl6qfMessage `json:"messages"`
	Result   SingleResponsePxUl6qfResult    `json:"result"`
	// Whether the API call was successful
	Success SingleResponsePxUl6qfSuccess `json:"success"`
	JSON    singleResponsePxUl6qfJSON    `json:"-"`
}

// singleResponsePxUl6qfJSON contains the JSON metadata for the struct
// [SingleResponsePxUl6qf]
type singleResponsePxUl6qfJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponsePxUl6qf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponsePxUl6qfError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    singleResponsePxUl6qfErrorJSON `json:"-"`
}

// singleResponsePxUl6qfErrorJSON contains the JSON metadata for the struct
// [SingleResponsePxUl6qfError]
type singleResponsePxUl6qfErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponsePxUl6qfError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponsePxUl6qfMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    singleResponsePxUl6qfMessageJSON `json:"-"`
}

// singleResponsePxUl6qfMessageJSON contains the JSON metadata for the struct
// [SingleResponsePxUl6qfMessage]
type singleResponsePxUl6qfMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponsePxUl6qfMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponsePxUl6qfResult struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string    `json:"auth_domain"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                   `json:"is_ui_read_only"`
	LoginDesign  SingleResponsePxUl6qfResultLoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string                          `json:"user_seat_expiration_inactive_time"`
	JSON                           singleResponsePxUl6qfResultJSON `json:"-"`
}

// singleResponsePxUl6qfResultJSON contains the JSON metadata for the struct
// [SingleResponsePxUl6qfResult]
type singleResponsePxUl6qfResultJSON struct {
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

func (r *SingleResponsePxUl6qfResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponsePxUl6qfResultLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                     `json:"text_color"`
	JSON      singleResponsePxUl6qfResultLoginDesignJSON `json:"-"`
}

// singleResponsePxUl6qfResultLoginDesignJSON contains the JSON metadata for the
// struct [SingleResponsePxUl6qfResultLoginDesign]
type singleResponsePxUl6qfResultLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SingleResponsePxUl6qfResultLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SingleResponsePxUl6qfSuccess bool

const (
	SingleResponsePxUl6qfSuccessTrue SingleResponsePxUl6qfSuccess = true
)

type AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParams struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain,required"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name,required"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]                                                                                        `json:"is_ui_read_only"`
	LoginDesign  param.Field[AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign] `json:"login_design"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
}

func (r AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign struct {
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

func (r AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]                                                                                           `json:"is_ui_read_only"`
	LoginDesign  param.Field[AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign] `json:"login_design"`
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

func (r AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign struct {
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

func (r AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

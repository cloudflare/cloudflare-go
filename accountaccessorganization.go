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
func (r *AccountAccessOrganizationService) ZeroTrustOrganizationNewYourZeroTrustOrganization(ctx context.Context, identifier string, body AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParams, opts ...option.RequestOption) (res *AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/organizations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the configuration for your Zero Trust organization.
func (r *AccountAccessOrganizationService) ZeroTrustOrganizationGetYourZeroTrustOrganization(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/organizations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the configuration for your Zero Trust organization.
func (r *AccountAccessOrganizationService) ZeroTrustOrganizationUpdateYourZeroTrustOrganization(ctx context.Context, identifier string, body AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams, opts ...option.RequestOption) (res *AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/organizations", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse struct {
	Errors   []AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseError   `json:"errors"`
	Messages []AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessage `json:"messages"`
	Result   AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseSuccess `json:"success"`
	JSON    accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseJSON    `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse]
type accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseError struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseErrorJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseError]
type accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessage struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessageJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessage]
type accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResult struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                                                                                `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                                                                           `json:"created_at" format:"date-time"`
	CustomPages            AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                                                                                `json:"is_ui_read_only"`
	LoginDesign  AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration string `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string                                                                                       `json:"user_seat_expiration_inactive_time"`
	JSON                           accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResult]
type accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultJSON struct {
	AuthDomain                     apijson.Field
	AutoRedirectToIdentity         apijson.Field
	CreatedAt                      apijson.Field
	CustomPages                    apijson.Field
	IsUiReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	SessionDuration                apijson.Field
	UiReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                                                                                  `json:"identity_denied"`
	JSON           accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultCustomPagesJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultCustomPagesJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultCustomPages]
type accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                                                                  `json:"text_color"`
	JSON      accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesignJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesignJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesign]
type accountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseSuccess bool

const (
	AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseSuccessTrue AccountAccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseSuccess = true
)

type AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse struct {
	Errors   []AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseError   `json:"errors"`
	Messages []AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessage `json:"messages"`
	Result   AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseSuccess `json:"success"`
	JSON    accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseJSON    `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse]
type accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseError struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseErrorJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseError]
type accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessage struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessageJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessage]
type accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResult struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                                                                                `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                                                                           `json:"created_at" format:"date-time"`
	CustomPages            AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                                                                                `json:"is_ui_read_only"`
	LoginDesign  AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration string `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string                                                                                       `json:"user_seat_expiration_inactive_time"`
	JSON                           accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResult]
type accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultJSON struct {
	AuthDomain                     apijson.Field
	AutoRedirectToIdentity         apijson.Field
	CreatedAt                      apijson.Field
	CustomPages                    apijson.Field
	IsUiReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	SessionDuration                apijson.Field
	UiReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                                                                                  `json:"identity_denied"`
	JSON           accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultCustomPagesJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultCustomPagesJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultCustomPages]
type accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                                                                  `json:"text_color"`
	JSON      accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesignJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesignJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesign]
type accountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseSuccess bool

const (
	AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseSuccessTrue AccountAccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseSuccess = true
)

type AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse struct {
	Errors   []AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseError   `json:"errors"`
	Messages []AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessage `json:"messages"`
	Result   AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseSuccess `json:"success"`
	JSON    accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseJSON    `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse]
type accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseError struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseErrorJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseError]
type accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessage struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessageJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessage]
type accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResult struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                                                                                   `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                                                                              `json:"created_at" format:"date-time"`
	CustomPages            AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                                                                                   `json:"is_ui_read_only"`
	LoginDesign  AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration string `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UiReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Must be in the format `300ms` or
	// `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string                                                                                          `json:"user_seat_expiration_inactive_time"`
	JSON                           accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResult]
type accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultJSON struct {
	AuthDomain                     apijson.Field
	AutoRedirectToIdentity         apijson.Field
	CreatedAt                      apijson.Field
	CustomPages                    apijson.Field
	IsUiReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	SessionDuration                apijson.Field
	UiReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                                                                                     `json:"identity_denied"`
	JSON           accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultCustomPagesJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultCustomPagesJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultCustomPages]
type accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                                                                     `json:"text_color"`
	JSON      accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesignJSON `json:"-"`
}

// accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesignJSON
// contains the JSON metadata for the struct
// [AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesign]
type accountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseSuccess bool

const (
	AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseSuccessTrue AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseSuccess = true
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
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration param.Field[string] `json:"session_duration"`
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
	AutoRedirectToIdentity param.Field[bool]                                                                                           `json:"auto_redirect_to_identity"`
	CustomPages            param.Field[AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsCustomPages] `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]                                                                                           `json:"is_ui_read_only"`
	LoginDesign  param.Field[AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign] `json:"login_design"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration param.Field[string] `json:"session_duration"`
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

type AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden param.Field[string] `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied param.Field[string] `json:"identity_denied"`
}

func (r AccountAccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsCustomPages) MarshalJSON() (data []byte, err error) {
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

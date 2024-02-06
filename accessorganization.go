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

// AccessOrganizationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessOrganizationService] method
// instead.
type AccessOrganizationService struct {
	Options []option.RequestOption
}

// NewAccessOrganizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessOrganizationService(opts ...option.RequestOption) (r *AccessOrganizationService) {
	r = &AccessOrganizationService{}
	r.Options = opts
	return
}

// Sets up a Zero Trust organization for your account or zone.
func (r *AccessOrganizationService) ZeroTrustOrganizationNewYourZeroTrustOrganization(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParams, opts ...option.RequestOption) (res *AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the configuration for your Zero Trust organization.
func (r *AccessOrganizationService) ZeroTrustOrganizationGetYourZeroTrustOrganization(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the configuration for your Zero Trust organization.
func (r *AccessOrganizationService) ZeroTrustOrganizationUpdateYourZeroTrustOrganization(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams, opts ...option.RequestOption) (res *AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse struct {
	Errors   []AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseError   `json:"errors"`
	Messages []AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessage `json:"messages"`
	Result   AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseSuccess `json:"success"`
	JSON    accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseJSON    `json:"-"`
}

// accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse]
type accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseError struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseErrorJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseError]
type accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessage struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessageJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessage]
type accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResult struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                                                                         `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                                                                    `json:"created_at" format:"date-time"`
	CustomPages            AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                                                                         `json:"is_ui_read_only"`
	LoginDesign  AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesign `json:"login_design"`
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
	UserSeatExpirationInactiveTime string `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration string                                                                                `json:"warp_auth_session_duration"`
	JSON                    accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResult]
type accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultJSON struct {
	AllowAuthenticateViaWarp       apijson.Field
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
	WarpAuthSessionDuration        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                                                                           `json:"identity_denied"`
	JSON           accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultCustomPagesJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultCustomPagesJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultCustomPages]
type accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                                                           `json:"text_color"`
	JSON      accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesignJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesignJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesign]
type accessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseResultLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseSuccess bool

const (
	AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseSuccessTrue AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationResponseSuccess = true
)

type AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse struct {
	Errors   []AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseError   `json:"errors"`
	Messages []AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessage `json:"messages"`
	Result   AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseSuccess `json:"success"`
	JSON    accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseJSON    `json:"-"`
}

// accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse]
type accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseError struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseErrorJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseError]
type accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessage struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessageJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessage]
type accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResult struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                                                                         `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                                                                    `json:"created_at" format:"date-time"`
	CustomPages            AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                                                                         `json:"is_ui_read_only"`
	LoginDesign  AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesign `json:"login_design"`
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
	UserSeatExpirationInactiveTime string `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration string                                                                                `json:"warp_auth_session_duration"`
	JSON                    accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResult]
type accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultJSON struct {
	AllowAuthenticateViaWarp       apijson.Field
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
	WarpAuthSessionDuration        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                                                                           `json:"identity_denied"`
	JSON           accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultCustomPagesJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultCustomPagesJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultCustomPages]
type accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                                                           `json:"text_color"`
	JSON      accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesignJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesignJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesign]
type accessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseResultLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseSuccess bool

const (
	AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseSuccessTrue AccessOrganizationZeroTrustOrganizationGetYourZeroTrustOrganizationResponseSuccess = true
)

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse struct {
	Errors   []AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseError   `json:"errors"`
	Messages []AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessage `json:"messages"`
	Result   AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseSuccess `json:"success"`
	JSON    accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseJSON    `json:"-"`
}

// accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse]
type accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseError struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseErrorJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseError]
type accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessage struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessageJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessage]
type accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResult struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                                                                                            `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time                                                                                       `json:"created_at" format:"date-time"`
	CustomPages            AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly bool                                                                                            `json:"is_ui_read_only"`
	LoginDesign  AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesign `json:"login_design"`
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
	UserSeatExpirationInactiveTime string `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration string                                                                                   `json:"warp_auth_session_duration"`
	JSON                    accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResult]
type accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultJSON struct {
	AllowAuthenticateViaWarp       apijson.Field
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
	WarpAuthSessionDuration        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                                                                                              `json:"identity_denied"`
	JSON           accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultCustomPagesJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultCustomPagesJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultCustomPages]
type accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string                                                                                              `json:"text_color"`
	JSON      accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesignJSON `json:"-"`
}

// accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesignJSON
// contains the JSON metadata for the struct
// [AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesign]
type accessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseResultLoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseSuccess bool

const (
	AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseSuccessTrue AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationResponseSuccess = true
)

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParams struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain,required"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name,required"`
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]                                                                                 `json:"is_ui_read_only"`
	LoginDesign  param.Field[AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign] `json:"login_design"`
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
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration param.Field[string] `json:"warp_auth_session_duration"`
}

func (r AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign struct {
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

func (r AccessOrganizationZeroTrustOrganizationNewYourZeroTrustOrganizationParamsLoginDesign) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity param.Field[bool]                                                                                    `json:"auto_redirect_to_identity"`
	CustomPages            param.Field[AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsCustomPages] `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUiReadOnly param.Field[bool]                                                                                    `json:"is_ui_read_only"`
	LoginDesign  param.Field[AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign] `json:"login_design"`
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
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration param.Field[string] `json:"warp_auth_session_duration"`
}

func (r AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden param.Field[string] `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied param.Field[string] `json:"identity_denied"`
}

func (r AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsCustomPages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign struct {
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

func (r AccessOrganizationZeroTrustOrganizationUpdateYourZeroTrustOrganizationParamsLoginDesign) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

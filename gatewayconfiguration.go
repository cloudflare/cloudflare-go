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

// GatewayConfigurationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayConfigurationService]
// method instead.
type GatewayConfigurationService struct {
	Options []option.RequestOption
}

// NewGatewayConfigurationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGatewayConfigurationService(opts ...option.RequestOption) (r *GatewayConfigurationService) {
	r = &GatewayConfigurationService{}
	r.Options = opts
	return
}

// Fetches the current Zero Trust account configuration.
func (r *GatewayConfigurationService) ZeroTrustAccountsGetZeroTrustAccountConfiguration(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/configuration", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patches the current Zero Trust account configuration. This endpoint can update a
// single subcollection of settings such as `antivirus`, `tls_decrypt`,
// `activity_log`, `block_page`, `browser_isolation`, `fips`, `body_scanning`, or
// `custom_certificate`, without updating the entire configuration object. Returns
// an error if any collection of settings is not properly configured.
func (r *GatewayConfigurationService) ZeroTrustAccountsPatchZeroTrustAccountConfiguration(ctx context.Context, accountID interface{}, body GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParams, opts ...option.RequestOption) (res *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/configuration", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the current Zero Trust account configuration.
func (r *GatewayConfigurationService) ZeroTrustAccountsUpdateZeroTrustAccountConfiguration(ctx context.Context, accountID interface{}, body GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParams, opts ...option.RequestOption) (res *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/configuration", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// account settings.
type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettings `json:"settings"`
	UpdatedAt time.Time                                                                             `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseJSON     `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponse]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// account settings.
type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettings struct {
	// Activity log settings.
	ActivityLog GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsCustomCertificate `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatching `json:"extended_email_matching"`
	// FIPS settings.
	Fips GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsTLSDecrypt `json:"tls_decrypt"`
	JSON       gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsJSON       `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettings]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsJSON struct {
	ActivityLog           apijson.Field
	Antivirus             apijson.Field
	BlockPage             apijson.Field
	BodyScanning          apijson.Field
	BrowserIsolation      apijson.Field
	CustomCertificate     apijson.Field
	ExtendedEmailMatching apijson.Field
	Fips                  apijson.Field
	ProtocolDetection     apijson.Field
	TLSDecrypt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Activity log settings.
type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                                                                                 `json:"enabled"`
	JSON    gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsActivityLogJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsActivityLogJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsActivityLog]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Anti-virus settings.
type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettings `json:"notification_settings"`
	JSON                 gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsAntivirusJSON                 `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsAntivirusJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsAntivirus]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                                                                 `json:"support_url"`
	JSON       gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettingsJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettingsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettings]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Block page layout settings.
type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBlockPage struct {
	// Block page background color in #rrggbb format.
	BackgroundColor string `json:"background_color"`
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Enabled bool `json:"enabled"`
	// Block page footer text.
	FooterText string `json:"footer_text"`
	// Block page header text.
	HeaderText string `json:"header_text"`
	// Full URL to the logo file.
	LogoPath string `json:"logo_path"`
	// Admin email for users to contact.
	MailtoAddress string `json:"mailto_address"`
	// Subject line for emails created from block page.
	MailtoSubject string `json:"mailto_subject"`
	// Block page title.
	Name string `json:"name"`
	// Suppress detailed info at the bottom of the block page.
	SuppressFooter bool                                                                                               `json:"suppress_footer"`
	JSON           gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBlockPageJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBlockPageJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBlockPage]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBlockPageJSON struct {
	BackgroundColor apijson.Field
	Enabled         apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	MailtoAddress   apijson.Field
	MailtoSubject   apijson.Field
	Name            apijson.Field
	SuppressFooter  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DLP body scanning settings.
type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                                                                                `json:"inspection_mode"`
	JSON           gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBodyScanningJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBodyScanningJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBodyScanning]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser isolation settings.
type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                                                                      `json:"url_browser_isolation_enabled"`
	JSON                       gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBrowserIsolationJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBrowserIsolationJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBrowserIsolation]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom certificate settings for BYO-PKI.
type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                                                                                     `json:"binding_status"`
	UpdatedAt     time.Time                                                                                                  `json:"updated_at" format:"date-time"`
	JSON          gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsCustomCertificateJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsCustomCertificateJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsCustomCertificate]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Extended e-mail matching settings.
type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool                                                                                                           `json:"enabled"`
	JSON    gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatchingJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatchingJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatching]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatchingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// FIPS settings.
type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS  bool                                                                                          `json:"tls"`
	JSON gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsFipsJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsFipsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsFips]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsFipsJSON struct {
	TLS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol Detection settings.
type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                                                                       `json:"enabled"`
	JSON    gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsProtocolDetectionJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsProtocolDetectionJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsProtocolDetection]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// TLS interception settings.
type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                                                                                `json:"enabled"`
	JSON    gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsTLSDecryptJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsTLSDecryptJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsTLSDecrypt]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsTLSDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSettingsTLSDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// account settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettings `json:"settings"`
	UpdatedAt time.Time                                                                               `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseJSON     `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponse]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// account settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettings struct {
	// Activity log settings.
	ActivityLog GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsCustomCertificate `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatching `json:"extended_email_matching"`
	// FIPS settings.
	Fips GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsTLSDecrypt `json:"tls_decrypt"`
	JSON       gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsJSON       `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettings]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsJSON struct {
	ActivityLog           apijson.Field
	Antivirus             apijson.Field
	BlockPage             apijson.Field
	BodyScanning          apijson.Field
	BrowserIsolation      apijson.Field
	CustomCertificate     apijson.Field
	ExtendedEmailMatching apijson.Field
	Fips                  apijson.Field
	ProtocolDetection     apijson.Field
	TLSDecrypt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Activity log settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                                                                                   `json:"enabled"`
	JSON    gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsActivityLogJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsActivityLogJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsActivityLog]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Anti-virus settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettings `json:"notification_settings"`
	JSON                 gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsAntivirusJSON                 `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsAntivirusJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsAntivirus]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                                                                   `json:"support_url"`
	JSON       gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettingsJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettingsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettings]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Block page layout settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBlockPage struct {
	// Block page background color in #rrggbb format.
	BackgroundColor string `json:"background_color"`
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Enabled bool `json:"enabled"`
	// Block page footer text.
	FooterText string `json:"footer_text"`
	// Block page header text.
	HeaderText string `json:"header_text"`
	// Full URL to the logo file.
	LogoPath string `json:"logo_path"`
	// Admin email for users to contact.
	MailtoAddress string `json:"mailto_address"`
	// Subject line for emails created from block page.
	MailtoSubject string `json:"mailto_subject"`
	// Block page title.
	Name string `json:"name"`
	// Suppress detailed info at the bottom of the block page.
	SuppressFooter bool                                                                                                 `json:"suppress_footer"`
	JSON           gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBlockPageJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBlockPageJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBlockPage]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBlockPageJSON struct {
	BackgroundColor apijson.Field
	Enabled         apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	MailtoAddress   apijson.Field
	MailtoSubject   apijson.Field
	Name            apijson.Field
	SuppressFooter  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DLP body scanning settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                                                                                  `json:"inspection_mode"`
	JSON           gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBodyScanningJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBodyScanningJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBodyScanning]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser isolation settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                                                                        `json:"url_browser_isolation_enabled"`
	JSON                       gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBrowserIsolationJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBrowserIsolationJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBrowserIsolation]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom certificate settings for BYO-PKI.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                                                                                       `json:"binding_status"`
	UpdatedAt     time.Time                                                                                                    `json:"updated_at" format:"date-time"`
	JSON          gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsCustomCertificateJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsCustomCertificateJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsCustomCertificate]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Extended e-mail matching settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool                                                                                                             `json:"enabled"`
	JSON    gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatchingJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatchingJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatching]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatchingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// FIPS settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS  bool                                                                                            `json:"tls"`
	JSON gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsFipsJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsFipsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsFips]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsFipsJSON struct {
	TLS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol Detection settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                                                                         `json:"enabled"`
	JSON    gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsProtocolDetectionJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsProtocolDetectionJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsProtocolDetection]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// TLS interception settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                                                                                  `json:"enabled"`
	JSON    gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsTLSDecryptJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsTLSDecryptJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsTLSDecrypt]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsTLSDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSettingsTLSDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// account settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettings `json:"settings"`
	UpdatedAt time.Time                                                                                `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseJSON     `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponse]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// account settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettings struct {
	// Activity log settings.
	ActivityLog GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsCustomCertificate `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatching `json:"extended_email_matching"`
	// FIPS settings.
	Fips GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsTLSDecrypt `json:"tls_decrypt"`
	JSON       gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsJSON       `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettings]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsJSON struct {
	ActivityLog           apijson.Field
	Antivirus             apijson.Field
	BlockPage             apijson.Field
	BodyScanning          apijson.Field
	BrowserIsolation      apijson.Field
	CustomCertificate     apijson.Field
	ExtendedEmailMatching apijson.Field
	Fips                  apijson.Field
	ProtocolDetection     apijson.Field
	TLSDecrypt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Activity log settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                                                                                    `json:"enabled"`
	JSON    gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsActivityLogJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsActivityLogJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsActivityLog]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Anti-virus settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettings `json:"notification_settings"`
	JSON                 gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsAntivirusJSON                 `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsAntivirusJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsAntivirus]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                                                                    `json:"support_url"`
	JSON       gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettingsJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettingsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettings]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsAntivirusNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Block page layout settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBlockPage struct {
	// Block page background color in #rrggbb format.
	BackgroundColor string `json:"background_color"`
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Enabled bool `json:"enabled"`
	// Block page footer text.
	FooterText string `json:"footer_text"`
	// Block page header text.
	HeaderText string `json:"header_text"`
	// Full URL to the logo file.
	LogoPath string `json:"logo_path"`
	// Admin email for users to contact.
	MailtoAddress string `json:"mailto_address"`
	// Subject line for emails created from block page.
	MailtoSubject string `json:"mailto_subject"`
	// Block page title.
	Name string `json:"name"`
	// Suppress detailed info at the bottom of the block page.
	SuppressFooter bool                                                                                                  `json:"suppress_footer"`
	JSON           gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBlockPageJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBlockPageJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBlockPage]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBlockPageJSON struct {
	BackgroundColor apijson.Field
	Enabled         apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	MailtoAddress   apijson.Field
	MailtoSubject   apijson.Field
	Name            apijson.Field
	SuppressFooter  apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DLP body scanning settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                                                                                   `json:"inspection_mode"`
	JSON           gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBodyScanningJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBodyScanningJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBodyScanning]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser isolation settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                                                                         `json:"url_browser_isolation_enabled"`
	JSON                       gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBrowserIsolationJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBrowserIsolationJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBrowserIsolation]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom certificate settings for BYO-PKI.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                                                                                        `json:"binding_status"`
	UpdatedAt     time.Time                                                                                                     `json:"updated_at" format:"date-time"`
	JSON          gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsCustomCertificateJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsCustomCertificateJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsCustomCertificate]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Extended e-mail matching settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool                                                                                                              `json:"enabled"`
	JSON    gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatchingJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatchingJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatching]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatchingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// FIPS settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS  bool                                                                                             `json:"tls"`
	JSON gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsFipsJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsFipsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsFips]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsFipsJSON struct {
	TLS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol Detection settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                                                                          `json:"enabled"`
	JSON    gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsProtocolDetectionJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsProtocolDetectionJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsProtocolDetection]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// TLS interception settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                                                                                   `json:"enabled"`
	JSON    gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsTLSDecryptJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsTLSDecryptJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsTLSDecrypt]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsTLSDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSettingsTLSDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelope struct {
	Errors   []GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeMessages `json:"messages,required"`
	// account settings.
	Result GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponse `json:"result,required"`
	// Whether the API call was successful
	Success GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeJSON    `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelope]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeErrors]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeMessages]
type gatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeSuccess bool

const (
	GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeSuccessTrue GatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseEnvelopeSuccess = true
)

type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParams struct {
	// account settings.
	Settings param.Field[GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettings] `json:"settings"`
}

func (r GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// account settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettings struct {
	// Activity log settings.
	ActivityLog param.Field[GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsActivityLog] `json:"activity_log"`
	// Anti-virus settings.
	Antivirus param.Field[GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsAntivirus] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBlockPage] `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning param.Field[GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBodyScanning] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBrowserIsolation] `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate param.Field[GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsCustomCertificate] `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching param.Field[GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsExtendedEmailMatching] `json:"extended_email_matching"`
	// FIPS settings.
	Fips param.Field[GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsFips] `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection param.Field[GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsProtocolDetection] `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt param.Field[GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsTLSDecrypt] `json:"tls_decrypt"`
}

func (r GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Activity log settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsActivityLog struct {
	// Enable activity logging.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsActivityLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Anti-virus settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase param.Field[bool] `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed param.Field[bool] `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings param.Field[GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsAntivirusNotificationSettings] `json:"notification_settings"`
}

func (r GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsAntivirus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsAntivirusNotificationSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Block page layout settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBlockPage struct {
	// Block page background color in #rrggbb format.
	BackgroundColor param.Field[string] `json:"background_color"`
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Enabled param.Field[bool] `json:"enabled"`
	// Block page footer text.
	FooterText param.Field[string] `json:"footer_text"`
	// Block page header text.
	HeaderText param.Field[string] `json:"header_text"`
	// Full URL to the logo file.
	LogoPath param.Field[string] `json:"logo_path"`
	// Admin email for users to contact.
	MailtoAddress param.Field[string] `json:"mailto_address"`
	// Subject line for emails created from block page.
	MailtoSubject param.Field[string] `json:"mailto_subject"`
	// Block page title.
	Name param.Field[string] `json:"name"`
	// Suppress detailed info at the bottom of the block page.
	SuppressFooter param.Field[bool] `json:"suppress_footer"`
}

func (r GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBlockPage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLP body scanning settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode param.Field[string] `json:"inspection_mode"`
}

func (r GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBodyScanning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled param.Field[bool] `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBrowserIsolation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom certificate settings for BYO-PKI.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID param.Field[string] `json:"id"`
}

func (r GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsCustomCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Extended e-mail matching settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsExtendedEmailMatching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// FIPS settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS param.Field[bool] `json:"tls"`
}

func (r GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsFips) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Protocol Detection settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsProtocolDetection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TLS interception settings.
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsTLSDecrypt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelope struct {
	Errors   []GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeMessages `json:"messages,required"`
	// account settings.
	Result GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponse `json:"result,required"`
	// Whether the API call was successful
	Success GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeJSON    `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelope]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeErrors]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeMessages]
type gatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeSuccess bool

const (
	GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeSuccessTrue GatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseEnvelopeSuccess = true
)

type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParams struct {
	// account settings.
	Settings param.Field[GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettings] `json:"settings"`
}

func (r GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// account settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettings struct {
	// Activity log settings.
	ActivityLog param.Field[GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsActivityLog] `json:"activity_log"`
	// Anti-virus settings.
	Antivirus param.Field[GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsAntivirus] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBlockPage] `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning param.Field[GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBodyScanning] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBrowserIsolation] `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate param.Field[GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsCustomCertificate] `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching param.Field[GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsExtendedEmailMatching] `json:"extended_email_matching"`
	// FIPS settings.
	Fips param.Field[GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsFips] `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection param.Field[GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsProtocolDetection] `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt param.Field[GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsTLSDecrypt] `json:"tls_decrypt"`
}

func (r GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Activity log settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsActivityLog struct {
	// Enable activity logging.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsActivityLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Anti-virus settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase param.Field[bool] `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed param.Field[bool] `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings param.Field[GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsAntivirusNotificationSettings] `json:"notification_settings"`
}

func (r GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsAntivirus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsAntivirusNotificationSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Block page layout settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBlockPage struct {
	// Block page background color in #rrggbb format.
	BackgroundColor param.Field[string] `json:"background_color"`
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Enabled param.Field[bool] `json:"enabled"`
	// Block page footer text.
	FooterText param.Field[string] `json:"footer_text"`
	// Block page header text.
	HeaderText param.Field[string] `json:"header_text"`
	// Full URL to the logo file.
	LogoPath param.Field[string] `json:"logo_path"`
	// Admin email for users to contact.
	MailtoAddress param.Field[string] `json:"mailto_address"`
	// Subject line for emails created from block page.
	MailtoSubject param.Field[string] `json:"mailto_subject"`
	// Block page title.
	Name param.Field[string] `json:"name"`
	// Suppress detailed info at the bottom of the block page.
	SuppressFooter param.Field[bool] `json:"suppress_footer"`
}

func (r GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBlockPage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLP body scanning settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode param.Field[string] `json:"inspection_mode"`
}

func (r GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBodyScanning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled param.Field[bool] `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBrowserIsolation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom certificate settings for BYO-PKI.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID param.Field[string] `json:"id"`
}

func (r GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsCustomCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Extended e-mail matching settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsExtendedEmailMatching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// FIPS settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS param.Field[bool] `json:"tls"`
}

func (r GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsFips) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Protocol Detection settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsProtocolDetection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TLS interception settings.
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsTLSDecrypt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelope struct {
	Errors   []GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeMessages `json:"messages,required"`
	// account settings.
	Result GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponse `json:"result,required"`
	// Whether the API call was successful
	Success GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeJSON    `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelope]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                              `json:"code,required"`
	Message string                                                                                             `json:"message,required"`
	JSON    gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeErrors]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeMessages]
type gatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeSuccess bool

const (
	GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeSuccessTrue GatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseEnvelopeSuccess = true
)

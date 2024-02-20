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

// Patches the current Zero Trust account configuration. This endpoint can update a
// single subcollection of settings such as `antivirus`, `tls_decrypt`,
// `activity_log`, `block_page`, `browser_isolation`, `fips`, `body_scanning`, or
// `custom_certificate`, without updating the entire configuration object. Returns
// an error if any collection of settings is not properly configured.
func (r *GatewayConfigurationService) Update(ctx context.Context, accountID interface{}, body GatewayConfigurationUpdateParams, opts ...option.RequestOption) (res *GatewayConfigurationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayConfigurationUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/configuration", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the current Zero Trust account configuration.
func (r *GatewayConfigurationService) Get(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *GatewayConfigurationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayConfigurationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/configuration", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the current Zero Trust account configuration.
func (r *GatewayConfigurationService) Replace(ctx context.Context, accountID interface{}, body GatewayConfigurationReplaceParams, opts ...option.RequestOption) (res *GatewayConfigurationReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayConfigurationReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/configuration", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// account settings.
type GatewayConfigurationUpdateResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  GatewayConfigurationUpdateResponseSettings `json:"settings"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationUpdateResponseJSON     `json:"-"`
}

// gatewayConfigurationUpdateResponseJSON contains the JSON metadata for the struct
// [GatewayConfigurationUpdateResponse]
type gatewayConfigurationUpdateResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// account settings.
type GatewayConfigurationUpdateResponseSettings struct {
	// Activity log settings.
	ActivityLog GatewayConfigurationUpdateResponseSettingsActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus GatewayConfigurationUpdateResponseSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage GatewayConfigurationUpdateResponseSettingsBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning GatewayConfigurationUpdateResponseSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation GatewayConfigurationUpdateResponseSettingsBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate GatewayConfigurationUpdateResponseSettingsCustomCertificate `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching GatewayConfigurationUpdateResponseSettingsExtendedEmailMatching `json:"extended_email_matching"`
	// FIPS settings.
	Fips GatewayConfigurationUpdateResponseSettingsFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection GatewayConfigurationUpdateResponseSettingsProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt GatewayConfigurationUpdateResponseSettingsTLSDecrypt `json:"tls_decrypt"`
	JSON       gatewayConfigurationUpdateResponseSettingsJSON       `json:"-"`
}

// gatewayConfigurationUpdateResponseSettingsJSON contains the JSON metadata for
// the struct [GatewayConfigurationUpdateResponseSettings]
type gatewayConfigurationUpdateResponseSettingsJSON struct {
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

func (r *GatewayConfigurationUpdateResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Activity log settings.
type GatewayConfigurationUpdateResponseSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                                      `json:"enabled"`
	JSON    gatewayConfigurationUpdateResponseSettingsActivityLogJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseSettingsActivityLogJSON contains the JSON
// metadata for the struct [GatewayConfigurationUpdateResponseSettingsActivityLog]
type gatewayConfigurationUpdateResponseSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Anti-virus settings.
type GatewayConfigurationUpdateResponseSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings GatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettings `json:"notification_settings"`
	JSON                 gatewayConfigurationUpdateResponseSettingsAntivirusJSON                 `json:"-"`
}

// gatewayConfigurationUpdateResponseSettingsAntivirusJSON contains the JSON
// metadata for the struct [GatewayConfigurationUpdateResponseSettingsAntivirus]
type gatewayConfigurationUpdateResponseSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                      `json:"support_url"`
	JSON       gatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettingsJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettingsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettings]
type gatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Block page layout settings.
type GatewayConfigurationUpdateResponseSettingsBlockPage struct {
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
	SuppressFooter bool                                                    `json:"suppress_footer"`
	JSON           gatewayConfigurationUpdateResponseSettingsBlockPageJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseSettingsBlockPageJSON contains the JSON
// metadata for the struct [GatewayConfigurationUpdateResponseSettingsBlockPage]
type gatewayConfigurationUpdateResponseSettingsBlockPageJSON struct {
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

func (r *GatewayConfigurationUpdateResponseSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DLP body scanning settings.
type GatewayConfigurationUpdateResponseSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                                     `json:"inspection_mode"`
	JSON           gatewayConfigurationUpdateResponseSettingsBodyScanningJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseSettingsBodyScanningJSON contains the JSON
// metadata for the struct [GatewayConfigurationUpdateResponseSettingsBodyScanning]
type gatewayConfigurationUpdateResponseSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser isolation settings.
type GatewayConfigurationUpdateResponseSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                           `json:"url_browser_isolation_enabled"`
	JSON                       gatewayConfigurationUpdateResponseSettingsBrowserIsolationJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseSettingsBrowserIsolationJSON contains the JSON
// metadata for the struct
// [GatewayConfigurationUpdateResponseSettingsBrowserIsolation]
type gatewayConfigurationUpdateResponseSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom certificate settings for BYO-PKI.
type GatewayConfigurationUpdateResponseSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                                          `json:"binding_status"`
	UpdatedAt     time.Time                                                       `json:"updated_at" format:"date-time"`
	JSON          gatewayConfigurationUpdateResponseSettingsCustomCertificateJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseSettingsCustomCertificateJSON contains the
// JSON metadata for the struct
// [GatewayConfigurationUpdateResponseSettingsCustomCertificate]
type gatewayConfigurationUpdateResponseSettingsCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseSettingsCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Extended e-mail matching settings.
type GatewayConfigurationUpdateResponseSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool                                                                `json:"enabled"`
	JSON    gatewayConfigurationUpdateResponseSettingsExtendedEmailMatchingJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseSettingsExtendedEmailMatchingJSON contains the
// JSON metadata for the struct
// [GatewayConfigurationUpdateResponseSettingsExtendedEmailMatching]
type gatewayConfigurationUpdateResponseSettingsExtendedEmailMatchingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseSettingsExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// FIPS settings.
type GatewayConfigurationUpdateResponseSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS  bool                                               `json:"tls"`
	JSON gatewayConfigurationUpdateResponseSettingsFipsJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseSettingsFipsJSON contains the JSON metadata
// for the struct [GatewayConfigurationUpdateResponseSettingsFips]
type gatewayConfigurationUpdateResponseSettingsFipsJSON struct {
	TLS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol Detection settings.
type GatewayConfigurationUpdateResponseSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                            `json:"enabled"`
	JSON    gatewayConfigurationUpdateResponseSettingsProtocolDetectionJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseSettingsProtocolDetectionJSON contains the
// JSON metadata for the struct
// [GatewayConfigurationUpdateResponseSettingsProtocolDetection]
type gatewayConfigurationUpdateResponseSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// TLS interception settings.
type GatewayConfigurationUpdateResponseSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                                     `json:"enabled"`
	JSON    gatewayConfigurationUpdateResponseSettingsTLSDecryptJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseSettingsTLSDecryptJSON contains the JSON
// metadata for the struct [GatewayConfigurationUpdateResponseSettingsTLSDecrypt]
type gatewayConfigurationUpdateResponseSettingsTLSDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseSettingsTLSDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// account settings.
type GatewayConfigurationGetResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  GatewayConfigurationGetResponseSettings `json:"settings"`
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationGetResponseJSON     `json:"-"`
}

// gatewayConfigurationGetResponseJSON contains the JSON metadata for the struct
// [GatewayConfigurationGetResponse]
type gatewayConfigurationGetResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// account settings.
type GatewayConfigurationGetResponseSettings struct {
	// Activity log settings.
	ActivityLog GatewayConfigurationGetResponseSettingsActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus GatewayConfigurationGetResponseSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage GatewayConfigurationGetResponseSettingsBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning GatewayConfigurationGetResponseSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation GatewayConfigurationGetResponseSettingsBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate GatewayConfigurationGetResponseSettingsCustomCertificate `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching GatewayConfigurationGetResponseSettingsExtendedEmailMatching `json:"extended_email_matching"`
	// FIPS settings.
	Fips GatewayConfigurationGetResponseSettingsFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection GatewayConfigurationGetResponseSettingsProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt GatewayConfigurationGetResponseSettingsTLSDecrypt `json:"tls_decrypt"`
	JSON       gatewayConfigurationGetResponseSettingsJSON       `json:"-"`
}

// gatewayConfigurationGetResponseSettingsJSON contains the JSON metadata for the
// struct [GatewayConfigurationGetResponseSettings]
type gatewayConfigurationGetResponseSettingsJSON struct {
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

func (r *GatewayConfigurationGetResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Activity log settings.
type GatewayConfigurationGetResponseSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                                   `json:"enabled"`
	JSON    gatewayConfigurationGetResponseSettingsActivityLogJSON `json:"-"`
}

// gatewayConfigurationGetResponseSettingsActivityLogJSON contains the JSON
// metadata for the struct [GatewayConfigurationGetResponseSettingsActivityLog]
type gatewayConfigurationGetResponseSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Anti-virus settings.
type GatewayConfigurationGetResponseSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings GatewayConfigurationGetResponseSettingsAntivirusNotificationSettings `json:"notification_settings"`
	JSON                 gatewayConfigurationGetResponseSettingsAntivirusJSON                 `json:"-"`
}

// gatewayConfigurationGetResponseSettingsAntivirusJSON contains the JSON metadata
// for the struct [GatewayConfigurationGetResponseSettingsAntivirus]
type gatewayConfigurationGetResponseSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayConfigurationGetResponseSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                   `json:"support_url"`
	JSON       gatewayConfigurationGetResponseSettingsAntivirusNotificationSettingsJSON `json:"-"`
}

// gatewayConfigurationGetResponseSettingsAntivirusNotificationSettingsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationGetResponseSettingsAntivirusNotificationSettings]
type gatewayConfigurationGetResponseSettingsAntivirusNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseSettingsAntivirusNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Block page layout settings.
type GatewayConfigurationGetResponseSettingsBlockPage struct {
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
	SuppressFooter bool                                                 `json:"suppress_footer"`
	JSON           gatewayConfigurationGetResponseSettingsBlockPageJSON `json:"-"`
}

// gatewayConfigurationGetResponseSettingsBlockPageJSON contains the JSON metadata
// for the struct [GatewayConfigurationGetResponseSettingsBlockPage]
type gatewayConfigurationGetResponseSettingsBlockPageJSON struct {
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

func (r *GatewayConfigurationGetResponseSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DLP body scanning settings.
type GatewayConfigurationGetResponseSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                                  `json:"inspection_mode"`
	JSON           gatewayConfigurationGetResponseSettingsBodyScanningJSON `json:"-"`
}

// gatewayConfigurationGetResponseSettingsBodyScanningJSON contains the JSON
// metadata for the struct [GatewayConfigurationGetResponseSettingsBodyScanning]
type gatewayConfigurationGetResponseSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser isolation settings.
type GatewayConfigurationGetResponseSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                        `json:"url_browser_isolation_enabled"`
	JSON                       gatewayConfigurationGetResponseSettingsBrowserIsolationJSON `json:"-"`
}

// gatewayConfigurationGetResponseSettingsBrowserIsolationJSON contains the JSON
// metadata for the struct
// [GatewayConfigurationGetResponseSettingsBrowserIsolation]
type gatewayConfigurationGetResponseSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom certificate settings for BYO-PKI.
type GatewayConfigurationGetResponseSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                                       `json:"binding_status"`
	UpdatedAt     time.Time                                                    `json:"updated_at" format:"date-time"`
	JSON          gatewayConfigurationGetResponseSettingsCustomCertificateJSON `json:"-"`
}

// gatewayConfigurationGetResponseSettingsCustomCertificateJSON contains the JSON
// metadata for the struct
// [GatewayConfigurationGetResponseSettingsCustomCertificate]
type gatewayConfigurationGetResponseSettingsCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseSettingsCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Extended e-mail matching settings.
type GatewayConfigurationGetResponseSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool                                                             `json:"enabled"`
	JSON    gatewayConfigurationGetResponseSettingsExtendedEmailMatchingJSON `json:"-"`
}

// gatewayConfigurationGetResponseSettingsExtendedEmailMatchingJSON contains the
// JSON metadata for the struct
// [GatewayConfigurationGetResponseSettingsExtendedEmailMatching]
type gatewayConfigurationGetResponseSettingsExtendedEmailMatchingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseSettingsExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// FIPS settings.
type GatewayConfigurationGetResponseSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS  bool                                            `json:"tls"`
	JSON gatewayConfigurationGetResponseSettingsFipsJSON `json:"-"`
}

// gatewayConfigurationGetResponseSettingsFipsJSON contains the JSON metadata for
// the struct [GatewayConfigurationGetResponseSettingsFips]
type gatewayConfigurationGetResponseSettingsFipsJSON struct {
	TLS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol Detection settings.
type GatewayConfigurationGetResponseSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                         `json:"enabled"`
	JSON    gatewayConfigurationGetResponseSettingsProtocolDetectionJSON `json:"-"`
}

// gatewayConfigurationGetResponseSettingsProtocolDetectionJSON contains the JSON
// metadata for the struct
// [GatewayConfigurationGetResponseSettingsProtocolDetection]
type gatewayConfigurationGetResponseSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// TLS interception settings.
type GatewayConfigurationGetResponseSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                                  `json:"enabled"`
	JSON    gatewayConfigurationGetResponseSettingsTLSDecryptJSON `json:"-"`
}

// gatewayConfigurationGetResponseSettingsTLSDecryptJSON contains the JSON metadata
// for the struct [GatewayConfigurationGetResponseSettingsTLSDecrypt]
type gatewayConfigurationGetResponseSettingsTLSDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseSettingsTLSDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// account settings.
type GatewayConfigurationReplaceResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  GatewayConfigurationReplaceResponseSettings `json:"settings"`
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationReplaceResponseJSON     `json:"-"`
}

// gatewayConfigurationReplaceResponseJSON contains the JSON metadata for the
// struct [GatewayConfigurationReplaceResponse]
type gatewayConfigurationReplaceResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// account settings.
type GatewayConfigurationReplaceResponseSettings struct {
	// Activity log settings.
	ActivityLog GatewayConfigurationReplaceResponseSettingsActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus GatewayConfigurationReplaceResponseSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage GatewayConfigurationReplaceResponseSettingsBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning GatewayConfigurationReplaceResponseSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation GatewayConfigurationReplaceResponseSettingsBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate GatewayConfigurationReplaceResponseSettingsCustomCertificate `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching GatewayConfigurationReplaceResponseSettingsExtendedEmailMatching `json:"extended_email_matching"`
	// FIPS settings.
	Fips GatewayConfigurationReplaceResponseSettingsFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection GatewayConfigurationReplaceResponseSettingsProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt GatewayConfigurationReplaceResponseSettingsTLSDecrypt `json:"tls_decrypt"`
	JSON       gatewayConfigurationReplaceResponseSettingsJSON       `json:"-"`
}

// gatewayConfigurationReplaceResponseSettingsJSON contains the JSON metadata for
// the struct [GatewayConfigurationReplaceResponseSettings]
type gatewayConfigurationReplaceResponseSettingsJSON struct {
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

func (r *GatewayConfigurationReplaceResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Activity log settings.
type GatewayConfigurationReplaceResponseSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                                       `json:"enabled"`
	JSON    gatewayConfigurationReplaceResponseSettingsActivityLogJSON `json:"-"`
}

// gatewayConfigurationReplaceResponseSettingsActivityLogJSON contains the JSON
// metadata for the struct [GatewayConfigurationReplaceResponseSettingsActivityLog]
type gatewayConfigurationReplaceResponseSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationReplaceResponseSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Anti-virus settings.
type GatewayConfigurationReplaceResponseSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings GatewayConfigurationReplaceResponseSettingsAntivirusNotificationSettings `json:"notification_settings"`
	JSON                 gatewayConfigurationReplaceResponseSettingsAntivirusJSON                 `json:"-"`
}

// gatewayConfigurationReplaceResponseSettingsAntivirusJSON contains the JSON
// metadata for the struct [GatewayConfigurationReplaceResponseSettingsAntivirus]
type gatewayConfigurationReplaceResponseSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *GatewayConfigurationReplaceResponseSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayConfigurationReplaceResponseSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                       `json:"support_url"`
	JSON       gatewayConfigurationReplaceResponseSettingsAntivirusNotificationSettingsJSON `json:"-"`
}

// gatewayConfigurationReplaceResponseSettingsAntivirusNotificationSettingsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationReplaceResponseSettingsAntivirusNotificationSettings]
type gatewayConfigurationReplaceResponseSettingsAntivirusNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationReplaceResponseSettingsAntivirusNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Block page layout settings.
type GatewayConfigurationReplaceResponseSettingsBlockPage struct {
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
	SuppressFooter bool                                                     `json:"suppress_footer"`
	JSON           gatewayConfigurationReplaceResponseSettingsBlockPageJSON `json:"-"`
}

// gatewayConfigurationReplaceResponseSettingsBlockPageJSON contains the JSON
// metadata for the struct [GatewayConfigurationReplaceResponseSettingsBlockPage]
type gatewayConfigurationReplaceResponseSettingsBlockPageJSON struct {
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

func (r *GatewayConfigurationReplaceResponseSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DLP body scanning settings.
type GatewayConfigurationReplaceResponseSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                                      `json:"inspection_mode"`
	JSON           gatewayConfigurationReplaceResponseSettingsBodyScanningJSON `json:"-"`
}

// gatewayConfigurationReplaceResponseSettingsBodyScanningJSON contains the JSON
// metadata for the struct
// [GatewayConfigurationReplaceResponseSettingsBodyScanning]
type gatewayConfigurationReplaceResponseSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayConfigurationReplaceResponseSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser isolation settings.
type GatewayConfigurationReplaceResponseSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                            `json:"url_browser_isolation_enabled"`
	JSON                       gatewayConfigurationReplaceResponseSettingsBrowserIsolationJSON `json:"-"`
}

// gatewayConfigurationReplaceResponseSettingsBrowserIsolationJSON contains the
// JSON metadata for the struct
// [GatewayConfigurationReplaceResponseSettingsBrowserIsolation]
type gatewayConfigurationReplaceResponseSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayConfigurationReplaceResponseSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom certificate settings for BYO-PKI.
type GatewayConfigurationReplaceResponseSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                                           `json:"binding_status"`
	UpdatedAt     time.Time                                                        `json:"updated_at" format:"date-time"`
	JSON          gatewayConfigurationReplaceResponseSettingsCustomCertificateJSON `json:"-"`
}

// gatewayConfigurationReplaceResponseSettingsCustomCertificateJSON contains the
// JSON metadata for the struct
// [GatewayConfigurationReplaceResponseSettingsCustomCertificate]
type gatewayConfigurationReplaceResponseSettingsCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayConfigurationReplaceResponseSettingsCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Extended e-mail matching settings.
type GatewayConfigurationReplaceResponseSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool                                                                 `json:"enabled"`
	JSON    gatewayConfigurationReplaceResponseSettingsExtendedEmailMatchingJSON `json:"-"`
}

// gatewayConfigurationReplaceResponseSettingsExtendedEmailMatchingJSON contains
// the JSON metadata for the struct
// [GatewayConfigurationReplaceResponseSettingsExtendedEmailMatching]
type gatewayConfigurationReplaceResponseSettingsExtendedEmailMatchingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationReplaceResponseSettingsExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// FIPS settings.
type GatewayConfigurationReplaceResponseSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS  bool                                                `json:"tls"`
	JSON gatewayConfigurationReplaceResponseSettingsFipsJSON `json:"-"`
}

// gatewayConfigurationReplaceResponseSettingsFipsJSON contains the JSON metadata
// for the struct [GatewayConfigurationReplaceResponseSettingsFips]
type gatewayConfigurationReplaceResponseSettingsFipsJSON struct {
	TLS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationReplaceResponseSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol Detection settings.
type GatewayConfigurationReplaceResponseSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                             `json:"enabled"`
	JSON    gatewayConfigurationReplaceResponseSettingsProtocolDetectionJSON `json:"-"`
}

// gatewayConfigurationReplaceResponseSettingsProtocolDetectionJSON contains the
// JSON metadata for the struct
// [GatewayConfigurationReplaceResponseSettingsProtocolDetection]
type gatewayConfigurationReplaceResponseSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationReplaceResponseSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// TLS interception settings.
type GatewayConfigurationReplaceResponseSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                                      `json:"enabled"`
	JSON    gatewayConfigurationReplaceResponseSettingsTLSDecryptJSON `json:"-"`
}

// gatewayConfigurationReplaceResponseSettingsTLSDecryptJSON contains the JSON
// metadata for the struct [GatewayConfigurationReplaceResponseSettingsTLSDecrypt]
type gatewayConfigurationReplaceResponseSettingsTLSDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationReplaceResponseSettingsTLSDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayConfigurationUpdateParams struct {
	// account settings.
	Settings param.Field[GatewayConfigurationUpdateParamsSettings] `json:"settings"`
}

func (r GatewayConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// account settings.
type GatewayConfigurationUpdateParamsSettings struct {
	// Activity log settings.
	ActivityLog param.Field[GatewayConfigurationUpdateParamsSettingsActivityLog] `json:"activity_log"`
	// Anti-virus settings.
	Antivirus param.Field[GatewayConfigurationUpdateParamsSettingsAntivirus] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[GatewayConfigurationUpdateParamsSettingsBlockPage] `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning param.Field[GatewayConfigurationUpdateParamsSettingsBodyScanning] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[GatewayConfigurationUpdateParamsSettingsBrowserIsolation] `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate param.Field[GatewayConfigurationUpdateParamsSettingsCustomCertificate] `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching param.Field[GatewayConfigurationUpdateParamsSettingsExtendedEmailMatching] `json:"extended_email_matching"`
	// FIPS settings.
	Fips param.Field[GatewayConfigurationUpdateParamsSettingsFips] `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection param.Field[GatewayConfigurationUpdateParamsSettingsProtocolDetection] `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt param.Field[GatewayConfigurationUpdateParamsSettingsTLSDecrypt] `json:"tls_decrypt"`
}

func (r GatewayConfigurationUpdateParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Activity log settings.
type GatewayConfigurationUpdateParamsSettingsActivityLog struct {
	// Enable activity logging.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationUpdateParamsSettingsActivityLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Anti-virus settings.
type GatewayConfigurationUpdateParamsSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase param.Field[bool] `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed param.Field[bool] `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings param.Field[GatewayConfigurationUpdateParamsSettingsAntivirusNotificationSettings] `json:"notification_settings"`
}

func (r GatewayConfigurationUpdateParamsSettingsAntivirus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayConfigurationUpdateParamsSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r GatewayConfigurationUpdateParamsSettingsAntivirusNotificationSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Block page layout settings.
type GatewayConfigurationUpdateParamsSettingsBlockPage struct {
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

func (r GatewayConfigurationUpdateParamsSettingsBlockPage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLP body scanning settings.
type GatewayConfigurationUpdateParamsSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode param.Field[string] `json:"inspection_mode"`
}

func (r GatewayConfigurationUpdateParamsSettingsBodyScanning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type GatewayConfigurationUpdateParamsSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled param.Field[bool] `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r GatewayConfigurationUpdateParamsSettingsBrowserIsolation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom certificate settings for BYO-PKI.
type GatewayConfigurationUpdateParamsSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID param.Field[string] `json:"id"`
}

func (r GatewayConfigurationUpdateParamsSettingsCustomCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Extended e-mail matching settings.
type GatewayConfigurationUpdateParamsSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationUpdateParamsSettingsExtendedEmailMatching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// FIPS settings.
type GatewayConfigurationUpdateParamsSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS param.Field[bool] `json:"tls"`
}

func (r GatewayConfigurationUpdateParamsSettingsFips) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Protocol Detection settings.
type GatewayConfigurationUpdateParamsSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationUpdateParamsSettingsProtocolDetection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TLS interception settings.
type GatewayConfigurationUpdateParamsSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationUpdateParamsSettingsTLSDecrypt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayConfigurationUpdateResponseEnvelope struct {
	Errors   []GatewayConfigurationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayConfigurationUpdateResponseEnvelopeMessages `json:"messages,required"`
	// account settings.
	Result GatewayConfigurationUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success GatewayConfigurationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayConfigurationUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayConfigurationUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayConfigurationUpdateResponseEnvelope]
type gatewayConfigurationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayConfigurationUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    gatewayConfigurationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [GatewayConfigurationUpdateResponseEnvelopeErrors]
type gatewayConfigurationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayConfigurationUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    gatewayConfigurationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [GatewayConfigurationUpdateResponseEnvelopeMessages]
type gatewayConfigurationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayConfigurationUpdateResponseEnvelopeSuccess bool

const (
	GatewayConfigurationUpdateResponseEnvelopeSuccessTrue GatewayConfigurationUpdateResponseEnvelopeSuccess = true
)

type GatewayConfigurationGetResponseEnvelope struct {
	Errors   []GatewayConfigurationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayConfigurationGetResponseEnvelopeMessages `json:"messages,required"`
	// account settings.
	Result GatewayConfigurationGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success GatewayConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayConfigurationGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayConfigurationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayConfigurationGetResponseEnvelope]
type gatewayConfigurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayConfigurationGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    gatewayConfigurationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayConfigurationGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [GatewayConfigurationGetResponseEnvelopeErrors]
type gatewayConfigurationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayConfigurationGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    gatewayConfigurationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayConfigurationGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [GatewayConfigurationGetResponseEnvelopeMessages]
type gatewayConfigurationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayConfigurationGetResponseEnvelopeSuccess bool

const (
	GatewayConfigurationGetResponseEnvelopeSuccessTrue GatewayConfigurationGetResponseEnvelopeSuccess = true
)

type GatewayConfigurationReplaceParams struct {
	// account settings.
	Settings param.Field[GatewayConfigurationReplaceParamsSettings] `json:"settings"`
}

func (r GatewayConfigurationReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// account settings.
type GatewayConfigurationReplaceParamsSettings struct {
	// Activity log settings.
	ActivityLog param.Field[GatewayConfigurationReplaceParamsSettingsActivityLog] `json:"activity_log"`
	// Anti-virus settings.
	Antivirus param.Field[GatewayConfigurationReplaceParamsSettingsAntivirus] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[GatewayConfigurationReplaceParamsSettingsBlockPage] `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning param.Field[GatewayConfigurationReplaceParamsSettingsBodyScanning] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[GatewayConfigurationReplaceParamsSettingsBrowserIsolation] `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate param.Field[GatewayConfigurationReplaceParamsSettingsCustomCertificate] `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching param.Field[GatewayConfigurationReplaceParamsSettingsExtendedEmailMatching] `json:"extended_email_matching"`
	// FIPS settings.
	Fips param.Field[GatewayConfigurationReplaceParamsSettingsFips] `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection param.Field[GatewayConfigurationReplaceParamsSettingsProtocolDetection] `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt param.Field[GatewayConfigurationReplaceParamsSettingsTLSDecrypt] `json:"tls_decrypt"`
}

func (r GatewayConfigurationReplaceParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Activity log settings.
type GatewayConfigurationReplaceParamsSettingsActivityLog struct {
	// Enable activity logging.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationReplaceParamsSettingsActivityLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Anti-virus settings.
type GatewayConfigurationReplaceParamsSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase param.Field[bool] `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed param.Field[bool] `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings param.Field[GatewayConfigurationReplaceParamsSettingsAntivirusNotificationSettings] `json:"notification_settings"`
}

func (r GatewayConfigurationReplaceParamsSettingsAntivirus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayConfigurationReplaceParamsSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r GatewayConfigurationReplaceParamsSettingsAntivirusNotificationSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Block page layout settings.
type GatewayConfigurationReplaceParamsSettingsBlockPage struct {
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

func (r GatewayConfigurationReplaceParamsSettingsBlockPage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLP body scanning settings.
type GatewayConfigurationReplaceParamsSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode param.Field[string] `json:"inspection_mode"`
}

func (r GatewayConfigurationReplaceParamsSettingsBodyScanning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type GatewayConfigurationReplaceParamsSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled param.Field[bool] `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r GatewayConfigurationReplaceParamsSettingsBrowserIsolation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom certificate settings for BYO-PKI.
type GatewayConfigurationReplaceParamsSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID param.Field[string] `json:"id"`
}

func (r GatewayConfigurationReplaceParamsSettingsCustomCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Extended e-mail matching settings.
type GatewayConfigurationReplaceParamsSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationReplaceParamsSettingsExtendedEmailMatching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// FIPS settings.
type GatewayConfigurationReplaceParamsSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS param.Field[bool] `json:"tls"`
}

func (r GatewayConfigurationReplaceParamsSettingsFips) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Protocol Detection settings.
type GatewayConfigurationReplaceParamsSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationReplaceParamsSettingsProtocolDetection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TLS interception settings.
type GatewayConfigurationReplaceParamsSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationReplaceParamsSettingsTLSDecrypt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayConfigurationReplaceResponseEnvelope struct {
	Errors   []GatewayConfigurationReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayConfigurationReplaceResponseEnvelopeMessages `json:"messages,required"`
	// account settings.
	Result GatewayConfigurationReplaceResponse `json:"result,required"`
	// Whether the API call was successful
	Success GatewayConfigurationReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayConfigurationReplaceResponseEnvelopeJSON    `json:"-"`
}

// gatewayConfigurationReplaceResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayConfigurationReplaceResponseEnvelope]
type gatewayConfigurationReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayConfigurationReplaceResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    gatewayConfigurationReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayConfigurationReplaceResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [GatewayConfigurationReplaceResponseEnvelopeErrors]
type gatewayConfigurationReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayConfigurationReplaceResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    gatewayConfigurationReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayConfigurationReplaceResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [GatewayConfigurationReplaceResponseEnvelopeMessages]
type gatewayConfigurationReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayConfigurationReplaceResponseEnvelopeSuccess bool

const (
	GatewayConfigurationReplaceResponseEnvelopeSuccessTrue GatewayConfigurationReplaceResponseEnvelopeSuccess = true
)

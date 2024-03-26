// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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

// Updates the current Zero Trust account configuration.
func (r *GatewayConfigurationService) Update(ctx context.Context, params GatewayConfigurationUpdateParams, opts ...option.RequestOption) (res *GatewayConfigurationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayConfigurationUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/configuration", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
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
func (r *GatewayConfigurationService) Edit(ctx context.Context, params GatewayConfigurationEditParams, opts ...option.RequestOption) (res *GatewayConfigurationEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayConfigurationEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/configuration", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the current Zero Trust account configuration.
func (r *GatewayConfigurationService) Get(ctx context.Context, query GatewayConfigurationGetParams, opts ...option.RequestOption) (res *GatewayConfigurationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayConfigurationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/configuration", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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

func (r gatewayConfigurationUpdateResponseJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationUpdateResponseSettingsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationUpdateResponseSettingsActivityLogJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationUpdateResponseSettingsAntivirusJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettingsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationUpdateResponseSettingsBlockPageJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationUpdateResponseSettingsBodyScanningJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationUpdateResponseSettingsBrowserIsolationJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationUpdateResponseSettingsCustomCertificateJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationUpdateResponseSettingsExtendedEmailMatchingJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationUpdateResponseSettingsFipsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationUpdateResponseSettingsProtocolDetectionJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationUpdateResponseSettingsTLSDecryptJSON) RawJSON() string {
	return r.raw
}

// account settings.
type GatewayConfigurationEditResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  GatewayConfigurationEditResponseSettings `json:"settings"`
	UpdatedAt time.Time                                `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationEditResponseJSON     `json:"-"`
}

// gatewayConfigurationEditResponseJSON contains the JSON metadata for the struct
// [GatewayConfigurationEditResponse]
type gatewayConfigurationEditResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseJSON) RawJSON() string {
	return r.raw
}

// account settings.
type GatewayConfigurationEditResponseSettings struct {
	// Activity log settings.
	ActivityLog GatewayConfigurationEditResponseSettingsActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus GatewayConfigurationEditResponseSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage GatewayConfigurationEditResponseSettingsBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning GatewayConfigurationEditResponseSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation GatewayConfigurationEditResponseSettingsBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate GatewayConfigurationEditResponseSettingsCustomCertificate `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching GatewayConfigurationEditResponseSettingsExtendedEmailMatching `json:"extended_email_matching"`
	// FIPS settings.
	Fips GatewayConfigurationEditResponseSettingsFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection GatewayConfigurationEditResponseSettingsProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt GatewayConfigurationEditResponseSettingsTLSDecrypt `json:"tls_decrypt"`
	JSON       gatewayConfigurationEditResponseSettingsJSON       `json:"-"`
}

// gatewayConfigurationEditResponseSettingsJSON contains the JSON metadata for the
// struct [GatewayConfigurationEditResponseSettings]
type gatewayConfigurationEditResponseSettingsJSON struct {
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

func (r *GatewayConfigurationEditResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseSettingsJSON) RawJSON() string {
	return r.raw
}

// Activity log settings.
type GatewayConfigurationEditResponseSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                                    `json:"enabled"`
	JSON    gatewayConfigurationEditResponseSettingsActivityLogJSON `json:"-"`
}

// gatewayConfigurationEditResponseSettingsActivityLogJSON contains the JSON
// metadata for the struct [GatewayConfigurationEditResponseSettingsActivityLog]
type gatewayConfigurationEditResponseSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseSettingsActivityLogJSON) RawJSON() string {
	return r.raw
}

// Anti-virus settings.
type GatewayConfigurationEditResponseSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings GatewayConfigurationEditResponseSettingsAntivirusNotificationSettings `json:"notification_settings"`
	JSON                 gatewayConfigurationEditResponseSettingsAntivirusJSON                 `json:"-"`
}

// gatewayConfigurationEditResponseSettingsAntivirusJSON contains the JSON metadata
// for the struct [GatewayConfigurationEditResponseSettingsAntivirus]
type gatewayConfigurationEditResponseSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseSettingsAntivirusJSON) RawJSON() string {
	return r.raw
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayConfigurationEditResponseSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                    `json:"support_url"`
	JSON       gatewayConfigurationEditResponseSettingsAntivirusNotificationSettingsJSON `json:"-"`
}

// gatewayConfigurationEditResponseSettingsAntivirusNotificationSettingsJSON
// contains the JSON metadata for the struct
// [GatewayConfigurationEditResponseSettingsAntivirusNotificationSettings]
type gatewayConfigurationEditResponseSettingsAntivirusNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseSettingsAntivirusNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseSettingsAntivirusNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Block page layout settings.
type GatewayConfigurationEditResponseSettingsBlockPage struct {
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
	SuppressFooter bool                                                  `json:"suppress_footer"`
	JSON           gatewayConfigurationEditResponseSettingsBlockPageJSON `json:"-"`
}

// gatewayConfigurationEditResponseSettingsBlockPageJSON contains the JSON metadata
// for the struct [GatewayConfigurationEditResponseSettingsBlockPage]
type gatewayConfigurationEditResponseSettingsBlockPageJSON struct {
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

func (r *GatewayConfigurationEditResponseSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseSettingsBlockPageJSON) RawJSON() string {
	return r.raw
}

// DLP body scanning settings.
type GatewayConfigurationEditResponseSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                                   `json:"inspection_mode"`
	JSON           gatewayConfigurationEditResponseSettingsBodyScanningJSON `json:"-"`
}

// gatewayConfigurationEditResponseSettingsBodyScanningJSON contains the JSON
// metadata for the struct [GatewayConfigurationEditResponseSettingsBodyScanning]
type gatewayConfigurationEditResponseSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseSettingsBodyScanningJSON) RawJSON() string {
	return r.raw
}

// Browser isolation settings.
type GatewayConfigurationEditResponseSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                         `json:"url_browser_isolation_enabled"`
	JSON                       gatewayConfigurationEditResponseSettingsBrowserIsolationJSON `json:"-"`
}

// gatewayConfigurationEditResponseSettingsBrowserIsolationJSON contains the JSON
// metadata for the struct
// [GatewayConfigurationEditResponseSettingsBrowserIsolation]
type gatewayConfigurationEditResponseSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseSettingsBrowserIsolationJSON) RawJSON() string {
	return r.raw
}

// Custom certificate settings for BYO-PKI.
type GatewayConfigurationEditResponseSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                                        `json:"binding_status"`
	UpdatedAt     time.Time                                                     `json:"updated_at" format:"date-time"`
	JSON          gatewayConfigurationEditResponseSettingsCustomCertificateJSON `json:"-"`
}

// gatewayConfigurationEditResponseSettingsCustomCertificateJSON contains the JSON
// metadata for the struct
// [GatewayConfigurationEditResponseSettingsCustomCertificate]
type gatewayConfigurationEditResponseSettingsCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseSettingsCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseSettingsCustomCertificateJSON) RawJSON() string {
	return r.raw
}

// Extended e-mail matching settings.
type GatewayConfigurationEditResponseSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool                                                              `json:"enabled"`
	JSON    gatewayConfigurationEditResponseSettingsExtendedEmailMatchingJSON `json:"-"`
}

// gatewayConfigurationEditResponseSettingsExtendedEmailMatchingJSON contains the
// JSON metadata for the struct
// [GatewayConfigurationEditResponseSettingsExtendedEmailMatching]
type gatewayConfigurationEditResponseSettingsExtendedEmailMatchingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseSettingsExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseSettingsExtendedEmailMatchingJSON) RawJSON() string {
	return r.raw
}

// FIPS settings.
type GatewayConfigurationEditResponseSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS  bool                                             `json:"tls"`
	JSON gatewayConfigurationEditResponseSettingsFipsJSON `json:"-"`
}

// gatewayConfigurationEditResponseSettingsFipsJSON contains the JSON metadata for
// the struct [GatewayConfigurationEditResponseSettingsFips]
type gatewayConfigurationEditResponseSettingsFipsJSON struct {
	TLS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseSettingsFipsJSON) RawJSON() string {
	return r.raw
}

// Protocol Detection settings.
type GatewayConfigurationEditResponseSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                          `json:"enabled"`
	JSON    gatewayConfigurationEditResponseSettingsProtocolDetectionJSON `json:"-"`
}

// gatewayConfigurationEditResponseSettingsProtocolDetectionJSON contains the JSON
// metadata for the struct
// [GatewayConfigurationEditResponseSettingsProtocolDetection]
type gatewayConfigurationEditResponseSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseSettingsProtocolDetectionJSON) RawJSON() string {
	return r.raw
}

// TLS interception settings.
type GatewayConfigurationEditResponseSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                                   `json:"enabled"`
	JSON    gatewayConfigurationEditResponseSettingsTLSDecryptJSON `json:"-"`
}

// gatewayConfigurationEditResponseSettingsTLSDecryptJSON contains the JSON
// metadata for the struct [GatewayConfigurationEditResponseSettingsTLSDecrypt]
type gatewayConfigurationEditResponseSettingsTLSDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseSettingsTLSDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseSettingsTLSDecryptJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseSettingsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseSettingsActivityLogJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseSettingsAntivirusJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseSettingsAntivirusNotificationSettingsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseSettingsBlockPageJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseSettingsBodyScanningJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseSettingsBrowserIsolationJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseSettingsCustomCertificateJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseSettingsExtendedEmailMatchingJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseSettingsFipsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseSettingsProtocolDetectionJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseSettingsTLSDecryptJSON) RawJSON() string {
	return r.raw
}

type GatewayConfigurationUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
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

func (r gatewayConfigurationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayConfigurationUpdateResponseEnvelopeSuccess bool

const (
	GatewayConfigurationUpdateResponseEnvelopeSuccessTrue GatewayConfigurationUpdateResponseEnvelopeSuccess = true
)

func (r GatewayConfigurationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayConfigurationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayConfigurationEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// account settings.
	Settings param.Field[GatewayConfigurationEditParamsSettings] `json:"settings"`
}

func (r GatewayConfigurationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// account settings.
type GatewayConfigurationEditParamsSettings struct {
	// Activity log settings.
	ActivityLog param.Field[GatewayConfigurationEditParamsSettingsActivityLog] `json:"activity_log"`
	// Anti-virus settings.
	Antivirus param.Field[GatewayConfigurationEditParamsSettingsAntivirus] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[GatewayConfigurationEditParamsSettingsBlockPage] `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning param.Field[GatewayConfigurationEditParamsSettingsBodyScanning] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[GatewayConfigurationEditParamsSettingsBrowserIsolation] `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate param.Field[GatewayConfigurationEditParamsSettingsCustomCertificate] `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching param.Field[GatewayConfigurationEditParamsSettingsExtendedEmailMatching] `json:"extended_email_matching"`
	// FIPS settings.
	Fips param.Field[GatewayConfigurationEditParamsSettingsFips] `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection param.Field[GatewayConfigurationEditParamsSettingsProtocolDetection] `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt param.Field[GatewayConfigurationEditParamsSettingsTLSDecrypt] `json:"tls_decrypt"`
}

func (r GatewayConfigurationEditParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Activity log settings.
type GatewayConfigurationEditParamsSettingsActivityLog struct {
	// Enable activity logging.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationEditParamsSettingsActivityLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Anti-virus settings.
type GatewayConfigurationEditParamsSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase param.Field[bool] `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed param.Field[bool] `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings param.Field[GatewayConfigurationEditParamsSettingsAntivirusNotificationSettings] `json:"notification_settings"`
}

func (r GatewayConfigurationEditParamsSettingsAntivirus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type GatewayConfigurationEditParamsSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r GatewayConfigurationEditParamsSettingsAntivirusNotificationSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Block page layout settings.
type GatewayConfigurationEditParamsSettingsBlockPage struct {
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

func (r GatewayConfigurationEditParamsSettingsBlockPage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLP body scanning settings.
type GatewayConfigurationEditParamsSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode param.Field[string] `json:"inspection_mode"`
}

func (r GatewayConfigurationEditParamsSettingsBodyScanning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type GatewayConfigurationEditParamsSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled param.Field[bool] `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r GatewayConfigurationEditParamsSettingsBrowserIsolation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom certificate settings for BYO-PKI.
type GatewayConfigurationEditParamsSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID param.Field[string] `json:"id"`
}

func (r GatewayConfigurationEditParamsSettingsCustomCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Extended e-mail matching settings.
type GatewayConfigurationEditParamsSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationEditParamsSettingsExtendedEmailMatching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// FIPS settings.
type GatewayConfigurationEditParamsSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS param.Field[bool] `json:"tls"`
}

func (r GatewayConfigurationEditParamsSettingsFips) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Protocol Detection settings.
type GatewayConfigurationEditParamsSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationEditParamsSettingsProtocolDetection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TLS interception settings.
type GatewayConfigurationEditParamsSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationEditParamsSettingsTLSDecrypt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayConfigurationEditResponseEnvelope struct {
	Errors   []GatewayConfigurationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayConfigurationEditResponseEnvelopeMessages `json:"messages,required"`
	// account settings.
	Result GatewayConfigurationEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success GatewayConfigurationEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayConfigurationEditResponseEnvelopeJSON    `json:"-"`
}

// gatewayConfigurationEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayConfigurationEditResponseEnvelope]
type gatewayConfigurationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayConfigurationEditResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    gatewayConfigurationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayConfigurationEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [GatewayConfigurationEditResponseEnvelopeErrors]
type gatewayConfigurationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayConfigurationEditResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    gatewayConfigurationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayConfigurationEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [GatewayConfigurationEditResponseEnvelopeMessages]
type gatewayConfigurationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayConfigurationEditResponseEnvelopeSuccess bool

const (
	GatewayConfigurationEditResponseEnvelopeSuccessTrue GatewayConfigurationEditResponseEnvelopeSuccess = true
)

func (r GatewayConfigurationEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayConfigurationEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayConfigurationGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

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

func (r gatewayConfigurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayConfigurationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayConfigurationGetResponseEnvelopeSuccess bool

const (
	GatewayConfigurationGetResponseEnvelopeSuccessTrue GatewayConfigurationGetResponseEnvelopeSuccess = true
)

func (r GatewayConfigurationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayConfigurationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

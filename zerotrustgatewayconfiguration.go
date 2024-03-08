// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ZeroTrustGatewayConfigurationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustGatewayConfigurationService] method instead.
type ZeroTrustGatewayConfigurationService struct {
	Options []option.RequestOption
}

// NewZeroTrustGatewayConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustGatewayConfigurationService(opts ...option.RequestOption) (r *ZeroTrustGatewayConfigurationService) {
	r = &ZeroTrustGatewayConfigurationService{}
	r.Options = opts
	return
}

// Updates the current Zero Trust account configuration.
func (r *ZeroTrustGatewayConfigurationService) Update(ctx context.Context, params ZeroTrustGatewayConfigurationUpdateParams, opts ...option.RequestOption) (res *ZeroTrustGatewayConfigurationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayConfigurationUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/configuration", params.AccountID)
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
func (r *ZeroTrustGatewayConfigurationService) Edit(ctx context.Context, params ZeroTrustGatewayConfigurationEditParams, opts ...option.RequestOption) (res *ZeroTrustGatewayConfigurationEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayConfigurationEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/configuration", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the current Zero Trust account configuration.
func (r *ZeroTrustGatewayConfigurationService) Get(ctx context.Context, query ZeroTrustGatewayConfigurationGetParams, opts ...option.RequestOption) (res *ZeroTrustGatewayConfigurationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayConfigurationGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/configuration", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// account settings.
type ZeroTrustGatewayConfigurationUpdateResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  ZeroTrustGatewayConfigurationUpdateResponseSettings `json:"settings"`
	UpdatedAt time.Time                                           `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayConfigurationUpdateResponseJSON     `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayConfigurationUpdateResponse]
type zeroTrustGatewayConfigurationUpdateResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// account settings.
type ZeroTrustGatewayConfigurationUpdateResponseSettings struct {
	// Activity log settings.
	ActivityLog ZeroTrustGatewayConfigurationUpdateResponseSettingsActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus ZeroTrustGatewayConfigurationUpdateResponseSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage ZeroTrustGatewayConfigurationUpdateResponseSettingsBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning ZeroTrustGatewayConfigurationUpdateResponseSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation ZeroTrustGatewayConfigurationUpdateResponseSettingsBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate ZeroTrustGatewayConfigurationUpdateResponseSettingsCustomCertificate `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching ZeroTrustGatewayConfigurationUpdateResponseSettingsExtendedEmailMatching `json:"extended_email_matching"`
	// FIPS settings.
	Fips ZeroTrustGatewayConfigurationUpdateResponseSettingsFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection ZeroTrustGatewayConfigurationUpdateResponseSettingsProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt ZeroTrustGatewayConfigurationUpdateResponseSettingsTLSDecrypt `json:"tls_decrypt"`
	JSON       zeroTrustGatewayConfigurationUpdateResponseSettingsJSON       `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseSettingsJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayConfigurationUpdateResponseSettings]
type zeroTrustGatewayConfigurationUpdateResponseSettingsJSON struct {
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

func (r *ZeroTrustGatewayConfigurationUpdateResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseSettingsJSON) RawJSON() string {
	return r.raw
}

// Activity log settings.
type ZeroTrustGatewayConfigurationUpdateResponseSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                                               `json:"enabled"`
	JSON    zeroTrustGatewayConfigurationUpdateResponseSettingsActivityLogJSON `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseSettingsActivityLogJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayConfigurationUpdateResponseSettingsActivityLog]
type zeroTrustGatewayConfigurationUpdateResponseSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationUpdateResponseSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseSettingsActivityLogJSON) RawJSON() string {
	return r.raw
}

// Anti-virus settings.
type ZeroTrustGatewayConfigurationUpdateResponseSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings ZeroTrustGatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettings `json:"notification_settings"`
	JSON                 zeroTrustGatewayConfigurationUpdateResponseSettingsAntivirusJSON                 `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseSettingsAntivirusJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayConfigurationUpdateResponseSettingsAntivirus]
type zeroTrustGatewayConfigurationUpdateResponseSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationUpdateResponseSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseSettingsAntivirusJSON) RawJSON() string {
	return r.raw
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type ZeroTrustGatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                               `json:"support_url"`
	JSON       zeroTrustGatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettingsJSON `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettingsJSON
// contains the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettings]
type zeroTrustGatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseSettingsAntivirusNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Block page layout settings.
type ZeroTrustGatewayConfigurationUpdateResponseSettingsBlockPage struct {
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
	SuppressFooter bool                                                             `json:"suppress_footer"`
	JSON           zeroTrustGatewayConfigurationUpdateResponseSettingsBlockPageJSON `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseSettingsBlockPageJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayConfigurationUpdateResponseSettingsBlockPage]
type zeroTrustGatewayConfigurationUpdateResponseSettingsBlockPageJSON struct {
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

func (r *ZeroTrustGatewayConfigurationUpdateResponseSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseSettingsBlockPageJSON) RawJSON() string {
	return r.raw
}

// DLP body scanning settings.
type ZeroTrustGatewayConfigurationUpdateResponseSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                                              `json:"inspection_mode"`
	JSON           zeroTrustGatewayConfigurationUpdateResponseSettingsBodyScanningJSON `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseSettingsBodyScanningJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayConfigurationUpdateResponseSettingsBodyScanning]
type zeroTrustGatewayConfigurationUpdateResponseSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationUpdateResponseSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseSettingsBodyScanningJSON) RawJSON() string {
	return r.raw
}

// Browser isolation settings.
type ZeroTrustGatewayConfigurationUpdateResponseSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                                    `json:"url_browser_isolation_enabled"`
	JSON                       zeroTrustGatewayConfigurationUpdateResponseSettingsBrowserIsolationJSON `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseSettingsBrowserIsolationJSON contains
// the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationUpdateResponseSettingsBrowserIsolation]
type zeroTrustGatewayConfigurationUpdateResponseSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationUpdateResponseSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseSettingsBrowserIsolationJSON) RawJSON() string {
	return r.raw
}

// Custom certificate settings for BYO-PKI.
type ZeroTrustGatewayConfigurationUpdateResponseSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                                                   `json:"binding_status"`
	UpdatedAt     time.Time                                                                `json:"updated_at" format:"date-time"`
	JSON          zeroTrustGatewayConfigurationUpdateResponseSettingsCustomCertificateJSON `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseSettingsCustomCertificateJSON
// contains the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationUpdateResponseSettingsCustomCertificate]
type zeroTrustGatewayConfigurationUpdateResponseSettingsCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationUpdateResponseSettingsCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseSettingsCustomCertificateJSON) RawJSON() string {
	return r.raw
}

// Extended e-mail matching settings.
type ZeroTrustGatewayConfigurationUpdateResponseSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool                                                                         `json:"enabled"`
	JSON    zeroTrustGatewayConfigurationUpdateResponseSettingsExtendedEmailMatchingJSON `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseSettingsExtendedEmailMatchingJSON
// contains the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationUpdateResponseSettingsExtendedEmailMatching]
type zeroTrustGatewayConfigurationUpdateResponseSettingsExtendedEmailMatchingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationUpdateResponseSettingsExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseSettingsExtendedEmailMatchingJSON) RawJSON() string {
	return r.raw
}

// FIPS settings.
type ZeroTrustGatewayConfigurationUpdateResponseSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS  bool                                                        `json:"tls"`
	JSON zeroTrustGatewayConfigurationUpdateResponseSettingsFipsJSON `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseSettingsFipsJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayConfigurationUpdateResponseSettingsFips]
type zeroTrustGatewayConfigurationUpdateResponseSettingsFipsJSON struct {
	TLS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationUpdateResponseSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseSettingsFipsJSON) RawJSON() string {
	return r.raw
}

// Protocol Detection settings.
type ZeroTrustGatewayConfigurationUpdateResponseSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                                     `json:"enabled"`
	JSON    zeroTrustGatewayConfigurationUpdateResponseSettingsProtocolDetectionJSON `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseSettingsProtocolDetectionJSON
// contains the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationUpdateResponseSettingsProtocolDetection]
type zeroTrustGatewayConfigurationUpdateResponseSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationUpdateResponseSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseSettingsProtocolDetectionJSON) RawJSON() string {
	return r.raw
}

// TLS interception settings.
type ZeroTrustGatewayConfigurationUpdateResponseSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                                              `json:"enabled"`
	JSON    zeroTrustGatewayConfigurationUpdateResponseSettingsTLSDecryptJSON `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseSettingsTLSDecryptJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayConfigurationUpdateResponseSettingsTLSDecrypt]
type zeroTrustGatewayConfigurationUpdateResponseSettingsTLSDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationUpdateResponseSettingsTLSDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseSettingsTLSDecryptJSON) RawJSON() string {
	return r.raw
}

// account settings.
type ZeroTrustGatewayConfigurationEditResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  ZeroTrustGatewayConfigurationEditResponseSettings `json:"settings"`
	UpdatedAt time.Time                                         `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayConfigurationEditResponseJSON     `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayConfigurationEditResponse]
type zeroTrustGatewayConfigurationEditResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseJSON) RawJSON() string {
	return r.raw
}

// account settings.
type ZeroTrustGatewayConfigurationEditResponseSettings struct {
	// Activity log settings.
	ActivityLog ZeroTrustGatewayConfigurationEditResponseSettingsActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus ZeroTrustGatewayConfigurationEditResponseSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage ZeroTrustGatewayConfigurationEditResponseSettingsBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning ZeroTrustGatewayConfigurationEditResponseSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation ZeroTrustGatewayConfigurationEditResponseSettingsBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate ZeroTrustGatewayConfigurationEditResponseSettingsCustomCertificate `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching ZeroTrustGatewayConfigurationEditResponseSettingsExtendedEmailMatching `json:"extended_email_matching"`
	// FIPS settings.
	Fips ZeroTrustGatewayConfigurationEditResponseSettingsFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection ZeroTrustGatewayConfigurationEditResponseSettingsProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt ZeroTrustGatewayConfigurationEditResponseSettingsTLSDecrypt `json:"tls_decrypt"`
	JSON       zeroTrustGatewayConfigurationEditResponseSettingsJSON       `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseSettingsJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayConfigurationEditResponseSettings]
type zeroTrustGatewayConfigurationEditResponseSettingsJSON struct {
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

func (r *ZeroTrustGatewayConfigurationEditResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseSettingsJSON) RawJSON() string {
	return r.raw
}

// Activity log settings.
type ZeroTrustGatewayConfigurationEditResponseSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                                             `json:"enabled"`
	JSON    zeroTrustGatewayConfigurationEditResponseSettingsActivityLogJSON `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseSettingsActivityLogJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayConfigurationEditResponseSettingsActivityLog]
type zeroTrustGatewayConfigurationEditResponseSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationEditResponseSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseSettingsActivityLogJSON) RawJSON() string {
	return r.raw
}

// Anti-virus settings.
type ZeroTrustGatewayConfigurationEditResponseSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings ZeroTrustGatewayConfigurationEditResponseSettingsAntivirusNotificationSettings `json:"notification_settings"`
	JSON                 zeroTrustGatewayConfigurationEditResponseSettingsAntivirusJSON                 `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseSettingsAntivirusJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayConfigurationEditResponseSettingsAntivirus]
type zeroTrustGatewayConfigurationEditResponseSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationEditResponseSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseSettingsAntivirusJSON) RawJSON() string {
	return r.raw
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type ZeroTrustGatewayConfigurationEditResponseSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                             `json:"support_url"`
	JSON       zeroTrustGatewayConfigurationEditResponseSettingsAntivirusNotificationSettingsJSON `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseSettingsAntivirusNotificationSettingsJSON
// contains the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationEditResponseSettingsAntivirusNotificationSettings]
type zeroTrustGatewayConfigurationEditResponseSettingsAntivirusNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationEditResponseSettingsAntivirusNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseSettingsAntivirusNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Block page layout settings.
type ZeroTrustGatewayConfigurationEditResponseSettingsBlockPage struct {
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
	SuppressFooter bool                                                           `json:"suppress_footer"`
	JSON           zeroTrustGatewayConfigurationEditResponseSettingsBlockPageJSON `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseSettingsBlockPageJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayConfigurationEditResponseSettingsBlockPage]
type zeroTrustGatewayConfigurationEditResponseSettingsBlockPageJSON struct {
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

func (r *ZeroTrustGatewayConfigurationEditResponseSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseSettingsBlockPageJSON) RawJSON() string {
	return r.raw
}

// DLP body scanning settings.
type ZeroTrustGatewayConfigurationEditResponseSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                                            `json:"inspection_mode"`
	JSON           zeroTrustGatewayConfigurationEditResponseSettingsBodyScanningJSON `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseSettingsBodyScanningJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayConfigurationEditResponseSettingsBodyScanning]
type zeroTrustGatewayConfigurationEditResponseSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationEditResponseSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseSettingsBodyScanningJSON) RawJSON() string {
	return r.raw
}

// Browser isolation settings.
type ZeroTrustGatewayConfigurationEditResponseSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                                  `json:"url_browser_isolation_enabled"`
	JSON                       zeroTrustGatewayConfigurationEditResponseSettingsBrowserIsolationJSON `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseSettingsBrowserIsolationJSON contains
// the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationEditResponseSettingsBrowserIsolation]
type zeroTrustGatewayConfigurationEditResponseSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationEditResponseSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseSettingsBrowserIsolationJSON) RawJSON() string {
	return r.raw
}

// Custom certificate settings for BYO-PKI.
type ZeroTrustGatewayConfigurationEditResponseSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                                                 `json:"binding_status"`
	UpdatedAt     time.Time                                                              `json:"updated_at" format:"date-time"`
	JSON          zeroTrustGatewayConfigurationEditResponseSettingsCustomCertificateJSON `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseSettingsCustomCertificateJSON contains
// the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationEditResponseSettingsCustomCertificate]
type zeroTrustGatewayConfigurationEditResponseSettingsCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationEditResponseSettingsCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseSettingsCustomCertificateJSON) RawJSON() string {
	return r.raw
}

// Extended e-mail matching settings.
type ZeroTrustGatewayConfigurationEditResponseSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool                                                                       `json:"enabled"`
	JSON    zeroTrustGatewayConfigurationEditResponseSettingsExtendedEmailMatchingJSON `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseSettingsExtendedEmailMatchingJSON
// contains the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationEditResponseSettingsExtendedEmailMatching]
type zeroTrustGatewayConfigurationEditResponseSettingsExtendedEmailMatchingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationEditResponseSettingsExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseSettingsExtendedEmailMatchingJSON) RawJSON() string {
	return r.raw
}

// FIPS settings.
type ZeroTrustGatewayConfigurationEditResponseSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS  bool                                                      `json:"tls"`
	JSON zeroTrustGatewayConfigurationEditResponseSettingsFipsJSON `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseSettingsFipsJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayConfigurationEditResponseSettingsFips]
type zeroTrustGatewayConfigurationEditResponseSettingsFipsJSON struct {
	TLS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationEditResponseSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseSettingsFipsJSON) RawJSON() string {
	return r.raw
}

// Protocol Detection settings.
type ZeroTrustGatewayConfigurationEditResponseSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                                   `json:"enabled"`
	JSON    zeroTrustGatewayConfigurationEditResponseSettingsProtocolDetectionJSON `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseSettingsProtocolDetectionJSON contains
// the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationEditResponseSettingsProtocolDetection]
type zeroTrustGatewayConfigurationEditResponseSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationEditResponseSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseSettingsProtocolDetectionJSON) RawJSON() string {
	return r.raw
}

// TLS interception settings.
type ZeroTrustGatewayConfigurationEditResponseSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                                            `json:"enabled"`
	JSON    zeroTrustGatewayConfigurationEditResponseSettingsTLSDecryptJSON `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseSettingsTLSDecryptJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayConfigurationEditResponseSettingsTLSDecrypt]
type zeroTrustGatewayConfigurationEditResponseSettingsTLSDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationEditResponseSettingsTLSDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseSettingsTLSDecryptJSON) RawJSON() string {
	return r.raw
}

// account settings.
type ZeroTrustGatewayConfigurationGetResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  ZeroTrustGatewayConfigurationGetResponseSettings `json:"settings"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayConfigurationGetResponseJSON     `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayConfigurationGetResponse]
type zeroTrustGatewayConfigurationGetResponseJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseJSON) RawJSON() string {
	return r.raw
}

// account settings.
type ZeroTrustGatewayConfigurationGetResponseSettings struct {
	// Activity log settings.
	ActivityLog ZeroTrustGatewayConfigurationGetResponseSettingsActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus ZeroTrustGatewayConfigurationGetResponseSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage ZeroTrustGatewayConfigurationGetResponseSettingsBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning ZeroTrustGatewayConfigurationGetResponseSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation ZeroTrustGatewayConfigurationGetResponseSettingsBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate ZeroTrustGatewayConfigurationGetResponseSettingsCustomCertificate `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching ZeroTrustGatewayConfigurationGetResponseSettingsExtendedEmailMatching `json:"extended_email_matching"`
	// FIPS settings.
	Fips ZeroTrustGatewayConfigurationGetResponseSettingsFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection ZeroTrustGatewayConfigurationGetResponseSettingsProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt ZeroTrustGatewayConfigurationGetResponseSettingsTLSDecrypt `json:"tls_decrypt"`
	JSON       zeroTrustGatewayConfigurationGetResponseSettingsJSON       `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseSettingsJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayConfigurationGetResponseSettings]
type zeroTrustGatewayConfigurationGetResponseSettingsJSON struct {
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

func (r *ZeroTrustGatewayConfigurationGetResponseSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseSettingsJSON) RawJSON() string {
	return r.raw
}

// Activity log settings.
type ZeroTrustGatewayConfigurationGetResponseSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                                            `json:"enabled"`
	JSON    zeroTrustGatewayConfigurationGetResponseSettingsActivityLogJSON `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseSettingsActivityLogJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayConfigurationGetResponseSettingsActivityLog]
type zeroTrustGatewayConfigurationGetResponseSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationGetResponseSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseSettingsActivityLogJSON) RawJSON() string {
	return r.raw
}

// Anti-virus settings.
type ZeroTrustGatewayConfigurationGetResponseSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings ZeroTrustGatewayConfigurationGetResponseSettingsAntivirusNotificationSettings `json:"notification_settings"`
	JSON                 zeroTrustGatewayConfigurationGetResponseSettingsAntivirusJSON                 `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseSettingsAntivirusJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayConfigurationGetResponseSettingsAntivirus]
type zeroTrustGatewayConfigurationGetResponseSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationGetResponseSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseSettingsAntivirusJSON) RawJSON() string {
	return r.raw
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type ZeroTrustGatewayConfigurationGetResponseSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                            `json:"support_url"`
	JSON       zeroTrustGatewayConfigurationGetResponseSettingsAntivirusNotificationSettingsJSON `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseSettingsAntivirusNotificationSettingsJSON
// contains the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationGetResponseSettingsAntivirusNotificationSettings]
type zeroTrustGatewayConfigurationGetResponseSettingsAntivirusNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationGetResponseSettingsAntivirusNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseSettingsAntivirusNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Block page layout settings.
type ZeroTrustGatewayConfigurationGetResponseSettingsBlockPage struct {
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
	SuppressFooter bool                                                          `json:"suppress_footer"`
	JSON           zeroTrustGatewayConfigurationGetResponseSettingsBlockPageJSON `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseSettingsBlockPageJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayConfigurationGetResponseSettingsBlockPage]
type zeroTrustGatewayConfigurationGetResponseSettingsBlockPageJSON struct {
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

func (r *ZeroTrustGatewayConfigurationGetResponseSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseSettingsBlockPageJSON) RawJSON() string {
	return r.raw
}

// DLP body scanning settings.
type ZeroTrustGatewayConfigurationGetResponseSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                                           `json:"inspection_mode"`
	JSON           zeroTrustGatewayConfigurationGetResponseSettingsBodyScanningJSON `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseSettingsBodyScanningJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayConfigurationGetResponseSettingsBodyScanning]
type zeroTrustGatewayConfigurationGetResponseSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationGetResponseSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseSettingsBodyScanningJSON) RawJSON() string {
	return r.raw
}

// Browser isolation settings.
type ZeroTrustGatewayConfigurationGetResponseSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                                 `json:"url_browser_isolation_enabled"`
	JSON                       zeroTrustGatewayConfigurationGetResponseSettingsBrowserIsolationJSON `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseSettingsBrowserIsolationJSON contains
// the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationGetResponseSettingsBrowserIsolation]
type zeroTrustGatewayConfigurationGetResponseSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationGetResponseSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseSettingsBrowserIsolationJSON) RawJSON() string {
	return r.raw
}

// Custom certificate settings for BYO-PKI.
type ZeroTrustGatewayConfigurationGetResponseSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                                                `json:"binding_status"`
	UpdatedAt     time.Time                                                             `json:"updated_at" format:"date-time"`
	JSON          zeroTrustGatewayConfigurationGetResponseSettingsCustomCertificateJSON `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseSettingsCustomCertificateJSON contains
// the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationGetResponseSettingsCustomCertificate]
type zeroTrustGatewayConfigurationGetResponseSettingsCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationGetResponseSettingsCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseSettingsCustomCertificateJSON) RawJSON() string {
	return r.raw
}

// Extended e-mail matching settings.
type ZeroTrustGatewayConfigurationGetResponseSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool                                                                      `json:"enabled"`
	JSON    zeroTrustGatewayConfigurationGetResponseSettingsExtendedEmailMatchingJSON `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseSettingsExtendedEmailMatchingJSON
// contains the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationGetResponseSettingsExtendedEmailMatching]
type zeroTrustGatewayConfigurationGetResponseSettingsExtendedEmailMatchingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationGetResponseSettingsExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseSettingsExtendedEmailMatchingJSON) RawJSON() string {
	return r.raw
}

// FIPS settings.
type ZeroTrustGatewayConfigurationGetResponseSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS  bool                                                     `json:"tls"`
	JSON zeroTrustGatewayConfigurationGetResponseSettingsFipsJSON `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseSettingsFipsJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayConfigurationGetResponseSettingsFips]
type zeroTrustGatewayConfigurationGetResponseSettingsFipsJSON struct {
	TLS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationGetResponseSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseSettingsFipsJSON) RawJSON() string {
	return r.raw
}

// Protocol Detection settings.
type ZeroTrustGatewayConfigurationGetResponseSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                                  `json:"enabled"`
	JSON    zeroTrustGatewayConfigurationGetResponseSettingsProtocolDetectionJSON `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseSettingsProtocolDetectionJSON contains
// the JSON metadata for the struct
// [ZeroTrustGatewayConfigurationGetResponseSettingsProtocolDetection]
type zeroTrustGatewayConfigurationGetResponseSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationGetResponseSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseSettingsProtocolDetectionJSON) RawJSON() string {
	return r.raw
}

// TLS interception settings.
type ZeroTrustGatewayConfigurationGetResponseSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                                           `json:"enabled"`
	JSON    zeroTrustGatewayConfigurationGetResponseSettingsTLSDecryptJSON `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseSettingsTLSDecryptJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayConfigurationGetResponseSettingsTLSDecrypt]
type zeroTrustGatewayConfigurationGetResponseSettingsTLSDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationGetResponseSettingsTLSDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseSettingsTLSDecryptJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayConfigurationUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// account settings.
	Settings param.Field[ZeroTrustGatewayConfigurationUpdateParamsSettings] `json:"settings"`
}

func (r ZeroTrustGatewayConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// account settings.
type ZeroTrustGatewayConfigurationUpdateParamsSettings struct {
	// Activity log settings.
	ActivityLog param.Field[ZeroTrustGatewayConfigurationUpdateParamsSettingsActivityLog] `json:"activity_log"`
	// Anti-virus settings.
	Antivirus param.Field[ZeroTrustGatewayConfigurationUpdateParamsSettingsAntivirus] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[ZeroTrustGatewayConfigurationUpdateParamsSettingsBlockPage] `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning param.Field[ZeroTrustGatewayConfigurationUpdateParamsSettingsBodyScanning] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[ZeroTrustGatewayConfigurationUpdateParamsSettingsBrowserIsolation] `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate param.Field[ZeroTrustGatewayConfigurationUpdateParamsSettingsCustomCertificate] `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching param.Field[ZeroTrustGatewayConfigurationUpdateParamsSettingsExtendedEmailMatching] `json:"extended_email_matching"`
	// FIPS settings.
	Fips param.Field[ZeroTrustGatewayConfigurationUpdateParamsSettingsFips] `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection param.Field[ZeroTrustGatewayConfigurationUpdateParamsSettingsProtocolDetection] `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt param.Field[ZeroTrustGatewayConfigurationUpdateParamsSettingsTLSDecrypt] `json:"tls_decrypt"`
}

func (r ZeroTrustGatewayConfigurationUpdateParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Activity log settings.
type ZeroTrustGatewayConfigurationUpdateParamsSettingsActivityLog struct {
	// Enable activity logging.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZeroTrustGatewayConfigurationUpdateParamsSettingsActivityLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Anti-virus settings.
type ZeroTrustGatewayConfigurationUpdateParamsSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase param.Field[bool] `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed param.Field[bool] `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings param.Field[ZeroTrustGatewayConfigurationUpdateParamsSettingsAntivirusNotificationSettings] `json:"notification_settings"`
}

func (r ZeroTrustGatewayConfigurationUpdateParamsSettingsAntivirus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type ZeroTrustGatewayConfigurationUpdateParamsSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r ZeroTrustGatewayConfigurationUpdateParamsSettingsAntivirusNotificationSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Block page layout settings.
type ZeroTrustGatewayConfigurationUpdateParamsSettingsBlockPage struct {
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

func (r ZeroTrustGatewayConfigurationUpdateParamsSettingsBlockPage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLP body scanning settings.
type ZeroTrustGatewayConfigurationUpdateParamsSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode param.Field[string] `json:"inspection_mode"`
}

func (r ZeroTrustGatewayConfigurationUpdateParamsSettingsBodyScanning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type ZeroTrustGatewayConfigurationUpdateParamsSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled param.Field[bool] `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r ZeroTrustGatewayConfigurationUpdateParamsSettingsBrowserIsolation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom certificate settings for BYO-PKI.
type ZeroTrustGatewayConfigurationUpdateParamsSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID param.Field[string] `json:"id"`
}

func (r ZeroTrustGatewayConfigurationUpdateParamsSettingsCustomCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Extended e-mail matching settings.
type ZeroTrustGatewayConfigurationUpdateParamsSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZeroTrustGatewayConfigurationUpdateParamsSettingsExtendedEmailMatching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// FIPS settings.
type ZeroTrustGatewayConfigurationUpdateParamsSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS param.Field[bool] `json:"tls"`
}

func (r ZeroTrustGatewayConfigurationUpdateParamsSettingsFips) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Protocol Detection settings.
type ZeroTrustGatewayConfigurationUpdateParamsSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZeroTrustGatewayConfigurationUpdateParamsSettingsProtocolDetection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TLS interception settings.
type ZeroTrustGatewayConfigurationUpdateParamsSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZeroTrustGatewayConfigurationUpdateParamsSettingsTLSDecrypt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayConfigurationUpdateResponseEnvelope struct {
	Errors   []ZeroTrustGatewayConfigurationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayConfigurationUpdateResponseEnvelopeMessages `json:"messages,required"`
	// account settings.
	Result ZeroTrustGatewayConfigurationUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayConfigurationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayConfigurationUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayConfigurationUpdateResponseEnvelope]
type zeroTrustGatewayConfigurationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayConfigurationUpdateResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustGatewayConfigurationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayConfigurationUpdateResponseEnvelopeErrors]
type zeroTrustGatewayConfigurationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayConfigurationUpdateResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustGatewayConfigurationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayConfigurationUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayConfigurationUpdateResponseEnvelopeMessages]
type zeroTrustGatewayConfigurationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustGatewayConfigurationUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayConfigurationUpdateResponseEnvelopeSuccessTrue ZeroTrustGatewayConfigurationUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayConfigurationEditParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// account settings.
	Settings param.Field[ZeroTrustGatewayConfigurationEditParamsSettings] `json:"settings"`
}

func (r ZeroTrustGatewayConfigurationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// account settings.
type ZeroTrustGatewayConfigurationEditParamsSettings struct {
	// Activity log settings.
	ActivityLog param.Field[ZeroTrustGatewayConfigurationEditParamsSettingsActivityLog] `json:"activity_log"`
	// Anti-virus settings.
	Antivirus param.Field[ZeroTrustGatewayConfigurationEditParamsSettingsAntivirus] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[ZeroTrustGatewayConfigurationEditParamsSettingsBlockPage] `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning param.Field[ZeroTrustGatewayConfigurationEditParamsSettingsBodyScanning] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[ZeroTrustGatewayConfigurationEditParamsSettingsBrowserIsolation] `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate param.Field[ZeroTrustGatewayConfigurationEditParamsSettingsCustomCertificate] `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching param.Field[ZeroTrustGatewayConfigurationEditParamsSettingsExtendedEmailMatching] `json:"extended_email_matching"`
	// FIPS settings.
	Fips param.Field[ZeroTrustGatewayConfigurationEditParamsSettingsFips] `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection param.Field[ZeroTrustGatewayConfigurationEditParamsSettingsProtocolDetection] `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt param.Field[ZeroTrustGatewayConfigurationEditParamsSettingsTLSDecrypt] `json:"tls_decrypt"`
}

func (r ZeroTrustGatewayConfigurationEditParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Activity log settings.
type ZeroTrustGatewayConfigurationEditParamsSettingsActivityLog struct {
	// Enable activity logging.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZeroTrustGatewayConfigurationEditParamsSettingsActivityLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Anti-virus settings.
type ZeroTrustGatewayConfigurationEditParamsSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase param.Field[bool] `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed param.Field[bool] `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings param.Field[ZeroTrustGatewayConfigurationEditParamsSettingsAntivirusNotificationSettings] `json:"notification_settings"`
}

func (r ZeroTrustGatewayConfigurationEditParamsSettingsAntivirus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type ZeroTrustGatewayConfigurationEditParamsSettingsAntivirusNotificationSettings struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r ZeroTrustGatewayConfigurationEditParamsSettingsAntivirusNotificationSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Block page layout settings.
type ZeroTrustGatewayConfigurationEditParamsSettingsBlockPage struct {
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

func (r ZeroTrustGatewayConfigurationEditParamsSettingsBlockPage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLP body scanning settings.
type ZeroTrustGatewayConfigurationEditParamsSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode param.Field[string] `json:"inspection_mode"`
}

func (r ZeroTrustGatewayConfigurationEditParamsSettingsBodyScanning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type ZeroTrustGatewayConfigurationEditParamsSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled param.Field[bool] `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r ZeroTrustGatewayConfigurationEditParamsSettingsBrowserIsolation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom certificate settings for BYO-PKI.
type ZeroTrustGatewayConfigurationEditParamsSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID param.Field[string] `json:"id"`
}

func (r ZeroTrustGatewayConfigurationEditParamsSettingsCustomCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Extended e-mail matching settings.
type ZeroTrustGatewayConfigurationEditParamsSettingsExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZeroTrustGatewayConfigurationEditParamsSettingsExtendedEmailMatching) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// FIPS settings.
type ZeroTrustGatewayConfigurationEditParamsSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS param.Field[bool] `json:"tls"`
}

func (r ZeroTrustGatewayConfigurationEditParamsSettingsFips) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Protocol Detection settings.
type ZeroTrustGatewayConfigurationEditParamsSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZeroTrustGatewayConfigurationEditParamsSettingsProtocolDetection) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TLS interception settings.
type ZeroTrustGatewayConfigurationEditParamsSettingsTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZeroTrustGatewayConfigurationEditParamsSettingsTLSDecrypt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayConfigurationEditResponseEnvelope struct {
	Errors   []ZeroTrustGatewayConfigurationEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayConfigurationEditResponseEnvelopeMessages `json:"messages,required"`
	// account settings.
	Result ZeroTrustGatewayConfigurationEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayConfigurationEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayConfigurationEditResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayConfigurationEditResponseEnvelope]
type zeroTrustGatewayConfigurationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayConfigurationEditResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustGatewayConfigurationEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayConfigurationEditResponseEnvelopeErrors]
type zeroTrustGatewayConfigurationEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayConfigurationEditResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustGatewayConfigurationEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayConfigurationEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayConfigurationEditResponseEnvelopeMessages]
type zeroTrustGatewayConfigurationEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustGatewayConfigurationEditResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayConfigurationEditResponseEnvelopeSuccessTrue ZeroTrustGatewayConfigurationEditResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayConfigurationGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayConfigurationGetResponseEnvelope struct {
	Errors   []ZeroTrustGatewayConfigurationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayConfigurationGetResponseEnvelopeMessages `json:"messages,required"`
	// account settings.
	Result ZeroTrustGatewayConfigurationGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayConfigurationGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayConfigurationGetResponseEnvelope]
type zeroTrustGatewayConfigurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayConfigurationGetResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustGatewayConfigurationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayConfigurationGetResponseEnvelopeErrors]
type zeroTrustGatewayConfigurationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayConfigurationGetResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustGatewayConfigurationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayConfigurationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayConfigurationGetResponseEnvelopeMessages]
type zeroTrustGatewayConfigurationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayConfigurationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayConfigurationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustGatewayConfigurationGetResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayConfigurationGetResponseEnvelopeSuccessTrue ZeroTrustGatewayConfigurationGetResponseEnvelopeSuccess = true
)

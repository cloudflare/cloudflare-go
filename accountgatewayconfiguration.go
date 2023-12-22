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

// AccountGatewayConfigurationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountGatewayConfigurationService] method instead.
type AccountGatewayConfigurationService struct {
	Options []option.RequestOption
}

// NewAccountGatewayConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountGatewayConfigurationService(opts ...option.RequestOption) (r *AccountGatewayConfigurationService) {
	r = &AccountGatewayConfigurationService{}
	r.Options = opts
	return
}

// Describes current Zero Trust account configuration.
func (r *AccountGatewayConfigurationService) ZeroTrustAccountsGetZeroTrustAccountConfiguration(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *GatewayAccountConfig, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/configuration", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patches the current Zero Trust account configuration. This end point helps you
// to update single sub collection of settings like antivirus, tls_decrypt,
// activity_log, block_page, browser_isolation, fips, body_scanning or
// custom_certificate, without updating the whole configuration object. Returns
// error if a single collection of setting is not properly configured.
func (r *AccountGatewayConfigurationService) ZeroTrustAccountsPatchZeroTrustAccountConfiguration(ctx context.Context, identifier interface{}, body AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParams, opts ...option.RequestOption) (res *GatewayAccountConfig, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/configuration", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Updates the current Zero Trust account configuration.
func (r *AccountGatewayConfigurationService) ZeroTrustAccountsUpdateZeroTrustAccountConfiguration(ctx context.Context, identifier interface{}, body AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParams, opts ...option.RequestOption) (res *GatewayAccountConfig, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/configuration", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type GatewayAccountConfig struct {
	Errors   []GatewayAccountConfigError   `json:"errors"`
	Messages []GatewayAccountConfigMessage `json:"messages"`
	Result   GatewayAccountConfigResult    `json:"result"`
	// Whether the API call was successful
	Success GatewayAccountConfigSuccess `json:"success"`
	JSON    gatewayAccountConfigJSON    `json:"-"`
}

// gatewayAccountConfigJSON contains the JSON metadata for the struct
// [GatewayAccountConfig]
type gatewayAccountConfigJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAccountConfigError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    gatewayAccountConfigErrorJSON `json:"-"`
}

// gatewayAccountConfigErrorJSON contains the JSON metadata for the struct
// [GatewayAccountConfigError]
type gatewayAccountConfigErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAccountConfigMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    gatewayAccountConfigMessageJSON `json:"-"`
}

// gatewayAccountConfigMessageJSON contains the JSON metadata for the struct
// [GatewayAccountConfigMessage]
type gatewayAccountConfigMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAccountConfigResult struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  GatewayAccountConfigResultSettings `json:"settings"`
	UpdatedAt time.Time                          `json:"updated_at" format:"date-time"`
	JSON      gatewayAccountConfigResultJSON     `json:"-"`
}

// gatewayAccountConfigResultJSON contains the JSON metadata for the struct
// [GatewayAccountConfigResult]
type gatewayAccountConfigResultJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// account settings.
type GatewayAccountConfigResultSettings struct {
	// Activity log settings.
	ActivityLog GatewayAccountConfigResultSettingsActivityLog `json:"activity_log"`
	// Anti virus settings.
	Antivirus GatewayAccountConfigResultSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage GatewayAccountConfigResultSettingsBlockPage `json:"block_page"`
	// DLP body scanning setting
	BodyScanning GatewayAccountConfigResultSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation GatewayAccountConfigResultSettingsBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate GatewayAccountConfigResultSettingsCustomCertificate `json:"custom_certificate"`
	// FIPS settings.
	Fips GatewayAccountConfigResultSettingsFips `json:"fips"`
	// TLS interception settings.
	TlsDecrypt GatewayAccountConfigResultSettingsTlsDecrypt `json:"tls_decrypt"`
	JSON       gatewayAccountConfigResultSettingsJSON       `json:"-"`
}

// gatewayAccountConfigResultSettingsJSON contains the JSON metadata for the struct
// [GatewayAccountConfigResultSettings]
type gatewayAccountConfigResultSettingsJSON struct {
	ActivityLog       apijson.Field
	Antivirus         apijson.Field
	BlockPage         apijson.Field
	BodyScanning      apijson.Field
	BrowserIsolation  apijson.Field
	CustomCertificate apijson.Field
	Fips              apijson.Field
	TlsDecrypt        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Activity log settings.
type GatewayAccountConfigResultSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                              `json:"enabled"`
	JSON    gatewayAccountConfigResultSettingsActivityLogJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsActivityLogJSON contains the JSON metadata for
// the struct [GatewayAccountConfigResultSettingsActivityLog]
type gatewayAccountConfigResultSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Anti virus settings.
type GatewayAccountConfigResultSettingsAntivirus struct {
	// Set to enable antivirus scan on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Set to enable antivirus scan on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool                                            `json:"fail_closed"`
	JSON       gatewayAccountConfigResultSettingsAntivirusJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsAntivirusJSON contains the JSON metadata for
// the struct [GatewayAccountConfigResultSettingsAntivirus]
type gatewayAccountConfigResultSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Block page layout settings.
type GatewayAccountConfigResultSettingsBlockPage struct {
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
	SuppressFooter bool                                            `json:"suppress_footer"`
	JSON           gatewayAccountConfigResultSettingsBlockPageJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsBlockPageJSON contains the JSON metadata for
// the struct [GatewayAccountConfigResultSettingsBlockPage]
type gatewayAccountConfigResultSettingsBlockPageJSON struct {
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

func (r *GatewayAccountConfigResultSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DLP body scanning setting
type GatewayAccountConfigResultSettingsBodyScanning struct {
	// Inspection mode. One of deep or shallow
	InspectionMode string                                             `json:"inspection_mode"`
	JSON           gatewayAccountConfigResultSettingsBodyScanningJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsBodyScanningJSON contains the JSON metadata
// for the struct [GatewayAccountConfigResultSettingsBodyScanning]
type gatewayAccountConfigResultSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser isolation settings.
type GatewayAccountConfigResultSettingsBrowserIsolation struct {
	// Enable Browser Isolation.
	URLBrowserIsolationEnabled bool                                                   `json:"url_browser_isolation_enabled"`
	JSON                       gatewayAccountConfigResultSettingsBrowserIsolationJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsBrowserIsolationJSON contains the JSON
// metadata for the struct [GatewayAccountConfigResultSettingsBrowserIsolation]
type gatewayAccountConfigResultSettingsBrowserIsolationJSON struct {
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom certificate settings for BYO-PKI.
type GatewayAccountConfigResultSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store)
	ID string `json:"id"`
	// Certificate status (internal)
	BindingStatus string                                                  `json:"binding_status"`
	UpdatedAt     time.Time                                               `json:"updated_at" format:"date-time"`
	JSON          gatewayAccountConfigResultSettingsCustomCertificateJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsCustomCertificateJSON contains the JSON
// metadata for the struct [GatewayAccountConfigResultSettingsCustomCertificate]
type gatewayAccountConfigResultSettingsCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// FIPS settings.
type GatewayAccountConfigResultSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Tls  bool                                       `json:"tls"`
	JSON gatewayAccountConfigResultSettingsFipsJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsFipsJSON contains the JSON metadata for the
// struct [GatewayAccountConfigResultSettingsFips]
type gatewayAccountConfigResultSettingsFipsJSON struct {
	Tls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// TLS interception settings.
type GatewayAccountConfigResultSettingsTlsDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                             `json:"enabled"`
	JSON    gatewayAccountConfigResultSettingsTlsDecryptJSON `json:"-"`
}

// gatewayAccountConfigResultSettingsTlsDecryptJSON contains the JSON metadata for
// the struct [GatewayAccountConfigResultSettingsTlsDecrypt]
type gatewayAccountConfigResultSettingsTlsDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountConfigResultSettingsTlsDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayAccountConfigSuccess bool

const (
	GatewayAccountConfigSuccessTrue GatewayAccountConfigSuccess = true
)

type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParams struct {
	// account settings.
	Settings param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettings] `json:"settings"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// account settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettings struct {
	// Activity log settings.
	ActivityLog param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsActivityLog] `json:"activity_log"`
	// Anti virus settings.
	Antivirus param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsAntivirus] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBlockPage] `json:"block_page"`
	// DLP body scanning setting
	BodyScanning param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBodyScanning] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBrowserIsolation] `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsCustomCertificate] `json:"custom_certificate"`
	// FIPS settings.
	Fips param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsFips] `json:"fips"`
	// TLS interception settings.
	TlsDecrypt param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsTlsDecrypt] `json:"tls_decrypt"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Activity log settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsActivityLog struct {
	// Enable activity logging.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsActivityLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Anti virus settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsAntivirus struct {
	// Set to enable antivirus scan on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Set to enable antivirus scan on uploads.
	EnabledUploadPhase param.Field[bool] `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed param.Field[bool] `json:"fail_closed"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsAntivirus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Block page layout settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBlockPage struct {
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

func (r AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBlockPage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLP body scanning setting
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBodyScanning struct {
	// Inspection mode. One of deep or shallow
	InspectionMode param.Field[string] `json:"inspection_mode"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBodyScanning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBrowserIsolation struct {
	// Enable Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBrowserIsolation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom certificate settings for BYO-PKI.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store)
	ID param.Field[string] `json:"id"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsCustomCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// FIPS settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Tls param.Field[bool] `json:"tls"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsFips) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TLS interception settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsTlsDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsTlsDecrypt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParams struct {
	// account settings.
	Settings param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettings] `json:"settings"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// account settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettings struct {
	// Activity log settings.
	ActivityLog param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsActivityLog] `json:"activity_log"`
	// Anti virus settings.
	Antivirus param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsAntivirus] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBlockPage] `json:"block_page"`
	// DLP body scanning setting
	BodyScanning param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBodyScanning] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBrowserIsolation] `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsCustomCertificate] `json:"custom_certificate"`
	// FIPS settings.
	Fips param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsFips] `json:"fips"`
	// TLS interception settings.
	TlsDecrypt param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsTlsDecrypt] `json:"tls_decrypt"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Activity log settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsActivityLog struct {
	// Enable activity logging.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsActivityLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Anti virus settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsAntivirus struct {
	// Set to enable antivirus scan on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Set to enable antivirus scan on uploads.
	EnabledUploadPhase param.Field[bool] `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed param.Field[bool] `json:"fail_closed"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsAntivirus) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Block page layout settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBlockPage struct {
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

func (r AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBlockPage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLP body scanning setting
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBodyScanning struct {
	// Inspection mode. One of deep or shallow
	InspectionMode param.Field[string] `json:"inspection_mode"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBodyScanning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBrowserIsolation struct {
	// Enable Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBrowserIsolation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom certificate settings for BYO-PKI.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store)
	ID param.Field[string] `json:"id"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsCustomCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// FIPS settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Tls param.Field[bool] `json:"tls"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsFips) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TLS interception settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsTlsDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsTlsDecrypt) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

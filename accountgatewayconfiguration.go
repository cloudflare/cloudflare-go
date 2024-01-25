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

// Fetches the current Zero Trust account configuration.
func (r *AccountGatewayConfigurationService) ZeroTrustAccountsGetZeroTrustAccountConfiguration(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/configuration", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Patches the current Zero Trust account configuration. This endpoint can update a
// single subcollection of settings such as `antivirus`, `tls_decrypt`,
// `activity_log`, `block_page`, `browser_isolation`, `fips`, `body_scanning`, or
// `custom_certificate`, without updating the entire configuration object. Returns
// an error if any collection of settings is not properly configured.
func (r *AccountGatewayConfigurationService) ZeroTrustAccountsPatchZeroTrustAccountConfiguration(ctx context.Context, identifier interface{}, body AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParams, opts ...option.RequestOption) (res *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/configuration", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Updates the current Zero Trust account configuration.
func (r *AccountGatewayConfigurationService) ZeroTrustAccountsUpdateZeroTrustAccountConfiguration(ctx context.Context, identifier interface{}, body AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParams, opts ...option.RequestOption) (res *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/configuration", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponse struct {
	Errors   []AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseError   `json:"errors"`
	Messages []AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseMessage `json:"messages"`
	Result   AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSuccess `json:"success"`
	JSON    accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseJSON    `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponse]
type accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseError struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseErrorJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseError]
type accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseMessage struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseMessageJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseMessage]
type accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResult struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettings `json:"settings"`
	UpdatedAt time.Time                                                                                          `json:"updated_at" format:"date-time"`
	JSON      accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultJSON     `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResult]
type accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// account settings.
type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettings struct {
	// Activity log settings.
	ActivityLog AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsCustomCertificate `json:"custom_certificate"`
	// FIPS settings.
	Fips AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TlsDecrypt AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsTlsDecrypt `json:"tls_decrypt"`
	JSON       accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsJSON       `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettings]
type accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsJSON struct {
	ActivityLog       apijson.Field
	Antivirus         apijson.Field
	BlockPage         apijson.Field
	BodyScanning      apijson.Field
	BrowserIsolation  apijson.Field
	CustomCertificate apijson.Field
	Fips              apijson.Field
	ProtocolDetection apijson.Field
	TlsDecrypt        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Activity log settings.
type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                                                                                              `json:"enabled"`
	JSON    accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsActivityLogJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsActivityLogJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsActivityLog]
type accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Anti-virus settings.
type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool                                                                                                            `json:"fail_closed"`
	JSON       accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsAntivirusJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsAntivirusJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsAntivirus]
type accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Block page layout settings.
type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBlockPage struct {
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
	SuppressFooter bool                                                                                                            `json:"suppress_footer"`
	JSON           accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBlockPageJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBlockPageJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBlockPage]
type accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBlockPageJSON struct {
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

func (r *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DLP body scanning settings.
type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                                                                                             `json:"inspection_mode"`
	JSON           accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBodyScanningJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBodyScanningJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBodyScanning]
type accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser isolation settings.
type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                                                                                   `json:"url_browser_isolation_enabled"`
	JSON                       accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolationJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolationJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolation]
type accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom certificate settings for BYO-PKI.
type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                                                                                                  `json:"binding_status"`
	UpdatedAt     time.Time                                                                                                               `json:"updated_at" format:"date-time"`
	JSON          accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsCustomCertificateJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsCustomCertificateJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsCustomCertificate]
type accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// FIPS settings.
type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Tls  bool                                                                                                       `json:"tls"`
	JSON accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsFipsJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsFipsJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsFips]
type accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsFipsJSON struct {
	Tls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol Detection settings.
type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                                                                                    `json:"enabled"`
	JSON    accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsProtocolDetectionJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsProtocolDetectionJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsProtocolDetection]
type accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// TLS interception settings.
type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsTlsDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                                                                                             `json:"enabled"`
	JSON    accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsTlsDecryptJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsTlsDecryptJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsTlsDecrypt]
type accountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsTlsDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseResultSettingsTlsDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSuccess bool

const (
	AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSuccessTrue AccountGatewayConfigurationZeroTrustAccountsGetZeroTrustAccountConfigurationResponseSuccess = true
)

type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponse struct {
	Errors   []AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseError   `json:"errors"`
	Messages []AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseMessage `json:"messages"`
	Result   AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSuccess `json:"success"`
	JSON    accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseJSON    `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponse]
type accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseError struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseErrorJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseError]
type accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseMessage struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseMessageJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseMessage]
type accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResult struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettings `json:"settings"`
	UpdatedAt time.Time                                                                                            `json:"updated_at" format:"date-time"`
	JSON      accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultJSON     `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResult]
type accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// account settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettings struct {
	// Activity log settings.
	ActivityLog AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsCustomCertificate `json:"custom_certificate"`
	// FIPS settings.
	Fips AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TlsDecrypt AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsTlsDecrypt `json:"tls_decrypt"`
	JSON       accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsJSON       `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettings]
type accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsJSON struct {
	ActivityLog       apijson.Field
	Antivirus         apijson.Field
	BlockPage         apijson.Field
	BodyScanning      apijson.Field
	BrowserIsolation  apijson.Field
	CustomCertificate apijson.Field
	Fips              apijson.Field
	ProtocolDetection apijson.Field
	TlsDecrypt        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Activity log settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                                                                                                `json:"enabled"`
	JSON    accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsActivityLogJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsActivityLogJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsActivityLog]
type accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Anti-virus settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool                                                                                                              `json:"fail_closed"`
	JSON       accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsAntivirusJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsAntivirusJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsAntivirus]
type accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Block page layout settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBlockPage struct {
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
	SuppressFooter bool                                                                                                              `json:"suppress_footer"`
	JSON           accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBlockPageJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBlockPageJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBlockPage]
type accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBlockPageJSON struct {
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

func (r *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DLP body scanning settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                                                                                               `json:"inspection_mode"`
	JSON           accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBodyScanningJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBodyScanningJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBodyScanning]
type accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser isolation settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                                                                                     `json:"url_browser_isolation_enabled"`
	JSON                       accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolationJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolationJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolation]
type accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom certificate settings for BYO-PKI.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                                                                                                    `json:"binding_status"`
	UpdatedAt     time.Time                                                                                                                 `json:"updated_at" format:"date-time"`
	JSON          accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsCustomCertificateJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsCustomCertificateJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsCustomCertificate]
type accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// FIPS settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Tls  bool                                                                                                         `json:"tls"`
	JSON accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsFipsJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsFipsJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsFips]
type accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsFipsJSON struct {
	Tls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol Detection settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                                                                                      `json:"enabled"`
	JSON    accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsProtocolDetectionJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsProtocolDetectionJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsProtocolDetection]
type accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// TLS interception settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsTlsDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                                                                                               `json:"enabled"`
	JSON    accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsTlsDecryptJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsTlsDecryptJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsTlsDecrypt]
type accountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsTlsDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseResultSettingsTlsDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSuccess bool

const (
	AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSuccessTrue AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationResponseSuccess = true
)

type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponse struct {
	Errors   []AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseError   `json:"errors"`
	Messages []AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseMessage `json:"messages"`
	Result   AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSuccess `json:"success"`
	JSON    accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseJSON    `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponse]
type accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseError struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseErrorJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseError]
type accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseMessage struct {
	Code    int64                                                                                              `json:"code,required"`
	Message string                                                                                             `json:"message,required"`
	JSON    accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseMessageJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseMessage]
type accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResult struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettings `json:"settings"`
	UpdatedAt time.Time                                                                                             `json:"updated_at" format:"date-time"`
	JSON      accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultJSON     `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResult]
type accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultJSON struct {
	CreatedAt   apijson.Field
	Settings    apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// account settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettings struct {
	// Activity log settings.
	ActivityLog AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsCustomCertificate `json:"custom_certificate"`
	// FIPS settings.
	Fips AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TlsDecrypt AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsTlsDecrypt `json:"tls_decrypt"`
	JSON       accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsJSON       `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettings]
type accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsJSON struct {
	ActivityLog       apijson.Field
	Antivirus         apijson.Field
	BlockPage         apijson.Field
	BodyScanning      apijson.Field
	BrowserIsolation  apijson.Field
	CustomCertificate apijson.Field
	Fips              apijson.Field
	ProtocolDetection apijson.Field
	TlsDecrypt        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Activity log settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsActivityLog struct {
	// Enable activity logging.
	Enabled bool                                                                                                                 `json:"enabled"`
	JSON    accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsActivityLogJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsActivityLogJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsActivityLog]
type accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Anti-virus settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool                                                                                                               `json:"fail_closed"`
	JSON       accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsAntivirusJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsAntivirusJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsAntivirus]
type accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Block page layout settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBlockPage struct {
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
	SuppressFooter bool                                                                                                               `json:"suppress_footer"`
	JSON           accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBlockPageJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBlockPageJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBlockPage]
type accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBlockPageJSON struct {
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

func (r *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DLP body scanning settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                                                                                                `json:"inspection_mode"`
	JSON           accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBodyScanningJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBodyScanningJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBodyScanning]
type accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Browser isolation settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                                                                                      `json:"url_browser_isolation_enabled"`
	JSON                       accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolationJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolationJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolation]
type accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Custom certificate settings for BYO-PKI.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                                                                                                     `json:"binding_status"`
	UpdatedAt     time.Time                                                                                                                  `json:"updated_at" format:"date-time"`
	JSON          accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsCustomCertificateJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsCustomCertificateJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsCustomCertificate]
type accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// FIPS settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	Tls  bool                                                                                                          `json:"tls"`
	JSON accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsFipsJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsFipsJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsFips]
type accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsFipsJSON struct {
	Tls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol Detection settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                                                                                       `json:"enabled"`
	JSON    accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsProtocolDetectionJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsProtocolDetectionJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsProtocolDetection]
type accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// TLS interception settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsTlsDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                                                                                                `json:"enabled"`
	JSON    accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsTlsDecryptJSON `json:"-"`
}

// accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsTlsDecryptJSON
// contains the JSON metadata for the struct
// [AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsTlsDecrypt]
type accountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsTlsDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseResultSettingsTlsDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSuccess bool

const (
	AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSuccessTrue AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationResponseSuccess = true
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
	// Anti-virus settings.
	Antivirus param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsAntivirus] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBlockPage] `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBodyScanning] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBrowserIsolation] `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsCustomCertificate] `json:"custom_certificate"`
	// FIPS settings.
	Fips param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsFips] `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection param.Field[AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsProtocolDetection] `json:"protocol_detection"`
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

// Anti-virus settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
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

// DLP body scanning settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode param.Field[string] `json:"inspection_mode"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBodyScanning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled param.Field[bool] `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsBrowserIsolation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom certificate settings for BYO-PKI.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
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

// Protocol Detection settings.
type AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsPatchZeroTrustAccountConfigurationParamsSettingsProtocolDetection) MarshalJSON() (data []byte, err error) {
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
	// Anti-virus settings.
	Antivirus param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsAntivirus] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBlockPage] `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBodyScanning] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBrowserIsolation] `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsCustomCertificate] `json:"custom_certificate"`
	// FIPS settings.
	Fips param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsFips] `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection param.Field[AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsProtocolDetection] `json:"protocol_detection"`
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

// Anti-virus settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
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

// DLP body scanning settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode param.Field[string] `json:"inspection_mode"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBodyScanning) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled param.Field[bool] `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsBrowserIsolation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom certificate settings for BYO-PKI.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
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

// Protocol Detection settings.
type AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountGatewayConfigurationZeroTrustAccountsUpdateZeroTrustAccountConfigurationParamsSettingsProtocolDetection) MarshalJSON() (data []byte, err error) {
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

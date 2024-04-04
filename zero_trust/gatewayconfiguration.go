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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
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
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182e struct {
	// Activity log settings.
	ActivityLog UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eActivityLog `json:"activity_log"`
	// Anti-virus settings.
	Antivirus UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirus `json:"antivirus"`
	// Block page layout settings.
	BlockPage UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBlockPage `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBodyScanning `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBrowserIsolation `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eCustomCertificate `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eExtendedEmailMatching `json:"extended_email_matching"`
	// FIPS settings.
	Fips UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eFips `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eProtocolDetection `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eTLSDecrypt `json:"tls_decrypt"`
	JSON       unnamedSchemaRef055aaf3918bf29f81c09d394a864182eJSON       `json:"-"`
}

// unnamedSchemaRef055aaf3918bf29f81c09d394a864182eJSON contains the JSON metadata
// for the struct [UnnamedSchemaRef055aaf3918bf29f81c09d394a864182e]
type unnamedSchemaRef055aaf3918bf29f81c09d394a864182eJSON struct {
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

func (r *UnnamedSchemaRef055aaf3918bf29f81c09d394a864182e) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef055aaf3918bf29f81c09d394a864182eJSON) RawJSON() string {
	return r.raw
}

// Activity log settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eActivityLog struct {
	// Enable activity logging.
	Enabled bool                                                            `json:"enabled"`
	JSON    unnamedSchemaRef055aaf3918bf29f81c09d394a864182eActivityLogJSON `json:"-"`
}

// unnamedSchemaRef055aaf3918bf29f81c09d394a864182eActivityLogJSON contains the
// JSON metadata for the struct
// [UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eActivityLog]
type unnamedSchemaRef055aaf3918bf29f81c09d394a864182eActivityLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eActivityLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef055aaf3918bf29f81c09d394a864182eActivityLogJSON) RawJSON() string {
	return r.raw
}

// Anti-virus settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirus struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed bool `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusNotificationSettings `json:"notification_settings"`
	JSON                 unnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusJSON                 `json:"-"`
}

// unnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusJSON contains the JSON
// metadata for the struct
// [UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirus]
type unnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusJSON) RawJSON() string {
	return r.raw
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                                            `json:"support_url"`
	JSON       unnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusNotificationSettingsJSON `json:"-"`
}

// unnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusNotificationSettingsJSON
// contains the JSON metadata for the struct
// [UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusNotificationSettings]
type unnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Block page layout settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBlockPage struct {
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
	JSON           unnamedSchemaRef055aaf3918bf29f81c09d394a864182eBlockPageJSON `json:"-"`
}

// unnamedSchemaRef055aaf3918bf29f81c09d394a864182eBlockPageJSON contains the JSON
// metadata for the struct
// [UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBlockPage]
type unnamedSchemaRef055aaf3918bf29f81c09d394a864182eBlockPageJSON struct {
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

func (r *UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef055aaf3918bf29f81c09d394a864182eBlockPageJSON) RawJSON() string {
	return r.raw
}

// DLP body scanning settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBodyScanning struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode string                                                           `json:"inspection_mode"`
	JSON           unnamedSchemaRef055aaf3918bf29f81c09d394a864182eBodyScanningJSON `json:"-"`
}

// unnamedSchemaRef055aaf3918bf29f81c09d394a864182eBodyScanningJSON contains the
// JSON metadata for the struct
// [UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBodyScanning]
type unnamedSchemaRef055aaf3918bf29f81c09d394a864182eBodyScanningJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBodyScanning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef055aaf3918bf29f81c09d394a864182eBodyScanningJSON) RawJSON() string {
	return r.raw
}

// Browser isolation settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBrowserIsolation struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                                                                 `json:"url_browser_isolation_enabled"`
	JSON                       unnamedSchemaRef055aaf3918bf29f81c09d394a864182eBrowserIsolationJSON `json:"-"`
}

// unnamedSchemaRef055aaf3918bf29f81c09d394a864182eBrowserIsolationJSON contains
// the JSON metadata for the struct
// [UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBrowserIsolation]
type unnamedSchemaRef055aaf3918bf29f81c09d394a864182eBrowserIsolationJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBrowserIsolation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef055aaf3918bf29f81c09d394a864182eBrowserIsolationJSON) RawJSON() string {
	return r.raw
}

// Custom certificate settings for BYO-PKI.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eCustomCertificate struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled bool `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Certificate status (internal).
	BindingStatus string                                                                `json:"binding_status"`
	UpdatedAt     time.Time                                                             `json:"updated_at" format:"date-time"`
	JSON          unnamedSchemaRef055aaf3918bf29f81c09d394a864182eCustomCertificateJSON `json:"-"`
}

// unnamedSchemaRef055aaf3918bf29f81c09d394a864182eCustomCertificateJSON contains
// the JSON metadata for the struct
// [UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eCustomCertificate]
type unnamedSchemaRef055aaf3918bf29f81c09d394a864182eCustomCertificateJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eCustomCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef055aaf3918bf29f81c09d394a864182eCustomCertificateJSON) RawJSON() string {
	return r.raw
}

// Extended e-mail matching settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eExtendedEmailMatching struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled bool                                                                      `json:"enabled"`
	JSON    unnamedSchemaRef055aaf3918bf29f81c09d394a864182eExtendedEmailMatchingJSON `json:"-"`
}

// unnamedSchemaRef055aaf3918bf29f81c09d394a864182eExtendedEmailMatchingJSON
// contains the JSON metadata for the struct
// [UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eExtendedEmailMatching]
type unnamedSchemaRef055aaf3918bf29f81c09d394a864182eExtendedEmailMatchingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef055aaf3918bf29f81c09d394a864182eExtendedEmailMatchingJSON) RawJSON() string {
	return r.raw
}

// FIPS settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eFips struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS  bool                                                     `json:"tls"`
	JSON unnamedSchemaRef055aaf3918bf29f81c09d394a864182eFipsJSON `json:"-"`
}

// unnamedSchemaRef055aaf3918bf29f81c09d394a864182eFipsJSON contains the JSON
// metadata for the struct [UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eFips]
type unnamedSchemaRef055aaf3918bf29f81c09d394a864182eFipsJSON struct {
	TLS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eFips) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef055aaf3918bf29f81c09d394a864182eFipsJSON) RawJSON() string {
	return r.raw
}

// Protocol Detection settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eProtocolDetection struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled bool                                                                  `json:"enabled"`
	JSON    unnamedSchemaRef055aaf3918bf29f81c09d394a864182eProtocolDetectionJSON `json:"-"`
}

// unnamedSchemaRef055aaf3918bf29f81c09d394a864182eProtocolDetectionJSON contains
// the JSON metadata for the struct
// [UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eProtocolDetection]
type unnamedSchemaRef055aaf3918bf29f81c09d394a864182eProtocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef055aaf3918bf29f81c09d394a864182eProtocolDetectionJSON) RawJSON() string {
	return r.raw
}

// TLS interception settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eTLSDecrypt struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled bool                                                           `json:"enabled"`
	JSON    unnamedSchemaRef055aaf3918bf29f81c09d394a864182eTLSDecryptJSON `json:"-"`
}

// unnamedSchemaRef055aaf3918bf29f81c09d394a864182eTLSDecryptJSON contains the JSON
// metadata for the struct
// [UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eTLSDecrypt]
type unnamedSchemaRef055aaf3918bf29f81c09d394a864182eTLSDecryptJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eTLSDecrypt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef055aaf3918bf29f81c09d394a864182eTLSDecryptJSON) RawJSON() string {
	return r.raw
}

// account settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eParam struct {
	// Activity log settings.
	ActivityLog param.Field[UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eActivityLogParam] `json:"activity_log"`
	// Anti-virus settings.
	Antivirus param.Field[UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusParam] `json:"antivirus"`
	// Block page layout settings.
	BlockPage param.Field[UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBlockPageParam] `json:"block_page"`
	// DLP body scanning settings.
	BodyScanning param.Field[UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBodyScanningParam] `json:"body_scanning"`
	// Browser isolation settings.
	BrowserIsolation param.Field[UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBrowserIsolationParam] `json:"browser_isolation"`
	// Custom certificate settings for BYO-PKI.
	CustomCertificate param.Field[UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eCustomCertificateParam] `json:"custom_certificate"`
	// Extended e-mail matching settings.
	ExtendedEmailMatching param.Field[UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eExtendedEmailMatchingParam] `json:"extended_email_matching"`
	// FIPS settings.
	Fips param.Field[UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eFipsParam] `json:"fips"`
	// Protocol Detection settings.
	ProtocolDetection param.Field[UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eProtocolDetectionParam] `json:"protocol_detection"`
	// TLS interception settings.
	TLSDecrypt param.Field[UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eTLSDecryptParam] `json:"tls_decrypt"`
}

func (r UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Activity log settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eActivityLogParam struct {
	// Enable activity logging.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eActivityLogParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Anti-virus settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusParam struct {
	// Enable anti-virus scanning on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Enable anti-virus scanning on uploads.
	EnabledUploadPhase param.Field[bool] `json:"enabled_upload_phase"`
	// Block requests for files that cannot be scanned.
	FailClosed param.Field[bool] `json:"fail_closed"`
	// Configure a message to display on the user's device when an antivirus search is
	// performed.
	NotificationSettings param.Field[UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusNotificationSettingsParam] `json:"notification_settings"`
}

func (r UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a message to display on the user's device when an antivirus search is
// performed.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusNotificationSettingsParam struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eAntivirusNotificationSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Block page layout settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBlockPageParam struct {
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

func (r UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBlockPageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLP body scanning settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBodyScanningParam struct {
	// Set the inspection mode to either `deep` or `shallow`.
	InspectionMode param.Field[string] `json:"inspection_mode"`
}

func (r UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBodyScanningParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Browser isolation settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBrowserIsolationParam struct {
	// Enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled param.Field[bool] `json:"non_identity_enabled"`
	// Enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eBrowserIsolationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom certificate settings for BYO-PKI.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eCustomCertificateParam struct {
	// Enable use of custom certificate authority for signing Gateway traffic.
	Enabled param.Field[bool] `json:"enabled,required"`
	// UUID of certificate (ID from MTLS certificate store).
	ID param.Field[string] `json:"id"`
}

func (r UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eCustomCertificateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Extended e-mail matching settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eExtendedEmailMatchingParam struct {
	// Enable matching all variants of user emails (with + or . modifiers) used as
	// criteria in Firewall policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eExtendedEmailMatchingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// FIPS settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eFipsParam struct {
	// Enable only cipher suites and TLS versions compliant with FIPS 140-2.
	TLS param.Field[bool] `json:"tls"`
}

func (r UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eFipsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Protocol Detection settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eProtocolDetectionParam struct {
	// Enable detecting protocol on initial bytes of client traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eProtocolDetectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TLS interception settings.
type UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eTLSDecryptParam struct {
	// Enable inspecting encrypted HTTP traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eTLSDecryptParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// account settings.
type GatewayConfigurationUpdateResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  UnnamedSchemaRef055aaf3918bf29f81c09d394a864182e `json:"settings"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationUpdateResponseJSON           `json:"-"`
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
type GatewayConfigurationEditResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  UnnamedSchemaRef055aaf3918bf29f81c09d394a864182e `json:"settings"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationEditResponseJSON             `json:"-"`
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
type GatewayConfigurationGetResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// account settings.
	Settings  UnnamedSchemaRef055aaf3918bf29f81c09d394a864182e `json:"settings"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationGetResponseJSON              `json:"-"`
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

type GatewayConfigurationUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// account settings.
	Settings param.Field[UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eParam] `json:"settings"`
}

func (r GatewayConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayConfigurationUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
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
	Settings param.Field[UnnamedSchemaRef055aaf3918bf29f81c09d394a864182eParam] `json:"settings"`
}

func (r GatewayConfigurationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayConfigurationEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
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
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
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

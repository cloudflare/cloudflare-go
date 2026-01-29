// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// GatewayConfigurationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayConfigurationService] method instead.
type GatewayConfigurationService struct {
	Options           []option.RequestOption
	CustomCertificate *GatewayConfigurationCustomCertificateService
}

// NewGatewayConfigurationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGatewayConfigurationService(opts ...option.RequestOption) (r *GatewayConfigurationService) {
	r = &GatewayConfigurationService{}
	r.Options = opts
	r.CustomCertificate = NewGatewayConfigurationCustomCertificateService(opts...)
	return
}

// Update the current Zero Trust account configuration.
func (r *GatewayConfigurationService) Update(ctx context.Context, params GatewayConfigurationUpdateParams, opts ...option.RequestOption) (res *GatewayConfigurationUpdateResponse, err error) {
	var env GatewayConfigurationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/configuration", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update (PATCH) a single subcollection of settings such as `antivirus`,
// `tls_decrypt`, `activity_log`, `block_page`, `browser_isolation`, `fips`,
// `body_scanning`, or `certificate` without updating the entire configuration
// object. This endpoint returns an error if any settings collection lacks proper
// configuration.
func (r *GatewayConfigurationService) Edit(ctx context.Context, params GatewayConfigurationEditParams, opts ...option.RequestOption) (res *GatewayConfigurationEditResponse, err error) {
	var env GatewayConfigurationEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/configuration", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve the current Zero Trust account configuration.
func (r *GatewayConfigurationService) Get(ctx context.Context, query GatewayConfigurationGetParams, opts ...option.RequestOption) (res *GatewayConfigurationGetResponse, err error) {
	var env GatewayConfigurationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/configuration", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Specify activity log settings.
type ActivityLogSettings struct {
	// Specify whether to log activity.
	Enabled bool                    `json:"enabled,nullable"`
	JSON    activityLogSettingsJSON `json:"-"`
}

// activityLogSettingsJSON contains the JSON metadata for the struct
// [ActivityLogSettings]
type activityLogSettingsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActivityLogSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r activityLogSettingsJSON) RawJSON() string {
	return r.raw
}

// Specify activity log settings.
type ActivityLogSettingsParam struct {
	// Specify whether to log activity.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ActivityLogSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify anti-virus settings.
type AntiVirusSettings struct {
	// Specify whether to enable anti-virus scanning on downloads.
	EnabledDownloadPhase bool `json:"enabled_download_phase,nullable"`
	// Specify whether to enable anti-virus scanning on uploads.
	EnabledUploadPhase bool `json:"enabled_upload_phase,nullable"`
	// Specify whether to block requests for unscannable files.
	FailClosed bool `json:"fail_closed,nullable"`
	// Configure the message the user's device shows during an antivirus scan.
	NotificationSettings NotificationSettings  `json:"notification_settings,nullable"`
	JSON                 antiVirusSettingsJSON `json:"-"`
}

// antiVirusSettingsJSON contains the JSON metadata for the struct
// [AntiVirusSettings]
type antiVirusSettingsJSON struct {
	EnabledDownloadPhase apijson.Field
	EnabledUploadPhase   apijson.Field
	FailClosed           apijson.Field
	NotificationSettings apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *AntiVirusSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r antiVirusSettingsJSON) RawJSON() string {
	return r.raw
}

// Specify anti-virus settings.
type AntiVirusSettingsParam struct {
	// Specify whether to enable anti-virus scanning on downloads.
	EnabledDownloadPhase param.Field[bool] `json:"enabled_download_phase"`
	// Specify whether to enable anti-virus scanning on uploads.
	EnabledUploadPhase param.Field[bool] `json:"enabled_upload_phase"`
	// Specify whether to block requests for unscannable files.
	FailClosed param.Field[bool] `json:"fail_closed"`
	// Configure the message the user's device shows during an antivirus scan.
	NotificationSettings param.Field[NotificationSettingsParam] `json:"notification_settings"`
}

func (r AntiVirusSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify block page layout settings.
type BlockPageSettings struct {
	// Specify the block page background color in `#rrggbb` format when the mode is
	// customized_block_page.
	BackgroundColor string `json:"background_color"`
	// Specify whether to enable the custom block page.
	Enabled bool `json:"enabled,nullable"`
	// Specify the block page footer text when the mode is customized_block_page.
	FooterText string `json:"footer_text"`
	// Specify the block page header text when the mode is customized_block_page.
	HeaderText string `json:"header_text"`
	// Specify whether to append context to target_uri as query parameters. This
	// applies only when the mode is redirect_uri.
	IncludeContext bool `json:"include_context"`
	// Specify the full URL to the logo file when the mode is customized_block_page.
	LogoPath string `json:"logo_path"`
	// Specify the admin email for users to contact when the mode is
	// customized_block_page.
	MailtoAddress string `json:"mailto_address"`
	// Specify the subject line for emails created from the block page when the mode is
	// customized_block_page.
	MailtoSubject string `json:"mailto_subject"`
	// Specify whether to redirect users to a Cloudflare-hosted block page or a
	// customer-provided URI.
	Mode BlockPageSettingsMode `json:"mode"`
	// Specify the block page title when the mode is customized_block_page.
	Name string `json:"name"`
	// Indicate that this setting was shared via the Orgs API and read only for the
	// current account.
	ReadOnly bool `json:"read_only,nullable"`
	// Indicate the account tag of the account that shared this setting.
	SourceAccount string `json:"source_account,nullable"`
	// Specify whether to suppress detailed information at the bottom of the block page
	// when the mode is customized_block_page.
	SuppressFooter bool `json:"suppress_footer"`
	// Specify the URI to redirect users to when the mode is redirect_uri.
	TargetURI string `json:"target_uri" format:"uri"`
	// Indicate the version number of the setting.
	Version int64                 `json:"version,nullable"`
	JSON    blockPageSettingsJSON `json:"-"`
}

// blockPageSettingsJSON contains the JSON metadata for the struct
// [BlockPageSettings]
type blockPageSettingsJSON struct {
	BackgroundColor apijson.Field
	Enabled         apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	IncludeContext  apijson.Field
	LogoPath        apijson.Field
	MailtoAddress   apijson.Field
	MailtoSubject   apijson.Field
	Mode            apijson.Field
	Name            apijson.Field
	ReadOnly        apijson.Field
	SourceAccount   apijson.Field
	SuppressFooter  apijson.Field
	TargetURI       apijson.Field
	Version         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *BlockPageSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockPageSettingsJSON) RawJSON() string {
	return r.raw
}

// Specify whether to redirect users to a Cloudflare-hosted block page or a
// customer-provided URI.
type BlockPageSettingsMode string

const (
	BlockPageSettingsModeEmpty               BlockPageSettingsMode = ""
	BlockPageSettingsModeCustomizedBlockPage BlockPageSettingsMode = "customized_block_page"
	BlockPageSettingsModeRedirectURI         BlockPageSettingsMode = "redirect_uri"
)

func (r BlockPageSettingsMode) IsKnown() bool {
	switch r {
	case BlockPageSettingsModeEmpty, BlockPageSettingsModeCustomizedBlockPage, BlockPageSettingsModeRedirectURI:
		return true
	}
	return false
}

// Specify block page layout settings.
type BlockPageSettingsParam struct {
	// Specify the block page background color in `#rrggbb` format when the mode is
	// customized_block_page.
	BackgroundColor param.Field[string] `json:"background_color"`
	// Specify whether to enable the custom block page.
	Enabled param.Field[bool] `json:"enabled"`
	// Specify the block page footer text when the mode is customized_block_page.
	FooterText param.Field[string] `json:"footer_text"`
	// Specify the block page header text when the mode is customized_block_page.
	HeaderText param.Field[string] `json:"header_text"`
	// Specify whether to append context to target_uri as query parameters. This
	// applies only when the mode is redirect_uri.
	IncludeContext param.Field[bool] `json:"include_context"`
	// Specify the full URL to the logo file when the mode is customized_block_page.
	LogoPath param.Field[string] `json:"logo_path"`
	// Specify the admin email for users to contact when the mode is
	// customized_block_page.
	MailtoAddress param.Field[string] `json:"mailto_address"`
	// Specify the subject line for emails created from the block page when the mode is
	// customized_block_page.
	MailtoSubject param.Field[string] `json:"mailto_subject"`
	// Specify whether to redirect users to a Cloudflare-hosted block page or a
	// customer-provided URI.
	Mode param.Field[BlockPageSettingsMode] `json:"mode"`
	// Specify the block page title when the mode is customized_block_page.
	Name param.Field[string] `json:"name"`
	// Specify whether to suppress detailed information at the bottom of the block page
	// when the mode is customized_block_page.
	SuppressFooter param.Field[bool] `json:"suppress_footer"`
	// Specify the URI to redirect users to when the mode is redirect_uri.
	TargetURI param.Field[string] `json:"target_uri" format:"uri"`
}

func (r BlockPageSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify the DLP inspection mode.
type BodyScanningSettings struct {
	// Specify the inspection mode as either `deep` or `shallow`.
	InspectionMode BodyScanningSettingsInspectionMode `json:"inspection_mode"`
	JSON           bodyScanningSettingsJSON           `json:"-"`
}

// bodyScanningSettingsJSON contains the JSON metadata for the struct
// [BodyScanningSettings]
type bodyScanningSettingsJSON struct {
	InspectionMode apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BodyScanningSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bodyScanningSettingsJSON) RawJSON() string {
	return r.raw
}

// Specify the inspection mode as either `deep` or `shallow`.
type BodyScanningSettingsInspectionMode string

const (
	BodyScanningSettingsInspectionModeDeep    BodyScanningSettingsInspectionMode = "deep"
	BodyScanningSettingsInspectionModeShallow BodyScanningSettingsInspectionMode = "shallow"
)

func (r BodyScanningSettingsInspectionMode) IsKnown() bool {
	switch r {
	case BodyScanningSettingsInspectionModeDeep, BodyScanningSettingsInspectionModeShallow:
		return true
	}
	return false
}

// Specify the DLP inspection mode.
type BodyScanningSettingsParam struct {
	// Specify the inspection mode as either `deep` or `shallow`.
	InspectionMode param.Field[BodyScanningSettingsInspectionMode] `json:"inspection_mode"`
}

func (r BodyScanningSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify Clientless Browser Isolation settings.
type BrowserIsolationSettings struct {
	// Specify whether to enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled bool `json:"non_identity_enabled"`
	// Specify whether to enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled bool                         `json:"url_browser_isolation_enabled"`
	JSON                       browserIsolationSettingsJSON `json:"-"`
}

// browserIsolationSettingsJSON contains the JSON metadata for the struct
// [BrowserIsolationSettings]
type browserIsolationSettingsJSON struct {
	NonIdentityEnabled         apijson.Field
	URLBrowserIsolationEnabled apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *BrowserIsolationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r browserIsolationSettingsJSON) RawJSON() string {
	return r.raw
}

// Specify Clientless Browser Isolation settings.
type BrowserIsolationSettingsParam struct {
	// Specify whether to enable non-identity onramp support for Browser Isolation.
	NonIdentityEnabled param.Field[bool] `json:"non_identity_enabled"`
	// Specify whether to enable Clientless Browser Isolation.
	URLBrowserIsolationEnabled param.Field[bool] `json:"url_browser_isolation_enabled"`
}

func (r BrowserIsolationSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify custom certificate settings for BYO-PKI. This field is deprecated; use
// `certificate` instead.
//
// Deprecated: deprecated
type CustomCertificateSettings struct {
	// Specify whether to enable a custom certificate authority for signing Gateway
	// traffic.
	Enabled bool `json:"enabled,required,nullable"`
	// Specify the UUID of the certificate (ID from MTLS certificate store).
	ID string `json:"id"`
	// Indicate the internal certificate status.
	BindingStatus string                        `json:"binding_status"`
	UpdatedAt     time.Time                     `json:"updated_at" format:"date-time"`
	JSON          customCertificateSettingsJSON `json:"-"`
}

// customCertificateSettingsJSON contains the JSON metadata for the struct
// [CustomCertificateSettings]
type customCertificateSettingsJSON struct {
	Enabled       apijson.Field
	ID            apijson.Field
	BindingStatus apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomCertificateSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customCertificateSettingsJSON) RawJSON() string {
	return r.raw
}

// Specify custom certificate settings for BYO-PKI. This field is deprecated; use
// `certificate` instead.
//
// Deprecated: deprecated
type CustomCertificateSettingsParam struct {
	// Specify whether to enable a custom certificate authority for signing Gateway
	// traffic.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Specify the UUID of the certificate (ID from MTLS certificate store).
	ID param.Field[string] `json:"id"`
}

func (r CustomCertificateSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures user email settings for firewall policies. When you enable this, the
// system standardizes email addresses in the identity portion of the rule to match
// extended email variants in firewall policies. When you disable this setting, the
// system matches email addresses exactly as you provide them. Enable this setting
// if your email uses `.` or `+` modifiers.
type ExtendedEmailMatching struct {
	// Specify whether to match all variants of user emails (with + or . modifiers)
	// used as criteria in Firewall policies.
	Enabled bool `json:"enabled,nullable"`
	// Indicate that this setting was shared via the Orgs API and read only for the
	// current account.
	ReadOnly bool `json:"read_only"`
	// Indicate the account tag of the account that shared this setting.
	SourceAccount string `json:"source_account"`
	// Indicate the version number of the setting.
	Version int64                     `json:"version"`
	JSON    extendedEmailMatchingJSON `json:"-"`
}

// extendedEmailMatchingJSON contains the JSON metadata for the struct
// [ExtendedEmailMatching]
type extendedEmailMatchingJSON struct {
	Enabled       apijson.Field
	ReadOnly      apijson.Field
	SourceAccount apijson.Field
	Version       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ExtendedEmailMatching) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extendedEmailMatchingJSON) RawJSON() string {
	return r.raw
}

// Configures user email settings for firewall policies. When you enable this, the
// system standardizes email addresses in the identity portion of the rule to match
// extended email variants in firewall policies. When you disable this setting, the
// system matches email addresses exactly as you provide them. Enable this setting
// if your email uses `.` or `+` modifiers.
type ExtendedEmailMatchingParam struct {
	// Specify whether to match all variants of user emails (with + or . modifiers)
	// used as criteria in Firewall policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ExtendedEmailMatchingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify FIPS settings.
type FipsSettings struct {
	// Enforce cipher suites and TLS versions compliant with FIPS 140-2.
	TLS  bool             `json:"tls"`
	JSON fipsSettingsJSON `json:"-"`
}

// fipsSettingsJSON contains the JSON metadata for the struct [FipsSettings]
type fipsSettingsJSON struct {
	TLS         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FipsSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fipsSettingsJSON) RawJSON() string {
	return r.raw
}

// Specify FIPS settings.
type FipsSettingsParam struct {
	// Enforce cipher suites and TLS versions compliant with FIPS 140-2.
	TLS param.Field[bool] `json:"tls"`
}

func (r FipsSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify account settings.
type GatewayConfigurationSettings struct {
	// Specify activity log settings.
	ActivityLog ActivityLogSettings `json:"activity_log,nullable"`
	// Specify anti-virus settings.
	Antivirus AntiVirusSettings `json:"antivirus,nullable"`
	// Specify block page layout settings.
	BlockPage BlockPageSettings `json:"block_page,nullable"`
	// Specify the DLP inspection mode.
	BodyScanning BodyScanningSettings `json:"body_scanning,nullable"`
	// Specify Clientless Browser Isolation settings.
	BrowserIsolation BrowserIsolationSettings `json:"browser_isolation,nullable"`
	// Specify certificate settings for Gateway TLS interception. If unset, the
	// Cloudflare Root CA handles interception.
	Certificate GatewayConfigurationSettingsCertificate `json:"certificate,nullable"`
	// Specify custom certificate settings for BYO-PKI. This field is deprecated; use
	// `certificate` instead.
	//
	// Deprecated: deprecated
	CustomCertificate CustomCertificateSettings `json:"custom_certificate,nullable"`
	// Configures user email settings for firewall policies. When you enable this, the
	// system standardizes email addresses in the identity portion of the rule to match
	// extended email variants in firewall policies. When you disable this setting, the
	// system matches email addresses exactly as you provide them. Enable this setting
	// if your email uses `.` or `+` modifiers.
	ExtendedEmailMatching ExtendedEmailMatching `json:"extended_email_matching,nullable"`
	// Specify FIPS settings.
	Fips FipsSettings `json:"fips,nullable"`
	// Enable host selection in egress policies.
	HostSelector GatewayConfigurationSettingsHostSelector `json:"host_selector,nullable"`
	// Define the proxy inspection mode.
	Inspection GatewayConfigurationSettingsInspection `json:"inspection,nullable"`
	// Specify whether to detect protocols from the initial bytes of client traffic.
	ProtocolDetection ProtocolDetection `json:"protocol_detection,nullable"`
	// Specify whether to enable the sandbox.
	Sandbox GatewayConfigurationSettingsSandbox `json:"sandbox,nullable"`
	// Specify whether to inspect encrypted HTTP traffic.
	TLSDecrypt TLSSettings                      `json:"tls_decrypt,nullable"`
	JSON       gatewayConfigurationSettingsJSON `json:"-"`
}

// gatewayConfigurationSettingsJSON contains the JSON metadata for the struct
// [GatewayConfigurationSettings]
type gatewayConfigurationSettingsJSON struct {
	ActivityLog           apijson.Field
	Antivirus             apijson.Field
	BlockPage             apijson.Field
	BodyScanning          apijson.Field
	BrowserIsolation      apijson.Field
	Certificate           apijson.Field
	CustomCertificate     apijson.Field
	ExtendedEmailMatching apijson.Field
	Fips                  apijson.Field
	HostSelector          apijson.Field
	Inspection            apijson.Field
	ProtocolDetection     apijson.Field
	Sandbox               apijson.Field
	TLSDecrypt            apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *GatewayConfigurationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationSettingsJSON) RawJSON() string {
	return r.raw
}

// Specify certificate settings for Gateway TLS interception. If unset, the
// Cloudflare Root CA handles interception.
type GatewayConfigurationSettingsCertificate struct {
	// Specify the UUID of the certificate used for interception. Ensure the
	// certificate is available at the edge(previously called 'active'). A nil UUID
	// directs Cloudflare to use the Root CA.
	ID   string                                      `json:"id,required"`
	JSON gatewayConfigurationSettingsCertificateJSON `json:"-"`
}

// gatewayConfigurationSettingsCertificateJSON contains the JSON metadata for the
// struct [GatewayConfigurationSettingsCertificate]
type gatewayConfigurationSettingsCertificateJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationSettingsCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationSettingsCertificateJSON) RawJSON() string {
	return r.raw
}

// Enable host selection in egress policies.
type GatewayConfigurationSettingsHostSelector struct {
	// Specify whether to enable filtering via hosts for egress policies.
	Enabled bool                                         `json:"enabled,nullable"`
	JSON    gatewayConfigurationSettingsHostSelectorJSON `json:"-"`
}

// gatewayConfigurationSettingsHostSelectorJSON contains the JSON metadata for the
// struct [GatewayConfigurationSettingsHostSelector]
type gatewayConfigurationSettingsHostSelectorJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationSettingsHostSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationSettingsHostSelectorJSON) RawJSON() string {
	return r.raw
}

// Define the proxy inspection mode.
type GatewayConfigurationSettingsInspection struct {
	// Define the proxy inspection mode. 1. static: Gateway applies static inspection
	// to HTTP on TCP(80). With TLS decryption on, Gateway inspects HTTPS traffic on
	// TCP(443) and UDP(443). 2. dynamic: Gateway applies protocol detection to inspect
	// HTTP and HTTPS traffic on any port. TLS decryption must remain on to inspect
	// HTTPS traffic.
	Mode GatewayConfigurationSettingsInspectionMode `json:"mode"`
	JSON gatewayConfigurationSettingsInspectionJSON `json:"-"`
}

// gatewayConfigurationSettingsInspectionJSON contains the JSON metadata for the
// struct [GatewayConfigurationSettingsInspection]
type gatewayConfigurationSettingsInspectionJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationSettingsInspection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationSettingsInspectionJSON) RawJSON() string {
	return r.raw
}

// Define the proxy inspection mode. 1. static: Gateway applies static inspection
// to HTTP on TCP(80). With TLS decryption on, Gateway inspects HTTPS traffic on
// TCP(443) and UDP(443). 2. dynamic: Gateway applies protocol detection to inspect
// HTTP and HTTPS traffic on any port. TLS decryption must remain on to inspect
// HTTPS traffic.
type GatewayConfigurationSettingsInspectionMode string

const (
	GatewayConfigurationSettingsInspectionModeStatic  GatewayConfigurationSettingsInspectionMode = "static"
	GatewayConfigurationSettingsInspectionModeDynamic GatewayConfigurationSettingsInspectionMode = "dynamic"
)

func (r GatewayConfigurationSettingsInspectionMode) IsKnown() bool {
	switch r {
	case GatewayConfigurationSettingsInspectionModeStatic, GatewayConfigurationSettingsInspectionModeDynamic:
		return true
	}
	return false
}

// Specify whether to enable the sandbox.
type GatewayConfigurationSettingsSandbox struct {
	// Specify whether to enable the sandbox.
	Enabled bool `json:"enabled,nullable"`
	// Specify the action to take when the system cannot scan the file.
	FallbackAction GatewayConfigurationSettingsSandboxFallbackAction `json:"fallback_action"`
	JSON           gatewayConfigurationSettingsSandboxJSON           `json:"-"`
}

// gatewayConfigurationSettingsSandboxJSON contains the JSON metadata for the
// struct [GatewayConfigurationSettingsSandbox]
type gatewayConfigurationSettingsSandboxJSON struct {
	Enabled        apijson.Field
	FallbackAction apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *GatewayConfigurationSettingsSandbox) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationSettingsSandboxJSON) RawJSON() string {
	return r.raw
}

// Specify the action to take when the system cannot scan the file.
type GatewayConfigurationSettingsSandboxFallbackAction string

const (
	GatewayConfigurationSettingsSandboxFallbackActionAllow GatewayConfigurationSettingsSandboxFallbackAction = "allow"
	GatewayConfigurationSettingsSandboxFallbackActionBlock GatewayConfigurationSettingsSandboxFallbackAction = "block"
)

func (r GatewayConfigurationSettingsSandboxFallbackAction) IsKnown() bool {
	switch r {
	case GatewayConfigurationSettingsSandboxFallbackActionAllow, GatewayConfigurationSettingsSandboxFallbackActionBlock:
		return true
	}
	return false
}

// Specify account settings.
type GatewayConfigurationSettingsParam struct {
	// Specify activity log settings.
	ActivityLog param.Field[ActivityLogSettingsParam] `json:"activity_log"`
	// Specify anti-virus settings.
	Antivirus param.Field[AntiVirusSettingsParam] `json:"antivirus"`
	// Specify block page layout settings.
	BlockPage param.Field[BlockPageSettingsParam] `json:"block_page"`
	// Specify the DLP inspection mode.
	BodyScanning param.Field[BodyScanningSettingsParam] `json:"body_scanning"`
	// Specify Clientless Browser Isolation settings.
	BrowserIsolation param.Field[BrowserIsolationSettingsParam] `json:"browser_isolation"`
	// Specify certificate settings for Gateway TLS interception. If unset, the
	// Cloudflare Root CA handles interception.
	Certificate param.Field[GatewayConfigurationSettingsCertificateParam] `json:"certificate"`
	// Specify custom certificate settings for BYO-PKI. This field is deprecated; use
	// `certificate` instead.
	//
	// Deprecated: deprecated
	CustomCertificate param.Field[CustomCertificateSettingsParam] `json:"custom_certificate"`
	// Configures user email settings for firewall policies. When you enable this, the
	// system standardizes email addresses in the identity portion of the rule to match
	// extended email variants in firewall policies. When you disable this setting, the
	// system matches email addresses exactly as you provide them. Enable this setting
	// if your email uses `.` or `+` modifiers.
	ExtendedEmailMatching param.Field[ExtendedEmailMatchingParam] `json:"extended_email_matching"`
	// Specify FIPS settings.
	Fips param.Field[FipsSettingsParam] `json:"fips"`
	// Enable host selection in egress policies.
	HostSelector param.Field[GatewayConfigurationSettingsHostSelectorParam] `json:"host_selector"`
	// Define the proxy inspection mode.
	Inspection param.Field[GatewayConfigurationSettingsInspectionParam] `json:"inspection"`
	// Specify whether to detect protocols from the initial bytes of client traffic.
	ProtocolDetection param.Field[ProtocolDetectionParam] `json:"protocol_detection"`
	// Specify whether to enable the sandbox.
	Sandbox param.Field[GatewayConfigurationSettingsSandboxParam] `json:"sandbox"`
	// Specify whether to inspect encrypted HTTP traffic.
	TLSDecrypt param.Field[TLSSettingsParam] `json:"tls_decrypt"`
}

func (r GatewayConfigurationSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify certificate settings for Gateway TLS interception. If unset, the
// Cloudflare Root CA handles interception.
type GatewayConfigurationSettingsCertificateParam struct {
	// Specify the UUID of the certificate used for interception. Ensure the
	// certificate is available at the edge(previously called 'active'). A nil UUID
	// directs Cloudflare to use the Root CA.
	ID param.Field[string] `json:"id,required"`
}

func (r GatewayConfigurationSettingsCertificateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable host selection in egress policies.
type GatewayConfigurationSettingsHostSelectorParam struct {
	// Specify whether to enable filtering via hosts for egress policies.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayConfigurationSettingsHostSelectorParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the proxy inspection mode.
type GatewayConfigurationSettingsInspectionParam struct {
	// Define the proxy inspection mode. 1. static: Gateway applies static inspection
	// to HTTP on TCP(80). With TLS decryption on, Gateway inspects HTTPS traffic on
	// TCP(443) and UDP(443). 2. dynamic: Gateway applies protocol detection to inspect
	// HTTP and HTTPS traffic on any port. TLS decryption must remain on to inspect
	// HTTPS traffic.
	Mode param.Field[GatewayConfigurationSettingsInspectionMode] `json:"mode"`
}

func (r GatewayConfigurationSettingsInspectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify whether to enable the sandbox.
type GatewayConfigurationSettingsSandboxParam struct {
	// Specify whether to enable the sandbox.
	Enabled param.Field[bool] `json:"enabled"`
	// Specify the action to take when the system cannot scan the file.
	FallbackAction param.Field[GatewayConfigurationSettingsSandboxFallbackAction] `json:"fallback_action"`
}

func (r GatewayConfigurationSettingsSandboxParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure the message the user's device shows during an antivirus scan.
type NotificationSettings struct {
	// Specify whether to enable notifications.
	Enabled bool `json:"enabled"`
	// Specify whether to include context information as query parameters.
	IncludeContext bool `json:"include_context"`
	// Specify the message to show in the notification.
	Msg string `json:"msg"`
	// Specify a URL that directs users to more information. If unset, the notification
	// opens a block page.
	SupportURL string                   `json:"support_url"`
	JSON       notificationSettingsJSON `json:"-"`
}

// notificationSettingsJSON contains the JSON metadata for the struct
// [NotificationSettings]
type notificationSettingsJSON struct {
	Enabled        apijson.Field
	IncludeContext apijson.Field
	Msg            apijson.Field
	SupportURL     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Configure the message the user's device shows during an antivirus scan.
type NotificationSettingsParam struct {
	// Specify whether to enable notifications.
	Enabled param.Field[bool] `json:"enabled"`
	// Specify whether to include context information as query parameters.
	IncludeContext param.Field[bool] `json:"include_context"`
	// Specify the message to show in the notification.
	Msg param.Field[string] `json:"msg"`
	// Specify a URL that directs users to more information. If unset, the notification
	// opens a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r NotificationSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify whether to detect protocols from the initial bytes of client traffic.
type ProtocolDetection struct {
	// Specify whether to detect protocols from the initial bytes of client traffic.
	Enabled bool                  `json:"enabled,nullable"`
	JSON    protocolDetectionJSON `json:"-"`
}

// protocolDetectionJSON contains the JSON metadata for the struct
// [ProtocolDetection]
type protocolDetectionJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProtocolDetection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r protocolDetectionJSON) RawJSON() string {
	return r.raw
}

// Specify whether to detect protocols from the initial bytes of client traffic.
type ProtocolDetectionParam struct {
	// Specify whether to detect protocols from the initial bytes of client traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ProtocolDetectionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify whether to inspect encrypted HTTP traffic.
type TLSSettings struct {
	// Specify whether to inspect encrypted HTTP traffic.
	Enabled bool            `json:"enabled"`
	JSON    tlsSettingsJSON `json:"-"`
}

// tlsSettingsJSON contains the JSON metadata for the struct [TLSSettings]
type tlsSettingsJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TLSSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsSettingsJSON) RawJSON() string {
	return r.raw
}

// Specify whether to inspect encrypted HTTP traffic.
type TLSSettingsParam struct {
	// Specify whether to inspect encrypted HTTP traffic.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r TLSSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify account settings.
type GatewayConfigurationUpdateResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specify account settings.
	Settings  GatewayConfigurationSettings           `json:"settings"`
	UpdatedAt time.Time                              `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationUpdateResponseJSON `json:"-"`
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

// Specify account settings.
type GatewayConfigurationEditResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specify account settings.
	Settings  GatewayConfigurationSettings         `json:"settings"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationEditResponseJSON `json:"-"`
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

// Specify account settings.
type GatewayConfigurationGetResponse struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Specify account settings.
	Settings  GatewayConfigurationSettings        `json:"settings"`
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      gatewayConfigurationGetResponseJSON `json:"-"`
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
	// Specify account settings.
	Settings param.Field[GatewayConfigurationSettingsParam] `json:"settings"`
}

func (r GatewayConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayConfigurationUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayConfigurationUpdateResponseEnvelopeSuccess `json:"success,required"`
	// Specify account settings.
	Result GatewayConfigurationUpdateResponse             `json:"result"`
	JSON   gatewayConfigurationUpdateResponseEnvelopeJSON `json:"-"`
}

// gatewayConfigurationUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayConfigurationUpdateResponseEnvelope]
type gatewayConfigurationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
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
	// Specify account settings.
	Settings param.Field[GatewayConfigurationSettingsParam] `json:"settings"`
}

func (r GatewayConfigurationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayConfigurationEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayConfigurationEditResponseEnvelopeSuccess `json:"success,required"`
	// Specify account settings.
	Result GatewayConfigurationEditResponse             `json:"result"`
	JSON   gatewayConfigurationEditResponseEnvelopeJSON `json:"-"`
}

// gatewayConfigurationEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayConfigurationEditResponseEnvelope]
type gatewayConfigurationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	// Specify account settings.
	Result GatewayConfigurationGetResponse             `json:"result"`
	JSON   gatewayConfigurationGetResponseEnvelopeJSON `json:"-"`
}

// gatewayConfigurationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayConfigurationGetResponseEnvelope]
type gatewayConfigurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayConfigurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
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

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostnames

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// CustomHostnameService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomHostnameService] method
// instead.
type CustomHostnameService struct {
	Options        []option.RequestOption
	FallbackOrigin *FallbackOriginService
}

// NewCustomHostnameService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomHostnameService(opts ...option.RequestOption) (r *CustomHostnameService) {
	r = &CustomHostnameService{}
	r.Options = opts
	r.FallbackOrigin = NewFallbackOriginService(opts...)
	return
}

// Add a new custom hostname and request that an SSL certificate be issued for it.
// One of three validation methods—http, txt, email—should be used, with 'http'
// recommended if the CNAME is already in place (or will be soon). Specifying
// 'email' will send an email to the WHOIS contacts on file for the base domain
// plus hostmaster, postmaster, webmaster, admin, administrator. If http is used
// and the domain is not already pointing to the Managed CNAME host, the PATCH
// method must be used once it is (to complete validation).
func (r *CustomHostnameService) New(ctx context.Context, params CustomHostnameNewParams, opts ...option.RequestOption) (res *CustomHostnameNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomHostnameNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, search, sort, and filter all of your custom hostnames.
func (r *CustomHostnameService) List(ctx context.Context, params CustomHostnameListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[CustomHostnameListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List, search, sort, and filter all of your custom hostnames.
func (r *CustomHostnameService) ListAutoPaging(ctx context.Context, params CustomHostnameListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[CustomHostnameListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete Custom Hostname (and any issued SSL certificates)
func (r *CustomHostnameService) Delete(ctx context.Context, customHostnameID string, params CustomHostnameDeleteParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef8900f4cb9dca9b9ed0ac41ad571e6837, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", params.ZoneID, customHostnameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Modify SSL configuration for a custom hostname. When sent with SSL config that
// matches existing config, used to indicate that hostname should pass domain
// control validation (DCV). Can also be used to change validation type, e.g., from
// 'http' to 'email'.
func (r *CustomHostnameService) Edit(ctx context.Context, customHostnameID string, params CustomHostnameEditParams, opts ...option.RequestOption) (res *CustomHostnameEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomHostnameEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", params.ZoneID, customHostnameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Custom Hostname Details
func (r *CustomHostnameService) Get(ctx context.Context, customHostnameID string, query CustomHostnameGetParams, opts ...option.RequestOption) (res *CustomHostnameGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomHostnameGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", query.ZoneID, customHostnameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// SSL properties for the custom hostname.
type SSL struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod UnnamedSchemaRef16aca57bde2963201c7e6e895436c1c1 `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority SSLCertificateAuthority `json:"certificate_authority"`
	// If a custom uploaded certificate is used.
	CustomCertificate string `json:"custom_certificate"`
	// The identifier for the Custom CSR that was used.
	CustomCsrID string `json:"custom_csr_id"`
	// The key for a custom uploaded certificate.
	CustomKey string `json:"custom_key"`
	// The time the custom certificate expires on.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// A list of Hostnames on a custom uploaded certificate.
	Hosts []interface{} `json:"hosts"`
	// The issuer on a custom uploaded certificate.
	Issuer string `json:"issuer"`
	// Domain control validation (DCV) method used for this hostname.
	Method UnnamedSchemaRef78adb375f06c6d462dd92b99e2ecf510 `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// SSL specific settings.
	Settings SSLSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status SSLStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type UnnamedSchemaRef9a9935a9a770967bb604ae41a81e42e1 `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []SSLValidationError  `json:"validation_errors"`
	ValidationRecords []SSLValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool    `json:"wildcard"`
	JSON     sslJSON `json:"-"`
}

// sslJSON contains the JSON metadata for the struct [SSL]
type sslJSON struct {
	ID                   apijson.Field
	BundleMethod         apijson.Field
	CertificateAuthority apijson.Field
	CustomCertificate    apijson.Field
	CustomCsrID          apijson.Field
	CustomKey            apijson.Field
	ExpiresOn            apijson.Field
	Hosts                apijson.Field
	Issuer               apijson.Field
	Method               apijson.Field
	SerialNumber         apijson.Field
	Settings             apijson.Field
	Signature            apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	UploadedOn           apijson.Field
	ValidationErrors     apijson.Field
	ValidationRecords    apijson.Field
	Wildcard             apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sslJSON) RawJSON() string {
	return r.raw
}

// The Certificate Authority that will issue the certificate
type SSLCertificateAuthority string

const (
	SSLCertificateAuthorityDigicert    SSLCertificateAuthority = "digicert"
	SSLCertificateAuthorityGoogle      SSLCertificateAuthority = "google"
	SSLCertificateAuthorityLetsEncrypt SSLCertificateAuthority = "lets_encrypt"
)

func (r SSLCertificateAuthority) IsKnown() bool {
	switch r {
	case SSLCertificateAuthorityDigicert, SSLCertificateAuthorityGoogle, SSLCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

// SSL specific settings.
type SSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints SSLSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 SSLSettingsHTTP2 `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion SSLSettingsMinTLSVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 SSLSettingsTLS1_3 `json:"tls_1_3"`
	JSON   sslSettingsJSON   `json:"-"`
}

// sslSettingsJSON contains the JSON metadata for the struct [SSLSettings]
type sslSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	HTTP2         apijson.Field
	MinTLSVersion apijson.Field
	TLS1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SSLSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sslSettingsJSON) RawJSON() string {
	return r.raw
}

// Whether or not Early Hints is enabled.
type SSLSettingsEarlyHints string

const (
	SSLSettingsEarlyHintsOn  SSLSettingsEarlyHints = "on"
	SSLSettingsEarlyHintsOff SSLSettingsEarlyHints = "off"
)

func (r SSLSettingsEarlyHints) IsKnown() bool {
	switch r {
	case SSLSettingsEarlyHintsOn, SSLSettingsEarlyHintsOff:
		return true
	}
	return false
}

// Whether or not HTTP2 is enabled.
type SSLSettingsHTTP2 string

const (
	SSLSettingsHTTP2On  SSLSettingsHTTP2 = "on"
	SSLSettingsHTTP2Off SSLSettingsHTTP2 = "off"
)

func (r SSLSettingsHTTP2) IsKnown() bool {
	switch r {
	case SSLSettingsHTTP2On, SSLSettingsHTTP2Off:
		return true
	}
	return false
}

// The minimum TLS version supported.
type SSLSettingsMinTLSVersion string

const (
	SSLSettingsMinTLSVersion1_0 SSLSettingsMinTLSVersion = "1.0"
	SSLSettingsMinTLSVersion1_1 SSLSettingsMinTLSVersion = "1.1"
	SSLSettingsMinTLSVersion1_2 SSLSettingsMinTLSVersion = "1.2"
	SSLSettingsMinTLSVersion1_3 SSLSettingsMinTLSVersion = "1.3"
)

func (r SSLSettingsMinTLSVersion) IsKnown() bool {
	switch r {
	case SSLSettingsMinTLSVersion1_0, SSLSettingsMinTLSVersion1_1, SSLSettingsMinTLSVersion1_2, SSLSettingsMinTLSVersion1_3:
		return true
	}
	return false
}

// Whether or not TLS 1.3 is enabled.
type SSLSettingsTLS1_3 string

const (
	SSLSettingsTLS1_3On  SSLSettingsTLS1_3 = "on"
	SSLSettingsTLS1_3Off SSLSettingsTLS1_3 = "off"
)

func (r SSLSettingsTLS1_3) IsKnown() bool {
	switch r {
	case SSLSettingsTLS1_3On, SSLSettingsTLS1_3Off:
		return true
	}
	return false
}

// Status of the hostname's SSL certificates.
type SSLStatus string

const (
	SSLStatusInitializing         SSLStatus = "initializing"
	SSLStatusPendingValidation    SSLStatus = "pending_validation"
	SSLStatusDeleted              SSLStatus = "deleted"
	SSLStatusPendingIssuance      SSLStatus = "pending_issuance"
	SSLStatusPendingDeployment    SSLStatus = "pending_deployment"
	SSLStatusPendingDeletion      SSLStatus = "pending_deletion"
	SSLStatusPendingExpiration    SSLStatus = "pending_expiration"
	SSLStatusExpired              SSLStatus = "expired"
	SSLStatusActive               SSLStatus = "active"
	SSLStatusInitializingTimedOut SSLStatus = "initializing_timed_out"
	SSLStatusValidationTimedOut   SSLStatus = "validation_timed_out"
	SSLStatusIssuanceTimedOut     SSLStatus = "issuance_timed_out"
	SSLStatusDeploymentTimedOut   SSLStatus = "deployment_timed_out"
	SSLStatusDeletionTimedOut     SSLStatus = "deletion_timed_out"
	SSLStatusPendingCleanup       SSLStatus = "pending_cleanup"
	SSLStatusStagingDeployment    SSLStatus = "staging_deployment"
	SSLStatusStagingActive        SSLStatus = "staging_active"
	SSLStatusDeactivating         SSLStatus = "deactivating"
	SSLStatusInactive             SSLStatus = "inactive"
	SSLStatusBackupIssued         SSLStatus = "backup_issued"
	SSLStatusHoldingDeployment    SSLStatus = "holding_deployment"
)

func (r SSLStatus) IsKnown() bool {
	switch r {
	case SSLStatusInitializing, SSLStatusPendingValidation, SSLStatusDeleted, SSLStatusPendingIssuance, SSLStatusPendingDeployment, SSLStatusPendingDeletion, SSLStatusPendingExpiration, SSLStatusExpired, SSLStatusActive, SSLStatusInitializingTimedOut, SSLStatusValidationTimedOut, SSLStatusIssuanceTimedOut, SSLStatusDeploymentTimedOut, SSLStatusDeletionTimedOut, SSLStatusPendingCleanup, SSLStatusStagingDeployment, SSLStatusStagingActive, SSLStatusDeactivating, SSLStatusInactive, SSLStatusBackupIssued, SSLStatusHoldingDeployment:
		return true
	}
	return false
}

type SSLValidationError struct {
	// A domain validation error.
	Message string                 `json:"message"`
	JSON    sslValidationErrorJSON `json:"-"`
}

// sslValidationErrorJSON contains the JSON metadata for the struct
// [SSLValidationError]
type sslValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sslValidationErrorJSON) RawJSON() string {
	return r.raw
}

// Certificate's required validation record.
type SSLValidationRecord struct {
	// The set of email addresses that the certificate authority (CA) will use to
	// complete domain validation.
	Emails []interface{} `json:"emails"`
	// The content that the certificate authority (CA) will expect to find at the
	// http_url during the domain validation.
	HTTPBody string `json:"http_body"`
	// The url that will be checked during domain validation.
	HTTPURL string `json:"http_url"`
	// The hostname that the certificate authority (CA) will check for a TXT record
	// during domain validation .
	TXTName string `json:"txt_name"`
	// The TXT record that the certificate authority (CA) will check during domain
	// validation.
	TXTValue string                  `json:"txt_value"`
	JSON     sslValidationRecordJSON `json:"-"`
}

// sslValidationRecordJSON contains the JSON metadata for the struct
// [SSLValidationRecord]
type sslValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TXTName     apijson.Field
	TXTValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SSLValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sslValidationRecordJSON) RawJSON() string {
	return r.raw
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type UnnamedSchemaRef16aca57bde2963201c7e6e895436c1c1 string

const (
	UnnamedSchemaRef16aca57bde2963201c7e6e895436c1c1Ubiquitous UnnamedSchemaRef16aca57bde2963201c7e6e895436c1c1 = "ubiquitous"
	UnnamedSchemaRef16aca57bde2963201c7e6e895436c1c1Optimal    UnnamedSchemaRef16aca57bde2963201c7e6e895436c1c1 = "optimal"
	UnnamedSchemaRef16aca57bde2963201c7e6e895436c1c1Force      UnnamedSchemaRef16aca57bde2963201c7e6e895436c1c1 = "force"
)

func (r UnnamedSchemaRef16aca57bde2963201c7e6e895436c1c1) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef16aca57bde2963201c7e6e895436c1c1Ubiquitous, UnnamedSchemaRef16aca57bde2963201c7e6e895436c1c1Optimal, UnnamedSchemaRef16aca57bde2963201c7e6e895436c1c1Force:
		return true
	}
	return false
}

// Domain control validation (DCV) method used for this hostname.
type UnnamedSchemaRef78adb375f06c6d462dd92b99e2ecf510 string

const (
	UnnamedSchemaRef78adb375f06c6d462dd92b99e2ecf510HTTP  UnnamedSchemaRef78adb375f06c6d462dd92b99e2ecf510 = "http"
	UnnamedSchemaRef78adb375f06c6d462dd92b99e2ecf510TXT   UnnamedSchemaRef78adb375f06c6d462dd92b99e2ecf510 = "txt"
	UnnamedSchemaRef78adb375f06c6d462dd92b99e2ecf510Email UnnamedSchemaRef78adb375f06c6d462dd92b99e2ecf510 = "email"
)

func (r UnnamedSchemaRef78adb375f06c6d462dd92b99e2ecf510) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef78adb375f06c6d462dd92b99e2ecf510HTTP, UnnamedSchemaRef78adb375f06c6d462dd92b99e2ecf510TXT, UnnamedSchemaRef78adb375f06c6d462dd92b99e2ecf510Email:
		return true
	}
	return false
}

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type UnnamedSchemaRef9a9935a9a770967bb604ae41a81e42e1 string

const (
	UnnamedSchemaRef9a9935a9a770967bb604ae41a81e42e1Dv UnnamedSchemaRef9a9935a9a770967bb604ae41a81e42e1 = "dv"
)

func (r UnnamedSchemaRef9a9935a9a770967bb604ae41a81e42e1) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef9a9935a9a770967bb604ae41a81e42e1Dv:
		return true
	}
	return false
}

type CustomHostnameNewResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string `json:"hostname,required"`
	// SSL properties for the custom hostname.
	SSL  SSL                           `json:"ssl,required"`
	JSON customHostnameNewResponseJSON `json:"-"`
}

// customHostnameNewResponseJSON contains the JSON metadata for the struct
// [CustomHostnameNewResponse]
type customHostnameNewResponseJSON struct {
	ID          apijson.Field
	Hostname    apijson.Field
	SSL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameNewResponseJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string `json:"hostname,required"`
	// SSL properties for the custom hostname.
	SSL  SSL                            `json:"ssl,required"`
	JSON customHostnameListResponseJSON `json:"-"`
}

// customHostnameListResponseJSON contains the JSON metadata for the struct
// [CustomHostnameListResponse]
type customHostnameListResponseJSON struct {
	ID          apijson.Field
	Hostname    apijson.Field
	SSL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameListResponseJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameEditResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string `json:"hostname,required"`
	// SSL properties for the custom hostname.
	SSL  SSL                            `json:"ssl,required"`
	JSON customHostnameEditResponseJSON `json:"-"`
}

// customHostnameEditResponseJSON contains the JSON metadata for the struct
// [CustomHostnameEditResponse]
type customHostnameEditResponseJSON struct {
	ID          apijson.Field
	Hostname    apijson.Field
	SSL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameEditResponseJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameGetResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string `json:"hostname,required"`
	// SSL properties for the custom hostname.
	SSL  SSL                           `json:"ssl,required"`
	JSON customHostnameGetResponseJSON `json:"-"`
}

// customHostnameGetResponseJSON contains the JSON metadata for the struct
// [CustomHostnameGetResponse]
type customHostnameGetResponseJSON struct {
	ID          apijson.Field
	Hostname    apijson.Field
	SSL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameGetResponseJSON) RawJSON() string {
	return r.raw
}

type CustomHostnameNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname param.Field[string] `json:"hostname,required"`
	// SSL properties used when creating the custom hostname.
	SSL param.Field[CustomHostnameNewParamsSSL] `json:"ssl,required"`
	// These are per-hostname (customer) settings.
	CustomMetadata param.Field[CustomHostnameNewParamsCustomMetadata] `json:"custom_metadata"`
}

func (r CustomHostnameNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SSL properties used when creating the custom hostname.
type CustomHostnameNewParamsSSL struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[UnnamedSchemaRef16aca57bde2963201c7e6e895436c1c1] `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority param.Field[CustomHostnameNewParamsSSLCertificateAuthority] `json:"certificate_authority"`
	// If a custom uploaded certificate is used.
	CustomCertificate param.Field[string] `json:"custom_certificate"`
	// The key for a custom uploaded certificate.
	CustomKey param.Field[string] `json:"custom_key"`
	// Domain control validation (DCV) method used for this hostname.
	Method param.Field[UnnamedSchemaRef78adb375f06c6d462dd92b99e2ecf510] `json:"method"`
	// SSL specific settings.
	Settings param.Field[CustomHostnameNewParamsSSLSettings] `json:"settings"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type param.Field[UnnamedSchemaRef9a9935a9a770967bb604ae41a81e42e1] `json:"type"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard param.Field[bool] `json:"wildcard"`
}

func (r CustomHostnameNewParamsSSL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Certificate Authority that will issue the certificate
type CustomHostnameNewParamsSSLCertificateAuthority string

const (
	CustomHostnameNewParamsSSLCertificateAuthorityDigicert    CustomHostnameNewParamsSSLCertificateAuthority = "digicert"
	CustomHostnameNewParamsSSLCertificateAuthorityGoogle      CustomHostnameNewParamsSSLCertificateAuthority = "google"
	CustomHostnameNewParamsSSLCertificateAuthorityLetsEncrypt CustomHostnameNewParamsSSLCertificateAuthority = "lets_encrypt"
)

func (r CustomHostnameNewParamsSSLCertificateAuthority) IsKnown() bool {
	switch r {
	case CustomHostnameNewParamsSSLCertificateAuthorityDigicert, CustomHostnameNewParamsSSLCertificateAuthorityGoogle, CustomHostnameNewParamsSSLCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

// SSL specific settings.
type CustomHostnameNewParamsSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers param.Field[[]string] `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints param.Field[CustomHostnameNewParamsSSLSettingsEarlyHints] `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 param.Field[CustomHostnameNewParamsSSLSettingsHTTP2] `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion param.Field[CustomHostnameNewParamsSSLSettingsMinTLSVersion] `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 param.Field[CustomHostnameNewParamsSSLSettingsTLS1_3] `json:"tls_1_3"`
}

func (r CustomHostnameNewParamsSSLSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not Early Hints is enabled.
type CustomHostnameNewParamsSSLSettingsEarlyHints string

const (
	CustomHostnameNewParamsSSLSettingsEarlyHintsOn  CustomHostnameNewParamsSSLSettingsEarlyHints = "on"
	CustomHostnameNewParamsSSLSettingsEarlyHintsOff CustomHostnameNewParamsSSLSettingsEarlyHints = "off"
)

func (r CustomHostnameNewParamsSSLSettingsEarlyHints) IsKnown() bool {
	switch r {
	case CustomHostnameNewParamsSSLSettingsEarlyHintsOn, CustomHostnameNewParamsSSLSettingsEarlyHintsOff:
		return true
	}
	return false
}

// Whether or not HTTP2 is enabled.
type CustomHostnameNewParamsSSLSettingsHTTP2 string

const (
	CustomHostnameNewParamsSSLSettingsHTTP2On  CustomHostnameNewParamsSSLSettingsHTTP2 = "on"
	CustomHostnameNewParamsSSLSettingsHTTP2Off CustomHostnameNewParamsSSLSettingsHTTP2 = "off"
)

func (r CustomHostnameNewParamsSSLSettingsHTTP2) IsKnown() bool {
	switch r {
	case CustomHostnameNewParamsSSLSettingsHTTP2On, CustomHostnameNewParamsSSLSettingsHTTP2Off:
		return true
	}
	return false
}

// The minimum TLS version supported.
type CustomHostnameNewParamsSSLSettingsMinTLSVersion string

const (
	CustomHostnameNewParamsSSLSettingsMinTLSVersion1_0 CustomHostnameNewParamsSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameNewParamsSSLSettingsMinTLSVersion1_1 CustomHostnameNewParamsSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameNewParamsSSLSettingsMinTLSVersion1_2 CustomHostnameNewParamsSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameNewParamsSSLSettingsMinTLSVersion1_3 CustomHostnameNewParamsSSLSettingsMinTLSVersion = "1.3"
)

func (r CustomHostnameNewParamsSSLSettingsMinTLSVersion) IsKnown() bool {
	switch r {
	case CustomHostnameNewParamsSSLSettingsMinTLSVersion1_0, CustomHostnameNewParamsSSLSettingsMinTLSVersion1_1, CustomHostnameNewParamsSSLSettingsMinTLSVersion1_2, CustomHostnameNewParamsSSLSettingsMinTLSVersion1_3:
		return true
	}
	return false
}

// Whether or not TLS 1.3 is enabled.
type CustomHostnameNewParamsSSLSettingsTLS1_3 string

const (
	CustomHostnameNewParamsSSLSettingsTLS1_3On  CustomHostnameNewParamsSSLSettingsTLS1_3 = "on"
	CustomHostnameNewParamsSSLSettingsTLS1_3Off CustomHostnameNewParamsSSLSettingsTLS1_3 = "off"
)

func (r CustomHostnameNewParamsSSLSettingsTLS1_3) IsKnown() bool {
	switch r {
	case CustomHostnameNewParamsSSLSettingsTLS1_3On, CustomHostnameNewParamsSSLSettingsTLS1_3Off:
		return true
	}
	return false
}

// These are per-hostname (customer) settings.
type CustomHostnameNewParamsCustomMetadata struct {
	// Unique metadata for this hostname.
	Key param.Field[string] `json:"key"`
}

func (r CustomHostnameNewParamsCustomMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomHostnameNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   CustomHostnameNewResponse                                 `json:"result,required"`
	// Whether the API call was successful
	Success CustomHostnameNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    customHostnameNewResponseEnvelopeJSON    `json:"-"`
}

// customHostnameNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomHostnameNewResponseEnvelope]
type customHostnameNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomHostnameNewResponseEnvelopeSuccess bool

const (
	CustomHostnameNewResponseEnvelopeSuccessTrue CustomHostnameNewResponseEnvelopeSuccess = true
)

func (r CustomHostnameNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomHostnameNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomHostnameListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Hostname ID to match against. This ID was generated and returned during the
	// initial custom_hostname creation. This parameter cannot be used with the
	// 'hostname' parameter.
	ID param.Field[string] `query:"id"`
	// Direction to order hostnames.
	Direction param.Field[CustomHostnameListParamsDirection] `query:"direction"`
	// Fully qualified domain name to match against. This parameter cannot be used with
	// the 'id' parameter.
	Hostname param.Field[string] `query:"hostname"`
	// Field to order hostnames by.
	Order param.Field[CustomHostnameListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of hostnames per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Whether to filter hostnames based on if they have SSL enabled.
	SSL param.Field[CustomHostnameListParamsSSL] `query:"ssl"`
}

// URLQuery serializes [CustomHostnameListParams]'s query parameters as
// `url.Values`.
func (r CustomHostnameListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order hostnames.
type CustomHostnameListParamsDirection string

const (
	CustomHostnameListParamsDirectionAsc  CustomHostnameListParamsDirection = "asc"
	CustomHostnameListParamsDirectionDesc CustomHostnameListParamsDirection = "desc"
)

func (r CustomHostnameListParamsDirection) IsKnown() bool {
	switch r {
	case CustomHostnameListParamsDirectionAsc, CustomHostnameListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order hostnames by.
type CustomHostnameListParamsOrder string

const (
	CustomHostnameListParamsOrderSSL       CustomHostnameListParamsOrder = "ssl"
	CustomHostnameListParamsOrderSSLStatus CustomHostnameListParamsOrder = "ssl_status"
)

func (r CustomHostnameListParamsOrder) IsKnown() bool {
	switch r {
	case CustomHostnameListParamsOrderSSL, CustomHostnameListParamsOrderSSLStatus:
		return true
	}
	return false
}

// Whether to filter hostnames based on if they have SSL enabled.
type CustomHostnameListParamsSSL float64

const (
	CustomHostnameListParamsSSL0 CustomHostnameListParamsSSL = 0
	CustomHostnameListParamsSSL1 CustomHostnameListParamsSSL = 1
)

func (r CustomHostnameListParamsSSL) IsKnown() bool {
	switch r {
	case CustomHostnameListParamsSSL0, CustomHostnameListParamsSSL1:
		return true
	}
	return false
}

type CustomHostnameDeleteParams struct {
	// Identifier
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r CustomHostnameDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CustomHostnameEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// These are per-hostname (customer) settings.
	CustomMetadata param.Field[CustomHostnameEditParamsCustomMetadata] `json:"custom_metadata"`
	// a valid hostname that’s been added to your DNS zone as an A, AAAA, or CNAME
	// record.
	CustomOriginServer param.Field[string] `json:"custom_origin_server"`
	// A hostname that will be sent to your custom origin server as SNI for TLS
	// handshake. This can be a valid subdomain of the zone or custom origin server
	// name or the string ':request_host_header:' which will cause the host header in
	// the request to be used as SNI. Not configurable with default/fallback origin
	// server.
	CustomOriginSni param.Field[string] `json:"custom_origin_sni"`
	// SSL properties used when creating the custom hostname.
	SSL param.Field[CustomHostnameEditParamsSSL] `json:"ssl"`
}

func (r CustomHostnameEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// These are per-hostname (customer) settings.
type CustomHostnameEditParamsCustomMetadata struct {
	// Unique metadata for this hostname.
	Key param.Field[string] `json:"key"`
}

func (r CustomHostnameEditParamsCustomMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SSL properties used when creating the custom hostname.
type CustomHostnameEditParamsSSL struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[UnnamedSchemaRef16aca57bde2963201c7e6e895436c1c1] `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority param.Field[CustomHostnameEditParamsSSLCertificateAuthority] `json:"certificate_authority"`
	// If a custom uploaded certificate is used.
	CustomCertificate param.Field[string] `json:"custom_certificate"`
	// The key for a custom uploaded certificate.
	CustomKey param.Field[string] `json:"custom_key"`
	// Domain control validation (DCV) method used for this hostname.
	Method param.Field[UnnamedSchemaRef78adb375f06c6d462dd92b99e2ecf510] `json:"method"`
	// SSL specific settings.
	Settings param.Field[CustomHostnameEditParamsSSLSettings] `json:"settings"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type param.Field[UnnamedSchemaRef9a9935a9a770967bb604ae41a81e42e1] `json:"type"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard param.Field[bool] `json:"wildcard"`
}

func (r CustomHostnameEditParamsSSL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The Certificate Authority that will issue the certificate
type CustomHostnameEditParamsSSLCertificateAuthority string

const (
	CustomHostnameEditParamsSSLCertificateAuthorityDigicert    CustomHostnameEditParamsSSLCertificateAuthority = "digicert"
	CustomHostnameEditParamsSSLCertificateAuthorityGoogle      CustomHostnameEditParamsSSLCertificateAuthority = "google"
	CustomHostnameEditParamsSSLCertificateAuthorityLetsEncrypt CustomHostnameEditParamsSSLCertificateAuthority = "lets_encrypt"
)

func (r CustomHostnameEditParamsSSLCertificateAuthority) IsKnown() bool {
	switch r {
	case CustomHostnameEditParamsSSLCertificateAuthorityDigicert, CustomHostnameEditParamsSSLCertificateAuthorityGoogle, CustomHostnameEditParamsSSLCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

// SSL specific settings.
type CustomHostnameEditParamsSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers param.Field[[]string] `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints param.Field[CustomHostnameEditParamsSSLSettingsEarlyHints] `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 param.Field[CustomHostnameEditParamsSSLSettingsHTTP2] `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion param.Field[CustomHostnameEditParamsSSLSettingsMinTLSVersion] `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 param.Field[CustomHostnameEditParamsSSLSettingsTLS1_3] `json:"tls_1_3"`
}

func (r CustomHostnameEditParamsSSLSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not Early Hints is enabled.
type CustomHostnameEditParamsSSLSettingsEarlyHints string

const (
	CustomHostnameEditParamsSSLSettingsEarlyHintsOn  CustomHostnameEditParamsSSLSettingsEarlyHints = "on"
	CustomHostnameEditParamsSSLSettingsEarlyHintsOff CustomHostnameEditParamsSSLSettingsEarlyHints = "off"
)

func (r CustomHostnameEditParamsSSLSettingsEarlyHints) IsKnown() bool {
	switch r {
	case CustomHostnameEditParamsSSLSettingsEarlyHintsOn, CustomHostnameEditParamsSSLSettingsEarlyHintsOff:
		return true
	}
	return false
}

// Whether or not HTTP2 is enabled.
type CustomHostnameEditParamsSSLSettingsHTTP2 string

const (
	CustomHostnameEditParamsSSLSettingsHTTP2On  CustomHostnameEditParamsSSLSettingsHTTP2 = "on"
	CustomHostnameEditParamsSSLSettingsHTTP2Off CustomHostnameEditParamsSSLSettingsHTTP2 = "off"
)

func (r CustomHostnameEditParamsSSLSettingsHTTP2) IsKnown() bool {
	switch r {
	case CustomHostnameEditParamsSSLSettingsHTTP2On, CustomHostnameEditParamsSSLSettingsHTTP2Off:
		return true
	}
	return false
}

// The minimum TLS version supported.
type CustomHostnameEditParamsSSLSettingsMinTLSVersion string

const (
	CustomHostnameEditParamsSSLSettingsMinTLSVersion1_0 CustomHostnameEditParamsSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameEditParamsSSLSettingsMinTLSVersion1_1 CustomHostnameEditParamsSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameEditParamsSSLSettingsMinTLSVersion1_2 CustomHostnameEditParamsSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameEditParamsSSLSettingsMinTLSVersion1_3 CustomHostnameEditParamsSSLSettingsMinTLSVersion = "1.3"
)

func (r CustomHostnameEditParamsSSLSettingsMinTLSVersion) IsKnown() bool {
	switch r {
	case CustomHostnameEditParamsSSLSettingsMinTLSVersion1_0, CustomHostnameEditParamsSSLSettingsMinTLSVersion1_1, CustomHostnameEditParamsSSLSettingsMinTLSVersion1_2, CustomHostnameEditParamsSSLSettingsMinTLSVersion1_3:
		return true
	}
	return false
}

// Whether or not TLS 1.3 is enabled.
type CustomHostnameEditParamsSSLSettingsTLS1_3 string

const (
	CustomHostnameEditParamsSSLSettingsTLS1_3On  CustomHostnameEditParamsSSLSettingsTLS1_3 = "on"
	CustomHostnameEditParamsSSLSettingsTLS1_3Off CustomHostnameEditParamsSSLSettingsTLS1_3 = "off"
)

func (r CustomHostnameEditParamsSSLSettingsTLS1_3) IsKnown() bool {
	switch r {
	case CustomHostnameEditParamsSSLSettingsTLS1_3On, CustomHostnameEditParamsSSLSettingsTLS1_3Off:
		return true
	}
	return false
}

type CustomHostnameEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   CustomHostnameEditResponse                                `json:"result,required"`
	// Whether the API call was successful
	Success CustomHostnameEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    customHostnameEditResponseEnvelopeJSON    `json:"-"`
}

// customHostnameEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomHostnameEditResponseEnvelope]
type customHostnameEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomHostnameEditResponseEnvelopeSuccess bool

const (
	CustomHostnameEditResponseEnvelopeSuccessTrue CustomHostnameEditResponseEnvelopeSuccess = true
)

func (r CustomHostnameEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomHostnameEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomHostnameGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type CustomHostnameGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   CustomHostnameGetResponse                                 `json:"result,required"`
	// Whether the API call was successful
	Success CustomHostnameGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    customHostnameGetResponseEnvelopeJSON    `json:"-"`
}

// customHostnameGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CustomHostnameGetResponseEnvelope]
type customHostnameGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customHostnameGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CustomHostnameGetResponseEnvelopeSuccess bool

const (
	CustomHostnameGetResponseEnvelopeSuccessTrue CustomHostnameGetResponseEnvelopeSuccess = true
)

func (r CustomHostnameGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomHostnameGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

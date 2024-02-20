// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// CustomHostnameService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomHostnameService] method
// instead.
type CustomHostnameService struct {
	Options        []option.RequestOption
	FallbackOrigin *CustomHostnameFallbackOriginService
}

// NewCustomHostnameService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomHostnameService(opts ...option.RequestOption) (r *CustomHostnameService) {
	r = &CustomHostnameService{}
	r.Options = opts
	r.FallbackOrigin = NewCustomHostnameFallbackOriginService(opts...)
	return
}

// Add a new custom hostname and request that an SSL certificate be issued for it.
// One of three validation methods—http, txt, email—should be used, with 'http'
// recommended if the CNAME is already in place (or will be soon). Specifying
// 'email' will send an email to the WHOIS contacts on file for the base domain
// plus hostmaster, postmaster, webmaster, admin, administrator. If http is used
// and the domain is not already pointing to the Managed CNAME host, the PATCH
// method must be used once it is (to complete validation).
func (r *CustomHostnameService) New(ctx context.Context, zoneID string, body CustomHostnameNewParams, opts ...option.RequestOption) (res *CustomHostnameNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomHostnameNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify SSL configuration for a custom hostname. When sent with SSL config that
// matches existing config, used to indicate that hostname should pass domain
// control validation (DCV). Can also be used to change validation type, e.g., from
// 'http' to 'email'.
func (r *CustomHostnameService) Update(ctx context.Context, zoneID string, customHostnameID string, body CustomHostnameUpdateParams, opts ...option.RequestOption) (res *CustomHostnameUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomHostnameUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", zoneID, customHostnameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, search, sort, and filter all of your custom hostnames.
func (r *CustomHostnameService) List(ctx context.Context, zoneID string, query CustomHostnameListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[CustomHostnameListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames", zoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *CustomHostnameService) ListAutoPaging(ctx context.Context, zoneID string, query CustomHostnameListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[CustomHostnameListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneID, query, opts...))
}

// Delete Custom Hostname (and any issued SSL certificates)
func (r *CustomHostnameService) Delete(ctx context.Context, zoneID string, customHostnameID string, opts ...option.RequestOption) (res *CustomHostnameDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", zoneID, customHostnameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Custom Hostname Details
func (r *CustomHostnameService) Get(ctx context.Context, zoneID string, customHostnameID string, opts ...option.RequestOption) (res *CustomHostnameGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomHostnameGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", zoneID, customHostnameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CustomHostnameNewResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string `json:"hostname,required"`
	// SSL properties for the custom hostname.
	SSL  CustomHostnameNewResponseSSL  `json:"ssl,required"`
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

// SSL properties for the custom hostname.
type CustomHostnameNewResponseSSL struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod CustomHostnameNewResponseSSLBundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority CustomHostnameNewResponseSSLCertificateAuthority `json:"certificate_authority"`
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
	Method CustomHostnameNewResponseSSLMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// SSL specific settings.
	Settings CustomHostnameNewResponseSSLSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status CustomHostnameNewResponseSSLStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type CustomHostnameNewResponseSSLType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []CustomHostnameNewResponseSSLValidationError  `json:"validation_errors"`
	ValidationRecords []CustomHostnameNewResponseSSLValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                             `json:"wildcard"`
	JSON     customHostnameNewResponseSSLJSON `json:"-"`
}

// customHostnameNewResponseSSLJSON contains the JSON metadata for the struct
// [CustomHostnameNewResponseSSL]
type customHostnameNewResponseSSLJSON struct {
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

func (r *CustomHostnameNewResponseSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomHostnameNewResponseSSLBundleMethod string

const (
	CustomHostnameNewResponseSSLBundleMethodUbiquitous CustomHostnameNewResponseSSLBundleMethod = "ubiquitous"
	CustomHostnameNewResponseSSLBundleMethodOptimal    CustomHostnameNewResponseSSLBundleMethod = "optimal"
	CustomHostnameNewResponseSSLBundleMethodForce      CustomHostnameNewResponseSSLBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type CustomHostnameNewResponseSSLCertificateAuthority string

const (
	CustomHostnameNewResponseSSLCertificateAuthorityDigicert    CustomHostnameNewResponseSSLCertificateAuthority = "digicert"
	CustomHostnameNewResponseSSLCertificateAuthorityGoogle      CustomHostnameNewResponseSSLCertificateAuthority = "google"
	CustomHostnameNewResponseSSLCertificateAuthorityLetsEncrypt CustomHostnameNewResponseSSLCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type CustomHostnameNewResponseSSLMethod string

const (
	CustomHostnameNewResponseSSLMethodHTTP  CustomHostnameNewResponseSSLMethod = "http"
	CustomHostnameNewResponseSSLMethodTxt   CustomHostnameNewResponseSSLMethod = "txt"
	CustomHostnameNewResponseSSLMethodEmail CustomHostnameNewResponseSSLMethod = "email"
)

// SSL specific settings.
type CustomHostnameNewResponseSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints CustomHostnameNewResponseSSLSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 CustomHostnameNewResponseSSLSettingsHTTP2 `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion CustomHostnameNewResponseSSLSettingsMinTLSVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 CustomHostnameNewResponseSSLSettingsTLS1_3 `json:"tls_1_3"`
	JSON   customHostnameNewResponseSSLSettingsJSON   `json:"-"`
}

// customHostnameNewResponseSSLSettingsJSON contains the JSON metadata for the
// struct [CustomHostnameNewResponseSSLSettings]
type customHostnameNewResponseSSLSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	HTTP2         apijson.Field
	MinTLSVersion apijson.Field
	TLS1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomHostnameNewResponseSSLSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not Early Hints is enabled.
type CustomHostnameNewResponseSSLSettingsEarlyHints string

const (
	CustomHostnameNewResponseSSLSettingsEarlyHintsOn  CustomHostnameNewResponseSSLSettingsEarlyHints = "on"
	CustomHostnameNewResponseSSLSettingsEarlyHintsOff CustomHostnameNewResponseSSLSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type CustomHostnameNewResponseSSLSettingsHTTP2 string

const (
	CustomHostnameNewResponseSSLSettingsHTTP2On  CustomHostnameNewResponseSSLSettingsHTTP2 = "on"
	CustomHostnameNewResponseSSLSettingsHTTP2Off CustomHostnameNewResponseSSLSettingsHTTP2 = "off"
)

// The minimum TLS version supported.
type CustomHostnameNewResponseSSLSettingsMinTLSVersion string

const (
	CustomHostnameNewResponseSSLSettingsMinTLSVersion1_0 CustomHostnameNewResponseSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameNewResponseSSLSettingsMinTLSVersion1_1 CustomHostnameNewResponseSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameNewResponseSSLSettingsMinTLSVersion1_2 CustomHostnameNewResponseSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameNewResponseSSLSettingsMinTLSVersion1_3 CustomHostnameNewResponseSSLSettingsMinTLSVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type CustomHostnameNewResponseSSLSettingsTLS1_3 string

const (
	CustomHostnameNewResponseSSLSettingsTLS1_3On  CustomHostnameNewResponseSSLSettingsTLS1_3 = "on"
	CustomHostnameNewResponseSSLSettingsTLS1_3Off CustomHostnameNewResponseSSLSettingsTLS1_3 = "off"
)

// Status of the hostname's SSL certificates.
type CustomHostnameNewResponseSSLStatus string

const (
	CustomHostnameNewResponseSSLStatusInitializing         CustomHostnameNewResponseSSLStatus = "initializing"
	CustomHostnameNewResponseSSLStatusPendingValidation    CustomHostnameNewResponseSSLStatus = "pending_validation"
	CustomHostnameNewResponseSSLStatusDeleted              CustomHostnameNewResponseSSLStatus = "deleted"
	CustomHostnameNewResponseSSLStatusPendingIssuance      CustomHostnameNewResponseSSLStatus = "pending_issuance"
	CustomHostnameNewResponseSSLStatusPendingDeployment    CustomHostnameNewResponseSSLStatus = "pending_deployment"
	CustomHostnameNewResponseSSLStatusPendingDeletion      CustomHostnameNewResponseSSLStatus = "pending_deletion"
	CustomHostnameNewResponseSSLStatusPendingExpiration    CustomHostnameNewResponseSSLStatus = "pending_expiration"
	CustomHostnameNewResponseSSLStatusExpired              CustomHostnameNewResponseSSLStatus = "expired"
	CustomHostnameNewResponseSSLStatusActive               CustomHostnameNewResponseSSLStatus = "active"
	CustomHostnameNewResponseSSLStatusInitializingTimedOut CustomHostnameNewResponseSSLStatus = "initializing_timed_out"
	CustomHostnameNewResponseSSLStatusValidationTimedOut   CustomHostnameNewResponseSSLStatus = "validation_timed_out"
	CustomHostnameNewResponseSSLStatusIssuanceTimedOut     CustomHostnameNewResponseSSLStatus = "issuance_timed_out"
	CustomHostnameNewResponseSSLStatusDeploymentTimedOut   CustomHostnameNewResponseSSLStatus = "deployment_timed_out"
	CustomHostnameNewResponseSSLStatusDeletionTimedOut     CustomHostnameNewResponseSSLStatus = "deletion_timed_out"
	CustomHostnameNewResponseSSLStatusPendingCleanup       CustomHostnameNewResponseSSLStatus = "pending_cleanup"
	CustomHostnameNewResponseSSLStatusStagingDeployment    CustomHostnameNewResponseSSLStatus = "staging_deployment"
	CustomHostnameNewResponseSSLStatusStagingActive        CustomHostnameNewResponseSSLStatus = "staging_active"
	CustomHostnameNewResponseSSLStatusDeactivating         CustomHostnameNewResponseSSLStatus = "deactivating"
	CustomHostnameNewResponseSSLStatusInactive             CustomHostnameNewResponseSSLStatus = "inactive"
	CustomHostnameNewResponseSSLStatusBackupIssued         CustomHostnameNewResponseSSLStatus = "backup_issued"
	CustomHostnameNewResponseSSLStatusHoldingDeployment    CustomHostnameNewResponseSSLStatus = "holding_deployment"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type CustomHostnameNewResponseSSLType string

const (
	CustomHostnameNewResponseSSLTypeDv CustomHostnameNewResponseSSLType = "dv"
)

type CustomHostnameNewResponseSSLValidationError struct {
	// A domain validation error.
	Message string                                          `json:"message"`
	JSON    customHostnameNewResponseSSLValidationErrorJSON `json:"-"`
}

// customHostnameNewResponseSSLValidationErrorJSON contains the JSON metadata for
// the struct [CustomHostnameNewResponseSSLValidationError]
type customHostnameNewResponseSSLValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameNewResponseSSLValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate's required validation record.
type CustomHostnameNewResponseSSLValidationRecord struct {
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
	TxtName string `json:"txt_name"`
	// The TXT record that the certificate authority (CA) will check during domain
	// validation.
	TxtValue string                                           `json:"txt_value"`
	JSON     customHostnameNewResponseSSLValidationRecordJSON `json:"-"`
}

// customHostnameNewResponseSSLValidationRecordJSON contains the JSON metadata for
// the struct [CustomHostnameNewResponseSSLValidationRecord]
type customHostnameNewResponseSSLValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TxtName     apijson.Field
	TxtValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameNewResponseSSLValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameUpdateResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string `json:"hostname,required"`
	// SSL properties for the custom hostname.
	SSL  CustomHostnameUpdateResponseSSL  `json:"ssl,required"`
	JSON customHostnameUpdateResponseJSON `json:"-"`
}

// customHostnameUpdateResponseJSON contains the JSON metadata for the struct
// [CustomHostnameUpdateResponse]
type customHostnameUpdateResponseJSON struct {
	ID          apijson.Field
	Hostname    apijson.Field
	SSL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// SSL properties for the custom hostname.
type CustomHostnameUpdateResponseSSL struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod CustomHostnameUpdateResponseSSLBundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority CustomHostnameUpdateResponseSSLCertificateAuthority `json:"certificate_authority"`
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
	Method CustomHostnameUpdateResponseSSLMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// SSL specific settings.
	Settings CustomHostnameUpdateResponseSSLSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status CustomHostnameUpdateResponseSSLStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type CustomHostnameUpdateResponseSSLType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []CustomHostnameUpdateResponseSSLValidationError  `json:"validation_errors"`
	ValidationRecords []CustomHostnameUpdateResponseSSLValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                                `json:"wildcard"`
	JSON     customHostnameUpdateResponseSSLJSON `json:"-"`
}

// customHostnameUpdateResponseSSLJSON contains the JSON metadata for the struct
// [CustomHostnameUpdateResponseSSL]
type customHostnameUpdateResponseSSLJSON struct {
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

func (r *CustomHostnameUpdateResponseSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomHostnameUpdateResponseSSLBundleMethod string

const (
	CustomHostnameUpdateResponseSSLBundleMethodUbiquitous CustomHostnameUpdateResponseSSLBundleMethod = "ubiquitous"
	CustomHostnameUpdateResponseSSLBundleMethodOptimal    CustomHostnameUpdateResponseSSLBundleMethod = "optimal"
	CustomHostnameUpdateResponseSSLBundleMethodForce      CustomHostnameUpdateResponseSSLBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type CustomHostnameUpdateResponseSSLCertificateAuthority string

const (
	CustomHostnameUpdateResponseSSLCertificateAuthorityDigicert    CustomHostnameUpdateResponseSSLCertificateAuthority = "digicert"
	CustomHostnameUpdateResponseSSLCertificateAuthorityGoogle      CustomHostnameUpdateResponseSSLCertificateAuthority = "google"
	CustomHostnameUpdateResponseSSLCertificateAuthorityLetsEncrypt CustomHostnameUpdateResponseSSLCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type CustomHostnameUpdateResponseSSLMethod string

const (
	CustomHostnameUpdateResponseSSLMethodHTTP  CustomHostnameUpdateResponseSSLMethod = "http"
	CustomHostnameUpdateResponseSSLMethodTxt   CustomHostnameUpdateResponseSSLMethod = "txt"
	CustomHostnameUpdateResponseSSLMethodEmail CustomHostnameUpdateResponseSSLMethod = "email"
)

// SSL specific settings.
type CustomHostnameUpdateResponseSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints CustomHostnameUpdateResponseSSLSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 CustomHostnameUpdateResponseSSLSettingsHTTP2 `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion CustomHostnameUpdateResponseSSLSettingsMinTLSVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 CustomHostnameUpdateResponseSSLSettingsTLS1_3 `json:"tls_1_3"`
	JSON   customHostnameUpdateResponseSSLSettingsJSON   `json:"-"`
}

// customHostnameUpdateResponseSSLSettingsJSON contains the JSON metadata for the
// struct [CustomHostnameUpdateResponseSSLSettings]
type customHostnameUpdateResponseSSLSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	HTTP2         apijson.Field
	MinTLSVersion apijson.Field
	TLS1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomHostnameUpdateResponseSSLSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not Early Hints is enabled.
type CustomHostnameUpdateResponseSSLSettingsEarlyHints string

const (
	CustomHostnameUpdateResponseSSLSettingsEarlyHintsOn  CustomHostnameUpdateResponseSSLSettingsEarlyHints = "on"
	CustomHostnameUpdateResponseSSLSettingsEarlyHintsOff CustomHostnameUpdateResponseSSLSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type CustomHostnameUpdateResponseSSLSettingsHTTP2 string

const (
	CustomHostnameUpdateResponseSSLSettingsHTTP2On  CustomHostnameUpdateResponseSSLSettingsHTTP2 = "on"
	CustomHostnameUpdateResponseSSLSettingsHTTP2Off CustomHostnameUpdateResponseSSLSettingsHTTP2 = "off"
)

// The minimum TLS version supported.
type CustomHostnameUpdateResponseSSLSettingsMinTLSVersion string

const (
	CustomHostnameUpdateResponseSSLSettingsMinTLSVersion1_0 CustomHostnameUpdateResponseSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameUpdateResponseSSLSettingsMinTLSVersion1_1 CustomHostnameUpdateResponseSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameUpdateResponseSSLSettingsMinTLSVersion1_2 CustomHostnameUpdateResponseSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameUpdateResponseSSLSettingsMinTLSVersion1_3 CustomHostnameUpdateResponseSSLSettingsMinTLSVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type CustomHostnameUpdateResponseSSLSettingsTLS1_3 string

const (
	CustomHostnameUpdateResponseSSLSettingsTLS1_3On  CustomHostnameUpdateResponseSSLSettingsTLS1_3 = "on"
	CustomHostnameUpdateResponseSSLSettingsTLS1_3Off CustomHostnameUpdateResponseSSLSettingsTLS1_3 = "off"
)

// Status of the hostname's SSL certificates.
type CustomHostnameUpdateResponseSSLStatus string

const (
	CustomHostnameUpdateResponseSSLStatusInitializing         CustomHostnameUpdateResponseSSLStatus = "initializing"
	CustomHostnameUpdateResponseSSLStatusPendingValidation    CustomHostnameUpdateResponseSSLStatus = "pending_validation"
	CustomHostnameUpdateResponseSSLStatusDeleted              CustomHostnameUpdateResponseSSLStatus = "deleted"
	CustomHostnameUpdateResponseSSLStatusPendingIssuance      CustomHostnameUpdateResponseSSLStatus = "pending_issuance"
	CustomHostnameUpdateResponseSSLStatusPendingDeployment    CustomHostnameUpdateResponseSSLStatus = "pending_deployment"
	CustomHostnameUpdateResponseSSLStatusPendingDeletion      CustomHostnameUpdateResponseSSLStatus = "pending_deletion"
	CustomHostnameUpdateResponseSSLStatusPendingExpiration    CustomHostnameUpdateResponseSSLStatus = "pending_expiration"
	CustomHostnameUpdateResponseSSLStatusExpired              CustomHostnameUpdateResponseSSLStatus = "expired"
	CustomHostnameUpdateResponseSSLStatusActive               CustomHostnameUpdateResponseSSLStatus = "active"
	CustomHostnameUpdateResponseSSLStatusInitializingTimedOut CustomHostnameUpdateResponseSSLStatus = "initializing_timed_out"
	CustomHostnameUpdateResponseSSLStatusValidationTimedOut   CustomHostnameUpdateResponseSSLStatus = "validation_timed_out"
	CustomHostnameUpdateResponseSSLStatusIssuanceTimedOut     CustomHostnameUpdateResponseSSLStatus = "issuance_timed_out"
	CustomHostnameUpdateResponseSSLStatusDeploymentTimedOut   CustomHostnameUpdateResponseSSLStatus = "deployment_timed_out"
	CustomHostnameUpdateResponseSSLStatusDeletionTimedOut     CustomHostnameUpdateResponseSSLStatus = "deletion_timed_out"
	CustomHostnameUpdateResponseSSLStatusPendingCleanup       CustomHostnameUpdateResponseSSLStatus = "pending_cleanup"
	CustomHostnameUpdateResponseSSLStatusStagingDeployment    CustomHostnameUpdateResponseSSLStatus = "staging_deployment"
	CustomHostnameUpdateResponseSSLStatusStagingActive        CustomHostnameUpdateResponseSSLStatus = "staging_active"
	CustomHostnameUpdateResponseSSLStatusDeactivating         CustomHostnameUpdateResponseSSLStatus = "deactivating"
	CustomHostnameUpdateResponseSSLStatusInactive             CustomHostnameUpdateResponseSSLStatus = "inactive"
	CustomHostnameUpdateResponseSSLStatusBackupIssued         CustomHostnameUpdateResponseSSLStatus = "backup_issued"
	CustomHostnameUpdateResponseSSLStatusHoldingDeployment    CustomHostnameUpdateResponseSSLStatus = "holding_deployment"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type CustomHostnameUpdateResponseSSLType string

const (
	CustomHostnameUpdateResponseSSLTypeDv CustomHostnameUpdateResponseSSLType = "dv"
)

type CustomHostnameUpdateResponseSSLValidationError struct {
	// A domain validation error.
	Message string                                             `json:"message"`
	JSON    customHostnameUpdateResponseSSLValidationErrorJSON `json:"-"`
}

// customHostnameUpdateResponseSSLValidationErrorJSON contains the JSON metadata
// for the struct [CustomHostnameUpdateResponseSSLValidationError]
type customHostnameUpdateResponseSSLValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameUpdateResponseSSLValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate's required validation record.
type CustomHostnameUpdateResponseSSLValidationRecord struct {
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
	TxtName string `json:"txt_name"`
	// The TXT record that the certificate authority (CA) will check during domain
	// validation.
	TxtValue string                                              `json:"txt_value"`
	JSON     customHostnameUpdateResponseSSLValidationRecordJSON `json:"-"`
}

// customHostnameUpdateResponseSSLValidationRecordJSON contains the JSON metadata
// for the struct [CustomHostnameUpdateResponseSSLValidationRecord]
type customHostnameUpdateResponseSSLValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TxtName     apijson.Field
	TxtValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameUpdateResponseSSLValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string `json:"hostname,required"`
	// SSL properties for the custom hostname.
	SSL  CustomHostnameListResponseSSL  `json:"ssl,required"`
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

// SSL properties for the custom hostname.
type CustomHostnameListResponseSSL struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod CustomHostnameListResponseSSLBundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority CustomHostnameListResponseSSLCertificateAuthority `json:"certificate_authority"`
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
	Method CustomHostnameListResponseSSLMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// SSL specific settings.
	Settings CustomHostnameListResponseSSLSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status CustomHostnameListResponseSSLStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type CustomHostnameListResponseSSLType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []CustomHostnameListResponseSSLValidationError  `json:"validation_errors"`
	ValidationRecords []CustomHostnameListResponseSSLValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                              `json:"wildcard"`
	JSON     customHostnameListResponseSSLJSON `json:"-"`
}

// customHostnameListResponseSSLJSON contains the JSON metadata for the struct
// [CustomHostnameListResponseSSL]
type customHostnameListResponseSSLJSON struct {
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

func (r *CustomHostnameListResponseSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomHostnameListResponseSSLBundleMethod string

const (
	CustomHostnameListResponseSSLBundleMethodUbiquitous CustomHostnameListResponseSSLBundleMethod = "ubiquitous"
	CustomHostnameListResponseSSLBundleMethodOptimal    CustomHostnameListResponseSSLBundleMethod = "optimal"
	CustomHostnameListResponseSSLBundleMethodForce      CustomHostnameListResponseSSLBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type CustomHostnameListResponseSSLCertificateAuthority string

const (
	CustomHostnameListResponseSSLCertificateAuthorityDigicert    CustomHostnameListResponseSSLCertificateAuthority = "digicert"
	CustomHostnameListResponseSSLCertificateAuthorityGoogle      CustomHostnameListResponseSSLCertificateAuthority = "google"
	CustomHostnameListResponseSSLCertificateAuthorityLetsEncrypt CustomHostnameListResponseSSLCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type CustomHostnameListResponseSSLMethod string

const (
	CustomHostnameListResponseSSLMethodHTTP  CustomHostnameListResponseSSLMethod = "http"
	CustomHostnameListResponseSSLMethodTxt   CustomHostnameListResponseSSLMethod = "txt"
	CustomHostnameListResponseSSLMethodEmail CustomHostnameListResponseSSLMethod = "email"
)

// SSL specific settings.
type CustomHostnameListResponseSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints CustomHostnameListResponseSSLSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 CustomHostnameListResponseSSLSettingsHTTP2 `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion CustomHostnameListResponseSSLSettingsMinTLSVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 CustomHostnameListResponseSSLSettingsTLS1_3 `json:"tls_1_3"`
	JSON   customHostnameListResponseSSLSettingsJSON   `json:"-"`
}

// customHostnameListResponseSSLSettingsJSON contains the JSON metadata for the
// struct [CustomHostnameListResponseSSLSettings]
type customHostnameListResponseSSLSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	HTTP2         apijson.Field
	MinTLSVersion apijson.Field
	TLS1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomHostnameListResponseSSLSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not Early Hints is enabled.
type CustomHostnameListResponseSSLSettingsEarlyHints string

const (
	CustomHostnameListResponseSSLSettingsEarlyHintsOn  CustomHostnameListResponseSSLSettingsEarlyHints = "on"
	CustomHostnameListResponseSSLSettingsEarlyHintsOff CustomHostnameListResponseSSLSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type CustomHostnameListResponseSSLSettingsHTTP2 string

const (
	CustomHostnameListResponseSSLSettingsHTTP2On  CustomHostnameListResponseSSLSettingsHTTP2 = "on"
	CustomHostnameListResponseSSLSettingsHTTP2Off CustomHostnameListResponseSSLSettingsHTTP2 = "off"
)

// The minimum TLS version supported.
type CustomHostnameListResponseSSLSettingsMinTLSVersion string

const (
	CustomHostnameListResponseSSLSettingsMinTLSVersion1_0 CustomHostnameListResponseSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameListResponseSSLSettingsMinTLSVersion1_1 CustomHostnameListResponseSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameListResponseSSLSettingsMinTLSVersion1_2 CustomHostnameListResponseSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameListResponseSSLSettingsMinTLSVersion1_3 CustomHostnameListResponseSSLSettingsMinTLSVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type CustomHostnameListResponseSSLSettingsTLS1_3 string

const (
	CustomHostnameListResponseSSLSettingsTLS1_3On  CustomHostnameListResponseSSLSettingsTLS1_3 = "on"
	CustomHostnameListResponseSSLSettingsTLS1_3Off CustomHostnameListResponseSSLSettingsTLS1_3 = "off"
)

// Status of the hostname's SSL certificates.
type CustomHostnameListResponseSSLStatus string

const (
	CustomHostnameListResponseSSLStatusInitializing         CustomHostnameListResponseSSLStatus = "initializing"
	CustomHostnameListResponseSSLStatusPendingValidation    CustomHostnameListResponseSSLStatus = "pending_validation"
	CustomHostnameListResponseSSLStatusDeleted              CustomHostnameListResponseSSLStatus = "deleted"
	CustomHostnameListResponseSSLStatusPendingIssuance      CustomHostnameListResponseSSLStatus = "pending_issuance"
	CustomHostnameListResponseSSLStatusPendingDeployment    CustomHostnameListResponseSSLStatus = "pending_deployment"
	CustomHostnameListResponseSSLStatusPendingDeletion      CustomHostnameListResponseSSLStatus = "pending_deletion"
	CustomHostnameListResponseSSLStatusPendingExpiration    CustomHostnameListResponseSSLStatus = "pending_expiration"
	CustomHostnameListResponseSSLStatusExpired              CustomHostnameListResponseSSLStatus = "expired"
	CustomHostnameListResponseSSLStatusActive               CustomHostnameListResponseSSLStatus = "active"
	CustomHostnameListResponseSSLStatusInitializingTimedOut CustomHostnameListResponseSSLStatus = "initializing_timed_out"
	CustomHostnameListResponseSSLStatusValidationTimedOut   CustomHostnameListResponseSSLStatus = "validation_timed_out"
	CustomHostnameListResponseSSLStatusIssuanceTimedOut     CustomHostnameListResponseSSLStatus = "issuance_timed_out"
	CustomHostnameListResponseSSLStatusDeploymentTimedOut   CustomHostnameListResponseSSLStatus = "deployment_timed_out"
	CustomHostnameListResponseSSLStatusDeletionTimedOut     CustomHostnameListResponseSSLStatus = "deletion_timed_out"
	CustomHostnameListResponseSSLStatusPendingCleanup       CustomHostnameListResponseSSLStatus = "pending_cleanup"
	CustomHostnameListResponseSSLStatusStagingDeployment    CustomHostnameListResponseSSLStatus = "staging_deployment"
	CustomHostnameListResponseSSLStatusStagingActive        CustomHostnameListResponseSSLStatus = "staging_active"
	CustomHostnameListResponseSSLStatusDeactivating         CustomHostnameListResponseSSLStatus = "deactivating"
	CustomHostnameListResponseSSLStatusInactive             CustomHostnameListResponseSSLStatus = "inactive"
	CustomHostnameListResponseSSLStatusBackupIssued         CustomHostnameListResponseSSLStatus = "backup_issued"
	CustomHostnameListResponseSSLStatusHoldingDeployment    CustomHostnameListResponseSSLStatus = "holding_deployment"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type CustomHostnameListResponseSSLType string

const (
	CustomHostnameListResponseSSLTypeDv CustomHostnameListResponseSSLType = "dv"
)

type CustomHostnameListResponseSSLValidationError struct {
	// A domain validation error.
	Message string                                           `json:"message"`
	JSON    customHostnameListResponseSSLValidationErrorJSON `json:"-"`
}

// customHostnameListResponseSSLValidationErrorJSON contains the JSON metadata for
// the struct [CustomHostnameListResponseSSLValidationError]
type customHostnameListResponseSSLValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameListResponseSSLValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate's required validation record.
type CustomHostnameListResponseSSLValidationRecord struct {
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
	TxtName string `json:"txt_name"`
	// The TXT record that the certificate authority (CA) will check during domain
	// validation.
	TxtValue string                                            `json:"txt_value"`
	JSON     customHostnameListResponseSSLValidationRecordJSON `json:"-"`
}

// customHostnameListResponseSSLValidationRecordJSON contains the JSON metadata for
// the struct [CustomHostnameListResponseSSLValidationRecord]
type customHostnameListResponseSSLValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TxtName     apijson.Field
	TxtValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameListResponseSSLValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameDeleteResponse struct {
	// Identifier
	ID   string                           `json:"id"`
	JSON customHostnameDeleteResponseJSON `json:"-"`
}

// customHostnameDeleteResponseJSON contains the JSON metadata for the struct
// [CustomHostnameDeleteResponse]
type customHostnameDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameGetResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string `json:"hostname,required"`
	// SSL properties for the custom hostname.
	SSL  CustomHostnameGetResponseSSL  `json:"ssl,required"`
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

// SSL properties for the custom hostname.
type CustomHostnameGetResponseSSL struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod CustomHostnameGetResponseSSLBundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority CustomHostnameGetResponseSSLCertificateAuthority `json:"certificate_authority"`
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
	Method CustomHostnameGetResponseSSLMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// SSL specific settings.
	Settings CustomHostnameGetResponseSSLSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status CustomHostnameGetResponseSSLStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type CustomHostnameGetResponseSSLType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []CustomHostnameGetResponseSSLValidationError  `json:"validation_errors"`
	ValidationRecords []CustomHostnameGetResponseSSLValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                             `json:"wildcard"`
	JSON     customHostnameGetResponseSSLJSON `json:"-"`
}

// customHostnameGetResponseSSLJSON contains the JSON metadata for the struct
// [CustomHostnameGetResponseSSL]
type customHostnameGetResponseSSLJSON struct {
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

func (r *CustomHostnameGetResponseSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomHostnameGetResponseSSLBundleMethod string

const (
	CustomHostnameGetResponseSSLBundleMethodUbiquitous CustomHostnameGetResponseSSLBundleMethod = "ubiquitous"
	CustomHostnameGetResponseSSLBundleMethodOptimal    CustomHostnameGetResponseSSLBundleMethod = "optimal"
	CustomHostnameGetResponseSSLBundleMethodForce      CustomHostnameGetResponseSSLBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type CustomHostnameGetResponseSSLCertificateAuthority string

const (
	CustomHostnameGetResponseSSLCertificateAuthorityDigicert    CustomHostnameGetResponseSSLCertificateAuthority = "digicert"
	CustomHostnameGetResponseSSLCertificateAuthorityGoogle      CustomHostnameGetResponseSSLCertificateAuthority = "google"
	CustomHostnameGetResponseSSLCertificateAuthorityLetsEncrypt CustomHostnameGetResponseSSLCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type CustomHostnameGetResponseSSLMethod string

const (
	CustomHostnameGetResponseSSLMethodHTTP  CustomHostnameGetResponseSSLMethod = "http"
	CustomHostnameGetResponseSSLMethodTxt   CustomHostnameGetResponseSSLMethod = "txt"
	CustomHostnameGetResponseSSLMethodEmail CustomHostnameGetResponseSSLMethod = "email"
)

// SSL specific settings.
type CustomHostnameGetResponseSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints CustomHostnameGetResponseSSLSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 CustomHostnameGetResponseSSLSettingsHTTP2 `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion CustomHostnameGetResponseSSLSettingsMinTLSVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 CustomHostnameGetResponseSSLSettingsTLS1_3 `json:"tls_1_3"`
	JSON   customHostnameGetResponseSSLSettingsJSON   `json:"-"`
}

// customHostnameGetResponseSSLSettingsJSON contains the JSON metadata for the
// struct [CustomHostnameGetResponseSSLSettings]
type customHostnameGetResponseSSLSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	HTTP2         apijson.Field
	MinTLSVersion apijson.Field
	TLS1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomHostnameGetResponseSSLSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not Early Hints is enabled.
type CustomHostnameGetResponseSSLSettingsEarlyHints string

const (
	CustomHostnameGetResponseSSLSettingsEarlyHintsOn  CustomHostnameGetResponseSSLSettingsEarlyHints = "on"
	CustomHostnameGetResponseSSLSettingsEarlyHintsOff CustomHostnameGetResponseSSLSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type CustomHostnameGetResponseSSLSettingsHTTP2 string

const (
	CustomHostnameGetResponseSSLSettingsHTTP2On  CustomHostnameGetResponseSSLSettingsHTTP2 = "on"
	CustomHostnameGetResponseSSLSettingsHTTP2Off CustomHostnameGetResponseSSLSettingsHTTP2 = "off"
)

// The minimum TLS version supported.
type CustomHostnameGetResponseSSLSettingsMinTLSVersion string

const (
	CustomHostnameGetResponseSSLSettingsMinTLSVersion1_0 CustomHostnameGetResponseSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameGetResponseSSLSettingsMinTLSVersion1_1 CustomHostnameGetResponseSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameGetResponseSSLSettingsMinTLSVersion1_2 CustomHostnameGetResponseSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameGetResponseSSLSettingsMinTLSVersion1_3 CustomHostnameGetResponseSSLSettingsMinTLSVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type CustomHostnameGetResponseSSLSettingsTLS1_3 string

const (
	CustomHostnameGetResponseSSLSettingsTLS1_3On  CustomHostnameGetResponseSSLSettingsTLS1_3 = "on"
	CustomHostnameGetResponseSSLSettingsTLS1_3Off CustomHostnameGetResponseSSLSettingsTLS1_3 = "off"
)

// Status of the hostname's SSL certificates.
type CustomHostnameGetResponseSSLStatus string

const (
	CustomHostnameGetResponseSSLStatusInitializing         CustomHostnameGetResponseSSLStatus = "initializing"
	CustomHostnameGetResponseSSLStatusPendingValidation    CustomHostnameGetResponseSSLStatus = "pending_validation"
	CustomHostnameGetResponseSSLStatusDeleted              CustomHostnameGetResponseSSLStatus = "deleted"
	CustomHostnameGetResponseSSLStatusPendingIssuance      CustomHostnameGetResponseSSLStatus = "pending_issuance"
	CustomHostnameGetResponseSSLStatusPendingDeployment    CustomHostnameGetResponseSSLStatus = "pending_deployment"
	CustomHostnameGetResponseSSLStatusPendingDeletion      CustomHostnameGetResponseSSLStatus = "pending_deletion"
	CustomHostnameGetResponseSSLStatusPendingExpiration    CustomHostnameGetResponseSSLStatus = "pending_expiration"
	CustomHostnameGetResponseSSLStatusExpired              CustomHostnameGetResponseSSLStatus = "expired"
	CustomHostnameGetResponseSSLStatusActive               CustomHostnameGetResponseSSLStatus = "active"
	CustomHostnameGetResponseSSLStatusInitializingTimedOut CustomHostnameGetResponseSSLStatus = "initializing_timed_out"
	CustomHostnameGetResponseSSLStatusValidationTimedOut   CustomHostnameGetResponseSSLStatus = "validation_timed_out"
	CustomHostnameGetResponseSSLStatusIssuanceTimedOut     CustomHostnameGetResponseSSLStatus = "issuance_timed_out"
	CustomHostnameGetResponseSSLStatusDeploymentTimedOut   CustomHostnameGetResponseSSLStatus = "deployment_timed_out"
	CustomHostnameGetResponseSSLStatusDeletionTimedOut     CustomHostnameGetResponseSSLStatus = "deletion_timed_out"
	CustomHostnameGetResponseSSLStatusPendingCleanup       CustomHostnameGetResponseSSLStatus = "pending_cleanup"
	CustomHostnameGetResponseSSLStatusStagingDeployment    CustomHostnameGetResponseSSLStatus = "staging_deployment"
	CustomHostnameGetResponseSSLStatusStagingActive        CustomHostnameGetResponseSSLStatus = "staging_active"
	CustomHostnameGetResponseSSLStatusDeactivating         CustomHostnameGetResponseSSLStatus = "deactivating"
	CustomHostnameGetResponseSSLStatusInactive             CustomHostnameGetResponseSSLStatus = "inactive"
	CustomHostnameGetResponseSSLStatusBackupIssued         CustomHostnameGetResponseSSLStatus = "backup_issued"
	CustomHostnameGetResponseSSLStatusHoldingDeployment    CustomHostnameGetResponseSSLStatus = "holding_deployment"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type CustomHostnameGetResponseSSLType string

const (
	CustomHostnameGetResponseSSLTypeDv CustomHostnameGetResponseSSLType = "dv"
)

type CustomHostnameGetResponseSSLValidationError struct {
	// A domain validation error.
	Message string                                          `json:"message"`
	JSON    customHostnameGetResponseSSLValidationErrorJSON `json:"-"`
}

// customHostnameGetResponseSSLValidationErrorJSON contains the JSON metadata for
// the struct [CustomHostnameGetResponseSSLValidationError]
type customHostnameGetResponseSSLValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameGetResponseSSLValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate's required validation record.
type CustomHostnameGetResponseSSLValidationRecord struct {
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
	TxtName string `json:"txt_name"`
	// The TXT record that the certificate authority (CA) will check during domain
	// validation.
	TxtValue string                                           `json:"txt_value"`
	JSON     customHostnameGetResponseSSLValidationRecordJSON `json:"-"`
}

// customHostnameGetResponseSSLValidationRecordJSON contains the JSON metadata for
// the struct [CustomHostnameGetResponseSSLValidationRecord]
type customHostnameGetResponseSSLValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TxtName     apijson.Field
	TxtValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameGetResponseSSLValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameNewParams struct {
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
	BundleMethod param.Field[CustomHostnameNewParamsSSLBundleMethod] `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority param.Field[CustomHostnameNewParamsSSLCertificateAuthority] `json:"certificate_authority"`
	// If a custom uploaded certificate is used.
	CustomCertificate param.Field[string] `json:"custom_certificate"`
	// The key for a custom uploaded certificate.
	CustomKey param.Field[string] `json:"custom_key"`
	// Domain control validation (DCV) method used for this hostname.
	Method param.Field[CustomHostnameNewParamsSSLMethod] `json:"method"`
	// SSL specific settings.
	Settings param.Field[CustomHostnameNewParamsSSLSettings] `json:"settings"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type param.Field[CustomHostnameNewParamsSSLType] `json:"type"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard param.Field[bool] `json:"wildcard"`
}

func (r CustomHostnameNewParamsSSL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomHostnameNewParamsSSLBundleMethod string

const (
	CustomHostnameNewParamsSSLBundleMethodUbiquitous CustomHostnameNewParamsSSLBundleMethod = "ubiquitous"
	CustomHostnameNewParamsSSLBundleMethodOptimal    CustomHostnameNewParamsSSLBundleMethod = "optimal"
	CustomHostnameNewParamsSSLBundleMethodForce      CustomHostnameNewParamsSSLBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type CustomHostnameNewParamsSSLCertificateAuthority string

const (
	CustomHostnameNewParamsSSLCertificateAuthorityDigicert    CustomHostnameNewParamsSSLCertificateAuthority = "digicert"
	CustomHostnameNewParamsSSLCertificateAuthorityGoogle      CustomHostnameNewParamsSSLCertificateAuthority = "google"
	CustomHostnameNewParamsSSLCertificateAuthorityLetsEncrypt CustomHostnameNewParamsSSLCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type CustomHostnameNewParamsSSLMethod string

const (
	CustomHostnameNewParamsSSLMethodHTTP  CustomHostnameNewParamsSSLMethod = "http"
	CustomHostnameNewParamsSSLMethodTxt   CustomHostnameNewParamsSSLMethod = "txt"
	CustomHostnameNewParamsSSLMethodEmail CustomHostnameNewParamsSSLMethod = "email"
)

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

// Whether or not HTTP2 is enabled.
type CustomHostnameNewParamsSSLSettingsHTTP2 string

const (
	CustomHostnameNewParamsSSLSettingsHTTP2On  CustomHostnameNewParamsSSLSettingsHTTP2 = "on"
	CustomHostnameNewParamsSSLSettingsHTTP2Off CustomHostnameNewParamsSSLSettingsHTTP2 = "off"
)

// The minimum TLS version supported.
type CustomHostnameNewParamsSSLSettingsMinTLSVersion string

const (
	CustomHostnameNewParamsSSLSettingsMinTLSVersion1_0 CustomHostnameNewParamsSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameNewParamsSSLSettingsMinTLSVersion1_1 CustomHostnameNewParamsSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameNewParamsSSLSettingsMinTLSVersion1_2 CustomHostnameNewParamsSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameNewParamsSSLSettingsMinTLSVersion1_3 CustomHostnameNewParamsSSLSettingsMinTLSVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type CustomHostnameNewParamsSSLSettingsTLS1_3 string

const (
	CustomHostnameNewParamsSSLSettingsTLS1_3On  CustomHostnameNewParamsSSLSettingsTLS1_3 = "on"
	CustomHostnameNewParamsSSLSettingsTLS1_3Off CustomHostnameNewParamsSSLSettingsTLS1_3 = "off"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type CustomHostnameNewParamsSSLType string

const (
	CustomHostnameNewParamsSSLTypeDv CustomHostnameNewParamsSSLType = "dv"
)

// These are per-hostname (customer) settings.
type CustomHostnameNewParamsCustomMetadata struct {
	// Unique metadata for this hostname.
	Key param.Field[string] `json:"key"`
}

func (r CustomHostnameNewParamsCustomMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomHostnameNewResponseEnvelope struct {
	Errors   []CustomHostnameNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomHostnameNewResponseEnvelopeMessages `json:"messages,required"`
	Result   CustomHostnameNewResponse                   `json:"result,required"`
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

type CustomHostnameNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    customHostnameNewResponseEnvelopeErrorsJSON `json:"-"`
}

// customHostnameNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomHostnameNewResponseEnvelopeErrors]
type customHostnameNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    customHostnameNewResponseEnvelopeMessagesJSON `json:"-"`
}

// customHostnameNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CustomHostnameNewResponseEnvelopeMessages]
type customHostnameNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomHostnameNewResponseEnvelopeSuccess bool

const (
	CustomHostnameNewResponseEnvelopeSuccessTrue CustomHostnameNewResponseEnvelopeSuccess = true
)

type CustomHostnameUpdateParams struct {
	// These are per-hostname (customer) settings.
	CustomMetadata param.Field[CustomHostnameUpdateParamsCustomMetadata] `json:"custom_metadata"`
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
	SSL param.Field[CustomHostnameUpdateParamsSSL] `json:"ssl"`
}

func (r CustomHostnameUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// These are per-hostname (customer) settings.
type CustomHostnameUpdateParamsCustomMetadata struct {
	// Unique metadata for this hostname.
	Key param.Field[string] `json:"key"`
}

func (r CustomHostnameUpdateParamsCustomMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SSL properties used when creating the custom hostname.
type CustomHostnameUpdateParamsSSL struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[CustomHostnameUpdateParamsSSLBundleMethod] `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority param.Field[CustomHostnameUpdateParamsSSLCertificateAuthority] `json:"certificate_authority"`
	// If a custom uploaded certificate is used.
	CustomCertificate param.Field[string] `json:"custom_certificate"`
	// The key for a custom uploaded certificate.
	CustomKey param.Field[string] `json:"custom_key"`
	// Domain control validation (DCV) method used for this hostname.
	Method param.Field[CustomHostnameUpdateParamsSSLMethod] `json:"method"`
	// SSL specific settings.
	Settings param.Field[CustomHostnameUpdateParamsSSLSettings] `json:"settings"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type param.Field[CustomHostnameUpdateParamsSSLType] `json:"type"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard param.Field[bool] `json:"wildcard"`
}

func (r CustomHostnameUpdateParamsSSL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomHostnameUpdateParamsSSLBundleMethod string

const (
	CustomHostnameUpdateParamsSSLBundleMethodUbiquitous CustomHostnameUpdateParamsSSLBundleMethod = "ubiquitous"
	CustomHostnameUpdateParamsSSLBundleMethodOptimal    CustomHostnameUpdateParamsSSLBundleMethod = "optimal"
	CustomHostnameUpdateParamsSSLBundleMethodForce      CustomHostnameUpdateParamsSSLBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type CustomHostnameUpdateParamsSSLCertificateAuthority string

const (
	CustomHostnameUpdateParamsSSLCertificateAuthorityDigicert    CustomHostnameUpdateParamsSSLCertificateAuthority = "digicert"
	CustomHostnameUpdateParamsSSLCertificateAuthorityGoogle      CustomHostnameUpdateParamsSSLCertificateAuthority = "google"
	CustomHostnameUpdateParamsSSLCertificateAuthorityLetsEncrypt CustomHostnameUpdateParamsSSLCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type CustomHostnameUpdateParamsSSLMethod string

const (
	CustomHostnameUpdateParamsSSLMethodHTTP  CustomHostnameUpdateParamsSSLMethod = "http"
	CustomHostnameUpdateParamsSSLMethodTxt   CustomHostnameUpdateParamsSSLMethod = "txt"
	CustomHostnameUpdateParamsSSLMethodEmail CustomHostnameUpdateParamsSSLMethod = "email"
)

// SSL specific settings.
type CustomHostnameUpdateParamsSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers param.Field[[]string] `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints param.Field[CustomHostnameUpdateParamsSSLSettingsEarlyHints] `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 param.Field[CustomHostnameUpdateParamsSSLSettingsHTTP2] `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion param.Field[CustomHostnameUpdateParamsSSLSettingsMinTLSVersion] `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 param.Field[CustomHostnameUpdateParamsSSLSettingsTLS1_3] `json:"tls_1_3"`
}

func (r CustomHostnameUpdateParamsSSLSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not Early Hints is enabled.
type CustomHostnameUpdateParamsSSLSettingsEarlyHints string

const (
	CustomHostnameUpdateParamsSSLSettingsEarlyHintsOn  CustomHostnameUpdateParamsSSLSettingsEarlyHints = "on"
	CustomHostnameUpdateParamsSSLSettingsEarlyHintsOff CustomHostnameUpdateParamsSSLSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type CustomHostnameUpdateParamsSSLSettingsHTTP2 string

const (
	CustomHostnameUpdateParamsSSLSettingsHTTP2On  CustomHostnameUpdateParamsSSLSettingsHTTP2 = "on"
	CustomHostnameUpdateParamsSSLSettingsHTTP2Off CustomHostnameUpdateParamsSSLSettingsHTTP2 = "off"
)

// The minimum TLS version supported.
type CustomHostnameUpdateParamsSSLSettingsMinTLSVersion string

const (
	CustomHostnameUpdateParamsSSLSettingsMinTLSVersion1_0 CustomHostnameUpdateParamsSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameUpdateParamsSSLSettingsMinTLSVersion1_1 CustomHostnameUpdateParamsSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameUpdateParamsSSLSettingsMinTLSVersion1_2 CustomHostnameUpdateParamsSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameUpdateParamsSSLSettingsMinTLSVersion1_3 CustomHostnameUpdateParamsSSLSettingsMinTLSVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type CustomHostnameUpdateParamsSSLSettingsTLS1_3 string

const (
	CustomHostnameUpdateParamsSSLSettingsTLS1_3On  CustomHostnameUpdateParamsSSLSettingsTLS1_3 = "on"
	CustomHostnameUpdateParamsSSLSettingsTLS1_3Off CustomHostnameUpdateParamsSSLSettingsTLS1_3 = "off"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type CustomHostnameUpdateParamsSSLType string

const (
	CustomHostnameUpdateParamsSSLTypeDv CustomHostnameUpdateParamsSSLType = "dv"
)

type CustomHostnameUpdateResponseEnvelope struct {
	Errors   []CustomHostnameUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomHostnameUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   CustomHostnameUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CustomHostnameUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    customHostnameUpdateResponseEnvelopeJSON    `json:"-"`
}

// customHostnameUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomHostnameUpdateResponseEnvelope]
type customHostnameUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    customHostnameUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// customHostnameUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CustomHostnameUpdateResponseEnvelopeErrors]
type customHostnameUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    customHostnameUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// customHostnameUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CustomHostnameUpdateResponseEnvelopeMessages]
type customHostnameUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomHostnameUpdateResponseEnvelopeSuccess bool

const (
	CustomHostnameUpdateResponseEnvelopeSuccessTrue CustomHostnameUpdateResponseEnvelopeSuccess = true
)

type CustomHostnameListParams struct {
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order hostnames.
type CustomHostnameListParamsDirection string

const (
	CustomHostnameListParamsDirectionAsc  CustomHostnameListParamsDirection = "asc"
	CustomHostnameListParamsDirectionDesc CustomHostnameListParamsDirection = "desc"
)

// Field to order hostnames by.
type CustomHostnameListParamsOrder string

const (
	CustomHostnameListParamsOrderSSL       CustomHostnameListParamsOrder = "ssl"
	CustomHostnameListParamsOrderSSLStatus CustomHostnameListParamsOrder = "ssl_status"
)

// Whether to filter hostnames based on if they have SSL enabled.
type CustomHostnameListParamsSSL float64

const (
	CustomHostnameListParamsSSL0 CustomHostnameListParamsSSL = 0
	CustomHostnameListParamsSSL1 CustomHostnameListParamsSSL = 1
)

type CustomHostnameGetResponseEnvelope struct {
	Errors   []CustomHostnameGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomHostnameGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CustomHostnameGetResponse                   `json:"result,required"`
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

type CustomHostnameGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    customHostnameGetResponseEnvelopeErrorsJSON `json:"-"`
}

// customHostnameGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomHostnameGetResponseEnvelopeErrors]
type customHostnameGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    customHostnameGetResponseEnvelopeMessagesJSON `json:"-"`
}

// customHostnameGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [CustomHostnameGetResponseEnvelopeMessages]
type customHostnameGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomHostnameGetResponseEnvelopeSuccess bool

const (
	CustomHostnameGetResponseEnvelopeSuccessTrue CustomHostnameGetResponseEnvelopeSuccess = true
)

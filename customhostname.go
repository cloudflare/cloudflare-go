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
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// CustomHostnameService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCustomHostnameService] method
// instead.
type CustomHostnameService struct {
	Options         []option.RequestOption
	FallbackOrigins *CustomHostnameFallbackOriginService
}

// NewCustomHostnameService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomHostnameService(opts ...option.RequestOption) (r *CustomHostnameService) {
	r = &CustomHostnameService{}
	r.Options = opts
	r.FallbackOrigins = NewCustomHostnameFallbackOriginService(opts...)
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

// Delete Custom Hostname (and any issued SSL certificates)
func (r *CustomHostnameService) Delete(ctx context.Context, zoneID string, customHostnameID string, opts ...option.RequestOption) (res *CustomHostnameDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", zoneID, customHostnameID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add a new custom hostname and request that an SSL certificate be issued for it.
// One of three validation methods—http, txt, email—should be used, with 'http'
// recommended if the CNAME is already in place (or will be soon). Specifying
// 'email' will send an email to the WHOIS contacts on file for the base domain
// plus hostmaster, postmaster, webmaster, admin, administrator. If http is used
// and the domain is not already pointing to the Managed CNAME host, the PATCH
// method must be used once it is (to complete validation).
func (r *CustomHostnameService) CustomHostnameForAZoneNewCustomHostname(ctx context.Context, zoneID string, body CustomHostnameCustomHostnameForAZoneNewCustomHostnameParams, opts ...option.RequestOption) (res *CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List, search, sort, and filter all of your custom hostnames.
func (r *CustomHostnameService) CustomHostnameForAZoneListCustomHostnames(ctx context.Context, zoneID string, query CustomHostnameCustomHostnameForAZoneListCustomHostnamesParams, opts ...option.RequestOption) (res *[]CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelope
	path := fmt.Sprintf("zones/%s/custom_hostnames", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string `json:"hostname,required"`
	// SSL properties for the custom hostname.
	SSL  CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSL  `json:"ssl,required"`
	JSON customHostnameCustomHostnameForAZoneNewCustomHostnameResponseJSON `json:"-"`
}

// customHostnameCustomHostnameForAZoneNewCustomHostnameResponseJSON contains the
// JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponse]
type customHostnameCustomHostnameForAZoneNewCustomHostnameResponseJSON struct {
	ID          apijson.Field
	Hostname    apijson.Field
	SSL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// SSL properties for the custom hostname.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSL struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLBundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLCertificateAuthority `json:"certificate_authority"`
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
	Method CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// SSL specific settings.
	Settings CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLValidationError  `json:"validation_errors"`
	ValidationRecords []CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                                                                 `json:"wildcard"`
	JSON     customHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLJSON `json:"-"`
}

// customHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLJSON contains
// the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSL]
type customHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLJSON struct {
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

func (r *CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLBundleMethod string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLBundleMethodUbiquitous CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLBundleMethod = "ubiquitous"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLBundleMethodOptimal    CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLBundleMethod = "optimal"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLBundleMethodForce      CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLCertificateAuthority string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLCertificateAuthorityDigicert    CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLCertificateAuthority = "digicert"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLCertificateAuthorityGoogle      CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLCertificateAuthority = "google"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLCertificateAuthorityLetsEncrypt CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLMethod string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLMethodHTTP  CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLMethod = "http"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLMethodTxt   CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLMethod = "txt"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLMethodEmail CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLMethod = "email"
)

// SSL specific settings.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsHTTP2 `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsMinTLSVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsTLS1_3 `json:"tls_1_3"`
	JSON   customHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsJSON   `json:"-"`
}

// customHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsJSON
// contains the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettings]
type customHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	HTTP2         apijson.Field
	MinTLSVersion apijson.Field
	TLS1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not Early Hints is enabled.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsEarlyHints string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsEarlyHintsOn  CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsEarlyHints = "on"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsEarlyHintsOff CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsHTTP2 string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsHTTP2On  CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsHTTP2 = "on"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsHTTP2Off CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsHTTP2 = "off"
)

// The minimum TLS version supported.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsMinTLSVersion string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsMinTLSVersion1_0 CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsMinTLSVersion1_1 CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsMinTLSVersion1_2 CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsMinTLSVersion1_3 CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsMinTLSVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsTLS1_3 string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsTLS1_3On  CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsTLS1_3 = "on"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsTLS1_3Off CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLSettingsTLS1_3 = "off"
)

// Status of the hostname's SSL certificates.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusInitializing         CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "initializing"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusPendingValidation    CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "pending_validation"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusDeleted              CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "deleted"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusPendingIssuance      CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "pending_issuance"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusPendingDeployment    CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "pending_deployment"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusPendingDeletion      CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "pending_deletion"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusPendingExpiration    CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "pending_expiration"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusExpired              CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "expired"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusActive               CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "active"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusInitializingTimedOut CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "initializing_timed_out"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusValidationTimedOut   CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "validation_timed_out"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusIssuanceTimedOut     CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "issuance_timed_out"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusDeploymentTimedOut   CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "deployment_timed_out"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusDeletionTimedOut     CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "deletion_timed_out"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusPendingCleanup       CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "pending_cleanup"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusStagingDeployment    CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "staging_deployment"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusStagingActive        CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "staging_active"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusDeactivating         CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "deactivating"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusInactive             CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "inactive"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusBackupIssued         CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "backup_issued"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatusHoldingDeployment    CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLStatus = "holding_deployment"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLType string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLTypeDv CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLType = "dv"
)

type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLValidationError struct {
	// A domain validation error.
	Message string                                                                              `json:"message"`
	JSON    customHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLValidationErrorJSON `json:"-"`
}

// customHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLValidationErrorJSON
// contains the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLValidationError]
type customHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate's required validation record.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLValidationRecord struct {
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
	TxtValue string                                                                               `json:"txt_value"`
	JSON     customHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLValidationRecordJSON `json:"-"`
}

// customHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLValidationRecordJSON
// contains the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLValidationRecord]
type customHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TxtName     apijson.Field
	TxtValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSSLValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string `json:"hostname,required"`
	// SSL properties for the custom hostname.
	SSL  CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSL  `json:"ssl,required"`
	JSON customHostnameCustomHostnameForAZoneListCustomHostnamesResponseJSON `json:"-"`
}

// customHostnameCustomHostnameForAZoneListCustomHostnamesResponseJSON contains the
// JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponse]
type customHostnameCustomHostnameForAZoneListCustomHostnamesResponseJSON struct {
	ID          apijson.Field
	Hostname    apijson.Field
	SSL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// SSL properties for the custom hostname.
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSL struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLBundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLCertificateAuthority `json:"certificate_authority"`
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
	Method CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// SSL specific settings.
	Settings CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLValidationError  `json:"validation_errors"`
	ValidationRecords []CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                                                                   `json:"wildcard"`
	JSON     customHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLJSON `json:"-"`
}

// customHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLJSON contains
// the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSL]
type customHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLJSON struct {
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

func (r *CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLBundleMethod string

const (
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLBundleMethodUbiquitous CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLBundleMethod = "ubiquitous"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLBundleMethodOptimal    CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLBundleMethod = "optimal"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLBundleMethodForce      CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLCertificateAuthority string

const (
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLCertificateAuthorityDigicert    CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLCertificateAuthority = "digicert"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLCertificateAuthorityGoogle      CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLCertificateAuthority = "google"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLCertificateAuthorityLetsEncrypt CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLMethod string

const (
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLMethodHTTP  CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLMethod = "http"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLMethodTxt   CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLMethod = "txt"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLMethodEmail CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLMethod = "email"
)

// SSL specific settings.
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsHTTP2 `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsMinTLSVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsTLS1_3 `json:"tls_1_3"`
	JSON   customHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsJSON   `json:"-"`
}

// customHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsJSON
// contains the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettings]
type customHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	HTTP2         apijson.Field
	MinTLSVersion apijson.Field
	TLS1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not Early Hints is enabled.
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsEarlyHints string

const (
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsEarlyHintsOn  CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsEarlyHints = "on"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsEarlyHintsOff CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsHTTP2 string

const (
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsHTTP2On  CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsHTTP2 = "on"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsHTTP2Off CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsHTTP2 = "off"
)

// The minimum TLS version supported.
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsMinTLSVersion string

const (
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsMinTLSVersion1_0 CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsMinTLSVersion1_1 CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsMinTLSVersion1_2 CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsMinTLSVersion1_3 CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsMinTLSVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsTLS1_3 string

const (
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsTLS1_3On  CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsTLS1_3 = "on"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsTLS1_3Off CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLSettingsTLS1_3 = "off"
)

// Status of the hostname's SSL certificates.
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus string

const (
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusInitializing         CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "initializing"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusPendingValidation    CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "pending_validation"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusDeleted              CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "deleted"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusPendingIssuance      CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "pending_issuance"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusPendingDeployment    CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "pending_deployment"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusPendingDeletion      CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "pending_deletion"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusPendingExpiration    CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "pending_expiration"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusExpired              CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "expired"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusActive               CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "active"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusInitializingTimedOut CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "initializing_timed_out"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusValidationTimedOut   CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "validation_timed_out"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusIssuanceTimedOut     CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "issuance_timed_out"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusDeploymentTimedOut   CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "deployment_timed_out"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusDeletionTimedOut     CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "deletion_timed_out"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusPendingCleanup       CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "pending_cleanup"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusStagingDeployment    CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "staging_deployment"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusStagingActive        CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "staging_active"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusDeactivating         CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "deactivating"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusInactive             CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "inactive"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusBackupIssued         CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "backup_issued"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatusHoldingDeployment    CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLStatus = "holding_deployment"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLType string

const (
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLTypeDv CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLType = "dv"
)

type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLValidationError struct {
	// A domain validation error.
	Message string                                                                                `json:"message"`
	JSON    customHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLValidationErrorJSON `json:"-"`
}

// customHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLValidationErrorJSON
// contains the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLValidationError]
type customHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate's required validation record.
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLValidationRecord struct {
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
	TxtValue string                                                                                 `json:"txt_value"`
	JSON     customHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLValidationRecordJSON `json:"-"`
}

// customHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLValidationRecordJSON
// contains the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLValidationRecord]
type customHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TxtName     apijson.Field
	TxtValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSSLValidationRecord) UnmarshalJSON(data []byte) (err error) {
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

type CustomHostnameCustomHostnameForAZoneNewCustomHostnameParams struct {
	// The custom hostname that will point to your hostname via CNAME.
	Hostname param.Field[string] `json:"hostname,required"`
	// SSL properties used when creating the custom hostname.
	SSL param.Field[CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSL] `json:"ssl,required"`
	// These are per-hostname (customer) settings.
	CustomMetadata param.Field[CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsCustomMetadata] `json:"custom_metadata"`
}

func (r CustomHostnameCustomHostnameForAZoneNewCustomHostnameParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SSL properties used when creating the custom hostname.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSL struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLBundleMethod] `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority param.Field[CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLCertificateAuthority] `json:"certificate_authority"`
	// If a custom uploaded certificate is used.
	CustomCertificate param.Field[string] `json:"custom_certificate"`
	// The key for a custom uploaded certificate.
	CustomKey param.Field[string] `json:"custom_key"`
	// Domain control validation (DCV) method used for this hostname.
	Method param.Field[CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLMethod] `json:"method"`
	// SSL specific settings.
	Settings param.Field[CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettings] `json:"settings"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type param.Field[CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLType] `json:"type"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard param.Field[bool] `json:"wildcard"`
}

func (r CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLBundleMethod string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLBundleMethodUbiquitous CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLBundleMethod = "ubiquitous"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLBundleMethodOptimal    CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLBundleMethod = "optimal"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLBundleMethodForce      CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLCertificateAuthority string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLCertificateAuthorityDigicert    CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLCertificateAuthority = "digicert"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLCertificateAuthorityGoogle      CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLCertificateAuthority = "google"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLCertificateAuthorityLetsEncrypt CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLMethod string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLMethodHTTP  CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLMethod = "http"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLMethodTxt   CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLMethod = "txt"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLMethodEmail CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLMethod = "email"
)

// SSL specific settings.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers param.Field[[]string] `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints param.Field[CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsEarlyHints] `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	HTTP2 param.Field[CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsHTTP2] `json:"http2"`
	// The minimum TLS version supported.
	MinTLSVersion param.Field[CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsMinTLSVersion] `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	TLS1_3 param.Field[CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsTLS1_3] `json:"tls_1_3"`
}

func (r CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not Early Hints is enabled.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsEarlyHints string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsEarlyHintsOn  CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsEarlyHints = "on"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsEarlyHintsOff CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsHTTP2 string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsHTTP2On  CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsHTTP2 = "on"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsHTTP2Off CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsHTTP2 = "off"
)

// The minimum TLS version supported.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsMinTLSVersion string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsMinTLSVersion1_0 CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsMinTLSVersion = "1.0"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsMinTLSVersion1_1 CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsMinTLSVersion = "1.1"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsMinTLSVersion1_2 CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsMinTLSVersion = "1.2"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsMinTLSVersion1_3 CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsMinTLSVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsTLS1_3 string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsTLS1_3On  CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsTLS1_3 = "on"
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsTLS1_3Off CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLSettingsTLS1_3 = "off"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLType string

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLTypeDv CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSSLType = "dv"
)

// These are per-hostname (customer) settings.
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsCustomMetadata struct {
	// Unique metadata for this hostname.
	Key param.Field[string] `json:"key"`
}

func (r CustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsCustomMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelope struct {
	Errors   []CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeMessages `json:"messages,required"`
	Result   CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeSuccess `json:"success,required"`
	JSON    customHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeJSON    `json:"-"`
}

// customHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelope]
type customHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeErrors struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    customHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeErrorsJSON `json:"-"`
}

// customHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeErrors]
type customHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeMessages struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    customHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeMessagesJSON `json:"-"`
}

// customHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeMessages]
type customHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeSuccess bool

const (
	CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeSuccessTrue CustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseEnvelopeSuccess = true
)

type CustomHostnameCustomHostnameForAZoneListCustomHostnamesParams struct {
	// Hostname ID to match against. This ID was generated and returned during the
	// initial custom_hostname creation. This parameter cannot be used with the
	// 'hostname' parameter.
	ID param.Field[string] `query:"id"`
	// Direction to order hostnames.
	Direction param.Field[CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsDirection] `query:"direction"`
	// Fully qualified domain name to match against. This parameter cannot be used with
	// the 'id' parameter.
	Hostname param.Field[string] `query:"hostname"`
	// Field to order hostnames by.
	Order param.Field[CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of hostnames per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Whether to filter hostnames based on if they have SSL enabled.
	SSL param.Field[CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsSSL] `query:"ssl"`
}

// URLQuery serializes
// [CustomHostnameCustomHostnameForAZoneListCustomHostnamesParams]'s query
// parameters as `url.Values`.
func (r CustomHostnameCustomHostnameForAZoneListCustomHostnamesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order hostnames.
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsDirection string

const (
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsDirectionAsc  CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsDirection = "asc"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsDirectionDesc CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsDirection = "desc"
)

// Field to order hostnames by.
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsOrder string

const (
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsOrderSSL       CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsOrder = "ssl"
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsOrderSSLStatus CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsOrder = "ssl_status"
)

// Whether to filter hostnames based on if they have SSL enabled.
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsSSL float64

const (
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsSSL0 CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsSSL = 0
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsSSL1 CustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsSSL = 1
)

type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelope struct {
	Errors   []CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeMessages `json:"messages,required"`
	Result   []CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       customHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeJSON       `json:"-"`
}

// customHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelope]
type customHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeErrors struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    customHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeErrorsJSON `json:"-"`
}

// customHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeErrors]
type customHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeMessages struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    customHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeMessagesJSON `json:"-"`
}

// customHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeMessages]
type customHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeSuccess bool

const (
	CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeSuccessTrue CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeSuccess = true
)

type CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                               `json:"total_count"`
	JSON       customHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeResultInfoJSON `json:"-"`
}

// customHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeResultInfo]
type customHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

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

// ZoneCustomHostnameService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneCustomHostnameService] method
// instead.
type ZoneCustomHostnameService struct {
	Options         []option.RequestOption
	FallbackOrigins *ZoneCustomHostnameFallbackOriginService
}

// NewZoneCustomHostnameService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneCustomHostnameService(opts ...option.RequestOption) (r *ZoneCustomHostnameService) {
	r = &ZoneCustomHostnameService{}
	r.Options = opts
	r.FallbackOrigins = NewZoneCustomHostnameFallbackOriginService(opts...)
	return
}

// Custom Hostname Details
func (r *ZoneCustomHostnameService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *CustomHostnameResponseSingle4Byp7wot, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify SSL configuration for a custom hostname. When sent with SSL config that
// matches existing config, used to indicate that hostname should pass domain
// control validation (DCV). Can also be used to change validation type, e.g., from
// 'http' to 'email'.
func (r *ZoneCustomHostnameService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneCustomHostnameUpdateParams, opts ...option.RequestOption) (res *CustomHostnameResponseSingle4Byp7wot, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete Custom Hostname (and any issued SSL certificates)
func (r *ZoneCustomHostnameService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneCustomHostnameDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", zoneIdentifier, identifier)
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
func (r *ZoneCustomHostnameService) CustomHostnameForAZoneNewCustomHostname(ctx context.Context, zoneIdentifier string, body ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParams, opts ...option.RequestOption) (res *CustomHostnameResponseSingle4Byp7wot, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List, search, sort, and filter all of your custom hostnames.
func (r *ZoneCustomHostnameService) CustomHostnameForAZoneListCustomHostnames(ctx context.Context, zoneIdentifier string, query ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParams, opts ...option.RequestOption) (res *CustomHostnameResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type CustomHostnameResponseCollection struct {
	Errors     []CustomHostnameResponseCollectionError    `json:"errors"`
	Messages   []CustomHostnameResponseCollectionMessage  `json:"messages"`
	Result     []CustomHostnameResponseCollectionResult   `json:"result"`
	ResultInfo CustomHostnameResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success CustomHostnameResponseCollectionSuccess `json:"success"`
	JSON    customHostnameResponseCollectionJSON    `json:"-"`
}

// customHostnameResponseCollectionJSON contains the JSON metadata for the struct
// [CustomHostnameResponseCollection]
type customHostnameResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameResponseCollectionError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    customHostnameResponseCollectionErrorJSON `json:"-"`
}

// customHostnameResponseCollectionErrorJSON contains the JSON metadata for the
// struct [CustomHostnameResponseCollectionError]
type customHostnameResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameResponseCollectionMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    customHostnameResponseCollectionMessageJSON `json:"-"`
}

// customHostnameResponseCollectionMessageJSON contains the JSON metadata for the
// struct [CustomHostnameResponseCollectionMessage]
type customHostnameResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameResponseCollectionResult struct {
	// Identifier
	ID string `json:"id"`
	// This is the time the hostname was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// These are per-hostname (customer) settings.
	CustomMetadata CustomHostnameResponseCollectionResultCustomMetadata `json:"custom_metadata"`
	// a valid hostname that’s been added to your DNS zone as an A, AAAA, or CNAME
	// record.
	CustomOriginServer string `json:"custom_origin_server"`
	// A hostname that will be sent to your custom origin server as SNI for TLS
	// handshake. This can be a valid subdomain of the zone or custom origin server
	// name or the string ':request_host_header:' which will cause the host header in
	// the request to be used as SNI. Not configurable with default/fallback origin
	// server.
	CustomOriginSni string `json:"custom_origin_sni"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string `json:"hostname"`
	// This is a record which can be placed to activate a hostname.
	OwnershipVerification CustomHostnameResponseCollectionResultOwnershipVerification `json:"ownership_verification"`
	// This presents the token to be served by the given http url to activate a
	// hostname.
	OwnershipVerificationHTTP CustomHostnameResponseCollectionResultOwnershipVerificationHTTP `json:"ownership_verification_http"`
	// SSL properties for the custom hostname.
	Ssl CustomHostnameResponseCollectionResultSsl `json:"ssl"`
	// Status of the hostname's activation.
	Status CustomHostnameResponseCollectionResultStatus `json:"status"`
	// These are errors that were encountered while trying to activate a hostname.
	VerificationErrors []interface{}                              `json:"verification_errors"`
	JSON               customHostnameResponseCollectionResultJSON `json:"-"`
}

// customHostnameResponseCollectionResultJSON contains the JSON metadata for the
// struct [CustomHostnameResponseCollectionResult]
type customHostnameResponseCollectionResultJSON struct {
	ID                        apijson.Field
	CreatedAt                 apijson.Field
	CustomMetadata            apijson.Field
	CustomOriginServer        apijson.Field
	CustomOriginSni           apijson.Field
	Hostname                  apijson.Field
	OwnershipVerification     apijson.Field
	OwnershipVerificationHTTP apijson.Field
	Ssl                       apijson.Field
	Status                    apijson.Field
	VerificationErrors        apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CustomHostnameResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// These are per-hostname (customer) settings.
type CustomHostnameResponseCollectionResultCustomMetadata struct {
	// Unique metadata for this hostname.
	Key  string                                                   `json:"key"`
	JSON customHostnameResponseCollectionResultCustomMetadataJSON `json:"-"`
}

// customHostnameResponseCollectionResultCustomMetadataJSON contains the JSON
// metadata for the struct [CustomHostnameResponseCollectionResultCustomMetadata]
type customHostnameResponseCollectionResultCustomMetadataJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseCollectionResultCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This is a record which can be placed to activate a hostname.
type CustomHostnameResponseCollectionResultOwnershipVerification struct {
	// DNS Name for record.
	Name string `json:"name"`
	// DNS Record type.
	Type CustomHostnameResponseCollectionResultOwnershipVerificationType `json:"type"`
	// Content for the record.
	Value string                                                          `json:"value"`
	JSON  customHostnameResponseCollectionResultOwnershipVerificationJSON `json:"-"`
}

// customHostnameResponseCollectionResultOwnershipVerificationJSON contains the
// JSON metadata for the struct
// [CustomHostnameResponseCollectionResultOwnershipVerification]
type customHostnameResponseCollectionResultOwnershipVerificationJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseCollectionResultOwnershipVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS Record type.
type CustomHostnameResponseCollectionResultOwnershipVerificationType string

const (
	CustomHostnameResponseCollectionResultOwnershipVerificationTypeTxt CustomHostnameResponseCollectionResultOwnershipVerificationType = "txt"
)

// This presents the token to be served by the given http url to activate a
// hostname.
type CustomHostnameResponseCollectionResultOwnershipVerificationHTTP struct {
	// Token to be served.
	HTTPBody string `json:"http_body"`
	// The HTTP URL that will be checked during custom hostname verification and where
	// the customer should host the token.
	HTTPURL string                                                              `json:"http_url"`
	JSON    customHostnameResponseCollectionResultOwnershipVerificationHTTPJSON `json:"-"`
}

// customHostnameResponseCollectionResultOwnershipVerificationHTTPJSON contains the
// JSON metadata for the struct
// [CustomHostnameResponseCollectionResultOwnershipVerificationHTTP]
type customHostnameResponseCollectionResultOwnershipVerificationHTTPJSON struct {
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseCollectionResultOwnershipVerificationHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// SSL properties for the custom hostname.
type CustomHostnameResponseCollectionResultSsl struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod CustomHostnameResponseCollectionResultSslBundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority CustomHostnameResponseCollectionResultSslCertificateAuthority `json:"certificate_authority"`
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
	Method CustomHostnameResponseCollectionResultSslMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// SSL specific settings.
	Settings CustomHostnameResponseCollectionResultSslSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status CustomHostnameResponseCollectionResultSslStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type CustomHostnameResponseCollectionResultSslType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []CustomHostnameResponseCollectionResultSslValidationError  `json:"validation_errors"`
	ValidationRecords []CustomHostnameResponseCollectionResultSslValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                                          `json:"wildcard"`
	JSON     customHostnameResponseCollectionResultSslJSON `json:"-"`
}

// customHostnameResponseCollectionResultSslJSON contains the JSON metadata for the
// struct [CustomHostnameResponseCollectionResultSsl]
type customHostnameResponseCollectionResultSslJSON struct {
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

func (r *CustomHostnameResponseCollectionResultSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomHostnameResponseCollectionResultSslBundleMethod string

const (
	CustomHostnameResponseCollectionResultSslBundleMethodUbiquitous CustomHostnameResponseCollectionResultSslBundleMethod = "ubiquitous"
	CustomHostnameResponseCollectionResultSslBundleMethodOptimal    CustomHostnameResponseCollectionResultSslBundleMethod = "optimal"
	CustomHostnameResponseCollectionResultSslBundleMethodForce      CustomHostnameResponseCollectionResultSslBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type CustomHostnameResponseCollectionResultSslCertificateAuthority string

const (
	CustomHostnameResponseCollectionResultSslCertificateAuthorityDigicert    CustomHostnameResponseCollectionResultSslCertificateAuthority = "digicert"
	CustomHostnameResponseCollectionResultSslCertificateAuthorityGoogle      CustomHostnameResponseCollectionResultSslCertificateAuthority = "google"
	CustomHostnameResponseCollectionResultSslCertificateAuthorityLetsEncrypt CustomHostnameResponseCollectionResultSslCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type CustomHostnameResponseCollectionResultSslMethod string

const (
	CustomHostnameResponseCollectionResultSslMethodHTTP  CustomHostnameResponseCollectionResultSslMethod = "http"
	CustomHostnameResponseCollectionResultSslMethodTxt   CustomHostnameResponseCollectionResultSslMethod = "txt"
	CustomHostnameResponseCollectionResultSslMethodEmail CustomHostnameResponseCollectionResultSslMethod = "email"
)

// SSL specific settings.
type CustomHostnameResponseCollectionResultSslSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints CustomHostnameResponseCollectionResultSslSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	Http2 CustomHostnameResponseCollectionResultSslSettingsHttp2 `json:"http2"`
	// The minimum TLS version supported.
	MinTlsVersion CustomHostnameResponseCollectionResultSslSettingsMinTlsVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	Tls1_3 CustomHostnameResponseCollectionResultSslSettingsTls1_3 `json:"tls_1_3"`
	JSON   customHostnameResponseCollectionResultSslSettingsJSON   `json:"-"`
}

// customHostnameResponseCollectionResultSslSettingsJSON contains the JSON metadata
// for the struct [CustomHostnameResponseCollectionResultSslSettings]
type customHostnameResponseCollectionResultSslSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	Http2         apijson.Field
	MinTlsVersion apijson.Field
	Tls1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomHostnameResponseCollectionResultSslSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not Early Hints is enabled.
type CustomHostnameResponseCollectionResultSslSettingsEarlyHints string

const (
	CustomHostnameResponseCollectionResultSslSettingsEarlyHintsOn  CustomHostnameResponseCollectionResultSslSettingsEarlyHints = "on"
	CustomHostnameResponseCollectionResultSslSettingsEarlyHintsOff CustomHostnameResponseCollectionResultSslSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type CustomHostnameResponseCollectionResultSslSettingsHttp2 string

const (
	CustomHostnameResponseCollectionResultSslSettingsHttp2On  CustomHostnameResponseCollectionResultSslSettingsHttp2 = "on"
	CustomHostnameResponseCollectionResultSslSettingsHttp2Off CustomHostnameResponseCollectionResultSslSettingsHttp2 = "off"
)

// The minimum TLS version supported.
type CustomHostnameResponseCollectionResultSslSettingsMinTlsVersion string

const (
	CustomHostnameResponseCollectionResultSslSettingsMinTlsVersion1_0 CustomHostnameResponseCollectionResultSslSettingsMinTlsVersion = "1.0"
	CustomHostnameResponseCollectionResultSslSettingsMinTlsVersion1_1 CustomHostnameResponseCollectionResultSslSettingsMinTlsVersion = "1.1"
	CustomHostnameResponseCollectionResultSslSettingsMinTlsVersion1_2 CustomHostnameResponseCollectionResultSslSettingsMinTlsVersion = "1.2"
	CustomHostnameResponseCollectionResultSslSettingsMinTlsVersion1_3 CustomHostnameResponseCollectionResultSslSettingsMinTlsVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type CustomHostnameResponseCollectionResultSslSettingsTls1_3 string

const (
	CustomHostnameResponseCollectionResultSslSettingsTls1_3On  CustomHostnameResponseCollectionResultSslSettingsTls1_3 = "on"
	CustomHostnameResponseCollectionResultSslSettingsTls1_3Off CustomHostnameResponseCollectionResultSslSettingsTls1_3 = "off"
)

// Status of the hostname's SSL certificates.
type CustomHostnameResponseCollectionResultSslStatus string

const (
	CustomHostnameResponseCollectionResultSslStatusInitializing         CustomHostnameResponseCollectionResultSslStatus = "initializing"
	CustomHostnameResponseCollectionResultSslStatusPendingValidation    CustomHostnameResponseCollectionResultSslStatus = "pending_validation"
	CustomHostnameResponseCollectionResultSslStatusDeleted              CustomHostnameResponseCollectionResultSslStatus = "deleted"
	CustomHostnameResponseCollectionResultSslStatusPendingIssuance      CustomHostnameResponseCollectionResultSslStatus = "pending_issuance"
	CustomHostnameResponseCollectionResultSslStatusPendingDeployment    CustomHostnameResponseCollectionResultSslStatus = "pending_deployment"
	CustomHostnameResponseCollectionResultSslStatusPendingDeletion      CustomHostnameResponseCollectionResultSslStatus = "pending_deletion"
	CustomHostnameResponseCollectionResultSslStatusPendingExpiration    CustomHostnameResponseCollectionResultSslStatus = "pending_expiration"
	CustomHostnameResponseCollectionResultSslStatusExpired              CustomHostnameResponseCollectionResultSslStatus = "expired"
	CustomHostnameResponseCollectionResultSslStatusActive               CustomHostnameResponseCollectionResultSslStatus = "active"
	CustomHostnameResponseCollectionResultSslStatusInitializingTimedOut CustomHostnameResponseCollectionResultSslStatus = "initializing_timed_out"
	CustomHostnameResponseCollectionResultSslStatusValidationTimedOut   CustomHostnameResponseCollectionResultSslStatus = "validation_timed_out"
	CustomHostnameResponseCollectionResultSslStatusIssuanceTimedOut     CustomHostnameResponseCollectionResultSslStatus = "issuance_timed_out"
	CustomHostnameResponseCollectionResultSslStatusDeploymentTimedOut   CustomHostnameResponseCollectionResultSslStatus = "deployment_timed_out"
	CustomHostnameResponseCollectionResultSslStatusDeletionTimedOut     CustomHostnameResponseCollectionResultSslStatus = "deletion_timed_out"
	CustomHostnameResponseCollectionResultSslStatusPendingCleanup       CustomHostnameResponseCollectionResultSslStatus = "pending_cleanup"
	CustomHostnameResponseCollectionResultSslStatusStagingDeployment    CustomHostnameResponseCollectionResultSslStatus = "staging_deployment"
	CustomHostnameResponseCollectionResultSslStatusStagingActive        CustomHostnameResponseCollectionResultSslStatus = "staging_active"
	CustomHostnameResponseCollectionResultSslStatusDeactivating         CustomHostnameResponseCollectionResultSslStatus = "deactivating"
	CustomHostnameResponseCollectionResultSslStatusInactive             CustomHostnameResponseCollectionResultSslStatus = "inactive"
	CustomHostnameResponseCollectionResultSslStatusBackupIssued         CustomHostnameResponseCollectionResultSslStatus = "backup_issued"
	CustomHostnameResponseCollectionResultSslStatusHoldingDeployment    CustomHostnameResponseCollectionResultSslStatus = "holding_deployment"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type CustomHostnameResponseCollectionResultSslType string

const (
	CustomHostnameResponseCollectionResultSslTypeDv CustomHostnameResponseCollectionResultSslType = "dv"
)

type CustomHostnameResponseCollectionResultSslValidationError struct {
	// A domain validation error.
	Message string                                                       `json:"message"`
	JSON    customHostnameResponseCollectionResultSslValidationErrorJSON `json:"-"`
}

// customHostnameResponseCollectionResultSslValidationErrorJSON contains the JSON
// metadata for the struct
// [CustomHostnameResponseCollectionResultSslValidationError]
type customHostnameResponseCollectionResultSslValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseCollectionResultSslValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate's required validation record.
type CustomHostnameResponseCollectionResultSslValidationRecord struct {
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
	TxtValue string                                                        `json:"txt_value"`
	JSON     customHostnameResponseCollectionResultSslValidationRecordJSON `json:"-"`
}

// customHostnameResponseCollectionResultSslValidationRecordJSON contains the JSON
// metadata for the struct
// [CustomHostnameResponseCollectionResultSslValidationRecord]
type customHostnameResponseCollectionResultSslValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TxtName     apijson.Field
	TxtValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseCollectionResultSslValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type CustomHostnameResponseCollectionResultStatus string

const (
	CustomHostnameResponseCollectionResultStatusActive             CustomHostnameResponseCollectionResultStatus = "active"
	CustomHostnameResponseCollectionResultStatusPending            CustomHostnameResponseCollectionResultStatus = "pending"
	CustomHostnameResponseCollectionResultStatusActiveRedeploying  CustomHostnameResponseCollectionResultStatus = "active_redeploying"
	CustomHostnameResponseCollectionResultStatusMoved              CustomHostnameResponseCollectionResultStatus = "moved"
	CustomHostnameResponseCollectionResultStatusPendingDeletion    CustomHostnameResponseCollectionResultStatus = "pending_deletion"
	CustomHostnameResponseCollectionResultStatusDeleted            CustomHostnameResponseCollectionResultStatus = "deleted"
	CustomHostnameResponseCollectionResultStatusPendingBlocked     CustomHostnameResponseCollectionResultStatus = "pending_blocked"
	CustomHostnameResponseCollectionResultStatusPendingMigration   CustomHostnameResponseCollectionResultStatus = "pending_migration"
	CustomHostnameResponseCollectionResultStatusPendingProvisioned CustomHostnameResponseCollectionResultStatus = "pending_provisioned"
	CustomHostnameResponseCollectionResultStatusTestPending        CustomHostnameResponseCollectionResultStatus = "test_pending"
	CustomHostnameResponseCollectionResultStatusTestActive         CustomHostnameResponseCollectionResultStatus = "test_active"
	CustomHostnameResponseCollectionResultStatusTestActiveApex     CustomHostnameResponseCollectionResultStatus = "test_active_apex"
	CustomHostnameResponseCollectionResultStatusTestBlocked        CustomHostnameResponseCollectionResultStatus = "test_blocked"
	CustomHostnameResponseCollectionResultStatusTestFailed         CustomHostnameResponseCollectionResultStatus = "test_failed"
	CustomHostnameResponseCollectionResultStatusProvisioned        CustomHostnameResponseCollectionResultStatus = "provisioned"
	CustomHostnameResponseCollectionResultStatusBlocked            CustomHostnameResponseCollectionResultStatus = "blocked"
)

type CustomHostnameResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       customHostnameResponseCollectionResultInfoJSON `json:"-"`
}

// customHostnameResponseCollectionResultInfoJSON contains the JSON metadata for
// the struct [CustomHostnameResponseCollectionResultInfo]
type customHostnameResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomHostnameResponseCollectionSuccess bool

const (
	CustomHostnameResponseCollectionSuccessTrue CustomHostnameResponseCollectionSuccess = true
)

type CustomHostnameResponseSingle4Byp7wot struct {
	Errors   []CustomHostnameResponseSingle4Byp7wotError   `json:"errors"`
	Messages []CustomHostnameResponseSingle4Byp7wotMessage `json:"messages"`
	Result   CustomHostnameResponseSingle4Byp7wotResult    `json:"result"`
	// Whether the API call was successful
	Success CustomHostnameResponseSingle4Byp7wotSuccess `json:"success"`
	JSON    customHostnameResponseSingle4Byp7wotJSON    `json:"-"`
}

// customHostnameResponseSingle4Byp7wotJSON contains the JSON metadata for the
// struct [CustomHostnameResponseSingle4Byp7wot]
type customHostnameResponseSingle4Byp7wotJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseSingle4Byp7wot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameResponseSingle4Byp7wotError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    customHostnameResponseSingle4Byp7wotErrorJSON `json:"-"`
}

// customHostnameResponseSingle4Byp7wotErrorJSON contains the JSON metadata for the
// struct [CustomHostnameResponseSingle4Byp7wotError]
type customHostnameResponseSingle4Byp7wotErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseSingle4Byp7wotError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameResponseSingle4Byp7wotMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    customHostnameResponseSingle4Byp7wotMessageJSON `json:"-"`
}

// customHostnameResponseSingle4Byp7wotMessageJSON contains the JSON metadata for
// the struct [CustomHostnameResponseSingle4Byp7wotMessage]
type customHostnameResponseSingle4Byp7wotMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseSingle4Byp7wotMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomHostnameResponseSingle4Byp7wotResult struct {
	// Identifier
	ID string `json:"id"`
	// This is the time the hostname was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// These are per-hostname (customer) settings.
	CustomMetadata CustomHostnameResponseSingle4Byp7wotResultCustomMetadata `json:"custom_metadata"`
	// a valid hostname that’s been added to your DNS zone as an A, AAAA, or CNAME
	// record.
	CustomOriginServer string `json:"custom_origin_server"`
	// A hostname that will be sent to your custom origin server as SNI for TLS
	// handshake. This can be a valid subdomain of the zone or custom origin server
	// name or the string ':request_host_header:' which will cause the host header in
	// the request to be used as SNI. Not configurable with default/fallback origin
	// server.
	CustomOriginSni string `json:"custom_origin_sni"`
	// The custom hostname that will point to your hostname via CNAME.
	Hostname string `json:"hostname"`
	// This is a record which can be placed to activate a hostname.
	OwnershipVerification CustomHostnameResponseSingle4Byp7wotResultOwnershipVerification `json:"ownership_verification"`
	// This presents the token to be served by the given http url to activate a
	// hostname.
	OwnershipVerificationHTTP CustomHostnameResponseSingle4Byp7wotResultOwnershipVerificationHTTP `json:"ownership_verification_http"`
	// SSL properties for the custom hostname.
	Ssl CustomHostnameResponseSingle4Byp7wotResultSsl `json:"ssl"`
	// Status of the hostname's activation.
	Status CustomHostnameResponseSingle4Byp7wotResultStatus `json:"status"`
	// These are errors that were encountered while trying to activate a hostname.
	VerificationErrors []interface{}                                  `json:"verification_errors"`
	JSON               customHostnameResponseSingle4Byp7wotResultJSON `json:"-"`
}

// customHostnameResponseSingle4Byp7wotResultJSON contains the JSON metadata for
// the struct [CustomHostnameResponseSingle4Byp7wotResult]
type customHostnameResponseSingle4Byp7wotResultJSON struct {
	ID                        apijson.Field
	CreatedAt                 apijson.Field
	CustomMetadata            apijson.Field
	CustomOriginServer        apijson.Field
	CustomOriginSni           apijson.Field
	Hostname                  apijson.Field
	OwnershipVerification     apijson.Field
	OwnershipVerificationHTTP apijson.Field
	Ssl                       apijson.Field
	Status                    apijson.Field
	VerificationErrors        apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *CustomHostnameResponseSingle4Byp7wotResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// These are per-hostname (customer) settings.
type CustomHostnameResponseSingle4Byp7wotResultCustomMetadata struct {
	// Unique metadata for this hostname.
	Key  string                                                       `json:"key"`
	JSON customHostnameResponseSingle4Byp7wotResultCustomMetadataJSON `json:"-"`
}

// customHostnameResponseSingle4Byp7wotResultCustomMetadataJSON contains the JSON
// metadata for the struct
// [CustomHostnameResponseSingle4Byp7wotResultCustomMetadata]
type customHostnameResponseSingle4Byp7wotResultCustomMetadataJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseSingle4Byp7wotResultCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This is a record which can be placed to activate a hostname.
type CustomHostnameResponseSingle4Byp7wotResultOwnershipVerification struct {
	// DNS Name for record.
	Name string `json:"name"`
	// DNS Record type.
	Type CustomHostnameResponseSingle4Byp7wotResultOwnershipVerificationType `json:"type"`
	// Content for the record.
	Value string                                                              `json:"value"`
	JSON  customHostnameResponseSingle4Byp7wotResultOwnershipVerificationJSON `json:"-"`
}

// customHostnameResponseSingle4Byp7wotResultOwnershipVerificationJSON contains the
// JSON metadata for the struct
// [CustomHostnameResponseSingle4Byp7wotResultOwnershipVerification]
type customHostnameResponseSingle4Byp7wotResultOwnershipVerificationJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseSingle4Byp7wotResultOwnershipVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS Record type.
type CustomHostnameResponseSingle4Byp7wotResultOwnershipVerificationType string

const (
	CustomHostnameResponseSingle4Byp7wotResultOwnershipVerificationTypeTxt CustomHostnameResponseSingle4Byp7wotResultOwnershipVerificationType = "txt"
)

// This presents the token to be served by the given http url to activate a
// hostname.
type CustomHostnameResponseSingle4Byp7wotResultOwnershipVerificationHTTP struct {
	// Token to be served.
	HTTPBody string `json:"http_body"`
	// The HTTP URL that will be checked during custom hostname verification and where
	// the customer should host the token.
	HTTPURL string                                                                  `json:"http_url"`
	JSON    customHostnameResponseSingle4Byp7wotResultOwnershipVerificationHTTPJSON `json:"-"`
}

// customHostnameResponseSingle4Byp7wotResultOwnershipVerificationHTTPJSON contains
// the JSON metadata for the struct
// [CustomHostnameResponseSingle4Byp7wotResultOwnershipVerificationHTTP]
type customHostnameResponseSingle4Byp7wotResultOwnershipVerificationHTTPJSON struct {
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseSingle4Byp7wotResultOwnershipVerificationHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// SSL properties for the custom hostname.
type CustomHostnameResponseSingle4Byp7wotResultSsl struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod CustomHostnameResponseSingle4Byp7wotResultSslBundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority CustomHostnameResponseSingle4Byp7wotResultSslCertificateAuthority `json:"certificate_authority"`
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
	Method CustomHostnameResponseSingle4Byp7wotResultSslMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// SSL specific settings.
	Settings CustomHostnameResponseSingle4Byp7wotResultSslSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status CustomHostnameResponseSingle4Byp7wotResultSslStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type CustomHostnameResponseSingle4Byp7wotResultSslType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []CustomHostnameResponseSingle4Byp7wotResultSslValidationError  `json:"validation_errors"`
	ValidationRecords []CustomHostnameResponseSingle4Byp7wotResultSslValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                                              `json:"wildcard"`
	JSON     customHostnameResponseSingle4Byp7wotResultSslJSON `json:"-"`
}

// customHostnameResponseSingle4Byp7wotResultSslJSON contains the JSON metadata for
// the struct [CustomHostnameResponseSingle4Byp7wotResultSsl]
type customHostnameResponseSingle4Byp7wotResultSslJSON struct {
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

func (r *CustomHostnameResponseSingle4Byp7wotResultSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type CustomHostnameResponseSingle4Byp7wotResultSslBundleMethod string

const (
	CustomHostnameResponseSingle4Byp7wotResultSslBundleMethodUbiquitous CustomHostnameResponseSingle4Byp7wotResultSslBundleMethod = "ubiquitous"
	CustomHostnameResponseSingle4Byp7wotResultSslBundleMethodOptimal    CustomHostnameResponseSingle4Byp7wotResultSslBundleMethod = "optimal"
	CustomHostnameResponseSingle4Byp7wotResultSslBundleMethodForce      CustomHostnameResponseSingle4Byp7wotResultSslBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type CustomHostnameResponseSingle4Byp7wotResultSslCertificateAuthority string

const (
	CustomHostnameResponseSingle4Byp7wotResultSslCertificateAuthorityDigicert    CustomHostnameResponseSingle4Byp7wotResultSslCertificateAuthority = "digicert"
	CustomHostnameResponseSingle4Byp7wotResultSslCertificateAuthorityGoogle      CustomHostnameResponseSingle4Byp7wotResultSslCertificateAuthority = "google"
	CustomHostnameResponseSingle4Byp7wotResultSslCertificateAuthorityLetsEncrypt CustomHostnameResponseSingle4Byp7wotResultSslCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type CustomHostnameResponseSingle4Byp7wotResultSslMethod string

const (
	CustomHostnameResponseSingle4Byp7wotResultSslMethodHTTP  CustomHostnameResponseSingle4Byp7wotResultSslMethod = "http"
	CustomHostnameResponseSingle4Byp7wotResultSslMethodTxt   CustomHostnameResponseSingle4Byp7wotResultSslMethod = "txt"
	CustomHostnameResponseSingle4Byp7wotResultSslMethodEmail CustomHostnameResponseSingle4Byp7wotResultSslMethod = "email"
)

// SSL specific settings.
type CustomHostnameResponseSingle4Byp7wotResultSslSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints CustomHostnameResponseSingle4Byp7wotResultSslSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	Http2 CustomHostnameResponseSingle4Byp7wotResultSslSettingsHttp2 `json:"http2"`
	// The minimum TLS version supported.
	MinTlsVersion CustomHostnameResponseSingle4Byp7wotResultSslSettingsMinTlsVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	Tls1_3 CustomHostnameResponseSingle4Byp7wotResultSslSettingsTls1_3 `json:"tls_1_3"`
	JSON   customHostnameResponseSingle4Byp7wotResultSslSettingsJSON   `json:"-"`
}

// customHostnameResponseSingle4Byp7wotResultSslSettingsJSON contains the JSON
// metadata for the struct [CustomHostnameResponseSingle4Byp7wotResultSslSettings]
type customHostnameResponseSingle4Byp7wotResultSslSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	Http2         apijson.Field
	MinTlsVersion apijson.Field
	Tls1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CustomHostnameResponseSingle4Byp7wotResultSslSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not Early Hints is enabled.
type CustomHostnameResponseSingle4Byp7wotResultSslSettingsEarlyHints string

const (
	CustomHostnameResponseSingle4Byp7wotResultSslSettingsEarlyHintsOn  CustomHostnameResponseSingle4Byp7wotResultSslSettingsEarlyHints = "on"
	CustomHostnameResponseSingle4Byp7wotResultSslSettingsEarlyHintsOff CustomHostnameResponseSingle4Byp7wotResultSslSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type CustomHostnameResponseSingle4Byp7wotResultSslSettingsHttp2 string

const (
	CustomHostnameResponseSingle4Byp7wotResultSslSettingsHttp2On  CustomHostnameResponseSingle4Byp7wotResultSslSettingsHttp2 = "on"
	CustomHostnameResponseSingle4Byp7wotResultSslSettingsHttp2Off CustomHostnameResponseSingle4Byp7wotResultSslSettingsHttp2 = "off"
)

// The minimum TLS version supported.
type CustomHostnameResponseSingle4Byp7wotResultSslSettingsMinTlsVersion string

const (
	CustomHostnameResponseSingle4Byp7wotResultSslSettingsMinTlsVersion1_0 CustomHostnameResponseSingle4Byp7wotResultSslSettingsMinTlsVersion = "1.0"
	CustomHostnameResponseSingle4Byp7wotResultSslSettingsMinTlsVersion1_1 CustomHostnameResponseSingle4Byp7wotResultSslSettingsMinTlsVersion = "1.1"
	CustomHostnameResponseSingle4Byp7wotResultSslSettingsMinTlsVersion1_2 CustomHostnameResponseSingle4Byp7wotResultSslSettingsMinTlsVersion = "1.2"
	CustomHostnameResponseSingle4Byp7wotResultSslSettingsMinTlsVersion1_3 CustomHostnameResponseSingle4Byp7wotResultSslSettingsMinTlsVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type CustomHostnameResponseSingle4Byp7wotResultSslSettingsTls1_3 string

const (
	CustomHostnameResponseSingle4Byp7wotResultSslSettingsTls1_3On  CustomHostnameResponseSingle4Byp7wotResultSslSettingsTls1_3 = "on"
	CustomHostnameResponseSingle4Byp7wotResultSslSettingsTls1_3Off CustomHostnameResponseSingle4Byp7wotResultSslSettingsTls1_3 = "off"
)

// Status of the hostname's SSL certificates.
type CustomHostnameResponseSingle4Byp7wotResultSslStatus string

const (
	CustomHostnameResponseSingle4Byp7wotResultSslStatusInitializing         CustomHostnameResponseSingle4Byp7wotResultSslStatus = "initializing"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusPendingValidation    CustomHostnameResponseSingle4Byp7wotResultSslStatus = "pending_validation"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusDeleted              CustomHostnameResponseSingle4Byp7wotResultSslStatus = "deleted"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusPendingIssuance      CustomHostnameResponseSingle4Byp7wotResultSslStatus = "pending_issuance"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusPendingDeployment    CustomHostnameResponseSingle4Byp7wotResultSslStatus = "pending_deployment"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusPendingDeletion      CustomHostnameResponseSingle4Byp7wotResultSslStatus = "pending_deletion"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusPendingExpiration    CustomHostnameResponseSingle4Byp7wotResultSslStatus = "pending_expiration"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusExpired              CustomHostnameResponseSingle4Byp7wotResultSslStatus = "expired"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusActive               CustomHostnameResponseSingle4Byp7wotResultSslStatus = "active"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusInitializingTimedOut CustomHostnameResponseSingle4Byp7wotResultSslStatus = "initializing_timed_out"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusValidationTimedOut   CustomHostnameResponseSingle4Byp7wotResultSslStatus = "validation_timed_out"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusIssuanceTimedOut     CustomHostnameResponseSingle4Byp7wotResultSslStatus = "issuance_timed_out"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusDeploymentTimedOut   CustomHostnameResponseSingle4Byp7wotResultSslStatus = "deployment_timed_out"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusDeletionTimedOut     CustomHostnameResponseSingle4Byp7wotResultSslStatus = "deletion_timed_out"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusPendingCleanup       CustomHostnameResponseSingle4Byp7wotResultSslStatus = "pending_cleanup"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusStagingDeployment    CustomHostnameResponseSingle4Byp7wotResultSslStatus = "staging_deployment"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusStagingActive        CustomHostnameResponseSingle4Byp7wotResultSslStatus = "staging_active"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusDeactivating         CustomHostnameResponseSingle4Byp7wotResultSslStatus = "deactivating"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusInactive             CustomHostnameResponseSingle4Byp7wotResultSslStatus = "inactive"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusBackupIssued         CustomHostnameResponseSingle4Byp7wotResultSslStatus = "backup_issued"
	CustomHostnameResponseSingle4Byp7wotResultSslStatusHoldingDeployment    CustomHostnameResponseSingle4Byp7wotResultSslStatus = "holding_deployment"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type CustomHostnameResponseSingle4Byp7wotResultSslType string

const (
	CustomHostnameResponseSingle4Byp7wotResultSslTypeDv CustomHostnameResponseSingle4Byp7wotResultSslType = "dv"
)

type CustomHostnameResponseSingle4Byp7wotResultSslValidationError struct {
	// A domain validation error.
	Message string                                                           `json:"message"`
	JSON    customHostnameResponseSingle4Byp7wotResultSslValidationErrorJSON `json:"-"`
}

// customHostnameResponseSingle4Byp7wotResultSslValidationErrorJSON contains the
// JSON metadata for the struct
// [CustomHostnameResponseSingle4Byp7wotResultSslValidationError]
type customHostnameResponseSingle4Byp7wotResultSslValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseSingle4Byp7wotResultSslValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate's required validation record.
type CustomHostnameResponseSingle4Byp7wotResultSslValidationRecord struct {
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
	TxtValue string                                                            `json:"txt_value"`
	JSON     customHostnameResponseSingle4Byp7wotResultSslValidationRecordJSON `json:"-"`
}

// customHostnameResponseSingle4Byp7wotResultSslValidationRecordJSON contains the
// JSON metadata for the struct
// [CustomHostnameResponseSingle4Byp7wotResultSslValidationRecord]
type customHostnameResponseSingle4Byp7wotResultSslValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TxtName     apijson.Field
	TxtValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomHostnameResponseSingle4Byp7wotResultSslValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type CustomHostnameResponseSingle4Byp7wotResultStatus string

const (
	CustomHostnameResponseSingle4Byp7wotResultStatusActive             CustomHostnameResponseSingle4Byp7wotResultStatus = "active"
	CustomHostnameResponseSingle4Byp7wotResultStatusPending            CustomHostnameResponseSingle4Byp7wotResultStatus = "pending"
	CustomHostnameResponseSingle4Byp7wotResultStatusActiveRedeploying  CustomHostnameResponseSingle4Byp7wotResultStatus = "active_redeploying"
	CustomHostnameResponseSingle4Byp7wotResultStatusMoved              CustomHostnameResponseSingle4Byp7wotResultStatus = "moved"
	CustomHostnameResponseSingle4Byp7wotResultStatusPendingDeletion    CustomHostnameResponseSingle4Byp7wotResultStatus = "pending_deletion"
	CustomHostnameResponseSingle4Byp7wotResultStatusDeleted            CustomHostnameResponseSingle4Byp7wotResultStatus = "deleted"
	CustomHostnameResponseSingle4Byp7wotResultStatusPendingBlocked     CustomHostnameResponseSingle4Byp7wotResultStatus = "pending_blocked"
	CustomHostnameResponseSingle4Byp7wotResultStatusPendingMigration   CustomHostnameResponseSingle4Byp7wotResultStatus = "pending_migration"
	CustomHostnameResponseSingle4Byp7wotResultStatusPendingProvisioned CustomHostnameResponseSingle4Byp7wotResultStatus = "pending_provisioned"
	CustomHostnameResponseSingle4Byp7wotResultStatusTestPending        CustomHostnameResponseSingle4Byp7wotResultStatus = "test_pending"
	CustomHostnameResponseSingle4Byp7wotResultStatusTestActive         CustomHostnameResponseSingle4Byp7wotResultStatus = "test_active"
	CustomHostnameResponseSingle4Byp7wotResultStatusTestActiveApex     CustomHostnameResponseSingle4Byp7wotResultStatus = "test_active_apex"
	CustomHostnameResponseSingle4Byp7wotResultStatusTestBlocked        CustomHostnameResponseSingle4Byp7wotResultStatus = "test_blocked"
	CustomHostnameResponseSingle4Byp7wotResultStatusTestFailed         CustomHostnameResponseSingle4Byp7wotResultStatus = "test_failed"
	CustomHostnameResponseSingle4Byp7wotResultStatusProvisioned        CustomHostnameResponseSingle4Byp7wotResultStatus = "provisioned"
	CustomHostnameResponseSingle4Byp7wotResultStatusBlocked            CustomHostnameResponseSingle4Byp7wotResultStatus = "blocked"
)

// Whether the API call was successful
type CustomHostnameResponseSingle4Byp7wotSuccess bool

const (
	CustomHostnameResponseSingle4Byp7wotSuccessTrue CustomHostnameResponseSingle4Byp7wotSuccess = true
)

type ZoneCustomHostnameDeleteResponse struct {
	// Identifier
	ID   string                               `json:"id"`
	JSON zoneCustomHostnameDeleteResponseJSON `json:"-"`
}

// zoneCustomHostnameDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneCustomHostnameDeleteResponse]
type zoneCustomHostnameDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameUpdateParams struct {
	// These are per-hostname (customer) settings.
	CustomMetadata param.Field[ZoneCustomHostnameUpdateParamsCustomMetadata] `json:"custom_metadata"`
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
	Ssl param.Field[ZoneCustomHostnameUpdateParamsSsl] `json:"ssl"`
}

func (r ZoneCustomHostnameUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// These are per-hostname (customer) settings.
type ZoneCustomHostnameUpdateParamsCustomMetadata struct {
	// Unique metadata for this hostname.
	Key param.Field[string] `json:"key"`
}

func (r ZoneCustomHostnameUpdateParamsCustomMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SSL properties used when creating the custom hostname.
type ZoneCustomHostnameUpdateParamsSsl struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[ZoneCustomHostnameUpdateParamsSslBundleMethod] `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority param.Field[ZoneCustomHostnameUpdateParamsSslCertificateAuthority] `json:"certificate_authority"`
	// If a custom uploaded certificate is used.
	CustomCertificate param.Field[string] `json:"custom_certificate"`
	// The key for a custom uploaded certificate.
	CustomKey param.Field[string] `json:"custom_key"`
	// Domain control validation (DCV) method used for this hostname.
	Method param.Field[ZoneCustomHostnameUpdateParamsSslMethod] `json:"method"`
	// SSL specific settings.
	Settings param.Field[ZoneCustomHostnameUpdateParamsSslSettings] `json:"settings"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type param.Field[ZoneCustomHostnameUpdateParamsSslType] `json:"type"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard param.Field[bool] `json:"wildcard"`
}

func (r ZoneCustomHostnameUpdateParamsSsl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type ZoneCustomHostnameUpdateParamsSslBundleMethod string

const (
	ZoneCustomHostnameUpdateParamsSslBundleMethodUbiquitous ZoneCustomHostnameUpdateParamsSslBundleMethod = "ubiquitous"
	ZoneCustomHostnameUpdateParamsSslBundleMethodOptimal    ZoneCustomHostnameUpdateParamsSslBundleMethod = "optimal"
	ZoneCustomHostnameUpdateParamsSslBundleMethodForce      ZoneCustomHostnameUpdateParamsSslBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type ZoneCustomHostnameUpdateParamsSslCertificateAuthority string

const (
	ZoneCustomHostnameUpdateParamsSslCertificateAuthorityDigicert    ZoneCustomHostnameUpdateParamsSslCertificateAuthority = "digicert"
	ZoneCustomHostnameUpdateParamsSslCertificateAuthorityGoogle      ZoneCustomHostnameUpdateParamsSslCertificateAuthority = "google"
	ZoneCustomHostnameUpdateParamsSslCertificateAuthorityLetsEncrypt ZoneCustomHostnameUpdateParamsSslCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type ZoneCustomHostnameUpdateParamsSslMethod string

const (
	ZoneCustomHostnameUpdateParamsSslMethodHTTP  ZoneCustomHostnameUpdateParamsSslMethod = "http"
	ZoneCustomHostnameUpdateParamsSslMethodTxt   ZoneCustomHostnameUpdateParamsSslMethod = "txt"
	ZoneCustomHostnameUpdateParamsSslMethodEmail ZoneCustomHostnameUpdateParamsSslMethod = "email"
)

// SSL specific settings.
type ZoneCustomHostnameUpdateParamsSslSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers param.Field[[]string] `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints param.Field[ZoneCustomHostnameUpdateParamsSslSettingsEarlyHints] `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	Http2 param.Field[ZoneCustomHostnameUpdateParamsSslSettingsHttp2] `json:"http2"`
	// The minimum TLS version supported.
	MinTlsVersion param.Field[ZoneCustomHostnameUpdateParamsSslSettingsMinTlsVersion] `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	Tls1_3 param.Field[ZoneCustomHostnameUpdateParamsSslSettingsTls1_3] `json:"tls_1_3"`
}

func (r ZoneCustomHostnameUpdateParamsSslSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not Early Hints is enabled.
type ZoneCustomHostnameUpdateParamsSslSettingsEarlyHints string

const (
	ZoneCustomHostnameUpdateParamsSslSettingsEarlyHintsOn  ZoneCustomHostnameUpdateParamsSslSettingsEarlyHints = "on"
	ZoneCustomHostnameUpdateParamsSslSettingsEarlyHintsOff ZoneCustomHostnameUpdateParamsSslSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type ZoneCustomHostnameUpdateParamsSslSettingsHttp2 string

const (
	ZoneCustomHostnameUpdateParamsSslSettingsHttp2On  ZoneCustomHostnameUpdateParamsSslSettingsHttp2 = "on"
	ZoneCustomHostnameUpdateParamsSslSettingsHttp2Off ZoneCustomHostnameUpdateParamsSslSettingsHttp2 = "off"
)

// The minimum TLS version supported.
type ZoneCustomHostnameUpdateParamsSslSettingsMinTlsVersion string

const (
	ZoneCustomHostnameUpdateParamsSslSettingsMinTlsVersion1_0 ZoneCustomHostnameUpdateParamsSslSettingsMinTlsVersion = "1.0"
	ZoneCustomHostnameUpdateParamsSslSettingsMinTlsVersion1_1 ZoneCustomHostnameUpdateParamsSslSettingsMinTlsVersion = "1.1"
	ZoneCustomHostnameUpdateParamsSslSettingsMinTlsVersion1_2 ZoneCustomHostnameUpdateParamsSslSettingsMinTlsVersion = "1.2"
	ZoneCustomHostnameUpdateParamsSslSettingsMinTlsVersion1_3 ZoneCustomHostnameUpdateParamsSslSettingsMinTlsVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type ZoneCustomHostnameUpdateParamsSslSettingsTls1_3 string

const (
	ZoneCustomHostnameUpdateParamsSslSettingsTls1_3On  ZoneCustomHostnameUpdateParamsSslSettingsTls1_3 = "on"
	ZoneCustomHostnameUpdateParamsSslSettingsTls1_3Off ZoneCustomHostnameUpdateParamsSslSettingsTls1_3 = "off"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type ZoneCustomHostnameUpdateParamsSslType string

const (
	ZoneCustomHostnameUpdateParamsSslTypeDv ZoneCustomHostnameUpdateParamsSslType = "dv"
)

type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParams struct {
	// The custom hostname that will point to your hostname via CNAME.
	Hostname param.Field[string] `json:"hostname,required"`
	// SSL properties used when creating the custom hostname.
	Ssl param.Field[ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSsl] `json:"ssl,required"`
	// These are per-hostname (customer) settings.
	CustomMetadata param.Field[ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsCustomMetadata] `json:"custom_metadata"`
}

func (r ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SSL properties used when creating the custom hostname.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSsl struct {
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslBundleMethod] `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority param.Field[ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslCertificateAuthority] `json:"certificate_authority"`
	// If a custom uploaded certificate is used.
	CustomCertificate param.Field[string] `json:"custom_certificate"`
	// The key for a custom uploaded certificate.
	CustomKey param.Field[string] `json:"custom_key"`
	// Domain control validation (DCV) method used for this hostname.
	Method param.Field[ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslMethod] `json:"method"`
	// SSL specific settings.
	Settings param.Field[ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettings] `json:"settings"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type param.Field[ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslType] `json:"type"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard param.Field[bool] `json:"wildcard"`
}

func (r ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSsl) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslBundleMethod string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslBundleMethodUbiquitous ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslBundleMethod = "ubiquitous"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslBundleMethodOptimal    ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslBundleMethod = "optimal"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslBundleMethodForce      ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslCertificateAuthority string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslCertificateAuthorityDigicert    ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslCertificateAuthority = "digicert"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslCertificateAuthorityGoogle      ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslCertificateAuthority = "google"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslCertificateAuthorityLetsEncrypt ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslMethod string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslMethodHTTP  ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslMethod = "http"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslMethodTxt   ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslMethod = "txt"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslMethodEmail ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslMethod = "email"
)

// SSL specific settings.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers param.Field[[]string] `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints param.Field[ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsEarlyHints] `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	Http2 param.Field[ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsHttp2] `json:"http2"`
	// The minimum TLS version supported.
	MinTlsVersion param.Field[ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsMinTlsVersion] `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	Tls1_3 param.Field[ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsTls1_3] `json:"tls_1_3"`
}

func (r ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether or not Early Hints is enabled.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsEarlyHints string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsEarlyHintsOn  ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsEarlyHints = "on"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsEarlyHintsOff ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsHttp2 string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsHttp2On  ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsHttp2 = "on"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsHttp2Off ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsHttp2 = "off"
)

// The minimum TLS version supported.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsMinTlsVersion string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsMinTlsVersion1_0 ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsMinTlsVersion = "1.0"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsMinTlsVersion1_1 ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsMinTlsVersion = "1.1"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsMinTlsVersion1_2 ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsMinTlsVersion = "1.2"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsMinTlsVersion1_3 ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsMinTlsVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsTls1_3 string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsTls1_3On  ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsTls1_3 = "on"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsTls1_3Off ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslSettingsTls1_3 = "off"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslType string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslTypeDv ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsSslType = "dv"
)

// These are per-hostname (customer) settings.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsCustomMetadata struct {
	// Unique metadata for this hostname.
	Key param.Field[string] `json:"key"`
}

func (r ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParamsCustomMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParams struct {
	// Hostname ID to match against. This ID was generated and returned during the
	// initial custom_hostname creation. This parameter cannot be used with the
	// 'hostname' parameter.
	ID param.Field[string] `query:"id"`
	// Direction to order hostnames.
	Direction param.Field[ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsDirection] `query:"direction"`
	// Fully qualified domain name to match against. This parameter cannot be used with
	// the 'id' parameter.
	Hostname param.Field[string] `query:"hostname"`
	// Field to order hostnames by.
	Order param.Field[ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of hostnames per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Whether to filter hostnames based on if they have SSL enabled.
	Ssl param.Field[ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsSsl] `query:"ssl"`
}

// URLQuery serializes
// [ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParams]'s query
// parameters as `url.Values`.
func (r ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order hostnames.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsDirection string

const (
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsDirectionAsc  ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsDirection = "asc"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsDirectionDesc ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsDirection = "desc"
)

// Field to order hostnames by.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsOrder string

const (
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsOrderSsl       ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsOrder = "ssl"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsOrderSslStatus ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsOrder = "ssl_status"
)

// Whether to filter hostnames based on if they have SSL enabled.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsSsl float64

const (
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsSsl0 ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsSsl = 0
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsSsl1 ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParamsSsl = 1
)

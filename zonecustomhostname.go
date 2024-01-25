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
func (r *ZoneCustomHostnameService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneCustomHostnameGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Modify SSL configuration for a custom hostname. When sent with SSL config that
// matches existing config, used to indicate that hostname should pass domain
// control validation (DCV). Can also be used to change validation type, e.g., from
// 'http' to 'email'.
func (r *ZoneCustomHostnameService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneCustomHostnameUpdateParams, opts ...option.RequestOption) (res *ZoneCustomHostnameUpdateResponse, err error) {
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
func (r *ZoneCustomHostnameService) CustomHostnameForAZoneNewCustomHostname(ctx context.Context, zoneIdentifier string, body ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameParams, opts ...option.RequestOption) (res *ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List, search, sort, and filter all of your custom hostnames.
func (r *ZoneCustomHostnameService) CustomHostnameForAZoneListCustomHostnames(ctx context.Context, zoneIdentifier string, query ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesParams, opts ...option.RequestOption) (res *shared.Page[ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/custom_hostnames", zoneIdentifier)
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

type ZoneCustomHostnameGetResponse struct {
	Errors   []ZoneCustomHostnameGetResponseError   `json:"errors"`
	Messages []ZoneCustomHostnameGetResponseMessage `json:"messages"`
	Result   ZoneCustomHostnameGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneCustomHostnameGetResponseSuccess `json:"success"`
	JSON    zoneCustomHostnameGetResponseJSON    `json:"-"`
}

// zoneCustomHostnameGetResponseJSON contains the JSON metadata for the struct
// [ZoneCustomHostnameGetResponse]
type zoneCustomHostnameGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameGetResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneCustomHostnameGetResponseErrorJSON `json:"-"`
}

// zoneCustomHostnameGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneCustomHostnameGetResponseError]
type zoneCustomHostnameGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameGetResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneCustomHostnameGetResponseMessageJSON `json:"-"`
}

// zoneCustomHostnameGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneCustomHostnameGetResponseMessage]
type zoneCustomHostnameGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameGetResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// This is the time the hostname was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// These are per-hostname (customer) settings.
	CustomMetadata ZoneCustomHostnameGetResponseResultCustomMetadata `json:"custom_metadata"`
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
	OwnershipVerification ZoneCustomHostnameGetResponseResultOwnershipVerification `json:"ownership_verification"`
	// This presents the token to be served by the given http url to activate a
	// hostname.
	OwnershipVerificationHTTP ZoneCustomHostnameGetResponseResultOwnershipVerificationHTTP `json:"ownership_verification_http"`
	// SSL properties for the custom hostname.
	Ssl ZoneCustomHostnameGetResponseResultSsl `json:"ssl"`
	// Status of the hostname's activation.
	Status ZoneCustomHostnameGetResponseResultStatus `json:"status"`
	// These are errors that were encountered while trying to activate a hostname.
	VerificationErrors []interface{}                           `json:"verification_errors"`
	JSON               zoneCustomHostnameGetResponseResultJSON `json:"-"`
}

// zoneCustomHostnameGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneCustomHostnameGetResponseResult]
type zoneCustomHostnameGetResponseResultJSON struct {
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

func (r *ZoneCustomHostnameGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// These are per-hostname (customer) settings.
type ZoneCustomHostnameGetResponseResultCustomMetadata struct {
	// Unique metadata for this hostname.
	Key  string                                                `json:"key"`
	JSON zoneCustomHostnameGetResponseResultCustomMetadataJSON `json:"-"`
}

// zoneCustomHostnameGetResponseResultCustomMetadataJSON contains the JSON metadata
// for the struct [ZoneCustomHostnameGetResponseResultCustomMetadata]
type zoneCustomHostnameGetResponseResultCustomMetadataJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameGetResponseResultCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This is a record which can be placed to activate a hostname.
type ZoneCustomHostnameGetResponseResultOwnershipVerification struct {
	// DNS Name for record.
	Name string `json:"name"`
	// DNS Record type.
	Type ZoneCustomHostnameGetResponseResultOwnershipVerificationType `json:"type"`
	// Content for the record.
	Value string                                                       `json:"value"`
	JSON  zoneCustomHostnameGetResponseResultOwnershipVerificationJSON `json:"-"`
}

// zoneCustomHostnameGetResponseResultOwnershipVerificationJSON contains the JSON
// metadata for the struct
// [ZoneCustomHostnameGetResponseResultOwnershipVerification]
type zoneCustomHostnameGetResponseResultOwnershipVerificationJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameGetResponseResultOwnershipVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS Record type.
type ZoneCustomHostnameGetResponseResultOwnershipVerificationType string

const (
	ZoneCustomHostnameGetResponseResultOwnershipVerificationTypeTxt ZoneCustomHostnameGetResponseResultOwnershipVerificationType = "txt"
)

// This presents the token to be served by the given http url to activate a
// hostname.
type ZoneCustomHostnameGetResponseResultOwnershipVerificationHTTP struct {
	// Token to be served.
	HTTPBody string `json:"http_body"`
	// The HTTP URL that will be checked during custom hostname verification and where
	// the customer should host the token.
	HTTPURL string                                                           `json:"http_url"`
	JSON    zoneCustomHostnameGetResponseResultOwnershipVerificationHTTPJSON `json:"-"`
}

// zoneCustomHostnameGetResponseResultOwnershipVerificationHTTPJSON contains the
// JSON metadata for the struct
// [ZoneCustomHostnameGetResponseResultOwnershipVerificationHTTP]
type zoneCustomHostnameGetResponseResultOwnershipVerificationHTTPJSON struct {
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameGetResponseResultOwnershipVerificationHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// SSL properties for the custom hostname.
type ZoneCustomHostnameGetResponseResultSsl struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod ZoneCustomHostnameGetResponseResultSslBundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority ZoneCustomHostnameGetResponseResultSslCertificateAuthority `json:"certificate_authority"`
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
	Method ZoneCustomHostnameGetResponseResultSslMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// SSL specific settings.
	Settings ZoneCustomHostnameGetResponseResultSslSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status ZoneCustomHostnameGetResponseResultSslStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type ZoneCustomHostnameGetResponseResultSslType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []ZoneCustomHostnameGetResponseResultSslValidationError  `json:"validation_errors"`
	ValidationRecords []ZoneCustomHostnameGetResponseResultSslValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                                       `json:"wildcard"`
	JSON     zoneCustomHostnameGetResponseResultSslJSON `json:"-"`
}

// zoneCustomHostnameGetResponseResultSslJSON contains the JSON metadata for the
// struct [ZoneCustomHostnameGetResponseResultSsl]
type zoneCustomHostnameGetResponseResultSslJSON struct {
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

func (r *ZoneCustomHostnameGetResponseResultSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type ZoneCustomHostnameGetResponseResultSslBundleMethod string

const (
	ZoneCustomHostnameGetResponseResultSslBundleMethodUbiquitous ZoneCustomHostnameGetResponseResultSslBundleMethod = "ubiquitous"
	ZoneCustomHostnameGetResponseResultSslBundleMethodOptimal    ZoneCustomHostnameGetResponseResultSslBundleMethod = "optimal"
	ZoneCustomHostnameGetResponseResultSslBundleMethodForce      ZoneCustomHostnameGetResponseResultSslBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type ZoneCustomHostnameGetResponseResultSslCertificateAuthority string

const (
	ZoneCustomHostnameGetResponseResultSslCertificateAuthorityDigicert    ZoneCustomHostnameGetResponseResultSslCertificateAuthority = "digicert"
	ZoneCustomHostnameGetResponseResultSslCertificateAuthorityGoogle      ZoneCustomHostnameGetResponseResultSslCertificateAuthority = "google"
	ZoneCustomHostnameGetResponseResultSslCertificateAuthorityLetsEncrypt ZoneCustomHostnameGetResponseResultSslCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type ZoneCustomHostnameGetResponseResultSslMethod string

const (
	ZoneCustomHostnameGetResponseResultSslMethodHTTP  ZoneCustomHostnameGetResponseResultSslMethod = "http"
	ZoneCustomHostnameGetResponseResultSslMethodTxt   ZoneCustomHostnameGetResponseResultSslMethod = "txt"
	ZoneCustomHostnameGetResponseResultSslMethodEmail ZoneCustomHostnameGetResponseResultSslMethod = "email"
)

// SSL specific settings.
type ZoneCustomHostnameGetResponseResultSslSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints ZoneCustomHostnameGetResponseResultSslSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	Http2 ZoneCustomHostnameGetResponseResultSslSettingsHttp2 `json:"http2"`
	// The minimum TLS version supported.
	MinTlsVersion ZoneCustomHostnameGetResponseResultSslSettingsMinTlsVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	Tls1_3 ZoneCustomHostnameGetResponseResultSslSettingsTls1_3 `json:"tls_1_3"`
	JSON   zoneCustomHostnameGetResponseResultSslSettingsJSON   `json:"-"`
}

// zoneCustomHostnameGetResponseResultSslSettingsJSON contains the JSON metadata
// for the struct [ZoneCustomHostnameGetResponseResultSslSettings]
type zoneCustomHostnameGetResponseResultSslSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	Http2         apijson.Field
	MinTlsVersion apijson.Field
	Tls1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneCustomHostnameGetResponseResultSslSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not Early Hints is enabled.
type ZoneCustomHostnameGetResponseResultSslSettingsEarlyHints string

const (
	ZoneCustomHostnameGetResponseResultSslSettingsEarlyHintsOn  ZoneCustomHostnameGetResponseResultSslSettingsEarlyHints = "on"
	ZoneCustomHostnameGetResponseResultSslSettingsEarlyHintsOff ZoneCustomHostnameGetResponseResultSslSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type ZoneCustomHostnameGetResponseResultSslSettingsHttp2 string

const (
	ZoneCustomHostnameGetResponseResultSslSettingsHttp2On  ZoneCustomHostnameGetResponseResultSslSettingsHttp2 = "on"
	ZoneCustomHostnameGetResponseResultSslSettingsHttp2Off ZoneCustomHostnameGetResponseResultSslSettingsHttp2 = "off"
)

// The minimum TLS version supported.
type ZoneCustomHostnameGetResponseResultSslSettingsMinTlsVersion string

const (
	ZoneCustomHostnameGetResponseResultSslSettingsMinTlsVersion1_0 ZoneCustomHostnameGetResponseResultSslSettingsMinTlsVersion = "1.0"
	ZoneCustomHostnameGetResponseResultSslSettingsMinTlsVersion1_1 ZoneCustomHostnameGetResponseResultSslSettingsMinTlsVersion = "1.1"
	ZoneCustomHostnameGetResponseResultSslSettingsMinTlsVersion1_2 ZoneCustomHostnameGetResponseResultSslSettingsMinTlsVersion = "1.2"
	ZoneCustomHostnameGetResponseResultSslSettingsMinTlsVersion1_3 ZoneCustomHostnameGetResponseResultSslSettingsMinTlsVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type ZoneCustomHostnameGetResponseResultSslSettingsTls1_3 string

const (
	ZoneCustomHostnameGetResponseResultSslSettingsTls1_3On  ZoneCustomHostnameGetResponseResultSslSettingsTls1_3 = "on"
	ZoneCustomHostnameGetResponseResultSslSettingsTls1_3Off ZoneCustomHostnameGetResponseResultSslSettingsTls1_3 = "off"
)

// Status of the hostname's SSL certificates.
type ZoneCustomHostnameGetResponseResultSslStatus string

const (
	ZoneCustomHostnameGetResponseResultSslStatusInitializing         ZoneCustomHostnameGetResponseResultSslStatus = "initializing"
	ZoneCustomHostnameGetResponseResultSslStatusPendingValidation    ZoneCustomHostnameGetResponseResultSslStatus = "pending_validation"
	ZoneCustomHostnameGetResponseResultSslStatusDeleted              ZoneCustomHostnameGetResponseResultSslStatus = "deleted"
	ZoneCustomHostnameGetResponseResultSslStatusPendingIssuance      ZoneCustomHostnameGetResponseResultSslStatus = "pending_issuance"
	ZoneCustomHostnameGetResponseResultSslStatusPendingDeployment    ZoneCustomHostnameGetResponseResultSslStatus = "pending_deployment"
	ZoneCustomHostnameGetResponseResultSslStatusPendingDeletion      ZoneCustomHostnameGetResponseResultSslStatus = "pending_deletion"
	ZoneCustomHostnameGetResponseResultSslStatusPendingExpiration    ZoneCustomHostnameGetResponseResultSslStatus = "pending_expiration"
	ZoneCustomHostnameGetResponseResultSslStatusExpired              ZoneCustomHostnameGetResponseResultSslStatus = "expired"
	ZoneCustomHostnameGetResponseResultSslStatusActive               ZoneCustomHostnameGetResponseResultSslStatus = "active"
	ZoneCustomHostnameGetResponseResultSslStatusInitializingTimedOut ZoneCustomHostnameGetResponseResultSslStatus = "initializing_timed_out"
	ZoneCustomHostnameGetResponseResultSslStatusValidationTimedOut   ZoneCustomHostnameGetResponseResultSslStatus = "validation_timed_out"
	ZoneCustomHostnameGetResponseResultSslStatusIssuanceTimedOut     ZoneCustomHostnameGetResponseResultSslStatus = "issuance_timed_out"
	ZoneCustomHostnameGetResponseResultSslStatusDeploymentTimedOut   ZoneCustomHostnameGetResponseResultSslStatus = "deployment_timed_out"
	ZoneCustomHostnameGetResponseResultSslStatusDeletionTimedOut     ZoneCustomHostnameGetResponseResultSslStatus = "deletion_timed_out"
	ZoneCustomHostnameGetResponseResultSslStatusPendingCleanup       ZoneCustomHostnameGetResponseResultSslStatus = "pending_cleanup"
	ZoneCustomHostnameGetResponseResultSslStatusStagingDeployment    ZoneCustomHostnameGetResponseResultSslStatus = "staging_deployment"
	ZoneCustomHostnameGetResponseResultSslStatusStagingActive        ZoneCustomHostnameGetResponseResultSslStatus = "staging_active"
	ZoneCustomHostnameGetResponseResultSslStatusDeactivating         ZoneCustomHostnameGetResponseResultSslStatus = "deactivating"
	ZoneCustomHostnameGetResponseResultSslStatusInactive             ZoneCustomHostnameGetResponseResultSslStatus = "inactive"
	ZoneCustomHostnameGetResponseResultSslStatusBackupIssued         ZoneCustomHostnameGetResponseResultSslStatus = "backup_issued"
	ZoneCustomHostnameGetResponseResultSslStatusHoldingDeployment    ZoneCustomHostnameGetResponseResultSslStatus = "holding_deployment"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type ZoneCustomHostnameGetResponseResultSslType string

const (
	ZoneCustomHostnameGetResponseResultSslTypeDv ZoneCustomHostnameGetResponseResultSslType = "dv"
)

type ZoneCustomHostnameGetResponseResultSslValidationError struct {
	// A domain validation error.
	Message string                                                    `json:"message"`
	JSON    zoneCustomHostnameGetResponseResultSslValidationErrorJSON `json:"-"`
}

// zoneCustomHostnameGetResponseResultSslValidationErrorJSON contains the JSON
// metadata for the struct [ZoneCustomHostnameGetResponseResultSslValidationError]
type zoneCustomHostnameGetResponseResultSslValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameGetResponseResultSslValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate's required validation record.
type ZoneCustomHostnameGetResponseResultSslValidationRecord struct {
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
	TxtValue string                                                     `json:"txt_value"`
	JSON     zoneCustomHostnameGetResponseResultSslValidationRecordJSON `json:"-"`
}

// zoneCustomHostnameGetResponseResultSslValidationRecordJSON contains the JSON
// metadata for the struct [ZoneCustomHostnameGetResponseResultSslValidationRecord]
type zoneCustomHostnameGetResponseResultSslValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TxtName     apijson.Field
	TxtValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameGetResponseResultSslValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type ZoneCustomHostnameGetResponseResultStatus string

const (
	ZoneCustomHostnameGetResponseResultStatusActive             ZoneCustomHostnameGetResponseResultStatus = "active"
	ZoneCustomHostnameGetResponseResultStatusPending            ZoneCustomHostnameGetResponseResultStatus = "pending"
	ZoneCustomHostnameGetResponseResultStatusActiveRedeploying  ZoneCustomHostnameGetResponseResultStatus = "active_redeploying"
	ZoneCustomHostnameGetResponseResultStatusMoved              ZoneCustomHostnameGetResponseResultStatus = "moved"
	ZoneCustomHostnameGetResponseResultStatusPendingDeletion    ZoneCustomHostnameGetResponseResultStatus = "pending_deletion"
	ZoneCustomHostnameGetResponseResultStatusDeleted            ZoneCustomHostnameGetResponseResultStatus = "deleted"
	ZoneCustomHostnameGetResponseResultStatusPendingBlocked     ZoneCustomHostnameGetResponseResultStatus = "pending_blocked"
	ZoneCustomHostnameGetResponseResultStatusPendingMigration   ZoneCustomHostnameGetResponseResultStatus = "pending_migration"
	ZoneCustomHostnameGetResponseResultStatusPendingProvisioned ZoneCustomHostnameGetResponseResultStatus = "pending_provisioned"
	ZoneCustomHostnameGetResponseResultStatusTestPending        ZoneCustomHostnameGetResponseResultStatus = "test_pending"
	ZoneCustomHostnameGetResponseResultStatusTestActive         ZoneCustomHostnameGetResponseResultStatus = "test_active"
	ZoneCustomHostnameGetResponseResultStatusTestActiveApex     ZoneCustomHostnameGetResponseResultStatus = "test_active_apex"
	ZoneCustomHostnameGetResponseResultStatusTestBlocked        ZoneCustomHostnameGetResponseResultStatus = "test_blocked"
	ZoneCustomHostnameGetResponseResultStatusTestFailed         ZoneCustomHostnameGetResponseResultStatus = "test_failed"
	ZoneCustomHostnameGetResponseResultStatusProvisioned        ZoneCustomHostnameGetResponseResultStatus = "provisioned"
	ZoneCustomHostnameGetResponseResultStatusBlocked            ZoneCustomHostnameGetResponseResultStatus = "blocked"
)

// Whether the API call was successful
type ZoneCustomHostnameGetResponseSuccess bool

const (
	ZoneCustomHostnameGetResponseSuccessTrue ZoneCustomHostnameGetResponseSuccess = true
)

type ZoneCustomHostnameUpdateResponse struct {
	Errors   []ZoneCustomHostnameUpdateResponseError   `json:"errors"`
	Messages []ZoneCustomHostnameUpdateResponseMessage `json:"messages"`
	Result   ZoneCustomHostnameUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneCustomHostnameUpdateResponseSuccess `json:"success"`
	JSON    zoneCustomHostnameUpdateResponseJSON    `json:"-"`
}

// zoneCustomHostnameUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneCustomHostnameUpdateResponse]
type zoneCustomHostnameUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameUpdateResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneCustomHostnameUpdateResponseErrorJSON `json:"-"`
}

// zoneCustomHostnameUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneCustomHostnameUpdateResponseError]
type zoneCustomHostnameUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameUpdateResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneCustomHostnameUpdateResponseMessageJSON `json:"-"`
}

// zoneCustomHostnameUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneCustomHostnameUpdateResponseMessage]
type zoneCustomHostnameUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameUpdateResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// This is the time the hostname was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// These are per-hostname (customer) settings.
	CustomMetadata ZoneCustomHostnameUpdateResponseResultCustomMetadata `json:"custom_metadata"`
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
	OwnershipVerification ZoneCustomHostnameUpdateResponseResultOwnershipVerification `json:"ownership_verification"`
	// This presents the token to be served by the given http url to activate a
	// hostname.
	OwnershipVerificationHTTP ZoneCustomHostnameUpdateResponseResultOwnershipVerificationHTTP `json:"ownership_verification_http"`
	// SSL properties for the custom hostname.
	Ssl ZoneCustomHostnameUpdateResponseResultSsl `json:"ssl"`
	// Status of the hostname's activation.
	Status ZoneCustomHostnameUpdateResponseResultStatus `json:"status"`
	// These are errors that were encountered while trying to activate a hostname.
	VerificationErrors []interface{}                              `json:"verification_errors"`
	JSON               zoneCustomHostnameUpdateResponseResultJSON `json:"-"`
}

// zoneCustomHostnameUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneCustomHostnameUpdateResponseResult]
type zoneCustomHostnameUpdateResponseResultJSON struct {
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

func (r *ZoneCustomHostnameUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// These are per-hostname (customer) settings.
type ZoneCustomHostnameUpdateResponseResultCustomMetadata struct {
	// Unique metadata for this hostname.
	Key  string                                                   `json:"key"`
	JSON zoneCustomHostnameUpdateResponseResultCustomMetadataJSON `json:"-"`
}

// zoneCustomHostnameUpdateResponseResultCustomMetadataJSON contains the JSON
// metadata for the struct [ZoneCustomHostnameUpdateResponseResultCustomMetadata]
type zoneCustomHostnameUpdateResponseResultCustomMetadataJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameUpdateResponseResultCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This is a record which can be placed to activate a hostname.
type ZoneCustomHostnameUpdateResponseResultOwnershipVerification struct {
	// DNS Name for record.
	Name string `json:"name"`
	// DNS Record type.
	Type ZoneCustomHostnameUpdateResponseResultOwnershipVerificationType `json:"type"`
	// Content for the record.
	Value string                                                          `json:"value"`
	JSON  zoneCustomHostnameUpdateResponseResultOwnershipVerificationJSON `json:"-"`
}

// zoneCustomHostnameUpdateResponseResultOwnershipVerificationJSON contains the
// JSON metadata for the struct
// [ZoneCustomHostnameUpdateResponseResultOwnershipVerification]
type zoneCustomHostnameUpdateResponseResultOwnershipVerificationJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameUpdateResponseResultOwnershipVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS Record type.
type ZoneCustomHostnameUpdateResponseResultOwnershipVerificationType string

const (
	ZoneCustomHostnameUpdateResponseResultOwnershipVerificationTypeTxt ZoneCustomHostnameUpdateResponseResultOwnershipVerificationType = "txt"
)

// This presents the token to be served by the given http url to activate a
// hostname.
type ZoneCustomHostnameUpdateResponseResultOwnershipVerificationHTTP struct {
	// Token to be served.
	HTTPBody string `json:"http_body"`
	// The HTTP URL that will be checked during custom hostname verification and where
	// the customer should host the token.
	HTTPURL string                                                              `json:"http_url"`
	JSON    zoneCustomHostnameUpdateResponseResultOwnershipVerificationHTTPJSON `json:"-"`
}

// zoneCustomHostnameUpdateResponseResultOwnershipVerificationHTTPJSON contains the
// JSON metadata for the struct
// [ZoneCustomHostnameUpdateResponseResultOwnershipVerificationHTTP]
type zoneCustomHostnameUpdateResponseResultOwnershipVerificationHTTPJSON struct {
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameUpdateResponseResultOwnershipVerificationHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// SSL properties for the custom hostname.
type ZoneCustomHostnameUpdateResponseResultSsl struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod ZoneCustomHostnameUpdateResponseResultSslBundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority ZoneCustomHostnameUpdateResponseResultSslCertificateAuthority `json:"certificate_authority"`
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
	Method ZoneCustomHostnameUpdateResponseResultSslMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// SSL specific settings.
	Settings ZoneCustomHostnameUpdateResponseResultSslSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status ZoneCustomHostnameUpdateResponseResultSslStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type ZoneCustomHostnameUpdateResponseResultSslType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []ZoneCustomHostnameUpdateResponseResultSslValidationError  `json:"validation_errors"`
	ValidationRecords []ZoneCustomHostnameUpdateResponseResultSslValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                                          `json:"wildcard"`
	JSON     zoneCustomHostnameUpdateResponseResultSslJSON `json:"-"`
}

// zoneCustomHostnameUpdateResponseResultSslJSON contains the JSON metadata for the
// struct [ZoneCustomHostnameUpdateResponseResultSsl]
type zoneCustomHostnameUpdateResponseResultSslJSON struct {
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

func (r *ZoneCustomHostnameUpdateResponseResultSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type ZoneCustomHostnameUpdateResponseResultSslBundleMethod string

const (
	ZoneCustomHostnameUpdateResponseResultSslBundleMethodUbiquitous ZoneCustomHostnameUpdateResponseResultSslBundleMethod = "ubiquitous"
	ZoneCustomHostnameUpdateResponseResultSslBundleMethodOptimal    ZoneCustomHostnameUpdateResponseResultSslBundleMethod = "optimal"
	ZoneCustomHostnameUpdateResponseResultSslBundleMethodForce      ZoneCustomHostnameUpdateResponseResultSslBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type ZoneCustomHostnameUpdateResponseResultSslCertificateAuthority string

const (
	ZoneCustomHostnameUpdateResponseResultSslCertificateAuthorityDigicert    ZoneCustomHostnameUpdateResponseResultSslCertificateAuthority = "digicert"
	ZoneCustomHostnameUpdateResponseResultSslCertificateAuthorityGoogle      ZoneCustomHostnameUpdateResponseResultSslCertificateAuthority = "google"
	ZoneCustomHostnameUpdateResponseResultSslCertificateAuthorityLetsEncrypt ZoneCustomHostnameUpdateResponseResultSslCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type ZoneCustomHostnameUpdateResponseResultSslMethod string

const (
	ZoneCustomHostnameUpdateResponseResultSslMethodHTTP  ZoneCustomHostnameUpdateResponseResultSslMethod = "http"
	ZoneCustomHostnameUpdateResponseResultSslMethodTxt   ZoneCustomHostnameUpdateResponseResultSslMethod = "txt"
	ZoneCustomHostnameUpdateResponseResultSslMethodEmail ZoneCustomHostnameUpdateResponseResultSslMethod = "email"
)

// SSL specific settings.
type ZoneCustomHostnameUpdateResponseResultSslSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints ZoneCustomHostnameUpdateResponseResultSslSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	Http2 ZoneCustomHostnameUpdateResponseResultSslSettingsHttp2 `json:"http2"`
	// The minimum TLS version supported.
	MinTlsVersion ZoneCustomHostnameUpdateResponseResultSslSettingsMinTlsVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	Tls1_3 ZoneCustomHostnameUpdateResponseResultSslSettingsTls1_3 `json:"tls_1_3"`
	JSON   zoneCustomHostnameUpdateResponseResultSslSettingsJSON   `json:"-"`
}

// zoneCustomHostnameUpdateResponseResultSslSettingsJSON contains the JSON metadata
// for the struct [ZoneCustomHostnameUpdateResponseResultSslSettings]
type zoneCustomHostnameUpdateResponseResultSslSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	Http2         apijson.Field
	MinTlsVersion apijson.Field
	Tls1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneCustomHostnameUpdateResponseResultSslSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not Early Hints is enabled.
type ZoneCustomHostnameUpdateResponseResultSslSettingsEarlyHints string

const (
	ZoneCustomHostnameUpdateResponseResultSslSettingsEarlyHintsOn  ZoneCustomHostnameUpdateResponseResultSslSettingsEarlyHints = "on"
	ZoneCustomHostnameUpdateResponseResultSslSettingsEarlyHintsOff ZoneCustomHostnameUpdateResponseResultSslSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type ZoneCustomHostnameUpdateResponseResultSslSettingsHttp2 string

const (
	ZoneCustomHostnameUpdateResponseResultSslSettingsHttp2On  ZoneCustomHostnameUpdateResponseResultSslSettingsHttp2 = "on"
	ZoneCustomHostnameUpdateResponseResultSslSettingsHttp2Off ZoneCustomHostnameUpdateResponseResultSslSettingsHttp2 = "off"
)

// The minimum TLS version supported.
type ZoneCustomHostnameUpdateResponseResultSslSettingsMinTlsVersion string

const (
	ZoneCustomHostnameUpdateResponseResultSslSettingsMinTlsVersion1_0 ZoneCustomHostnameUpdateResponseResultSslSettingsMinTlsVersion = "1.0"
	ZoneCustomHostnameUpdateResponseResultSslSettingsMinTlsVersion1_1 ZoneCustomHostnameUpdateResponseResultSslSettingsMinTlsVersion = "1.1"
	ZoneCustomHostnameUpdateResponseResultSslSettingsMinTlsVersion1_2 ZoneCustomHostnameUpdateResponseResultSslSettingsMinTlsVersion = "1.2"
	ZoneCustomHostnameUpdateResponseResultSslSettingsMinTlsVersion1_3 ZoneCustomHostnameUpdateResponseResultSslSettingsMinTlsVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type ZoneCustomHostnameUpdateResponseResultSslSettingsTls1_3 string

const (
	ZoneCustomHostnameUpdateResponseResultSslSettingsTls1_3On  ZoneCustomHostnameUpdateResponseResultSslSettingsTls1_3 = "on"
	ZoneCustomHostnameUpdateResponseResultSslSettingsTls1_3Off ZoneCustomHostnameUpdateResponseResultSslSettingsTls1_3 = "off"
)

// Status of the hostname's SSL certificates.
type ZoneCustomHostnameUpdateResponseResultSslStatus string

const (
	ZoneCustomHostnameUpdateResponseResultSslStatusInitializing         ZoneCustomHostnameUpdateResponseResultSslStatus = "initializing"
	ZoneCustomHostnameUpdateResponseResultSslStatusPendingValidation    ZoneCustomHostnameUpdateResponseResultSslStatus = "pending_validation"
	ZoneCustomHostnameUpdateResponseResultSslStatusDeleted              ZoneCustomHostnameUpdateResponseResultSslStatus = "deleted"
	ZoneCustomHostnameUpdateResponseResultSslStatusPendingIssuance      ZoneCustomHostnameUpdateResponseResultSslStatus = "pending_issuance"
	ZoneCustomHostnameUpdateResponseResultSslStatusPendingDeployment    ZoneCustomHostnameUpdateResponseResultSslStatus = "pending_deployment"
	ZoneCustomHostnameUpdateResponseResultSslStatusPendingDeletion      ZoneCustomHostnameUpdateResponseResultSslStatus = "pending_deletion"
	ZoneCustomHostnameUpdateResponseResultSslStatusPendingExpiration    ZoneCustomHostnameUpdateResponseResultSslStatus = "pending_expiration"
	ZoneCustomHostnameUpdateResponseResultSslStatusExpired              ZoneCustomHostnameUpdateResponseResultSslStatus = "expired"
	ZoneCustomHostnameUpdateResponseResultSslStatusActive               ZoneCustomHostnameUpdateResponseResultSslStatus = "active"
	ZoneCustomHostnameUpdateResponseResultSslStatusInitializingTimedOut ZoneCustomHostnameUpdateResponseResultSslStatus = "initializing_timed_out"
	ZoneCustomHostnameUpdateResponseResultSslStatusValidationTimedOut   ZoneCustomHostnameUpdateResponseResultSslStatus = "validation_timed_out"
	ZoneCustomHostnameUpdateResponseResultSslStatusIssuanceTimedOut     ZoneCustomHostnameUpdateResponseResultSslStatus = "issuance_timed_out"
	ZoneCustomHostnameUpdateResponseResultSslStatusDeploymentTimedOut   ZoneCustomHostnameUpdateResponseResultSslStatus = "deployment_timed_out"
	ZoneCustomHostnameUpdateResponseResultSslStatusDeletionTimedOut     ZoneCustomHostnameUpdateResponseResultSslStatus = "deletion_timed_out"
	ZoneCustomHostnameUpdateResponseResultSslStatusPendingCleanup       ZoneCustomHostnameUpdateResponseResultSslStatus = "pending_cleanup"
	ZoneCustomHostnameUpdateResponseResultSslStatusStagingDeployment    ZoneCustomHostnameUpdateResponseResultSslStatus = "staging_deployment"
	ZoneCustomHostnameUpdateResponseResultSslStatusStagingActive        ZoneCustomHostnameUpdateResponseResultSslStatus = "staging_active"
	ZoneCustomHostnameUpdateResponseResultSslStatusDeactivating         ZoneCustomHostnameUpdateResponseResultSslStatus = "deactivating"
	ZoneCustomHostnameUpdateResponseResultSslStatusInactive             ZoneCustomHostnameUpdateResponseResultSslStatus = "inactive"
	ZoneCustomHostnameUpdateResponseResultSslStatusBackupIssued         ZoneCustomHostnameUpdateResponseResultSslStatus = "backup_issued"
	ZoneCustomHostnameUpdateResponseResultSslStatusHoldingDeployment    ZoneCustomHostnameUpdateResponseResultSslStatus = "holding_deployment"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type ZoneCustomHostnameUpdateResponseResultSslType string

const (
	ZoneCustomHostnameUpdateResponseResultSslTypeDv ZoneCustomHostnameUpdateResponseResultSslType = "dv"
)

type ZoneCustomHostnameUpdateResponseResultSslValidationError struct {
	// A domain validation error.
	Message string                                                       `json:"message"`
	JSON    zoneCustomHostnameUpdateResponseResultSslValidationErrorJSON `json:"-"`
}

// zoneCustomHostnameUpdateResponseResultSslValidationErrorJSON contains the JSON
// metadata for the struct
// [ZoneCustomHostnameUpdateResponseResultSslValidationError]
type zoneCustomHostnameUpdateResponseResultSslValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameUpdateResponseResultSslValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate's required validation record.
type ZoneCustomHostnameUpdateResponseResultSslValidationRecord struct {
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
	JSON     zoneCustomHostnameUpdateResponseResultSslValidationRecordJSON `json:"-"`
}

// zoneCustomHostnameUpdateResponseResultSslValidationRecordJSON contains the JSON
// metadata for the struct
// [ZoneCustomHostnameUpdateResponseResultSslValidationRecord]
type zoneCustomHostnameUpdateResponseResultSslValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TxtName     apijson.Field
	TxtValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameUpdateResponseResultSslValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type ZoneCustomHostnameUpdateResponseResultStatus string

const (
	ZoneCustomHostnameUpdateResponseResultStatusActive             ZoneCustomHostnameUpdateResponseResultStatus = "active"
	ZoneCustomHostnameUpdateResponseResultStatusPending            ZoneCustomHostnameUpdateResponseResultStatus = "pending"
	ZoneCustomHostnameUpdateResponseResultStatusActiveRedeploying  ZoneCustomHostnameUpdateResponseResultStatus = "active_redeploying"
	ZoneCustomHostnameUpdateResponseResultStatusMoved              ZoneCustomHostnameUpdateResponseResultStatus = "moved"
	ZoneCustomHostnameUpdateResponseResultStatusPendingDeletion    ZoneCustomHostnameUpdateResponseResultStatus = "pending_deletion"
	ZoneCustomHostnameUpdateResponseResultStatusDeleted            ZoneCustomHostnameUpdateResponseResultStatus = "deleted"
	ZoneCustomHostnameUpdateResponseResultStatusPendingBlocked     ZoneCustomHostnameUpdateResponseResultStatus = "pending_blocked"
	ZoneCustomHostnameUpdateResponseResultStatusPendingMigration   ZoneCustomHostnameUpdateResponseResultStatus = "pending_migration"
	ZoneCustomHostnameUpdateResponseResultStatusPendingProvisioned ZoneCustomHostnameUpdateResponseResultStatus = "pending_provisioned"
	ZoneCustomHostnameUpdateResponseResultStatusTestPending        ZoneCustomHostnameUpdateResponseResultStatus = "test_pending"
	ZoneCustomHostnameUpdateResponseResultStatusTestActive         ZoneCustomHostnameUpdateResponseResultStatus = "test_active"
	ZoneCustomHostnameUpdateResponseResultStatusTestActiveApex     ZoneCustomHostnameUpdateResponseResultStatus = "test_active_apex"
	ZoneCustomHostnameUpdateResponseResultStatusTestBlocked        ZoneCustomHostnameUpdateResponseResultStatus = "test_blocked"
	ZoneCustomHostnameUpdateResponseResultStatusTestFailed         ZoneCustomHostnameUpdateResponseResultStatus = "test_failed"
	ZoneCustomHostnameUpdateResponseResultStatusProvisioned        ZoneCustomHostnameUpdateResponseResultStatus = "provisioned"
	ZoneCustomHostnameUpdateResponseResultStatusBlocked            ZoneCustomHostnameUpdateResponseResultStatus = "blocked"
)

// Whether the API call was successful
type ZoneCustomHostnameUpdateResponseSuccess bool

const (
	ZoneCustomHostnameUpdateResponseSuccessTrue ZoneCustomHostnameUpdateResponseSuccess = true
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

type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponse struct {
	Errors   []ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseError   `json:"errors"`
	Messages []ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseMessage `json:"messages"`
	Result   ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSuccess `json:"success"`
	JSON    zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseJSON    `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseJSON contains
// the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponse]
type zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseError struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseErrorJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseError]
type zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseMessage struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseMessageJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseMessage]
type zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// This is the time the hostname was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// These are per-hostname (customer) settings.
	CustomMetadata ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultCustomMetadata `json:"custom_metadata"`
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
	OwnershipVerification ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerification `json:"ownership_verification"`
	// This presents the token to be served by the given http url to activate a
	// hostname.
	OwnershipVerificationHTTP ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerificationHTTP `json:"ownership_verification_http"`
	// SSL properties for the custom hostname.
	Ssl ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSsl `json:"ssl"`
	// Status of the hostname's activation.
	Status ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus `json:"status"`
	// These are errors that were encountered while trying to activate a hostname.
	VerificationErrors []interface{}                                                               `json:"verification_errors"`
	JSON               zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResult]
type zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultJSON struct {
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

func (r *ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// These are per-hostname (customer) settings.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultCustomMetadata struct {
	// Unique metadata for this hostname.
	Key  string                                                                                    `json:"key"`
	JSON zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultCustomMetadataJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultCustomMetadataJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultCustomMetadata]
type zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultCustomMetadataJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This is a record which can be placed to activate a hostname.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerification struct {
	// DNS Name for record.
	Name string `json:"name"`
	// DNS Record type.
	Type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerificationType `json:"type"`
	// Content for the record.
	Value string                                                                                           `json:"value"`
	JSON  zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerificationJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerificationJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerification]
type zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerificationJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS Record type.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerificationType string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerificationTypeTxt ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerificationType = "txt"
)

// This presents the token to be served by the given http url to activate a
// hostname.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerificationHTTP struct {
	// Token to be served.
	HTTPBody string `json:"http_body"`
	// The HTTP URL that will be checked during custom hostname verification and where
	// the customer should host the token.
	HTTPURL string                                                                                               `json:"http_url"`
	JSON    zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerificationHTTPJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerificationHTTPJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerificationHTTP]
type zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerificationHTTPJSON struct {
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultOwnershipVerificationHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// SSL properties for the custom hostname.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSsl struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslBundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslCertificateAuthority `json:"certificate_authority"`
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
	Method ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// SSL specific settings.
	Settings ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslValidationError  `json:"validation_errors"`
	ValidationRecords []ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                                                                           `json:"wildcard"`
	JSON     zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSsl]
type zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslJSON struct {
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

func (r *ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslBundleMethod string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslBundleMethodUbiquitous ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslBundleMethod = "ubiquitous"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslBundleMethodOptimal    ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslBundleMethod = "optimal"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslBundleMethodForce      ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslCertificateAuthority string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslCertificateAuthorityDigicert    ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslCertificateAuthority = "digicert"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslCertificateAuthorityGoogle      ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslCertificateAuthority = "google"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslCertificateAuthorityLetsEncrypt ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslMethod string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslMethodHTTP  ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslMethod = "http"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslMethodTxt   ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslMethod = "txt"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslMethodEmail ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslMethod = "email"
)

// SSL specific settings.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	Http2 ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsHttp2 `json:"http2"`
	// The minimum TLS version supported.
	MinTlsVersion ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsMinTlsVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	Tls1_3 ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsTls1_3 `json:"tls_1_3"`
	JSON   zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsJSON   `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettings]
type zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	Http2         apijson.Field
	MinTlsVersion apijson.Field
	Tls1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not Early Hints is enabled.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsEarlyHints string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsEarlyHintsOn  ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsEarlyHints = "on"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsEarlyHintsOff ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsHttp2 string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsHttp2On  ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsHttp2 = "on"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsHttp2Off ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsHttp2 = "off"
)

// The minimum TLS version supported.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsMinTlsVersion string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsMinTlsVersion1_0 ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsMinTlsVersion = "1.0"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsMinTlsVersion1_1 ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsMinTlsVersion = "1.1"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsMinTlsVersion1_2 ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsMinTlsVersion = "1.2"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsMinTlsVersion1_3 ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsMinTlsVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsTls1_3 string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsTls1_3On  ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsTls1_3 = "on"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsTls1_3Off ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslSettingsTls1_3 = "off"
)

// Status of the hostname's SSL certificates.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusInitializing         ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "initializing"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusPendingValidation    ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "pending_validation"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusDeleted              ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "deleted"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusPendingIssuance      ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "pending_issuance"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusPendingDeployment    ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "pending_deployment"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusPendingDeletion      ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "pending_deletion"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusPendingExpiration    ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "pending_expiration"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusExpired              ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "expired"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusActive               ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "active"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusInitializingTimedOut ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "initializing_timed_out"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusValidationTimedOut   ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "validation_timed_out"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusIssuanceTimedOut     ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "issuance_timed_out"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusDeploymentTimedOut   ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "deployment_timed_out"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusDeletionTimedOut     ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "deletion_timed_out"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusPendingCleanup       ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "pending_cleanup"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusStagingDeployment    ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "staging_deployment"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusStagingActive        ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "staging_active"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusDeactivating         ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "deactivating"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusInactive             ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "inactive"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusBackupIssued         ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "backup_issued"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatusHoldingDeployment    ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslStatus = "holding_deployment"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslType string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslTypeDv ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslType = "dv"
)

type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslValidationError struct {
	// A domain validation error.
	Message string                                                                                        `json:"message"`
	JSON    zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslValidationErrorJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslValidationErrorJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslValidationError]
type zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate's required validation record.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslValidationRecord struct {
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
	TxtValue string                                                                                         `json:"txt_value"`
	JSON     zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslValidationRecordJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslValidationRecordJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslValidationRecord]
type zoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TxtName     apijson.Field
	TxtValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultSslValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus string

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusActive             ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "active"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusPending            ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "pending"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusActiveRedeploying  ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "active_redeploying"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusMoved              ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "moved"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusPendingDeletion    ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "pending_deletion"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusDeleted            ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "deleted"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusPendingBlocked     ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "pending_blocked"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusPendingMigration   ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "pending_migration"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusPendingProvisioned ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "pending_provisioned"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusTestPending        ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "test_pending"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusTestActive         ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "test_active"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusTestActiveApex     ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "test_active_apex"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusTestBlocked        ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "test_blocked"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusTestFailed         ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "test_failed"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusProvisioned        ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "provisioned"
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatusBlocked            ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseResultStatus = "blocked"
)

// Whether the API call was successful
type ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSuccess bool

const (
	ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSuccessTrue ZoneCustomHostnameCustomHostnameForAZoneNewCustomHostnameResponseSuccess = true
)

type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponse struct {
	// Identifier
	ID string `json:"id"`
	// This is the time the hostname was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// These are per-hostname (customer) settings.
	CustomMetadata ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseCustomMetadata `json:"custom_metadata"`
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
	OwnershipVerification ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerification `json:"ownership_verification"`
	// This presents the token to be served by the given http url to activate a
	// hostname.
	OwnershipVerificationHTTP ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerificationHTTP `json:"ownership_verification_http"`
	// SSL properties for the custom hostname.
	Ssl ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSsl `json:"ssl"`
	// Status of the hostname's activation.
	Status ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus `json:"status"`
	// These are errors that were encountered while trying to activate a hostname.
	VerificationErrors []interface{}                                                           `json:"verification_errors"`
	JSON               zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseJSON contains
// the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponse]
type zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseJSON struct {
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

func (r *ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// These are per-hostname (customer) settings.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseCustomMetadata struct {
	// Unique metadata for this hostname.
	Key  string                                                                                `json:"key"`
	JSON zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseCustomMetadataJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseCustomMetadataJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseCustomMetadata]
type zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseCustomMetadataJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseCustomMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This is a record which can be placed to activate a hostname.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerification struct {
	// DNS Name for record.
	Name string `json:"name"`
	// DNS Record type.
	Type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerificationType `json:"type"`
	// Content for the record.
	Value string                                                                                       `json:"value"`
	JSON  zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerificationJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerificationJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerification]
type zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerificationJSON struct {
	Name        apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerification) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// DNS Record type.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerificationType string

const (
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerificationTypeTxt ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerificationType = "txt"
)

// This presents the token to be served by the given http url to activate a
// hostname.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerificationHTTP struct {
	// Token to be served.
	HTTPBody string `json:"http_body"`
	// The HTTP URL that will be checked during custom hostname verification and where
	// the customer should host the token.
	HTTPURL string                                                                                           `json:"http_url"`
	JSON    zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerificationHTTPJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerificationHTTPJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerificationHTTP]
type zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerificationHTTPJSON struct {
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseOwnershipVerificationHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// SSL properties for the custom hostname.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSsl struct {
	// Custom hostname SSL identifier tag.
	ID string `json:"id"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslBundleMethod `json:"bundle_method"`
	// The Certificate Authority that will issue the certificate
	CertificateAuthority ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslCertificateAuthority `json:"certificate_authority"`
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
	Method ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslMethod `json:"method"`
	// The serial number on a custom uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// SSL specific settings.
	Settings ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettings `json:"settings"`
	// The signature on a custom uploaded certificate.
	Signature string `json:"signature"`
	// Status of the hostname's SSL certificates.
	Status ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus `json:"status"`
	// Level of validation to be used for this hostname. Domain validation (dv) must be
	// used.
	Type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslType `json:"type"`
	// The time the custom certificate was uploaded.
	UploadedOn time.Time `json:"uploaded_on" format:"date-time"`
	// Domain validation errors that have been received by the certificate authority
	// (CA).
	ValidationErrors  []ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslValidationError  `json:"validation_errors"`
	ValidationRecords []ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslValidationRecord `json:"validation_records"`
	// Indicates whether the certificate covers a wildcard.
	Wildcard bool                                                                       `json:"wildcard"`
	JSON     zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSsl]
type zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslJSON struct {
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

func (r *ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslBundleMethod string

const (
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslBundleMethodUbiquitous ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslBundleMethod = "ubiquitous"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslBundleMethodOptimal    ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslBundleMethod = "optimal"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslBundleMethodForce      ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslBundleMethod = "force"
)

// The Certificate Authority that will issue the certificate
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslCertificateAuthority string

const (
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslCertificateAuthorityDigicert    ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslCertificateAuthority = "digicert"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslCertificateAuthorityGoogle      ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslCertificateAuthority = "google"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslCertificateAuthorityLetsEncrypt ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslCertificateAuthority = "lets_encrypt"
)

// Domain control validation (DCV) method used for this hostname.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslMethod string

const (
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslMethodHTTP  ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslMethod = "http"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslMethodTxt   ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslMethod = "txt"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslMethodEmail ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslMethod = "email"
)

// SSL specific settings.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettings struct {
	// An allowlist of ciphers for TLS termination. These ciphers must be in the
	// BoringSSL format.
	Ciphers []string `json:"ciphers"`
	// Whether or not Early Hints is enabled.
	EarlyHints ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsEarlyHints `json:"early_hints"`
	// Whether or not HTTP2 is enabled.
	Http2 ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsHttp2 `json:"http2"`
	// The minimum TLS version supported.
	MinTlsVersion ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsMinTlsVersion `json:"min_tls_version"`
	// Whether or not TLS 1.3 is enabled.
	Tls1_3 ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsTls1_3 `json:"tls_1_3"`
	JSON   zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsJSON   `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettings]
type zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsJSON struct {
	Ciphers       apijson.Field
	EarlyHints    apijson.Field
	Http2         apijson.Field
	MinTlsVersion apijson.Field
	Tls1_3        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether or not Early Hints is enabled.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsEarlyHints string

const (
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsEarlyHintsOn  ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsEarlyHints = "on"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsEarlyHintsOff ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsEarlyHints = "off"
)

// Whether or not HTTP2 is enabled.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsHttp2 string

const (
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsHttp2On  ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsHttp2 = "on"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsHttp2Off ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsHttp2 = "off"
)

// The minimum TLS version supported.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsMinTlsVersion string

const (
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsMinTlsVersion1_0 ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsMinTlsVersion = "1.0"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsMinTlsVersion1_1 ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsMinTlsVersion = "1.1"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsMinTlsVersion1_2 ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsMinTlsVersion = "1.2"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsMinTlsVersion1_3 ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsMinTlsVersion = "1.3"
)

// Whether or not TLS 1.3 is enabled.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsTls1_3 string

const (
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsTls1_3On  ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsTls1_3 = "on"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsTls1_3Off ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslSettingsTls1_3 = "off"
)

// Status of the hostname's SSL certificates.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus string

const (
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusInitializing         ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "initializing"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusPendingValidation    ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "pending_validation"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusDeleted              ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "deleted"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusPendingIssuance      ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "pending_issuance"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusPendingDeployment    ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "pending_deployment"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusPendingDeletion      ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "pending_deletion"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusPendingExpiration    ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "pending_expiration"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusExpired              ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "expired"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusActive               ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "active"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusInitializingTimedOut ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "initializing_timed_out"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusValidationTimedOut   ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "validation_timed_out"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusIssuanceTimedOut     ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "issuance_timed_out"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusDeploymentTimedOut   ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "deployment_timed_out"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusDeletionTimedOut     ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "deletion_timed_out"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusPendingCleanup       ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "pending_cleanup"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusStagingDeployment    ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "staging_deployment"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusStagingActive        ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "staging_active"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusDeactivating         ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "deactivating"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusInactive             ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "inactive"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusBackupIssued         ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "backup_issued"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatusHoldingDeployment    ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslStatus = "holding_deployment"
)

// Level of validation to be used for this hostname. Domain validation (dv) must be
// used.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslType string

const (
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslTypeDv ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslType = "dv"
)

type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslValidationError struct {
	// A domain validation error.
	Message string                                                                                    `json:"message"`
	JSON    zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslValidationErrorJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslValidationErrorJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslValidationError]
type zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslValidationErrorJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslValidationError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Certificate's required validation record.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslValidationRecord struct {
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
	TxtValue string                                                                                     `json:"txt_value"`
	JSON     zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslValidationRecordJSON `json:"-"`
}

// zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslValidationRecordJSON
// contains the JSON metadata for the struct
// [ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslValidationRecord]
type zoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslValidationRecordJSON struct {
	Emails      apijson.Field
	HTTPBody    apijson.Field
	HTTPURL     apijson.Field
	TxtName     apijson.Field
	TxtValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseSslValidationRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the hostname's activation.
type ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus string

const (
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusActive             ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "active"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusPending            ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "pending"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusActiveRedeploying  ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "active_redeploying"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusMoved              ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "moved"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusPendingDeletion    ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "pending_deletion"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusDeleted            ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "deleted"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusPendingBlocked     ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "pending_blocked"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusPendingMigration   ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "pending_migration"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusPendingProvisioned ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "pending_provisioned"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusTestPending        ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "test_pending"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusTestActive         ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "test_active"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusTestActiveApex     ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "test_active_apex"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusTestBlocked        ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "test_blocked"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusTestFailed         ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "test_failed"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusProvisioned        ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "provisioned"
	ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatusBlocked            ZoneCustomHostnameCustomHostnameForAZoneListCustomHostnamesResponseStatus = "blocked"
)

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

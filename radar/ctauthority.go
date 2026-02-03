// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// CTAuthorityService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCTAuthorityService] method instead.
type CTAuthorityService struct {
	Options []option.RequestOption
}

// NewCTAuthorityService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCTAuthorityService(opts ...option.RequestOption) (r *CTAuthorityService) {
	r = &CTAuthorityService{}
	r.Options = opts
	return
}

// Retrieves a list of certificate authorities.
func (r *CTAuthorityService) List(ctx context.Context, query CTAuthorityListParams, opts ...option.RequestOption) (res *CTAuthorityListResponse, err error) {
	var env CTAuthorityListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/ct/authorities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the requested CA information.
func (r *CTAuthorityService) Get(ctx context.Context, caSlug string, query CTAuthorityGetParams, opts ...option.RequestOption) (res *CTAuthorityGetResponse, err error) {
	var env CTAuthorityGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if caSlug == "" {
		err = errors.New("missing required ca_slug parameter")
		return
	}
	path := fmt.Sprintf("radar/ct/authorities/%s", caSlug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CTAuthorityListResponse struct {
	CertificateAuthorities []CTAuthorityListResponseCertificateAuthority `json:"certificateAuthorities,required"`
	JSON                   ctAuthorityListResponseJSON                   `json:"-"`
}

// ctAuthorityListResponseJSON contains the JSON metadata for the struct
// [CTAuthorityListResponse]
type ctAuthorityListResponseJSON struct {
	CertificateAuthorities apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CTAuthorityListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAuthorityListResponseJSON) RawJSON() string {
	return r.raw
}

type CTAuthorityListResponseCertificateAuthority struct {
	// Specifies the type of certificate in the trust chain.
	CertificateRecordType CTAuthorityListResponseCertificateAuthoritiesCertificateRecordType `json:"certificateRecordType,required"`
	// The two-letter ISO country code where the CA organization is based.
	Country string `json:"country,required"`
	// The full country name corresponding to the country code.
	CountryName string `json:"countryName,required"`
	// The full name of the certificate authority (CA).
	Name string `json:"name,required"`
	// The organization that owns and operates the CA.
	Owner string `json:"owner,required"`
	// The name of the parent/root certificate authority that issued this intermediate
	// certificate.
	ParentName string `json:"parentName,required"`
	// The SHA-256 fingerprint of the parent certificate.
	ParentSha256Fingerprint string `json:"parentSha256Fingerprint,required"`
	// The current revocation status of a Certificate Authority (CA) certificate.
	RevocationStatus CTAuthorityListResponseCertificateAuthoritiesRevocationStatus `json:"revocationStatus,required"`
	// The SHA-256 fingerprint of the intermediate certificate.
	Sha256Fingerprint string                                          `json:"sha256Fingerprint,required"`
	JSON              ctAuthorityListResponseCertificateAuthorityJSON `json:"-"`
}

// ctAuthorityListResponseCertificateAuthorityJSON contains the JSON metadata for
// the struct [CTAuthorityListResponseCertificateAuthority]
type ctAuthorityListResponseCertificateAuthorityJSON struct {
	CertificateRecordType   apijson.Field
	Country                 apijson.Field
	CountryName             apijson.Field
	Name                    apijson.Field
	Owner                   apijson.Field
	ParentName              apijson.Field
	ParentSha256Fingerprint apijson.Field
	RevocationStatus        apijson.Field
	Sha256Fingerprint       apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CTAuthorityListResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAuthorityListResponseCertificateAuthorityJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of certificate in the trust chain.
type CTAuthorityListResponseCertificateAuthoritiesCertificateRecordType string

const (
	CTAuthorityListResponseCertificateAuthoritiesCertificateRecordTypeRootCertificate         CTAuthorityListResponseCertificateAuthoritiesCertificateRecordType = "ROOT_CERTIFICATE"
	CTAuthorityListResponseCertificateAuthoritiesCertificateRecordTypeIntermediateCertificate CTAuthorityListResponseCertificateAuthoritiesCertificateRecordType = "INTERMEDIATE_CERTIFICATE"
)

func (r CTAuthorityListResponseCertificateAuthoritiesCertificateRecordType) IsKnown() bool {
	switch r {
	case CTAuthorityListResponseCertificateAuthoritiesCertificateRecordTypeRootCertificate, CTAuthorityListResponseCertificateAuthoritiesCertificateRecordTypeIntermediateCertificate:
		return true
	}
	return false
}

// The current revocation status of a Certificate Authority (CA) certificate.
type CTAuthorityListResponseCertificateAuthoritiesRevocationStatus string

const (
	CTAuthorityListResponseCertificateAuthoritiesRevocationStatusNotRevoked        CTAuthorityListResponseCertificateAuthoritiesRevocationStatus = "NOT_REVOKED"
	CTAuthorityListResponseCertificateAuthoritiesRevocationStatusRevoked           CTAuthorityListResponseCertificateAuthoritiesRevocationStatus = "REVOKED"
	CTAuthorityListResponseCertificateAuthoritiesRevocationStatusParentCERTRevoked CTAuthorityListResponseCertificateAuthoritiesRevocationStatus = "PARENT_CERT_REVOKED"
)

func (r CTAuthorityListResponseCertificateAuthoritiesRevocationStatus) IsKnown() bool {
	switch r {
	case CTAuthorityListResponseCertificateAuthoritiesRevocationStatusNotRevoked, CTAuthorityListResponseCertificateAuthoritiesRevocationStatusRevoked, CTAuthorityListResponseCertificateAuthoritiesRevocationStatusParentCERTRevoked:
		return true
	}
	return false
}

type CTAuthorityGetResponse struct {
	CertificateAuthority CTAuthorityGetResponseCertificateAuthority `json:"certificateAuthority,required"`
	JSON                 ctAuthorityGetResponseJSON                 `json:"-"`
}

// ctAuthorityGetResponseJSON contains the JSON metadata for the struct
// [CTAuthorityGetResponse]
type ctAuthorityGetResponseJSON struct {
	CertificateAuthority apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CTAuthorityGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAuthorityGetResponseJSON) RawJSON() string {
	return r.raw
}

type CTAuthorityGetResponseCertificateAuthority struct {
	// The inclusion status of a Certificate Authority (CA) in the trust store.
	AppleStatus CTAuthorityGetResponseCertificateAuthorityAppleStatus `json:"appleStatus,required"`
	// The authorityKeyIdentifier value extracted from the certificate PEM.
	AuthorityKeyIdentifier string `json:"authorityKeyIdentifier,required"`
	// Specifies the type of certificate in the trust chain.
	CertificateRecordType CTAuthorityGetResponseCertificateAuthorityCertificateRecordType `json:"certificateRecordType,required"`
	// The inclusion status of a Certificate Authority (CA) in the trust store.
	ChromeStatus CTAuthorityGetResponseCertificateAuthorityChromeStatus `json:"chromeStatus,required"`
	// The two-letter ISO country code where the CA organization is based.
	Country string `json:"country,required"`
	// The full country name corresponding to the country code.
	CountryName string `json:"countryName,required"`
	// The inclusion status of a Certificate Authority (CA) in the trust store.
	MicrosoftStatus CTAuthorityGetResponseCertificateAuthorityMicrosoftStatus `json:"microsoftStatus,required"`
	// The inclusion status of a Certificate Authority (CA) in the trust store.
	MozillaStatus CTAuthorityGetResponseCertificateAuthorityMozillaStatus `json:"mozillaStatus,required"`
	// The full name of the certificate authority (CA).
	Name string `json:"name,required"`
	// The organization that owns and operates the CA.
	Owner string `json:"owner,required"`
	// The name of the parent/root certificate authority that issued this intermediate
	// certificate.
	ParentName string `json:"parentName,required"`
	// The SHA-256 fingerprint of the parent certificate.
	ParentSha256Fingerprint string `json:"parentSha256Fingerprint,required"`
	// CAs from the same owner.
	Related []CTAuthorityGetResponseCertificateAuthorityRelated `json:"related,required"`
	// The current revocation status of a Certificate Authority (CA) certificate.
	RevocationStatus CTAuthorityGetResponseCertificateAuthorityRevocationStatus `json:"revocationStatus,required"`
	// The SHA-256 fingerprint of the intermediate certificate.
	Sha256Fingerprint string `json:"sha256Fingerprint,required"`
	// The subjectKeyIdentifier value extracted from the certificate PEM.
	SubjectKeyIdentifier string `json:"subjectKeyIdentifier,required"`
	// The start date of the certificate’s validity period (ISO format).
	ValidFrom string `json:"validFrom,required"`
	// The end date of the certificate’s validity period (ISO format).
	ValidTo string                                         `json:"validTo,required"`
	JSON    ctAuthorityGetResponseCertificateAuthorityJSON `json:"-"`
}

// ctAuthorityGetResponseCertificateAuthorityJSON contains the JSON metadata for
// the struct [CTAuthorityGetResponseCertificateAuthority]
type ctAuthorityGetResponseCertificateAuthorityJSON struct {
	AppleStatus             apijson.Field
	AuthorityKeyIdentifier  apijson.Field
	CertificateRecordType   apijson.Field
	ChromeStatus            apijson.Field
	Country                 apijson.Field
	CountryName             apijson.Field
	MicrosoftStatus         apijson.Field
	MozillaStatus           apijson.Field
	Name                    apijson.Field
	Owner                   apijson.Field
	ParentName              apijson.Field
	ParentSha256Fingerprint apijson.Field
	Related                 apijson.Field
	RevocationStatus        apijson.Field
	Sha256Fingerprint       apijson.Field
	SubjectKeyIdentifier    apijson.Field
	ValidFrom               apijson.Field
	ValidTo                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *CTAuthorityGetResponseCertificateAuthority) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAuthorityGetResponseCertificateAuthorityJSON) RawJSON() string {
	return r.raw
}

// The inclusion status of a Certificate Authority (CA) in the trust store.
type CTAuthorityGetResponseCertificateAuthorityAppleStatus string

const (
	CTAuthorityGetResponseCertificateAuthorityAppleStatusIncluded       CTAuthorityGetResponseCertificateAuthorityAppleStatus = "INCLUDED"
	CTAuthorityGetResponseCertificateAuthorityAppleStatusNotYetIncluded CTAuthorityGetResponseCertificateAuthorityAppleStatus = "NOT_YET_INCLUDED"
	CTAuthorityGetResponseCertificateAuthorityAppleStatusNotIncluded    CTAuthorityGetResponseCertificateAuthorityAppleStatus = "NOT_INCLUDED"
	CTAuthorityGetResponseCertificateAuthorityAppleStatusNotBefore      CTAuthorityGetResponseCertificateAuthorityAppleStatus = "NOT_BEFORE"
	CTAuthorityGetResponseCertificateAuthorityAppleStatusRemoved        CTAuthorityGetResponseCertificateAuthorityAppleStatus = "REMOVED"
	CTAuthorityGetResponseCertificateAuthorityAppleStatusDisabled       CTAuthorityGetResponseCertificateAuthorityAppleStatus = "DISABLED"
	CTAuthorityGetResponseCertificateAuthorityAppleStatusBlocked        CTAuthorityGetResponseCertificateAuthorityAppleStatus = "BLOCKED"
)

func (r CTAuthorityGetResponseCertificateAuthorityAppleStatus) IsKnown() bool {
	switch r {
	case CTAuthorityGetResponseCertificateAuthorityAppleStatusIncluded, CTAuthorityGetResponseCertificateAuthorityAppleStatusNotYetIncluded, CTAuthorityGetResponseCertificateAuthorityAppleStatusNotIncluded, CTAuthorityGetResponseCertificateAuthorityAppleStatusNotBefore, CTAuthorityGetResponseCertificateAuthorityAppleStatusRemoved, CTAuthorityGetResponseCertificateAuthorityAppleStatusDisabled, CTAuthorityGetResponseCertificateAuthorityAppleStatusBlocked:
		return true
	}
	return false
}

// Specifies the type of certificate in the trust chain.
type CTAuthorityGetResponseCertificateAuthorityCertificateRecordType string

const (
	CTAuthorityGetResponseCertificateAuthorityCertificateRecordTypeRootCertificate         CTAuthorityGetResponseCertificateAuthorityCertificateRecordType = "ROOT_CERTIFICATE"
	CTAuthorityGetResponseCertificateAuthorityCertificateRecordTypeIntermediateCertificate CTAuthorityGetResponseCertificateAuthorityCertificateRecordType = "INTERMEDIATE_CERTIFICATE"
)

func (r CTAuthorityGetResponseCertificateAuthorityCertificateRecordType) IsKnown() bool {
	switch r {
	case CTAuthorityGetResponseCertificateAuthorityCertificateRecordTypeRootCertificate, CTAuthorityGetResponseCertificateAuthorityCertificateRecordTypeIntermediateCertificate:
		return true
	}
	return false
}

// The inclusion status of a Certificate Authority (CA) in the trust store.
type CTAuthorityGetResponseCertificateAuthorityChromeStatus string

const (
	CTAuthorityGetResponseCertificateAuthorityChromeStatusIncluded       CTAuthorityGetResponseCertificateAuthorityChromeStatus = "INCLUDED"
	CTAuthorityGetResponseCertificateAuthorityChromeStatusNotYetIncluded CTAuthorityGetResponseCertificateAuthorityChromeStatus = "NOT_YET_INCLUDED"
	CTAuthorityGetResponseCertificateAuthorityChromeStatusNotIncluded    CTAuthorityGetResponseCertificateAuthorityChromeStatus = "NOT_INCLUDED"
	CTAuthorityGetResponseCertificateAuthorityChromeStatusNotBefore      CTAuthorityGetResponseCertificateAuthorityChromeStatus = "NOT_BEFORE"
	CTAuthorityGetResponseCertificateAuthorityChromeStatusRemoved        CTAuthorityGetResponseCertificateAuthorityChromeStatus = "REMOVED"
	CTAuthorityGetResponseCertificateAuthorityChromeStatusDisabled       CTAuthorityGetResponseCertificateAuthorityChromeStatus = "DISABLED"
	CTAuthorityGetResponseCertificateAuthorityChromeStatusBlocked        CTAuthorityGetResponseCertificateAuthorityChromeStatus = "BLOCKED"
)

func (r CTAuthorityGetResponseCertificateAuthorityChromeStatus) IsKnown() bool {
	switch r {
	case CTAuthorityGetResponseCertificateAuthorityChromeStatusIncluded, CTAuthorityGetResponseCertificateAuthorityChromeStatusNotYetIncluded, CTAuthorityGetResponseCertificateAuthorityChromeStatusNotIncluded, CTAuthorityGetResponseCertificateAuthorityChromeStatusNotBefore, CTAuthorityGetResponseCertificateAuthorityChromeStatusRemoved, CTAuthorityGetResponseCertificateAuthorityChromeStatusDisabled, CTAuthorityGetResponseCertificateAuthorityChromeStatusBlocked:
		return true
	}
	return false
}

// The inclusion status of a Certificate Authority (CA) in the trust store.
type CTAuthorityGetResponseCertificateAuthorityMicrosoftStatus string

const (
	CTAuthorityGetResponseCertificateAuthorityMicrosoftStatusIncluded       CTAuthorityGetResponseCertificateAuthorityMicrosoftStatus = "INCLUDED"
	CTAuthorityGetResponseCertificateAuthorityMicrosoftStatusNotYetIncluded CTAuthorityGetResponseCertificateAuthorityMicrosoftStatus = "NOT_YET_INCLUDED"
	CTAuthorityGetResponseCertificateAuthorityMicrosoftStatusNotIncluded    CTAuthorityGetResponseCertificateAuthorityMicrosoftStatus = "NOT_INCLUDED"
	CTAuthorityGetResponseCertificateAuthorityMicrosoftStatusNotBefore      CTAuthorityGetResponseCertificateAuthorityMicrosoftStatus = "NOT_BEFORE"
	CTAuthorityGetResponseCertificateAuthorityMicrosoftStatusRemoved        CTAuthorityGetResponseCertificateAuthorityMicrosoftStatus = "REMOVED"
	CTAuthorityGetResponseCertificateAuthorityMicrosoftStatusDisabled       CTAuthorityGetResponseCertificateAuthorityMicrosoftStatus = "DISABLED"
	CTAuthorityGetResponseCertificateAuthorityMicrosoftStatusBlocked        CTAuthorityGetResponseCertificateAuthorityMicrosoftStatus = "BLOCKED"
)

func (r CTAuthorityGetResponseCertificateAuthorityMicrosoftStatus) IsKnown() bool {
	switch r {
	case CTAuthorityGetResponseCertificateAuthorityMicrosoftStatusIncluded, CTAuthorityGetResponseCertificateAuthorityMicrosoftStatusNotYetIncluded, CTAuthorityGetResponseCertificateAuthorityMicrosoftStatusNotIncluded, CTAuthorityGetResponseCertificateAuthorityMicrosoftStatusNotBefore, CTAuthorityGetResponseCertificateAuthorityMicrosoftStatusRemoved, CTAuthorityGetResponseCertificateAuthorityMicrosoftStatusDisabled, CTAuthorityGetResponseCertificateAuthorityMicrosoftStatusBlocked:
		return true
	}
	return false
}

// The inclusion status of a Certificate Authority (CA) in the trust store.
type CTAuthorityGetResponseCertificateAuthorityMozillaStatus string

const (
	CTAuthorityGetResponseCertificateAuthorityMozillaStatusIncluded       CTAuthorityGetResponseCertificateAuthorityMozillaStatus = "INCLUDED"
	CTAuthorityGetResponseCertificateAuthorityMozillaStatusNotYetIncluded CTAuthorityGetResponseCertificateAuthorityMozillaStatus = "NOT_YET_INCLUDED"
	CTAuthorityGetResponseCertificateAuthorityMozillaStatusNotIncluded    CTAuthorityGetResponseCertificateAuthorityMozillaStatus = "NOT_INCLUDED"
	CTAuthorityGetResponseCertificateAuthorityMozillaStatusNotBefore      CTAuthorityGetResponseCertificateAuthorityMozillaStatus = "NOT_BEFORE"
	CTAuthorityGetResponseCertificateAuthorityMozillaStatusRemoved        CTAuthorityGetResponseCertificateAuthorityMozillaStatus = "REMOVED"
	CTAuthorityGetResponseCertificateAuthorityMozillaStatusDisabled       CTAuthorityGetResponseCertificateAuthorityMozillaStatus = "DISABLED"
	CTAuthorityGetResponseCertificateAuthorityMozillaStatusBlocked        CTAuthorityGetResponseCertificateAuthorityMozillaStatus = "BLOCKED"
)

func (r CTAuthorityGetResponseCertificateAuthorityMozillaStatus) IsKnown() bool {
	switch r {
	case CTAuthorityGetResponseCertificateAuthorityMozillaStatusIncluded, CTAuthorityGetResponseCertificateAuthorityMozillaStatusNotYetIncluded, CTAuthorityGetResponseCertificateAuthorityMozillaStatusNotIncluded, CTAuthorityGetResponseCertificateAuthorityMozillaStatusNotBefore, CTAuthorityGetResponseCertificateAuthorityMozillaStatusRemoved, CTAuthorityGetResponseCertificateAuthorityMozillaStatusDisabled, CTAuthorityGetResponseCertificateAuthorityMozillaStatusBlocked:
		return true
	}
	return false
}

type CTAuthorityGetResponseCertificateAuthorityRelated struct {
	// Specifies the type of certificate in the trust chain.
	CertificateRecordType CTAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordType `json:"certificateRecordType,required"`
	// The full name of the certificate authority (CA).
	Name string `json:"name,required"`
	// The current revocation status of a Certificate Authority (CA) certificate.
	RevocationStatus CTAuthorityGetResponseCertificateAuthorityRelatedRevocationStatus `json:"revocationStatus,required"`
	// The SHA-256 fingerprint of the intermediate certificate.
	Sha256Fingerprint string                                                `json:"sha256Fingerprint,required"`
	JSON              ctAuthorityGetResponseCertificateAuthorityRelatedJSON `json:"-"`
}

// ctAuthorityGetResponseCertificateAuthorityRelatedJSON contains the JSON metadata
// for the struct [CTAuthorityGetResponseCertificateAuthorityRelated]
type ctAuthorityGetResponseCertificateAuthorityRelatedJSON struct {
	CertificateRecordType apijson.Field
	Name                  apijson.Field
	RevocationStatus      apijson.Field
	Sha256Fingerprint     apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CTAuthorityGetResponseCertificateAuthorityRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAuthorityGetResponseCertificateAuthorityRelatedJSON) RawJSON() string {
	return r.raw
}

// Specifies the type of certificate in the trust chain.
type CTAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordType string

const (
	CTAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordTypeRootCertificate         CTAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordType = "ROOT_CERTIFICATE"
	CTAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordTypeIntermediateCertificate CTAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordType = "INTERMEDIATE_CERTIFICATE"
)

func (r CTAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordType) IsKnown() bool {
	switch r {
	case CTAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordTypeRootCertificate, CTAuthorityGetResponseCertificateAuthorityRelatedCertificateRecordTypeIntermediateCertificate:
		return true
	}
	return false
}

// The current revocation status of a Certificate Authority (CA) certificate.
type CTAuthorityGetResponseCertificateAuthorityRelatedRevocationStatus string

const (
	CTAuthorityGetResponseCertificateAuthorityRelatedRevocationStatusNotRevoked        CTAuthorityGetResponseCertificateAuthorityRelatedRevocationStatus = "NOT_REVOKED"
	CTAuthorityGetResponseCertificateAuthorityRelatedRevocationStatusRevoked           CTAuthorityGetResponseCertificateAuthorityRelatedRevocationStatus = "REVOKED"
	CTAuthorityGetResponseCertificateAuthorityRelatedRevocationStatusParentCERTRevoked CTAuthorityGetResponseCertificateAuthorityRelatedRevocationStatus = "PARENT_CERT_REVOKED"
)

func (r CTAuthorityGetResponseCertificateAuthorityRelatedRevocationStatus) IsKnown() bool {
	switch r {
	case CTAuthorityGetResponseCertificateAuthorityRelatedRevocationStatusNotRevoked, CTAuthorityGetResponseCertificateAuthorityRelatedRevocationStatusRevoked, CTAuthorityGetResponseCertificateAuthorityRelatedRevocationStatusParentCERTRevoked:
		return true
	}
	return false
}

// The current revocation status of a Certificate Authority (CA) certificate.
type CTAuthorityGetResponseCertificateAuthorityRevocationStatus string

const (
	CTAuthorityGetResponseCertificateAuthorityRevocationStatusNotRevoked        CTAuthorityGetResponseCertificateAuthorityRevocationStatus = "NOT_REVOKED"
	CTAuthorityGetResponseCertificateAuthorityRevocationStatusRevoked           CTAuthorityGetResponseCertificateAuthorityRevocationStatus = "REVOKED"
	CTAuthorityGetResponseCertificateAuthorityRevocationStatusParentCERTRevoked CTAuthorityGetResponseCertificateAuthorityRevocationStatus = "PARENT_CERT_REVOKED"
)

func (r CTAuthorityGetResponseCertificateAuthorityRevocationStatus) IsKnown() bool {
	switch r {
	case CTAuthorityGetResponseCertificateAuthorityRevocationStatusNotRevoked, CTAuthorityGetResponseCertificateAuthorityRevocationStatusRevoked, CTAuthorityGetResponseCertificateAuthorityRevocationStatusParentCERTRevoked:
		return true
	}
	return false
}

type CTAuthorityListParams struct {
	// Format in which results will be returned.
	Format param.Field[CTAuthorityListParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [CTAuthorityListParams]'s query parameters as `url.Values`.
func (r CTAuthorityListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type CTAuthorityListParamsFormat string

const (
	CTAuthorityListParamsFormatJson CTAuthorityListParamsFormat = "JSON"
	CTAuthorityListParamsFormatCsv  CTAuthorityListParamsFormat = "CSV"
)

func (r CTAuthorityListParamsFormat) IsKnown() bool {
	switch r {
	case CTAuthorityListParamsFormatJson, CTAuthorityListParamsFormatCsv:
		return true
	}
	return false
}

type CTAuthorityListResponseEnvelope struct {
	Result  CTAuthorityListResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    ctAuthorityListResponseEnvelopeJSON `json:"-"`
}

// ctAuthorityListResponseEnvelopeJSON contains the JSON metadata for the struct
// [CTAuthorityListResponseEnvelope]
type ctAuthorityListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTAuthorityListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAuthorityListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CTAuthorityGetParams struct {
	// Format in which results will be returned.
	Format param.Field[CTAuthorityGetParamsFormat] `query:"format"`
}

// URLQuery serializes [CTAuthorityGetParams]'s query parameters as `url.Values`.
func (r CTAuthorityGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type CTAuthorityGetParamsFormat string

const (
	CTAuthorityGetParamsFormatJson CTAuthorityGetParamsFormat = "JSON"
	CTAuthorityGetParamsFormatCsv  CTAuthorityGetParamsFormat = "CSV"
)

func (r CTAuthorityGetParamsFormat) IsKnown() bool {
	switch r {
	case CTAuthorityGetParamsFormatJson, CTAuthorityGetParamsFormatCsv:
		return true
	}
	return false
}

type CTAuthorityGetResponseEnvelope struct {
	Result  CTAuthorityGetResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    ctAuthorityGetResponseEnvelopeJSON `json:"-"`
}

// ctAuthorityGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CTAuthorityGetResponseEnvelope]
type ctAuthorityGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CTAuthorityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctAuthorityGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

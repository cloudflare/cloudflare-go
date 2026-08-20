// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// AccessSAMLCertificateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessSAMLCertificateService] method instead.
type AccessSAMLCertificateService struct {
	Options []option.RequestOption
}

// NewAccessSAMLCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessSAMLCertificateService(opts ...option.RequestOption) (r *AccessSAMLCertificateService) {
	r = &AccessSAMLCertificateService{}
	r.Options = opts
	return
}

// Returns a paginated list of the organization's SAML encryption certificate sets.
// Each certificate set includes the current and (if present) previous
// certificates.
func (r *AccessSAMLCertificateService) List(ctx context.Context, params AccessSAMLCertificateListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessSAMLCertificateListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/saml_certificates", params.AccountID)
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

// Returns a paginated list of the organization's SAML encryption certificate sets.
// Each certificate set includes the current and (if present) previous
// certificates.
func (r *AccessSAMLCertificateService) ListAutoPaging(ctx context.Context, params AccessSAMLCertificateListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessSAMLCertificateListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Retrieves a specific SAML encryption certificate set by its UID, including both
// current and previous certificates if available.
func (r *AccessSAMLCertificateService) Get(ctx context.Context, samlCERTSetID string, query AccessSAMLCertificateGetParams, opts ...option.RequestOption) (res *AccessSAMLCertificateGetResponse, err error) {
	var env AccessSAMLCertificateGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if samlCERTSetID == "" {
		err = errors.New("missing required saml_cert_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/saml_certificates/%s", query.AccountID, samlCERTSetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Downloads the current SAML encryption certificate's public key in PEM format for
// the specified certificate set. This endpoint is useful for providing the
// certificate to Identity Providers for SAML assertion encryption configuration.
func (r *AccessSAMLCertificateService) GetPem(ctx context.Context, samlCERTSetID string, query AccessSAMLCertificateGetPemParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/x-pem-file")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if samlCERTSetID == "" {
		err = errors.New("missing required saml_cert_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/saml_certificates/%s/pem", query.AccountID, samlCERTSetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Rotates the SAML encryption certificates within the specified certificate set.
// This generates a new certificate and moves the current certificate to the
// previous slot. If a previous certificate exists, it will be deactivated and
// removed.
//
// This endpoint ensures zero-downtime rotation by maintaining both current and
// previous certificates during the transition period, allowing IdPs time to update
// their configurations. Automated rotation happens 30 days before a current
// certificate's expiration.
func (r *AccessSAMLCertificateService) Rotate(ctx context.Context, samlCERTSetID string, body AccessSAMLCertificateRotateParams, opts ...option.RequestOption) (res *AccessSAMLCertificateRotateResponse, err error) {
	var env AccessSAMLCertificateRotateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if samlCERTSetID == "" {
		err = errors.New("missing required saml_cert_set_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/access/saml_certificates/%s/rotate", body.AccountID, samlCERTSetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AccessSAMLCertificateListResponse struct {
	// When the certificate set was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Unique identifier for the certificate set
	UID string `json:"uid" api:"required"`
	// When the certificate set was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The current active certificate
	CurrentCertificate AccessSAMLCertificateListResponseCurrentCertificate `json:"current_certificate"`
	// The previous certificate (maintained during rotation period). May be null when
	// no rotation has occurred. Mirrors the structure of `saml_certificate`.
	PreviousCertificate interface{}                           `json:"previous_certificate" api:"nullable"`
	JSON                accessSAMLCertificateListResponseJSON `json:"-"`
}

// accessSAMLCertificateListResponseJSON contains the JSON metadata for the struct
// [AccessSAMLCertificateListResponse]
type accessSAMLCertificateListResponseJSON struct {
	CreatedAt           apijson.Field
	UID                 apijson.Field
	UpdatedAt           apijson.Field
	CurrentCertificate  apijson.Field
	PreviousCertificate apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessSAMLCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateListResponseJSON) RawJSON() string {
	return r.raw
}

// AccessSAMLCertificateListResponseCurrentCertificate is the current active certificate
type AccessSAMLCertificateListResponseCurrentCertificate struct {
	// Indicates whether the certificate can be used for IdP configuration.
	IsCurrent bool `json:"is_current" api:"required"`
	// Certificate expiration date
	NotAfter time.Time `json:"not_after" api:"required" format:"date-time"`
	// The public certificate in PEM format
	PublicCertificate string `json:"public_certificate" api:"required"`
	// Unique identifier for the certificate
	UID  string                                                  `json:"uid" api:"required"`
	JSON accessSAMLCertificateListResponseCurrentCertificateJSON `json:"-"`
}

// accessSAMLCertificateListResponseCurrentCertificateJSON contains the JSON
// metadata for the struct [AccessSAMLCertificateListResponseCurrentCertificate]
type accessSAMLCertificateListResponseCurrentCertificateJSON struct {
	IsCurrent         apijson.Field
	NotAfter          apijson.Field
	PublicCertificate apijson.Field
	UID               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessSAMLCertificateListResponseCurrentCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateListResponseCurrentCertificateJSON) RawJSON() string {
	return r.raw
}

type AccessSAMLCertificateGetResponse struct {
	// When the certificate set was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Unique identifier for the certificate set
	UID string `json:"uid" api:"required"`
	// When the certificate set was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The current active certificate
	CurrentCertificate AccessSAMLCertificateGetResponseCurrentCertificate `json:"current_certificate"`
	// The previous certificate (maintained during rotation period). May be null when
	// no rotation has occurred. Mirrors the structure of `saml_certificate`.
	PreviousCertificate interface{}                          `json:"previous_certificate" api:"nullable"`
	JSON                accessSAMLCertificateGetResponseJSON `json:"-"`
}

// accessSAMLCertificateGetResponseJSON contains the JSON metadata for the struct
// [AccessSAMLCertificateGetResponse]
type accessSAMLCertificateGetResponseJSON struct {
	CreatedAt           apijson.Field
	UID                 apijson.Field
	UpdatedAt           apijson.Field
	CurrentCertificate  apijson.Field
	PreviousCertificate apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessSAMLCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateGetResponseJSON) RawJSON() string {
	return r.raw
}

// AccessSAMLCertificateGetResponseCurrentCertificate is the current active certificate
type AccessSAMLCertificateGetResponseCurrentCertificate struct {
	// Indicates whether the certificate can be used for IdP configuration.
	IsCurrent bool `json:"is_current" api:"required"`
	// Certificate expiration date
	NotAfter time.Time `json:"not_after" api:"required" format:"date-time"`
	// The public certificate in PEM format
	PublicCertificate string `json:"public_certificate" api:"required"`
	// Unique identifier for the certificate
	UID  string                                                 `json:"uid" api:"required"`
	JSON accessSAMLCertificateGetResponseCurrentCertificateJSON `json:"-"`
}

// accessSAMLCertificateGetResponseCurrentCertificateJSON contains the JSON
// metadata for the struct [AccessSAMLCertificateGetResponseCurrentCertificate]
type accessSAMLCertificateGetResponseCurrentCertificateJSON struct {
	IsCurrent         apijson.Field
	NotAfter          apijson.Field
	PublicCertificate apijson.Field
	UID               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessSAMLCertificateGetResponseCurrentCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateGetResponseCurrentCertificateJSON) RawJSON() string {
	return r.raw
}

type AccessSAMLCertificateRotateResponse struct {
	// When the certificate set was created
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Unique identifier for the certificate set
	UID string `json:"uid" api:"required"`
	// When the certificate set was last updated
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// The current active certificate
	CurrentCertificate AccessSAMLCertificateRotateResponseCurrentCertificate `json:"current_certificate"`
	// The previous certificate (maintained during rotation period). May be null when
	// no rotation has occurred. Mirrors the structure of `saml_certificate`.
	PreviousCertificate interface{}                             `json:"previous_certificate" api:"nullable"`
	JSON                accessSAMLCertificateRotateResponseJSON `json:"-"`
}

// accessSAMLCertificateRotateResponseJSON contains the JSON metadata for the
// struct [AccessSAMLCertificateRotateResponse]
type accessSAMLCertificateRotateResponseJSON struct {
	CreatedAt           apijson.Field
	UID                 apijson.Field
	UpdatedAt           apijson.Field
	CurrentCertificate  apijson.Field
	PreviousCertificate apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessSAMLCertificateRotateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateRotateResponseJSON) RawJSON() string {
	return r.raw
}

// AccessSAMLCertificateRotateResponseCurrentCertificate is the current active certificate
type AccessSAMLCertificateRotateResponseCurrentCertificate struct {
	// Indicates whether the certificate can be used for IdP configuration.
	IsCurrent bool `json:"is_current" api:"required"`
	// Certificate expiration date
	NotAfter time.Time `json:"not_after" api:"required" format:"date-time"`
	// The public certificate in PEM format
	PublicCertificate string `json:"public_certificate" api:"required"`
	// Unique identifier for the certificate
	UID  string                                                    `json:"uid" api:"required"`
	JSON accessSAMLCertificateRotateResponseCurrentCertificateJSON `json:"-"`
}

// accessSAMLCertificateRotateResponseCurrentCertificateJSON contains the JSON
// metadata for the struct [AccessSAMLCertificateRotateResponseCurrentCertificate]
type accessSAMLCertificateRotateResponseCurrentCertificateJSON struct {
	IsCurrent         apijson.Field
	NotAfter          apijson.Field
	PublicCertificate apijson.Field
	UID               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessSAMLCertificateRotateResponseCurrentCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateRotateResponseCurrentCertificateJSON) RawJSON() string {
	return r.raw
}

type AccessSAMLCertificateListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter by SAML certificate set UID. Accepts a comma-separated list of UIDs.
	ID param.Field[string] `query:"id"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [AccessSAMLCertificateListParams]'s query parameters as
// `url.Values`.
func (r AccessSAMLCertificateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessSAMLCertificateGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessSAMLCertificateGetResponseEnvelope struct {
	Errors   []AccessSAMLCertificateGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessSAMLCertificateGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessSAMLCertificateGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessSAMLCertificateGetResponse                `json:"result"`
	JSON    accessSAMLCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// accessSAMLCertificateGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessSAMLCertificateGetResponseEnvelope]
type accessSAMLCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSAMLCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessSAMLCertificateGetResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           AccessSAMLCertificateGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessSAMLCertificateGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessSAMLCertificateGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessSAMLCertificateGetResponseEnvelopeErrors]
type accessSAMLCertificateGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessSAMLCertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessSAMLCertificateGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    accessSAMLCertificateGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessSAMLCertificateGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessSAMLCertificateGetResponseEnvelopeErrorsSource]
type accessSAMLCertificateGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSAMLCertificateGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessSAMLCertificateGetResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccessSAMLCertificateGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessSAMLCertificateGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessSAMLCertificateGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessSAMLCertificateGetResponseEnvelopeMessages]
type accessSAMLCertificateGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessSAMLCertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessSAMLCertificateGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accessSAMLCertificateGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessSAMLCertificateGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessSAMLCertificateGetResponseEnvelopeMessagesSource]
type accessSAMLCertificateGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSAMLCertificateGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// AccessSAMLCertificateGetResponseEnvelopeSuccess indicates whether the API call was successful.
type AccessSAMLCertificateGetResponseEnvelopeSuccess bool

const (
	AccessSAMLCertificateGetResponseEnvelopeSuccessTrue AccessSAMLCertificateGetResponseEnvelopeSuccess = true
)

func (r AccessSAMLCertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessSAMLCertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessSAMLCertificateGetPemParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessSAMLCertificateRotateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type AccessSAMLCertificateRotateResponseEnvelope struct {
	Errors   []AccessSAMLCertificateRotateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []AccessSAMLCertificateRotateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success AccessSAMLCertificateRotateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  AccessSAMLCertificateRotateResponse                `json:"result"`
	JSON    accessSAMLCertificateRotateResponseEnvelopeJSON    `json:"-"`
}

// accessSAMLCertificateRotateResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessSAMLCertificateRotateResponseEnvelope]
type accessSAMLCertificateRotateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSAMLCertificateRotateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateRotateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessSAMLCertificateRotateResponseEnvelopeErrors struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           AccessSAMLCertificateRotateResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessSAMLCertificateRotateResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessSAMLCertificateRotateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessSAMLCertificateRotateResponseEnvelopeErrors]
type accessSAMLCertificateRotateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessSAMLCertificateRotateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateRotateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessSAMLCertificateRotateResponseEnvelopeErrorsSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    accessSAMLCertificateRotateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessSAMLCertificateRotateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessSAMLCertificateRotateResponseEnvelopeErrorsSource]
type accessSAMLCertificateRotateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSAMLCertificateRotateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateRotateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessSAMLCertificateRotateResponseEnvelopeMessages struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AccessSAMLCertificateRotateResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessSAMLCertificateRotateResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessSAMLCertificateRotateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessSAMLCertificateRotateResponseEnvelopeMessages]
type accessSAMLCertificateRotateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessSAMLCertificateRotateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateRotateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessSAMLCertificateRotateResponseEnvelopeMessagesSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    accessSAMLCertificateRotateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessSAMLCertificateRotateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessSAMLCertificateRotateResponseEnvelopeMessagesSource]
type accessSAMLCertificateRotateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessSAMLCertificateRotateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessSAMLCertificateRotateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// AccessSAMLCertificateRotateResponseEnvelopeSuccess indicates whether the API call was successful.
type AccessSAMLCertificateRotateResponseEnvelopeSuccess bool

const (
	AccessSAMLCertificateRotateResponseEnvelopeSuccessTrue AccessSAMLCertificateRotateResponseEnvelopeSuccess = true
)

func (r AccessSAMLCertificateRotateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessSAMLCertificateRotateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

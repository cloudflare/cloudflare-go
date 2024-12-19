// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificates

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// MTLSCertificateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMTLSCertificateService] method instead.
type MTLSCertificateService struct {
	Options      []option.RequestOption
	Associations *AssociationService
}

// NewMTLSCertificateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMTLSCertificateService(opts ...option.RequestOption) (r *MTLSCertificateService) {
	r = &MTLSCertificateService{}
	r.Options = opts
	r.Associations = NewAssociationService(opts...)
	return
}

// Upload a certificate that you want to use with mTLS-enabled Cloudflare services.
func (r *MTLSCertificateService) New(ctx context.Context, params MTLSCertificateNewParams, opts ...option.RequestOption) (res *MTLSCertificateNewResponse, err error) {
	var env MTLSCertificateNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mtls_certificates", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all mTLS certificates.
func (r *MTLSCertificateService) List(ctx context.Context, query MTLSCertificateListParams, opts ...option.RequestOption) (res *pagination.SinglePage[MTLSCertificate], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mtls_certificates", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists all mTLS certificates.
func (r *MTLSCertificateService) ListAutoPaging(ctx context.Context, query MTLSCertificateListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[MTLSCertificate] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes the mTLS certificate unless the certificate is in use by one or more
// Cloudflare services.
func (r *MTLSCertificateService) Delete(ctx context.Context, mtlsCertificateID string, body MTLSCertificateDeleteParams, opts ...option.RequestOption) (res *MTLSCertificate, err error) {
	var env MTLSCertificateDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if mtlsCertificateID == "" {
		err = errors.New("missing required mtls_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s", body.AccountID, mtlsCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single mTLS certificate.
func (r *MTLSCertificateService) Get(ctx context.Context, mtlsCertificateID string, query MTLSCertificateGetParams, opts ...option.RequestOption) (res *MTLSCertificate, err error) {
	var env MTLSCertificateGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if mtlsCertificateID == "" {
		err = errors.New("missing required mtls_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s", query.AccountID, mtlsCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MTLSCertificate struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether the certificate is a CA or leaf certificate.
	CA bool `json:"ca"`
	// The uploaded root CA certificate.
	Certificates string `json:"certificates"`
	// When the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// Optional unique name for the certificate. Only used for human readability.
	Name string `json:"name"`
	// The certificate serial number.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// This is the time the certificate was uploaded.
	UploadedOn time.Time           `json:"uploaded_on" format:"date-time"`
	JSON       mtlsCertificateJSON `json:"-"`
}

// mtlsCertificateJSON contains the JSON metadata for the struct [MTLSCertificate]
type mtlsCertificateJSON struct {
	ID           apijson.Field
	CA           apijson.Field
	Certificates apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	Name         apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MTLSCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mtlsCertificateJSON) RawJSON() string {
	return r.raw
}

type MTLSCertificateNewResponse struct {
	// Identifier
	ID string `json:"id"`
	// Indicates whether the certificate is a CA or leaf certificate.
	CA bool `json:"ca"`
	// The uploaded root CA certificate.
	Certificates string `json:"certificates"`
	// When the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// Optional unique name for the certificate. Only used for human readability.
	Name string `json:"name"`
	// The certificate serial number.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// This is the time the certificate was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// This is the time the certificate was uploaded.
	UploadedOn time.Time                      `json:"uploaded_on" format:"date-time"`
	JSON       mtlsCertificateNewResponseJSON `json:"-"`
}

// mtlsCertificateNewResponseJSON contains the JSON metadata for the struct
// [MTLSCertificateNewResponse]
type mtlsCertificateNewResponseJSON struct {
	ID           apijson.Field
	CA           apijson.Field
	Certificates apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	Name         apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	UpdatedAt    apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *MTLSCertificateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mtlsCertificateNewResponseJSON) RawJSON() string {
	return r.raw
}

type MTLSCertificateNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Indicates whether the certificate is a CA or leaf certificate.
	CA param.Field[bool] `json:"ca,required"`
	// The uploaded root CA certificate.
	Certificates param.Field[string] `json:"certificates,required"`
	// Optional unique name for the certificate. Only used for human readability.
	Name param.Field[string] `json:"name"`
	// The private key for the certificate
	PrivateKey param.Field[string] `json:"private_key"`
}

func (r MTLSCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MTLSCertificateNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success MTLSCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	Result  MTLSCertificateNewResponse                `json:"result"`
	JSON    mtlsCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// mtlsCertificateNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MTLSCertificateNewResponseEnvelope]
type mtlsCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MTLSCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mtlsCertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MTLSCertificateNewResponseEnvelopeSuccess bool

const (
	MTLSCertificateNewResponseEnvelopeSuccessTrue MTLSCertificateNewResponseEnvelopeSuccess = true
)

func (r MTLSCertificateNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MTLSCertificateNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MTLSCertificateListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type MTLSCertificateDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type MTLSCertificateDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success MTLSCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  MTLSCertificate                              `json:"result"`
	JSON    mtlsCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// mtlsCertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [MTLSCertificateDeleteResponseEnvelope]
type mtlsCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MTLSCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mtlsCertificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MTLSCertificateDeleteResponseEnvelopeSuccess bool

const (
	MTLSCertificateDeleteResponseEnvelopeSuccessTrue MTLSCertificateDeleteResponseEnvelopeSuccess = true
)

func (r MTLSCertificateDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MTLSCertificateDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MTLSCertificateGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type MTLSCertificateGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success MTLSCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	Result  MTLSCertificate                           `json:"result"`
	JSON    mtlsCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// mtlsCertificateGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MTLSCertificateGetResponseEnvelope]
type mtlsCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MTLSCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mtlsCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MTLSCertificateGetResponseEnvelopeSuccess bool

const (
	MTLSCertificateGetResponseEnvelopeSuccessTrue MTLSCertificateGetResponseEnvelopeSuccess = true
)

func (r MTLSCertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MTLSCertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

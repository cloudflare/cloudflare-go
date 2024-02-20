// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// CertificateAuthorityHostnameAssociationService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCertificateAuthorityHostnameAssociationService] method instead.
type CertificateAuthorityHostnameAssociationService struct {
	Options []option.RequestOption
}

// NewCertificateAuthorityHostnameAssociationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewCertificateAuthorityHostnameAssociationService(opts ...option.RequestOption) (r *CertificateAuthorityHostnameAssociationService) {
	r = &CertificateAuthorityHostnameAssociationService{}
	r.Options = opts
	return
}

// Replace Hostname Associations
func (r *CertificateAuthorityHostnameAssociationService) Update(ctx context.Context, zoneID string, body CertificateAuthorityHostnameAssociationUpdateParams, opts ...option.RequestOption) (res *CertificateAuthorityHostnameAssociationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificateAuthorityHostnameAssociationUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/certificate_authorities/hostname_associations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Hostname Associations
func (r *CertificateAuthorityHostnameAssociationService) List(ctx context.Context, zoneID string, query CertificateAuthorityHostnameAssociationListParams, opts ...option.RequestOption) (res *CertificateAuthorityHostnameAssociationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificateAuthorityHostnameAssociationListResponseEnvelope
	path := fmt.Sprintf("zones/%s/certificate_authorities/hostname_associations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CertificateAuthorityHostnameAssociationUpdateResponse struct {
	Hostnames []string `json:"hostnames"`
	// The UUID for a certificate that was uploaded to the mTLS Certificate Management
	// endpoint. If no mtls_certificate_id is given, the hostnames will be associated
	// to your active Cloudflare Managed CA.
	MtlsCertificateID string                                                    `json:"mtls_certificate_id"`
	JSON              certificateAuthorityHostnameAssociationUpdateResponseJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationUpdateResponseJSON contains the JSON
// metadata for the struct [CertificateAuthorityHostnameAssociationUpdateResponse]
type certificateAuthorityHostnameAssociationUpdateResponseJSON struct {
	Hostnames         apijson.Field
	MtlsCertificateID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateAuthorityHostnameAssociationListResponse struct {
	Hostnames []string `json:"hostnames"`
	// The UUID for a certificate that was uploaded to the mTLS Certificate Management
	// endpoint. If no mtls_certificate_id is given, the hostnames will be associated
	// to your active Cloudflare Managed CA.
	MtlsCertificateID string                                                  `json:"mtls_certificate_id"`
	JSON              certificateAuthorityHostnameAssociationListResponseJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationListResponseJSON contains the JSON
// metadata for the struct [CertificateAuthorityHostnameAssociationListResponse]
type certificateAuthorityHostnameAssociationListResponseJSON struct {
	Hostnames         apijson.Field
	MtlsCertificateID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateAuthorityHostnameAssociationUpdateParams struct {
	Hostnames param.Field[[]string] `json:"hostnames"`
	// The UUID for a certificate that was uploaded to the mTLS Certificate Management
	// endpoint. If no mtls_certificate_id is given, the hostnames will be associated
	// to your active Cloudflare Managed CA.
	MtlsCertificateID param.Field[string] `json:"mtls_certificate_id"`
}

func (r CertificateAuthorityHostnameAssociationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CertificateAuthorityHostnameAssociationUpdateResponseEnvelope struct {
	Errors   []CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   CertificateAuthorityHostnameAssociationUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    certificateAuthorityHostnameAssociationUpdateResponseEnvelopeJSON    `json:"-"`
}

// certificateAuthorityHostnameAssociationUpdateResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationUpdateResponseEnvelope]
type certificateAuthorityHostnameAssociationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeErrors struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    certificateAuthorityHostnameAssociationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationUpdateResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeErrors]
type certificateAuthorityHostnameAssociationUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeMessages struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    certificateAuthorityHostnameAssociationUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationUpdateResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeMessages]
type certificateAuthorityHostnameAssociationUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeSuccess bool

const (
	CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeSuccessTrue CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeSuccess = true
)

type CertificateAuthorityHostnameAssociationListParams struct {
	// The UUID to match against for a certificate that was uploaded to the mTLS
	// Certificate Management endpoint. If no mtls_certificate_id is given, the results
	// will be the hostnames associated to your active Cloudflare Managed CA.
	MtlsCertificateID param.Field[string] `query:"mtls_certificate_id"`
}

// URLQuery serializes [CertificateAuthorityHostnameAssociationListParams]'s query
// parameters as `url.Values`.
func (r CertificateAuthorityHostnameAssociationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CertificateAuthorityHostnameAssociationListResponseEnvelope struct {
	Errors   []CertificateAuthorityHostnameAssociationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificateAuthorityHostnameAssociationListResponseEnvelopeMessages `json:"messages,required"`
	Result   CertificateAuthorityHostnameAssociationListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CertificateAuthorityHostnameAssociationListResponseEnvelopeSuccess `json:"success,required"`
	JSON    certificateAuthorityHostnameAssociationListResponseEnvelopeJSON    `json:"-"`
}

// certificateAuthorityHostnameAssociationListResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationListResponseEnvelope]
type certificateAuthorityHostnameAssociationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateAuthorityHostnameAssociationListResponseEnvelopeErrors struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    certificateAuthorityHostnameAssociationListResponseEnvelopeErrorsJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationListResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationListResponseEnvelopeErrors]
type certificateAuthorityHostnameAssociationListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateAuthorityHostnameAssociationListResponseEnvelopeMessages struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    certificateAuthorityHostnameAssociationListResponseEnvelopeMessagesJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationListResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationListResponseEnvelopeMessages]
type certificateAuthorityHostnameAssociationListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificateAuthorityHostnameAssociationListResponseEnvelopeSuccess bool

const (
	CertificateAuthorityHostnameAssociationListResponseEnvelopeSuccessTrue CertificateAuthorityHostnameAssociationListResponseEnvelopeSuccess = true
)

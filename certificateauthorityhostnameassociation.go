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

// List Hostname Associations
func (r *CertificateAuthorityHostnameAssociationService) ClientCertificateForAZoneListHostnameAssociations(ctx context.Context, zoneID string, query CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsParams, opts ...option.RequestOption) (res *CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelope
	path := fmt.Sprintf("zones/%s/certificate_authorities/hostname_associations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Replace Hostname Associations
func (r *CertificateAuthorityHostnameAssociationService) ClientCertificateForAZonePutHostnameAssociations(ctx context.Context, zoneID string, body CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsParams, opts ...option.RequestOption) (res *CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelope
	path := fmt.Sprintf("zones/%s/certificate_authorities/hostname_associations", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponse struct {
	Hostnames []string `json:"hostnames"`
	// The UUID for a certificate that was uploaded to the mTLS Certificate Management
	// endpoint. If no mtls_certificate_id is given, the hostnames will be associated
	// to your active Cloudflare Managed CA.
	MtlsCertificateID string                                                                                               `json:"mtls_certificate_id"`
	JSON              certificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseJSON
// contains the JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponse]
type certificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseJSON struct {
	Hostnames         apijson.Field
	MtlsCertificateID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponse struct {
	Hostnames []string `json:"hostnames"`
	// The UUID for a certificate that was uploaded to the mTLS Certificate Management
	// endpoint. If no mtls_certificate_id is given, the hostnames will be associated
	// to your active Cloudflare Managed CA.
	MtlsCertificateID string                                                                                              `json:"mtls_certificate_id"`
	JSON              certificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseJSON
// contains the JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponse]
type certificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseJSON struct {
	Hostnames         apijson.Field
	MtlsCertificateID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsParams struct {
	// The UUID to match against for a certificate that was uploaded to the mTLS
	// Certificate Management endpoint. If no mtls_certificate_id is given, the results
	// will be the hostnames associated to your active Cloudflare Managed CA.
	MtlsCertificateID param.Field[string] `query:"mtls_certificate_id"`
}

// URLQuery serializes
// [CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsParams]'s
// query parameters as `url.Values`.
func (r CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelope struct {
	Errors   []CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeMessages `json:"messages,required"`
	Result   CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeSuccess `json:"success,required"`
	JSON    certificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeJSON    `json:"-"`
}

// certificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelope]
type certificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeErrors struct {
	Code    int64                                                                                                              `json:"code,required"`
	Message string                                                                                                             `json:"message,required"`
	JSON    certificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeErrorsJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeErrors]
type certificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeMessages struct {
	Code    int64                                                                                                                `json:"code,required"`
	Message string                                                                                                               `json:"message,required"`
	JSON    certificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeMessagesJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeMessages]
type certificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeSuccess bool

const (
	CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeSuccessTrue CertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseEnvelopeSuccess = true
)

type CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsParams struct {
	Hostnames param.Field[[]string] `json:"hostnames"`
	// The UUID for a certificate that was uploaded to the mTLS Certificate Management
	// endpoint. If no mtls_certificate_id is given, the hostnames will be associated
	// to your active Cloudflare Managed CA.
	MtlsCertificateID param.Field[string] `json:"mtls_certificate_id"`
}

func (r CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelope struct {
	Errors   []CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeMessages `json:"messages,required"`
	Result   CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeSuccess `json:"success,required"`
	JSON    certificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeJSON    `json:"-"`
}

// certificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelope]
type certificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeErrors struct {
	Code    int64                                                                                                             `json:"code,required"`
	Message string                                                                                                            `json:"message,required"`
	JSON    certificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeErrorsJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeErrors]
type certificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeMessages struct {
	Code    int64                                                                                                               `json:"code,required"`
	Message string                                                                                                              `json:"message,required"`
	JSON    certificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeMessagesJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeMessages]
type certificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeSuccess bool

const (
	CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeSuccessTrue CertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseEnvelopeSuccess = true
)

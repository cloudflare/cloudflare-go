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
func (r *CertificateAuthorityHostnameAssociationService) Update(ctx context.Context, params CertificateAuthorityHostnameAssociationUpdateParams, opts ...option.RequestOption) (res *CertificateAuthorityHostnameAssociationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificateAuthorityHostnameAssociationUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/certificate_authorities/hostname_associations", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Hostname Associations
func (r *CertificateAuthorityHostnameAssociationService) Get(ctx context.Context, params CertificateAuthorityHostnameAssociationGetParams, opts ...option.RequestOption) (res *CertificateAuthorityHostnameAssociationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env CertificateAuthorityHostnameAssociationGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/certificate_authorities/hostname_associations", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
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
	MTLSCertificateID string                                                    `json:"mtls_certificate_id"`
	JSON              certificateAuthorityHostnameAssociationUpdateResponseJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationUpdateResponseJSON contains the JSON
// metadata for the struct [CertificateAuthorityHostnameAssociationUpdateResponse]
type certificateAuthorityHostnameAssociationUpdateResponseJSON struct {
	Hostnames         apijson.Field
	MTLSCertificateID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateAuthorityHostnameAssociationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type CertificateAuthorityHostnameAssociationGetResponse struct {
	Hostnames []string `json:"hostnames"`
	// The UUID for a certificate that was uploaded to the mTLS Certificate Management
	// endpoint. If no mtls_certificate_id is given, the hostnames will be associated
	// to your active Cloudflare Managed CA.
	MTLSCertificateID string                                                 `json:"mtls_certificate_id"`
	JSON              certificateAuthorityHostnameAssociationGetResponseJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationGetResponseJSON contains the JSON
// metadata for the struct [CertificateAuthorityHostnameAssociationGetResponse]
type certificateAuthorityHostnameAssociationGetResponseJSON struct {
	Hostnames         apijson.Field
	MTLSCertificateID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateAuthorityHostnameAssociationGetResponseJSON) RawJSON() string {
	return r.raw
}

type CertificateAuthorityHostnameAssociationUpdateParams struct {
	// Identifier
	ZoneID    param.Field[string]   `path:"zone_id,required"`
	Hostnames param.Field[[]string] `json:"hostnames"`
	// The UUID for a certificate that was uploaded to the mTLS Certificate Management
	// endpoint. If no mtls_certificate_id is given, the hostnames will be associated
	// to your active Cloudflare Managed CA.
	MTLSCertificateID param.Field[string] `json:"mtls_certificate_id"`
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

func (r certificateAuthorityHostnameAssociationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r certificateAuthorityHostnameAssociationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r certificateAuthorityHostnameAssociationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeSuccess bool

const (
	CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeSuccessTrue CertificateAuthorityHostnameAssociationUpdateResponseEnvelopeSuccess = true
)

type CertificateAuthorityHostnameAssociationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The UUID to match against for a certificate that was uploaded to the mTLS
	// Certificate Management endpoint. If no mtls_certificate_id is given, the results
	// will be the hostnames associated to your active Cloudflare Managed CA.
	MTLSCertificateID param.Field[string] `query:"mtls_certificate_id"`
}

// URLQuery serializes [CertificateAuthorityHostnameAssociationGetParams]'s query
// parameters as `url.Values`.
func (r CertificateAuthorityHostnameAssociationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CertificateAuthorityHostnameAssociationGetResponseEnvelope struct {
	Errors   []CertificateAuthorityHostnameAssociationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CertificateAuthorityHostnameAssociationGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CertificateAuthorityHostnameAssociationGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success CertificateAuthorityHostnameAssociationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    certificateAuthorityHostnameAssociationGetResponseEnvelopeJSON    `json:"-"`
}

// certificateAuthorityHostnameAssociationGetResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [CertificateAuthorityHostnameAssociationGetResponseEnvelope]
type certificateAuthorityHostnameAssociationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateAuthorityHostnameAssociationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CertificateAuthorityHostnameAssociationGetResponseEnvelopeErrors struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    certificateAuthorityHostnameAssociationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationGetResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationGetResponseEnvelopeErrors]
type certificateAuthorityHostnameAssociationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateAuthorityHostnameAssociationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CertificateAuthorityHostnameAssociationGetResponseEnvelopeMessages struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    certificateAuthorityHostnameAssociationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// certificateAuthorityHostnameAssociationGetResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [CertificateAuthorityHostnameAssociationGetResponseEnvelopeMessages]
type certificateAuthorityHostnameAssociationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAuthorityHostnameAssociationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateAuthorityHostnameAssociationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CertificateAuthorityHostnameAssociationGetResponseEnvelopeSuccess bool

const (
	CertificateAuthorityHostnameAssociationGetResponseEnvelopeSuccessTrue CertificateAuthorityHostnameAssociationGetResponseEnvelopeSuccess = true
)

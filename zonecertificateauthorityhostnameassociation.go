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

// ZoneCertificateAuthorityHostnameAssociationService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCertificateAuthorityHostnameAssociationService] method instead.
type ZoneCertificateAuthorityHostnameAssociationService struct {
	Options []option.RequestOption
}

// NewZoneCertificateAuthorityHostnameAssociationService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewZoneCertificateAuthorityHostnameAssociationService(opts ...option.RequestOption) (r *ZoneCertificateAuthorityHostnameAssociationService) {
	r = &ZoneCertificateAuthorityHostnameAssociationService{}
	r.Options = opts
	return
}

// List Hostname Associations
func (r *ZoneCertificateAuthorityHostnameAssociationService) ClientCertificateForAZoneListHostnameAssociations(ctx context.Context, zoneIdentifier string, query ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsParams, opts ...option.RequestOption) (res *ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/certificate_authorities/hostname_associations", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Replace Hostname Associations
func (r *ZoneCertificateAuthorityHostnameAssociationService) ClientCertificateForAZonePutHostnameAssociations(ctx context.Context, zoneIdentifier string, body ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsParams, opts ...option.RequestOption) (res *ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/certificate_authorities/hostname_associations", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponse struct {
	Errors   []ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseError   `json:"errors"`
	Messages []ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseMessage `json:"messages"`
	Result   ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseSuccess `json:"success"`
	JSON    zoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseJSON    `json:"-"`
}

// zoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseJSON
// contains the JSON metadata for the struct
// [ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponse]
type zoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseError struct {
	Code    int64                                                                                                         `json:"code,required"`
	Message string                                                                                                        `json:"message,required"`
	JSON    zoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseErrorJSON `json:"-"`
}

// zoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseError]
type zoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseMessage struct {
	Code    int64                                                                                                           `json:"code,required"`
	Message string                                                                                                          `json:"message,required"`
	JSON    zoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseMessageJSON `json:"-"`
}

// zoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseMessage]
type zoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseResult struct {
	Hostnames []string `json:"hostnames"`
	// The UUID for a certificate that was uploaded to the mTLS Certificate Management
	// endpoint. If no mtls_certificate_id is given, the hostnames will be associated
	// to your active Cloudflare Managed CA.
	MtlsCertificateID string                                                                                                         `json:"mtls_certificate_id"`
	JSON              zoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseResultJSON `json:"-"`
}

// zoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseResult]
type zoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseResultJSON struct {
	Hostnames         apijson.Field
	MtlsCertificateID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseSuccess bool

const (
	ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseSuccessTrue ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsResponseSuccess = true
)

type ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponse struct {
	Errors   []ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseError   `json:"errors"`
	Messages []ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseMessage `json:"messages"`
	Result   ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseSuccess `json:"success"`
	JSON    zoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseJSON    `json:"-"`
}

// zoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseJSON
// contains the JSON metadata for the struct
// [ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponse]
type zoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseError struct {
	Code    int64                                                                                                        `json:"code,required"`
	Message string                                                                                                       `json:"message,required"`
	JSON    zoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseErrorJSON `json:"-"`
}

// zoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseError]
type zoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseMessage struct {
	Code    int64                                                                                                          `json:"code,required"`
	Message string                                                                                                         `json:"message,required"`
	JSON    zoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseMessageJSON `json:"-"`
}

// zoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseMessage]
type zoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseResult struct {
	Hostnames []string `json:"hostnames"`
	// The UUID for a certificate that was uploaded to the mTLS Certificate Management
	// endpoint. If no mtls_certificate_id is given, the hostnames will be associated
	// to your active Cloudflare Managed CA.
	MtlsCertificateID string                                                                                                        `json:"mtls_certificate_id"`
	JSON              zoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseResultJSON `json:"-"`
}

// zoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseResult]
type zoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseResultJSON struct {
	Hostnames         apijson.Field
	MtlsCertificateID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseSuccess bool

const (
	ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseSuccessTrue ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsResponseSuccess = true
)

type ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsParams struct {
	// The UUID to match against for a certificate that was uploaded to the mTLS
	// Certificate Management endpoint. If no mtls_certificate_id is given, the results
	// will be the hostnames associated to your active Cloudflare Managed CA.
	MtlsCertificateID param.Field[string] `query:"mtls_certificate_id"`
}

// URLQuery serializes
// [ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsParams]'s
// query parameters as `url.Values`.
func (r ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZoneListHostnameAssociationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsParams struct {
	Hostnames param.Field[[]string] `json:"hostnames"`
	// The UUID for a certificate that was uploaded to the mTLS Certificate Management
	// endpoint. If no mtls_certificate_id is given, the hostnames will be associated
	// to your active Cloudflare Managed CA.
	MtlsCertificateID param.Field[string] `json:"mtls_certificate_id"`
}

func (r ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

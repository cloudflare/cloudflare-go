// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
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
func (r *ZoneCertificateAuthorityHostnameAssociationService) ClientCertificateForAZoneListHostnameAssociations(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *HostnameAssociationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/certificate_authorities/hostname_associations", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Replace Hostname Associations
func (r *ZoneCertificateAuthorityHostnameAssociationService) ClientCertificateForAZonePutHostnameAssociations(ctx context.Context, zoneIdentifier string, body ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsParams, opts ...option.RequestOption) (res *HostnameAssociationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/certificate_authorities/hostname_associations", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type HostnameAssociationsResponse struct {
	Errors   []HostnameAssociationsResponseError   `json:"errors"`
	Messages []HostnameAssociationsResponseMessage `json:"messages"`
	Result   HostnameAssociationsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success HostnameAssociationsResponseSuccess `json:"success"`
	JSON    hostnameAssociationsResponseJSON    `json:"-"`
}

// hostnameAssociationsResponseJSON contains the JSON metadata for the struct
// [HostnameAssociationsResponse]
type hostnameAssociationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAssociationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameAssociationsResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    hostnameAssociationsResponseErrorJSON `json:"-"`
}

// hostnameAssociationsResponseErrorJSON contains the JSON metadata for the struct
// [HostnameAssociationsResponseError]
type hostnameAssociationsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAssociationsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameAssociationsResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    hostnameAssociationsResponseMessageJSON `json:"-"`
}

// hostnameAssociationsResponseMessageJSON contains the JSON metadata for the
// struct [HostnameAssociationsResponseMessage]
type hostnameAssociationsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAssociationsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HostnameAssociationsResponseResult struct {
	Hostnames []string                               `json:"hostnames"`
	JSON      hostnameAssociationsResponseResultJSON `json:"-"`
}

// hostnameAssociationsResponseResultJSON contains the JSON metadata for the struct
// [HostnameAssociationsResponseResult]
type hostnameAssociationsResponseResultJSON struct {
	Hostnames   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameAssociationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HostnameAssociationsResponseSuccess bool

const (
	HostnameAssociationsResponseSuccessTrue HostnameAssociationsResponseSuccess = true
)

type ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsParams struct {
	Hostnames param.Field[[]string] `json:"hostnames"`
}

func (r ZoneCertificateAuthorityHostnameAssociationClientCertificateForAZonePutHostnameAssociationsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

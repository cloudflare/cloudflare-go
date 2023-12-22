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

// ZoneAccessCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAccessCertificateService]
// method instead.
type ZoneAccessCertificateService struct {
	Options  []option.RequestOption
	Settings *ZoneAccessCertificateSettingService
}

// NewZoneAccessCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAccessCertificateService(opts ...option.RequestOption) (r *ZoneAccessCertificateService) {
	r = &ZoneAccessCertificateService{}
	r.Options = opts
	r.Settings = NewZoneAccessCertificateSettingService(opts...)
	return
}

// Fetches a single mTLS certificate.
func (r *ZoneAccessCertificateService) Get(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *CertificatesSingleResponseBEbBQusK, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/certificates/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured mTLS certificate.
func (r *ZoneAccessCertificateService) Update(ctx context.Context, identifier string, uuid string, body ZoneAccessCertificateUpdateParams, opts ...option.RequestOption) (res *CertificatesSingleResponseBEbBQusK, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/certificates/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all mTLS certificates.
func (r *ZoneAccessCertificateService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *CertificatesResponseCollectionFGyg2vQi, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/certificates", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an mTLS certificate.
func (r *ZoneAccessCertificateService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ZoneAccessCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/certificates/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new mTLS root certificate to Access.
func (r *ZoneAccessCertificateService) Add(ctx context.Context, identifier string, body ZoneAccessCertificateAddParams, opts ...option.RequestOption) (res *CertificatesSingleResponseBEbBQusK, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/certificates", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneAccessCertificateDeleteResponse struct {
	Errors   []ZoneAccessCertificateDeleteResponseError   `json:"errors"`
	Messages []ZoneAccessCertificateDeleteResponseMessage `json:"messages"`
	Result   ZoneAccessCertificateDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessCertificateDeleteResponseSuccess `json:"success"`
	JSON    zoneAccessCertificateDeleteResponseJSON    `json:"-"`
}

// zoneAccessCertificateDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateDeleteResponse]
type zoneAccessCertificateDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateDeleteResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneAccessCertificateDeleteResponseErrorJSON `json:"-"`
}

// zoneAccessCertificateDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateDeleteResponseError]
type zoneAccessCertificateDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateDeleteResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneAccessCertificateDeleteResponseMessageJSON `json:"-"`
}

// zoneAccessCertificateDeleteResponseMessageJSON contains the JSON metadata for
// the struct [ZoneAccessCertificateDeleteResponseMessage]
type zoneAccessCertificateDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateDeleteResponseResult struct {
	// UUID
	ID   string                                        `json:"id"`
	JSON zoneAccessCertificateDeleteResponseResultJSON `json:"-"`
}

// zoneAccessCertificateDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateDeleteResponseResult]
type zoneAccessCertificateDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessCertificateDeleteResponseSuccess bool

const (
	ZoneAccessCertificateDeleteResponseSuccessTrue ZoneAccessCertificateDeleteResponseSuccess = true
)

type ZoneAccessCertificateUpdateParams struct {
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneAccessCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneAccessCertificateAddParams struct {
	// The certificate content.
	Certificate param.Field[string] `json:"certificate,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name,required"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames"`
}

func (r ZoneAccessCertificateAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

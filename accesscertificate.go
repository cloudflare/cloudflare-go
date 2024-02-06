// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccessCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessCertificateService] method
// instead.
type AccessCertificateService struct {
	Options []option.RequestOption
}

// NewAccessCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessCertificateService(opts ...option.RequestOption) (r *AccessCertificateService) {
	r = &AccessCertificateService{}
	r.Options = opts
	return
}

// Fetches a single mTLS certificate.
func (r *AccessCertificateService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessCertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/certificates/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured mTLS certificate.
func (r *AccessCertificateService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, body AccessCertificateUpdateParams, opts ...option.RequestOption) (res *AccessCertificateUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/certificates/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an mTLS certificate.
func (r *AccessCertificateService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/certificates/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new mTLS root certificate to Access.
func (r *AccessCertificateService) AccessMTLSAuthenticationAddAnMTLSCertificate(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessCertificateAccessMTLSAuthenticationAddAnMTLSCertificateParams, opts ...option.RequestOption) (res *AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/certificates", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all mTLS root certificates.
func (r *AccessCertificateService) AccessMTLSAuthenticationListMTLSCertificates(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/certificates", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessCertificateGetResponse struct {
	Errors   []AccessCertificateGetResponseError   `json:"errors"`
	Messages []AccessCertificateGetResponseMessage `json:"messages"`
	Result   AccessCertificateGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessCertificateGetResponseSuccess `json:"success"`
	JSON    accessCertificateGetResponseJSON    `json:"-"`
}

// accessCertificateGetResponseJSON contains the JSON metadata for the struct
// [AccessCertificateGetResponse]
type accessCertificateGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateGetResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accessCertificateGetResponseErrorJSON `json:"-"`
}

// accessCertificateGetResponseErrorJSON contains the JSON metadata for the struct
// [AccessCertificateGetResponseError]
type accessCertificateGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateGetResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accessCertificateGetResponseMessageJSON `json:"-"`
}

// accessCertificateGetResponseMessageJSON contains the JSON metadata for the
// struct [AccessCertificateGetResponseMessage]
type accessCertificateGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateGetResponseResult struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                 `json:"name"`
	UpdatedAt time.Time                              `json:"updated_at" format:"date-time"`
	JSON      accessCertificateGetResponseResultJSON `json:"-"`
}

// accessCertificateGetResponseResultJSON contains the JSON metadata for the struct
// [AccessCertificateGetResponseResult]
type accessCertificateGetResponseResultJSON struct {
	ID                  apijson.Field
	AssociatedHostnames apijson.Field
	CreatedAt           apijson.Field
	ExpiresOn           apijson.Field
	Fingerprint         apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessCertificateGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCertificateGetResponseSuccess bool

const (
	AccessCertificateGetResponseSuccessTrue AccessCertificateGetResponseSuccess = true
)

type AccessCertificateUpdateResponse struct {
	Errors   []AccessCertificateUpdateResponseError   `json:"errors"`
	Messages []AccessCertificateUpdateResponseMessage `json:"messages"`
	Result   AccessCertificateUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessCertificateUpdateResponseSuccess `json:"success"`
	JSON    accessCertificateUpdateResponseJSON    `json:"-"`
}

// accessCertificateUpdateResponseJSON contains the JSON metadata for the struct
// [AccessCertificateUpdateResponse]
type accessCertificateUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateUpdateResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessCertificateUpdateResponseErrorJSON `json:"-"`
}

// accessCertificateUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccessCertificateUpdateResponseError]
type accessCertificateUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateUpdateResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accessCertificateUpdateResponseMessageJSON `json:"-"`
}

// accessCertificateUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccessCertificateUpdateResponseMessage]
type accessCertificateUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateUpdateResponseResult struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                    `json:"name"`
	UpdatedAt time.Time                                 `json:"updated_at" format:"date-time"`
	JSON      accessCertificateUpdateResponseResultJSON `json:"-"`
}

// accessCertificateUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccessCertificateUpdateResponseResult]
type accessCertificateUpdateResponseResultJSON struct {
	ID                  apijson.Field
	AssociatedHostnames apijson.Field
	CreatedAt           apijson.Field
	ExpiresOn           apijson.Field
	Fingerprint         apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessCertificateUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCertificateUpdateResponseSuccess bool

const (
	AccessCertificateUpdateResponseSuccessTrue AccessCertificateUpdateResponseSuccess = true
)

type AccessCertificateDeleteResponse struct {
	Errors   []AccessCertificateDeleteResponseError   `json:"errors"`
	Messages []AccessCertificateDeleteResponseMessage `json:"messages"`
	Result   AccessCertificateDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessCertificateDeleteResponseSuccess `json:"success"`
	JSON    accessCertificateDeleteResponseJSON    `json:"-"`
}

// accessCertificateDeleteResponseJSON contains the JSON metadata for the struct
// [AccessCertificateDeleteResponse]
type accessCertificateDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateDeleteResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessCertificateDeleteResponseErrorJSON `json:"-"`
}

// accessCertificateDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccessCertificateDeleteResponseError]
type accessCertificateDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateDeleteResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accessCertificateDeleteResponseMessageJSON `json:"-"`
}

// accessCertificateDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccessCertificateDeleteResponseMessage]
type accessCertificateDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateDeleteResponseResult struct {
	// UUID
	ID   string                                    `json:"id"`
	JSON accessCertificateDeleteResponseResultJSON `json:"-"`
}

// accessCertificateDeleteResponseResultJSON contains the JSON metadata for the
// struct [AccessCertificateDeleteResponseResult]
type accessCertificateDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCertificateDeleteResponseSuccess bool

const (
	AccessCertificateDeleteResponseSuccessTrue AccessCertificateDeleteResponseSuccess = true
)

type AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponse struct {
	Errors   []AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseError   `json:"errors"`
	Messages []AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseMessage `json:"messages"`
	Result   AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseSuccess `json:"success"`
	JSON    accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseJSON    `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponse]
type accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseError struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseErrorJSON `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseError]
type accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseMessage struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseMessageJSON `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseMessage]
type accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseResult struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                                                          `json:"name"`
	UpdatedAt time.Time                                                                       `json:"updated_at" format:"date-time"`
	JSON      accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseResultJSON `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseResultJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseResult]
type accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseResultJSON struct {
	ID                  apijson.Field
	AssociatedHostnames apijson.Field
	CreatedAt           apijson.Field
	ExpiresOn           apijson.Field
	Fingerprint         apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseSuccess bool

const (
	AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseSuccessTrue AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseSuccess = true
)

type AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponse struct {
	Errors     []AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseError    `json:"errors"`
	Messages   []AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseMessage  `json:"messages"`
	Result     []AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseResult   `json:"result"`
	ResultInfo AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseSuccess `json:"success"`
	JSON    accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseJSON    `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponse]
type accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseError struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseErrorJSON `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseError]
type accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseMessage struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseMessageJSON `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseMessage]
type accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseResult struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                                                          `json:"name"`
	UpdatedAt time.Time                                                                       `json:"updated_at" format:"date-time"`
	JSON      accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseResultJSON `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseResultJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseResult]
type accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseResultJSON struct {
	ID                  apijson.Field
	AssociatedHostnames apijson.Field
	CreatedAt           apijson.Field
	ExpiresOn           apijson.Field
	Fingerprint         apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                             `json:"total_count"`
	JSON       accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseResultInfoJSON `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseResultInfo]
type accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseSuccess bool

const (
	AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseSuccessTrue AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseSuccess = true
)

type AccessCertificateUpdateParams struct {
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name"`
}

func (r AccessCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateAccessMTLSAuthenticationAddAnMTLSCertificateParams struct {
	// The certificate content.
	Certificate param.Field[string] `json:"certificate,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name,required"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames"`
}

func (r AccessCertificateAccessMTLSAuthenticationAddAnMTLSCertificateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

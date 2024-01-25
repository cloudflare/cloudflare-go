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

// AccountAccessCertificateService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountAccessCertificateService] method instead.
type AccountAccessCertificateService struct {
	Options  []option.RequestOption
	Settings *AccountAccessCertificateSettingService
}

// NewAccountAccessCertificateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessCertificateService(opts ...option.RequestOption) (r *AccountAccessCertificateService) {
	r = &AccountAccessCertificateService{}
	r.Options = opts
	r.Settings = NewAccountAccessCertificateSettingService(opts...)
	return
}

// Fetches a single mTLS certificate.
func (r *AccountAccessCertificateService) Get(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccountAccessCertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/certificates/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured mTLS certificate.
func (r *AccountAccessCertificateService) Update(ctx context.Context, identifier string, uuid string, body AccountAccessCertificateUpdateParams, opts ...option.RequestOption) (res *AccountAccessCertificateUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/certificates/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an mTLS certificate.
func (r *AccountAccessCertificateService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *AccountAccessCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/certificates/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new mTLS root certificate to Access.
func (r *AccountAccessCertificateService) AccessMTlsAuthenticationAddAnMTlsCertificate(ctx context.Context, identifier string, body AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateParams, opts ...option.RequestOption) (res *AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/certificates", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all mTLS root certificates.
func (r *AccountAccessCertificateService) AccessMTlsAuthenticationListMTlsCertificates(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/certificates", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAccessCertificateGetResponse struct {
	Errors   []AccountAccessCertificateGetResponseError   `json:"errors"`
	Messages []AccountAccessCertificateGetResponseMessage `json:"messages"`
	Result   AccountAccessCertificateGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessCertificateGetResponseSuccess `json:"success"`
	JSON    accountAccessCertificateGetResponseJSON    `json:"-"`
}

// accountAccessCertificateGetResponseJSON contains the JSON metadata for the
// struct [AccountAccessCertificateGetResponse]
type accountAccessCertificateGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateGetResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountAccessCertificateGetResponseErrorJSON `json:"-"`
}

// accountAccessCertificateGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountAccessCertificateGetResponseError]
type accountAccessCertificateGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateGetResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountAccessCertificateGetResponseMessageJSON `json:"-"`
}

// accountAccessCertificateGetResponseMessageJSON contains the JSON metadata for
// the struct [AccountAccessCertificateGetResponseMessage]
type accountAccessCertificateGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateGetResponseResult struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                        `json:"name"`
	UpdatedAt time.Time                                     `json:"updated_at" format:"date-time"`
	JSON      accountAccessCertificateGetResponseResultJSON `json:"-"`
}

// accountAccessCertificateGetResponseResultJSON contains the JSON metadata for the
// struct [AccountAccessCertificateGetResponseResult]
type accountAccessCertificateGetResponseResultJSON struct {
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

func (r *AccountAccessCertificateGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessCertificateGetResponseSuccess bool

const (
	AccountAccessCertificateGetResponseSuccessTrue AccountAccessCertificateGetResponseSuccess = true
)

type AccountAccessCertificateUpdateResponse struct {
	Errors   []AccountAccessCertificateUpdateResponseError   `json:"errors"`
	Messages []AccountAccessCertificateUpdateResponseMessage `json:"messages"`
	Result   AccountAccessCertificateUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessCertificateUpdateResponseSuccess `json:"success"`
	JSON    accountAccessCertificateUpdateResponseJSON    `json:"-"`
}

// accountAccessCertificateUpdateResponseJSON contains the JSON metadata for the
// struct [AccountAccessCertificateUpdateResponse]
type accountAccessCertificateUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateUpdateResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountAccessCertificateUpdateResponseErrorJSON `json:"-"`
}

// accountAccessCertificateUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccountAccessCertificateUpdateResponseError]
type accountAccessCertificateUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateUpdateResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountAccessCertificateUpdateResponseMessageJSON `json:"-"`
}

// accountAccessCertificateUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccountAccessCertificateUpdateResponseMessage]
type accountAccessCertificateUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateUpdateResponseResult struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                           `json:"name"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	JSON      accountAccessCertificateUpdateResponseResultJSON `json:"-"`
}

// accountAccessCertificateUpdateResponseResultJSON contains the JSON metadata for
// the struct [AccountAccessCertificateUpdateResponseResult]
type accountAccessCertificateUpdateResponseResultJSON struct {
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

func (r *AccountAccessCertificateUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessCertificateUpdateResponseSuccess bool

const (
	AccountAccessCertificateUpdateResponseSuccessTrue AccountAccessCertificateUpdateResponseSuccess = true
)

type AccountAccessCertificateDeleteResponse struct {
	Errors   []AccountAccessCertificateDeleteResponseError   `json:"errors"`
	Messages []AccountAccessCertificateDeleteResponseMessage `json:"messages"`
	Result   AccountAccessCertificateDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessCertificateDeleteResponseSuccess `json:"success"`
	JSON    accountAccessCertificateDeleteResponseJSON    `json:"-"`
}

// accountAccessCertificateDeleteResponseJSON contains the JSON metadata for the
// struct [AccountAccessCertificateDeleteResponse]
type accountAccessCertificateDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateDeleteResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountAccessCertificateDeleteResponseErrorJSON `json:"-"`
}

// accountAccessCertificateDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountAccessCertificateDeleteResponseError]
type accountAccessCertificateDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateDeleteResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountAccessCertificateDeleteResponseMessageJSON `json:"-"`
}

// accountAccessCertificateDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountAccessCertificateDeleteResponseMessage]
type accountAccessCertificateDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateDeleteResponseResult struct {
	// UUID
	ID   string                                           `json:"id"`
	JSON accountAccessCertificateDeleteResponseResultJSON `json:"-"`
}

// accountAccessCertificateDeleteResponseResultJSON contains the JSON metadata for
// the struct [AccountAccessCertificateDeleteResponseResult]
type accountAccessCertificateDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessCertificateDeleteResponseSuccess bool

const (
	AccountAccessCertificateDeleteResponseSuccessTrue AccountAccessCertificateDeleteResponseSuccess = true
)

type AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponse struct {
	Errors   []AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseError   `json:"errors"`
	Messages []AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseMessage `json:"messages"`
	Result   AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseSuccess `json:"success"`
	JSON    accountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseJSON    `json:"-"`
}

// accountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponse]
type accountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseError struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseErrorJSON `json:"-"`
}

// accountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseError]
type accountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseMessage struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    accountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseMessageJSON `json:"-"`
}

// accountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseMessage]
type accountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseResult struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                                                                 `json:"name"`
	UpdatedAt time.Time                                                                              `json:"updated_at" format:"date-time"`
	JSON      accountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseResultJSON `json:"-"`
}

// accountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseResult]
type accountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseResultJSON struct {
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

func (r *AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseSuccess bool

const (
	AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseSuccessTrue AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateResponseSuccess = true
)

type AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponse struct {
	Errors     []AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseError    `json:"errors"`
	Messages   []AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseMessage  `json:"messages"`
	Result     []AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseResult   `json:"result"`
	ResultInfo AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseSuccess `json:"success"`
	JSON    accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseJSON    `json:"-"`
}

// accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseJSON
// contains the JSON metadata for the struct
// [AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponse]
type accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseError struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseErrorJSON `json:"-"`
}

// accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseError]
type accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseMessage struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseMessageJSON `json:"-"`
}

// accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseMessage]
type accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseResult struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                                                                 `json:"name"`
	UpdatedAt time.Time                                                                              `json:"updated_at" format:"date-time"`
	JSON      accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseResultJSON `json:"-"`
}

// accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseResult]
type accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseResultJSON struct {
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

func (r *AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                    `json:"total_count"`
	JSON       accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseResultInfoJSON `json:"-"`
}

// accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseResultInfo]
type accountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseSuccess bool

const (
	AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseSuccessTrue AccountAccessCertificateAccessMTlsAuthenticationListMTlsCertificatesResponseSuccess = true
)

type AccountAccessCertificateUpdateParams struct {
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name"`
}

func (r AccountAccessCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateParams struct {
	// The certificate content.
	Certificate param.Field[string] `json:"certificate,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name,required"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames"`
}

func (r AccountAccessCertificateAccessMTlsAuthenticationAddAnMTlsCertificateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

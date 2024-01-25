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
func (r *ZoneAccessCertificateService) Get(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ZoneAccessCertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/certificates/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured mTLS certificate.
func (r *ZoneAccessCertificateService) Update(ctx context.Context, identifier string, uuid string, body ZoneAccessCertificateUpdateParams, opts ...option.RequestOption) (res *ZoneAccessCertificateUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/certificates/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all mTLS certificates.
func (r *ZoneAccessCertificateService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ZoneAccessCertificateListResponse, err error) {
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
func (r *ZoneAccessCertificateService) Add(ctx context.Context, identifier string, body ZoneAccessCertificateAddParams, opts ...option.RequestOption) (res *ZoneAccessCertificateAddResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/certificates", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneAccessCertificateGetResponse struct {
	Errors   []ZoneAccessCertificateGetResponseError   `json:"errors"`
	Messages []ZoneAccessCertificateGetResponseMessage `json:"messages"`
	Result   ZoneAccessCertificateGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessCertificateGetResponseSuccess `json:"success"`
	JSON    zoneAccessCertificateGetResponseJSON    `json:"-"`
}

// zoneAccessCertificateGetResponseJSON contains the JSON metadata for the struct
// [ZoneAccessCertificateGetResponse]
type zoneAccessCertificateGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateGetResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneAccessCertificateGetResponseErrorJSON `json:"-"`
}

// zoneAccessCertificateGetResponseErrorJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateGetResponseError]
type zoneAccessCertificateGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateGetResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneAccessCertificateGetResponseMessageJSON `json:"-"`
}

// zoneAccessCertificateGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateGetResponseMessage]
type zoneAccessCertificateGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateGetResponseResult struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                     `json:"name"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      zoneAccessCertificateGetResponseResultJSON `json:"-"`
}

// zoneAccessCertificateGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateGetResponseResult]
type zoneAccessCertificateGetResponseResultJSON struct {
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

func (r *ZoneAccessCertificateGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessCertificateGetResponseSuccess bool

const (
	ZoneAccessCertificateGetResponseSuccessTrue ZoneAccessCertificateGetResponseSuccess = true
)

type ZoneAccessCertificateUpdateResponse struct {
	Errors   []ZoneAccessCertificateUpdateResponseError   `json:"errors"`
	Messages []ZoneAccessCertificateUpdateResponseMessage `json:"messages"`
	Result   ZoneAccessCertificateUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessCertificateUpdateResponseSuccess `json:"success"`
	JSON    zoneAccessCertificateUpdateResponseJSON    `json:"-"`
}

// zoneAccessCertificateUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateUpdateResponse]
type zoneAccessCertificateUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateUpdateResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneAccessCertificateUpdateResponseErrorJSON `json:"-"`
}

// zoneAccessCertificateUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateUpdateResponseError]
type zoneAccessCertificateUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateUpdateResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneAccessCertificateUpdateResponseMessageJSON `json:"-"`
}

// zoneAccessCertificateUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneAccessCertificateUpdateResponseMessage]
type zoneAccessCertificateUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateUpdateResponseResult struct {
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
	JSON      zoneAccessCertificateUpdateResponseResultJSON `json:"-"`
}

// zoneAccessCertificateUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateUpdateResponseResult]
type zoneAccessCertificateUpdateResponseResultJSON struct {
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

func (r *ZoneAccessCertificateUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessCertificateUpdateResponseSuccess bool

const (
	ZoneAccessCertificateUpdateResponseSuccessTrue ZoneAccessCertificateUpdateResponseSuccess = true
)

type ZoneAccessCertificateListResponse struct {
	Errors     []ZoneAccessCertificateListResponseError    `json:"errors"`
	Messages   []ZoneAccessCertificateListResponseMessage  `json:"messages"`
	Result     []ZoneAccessCertificateListResponseResult   `json:"result"`
	ResultInfo ZoneAccessCertificateListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneAccessCertificateListResponseSuccess `json:"success"`
	JSON    zoneAccessCertificateListResponseJSON    `json:"-"`
}

// zoneAccessCertificateListResponseJSON contains the JSON metadata for the struct
// [ZoneAccessCertificateListResponse]
type zoneAccessCertificateListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateListResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneAccessCertificateListResponseErrorJSON `json:"-"`
}

// zoneAccessCertificateListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateListResponseError]
type zoneAccessCertificateListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateListResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneAccessCertificateListResponseMessageJSON `json:"-"`
}

// zoneAccessCertificateListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateListResponseMessage]
type zoneAccessCertificateListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateListResponseResult struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                      `json:"name"`
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	JSON      zoneAccessCertificateListResponseResultJSON `json:"-"`
}

// zoneAccessCertificateListResponseResultJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateListResponseResult]
type zoneAccessCertificateListResponseResultJSON struct {
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

func (r *ZoneAccessCertificateListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       zoneAccessCertificateListResponseResultInfoJSON `json:"-"`
}

// zoneAccessCertificateListResponseResultInfoJSON contains the JSON metadata for
// the struct [ZoneAccessCertificateListResponseResultInfo]
type zoneAccessCertificateListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessCertificateListResponseSuccess bool

const (
	ZoneAccessCertificateListResponseSuccessTrue ZoneAccessCertificateListResponseSuccess = true
)

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

type ZoneAccessCertificateAddResponse struct {
	Errors   []ZoneAccessCertificateAddResponseError   `json:"errors"`
	Messages []ZoneAccessCertificateAddResponseMessage `json:"messages"`
	Result   ZoneAccessCertificateAddResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneAccessCertificateAddResponseSuccess `json:"success"`
	JSON    zoneAccessCertificateAddResponseJSON    `json:"-"`
}

// zoneAccessCertificateAddResponseJSON contains the JSON metadata for the struct
// [ZoneAccessCertificateAddResponse]
type zoneAccessCertificateAddResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateAddResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneAccessCertificateAddResponseErrorJSON `json:"-"`
}

// zoneAccessCertificateAddResponseErrorJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateAddResponseError]
type zoneAccessCertificateAddResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateAddResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateAddResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneAccessCertificateAddResponseMessageJSON `json:"-"`
}

// zoneAccessCertificateAddResponseMessageJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateAddResponseMessage]
type zoneAccessCertificateAddResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccessCertificateAddResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAccessCertificateAddResponseResult struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                     `json:"name"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      zoneAccessCertificateAddResponseResultJSON `json:"-"`
}

// zoneAccessCertificateAddResponseResultJSON contains the JSON metadata for the
// struct [ZoneAccessCertificateAddResponseResult]
type zoneAccessCertificateAddResponseResultJSON struct {
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

func (r *ZoneAccessCertificateAddResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneAccessCertificateAddResponseSuccess bool

const (
	ZoneAccessCertificateAddResponseSuccessTrue ZoneAccessCertificateAddResponseSuccess = true
)

type ZoneAccessCertificateUpdateParams struct {
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name"`
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

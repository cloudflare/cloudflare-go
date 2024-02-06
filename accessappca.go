// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccessAppCaService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessAppCaService] method
// instead.
type AccessAppCaService struct {
	Options []option.RequestOption
}

// NewAccessAppCaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessAppCaService(opts ...option.RequestOption) (r *AccessAppCaService) {
	r = &AccessAppCaService{}
	r.Options = opts
	return
}

// Deletes a short-lived certificate CA.
func (r *AccessAppCaService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessAppCaDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Generates a new short-lived certificate CA and public key.
func (r *AccessAppCaService) AccessShortLivedCertificateCAsNewAShortLivedCertificateCa(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Fetches a short-lived certificate CA and its public key.
func (r *AccessAppCaService) AccessShortLivedCertificateCAsGetAShortLivedCertificateCa(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists short-lived certificate CAs and their public keys.
func (r *AccessAppCaService) AccessShortLivedCertificateCAsListShortLivedCertificateCAs(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/apps/ca", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessAppCaDeleteResponse struct {
	Errors   []AccessAppCaDeleteResponseError   `json:"errors"`
	Messages []AccessAppCaDeleteResponseMessage `json:"messages"`
	Result   AccessAppCaDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccessAppCaDeleteResponseSuccess `json:"success"`
	JSON    accessAppCaDeleteResponseJSON    `json:"-"`
}

// accessAppCaDeleteResponseJSON contains the JSON metadata for the struct
// [AccessAppCaDeleteResponse]
type accessAppCaDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaDeleteResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    accessAppCaDeleteResponseErrorJSON `json:"-"`
}

// accessAppCaDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccessAppCaDeleteResponseError]
type accessAppCaDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaDeleteResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accessAppCaDeleteResponseMessageJSON `json:"-"`
}

// accessAppCaDeleteResponseMessageJSON contains the JSON metadata for the struct
// [AccessAppCaDeleteResponseMessage]
type accessAppCaDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaDeleteResponseResult struct {
	// The ID of the CA.
	ID   string                              `json:"id"`
	JSON accessAppCaDeleteResponseResultJSON `json:"-"`
}

// accessAppCaDeleteResponseResultJSON contains the JSON metadata for the struct
// [AccessAppCaDeleteResponseResult]
type accessAppCaDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppCaDeleteResponseSuccess bool

const (
	AccessAppCaDeleteResponseSuccessTrue AccessAppCaDeleteResponseSuccess = true
)

type AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponse struct {
	Errors   []AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseError   `json:"errors"`
	Messages []AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseMessage `json:"messages"`
	Result   interface{}                                                                           `json:"result"`
	// Whether the API call was successful
	Success AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseSuccess `json:"success"`
	JSON    accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseJSON    `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponse]
type accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseError struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseErrorJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseError]
type accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseMessage struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseMessageJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseMessage]
type accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseSuccess bool

const (
	AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseSuccessTrue AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseSuccess = true
)

type AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponse struct {
	Errors   []AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseError   `json:"errors"`
	Messages []AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseMessage `json:"messages"`
	Result   interface{}                                                                           `json:"result"`
	// Whether the API call was successful
	Success AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseSuccess `json:"success"`
	JSON    accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseJSON    `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponse]
type accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseError struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseErrorJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseError]
type accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseMessage struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseMessageJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseMessage]
type accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseSuccess bool

const (
	AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseSuccessTrue AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseSuccess = true
)

type AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponse struct {
	Errors     []AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseError    `json:"errors"`
	Messages   []AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseMessage  `json:"messages"`
	Result     []AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseResult   `json:"result"`
	ResultInfo AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseSuccess `json:"success"`
	JSON    accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseJSON    `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponse]
type accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseError struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseErrorJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseError]
type accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseMessage struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseMessageJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseMessage]
type accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseResult struct {
	// The ID of the CA.
	ID string `json:"id"`
	// The Application Audience (AUD) tag. Identifies the application associated with
	// the CA.
	Aud string `json:"aud"`
	// The public key to add to your SSH server configuration.
	PublicKey string                                                                                  `json:"public_key"`
	JSON      accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseResultJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseResultJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseResult]
type accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseResultJSON struct {
	ID          apijson.Field
	Aud         apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                     `json:"total_count"`
	JSON       accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseResultInfoJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseResultInfo]
type accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseSuccess bool

const (
	AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseSuccessTrue AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseSuccess = true
)

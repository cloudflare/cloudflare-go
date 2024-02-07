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

// AccessCertificateSettingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccessCertificateSettingService] method instead.
type AccessCertificateSettingService struct {
	Options []option.RequestOption
}

// NewAccessCertificateSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessCertificateSettingService(opts ...option.RequestOption) (r *AccessCertificateSettingService) {
	r = &AccessCertificateSettingService{}
	r.Options = opts
	return
}

// Updates an mTLS certificate's hostname settings.
func (r *AccessCertificateSettingService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessCertificateSettingUpdateParams, opts ...option.RequestOption) (res *AccessCertificateSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/certificates/settings", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all mTLS hostname settings for this account or zone.
func (r *AccessCertificateSettingService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *AccessCertificateSettingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("%s/%s/access/certificates/settings", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccessCertificateSettingUpdateResponse struct {
	Errors     []AccessCertificateSettingUpdateResponseError    `json:"errors"`
	Messages   []AccessCertificateSettingUpdateResponseMessage  `json:"messages"`
	Result     []AccessCertificateSettingUpdateResponseResult   `json:"result"`
	ResultInfo AccessCertificateSettingUpdateResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessCertificateSettingUpdateResponseSuccess `json:"success"`
	JSON    accessCertificateSettingUpdateResponseJSON    `json:"-"`
}

// accessCertificateSettingUpdateResponseJSON contains the JSON metadata for the
// struct [AccessCertificateSettingUpdateResponse]
type accessCertificateSettingUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateSettingUpdateResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accessCertificateSettingUpdateResponseErrorJSON `json:"-"`
}

// accessCertificateSettingUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccessCertificateSettingUpdateResponseError]
type accessCertificateSettingUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateSettingUpdateResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accessCertificateSettingUpdateResponseMessageJSON `json:"-"`
}

// accessCertificateSettingUpdateResponseMessageJSON contains the JSON metadata for
// the struct [AccessCertificateSettingUpdateResponseMessage]
type accessCertificateSettingUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateSettingUpdateResponseResult struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork bool `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding bool `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname string                                           `json:"hostname,required"`
	JSON     accessCertificateSettingUpdateResponseResultJSON `json:"-"`
}

// accessCertificateSettingUpdateResponseResultJSON contains the JSON metadata for
// the struct [AccessCertificateSettingUpdateResponseResult]
type accessCertificateSettingUpdateResponseResultJSON struct {
	ChinaNetwork                apijson.Field
	ClientCertificateForwarding apijson.Field
	Hostname                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AccessCertificateSettingUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateSettingUpdateResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       accessCertificateSettingUpdateResponseResultInfoJSON `json:"-"`
}

// accessCertificateSettingUpdateResponseResultInfoJSON contains the JSON metadata
// for the struct [AccessCertificateSettingUpdateResponseResultInfo]
type accessCertificateSettingUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCertificateSettingUpdateResponseSuccess bool

const (
	AccessCertificateSettingUpdateResponseSuccessTrue AccessCertificateSettingUpdateResponseSuccess = true
)

type AccessCertificateSettingListResponse struct {
	Errors     []AccessCertificateSettingListResponseError    `json:"errors"`
	Messages   []AccessCertificateSettingListResponseMessage  `json:"messages"`
	Result     []AccessCertificateSettingListResponseResult   `json:"result"`
	ResultInfo AccessCertificateSettingListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessCertificateSettingListResponseSuccess `json:"success"`
	JSON    accessCertificateSettingListResponseJSON    `json:"-"`
}

// accessCertificateSettingListResponseJSON contains the JSON metadata for the
// struct [AccessCertificateSettingListResponse]
type accessCertificateSettingListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateSettingListResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accessCertificateSettingListResponseErrorJSON `json:"-"`
}

// accessCertificateSettingListResponseErrorJSON contains the JSON metadata for the
// struct [AccessCertificateSettingListResponseError]
type accessCertificateSettingListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateSettingListResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accessCertificateSettingListResponseMessageJSON `json:"-"`
}

// accessCertificateSettingListResponseMessageJSON contains the JSON metadata for
// the struct [AccessCertificateSettingListResponseMessage]
type accessCertificateSettingListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateSettingListResponseResult struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork bool `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding bool `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname string                                         `json:"hostname,required"`
	JSON     accessCertificateSettingListResponseResultJSON `json:"-"`
}

// accessCertificateSettingListResponseResultJSON contains the JSON metadata for
// the struct [AccessCertificateSettingListResponseResult]
type accessCertificateSettingListResponseResultJSON struct {
	ChinaNetwork                apijson.Field
	ClientCertificateForwarding apijson.Field
	Hostname                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AccessCertificateSettingListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateSettingListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       accessCertificateSettingListResponseResultInfoJSON `json:"-"`
}

// accessCertificateSettingListResponseResultInfoJSON contains the JSON metadata
// for the struct [AccessCertificateSettingListResponseResultInfo]
type accessCertificateSettingListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateSettingListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCertificateSettingListResponseSuccess bool

const (
	AccessCertificateSettingListResponseSuccessTrue AccessCertificateSettingListResponseSuccess = true
)

type AccessCertificateSettingUpdateParams struct {
	Settings param.Field[[]AccessCertificateSettingUpdateParamsSetting] `json:"settings,required"`
}

func (r AccessCertificateSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateSettingUpdateParamsSetting struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork param.Field[bool] `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding param.Field[bool] `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname param.Field[string] `json:"hostname,required"`
}

func (r AccessCertificateSettingUpdateParamsSetting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

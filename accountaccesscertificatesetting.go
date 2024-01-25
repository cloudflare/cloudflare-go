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

// AccountAccessCertificateSettingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAccessCertificateSettingService] method instead.
type AccountAccessCertificateSettingService struct {
	Options []option.RequestOption
}

// NewAccountAccessCertificateSettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAccessCertificateSettingService(opts ...option.RequestOption) (r *AccountAccessCertificateSettingService) {
	r = &AccountAccessCertificateSettingService{}
	r.Options = opts
	return
}

// Updates an mTLS certificate's hostname settings.
func (r *AccountAccessCertificateSettingService) Update(ctx context.Context, identifier string, body AccountAccessCertificateSettingUpdateParams, opts ...option.RequestOption) (res *AccountAccessCertificateSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/access/certificates/settings", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all mTLS hostname settings for this account.
func (r *AccountAccessCertificateSettingService) List(ctx context.Context, identifier string, opts ...option.RequestOption) (res *AccountAccessCertificateSettingListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/access/certificates/settings", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAccessCertificateSettingUpdateResponse struct {
	Errors     []AccountAccessCertificateSettingUpdateResponseError    `json:"errors"`
	Messages   []AccountAccessCertificateSettingUpdateResponseMessage  `json:"messages"`
	Result     []AccountAccessCertificateSettingUpdateResponseResult   `json:"result"`
	ResultInfo AccountAccessCertificateSettingUpdateResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAccessCertificateSettingUpdateResponseSuccess `json:"success"`
	JSON    accountAccessCertificateSettingUpdateResponseJSON    `json:"-"`
}

// accountAccessCertificateSettingUpdateResponseJSON contains the JSON metadata for
// the struct [AccountAccessCertificateSettingUpdateResponse]
type accountAccessCertificateSettingUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateSettingUpdateResponseError struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountAccessCertificateSettingUpdateResponseErrorJSON `json:"-"`
}

// accountAccessCertificateSettingUpdateResponseErrorJSON contains the JSON
// metadata for the struct [AccountAccessCertificateSettingUpdateResponseError]
type accountAccessCertificateSettingUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateSettingUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateSettingUpdateResponseMessage struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountAccessCertificateSettingUpdateResponseMessageJSON `json:"-"`
}

// accountAccessCertificateSettingUpdateResponseMessageJSON contains the JSON
// metadata for the struct [AccountAccessCertificateSettingUpdateResponseMessage]
type accountAccessCertificateSettingUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateSettingUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateSettingUpdateResponseResult struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork bool `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding bool `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname string                                                  `json:"hostname,required"`
	JSON     accountAccessCertificateSettingUpdateResponseResultJSON `json:"-"`
}

// accountAccessCertificateSettingUpdateResponseResultJSON contains the JSON
// metadata for the struct [AccountAccessCertificateSettingUpdateResponseResult]
type accountAccessCertificateSettingUpdateResponseResultJSON struct {
	ChinaNetwork                apijson.Field
	ClientCertificateForwarding apijson.Field
	Hostname                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AccountAccessCertificateSettingUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateSettingUpdateResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                     `json:"total_count"`
	JSON       accountAccessCertificateSettingUpdateResponseResultInfoJSON `json:"-"`
}

// accountAccessCertificateSettingUpdateResponseResultInfoJSON contains the JSON
// metadata for the struct
// [AccountAccessCertificateSettingUpdateResponseResultInfo]
type accountAccessCertificateSettingUpdateResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateSettingUpdateResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessCertificateSettingUpdateResponseSuccess bool

const (
	AccountAccessCertificateSettingUpdateResponseSuccessTrue AccountAccessCertificateSettingUpdateResponseSuccess = true
)

type AccountAccessCertificateSettingListResponse struct {
	Errors     []AccountAccessCertificateSettingListResponseError    `json:"errors"`
	Messages   []AccountAccessCertificateSettingListResponseMessage  `json:"messages"`
	Result     []AccountAccessCertificateSettingListResponseResult   `json:"result"`
	ResultInfo AccountAccessCertificateSettingListResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountAccessCertificateSettingListResponseSuccess `json:"success"`
	JSON    accountAccessCertificateSettingListResponseJSON    `json:"-"`
}

// accountAccessCertificateSettingListResponseJSON contains the JSON metadata for
// the struct [AccountAccessCertificateSettingListResponse]
type accountAccessCertificateSettingListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateSettingListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateSettingListResponseError struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountAccessCertificateSettingListResponseErrorJSON `json:"-"`
}

// accountAccessCertificateSettingListResponseErrorJSON contains the JSON metadata
// for the struct [AccountAccessCertificateSettingListResponseError]
type accountAccessCertificateSettingListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateSettingListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateSettingListResponseMessage struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    accountAccessCertificateSettingListResponseMessageJSON `json:"-"`
}

// accountAccessCertificateSettingListResponseMessageJSON contains the JSON
// metadata for the struct [AccountAccessCertificateSettingListResponseMessage]
type accountAccessCertificateSettingListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateSettingListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateSettingListResponseResult struct {
	// Request client certificates for this hostname in China. Can only be set to true
	// if this zone is china network enabled.
	ChinaNetwork bool `json:"china_network,required"`
	// Client Certificate Forwarding is a feature that takes the client cert provided
	// by the eyeball to the edge, and forwards it to the origin as a HTTP header to
	// allow logging on the origin.
	ClientCertificateForwarding bool `json:"client_certificate_forwarding,required"`
	// The hostname that these settings apply to.
	Hostname string                                                `json:"hostname,required"`
	JSON     accountAccessCertificateSettingListResponseResultJSON `json:"-"`
}

// accountAccessCertificateSettingListResponseResultJSON contains the JSON metadata
// for the struct [AccountAccessCertificateSettingListResponseResult]
type accountAccessCertificateSettingListResponseResultJSON struct {
	ChinaNetwork                apijson.Field
	ClientCertificateForwarding apijson.Field
	Hostname                    apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AccountAccessCertificateSettingListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAccessCertificateSettingListResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                   `json:"total_count"`
	JSON       accountAccessCertificateSettingListResponseResultInfoJSON `json:"-"`
}

// accountAccessCertificateSettingListResponseResultInfoJSON contains the JSON
// metadata for the struct [AccountAccessCertificateSettingListResponseResultInfo]
type accountAccessCertificateSettingListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAccessCertificateSettingListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAccessCertificateSettingListResponseSuccess bool

const (
	AccountAccessCertificateSettingListResponseSuccessTrue AccountAccessCertificateSettingListResponseSuccess = true
)

type AccountAccessCertificateSettingUpdateParams struct {
	Settings param.Field[[]AccountAccessCertificateSettingUpdateParamsSetting] `json:"settings,required"`
}

func (r AccountAccessCertificateSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccessCertificateSettingUpdateParamsSetting struct {
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

func (r AccountAccessCertificateSettingUpdateParamsSetting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

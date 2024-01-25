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

// AccountImageV1StatService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountImageV1StatService] method
// instead.
type AccountImageV1StatService struct {
	Options []option.RequestOption
}

// NewAccountImageV1StatService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountImageV1StatService(opts ...option.RequestOption) (r *AccountImageV1StatService) {
	r = &AccountImageV1StatService{}
	r.Options = opts
	return
}

// Fetch usage statistics details for Cloudflare Images.
func (r *AccountImageV1StatService) CloudflareImagesImagesUsageStatistics(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/stats", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponse struct {
	Errors   []AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseError   `json:"errors"`
	Messages []AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseMessage `json:"messages"`
	Result   AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseSuccess `json:"success"`
	JSON    accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseJSON    `json:"-"`
}

// accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseJSON contains the
// JSON metadata for the struct
// [AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponse]
type accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseError struct {
	Code    int64                                                                    `json:"code,required"`
	Message string                                                                   `json:"message,required"`
	JSON    accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseErrorJSON `json:"-"`
}

// accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseError]
type accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseMessage struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseMessageJSON `json:"-"`
}

// accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseMessage]
type accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseResult struct {
	Count AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseResultCount `json:"count"`
	JSON  accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseResultJSON  `json:"-"`
}

// accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseResult]
type accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseResultJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseResultCount struct {
	// Cloudflare Images allowed usage.
	Allowed float64 `json:"allowed"`
	// Cloudflare Images current usage.
	Current float64                                                                        `json:"current"`
	JSON    accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseResultCountJSON `json:"-"`
}

// accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseResultCountJSON
// contains the JSON metadata for the struct
// [AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseResultCount]
type accountImageV1StatCloudflareImagesImagesUsageStatisticsResponseResultCountJSON struct {
	Allowed     apijson.Field
	Current     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseResultCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseSuccess bool

const (
	AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseSuccessTrue AccountImageV1StatCloudflareImagesImagesUsageStatisticsResponseSuccess = true
)

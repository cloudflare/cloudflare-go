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

// ImageV1StatService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewImageV1StatService] method
// instead.
type ImageV1StatService struct {
	Options []option.RequestOption
}

// NewImageV1StatService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewImageV1StatService(opts ...option.RequestOption) (r *ImageV1StatService) {
	r = &ImageV1StatService{}
	r.Options = opts
	return
}

// Fetch usage statistics details for Cloudflare Images.
func (r *ImageV1StatService) CloudflareImagesImagesUsageStatistics(ctx context.Context, accountID string, opts ...option.RequestOption) (res *ImageV1StatCloudflareImagesImagesUsageStatisticsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/stats", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ImageV1StatCloudflareImagesImagesUsageStatisticsResponse struct {
	Count ImageV1StatCloudflareImagesImagesUsageStatisticsResponseCount `json:"count"`
	JSON  imageV1StatCloudflareImagesImagesUsageStatisticsResponseJSON  `json:"-"`
}

// imageV1StatCloudflareImagesImagesUsageStatisticsResponseJSON contains the JSON
// metadata for the struct
// [ImageV1StatCloudflareImagesImagesUsageStatisticsResponse]
type imageV1StatCloudflareImagesImagesUsageStatisticsResponseJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1StatCloudflareImagesImagesUsageStatisticsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1StatCloudflareImagesImagesUsageStatisticsResponseCount struct {
	// Cloudflare Images allowed usage.
	Allowed float64 `json:"allowed"`
	// Cloudflare Images current usage.
	Current float64                                                           `json:"current"`
	JSON    imageV1StatCloudflareImagesImagesUsageStatisticsResponseCountJSON `json:"-"`
}

// imageV1StatCloudflareImagesImagesUsageStatisticsResponseCountJSON contains the
// JSON metadata for the struct
// [ImageV1StatCloudflareImagesImagesUsageStatisticsResponseCount]
type imageV1StatCloudflareImagesImagesUsageStatisticsResponseCountJSON struct {
	Allowed     apijson.Field
	Current     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1StatCloudflareImagesImagesUsageStatisticsResponseCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelope struct {
	Errors   []ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageV1StatCloudflareImagesImagesUsageStatisticsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeJSON    `json:"-"`
}

// imageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelope]
type imageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeErrors struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    imageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeErrors]
type imageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeMessages struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    imageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeMessages]
type imageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeSuccess bool

const (
	ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeSuccessTrue ImageV1StatCloudflareImagesImagesUsageStatisticsResponseEnvelopeSuccess = true
)

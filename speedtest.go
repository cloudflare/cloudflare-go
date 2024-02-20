// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SpeedTestService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSpeedTestService] method instead.
type SpeedTestService struct {
	Options []option.RequestOption
}

// NewSpeedTestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpeedTestService(opts ...option.RequestOption) (r *SpeedTestService) {
	r = &SpeedTestService{}
	r.Options = opts
	return
}

// Test history (list of tests) for a specific webpage.
func (r *SpeedTestService) List(ctx context.Context, zoneID string, url string, query SpeedTestListParams, opts ...option.RequestOption) (res *SpeedTestListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/speed_api/pages/%s/tests", zoneID, url)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SpeedTestListResponse struct {
	Errors   []SpeedTestListResponseError   `json:"errors,required"`
	Messages []SpeedTestListResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success    bool                            `json:"success,required"`
	ResultInfo SpeedTestListResponseResultInfo `json:"result_info"`
	JSON       speedTestListResponseJSON       `json:"-"`
}

// speedTestListResponseJSON contains the JSON metadata for the struct
// [SpeedTestListResponse]
type speedTestListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedTestListResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    speedTestListResponseErrorJSON `json:"-"`
}

// speedTestListResponseErrorJSON contains the JSON metadata for the struct
// [SpeedTestListResponseError]
type speedTestListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedTestListResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    speedTestListResponseMessageJSON `json:"-"`
}

// speedTestListResponseMessageJSON contains the JSON metadata for the struct
// [SpeedTestListResponseMessage]
type speedTestListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedTestListResponseResultInfo struct {
	Count      int64                               `json:"count"`
	Page       int64                               `json:"page"`
	PerPage    int64                               `json:"per_page"`
	TotalCount int64                               `json:"total_count"`
	JSON       speedTestListResponseResultInfoJSON `json:"-"`
}

// speedTestListResponseResultInfoJSON contains the JSON metadata for the struct
// [SpeedTestListResponseResultInfo]
type speedTestListResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpeedTestListResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpeedTestListParams struct {
	Page    param.Field[int64] `query:"page"`
	PerPage param.Field[int64] `query:"per_page"`
	// A test region.
	Region param.Field[SpeedTestListParamsRegion] `query:"region"`
}

// URLQuery serializes [SpeedTestListParams]'s query parameters as `url.Values`.
func (r SpeedTestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A test region.
type SpeedTestListParamsRegion string

const (
	SpeedTestListParamsRegionAsiaEast1           SpeedTestListParamsRegion = "asia-east1"
	SpeedTestListParamsRegionAsiaNortheast1      SpeedTestListParamsRegion = "asia-northeast1"
	SpeedTestListParamsRegionAsiaNortheast2      SpeedTestListParamsRegion = "asia-northeast2"
	SpeedTestListParamsRegionAsiaSouth1          SpeedTestListParamsRegion = "asia-south1"
	SpeedTestListParamsRegionAsiaSoutheast1      SpeedTestListParamsRegion = "asia-southeast1"
	SpeedTestListParamsRegionAustraliaSoutheast1 SpeedTestListParamsRegion = "australia-southeast1"
	SpeedTestListParamsRegionEuropeNorth1        SpeedTestListParamsRegion = "europe-north1"
	SpeedTestListParamsRegionEuropeSouthwest1    SpeedTestListParamsRegion = "europe-southwest1"
	SpeedTestListParamsRegionEuropeWest1         SpeedTestListParamsRegion = "europe-west1"
	SpeedTestListParamsRegionEuropeWest2         SpeedTestListParamsRegion = "europe-west2"
	SpeedTestListParamsRegionEuropeWest3         SpeedTestListParamsRegion = "europe-west3"
	SpeedTestListParamsRegionEuropeWest4         SpeedTestListParamsRegion = "europe-west4"
	SpeedTestListParamsRegionEuropeWest8         SpeedTestListParamsRegion = "europe-west8"
	SpeedTestListParamsRegionEuropeWest9         SpeedTestListParamsRegion = "europe-west9"
	SpeedTestListParamsRegionMeWest1             SpeedTestListParamsRegion = "me-west1"
	SpeedTestListParamsRegionSouthamericaEast1   SpeedTestListParamsRegion = "southamerica-east1"
	SpeedTestListParamsRegionUsCentral1          SpeedTestListParamsRegion = "us-central1"
	SpeedTestListParamsRegionUsEast1             SpeedTestListParamsRegion = "us-east1"
	SpeedTestListParamsRegionUsEast4             SpeedTestListParamsRegion = "us-east4"
	SpeedTestListParamsRegionUsSouth1            SpeedTestListParamsRegion = "us-south1"
	SpeedTestListParamsRegionUsWest1             SpeedTestListParamsRegion = "us-west1"
)

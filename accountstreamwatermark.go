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

// AccountStreamWatermarkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStreamWatermarkService]
// method instead.
type AccountStreamWatermarkService struct {
	Options []option.RequestOption
}

// NewAccountStreamWatermarkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStreamWatermarkService(opts ...option.RequestOption) (r *AccountStreamWatermarkService) {
	r = &AccountStreamWatermarkService{}
	r.Options = opts
	return
}

// Retrieves details for a single watermark profile.
func (r *AccountStreamWatermarkService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *WatermarkResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/watermarks/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a watermark profile.
func (r *AccountStreamWatermarkService) Delete(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountStreamWatermarkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/watermarks/%s", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates watermark profiles using a single `HTTP POST multipart/form-data`
// request.
func (r *AccountStreamWatermarkService) StreamWatermarkProfileNewWatermarkProfilesViaBasicUpload(ctx context.Context, accountIdentifier string, body AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadParams, opts ...option.RequestOption) (res *WatermarkResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/watermarks", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all watermark profiles for an account.
func (r *AccountStreamWatermarkService) StreamWatermarkProfileListWatermarkProfiles(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *WatermarkResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/watermarks", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type WatermarkResponseCollection struct {
	Errors     []WatermarkResponseCollectionError    `json:"errors"`
	Messages   []WatermarkResponseCollectionMessage  `json:"messages"`
	Result     []WatermarkResponseCollectionResult   `json:"result"`
	ResultInfo WatermarkResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success WatermarkResponseCollectionSuccess `json:"success"`
	JSON    watermarkResponseCollectionJSON    `json:"-"`
}

// watermarkResponseCollectionJSON contains the JSON metadata for the struct
// [WatermarkResponseCollection]
type watermarkResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatermarkResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WatermarkResponseCollectionError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    watermarkResponseCollectionErrorJSON `json:"-"`
}

// watermarkResponseCollectionErrorJSON contains the JSON metadata for the struct
// [WatermarkResponseCollectionError]
type watermarkResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatermarkResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WatermarkResponseCollectionMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    watermarkResponseCollectionMessageJSON `json:"-"`
}

// watermarkResponseCollectionMessageJSON contains the JSON metadata for the struct
// [WatermarkResponseCollectionMessage]
type watermarkResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatermarkResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WatermarkResponseCollectionResult struct {
	// The date and a time a watermark profile was created.
	Created time.Time `json:"created" format:"date-time"`
	// The source URL for a downloaded image. If the watermark profile was created via
	// direct upload, this field is null.
	DownloadedFrom string `json:"downloadedFrom"`
	// The height of the image in pixels.
	Height int64 `json:"height"`
	// A short description of the watermark profile.
	Name string `json:"name"`
	// The translucency of the image. A value of `0.0` makes the image completely
	// transparent, and `1.0` makes the image completely opaque. Note that if the image
	// is already semi-transparent, setting this to `1.0` will not make the image
	// completely opaque.
	Opacity float64 `json:"opacity"`
	// The whitespace between the adjacent edges (determined by position) of the video
	// and the image. `0.0` indicates no padding, and `1.0` indicates a fully padded
	// video width or length, as determined by the algorithm.
	Padding float64 `json:"padding"`
	// The location of the image. Valid positions are: `upperRight`, `upperLeft`,
	// `lowerLeft`, `lowerRight`, and `center`. Note that `center` ignores the
	// `padding` parameter.
	Position string `json:"position"`
	// The size of the image relative to the overall size of the video. This parameter
	// will adapt to horizontal and vertical videos automatically. `0.0` indicates no
	// scaling (use the size of the image as-is), and `1.0 `fills the entire video.
	Scale float64 `json:"scale"`
	// The size of the image in bytes.
	Size float64 `json:"size"`
	// The unique identifier for a watermark profile.
	Uid string `json:"uid"`
	// The width of the image in pixels.
	Width int64                                 `json:"width"`
	JSON  watermarkResponseCollectionResultJSON `json:"-"`
}

// watermarkResponseCollectionResultJSON contains the JSON metadata for the struct
// [WatermarkResponseCollectionResult]
type watermarkResponseCollectionResultJSON struct {
	Created        apijson.Field
	DownloadedFrom apijson.Field
	Height         apijson.Field
	Name           apijson.Field
	Opacity        apijson.Field
	Padding        apijson.Field
	Position       apijson.Field
	Scale          apijson.Field
	Size           apijson.Field
	Uid            apijson.Field
	Width          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *WatermarkResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WatermarkResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       watermarkResponseCollectionResultInfoJSON `json:"-"`
}

// watermarkResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [WatermarkResponseCollectionResultInfo]
type watermarkResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatermarkResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WatermarkResponseCollectionSuccess bool

const (
	WatermarkResponseCollectionSuccessTrue WatermarkResponseCollectionSuccess = true
)

type WatermarkResponseSingle struct {
	Errors   []WatermarkResponseSingleError   `json:"errors"`
	Messages []WatermarkResponseSingleMessage `json:"messages"`
	Result   interface{}                      `json:"result"`
	// Whether the API call was successful
	Success WatermarkResponseSingleSuccess `json:"success"`
	JSON    watermarkResponseSingleJSON    `json:"-"`
}

// watermarkResponseSingleJSON contains the JSON metadata for the struct
// [WatermarkResponseSingle]
type watermarkResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatermarkResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WatermarkResponseSingleError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    watermarkResponseSingleErrorJSON `json:"-"`
}

// watermarkResponseSingleErrorJSON contains the JSON metadata for the struct
// [WatermarkResponseSingleError]
type watermarkResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatermarkResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WatermarkResponseSingleMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    watermarkResponseSingleMessageJSON `json:"-"`
}

// watermarkResponseSingleMessageJSON contains the JSON metadata for the struct
// [WatermarkResponseSingleMessage]
type watermarkResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WatermarkResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WatermarkResponseSingleSuccess bool

const (
	WatermarkResponseSingleSuccessTrue WatermarkResponseSingleSuccess = true
)

type AccountStreamWatermarkDeleteResponse struct {
	Errors   []AccountStreamWatermarkDeleteResponseError   `json:"errors"`
	Messages []AccountStreamWatermarkDeleteResponseMessage `json:"messages"`
	Result   string                                        `json:"result"`
	// Whether the API call was successful
	Success AccountStreamWatermarkDeleteResponseSuccess `json:"success"`
	JSON    accountStreamWatermarkDeleteResponseJSON    `json:"-"`
}

// accountStreamWatermarkDeleteResponseJSON contains the JSON metadata for the
// struct [AccountStreamWatermarkDeleteResponse]
type accountStreamWatermarkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWatermarkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWatermarkDeleteResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountStreamWatermarkDeleteResponseErrorJSON `json:"-"`
}

// accountStreamWatermarkDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountStreamWatermarkDeleteResponseError]
type accountStreamWatermarkDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWatermarkDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWatermarkDeleteResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountStreamWatermarkDeleteResponseMessageJSON `json:"-"`
}

// accountStreamWatermarkDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountStreamWatermarkDeleteResponseMessage]
type accountStreamWatermarkDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWatermarkDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamWatermarkDeleteResponseSuccess bool

const (
	AccountStreamWatermarkDeleteResponseSuccessTrue AccountStreamWatermarkDeleteResponseSuccess = true
)

type AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadParams struct {
	// The image file to upload.
	File param.Field[string] `json:"file,required"`
}

func (r AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

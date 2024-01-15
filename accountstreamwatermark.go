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
func (r *AccountStreamWatermarkService) Get(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *AccountStreamWatermarkGetResponse, err error) {
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
func (r *AccountStreamWatermarkService) StreamWatermarkProfileNewWatermarkProfilesViaBasicUpload(ctx context.Context, accountIdentifier string, body AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadParams, opts ...option.RequestOption) (res *AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/watermarks", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Lists all watermark profiles for an account.
func (r *AccountStreamWatermarkService) StreamWatermarkProfileListWatermarkProfiles(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/stream/watermarks", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountStreamWatermarkGetResponse struct {
	Errors   []AccountStreamWatermarkGetResponseError   `json:"errors"`
	Messages []AccountStreamWatermarkGetResponseMessage `json:"messages"`
	Result   interface{}                                `json:"result"`
	// Whether the API call was successful
	Success AccountStreamWatermarkGetResponseSuccess `json:"success"`
	JSON    accountStreamWatermarkGetResponseJSON    `json:"-"`
}

// accountStreamWatermarkGetResponseJSON contains the JSON metadata for the struct
// [AccountStreamWatermarkGetResponse]
type accountStreamWatermarkGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWatermarkGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWatermarkGetResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountStreamWatermarkGetResponseErrorJSON `json:"-"`
}

// accountStreamWatermarkGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountStreamWatermarkGetResponseError]
type accountStreamWatermarkGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWatermarkGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWatermarkGetResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountStreamWatermarkGetResponseMessageJSON `json:"-"`
}

// accountStreamWatermarkGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountStreamWatermarkGetResponseMessage]
type accountStreamWatermarkGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWatermarkGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamWatermarkGetResponseSuccess bool

const (
	AccountStreamWatermarkGetResponseSuccessTrue AccountStreamWatermarkGetResponseSuccess = true
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

type AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponse struct {
	Errors   []AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseError   `json:"errors"`
	Messages []AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseMessage `json:"messages"`
	Result   interface{}                                                                                     `json:"result"`
	// Whether the API call was successful
	Success AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseSuccess `json:"success"`
	JSON    accountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseJSON    `json:"-"`
}

// accountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseJSON
// contains the JSON metadata for the struct
// [AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponse]
type accountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseError struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    accountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseErrorJSON `json:"-"`
}

// accountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseError]
type accountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseMessage struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    accountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseMessageJSON `json:"-"`
}

// accountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseMessage]
type accountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseSuccess bool

const (
	AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseSuccessTrue AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadResponseSuccess = true
)

type AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponse struct {
	Errors   []AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseError   `json:"errors"`
	Messages []AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseMessage `json:"messages"`
	Result   []AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseSuccess `json:"success"`
	JSON    accountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseJSON    `json:"-"`
}

// accountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseJSON
// contains the JSON metadata for the struct
// [AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponse]
type accountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseError struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    accountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseErrorJSON `json:"-"`
}

// accountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseError]
type accountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseMessage struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    accountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseMessageJSON `json:"-"`
}

// accountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseMessage]
type accountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseResult struct {
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
	Width int64                                                                               `json:"width"`
	JSON  accountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseResultJSON `json:"-"`
}

// accountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseResultJSON
// contains the JSON metadata for the struct
// [AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseResult]
type accountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseResultJSON struct {
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

func (r *AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseSuccess bool

const (
	AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseSuccessTrue AccountStreamWatermarkStreamWatermarkProfileListWatermarkProfilesResponseSuccess = true
)

type AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadParams struct {
	// The image file to upload.
	File param.Field[string] `json:"file,required"`
	// A short description of the watermark profile.
	Name param.Field[string] `json:"name"`
	// The translucency of the image. A value of `0.0` makes the image completely
	// transparent, and `1.0` makes the image completely opaque. Note that if the image
	// is already semi-transparent, setting this to `1.0` will not make the image
	// completely opaque.
	Opacity param.Field[float64] `json:"opacity"`
	// The whitespace between the adjacent edges (determined by position) of the video
	// and the image. `0.0` indicates no padding, and `1.0` indicates a fully padded
	// video width or length, as determined by the algorithm.
	Padding param.Field[float64] `json:"padding"`
	// The location of the image. Valid positions are: `upperRight`, `upperLeft`,
	// `lowerLeft`, `lowerRight`, and `center`. Note that `center` ignores the
	// `padding` parameter.
	Position param.Field[string] `json:"position"`
	// The size of the image relative to the overall size of the video. This parameter
	// will adapt to horizontal and vertical videos automatically. `0.0` indicates no
	// scaling (use the size of the image as-is), and `1.0 `fills the entire video.
	Scale param.Field[float64] `json:"scale"`
}

func (r AccountStreamWatermarkStreamWatermarkProfileNewWatermarkProfilesViaBasicUploadParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

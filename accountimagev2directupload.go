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

// AccountImageV2DirectUploadService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountImageV2DirectUploadService] method instead.
type AccountImageV2DirectUploadService struct {
	Options []option.RequestOption
}

// NewAccountImageV2DirectUploadService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountImageV2DirectUploadService(opts ...option.RequestOption) (r *AccountImageV2DirectUploadService) {
	r = &AccountImageV2DirectUploadService{}
	r.Options = opts
	return
}

// Direct uploads allow users to upload images without API keys. A common use case
// are web apps, client-side applications, or mobile devices where users upload
// content directly to Cloudflare Images. This method creates a draft record for a
// future image. It returns an upload URL and an image identifier. To verify if the
// image itself has been uploaded, send an image details request
// (accounts/:account_identifier/images/v1/:identifier), and check that the
// `draft: true` property is not present.
func (r *AccountImageV2DirectUploadService) CloudflareImagesNewAuthenticatedDirectUploadURLV2(ctx context.Context, accountIdentifier string, body AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadURLV2Params, opts ...option.RequestOption) (res *AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2Response, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v2/direct_upload", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2Response struct {
	Errors   []AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseError   `json:"errors"`
	Messages []AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseMessage `json:"messages"`
	Result   AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseSuccess `json:"success"`
	JSON    accountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseJSON    `json:"-"`
}

// accountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseJSON
// contains the JSON metadata for the struct
// [AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2Response]
type accountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseError struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    accountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseErrorJSON `json:"-"`
}

// accountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseError]
type accountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseMessage struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    accountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseMessageJSON `json:"-"`
}

// accountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseMessage]
type accountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseResult struct {
	// Image unique identifier.
	ID string `json:"id"`
	// The URL the unauthenticated upload can be performed to using a single HTTP POST
	// (multipart/form-data) request.
	UploadURL string                                                                                        `json:"uploadURL"`
	JSON      accountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseResultJSON `json:"-"`
}

// accountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseResultJSON
// contains the JSON metadata for the struct
// [AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseResult]
type accountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseResultJSON struct {
	ID          apijson.Field
	UploadURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseSuccess bool

const (
	AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseSuccessTrue AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseSuccess = true
)

type AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadURLV2Params struct {
	// The date after which the upload will not be accepted. Minimum: Now + 2 minutes.
	// Maximum: Now + 6 hours.
	Expiry param.Field[time.Time] `json:"expiry" format:"date-time"`
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record, for managing images.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Indicates whether the image requires a signature token to be accessed.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
}

func (r AccountImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadURLV2Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

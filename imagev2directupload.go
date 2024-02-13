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

// ImageV2DirectUploadService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewImageV2DirectUploadService]
// method instead.
type ImageV2DirectUploadService struct {
	Options []option.RequestOption
}

// NewImageV2DirectUploadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewImageV2DirectUploadService(opts ...option.RequestOption) (r *ImageV2DirectUploadService) {
	r = &ImageV2DirectUploadService{}
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
func (r *ImageV2DirectUploadService) CloudflareImagesNewAuthenticatedDirectUploadURLV2(ctx context.Context, accountID string, body ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadURLV2Params, opts ...option.RequestOption) (res *ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2Response, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v2/direct_upload", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2Response struct {
	// Image unique identifier.
	ID string `json:"id"`
	// The URL the unauthenticated upload can be performed to using a single HTTP POST
	// (multipart/form-data) request.
	UploadURL string                                                                           `json:"uploadURL"`
	JSON      imageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseJSON `json:"-"`
}

// imageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseJSON
// contains the JSON metadata for the struct
// [ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2Response]
type imageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseJSON struct {
	ID          apijson.Field
	UploadURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2Response) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadURLV2Params struct {
	// The date after which the upload will not be accepted. Minimum: Now + 2 minutes.
	// Maximum: Now + 6 hours.
	Expiry param.Field[time.Time] `json:"expiry" format:"date-time"`
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record, for managing images.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Indicates whether the image requires a signature token to be accessed.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
}

func (r ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadURLV2Params) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelope struct {
	Errors   []ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2Response                   `json:"result,required"`
	// Whether the API call was successful
	Success ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeJSON    `json:"-"`
}

// imageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelope]
type imageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeErrors struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    imageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeErrors]
type imageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeMessages struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    imageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeMessages]
type imageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeSuccess bool

const (
	ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeSuccessTrue ImageV2DirectUploadCloudflareImagesNewAuthenticatedDirectUploadUrlv2ResponseEnvelopeSuccess = true
)

// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *ImageV2DirectUploadService) New(ctx context.Context, params ImageV2DirectUploadNewParams, opts ...option.RequestOption) (res *ImageV2DirectUploadNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV2DirectUploadNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v2/direct_upload", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ImageV2DirectUploadNewResponse struct {
	// Image unique identifier.
	ID string `json:"id"`
	// The URL the unauthenticated upload can be performed to using a single HTTP POST
	// (multipart/form-data) request.
	UploadURL string                             `json:"uploadURL"`
	JSON      imageV2DirectUploadNewResponseJSON `json:"-"`
}

// imageV2DirectUploadNewResponseJSON contains the JSON metadata for the struct
// [ImageV2DirectUploadNewResponse]
type imageV2DirectUploadNewResponseJSON struct {
	ID          apijson.Field
	UploadURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV2DirectUploadNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV2DirectUploadNewResponseJSON) RawJSON() string {
	return r.raw
}

type ImageV2DirectUploadNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The date after which the upload will not be accepted. Minimum: Now + 2 minutes.
	// Maximum: Now + 6 hours.
	Expiry param.Field[time.Time] `json:"expiry" format:"date-time"`
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record, for managing images.
	Metadata param.Field[interface{}] `json:"metadata"`
	// Indicates whether the image requires a signature token to be accessed.
	RequireSignedURLs param.Field[bool] `json:"requireSignedURLs"`
}

func (r ImageV2DirectUploadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ImageV2DirectUploadNewResponseEnvelope struct {
	Errors   []ImageV2DirectUploadNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV2DirectUploadNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageV2DirectUploadNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ImageV2DirectUploadNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV2DirectUploadNewResponseEnvelopeJSON    `json:"-"`
}

// imageV2DirectUploadNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ImageV2DirectUploadNewResponseEnvelope]
type imageV2DirectUploadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV2DirectUploadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV2DirectUploadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ImageV2DirectUploadNewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    imageV2DirectUploadNewResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV2DirectUploadNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ImageV2DirectUploadNewResponseEnvelopeErrors]
type imageV2DirectUploadNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV2DirectUploadNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV2DirectUploadNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ImageV2DirectUploadNewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    imageV2DirectUploadNewResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV2DirectUploadNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ImageV2DirectUploadNewResponseEnvelopeMessages]
type imageV2DirectUploadNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV2DirectUploadNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV2DirectUploadNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ImageV2DirectUploadNewResponseEnvelopeSuccess bool

const (
	ImageV2DirectUploadNewResponseEnvelopeSuccessTrue ImageV2DirectUploadNewResponseEnvelopeSuccess = true
)

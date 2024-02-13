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

// ImageV1KeyService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewImageV1KeyService] method instead.
type ImageV1KeyService struct {
	Options []option.RequestOption
}

// NewImageV1KeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewImageV1KeyService(opts ...option.RequestOption) (r *ImageV1KeyService) {
	r = &ImageV1KeyService{}
	r.Options = opts
	return
}

// Lists your signing keys. These can be found on your Cloudflare Images dashboard.
func (r *ImageV1KeyService) CloudflareImagesKeysListSigningKeys(ctx context.Context, accountID string, opts ...option.RequestOption) (res *ImageV1KeyCloudflareImagesKeysListSigningKeysResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/keys", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ImageV1KeyCloudflareImagesKeysListSigningKeysResponse struct {
	Keys []ImageV1KeyCloudflareImagesKeysListSigningKeysResponseKey `json:"keys"`
	JSON imageV1KeyCloudflareImagesKeysListSigningKeysResponseJSON  `json:"-"`
}

// imageV1KeyCloudflareImagesKeysListSigningKeysResponseJSON contains the JSON
// metadata for the struct [ImageV1KeyCloudflareImagesKeysListSigningKeysResponse]
type imageV1KeyCloudflareImagesKeysListSigningKeysResponseJSON struct {
	Keys        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1KeyCloudflareImagesKeysListSigningKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1KeyCloudflareImagesKeysListSigningKeysResponseKey struct {
	// Key name.
	Name string `json:"name"`
	// Key value.
	Value string                                                       `json:"value"`
	JSON  imageV1KeyCloudflareImagesKeysListSigningKeysResponseKeyJSON `json:"-"`
}

// imageV1KeyCloudflareImagesKeysListSigningKeysResponseKeyJSON contains the JSON
// metadata for the struct
// [ImageV1KeyCloudflareImagesKeysListSigningKeysResponseKey]
type imageV1KeyCloudflareImagesKeysListSigningKeysResponseKeyJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1KeyCloudflareImagesKeysListSigningKeysResponseKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelope struct {
	Errors   []ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageV1KeyCloudflareImagesKeysListSigningKeysResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeJSON    `json:"-"`
}

// imageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelope]
type imageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeErrors struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    imageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeErrors]
type imageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeMessages struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    imageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeMessages]
type imageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeSuccess bool

const (
	ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeSuccessTrue ImageV1KeyCloudflareImagesKeysListSigningKeysResponseEnvelopeSuccess = true
)

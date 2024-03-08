// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *ImageV1KeyService) List(ctx context.Context, query ImageV1KeyListParams, opts ...option.RequestOption) (res *ImageV1KeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ImageV1KeyListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v1/keys", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ImageV1KeyListResponse struct {
	Keys []ImageV1KeyListResponseKey `json:"keys"`
	JSON imageV1KeyListResponseJSON  `json:"-"`
}

// imageV1KeyListResponseJSON contains the JSON metadata for the struct
// [ImageV1KeyListResponse]
type imageV1KeyListResponseJSON struct {
	Keys        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1KeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1KeyListResponseJSON) RawJSON() string {
	return r.raw
}

type ImageV1KeyListResponseKey struct {
	// Key name.
	Name string `json:"name"`
	// Key value.
	Value string                        `json:"value"`
	JSON  imageV1KeyListResponseKeyJSON `json:"-"`
}

// imageV1KeyListResponseKeyJSON contains the JSON metadata for the struct
// [ImageV1KeyListResponseKey]
type imageV1KeyListResponseKeyJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1KeyListResponseKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1KeyListResponseKeyJSON) RawJSON() string {
	return r.raw
}

type ImageV1KeyListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ImageV1KeyListResponseEnvelope struct {
	Errors   []ImageV1KeyListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ImageV1KeyListResponseEnvelopeMessages `json:"messages,required"`
	Result   ImageV1KeyListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ImageV1KeyListResponseEnvelopeSuccess `json:"success,required"`
	JSON    imageV1KeyListResponseEnvelopeJSON    `json:"-"`
}

// imageV1KeyListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ImageV1KeyListResponseEnvelope]
type imageV1KeyListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1KeyListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1KeyListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ImageV1KeyListResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    imageV1KeyListResponseEnvelopeErrorsJSON `json:"-"`
}

// imageV1KeyListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ImageV1KeyListResponseEnvelopeErrors]
type imageV1KeyListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1KeyListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1KeyListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ImageV1KeyListResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    imageV1KeyListResponseEnvelopeMessagesJSON `json:"-"`
}

// imageV1KeyListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ImageV1KeyListResponseEnvelopeMessages]
type imageV1KeyListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageV1KeyListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imageV1KeyListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ImageV1KeyListResponseEnvelopeSuccess bool

const (
	ImageV1KeyListResponseEnvelopeSuccessTrue ImageV1KeyListResponseEnvelopeSuccess = true
)

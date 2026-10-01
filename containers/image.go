// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package containers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ImageService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewImageService] method instead.
type ImageService struct {
	Options []option.RequestOption
}

// NewImageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewImageService(opts ...option.RequestOption) (r *ImageService) {
	r = &ImageService{}
	r.Options = opts
	return
}

// Idempotently starts or observes preparation of the runtime artifacts required to
// run one digest-pinned managed container image on Cloudflare's network. Returns
// 202 while durable preparation continues and 200 when the image is ready or
// preparation has reached a terminal error.
func (r *ImageService) Prepare(ctx context.Context, params ImagePrepareParams, opts ...option.RequestOption) (res *ImagePrepareResponse, err error) {
	var env ImagePrepareResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/containers/image-preparations", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Durable preparation state for a container image.
type ImagePrepareResponse struct {
	// Image url.
	Image string `json:"image" api:"required"`
	// Current durable preparation state for a container image.
	Status ImagePrepareResponseStatus `json:"status" api:"required"`
	// Digest of the prepared runtime artifact when status is ready.
	ArtifactDigest string `json:"artifact_digest"`
	// Human-readable pending or terminal error detail.
	Reason string                   `json:"reason"`
	JSON   imagePrepareResponseJSON `json:"-"`
}

// imagePrepareResponseJSON contains the JSON metadata for the struct
// [ImagePrepareResponse]
type imagePrepareResponseJSON struct {
	Image          apijson.Field
	Status         apijson.Field
	ArtifactDigest apijson.Field
	Reason         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ImagePrepareResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imagePrepareResponseJSON) RawJSON() string {
	return r.raw
}

// Current durable preparation state for a container image.
type ImagePrepareResponseStatus string

const (
	ImagePrepareResponseStatusPending ImagePrepareResponseStatus = "pending"
	ImagePrepareResponseStatusReady   ImagePrepareResponseStatus = "ready"
	ImagePrepareResponseStatusError   ImagePrepareResponseStatus = "error"
)

func (r ImagePrepareResponseStatus) IsKnown() bool {
	switch r {
	case ImagePrepareResponseStatusPending, ImagePrepareResponseStatusReady, ImagePrepareResponseStatusError:
		return true
	}
	return false
}

type ImagePrepareParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Image url.
	Image param.Field[string] `json:"image" api:"required"`
}

func (r ImagePrepareParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ImagePrepareResponseEnvelope struct {
	Errors   []ImagePrepareResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ImagePrepareResponseEnvelopeMessages `json:"messages" api:"required"`
	// Durable preparation state for a container image.
	Result ImagePrepareResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                             `json:"success" api:"required"`
	JSON    imagePrepareResponseEnvelopeJSON `json:"-"`
}

// imagePrepareResponseEnvelopeJSON contains the JSON metadata for the struct
// [ImagePrepareResponseEnvelope]
type imagePrepareResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImagePrepareResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imagePrepareResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ImagePrepareResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code" api:"required"`
	Message          string                                   `json:"message" api:"required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           ImagePrepareResponseEnvelopeErrorsSource `json:"source"`
	JSON             imagePrepareResponseEnvelopeErrorsJSON   `json:"-"`
}

// imagePrepareResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ImagePrepareResponseEnvelopeErrors]
type imagePrepareResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ImagePrepareResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imagePrepareResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ImagePrepareResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    imagePrepareResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// imagePrepareResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ImagePrepareResponseEnvelopeErrorsSource]
type imagePrepareResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImagePrepareResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imagePrepareResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ImagePrepareResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ImagePrepareResponseEnvelopeMessagesSource `json:"source"`
	JSON             imagePrepareResponseEnvelopeMessagesJSON   `json:"-"`
}

// imagePrepareResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ImagePrepareResponseEnvelopeMessages]
type imagePrepareResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ImagePrepareResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imagePrepareResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ImagePrepareResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    imagePrepareResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// imagePrepareResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ImagePrepareResponseEnvelopeMessagesSource]
type imagePrepareResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImagePrepareResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r imagePrepareResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

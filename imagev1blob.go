// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ImageV1BlobService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewImageV1BlobService] method
// instead.
type ImageV1BlobService struct {
	Options []option.RequestOption
}

// NewImageV1BlobService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewImageV1BlobService(opts ...option.RequestOption) (r *ImageV1BlobService) {
	r = &ImageV1BlobService{}
	r.Options = opts
	return
}

// Fetch base image. For most images this will be the originally uploaded file. For
// larger images it can be a near-lossless version of the original.
func (r *ImageV1BlobService) Get(ctx context.Context, accountID string, imageID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "image/*")}, opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/%s/blob", accountID, imageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

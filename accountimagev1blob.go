// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AccountImageV1BlobService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountImageV1BlobService] method
// instead.
type AccountImageV1BlobService struct {
	Options []option.RequestOption
}

// NewAccountImageV1BlobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountImageV1BlobService(opts ...option.RequestOption) (r *AccountImageV1BlobService) {
	r = &AccountImageV1BlobService{}
	r.Options = opts
	return
}

// Fetch base image. For most images this will be the originally uploaded file. For
// larger images it can be a near-lossless version of the original.
func (r *AccountImageV1BlobService) CloudflareImagesBaseImage(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *ImageResponseBlob, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/%s/blob", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Union satisfied by [shared.UnionString] or [ImageResponseBlobObject].
type ImageResponseBlob interface {
	ImplementsImageResponseBlob()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ImageResponseBlob)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

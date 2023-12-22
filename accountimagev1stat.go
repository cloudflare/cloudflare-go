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

// AccountImageV1StatService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountImageV1StatService] method
// instead.
type AccountImageV1StatService struct {
	Options []option.RequestOption
}

// NewAccountImageV1StatService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountImageV1StatService(opts ...option.RequestOption) (r *AccountImageV1StatService) {
	r = &AccountImageV1StatService{}
	r.Options = opts
	return
}

// Fetch usage statistics details for Cloudflare Images.
func (r *AccountImageV1StatService) CloudflareImagesImagesUsageStatistics(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *ImagesStat, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/stats", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ImagesStat struct {
	Count ImagesStatCount `json:"count"`
	JSON  imagesStatJSON  `json:"-"`
}

// imagesStatJSON contains the JSON metadata for the struct [ImagesStat]
type imagesStatJSON struct {
	Count       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImagesStat) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImagesStatCount struct {
	// Cloudflare Images allowed usage.
	Allowed float64 `json:"allowed"`
	// Cloudflare Images current usage.
	Current float64             `json:"current"`
	JSON    imagesStatCountJSON `json:"-"`
}

// imagesStatCountJSON contains the JSON metadata for the struct [ImagesStatCount]
type imagesStatCountJSON struct {
	Allowed     apijson.Field
	Current     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImagesStatCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

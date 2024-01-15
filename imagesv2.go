// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ImagesV2Service contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewImagesV2Service] method instead.
type ImagesV2Service struct {
	Options []option.RequestOption
}

// NewImagesV2Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewImagesV2Service(opts ...option.RequestOption) (r *ImagesV2Service) {
	r = &ImagesV2Service{}
	r.Options = opts
	return
}

// List up to 10000 images with one request. Use the optional parameters below to
// get a specific range of images. Endpoint returns continuation_token if more
// images are present.
func (r *ImagesV2Service) List(ctx context.Context, accountIdentifier string, query ImagesV2ListParams, opts ...option.RequestOption) (res *ImagesImagesListResponseV2, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v2", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ImagesImagesListResponseV2 struct {
	Errors   []ImagesImagesListResponseV2Error   `json:"errors"`
	Messages []ImagesImagesListResponseV2Message `json:"messages"`
	Result   ImagesImagesListResponseV2Result    `json:"result"`
	// Whether the API call was successful
	Success ImagesImagesListResponseV2Success `json:"success"`
	JSON    imagesImagesListResponseV2JSON    `json:"-"`
}

// imagesImagesListResponseV2JSON contains the JSON metadata for the struct
// [ImagesImagesListResponseV2]
type imagesImagesListResponseV2JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImagesImagesListResponseV2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImagesImagesListResponseV2Error struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    imagesImagesListResponseV2ErrorJSON `json:"-"`
}

// imagesImagesListResponseV2ErrorJSON contains the JSON metadata for the struct
// [ImagesImagesListResponseV2Error]
type imagesImagesListResponseV2ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImagesImagesListResponseV2Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImagesImagesListResponseV2Message struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    imagesImagesListResponseV2MessageJSON `json:"-"`
}

// imagesImagesListResponseV2MessageJSON contains the JSON metadata for the struct
// [ImagesImagesListResponseV2Message]
type imagesImagesListResponseV2MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImagesImagesListResponseV2Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImagesImagesListResponseV2Result struct {
	Images []ImagesImagesListResponseV2ResultImage `json:"images"`
	JSON   imagesImagesListResponseV2ResultJSON    `json:"-"`
}

// imagesImagesListResponseV2ResultJSON contains the JSON metadata for the struct
// [ImagesImagesListResponseV2Result]
type imagesImagesListResponseV2ResultJSON struct {
	Images      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImagesImagesListResponseV2Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImagesImagesListResponseV2ResultImage struct {
	// Image unique identifier.
	ID string `json:"id"`
	// Image file name.
	Filename string `json:"filename"`
	// User modifiable key-value store. Can be used for keeping references to another
	// system of record for managing images. Metadata must not exceed 1024 bytes.
	Meta interface{} `json:"meta"`
	// Indicates whether the image can be a accessed only using it's UID. If set to
	// true, a signed token needs to be generated with a signing key to view the image.
	RequireSignedURLs bool `json:"requireSignedURLs"`
	// When the media item was uploaded.
	Uploaded time.Time `json:"uploaded" format:"date-time"`
	// Object specifying available variants for an image.
	Variants []ImagesImagesListResponseV2ResultImagesVariant `json:"variants" format:"uri"`
	JSON     imagesImagesListResponseV2ResultImageJSON       `json:"-"`
}

// imagesImagesListResponseV2ResultImageJSON contains the JSON metadata for the
// struct [ImagesImagesListResponseV2ResultImage]
type imagesImagesListResponseV2ResultImageJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ImagesImagesListResponseV2ResultImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// URI to thumbnail variant for an image.
//
// Union satisfied by [shared.UnionString], [shared.UnionString] or
// [shared.UnionString].
type ImagesImagesListResponseV2ResultImagesVariant interface {
	ImplementsImagesImagesListResponseV2ResultImagesVariant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ImagesImagesListResponseV2ResultImagesVariant)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful
type ImagesImagesListResponseV2Success bool

const (
	ImagesImagesListResponseV2SuccessTrue ImagesImagesListResponseV2Success = true
)

type ImagesV2ListParams struct {
	// Continuation token for a next page. List images V2 returns continuation_token
	ContinuationToken param.Field[string] `query:"continuation_token"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Sorting order by upload time.
	SortOrder param.Field[ImagesV2ListParamsSortOrder] `query:"sort_order"`
}

// URLQuery serializes [ImagesV2ListParams]'s query parameters as `url.Values`.
func (r ImagesV2ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sorting order by upload time.
type ImagesV2ListParamsSortOrder string

const (
	ImagesV2ListParamsSortOrderAsc  ImagesV2ListParamsSortOrder = "asc"
	ImagesV2ListParamsSortOrderDesc ImagesV2ListParamsSortOrder = "desc"
)

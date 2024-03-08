// File generated from our OpenAPI spec by Stainless.

package images

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// V2Service contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewV2Service] method instead.
type V2Service struct {
	Options       []option.RequestOption
	DirectUploads *V2DirectUploadService
}

// NewV2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV2Service(opts ...option.RequestOption) (r *V2Service) {
	r = &V2Service{}
	r.Options = opts
	r.DirectUploads = NewV2DirectUploadService(opts...)
	return
}

// List up to 10000 images with one request. Use the optional parameters below to
// get a specific range of images. Endpoint returns continuation_token if more
// images are present.
func (r *V2Service) List(ctx context.Context, params V2ListParams, opts ...option.RequestOption) (res *V2ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V2ListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/images/v2", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type V2ListResponse struct {
	// Continuation token to fetch next page. Passed as a query param when requesting
	// List V2 api endpoint.
	ContinuationToken string                `json:"continuation_token,nullable"`
	Images            []V2ListResponseImage `json:"images"`
	JSON              v2ListResponseJSON    `json:"-"`
}

// v2ListResponseJSON contains the JSON metadata for the struct [V2ListResponse]
type v2ListResponseJSON struct {
	ContinuationToken apijson.Field
	Images            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *V2ListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ListResponseJSON) RawJSON() string {
	return r.raw
}

type V2ListResponseImage struct {
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
	Variants []V2ListResponseImagesVariant `json:"variants" format:"uri"`
	JSON     v2ListResponseImageJSON       `json:"-"`
}

// v2ListResponseImageJSON contains the JSON metadata for the struct
// [V2ListResponseImage]
type v2ListResponseImageJSON struct {
	ID                apijson.Field
	Filename          apijson.Field
	Meta              apijson.Field
	RequireSignedURLs apijson.Field
	Uploaded          apijson.Field
	Variants          apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *V2ListResponseImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ListResponseImageJSON) RawJSON() string {
	return r.raw
}

// URI to thumbnail variant for an image.
//
// Union satisfied by [shared.UnionString], [shared.UnionString] or
// [shared.UnionString].
type V2ListResponseImagesVariant interface {
	ImplementsImagesV2ListResponseImagesVariant()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V2ListResponseImagesVariant)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type V2ListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Continuation token for a next page. List images V2 returns continuation_token
	ContinuationToken param.Field[string] `query:"continuation_token"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Sorting order by upload time.
	SortOrder param.Field[V2ListParamsSortOrder] `query:"sort_order"`
}

// URLQuery serializes [V2ListParams]'s query parameters as `url.Values`.
func (r V2ListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sorting order by upload time.
type V2ListParamsSortOrder string

const (
	V2ListParamsSortOrderAsc  V2ListParamsSortOrder = "asc"
	V2ListParamsSortOrderDesc V2ListParamsSortOrder = "desc"
)

type V2ListResponseEnvelope struct {
	Errors   []V2ListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V2ListResponseEnvelopeMessages `json:"messages,required"`
	Result   V2ListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V2ListResponseEnvelopeSuccess `json:"success,required"`
	JSON    v2ListResponseEnvelopeJSON    `json:"-"`
}

// v2ListResponseEnvelopeJSON contains the JSON metadata for the struct
// [V2ListResponseEnvelope]
type v2ListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V2ListResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    v2ListResponseEnvelopeErrorsJSON `json:"-"`
}

// v2ListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V2ListResponseEnvelopeErrors]
type v2ListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V2ListResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    v2ListResponseEnvelopeMessagesJSON `json:"-"`
}

// v2ListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [V2ListResponseEnvelopeMessages]
type v2ListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V2ListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v2ListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V2ListResponseEnvelopeSuccess bool

const (
	V2ListResponseEnvelopeSuccessTrue V2ListResponseEnvelopeSuccess = true
)

// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZoneSnippetService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneSnippetService] method
// instead.
type ZoneSnippetService struct {
	Options []option.RequestOption
}

// NewZoneSnippetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneSnippetService(opts ...option.RequestOption) (r *ZoneSnippetService) {
	r = &ZoneSnippetService{}
	r.Options = opts
	return
}

// Snippet
func (r *ZoneSnippetService) Get(ctx context.Context, zoneIdentifier string, snippetName string, opts ...option.RequestOption) (res *ZoneSnippetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Put Snippet
func (r *ZoneSnippetService) Update(ctx context.Context, zoneIdentifier string, snippetName string, body ZoneSnippetUpdateParams, opts ...option.RequestOption) (res *ZoneSnippetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// All Snippets
func (r *ZoneSnippetService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSnippetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/snippets", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Snippet
func (r *ZoneSnippetService) Delete(ctx context.Context, zoneIdentifier string, snippetName string, opts ...option.RequestOption) (res *ZoneSnippetDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Snippet Content
func (r *ZoneSnippetService) Content(ctx context.Context, zoneIdentifier string, snippetName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "multipart/form-data")}, opts...)
	path := fmt.Sprintf("zones/%s/snippets/%s/content", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSnippetGetResponse struct {
	Errors   []ZoneSnippetGetResponseError   `json:"errors"`
	Messages []ZoneSnippetGetResponseMessage `json:"messages"`
	// Snippet Information
	Result ZoneSnippetGetResponseResult `json:"result"`
	// Whether the API call was successful
	Success ZoneSnippetGetResponseSuccess `json:"success"`
	JSON    zoneSnippetGetResponseJSON    `json:"-"`
}

// zoneSnippetGetResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetGetResponse]
type zoneSnippetGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSnippetGetResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    zoneSnippetGetResponseErrorJSON `json:"-"`
}

// zoneSnippetGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSnippetGetResponseError]
type zoneSnippetGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSnippetGetResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    zoneSnippetGetResponseMessageJSON `json:"-"`
}

// zoneSnippetGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSnippetGetResponseMessage]
type zoneSnippetGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Snippet Information
type ZoneSnippetGetResponseResult struct {
	// Creation time of the snippet
	CreatedOn string `json:"created_on"`
	// Modification time of the snippet
	ModifiedOn string `json:"modified_on"`
	// Snippet identifying name
	SnippetName string                           `json:"snippet_name"`
	JSON        zoneSnippetGetResponseResultJSON `json:"-"`
}

// zoneSnippetGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneSnippetGetResponseResult]
type zoneSnippetGetResponseResultJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSnippetGetResponseSuccess bool

const (
	ZoneSnippetGetResponseSuccessTrue ZoneSnippetGetResponseSuccess = true
)

type ZoneSnippetUpdateResponse struct {
	Errors   []ZoneSnippetUpdateResponseError   `json:"errors"`
	Messages []ZoneSnippetUpdateResponseMessage `json:"messages"`
	// Snippet Information
	Result ZoneSnippetUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success ZoneSnippetUpdateResponseSuccess `json:"success"`
	JSON    zoneSnippetUpdateResponseJSON    `json:"-"`
}

// zoneSnippetUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetUpdateResponse]
type zoneSnippetUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSnippetUpdateResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    zoneSnippetUpdateResponseErrorJSON `json:"-"`
}

// zoneSnippetUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSnippetUpdateResponseError]
type zoneSnippetUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSnippetUpdateResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneSnippetUpdateResponseMessageJSON `json:"-"`
}

// zoneSnippetUpdateResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSnippetUpdateResponseMessage]
type zoneSnippetUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Snippet Information
type ZoneSnippetUpdateResponseResult struct {
	// Creation time of the snippet
	CreatedOn string `json:"created_on"`
	// Modification time of the snippet
	ModifiedOn string `json:"modified_on"`
	// Snippet identifying name
	SnippetName string                              `json:"snippet_name"`
	JSON        zoneSnippetUpdateResponseResultJSON `json:"-"`
}

// zoneSnippetUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZoneSnippetUpdateResponseResult]
type zoneSnippetUpdateResponseResultJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSnippetUpdateResponseSuccess bool

const (
	ZoneSnippetUpdateResponseSuccessTrue ZoneSnippetUpdateResponseSuccess = true
)

type ZoneSnippetListResponse struct {
	Errors   []ZoneSnippetListResponseError   `json:"errors"`
	Messages []ZoneSnippetListResponseMessage `json:"messages"`
	// List of all zone snippets
	Result []ZoneSnippetListResponseResult `json:"result"`
	// Whether the API call was successful
	Success ZoneSnippetListResponseSuccess `json:"success"`
	JSON    zoneSnippetListResponseJSON    `json:"-"`
}

// zoneSnippetListResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetListResponse]
type zoneSnippetListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSnippetListResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    zoneSnippetListResponseErrorJSON `json:"-"`
}

// zoneSnippetListResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSnippetListResponseError]
type zoneSnippetListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSnippetListResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    zoneSnippetListResponseMessageJSON `json:"-"`
}

// zoneSnippetListResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSnippetListResponseMessage]
type zoneSnippetListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Snippet Information
type ZoneSnippetListResponseResult struct {
	// Creation time of the snippet
	CreatedOn string `json:"created_on"`
	// Modification time of the snippet
	ModifiedOn string `json:"modified_on"`
	// Snippet identifying name
	SnippetName string                            `json:"snippet_name"`
	JSON        zoneSnippetListResponseResultJSON `json:"-"`
}

// zoneSnippetListResponseResultJSON contains the JSON metadata for the struct
// [ZoneSnippetListResponseResult]
type zoneSnippetListResponseResultJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSnippetListResponseSuccess bool

const (
	ZoneSnippetListResponseSuccessTrue ZoneSnippetListResponseSuccess = true
)

type ZoneSnippetDeleteResponse struct {
	Errors   []ZoneSnippetDeleteResponseError   `json:"errors,required"`
	Messages []ZoneSnippetDeleteResponseMessage `json:"messages,required"`
	Result   ZoneSnippetDeleteResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success ZoneSnippetDeleteResponseSuccess `json:"success,required"`
	JSON    zoneSnippetDeleteResponseJSON    `json:"-"`
}

// zoneSnippetDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneSnippetDeleteResponse]
type zoneSnippetDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSnippetDeleteResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    zoneSnippetDeleteResponseErrorJSON `json:"-"`
}

// zoneSnippetDeleteResponseErrorJSON contains the JSON metadata for the struct
// [ZoneSnippetDeleteResponseError]
type zoneSnippetDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSnippetDeleteResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneSnippetDeleteResponseMessageJSON `json:"-"`
}

// zoneSnippetDeleteResponseMessageJSON contains the JSON metadata for the struct
// [ZoneSnippetDeleteResponseMessage]
type zoneSnippetDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSnippetDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZoneSnippetDeleteResponseResultUnknown],
// [ZoneSnippetDeleteResponseResultArray] or [shared.UnionString].
type ZoneSnippetDeleteResponseResult interface {
	ImplementsZoneSnippetDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneSnippetDeleteResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZoneSnippetDeleteResponseResultArray []interface{}

func (r ZoneSnippetDeleteResponseResultArray) ImplementsZoneSnippetDeleteResponseResult() {}

// Whether the API call was successful
type ZoneSnippetDeleteResponseSuccess bool

const (
	ZoneSnippetDeleteResponseSuccessTrue ZoneSnippetDeleteResponseSuccess = true
)

type ZoneSnippetUpdateParams struct {
	// Content files of uploaded snippet
	Files    param.Field[string]                          `json:"files"`
	Metadata param.Field[ZoneSnippetUpdateParamsMetadata] `json:"metadata"`
}

func (r ZoneSnippetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneSnippetUpdateParamsMetadata struct {
	// Main module name of uploaded snippet
	MainModule param.Field[string] `json:"main_module"`
}

func (r ZoneSnippetUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

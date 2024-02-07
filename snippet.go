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

// SnippetService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewSnippetService] method instead.
type SnippetService struct {
	Options      []option.RequestOption
	Content      *SnippetContentService
	SnippetRules *SnippetSnippetRuleService
}

// NewSnippetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSnippetService(opts ...option.RequestOption) (r *SnippetService) {
	r = &SnippetService{}
	r.Options = opts
	r.Content = NewSnippetContentService(opts...)
	r.SnippetRules = NewSnippetSnippetRuleService(opts...)
	return
}

// Snippet
func (r *SnippetService) Get(ctx context.Context, zoneIdentifier string, snippetName string, opts ...option.RequestOption) (res *SnippetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Put Snippet
func (r *SnippetService) Update(ctx context.Context, zoneIdentifier string, snippetName string, body SnippetUpdateParams, opts ...option.RequestOption) (res *SnippetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// All Snippets
func (r *SnippetService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *SnippetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/snippets", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Snippet
func (r *SnippetService) Delete(ctx context.Context, zoneIdentifier string, snippetName string, opts ...option.RequestOption) (res *SnippetDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type SnippetGetResponse struct {
	Errors   []SnippetGetResponseError   `json:"errors"`
	Messages []SnippetGetResponseMessage `json:"messages"`
	// Snippet Information
	Result SnippetGetResponseResult `json:"result"`
	// Whether the API call was successful
	Success SnippetGetResponseSuccess `json:"success"`
	JSON    snippetGetResponseJSON    `json:"-"`
}

// snippetGetResponseJSON contains the JSON metadata for the struct
// [SnippetGetResponse]
type snippetGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetGetResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    snippetGetResponseErrorJSON `json:"-"`
}

// snippetGetResponseErrorJSON contains the JSON metadata for the struct
// [SnippetGetResponseError]
type snippetGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetGetResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    snippetGetResponseMessageJSON `json:"-"`
}

// snippetGetResponseMessageJSON contains the JSON metadata for the struct
// [SnippetGetResponseMessage]
type snippetGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Snippet Information
type SnippetGetResponseResult struct {
	// Creation time of the snippet
	CreatedOn string `json:"created_on"`
	// Modification time of the snippet
	ModifiedOn string `json:"modified_on"`
	// Snippet identifying name
	SnippetName string                       `json:"snippet_name"`
	JSON        snippetGetResponseResultJSON `json:"-"`
}

// snippetGetResponseResultJSON contains the JSON metadata for the struct
// [SnippetGetResponseResult]
type snippetGetResponseResultJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SnippetGetResponseSuccess bool

const (
	SnippetGetResponseSuccessTrue SnippetGetResponseSuccess = true
)

type SnippetUpdateResponse struct {
	Errors   []SnippetUpdateResponseError   `json:"errors"`
	Messages []SnippetUpdateResponseMessage `json:"messages"`
	// Snippet Information
	Result SnippetUpdateResponseResult `json:"result"`
	// Whether the API call was successful
	Success SnippetUpdateResponseSuccess `json:"success"`
	JSON    snippetUpdateResponseJSON    `json:"-"`
}

// snippetUpdateResponseJSON contains the JSON metadata for the struct
// [SnippetUpdateResponse]
type snippetUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetUpdateResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    snippetUpdateResponseErrorJSON `json:"-"`
}

// snippetUpdateResponseErrorJSON contains the JSON metadata for the struct
// [SnippetUpdateResponseError]
type snippetUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetUpdateResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    snippetUpdateResponseMessageJSON `json:"-"`
}

// snippetUpdateResponseMessageJSON contains the JSON metadata for the struct
// [SnippetUpdateResponseMessage]
type snippetUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Snippet Information
type SnippetUpdateResponseResult struct {
	// Creation time of the snippet
	CreatedOn string `json:"created_on"`
	// Modification time of the snippet
	ModifiedOn string `json:"modified_on"`
	// Snippet identifying name
	SnippetName string                          `json:"snippet_name"`
	JSON        snippetUpdateResponseResultJSON `json:"-"`
}

// snippetUpdateResponseResultJSON contains the JSON metadata for the struct
// [SnippetUpdateResponseResult]
type snippetUpdateResponseResultJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SnippetUpdateResponseSuccess bool

const (
	SnippetUpdateResponseSuccessTrue SnippetUpdateResponseSuccess = true
)

type SnippetListResponse struct {
	Errors   []SnippetListResponseError   `json:"errors"`
	Messages []SnippetListResponseMessage `json:"messages"`
	// List of all zone snippets
	Result []SnippetListResponseResult `json:"result"`
	// Whether the API call was successful
	Success SnippetListResponseSuccess `json:"success"`
	JSON    snippetListResponseJSON    `json:"-"`
}

// snippetListResponseJSON contains the JSON metadata for the struct
// [SnippetListResponse]
type snippetListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetListResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    snippetListResponseErrorJSON `json:"-"`
}

// snippetListResponseErrorJSON contains the JSON metadata for the struct
// [SnippetListResponseError]
type snippetListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetListResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    snippetListResponseMessageJSON `json:"-"`
}

// snippetListResponseMessageJSON contains the JSON metadata for the struct
// [SnippetListResponseMessage]
type snippetListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Snippet Information
type SnippetListResponseResult struct {
	// Creation time of the snippet
	CreatedOn string `json:"created_on"`
	// Modification time of the snippet
	ModifiedOn string `json:"modified_on"`
	// Snippet identifying name
	SnippetName string                        `json:"snippet_name"`
	JSON        snippetListResponseResultJSON `json:"-"`
}

// snippetListResponseResultJSON contains the JSON metadata for the struct
// [SnippetListResponseResult]
type snippetListResponseResultJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SnippetListResponseSuccess bool

const (
	SnippetListResponseSuccessTrue SnippetListResponseSuccess = true
)

type SnippetDeleteResponse struct {
	Errors   []SnippetDeleteResponseError   `json:"errors,required"`
	Messages []SnippetDeleteResponseMessage `json:"messages,required"`
	Result   SnippetDeleteResponseResult    `json:"result,required"`
	// Whether the API call was successful
	Success SnippetDeleteResponseSuccess `json:"success,required"`
	JSON    snippetDeleteResponseJSON    `json:"-"`
}

// snippetDeleteResponseJSON contains the JSON metadata for the struct
// [SnippetDeleteResponse]
type snippetDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetDeleteResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    snippetDeleteResponseErrorJSON `json:"-"`
}

// snippetDeleteResponseErrorJSON contains the JSON metadata for the struct
// [SnippetDeleteResponseError]
type snippetDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetDeleteResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    snippetDeleteResponseMessageJSON `json:"-"`
}

// snippetDeleteResponseMessageJSON contains the JSON metadata for the struct
// [SnippetDeleteResponseMessage]
type snippetDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [SnippetDeleteResponseResultUnknown],
// [SnippetDeleteResponseResultArray] or [shared.UnionString].
type SnippetDeleteResponseResult interface {
	ImplementsSnippetDeleteResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SnippetDeleteResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SnippetDeleteResponseResultArray []interface{}

func (r SnippetDeleteResponseResultArray) ImplementsSnippetDeleteResponseResult() {}

// Whether the API call was successful
type SnippetDeleteResponseSuccess bool

const (
	SnippetDeleteResponseSuccessTrue SnippetDeleteResponseSuccess = true
)

type SnippetUpdateParams struct {
	// Content files of uploaded snippet
	Files    param.Field[string]                      `json:"files"`
	Metadata param.Field[SnippetUpdateParamsMetadata] `json:"metadata"`
}

func (r SnippetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnippetUpdateParamsMetadata struct {
	// Main module name of uploaded snippet
	MainModule param.Field[string] `json:"main_module"`
}

func (r SnippetUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

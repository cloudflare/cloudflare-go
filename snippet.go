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

// All Snippets
func (r *SnippetService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *[]SnippetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SnippetListResponseEnvelope
	path := fmt.Sprintf("zones/%s/snippets", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Snippet
func (r *SnippetService) Delete(ctx context.Context, zoneIdentifier string, snippetName string, opts ...option.RequestOption) (res *SnippetDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SnippetDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Snippet
func (r *SnippetService) Get(ctx context.Context, zoneIdentifier string, snippetName string, opts ...option.RequestOption) (res *SnippetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SnippetGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Put Snippet
func (r *SnippetService) Replace(ctx context.Context, zoneIdentifier string, snippetName string, body SnippetReplaceParams, opts ...option.RequestOption) (res *SnippetReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SnippetReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Snippet Information
type SnippetListResponse struct {
	// Creation time of the snippet
	CreatedOn string `json:"created_on"`
	// Modification time of the snippet
	ModifiedOn string `json:"modified_on"`
	// Snippet identifying name
	SnippetName string                  `json:"snippet_name"`
	JSON        snippetListResponseJSON `json:"-"`
}

// snippetListResponseJSON contains the JSON metadata for the struct
// [SnippetListResponse]
type snippetListResponseJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [SnippetDeleteResponseUnknown], [SnippetDeleteResponseArray]
// or [shared.UnionString].
type SnippetDeleteResponse interface {
	ImplementsSnippetDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SnippetDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SnippetDeleteResponseArray []interface{}

func (r SnippetDeleteResponseArray) ImplementsSnippetDeleteResponse() {}

// Snippet Information
type SnippetGetResponse struct {
	// Creation time of the snippet
	CreatedOn string `json:"created_on"`
	// Modification time of the snippet
	ModifiedOn string `json:"modified_on"`
	// Snippet identifying name
	SnippetName string                 `json:"snippet_name"`
	JSON        snippetGetResponseJSON `json:"-"`
}

// snippetGetResponseJSON contains the JSON metadata for the struct
// [SnippetGetResponse]
type snippetGetResponseJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Snippet Information
type SnippetReplaceResponse struct {
	// Creation time of the snippet
	CreatedOn string `json:"created_on"`
	// Modification time of the snippet
	ModifiedOn string `json:"modified_on"`
	// Snippet identifying name
	SnippetName string                     `json:"snippet_name"`
	JSON        snippetReplaceResponseJSON `json:"-"`
}

// snippetReplaceResponseJSON contains the JSON metadata for the struct
// [SnippetReplaceResponse]
type snippetReplaceResponseJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetListResponseEnvelope struct {
	Errors   []SnippetListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SnippetListResponseEnvelopeMessages `json:"messages,required"`
	// List of all zone snippets
	Result []SnippetListResponse `json:"result,required"`
	// Whether the API call was successful
	Success SnippetListResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetListResponseEnvelopeJSON    `json:"-"`
}

// snippetListResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetListResponseEnvelope]
type snippetListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    snippetListResponseEnvelopeErrorsJSON `json:"-"`
}

// snippetListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SnippetListResponseEnvelopeErrors]
type snippetListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    snippetListResponseEnvelopeMessagesJSON `json:"-"`
}

// snippetListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SnippetListResponseEnvelopeMessages]
type snippetListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SnippetListResponseEnvelopeSuccess bool

const (
	SnippetListResponseEnvelopeSuccessTrue SnippetListResponseEnvelopeSuccess = true
)

type SnippetDeleteResponseEnvelope struct {
	Errors   []SnippetDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SnippetDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SnippetDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SnippetDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetDeleteResponseEnvelopeJSON    `json:"-"`
}

// snippetDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetDeleteResponseEnvelope]
type snippetDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    snippetDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// snippetDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SnippetDeleteResponseEnvelopeErrors]
type snippetDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    snippetDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// snippetDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SnippetDeleteResponseEnvelopeMessages]
type snippetDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SnippetDeleteResponseEnvelopeSuccess bool

const (
	SnippetDeleteResponseEnvelopeSuccessTrue SnippetDeleteResponseEnvelopeSuccess = true
)

type SnippetGetResponseEnvelope struct {
	Errors   []SnippetGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SnippetGetResponseEnvelopeMessages `json:"messages,required"`
	// Snippet Information
	Result SnippetGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success SnippetGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetGetResponseEnvelopeJSON    `json:"-"`
}

// snippetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetGetResponseEnvelope]
type snippetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    snippetGetResponseEnvelopeErrorsJSON `json:"-"`
}

// snippetGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SnippetGetResponseEnvelopeErrors]
type snippetGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    snippetGetResponseEnvelopeMessagesJSON `json:"-"`
}

// snippetGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SnippetGetResponseEnvelopeMessages]
type snippetGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SnippetGetResponseEnvelopeSuccess bool

const (
	SnippetGetResponseEnvelopeSuccessTrue SnippetGetResponseEnvelopeSuccess = true
)

type SnippetReplaceParams struct {
	// Content files of uploaded snippet
	Files    param.Field[string]                       `json:"files"`
	Metadata param.Field[SnippetReplaceParamsMetadata] `json:"metadata"`
}

func (r SnippetReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnippetReplaceParamsMetadata struct {
	// Main module name of uploaded snippet
	MainModule param.Field[string] `json:"main_module"`
}

func (r SnippetReplaceParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SnippetReplaceResponseEnvelope struct {
	Errors   []SnippetReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SnippetReplaceResponseEnvelopeMessages `json:"messages,required"`
	// Snippet Information
	Result SnippetReplaceResponse `json:"result,required"`
	// Whether the API call was successful
	Success SnippetReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetReplaceResponseEnvelopeJSON    `json:"-"`
}

// snippetReplaceResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetReplaceResponseEnvelope]
type snippetReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetReplaceResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    snippetReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// snippetReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SnippetReplaceResponseEnvelopeErrors]
type snippetReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetReplaceResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    snippetReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// snippetReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SnippetReplaceResponseEnvelopeMessages]
type snippetReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SnippetReplaceResponseEnvelopeSuccess bool

const (
	SnippetReplaceResponseEnvelopeSuccessTrue SnippetReplaceResponseEnvelopeSuccess = true
)

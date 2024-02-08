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
func (r *SnippetService) Update(ctx context.Context, zoneIdentifier string, snippetName string, body SnippetUpdateParams, opts ...option.RequestOption) (res *SnippetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SnippetUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/snippets/%s", zoneIdentifier, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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
type SnippetUpdateResponse struct {
	// Creation time of the snippet
	CreatedOn string `json:"created_on"`
	// Modification time of the snippet
	ModifiedOn string `json:"modified_on"`
	// Snippet identifying name
	SnippetName string                    `json:"snippet_name"`
	JSON        snippetUpdateResponseJSON `json:"-"`
}

// snippetUpdateResponseJSON contains the JSON metadata for the struct
// [SnippetUpdateResponse]
type snippetUpdateResponseJSON struct {
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	SnippetName apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type SnippetUpdateResponseEnvelope struct {
	Errors   []SnippetUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SnippetUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Snippet Information
	Result SnippetUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success SnippetUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetUpdateResponseEnvelopeJSON    `json:"-"`
}

// snippetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetUpdateResponseEnvelope]
type snippetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetUpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    snippetUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// snippetUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SnippetUpdateResponseEnvelopeErrors]
type snippetUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SnippetUpdateResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    snippetUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// snippetUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SnippetUpdateResponseEnvelopeMessages]
type snippetUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SnippetUpdateResponseEnvelopeSuccess bool

const (
	SnippetUpdateResponseEnvelopeSuccessTrue SnippetUpdateResponseEnvelopeSuccess = true
)

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

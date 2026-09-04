// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar_sandbox

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// ExtensionService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtensionService] method instead.
type ExtensionService struct {
	Options []option.RequestOption
}

// NewExtensionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExtensionService(opts ...option.RequestOption) (r *ExtensionService) {
	r = &ExtensionService{}
	r.Options = opts
	return
}

// Returns metadata and JSON Schema documents describing the expected input
// structure for registration operations on each supported extension (TLD).
//
// This endpoint uses cursor-based pagination. Results are ordered by extension
// name by default. To fetch the next page, pass the `cursor` value from the
// `result_info` object in the response as the `cursor` query parameter in your
// next request. An empty `cursor` string indicates there are no more pages.
//
// Supports HTTP conditional GET via `ETag`. Include the `ETag` value from a
// previous response in an `If-None-Match` header to receive a `304 Not Modified`
// when the data has not changed.
func (r *ExtensionService) List(ctx context.Context, params ExtensionListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[ExtensionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar-sandbox/extensions", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns metadata and JSON Schema documents describing the expected input
// structure for registration operations on each supported extension (TLD).
//
// This endpoint uses cursor-based pagination. Results are ordered by extension
// name by default. To fetch the next page, pass the `cursor` value from the
// `result_info` object in the response as the `cursor` query parameter in your
// next request. An empty `cursor` string indicates there are no more pages.
//
// Supports HTTP conditional GET via `ETag`. Include the `ETag` value from a
// previous response in an `If-None-Match` header to receive a `304 Not Modified`
// when the data has not changed.
func (r *ExtensionService) ListAutoPaging(ctx context.Context, params ExtensionListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[ExtensionListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Returns metadata and JSON Schema documents describing the expected input
// structure for registration operations on a specific extension (TLD).
//
// Supports HTTP conditional GET via `ETag`. Include the `ETag` value from a
// previous response in an `If-None-Match` header to receive a `304 Not Modified`
// when the data has not changed.
func (r *ExtensionService) Get(ctx context.Context, extension string, query ExtensionGetParams, opts ...option.RequestOption) (res *ExtensionGetResponse, err error) {
	var env ExtensionGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if extension == "" {
		err = errors.New("missing required extension parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/registrar-sandbox/extensions/%s", query.AccountID, extension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Extension entry with metadata and JSON Schema documents for the registration
// operation.
type ExtensionListResponse struct {
	// Extension metadata.
	Metadata ExtensionListResponseMetadata `json:"metadata" api:"required"`
	// JSON Schema describing the expected input structure for registration operations
	// on this extension.
	RegistrationSchema interface{}               `json:"registration_schema" api:"required"`
	JSON               extensionListResponseJSON `json:"-"`
}

// extensionListResponseJSON contains the JSON metadata for the struct
// [ExtensionListResponse]
type extensionListResponseJSON struct {
	Metadata           apijson.Field
	RegistrationSchema apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ExtensionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionListResponseJSON) RawJSON() string {
	return r.raw
}

// Extension metadata.
type ExtensionListResponseMetadata struct {
	// The full name of the extension. For example, "co.uk", or "uk".
	Name string `json:"name" api:"required"`
	// The TLD of the extension. For example, for "co.uk", it is "uk". For "uk", it is
	// "uk".
	TLD  string                            `json:"tld" api:"required"`
	JSON extensionListResponseMetadataJSON `json:"-"`
}

// extensionListResponseMetadataJSON contains the JSON metadata for the struct
// [ExtensionListResponseMetadata]
type extensionListResponseMetadataJSON struct {
	Name        apijson.Field
	TLD         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtensionListResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionListResponseMetadataJSON) RawJSON() string {
	return r.raw
}

// Extension entry with metadata and JSON Schema documents for the registration
// operation.
type ExtensionGetResponse struct {
	// Extension metadata.
	Metadata ExtensionGetResponseMetadata `json:"metadata" api:"required"`
	// JSON Schema describing the expected input structure for registration operations
	// on this extension.
	RegistrationSchema interface{}              `json:"registration_schema" api:"required"`
	JSON               extensionGetResponseJSON `json:"-"`
}

// extensionGetResponseJSON contains the JSON metadata for the struct
// [ExtensionGetResponse]
type extensionGetResponseJSON struct {
	Metadata           apijson.Field
	RegistrationSchema apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ExtensionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Extension metadata.
type ExtensionGetResponseMetadata struct {
	// The full name of the extension. For example, "co.uk", or "uk".
	Name string `json:"name" api:"required"`
	// The TLD of the extension. For example, for "co.uk", it is "uk". For "uk", it is
	// "uk".
	TLD  string                           `json:"tld" api:"required"`
	JSON extensionGetResponseMetadataJSON `json:"-"`
}

// extensionGetResponseMetadataJSON contains the JSON metadata for the struct
// [ExtensionGetResponseMetadata]
type extensionGetResponseMetadataJSON struct {
	Name        apijson.Field
	TLD         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtensionGetResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionGetResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type ExtensionListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Opaque token from a previous response's `result_info.cursor`. Pass this value to
	// fetch the next page of results. Omit (or pass an empty string) for the first
	// page.
	Cursor param.Field[string] `query:"cursor"`
	// Sort direction for results. Defaults to ascending order.
	Direction param.Field[ExtensionListParamsDirection] `query:"direction"`
	// Filter extensions by exact name match. For example, `name=com` returns only the
	// `com` extension.
	Name param.Field[string] `query:"name"`
	// Number of items to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Column to sort results by. Defaults to `name` when omitted.
	SortBy param.Field[ExtensionListParamsSortBy] `query:"sort_by"`
}

// URLQuery serializes [ExtensionListParams]'s query parameters as `url.Values`.
func (r ExtensionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sort direction for results. Defaults to ascending order.
type ExtensionListParamsDirection string

const (
	ExtensionListParamsDirectionAsc  ExtensionListParamsDirection = "asc"
	ExtensionListParamsDirectionDesc ExtensionListParamsDirection = "desc"
)

func (r ExtensionListParamsDirection) IsKnown() bool {
	switch r {
	case ExtensionListParamsDirectionAsc, ExtensionListParamsDirectionDesc:
		return true
	}
	return false
}

// Column to sort results by. Defaults to `name` when omitted.
type ExtensionListParamsSortBy string

const (
	ExtensionListParamsSortByName      ExtensionListParamsSortBy = "name"
	ExtensionListParamsSortByCreatedAt ExtensionListParamsSortBy = "created_at"
	ExtensionListParamsSortByUpdatedAt ExtensionListParamsSortBy = "updated_at"
)

func (r ExtensionListParamsSortBy) IsKnown() bool {
	switch r {
	case ExtensionListParamsSortByName, ExtensionListParamsSortByCreatedAt, ExtensionListParamsSortByUpdatedAt:
		return true
	}
	return false
}

type ExtensionGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ExtensionGetResponseEnvelope struct {
	Errors   []ExtensionGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ExtensionGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Extension entry with metadata and JSON Schema documents for the registration
	// operation.
	Result ExtensionGetResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success ExtensionGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    extensionGetResponseEnvelopeJSON    `json:"-"`
}

// extensionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ExtensionGetResponseEnvelope]
type extensionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtensionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ExtensionGetResponseEnvelopeErrors struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source ExtensionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   extensionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// extensionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ExtensionGetResponseEnvelopeErrors]
type extensionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtensionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type ExtensionGetResponseEnvelopeErrorsSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                       `json:"pointer" api:"required"`
	JSON    extensionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// extensionGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ExtensionGetResponseEnvelopeErrorsSource]
type extensionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtensionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ExtensionGetResponseEnvelopeMessages struct {
	Code    int64  `json:"code" api:"required"`
	Message string `json:"message" api:"required"`
	// Location of the invalid value that caused the error.
	Source ExtensionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   extensionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// extensionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ExtensionGetResponseEnvelopeMessages]
type extensionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtensionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Location of the invalid value that caused the error.
type ExtensionGetResponseEnvelopeMessagesSource struct {
	// JSON Pointer to the invalid or missing request value.
	Pointer string                                         `json:"pointer" api:"required"`
	JSON    extensionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// extensionGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ExtensionGetResponseEnvelopeMessagesSource]
type extensionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExtensionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r extensionGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ExtensionGetResponseEnvelopeSuccess bool

const (
	ExtensionGetResponseEnvelopeSuccessTrue ExtensionGetResponseEnvelopeSuccess = true
)

func (r ExtensionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ExtensionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

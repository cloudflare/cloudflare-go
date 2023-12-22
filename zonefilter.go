// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneFilterService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewZoneFilterService] method instead.
type ZoneFilterService struct {
	Options []option.RequestOption
}

// NewZoneFilterService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneFilterService(opts ...option.RequestOption) (r *ZoneFilterService) {
	r = &ZoneFilterService{}
	r.Options = opts
	return
}

// Fetches the details of a filter.
func (r *ZoneFilterService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *SchemasFilterResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/filters/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing filter.
func (r *ZoneFilterService) Update(ctx context.Context, zoneIdentifier string, id string, body ZoneFilterUpdateParams, opts ...option.RequestOption) (res *SchemasFilterResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/filters/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an existing filter.
func (r *ZoneFilterService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FilterDeleteResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/filters/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates one or more filters.
func (r *ZoneFilterService) FiltersNewFilters(ctx context.Context, zoneIdentifier string, body ZoneFilterFiltersNewFiltersParams, opts ...option.RequestOption) (res *SchemasFilterResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/filters", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches filters in a zone. You can filter the results using several optional
// parameters.
func (r *ZoneFilterService) FiltersListFilters(ctx context.Context, zoneIdentifier string, query ZoneFilterFiltersListFiltersParams, opts ...option.RequestOption) (res *SchemasFilterResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/filters", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates one or more existing filters.
func (r *ZoneFilterService) FiltersUpdateFilters(ctx context.Context, zoneIdentifier string, body ZoneFilterFiltersUpdateFiltersParams, opts ...option.RequestOption) (res *SchemasFilterResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/filters", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type FilterDeleteResponseSingle struct {
	Errors   []FilterDeleteResponseSingleError   `json:"errors"`
	Messages []FilterDeleteResponseSingleMessage `json:"messages"`
	Result   FilterDeleteResponseSingleResult    `json:"result"`
	// Whether the API call was successful
	Success FilterDeleteResponseSingleSuccess `json:"success"`
	JSON    filterDeleteResponseSingleJSON    `json:"-"`
}

// filterDeleteResponseSingleJSON contains the JSON metadata for the struct
// [FilterDeleteResponseSingle]
type filterDeleteResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterDeleteResponseSingleError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    filterDeleteResponseSingleErrorJSON `json:"-"`
}

// filterDeleteResponseSingleErrorJSON contains the JSON metadata for the struct
// [FilterDeleteResponseSingleError]
type filterDeleteResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterDeleteResponseSingleMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    filterDeleteResponseSingleMessageJSON `json:"-"`
}

// filterDeleteResponseSingleMessageJSON contains the JSON metadata for the struct
// [FilterDeleteResponseSingleMessage]
type filterDeleteResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterDeleteResponseSingleResult struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                               `json:"ref"`
	JSON filterDeleteResponseSingleResultJSON `json:"-"`
}

// filterDeleteResponseSingleResultJSON contains the JSON metadata for the struct
// [FilterDeleteResponseSingleResult]
type filterDeleteResponseSingleResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FilterDeleteResponseSingleSuccess bool

const (
	FilterDeleteResponseSingleSuccessTrue FilterDeleteResponseSingleSuccess = true
)

type SchemasFilterResponseCollection struct {
	Errors     []SchemasFilterResponseCollectionError    `json:"errors"`
	Messages   []SchemasFilterResponseCollectionMessage  `json:"messages"`
	Result     []SchemasFilterResponseCollectionResult   `json:"result"`
	ResultInfo SchemasFilterResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SchemasFilterResponseCollectionSuccess `json:"success"`
	JSON    schemasFilterResponseCollectionJSON    `json:"-"`
}

// schemasFilterResponseCollectionJSON contains the JSON metadata for the struct
// [SchemasFilterResponseCollection]
type schemasFilterResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasFilterResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasFilterResponseCollectionError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    schemasFilterResponseCollectionErrorJSON `json:"-"`
}

// schemasFilterResponseCollectionErrorJSON contains the JSON metadata for the
// struct [SchemasFilterResponseCollectionError]
type schemasFilterResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasFilterResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasFilterResponseCollectionMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    schemasFilterResponseCollectionMessageJSON `json:"-"`
}

// schemasFilterResponseCollectionMessageJSON contains the JSON metadata for the
// struct [SchemasFilterResponseCollectionMessage]
type schemasFilterResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasFilterResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasFilterResponseCollectionResult struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                                    `json:"ref"`
	JSON schemasFilterResponseCollectionResultJSON `json:"-"`
}

// schemasFilterResponseCollectionResultJSON contains the JSON metadata for the
// struct [SchemasFilterResponseCollectionResult]
type schemasFilterResponseCollectionResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasFilterResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasFilterResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       schemasFilterResponseCollectionResultInfoJSON `json:"-"`
}

// schemasFilterResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [SchemasFilterResponseCollectionResultInfo]
type schemasFilterResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasFilterResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasFilterResponseCollectionSuccess bool

const (
	SchemasFilterResponseCollectionSuccessTrue SchemasFilterResponseCollectionSuccess = true
)

type SchemasFilterResponseSingle struct {
	Errors   []SchemasFilterResponseSingleError   `json:"errors"`
	Messages []SchemasFilterResponseSingleMessage `json:"messages"`
	Result   SchemasFilterResponseSingleResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success SchemasFilterResponseSingleSuccess `json:"success"`
	JSON    schemasFilterResponseSingleJSON    `json:"-"`
}

// schemasFilterResponseSingleJSON contains the JSON metadata for the struct
// [SchemasFilterResponseSingle]
type schemasFilterResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasFilterResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasFilterResponseSingleError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    schemasFilterResponseSingleErrorJSON `json:"-"`
}

// schemasFilterResponseSingleErrorJSON contains the JSON metadata for the struct
// [SchemasFilterResponseSingleError]
type schemasFilterResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasFilterResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasFilterResponseSingleMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    schemasFilterResponseSingleMessageJSON `json:"-"`
}

// schemasFilterResponseSingleMessageJSON contains the JSON metadata for the struct
// [SchemasFilterResponseSingleMessage]
type schemasFilterResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasFilterResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasFilterResponseSingleResult struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                                `json:"ref"`
	JSON schemasFilterResponseSingleResultJSON `json:"-"`
}

// schemasFilterResponseSingleResultJSON contains the JSON metadata for the struct
// [SchemasFilterResponseSingleResult]
type schemasFilterResponseSingleResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasFilterResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasFilterResponseSingleSuccess bool

const (
	SchemasFilterResponseSingleSuccessTrue SchemasFilterResponseSingleSuccess = true
)

type ZoneFilterUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFilterUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFilterFiltersNewFiltersParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFilterFiltersNewFiltersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFilterFiltersListFiltersParams struct {
	// A case-insensitive string to find in the description.
	Description param.Field[string] `query:"description"`
	// A case-insensitive string to find in the expression.
	Expression param.Field[string] `query:"expression"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// When true, indicates that the filter is currently paused.
	Paused param.Field[bool] `query:"paused"`
	// Number of filters per page.
	PerPage param.Field[float64] `query:"per_page"`
	// The filter ref (a short reference tag) to search for. Must be an exact match.
	Ref param.Field[string] `query:"ref"`
}

// URLQuery serializes [ZoneFilterFiltersListFiltersParams]'s query parameters as
// `url.Values`.
func (r ZoneFilterFiltersListFiltersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneFilterFiltersUpdateFiltersParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFilterFiltersUpdateFiltersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

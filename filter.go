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

// FilterService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewFilterService] method instead.
type FilterService struct {
	Options []option.RequestOption
}

// NewFilterService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFilterService(opts ...option.RequestOption) (r *FilterService) {
	r = &FilterService{}
	r.Options = opts
	return
}

// Updates an existing filter.
func (r *FilterService) Update(ctx context.Context, zoneIdentifier string, id string, body FilterUpdateParams, opts ...option.RequestOption) (res *FilterUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FilterUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/filters/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing filter.
func (r *FilterService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FilterDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FilterDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/filters/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates one or more filters.
func (r *FilterService) FiltersNewFilters(ctx context.Context, zoneIdentifier string, body FilterFiltersNewFiltersParams, opts ...option.RequestOption) (res *[]FilterFiltersNewFiltersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FilterFiltersNewFiltersResponseEnvelope
	path := fmt.Sprintf("zones/%s/filters", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches filters in a zone. You can filter the results using several optional
// parameters.
func (r *FilterService) FiltersListFilters(ctx context.Context, zoneIdentifier string, query FilterFiltersListFiltersParams, opts ...option.RequestOption) (res *[]FilterFiltersListFiltersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FilterFiltersListFiltersResponseEnvelope
	path := fmt.Sprintf("zones/%s/filters", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates one or more existing filters.
func (r *FilterService) FiltersUpdateFilters(ctx context.Context, zoneIdentifier string, body FilterFiltersUpdateFiltersParams, opts ...option.RequestOption) (res *[]FilterFiltersUpdateFiltersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FilterFiltersUpdateFiltersResponseEnvelope
	path := fmt.Sprintf("zones/%s/filters", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a filter.
func (r *FilterService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FilterGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FilterGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/filters/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FilterUpdateResponse struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression,required"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                   `json:"ref"`
	JSON filterUpdateResponseJSON `json:"-"`
}

// filterUpdateResponseJSON contains the JSON metadata for the struct
// [FilterUpdateResponse]
type filterUpdateResponseJSON struct {
	ID          apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Description apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterDeleteResponse struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                   `json:"ref"`
	JSON filterDeleteResponseJSON `json:"-"`
}

// filterDeleteResponseJSON contains the JSON metadata for the struct
// [FilterDeleteResponse]
type filterDeleteResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterFiltersNewFiltersResponse struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression,required"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                              `json:"ref"`
	JSON filterFiltersNewFiltersResponseJSON `json:"-"`
}

// filterFiltersNewFiltersResponseJSON contains the JSON metadata for the struct
// [FilterFiltersNewFiltersResponse]
type filterFiltersNewFiltersResponseJSON struct {
	ID          apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Description apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersNewFiltersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterFiltersListFiltersResponse struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression,required"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                               `json:"ref"`
	JSON filterFiltersListFiltersResponseJSON `json:"-"`
}

// filterFiltersListFiltersResponseJSON contains the JSON metadata for the struct
// [FilterFiltersListFiltersResponse]
type filterFiltersListFiltersResponseJSON struct {
	ID          apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Description apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersListFiltersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterFiltersUpdateFiltersResponse struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression,required"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                                 `json:"ref"`
	JSON filterFiltersUpdateFiltersResponseJSON `json:"-"`
}

// filterFiltersUpdateFiltersResponseJSON contains the JSON metadata for the struct
// [FilterFiltersUpdateFiltersResponse]
type filterFiltersUpdateFiltersResponseJSON struct {
	ID          apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Description apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersUpdateFiltersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterGetResponse struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression,required"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused,required"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                `json:"ref"`
	JSON filterGetResponseJSON `json:"-"`
}

// filterGetResponseJSON contains the JSON metadata for the struct
// [FilterGetResponse]
type filterGetResponseJSON struct {
	ID          apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Description apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FilterUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FilterUpdateResponseEnvelope struct {
	Errors   []FilterUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FilterUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   FilterUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FilterUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    filterUpdateResponseEnvelopeJSON    `json:"-"`
}

// filterUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [FilterUpdateResponseEnvelope]
type filterUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterUpdateResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    filterUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// filterUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [FilterUpdateResponseEnvelopeErrors]
type filterUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterUpdateResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    filterUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// filterUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FilterUpdateResponseEnvelopeMessages]
type filterUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FilterUpdateResponseEnvelopeSuccess bool

const (
	FilterUpdateResponseEnvelopeSuccessTrue FilterUpdateResponseEnvelopeSuccess = true
)

type FilterDeleteResponseEnvelope struct {
	Errors   []FilterDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FilterDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   FilterDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FilterDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    filterDeleteResponseEnvelopeJSON    `json:"-"`
}

// filterDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [FilterDeleteResponseEnvelope]
type filterDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterDeleteResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    filterDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// filterDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [FilterDeleteResponseEnvelopeErrors]
type filterDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterDeleteResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    filterDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// filterDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FilterDeleteResponseEnvelopeMessages]
type filterDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FilterDeleteResponseEnvelopeSuccess bool

const (
	FilterDeleteResponseEnvelopeSuccessTrue FilterDeleteResponseEnvelopeSuccess = true
)

type FilterFiltersNewFiltersParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FilterFiltersNewFiltersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FilterFiltersNewFiltersResponseEnvelope struct {
	Errors   []FilterFiltersNewFiltersResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FilterFiltersNewFiltersResponseEnvelopeMessages `json:"messages,required"`
	Result   []FilterFiltersNewFiltersResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FilterFiltersNewFiltersResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FilterFiltersNewFiltersResponseEnvelopeResultInfo `json:"result_info"`
	JSON       filterFiltersNewFiltersResponseEnvelopeJSON       `json:"-"`
}

// filterFiltersNewFiltersResponseEnvelopeJSON contains the JSON metadata for the
// struct [FilterFiltersNewFiltersResponseEnvelope]
type filterFiltersNewFiltersResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersNewFiltersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterFiltersNewFiltersResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    filterFiltersNewFiltersResponseEnvelopeErrorsJSON `json:"-"`
}

// filterFiltersNewFiltersResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FilterFiltersNewFiltersResponseEnvelopeErrors]
type filterFiltersNewFiltersResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersNewFiltersResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterFiltersNewFiltersResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    filterFiltersNewFiltersResponseEnvelopeMessagesJSON `json:"-"`
}

// filterFiltersNewFiltersResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FilterFiltersNewFiltersResponseEnvelopeMessages]
type filterFiltersNewFiltersResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersNewFiltersResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FilterFiltersNewFiltersResponseEnvelopeSuccess bool

const (
	FilterFiltersNewFiltersResponseEnvelopeSuccessTrue FilterFiltersNewFiltersResponseEnvelopeSuccess = true
)

type FilterFiltersNewFiltersResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       filterFiltersNewFiltersResponseEnvelopeResultInfoJSON `json:"-"`
}

// filterFiltersNewFiltersResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [FilterFiltersNewFiltersResponseEnvelopeResultInfo]
type filterFiltersNewFiltersResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersNewFiltersResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterFiltersListFiltersParams struct {
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

// URLQuery serializes [FilterFiltersListFiltersParams]'s query parameters as
// `url.Values`.
func (r FilterFiltersListFiltersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FilterFiltersListFiltersResponseEnvelope struct {
	Errors   []FilterFiltersListFiltersResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FilterFiltersListFiltersResponseEnvelopeMessages `json:"messages,required"`
	Result   []FilterFiltersListFiltersResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FilterFiltersListFiltersResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FilterFiltersListFiltersResponseEnvelopeResultInfo `json:"result_info"`
	JSON       filterFiltersListFiltersResponseEnvelopeJSON       `json:"-"`
}

// filterFiltersListFiltersResponseEnvelopeJSON contains the JSON metadata for the
// struct [FilterFiltersListFiltersResponseEnvelope]
type filterFiltersListFiltersResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersListFiltersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterFiltersListFiltersResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    filterFiltersListFiltersResponseEnvelopeErrorsJSON `json:"-"`
}

// filterFiltersListFiltersResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [FilterFiltersListFiltersResponseEnvelopeErrors]
type filterFiltersListFiltersResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersListFiltersResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterFiltersListFiltersResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    filterFiltersListFiltersResponseEnvelopeMessagesJSON `json:"-"`
}

// filterFiltersListFiltersResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FilterFiltersListFiltersResponseEnvelopeMessages]
type filterFiltersListFiltersResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersListFiltersResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FilterFiltersListFiltersResponseEnvelopeSuccess bool

const (
	FilterFiltersListFiltersResponseEnvelopeSuccessTrue FilterFiltersListFiltersResponseEnvelopeSuccess = true
)

type FilterFiltersListFiltersResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       filterFiltersListFiltersResponseEnvelopeResultInfoJSON `json:"-"`
}

// filterFiltersListFiltersResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [FilterFiltersListFiltersResponseEnvelopeResultInfo]
type filterFiltersListFiltersResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersListFiltersResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterFiltersUpdateFiltersParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FilterFiltersUpdateFiltersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FilterFiltersUpdateFiltersResponseEnvelope struct {
	Errors   []FilterFiltersUpdateFiltersResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FilterFiltersUpdateFiltersResponseEnvelopeMessages `json:"messages,required"`
	Result   []FilterFiltersUpdateFiltersResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FilterFiltersUpdateFiltersResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FilterFiltersUpdateFiltersResponseEnvelopeResultInfo `json:"result_info"`
	JSON       filterFiltersUpdateFiltersResponseEnvelopeJSON       `json:"-"`
}

// filterFiltersUpdateFiltersResponseEnvelopeJSON contains the JSON metadata for
// the struct [FilterFiltersUpdateFiltersResponseEnvelope]
type filterFiltersUpdateFiltersResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersUpdateFiltersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterFiltersUpdateFiltersResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    filterFiltersUpdateFiltersResponseEnvelopeErrorsJSON `json:"-"`
}

// filterFiltersUpdateFiltersResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [FilterFiltersUpdateFiltersResponseEnvelopeErrors]
type filterFiltersUpdateFiltersResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersUpdateFiltersResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterFiltersUpdateFiltersResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    filterFiltersUpdateFiltersResponseEnvelopeMessagesJSON `json:"-"`
}

// filterFiltersUpdateFiltersResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [FilterFiltersUpdateFiltersResponseEnvelopeMessages]
type filterFiltersUpdateFiltersResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersUpdateFiltersResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FilterFiltersUpdateFiltersResponseEnvelopeSuccess bool

const (
	FilterFiltersUpdateFiltersResponseEnvelopeSuccessTrue FilterFiltersUpdateFiltersResponseEnvelopeSuccess = true
)

type FilterFiltersUpdateFiltersResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       filterFiltersUpdateFiltersResponseEnvelopeResultInfoJSON `json:"-"`
}

// filterFiltersUpdateFiltersResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [FilterFiltersUpdateFiltersResponseEnvelopeResultInfo]
type filterFiltersUpdateFiltersResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterFiltersUpdateFiltersResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterGetResponseEnvelope struct {
	Errors   []FilterGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FilterGetResponseEnvelopeMessages `json:"messages,required"`
	Result   FilterGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FilterGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    filterGetResponseEnvelopeJSON    `json:"-"`
}

// filterGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [FilterGetResponseEnvelope]
type filterGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterGetResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    filterGetResponseEnvelopeErrorsJSON `json:"-"`
}

// filterGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [FilterGetResponseEnvelopeErrors]
type filterGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterGetResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    filterGetResponseEnvelopeMessagesJSON `json:"-"`
}

// filterGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [FilterGetResponseEnvelopeMessages]
type filterGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FilterGetResponseEnvelopeSuccess bool

const (
	FilterGetResponseEnvelopeSuccessTrue FilterGetResponseEnvelopeSuccess = true
)

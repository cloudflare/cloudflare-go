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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
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

// Creates one or more filters.
func (r *FilterService) New(ctx context.Context, zoneIdentifier string, body FilterNewParams, opts ...option.RequestOption) (res *[]FilterNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FilterNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/filters", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

// Fetches filters in a zone. You can filter the results using several optional
// parameters.
func (r *FilterService) List(ctx context.Context, zoneIdentifier string, query FilterListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[FilterListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/filters", zoneIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Fetches filters in a zone. You can filter the results using several optional
// parameters.
func (r *FilterService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query FilterListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[FilterListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
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

type FilterNewResponse struct {
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
	JSON filterNewResponseJSON `json:"-"`
}

// filterNewResponseJSON contains the JSON metadata for the struct
// [FilterNewResponse]
type filterNewResponseJSON struct {
	ID          apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Description apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type FilterListResponse struct {
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
	Ref  string                 `json:"ref"`
	JSON filterListResponseJSON `json:"-"`
}

// filterListResponseJSON contains the JSON metadata for the struct
// [FilterListResponse]
type filterListResponseJSON struct {
	ID          apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Description apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterListResponse) UnmarshalJSON(data []byte) (err error) {
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

type FilterNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FilterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FilterNewResponseEnvelope struct {
	Errors   []FilterNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FilterNewResponseEnvelopeMessages `json:"messages,required"`
	Result   []FilterNewResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FilterNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FilterNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       filterNewResponseEnvelopeJSON       `json:"-"`
}

// filterNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [FilterNewResponseEnvelope]
type filterNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterNewResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    filterNewResponseEnvelopeErrorsJSON `json:"-"`
}

// filterNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [FilterNewResponseEnvelopeErrors]
type filterNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterNewResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    filterNewResponseEnvelopeMessagesJSON `json:"-"`
}

// filterNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [FilterNewResponseEnvelopeMessages]
type filterNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FilterNewResponseEnvelopeSuccess bool

const (
	FilterNewResponseEnvelopeSuccessTrue FilterNewResponseEnvelopeSuccess = true
)

type FilterNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       filterNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// filterNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [FilterNewResponseEnvelopeResultInfo]
type filterNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
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

type FilterListParams struct {
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

// URLQuery serializes [FilterListParams]'s query parameters as `url.Values`.
func (r FilterListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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

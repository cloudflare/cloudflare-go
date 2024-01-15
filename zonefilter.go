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
func (r *ZoneFilterService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *ZoneFilterGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/filters/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing filter.
func (r *ZoneFilterService) Update(ctx context.Context, zoneIdentifier string, id string, body ZoneFilterUpdateParams, opts ...option.RequestOption) (res *ZoneFilterUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/filters/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an existing filter.
func (r *ZoneFilterService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *ZoneFilterDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/filters/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates one or more filters.
func (r *ZoneFilterService) FiltersNewFilters(ctx context.Context, zoneIdentifier string, body ZoneFilterFiltersNewFiltersParams, opts ...option.RequestOption) (res *ZoneFilterFiltersNewFiltersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/filters", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches filters in a zone. You can filter the results using several optional
// parameters.
func (r *ZoneFilterService) FiltersListFilters(ctx context.Context, zoneIdentifier string, query ZoneFilterFiltersListFiltersParams, opts ...option.RequestOption) (res *shared.Page[ZoneFilterFiltersListFiltersResponse], err error) {
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

// Updates one or more existing filters.
func (r *ZoneFilterService) FiltersUpdateFilters(ctx context.Context, zoneIdentifier string, body ZoneFilterFiltersUpdateFiltersParams, opts ...option.RequestOption) (res *ZoneFilterFiltersUpdateFiltersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/filters", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneFilterGetResponse struct {
	Errors   []ZoneFilterGetResponseError   `json:"errors"`
	Messages []ZoneFilterGetResponseMessage `json:"messages"`
	Result   ZoneFilterGetResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneFilterGetResponseSuccess `json:"success"`
	JSON    zoneFilterGetResponseJSON    `json:"-"`
}

// zoneFilterGetResponseJSON contains the JSON metadata for the struct
// [ZoneFilterGetResponse]
type zoneFilterGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterGetResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    zoneFilterGetResponseErrorJSON `json:"-"`
}

// zoneFilterGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneFilterGetResponseError]
type zoneFilterGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterGetResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    zoneFilterGetResponseMessageJSON `json:"-"`
}

// zoneFilterGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneFilterGetResponseMessage]
type zoneFilterGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterGetResponseResult struct {
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
	Ref  string                          `json:"ref"`
	JSON zoneFilterGetResponseResultJSON `json:"-"`
}

// zoneFilterGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneFilterGetResponseResult]
type zoneFilterGetResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFilterGetResponseSuccess bool

const (
	ZoneFilterGetResponseSuccessTrue ZoneFilterGetResponseSuccess = true
)

type ZoneFilterUpdateResponse struct {
	Errors   []ZoneFilterUpdateResponseError   `json:"errors"`
	Messages []ZoneFilterUpdateResponseMessage `json:"messages"`
	Result   ZoneFilterUpdateResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZoneFilterUpdateResponseSuccess `json:"success"`
	JSON    zoneFilterUpdateResponseJSON    `json:"-"`
}

// zoneFilterUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneFilterUpdateResponse]
type zoneFilterUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterUpdateResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    zoneFilterUpdateResponseErrorJSON `json:"-"`
}

// zoneFilterUpdateResponseErrorJSON contains the JSON metadata for the struct
// [ZoneFilterUpdateResponseError]
type zoneFilterUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterUpdateResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneFilterUpdateResponseMessageJSON `json:"-"`
}

// zoneFilterUpdateResponseMessageJSON contains the JSON metadata for the struct
// [ZoneFilterUpdateResponseMessage]
type zoneFilterUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterUpdateResponseResult struct {
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
	Ref  string                             `json:"ref"`
	JSON zoneFilterUpdateResponseResultJSON `json:"-"`
}

// zoneFilterUpdateResponseResultJSON contains the JSON metadata for the struct
// [ZoneFilterUpdateResponseResult]
type zoneFilterUpdateResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFilterUpdateResponseSuccess bool

const (
	ZoneFilterUpdateResponseSuccessTrue ZoneFilterUpdateResponseSuccess = true
)

type ZoneFilterDeleteResponse struct {
	Errors   []ZoneFilterDeleteResponseError   `json:"errors"`
	Messages []ZoneFilterDeleteResponseMessage `json:"messages"`
	Result   ZoneFilterDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneFilterDeleteResponseSuccess `json:"success"`
	JSON    zoneFilterDeleteResponseJSON    `json:"-"`
}

// zoneFilterDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneFilterDeleteResponse]
type zoneFilterDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterDeleteResponseError struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    zoneFilterDeleteResponseErrorJSON `json:"-"`
}

// zoneFilterDeleteResponseErrorJSON contains the JSON metadata for the struct
// [ZoneFilterDeleteResponseError]
type zoneFilterDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterDeleteResponseMessage struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zoneFilterDeleteResponseMessageJSON `json:"-"`
}

// zoneFilterDeleteResponseMessageJSON contains the JSON metadata for the struct
// [ZoneFilterDeleteResponseMessage]
type zoneFilterDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterDeleteResponseResult struct {
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
	Ref  string                             `json:"ref"`
	JSON zoneFilterDeleteResponseResultJSON `json:"-"`
}

// zoneFilterDeleteResponseResultJSON contains the JSON metadata for the struct
// [ZoneFilterDeleteResponseResult]
type zoneFilterDeleteResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFilterDeleteResponseSuccess bool

const (
	ZoneFilterDeleteResponseSuccessTrue ZoneFilterDeleteResponseSuccess = true
)

type ZoneFilterFiltersNewFiltersResponse struct {
	Errors     []ZoneFilterFiltersNewFiltersResponseError    `json:"errors"`
	Messages   []ZoneFilterFiltersNewFiltersResponseMessage  `json:"messages"`
	Result     []ZoneFilterFiltersNewFiltersResponseResult   `json:"result"`
	ResultInfo ZoneFilterFiltersNewFiltersResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneFilterFiltersNewFiltersResponseSuccess `json:"success"`
	JSON    zoneFilterFiltersNewFiltersResponseJSON    `json:"-"`
}

// zoneFilterFiltersNewFiltersResponseJSON contains the JSON metadata for the
// struct [ZoneFilterFiltersNewFiltersResponse]
type zoneFilterFiltersNewFiltersResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterFiltersNewFiltersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterFiltersNewFiltersResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneFilterFiltersNewFiltersResponseErrorJSON `json:"-"`
}

// zoneFilterFiltersNewFiltersResponseErrorJSON contains the JSON metadata for the
// struct [ZoneFilterFiltersNewFiltersResponseError]
type zoneFilterFiltersNewFiltersResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterFiltersNewFiltersResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterFiltersNewFiltersResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zoneFilterFiltersNewFiltersResponseMessageJSON `json:"-"`
}

// zoneFilterFiltersNewFiltersResponseMessageJSON contains the JSON metadata for
// the struct [ZoneFilterFiltersNewFiltersResponseMessage]
type zoneFilterFiltersNewFiltersResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterFiltersNewFiltersResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterFiltersNewFiltersResponseResult struct {
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
	Ref  string                                        `json:"ref"`
	JSON zoneFilterFiltersNewFiltersResponseResultJSON `json:"-"`
}

// zoneFilterFiltersNewFiltersResponseResultJSON contains the JSON metadata for the
// struct [ZoneFilterFiltersNewFiltersResponseResult]
type zoneFilterFiltersNewFiltersResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterFiltersNewFiltersResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterFiltersNewFiltersResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                           `json:"total_count"`
	JSON       zoneFilterFiltersNewFiltersResponseResultInfoJSON `json:"-"`
}

// zoneFilterFiltersNewFiltersResponseResultInfoJSON contains the JSON metadata for
// the struct [ZoneFilterFiltersNewFiltersResponseResultInfo]
type zoneFilterFiltersNewFiltersResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterFiltersNewFiltersResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFilterFiltersNewFiltersResponseSuccess bool

const (
	ZoneFilterFiltersNewFiltersResponseSuccessTrue ZoneFilterFiltersNewFiltersResponseSuccess = true
)

type ZoneFilterFiltersListFiltersResponse struct {
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
	Ref  string                                   `json:"ref"`
	JSON zoneFilterFiltersListFiltersResponseJSON `json:"-"`
}

// zoneFilterFiltersListFiltersResponseJSON contains the JSON metadata for the
// struct [ZoneFilterFiltersListFiltersResponse]
type zoneFilterFiltersListFiltersResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterFiltersListFiltersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterFiltersUpdateFiltersResponse struct {
	Errors     []ZoneFilterFiltersUpdateFiltersResponseError    `json:"errors"`
	Messages   []ZoneFilterFiltersUpdateFiltersResponseMessage  `json:"messages"`
	Result     []ZoneFilterFiltersUpdateFiltersResponseResult   `json:"result"`
	ResultInfo ZoneFilterFiltersUpdateFiltersResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneFilterFiltersUpdateFiltersResponseSuccess `json:"success"`
	JSON    zoneFilterFiltersUpdateFiltersResponseJSON    `json:"-"`
}

// zoneFilterFiltersUpdateFiltersResponseJSON contains the JSON metadata for the
// struct [ZoneFilterFiltersUpdateFiltersResponse]
type zoneFilterFiltersUpdateFiltersResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterFiltersUpdateFiltersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterFiltersUpdateFiltersResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneFilterFiltersUpdateFiltersResponseErrorJSON `json:"-"`
}

// zoneFilterFiltersUpdateFiltersResponseErrorJSON contains the JSON metadata for
// the struct [ZoneFilterFiltersUpdateFiltersResponseError]
type zoneFilterFiltersUpdateFiltersResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterFiltersUpdateFiltersResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterFiltersUpdateFiltersResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zoneFilterFiltersUpdateFiltersResponseMessageJSON `json:"-"`
}

// zoneFilterFiltersUpdateFiltersResponseMessageJSON contains the JSON metadata for
// the struct [ZoneFilterFiltersUpdateFiltersResponseMessage]
type zoneFilterFiltersUpdateFiltersResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterFiltersUpdateFiltersResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterFiltersUpdateFiltersResponseResult struct {
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
	Ref  string                                           `json:"ref"`
	JSON zoneFilterFiltersUpdateFiltersResponseResultJSON `json:"-"`
}

// zoneFilterFiltersUpdateFiltersResponseResultJSON contains the JSON metadata for
// the struct [ZoneFilterFiltersUpdateFiltersResponseResult]
type zoneFilterFiltersUpdateFiltersResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterFiltersUpdateFiltersResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFilterFiltersUpdateFiltersResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       zoneFilterFiltersUpdateFiltersResponseResultInfoJSON `json:"-"`
}

// zoneFilterFiltersUpdateFiltersResponseResultInfoJSON contains the JSON metadata
// for the struct [ZoneFilterFiltersUpdateFiltersResponseResultInfo]
type zoneFilterFiltersUpdateFiltersResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFilterFiltersUpdateFiltersResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFilterFiltersUpdateFiltersResponseSuccess bool

const (
	ZoneFilterFiltersUpdateFiltersResponseSuccessTrue ZoneFilterFiltersUpdateFiltersResponseSuccess = true
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

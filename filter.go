// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
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

// Deletes one or more existing filters.
func (r *FilterService) Delete(ctx context.Context, zoneIdentifier string, body FilterDeleteParams, opts ...option.RequestOption) (res *FilterDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/filters", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type FilterDeleteResponse struct {
	Errors     []FilterDeleteResponseError    `json:"errors"`
	Messages   []FilterDeleteResponseMessage  `json:"messages"`
	Result     []FilterDeleteResponseResult   `json:"result"`
	ResultInfo FilterDeleteResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success FilterDeleteResponseSuccess `json:"success"`
	JSON    filterDeleteResponseJSON    `json:"-"`
}

// filterDeleteResponseJSON contains the JSON metadata for the struct
// [FilterDeleteResponse]
type filterDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterDeleteResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    filterDeleteResponseErrorJSON `json:"-"`
}

// filterDeleteResponseErrorJSON contains the JSON metadata for the struct
// [FilterDeleteResponseError]
type filterDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterDeleteResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    filterDeleteResponseMessageJSON `json:"-"`
}

// filterDeleteResponseMessageJSON contains the JSON metadata for the struct
// [FilterDeleteResponseMessage]
type filterDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterDeleteResponseResult struct {
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
	Ref  string                         `json:"ref"`
	JSON filterDeleteResponseResultJSON `json:"-"`
}

// filterDeleteResponseResultJSON contains the JSON metadata for the struct
// [FilterDeleteResponseResult]
type filterDeleteResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FilterDeleteResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                            `json:"total_count"`
	JSON       filterDeleteResponseResultInfoJSON `json:"-"`
}

// filterDeleteResponseResultInfoJSON contains the JSON metadata for the struct
// [FilterDeleteResponseResultInfo]
type filterDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FilterDeleteResponseSuccess bool

const (
	FilterDeleteResponseSuccessTrue FilterDeleteResponseSuccess = true
)

type FilterDeleteParams struct {
}

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

// PageShieldConnectionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPageShieldConnectionService]
// method instead.
type PageShieldConnectionService struct {
	Options []option.RequestOption
}

// NewPageShieldConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageShieldConnectionService(opts ...option.RequestOption) (r *PageShieldConnectionService) {
	r = &PageShieldConnectionService{}
	r.Options = opts
	return
}

// Lists all connections detected by Page Shield.
func (r *PageShieldConnectionService) List(ctx context.Context, params PageShieldConnectionListParams, opts ...option.RequestOption) (res *[]PageShieldConnectionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageShieldConnectionListResponseEnvelope
	path := fmt.Sprintf("zones/%s/page_shield/connections", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a connection detected by Page Shield by connection ID.
func (r *PageShieldConnectionService) Get(ctx context.Context, connectionID string, query PageShieldConnectionGetParams, opts ...option.RequestOption) (res *PageShieldConnectionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/connections/%s", query.ZoneID, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PageShieldConnectionListResponse struct {
	ID                      interface{}                          `json:"id"`
	AddedAt                 interface{}                          `json:"added_at"`
	DomainReportedMalicious interface{}                          `json:"domain_reported_malicious"`
	FirstPageURL            interface{}                          `json:"first_page_url"`
	FirstSeenAt             interface{}                          `json:"first_seen_at"`
	Host                    interface{}                          `json:"host"`
	LastSeenAt              interface{}                          `json:"last_seen_at"`
	PageURLs                interface{}                          `json:"page_urls"`
	URL                     interface{}                          `json:"url"`
	URLContainsCdnCgiPath   interface{}                          `json:"url_contains_cdn_cgi_path"`
	JSON                    pageShieldConnectionListResponseJSON `json:"-"`
}

// pageShieldConnectionListResponseJSON contains the JSON metadata for the struct
// [PageShieldConnectionListResponse]
type pageShieldConnectionListResponseJSON struct {
	ID                      apijson.Field
	AddedAt                 apijson.Field
	DomainReportedMalicious apijson.Field
	FirstPageURL            apijson.Field
	FirstSeenAt             apijson.Field
	Host                    apijson.Field
	LastSeenAt              apijson.Field
	PageURLs                apijson.Field
	URL                     apijson.Field
	URLContainsCdnCgiPath   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PageShieldConnectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldConnectionListResponseJSON) RawJSON() string {
	return r.raw
}

type PageShieldConnectionGetResponse struct {
	ID                      interface{}                         `json:"id"`
	AddedAt                 interface{}                         `json:"added_at"`
	DomainReportedMalicious interface{}                         `json:"domain_reported_malicious"`
	FirstPageURL            interface{}                         `json:"first_page_url"`
	FirstSeenAt             interface{}                         `json:"first_seen_at"`
	Host                    interface{}                         `json:"host"`
	LastSeenAt              interface{}                         `json:"last_seen_at"`
	PageURLs                interface{}                         `json:"page_urls"`
	URL                     interface{}                         `json:"url"`
	URLContainsCdnCgiPath   interface{}                         `json:"url_contains_cdn_cgi_path"`
	JSON                    pageShieldConnectionGetResponseJSON `json:"-"`
}

// pageShieldConnectionGetResponseJSON contains the JSON metadata for the struct
// [PageShieldConnectionGetResponse]
type pageShieldConnectionGetResponseJSON struct {
	ID                      apijson.Field
	AddedAt                 apijson.Field
	DomainReportedMalicious apijson.Field
	FirstPageURL            apijson.Field
	FirstSeenAt             apijson.Field
	Host                    apijson.Field
	LastSeenAt              apijson.Field
	PageURLs                apijson.Field
	URL                     apijson.Field
	URLContainsCdnCgiPath   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PageShieldConnectionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldConnectionGetResponseJSON) RawJSON() string {
	return r.raw
}

type PageShieldConnectionListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The direction used to sort returned connections.
	Direction param.Field[PageShieldConnectionListParamsDirection] `query:"direction"`
	// When true, excludes connections seen in a `/cdn-cgi` path from the returned
	// connections. The default value is true.
	ExcludeCdnCgi param.Field[bool] `query:"exclude_cdn_cgi"`
	// Excludes connections whose URL contains one of the URL-encoded URLs separated by
	// commas.
	ExcludeURLs param.Field[string] `query:"exclude_urls"`
	// Export the list of connections as a file. Cannot be used with per_page or page
	// options.
	Export param.Field[PageShieldConnectionListParamsExport] `query:"export"`
	// Includes connections that match one or more URL-encoded hostnames separated by
	// commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	Hosts param.Field[string] `query:"hosts"`
	// The field used to sort returned connections.
	OrderBy param.Field[PageShieldConnectionListParamsOrderBy] `query:"order_by"`
	// The current page number of the paginated results.
	//
	// We additionally support a special value "all". When "all" is used, the API will
	// return all the connections with the applied filters in a single page.
	// Additionally, when using this value, the API will not return the categorisation
	// data for the URL and domain of the connections. This feature is best-effort and
	// it may only work for zones with a low number of connections
	Page param.Field[string] `query:"page"`
	// Includes connections that match one or more page URLs (separated by commas)
	// where they were last seen
	//
	// Wildcards are supported at the start and end of each page URL to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	PageURL param.Field[string] `query:"page_url"`
	// The number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// When true, malicious connections appear first in the returned connections.
	PrioritizeMalicious param.Field[bool] `query:"prioritize_malicious"`
	// Filters the returned connections using a comma-separated list of connection
	// statuses. Accepted values: `active`, `infrequent`, and `inactive`. The default
	// value is `active`.
	Status param.Field[string] `query:"status"`
	// Includes connections whose URL contain one or more URL-encoded URLs separated by
	// commas.
	URLs param.Field[string] `query:"urls"`
}

// URLQuery serializes [PageShieldConnectionListParams]'s query parameters as
// `url.Values`.
func (r PageShieldConnectionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned connections.
type PageShieldConnectionListParamsDirection string

const (
	PageShieldConnectionListParamsDirectionAsc  PageShieldConnectionListParamsDirection = "asc"
	PageShieldConnectionListParamsDirectionDesc PageShieldConnectionListParamsDirection = "desc"
)

// Export the list of connections as a file. Cannot be used with per_page or page
// options.
type PageShieldConnectionListParamsExport string

const (
	PageShieldConnectionListParamsExportCsv PageShieldConnectionListParamsExport = "csv"
)

// The field used to sort returned connections.
type PageShieldConnectionListParamsOrderBy string

const (
	PageShieldConnectionListParamsOrderByFirstSeenAt PageShieldConnectionListParamsOrderBy = "first_seen_at"
	PageShieldConnectionListParamsOrderByLastSeenAt  PageShieldConnectionListParamsOrderBy = "last_seen_at"
)

type PageShieldConnectionListResponseEnvelope struct {
	Errors   []PageShieldConnectionListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageShieldConnectionListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PageShieldConnectionListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PageShieldConnectionListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PageShieldConnectionListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pageShieldConnectionListResponseEnvelopeJSON       `json:"-"`
}

// pageShieldConnectionListResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageShieldConnectionListResponseEnvelope]
type pageShieldConnectionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldConnectionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldConnectionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PageShieldConnectionListResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    pageShieldConnectionListResponseEnvelopeErrorsJSON `json:"-"`
}

// pageShieldConnectionListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [PageShieldConnectionListResponseEnvelopeErrors]
type pageShieldConnectionListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldConnectionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldConnectionListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageShieldConnectionListResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    pageShieldConnectionListResponseEnvelopeMessagesJSON `json:"-"`
}

// pageShieldConnectionListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [PageShieldConnectionListResponseEnvelopeMessages]
type pageShieldConnectionListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldConnectionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldConnectionListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageShieldConnectionListResponseEnvelopeSuccess bool

const (
	PageShieldConnectionListResponseEnvelopeSuccessTrue PageShieldConnectionListResponseEnvelopeSuccess = true
)

type PageShieldConnectionListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       pageShieldConnectionListResponseEnvelopeResultInfoJSON `json:"-"`
}

// pageShieldConnectionListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [PageShieldConnectionListResponseEnvelopeResultInfo]
type pageShieldConnectionListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldConnectionListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldConnectionListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type PageShieldConnectionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

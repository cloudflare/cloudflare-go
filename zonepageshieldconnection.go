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

// ZonePageShieldConnectionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZonePageShieldConnectionService] method instead.
type ZonePageShieldConnectionService struct {
	Options []option.RequestOption
}

// NewZonePageShieldConnectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZonePageShieldConnectionService(opts ...option.RequestOption) (r *ZonePageShieldConnectionService) {
	r = &ZonePageShieldConnectionService{}
	r.Options = opts
	return
}

// Fetches a connection detected by Page Shield by connection ID.
func (r *ZonePageShieldConnectionService) Get(ctx context.Context, zoneID string, connectionID string, opts ...option.RequestOption) (res *ZonePageShieldConnectionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/connections/%s", zoneID, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all connections detected by Page Shield.
func (r *ZonePageShieldConnectionService) PageShieldListPageShieldConnections(ctx context.Context, zoneID string, query ZonePageShieldConnectionPageShieldListPageShieldConnectionsParams, opts ...option.RequestOption) (res *ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/connections", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZonePageShieldConnectionGetResponse struct {
	ID                      interface{}                             `json:"id"`
	AddedAt                 interface{}                             `json:"added_at"`
	DomainReportedMalicious interface{}                             `json:"domain_reported_malicious"`
	FirstPageURL            interface{}                             `json:"first_page_url"`
	FirstSeenAt             interface{}                             `json:"first_seen_at"`
	Host                    interface{}                             `json:"host"`
	LastSeenAt              interface{}                             `json:"last_seen_at"`
	PageURLs                interface{}                             `json:"page_urls"`
	URL                     interface{}                             `json:"url"`
	URLContainsCdnCgiPath   interface{}                             `json:"url_contains_cdn_cgi_path"`
	JSON                    zonePageShieldConnectionGetResponseJSON `json:"-"`
}

// zonePageShieldConnectionGetResponseJSON contains the JSON metadata for the
// struct [ZonePageShieldConnectionGetResponse]
type zonePageShieldConnectionGetResponseJSON struct {
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

func (r *ZonePageShieldConnectionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponse struct {
	Errors     []ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseError    `json:"errors"`
	Messages   []ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseMessage  `json:"messages"`
	Result     []ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseResult   `json:"result"`
	ResultInfo ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseSuccess `json:"success"`
	JSON    zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseJSON    `json:"-"`
}

// zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseJSON contains
// the JSON metadata for the struct
// [ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponse]
type zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseError struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseErrorJSON `json:"-"`
}

// zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseError]
type zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseMessage struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseMessageJSON `json:"-"`
}

// zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseMessage]
type zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseResult struct {
	ID                      interface{}                                                                   `json:"id"`
	AddedAt                 interface{}                                                                   `json:"added_at"`
	DomainReportedMalicious interface{}                                                                   `json:"domain_reported_malicious"`
	FirstPageURL            interface{}                                                                   `json:"first_page_url"`
	FirstSeenAt             interface{}                                                                   `json:"first_seen_at"`
	Host                    interface{}                                                                   `json:"host"`
	LastSeenAt              interface{}                                                                   `json:"last_seen_at"`
	PageURLs                interface{}                                                                   `json:"page_urls"`
	URL                     interface{}                                                                   `json:"url"`
	URLContainsCdnCgiPath   interface{}                                                                   `json:"url_contains_cdn_cgi_path"`
	JSON                    zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseResultJSON `json:"-"`
}

// zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseResultJSON
// contains the JSON metadata for the struct
// [ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseResult]
type zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseResultJSON struct {
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

func (r *ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                           `json:"total_count"`
	JSON       zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseResultInfoJSON `json:"-"`
}

// zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseResultInfo]
type zonePageShieldConnectionPageShieldListPageShieldConnectionsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseSuccess bool

const (
	ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseSuccessTrue ZonePageShieldConnectionPageShieldListPageShieldConnectionsResponseSuccess = true
)

type ZonePageShieldConnectionPageShieldListPageShieldConnectionsParams struct {
	// The direction used to sort returned connections.
	Direction param.Field[ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsDirection] `query:"direction"`
	// When true, excludes connections seen in a `/cdn-cgi` path from the returned
	// connections. The default value is true.
	ExcludeCdnCgi param.Field[bool] `query:"exclude_cdn_cgi"`
	// Excludes connections whose URL contains one of the URL-encoded URLs separated by
	// commas.
	ExcludeURLs param.Field[string] `query:"exclude_urls"`
	// Export the list of connections as a file. Cannot be used with per_page or page
	// options.
	Export param.Field[ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsExport] `query:"export"`
	// Includes connections that match one or more URL-encoded hostnames separated by
	// commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	Hosts param.Field[string] `query:"hosts"`
	// The field used to sort returned connections.
	OrderBy param.Field[ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsOrderBy] `query:"order_by"`
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

// URLQuery serializes
// [ZonePageShieldConnectionPageShieldListPageShieldConnectionsParams]'s query
// parameters as `url.Values`.
func (r ZonePageShieldConnectionPageShieldListPageShieldConnectionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned connections.
type ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsDirection string

const (
	ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsDirectionAsc  ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsDirection = "asc"
	ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsDirectionDesc ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsDirection = "desc"
)

// Export the list of connections as a file. Cannot be used with per_page or page
// options.
type ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsExport string

const (
	ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsExportCsv ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsExport = "csv"
)

// The field used to sort returned connections.
type ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsOrderBy string

const (
	ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsOrderByFirstSeenAt ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsOrderBy = "first_seen_at"
	ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsOrderByLastSeenAt  ZonePageShieldConnectionPageShieldListPageShieldConnectionsParamsOrderBy = "last_seen_at"
)

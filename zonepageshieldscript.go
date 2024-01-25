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

// ZonePageShieldScriptService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZonePageShieldScriptService]
// method instead.
type ZonePageShieldScriptService struct {
	Options []option.RequestOption
}

// NewZonePageShieldScriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZonePageShieldScriptService(opts ...option.RequestOption) (r *ZonePageShieldScriptService) {
	r = &ZonePageShieldScriptService{}
	r.Options = opts
	return
}

// Fetches a script detected by Page Shield by script ID.
func (r *ZonePageShieldScriptService) Get(ctx context.Context, zoneID string, scriptID string, opts ...option.RequestOption) (res *ZonePageShieldScriptGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/scripts/%s", zoneID, scriptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all scripts detected by Page Shield.
func (r *ZonePageShieldScriptService) PageShieldListPageShieldScripts(ctx context.Context, zoneID string, query ZonePageShieldScriptPageShieldListPageShieldScriptsParams, opts ...option.RequestOption) (res *ZonePageShieldScriptPageShieldListPageShieldScriptsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/scripts", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZonePageShieldScriptGetResponse struct {
	ID                      interface{}                              `json:"id"`
	AddedAt                 interface{}                              `json:"added_at"`
	DomainReportedMalicious interface{}                              `json:"domain_reported_malicious"`
	FetchedAt               interface{}                              `json:"fetched_at"`
	FirstPageURL            interface{}                              `json:"first_page_url"`
	FirstSeenAt             interface{}                              `json:"first_seen_at"`
	Hash                    interface{}                              `json:"hash"`
	Host                    interface{}                              `json:"host"`
	JsIntegrityScore        interface{}                              `json:"js_integrity_score"`
	LastSeenAt              interface{}                              `json:"last_seen_at"`
	PageURLs                interface{}                              `json:"page_urls"`
	URL                     interface{}                              `json:"url"`
	URLContainsCdnCgiPath   interface{}                              `json:"url_contains_cdn_cgi_path"`
	Versions                []ZonePageShieldScriptGetResponseVersion `json:"versions,nullable"`
	JSON                    zonePageShieldScriptGetResponseJSON      `json:"-"`
}

// zonePageShieldScriptGetResponseJSON contains the JSON metadata for the struct
// [ZonePageShieldScriptGetResponse]
type zonePageShieldScriptGetResponseJSON struct {
	ID                      apijson.Field
	AddedAt                 apijson.Field
	DomainReportedMalicious apijson.Field
	FetchedAt               apijson.Field
	FirstPageURL            apijson.Field
	FirstSeenAt             apijson.Field
	Hash                    apijson.Field
	Host                    apijson.Field
	JsIntegrityScore        apijson.Field
	LastSeenAt              apijson.Field
	PageURLs                apijson.Field
	URL                     apijson.Field
	URLContainsCdnCgiPath   apijson.Field
	Versions                apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ZonePageShieldScriptGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The version of the analyzed script.
type ZonePageShieldScriptGetResponseVersion struct {
	// The timestamp of when the script was last fetched.
	FetchedAt string `json:"fetched_at,nullable"`
	// The computed hash of the analyzed script.
	Hash string `json:"hash,nullable"`
	// The integrity score of the JavaScript content.
	JsIntegrityScore int64                                      `json:"js_integrity_score,nullable"`
	JSON             zonePageShieldScriptGetResponseVersionJSON `json:"-"`
}

// zonePageShieldScriptGetResponseVersionJSON contains the JSON metadata for the
// struct [ZonePageShieldScriptGetResponseVersion]
type zonePageShieldScriptGetResponseVersionJSON struct {
	FetchedAt        apijson.Field
	Hash             apijson.Field
	JsIntegrityScore apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZonePageShieldScriptGetResponseVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldScriptPageShieldListPageShieldScriptsResponse struct {
	Errors     []ZonePageShieldScriptPageShieldListPageShieldScriptsResponseError    `json:"errors"`
	Messages   []ZonePageShieldScriptPageShieldListPageShieldScriptsResponseMessage  `json:"messages"`
	Result     []ZonePageShieldScriptPageShieldListPageShieldScriptsResponseResult   `json:"result"`
	ResultInfo ZonePageShieldScriptPageShieldListPageShieldScriptsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZonePageShieldScriptPageShieldListPageShieldScriptsResponseSuccess `json:"success"`
	JSON    zonePageShieldScriptPageShieldListPageShieldScriptsResponseJSON    `json:"-"`
}

// zonePageShieldScriptPageShieldListPageShieldScriptsResponseJSON contains the
// JSON metadata for the struct
// [ZonePageShieldScriptPageShieldListPageShieldScriptsResponse]
type zonePageShieldScriptPageShieldListPageShieldScriptsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldScriptPageShieldListPageShieldScriptsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldScriptPageShieldListPageShieldScriptsResponseError struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    zonePageShieldScriptPageShieldListPageShieldScriptsResponseErrorJSON `json:"-"`
}

// zonePageShieldScriptPageShieldListPageShieldScriptsResponseErrorJSON contains
// the JSON metadata for the struct
// [ZonePageShieldScriptPageShieldListPageShieldScriptsResponseError]
type zonePageShieldScriptPageShieldListPageShieldScriptsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldScriptPageShieldListPageShieldScriptsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldScriptPageShieldListPageShieldScriptsResponseMessage struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    zonePageShieldScriptPageShieldListPageShieldScriptsResponseMessageJSON `json:"-"`
}

// zonePageShieldScriptPageShieldListPageShieldScriptsResponseMessageJSON contains
// the JSON metadata for the struct
// [ZonePageShieldScriptPageShieldListPageShieldScriptsResponseMessage]
type zonePageShieldScriptPageShieldListPageShieldScriptsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldScriptPageShieldListPageShieldScriptsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldScriptPageShieldListPageShieldScriptsResponseResult struct {
	ID                      interface{}                                                           `json:"id"`
	AddedAt                 interface{}                                                           `json:"added_at"`
	DomainReportedMalicious interface{}                                                           `json:"domain_reported_malicious"`
	FetchedAt               interface{}                                                           `json:"fetched_at"`
	FirstPageURL            interface{}                                                           `json:"first_page_url"`
	FirstSeenAt             interface{}                                                           `json:"first_seen_at"`
	Hash                    interface{}                                                           `json:"hash"`
	Host                    interface{}                                                           `json:"host"`
	JsIntegrityScore        interface{}                                                           `json:"js_integrity_score"`
	LastSeenAt              interface{}                                                           `json:"last_seen_at"`
	PageURLs                interface{}                                                           `json:"page_urls"`
	URL                     interface{}                                                           `json:"url"`
	URLContainsCdnCgiPath   interface{}                                                           `json:"url_contains_cdn_cgi_path"`
	JSON                    zonePageShieldScriptPageShieldListPageShieldScriptsResponseResultJSON `json:"-"`
}

// zonePageShieldScriptPageShieldListPageShieldScriptsResponseResultJSON contains
// the JSON metadata for the struct
// [ZonePageShieldScriptPageShieldListPageShieldScriptsResponseResult]
type zonePageShieldScriptPageShieldListPageShieldScriptsResponseResultJSON struct {
	ID                      apijson.Field
	AddedAt                 apijson.Field
	DomainReportedMalicious apijson.Field
	FetchedAt               apijson.Field
	FirstPageURL            apijson.Field
	FirstSeenAt             apijson.Field
	Hash                    apijson.Field
	Host                    apijson.Field
	JsIntegrityScore        apijson.Field
	LastSeenAt              apijson.Field
	PageURLs                apijson.Field
	URL                     apijson.Field
	URLContainsCdnCgiPath   apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *ZonePageShieldScriptPageShieldListPageShieldScriptsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageShieldScriptPageShieldListPageShieldScriptsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                   `json:"total_count"`
	JSON       zonePageShieldScriptPageShieldListPageShieldScriptsResponseResultInfoJSON `json:"-"`
}

// zonePageShieldScriptPageShieldListPageShieldScriptsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZonePageShieldScriptPageShieldListPageShieldScriptsResponseResultInfo]
type zonePageShieldScriptPageShieldListPageShieldScriptsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageShieldScriptPageShieldListPageShieldScriptsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZonePageShieldScriptPageShieldListPageShieldScriptsResponseSuccess bool

const (
	ZonePageShieldScriptPageShieldListPageShieldScriptsResponseSuccessTrue ZonePageShieldScriptPageShieldListPageShieldScriptsResponseSuccess = true
)

type ZonePageShieldScriptPageShieldListPageShieldScriptsParams struct {
	// The direction used to sort returned scripts.
	Direction param.Field[ZonePageShieldScriptPageShieldListPageShieldScriptsParamsDirection] `query:"direction"`
	// When true, excludes scripts seen in a `/cdn-cgi` path from the returned scripts.
	// The default value is true.
	ExcludeCdnCgi param.Field[bool] `query:"exclude_cdn_cgi"`
	// When true, excludes duplicate scripts. We consider a script duplicate of another
	// if their javascript content matches and they share the same url host and zone
	// hostname. In such case, we return the most recent script for the URL host and
	// zone hostname combination.
	ExcludeDuplicates param.Field[bool] `query:"exclude_duplicates"`
	// Excludes scripts whose URL contains one of the URL-encoded URLs separated by
	// commas.
	ExcludeURLs param.Field[string] `query:"exclude_urls"`
	// Export the list of scripts as a file. Cannot be used with per_page or page
	// options.
	Export param.Field[ZonePageShieldScriptPageShieldListPageShieldScriptsParamsExport] `query:"export"`
	// Includes scripts that match one or more URL-encoded hostnames separated by
	// commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	Hosts param.Field[string] `query:"hosts"`
	// The field used to sort returned scripts.
	OrderBy param.Field[ZonePageShieldScriptPageShieldListPageShieldScriptsParamsOrderBy] `query:"order_by"`
	// The current page number of the paginated results.
	//
	// We additionally support a special value "all". When "all" is used, the API will
	// return all the scripts with the applied filters in a single page. Additionally,
	// when using this value, the API will not return the script versions or
	// categorisation data for the URL and domain of the scripts. This feature is
	// best-effort and it may only work for zones with a low number of scripts
	Page param.Field[string] `query:"page"`
	// Includes scripts that match one or more page URLs (separated by commas) where
	// they were last seen
	//
	// Wildcards are supported at the start and end of each page URL to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	PageURL param.Field[string] `query:"page_url"`
	// The number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// When true, malicious scripts appear first in the returned scripts.
	PrioritizeMalicious param.Field[bool] `query:"prioritize_malicious"`
	// Filters the returned scripts using a comma-separated list of scripts statuses.
	// Accepted values: `active`, `infrequent`, and `inactive`. The default value is
	// `active`.
	Status param.Field[string] `query:"status"`
	// Includes scripts whose URL contain one or more URL-encoded URLs separated by
	// commas.
	URLs param.Field[string] `query:"urls"`
}

// URLQuery serializes
// [ZonePageShieldScriptPageShieldListPageShieldScriptsParams]'s query parameters
// as `url.Values`.
func (r ZonePageShieldScriptPageShieldListPageShieldScriptsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned scripts.
type ZonePageShieldScriptPageShieldListPageShieldScriptsParamsDirection string

const (
	ZonePageShieldScriptPageShieldListPageShieldScriptsParamsDirectionAsc  ZonePageShieldScriptPageShieldListPageShieldScriptsParamsDirection = "asc"
	ZonePageShieldScriptPageShieldListPageShieldScriptsParamsDirectionDesc ZonePageShieldScriptPageShieldListPageShieldScriptsParamsDirection = "desc"
)

// Export the list of scripts as a file. Cannot be used with per_page or page
// options.
type ZonePageShieldScriptPageShieldListPageShieldScriptsParamsExport string

const (
	ZonePageShieldScriptPageShieldListPageShieldScriptsParamsExportCsv ZonePageShieldScriptPageShieldListPageShieldScriptsParamsExport = "csv"
)

// The field used to sort returned scripts.
type ZonePageShieldScriptPageShieldListPageShieldScriptsParamsOrderBy string

const (
	ZonePageShieldScriptPageShieldListPageShieldScriptsParamsOrderByFirstSeenAt ZonePageShieldScriptPageShieldListPageShieldScriptsParamsOrderBy = "first_seen_at"
	ZonePageShieldScriptPageShieldListPageShieldScriptsParamsOrderByLastSeenAt  ZonePageShieldScriptPageShieldListPageShieldScriptsParamsOrderBy = "last_seen_at"
)

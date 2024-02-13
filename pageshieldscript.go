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

// PageShieldScriptService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPageShieldScriptService] method
// instead.
type PageShieldScriptService struct {
	Options []option.RequestOption
}

// NewPageShieldScriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPageShieldScriptService(opts ...option.RequestOption) (r *PageShieldScriptService) {
	r = &PageShieldScriptService{}
	r.Options = opts
	return
}

// Fetches a script detected by Page Shield by script ID.
func (r *PageShieldScriptService) Get(ctx context.Context, zoneID string, scriptID string, opts ...option.RequestOption) (res *PageShieldScriptGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/scripts/%s", zoneID, scriptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all scripts detected by Page Shield.
func (r *PageShieldScriptService) PageShieldListPageShieldScripts(ctx context.Context, zoneID string, query PageShieldScriptPageShieldListPageShieldScriptsParams, opts ...option.RequestOption) (res *[]PageShieldScriptPageShieldListPageShieldScriptsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelope
	path := fmt.Sprintf("zones/%s/page_shield/scripts", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PageShieldScriptGetResponse struct {
	ID                      interface{}                          `json:"id"`
	AddedAt                 interface{}                          `json:"added_at"`
	DomainReportedMalicious interface{}                          `json:"domain_reported_malicious"`
	FetchedAt               interface{}                          `json:"fetched_at"`
	FirstPageURL            interface{}                          `json:"first_page_url"`
	FirstSeenAt             interface{}                          `json:"first_seen_at"`
	Hash                    interface{}                          `json:"hash"`
	Host                    interface{}                          `json:"host"`
	JsIntegrityScore        interface{}                          `json:"js_integrity_score"`
	LastSeenAt              interface{}                          `json:"last_seen_at"`
	PageURLs                interface{}                          `json:"page_urls"`
	URL                     interface{}                          `json:"url"`
	URLContainsCdnCgiPath   interface{}                          `json:"url_contains_cdn_cgi_path"`
	Versions                []PageShieldScriptGetResponseVersion `json:"versions,nullable"`
	JSON                    pageShieldScriptGetResponseJSON      `json:"-"`
}

// pageShieldScriptGetResponseJSON contains the JSON metadata for the struct
// [PageShieldScriptGetResponse]
type pageShieldScriptGetResponseJSON struct {
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

func (r *PageShieldScriptGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The version of the analyzed script.
type PageShieldScriptGetResponseVersion struct {
	// The timestamp of when the script was last fetched.
	FetchedAt string `json:"fetched_at,nullable"`
	// The computed hash of the analyzed script.
	Hash string `json:"hash,nullable"`
	// The integrity score of the JavaScript content.
	JsIntegrityScore int64                                  `json:"js_integrity_score,nullable"`
	JSON             pageShieldScriptGetResponseVersionJSON `json:"-"`
}

// pageShieldScriptGetResponseVersionJSON contains the JSON metadata for the struct
// [PageShieldScriptGetResponseVersion]
type pageShieldScriptGetResponseVersionJSON struct {
	FetchedAt        apijson.Field
	Hash             apijson.Field
	JsIntegrityScore apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PageShieldScriptGetResponseVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageShieldScriptPageShieldListPageShieldScriptsResponse struct {
	ID                      interface{}                                                 `json:"id"`
	AddedAt                 interface{}                                                 `json:"added_at"`
	DomainReportedMalicious interface{}                                                 `json:"domain_reported_malicious"`
	FetchedAt               interface{}                                                 `json:"fetched_at"`
	FirstPageURL            interface{}                                                 `json:"first_page_url"`
	FirstSeenAt             interface{}                                                 `json:"first_seen_at"`
	Hash                    interface{}                                                 `json:"hash"`
	Host                    interface{}                                                 `json:"host"`
	JsIntegrityScore        interface{}                                                 `json:"js_integrity_score"`
	LastSeenAt              interface{}                                                 `json:"last_seen_at"`
	PageURLs                interface{}                                                 `json:"page_urls"`
	URL                     interface{}                                                 `json:"url"`
	URLContainsCdnCgiPath   interface{}                                                 `json:"url_contains_cdn_cgi_path"`
	JSON                    pageShieldScriptPageShieldListPageShieldScriptsResponseJSON `json:"-"`
}

// pageShieldScriptPageShieldListPageShieldScriptsResponseJSON contains the JSON
// metadata for the struct
// [PageShieldScriptPageShieldListPageShieldScriptsResponse]
type pageShieldScriptPageShieldListPageShieldScriptsResponseJSON struct {
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

func (r *PageShieldScriptPageShieldListPageShieldScriptsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageShieldScriptPageShieldListPageShieldScriptsParams struct {
	// The direction used to sort returned scripts.
	Direction param.Field[PageShieldScriptPageShieldListPageShieldScriptsParamsDirection] `query:"direction"`
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
	Export param.Field[PageShieldScriptPageShieldListPageShieldScriptsParamsExport] `query:"export"`
	// Includes scripts that match one or more URL-encoded hostnames separated by
	// commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	Hosts param.Field[string] `query:"hosts"`
	// The field used to sort returned scripts.
	OrderBy param.Field[PageShieldScriptPageShieldListPageShieldScriptsParamsOrderBy] `query:"order_by"`
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

// URLQuery serializes [PageShieldScriptPageShieldListPageShieldScriptsParams]'s
// query parameters as `url.Values`.
func (r PageShieldScriptPageShieldListPageShieldScriptsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned scripts.
type PageShieldScriptPageShieldListPageShieldScriptsParamsDirection string

const (
	PageShieldScriptPageShieldListPageShieldScriptsParamsDirectionAsc  PageShieldScriptPageShieldListPageShieldScriptsParamsDirection = "asc"
	PageShieldScriptPageShieldListPageShieldScriptsParamsDirectionDesc PageShieldScriptPageShieldListPageShieldScriptsParamsDirection = "desc"
)

// Export the list of scripts as a file. Cannot be used with per_page or page
// options.
type PageShieldScriptPageShieldListPageShieldScriptsParamsExport string

const (
	PageShieldScriptPageShieldListPageShieldScriptsParamsExportCsv PageShieldScriptPageShieldListPageShieldScriptsParamsExport = "csv"
)

// The field used to sort returned scripts.
type PageShieldScriptPageShieldListPageShieldScriptsParamsOrderBy string

const (
	PageShieldScriptPageShieldListPageShieldScriptsParamsOrderByFirstSeenAt PageShieldScriptPageShieldListPageShieldScriptsParamsOrderBy = "first_seen_at"
	PageShieldScriptPageShieldListPageShieldScriptsParamsOrderByLastSeenAt  PageShieldScriptPageShieldListPageShieldScriptsParamsOrderBy = "last_seen_at"
)

type PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelope struct {
	Errors   []PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeMessages `json:"messages,required"`
	Result   []PageShieldScriptPageShieldListPageShieldScriptsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeJSON       `json:"-"`
}

// pageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelope]
type pageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeErrors struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    pageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeErrorsJSON `json:"-"`
}

// pageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeErrors]
type pageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeMessages struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    pageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeMessagesJSON `json:"-"`
}

// pageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeMessages]
type pageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeSuccess bool

const (
	PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeSuccessTrue PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeSuccess = true
)

type PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                       `json:"total_count"`
	JSON       pageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeResultInfoJSON `json:"-"`
}

// pageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeResultInfo]
type pageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldScriptPageShieldListPageShieldScriptsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

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

// Lists all scripts detected by Page Shield.
func (r *PageShieldScriptService) List(ctx context.Context, params PageShieldScriptListParams, opts ...option.RequestOption) (res *[]PageShieldScriptListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageShieldScriptListResponseEnvelope
	path := fmt.Sprintf("zones/%s/page_shield/scripts", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a script detected by Page Shield by script ID.
func (r *PageShieldScriptService) Get(ctx context.Context, scriptID string, query PageShieldScriptGetParams, opts ...option.RequestOption) (res *PageShieldScriptGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/scripts/%s", query.ZoneID, scriptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PageShieldScriptListResponse struct {
	ID                      interface{}                      `json:"id"`
	AddedAt                 interface{}                      `json:"added_at"`
	DomainReportedMalicious interface{}                      `json:"domain_reported_malicious"`
	FetchedAt               interface{}                      `json:"fetched_at"`
	FirstPageURL            interface{}                      `json:"first_page_url"`
	FirstSeenAt             interface{}                      `json:"first_seen_at"`
	Hash                    interface{}                      `json:"hash"`
	Host                    interface{}                      `json:"host"`
	JsIntegrityScore        interface{}                      `json:"js_integrity_score"`
	LastSeenAt              interface{}                      `json:"last_seen_at"`
	PageURLs                interface{}                      `json:"page_urls"`
	URL                     interface{}                      `json:"url"`
	URLContainsCdnCgiPath   interface{}                      `json:"url_contains_cdn_cgi_path"`
	JSON                    pageShieldScriptListResponseJSON `json:"-"`
}

// pageShieldScriptListResponseJSON contains the JSON metadata for the struct
// [PageShieldScriptListResponse]
type pageShieldScriptListResponseJSON struct {
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

func (r *PageShieldScriptListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldScriptListResponseJSON) RawJSON() string {
	return r.raw
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

func (r pageShieldScriptGetResponseJSON) RawJSON() string {
	return r.raw
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

func (r pageShieldScriptGetResponseVersionJSON) RawJSON() string {
	return r.raw
}

type PageShieldScriptListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The direction used to sort returned scripts.
	Direction param.Field[PageShieldScriptListParamsDirection] `query:"direction"`
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
	Export param.Field[PageShieldScriptListParamsExport] `query:"export"`
	// Includes scripts that match one or more URL-encoded hostnames separated by
	// commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	Hosts param.Field[string] `query:"hosts"`
	// The field used to sort returned scripts.
	OrderBy param.Field[PageShieldScriptListParamsOrderBy] `query:"order_by"`
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

// URLQuery serializes [PageShieldScriptListParams]'s query parameters as
// `url.Values`.
func (r PageShieldScriptListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned scripts.
type PageShieldScriptListParamsDirection string

const (
	PageShieldScriptListParamsDirectionAsc  PageShieldScriptListParamsDirection = "asc"
	PageShieldScriptListParamsDirectionDesc PageShieldScriptListParamsDirection = "desc"
)

// Export the list of scripts as a file. Cannot be used with per_page or page
// options.
type PageShieldScriptListParamsExport string

const (
	PageShieldScriptListParamsExportCsv PageShieldScriptListParamsExport = "csv"
)

// The field used to sort returned scripts.
type PageShieldScriptListParamsOrderBy string

const (
	PageShieldScriptListParamsOrderByFirstSeenAt PageShieldScriptListParamsOrderBy = "first_seen_at"
	PageShieldScriptListParamsOrderByLastSeenAt  PageShieldScriptListParamsOrderBy = "last_seen_at"
)

type PageShieldScriptListResponseEnvelope struct {
	Errors   []PageShieldScriptListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageShieldScriptListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PageShieldScriptListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    PageShieldScriptListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo PageShieldScriptListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       pageShieldScriptListResponseEnvelopeJSON       `json:"-"`
}

// pageShieldScriptListResponseEnvelopeJSON contains the JSON metadata for the
// struct [PageShieldScriptListResponseEnvelope]
type pageShieldScriptListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldScriptListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldScriptListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PageShieldScriptListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    pageShieldScriptListResponseEnvelopeErrorsJSON `json:"-"`
}

// pageShieldScriptListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [PageShieldScriptListResponseEnvelopeErrors]
type pageShieldScriptListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldScriptListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldScriptListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageShieldScriptListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    pageShieldScriptListResponseEnvelopeMessagesJSON `json:"-"`
}

// pageShieldScriptListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PageShieldScriptListResponseEnvelopeMessages]
type pageShieldScriptListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldScriptListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldScriptListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageShieldScriptListResponseEnvelopeSuccess bool

const (
	PageShieldScriptListResponseEnvelopeSuccessTrue PageShieldScriptListResponseEnvelopeSuccess = true
)

type PageShieldScriptListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       pageShieldScriptListResponseEnvelopeResultInfoJSON `json:"-"`
}

// pageShieldScriptListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [PageShieldScriptListResponseEnvelopeResultInfo]
type pageShieldScriptListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageShieldScriptListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldScriptListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type PageShieldScriptGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

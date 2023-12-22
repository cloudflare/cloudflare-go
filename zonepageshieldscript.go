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
func (r *ZonePageShieldScriptService) Get(ctx context.Context, zoneID string, id string, opts ...option.RequestOption) (res *GetZoneScriptResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/scripts/%s", zoneID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all scripts detected by Page Shield.
func (r *ZonePageShieldScriptService) PageShieldListPageShieldScripts(ctx context.Context, zoneID string, query ZonePageShieldScriptPageShieldListPageShieldScriptsParams, opts ...option.RequestOption) (res *ListZoneScriptsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/scripts", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type GetZoneScriptResponse struct {
	ID                      interface{}                    `json:"id"`
	AddedAt                 interface{}                    `json:"added_at"`
	DomainReportedMalicious interface{}                    `json:"domain_reported_malicious"`
	FetchedAt               interface{}                    `json:"fetched_at"`
	FirstPageURL            interface{}                    `json:"first_page_url"`
	FirstSeenAt             interface{}                    `json:"first_seen_at"`
	Hash                    interface{}                    `json:"hash"`
	Host                    interface{}                    `json:"host"`
	JsIntegrityScore        interface{}                    `json:"js_integrity_score"`
	LastSeenAt              interface{}                    `json:"last_seen_at"`
	PageURLs                interface{}                    `json:"page_urls"`
	URL                     interface{}                    `json:"url"`
	URLContainsCdnCgiPath   interface{}                    `json:"url_contains_cdn_cgi_path"`
	Versions                []GetZoneScriptResponseVersion `json:"versions,nullable"`
	JSON                    getZoneScriptResponseJSON      `json:"-"`
}

// getZoneScriptResponseJSON contains the JSON metadata for the struct
// [GetZoneScriptResponse]
type getZoneScriptResponseJSON struct {
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

func (r *GetZoneScriptResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The version of the analyzed script.
type GetZoneScriptResponseVersion struct {
	// The timestamp of when the script was last fetched.
	FetchedAt string `json:"fetched_at,nullable"`
	// The computed hash of the analyzed script.
	Hash string `json:"hash,nullable"`
	// The integrity score of the JavaScript content.
	JsIntegrityScore int64                            `json:"js_integrity_score,nullable"`
	JSON             getZoneScriptResponseVersionJSON `json:"-"`
}

// getZoneScriptResponseVersionJSON contains the JSON metadata for the struct
// [GetZoneScriptResponseVersion]
type getZoneScriptResponseVersionJSON struct {
	FetchedAt        apijson.Field
	Hash             apijson.Field
	JsIntegrityScore apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *GetZoneScriptResponseVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListZoneScriptsResponse struct {
	Errors     []ListZoneScriptsResponseError    `json:"errors"`
	Messages   []ListZoneScriptsResponseMessage  `json:"messages"`
	Result     []ListZoneScriptsResponseResult   `json:"result"`
	ResultInfo ListZoneScriptsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ListZoneScriptsResponseSuccess `json:"success"`
	JSON    listZoneScriptsResponseJSON    `json:"-"`
}

// listZoneScriptsResponseJSON contains the JSON metadata for the struct
// [ListZoneScriptsResponse]
type listZoneScriptsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListZoneScriptsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListZoneScriptsResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    listZoneScriptsResponseErrorJSON `json:"-"`
}

// listZoneScriptsResponseErrorJSON contains the JSON metadata for the struct
// [ListZoneScriptsResponseError]
type listZoneScriptsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListZoneScriptsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListZoneScriptsResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    listZoneScriptsResponseMessageJSON `json:"-"`
}

// listZoneScriptsResponseMessageJSON contains the JSON metadata for the struct
// [ListZoneScriptsResponseMessage]
type listZoneScriptsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListZoneScriptsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListZoneScriptsResponseResult struct {
	ID                      interface{}                       `json:"id"`
	AddedAt                 interface{}                       `json:"added_at"`
	DomainReportedMalicious interface{}                       `json:"domain_reported_malicious"`
	FetchedAt               interface{}                       `json:"fetched_at"`
	FirstPageURL            interface{}                       `json:"first_page_url"`
	FirstSeenAt             interface{}                       `json:"first_seen_at"`
	Hash                    interface{}                       `json:"hash"`
	Host                    interface{}                       `json:"host"`
	JsIntegrityScore        interface{}                       `json:"js_integrity_score"`
	LastSeenAt              interface{}                       `json:"last_seen_at"`
	PageURLs                interface{}                       `json:"page_urls"`
	URL                     interface{}                       `json:"url"`
	URLContainsCdnCgiPath   interface{}                       `json:"url_contains_cdn_cgi_path"`
	JSON                    listZoneScriptsResponseResultJSON `json:"-"`
}

// listZoneScriptsResponseResultJSON contains the JSON metadata for the struct
// [ListZoneScriptsResponseResult]
type listZoneScriptsResponseResultJSON struct {
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

func (r *ListZoneScriptsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ListZoneScriptsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                               `json:"total_count"`
	JSON       listZoneScriptsResponseResultInfoJSON `json:"-"`
}

// listZoneScriptsResponseResultInfoJSON contains the JSON metadata for the struct
// [ListZoneScriptsResponseResultInfo]
type listZoneScriptsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ListZoneScriptsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ListZoneScriptsResponseSuccess bool

const (
	ListZoneScriptsResponseSuccessTrue ListZoneScriptsResponseSuccess = true
)

type ZonePageShieldScriptPageShieldListPageShieldScriptsParams struct {
	// The direction used to sort returned scripts.
	Direction param.Field[ZonePageShieldScriptPageShieldListPageShieldScriptsParamsDirection] `query:"direction"`
	// When true, excludes scripts seen in a `/cdn-cgi` path from the returned scripts.
	// The default value is true.
	ExcludeCdnCgi param.Field[bool] `query:"exclude_cdn_cgi"`
	// Excludes scripts whose URL contains one the URL-encoded URL substrings
	// (separated by commas) from the returned scripts.
	ExcludeURLs param.Field[string] `query:"exclude_urls"`
	// Export the list of scripts as a file. Cannot be used with per_page or page
	// options.
	Export param.Field[ZonePageShieldScriptPageShieldListPageShieldScriptsParamsExport] `query:"export"`
	// Filters the returned scripts by one or more URL-encoded hostname substrings
	// separated by commas.
	Hosts param.Field[string] `query:"hosts"`
	// The field used to sort returned scripts.
	OrderBy param.Field[ZonePageShieldScriptPageShieldListPageShieldScriptsParamsOrderBy] `query:"order_by"`
	// The current page number of the paginated results.
	Page param.Field[float64] `query:"page"`
	// Filters the returned scripts by the page URL where they were last seen.
	PageURL param.Field[string] `query:"page_url"`
	// The number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// When true, malicious scripts appear first in the returned scripts.
	PrioritizeMalicious param.Field[bool] `query:"prioritize_malicious"`
	// Filters the returned scripts using a comma-separated list of scripts statuses.
	// Accepted values: `active`, `infrequent`, and `inactive`. The default value is
	// `active`.
	Status param.Field[string] `query:"status"`
	// Filters the returned scripts by one or more URL-encoded URL substrings separated
	// by commas.
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

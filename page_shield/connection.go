// File generated from our OpenAPI spec by Stainless.

package page_shield

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ConnectionService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewConnectionService] method instead.
type ConnectionService struct {
	Options []option.RequestOption
}

// NewConnectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectionService(opts ...option.RequestOption) (r *ConnectionService) {
	r = &ConnectionService{}
	r.Options = opts
	return
}

// Lists all connections detected by Page Shield.
func (r *ConnectionService) List(ctx context.Context, params ConnectionListParams, opts ...option.RequestOption) (res *[]PageShieldConnection, err error) {
	opts = append(r.Options[:], opts...)
	var env ConnectionListResponseEnvelope
	path := fmt.Sprintf("zones/%s/page_shield/connections", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a connection detected by Page Shield by connection ID.
func (r *ConnectionService) Get(ctx context.Context, connectionID string, query ConnectionGetParams, opts ...option.RequestOption) (res *PageShieldConnection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/page_shield/connections/%s", query.ZoneID, connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PageShieldConnection struct {
	ID                      interface{}              `json:"id"`
	AddedAt                 interface{}              `json:"added_at"`
	DomainReportedMalicious interface{}              `json:"domain_reported_malicious"`
	FirstPageURL            interface{}              `json:"first_page_url"`
	FirstSeenAt             interface{}              `json:"first_seen_at"`
	Host                    interface{}              `json:"host"`
	LastSeenAt              interface{}              `json:"last_seen_at"`
	PageURLs                interface{}              `json:"page_urls"`
	URL                     interface{}              `json:"url"`
	URLContainsCdnCgiPath   interface{}              `json:"url_contains_cdn_cgi_path"`
	JSON                    pageShieldConnectionJSON `json:"-"`
}

// pageShieldConnectionJSON contains the JSON metadata for the struct
// [PageShieldConnection]
type pageShieldConnectionJSON struct {
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

func (r *PageShieldConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageShieldConnectionJSON) RawJSON() string {
	return r.raw
}

type ConnectionListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The direction used to sort returned connections.
	Direction param.Field[ConnectionListParamsDirection] `query:"direction"`
	// When true, excludes connections seen in a `/cdn-cgi` path from the returned
	// connections. The default value is true.
	ExcludeCdnCgi param.Field[bool] `query:"exclude_cdn_cgi"`
	// Excludes connections whose URL contains one of the URL-encoded URLs separated by
	// commas.
	ExcludeURLs param.Field[string] `query:"exclude_urls"`
	// Export the list of connections as a file. Cannot be used with per_page or page
	// options.
	Export param.Field[ConnectionListParamsExport] `query:"export"`
	// Includes connections that match one or more URL-encoded hostnames separated by
	// commas.
	//
	// Wildcards are supported at the start and end of each hostname to support starts
	// with, ends with and contains. If no wildcards are used, results will be filtered
	// by exact match
	Hosts param.Field[string] `query:"hosts"`
	// The field used to sort returned connections.
	OrderBy param.Field[ConnectionListParamsOrderBy] `query:"order_by"`
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

// URLQuery serializes [ConnectionListParams]'s query parameters as `url.Values`.
func (r ConnectionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned connections.
type ConnectionListParamsDirection string

const (
	ConnectionListParamsDirectionAsc  ConnectionListParamsDirection = "asc"
	ConnectionListParamsDirectionDesc ConnectionListParamsDirection = "desc"
)

// Export the list of connections as a file. Cannot be used with per_page or page
// options.
type ConnectionListParamsExport string

const (
	ConnectionListParamsExportCsv ConnectionListParamsExport = "csv"
)

// The field used to sort returned connections.
type ConnectionListParamsOrderBy string

const (
	ConnectionListParamsOrderByFirstSeenAt ConnectionListParamsOrderBy = "first_seen_at"
	ConnectionListParamsOrderByLastSeenAt  ConnectionListParamsOrderBy = "last_seen_at"
)

type ConnectionListResponseEnvelope struct {
	Errors   []ConnectionListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConnectionListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PageShieldConnection                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ConnectionListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ConnectionListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       connectionListResponseEnvelopeJSON       `json:"-"`
}

// connectionListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectionListResponseEnvelope]
type connectionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectionListResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    connectionListResponseEnvelopeErrorsJSON `json:"-"`
}

// connectionListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConnectionListResponseEnvelopeErrors]
type connectionListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectionListResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    connectionListResponseEnvelopeMessagesJSON `json:"-"`
}

// connectionListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConnectionListResponseEnvelopeMessages]
type connectionListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConnectionListResponseEnvelopeSuccess bool

const (
	ConnectionListResponseEnvelopeSuccessTrue ConnectionListResponseEnvelopeSuccess = true
)

type ConnectionListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       connectionListResponseEnvelopeResultInfoJSON `json:"-"`
}

// connectionListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [ConnectionListResponseEnvelopeResultInfo]
type connectionListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectionListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectionListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ConnectionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

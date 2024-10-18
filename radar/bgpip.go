// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// BGPIPService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPIPService] method instead.
type BGPIPService struct {
	Options []option.RequestOption
}

// NewBGPIPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBGPIPService(opts ...option.RequestOption) (r *BGPIPService) {
	r = &BGPIPService{}
	r.Options = opts
	return
}

// Get time series data for the announced IP space count, represented as the number
// of IPv4 /24s and IPv6 /48s, for a given ASN.
func (r *BGPIPService) Timeseries(ctx context.Context, query BGPIPTimeseriesParams, opts ...option.RequestOption) (res *BgpipTimeseriesResponse, err error) {
	var env BgpipTimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/bgp/ips/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BgpipTimeseriesResponse struct {
	Meta     BgpipTimeseriesResponseMeta     `json:"meta,required"`
	Serie174 BgpipTimeseriesResponseSerie174 `json:"serie_174,required"`
	JSON     bgpipTimeseriesResponseJSON     `json:"-"`
}

// bgpipTimeseriesResponseJSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponse]
type bgpipTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie174    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseMeta struct {
	Queries []BgpipTimeseriesResponseMetaQuery `json:"queries,required"`
	JSON    bgpipTimeseriesResponseMetaJSON    `json:"-"`
}

// bgpipTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponseMeta]
type bgpipTimeseriesResponseMetaJSON struct {
	Queries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseMetaQuery struct {
	DateRange BgpipTimeseriesResponseMetaQueriesDateRange `json:"dateRange,required"`
	Entity    string                                      `json:"entity,required"`
	JSON      bgpipTimeseriesResponseMetaQueryJSON        `json:"-"`
}

// bgpipTimeseriesResponseMetaQueryJSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponseMetaQuery]
type bgpipTimeseriesResponseMetaQueryJSON struct {
	DateRange   apijson.Field
	Entity      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseMetaQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaQueryJSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseMetaQueriesDateRange struct {
	EndTime   string                                          `json:"endTime,required"`
	StartTime string                                          `json:"startTime,required"`
	JSON      bgpipTimeseriesResponseMetaQueriesDateRangeJSON `json:"-"`
}

// bgpipTimeseriesResponseMetaQueriesDateRangeJSON contains the JSON metadata for
// the struct [BgpipTimeseriesResponseMetaQueriesDateRange]
type bgpipTimeseriesResponseMetaQueriesDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseMetaQueriesDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseMetaQueriesDateRangeJSON) RawJSON() string {
	return r.raw
}

type BgpipTimeseriesResponseSerie174 struct {
	IPV4       []string                            `json:"ipv4,required"`
	IPV6       []string                            `json:"ipv6,required"`
	Timestamps []time.Time                         `json:"timestamps,required" format:"date-time"`
	JSON       bgpipTimeseriesResponseSerie174JSON `json:"-"`
}

// bgpipTimeseriesResponseSerie174JSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponseSerie174]
type bgpipTimeseriesResponseSerie174JSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseSerie174) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseSerie174JSON) RawJSON() string {
	return r.raw
}

type BGPIPTimeseriesParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[BgpipTimeseriesParamsFormat] `query:"format"`
	// Include data delay meta information
	IncludeDelay param.Field[bool] `query:"includeDelay"`
	// Filter for ip version.
	IPVersion param.Field[[]BgpipTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Array of locations (alpha-2 country codes).
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [BGPIPTimeseriesParams]'s query parameters as `url.Values`.
func (r BGPIPTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format results are returned in.
type BgpipTimeseriesParamsFormat string

const (
	BgpipTimeseriesParamsFormatJson BgpipTimeseriesParamsFormat = "JSON"
	BgpipTimeseriesParamsFormatCsv  BgpipTimeseriesParamsFormat = "CSV"
)

func (r BgpipTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case BgpipTimeseriesParamsFormatJson, BgpipTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type BgpipTimeseriesParamsIPVersion string

const (
	BgpipTimeseriesParamsIPVersionIPv4 BgpipTimeseriesParamsIPVersion = "IPv4"
	BgpipTimeseriesParamsIPVersionIPv6 BgpipTimeseriesParamsIPVersion = "IPv6"
)

func (r BgpipTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case BgpipTimeseriesParamsIPVersionIPv4, BgpipTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

type BgpipTimeseriesResponseEnvelope struct {
	Result  BgpipTimeseriesResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    bgpipTimeseriesResponseEnvelopeJSON `json:"-"`
}

// bgpipTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [BgpipTimeseriesResponseEnvelope]
type bgpipTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

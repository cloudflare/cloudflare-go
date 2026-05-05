// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// BGPRPKIRoaService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPRPKIRoaService] method instead.
type BGPRPKIRoaService struct {
	Options []option.RequestOption
}

// NewBGPRPKIRoaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBGPRPKIRoaService(opts ...option.RequestOption) (r *BGPRPKIRoaService) {
	r = &BGPRPKIRoaService{}
	r.Options = opts
	return
}

// Retrieves RPKI ROA (Route Origin Authorization) validation ratios over time.
// Returns the selected metric as a time series. Supports filtering by ASN or
// location (country code) — multiple values of the same filter type produce one
// series per value. If no ASN or location is specified, returns the global
// aggregate.
func (r *BGPRPKIRoaService) Timeseries(ctx context.Context, query BGPRPKIRoaTimeseriesParams, opts ...option.RequestOption) (res *BgprpkiRoaTimeseriesResponse, err error) {
	var env BgprpkiRoaTimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/bgp/rpki/roas/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type BgprpkiRoaTimeseriesResponse struct {
	Meta   BgprpkiRoaTimeseriesResponseMeta   `json:"meta" api:"required"`
	Serie0 BgprpkiRoaTimeseriesResponseSerie0 `json:"serie_0" api:"required"`
	JSON   bgprpkiRoaTimeseriesResponseJSON   `json:"-"`
}

// bgprpkiRoaTimeseriesResponseJSON contains the JSON metadata for the struct
// [BgprpkiRoaTimeseriesResponse]
type bgprpkiRoaTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiRoaTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiRoaTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type BgprpkiRoaTimeseriesResponseMeta struct {
	// Timestamp of the underlying data.
	DataTime time.Time `json:"dataTime" api:"required" format:"date-time"`
	// Timestamp when the query was executed.
	QueryTime time.Time                            `json:"queryTime" api:"required" format:"date-time"`
	JSON      bgprpkiRoaTimeseriesResponseMetaJSON `json:"-"`
}

// bgprpkiRoaTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [BgprpkiRoaTimeseriesResponseMeta]
type bgprpkiRoaTimeseriesResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiRoaTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiRoaTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BgprpkiRoaTimeseriesResponseSerie0 struct {
	Timestamps []time.Time                            `json:"timestamps" api:"required" format:"date-time"`
	Values     []string                               `json:"values" api:"required"`
	JSON       bgprpkiRoaTimeseriesResponseSerie0JSON `json:"-"`
}

// bgprpkiRoaTimeseriesResponseSerie0JSON contains the JSON metadata for the struct
// [BgprpkiRoaTimeseriesResponseSerie0]
type bgprpkiRoaTimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiRoaTimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiRoaTimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type BGPRPKIRoaTimeseriesParams struct {
	// Filters results by Autonomous System Number. Specify one or more ASNs. Multiple
	// values generate one series per ASN.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[BgprpkiRoaTimeseriesParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 location
	// codes.
	Location param.Field[[]string] `query:"location"`
	// Which RPKI ROA validation metric to return. validPfxsRatio = ratio of RPKI-valid
	// prefixes (IPv4+IPv6 combined). validPfxsV4Ratio / validPfxsV6Ratio = same, split
	// by IP version. validIpsRatio = ratio of RPKI-valid address space (IPv4 /24s +
	// IPv6 /48s). validIpsV4Ratio / validIpsV6Ratio = same, split by IP version.
	Metric param.Field[BgprpkiRoaTimeseriesParamsMetric] `query:"metric"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [BGPRPKIRoaTimeseriesParams]'s query parameters as
// `url.Values`.
func (r BGPRPKIRoaTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type BgprpkiRoaTimeseriesParamsFormat string

const (
	BgprpkiRoaTimeseriesParamsFormatJson BgprpkiRoaTimeseriesParamsFormat = "JSON"
	BgprpkiRoaTimeseriesParamsFormatCsv  BgprpkiRoaTimeseriesParamsFormat = "CSV"
)

func (r BgprpkiRoaTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case BgprpkiRoaTimeseriesParamsFormatJson, BgprpkiRoaTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

// Which RPKI ROA validation metric to return. validPfxsRatio = ratio of RPKI-valid
// prefixes (IPv4+IPv6 combined). validPfxsV4Ratio / validPfxsV6Ratio = same, split
// by IP version. validIpsRatio = ratio of RPKI-valid address space (IPv4 /24s +
// IPv6 /48s). validIpsV4Ratio / validIpsV6Ratio = same, split by IP version.
type BgprpkiRoaTimeseriesParamsMetric string

const (
	BgprpkiRoaTimeseriesParamsMetricValidPfxsRatio   BgprpkiRoaTimeseriesParamsMetric = "validPfxsRatio"
	BgprpkiRoaTimeseriesParamsMetricValidPfxsV4Ratio BgprpkiRoaTimeseriesParamsMetric = "validPfxsV4Ratio"
	BgprpkiRoaTimeseriesParamsMetricValidPfxsV6Ratio BgprpkiRoaTimeseriesParamsMetric = "validPfxsV6Ratio"
	BgprpkiRoaTimeseriesParamsMetricValidIPsRatio    BgprpkiRoaTimeseriesParamsMetric = "validIpsRatio"
	BgprpkiRoaTimeseriesParamsMetricValidIPsV4Ratio  BgprpkiRoaTimeseriesParamsMetric = "validIpsV4Ratio"
	BgprpkiRoaTimeseriesParamsMetricValidIPsV6Ratio  BgprpkiRoaTimeseriesParamsMetric = "validIpsV6Ratio"
)

func (r BgprpkiRoaTimeseriesParamsMetric) IsKnown() bool {
	switch r {
	case BgprpkiRoaTimeseriesParamsMetricValidPfxsRatio, BgprpkiRoaTimeseriesParamsMetricValidPfxsV4Ratio, BgprpkiRoaTimeseriesParamsMetricValidPfxsV6Ratio, BgprpkiRoaTimeseriesParamsMetricValidIPsRatio, BgprpkiRoaTimeseriesParamsMetricValidIPsV4Ratio, BgprpkiRoaTimeseriesParamsMetricValidIPsV6Ratio:
		return true
	}
	return false
}

type BgprpkiRoaTimeseriesResponseEnvelope struct {
	Result  BgprpkiRoaTimeseriesResponse             `json:"result" api:"required"`
	Success bool                                     `json:"success" api:"required"`
	JSON    bgprpkiRoaTimeseriesResponseEnvelopeJSON `json:"-"`
}

// bgprpkiRoaTimeseriesResponseEnvelopeJSON contains the JSON metadata for the
// struct [BgprpkiRoaTimeseriesResponseEnvelope]
type bgprpkiRoaTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiRoaTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiRoaTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

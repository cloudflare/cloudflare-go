// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// BGPRouteUpstreamService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPRouteUpstreamService] method instead.
type BGPRouteUpstreamService struct {
	Options []option.RequestOption
}

// NewBGPRouteUpstreamService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBGPRouteUpstreamService(opts ...option.RequestOption) (r *BGPRouteUpstreamService) {
	r = &BGPRouteUpstreamService{}
	r.Options = opts
	return
}

// Retrieves the share of an AS’s observed paths carried by each direct upstream
// over time, derived from RouteViews RIB snapshots across all collectors (the
// combined product). Each upstream ASN is returned as its own series of shares
// (0–1); the least-significant upstreams beyond the requested limit are grouped
// into an "OTHER" series. Series share a common set of timestamps.
func (r *BGPRouteUpstreamService) Timeseries(ctx context.Context, asn int64, query BGPRouteUpstreamTimeseriesParams, opts ...option.RequestOption) (res *BGPRouteUpstreamTimeseriesResponse, err error) {
	var env BGPRouteUpstreamTimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/bgp/routes/upstreams/%v/timeseries", asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type BGPRouteUpstreamTimeseriesResponse struct {
	Meta   BGPRouteUpstreamTimeseriesResponseMeta   `json:"meta" api:"required"`
	Serie0 BGPRouteUpstreamTimeseriesResponseSerie0 `json:"serie_0" api:"required"`
	JSON   bgpRouteUpstreamTimeseriesResponseJSON   `json:"-"`
}

// bgpRouteUpstreamTimeseriesResponseJSON contains the JSON metadata for the struct
// [BGPRouteUpstreamTimeseriesResponse]
type bgpRouteUpstreamTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteUpstreamTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteUpstreamTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type BGPRouteUpstreamTimeseriesResponseMeta struct {
	// Timestamp of the underlying RIB data.
	DataTime           time.Time `json:"dataTime" api:"required,nullable" format:"date-time"`
	EffectiveCollector string    `json:"effectiveCollector" api:"required,nullable"`
	// Timestamp when the query was executed.
	QueryTime time.Time                                  `json:"queryTime" api:"required,nullable" format:"date-time"`
	Stale     bool                                       `json:"stale" api:"required"`
	JSON      bgpRouteUpstreamTimeseriesResponseMetaJSON `json:"-"`
}

// bgpRouteUpstreamTimeseriesResponseMetaJSON contains the JSON metadata for the
// struct [BGPRouteUpstreamTimeseriesResponseMeta]
type bgpRouteUpstreamTimeseriesResponseMetaJSON struct {
	DataTime           apijson.Field
	EffectiveCollector apijson.Field
	QueryTime          apijson.Field
	Stale              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BGPRouteUpstreamTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteUpstreamTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BGPRouteUpstreamTimeseriesResponseSerie0 struct {
	Timestamps  []time.Time                                  `json:"timestamps" api:"required" format:"date-time"`
	ExtraFields map[string][]string                          `json:"-" api:"extrafields"`
	JSON        bgpRouteUpstreamTimeseriesResponseSerie0JSON `json:"-"`
}

// bgpRouteUpstreamTimeseriesResponseSerie0JSON contains the JSON metadata for the
// struct [BGPRouteUpstreamTimeseriesResponseSerie0]
type bgpRouteUpstreamTimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteUpstreamTimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteUpstreamTimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type BGPRouteUpstreamTimeseriesParams struct {
	// End of the date range (inclusive). Alternative to `dateRange`; provide together
	// with `dateStart`.
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Start of the date range (inclusive). Alternative to `dateRange`; provide
	// together with `dateEnd`.
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[BGPRouteUpstreamTimeseriesParamsFormat] `query:"format"`
	// Address family of the observed paths. Defaults to IPv4.
	IPVersion param.Field[BGPRouteUpstreamTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Number of upstream ASNs to return as separate series, ranked by the first
	// bucket. Remaining upstreams are grouped into an "OTHER" series. Defaults to 5.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [BGPRouteUpstreamTimeseriesParams]'s query parameters as
// `url.Values`.
func (r BGPRouteUpstreamTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type BGPRouteUpstreamTimeseriesParamsFormat string

const (
	BGPRouteUpstreamTimeseriesParamsFormatJson BGPRouteUpstreamTimeseriesParamsFormat = "JSON"
	BGPRouteUpstreamTimeseriesParamsFormatCsv  BGPRouteUpstreamTimeseriesParamsFormat = "CSV"
)

func (r BGPRouteUpstreamTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case BGPRouteUpstreamTimeseriesParamsFormatJson, BGPRouteUpstreamTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

// Address family of the observed paths. Defaults to IPv4.
type BGPRouteUpstreamTimeseriesParamsIPVersion string

const (
	BGPRouteUpstreamTimeseriesParamsIPVersionIPv4 BGPRouteUpstreamTimeseriesParamsIPVersion = "IPv4"
	BGPRouteUpstreamTimeseriesParamsIPVersionIPv6 BGPRouteUpstreamTimeseriesParamsIPVersion = "IPv6"
)

func (r BGPRouteUpstreamTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case BGPRouteUpstreamTimeseriesParamsIPVersionIPv4, BGPRouteUpstreamTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

type BGPRouteUpstreamTimeseriesResponseEnvelope struct {
	Result  BGPRouteUpstreamTimeseriesResponse             `json:"result" api:"required"`
	Success bool                                           `json:"success" api:"required"`
	JSON    bgpRouteUpstreamTimeseriesResponseEnvelopeJSON `json:"-"`
}

// bgpRouteUpstreamTimeseriesResponseEnvelopeJSON contains the JSON metadata for
// the struct [BGPRouteUpstreamTimeseriesResponseEnvelope]
type bgpRouteUpstreamTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRouteUpstreamTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRouteUpstreamTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

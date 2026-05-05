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

// BGPIPTopService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPIPTopService] method instead.
type BGPIPTopService struct {
	Options []option.RequestOption
}

// NewBGPIPTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBGPIPTopService(opts ...option.RequestOption) (r *BGPIPTopService) {
	r = &BGPIPTopService{}
	r.Options = opts
	return
}

// Returns the top-N autonomous systems by announced IP space at the nearest 8-hour
// RIB boundary at or before the requested date. The snapped boundary is returned
// as `anchor_ts`.
func (r *BGPIPTopService) Ases(ctx context.Context, query BGPIPTopAsesParams, opts ...option.RequestOption) (res *BgpipTopAsesResponse, err error) {
	var env BgpipTopAsesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/bgp/ips/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type BgpipTopAsesResponse struct {
	AnchorTs time.Time                 `json:"anchorTs" api:"required" format:"date-time"`
	ASNs     []BgpipTopAsesResponseASN `json:"asns" api:"required"`
	Country  string                    `json:"country" api:"required,nullable"`
	Metric   string                    `json:"metric" api:"required"`
	JSON     bgpipTopAsesResponseJSON  `json:"-"`
}

// bgpipTopAsesResponseJSON contains the JSON metadata for the struct
// [BgpipTopAsesResponse]
type bgpipTopAsesResponseJSON struct {
	AnchorTs    apijson.Field
	ASNs        apijson.Field
	Country     apijson.Field
	Metric      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTopAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTopAsesResponseJSON) RawJSON() string {
	return r.raw
}

type BgpipTopAsesResponseASN struct {
	ASN    int64                       `json:"asn" api:"required"`
	V4_24s int64                       `json:"v4_24s" api:"required"`
	V6_48s int64                       `json:"v6_48s" api:"required"`
	JSON   bgpipTopAsesResponseASNJSON `json:"-"`
}

// bgpipTopAsesResponseASNJSON contains the JSON metadata for the struct
// [BgpipTopAsesResponseASN]
type bgpipTopAsesResponseASNJSON struct {
	ASN         apijson.Field
	V4_24s      apijson.Field
	V6_48s      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTopAsesResponseASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTopAsesResponseASNJSON) RawJSON() string {
	return r.raw
}

type BGPIPTopAsesParams struct {
	// Optional ISO 3166-1 alpha-2 country filter. Omit for global top-N.
	Country param.Field[string] `query:"country"`
	// Filters results by the specified datetime (ISO 8601).
	Date param.Field[time.Time] `query:"date" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[BgpipTopAsesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Ranking metric: IPv4 /24 count or IPv6 /48 count.
	Metric param.Field[BgpipTopAsesParamsMetric] `query:"metric"`
}

// URLQuery serializes [BGPIPTopAsesParams]'s query parameters as `url.Values`.
func (r BGPIPTopAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type BgpipTopAsesParamsFormat string

const (
	BgpipTopAsesParamsFormatJson BgpipTopAsesParamsFormat = "JSON"
	BgpipTopAsesParamsFormatCsv  BgpipTopAsesParamsFormat = "CSV"
)

func (r BgpipTopAsesParamsFormat) IsKnown() bool {
	switch r {
	case BgpipTopAsesParamsFormatJson, BgpipTopAsesParamsFormatCsv:
		return true
	}
	return false
}

// Ranking metric: IPv4 /24 count or IPv6 /48 count.
type BgpipTopAsesParamsMetric string

const (
	BgpipTopAsesParamsMetricV4_24s BgpipTopAsesParamsMetric = "v4_24s"
	BgpipTopAsesParamsMetricV6_48s BgpipTopAsesParamsMetric = "v6_48s"
)

func (r BgpipTopAsesParamsMetric) IsKnown() bool {
	switch r {
	case BgpipTopAsesParamsMetricV4_24s, BgpipTopAsesParamsMetricV6_48s:
		return true
	}
	return false
}

type BgpipTopAsesResponseEnvelope struct {
	Result  BgpipTopAsesResponse             `json:"result" api:"required"`
	Success bool                             `json:"success" api:"required"`
	JSON    bgpipTopAsesResponseEnvelopeJSON `json:"-"`
}

// bgpipTopAsesResponseEnvelopeJSON contains the JSON metadata for the struct
// [BgpipTopAsesResponseEnvelope]
type bgpipTopAsesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgpipTopAsesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpipTopAsesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

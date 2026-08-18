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

// BGPRoutePathService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPRoutePathService] method instead.
type BGPRoutePathService struct {
	Options []option.RequestOption
}

// NewBGPRoutePathService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBGPRoutePathService(opts ...option.RequestOption) (r *BGPRoutePathService) {
	r = &BGPRoutePathService{}
	r.Options = opts
	return
}

// Retrieves the paths an AS uses to reach the tier-1 clique, derived from
// RouteViews RIB snapshots. Each entry is an ordered AS-path segment (from the
// queried AS toward a tier-1) with the number of observed paths and peers, and the
// collectors that observed it. By default segments are merged across all active
// collectors; pass "collector" to scope to one. The response also includes an
// "asnInfo" map (keyed by ASN) with the name and country for every ASN in the
// returned segments plus the queried ASN (best-effort; null when unavailable).
func (r *BGPRoutePathService) List(ctx context.Context, asn int64, query BGPRoutePathListParams, opts ...option.RequestOption) (res *BGPRoutePathListResponse, err error) {
	var env BGPRoutePathListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/bgp/routes/paths/%v", asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type BGPRoutePathListResponse struct {
	ASNInfo    map[string]BGPRoutePathListResponseASNInfo `json:"asnInfo" api:"required,nullable"`
	Collectors []string                                   `json:"collectors" api:"required"`
	Meta       BGPRoutePathListResponseMeta               `json:"meta" api:"required"`
	Paths      []BGPRoutePathListResponsePath             `json:"paths" api:"required"`
	JSON       bgpRoutePathListResponseJSON               `json:"-"`
}

// bgpRoutePathListResponseJSON contains the JSON metadata for the struct
// [BGPRoutePathListResponse]
type bgpRoutePathListResponseJSON struct {
	ASNInfo     apijson.Field
	Collectors  apijson.Field
	Meta        apijson.Field
	Paths       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRoutePathListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRoutePathListResponseJSON) RawJSON() string {
	return r.raw
}

type BGPRoutePathListResponseASNInfo struct {
	// ASN number.
	ASN int64 `json:"asn" api:"required"`
	// Alpha-2 country code.
	Country string `json:"country" api:"required,nullable"`
	// AS name.
	Name string                              `json:"name" api:"required,nullable"`
	JSON bgpRoutePathListResponseASNInfoJSON `json:"-"`
}

// bgpRoutePathListResponseASNInfoJSON contains the JSON metadata for the struct
// [BGPRoutePathListResponseASNInfo]
type bgpRoutePathListResponseASNInfoJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRoutePathListResponseASNInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRoutePathListResponseASNInfoJSON) RawJSON() string {
	return r.raw
}

type BGPRoutePathListResponseMeta struct {
	// Timestamp of the underlying RIB data.
	DataTime           time.Time `json:"dataTime" api:"required,nullable" format:"date-time"`
	EffectiveCollector string    `json:"effectiveCollector" api:"required,nullable"`
	// Timestamp when the query was executed.
	QueryTime time.Time                        `json:"queryTime" api:"required,nullable" format:"date-time"`
	Stale     bool                             `json:"stale" api:"required"`
	JSON      bgpRoutePathListResponseMetaJSON `json:"-"`
}

// bgpRoutePathListResponseMetaJSON contains the JSON metadata for the struct
// [BGPRoutePathListResponseMeta]
type bgpRoutePathListResponseMetaJSON struct {
	DataTime           apijson.Field
	EffectiveCollector apijson.Field
	QueryTime          apijson.Field
	Stale              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *BGPRoutePathListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRoutePathListResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BGPRoutePathListResponsePath struct {
	Collectors []string                         `json:"collectors" api:"required"`
	PathsCount int64                            `json:"pathsCount" api:"required"`
	PeersCount int64                            `json:"peersCount" api:"required"`
	Segment    []int64                          `json:"segment" api:"required"`
	JSON       bgpRoutePathListResponsePathJSON `json:"-"`
}

// bgpRoutePathListResponsePathJSON contains the JSON metadata for the struct
// [BGPRoutePathListResponsePath]
type bgpRoutePathListResponsePathJSON struct {
	Collectors  apijson.Field
	PathsCount  apijson.Field
	PeersCount  apijson.Field
	Segment     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRoutePathListResponsePath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRoutePathListResponsePathJSON) RawJSON() string {
	return r.raw
}

type BGPRoutePathListParams struct {
	// Scope to a single RouteViews collector (e.g. "route-views3"). Omit to merge
	// across all active collectors (identical path segments are deduplicated,
	// observation counts summed, and every contributing collector listed).
	Collector param.Field[string] `query:"collector"`
	// Format in which results will be returned.
	Format param.Field[BGPRoutePathListParamsFormat] `query:"format"`
	// Address family of the observed paths. Defaults to IPv4.
	IPVersion param.Field[BGPRoutePathListParamsIPVersion] `query:"ipVersion"`
}

// URLQuery serializes [BGPRoutePathListParams]'s query parameters as `url.Values`.
func (r BGPRoutePathListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type BGPRoutePathListParamsFormat string

const (
	BGPRoutePathListParamsFormatJson BGPRoutePathListParamsFormat = "JSON"
	BGPRoutePathListParamsFormatCsv  BGPRoutePathListParamsFormat = "CSV"
)

func (r BGPRoutePathListParamsFormat) IsKnown() bool {
	switch r {
	case BGPRoutePathListParamsFormatJson, BGPRoutePathListParamsFormatCsv:
		return true
	}
	return false
}

// Address family of the observed paths. Defaults to IPv4.
type BGPRoutePathListParamsIPVersion string

const (
	BGPRoutePathListParamsIPVersionIPv4 BGPRoutePathListParamsIPVersion = "IPv4"
	BGPRoutePathListParamsIPVersionIPv6 BGPRoutePathListParamsIPVersion = "IPv6"
)

func (r BGPRoutePathListParamsIPVersion) IsKnown() bool {
	switch r {
	case BGPRoutePathListParamsIPVersionIPv4, BGPRoutePathListParamsIPVersionIPv6:
		return true
	}
	return false
}

type BGPRoutePathListResponseEnvelope struct {
	Result  BGPRoutePathListResponse             `json:"result" api:"required"`
	Success bool                                 `json:"success" api:"required"`
	JSON    bgpRoutePathListResponseEnvelopeJSON `json:"-"`
}

// bgpRoutePathListResponseEnvelopeJSON contains the JSON metadata for the struct
// [BGPRoutePathListResponseEnvelope]
type bgpRoutePathListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPRoutePathListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpRoutePathListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

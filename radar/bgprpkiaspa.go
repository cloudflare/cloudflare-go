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

// BGPRPKIASPAService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBGPRPKIASPAService] method instead.
type BGPRPKIASPAService struct {
	Options []option.RequestOption
}

// NewBGPRPKIASPAService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBGPRPKIASPAService(opts ...option.RequestOption) (r *BGPRPKIASPAService) {
	r = &BGPRPKIASPAService{}
	r.Options = opts
	return
}

// Retrieves ASPA (Autonomous System Provider Authorization) changes over time.
// Returns daily aggregated changes including additions, removals, and
// modifications of ASPA objects.
func (r *BGPRPKIASPAService) Changes(ctx context.Context, query BGPRPKIASPAChangesParams, opts ...option.RequestOption) (res *BgprpkiaspaChangesResponse, err error) {
	var env BgprpkiaspaChangesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/bgp/rpki/aspa/changes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves current or historical ASPA (Autonomous System Provider Authorization)
// objects. ASPA objects define which ASNs are authorized upstream providers for a
// customer ASN.
func (r *BGPRPKIASPAService) Snapshot(ctx context.Context, query BGPRPKIASPASnapshotParams, opts ...option.RequestOption) (res *BgprpkiaspaSnapshotResponse, err error) {
	var env BgprpkiaspaSnapshotResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/bgp/rpki/aspa/snapshot"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves ASPA (Autonomous System Provider Authorization) object count over
// time. Supports filtering by RIR or location (country code) to generate multiple
// named series. If no RIR or location filter is specified, returns total count.
func (r *BGPRPKIASPAService) Timeseries(ctx context.Context, query BGPRPKIASPATimeseriesParams, opts ...option.RequestOption) (res *BgprpkiaspaTimeseriesResponse, err error) {
	var env BgprpkiaspaTimeseriesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/bgp/rpki/aspa/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BgprpkiaspaChangesResponse struct {
	ASNInfo BgprpkiaspaChangesResponseASNInfo  `json:"asnInfo,required"`
	Changes []BgprpkiaspaChangesResponseChange `json:"changes,required"`
	Meta    BgprpkiaspaChangesResponseMeta     `json:"meta,required"`
	JSON    bgprpkiaspaChangesResponseJSON     `json:"-"`
}

// bgprpkiaspaChangesResponseJSON contains the JSON metadata for the struct
// [BgprpkiaspaChangesResponse]
type bgprpkiaspaChangesResponseJSON struct {
	ASNInfo     apijson.Field
	Changes     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaChangesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaChangesResponseJSON) RawJSON() string {
	return r.raw
}

type BgprpkiaspaChangesResponseASNInfo struct {
	Number13335 BgprpkiaspaChangesResponseASNInfo13335 `json:"13335,required"`
	JSON        bgprpkiaspaChangesResponseASNInfoJSON  `json:"-"`
}

// bgprpkiaspaChangesResponseASNInfoJSON contains the JSON metadata for the struct
// [BgprpkiaspaChangesResponseASNInfo]
type bgprpkiaspaChangesResponseASNInfoJSON struct {
	Number13335 apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaChangesResponseASNInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaChangesResponseASNInfoJSON) RawJSON() string {
	return r.raw
}

type BgprpkiaspaChangesResponseASNInfo13335 struct {
	// ASN number.
	ASN int64 `json:"asn,required"`
	// Alpha-2 country code.
	Country string `json:"country,required"`
	// AS name.
	Name string                                     `json:"name,required"`
	JSON bgprpkiaspaChangesResponseASNInfo13335JSON `json:"-"`
}

// bgprpkiaspaChangesResponseASNInfo13335JSON contains the JSON metadata for the
// struct [BgprpkiaspaChangesResponseASNInfo13335]
type bgprpkiaspaChangesResponseASNInfo13335JSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaChangesResponseASNInfo13335) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaChangesResponseASNInfo13335JSON) RawJSON() string {
	return r.raw
}

type BgprpkiaspaChangesResponseChange struct {
	// Number of new ASPA objects created.
	CustomersAdded int64 `json:"customersAdded,required"`
	// Number of ASPA objects deleted.
	CustomersRemoved int64 `json:"customersRemoved,required"`
	// Date of the changes in ISO 8601 format.
	Date    time.Time                                `json:"date,required" format:"date-time"`
	Entries []BgprpkiaspaChangesResponseChangesEntry `json:"entries,required"`
	// Number of providers added to existing objects.
	ProvidersAdded int64 `json:"providersAdded,required"`
	// Number of providers removed from existing objects.
	ProvidersRemoved int64 `json:"providersRemoved,required"`
	// Running total of active ASPA objects after this day.
	TotalCount int64                                `json:"totalCount,required"`
	JSON       bgprpkiaspaChangesResponseChangeJSON `json:"-"`
}

// bgprpkiaspaChangesResponseChangeJSON contains the JSON metadata for the struct
// [BgprpkiaspaChangesResponseChange]
type bgprpkiaspaChangesResponseChangeJSON struct {
	CustomersAdded   apijson.Field
	CustomersRemoved apijson.Field
	Date             apijson.Field
	Entries          apijson.Field
	ProvidersAdded   apijson.Field
	ProvidersRemoved apijson.Field
	TotalCount       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BgprpkiaspaChangesResponseChange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaChangesResponseChangeJSON) RawJSON() string {
	return r.raw
}

type BgprpkiaspaChangesResponseChangesEntry struct {
	// The customer ASN affected.
	CustomerASN int64                                        `json:"customerAsn,required"`
	Providers   []int64                                      `json:"providers,required"`
	Type        BgprpkiaspaChangesResponseChangesEntriesType `json:"type,required"`
	JSON        bgprpkiaspaChangesResponseChangesEntryJSON   `json:"-"`
}

// bgprpkiaspaChangesResponseChangesEntryJSON contains the JSON metadata for the
// struct [BgprpkiaspaChangesResponseChangesEntry]
type bgprpkiaspaChangesResponseChangesEntryJSON struct {
	CustomerASN apijson.Field
	Providers   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaChangesResponseChangesEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaChangesResponseChangesEntryJSON) RawJSON() string {
	return r.raw
}

type BgprpkiaspaChangesResponseChangesEntriesType string

const (
	BgprpkiaspaChangesResponseChangesEntriesTypeCustomerAdded    BgprpkiaspaChangesResponseChangesEntriesType = "CustomerAdded"
	BgprpkiaspaChangesResponseChangesEntriesTypeCustomerRemoved  BgprpkiaspaChangesResponseChangesEntriesType = "CustomerRemoved"
	BgprpkiaspaChangesResponseChangesEntriesTypeProvidersAdded   BgprpkiaspaChangesResponseChangesEntriesType = "ProvidersAdded"
	BgprpkiaspaChangesResponseChangesEntriesTypeProvidersRemoved BgprpkiaspaChangesResponseChangesEntriesType = "ProvidersRemoved"
)

func (r BgprpkiaspaChangesResponseChangesEntriesType) IsKnown() bool {
	switch r {
	case BgprpkiaspaChangesResponseChangesEntriesTypeCustomerAdded, BgprpkiaspaChangesResponseChangesEntriesTypeCustomerRemoved, BgprpkiaspaChangesResponseChangesEntriesTypeProvidersAdded, BgprpkiaspaChangesResponseChangesEntriesTypeProvidersRemoved:
		return true
	}
	return false
}

type BgprpkiaspaChangesResponseMeta struct {
	// Timestamp of the underlying data.
	DataTime time.Time `json:"dataTime,required" format:"date-time"`
	// Timestamp when the query was executed.
	QueryTime time.Time                          `json:"queryTime,required" format:"date-time"`
	JSON      bgprpkiaspaChangesResponseMetaJSON `json:"-"`
}

// bgprpkiaspaChangesResponseMetaJSON contains the JSON metadata for the struct
// [BgprpkiaspaChangesResponseMeta]
type bgprpkiaspaChangesResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaChangesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaChangesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BgprpkiaspaSnapshotResponse struct {
	ASNInfo     BgprpkiaspaSnapshotResponseASNInfo      `json:"asnInfo,required"`
	ASPAObjects []BgprpkiaspaSnapshotResponseASPAObject `json:"aspaObjects,required"`
	Meta        BgprpkiaspaSnapshotResponseMeta         `json:"meta,required"`
	JSON        bgprpkiaspaSnapshotResponseJSON         `json:"-"`
}

// bgprpkiaspaSnapshotResponseJSON contains the JSON metadata for the struct
// [BgprpkiaspaSnapshotResponse]
type bgprpkiaspaSnapshotResponseJSON struct {
	ASNInfo     apijson.Field
	ASPAObjects apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaSnapshotResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaSnapshotResponseJSON) RawJSON() string {
	return r.raw
}

type BgprpkiaspaSnapshotResponseASNInfo struct {
	Number13335 BgprpkiaspaSnapshotResponseASNInfo13335 `json:"13335,required"`
	JSON        bgprpkiaspaSnapshotResponseASNInfoJSON  `json:"-"`
}

// bgprpkiaspaSnapshotResponseASNInfoJSON contains the JSON metadata for the struct
// [BgprpkiaspaSnapshotResponseASNInfo]
type bgprpkiaspaSnapshotResponseASNInfoJSON struct {
	Number13335 apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaSnapshotResponseASNInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaSnapshotResponseASNInfoJSON) RawJSON() string {
	return r.raw
}

type BgprpkiaspaSnapshotResponseASNInfo13335 struct {
	// ASN number.
	ASN int64 `json:"asn,required"`
	// Alpha-2 country code.
	Country string `json:"country,required"`
	// AS name.
	Name string                                      `json:"name,required"`
	JSON bgprpkiaspaSnapshotResponseASNInfo13335JSON `json:"-"`
}

// bgprpkiaspaSnapshotResponseASNInfo13335JSON contains the JSON metadata for the
// struct [BgprpkiaspaSnapshotResponseASNInfo13335]
type bgprpkiaspaSnapshotResponseASNInfo13335JSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaSnapshotResponseASNInfo13335) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaSnapshotResponseASNInfo13335JSON) RawJSON() string {
	return r.raw
}

type BgprpkiaspaSnapshotResponseASPAObject struct {
	// The customer ASN publishing the ASPA object.
	CustomerASN int64                                     `json:"customerAsn,required"`
	Providers   []int64                                   `json:"providers,required"`
	JSON        bgprpkiaspaSnapshotResponseASPAObjectJSON `json:"-"`
}

// bgprpkiaspaSnapshotResponseASPAObjectJSON contains the JSON metadata for the
// struct [BgprpkiaspaSnapshotResponseASPAObject]
type bgprpkiaspaSnapshotResponseASPAObjectJSON struct {
	CustomerASN apijson.Field
	Providers   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaSnapshotResponseASPAObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaSnapshotResponseASPAObjectJSON) RawJSON() string {
	return r.raw
}

type BgprpkiaspaSnapshotResponseMeta struct {
	// Timestamp of the underlying data.
	DataTime time.Time `json:"dataTime,required" format:"date-time"`
	// Timestamp when the query was executed.
	QueryTime time.Time `json:"queryTime,required" format:"date-time"`
	// Total number of ASPA objects.
	TotalCount int64                               `json:"totalCount,required"`
	JSON       bgprpkiaspaSnapshotResponseMetaJSON `json:"-"`
}

// bgprpkiaspaSnapshotResponseMetaJSON contains the JSON metadata for the struct
// [BgprpkiaspaSnapshotResponseMeta]
type bgprpkiaspaSnapshotResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaSnapshotResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaSnapshotResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BgprpkiaspaTimeseriesResponse struct {
	Meta   BgprpkiaspaTimeseriesResponseMeta   `json:"meta,required"`
	Serie0 BgprpkiaspaTimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   bgprpkiaspaTimeseriesResponseJSON   `json:"-"`
}

// bgprpkiaspaTimeseriesResponseJSON contains the JSON metadata for the struct
// [BgprpkiaspaTimeseriesResponse]
type bgprpkiaspaTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type BgprpkiaspaTimeseriesResponseMeta struct {
	// Timestamp of the underlying data.
	DataTime time.Time `json:"dataTime,required" format:"date-time"`
	// Timestamp when the query was executed.
	QueryTime time.Time                             `json:"queryTime,required" format:"date-time"`
	JSON      bgprpkiaspaTimeseriesResponseMetaJSON `json:"-"`
}

// bgprpkiaspaTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [BgprpkiaspaTimeseriesResponseMeta]
type bgprpkiaspaTimeseriesResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BgprpkiaspaTimeseriesResponseSerie0 struct {
	Timestamps []time.Time                             `json:"timestamps,required" format:"date-time"`
	Values     []string                                `json:"values,required"`
	JSON       bgprpkiaspaTimeseriesResponseSerie0JSON `json:"-"`
}

// bgprpkiaspaTimeseriesResponseSerie0JSON contains the JSON metadata for the
// struct [BgprpkiaspaTimeseriesResponseSerie0]
type bgprpkiaspaTimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaTimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaTimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type BGPRPKIASPAChangesParams struct {
	// Filter changes involving this ASN (as customer or provider).
	ASN param.Field[int64] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[BgprpkiaspaChangesParamsFormat] `query:"format"`
	// Include ASN metadata (name, country) in response.
	IncludeASNInfo param.Field[bool] `query:"includeAsnInfo"`
}

// URLQuery serializes [BGPRPKIASPAChangesParams]'s query parameters as
// `url.Values`.
func (r BGPRPKIASPAChangesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type BgprpkiaspaChangesParamsFormat string

const (
	BgprpkiaspaChangesParamsFormatJson BgprpkiaspaChangesParamsFormat = "JSON"
	BgprpkiaspaChangesParamsFormatCsv  BgprpkiaspaChangesParamsFormat = "CSV"
)

func (r BgprpkiaspaChangesParamsFormat) IsKnown() bool {
	switch r {
	case BgprpkiaspaChangesParamsFormatJson, BgprpkiaspaChangesParamsFormatCsv:
		return true
	}
	return false
}

type BgprpkiaspaChangesResponseEnvelope struct {
	Result  BgprpkiaspaChangesResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    bgprpkiaspaChangesResponseEnvelopeJSON `json:"-"`
}

// bgprpkiaspaChangesResponseEnvelopeJSON contains the JSON metadata for the struct
// [BgprpkiaspaChangesResponseEnvelope]
type bgprpkiaspaChangesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaChangesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaChangesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BGPRPKIASPASnapshotParams struct {
	// Filter by customer ASN (the ASN publishing the ASPA object).
	CustomerASN param.Field[int64] `query:"customerAsn"`
	// Filters results by the specified datetime (ISO 8601).
	Date param.Field[time.Time] `query:"date" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[BgprpkiaspaSnapshotParamsFormat] `query:"format"`
	// Include ASN metadata (name, country) in response.
	IncludeASNInfo param.Field[bool] `query:"includeAsnInfo"`
	// Filter by provider ASN (an authorized upstream provider in ASPA objects).
	ProviderASN param.Field[int64] `query:"providerAsn"`
}

// URLQuery serializes [BGPRPKIASPASnapshotParams]'s query parameters as
// `url.Values`.
func (r BGPRPKIASPASnapshotParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type BgprpkiaspaSnapshotParamsFormat string

const (
	BgprpkiaspaSnapshotParamsFormatJson BgprpkiaspaSnapshotParamsFormat = "JSON"
	BgprpkiaspaSnapshotParamsFormatCsv  BgprpkiaspaSnapshotParamsFormat = "CSV"
)

func (r BgprpkiaspaSnapshotParamsFormat) IsKnown() bool {
	switch r {
	case BgprpkiaspaSnapshotParamsFormatJson, BgprpkiaspaSnapshotParamsFormatCsv:
		return true
	}
	return false
}

type BgprpkiaspaSnapshotResponseEnvelope struct {
	Result  BgprpkiaspaSnapshotResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    bgprpkiaspaSnapshotResponseEnvelopeJSON `json:"-"`
}

// bgprpkiaspaSnapshotResponseEnvelopeJSON contains the JSON metadata for the
// struct [BgprpkiaspaSnapshotResponseEnvelope]
type bgprpkiaspaSnapshotResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaSnapshotResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaSnapshotResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BGPRPKIASPATimeseriesParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[BgprpkiaspaTimeseriesParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 location
	// codes.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filter by Regional Internet Registry (RIR). Multiple RIRs generate multiple
	// series.
	Rir param.Field[[]BgprpkiaspaTimeseriesParamsRir] `query:"rir"`
}

// URLQuery serializes [BGPRPKIASPATimeseriesParams]'s query parameters as
// `url.Values`.
func (r BGPRPKIASPATimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type BgprpkiaspaTimeseriesParamsFormat string

const (
	BgprpkiaspaTimeseriesParamsFormatJson BgprpkiaspaTimeseriesParamsFormat = "JSON"
	BgprpkiaspaTimeseriesParamsFormatCsv  BgprpkiaspaTimeseriesParamsFormat = "CSV"
)

func (r BgprpkiaspaTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case BgprpkiaspaTimeseriesParamsFormatJson, BgprpkiaspaTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type BgprpkiaspaTimeseriesParamsRir string

const (
	BgprpkiaspaTimeseriesParamsRirRipeNcc BgprpkiaspaTimeseriesParamsRir = "RIPE_NCC"
	BgprpkiaspaTimeseriesParamsRirArin    BgprpkiaspaTimeseriesParamsRir = "ARIN"
	BgprpkiaspaTimeseriesParamsRirApnic   BgprpkiaspaTimeseriesParamsRir = "APNIC"
	BgprpkiaspaTimeseriesParamsRirLacnic  BgprpkiaspaTimeseriesParamsRir = "LACNIC"
	BgprpkiaspaTimeseriesParamsRirAfrinic BgprpkiaspaTimeseriesParamsRir = "AFRINIC"
)

func (r BgprpkiaspaTimeseriesParamsRir) IsKnown() bool {
	switch r {
	case BgprpkiaspaTimeseriesParamsRirRipeNcc, BgprpkiaspaTimeseriesParamsRirArin, BgprpkiaspaTimeseriesParamsRirApnic, BgprpkiaspaTimeseriesParamsRirLacnic, BgprpkiaspaTimeseriesParamsRirAfrinic:
		return true
	}
	return false
}

type BgprpkiaspaTimeseriesResponseEnvelope struct {
	Result  BgprpkiaspaTimeseriesResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    bgprpkiaspaTimeseriesResponseEnvelopeJSON `json:"-"`
}

// bgprpkiaspaTimeseriesResponseEnvelopeJSON contains the JSON metadata for the
// struct [BgprpkiaspaTimeseriesResponseEnvelope]
type bgprpkiaspaTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BgprpkiaspaTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgprpkiaspaTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

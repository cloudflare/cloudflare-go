// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
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

// EntityASNService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityASNService] method instead.
type EntityASNService struct {
	Options []option.RequestOption
}

// NewEntityASNService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEntityASNService(opts ...option.RequestOption) (r *EntityASNService) {
	r = &EntityASNService{}
	r.Options = opts
	return
}

// Retrieves a list of autonomous systems.
func (r *EntityASNService) List(ctx context.Context, query EntityASNListParams, opts ...option.RequestOption) (res *EntityASNListResponse, err error) {
	var env EntityASNListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/entities/asns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves Internet Routing Registry AS-SETs that an AS is a member of.
func (r *EntityASNService) AsSet(ctx context.Context, asn int64, query EntityASNAsSetParams, opts ...option.RequestOption) (res *EntityASNAsSetResponse, err error) {
	var env EntityASNAsSetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/entities/asns/%v/as_set", asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a ranked list of Autonomous Systems based on their presence in the
// Cloudflare Botnet Threat Feed. Rankings can be sorted by offense count or number
// of bad IPs. Optionally compare to a previous date to see rank changes.
func (r *EntityASNService) BotnetThreatFeed(ctx context.Context, query EntityASNBotnetThreatFeedParams, opts ...option.RequestOption) (res *EntityASNBotnetThreatFeedResponse, err error) {
	var env EntityASNBotnetThreatFeedResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/entities/asns/botnet_threat_feed"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the requested autonomous system information. (A confidence level below
// `5` indicates a low level of confidence in the traffic data - normally this
// happens because Cloudflare has a small amount of traffic from/to this AS).
// Population estimates come from APNIC (refer to https://labs.apnic.net/?p=526).
func (r *EntityASNService) Get(ctx context.Context, asn int64, query EntityASNGetParams, opts ...option.RequestOption) (res *EntityASNGetResponse, err error) {
	var env EntityASNGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/entities/asns/%v", asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the requested autonomous system information based on IP address.
// Population estimates come from APNIC (refer to https://labs.apnic.net/?p=526).
func (r *EntityASNService) IP(ctx context.Context, query EntityASNIPParams, opts ...option.RequestOption) (res *EntityAsnipResponse, err error) {
	var env EntityAsnipResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "radar/entities/asns/ip"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves AS-level relationship for given networks.
func (r *EntityASNService) Rel(ctx context.Context, asn int64, query EntityASNRelParams, opts ...option.RequestOption) (res *EntityASNRelResponse, err error) {
	var env EntityASNRelResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/entities/asns/%v/rel", asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EntityASNListResponse struct {
	ASNs []EntityASNListResponseASN `json:"asns,required"`
	JSON entityASNListResponseJSON  `json:"-"`
}

// entityASNListResponseJSON contains the JSON metadata for the struct
// [EntityASNListResponse]
type entityASNListResponseJSON struct {
	ASNs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNListResponseJSON) RawJSON() string {
	return r.raw
}

type EntityASNListResponseASN struct {
	ASN         int64                        `json:"asn,required"`
	Country     string                       `json:"country,required"`
	CountryName string                       `json:"countryName,required"`
	Name        string                       `json:"name,required"`
	Aka         string                       `json:"aka"`
	OrgName     string                       `json:"orgName"`
	Website     string                       `json:"website"`
	JSON        entityASNListResponseASNJSON `json:"-"`
}

// entityASNListResponseASNJSON contains the JSON metadata for the struct
// [EntityASNListResponseASN]
type entityASNListResponseASNJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	CountryName apijson.Field
	Name        apijson.Field
	Aka         apijson.Field
	OrgName     apijson.Field
	Website     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNListResponseASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNListResponseASNJSON) RawJSON() string {
	return r.raw
}

type EntityASNAsSetResponse struct {
	AsSets []EntityASNAsSetResponseAsSet `json:"as_sets,required"`
	// Paths from the AS-SET that include the given AS to its upstreams recursively
	Paths [][]string                 `json:"paths,required"`
	JSON  entityASNAsSetResponseJSON `json:"-"`
}

// entityASNAsSetResponseJSON contains the JSON metadata for the struct
// [EntityASNAsSetResponse]
type entityASNAsSetResponseJSON struct {
	AsSets      apijson.Field
	Paths       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNAsSetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNAsSetResponseJSON) RawJSON() string {
	return r.raw
}

type EntityASNAsSetResponseAsSet struct {
	// The number of AS members in the AS-SET
	AsMembersCount int64 `json:"as_members_count,required"`
	// The number of AS-SET members in the AS-SET
	AsSetMembersCount int64 `json:"as_set_members_count,required"`
	// The number of recursive upstream AS-SETs
	AsSetUpstreamsCount int64 `json:"as_set_upstreams_count,required"`
	// The number of unique ASNs in the AS-SETs recursive downstream
	ASNConeSize int64 `json:"asn_cone_size,required"`
	// The IRR sources of the AS-SET
	IrrSources []string `json:"irr_sources,required"`
	// The name of the AS-SET
	Name string `json:"name,required"`
	// The AS number following hierarchical AS-SET name
	HierarchicalASN int64 `json:"hierarchical_asn"`
	// The inferred AS number of the AS-SET
	InferredASN int64 `json:"inferred_asn"`
	// The AS number matching PeeringDB record
	PeeringdbASN int64                           `json:"peeringdb_asn"`
	JSON         entityASNAsSetResponseAsSetJSON `json:"-"`
}

// entityASNAsSetResponseAsSetJSON contains the JSON metadata for the struct
// [EntityASNAsSetResponseAsSet]
type entityASNAsSetResponseAsSetJSON struct {
	AsMembersCount      apijson.Field
	AsSetMembersCount   apijson.Field
	AsSetUpstreamsCount apijson.Field
	ASNConeSize         apijson.Field
	IrrSources          apijson.Field
	Name                apijson.Field
	HierarchicalASN     apijson.Field
	InferredASN         apijson.Field
	PeeringdbASN        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *EntityASNAsSetResponseAsSet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNAsSetResponseAsSetJSON) RawJSON() string {
	return r.raw
}

type EntityASNBotnetThreatFeedResponse struct {
	Ases []EntityASNBotnetThreatFeedResponseAse `json:"ases,required"`
	Meta EntityASNBotnetThreatFeedResponseMeta  `json:"meta,required"`
	JSON entityASNBotnetThreatFeedResponseJSON  `json:"-"`
}

// entityASNBotnetThreatFeedResponseJSON contains the JSON metadata for the struct
// [EntityASNBotnetThreatFeedResponse]
type entityASNBotnetThreatFeedResponseJSON struct {
	Ases        apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNBotnetThreatFeedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNBotnetThreatFeedResponseJSON) RawJSON() string {
	return r.raw
}

type EntityASNBotnetThreatFeedResponseAse struct {
	ASN        int64                                    `json:"asn,required"`
	Country    string                                   `json:"country,required"`
	Name       string                                   `json:"name,required"`
	Rank       int64                                    `json:"rank,required"`
	RankChange int64                                    `json:"rankChange"`
	JSON       entityASNBotnetThreatFeedResponseAseJSON `json:"-"`
}

// entityASNBotnetThreatFeedResponseAseJSON contains the JSON metadata for the
// struct [EntityASNBotnetThreatFeedResponseAse]
type entityASNBotnetThreatFeedResponseAseJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	Name        apijson.Field
	Rank        apijson.Field
	RankChange  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNBotnetThreatFeedResponseAse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNBotnetThreatFeedResponseAseJSON) RawJSON() string {
	return r.raw
}

type EntityASNBotnetThreatFeedResponseMeta struct {
	Date        string                                    `json:"date,required"`
	Total       int64                                     `json:"total,required"`
	CompareDate string                                    `json:"compareDate"`
	JSON        entityASNBotnetThreatFeedResponseMetaJSON `json:"-"`
}

// entityASNBotnetThreatFeedResponseMetaJSON contains the JSON metadata for the
// struct [EntityASNBotnetThreatFeedResponseMeta]
type entityASNBotnetThreatFeedResponseMetaJSON struct {
	Date        apijson.Field
	Total       apijson.Field
	CompareDate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNBotnetThreatFeedResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNBotnetThreatFeedResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EntityASNGetResponse struct {
	ASN  EntityASNGetResponseASN  `json:"asn,required"`
	JSON entityASNGetResponseJSON `json:"-"`
}

// entityASNGetResponseJSON contains the JSON metadata for the struct
// [EntityASNGetResponse]
type entityASNGetResponseJSON struct {
	ASN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNGetResponseJSON) RawJSON() string {
	return r.raw
}

type EntityASNGetResponseASN struct {
	ASN             int64                                 `json:"asn,required"`
	ConfidenceLevel int64                                 `json:"confidenceLevel,required"`
	Country         string                                `json:"country,required"`
	CountryName     string                                `json:"countryName,required"`
	EstimatedUsers  EntityASNGetResponseASNEstimatedUsers `json:"estimatedUsers,required"`
	Name            string                                `json:"name,required"`
	OrgName         string                                `json:"orgName,required"`
	Related         []EntityASNGetResponseASNRelated      `json:"related,required"`
	// Regional Internet Registry.
	Source  string                      `json:"source,required"`
	Website string                      `json:"website,required"`
	Aka     string                      `json:"aka"`
	JSON    entityASNGetResponseASNJSON `json:"-"`
}

// entityASNGetResponseASNJSON contains the JSON metadata for the struct
// [EntityASNGetResponseASN]
type entityASNGetResponseASNJSON struct {
	ASN             apijson.Field
	ConfidenceLevel apijson.Field
	Country         apijson.Field
	CountryName     apijson.Field
	EstimatedUsers  apijson.Field
	Name            apijson.Field
	OrgName         apijson.Field
	Related         apijson.Field
	Source          apijson.Field
	Website         apijson.Field
	Aka             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EntityASNGetResponseASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNGetResponseASNJSON) RawJSON() string {
	return r.raw
}

type EntityASNGetResponseASNEstimatedUsers struct {
	Locations []EntityASNGetResponseASNEstimatedUsersLocation `json:"locations,required"`
	// Total estimated users.
	EstimatedUsers int64                                     `json:"estimatedUsers"`
	JSON           entityASNGetResponseASNEstimatedUsersJSON `json:"-"`
}

// entityASNGetResponseASNEstimatedUsersJSON contains the JSON metadata for the
// struct [EntityASNGetResponseASNEstimatedUsers]
type entityASNGetResponseASNEstimatedUsersJSON struct {
	Locations      apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityASNGetResponseASNEstimatedUsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNGetResponseASNEstimatedUsersJSON) RawJSON() string {
	return r.raw
}

type EntityASNGetResponseASNEstimatedUsersLocation struct {
	LocationAlpha2 string `json:"locationAlpha2,required"`
	LocationName   string `json:"locationName,required"`
	// Estimated users per location.
	EstimatedUsers int64                                             `json:"estimatedUsers"`
	JSON           entityASNGetResponseASNEstimatedUsersLocationJSON `json:"-"`
}

// entityASNGetResponseASNEstimatedUsersLocationJSON contains the JSON metadata for
// the struct [EntityASNGetResponseASNEstimatedUsersLocation]
type entityASNGetResponseASNEstimatedUsersLocationJSON struct {
	LocationAlpha2 apijson.Field
	LocationName   apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityASNGetResponseASNEstimatedUsersLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNGetResponseASNEstimatedUsersLocationJSON) RawJSON() string {
	return r.raw
}

type EntityASNGetResponseASNRelated struct {
	ASN  int64  `json:"asn,required"`
	Name string `json:"name,required"`
	Aka  string `json:"aka"`
	// Total estimated users.
	EstimatedUsers int64                              `json:"estimatedUsers"`
	JSON           entityASNGetResponseASNRelatedJSON `json:"-"`
}

// entityASNGetResponseASNRelatedJSON contains the JSON metadata for the struct
// [EntityASNGetResponseASNRelated]
type entityASNGetResponseASNRelatedJSON struct {
	ASN            apijson.Field
	Name           apijson.Field
	Aka            apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityASNGetResponseASNRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNGetResponseASNRelatedJSON) RawJSON() string {
	return r.raw
}

type EntityAsnipResponse struct {
	ASN  EntityAsnipResponseASN  `json:"asn,required"`
	JSON entityAsnipResponseJSON `json:"-"`
}

// entityAsnipResponseJSON contains the JSON metadata for the struct
// [EntityAsnipResponse]
type entityAsnipResponseJSON struct {
	ASN         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityAsnipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityAsnipResponseJSON) RawJSON() string {
	return r.raw
}

type EntityAsnipResponseASN struct {
	ASN            int64                                `json:"asn,required"`
	Country        string                               `json:"country,required"`
	CountryName    string                               `json:"countryName,required"`
	EstimatedUsers EntityAsnipResponseASNEstimatedUsers `json:"estimatedUsers,required"`
	Name           string                               `json:"name,required"`
	OrgName        string                               `json:"orgName,required"`
	Related        []EntityAsnipResponseASNRelated      `json:"related,required"`
	// Regional Internet Registry.
	Source  string                     `json:"source,required"`
	Website string                     `json:"website,required"`
	Aka     string                     `json:"aka"`
	JSON    entityAsnipResponseASNJSON `json:"-"`
}

// entityAsnipResponseASNJSON contains the JSON metadata for the struct
// [EntityAsnipResponseASN]
type entityAsnipResponseASNJSON struct {
	ASN            apijson.Field
	Country        apijson.Field
	CountryName    apijson.Field
	EstimatedUsers apijson.Field
	Name           apijson.Field
	OrgName        apijson.Field
	Related        apijson.Field
	Source         apijson.Field
	Website        apijson.Field
	Aka            apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityAsnipResponseASN) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityAsnipResponseASNJSON) RawJSON() string {
	return r.raw
}

type EntityAsnipResponseASNEstimatedUsers struct {
	Locations []EntityAsnipResponseASNEstimatedUsersLocation `json:"locations,required"`
	// Total estimated users.
	EstimatedUsers int64                                    `json:"estimatedUsers"`
	JSON           entityAsnipResponseASNEstimatedUsersJSON `json:"-"`
}

// entityAsnipResponseASNEstimatedUsersJSON contains the JSON metadata for the
// struct [EntityAsnipResponseASNEstimatedUsers]
type entityAsnipResponseASNEstimatedUsersJSON struct {
	Locations      apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityAsnipResponseASNEstimatedUsers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityAsnipResponseASNEstimatedUsersJSON) RawJSON() string {
	return r.raw
}

type EntityAsnipResponseASNEstimatedUsersLocation struct {
	LocationAlpha2 string `json:"locationAlpha2,required"`
	LocationName   string `json:"locationName,required"`
	// Estimated users per location.
	EstimatedUsers int64                                            `json:"estimatedUsers"`
	JSON           entityAsnipResponseASNEstimatedUsersLocationJSON `json:"-"`
}

// entityAsnipResponseASNEstimatedUsersLocationJSON contains the JSON metadata for
// the struct [EntityAsnipResponseASNEstimatedUsersLocation]
type entityAsnipResponseASNEstimatedUsersLocationJSON struct {
	LocationAlpha2 apijson.Field
	LocationName   apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityAsnipResponseASNEstimatedUsersLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityAsnipResponseASNEstimatedUsersLocationJSON) RawJSON() string {
	return r.raw
}

type EntityAsnipResponseASNRelated struct {
	ASN  int64  `json:"asn,required"`
	Name string `json:"name,required"`
	Aka  string `json:"aka"`
	// Total estimated users.
	EstimatedUsers int64                             `json:"estimatedUsers"`
	JSON           entityAsnipResponseASNRelatedJSON `json:"-"`
}

// entityAsnipResponseASNRelatedJSON contains the JSON metadata for the struct
// [EntityAsnipResponseASNRelated]
type entityAsnipResponseASNRelatedJSON struct {
	ASN            apijson.Field
	Name           apijson.Field
	Aka            apijson.Field
	EstimatedUsers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *EntityAsnipResponseASNRelated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityAsnipResponseASNRelatedJSON) RawJSON() string {
	return r.raw
}

type EntityASNRelResponse struct {
	Meta EntityASNRelResponseMeta  `json:"meta,required"`
	Rels []EntityASNRelResponseRel `json:"rels,required"`
	JSON entityASNRelResponseJSON  `json:"-"`
}

// entityASNRelResponseJSON contains the JSON metadata for the struct
// [EntityASNRelResponse]
type entityASNRelResponseJSON struct {
	Meta        apijson.Field
	Rels        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNRelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNRelResponseJSON) RawJSON() string {
	return r.raw
}

type EntityASNRelResponseMeta struct {
	DataTime   string                       `json:"data_time,required"`
	QueryTime  string                       `json:"query_time,required"`
	TotalPeers int64                        `json:"total_peers,required"`
	JSON       entityASNRelResponseMetaJSON `json:"-"`
}

// entityASNRelResponseMetaJSON contains the JSON metadata for the struct
// [EntityASNRelResponseMeta]
type entityASNRelResponseMetaJSON struct {
	DataTime    apijson.Field
	QueryTime   apijson.Field
	TotalPeers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNRelResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNRelResponseMetaJSON) RawJSON() string {
	return r.raw
}

type EntityASNRelResponseRel struct {
	Asn1        int64                       `json:"asn1,required"`
	Asn1Country string                      `json:"asn1_country,required"`
	Asn1Name    string                      `json:"asn1_name,required"`
	Asn2        int64                       `json:"asn2,required"`
	Asn2Country string                      `json:"asn2_country,required"`
	Asn2Name    string                      `json:"asn2_name,required"`
	Rel         string                      `json:"rel,required"`
	JSON        entityASNRelResponseRelJSON `json:"-"`
}

// entityASNRelResponseRelJSON contains the JSON metadata for the struct
// [EntityASNRelResponseRel]
type entityASNRelResponseRelJSON struct {
	Asn1        apijson.Field
	Asn1Country apijson.Field
	Asn1Name    apijson.Field
	Asn2        apijson.Field
	Asn2Country apijson.Field
	Asn2Name    apijson.Field
	Rel         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNRelResponseRel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNRelResponseRelJSON) RawJSON() string {
	return r.raw
}

type EntityASNListParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list.
	ASN param.Field[string] `query:"asn"`
	// Format in which results will be returned.
	Format param.Field[EntityASNListParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify an alpha-2 location code.
	Location param.Field[string] `query:"location"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
	// Specifies the metric to order the ASNs by.
	OrderBy param.Field[EntityASNListParamsOrderBy] `query:"orderBy"`
}

// URLQuery serializes [EntityASNListParams]'s query parameters as `url.Values`.
func (r EntityASNListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type EntityASNListParamsFormat string

const (
	EntityASNListParamsFormatJson EntityASNListParamsFormat = "JSON"
	EntityASNListParamsFormatCsv  EntityASNListParamsFormat = "CSV"
)

func (r EntityASNListParamsFormat) IsKnown() bool {
	switch r {
	case EntityASNListParamsFormatJson, EntityASNListParamsFormatCsv:
		return true
	}
	return false
}

// Specifies the metric to order the ASNs by.
type EntityASNListParamsOrderBy string

const (
	EntityASNListParamsOrderByASN        EntityASNListParamsOrderBy = "ASN"
	EntityASNListParamsOrderByPopulation EntityASNListParamsOrderBy = "POPULATION"
)

func (r EntityASNListParamsOrderBy) IsKnown() bool {
	switch r {
	case EntityASNListParamsOrderByASN, EntityASNListParamsOrderByPopulation:
		return true
	}
	return false
}

type EntityASNListResponseEnvelope struct {
	Result  EntityASNListResponse             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    entityASNListResponseEnvelopeJSON `json:"-"`
}

// entityASNListResponseEnvelopeJSON contains the JSON metadata for the struct
// [EntityASNListResponseEnvelope]
type entityASNListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EntityASNAsSetParams struct {
	// Format in which results will be returned.
	Format param.Field[EntityASNAsSetParamsFormat] `query:"format"`
}

// URLQuery serializes [EntityASNAsSetParams]'s query parameters as `url.Values`.
func (r EntityASNAsSetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type EntityASNAsSetParamsFormat string

const (
	EntityASNAsSetParamsFormatJson EntityASNAsSetParamsFormat = "JSON"
	EntityASNAsSetParamsFormatCsv  EntityASNAsSetParamsFormat = "CSV"
)

func (r EntityASNAsSetParamsFormat) IsKnown() bool {
	switch r {
	case EntityASNAsSetParamsFormatJson, EntityASNAsSetParamsFormatCsv:
		return true
	}
	return false
}

type EntityASNAsSetResponseEnvelope struct {
	Result  EntityASNAsSetResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    entityASNAsSetResponseEnvelopeJSON `json:"-"`
}

// entityASNAsSetResponseEnvelopeJSON contains the JSON metadata for the struct
// [EntityASNAsSetResponseEnvelope]
type entityASNAsSetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNAsSetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNAsSetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EntityASNBotnetThreatFeedParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Relative date range for rank change comparison (e.g., "1d", "7d", "30d").
	CompareDateRange param.Field[string] `query:"compareDateRange"`
	// The date to retrieve (YYYY-MM-DD format). If not specified, returns the most
	// recent available data. Note: This is the date the report was generated. The
	// report is generated from information collected from the previous day (e.g., the
	// 2026-02-23 entry contains data from 2026-02-22).
	Date param.Field[time.Time] `query:"date" format:"date"`
	// Format in which results will be returned.
	Format param.Field[EntityASNBotnetThreatFeedParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify an alpha-2 location code.
	Location param.Field[string] `query:"location"`
	// Metric to rank ASNs by.
	Metric param.Field[EntityASNBotnetThreatFeedParamsMetric] `query:"metric"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
	// Sort order.
	SortOrder param.Field[EntityASNBotnetThreatFeedParamsSortOrder] `query:"sortOrder"`
}

// URLQuery serializes [EntityASNBotnetThreatFeedParams]'s query parameters as
// `url.Values`.
func (r EntityASNBotnetThreatFeedParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type EntityASNBotnetThreatFeedParamsFormat string

const (
	EntityASNBotnetThreatFeedParamsFormatJson EntityASNBotnetThreatFeedParamsFormat = "JSON"
	EntityASNBotnetThreatFeedParamsFormatCsv  EntityASNBotnetThreatFeedParamsFormat = "CSV"
)

func (r EntityASNBotnetThreatFeedParamsFormat) IsKnown() bool {
	switch r {
	case EntityASNBotnetThreatFeedParamsFormatJson, EntityASNBotnetThreatFeedParamsFormatCsv:
		return true
	}
	return false
}

// Metric to rank ASNs by.
type EntityASNBotnetThreatFeedParamsMetric string

const (
	EntityASNBotnetThreatFeedParamsMetricOffenseCount         EntityASNBotnetThreatFeedParamsMetric = "OFFENSE_COUNT"
	EntityASNBotnetThreatFeedParamsMetricNumberOfOffendingIPs EntityASNBotnetThreatFeedParamsMetric = "NUMBER_OF_OFFENDING_IPS"
)

func (r EntityASNBotnetThreatFeedParamsMetric) IsKnown() bool {
	switch r {
	case EntityASNBotnetThreatFeedParamsMetricOffenseCount, EntityASNBotnetThreatFeedParamsMetricNumberOfOffendingIPs:
		return true
	}
	return false
}

// Sort order.
type EntityASNBotnetThreatFeedParamsSortOrder string

const (
	EntityASNBotnetThreatFeedParamsSortOrderAsc  EntityASNBotnetThreatFeedParamsSortOrder = "ASC"
	EntityASNBotnetThreatFeedParamsSortOrderDesc EntityASNBotnetThreatFeedParamsSortOrder = "DESC"
)

func (r EntityASNBotnetThreatFeedParamsSortOrder) IsKnown() bool {
	switch r {
	case EntityASNBotnetThreatFeedParamsSortOrderAsc, EntityASNBotnetThreatFeedParamsSortOrderDesc:
		return true
	}
	return false
}

type EntityASNBotnetThreatFeedResponseEnvelope struct {
	Result  EntityASNBotnetThreatFeedResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    entityASNBotnetThreatFeedResponseEnvelopeJSON `json:"-"`
}

// entityASNBotnetThreatFeedResponseEnvelopeJSON contains the JSON metadata for the
// struct [EntityASNBotnetThreatFeedResponseEnvelope]
type entityASNBotnetThreatFeedResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNBotnetThreatFeedResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNBotnetThreatFeedResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EntityASNGetParams struct {
	// Format in which results will be returned.
	Format param.Field[EntityASNGetParamsFormat] `query:"format"`
}

// URLQuery serializes [EntityASNGetParams]'s query parameters as `url.Values`.
func (r EntityASNGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type EntityASNGetParamsFormat string

const (
	EntityASNGetParamsFormatJson EntityASNGetParamsFormat = "JSON"
	EntityASNGetParamsFormatCsv  EntityASNGetParamsFormat = "CSV"
)

func (r EntityASNGetParamsFormat) IsKnown() bool {
	switch r {
	case EntityASNGetParamsFormatJson, EntityASNGetParamsFormatCsv:
		return true
	}
	return false
}

type EntityASNGetResponseEnvelope struct {
	Result  EntityASNGetResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    entityASNGetResponseEnvelopeJSON `json:"-"`
}

// entityASNGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [EntityASNGetResponseEnvelope]
type entityASNGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EntityASNIPParams struct {
	// IP address.
	IP param.Field[string] `query:"ip,required" format:"ip"`
	// Format in which results will be returned.
	Format param.Field[EntityAsnipParamsFormat] `query:"format"`
}

// URLQuery serializes [EntityASNIPParams]'s query parameters as `url.Values`.
func (r EntityASNIPParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type EntityAsnipParamsFormat string

const (
	EntityAsnipParamsFormatJson EntityAsnipParamsFormat = "JSON"
	EntityAsnipParamsFormatCsv  EntityAsnipParamsFormat = "CSV"
)

func (r EntityAsnipParamsFormat) IsKnown() bool {
	switch r {
	case EntityAsnipParamsFormatJson, EntityAsnipParamsFormatCsv:
		return true
	}
	return false
}

type EntityAsnipResponseEnvelope struct {
	Result  EntityAsnipResponse             `json:"result,required"`
	Success bool                            `json:"success,required"`
	JSON    entityAsnipResponseEnvelopeJSON `json:"-"`
}

// entityAsnipResponseEnvelopeJSON contains the JSON metadata for the struct
// [EntityAsnipResponseEnvelope]
type entityAsnipResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityAsnipResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityAsnipResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EntityASNRelParams struct {
	// Retrieves the AS relationship of ASN2 with respect to the given ASN.
	Asn2 param.Field[int64] `query:"asn2"`
	// Format in which results will be returned.
	Format param.Field[EntityASNRelParamsFormat] `query:"format"`
}

// URLQuery serializes [EntityASNRelParams]'s query parameters as `url.Values`.
func (r EntityASNRelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type EntityASNRelParamsFormat string

const (
	EntityASNRelParamsFormatJson EntityASNRelParamsFormat = "JSON"
	EntityASNRelParamsFormatCsv  EntityASNRelParamsFormat = "CSV"
)

func (r EntityASNRelParamsFormat) IsKnown() bool {
	switch r {
	case EntityASNRelParamsFormatJson, EntityASNRelParamsFormatCsv:
		return true
	}
	return false
}

type EntityASNRelResponseEnvelope struct {
	Result  EntityASNRelResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    entityASNRelResponseEnvelopeJSON `json:"-"`
}

// entityASNRelResponseEnvelopeJSON contains the JSON metadata for the struct
// [EntityASNRelResponseEnvelope]
type entityASNRelResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityASNRelResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityASNRelResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

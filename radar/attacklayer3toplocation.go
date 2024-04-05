// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AttackLayer3TopLocationService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAttackLayer3TopLocationService] method instead.
type AttackLayer3TopLocationService struct {
	Options []option.RequestOption
}

// NewAttackLayer3TopLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAttackLayer3TopLocationService(opts ...option.RequestOption) (r *AttackLayer3TopLocationService) {
	r = &AttackLayer3TopLocationService{}
	r.Options = opts
	return
}

// Get the origin locations of attacks.
func (r *AttackLayer3TopLocationService) Origin(ctx context.Context, query AttackLayer3TopLocationOriginParams, opts ...option.RequestOption) (res *AttackLayer3TopLocationOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer3TopLocationOriginResponseEnvelope
	path := "radar/attacks/layer3/top/locations/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the target locations of attacks.
func (r *AttackLayer3TopLocationService) Target(ctx context.Context, query AttackLayer3TopLocationTargetParams, opts ...option.RequestOption) (res *AttackLayer3TopLocationTargetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer3TopLocationTargetResponseEnvelope
	path := "radar/attacks/layer3/top/locations/target"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer3TopLocationOriginResponse struct {
	Meta AttackLayer3TopLocationOriginResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer3TopLocationOriginResponseTop0 `json:"top_0,required"`
	JSON attackLayer3TopLocationOriginResponseJSON   `json:"-"`
}

// attackLayer3TopLocationOriginResponseJSON contains the JSON metadata for the
// struct [AttackLayer3TopLocationOriginResponse]
type attackLayer3TopLocationOriginResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationOriginResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationOriginResponseMeta struct {
	DateRange      []UnnamedSchemaRefBaac9d7da12de53e99142f8ecd3982e5      `json:"dateRange,required"`
	LastUpdated    string                                                  `json:"lastUpdated,required"`
	ConfidenceInfo AttackLayer3TopLocationOriginResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer3TopLocationOriginResponseMetaJSON           `json:"-"`
}

// attackLayer3TopLocationOriginResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3TopLocationOriginResponseMeta]
type attackLayer3TopLocationOriginResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TopLocationOriginResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationOriginResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationOriginResponseMetaConfidenceInfo struct {
	Annotations []UnnamedSchemaRefB5f3bd1840490bc487ffef84567807b1          `json:"annotations"`
	Level       int64                                                       `json:"level"`
	JSON        attackLayer3TopLocationOriginResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TopLocationOriginResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer3TopLocationOriginResponseMetaConfidenceInfo]
type attackLayer3TopLocationOriginResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationOriginResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationOriginResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationOriginResponseTop0 struct {
	OriginCountryAlpha2 string                                        `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                        `json:"originCountryName,required"`
	Rank                float64                                       `json:"rank,required"`
	Value               string                                        `json:"value,required"`
	JSON                attackLayer3TopLocationOriginResponseTop0JSON `json:"-"`
}

// attackLayer3TopLocationOriginResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer3TopLocationOriginResponseTop0]
type attackLayer3TopLocationOriginResponseTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Rank                apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AttackLayer3TopLocationOriginResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationOriginResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationTargetResponse struct {
	Meta AttackLayer3TopLocationTargetResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer3TopLocationTargetResponseTop0 `json:"top_0,required"`
	JSON attackLayer3TopLocationTargetResponseJSON   `json:"-"`
}

// attackLayer3TopLocationTargetResponseJSON contains the JSON metadata for the
// struct [AttackLayer3TopLocationTargetResponse]
type attackLayer3TopLocationTargetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationTargetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationTargetResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationTargetResponseMeta struct {
	DateRange      []UnnamedSchemaRefBaac9d7da12de53e99142f8ecd3982e5      `json:"dateRange,required"`
	LastUpdated    string                                                  `json:"lastUpdated,required"`
	ConfidenceInfo AttackLayer3TopLocationTargetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer3TopLocationTargetResponseMetaJSON           `json:"-"`
}

// attackLayer3TopLocationTargetResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3TopLocationTargetResponseMeta]
type attackLayer3TopLocationTargetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TopLocationTargetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationTargetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationTargetResponseMetaConfidenceInfo struct {
	Annotations []UnnamedSchemaRefB5f3bd1840490bc487ffef84567807b1          `json:"annotations"`
	Level       int64                                                       `json:"level"`
	JSON        attackLayer3TopLocationTargetResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer3TopLocationTargetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer3TopLocationTargetResponseMetaConfidenceInfo]
type attackLayer3TopLocationTargetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationTargetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationTargetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationTargetResponseTop0 struct {
	Rank                float64                                       `json:"rank,required"`
	TargetCountryAlpha2 string                                        `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                        `json:"targetCountryName,required"`
	Value               string                                        `json:"value,required"`
	JSON                attackLayer3TopLocationTargetResponseTop0JSON `json:"-"`
}

// attackLayer3TopLocationTargetResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer3TopLocationTargetResponseTop0]
type attackLayer3TopLocationTargetResponseTop0JSON struct {
	Rank                apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AttackLayer3TopLocationTargetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationTargetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationOriginParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer3TopLocationOriginParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AttackLayer3TopLocationOriginParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3TopLocationOriginParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]AttackLayer3TopLocationOriginParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TopLocationOriginParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3TopLocationOriginParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttackLayer3TopLocationOriginParamsDateRange string

const (
	AttackLayer3TopLocationOriginParamsDateRange1d         AttackLayer3TopLocationOriginParamsDateRange = "1d"
	AttackLayer3TopLocationOriginParamsDateRange2d         AttackLayer3TopLocationOriginParamsDateRange = "2d"
	AttackLayer3TopLocationOriginParamsDateRange7d         AttackLayer3TopLocationOriginParamsDateRange = "7d"
	AttackLayer3TopLocationOriginParamsDateRange14d        AttackLayer3TopLocationOriginParamsDateRange = "14d"
	AttackLayer3TopLocationOriginParamsDateRange28d        AttackLayer3TopLocationOriginParamsDateRange = "28d"
	AttackLayer3TopLocationOriginParamsDateRange12w        AttackLayer3TopLocationOriginParamsDateRange = "12w"
	AttackLayer3TopLocationOriginParamsDateRange24w        AttackLayer3TopLocationOriginParamsDateRange = "24w"
	AttackLayer3TopLocationOriginParamsDateRange52w        AttackLayer3TopLocationOriginParamsDateRange = "52w"
	AttackLayer3TopLocationOriginParamsDateRange1dControl  AttackLayer3TopLocationOriginParamsDateRange = "1dControl"
	AttackLayer3TopLocationOriginParamsDateRange2dControl  AttackLayer3TopLocationOriginParamsDateRange = "2dControl"
	AttackLayer3TopLocationOriginParamsDateRange7dControl  AttackLayer3TopLocationOriginParamsDateRange = "7dControl"
	AttackLayer3TopLocationOriginParamsDateRange14dControl AttackLayer3TopLocationOriginParamsDateRange = "14dControl"
	AttackLayer3TopLocationOriginParamsDateRange28dControl AttackLayer3TopLocationOriginParamsDateRange = "28dControl"
	AttackLayer3TopLocationOriginParamsDateRange12wControl AttackLayer3TopLocationOriginParamsDateRange = "12wControl"
	AttackLayer3TopLocationOriginParamsDateRange24wControl AttackLayer3TopLocationOriginParamsDateRange = "24wControl"
)

func (r AttackLayer3TopLocationOriginParamsDateRange) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationOriginParamsDateRange1d, AttackLayer3TopLocationOriginParamsDateRange2d, AttackLayer3TopLocationOriginParamsDateRange7d, AttackLayer3TopLocationOriginParamsDateRange14d, AttackLayer3TopLocationOriginParamsDateRange28d, AttackLayer3TopLocationOriginParamsDateRange12w, AttackLayer3TopLocationOriginParamsDateRange24w, AttackLayer3TopLocationOriginParamsDateRange52w, AttackLayer3TopLocationOriginParamsDateRange1dControl, AttackLayer3TopLocationOriginParamsDateRange2dControl, AttackLayer3TopLocationOriginParamsDateRange7dControl, AttackLayer3TopLocationOriginParamsDateRange14dControl, AttackLayer3TopLocationOriginParamsDateRange28dControl, AttackLayer3TopLocationOriginParamsDateRange12wControl, AttackLayer3TopLocationOriginParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer3TopLocationOriginParamsFormat string

const (
	AttackLayer3TopLocationOriginParamsFormatJson AttackLayer3TopLocationOriginParamsFormat = "JSON"
	AttackLayer3TopLocationOriginParamsFormatCsv  AttackLayer3TopLocationOriginParamsFormat = "CSV"
)

func (r AttackLayer3TopLocationOriginParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationOriginParamsFormatJson, AttackLayer3TopLocationOriginParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TopLocationOriginParamsIPVersion string

const (
	AttackLayer3TopLocationOriginParamsIPVersionIPv4 AttackLayer3TopLocationOriginParamsIPVersion = "IPv4"
	AttackLayer3TopLocationOriginParamsIPVersionIPv6 AttackLayer3TopLocationOriginParamsIPVersion = "IPv6"
)

func (r AttackLayer3TopLocationOriginParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationOriginParamsIPVersionIPv4, AttackLayer3TopLocationOriginParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3TopLocationOriginParamsProtocol string

const (
	AttackLayer3TopLocationOriginParamsProtocolUdp  AttackLayer3TopLocationOriginParamsProtocol = "UDP"
	AttackLayer3TopLocationOriginParamsProtocolTCP  AttackLayer3TopLocationOriginParamsProtocol = "TCP"
	AttackLayer3TopLocationOriginParamsProtocolIcmp AttackLayer3TopLocationOriginParamsProtocol = "ICMP"
	AttackLayer3TopLocationOriginParamsProtocolGRE  AttackLayer3TopLocationOriginParamsProtocol = "GRE"
)

func (r AttackLayer3TopLocationOriginParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationOriginParamsProtocolUdp, AttackLayer3TopLocationOriginParamsProtocolTCP, AttackLayer3TopLocationOriginParamsProtocolIcmp, AttackLayer3TopLocationOriginParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TopLocationOriginResponseEnvelope struct {
	Result  AttackLayer3TopLocationOriginResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    attackLayer3TopLocationOriginResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TopLocationOriginResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer3TopLocationOriginResponseEnvelope]
type attackLayer3TopLocationOriginResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationOriginResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationOriginResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopLocationTargetParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer3TopLocationTargetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AttackLayer3TopLocationTargetParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3TopLocationTargetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]AttackLayer3TopLocationTargetParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TopLocationTargetParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3TopLocationTargetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttackLayer3TopLocationTargetParamsDateRange string

const (
	AttackLayer3TopLocationTargetParamsDateRange1d         AttackLayer3TopLocationTargetParamsDateRange = "1d"
	AttackLayer3TopLocationTargetParamsDateRange2d         AttackLayer3TopLocationTargetParamsDateRange = "2d"
	AttackLayer3TopLocationTargetParamsDateRange7d         AttackLayer3TopLocationTargetParamsDateRange = "7d"
	AttackLayer3TopLocationTargetParamsDateRange14d        AttackLayer3TopLocationTargetParamsDateRange = "14d"
	AttackLayer3TopLocationTargetParamsDateRange28d        AttackLayer3TopLocationTargetParamsDateRange = "28d"
	AttackLayer3TopLocationTargetParamsDateRange12w        AttackLayer3TopLocationTargetParamsDateRange = "12w"
	AttackLayer3TopLocationTargetParamsDateRange24w        AttackLayer3TopLocationTargetParamsDateRange = "24w"
	AttackLayer3TopLocationTargetParamsDateRange52w        AttackLayer3TopLocationTargetParamsDateRange = "52w"
	AttackLayer3TopLocationTargetParamsDateRange1dControl  AttackLayer3TopLocationTargetParamsDateRange = "1dControl"
	AttackLayer3TopLocationTargetParamsDateRange2dControl  AttackLayer3TopLocationTargetParamsDateRange = "2dControl"
	AttackLayer3TopLocationTargetParamsDateRange7dControl  AttackLayer3TopLocationTargetParamsDateRange = "7dControl"
	AttackLayer3TopLocationTargetParamsDateRange14dControl AttackLayer3TopLocationTargetParamsDateRange = "14dControl"
	AttackLayer3TopLocationTargetParamsDateRange28dControl AttackLayer3TopLocationTargetParamsDateRange = "28dControl"
	AttackLayer3TopLocationTargetParamsDateRange12wControl AttackLayer3TopLocationTargetParamsDateRange = "12wControl"
	AttackLayer3TopLocationTargetParamsDateRange24wControl AttackLayer3TopLocationTargetParamsDateRange = "24wControl"
)

func (r AttackLayer3TopLocationTargetParamsDateRange) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationTargetParamsDateRange1d, AttackLayer3TopLocationTargetParamsDateRange2d, AttackLayer3TopLocationTargetParamsDateRange7d, AttackLayer3TopLocationTargetParamsDateRange14d, AttackLayer3TopLocationTargetParamsDateRange28d, AttackLayer3TopLocationTargetParamsDateRange12w, AttackLayer3TopLocationTargetParamsDateRange24w, AttackLayer3TopLocationTargetParamsDateRange52w, AttackLayer3TopLocationTargetParamsDateRange1dControl, AttackLayer3TopLocationTargetParamsDateRange2dControl, AttackLayer3TopLocationTargetParamsDateRange7dControl, AttackLayer3TopLocationTargetParamsDateRange14dControl, AttackLayer3TopLocationTargetParamsDateRange28dControl, AttackLayer3TopLocationTargetParamsDateRange12wControl, AttackLayer3TopLocationTargetParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer3TopLocationTargetParamsFormat string

const (
	AttackLayer3TopLocationTargetParamsFormatJson AttackLayer3TopLocationTargetParamsFormat = "JSON"
	AttackLayer3TopLocationTargetParamsFormatCsv  AttackLayer3TopLocationTargetParamsFormat = "CSV"
)

func (r AttackLayer3TopLocationTargetParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationTargetParamsFormatJson, AttackLayer3TopLocationTargetParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer3TopLocationTargetParamsIPVersion string

const (
	AttackLayer3TopLocationTargetParamsIPVersionIPv4 AttackLayer3TopLocationTargetParamsIPVersion = "IPv4"
	AttackLayer3TopLocationTargetParamsIPVersionIPv6 AttackLayer3TopLocationTargetParamsIPVersion = "IPv6"
)

func (r AttackLayer3TopLocationTargetParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationTargetParamsIPVersionIPv4, AttackLayer3TopLocationTargetParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer3TopLocationTargetParamsProtocol string

const (
	AttackLayer3TopLocationTargetParamsProtocolUdp  AttackLayer3TopLocationTargetParamsProtocol = "UDP"
	AttackLayer3TopLocationTargetParamsProtocolTCP  AttackLayer3TopLocationTargetParamsProtocol = "TCP"
	AttackLayer3TopLocationTargetParamsProtocolIcmp AttackLayer3TopLocationTargetParamsProtocol = "ICMP"
	AttackLayer3TopLocationTargetParamsProtocolGRE  AttackLayer3TopLocationTargetParamsProtocol = "GRE"
)

func (r AttackLayer3TopLocationTargetParamsProtocol) IsKnown() bool {
	switch r {
	case AttackLayer3TopLocationTargetParamsProtocolUdp, AttackLayer3TopLocationTargetParamsProtocolTCP, AttackLayer3TopLocationTargetParamsProtocolIcmp, AttackLayer3TopLocationTargetParamsProtocolGRE:
		return true
	}
	return false
}

type AttackLayer3TopLocationTargetResponseEnvelope struct {
	Result  AttackLayer3TopLocationTargetResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    attackLayer3TopLocationTargetResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TopLocationTargetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer3TopLocationTargetResponseEnvelope]
type attackLayer3TopLocationTargetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopLocationTargetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopLocationTargetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

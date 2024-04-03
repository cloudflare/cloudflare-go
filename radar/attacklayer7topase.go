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

// AttackLayer7TopAseService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAttackLayer7TopAseService] method
// instead.
type AttackLayer7TopAseService struct {
	Options []option.RequestOption
}

// NewAttackLayer7TopAseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAttackLayer7TopAseService(opts ...option.RequestOption) (r *AttackLayer7TopAseService) {
	r = &AttackLayer7TopAseService{}
	r.Options = opts
	return
}

// Get the top origin Autonomous Systems of and by layer 7 attacks. Values are a
// percentage out of the total layer 7 attacks. The origin Autonomous Systems is
// determined by the client IP.
func (r *AttackLayer7TopAseService) Origin(ctx context.Context, query AttackLayer7TopAseOriginParams, opts ...option.RequestOption) (res *AttackLayer7TopAseOriginResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer7TopAseOriginResponseEnvelope
	path := "radar/attacks/layer7/top/ases/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer7TopAseOriginResponse struct {
	Meta AttackLayer7TopAseOriginResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer7TopAseOriginResponseTop0 `json:"top_0,required"`
	JSON attackLayer7TopAseOriginResponseJSON   `json:"-"`
}

// attackLayer7TopAseOriginResponseJSON contains the JSON metadata for the struct
// [AttackLayer7TopAseOriginResponse]
type attackLayer7TopAseOriginResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAseOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAseOriginResponseMeta struct {
	DateRange      []AttackLayer7TopAseOriginResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                             `json:"lastUpdated,required"`
	ConfidenceInfo AttackLayer7TopAseOriginResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer7TopAseOriginResponseMetaJSON           `json:"-"`
}

// attackLayer7TopAseOriginResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7TopAseOriginResponseMeta]
type attackLayer7TopAseOriginResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TopAseOriginResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAseOriginResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TopAseOriginResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TopAseOriginResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer7TopAseOriginResponseMetaDateRange]
type attackLayer7TopAseOriginResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAseOriginResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAseOriginResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                          `json:"level"`
	JSON        attackLayer7TopAseOriginResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer7TopAseOriginResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AttackLayer7TopAseOriginResponseMetaConfidenceInfo]
type attackLayer7TopAseOriginResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAseOriginResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                           `json:"dataSource,required"`
	Description     string                                                           `json:"description,required"`
	EventType       string                                                           `json:"eventType,required"`
	IsInstantaneous interface{}                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                        `json:"startTime" format:"date-time"`
	JSON            attackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotation]
type attackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AttackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAseOriginResponseTop0 struct {
	OriginASN     string                                   `json:"originAsn,required"`
	OriginASNName string                                   `json:"originAsnName,required"`
	Rank          float64                                  `json:"rank,required"`
	Value         string                                   `json:"value,required"`
	JSON          attackLayer7TopAseOriginResponseTop0JSON `json:"-"`
}

// attackLayer7TopAseOriginResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer7TopAseOriginResponseTop0]
type attackLayer7TopAseOriginResponseTop0JSON struct {
	OriginASN     apijson.Field
	OriginASNName apijson.Field
	Rank          apijson.Field
	Value         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AttackLayer7TopAseOriginResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopAseOriginParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer7TopAseOriginParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AttackLayer7TopAseOriginParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7TopAseOriginParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TopAseOriginParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttackLayer7TopAseOriginParamsDateRange string

const (
	AttackLayer7TopAseOriginParamsDateRange1d         AttackLayer7TopAseOriginParamsDateRange = "1d"
	AttackLayer7TopAseOriginParamsDateRange2d         AttackLayer7TopAseOriginParamsDateRange = "2d"
	AttackLayer7TopAseOriginParamsDateRange7d         AttackLayer7TopAseOriginParamsDateRange = "7d"
	AttackLayer7TopAseOriginParamsDateRange14d        AttackLayer7TopAseOriginParamsDateRange = "14d"
	AttackLayer7TopAseOriginParamsDateRange28d        AttackLayer7TopAseOriginParamsDateRange = "28d"
	AttackLayer7TopAseOriginParamsDateRange12w        AttackLayer7TopAseOriginParamsDateRange = "12w"
	AttackLayer7TopAseOriginParamsDateRange24w        AttackLayer7TopAseOriginParamsDateRange = "24w"
	AttackLayer7TopAseOriginParamsDateRange52w        AttackLayer7TopAseOriginParamsDateRange = "52w"
	AttackLayer7TopAseOriginParamsDateRange1dControl  AttackLayer7TopAseOriginParamsDateRange = "1dControl"
	AttackLayer7TopAseOriginParamsDateRange2dControl  AttackLayer7TopAseOriginParamsDateRange = "2dControl"
	AttackLayer7TopAseOriginParamsDateRange7dControl  AttackLayer7TopAseOriginParamsDateRange = "7dControl"
	AttackLayer7TopAseOriginParamsDateRange14dControl AttackLayer7TopAseOriginParamsDateRange = "14dControl"
	AttackLayer7TopAseOriginParamsDateRange28dControl AttackLayer7TopAseOriginParamsDateRange = "28dControl"
	AttackLayer7TopAseOriginParamsDateRange12wControl AttackLayer7TopAseOriginParamsDateRange = "12wControl"
	AttackLayer7TopAseOriginParamsDateRange24wControl AttackLayer7TopAseOriginParamsDateRange = "24wControl"
)

func (r AttackLayer7TopAseOriginParamsDateRange) IsKnown() bool {
	switch r {
	case AttackLayer7TopAseOriginParamsDateRange1d, AttackLayer7TopAseOriginParamsDateRange2d, AttackLayer7TopAseOriginParamsDateRange7d, AttackLayer7TopAseOriginParamsDateRange14d, AttackLayer7TopAseOriginParamsDateRange28d, AttackLayer7TopAseOriginParamsDateRange12w, AttackLayer7TopAseOriginParamsDateRange24w, AttackLayer7TopAseOriginParamsDateRange52w, AttackLayer7TopAseOriginParamsDateRange1dControl, AttackLayer7TopAseOriginParamsDateRange2dControl, AttackLayer7TopAseOriginParamsDateRange7dControl, AttackLayer7TopAseOriginParamsDateRange14dControl, AttackLayer7TopAseOriginParamsDateRange28dControl, AttackLayer7TopAseOriginParamsDateRange12wControl, AttackLayer7TopAseOriginParamsDateRange24wControl:
		return true
	}
	return false
}

// Format results are returned in.
type AttackLayer7TopAseOriginParamsFormat string

const (
	AttackLayer7TopAseOriginParamsFormatJson AttackLayer7TopAseOriginParamsFormat = "JSON"
	AttackLayer7TopAseOriginParamsFormatCsv  AttackLayer7TopAseOriginParamsFormat = "CSV"
)

func (r AttackLayer7TopAseOriginParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TopAseOriginParamsFormatJson, AttackLayer7TopAseOriginParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TopAseOriginResponseEnvelope struct {
	Result  AttackLayer7TopAseOriginResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    attackLayer7TopAseOriginResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TopAseOriginResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer7TopAseOriginResponseEnvelope]
type attackLayer7TopAseOriginResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopAseOriginResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopAseOriginResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

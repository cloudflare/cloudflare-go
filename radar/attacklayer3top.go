// File generated from our OpenAPI spec by Stainless.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// AttackLayer3TopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAttackLayer3TopService] method
// instead.
type AttackLayer3TopService struct {
	Options   []option.RequestOption
	Locations *AttackLayer3TopLocationService
}

// NewAttackLayer3TopService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAttackLayer3TopService(opts ...option.RequestOption) (r *AttackLayer3TopService) {
	r = &AttackLayer3TopService{}
	r.Options = opts
	r.Locations = NewAttackLayer3TopLocationService(opts...)
	return
}

// Get the top attacks from origin to target location. Values are a percentage out
// of the total layer 3 attacks (with billing country). You can optionally limit
// the number of attacks per origin/target location (useful if all the top attacks
// are from or to the same location).
func (r *AttackLayer3TopService) Attacks(ctx context.Context, query AttackLayer3TopAttacksParams, opts ...option.RequestOption) (res *AttackLayer3TopAttacksResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer3TopAttacksResponseEnvelope
	path := "radar/attacks/layer3/top/attacks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the Industry of attacks.
func (r *AttackLayer3TopService) Industry(ctx context.Context, query AttackLayer3TopIndustryParams, opts ...option.RequestOption) (res *AttackLayer3TopIndustryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer3TopIndustryResponseEnvelope
	path := "radar/attacks/layer3/top/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the Verticals of attacks.
func (r *AttackLayer3TopService) Vertical(ctx context.Context, query AttackLayer3TopVerticalParams, opts ...option.RequestOption) (res *AttackLayer3TopVerticalResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackLayer3TopVerticalResponseEnvelope
	path := "radar/attacks/layer3/top/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer3TopAttacksResponse struct {
	Meta AttackLayer3TopAttacksResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer3TopAttacksResponseTop0 `json:"top_0,required"`
	JSON attackLayer3TopAttacksResponseJSON   `json:"-"`
}

// attackLayer3TopAttacksResponseJSON contains the JSON metadata for the struct
// [AttackLayer3TopAttacksResponse]
type attackLayer3TopAttacksResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopAttacksResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopAttacksResponseMeta struct {
	DateRange      []AttackLayer3TopAttacksResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	ConfidenceInfo AttackLayer3TopAttacksResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer3TopAttacksResponseMetaJSON           `json:"-"`
}

// attackLayer3TopAttacksResponseMetaJSON contains the JSON metadata for the struct
// [AttackLayer3TopAttacksResponseMeta]
type attackLayer3TopAttacksResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TopAttacksResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopAttacksResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TopAttacksResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TopAttacksResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer3TopAttacksResponseMetaDateRange]
type attackLayer3TopAttacksResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopAttacksResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopAttacksResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        attackLayer3TopAttacksResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer3TopAttacksResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer3TopAttacksResponseMetaConfidenceInfo]
type attackLayer3TopAttacksResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopAttacksResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            attackLayer3TopAttacksResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TopAttacksResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotation]
type attackLayer3TopAttacksResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3TopAttacksResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopAttacksResponseTop0 struct {
	OriginCountryAlpha2 string                                 `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                 `json:"originCountryName,required"`
	Value               string                                 `json:"value,required"`
	JSON                attackLayer3TopAttacksResponseTop0JSON `json:"-"`
}

// attackLayer3TopAttacksResponseTop0JSON contains the JSON metadata for the struct
// [AttackLayer3TopAttacksResponseTop0]
type attackLayer3TopAttacksResponseTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AttackLayer3TopAttacksResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopIndustryResponse struct {
	Meta AttackLayer3TopIndustryResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer3TopIndustryResponseTop0 `json:"top_0,required"`
	JSON attackLayer3TopIndustryResponseJSON   `json:"-"`
}

// attackLayer3TopIndustryResponseJSON contains the JSON metadata for the struct
// [AttackLayer3TopIndustryResponse]
type attackLayer3TopIndustryResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopIndustryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopIndustryResponseMeta struct {
	DateRange      []AttackLayer3TopIndustryResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                            `json:"lastUpdated,required"`
	ConfidenceInfo AttackLayer3TopIndustryResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer3TopIndustryResponseMetaJSON           `json:"-"`
}

// attackLayer3TopIndustryResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3TopIndustryResponseMeta]
type attackLayer3TopIndustryResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TopIndustryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopIndustryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TopIndustryResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TopIndustryResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer3TopIndustryResponseMetaDateRange]
type attackLayer3TopIndustryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopIndustryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopIndustryResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                         `json:"level"`
	JSON        attackLayer3TopIndustryResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer3TopIndustryResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer3TopIndustryResponseMetaConfidenceInfo]
type attackLayer3TopIndustryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopIndustryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                          `json:"dataSource,required"`
	Description     string                                                          `json:"description,required"`
	EventType       string                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                       `json:"startTime" format:"date-time"`
	JSON            attackLayer3TopIndustryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TopIndustryResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotation]
type attackLayer3TopIndustryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3TopIndustryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopIndustryResponseTop0 struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  attackLayer3TopIndustryResponseTop0JSON `json:"-"`
}

// attackLayer3TopIndustryResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer3TopIndustryResponseTop0]
type attackLayer3TopIndustryResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopIndustryResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopVerticalResponse struct {
	Meta AttackLayer3TopVerticalResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer3TopVerticalResponseTop0 `json:"top_0,required"`
	JSON attackLayer3TopVerticalResponseJSON   `json:"-"`
}

// attackLayer3TopVerticalResponseJSON contains the JSON metadata for the struct
// [AttackLayer3TopVerticalResponse]
type attackLayer3TopVerticalResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopVerticalResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopVerticalResponseMeta struct {
	DateRange      []AttackLayer3TopVerticalResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                            `json:"lastUpdated,required"`
	ConfidenceInfo AttackLayer3TopVerticalResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           attackLayer3TopVerticalResponseMetaJSON           `json:"-"`
}

// attackLayer3TopVerticalResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer3TopVerticalResponseMeta]
type attackLayer3TopVerticalResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer3TopVerticalResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopVerticalResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      attackLayer3TopVerticalResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer3TopVerticalResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AttackLayer3TopVerticalResponseMetaDateRange]
type attackLayer3TopVerticalResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopVerticalResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopVerticalResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                         `json:"level"`
	JSON        attackLayer3TopVerticalResponseMetaConfidenceInfoJSON         `json:"-"`
}

// attackLayer3TopVerticalResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AttackLayer3TopVerticalResponseMetaConfidenceInfo]
type attackLayer3TopVerticalResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopVerticalResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                          `json:"dataSource,required"`
	Description     string                                                          `json:"description,required"`
	EventType       string                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                       `json:"startTime" format:"date-time"`
	JSON            attackLayer3TopVerticalResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer3TopVerticalResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotation]
type attackLayer3TopVerticalResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer3TopVerticalResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopVerticalResponseTop0 struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  attackLayer3TopVerticalResponseTop0JSON `json:"-"`
}

// attackLayer3TopVerticalResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer3TopVerticalResponseTop0]
type attackLayer3TopVerticalResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopVerticalResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopAttacksParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer3TopAttacksParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AttackLayer3TopAttacksParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3TopAttacksParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of attack origin/target location attack limits. Together with
	// `limitPerLocation`, limits how many objects will be fetched per origin/target
	// location.
	LimitDirection param.Field[AttackLayer3TopAttacksParamsLimitDirection] `query:"limitDirection"`
	// Limit the number of attacks per origin/target (refer to `limitDirection`
	// parameter) location.
	LimitPerLocation param.Field[int64] `query:"limitPerLocation"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]AttackLayer3TopAttacksParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TopAttacksParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3TopAttacksParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttackLayer3TopAttacksParamsDateRange string

const (
	AttackLayer3TopAttacksParamsDateRange1d         AttackLayer3TopAttacksParamsDateRange = "1d"
	AttackLayer3TopAttacksParamsDateRange2d         AttackLayer3TopAttacksParamsDateRange = "2d"
	AttackLayer3TopAttacksParamsDateRange7d         AttackLayer3TopAttacksParamsDateRange = "7d"
	AttackLayer3TopAttacksParamsDateRange14d        AttackLayer3TopAttacksParamsDateRange = "14d"
	AttackLayer3TopAttacksParamsDateRange28d        AttackLayer3TopAttacksParamsDateRange = "28d"
	AttackLayer3TopAttacksParamsDateRange12w        AttackLayer3TopAttacksParamsDateRange = "12w"
	AttackLayer3TopAttacksParamsDateRange24w        AttackLayer3TopAttacksParamsDateRange = "24w"
	AttackLayer3TopAttacksParamsDateRange52w        AttackLayer3TopAttacksParamsDateRange = "52w"
	AttackLayer3TopAttacksParamsDateRange1dControl  AttackLayer3TopAttacksParamsDateRange = "1dControl"
	AttackLayer3TopAttacksParamsDateRange2dControl  AttackLayer3TopAttacksParamsDateRange = "2dControl"
	AttackLayer3TopAttacksParamsDateRange7dControl  AttackLayer3TopAttacksParamsDateRange = "7dControl"
	AttackLayer3TopAttacksParamsDateRange14dControl AttackLayer3TopAttacksParamsDateRange = "14dControl"
	AttackLayer3TopAttacksParamsDateRange28dControl AttackLayer3TopAttacksParamsDateRange = "28dControl"
	AttackLayer3TopAttacksParamsDateRange12wControl AttackLayer3TopAttacksParamsDateRange = "12wControl"
	AttackLayer3TopAttacksParamsDateRange24wControl AttackLayer3TopAttacksParamsDateRange = "24wControl"
)

// Format results are returned in.
type AttackLayer3TopAttacksParamsFormat string

const (
	AttackLayer3TopAttacksParamsFormatJson AttackLayer3TopAttacksParamsFormat = "JSON"
	AttackLayer3TopAttacksParamsFormatCsv  AttackLayer3TopAttacksParamsFormat = "CSV"
)

type AttackLayer3TopAttacksParamsIPVersion string

const (
	AttackLayer3TopAttacksParamsIPVersionIPv4 AttackLayer3TopAttacksParamsIPVersion = "IPv4"
	AttackLayer3TopAttacksParamsIPVersionIPv6 AttackLayer3TopAttacksParamsIPVersion = "IPv6"
)

// Array of attack origin/target location attack limits. Together with
// `limitPerLocation`, limits how many objects will be fetched per origin/target
// location.
type AttackLayer3TopAttacksParamsLimitDirection string

const (
	AttackLayer3TopAttacksParamsLimitDirectionOrigin AttackLayer3TopAttacksParamsLimitDirection = "ORIGIN"
	AttackLayer3TopAttacksParamsLimitDirectionTarget AttackLayer3TopAttacksParamsLimitDirection = "TARGET"
)

type AttackLayer3TopAttacksParamsProtocol string

const (
	AttackLayer3TopAttacksParamsProtocolUdp  AttackLayer3TopAttacksParamsProtocol = "UDP"
	AttackLayer3TopAttacksParamsProtocolTcp  AttackLayer3TopAttacksParamsProtocol = "TCP"
	AttackLayer3TopAttacksParamsProtocolIcmp AttackLayer3TopAttacksParamsProtocol = "ICMP"
	AttackLayer3TopAttacksParamsProtocolGRE  AttackLayer3TopAttacksParamsProtocol = "GRE"
)

type AttackLayer3TopAttacksResponseEnvelope struct {
	Result  AttackLayer3TopAttacksResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    attackLayer3TopAttacksResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TopAttacksResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer3TopAttacksResponseEnvelope]
type attackLayer3TopAttacksResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopAttacksResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopAttacksResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopIndustryParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer3TopIndustryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AttackLayer3TopIndustryParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3TopIndustryParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]AttackLayer3TopIndustryParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TopIndustryParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3TopIndustryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttackLayer3TopIndustryParamsDateRange string

const (
	AttackLayer3TopIndustryParamsDateRange1d         AttackLayer3TopIndustryParamsDateRange = "1d"
	AttackLayer3TopIndustryParamsDateRange2d         AttackLayer3TopIndustryParamsDateRange = "2d"
	AttackLayer3TopIndustryParamsDateRange7d         AttackLayer3TopIndustryParamsDateRange = "7d"
	AttackLayer3TopIndustryParamsDateRange14d        AttackLayer3TopIndustryParamsDateRange = "14d"
	AttackLayer3TopIndustryParamsDateRange28d        AttackLayer3TopIndustryParamsDateRange = "28d"
	AttackLayer3TopIndustryParamsDateRange12w        AttackLayer3TopIndustryParamsDateRange = "12w"
	AttackLayer3TopIndustryParamsDateRange24w        AttackLayer3TopIndustryParamsDateRange = "24w"
	AttackLayer3TopIndustryParamsDateRange52w        AttackLayer3TopIndustryParamsDateRange = "52w"
	AttackLayer3TopIndustryParamsDateRange1dControl  AttackLayer3TopIndustryParamsDateRange = "1dControl"
	AttackLayer3TopIndustryParamsDateRange2dControl  AttackLayer3TopIndustryParamsDateRange = "2dControl"
	AttackLayer3TopIndustryParamsDateRange7dControl  AttackLayer3TopIndustryParamsDateRange = "7dControl"
	AttackLayer3TopIndustryParamsDateRange14dControl AttackLayer3TopIndustryParamsDateRange = "14dControl"
	AttackLayer3TopIndustryParamsDateRange28dControl AttackLayer3TopIndustryParamsDateRange = "28dControl"
	AttackLayer3TopIndustryParamsDateRange12wControl AttackLayer3TopIndustryParamsDateRange = "12wControl"
	AttackLayer3TopIndustryParamsDateRange24wControl AttackLayer3TopIndustryParamsDateRange = "24wControl"
)

// Format results are returned in.
type AttackLayer3TopIndustryParamsFormat string

const (
	AttackLayer3TopIndustryParamsFormatJson AttackLayer3TopIndustryParamsFormat = "JSON"
	AttackLayer3TopIndustryParamsFormatCsv  AttackLayer3TopIndustryParamsFormat = "CSV"
)

type AttackLayer3TopIndustryParamsIPVersion string

const (
	AttackLayer3TopIndustryParamsIPVersionIPv4 AttackLayer3TopIndustryParamsIPVersion = "IPv4"
	AttackLayer3TopIndustryParamsIPVersionIPv6 AttackLayer3TopIndustryParamsIPVersion = "IPv6"
)

type AttackLayer3TopIndustryParamsProtocol string

const (
	AttackLayer3TopIndustryParamsProtocolUdp  AttackLayer3TopIndustryParamsProtocol = "UDP"
	AttackLayer3TopIndustryParamsProtocolTcp  AttackLayer3TopIndustryParamsProtocol = "TCP"
	AttackLayer3TopIndustryParamsProtocolIcmp AttackLayer3TopIndustryParamsProtocol = "ICMP"
	AttackLayer3TopIndustryParamsProtocolGRE  AttackLayer3TopIndustryParamsProtocol = "GRE"
)

type AttackLayer3TopIndustryResponseEnvelope struct {
	Result  AttackLayer3TopIndustryResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    attackLayer3TopIndustryResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TopIndustryResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer3TopIndustryResponseEnvelope]
type attackLayer3TopIndustryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopIndustryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopIndustryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer3TopVerticalParams struct {
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]AttackLayer3TopVerticalParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[AttackLayer3TopVerticalParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]AttackLayer3TopVerticalParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]AttackLayer3TopVerticalParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [AttackLayer3TopVerticalParams]'s query parameters as
// `url.Values`.
func (r AttackLayer3TopVerticalParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AttackLayer3TopVerticalParamsDateRange string

const (
	AttackLayer3TopVerticalParamsDateRange1d         AttackLayer3TopVerticalParamsDateRange = "1d"
	AttackLayer3TopVerticalParamsDateRange2d         AttackLayer3TopVerticalParamsDateRange = "2d"
	AttackLayer3TopVerticalParamsDateRange7d         AttackLayer3TopVerticalParamsDateRange = "7d"
	AttackLayer3TopVerticalParamsDateRange14d        AttackLayer3TopVerticalParamsDateRange = "14d"
	AttackLayer3TopVerticalParamsDateRange28d        AttackLayer3TopVerticalParamsDateRange = "28d"
	AttackLayer3TopVerticalParamsDateRange12w        AttackLayer3TopVerticalParamsDateRange = "12w"
	AttackLayer3TopVerticalParamsDateRange24w        AttackLayer3TopVerticalParamsDateRange = "24w"
	AttackLayer3TopVerticalParamsDateRange52w        AttackLayer3TopVerticalParamsDateRange = "52w"
	AttackLayer3TopVerticalParamsDateRange1dControl  AttackLayer3TopVerticalParamsDateRange = "1dControl"
	AttackLayer3TopVerticalParamsDateRange2dControl  AttackLayer3TopVerticalParamsDateRange = "2dControl"
	AttackLayer3TopVerticalParamsDateRange7dControl  AttackLayer3TopVerticalParamsDateRange = "7dControl"
	AttackLayer3TopVerticalParamsDateRange14dControl AttackLayer3TopVerticalParamsDateRange = "14dControl"
	AttackLayer3TopVerticalParamsDateRange28dControl AttackLayer3TopVerticalParamsDateRange = "28dControl"
	AttackLayer3TopVerticalParamsDateRange12wControl AttackLayer3TopVerticalParamsDateRange = "12wControl"
	AttackLayer3TopVerticalParamsDateRange24wControl AttackLayer3TopVerticalParamsDateRange = "24wControl"
)

// Format results are returned in.
type AttackLayer3TopVerticalParamsFormat string

const (
	AttackLayer3TopVerticalParamsFormatJson AttackLayer3TopVerticalParamsFormat = "JSON"
	AttackLayer3TopVerticalParamsFormatCsv  AttackLayer3TopVerticalParamsFormat = "CSV"
)

type AttackLayer3TopVerticalParamsIPVersion string

const (
	AttackLayer3TopVerticalParamsIPVersionIPv4 AttackLayer3TopVerticalParamsIPVersion = "IPv4"
	AttackLayer3TopVerticalParamsIPVersionIPv6 AttackLayer3TopVerticalParamsIPVersion = "IPv6"
)

type AttackLayer3TopVerticalParamsProtocol string

const (
	AttackLayer3TopVerticalParamsProtocolUdp  AttackLayer3TopVerticalParamsProtocol = "UDP"
	AttackLayer3TopVerticalParamsProtocolTcp  AttackLayer3TopVerticalParamsProtocol = "TCP"
	AttackLayer3TopVerticalParamsProtocolIcmp AttackLayer3TopVerticalParamsProtocol = "ICMP"
	AttackLayer3TopVerticalParamsProtocolGRE  AttackLayer3TopVerticalParamsProtocol = "GRE"
)

type AttackLayer3TopVerticalResponseEnvelope struct {
	Result  AttackLayer3TopVerticalResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    attackLayer3TopVerticalResponseEnvelopeJSON `json:"-"`
}

// attackLayer3TopVerticalResponseEnvelopeJSON contains the JSON metadata for the
// struct [AttackLayer3TopVerticalResponseEnvelope]
type attackLayer3TopVerticalResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer3TopVerticalResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer3TopVerticalResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

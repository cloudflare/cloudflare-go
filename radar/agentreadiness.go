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

// AgentReadinessService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentReadinessService] method instead.
type AgentReadinessService struct {
	Options []option.RequestOption
}

// NewAgentReadinessService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentReadinessService(opts ...option.RequestOption) (r *AgentReadinessService) {
	r = &AgentReadinessService{}
	r.Options = opts
	return
}

// Returns a summary of AI agent readiness scores across scanned domains, grouped
// by the specified dimension. Data is sourced from weekly bulk scans. All values
// are raw domain counts.
func (r *AgentReadinessService) Summary(ctx context.Context, dimension AgentReadinessSummaryParamsDimension, query AgentReadinessSummaryParams, opts ...option.RequestOption) (res *AgentReadinessSummaryResponse, err error) {
	var env AgentReadinessSummaryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("radar/agent_readiness/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type AgentReadinessSummaryResponse struct {
	Meta     AgentReadinessSummaryResponseMeta `json:"meta" api:"required"`
	Summary0 map[string]string                 `json:"summary_0" api:"required"`
	JSON     agentReadinessSummaryResponseJSON `json:"-"`
}

// agentReadinessSummaryResponseJSON contains the JSON metadata for the struct
// [AgentReadinessSummaryResponse]
type agentReadinessSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentReadinessSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentReadinessSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type AgentReadinessSummaryResponseMeta struct {
	// Date of the returned scan (YYYY-MM-DD). May differ from the requested date if no
	// scan exists for that exact date.
	Date string `json:"date" api:"required"`
	// Available domain sub-categories with their scan counts. Use as filter options
	// for the domainCategory parameter.
	DomainCategories []AgentReadinessSummaryResponseMetaDomainCategory `json:"domainCategories" api:"required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated" api:"required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AgentReadinessSummaryResponseMetaNormalization `json:"normalization" api:"required"`
	// Domains successfully scanned (excludes errors).
	SuccessfulDomains int64 `json:"successfulDomains" api:"required"`
	// Total domains attempted in the scan.
	TotalDomains int64 `json:"totalDomains" api:"required"`
	// Measurement units for the results.
	Units []AgentReadinessSummaryResponseMetaUnit `json:"units" api:"required"`
	JSON  agentReadinessSummaryResponseMetaJSON   `json:"-"`
}

// agentReadinessSummaryResponseMetaJSON contains the JSON metadata for the struct
// [AgentReadinessSummaryResponseMeta]
type agentReadinessSummaryResponseMetaJSON struct {
	Date              apijson.Field
	DomainCategories  apijson.Field
	LastUpdated       apijson.Field
	Normalization     apijson.Field
	SuccessfulDomains apijson.Field
	TotalDomains      apijson.Field
	Units             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AgentReadinessSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentReadinessSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AgentReadinessSummaryResponseMetaDomainCategory struct {
	// Sub-category name.
	Name string `json:"name" api:"required"`
	// Number of successfully scanned domains in this sub-category.
	Value int64                                               `json:"value" api:"required"`
	JSON  agentReadinessSummaryResponseMetaDomainCategoryJSON `json:"-"`
}

// agentReadinessSummaryResponseMetaDomainCategoryJSON contains the JSON metadata
// for the struct [AgentReadinessSummaryResponseMetaDomainCategory]
type agentReadinessSummaryResponseMetaDomainCategoryJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentReadinessSummaryResponseMetaDomainCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentReadinessSummaryResponseMetaDomainCategoryJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AgentReadinessSummaryResponseMetaNormalization string

const (
	AgentReadinessSummaryResponseMetaNormalizationPercentage           AgentReadinessSummaryResponseMetaNormalization = "PERCENTAGE"
	AgentReadinessSummaryResponseMetaNormalizationMin0Max              AgentReadinessSummaryResponseMetaNormalization = "MIN0_MAX"
	AgentReadinessSummaryResponseMetaNormalizationMinMax               AgentReadinessSummaryResponseMetaNormalization = "MIN_MAX"
	AgentReadinessSummaryResponseMetaNormalizationRawValues            AgentReadinessSummaryResponseMetaNormalization = "RAW_VALUES"
	AgentReadinessSummaryResponseMetaNormalizationPercentageChange     AgentReadinessSummaryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AgentReadinessSummaryResponseMetaNormalizationRollingAverage       AgentReadinessSummaryResponseMetaNormalization = "ROLLING_AVERAGE"
	AgentReadinessSummaryResponseMetaNormalizationOverlappedPercentage AgentReadinessSummaryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AgentReadinessSummaryResponseMetaNormalizationRatio                AgentReadinessSummaryResponseMetaNormalization = "RATIO"
)

func (r AgentReadinessSummaryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AgentReadinessSummaryResponseMetaNormalizationPercentage, AgentReadinessSummaryResponseMetaNormalizationMin0Max, AgentReadinessSummaryResponseMetaNormalizationMinMax, AgentReadinessSummaryResponseMetaNormalizationRawValues, AgentReadinessSummaryResponseMetaNormalizationPercentageChange, AgentReadinessSummaryResponseMetaNormalizationRollingAverage, AgentReadinessSummaryResponseMetaNormalizationOverlappedPercentage, AgentReadinessSummaryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AgentReadinessSummaryResponseMetaUnit struct {
	Name  string                                    `json:"name" api:"required"`
	Value string                                    `json:"value" api:"required"`
	JSON  agentReadinessSummaryResponseMetaUnitJSON `json:"-"`
}

// agentReadinessSummaryResponseMetaUnitJSON contains the JSON metadata for the
// struct [AgentReadinessSummaryResponseMetaUnit]
type agentReadinessSummaryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentReadinessSummaryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentReadinessSummaryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AgentReadinessSummaryParams struct {
	// Filters results by the specified date.
	Date param.Field[time.Time] `query:"date" format:"date"`
	// Filters results by domain category.
	DomainCategory param.Field[[]string] `query:"domainCategory"`
	// Format in which results will be returned.
	Format param.Field[AgentReadinessSummaryParamsFormat] `query:"format"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AgentReadinessSummaryParams]'s query parameters as
// `url.Values`.
func (r AgentReadinessSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the agent readiness data dimension by which to group the results.
type AgentReadinessSummaryParamsDimension string

const (
	AgentReadinessSummaryParamsDimensionCheck AgentReadinessSummaryParamsDimension = "CHECK"
)

func (r AgentReadinessSummaryParamsDimension) IsKnown() bool {
	switch r {
	case AgentReadinessSummaryParamsDimensionCheck:
		return true
	}
	return false
}

// Format in which results will be returned.
type AgentReadinessSummaryParamsFormat string

const (
	AgentReadinessSummaryParamsFormatJson AgentReadinessSummaryParamsFormat = "JSON"
	AgentReadinessSummaryParamsFormatCsv  AgentReadinessSummaryParamsFormat = "CSV"
)

func (r AgentReadinessSummaryParamsFormat) IsKnown() bool {
	switch r {
	case AgentReadinessSummaryParamsFormatJson, AgentReadinessSummaryParamsFormatCsv:
		return true
	}
	return false
}

type AgentReadinessSummaryResponseEnvelope struct {
	Result  AgentReadinessSummaryResponse             `json:"result" api:"required"`
	Success bool                                      `json:"success" api:"required"`
	JSON    agentReadinessSummaryResponseEnvelopeJSON `json:"-"`
}

// agentReadinessSummaryResponseEnvelopeJSON contains the JSON metadata for the
// struct [AgentReadinessSummaryResponseEnvelope]
type agentReadinessSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AgentReadinessSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r agentReadinessSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

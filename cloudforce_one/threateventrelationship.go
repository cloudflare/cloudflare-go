// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ThreatEventRelationshipService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventRelationshipService] method instead.
type ThreatEventRelationshipService struct {
	Options []option.RequestOption
}

// NewThreatEventRelationshipService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventRelationshipService(opts ...option.RequestOption) (r *ThreatEventRelationshipService) {
	r = &ThreatEventRelationshipService{}
	r.Options = opts
	return
}

// The `event_id` must be defined (to list existing events (and their IDs), use the
// [`Filter and List Events`](https://developers.cloudflare.com/api/resources/cloudforce_one/subresources/threat_events/methods/list/)
// endpoint). Also, must provide query parameters.
func (r *ThreatEventRelationshipService) List(ctx context.Context, eventID string, params ThreatEventRelationshipListParams, opts ...option.RequestOption) (res *[]ThreatEventRelationshipListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s/relationships", params.AccountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type ThreatEventRelationshipListResponse struct {
	Attacker              string                                  `json:"attacker" api:"required"`
	AttackerCountry       string                                  `json:"attackerCountry" api:"required"`
	AttackerCountryAlpha3 string                                  `json:"attackerCountryAlpha3" api:"required"`
	Category              string                                  `json:"category" api:"required"`
	DatasetID             string                                  `json:"datasetId" api:"required"`
	Date                  string                                  `json:"date" api:"required"`
	Event                 string                                  `json:"event" api:"required"`
	HasChildren           bool                                    `json:"hasChildren" api:"required"`
	Indicator             string                                  `json:"indicator" api:"required"`
	IndicatorType         string                                  `json:"indicatorType" api:"required"`
	IndicatorTypeID       float64                                 `json:"indicatorTypeId" api:"required"`
	KillChain             float64                                 `json:"killChain" api:"required"`
	MitreAttack           []string                                `json:"mitreAttack" api:"required"`
	MitreCapec            []string                                `json:"mitreCapec" api:"required"`
	NumReferenced         float64                                 `json:"numReferenced" api:"required"`
	NumReferences         float64                                 `json:"numReferences" api:"required"`
	RawID                 string                                  `json:"rawId" api:"required"`
	Referenced            []string                                `json:"referenced" api:"required"`
	ReferencedIDs         []float64                               `json:"referencedIds" api:"required"`
	References            []string                                `json:"references" api:"required"`
	ReferencesIDs         []float64                               `json:"referencesIds" api:"required"`
	Tags                  []string                                `json:"tags" api:"required"`
	TargetCountry         string                                  `json:"targetCountry" api:"required"`
	TargetCountryAlpha3   string                                  `json:"targetCountryAlpha3" api:"required"`
	TargetIndustry        string                                  `json:"targetIndustry" api:"required"`
	TLP                   string                                  `json:"tlp" api:"required"`
	UUID                  string                                  `json:"uuid" api:"required"`
	Insight               string                                  `json:"insight"`
	ReleasabilityID       string                                  `json:"releasabilityId"`
	JSON                  threatEventRelationshipListResponseJSON `json:"-"`
}

// threatEventRelationshipListResponseJSON contains the JSON metadata for the
// struct [ThreatEventRelationshipListResponse]
type threatEventRelationshipListResponseJSON struct {
	Attacker              apijson.Field
	AttackerCountry       apijson.Field
	AttackerCountryAlpha3 apijson.Field
	Category              apijson.Field
	DatasetID             apijson.Field
	Date                  apijson.Field
	Event                 apijson.Field
	HasChildren           apijson.Field
	Indicator             apijson.Field
	IndicatorType         apijson.Field
	IndicatorTypeID       apijson.Field
	KillChain             apijson.Field
	MitreAttack           apijson.Field
	MitreCapec            apijson.Field
	NumReferenced         apijson.Field
	NumReferences         apijson.Field
	RawID                 apijson.Field
	Referenced            apijson.Field
	ReferencedIDs         apijson.Field
	References            apijson.Field
	ReferencesIDs         apijson.Field
	Tags                  apijson.Field
	TargetCountry         apijson.Field
	TargetCountryAlpha3   apijson.Field
	TargetIndustry        apijson.Field
	TLP                   apijson.Field
	UUID                  apijson.Field
	Insight               apijson.Field
	ReleasabilityID       apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ThreatEventRelationshipListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventRelationshipListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventRelationshipListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The dataset ID to search within.
	DatasetID param.Field[string] `query:"datasetId" api:"required"`
	// The direction to traverse the graph. Defaults to 'both' to search all.
	Direction param.Field[ThreatEventRelationshipListParamsDirection] `query:"direction"`
	// Whether to include the starting event in the results. Defaults to true.
	IncludeParent param.Field[bool] `query:"includeParent"`
	// An optional array of indicator type IDs to filter the results by.
	IndicatorTypeIDs param.Field[[]string] `query:"indicatorTypeIds"`
	// The maximum depth to traverse. Defaults to 5.
	MaxDepth param.Field[float64] `query:"maxDepth"`
	Page     param.Field[float64] `query:"page"`
	PageSize param.Field[float64] `query:"pageSize"`
	// An optional array of relationship types to filter by.
	RelationshipTypes param.Field[ThreatEventRelationshipListParamsRelationshipTypesUnion] `query:"relationshipTypes"`
}

// URLQuery serializes [ThreatEventRelationshipListParams]'s query parameters as
// `url.Values`.
func (r ThreatEventRelationshipListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The direction to traverse the graph. Defaults to 'both' to search all.
type ThreatEventRelationshipListParamsDirection string

const (
	ThreatEventRelationshipListParamsDirectionAncestors   ThreatEventRelationshipListParamsDirection = "ancestors"
	ThreatEventRelationshipListParamsDirectionDescendants ThreatEventRelationshipListParamsDirection = "descendants"
	ThreatEventRelationshipListParamsDirectionBoth        ThreatEventRelationshipListParamsDirection = "both"
)

func (r ThreatEventRelationshipListParamsDirection) IsKnown() bool {
	switch r {
	case ThreatEventRelationshipListParamsDirectionAncestors, ThreatEventRelationshipListParamsDirectionDescendants, ThreatEventRelationshipListParamsDirectionBoth:
		return true
	}
	return false
}

// An optional array of relationship types to filter by.
//
// Satisfied by [shared.UnionString],
// [cloudforce_one.ThreatEventRelationshipListParamsRelationshipTypesArray].
type ThreatEventRelationshipListParamsRelationshipTypesUnion interface {
	ImplementsThreatEventRelationshipListParamsRelationshipTypesUnion()
}

type ThreatEventRelationshipListParamsRelationshipTypesArray []string

func (r ThreatEventRelationshipListParamsRelationshipTypesArray) ImplementsThreatEventRelationshipListParamsRelationshipTypesUnion() {
}

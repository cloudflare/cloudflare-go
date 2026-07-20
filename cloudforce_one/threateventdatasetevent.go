// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ThreatEventDatasetEventService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventDatasetEventService] method instead.
type ThreatEventDatasetEventService struct {
	Options []option.RequestOption
}

// NewThreatEventDatasetEventService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventDatasetEventService(opts ...option.RequestOption) (r *ThreatEventDatasetEventService) {
	r = &ThreatEventDatasetEventService{}
	r.Options = opts
	return
}

// Retrieves a specific event by its UUID.
func (r *ThreatEventDatasetEventService) Get(ctx context.Context, datasetID string, eventID string, query ThreatEventDatasetEventGetParams, opts ...option.RequestOption) (res *ThreatEventDatasetEventGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return nil, err
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/dataset/%s/events/%s", query.AccountID, datasetID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ThreatEventDatasetEventGetResponse struct {
	Attacker              string                                 `json:"attacker" api:"required"`
	AttackerCountry       string                                 `json:"attackerCountry" api:"required"`
	AttackerCountryAlpha3 string                                 `json:"attackerCountryAlpha3" api:"required"`
	Category              string                                 `json:"category" api:"required"`
	DatasetID             string                                 `json:"datasetId" api:"required"`
	Date                  string                                 `json:"date" api:"required"`
	Event                 string                                 `json:"event" api:"required"`
	HasChildren           bool                                   `json:"hasChildren" api:"required"`
	Indicator             string                                 `json:"indicator" api:"required"`
	IndicatorType         string                                 `json:"indicatorType" api:"required"`
	IndicatorTypeID       float64                                `json:"indicatorTypeId" api:"required"`
	KillChain             float64                                `json:"killChain" api:"required"`
	MitreAttack           []string                               `json:"mitreAttack" api:"required"`
	MitreCapec            []string                               `json:"mitreCapec" api:"required"`
	NumReferenced         float64                                `json:"numReferenced" api:"required"`
	NumReferences         float64                                `json:"numReferences" api:"required"`
	RawID                 string                                 `json:"rawId" api:"required"`
	Referenced            []string                               `json:"referenced" api:"required"`
	ReferencedIDs         []float64                              `json:"referencedIds" api:"required"`
	References            []string                               `json:"references" api:"required"`
	ReferencesIDs         []float64                              `json:"referencesIds" api:"required"`
	Tags                  []string                               `json:"tags" api:"required"`
	TargetCountry         string                                 `json:"targetCountry" api:"required"`
	TargetCountryAlpha3   string                                 `json:"targetCountryAlpha3" api:"required"`
	TargetIndustry        string                                 `json:"targetIndustry" api:"required"`
	TLP                   string                                 `json:"tlp" api:"required"`
	UUID                  string                                 `json:"uuid" api:"required"`
	Insight               string                                 `json:"insight"`
	ReleasabilityID       string                                 `json:"releasabilityId"`
	JSON                  threatEventDatasetEventGetResponseJSON `json:"-"`
}

// threatEventDatasetEventGetResponseJSON contains the JSON metadata for the struct
// [ThreatEventDatasetEventGetResponse]
type threatEventDatasetEventGetResponseJSON struct {
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

func (r *ThreatEventDatasetEventGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetEventGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetEventGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

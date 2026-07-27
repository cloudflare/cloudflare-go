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

// ThreatEventTargetIndustryByDatasetService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventTargetIndustryByDatasetService] method instead.
type ThreatEventTargetIndustryByDatasetService struct {
	Options []option.RequestOption
}

// NewThreatEventTargetIndustryByDatasetService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewThreatEventTargetIndustryByDatasetService(opts ...option.RequestOption) (r *ThreatEventTargetIndustryByDatasetService) {
	r = &ThreatEventTargetIndustryByDatasetService{}
	r.Options = opts
	return
}

// List all target industries referenced in events for a specific dataset.
func (r *ThreatEventTargetIndustryByDatasetService) List(ctx context.Context, datasetID string, query ThreatEventTargetIndustryByDatasetListParams, opts ...option.RequestOption) (res *ThreatEventTargetIndustryByDatasetListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/dataset/%s/targetIndustries", query.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ThreatEventTargetIndustryByDatasetListResponse struct {
	Items ThreatEventTargetIndustryByDatasetListResponseItems `json:"items" api:"required"`
	Type  string                                              `json:"type" api:"required"`
	JSON  threatEventTargetIndustryByDatasetListResponseJSON  `json:"-"`
}

// threatEventTargetIndustryByDatasetListResponseJSON contains the JSON metadata
// for the struct [ThreatEventTargetIndustryByDatasetListResponse]
type threatEventTargetIndustryByDatasetListResponseJSON struct {
	Items       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTargetIndustryByDatasetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTargetIndustryByDatasetListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTargetIndustryByDatasetListResponseItems struct {
	Type string                                                  `json:"type" api:"required"`
	JSON threatEventTargetIndustryByDatasetListResponseItemsJSON `json:"-"`
}

// threatEventTargetIndustryByDatasetListResponseItemsJSON contains the JSON
// metadata for the struct [ThreatEventTargetIndustryByDatasetListResponseItems]
type threatEventTargetIndustryByDatasetListResponseItemsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTargetIndustryByDatasetListResponseItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTargetIndustryByDatasetListResponseItemsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTargetIndustryByDatasetListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

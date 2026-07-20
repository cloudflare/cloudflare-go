// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ThreatEventIndicatorByDatasetTagService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventIndicatorByDatasetTagService] method instead.
type ThreatEventIndicatorByDatasetTagService struct {
	Options []option.RequestOption
}

// NewThreatEventIndicatorByDatasetTagService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewThreatEventIndicatorByDatasetTagService(opts ...option.RequestOption) (r *ThreatEventIndicatorByDatasetTagService) {
	r = &ThreatEventIndicatorByDatasetTagService{}
	r.Options = opts
	return
}

// Returns all mirrored tags from the indicator dataset (DO mirror table). No
// pagination.
func (r *ThreatEventIndicatorByDatasetTagService) List(ctx context.Context, datasetID string, query ThreatEventIndicatorByDatasetTagListParams, opts ...option.RequestOption) (res *[]ThreatEventIndicatorByDatasetTagListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/dataset/%s/indicators/tags", query.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ThreatEventIndicatorByDatasetTagListResponse = interface{}

type ThreatEventIndicatorByDatasetTagListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

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

// ThreatEventTargetIndustryCatalogService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventTargetIndustryCatalogService] method instead.
type ThreatEventTargetIndustryCatalogService struct {
	Options []option.RequestOption
}

// NewThreatEventTargetIndustryCatalogService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewThreatEventTargetIndustryCatalogService(opts ...option.RequestOption) (r *ThreatEventTargetIndustryCatalogService) {
	r = &ThreatEventTargetIndustryCatalogService{}
	r.Options = opts
	return
}

// Lists all target industries from industry map catalog
func (r *ThreatEventTargetIndustryCatalogService) List(ctx context.Context, query ThreatEventTargetIndustryCatalogListParams, opts ...option.RequestOption) (res *ThreatEventTargetIndustryCatalogListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/targetIndustries/catalog", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ThreatEventTargetIndustryCatalogListResponse struct {
	Items ThreatEventTargetIndustryCatalogListResponseItems `json:"items" api:"required"`
	Type  string                                            `json:"type" api:"required"`
	JSON  threatEventTargetIndustryCatalogListResponseJSON  `json:"-"`
}

// threatEventTargetIndustryCatalogListResponseJSON contains the JSON metadata for
// the struct [ThreatEventTargetIndustryCatalogListResponse]
type threatEventTargetIndustryCatalogListResponseJSON struct {
	Items       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTargetIndustryCatalogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTargetIndustryCatalogListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTargetIndustryCatalogListResponseItems struct {
	Type string                                                `json:"type" api:"required"`
	JSON threatEventTargetIndustryCatalogListResponseItemsJSON `json:"-"`
}

// threatEventTargetIndustryCatalogListResponseItemsJSON contains the JSON metadata
// for the struct [ThreatEventTargetIndustryCatalogListResponseItems]
type threatEventTargetIndustryCatalogListResponseItemsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTargetIndustryCatalogListResponseItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTargetIndustryCatalogListResponseItemsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTargetIndustryCatalogListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

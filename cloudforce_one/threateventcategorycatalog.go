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

// ThreatEventCategoryCatalogService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventCategoryCatalogService] method instead.
type ThreatEventCategoryCatalogService struct {
	Options []option.RequestOption
}

// NewThreatEventCategoryCatalogService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewThreatEventCategoryCatalogService(opts ...option.RequestOption) (r *ThreatEventCategoryCatalogService) {
	r = &ThreatEventCategoryCatalogService{}
	r.Options = opts
	return
}

// List all categories stored in the account catalog.
func (r *ThreatEventCategoryCatalogService) List(ctx context.Context, query ThreatEventCategoryCatalogListParams, opts ...option.RequestOption) (res *[]ThreatEventCategoryCatalogListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/categories/catalog", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ThreatEventCategoryCatalogListResponse struct {
	KillChain   float64                                    `json:"killChain" api:"required"`
	Name        string                                     `json:"name" api:"required"`
	UUID        string                                     `json:"uuid" api:"required"`
	MitreAttack []string                                   `json:"mitreAttack"`
	MitreCapec  []string                                   `json:"mitreCapec"`
	Shortname   string                                     `json:"shortname"`
	JSON        threatEventCategoryCatalogListResponseJSON `json:"-"`
}

// threatEventCategoryCatalogListResponseJSON contains the JSON metadata for the
// struct [ThreatEventCategoryCatalogListResponse]
type threatEventCategoryCatalogListResponseJSON struct {
	KillChain   apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	MitreAttack apijson.Field
	MitreCapec  apijson.Field
	Shortname   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventCategoryCatalogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventCategoryCatalogListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventCategoryCatalogListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

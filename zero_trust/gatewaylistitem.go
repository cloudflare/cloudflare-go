// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// GatewayListItemService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayListItemService] method instead.
type GatewayListItemService struct {
	Options []option.RequestOption
}

// NewGatewayListItemService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayListItemService(opts ...option.RequestOption) (r *GatewayListItemService) {
	r = &GatewayListItemService{}
	r.Options = opts
	return
}

// Fetch all items in a single Zero Trust list.
func (r *GatewayListItemService) List(ctx context.Context, listID string, params GatewayListItemListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[GatewayItem], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if listID == "" {
		err = errors.New("missing required list_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s/items", params.AccountID, listID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetch all items in a single Zero Trust list.
func (r *GatewayListItemService) ListAutoPaging(ctx context.Context, listID string, params GatewayListItemListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[GatewayItem] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, listID, params, opts...))
}

type GatewayListItemListParams struct {
	// Specify the Cloudflare account identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [GatewayListItemListParams]'s query parameters as
// `url.Values`.
func (r GatewayListItemListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

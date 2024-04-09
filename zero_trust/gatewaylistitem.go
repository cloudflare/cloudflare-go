// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// GatewayListItemService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayListItemService] method
// instead.
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

// Fetches all items in a single Zero Trust list.
func (r *GatewayListItemService) List(ctx context.Context, listID string, query GatewayListItemListParams, opts ...option.RequestOption) (res *pagination.SinglePage[[]Lists], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/gateway/lists/%s/items", query.AccountID, listID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Fetches all items in a single Zero Trust list.
func (r *GatewayListItemService) ListAutoPaging(ctx context.Context, listID string, query GatewayListItemListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[[]Lists] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, listID, query, opts...))
}

type Lists struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The value of the item in a list.
	Value string    `json:"value"`
	JSON  listsJSON `json:"-"`
}

// listsJSON contains the JSON metadata for the struct [Lists]
type listsJSON struct {
	CreatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Lists) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r listsJSON) RawJSON() string {
	return r.raw
}

type ListsParam struct {
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r ListsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayListItemListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

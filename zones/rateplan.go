// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
)

// RatePlanService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRatePlanService] method instead.
type RatePlanService struct {
	Options []option.RequestOption
}

// NewRatePlanService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRatePlanService(opts ...option.RequestOption) (r *RatePlanService) {
	r = &RatePlanService{}
	r.Options = opts
	return
}

// Lists all rate plans the zone can subscribe to.
func (r *RatePlanService) Get(ctx context.Context, query RatePlanGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[RatePlanGetResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/available_rate_plans", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists all rate plans the zone can subscribe to.
func (r *RatePlanService) GetAutoPaging(ctx context.Context, query RatePlanGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RatePlanGetResponse] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, query, opts...))
}

type RatePlanGetResponse = interface{}

type RatePlanGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

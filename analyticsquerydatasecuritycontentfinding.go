// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// AnalyticsQueryDataSecurityContentFindingService contains methods and other
// services that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsQueryDataSecurityContentFindingService] method instead.
type AnalyticsQueryDataSecurityContentFindingService struct {
	Options []option.RequestOption
}

// NewAnalyticsQueryDataSecurityContentFindingService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAnalyticsQueryDataSecurityContentFindingService(opts ...option.RequestOption) (r *AnalyticsQueryDataSecurityContentFindingService) {
	r = &AnalyticsQueryDataSecurityContentFindingService{}
	r.Options = opts
	return
}

// Returns the top N integrations ranked by total content findings.
func (r *AnalyticsQueryDataSecurityContentFindingService) TopN(ctx context.Context, params AnalyticsQueryDataSecurityContentFindingTopNParams, opts ...option.RequestOption) (res *pagination.SinglePage[AnalyticsQueryDataSecurityContentFindingTopNResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/analytics/query/data-security/content-findings/top-n", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Returns the top N integrations ranked by total content findings.
func (r *AnalyticsQueryDataSecurityContentFindingService) TopNAutoPaging(ctx context.Context, params AnalyticsQueryDataSecurityContentFindingTopNParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AnalyticsQueryDataSecurityContentFindingTopNResponse] {
	return pagination.NewSinglePageAutoPager(r.TopN(ctx, params, opts...))
}

type AnalyticsQueryDataSecurityContentFindingTopNResponse map[string]interface{}

type AnalyticsQueryDataSecurityContentFindingTopNParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filters to apply. `findingType = content` is applied automatically for CASB
	// data.
	Filters param.Field[[]AnalyticsQueryDataSecurityContentFindingTopNParamsFilter] `json:"filters" api:"required"`
	// Start of the query time range (inclusive). RFC3339.
	From param.Field[time.Time] `json:"from" api:"required" format:"date-time"`
	// Maximum number of integrations to return.
	N param.Field[int64] `json:"n" api:"required"`
	// End of the query time range (exclusive). RFC3339.
	To param.Field[time.Time] `json:"to" api:"required" format:"date-time"`
}

func (r AnalyticsQueryDataSecurityContentFindingTopNParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A filter to apply to the query.
type AnalyticsQueryDataSecurityContentFindingTopNParamsFilter struct {
	// Specifies the column name to filter on. Requires a valid column for the target
	// dataset (e.g. `country`, `allowed`, `appId`).
	Name param.Field[string] `json:"name" api:"required"`
	// Filter operator. Common values: `eq`, `neq`, `in`, `not_in`, `gt`, `lt`, `gte`,
	// `lte`.
	Op param.Field[string] `json:"op" api:"required"`
	// Values to match against. Type depends on the column.
	Values param.Field[[]AnalyticsQueryDataSecurityContentFindingTopNParamsFiltersValueUnion] `json:"values" api:"required"`
}

func (r AnalyticsQueryDataSecurityContentFindingTopNParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionBool], [shared.UnionFloat].
type AnalyticsQueryDataSecurityContentFindingTopNParamsFiltersValueUnion interface {
	ImplementsAnalyticsQueryDataSecurityContentFindingTopNParamsFiltersValueUnion()
}

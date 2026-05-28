// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dls

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// RegionService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegionService] method instead.
type RegionService struct {
	Options []option.RequestOption
}

// NewRegionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRegionService(opts ...option.RequestOption) (r *RegionService) {
	r = &RegionService{}
	r.Options = opts
	return
}

// List DLS regions for an account
func (r *RegionService) List(ctx context.Context, params RegionListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[RegionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dls/regions", params.AccountID)
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

// List DLS regions for an account
func (r *RegionService) ListAutoPaging(ctx context.Context, params RegionListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[RegionListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Get a DLS region
func (r *RegionService) Get(ctx context.Context, regionID string, query RegionGetParams, opts ...option.RequestOption) (res *RegionGetResponse, err error) {
	var env RegionGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if regionID == "" {
		err = errors.New("missing required region_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dls/regions/%s", query.AccountID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type RegionListResponse struct {
	ID               string                 `json:"id" api:"required"`
	CreatedOn        time.Time              `json:"created_on" api:"required" format:"date-time"`
	ModifiedOn       time.Time              `json:"modified_on" api:"required" format:"date-time"`
	Name             string                 `json:"name" api:"required"`
	RegionKey        string                 `json:"region_key" api:"required"`
	Version          int64                  `json:"version" api:"required"`
	VersionCreatedOn time.Time              `json:"version_created_on" api:"required" format:"date-time"`
	JSON             regionListResponseJSON `json:"-"`
}

// regionListResponseJSON contains the JSON metadata for the struct
// [RegionListResponse]
type regionListResponseJSON struct {
	ID               apijson.Field
	CreatedOn        apijson.Field
	ModifiedOn       apijson.Field
	Name             apijson.Field
	RegionKey        apijson.Field
	Version          apijson.Field
	VersionCreatedOn apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionListResponseJSON) RawJSON() string {
	return r.raw
}

type RegionGetResponse struct {
	ID               string                `json:"id" api:"required"`
	CreatedOn        time.Time             `json:"created_on" api:"required" format:"date-time"`
	ModifiedOn       time.Time             `json:"modified_on" api:"required" format:"date-time"`
	Name             string                `json:"name" api:"required"`
	RegionKey        string                `json:"region_key" api:"required"`
	Version          int64                 `json:"version" api:"required"`
	VersionCreatedOn time.Time             `json:"version_created_on" api:"required" format:"date-time"`
	JSON             regionGetResponseJSON `json:"-"`
}

// regionGetResponseJSON contains the JSON metadata for the struct
// [RegionGetResponse]
type regionGetResponseJSON struct {
	ID               apijson.Field
	CreatedOn        apijson.Field
	ModifiedOn       apijson.Field
	Name             apijson.Field
	RegionKey        apijson.Field
	Version          apijson.Field
	VersionCreatedOn apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionGetResponseJSON) RawJSON() string {
	return r.raw
}

type RegionListParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Opaque token for cursor-based pagination. Omit for the first page. Pass the
	// value from a previous response to fetch the next page.
	Cursor  param.Field[string] `query:"cursor"`
	PerPage param.Field[int64]  `query:"per_page"`
	// Filter regions by type. Omit to return all regions.
	Type param.Field[RegionListParamsType] `query:"type"`
}

// URLQuery serializes [RegionListParams]'s query parameters as `url.Values`.
func (r RegionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter regions by type. Omit to return all regions.
type RegionListParamsType string

const (
	RegionListParamsTypeManaged RegionListParamsType = "managed"
	RegionListParamsTypeCustom  RegionListParamsType = "custom"
)

func (r RegionListParamsType) IsKnown() bool {
	switch r {
	case RegionListParamsTypeManaged, RegionListParamsTypeCustom:
		return true
	}
	return false
}

type RegionGetParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type RegionGetResponseEnvelope struct {
	Messages []shared.ResponseInfo         `json:"messages" api:"required"`
	Result   RegionGetResponse             `json:"result" api:"required"`
	Success  bool                          `json:"success" api:"required"`
	Errors   []shared.ResponseInfo         `json:"errors"`
	JSON     regionGetResponseEnvelopeJSON `json:"-"`
}

// regionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RegionGetResponseEnvelope]
type regionGetResponseEnvelopeJSON struct {
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

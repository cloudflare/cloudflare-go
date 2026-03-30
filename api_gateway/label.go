// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// LabelService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLabelService] method instead.
type LabelService struct {
	Options []option.RequestOption
	User    *LabelUserService
	Managed *LabelManagedService
}

// NewLabelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLabelService(opts ...option.RequestOption) (r *LabelService) {
	r = &LabelService{}
	r.Options = opts
	r.User = NewLabelUserService(opts...)
	r.Managed = NewLabelManagedService(opts...)
	return
}

// Retrieve all labels
func (r *LabelService) List(ctx context.Context, params LabelListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[LabelListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/api_gateway/labels", params.ZoneID)
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

// Retrieve all labels
func (r *LabelService) ListAutoPaging(ctx context.Context, params LabelListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[LabelListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type LabelListResponse struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The description of the label
	Description string    `json:"description" api:"required"`
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata" api:"required"`
	// The name of the label
	Name string `json:"name" api:"required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source LabelListResponseSource `json:"source" api:"required"`
	// Provides counts of what resources are linked to this label
	MappedResources interface{}           `json:"mapped_resources"`
	JSON            labelListResponseJSON `json:"-"`
}

// labelListResponseJSON contains the JSON metadata for the struct
// [LabelListResponse]
type labelListResponseJSON struct {
	CreatedAt       apijson.Field
	Description     apijson.Field
	LastUpdated     apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Source          apijson.Field
	MappedResources apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LabelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelListResponseJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type LabelListResponseSource string

const (
	LabelListResponseSourceUser    LabelListResponseSource = "user"
	LabelListResponseSourceManaged LabelListResponseSource = "managed"
)

func (r LabelListResponseSource) IsKnown() bool {
	switch r {
	case LabelListResponseSourceUser, LabelListResponseSourceManaged:
		return true
	}
	return false
}

type LabelListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Direction to order results.
	Direction param.Field[LabelListParamsDirection] `query:"direction"`
	// Filter for labels where the name or description matches using substring match
	Filter param.Field[string] `query:"filter"`
	// Field to order by
	Order param.Field[LabelListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter for labels with source
	Source param.Field[LabelListParamsSource] `query:"source"`
	// Include `mapped_resources` for each label
	WithMappedResourceCounts param.Field[bool] `query:"with_mapped_resource_counts"`
}

// URLQuery serializes [LabelListParams]'s query parameters as `url.Values`.
func (r LabelListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
type LabelListParamsDirection string

const (
	LabelListParamsDirectionAsc  LabelListParamsDirection = "asc"
	LabelListParamsDirectionDesc LabelListParamsDirection = "desc"
)

func (r LabelListParamsDirection) IsKnown() bool {
	switch r {
	case LabelListParamsDirectionAsc, LabelListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order by
type LabelListParamsOrder string

const (
	LabelListParamsOrderName                      LabelListParamsOrder = "name"
	LabelListParamsOrderDescription               LabelListParamsOrder = "description"
	LabelListParamsOrderCreatedAt                 LabelListParamsOrder = "created_at"
	LabelListParamsOrderLastUpdated               LabelListParamsOrder = "last_updated"
	LabelListParamsOrderMappedResourcesOperations LabelListParamsOrder = "mapped_resources.operations"
)

func (r LabelListParamsOrder) IsKnown() bool {
	switch r {
	case LabelListParamsOrderName, LabelListParamsOrderDescription, LabelListParamsOrderCreatedAt, LabelListParamsOrderLastUpdated, LabelListParamsOrderMappedResourcesOperations:
		return true
	}
	return false
}

// Filter for labels with source
type LabelListParamsSource string

const (
	LabelListParamsSourceUser    LabelListParamsSource = "user"
	LabelListParamsSourceManaged LabelListParamsSource = "managed"
)

func (r LabelListParamsSource) IsKnown() bool {
	switch r {
	case LabelListParamsSourceUser, LabelListParamsSourceManaged:
		return true
	}
	return false
}

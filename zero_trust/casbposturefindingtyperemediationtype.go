// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// CasbPostureFindingTypeRemediationTypeService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbPostureFindingTypeRemediationTypeService] method instead.
type CasbPostureFindingTypeRemediationTypeService struct {
	Options []option.RequestOption
}

// NewCasbPostureFindingTypeRemediationTypeService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewCasbPostureFindingTypeRemediationTypeService(opts ...option.RequestOption) (r *CasbPostureFindingTypeRemediationTypeService) {
	r = &CasbPostureFindingTypeRemediationTypeService{}
	r.Options = opts
	return
}

// List all remediation types for a given finding type. This endpoint supports both
// cursor and offset pagination. Note that `cursor` and `page` are mutually
// exclusive.
func (r *CasbPostureFindingTypeRemediationTypeService) List(ctx context.Context, findingTypeID string, params CasbPostureFindingTypeRemediationTypeListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[CasbPostureFindingTypeRemediationTypeListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if findingTypeID == "" {
		err = errors.New("missing required finding_type_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/finding_types/%s/remediation_types", params.AccountID, findingTypeID)
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

// List all remediation types for a given finding type. This endpoint supports both
// cursor and offset pagination. Note that `cursor` and `page` are mutually
// exclusive.
func (r *CasbPostureFindingTypeRemediationTypeService) ListAutoPaging(ctx context.Context, findingTypeID string, params CasbPostureFindingTypeRemediationTypeListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[CasbPostureFindingTypeRemediationTypeListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, findingTypeID, params, opts...))
}

// Information about a remediation type.
type CasbPostureFindingTypeRemediationTypeListResponse struct {
	// The identifier for the remediation type.
	ID string `json:"id" api:"required" format:"uuid"`
	// A description of the action(s) taken by the remediation type.
	Description string `json:"description" api:"required"`
	// The name of the remediation type as displayed in the cloudflare dashboard.
	DisplayName string `json:"display_name" api:"required"`
	// The identifier of the finding_type which this remediation type should remediate.
	FindingTypeID string `json:"finding_type_id" api:"required" format:"uuid"`
	// The name of the remediation type.
	RemediationType string                                                `json:"remediation_type" api:"required"`
	JSON            casbPostureFindingTypeRemediationTypeListResponseJSON `json:"-"`
}

// casbPostureFindingTypeRemediationTypeListResponseJSON contains the JSON metadata
// for the struct [CasbPostureFindingTypeRemediationTypeListResponse]
type casbPostureFindingTypeRemediationTypeListResponseJSON struct {
	ID              apijson.Field
	Description     apijson.Field
	DisplayName     apijson.Field
	FindingTypeID   apijson.Field
	RemediationType apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CasbPostureFindingTypeRemediationTypeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingTypeRemediationTypeListResponseJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingTypeRemediationTypeListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A cursor for pagination.
	Cursor param.Field[string] `query:"cursor"`
	// Filter by an integration ID
	IntegrationID param.Field[string] `query:"integration_id" format:"uuid"`
	// A page number within the paginated result set.
	Page param.Field[int64] `query:"page"`
	// Number of results to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [CasbPostureFindingTypeRemediationTypeListParams]'s query
// parameters as `url.Values`.
func (r CasbPostureFindingTypeRemediationTypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

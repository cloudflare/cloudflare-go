// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// OrganizationRailgunService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewOrganizationRailgunService]
// method instead.
type OrganizationRailgunService struct {
	Options []option.RequestOption
	Zones   *OrganizationRailgunZoneService
}

// NewOrganizationRailgunService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationRailgunService(opts ...option.RequestOption) (r *OrganizationRailgunService) {
	r = &OrganizationRailgunService{}
	r.Options = opts
	r.Zones = NewOrganizationRailgunZoneService(opts...)
	return
}

// Railgun details
func (r *OrganizationRailgunService) Get(ctx context.Context, organizationIdentifier string, identifier string, opts ...option.RequestOption) (res *RailgunResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/railguns/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable a Railgun for all zones connected to it.
func (r *OrganizationRailgunService) Update(ctx context.Context, organizationIdentifier string, identifier string, body OrganizationRailgunUpdateParams, opts ...option.RequestOption) (res *RailgunResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/railguns/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Disable and delete a Railgun. This will immediately disable the Railgun for any
// connected zones.
func (r *OrganizationRailgunService) Delete(ctx context.Context, organizationIdentifier string, identifier string, opts ...option.RequestOption) (res *RailgunResponseSingleID, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/railguns/%s", organizationIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create Railgun
func (r *OrganizationRailgunService) OrganizationRailgunNewRailgun(ctx context.Context, organizationIdentifier string, body OrganizationRailgunOrganizationRailgunNewRailgunParams, opts ...option.RequestOption) (res *RailgunResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/railguns", organizationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List, search, sort and filter your Railguns.
func (r *OrganizationRailgunService) OrganizationRailgunListRailguns(ctx context.Context, organizationIdentifier string, query OrganizationRailgunOrganizationRailgunListRailgunsParams, opts ...option.RequestOption) (res *RailgunResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("organizations/%s/railguns", organizationIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OrganizationRailgunUpdateParams struct {
	// Flag to determine if the Railgun is accepting connections.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r OrganizationRailgunUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationRailgunOrganizationRailgunNewRailgunParams struct {
	// Readable identifier of the Railgun.
	Name param.Field[string] `json:"name,required"`
}

func (r OrganizationRailgunOrganizationRailgunNewRailgunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationRailgunOrganizationRailgunListRailgunsParams struct {
	// Sort Railguns in ascending or descending order.
	Direction param.Field[OrganizationRailgunOrganizationRailgunListRailgunsParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [OrganizationRailgunOrganizationRailgunListRailgunsParams]'s
// query parameters as `url.Values`.
func (r OrganizationRailgunOrganizationRailgunListRailgunsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort Railguns in ascending or descending order.
type OrganizationRailgunOrganizationRailgunListRailgunsParamsDirection string

const (
	OrganizationRailgunOrganizationRailgunListRailgunsParamsDirectionAsc  OrganizationRailgunOrganizationRailgunListRailgunsParamsDirection = "asc"
	OrganizationRailgunOrganizationRailgunListRailgunsParamsDirectionDesc OrganizationRailgunOrganizationRailgunListRailgunsParamsDirection = "desc"
)

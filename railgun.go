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

// RailgunService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRailgunService] method instead.
type RailgunService struct {
	Options []option.RequestOption
	Zones   *RailgunZoneService
}

// NewRailgunService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRailgunService(opts ...option.RequestOption) (r *RailgunService) {
	r = &RailgunService{}
	r.Options = opts
	r.Zones = NewRailgunZoneService(opts...)
	return
}

// Railgun details
func (r *RailgunService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *RailgunResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("railguns/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Enable or disable a Railgun for all zones connected to it.
func (r *RailgunService) Update(ctx context.Context, identifier string, body RailgunUpdateParams, opts ...option.RequestOption) (res *RailgunResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("railguns/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Disable and delete a Railgun. This will immediately disable that Railgun for any
// connected zones.
func (r *RailgunService) Delete(ctx context.Context, identifier string, opts ...option.RequestOption) (res *RailgunResponseSingleID, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("railguns/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create Railgun
func (r *RailgunService) RailgunNewRailgun(ctx context.Context, body RailgunRailgunNewRailgunParams, opts ...option.RequestOption) (res *RailgunResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := "railguns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List, search, sort and filter your Railguns.
func (r *RailgunService) RailgunListRailguns(ctx context.Context, query RailgunRailgunListRailgunsParams, opts ...option.RequestOption) (res *RailgunResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := "railguns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RailgunUpdateParams struct {
	// Flag to determine if the Railgun is accepting connections.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r RailgunUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RailgunRailgunNewRailgunParams struct {
	// Readable identifier of the Railgun.
	Name param.Field[string] `json:"name,required"`
}

func (r RailgunRailgunNewRailgunParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RailgunRailgunListRailgunsParams struct {
	// Sort Railguns in ascending or descending order.
	Direction param.Field[RailgunRailgunListRailgunsParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RailgunRailgunListRailgunsParams]'s query parameters as
// `url.Values`.
func (r RailgunRailgunListRailgunsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort Railguns in ascending or descending order.
type RailgunRailgunListRailgunsParamsDirection string

const (
	RailgunRailgunListRailgunsParamsDirectionAsc  RailgunRailgunListRailgunsParamsDirection = "asc"
	RailgunRailgunListRailgunsParamsDirectionDesc RailgunRailgunListRailgunsParamsDirection = "desc"
)

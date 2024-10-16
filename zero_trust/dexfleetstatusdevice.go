// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
)

// DEXFleetStatusDeviceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXFleetStatusDeviceService] method instead.
type DEXFleetStatusDeviceService struct {
	Options []option.RequestOption
}

// NewDEXFleetStatusDeviceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDEXFleetStatusDeviceService(opts ...option.RequestOption) (r *DEXFleetStatusDeviceService) {
	r = &DEXFleetStatusDeviceService{}
	r.Options = opts
	return
}

// List details for devices using WARP
func (r *DEXFleetStatusDeviceService) List(ctx context.Context, params DEXFleetStatusDeviceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[DEXFleetStatusDeviceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/devices", params.AccountID)
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

// List details for devices using WARP
func (r *DEXFleetStatusDeviceService) ListAutoPaging(ctx context.Context, params DEXFleetStatusDeviceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[DEXFleetStatusDeviceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type DEXFleetStatusDeviceListResponse struct {
	// Cloudflare colo
	Colo string `json:"colo,required"`
	// Device identifier (UUID v4)
	DeviceID string `json:"deviceId,required"`
	// Operating system
	Platform string `json:"platform,required"`
	// Network status
	Status string `json:"status,required"`
	// WARP client version
	Version string `json:"version,required"`
	// Device identifier (human readable)
	DeviceName string `json:"deviceName"`
	// User contact email address
	PersonEmail string                               `json:"personEmail"`
	JSON        dexFleetStatusDeviceListResponseJSON `json:"-"`
}

// dexFleetStatusDeviceListResponseJSON contains the JSON metadata for the struct
// [DEXFleetStatusDeviceListResponse]
type dexFleetStatusDeviceListResponseJSON struct {
	Colo        apijson.Field
	DeviceID    apijson.Field
	Platform    apijson.Field
	Status      apijson.Field
	Version     apijson.Field
	DeviceName  apijson.Field
	PersonEmail apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXFleetStatusDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexFleetStatusDeviceListResponseJSON) RawJSON() string {
	return r.raw
}

type DEXFleetStatusDeviceListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Time range beginning in ISO format
	From param.Field[string] `query:"from,required"`
	// Page number
	Page param.Field[float64] `query:"page,required"`
	// Number of results per page
	PerPage param.Field[float64] `query:"per_page,required"`
	// Time range end in ISO format
	To param.Field[string] `query:"to,required"`
	// Cloudflare colo
	Colo param.Field[string] `query:"colo"`
	// Device-specific ID, given as UUID v4
	DeviceID param.Field[string] `query:"device_id"`
	// The mode under which the WARP client is run
	Mode param.Field[string] `query:"mode"`
	// Operating system
	Platform param.Field[string] `query:"platform"`
	// Dimension to sort results by
	SortBy param.Field[DEXFleetStatusDeviceListParamsSortBy] `query:"sort_by"`
	// Source:
	//
	// - `hourly` - device details aggregated hourly, up to 7 days prior
	// - `last_seen` - device details, up to 24 hours prior
	// - `raw` - device details, up to 7 days prior
	Source param.Field[DEXFleetStatusDeviceListParamsSource] `query:"source"`
	// Network status
	Status param.Field[string] `query:"status"`
	// WARP client version
	Version param.Field[string] `query:"version"`
}

// URLQuery serializes [DEXFleetStatusDeviceListParams]'s query parameters as
// `url.Values`.
func (r DEXFleetStatusDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Dimension to sort results by
type DEXFleetStatusDeviceListParamsSortBy string

const (
	DEXFleetStatusDeviceListParamsSortByColo      DEXFleetStatusDeviceListParamsSortBy = "colo"
	DEXFleetStatusDeviceListParamsSortByDeviceID  DEXFleetStatusDeviceListParamsSortBy = "device_id"
	DEXFleetStatusDeviceListParamsSortByMode      DEXFleetStatusDeviceListParamsSortBy = "mode"
	DEXFleetStatusDeviceListParamsSortByPlatform  DEXFleetStatusDeviceListParamsSortBy = "platform"
	DEXFleetStatusDeviceListParamsSortByStatus    DEXFleetStatusDeviceListParamsSortBy = "status"
	DEXFleetStatusDeviceListParamsSortByTimestamp DEXFleetStatusDeviceListParamsSortBy = "timestamp"
	DEXFleetStatusDeviceListParamsSortByVersion   DEXFleetStatusDeviceListParamsSortBy = "version"
)

func (r DEXFleetStatusDeviceListParamsSortBy) IsKnown() bool {
	switch r {
	case DEXFleetStatusDeviceListParamsSortByColo, DEXFleetStatusDeviceListParamsSortByDeviceID, DEXFleetStatusDeviceListParamsSortByMode, DEXFleetStatusDeviceListParamsSortByPlatform, DEXFleetStatusDeviceListParamsSortByStatus, DEXFleetStatusDeviceListParamsSortByTimestamp, DEXFleetStatusDeviceListParamsSortByVersion:
		return true
	}
	return false
}

// Source:
//
// - `hourly` - device details aggregated hourly, up to 7 days prior
// - `last_seen` - device details, up to 24 hours prior
// - `raw` - device details, up to 7 days prior
type DEXFleetStatusDeviceListParamsSource string

const (
	DEXFleetStatusDeviceListParamsSourceLastSeen DEXFleetStatusDeviceListParamsSource = "last_seen"
	DEXFleetStatusDeviceListParamsSourceHourly   DEXFleetStatusDeviceListParamsSource = "hourly"
	DEXFleetStatusDeviceListParamsSourceRaw      DEXFleetStatusDeviceListParamsSource = "raw"
)

func (r DEXFleetStatusDeviceListParamsSource) IsKnown() bool {
	switch r {
	case DEXFleetStatusDeviceListParamsSourceLastSeen, DEXFleetStatusDeviceListParamsSourceHourly, DEXFleetStatusDeviceListParamsSourceRaw:
		return true
	}
	return false
}

// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
)

// ZeroTrustDEXFleetStatusDeviceService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustDEXFleetStatusDeviceService] method instead.
type ZeroTrustDEXFleetStatusDeviceService struct {
	Options []option.RequestOption
}

// NewZeroTrustDEXFleetStatusDeviceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustDEXFleetStatusDeviceService(opts ...option.RequestOption) (r *ZeroTrustDEXFleetStatusDeviceService) {
	r = &ZeroTrustDEXFleetStatusDeviceService{}
	r.Options = opts
	return
}

// List details for devices using WARP
func (r *ZeroTrustDEXFleetStatusDeviceService) List(ctx context.Context, params ZeroTrustDEXFleetStatusDeviceListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[ZeroTrustDEXFleetStatusDeviceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
func (r *ZeroTrustDEXFleetStatusDeviceService) ListAutoPaging(ctx context.Context, params ZeroTrustDEXFleetStatusDeviceListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[ZeroTrustDEXFleetStatusDeviceListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type ZeroTrustDEXFleetStatusDeviceListResponse struct {
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
	PersonEmail string                                        `json:"personEmail"`
	JSON        zeroTrustDEXFleetStatusDeviceListResponseJSON `json:"-"`
}

// zeroTrustDEXFleetStatusDeviceListResponseJSON contains the JSON metadata for the
// struct [ZeroTrustDEXFleetStatusDeviceListResponse]
type zeroTrustDEXFleetStatusDeviceListResponseJSON struct {
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

func (r *ZeroTrustDEXFleetStatusDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustDEXFleetStatusDeviceListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustDEXFleetStatusDeviceListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number of paginated results
	Page param.Field[float64] `query:"page,required"`
	// Number of items per page
	PerPage param.Field[float64] `query:"per_page,required"`
	// Timestamp in ISO format
	TimeEnd param.Field[string] `query:"time_end,required"`
	// Timestamp in ISO format
	TimeStart param.Field[string] `query:"time_start,required"`
	// Cloudflare colo
	Colo param.Field[string] `query:"colo"`
	// Device-specific ID, given as UUID v4
	DeviceID param.Field[string] `query:"device_id"`
	// The mode under which the WARP client is run
	Mode param.Field[string] `query:"mode"`
	// Operating system
	Platform param.Field[string] `query:"platform"`
	// Dimension to sort results by
	SortBy param.Field[ZeroTrustDEXFleetStatusDeviceListParamsSortBy] `query:"sort_by"`
	// Network status
	Status param.Field[string] `query:"status"`
	// WARP client version
	Version param.Field[string] `query:"version"`
}

// URLQuery serializes [ZeroTrustDEXFleetStatusDeviceListParams]'s query parameters
// as `url.Values`.
func (r ZeroTrustDEXFleetStatusDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Dimension to sort results by
type ZeroTrustDEXFleetStatusDeviceListParamsSortBy string

const (
	ZeroTrustDEXFleetStatusDeviceListParamsSortByColo      ZeroTrustDEXFleetStatusDeviceListParamsSortBy = "colo"
	ZeroTrustDEXFleetStatusDeviceListParamsSortByDeviceID  ZeroTrustDEXFleetStatusDeviceListParamsSortBy = "device_id"
	ZeroTrustDEXFleetStatusDeviceListParamsSortByMode      ZeroTrustDEXFleetStatusDeviceListParamsSortBy = "mode"
	ZeroTrustDEXFleetStatusDeviceListParamsSortByPlatform  ZeroTrustDEXFleetStatusDeviceListParamsSortBy = "platform"
	ZeroTrustDEXFleetStatusDeviceListParamsSortByStatus    ZeroTrustDEXFleetStatusDeviceListParamsSortBy = "status"
	ZeroTrustDEXFleetStatusDeviceListParamsSortByTimestamp ZeroTrustDEXFleetStatusDeviceListParamsSortBy = "timestamp"
	ZeroTrustDEXFleetStatusDeviceListParamsSortByVersion   ZeroTrustDEXFleetStatusDeviceListParamsSortBy = "version"
)

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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountDexFleetStatusDeviceService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountDexFleetStatusDeviceService] method instead.
type AccountDexFleetStatusDeviceService struct {
	Options []option.RequestOption
}

// NewAccountDexFleetStatusDeviceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDexFleetStatusDeviceService(opts ...option.RequestOption) (r *AccountDexFleetStatusDeviceService) {
	r = &AccountDexFleetStatusDeviceService{}
	r.Options = opts
	return
}

// List details for devices using WARP
func (r *AccountDexFleetStatusDeviceService) List(ctx context.Context, accountIdentifier string, query AccountDexFleetStatusDeviceListParams, opts ...option.RequestOption) (res *shared.Page[AccountDexFleetStatusDeviceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/dex/fleet-status/devices", accountIdentifier)
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

type AccountDexFleetStatusDeviceListResponse struct {
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
	PersonEmail string                                      `json:"personEmail"`
	JSON        accountDexFleetStatusDeviceListResponseJSON `json:"-"`
}

// accountDexFleetStatusDeviceListResponseJSON contains the JSON metadata for the
// struct [AccountDexFleetStatusDeviceListResponse]
type accountDexFleetStatusDeviceListResponseJSON struct {
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

func (r *AccountDexFleetStatusDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDexFleetStatusDeviceListParams struct {
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
	SortBy param.Field[AccountDexFleetStatusDeviceListParamsSortBy] `query:"sort_by"`
	// Network status
	Status param.Field[string] `query:"status"`
	// WARP client version
	Version param.Field[string] `query:"version"`
}

// URLQuery serializes [AccountDexFleetStatusDeviceListParams]'s query parameters
// as `url.Values`.
func (r AccountDexFleetStatusDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Dimension to sort results by
type AccountDexFleetStatusDeviceListParamsSortBy string

const (
	AccountDexFleetStatusDeviceListParamsSortByColo      AccountDexFleetStatusDeviceListParamsSortBy = "colo"
	AccountDexFleetStatusDeviceListParamsSortByDeviceID  AccountDexFleetStatusDeviceListParamsSortBy = "device_id"
	AccountDexFleetStatusDeviceListParamsSortByMode      AccountDexFleetStatusDeviceListParamsSortBy = "mode"
	AccountDexFleetStatusDeviceListParamsSortByPlatform  AccountDexFleetStatusDeviceListParamsSortBy = "platform"
	AccountDexFleetStatusDeviceListParamsSortByStatus    AccountDexFleetStatusDeviceListParamsSortBy = "status"
	AccountDexFleetStatusDeviceListParamsSortByTimestamp AccountDexFleetStatusDeviceListParamsSortBy = "timestamp"
	AccountDexFleetStatusDeviceListParamsSortByVersion   AccountDexFleetStatusDeviceListParamsSortBy = "version"
)

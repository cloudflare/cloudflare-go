// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountDevicePolicyIncludeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountDevicePolicyIncludeService] method instead.
type AccountDevicePolicyIncludeService struct {
	Options []option.RequestOption
}

// NewAccountDevicePolicyIncludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDevicePolicyIncludeService(opts ...option.RequestOption) (r *AccountDevicePolicyIncludeService) {
	r = &AccountDevicePolicyIncludeService{}
	r.Options = opts
	return
}

// Get the list of routes included in the WARP client's tunnel.
func (r *AccountDevicePolicyIncludeService) DevicesGetSplitTunnelIncludeList(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *SplitTunnelIncludeResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/include", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the list of routes included in the WARP client's tunnel for one specific
// device settings policy.
func (r *AccountDevicePolicyIncludeService) DevicesGetSplitTunnelIncludeListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *SplitTunnelIncludeResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/include", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set the list of routes included in the WARP client's tunnel.
func (r *AccountDevicePolicyIncludeService) DevicesSetSplitTunnelIncludeList(ctx context.Context, identifier interface{}, body AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListParams, opts ...option.RequestOption) (res *SplitTunnelIncludeResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/include", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Set the list of routes included in the WARP client's tunnel for one specific
// device settings policy.
func (r *AccountDevicePolicyIncludeService) DevicesSetSplitTunnelIncludeListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, body AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParams, opts ...option.RequestOption) (res *SplitTunnelIncludeResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/include", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SplitTunnelIncludeResponseCollection struct {
	Errors     []SplitTunnelIncludeResponseCollectionError    `json:"errors"`
	Messages   []SplitTunnelIncludeResponseCollectionMessage  `json:"messages"`
	Result     []SplitTunnelIncludeResponseCollectionResult   `json:"result"`
	ResultInfo SplitTunnelIncludeResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SplitTunnelIncludeResponseCollectionSuccess `json:"success"`
	JSON    splitTunnelIncludeResponseCollectionJSON    `json:"-"`
}

// splitTunnelIncludeResponseCollectionJSON contains the JSON metadata for the
// struct [SplitTunnelIncludeResponseCollection]
type splitTunnelIncludeResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelIncludeResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SplitTunnelIncludeResponseCollectionError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    splitTunnelIncludeResponseCollectionErrorJSON `json:"-"`
}

// splitTunnelIncludeResponseCollectionErrorJSON contains the JSON metadata for the
// struct [SplitTunnelIncludeResponseCollectionError]
type splitTunnelIncludeResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelIncludeResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SplitTunnelIncludeResponseCollectionMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    splitTunnelIncludeResponseCollectionMessageJSON `json:"-"`
}

// splitTunnelIncludeResponseCollectionMessageJSON contains the JSON metadata for
// the struct [SplitTunnelIncludeResponseCollectionMessage]
type splitTunnelIncludeResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelIncludeResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SplitTunnelIncludeResponseCollectionResult struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host string                                         `json:"host"`
	JSON splitTunnelIncludeResponseCollectionResultJSON `json:"-"`
}

// splitTunnelIncludeResponseCollectionResultJSON contains the JSON metadata for
// the struct [SplitTunnelIncludeResponseCollectionResult]
type splitTunnelIncludeResponseCollectionResultJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelIncludeResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SplitTunnelIncludeResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                            `json:"total_count"`
	JSON       splitTunnelIncludeResponseCollectionResultInfoJSON `json:"-"`
}

// splitTunnelIncludeResponseCollectionResultInfoJSON contains the JSON metadata
// for the struct [SplitTunnelIncludeResponseCollectionResultInfo]
type splitTunnelIncludeResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelIncludeResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SplitTunnelIncludeResponseCollectionSuccess bool

const (
	SplitTunnelIncludeResponseCollectionSuccessTrue SplitTunnelIncludeResponseCollectionSuccess = true
)

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListParams struct {
	Body param.Field[[]AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListParamsBody] `json:"body,required"`
}

func (r AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListParamsBody struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host param.Field[string] `json:"host"`
}

func (r AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParams struct {
	Body param.Field[[]AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParamsBody] `json:"body,required"`
}

func (r AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParamsBody struct {
	// The address in CIDR format to include in the tunnel. If address is present, host
	// must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to include in the tunnel. If host is present, address must not
	// be present.
	Host param.Field[string] `json:"host"`
}

func (r AccountDevicePolicyIncludeDevicesSetSplitTunnelIncludeListForADeviceSettingsPolicyParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

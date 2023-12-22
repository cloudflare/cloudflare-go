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

// AccountDevicePolicyExcludeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountDevicePolicyExcludeService] method instead.
type AccountDevicePolicyExcludeService struct {
	Options []option.RequestOption
}

// NewAccountDevicePolicyExcludeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDevicePolicyExcludeService(opts ...option.RequestOption) (r *AccountDevicePolicyExcludeService) {
	r = &AccountDevicePolicyExcludeService{}
	r.Options = opts
	return
}

// Get the list of routes excluded from the WARP client's tunnel.
func (r *AccountDevicePolicyExcludeService) DevicesGetSplitTunnelExcludeList(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *SplitTunnelResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/exclude", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the list of routes excluded from the WARP client's tunnel for one specific
// device settings policy.
func (r *AccountDevicePolicyExcludeService) DevicesGetSplitTunnelExcludeListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *SplitTunnelResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/exclude", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Set the list of routes excluded from the WARP client's tunnel.
func (r *AccountDevicePolicyExcludeService) DevicesSetSplitTunnelExcludeList(ctx context.Context, identifier interface{}, body AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListParams, opts ...option.RequestOption) (res *SplitTunnelResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/exclude", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Set the list of routes excluded from the WARP client's tunnel for one specific
// device settings policy.
func (r *AccountDevicePolicyExcludeService) DevicesSetSplitTunnelExcludeListForADeviceSettingsPolicy(ctx context.Context, identifier interface{}, uuid string, body AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParams, opts ...option.RequestOption) (res *SplitTunnelResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/policy/%s/exclude", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SplitTunnelResponseCollection struct {
	Errors     []SplitTunnelResponseCollectionError    `json:"errors"`
	Messages   []SplitTunnelResponseCollectionMessage  `json:"messages"`
	Result     []SplitTunnelResponseCollectionResult   `json:"result"`
	ResultInfo SplitTunnelResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SplitTunnelResponseCollectionSuccess `json:"success"`
	JSON    splitTunnelResponseCollectionJSON    `json:"-"`
}

// splitTunnelResponseCollectionJSON contains the JSON metadata for the struct
// [SplitTunnelResponseCollection]
type splitTunnelResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SplitTunnelResponseCollectionError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    splitTunnelResponseCollectionErrorJSON `json:"-"`
}

// splitTunnelResponseCollectionErrorJSON contains the JSON metadata for the struct
// [SplitTunnelResponseCollectionError]
type splitTunnelResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SplitTunnelResponseCollectionMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    splitTunnelResponseCollectionMessageJSON `json:"-"`
}

// splitTunnelResponseCollectionMessageJSON contains the JSON metadata for the
// struct [SplitTunnelResponseCollectionMessage]
type splitTunnelResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SplitTunnelResponseCollectionResult struct {
	// The address in CIDR format to exclude from the tunnel. If address is present,
	// host must not be present.
	Address string `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description string `json:"description,required"`
	// The domain name to exclude from the tunnel. If host is present, address must not
	// be present.
	Host string                                  `json:"host"`
	JSON splitTunnelResponseCollectionResultJSON `json:"-"`
}

// splitTunnelResponseCollectionResultJSON contains the JSON metadata for the
// struct [SplitTunnelResponseCollectionResult]
type splitTunnelResponseCollectionResultJSON struct {
	Address     apijson.Field
	Description apijson.Field
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SplitTunnelResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       splitTunnelResponseCollectionResultInfoJSON `json:"-"`
}

// splitTunnelResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [SplitTunnelResponseCollectionResultInfo]
type splitTunnelResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SplitTunnelResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SplitTunnelResponseCollectionSuccess bool

const (
	SplitTunnelResponseCollectionSuccessTrue SplitTunnelResponseCollectionSuccess = true
)

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListParams struct {
	Body param.Field[[]AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListParamsBody] `json:"body,required"`
}

func (r AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListParamsBody struct {
	// The address in CIDR format to exclude from the tunnel. If address is present,
	// host must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to exclude from the tunnel. If host is present, address must not
	// be present.
	Host param.Field[string] `json:"host"`
}

func (r AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParams struct {
	Body param.Field[[]AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParamsBody] `json:"body,required"`
}

func (r AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParamsBody struct {
	// The address in CIDR format to exclude from the tunnel. If address is present,
	// host must not be present.
	Address param.Field[string] `json:"address,required"`
	// A description of the split tunnel item, displayed in the client UI.
	Description param.Field[string] `json:"description,required"`
	// The domain name to exclude from the tunnel. If host is present, address must not
	// be present.
	Host param.Field[string] `json:"host"`
}

func (r AccountDevicePolicyExcludeDevicesSetSplitTunnelExcludeListForADeviceSettingsPolicyParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

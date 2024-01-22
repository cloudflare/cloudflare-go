// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountDeviceOverrideCodeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountDeviceOverrideCodeService] method instead.
type AccountDeviceOverrideCodeService struct {
	Options []option.RequestOption
}

// NewAccountDeviceOverrideCodeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountDeviceOverrideCodeService(opts ...option.RequestOption) (r *AccountDeviceOverrideCodeService) {
	r = &AccountDeviceOverrideCodeService{}
	r.Options = opts
	return
}

// Fetches a one-time use admin override code for a device. This relies on the
// **Admin Override** setting being enabled in your device configuration.
func (r *AccountDeviceOverrideCodeService) DevicesListAdminOverrideCodeForDevice(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/%s/override_codes", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponse struct {
	Errors     []AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseError    `json:"errors"`
	Messages   []AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseMessage  `json:"messages"`
	Result     AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResult     `json:"result"`
	ResultInfo AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultInfo `json:"result_info"`
	// Whether the API call was successful.
	Success AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseSuccess `json:"success"`
	JSON    accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseJSON    `json:"-"`
}

// accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseJSON
// contains the JSON metadata for the struct
// [AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponse]
type accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseError struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseErrorJSON `json:"-"`
}

// accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseError]
type accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseMessage struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseMessageJSON `json:"-"`
}

// accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseMessage]
type accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResult struct {
	DisableForTime AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultDisableForTime `json:"disable_for_time"`
	JSON           accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultJSON           `json:"-"`
}

// accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultJSON
// contains the JSON metadata for the struct
// [AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResult]
type accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultJSON struct {
	DisableForTime apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultDisableForTime struct {
	// Override code that is valid for 1 hour.
	Number1 interface{} `json:"1"`
	// Override code that is valid for 12 hour2.
	Number12 interface{} `json:"12"`
	// Override code that is valid for 24 hour.2.
	Number24 interface{} `json:"24"`
	// Override code that is valid for 3 hours.
	Number3 interface{} `json:"3"`
	// Override code that is valid for 6 hours.
	Number6 interface{}                                                                                    `json:"6"`
	JSON    accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultDisableForTimeJSON `json:"-"`
}

// accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultDisableForTimeJSON
// contains the JSON metadata for the struct
// [AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultDisableForTime]
type accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultDisableForTimeJSON struct {
	Number1     apijson.Field
	Number12    apijson.Field
	Number24    apijson.Field
	Number3     apijson.Field
	Number6     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultDisableForTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                              `json:"total_count"`
	JSON       accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultInfoJSON `json:"-"`
}

// accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultInfo]
type accountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseSuccess bool

const (
	AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseSuccessTrue AccountDeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseSuccess = true
)

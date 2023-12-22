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

// Fetch one-time use admin override code for a device. This relies on the Admin
// Override setting being enabled in your device configuration.
func (r *AccountDeviceOverrideCodeService) DevicesListAdminOverrideCodeForDevice(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *OverrideCodesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/%s/override_codes", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type OverrideCodesResponse struct {
	Errors     []OverrideCodesResponseError    `json:"errors"`
	Messages   []OverrideCodesResponseMessage  `json:"messages"`
	Result     OverrideCodesResponseResult     `json:"result"`
	ResultInfo OverrideCodesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success OverrideCodesResponseSuccess `json:"success"`
	JSON    overrideCodesResponseJSON    `json:"-"`
}

// overrideCodesResponseJSON contains the JSON metadata for the struct
// [OverrideCodesResponse]
type overrideCodesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideCodesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OverrideCodesResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    overrideCodesResponseErrorJSON `json:"-"`
}

// overrideCodesResponseErrorJSON contains the JSON metadata for the struct
// [OverrideCodesResponseError]
type overrideCodesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideCodesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OverrideCodesResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    overrideCodesResponseMessageJSON `json:"-"`
}

// overrideCodesResponseMessageJSON contains the JSON metadata for the struct
// [OverrideCodesResponseMessage]
type overrideCodesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideCodesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OverrideCodesResponseResult struct {
	DisableForTime OverrideCodesResponseResultDisableForTime `json:"disable_for_time"`
	JSON           overrideCodesResponseResultJSON           `json:"-"`
}

// overrideCodesResponseResultJSON contains the JSON metadata for the struct
// [OverrideCodesResponseResult]
type overrideCodesResponseResultJSON struct {
	DisableForTime apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OverrideCodesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OverrideCodesResponseResultDisableForTime struct {
	// Override code that is valid for 1 hour.
	_1 interface{} `json:"1"`
	// Override code that is valid for 12 hour2.
	_12 interface{} `json:"12"`
	// Override code that is valid for 24 hour.2.
	_24 interface{} `json:"24"`
	// Override code that is valid for 3 hours.
	_3 interface{} `json:"3"`
	// Override code that is valid for 6 hours.
	_6   interface{}                                   `json:"6"`
	JSON overrideCodesResponseResultDisableForTimeJSON `json:"-"`
}

// overrideCodesResponseResultDisableForTimeJSON contains the JSON metadata for the
// struct [OverrideCodesResponseResultDisableForTime]
type overrideCodesResponseResultDisableForTimeJSON struct {
	_1          apijson.Field
	_12         apijson.Field
	_24         apijson.Field
	_3          apijson.Field
	_6          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideCodesResponseResultDisableForTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OverrideCodesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                             `json:"total_count"`
	JSON       overrideCodesResponseResultInfoJSON `json:"-"`
}

// overrideCodesResponseResultInfoJSON contains the JSON metadata for the struct
// [OverrideCodesResponseResultInfo]
type overrideCodesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OverrideCodesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type OverrideCodesResponseSuccess bool

const (
	OverrideCodesResponseSuccessTrue OverrideCodesResponseSuccess = true
)

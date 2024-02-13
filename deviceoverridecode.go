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

// DeviceOverrideCodeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDeviceOverrideCodeService] method
// instead.
type DeviceOverrideCodeService struct {
	Options []option.RequestOption
}

// NewDeviceOverrideCodeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeviceOverrideCodeService(opts ...option.RequestOption) (r *DeviceOverrideCodeService) {
	r = &DeviceOverrideCodeService{}
	r.Options = opts
	return
}

// Fetches a one-time use admin override code for a device. This relies on the
// **Admin Override** setting being enabled in your device configuration.
func (r *DeviceOverrideCodeService) DevicesListAdminOverrideCodeForDevice(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/%s/override_codes", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponse struct {
	DisableForTime DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseDisableForTime `json:"disable_for_time"`
	JSON           deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseJSON           `json:"-"`
}

// deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseJSON contains the
// JSON metadata for the struct
// [DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponse]
type deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseJSON struct {
	DisableForTime apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseDisableForTime struct {
	// Override code that is valid for 1 hour.
	Number1 interface{} `json:"1"`
	// Override code that is valid for 12 hour2.
	Number12 interface{} `json:"12"`
	// Override code that is valid for 24 hour.2.
	Number24 interface{} `json:"24"`
	// Override code that is valid for 3 hours.
	Number3 interface{} `json:"3"`
	// Override code that is valid for 6 hours.
	Number6 interface{}                                                                       `json:"6"`
	JSON    deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseDisableForTimeJSON `json:"-"`
}

// deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseDisableForTimeJSON
// contains the JSON metadata for the struct
// [DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseDisableForTime]
type deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseDisableForTimeJSON struct {
	Number1     apijson.Field
	Number12    apijson.Field
	Number24    apijson.Field
	Number3     apijson.Field
	Number6     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseDisableForTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelope struct {
	Errors   []DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeResultInfo `json:"result_info"`
	JSON       deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeJSON       `json:"-"`
}

// deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelope]
type deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeErrors struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeErrors]
type deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeMessages struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeMessages]
type deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeSuccess bool

const (
	DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeSuccessTrue DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeSuccess = true
)

type DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                               `json:"total_count"`
	JSON       deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeResultInfoJSON `json:"-"`
}

// deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeResultInfo]
type deviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeDevicesListAdminOverrideCodeForDeviceResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

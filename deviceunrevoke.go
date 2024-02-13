// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// DeviceUnrevokeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDeviceUnrevokeService] method
// instead.
type DeviceUnrevokeService struct {
	Options []option.RequestOption
}

// NewDeviceUnrevokeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceUnrevokeService(opts ...option.RequestOption) (r *DeviceUnrevokeService) {
	r = &DeviceUnrevokeService{}
	r.Options = opts
	return
}

// Unrevokes a list of devices.
func (r *DeviceUnrevokeService) DevicesUnrevokeDevices(ctx context.Context, identifier interface{}, body DeviceUnrevokeDevicesUnrevokeDevicesParams, opts ...option.RequestOption) (res *DeviceUnrevokeDevicesUnrevokeDevicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/unrevoke", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [DeviceUnrevokeDevicesUnrevokeDevicesResponseUnknown] or
// [shared.UnionString].
type DeviceUnrevokeDevicesUnrevokeDevicesResponse interface {
	ImplementsDeviceUnrevokeDevicesUnrevokeDevicesResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DeviceUnrevokeDevicesUnrevokeDevicesResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DeviceUnrevokeDevicesUnrevokeDevicesParams struct {
	// A list of device ids to unrevoke.
	Body param.Field[[]string] `json:"body,required"`
}

func (r DeviceUnrevokeDevicesUnrevokeDevicesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelope struct {
	Errors   []DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceUnrevokeDevicesUnrevokeDevicesResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeJSON    `json:"-"`
}

// deviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeJSON contains the JSON
// metadata for the struct [DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelope]
type deviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    deviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeErrors]
type deviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    deviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeMessages]
type deviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeSuccess bool

const (
	DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeSuccessTrue DeviceUnrevokeDevicesUnrevokeDevicesResponseEnvelopeSuccess = true
)

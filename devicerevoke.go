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

// DeviceRevokeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDeviceRevokeService] method
// instead.
type DeviceRevokeService struct {
	Options []option.RequestOption
}

// NewDeviceRevokeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceRevokeService(opts ...option.RequestOption) (r *DeviceRevokeService) {
	r = &DeviceRevokeService{}
	r.Options = opts
	return
}

// Revokes a list of devices.
func (r *DeviceRevokeService) DevicesRevokeDevices(ctx context.Context, identifier interface{}, body DeviceRevokeDevicesRevokeDevicesParams, opts ...option.RequestOption) (res *DeviceRevokeDevicesRevokeDevicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceRevokeDevicesRevokeDevicesResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/revoke", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [DeviceRevokeDevicesRevokeDevicesResponseUnknown] or
// [shared.UnionString].
type DeviceRevokeDevicesRevokeDevicesResponse interface {
	ImplementsDeviceRevokeDevicesRevokeDevicesResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DeviceRevokeDevicesRevokeDevicesResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DeviceRevokeDevicesRevokeDevicesParams struct {
	// A list of device ids to revoke.
	Body param.Field[[]string] `json:"body,required"`
}

func (r DeviceRevokeDevicesRevokeDevicesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DeviceRevokeDevicesRevokeDevicesResponseEnvelope struct {
	Errors   []DeviceRevokeDevicesRevokeDevicesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceRevokeDevicesRevokeDevicesResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceRevokeDevicesRevokeDevicesResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceRevokeDevicesRevokeDevicesResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceRevokeDevicesRevokeDevicesResponseEnvelopeJSON    `json:"-"`
}

// deviceRevokeDevicesRevokeDevicesResponseEnvelopeJSON contains the JSON metadata
// for the struct [DeviceRevokeDevicesRevokeDevicesResponseEnvelope]
type deviceRevokeDevicesRevokeDevicesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRevokeDevicesRevokeDevicesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRevokeDevicesRevokeDevicesResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    deviceRevokeDevicesRevokeDevicesResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceRevokeDevicesRevokeDevicesResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DeviceRevokeDevicesRevokeDevicesResponseEnvelopeErrors]
type deviceRevokeDevicesRevokeDevicesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRevokeDevicesRevokeDevicesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceRevokeDevicesRevokeDevicesResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    deviceRevokeDevicesRevokeDevicesResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceRevokeDevicesRevokeDevicesResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DeviceRevokeDevicesRevokeDevicesResponseEnvelopeMessages]
type deviceRevokeDevicesRevokeDevicesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRevokeDevicesRevokeDevicesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceRevokeDevicesRevokeDevicesResponseEnvelopeSuccess bool

const (
	DeviceRevokeDevicesRevokeDevicesResponseEnvelopeSuccessTrue DeviceRevokeDevicesRevokeDevicesResponseEnvelopeSuccess = true
)

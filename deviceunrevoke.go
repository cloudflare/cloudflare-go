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
func (r *DeviceUnrevokeService) New(ctx context.Context, identifier interface{}, body DeviceUnrevokeNewParams, opts ...option.RequestOption) (res *DeviceUnrevokeNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DeviceUnrevokeNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/unrevoke", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [DeviceUnrevokeNewResponseUnknown] or [shared.UnionString].
type DeviceUnrevokeNewResponse interface {
	ImplementsDeviceUnrevokeNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DeviceUnrevokeNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DeviceUnrevokeNewParams struct {
	// A list of device ids to unrevoke.
	Body param.Field[[]string] `json:"body,required"`
}

func (r DeviceUnrevokeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DeviceUnrevokeNewResponseEnvelope struct {
	Errors   []DeviceUnrevokeNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceUnrevokeNewResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceUnrevokeNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceUnrevokeNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceUnrevokeNewResponseEnvelopeJSON    `json:"-"`
}

// deviceUnrevokeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceUnrevokeNewResponseEnvelope]
type deviceUnrevokeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceUnrevokeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceUnrevokeNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    deviceUnrevokeNewResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceUnrevokeNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceUnrevokeNewResponseEnvelopeErrors]
type deviceUnrevokeNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceUnrevokeNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceUnrevokeNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    deviceUnrevokeNewResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceUnrevokeNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DeviceUnrevokeNewResponseEnvelopeMessages]
type deviceUnrevokeNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceUnrevokeNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type DeviceUnrevokeNewResponseEnvelopeSuccess bool

const (
	DeviceUnrevokeNewResponseEnvelopeSuccessTrue DeviceUnrevokeNewResponseEnvelopeSuccess = true
)

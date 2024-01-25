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

// AccountDeviceUnrevokeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDeviceUnrevokeService]
// method instead.
type AccountDeviceUnrevokeService struct {
	Options []option.RequestOption
}

// NewAccountDeviceUnrevokeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDeviceUnrevokeService(opts ...option.RequestOption) (r *AccountDeviceUnrevokeService) {
	r = &AccountDeviceUnrevokeService{}
	r.Options = opts
	return
}

// Unrevokes a list of devices.
func (r *AccountDeviceUnrevokeService) DevicesUnrevokeDevices(ctx context.Context, identifier interface{}, body AccountDeviceUnrevokeDevicesUnrevokeDevicesParams, opts ...option.RequestOption) (res *AccountDeviceUnrevokeDevicesUnrevokeDevicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/unrevoke", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountDeviceUnrevokeDevicesUnrevokeDevicesResponse struct {
	Errors   []AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseError   `json:"errors"`
	Messages []AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseMessage `json:"messages"`
	Result   AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseResult    `json:"result,nullable"`
	// Whether the API call was successful.
	Success AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseSuccess `json:"success"`
	JSON    accountDeviceUnrevokeDevicesUnrevokeDevicesResponseJSON    `json:"-"`
}

// accountDeviceUnrevokeDevicesUnrevokeDevicesResponseJSON contains the JSON
// metadata for the struct [AccountDeviceUnrevokeDevicesUnrevokeDevicesResponse]
type accountDeviceUnrevokeDevicesUnrevokeDevicesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceUnrevokeDevicesUnrevokeDevicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseError struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    accountDeviceUnrevokeDevicesUnrevokeDevicesResponseErrorJSON `json:"-"`
}

// accountDeviceUnrevokeDevicesUnrevokeDevicesResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseError]
type accountDeviceUnrevokeDevicesUnrevokeDevicesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseMessage struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    accountDeviceUnrevokeDevicesUnrevokeDevicesResponseMessageJSON `json:"-"`
}

// accountDeviceUnrevokeDevicesUnrevokeDevicesResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseMessage]
type accountDeviceUnrevokeDevicesUnrevokeDevicesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseResultUnknown] or
// [shared.UnionString].
type AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseResult interface {
	ImplementsAccountDeviceUnrevokeDevicesUnrevokeDevicesResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful.
type AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseSuccess bool

const (
	AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseSuccessTrue AccountDeviceUnrevokeDevicesUnrevokeDevicesResponseSuccess = true
)

type AccountDeviceUnrevokeDevicesUnrevokeDevicesParams struct {
	// A list of device ids to unrevoke.
	Body param.Field[[]string] `json:"body,required"`
}

func (r AccountDeviceUnrevokeDevicesUnrevokeDevicesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

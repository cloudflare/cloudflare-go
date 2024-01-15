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

// AccountDeviceRevokeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDeviceRevokeService]
// method instead.
type AccountDeviceRevokeService struct {
	Options []option.RequestOption
}

// NewAccountDeviceRevokeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDeviceRevokeService(opts ...option.RequestOption) (r *AccountDeviceRevokeService) {
	r = &AccountDeviceRevokeService{}
	r.Options = opts
	return
}

// Revokes a list of devices.
func (r *AccountDeviceRevokeService) DevicesRevokeDevices(ctx context.Context, identifier interface{}, body AccountDeviceRevokeDevicesRevokeDevicesParams, opts ...option.RequestOption) (res *AccountDeviceRevokeDevicesRevokeDevicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/revoke", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountDeviceRevokeDevicesRevokeDevicesResponse struct {
	Errors   []AccountDeviceRevokeDevicesRevokeDevicesResponseError   `json:"errors"`
	Messages []AccountDeviceRevokeDevicesRevokeDevicesResponseMessage `json:"messages"`
	Result   AccountDeviceRevokeDevicesRevokeDevicesResponseResult    `json:"result,nullable"`
	// Whether the API call was successful.
	Success AccountDeviceRevokeDevicesRevokeDevicesResponseSuccess `json:"success"`
	JSON    accountDeviceRevokeDevicesRevokeDevicesResponseJSON    `json:"-"`
}

// accountDeviceRevokeDevicesRevokeDevicesResponseJSON contains the JSON metadata
// for the struct [AccountDeviceRevokeDevicesRevokeDevicesResponse]
type accountDeviceRevokeDevicesRevokeDevicesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceRevokeDevicesRevokeDevicesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceRevokeDevicesRevokeDevicesResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountDeviceRevokeDevicesRevokeDevicesResponseErrorJSON `json:"-"`
}

// accountDeviceRevokeDevicesRevokeDevicesResponseErrorJSON contains the JSON
// metadata for the struct [AccountDeviceRevokeDevicesRevokeDevicesResponseError]
type accountDeviceRevokeDevicesRevokeDevicesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceRevokeDevicesRevokeDevicesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDeviceRevokeDevicesRevokeDevicesResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountDeviceRevokeDevicesRevokeDevicesResponseMessageJSON `json:"-"`
}

// accountDeviceRevokeDevicesRevokeDevicesResponseMessageJSON contains the JSON
// metadata for the struct [AccountDeviceRevokeDevicesRevokeDevicesResponseMessage]
type accountDeviceRevokeDevicesRevokeDevicesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDeviceRevokeDevicesRevokeDevicesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AccountDeviceRevokeDevicesRevokeDevicesResponseResultUnknown] or
// [shared.UnionString].
type AccountDeviceRevokeDevicesRevokeDevicesResponseResult interface {
	ImplementsAccountDeviceRevokeDevicesRevokeDevicesResponseResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccountDeviceRevokeDevicesRevokeDevicesResponseResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Whether the API call was successful.
type AccountDeviceRevokeDevicesRevokeDevicesResponseSuccess bool

const (
	AccountDeviceRevokeDevicesRevokeDevicesResponseSuccessTrue AccountDeviceRevokeDevicesRevokeDevicesResponseSuccess = true
)

type AccountDeviceRevokeDevicesRevokeDevicesParams struct {
	// A list of device ids to revoke.
	Body param.Field[[]string] `json:"body,required"`
}

func (r AccountDeviceRevokeDevicesRevokeDevicesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

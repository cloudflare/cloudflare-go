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

// Revoke a list of devices.
func (r *AccountDeviceRevokeService) DevicesRevokeDevices(ctx context.Context, identifier interface{}, body AccountDeviceRevokeDevicesRevokeDevicesParams, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/revoke", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountDeviceRevokeDevicesRevokeDevicesParams struct {
	// A list of device ids to revoke.
	Body param.Field[[]string] `json:"body,required"`
}

func (r AccountDeviceRevokeDevicesRevokeDevicesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

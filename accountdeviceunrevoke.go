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

// Unrevoke a list of devices.
func (r *AccountDeviceUnrevokeService) DevicesUnrevokeDevices(ctx context.Context, identifier interface{}, body AccountDeviceUnrevokeDevicesUnrevokeDevicesParams, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/devices/unrevoke", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountDeviceUnrevokeDevicesUnrevokeDevicesParams struct {
	// A list of device ids to unrevoke.
	Body param.Field[[]string] `json:"body,required"`
}

func (r AccountDeviceUnrevokeDevicesUnrevokeDevicesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

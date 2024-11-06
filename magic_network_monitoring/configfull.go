// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// ConfigFullService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigFullService] method instead.
type ConfigFullService struct {
	Options []option.RequestOption
}

// NewConfigFullService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConfigFullService(opts ...option.RequestOption) (r *ConfigFullService) {
	r = &ConfigFullService{}
	r.Options = opts
	return
}

// Lists default sampling, router IPs, warp devices, and rules for account.
func (r *ConfigFullService) Get(ctx context.Context, query ConfigFullGetParams, opts ...option.RequestOption) (res *ConfigFullGetResponse, err error) {
	var env ConfigFullGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/config/full", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConfigFullGetResponse struct {
	// Fallback sampling rate of flow messages being sent in packets per second. This
	// should match the packet sampling rate configured on the router.
	DefaultSampling float64 `json:"default_sampling,required"`
	// The account name.
	Name        string                            `json:"name,required"`
	RouterIPs   []string                          `json:"router_ips,required"`
	WARPDevices []ConfigFullGetResponseWARPDevice `json:"warp_devices,required"`
	JSON        configFullGetResponseJSON         `json:"-"`
}

// configFullGetResponseJSON contains the JSON metadata for the struct
// [ConfigFullGetResponse]
type configFullGetResponseJSON struct {
	DefaultSampling apijson.Field
	Name            apijson.Field
	RouterIPs       apijson.Field
	WARPDevices     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ConfigFullGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configFullGetResponseJSON) RawJSON() string {
	return r.raw
}

// Object representing a warp device with an ID and name.
type ConfigFullGetResponseWARPDevice struct {
	// Unique identifier for the warp device.
	ID string `json:"id,required"`
	// Name of the warp device.
	Name string `json:"name,required"`
	// IPv4 CIDR of the router sourcing flow data associated with this warp device.
	// Only /32 addresses are currently supported.
	RouterIP string                              `json:"router_ip,required"`
	JSON     configFullGetResponseWARPDeviceJSON `json:"-"`
}

// configFullGetResponseWARPDeviceJSON contains the JSON metadata for the struct
// [ConfigFullGetResponseWARPDevice]
type configFullGetResponseWARPDeviceJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	RouterIP    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigFullGetResponseWARPDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configFullGetResponseWARPDeviceJSON) RawJSON() string {
	return r.raw
}

type ConfigFullGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConfigFullGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ConfigFullGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success ConfigFullGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    configFullGetResponseEnvelopeJSON    `json:"-"`
}

// configFullGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConfigFullGetResponseEnvelope]
type configFullGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigFullGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r configFullGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ConfigFullGetResponseEnvelopeSuccess bool

const (
	ConfigFullGetResponseEnvelopeSuccessTrue ConfigFullGetResponseEnvelopeSuccess = true
)

func (r ConfigFullGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ConfigFullGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

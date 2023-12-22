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

// AccountCfdTunnelConfigurationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountCfdTunnelConfigurationService] method instead.
type AccountCfdTunnelConfigurationService struct {
	Options []option.RequestOption
}

// NewAccountCfdTunnelConfigurationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCfdTunnelConfigurationService(opts ...option.RequestOption) (r *AccountCfdTunnelConfigurationService) {
	r = &AccountCfdTunnelConfigurationService{}
	r.Options = opts
	return
}

// Gets the configuration for a remotely-managed tunnel
func (r *AccountCfdTunnelConfigurationService) CloudflareTunnelConfigurationGetConfiguration(ctx context.Context, accountIdentifier string, tunnelID string, opts ...option.RequestOption) (res *ConfigResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Adds or updates the configuration for a remotely-managed tunnel.
func (r *AccountCfdTunnelConfigurationService) CloudflareTunnelConfigurationPutConfiguration(ctx context.Context, accountIdentifier string, tunnelID string, body AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParams, opts ...option.RequestOption) (res *ConfigResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ConfigResponseSingle struct {
	Errors   []ConfigResponseSingleError   `json:"errors"`
	Messages []ConfigResponseSingleMessage `json:"messages"`
	Result   interface{}                   `json:"result"`
	// Whether the API call was successful
	Success ConfigResponseSingleSuccess `json:"success"`
	JSON    configResponseSingleJSON    `json:"-"`
}

// configResponseSingleJSON contains the JSON metadata for the struct
// [ConfigResponseSingle]
type configResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigResponseSingleError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    configResponseSingleErrorJSON `json:"-"`
}

// configResponseSingleErrorJSON contains the JSON metadata for the struct
// [ConfigResponseSingleError]
type configResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigResponseSingleMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    configResponseSingleMessageJSON `json:"-"`
}

// configResponseSingleMessageJSON contains the JSON metadata for the struct
// [ConfigResponseSingleMessage]
type configResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConfigResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ConfigResponseSingleSuccess bool

const (
	ConfigResponseSingleSuccessTrue ConfigResponseSingleSuccess = true
)

type AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParams struct {
	// The tunnel configuration and ingress rules in JSON format. For syntax and
	// parameters, refer to the
	// [configuration file documentation](https://developers.cloudflare.com/cloudflare-one/connections/connect-apps/install-and-setup/tunnel-guide/local/local-management/configuration-file/#file-structure).
	Config param.Field[interface{}] `json:"config"`
}

func (r AccountCfdTunnelConfigurationCloudflareTunnelConfigurationPutConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

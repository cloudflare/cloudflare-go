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

// AccountTunnelConnectionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountTunnelConnectionService] method instead.
type AccountTunnelConnectionService struct {
	Options []option.RequestOption
}

// NewAccountTunnelConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountTunnelConnectionService(opts ...option.RequestOption) (r *AccountTunnelConnectionService) {
	r = &AccountTunnelConnectionService{}
	r.Options = opts
	return
}

// Removes connections that are in a disconnected or pending reconnect state. We
// recommend running this command after shutting down a tunnel.
func (r *AccountTunnelConnectionService) Delete(ctx context.Context, accountIdentifier string, tunnelID string, body AccountTunnelConnectionDeleteParams, opts ...option.RequestOption) (res *AccountTunnelConnectionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/tunnels/%s/connections", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AccountTunnelConnectionDeleteResponse struct {
	Errors   []AccountTunnelConnectionDeleteResponseError   `json:"errors"`
	Messages []AccountTunnelConnectionDeleteResponseMessage `json:"messages"`
	Result   interface{}                                    `json:"result"`
	// Whether the API call was successful
	Success AccountTunnelConnectionDeleteResponseSuccess `json:"success"`
	JSON    accountTunnelConnectionDeleteResponseJSON    `json:"-"`
}

// accountTunnelConnectionDeleteResponseJSON contains the JSON metadata for the
// struct [AccountTunnelConnectionDeleteResponse]
type accountTunnelConnectionDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelConnectionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTunnelConnectionDeleteResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountTunnelConnectionDeleteResponseErrorJSON `json:"-"`
}

// accountTunnelConnectionDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountTunnelConnectionDeleteResponseError]
type accountTunnelConnectionDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelConnectionDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTunnelConnectionDeleteResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountTunnelConnectionDeleteResponseMessageJSON `json:"-"`
}

// accountTunnelConnectionDeleteResponseMessageJSON contains the JSON metadata for
// the struct [AccountTunnelConnectionDeleteResponseMessage]
type accountTunnelConnectionDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTunnelConnectionDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTunnelConnectionDeleteResponseSuccess bool

const (
	AccountTunnelConnectionDeleteResponseSuccessTrue AccountTunnelConnectionDeleteResponseSuccess = true
)

type AccountTunnelConnectionDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountTunnelConnectionDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

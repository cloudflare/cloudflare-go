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

// TunnelConnectionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewTunnelConnectionService] method
// instead.
type TunnelConnectionService struct {
	Options []option.RequestOption
}

// NewTunnelConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTunnelConnectionService(opts ...option.RequestOption) (r *TunnelConnectionService) {
	r = &TunnelConnectionService{}
	r.Options = opts
	return
}

// Removes connections that are in a disconnected or pending reconnect state. We
// recommend running this command after shutting down a tunnel.
func (r *TunnelConnectionService) Delete(ctx context.Context, accountID string, tunnelID string, body TunnelConnectionDeleteParams, opts ...option.RequestOption) (res *TunnelConnectionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/tunnels/%s/connections", accountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type TunnelConnectionDeleteResponse struct {
	Errors   []TunnelConnectionDeleteResponseError   `json:"errors"`
	Messages []TunnelConnectionDeleteResponseMessage `json:"messages"`
	Result   interface{}                             `json:"result"`
	// Whether the API call was successful
	Success TunnelConnectionDeleteResponseSuccess `json:"success"`
	JSON    tunnelConnectionDeleteResponseJSON    `json:"-"`
}

// tunnelConnectionDeleteResponseJSON contains the JSON metadata for the struct
// [TunnelConnectionDeleteResponse]
type tunnelConnectionDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectionDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    tunnelConnectionDeleteResponseErrorJSON `json:"-"`
}

// tunnelConnectionDeleteResponseErrorJSON contains the JSON metadata for the
// struct [TunnelConnectionDeleteResponseError]
type tunnelConnectionDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelConnectionDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    tunnelConnectionDeleteResponseMessageJSON `json:"-"`
}

// tunnelConnectionDeleteResponseMessageJSON contains the JSON metadata for the
// struct [TunnelConnectionDeleteResponseMessage]
type tunnelConnectionDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelConnectionDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelConnectionDeleteResponseSuccess bool

const (
	TunnelConnectionDeleteResponseSuccessTrue TunnelConnectionDeleteResponseSuccess = true
)

type TunnelConnectionDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r TunnelConnectionDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

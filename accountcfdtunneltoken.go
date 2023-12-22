// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountCfdTunnelTokenService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountCfdTunnelTokenService]
// method instead.
type AccountCfdTunnelTokenService struct {
	Options []option.RequestOption
}

// NewAccountCfdTunnelTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountCfdTunnelTokenService(opts ...option.RequestOption) (r *AccountCfdTunnelTokenService) {
	r = &AccountCfdTunnelTokenService{}
	r.Options = opts
	return
}

// Gets the token used to associate cloudflared with a specific tunnel.
func (r *AccountCfdTunnelTokenService) CloudflareTunnelGetACloudflareTunnelToken(ctx context.Context, accountIdentifier string, tunnelID string, opts ...option.RequestOption) (res *TunnelResponseToken, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/token", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TunnelResponseToken struct {
	Errors   []TunnelResponseTokenError   `json:"errors"`
	Messages []TunnelResponseTokenMessage `json:"messages"`
	Result   string                       `json:"result"`
	// Whether the API call was successful
	Success TunnelResponseTokenSuccess `json:"success"`
	JSON    tunnelResponseTokenJSON    `json:"-"`
}

// tunnelResponseTokenJSON contains the JSON metadata for the struct
// [TunnelResponseToken]
type tunnelResponseTokenJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelResponseToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelResponseTokenError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    tunnelResponseTokenErrorJSON `json:"-"`
}

// tunnelResponseTokenErrorJSON contains the JSON metadata for the struct
// [TunnelResponseTokenError]
type tunnelResponseTokenErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelResponseTokenError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TunnelResponseTokenMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    tunnelResponseTokenMessageJSON `json:"-"`
}

// tunnelResponseTokenMessageJSON contains the JSON metadata for the struct
// [TunnelResponseTokenMessage]
type tunnelResponseTokenMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelResponseTokenMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type TunnelResponseTokenSuccess bool

const (
	TunnelResponseTokenSuccessTrue TunnelResponseTokenSuccess = true
)

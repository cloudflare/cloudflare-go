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
func (r *AccountCfdTunnelTokenService) CloudflareTunnelGetACloudflareTunnelToken(ctx context.Context, accountIdentifier string, tunnelID string, opts ...option.RequestOption) (res *AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/token", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponse struct {
	Errors   []AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseError   `json:"errors"`
	Messages []AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseMessage `json:"messages"`
	Result   string                                                                          `json:"result"`
	// Whether the API call was successful
	Success AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseSuccess `json:"success"`
	JSON    accountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseJSON    `json:"-"`
}

// accountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponse]
type accountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseError struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseErrorJSON `json:"-"`
}

// accountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseError]
type accountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseMessage struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    accountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseMessageJSON `json:"-"`
}

// accountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseMessage]
type accountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseSuccess bool

const (
	AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseSuccessTrue AccountCfdTunnelTokenCloudflareTunnelGetACloudflareTunnelTokenResponseSuccess = true
)

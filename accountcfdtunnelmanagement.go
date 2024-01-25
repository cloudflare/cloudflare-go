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

// AccountCfdTunnelManagementService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountCfdTunnelManagementService] method instead.
type AccountCfdTunnelManagementService struct {
	Options []option.RequestOption
}

// NewAccountCfdTunnelManagementService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountCfdTunnelManagementService(opts ...option.RequestOption) (r *AccountCfdTunnelManagementService) {
	r = &AccountCfdTunnelManagementService{}
	r.Options = opts
	return
}

// Gets a management token used to access the management resources (i.e. Streaming
// Logs) of a tunnel.
func (r *AccountCfdTunnelManagementService) List(ctx context.Context, accountIdentifier string, tunnelID string, body AccountCfdTunnelManagementListParams, opts ...option.RequestOption) (res *AccountCfdTunnelManagementListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/management", accountIdentifier, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountCfdTunnelManagementListResponse struct {
	Errors   []AccountCfdTunnelManagementListResponseError   `json:"errors"`
	Messages []AccountCfdTunnelManagementListResponseMessage `json:"messages"`
	Result   string                                          `json:"result"`
	// Whether the API call was successful
	Success AccountCfdTunnelManagementListResponseSuccess `json:"success"`
	JSON    accountCfdTunnelManagementListResponseJSON    `json:"-"`
}

// accountCfdTunnelManagementListResponseJSON contains the JSON metadata for the
// struct [AccountCfdTunnelManagementListResponse]
type accountCfdTunnelManagementListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelManagementListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelManagementListResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountCfdTunnelManagementListResponseErrorJSON `json:"-"`
}

// accountCfdTunnelManagementListResponseErrorJSON contains the JSON metadata for
// the struct [AccountCfdTunnelManagementListResponseError]
type accountCfdTunnelManagementListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelManagementListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountCfdTunnelManagementListResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountCfdTunnelManagementListResponseMessageJSON `json:"-"`
}

// accountCfdTunnelManagementListResponseMessageJSON contains the JSON metadata for
// the struct [AccountCfdTunnelManagementListResponseMessage]
type accountCfdTunnelManagementListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountCfdTunnelManagementListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountCfdTunnelManagementListResponseSuccess bool

const (
	AccountCfdTunnelManagementListResponseSuccessTrue AccountCfdTunnelManagementListResponseSuccess = true
)

type AccountCfdTunnelManagementListParams struct {
	Resources param.Field[[]AccountCfdTunnelManagementListParamsResource] `json:"resources,required"`
}

func (r AccountCfdTunnelManagementListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Management resources the token will have access to.
type AccountCfdTunnelManagementListParamsResource string

const (
	AccountCfdTunnelManagementListParamsResourceLogs AccountCfdTunnelManagementListParamsResource = "logs"
)

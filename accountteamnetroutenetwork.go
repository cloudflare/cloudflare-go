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

// AccountTeamnetRouteNetworkService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountTeamnetRouteNetworkService] method instead.
type AccountTeamnetRouteNetworkService struct {
	Options []option.RequestOption
}

// NewAccountTeamnetRouteNetworkService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountTeamnetRouteNetworkService(opts ...option.RequestOption) (r *AccountTeamnetRouteNetworkService) {
	r = &AccountTeamnetRouteNetworkService{}
	r.Options = opts
	return
}

// Routes a private network through a Cloudflare Tunnel. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format.
func (r *AccountTeamnetRouteNetworkService) Update(ctx context.Context, accountIdentifier string, ipNetworkEncoded string, body AccountTeamnetRouteNetworkUpdateParams, opts ...option.RequestOption) (res *AccountTeamnetRouteNetworkUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", accountIdentifier, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes a private network route from an account. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format.
func (r *AccountTeamnetRouteNetworkService) Delete(ctx context.Context, accountIdentifier string, ipNetworkEncoded string, body AccountTeamnetRouteNetworkDeleteParams, opts ...option.RequestOption) (res *AccountTeamnetRouteNetworkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", accountIdentifier, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type AccountTeamnetRouteNetworkUpdateResponse struct {
	Errors   []AccountTeamnetRouteNetworkUpdateResponseError   `json:"errors"`
	Messages []AccountTeamnetRouteNetworkUpdateResponseMessage `json:"messages"`
	Result   interface{}                                       `json:"result"`
	// Whether the API call was successful
	Success AccountTeamnetRouteNetworkUpdateResponseSuccess `json:"success"`
	JSON    accountTeamnetRouteNetworkUpdateResponseJSON    `json:"-"`
}

// accountTeamnetRouteNetworkUpdateResponseJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteNetworkUpdateResponse]
type accountTeamnetRouteNetworkUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNetworkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteNetworkUpdateResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountTeamnetRouteNetworkUpdateResponseErrorJSON `json:"-"`
}

// accountTeamnetRouteNetworkUpdateResponseErrorJSON contains the JSON metadata for
// the struct [AccountTeamnetRouteNetworkUpdateResponseError]
type accountTeamnetRouteNetworkUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNetworkUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteNetworkUpdateResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountTeamnetRouteNetworkUpdateResponseMessageJSON `json:"-"`
}

// accountTeamnetRouteNetworkUpdateResponseMessageJSON contains the JSON metadata
// for the struct [AccountTeamnetRouteNetworkUpdateResponseMessage]
type accountTeamnetRouteNetworkUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNetworkUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTeamnetRouteNetworkUpdateResponseSuccess bool

const (
	AccountTeamnetRouteNetworkUpdateResponseSuccessTrue AccountTeamnetRouteNetworkUpdateResponseSuccess = true
)

type AccountTeamnetRouteNetworkDeleteResponse struct {
	Errors   []AccountTeamnetRouteNetworkDeleteResponseError   `json:"errors"`
	Messages []AccountTeamnetRouteNetworkDeleteResponseMessage `json:"messages"`
	Result   interface{}                                       `json:"result"`
	// Whether the API call was successful
	Success AccountTeamnetRouteNetworkDeleteResponseSuccess `json:"success"`
	JSON    accountTeamnetRouteNetworkDeleteResponseJSON    `json:"-"`
}

// accountTeamnetRouteNetworkDeleteResponseJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteNetworkDeleteResponse]
type accountTeamnetRouteNetworkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNetworkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteNetworkDeleteResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountTeamnetRouteNetworkDeleteResponseErrorJSON `json:"-"`
}

// accountTeamnetRouteNetworkDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountTeamnetRouteNetworkDeleteResponseError]
type accountTeamnetRouteNetworkDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNetworkDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteNetworkDeleteResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountTeamnetRouteNetworkDeleteResponseMessageJSON `json:"-"`
}

// accountTeamnetRouteNetworkDeleteResponseMessageJSON contains the JSON metadata
// for the struct [AccountTeamnetRouteNetworkDeleteResponseMessage]
type accountTeamnetRouteNetworkDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteNetworkDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTeamnetRouteNetworkDeleteResponseSuccess bool

const (
	AccountTeamnetRouteNetworkDeleteResponseSuccessTrue AccountTeamnetRouteNetworkDeleteResponseSuccess = true
)

type AccountTeamnetRouteNetworkUpdateParams struct {
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `json:"virtual_network_id"`
}

func (r AccountTeamnetRouteNetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountTeamnetRouteNetworkDeleteParams struct {
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `json:"virtual_network_id"`
}

func (r AccountTeamnetRouteNetworkDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountTeamnetRouteIPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountTeamnetRouteIPService]
// method instead.
type AccountTeamnetRouteIPService struct {
	Options []option.RequestOption
}

// NewAccountTeamnetRouteIPService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountTeamnetRouteIPService(opts ...option.RequestOption) (r *AccountTeamnetRouteIPService) {
	r = &AccountTeamnetRouteIPService{}
	r.Options = opts
	return
}

// Fetches routes that contain the given IP address.
func (r *AccountTeamnetRouteIPService) Get(ctx context.Context, accountIdentifier string, ip string, query AccountTeamnetRouteIPGetParams, opts ...option.RequestOption) (res *AccountTeamnetRouteIPGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/teamnet/routes/ip/%s", accountIdentifier, ip)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountTeamnetRouteIPGetResponse struct {
	Errors   []AccountTeamnetRouteIPGetResponseError   `json:"errors"`
	Messages []AccountTeamnetRouteIPGetResponseMessage `json:"messages"`
	Result   interface{}                               `json:"result"`
	// Whether the API call was successful
	Success AccountTeamnetRouteIPGetResponseSuccess `json:"success"`
	JSON    accountTeamnetRouteIPGetResponseJSON    `json:"-"`
}

// accountTeamnetRouteIPGetResponseJSON contains the JSON metadata for the struct
// [AccountTeamnetRouteIPGetResponse]
type accountTeamnetRouteIPGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteIPGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteIPGetResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountTeamnetRouteIPGetResponseErrorJSON `json:"-"`
}

// accountTeamnetRouteIPGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteIPGetResponseError]
type accountTeamnetRouteIPGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteIPGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountTeamnetRouteIPGetResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountTeamnetRouteIPGetResponseMessageJSON `json:"-"`
}

// accountTeamnetRouteIPGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountTeamnetRouteIPGetResponseMessage]
type accountTeamnetRouteIPGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountTeamnetRouteIPGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountTeamnetRouteIPGetResponseSuccess bool

const (
	AccountTeamnetRouteIPGetResponseSuccessTrue AccountTeamnetRouteIPGetResponseSuccess = true
)

type AccountTeamnetRouteIPGetParams struct {
	// UUID of the Tunnel Virtual Network this route belongs to. If no virtual networks
	// are configured, the route is assigned to the default virtual network of the
	// account.
	VirtualNetworkID param.Field[interface{}] `query:"virtual_network_id"`
}

// URLQuery serializes [AccountTeamnetRouteIPGetParams]'s query parameters as
// `url.Values`.
func (r AccountTeamnetRouteIPGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

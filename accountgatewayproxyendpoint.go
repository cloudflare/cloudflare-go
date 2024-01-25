// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountGatewayProxyEndpointService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountGatewayProxyEndpointService] method instead.
type AccountGatewayProxyEndpointService struct {
	Options []option.RequestOption
}

// NewAccountGatewayProxyEndpointService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountGatewayProxyEndpointService(opts ...option.RequestOption) (r *AccountGatewayProxyEndpointService) {
	r = &AccountGatewayProxyEndpointService{}
	r.Options = opts
	return
}

// Fetches all Zero Trust Gateway proxy endpoints for an account.
func (r *AccountGatewayProxyEndpointService) Get(ctx context.Context, identifier interface{}, uuid interface{}, opts ...option.RequestOption) (res *AccountGatewayProxyEndpointGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Zero Trust Gateway proxy endpoint.
func (r *AccountGatewayProxyEndpointService) Update(ctx context.Context, identifier interface{}, uuid interface{}, body AccountGatewayProxyEndpointUpdateParams, opts ...option.RequestOption) (res *AccountGatewayProxyEndpointUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes a configured Zero Trust Gateway proxy endpoint.
func (r *AccountGatewayProxyEndpointService) Delete(ctx context.Context, identifier interface{}, uuid interface{}, opts ...option.RequestOption) (res *AccountGatewayProxyEndpointDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new Zero Trust Gateway proxy endpoint.
func (r *AccountGatewayProxyEndpointService) ZeroTrustGatewayProxyEndpointsNewProxyEndpoint(ctx context.Context, identifier interface{}, body AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointParams, opts ...option.RequestOption) (res *AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches a single Zero Trust Gateway proxy endpoint.
func (r *AccountGatewayProxyEndpointService) ZeroTrustGatewayProxyEndpointsListProxyEndpoints(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountGatewayProxyEndpointGetResponse struct {
	Errors   []AccountGatewayProxyEndpointGetResponseError   `json:"errors"`
	Messages []AccountGatewayProxyEndpointGetResponseMessage `json:"messages"`
	Result   AccountGatewayProxyEndpointGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayProxyEndpointGetResponseSuccess `json:"success"`
	JSON    accountGatewayProxyEndpointGetResponseJSON    `json:"-"`
}

// accountGatewayProxyEndpointGetResponseJSON contains the JSON metadata for the
// struct [AccountGatewayProxyEndpointGetResponse]
type accountGatewayProxyEndpointGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointGetResponseError struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountGatewayProxyEndpointGetResponseErrorJSON `json:"-"`
}

// accountGatewayProxyEndpointGetResponseErrorJSON contains the JSON metadata for
// the struct [AccountGatewayProxyEndpointGetResponseError]
type accountGatewayProxyEndpointGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointGetResponseMessage struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountGatewayProxyEndpointGetResponseMessageJSON `json:"-"`
}

// accountGatewayProxyEndpointGetResponseMessageJSON contains the JSON metadata for
// the struct [AccountGatewayProxyEndpointGetResponseMessage]
type accountGatewayProxyEndpointGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointGetResponseResult struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                                           `json:"subdomain"`
	UpdatedAt time.Time                                        `json:"updated_at" format:"date-time"`
	JSON      accountGatewayProxyEndpointGetResponseResultJSON `json:"-"`
}

// accountGatewayProxyEndpointGetResponseResultJSON contains the JSON metadata for
// the struct [AccountGatewayProxyEndpointGetResponseResult]
type accountGatewayProxyEndpointGetResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayProxyEndpointGetResponseSuccess bool

const (
	AccountGatewayProxyEndpointGetResponseSuccessTrue AccountGatewayProxyEndpointGetResponseSuccess = true
)

type AccountGatewayProxyEndpointUpdateResponse struct {
	Errors   []AccountGatewayProxyEndpointUpdateResponseError   `json:"errors"`
	Messages []AccountGatewayProxyEndpointUpdateResponseMessage `json:"messages"`
	Result   AccountGatewayProxyEndpointUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayProxyEndpointUpdateResponseSuccess `json:"success"`
	JSON    accountGatewayProxyEndpointUpdateResponseJSON    `json:"-"`
}

// accountGatewayProxyEndpointUpdateResponseJSON contains the JSON metadata for the
// struct [AccountGatewayProxyEndpointUpdateResponse]
type accountGatewayProxyEndpointUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointUpdateResponseError struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accountGatewayProxyEndpointUpdateResponseErrorJSON `json:"-"`
}

// accountGatewayProxyEndpointUpdateResponseErrorJSON contains the JSON metadata
// for the struct [AccountGatewayProxyEndpointUpdateResponseError]
type accountGatewayProxyEndpointUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointUpdateResponseMessage struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountGatewayProxyEndpointUpdateResponseMessageJSON `json:"-"`
}

// accountGatewayProxyEndpointUpdateResponseMessageJSON contains the JSON metadata
// for the struct [AccountGatewayProxyEndpointUpdateResponseMessage]
type accountGatewayProxyEndpointUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointUpdateResponseResult struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                                              `json:"subdomain"`
	UpdatedAt time.Time                                           `json:"updated_at" format:"date-time"`
	JSON      accountGatewayProxyEndpointUpdateResponseResultJSON `json:"-"`
}

// accountGatewayProxyEndpointUpdateResponseResultJSON contains the JSON metadata
// for the struct [AccountGatewayProxyEndpointUpdateResponseResult]
type accountGatewayProxyEndpointUpdateResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayProxyEndpointUpdateResponseSuccess bool

const (
	AccountGatewayProxyEndpointUpdateResponseSuccessTrue AccountGatewayProxyEndpointUpdateResponseSuccess = true
)

type AccountGatewayProxyEndpointDeleteResponse struct {
	Errors   []AccountGatewayProxyEndpointDeleteResponseError   `json:"errors"`
	Messages []AccountGatewayProxyEndpointDeleteResponseMessage `json:"messages"`
	Result   interface{}                                        `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayProxyEndpointDeleteResponseSuccess `json:"success"`
	JSON    accountGatewayProxyEndpointDeleteResponseJSON    `json:"-"`
}

// accountGatewayProxyEndpointDeleteResponseJSON contains the JSON metadata for the
// struct [AccountGatewayProxyEndpointDeleteResponse]
type accountGatewayProxyEndpointDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointDeleteResponseError struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accountGatewayProxyEndpointDeleteResponseErrorJSON `json:"-"`
}

// accountGatewayProxyEndpointDeleteResponseErrorJSON contains the JSON metadata
// for the struct [AccountGatewayProxyEndpointDeleteResponseError]
type accountGatewayProxyEndpointDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointDeleteResponseMessage struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    accountGatewayProxyEndpointDeleteResponseMessageJSON `json:"-"`
}

// accountGatewayProxyEndpointDeleteResponseMessageJSON contains the JSON metadata
// for the struct [AccountGatewayProxyEndpointDeleteResponseMessage]
type accountGatewayProxyEndpointDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayProxyEndpointDeleteResponseSuccess bool

const (
	AccountGatewayProxyEndpointDeleteResponseSuccessTrue AccountGatewayProxyEndpointDeleteResponseSuccess = true
)

type AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponse struct {
	Errors   []AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseError   `json:"errors"`
	Messages []AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseMessage `json:"messages"`
	Result   AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseSuccess `json:"success"`
	JSON    accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseJSON    `json:"-"`
}

// accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseJSON
// contains the JSON metadata for the struct
// [AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponse]
type accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseError struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseErrorJSON `json:"-"`
}

// accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseError]
type accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseMessage struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseMessageJSON `json:"-"`
}

// accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseMessage]
type accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseResult struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                                                                                      `json:"subdomain"`
	UpdatedAt time.Time                                                                                   `json:"updated_at" format:"date-time"`
	JSON      accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseResultJSON `json:"-"`
}

// accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseResultJSON
// contains the JSON metadata for the struct
// [AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseResult]
type accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseSuccess bool

const (
	AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseSuccessTrue AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseSuccess = true
)

type AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponse struct {
	Errors     []AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseError    `json:"errors"`
	Messages   []AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseMessage  `json:"messages"`
	Result     []AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseResult   `json:"result"`
	ResultInfo AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseSuccess `json:"success"`
	JSON    accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseJSON    `json:"-"`
}

// accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseJSON
// contains the JSON metadata for the struct
// [AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponse]
type accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseError struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseErrorJSON `json:"-"`
}

// accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseError]
type accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseMessage struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseMessageJSON `json:"-"`
}

// accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseMessage]
type accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseResult struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                                                                                        `json:"subdomain"`
	UpdatedAt time.Time                                                                                     `json:"updated_at" format:"date-time"`
	JSON      accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseResultJSON `json:"-"`
}

// accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseResult]
type accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                           `json:"total_count"`
	JSON       accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseResultInfoJSON `json:"-"`
}

// accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseResultInfo]
type accountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseSuccess bool

const (
	AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseSuccessTrue AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseSuccess = true
)

type AccountGatewayProxyEndpointUpdateParams struct {
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]string] `json:"ips"`
	// The name of the proxy endpoint.
	Name param.Field[string] `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain param.Field[string] `json:"subdomain"`
}

func (r AccountGatewayProxyEndpointUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointParams struct {
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]string] `json:"ips,required"`
	// The name of the proxy endpoint.
	Name param.Field[string] `json:"name,required"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain param.Field[string] `json:"subdomain"`
}

func (r AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

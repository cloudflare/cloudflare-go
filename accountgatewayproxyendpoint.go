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

// List Zero Trust Gateway Proxy Endpoints for an account.
func (r *AccountGatewayProxyEndpointService) Get(ctx context.Context, identifier interface{}, uuid interface{}, opts ...option.RequestOption) (res *ProxyEndpointsSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a new Zero Trust Gateway Proxy Endpoint.
func (r *AccountGatewayProxyEndpointService) Update(ctx context.Context, identifier interface{}, uuid interface{}, body AccountGatewayProxyEndpointUpdateParams, opts ...option.RequestOption) (res *ProxyEndpointsSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete a Zero Trust Gateway Proxy Endpoint.
func (r *AccountGatewayProxyEndpointService) Delete(ctx context.Context, identifier interface{}, uuid interface{}, opts ...option.RequestOption) (res *EmptyResponseArJnNcMr, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new Zero Trust Gateway Proxy Endpoint.
func (r *AccountGatewayProxyEndpointService) ZeroTrustGatewayProxyEndpointsNewProxyEndpoint(ctx context.Context, identifier interface{}, body AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointParams, opts ...option.RequestOption) (res *ProxyEndpointsSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetch a single Zero Trust Gateway Proxy Endpoint.
func (r *AccountGatewayProxyEndpointService) ZeroTrustGatewayProxyEndpointsListProxyEndpoints(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *ProxyEndpointsResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ProxyEndpointsResponseCollection struct {
	Errors     []ProxyEndpointsResponseCollectionError    `json:"errors"`
	Messages   []ProxyEndpointsResponseCollectionMessage  `json:"messages"`
	Result     []ProxyEndpointsResponseCollectionResult   `json:"result"`
	ResultInfo ProxyEndpointsResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ProxyEndpointsResponseCollectionSuccess `json:"success"`
	JSON    proxyEndpointsResponseCollectionJSON    `json:"-"`
}

// proxyEndpointsResponseCollectionJSON contains the JSON metadata for the struct
// [ProxyEndpointsResponseCollection]
type proxyEndpointsResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyEndpointsResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProxyEndpointsResponseCollectionError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    proxyEndpointsResponseCollectionErrorJSON `json:"-"`
}

// proxyEndpointsResponseCollectionErrorJSON contains the JSON metadata for the
// struct [ProxyEndpointsResponseCollectionError]
type proxyEndpointsResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyEndpointsResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProxyEndpointsResponseCollectionMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    proxyEndpointsResponseCollectionMessageJSON `json:"-"`
}

// proxyEndpointsResponseCollectionMessageJSON contains the JSON metadata for the
// struct [ProxyEndpointsResponseCollectionMessage]
type proxyEndpointsResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyEndpointsResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProxyEndpointsResponseCollectionResult struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the Proxy Endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                                     `json:"subdomain"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      proxyEndpointsResponseCollectionResultJSON `json:"-"`
}

// proxyEndpointsResponseCollectionResultJSON contains the JSON metadata for the
// struct [ProxyEndpointsResponseCollectionResult]
type proxyEndpointsResponseCollectionResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyEndpointsResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProxyEndpointsResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       proxyEndpointsResponseCollectionResultInfoJSON `json:"-"`
}

// proxyEndpointsResponseCollectionResultInfoJSON contains the JSON metadata for
// the struct [ProxyEndpointsResponseCollectionResultInfo]
type proxyEndpointsResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyEndpointsResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ProxyEndpointsResponseCollectionSuccess bool

const (
	ProxyEndpointsResponseCollectionSuccessTrue ProxyEndpointsResponseCollectionSuccess = true
)

type ProxyEndpointsSingleResponse struct {
	Errors   []ProxyEndpointsSingleResponseError   `json:"errors"`
	Messages []ProxyEndpointsSingleResponseMessage `json:"messages"`
	Result   ProxyEndpointsSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ProxyEndpointsSingleResponseSuccess `json:"success"`
	JSON    proxyEndpointsSingleResponseJSON    `json:"-"`
}

// proxyEndpointsSingleResponseJSON contains the JSON metadata for the struct
// [ProxyEndpointsSingleResponse]
type proxyEndpointsSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyEndpointsSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProxyEndpointsSingleResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    proxyEndpointsSingleResponseErrorJSON `json:"-"`
}

// proxyEndpointsSingleResponseErrorJSON contains the JSON metadata for the struct
// [ProxyEndpointsSingleResponseError]
type proxyEndpointsSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyEndpointsSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProxyEndpointsSingleResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    proxyEndpointsSingleResponseMessageJSON `json:"-"`
}

// proxyEndpointsSingleResponseMessageJSON contains the JSON metadata for the
// struct [ProxyEndpointsSingleResponseMessage]
type proxyEndpointsSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyEndpointsSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProxyEndpointsSingleResponseResult struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the Proxy Endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                                 `json:"subdomain"`
	UpdatedAt time.Time                              `json:"updated_at" format:"date-time"`
	JSON      proxyEndpointsSingleResponseResultJSON `json:"-"`
}

// proxyEndpointsSingleResponseResultJSON contains the JSON metadata for the struct
// [ProxyEndpointsSingleResponseResult]
type proxyEndpointsSingleResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyEndpointsSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ProxyEndpointsSingleResponseSuccess bool

const (
	ProxyEndpointsSingleResponseSuccessTrue ProxyEndpointsSingleResponseSuccess = true
)

type AccountGatewayProxyEndpointUpdateParams struct {
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]string] `json:"ips"`
	// The name of the Proxy Endpoint.
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
	// The name of the Proxy Endpoint.
	Name param.Field[string] `json:"name,required"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain param.Field[string] `json:"subdomain"`
}

func (r AccountGatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

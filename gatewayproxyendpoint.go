// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// GatewayProxyEndpointService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayProxyEndpointService]
// method instead.
type GatewayProxyEndpointService struct {
	Options []option.RequestOption
}

// NewGatewayProxyEndpointService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGatewayProxyEndpointService(opts ...option.RequestOption) (r *GatewayProxyEndpointService) {
	r = &GatewayProxyEndpointService{}
	r.Options = opts
	return
}

// Updates a configured Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) Update(ctx context.Context, accountID interface{}, proxyEndpointID interface{}, body GatewayProxyEndpointUpdateParams, opts ...option.RequestOption) (res *GatewayProxyEndpointUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", accountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all Zero Trust Gateway proxy endpoints for an account.
func (r *GatewayProxyEndpointService) List(ctx context.Context, accountID interface{}, proxyEndpointID interface{}, opts ...option.RequestOption) (res *GatewayProxyEndpointListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", accountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a configured Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) Delete(ctx context.Context, accountID interface{}, proxyEndpointID interface{}, opts ...option.RequestOption) (res *GatewayProxyEndpointDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", accountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all Zero Trust Gateway proxy endpoints for an account.
func (r *GatewayProxyEndpointService) Get(ctx context.Context, accountID interface{}, proxyEndpointID interface{}, opts ...option.RequestOption) (res *GatewayProxyEndpointGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", accountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a new Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) ZeroTrustGatewayProxyEndpointsNewProxyEndpoint(ctx context.Context, accountID interface{}, body GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointParams, opts ...option.RequestOption) (res *GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) ZeroTrustGatewayProxyEndpointsListProxyEndpoints(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *[]GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayProxyEndpointUpdateResponse struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                                 `json:"subdomain"`
	UpdatedAt time.Time                              `json:"updated_at" format:"date-time"`
	JSON      gatewayProxyEndpointUpdateResponseJSON `json:"-"`
}

// gatewayProxyEndpointUpdateResponseJSON contains the JSON metadata for the struct
// [GatewayProxyEndpointUpdateResponse]
type gatewayProxyEndpointUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointListResponse struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                               `json:"subdomain"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      gatewayProxyEndpointListResponseJSON `json:"-"`
}

// gatewayProxyEndpointListResponseJSON contains the JSON metadata for the struct
// [GatewayProxyEndpointListResponse]
type gatewayProxyEndpointListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [GatewayProxyEndpointDeleteResponseUnknown] or
// [shared.UnionString].
type GatewayProxyEndpointDeleteResponse interface {
	ImplementsGatewayProxyEndpointDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GatewayProxyEndpointDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type GatewayProxyEndpointGetResponse struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                              `json:"subdomain"`
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      gatewayProxyEndpointGetResponseJSON `json:"-"`
}

// gatewayProxyEndpointGetResponseJSON contains the JSON metadata for the struct
// [GatewayProxyEndpointGetResponse]
type gatewayProxyEndpointGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponse struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                                                                         `json:"subdomain"`
	UpdatedAt time.Time                                                                      `json:"updated_at" format:"date-time"`
	JSON      gatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseJSON `json:"-"`
}

// gatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseJSON
// contains the JSON metadata for the struct
// [GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponse]
type gatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponse struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                                                                           `json:"subdomain"`
	UpdatedAt time.Time                                                                        `json:"updated_at" format:"date-time"`
	JSON      gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseJSON `json:"-"`
}

// gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseJSON
// contains the JSON metadata for the struct
// [GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponse]
type gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointUpdateParams struct {
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]string] `json:"ips"`
	// The name of the proxy endpoint.
	Name param.Field[string] `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain param.Field[string] `json:"subdomain"`
}

func (r GatewayProxyEndpointUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayProxyEndpointUpdateResponseEnvelope struct {
	Errors   []GatewayProxyEndpointUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayProxyEndpointUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayProxyEndpointUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayProxyEndpointUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayProxyEndpointUpdateResponseEnvelope]
type gatewayProxyEndpointUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    gatewayProxyEndpointUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayProxyEndpointUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [GatewayProxyEndpointUpdateResponseEnvelopeErrors]
type gatewayProxyEndpointUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    gatewayProxyEndpointUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayProxyEndpointUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [GatewayProxyEndpointUpdateResponseEnvelopeMessages]
type gatewayProxyEndpointUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayProxyEndpointUpdateResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointUpdateResponseEnvelopeSuccessTrue GatewayProxyEndpointUpdateResponseEnvelopeSuccess = true
)

type GatewayProxyEndpointListResponseEnvelope struct {
	Errors   []GatewayProxyEndpointListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayProxyEndpointListResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayProxyEndpointListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointListResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayProxyEndpointListResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointListResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayProxyEndpointListResponseEnvelope]
type gatewayProxyEndpointListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointListResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    gatewayProxyEndpointListResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayProxyEndpointListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [GatewayProxyEndpointListResponseEnvelopeErrors]
type gatewayProxyEndpointListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointListResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    gatewayProxyEndpointListResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayProxyEndpointListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [GatewayProxyEndpointListResponseEnvelopeMessages]
type gatewayProxyEndpointListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayProxyEndpointListResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointListResponseEnvelopeSuccessTrue GatewayProxyEndpointListResponseEnvelopeSuccess = true
)

type GatewayProxyEndpointDeleteResponseEnvelope struct {
	Errors   []GatewayProxyEndpointDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayProxyEndpointDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayProxyEndpointDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayProxyEndpointDeleteResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayProxyEndpointDeleteResponseEnvelope]
type gatewayProxyEndpointDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointDeleteResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    gatewayProxyEndpointDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayProxyEndpointDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [GatewayProxyEndpointDeleteResponseEnvelopeErrors]
type gatewayProxyEndpointDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointDeleteResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    gatewayProxyEndpointDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayProxyEndpointDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [GatewayProxyEndpointDeleteResponseEnvelopeMessages]
type gatewayProxyEndpointDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayProxyEndpointDeleteResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointDeleteResponseEnvelopeSuccessTrue GatewayProxyEndpointDeleteResponseEnvelopeSuccess = true
)

type GatewayProxyEndpointGetResponseEnvelope struct {
	Errors   []GatewayProxyEndpointGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayProxyEndpointGetResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayProxyEndpointGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayProxyEndpointGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayProxyEndpointGetResponseEnvelope]
type gatewayProxyEndpointGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    gatewayProxyEndpointGetResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayProxyEndpointGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [GatewayProxyEndpointGetResponseEnvelopeErrors]
type gatewayProxyEndpointGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    gatewayProxyEndpointGetResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayProxyEndpointGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [GatewayProxyEndpointGetResponseEnvelopeMessages]
type gatewayProxyEndpointGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayProxyEndpointGetResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointGetResponseEnvelopeSuccessTrue GatewayProxyEndpointGetResponseEnvelopeSuccess = true
)

type GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointParams struct {
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]string] `json:"ips,required"`
	// The name of the proxy endpoint.
	Name param.Field[string] `json:"name,required"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain param.Field[string] `json:"subdomain"`
}

func (r GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelope struct {
	Errors   []GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelope]
type gatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeErrors struct {
	Code    int64                                                                                        `json:"code,required"`
	Message string                                                                                       `json:"message,required"`
	JSON    gatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeErrors]
type gatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeMessages struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    gatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeMessages]
type gatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeSuccessTrue GatewayProxyEndpointZeroTrustGatewayProxyEndpointsNewProxyEndpointResponseEnvelopeSuccess = true
)

type GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelope struct {
	Errors   []GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeMessages `json:"messages,required"`
	Result   []GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeJSON       `json:"-"`
}

// gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelope]
type gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeErrors struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeErrors]
type gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeMessages struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeMessages]
type gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeSuccessTrue GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeSuccess = true
)

type GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                            `json:"total_count"`
	JSON       gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeResultInfoJSON `json:"-"`
}

// gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeResultInfo]
type gatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointZeroTrustGatewayProxyEndpointsListProxyEndpointsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

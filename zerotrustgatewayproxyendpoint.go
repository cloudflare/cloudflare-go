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

// ZeroTrustGatewayProxyEndpointService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustGatewayProxyEndpointService] method instead.
type ZeroTrustGatewayProxyEndpointService struct {
	Options []option.RequestOption
}

// NewZeroTrustGatewayProxyEndpointService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustGatewayProxyEndpointService(opts ...option.RequestOption) (r *ZeroTrustGatewayProxyEndpointService) {
	r = &ZeroTrustGatewayProxyEndpointService{}
	r.Options = opts
	return
}

// Creates a new Zero Trust Gateway proxy endpoint.
func (r *ZeroTrustGatewayProxyEndpointService) New(ctx context.Context, params ZeroTrustGatewayProxyEndpointNewParams, opts ...option.RequestOption) (res *ZeroTrustGatewayProxyEndpointNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayProxyEndpointNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust Gateway proxy endpoint.
func (r *ZeroTrustGatewayProxyEndpointService) List(ctx context.Context, query ZeroTrustGatewayProxyEndpointListParams, opts ...option.RequestOption) (res *[]ZeroTrustGatewayProxyEndpointListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayProxyEndpointListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a configured Zero Trust Gateway proxy endpoint.
func (r *ZeroTrustGatewayProxyEndpointService) Delete(ctx context.Context, proxyEndpointID interface{}, body ZeroTrustGatewayProxyEndpointDeleteParams, opts ...option.RequestOption) (res *ZeroTrustGatewayProxyEndpointDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayProxyEndpointDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", body.AccountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust Gateway proxy endpoint.
func (r *ZeroTrustGatewayProxyEndpointService) Edit(ctx context.Context, proxyEndpointID interface{}, params ZeroTrustGatewayProxyEndpointEditParams, opts ...option.RequestOption) (res *ZeroTrustGatewayProxyEndpointEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayProxyEndpointEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", params.AccountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all Zero Trust Gateway proxy endpoints for an account.
func (r *ZeroTrustGatewayProxyEndpointService) Get(ctx context.Context, proxyEndpointID interface{}, query ZeroTrustGatewayProxyEndpointGetParams, opts ...option.RequestOption) (res *ZeroTrustGatewayProxyEndpointGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayProxyEndpointGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", query.AccountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustGatewayProxyEndpointNewResponse struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                                       `json:"subdomain"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayProxyEndpointNewResponseJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointNewResponseJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayProxyEndpointNewResponse]
type zeroTrustGatewayProxyEndpointNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayProxyEndpointListResponse struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                                        `json:"subdomain"`
	UpdatedAt time.Time                                     `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayProxyEndpointListResponseJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointListResponseJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayProxyEndpointListResponse]
type zeroTrustGatewayProxyEndpointListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [ZeroTrustGatewayProxyEndpointDeleteResponseUnknown] or
// [shared.UnionString].
type ZeroTrustGatewayProxyEndpointDeleteResponse interface {
	ImplementsZeroTrustGatewayProxyEndpointDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustGatewayProxyEndpointDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustGatewayProxyEndpointEditResponse struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                                        `json:"subdomain"`
	UpdatedAt time.Time                                     `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayProxyEndpointEditResponseJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointEditResponseJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayProxyEndpointEditResponse]
type zeroTrustGatewayProxyEndpointEditResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayProxyEndpointGetResponse struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                                       `json:"subdomain"`
	UpdatedAt time.Time                                    `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayProxyEndpointGetResponseJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointGetResponseJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayProxyEndpointGetResponse]
type zeroTrustGatewayProxyEndpointGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayProxyEndpointNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]string] `json:"ips,required"`
	// The name of the proxy endpoint.
	Name param.Field[string] `json:"name,required"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain param.Field[string] `json:"subdomain"`
}

func (r ZeroTrustGatewayProxyEndpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayProxyEndpointNewResponseEnvelope struct {
	Errors   []ZeroTrustGatewayProxyEndpointNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayProxyEndpointNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayProxyEndpointNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayProxyEndpointNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayProxyEndpointNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayProxyEndpointNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayProxyEndpointNewResponseEnvelope]
type zeroTrustGatewayProxyEndpointNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayProxyEndpointNewResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustGatewayProxyEndpointNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayProxyEndpointNewResponseEnvelopeErrors]
type zeroTrustGatewayProxyEndpointNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayProxyEndpointNewResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustGatewayProxyEndpointNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayProxyEndpointNewResponseEnvelopeMessages]
type zeroTrustGatewayProxyEndpointNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayProxyEndpointNewResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayProxyEndpointNewResponseEnvelopeSuccessTrue ZeroTrustGatewayProxyEndpointNewResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayProxyEndpointListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayProxyEndpointListResponseEnvelope struct {
	Errors   []ZeroTrustGatewayProxyEndpointListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayProxyEndpointListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustGatewayProxyEndpointListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustGatewayProxyEndpointListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustGatewayProxyEndpointListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustGatewayProxyEndpointListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustGatewayProxyEndpointListResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayProxyEndpointListResponseEnvelope]
type zeroTrustGatewayProxyEndpointListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayProxyEndpointListResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustGatewayProxyEndpointListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayProxyEndpointListResponseEnvelopeErrors]
type zeroTrustGatewayProxyEndpointListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayProxyEndpointListResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustGatewayProxyEndpointListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayProxyEndpointListResponseEnvelopeMessages]
type zeroTrustGatewayProxyEndpointListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayProxyEndpointListResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayProxyEndpointListResponseEnvelopeSuccessTrue ZeroTrustGatewayProxyEndpointListResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayProxyEndpointListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       zeroTrustGatewayProxyEndpointListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointListResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayProxyEndpointListResponseEnvelopeResultInfo]
type zeroTrustGatewayProxyEndpointListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayProxyEndpointDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayProxyEndpointDeleteResponseEnvelope struct {
	Errors   []ZeroTrustGatewayProxyEndpointDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayProxyEndpointDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayProxyEndpointDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayProxyEndpointDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayProxyEndpointDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayProxyEndpointDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayProxyEndpointDeleteResponseEnvelope]
type zeroTrustGatewayProxyEndpointDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayProxyEndpointDeleteResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustGatewayProxyEndpointDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayProxyEndpointDeleteResponseEnvelopeErrors]
type zeroTrustGatewayProxyEndpointDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayProxyEndpointDeleteResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustGatewayProxyEndpointDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointDeleteResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayProxyEndpointDeleteResponseEnvelopeMessages]
type zeroTrustGatewayProxyEndpointDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayProxyEndpointDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayProxyEndpointDeleteResponseEnvelopeSuccessTrue ZeroTrustGatewayProxyEndpointDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayProxyEndpointEditParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]string] `json:"ips"`
	// The name of the proxy endpoint.
	Name param.Field[string] `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain param.Field[string] `json:"subdomain"`
}

func (r ZeroTrustGatewayProxyEndpointEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayProxyEndpointEditResponseEnvelope struct {
	Errors   []ZeroTrustGatewayProxyEndpointEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayProxyEndpointEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayProxyEndpointEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayProxyEndpointEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayProxyEndpointEditResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayProxyEndpointEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayProxyEndpointEditResponseEnvelope]
type zeroTrustGatewayProxyEndpointEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayProxyEndpointEditResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustGatewayProxyEndpointEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayProxyEndpointEditResponseEnvelopeErrors]
type zeroTrustGatewayProxyEndpointEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayProxyEndpointEditResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustGatewayProxyEndpointEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayProxyEndpointEditResponseEnvelopeMessages]
type zeroTrustGatewayProxyEndpointEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayProxyEndpointEditResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayProxyEndpointEditResponseEnvelopeSuccessTrue ZeroTrustGatewayProxyEndpointEditResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayProxyEndpointGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayProxyEndpointGetResponseEnvelope struct {
	Errors   []ZeroTrustGatewayProxyEndpointGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayProxyEndpointGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayProxyEndpointGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayProxyEndpointGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayProxyEndpointGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayProxyEndpointGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayProxyEndpointGetResponseEnvelope]
type zeroTrustGatewayProxyEndpointGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayProxyEndpointGetResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustGatewayProxyEndpointGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayProxyEndpointGetResponseEnvelopeErrors]
type zeroTrustGatewayProxyEndpointGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayProxyEndpointGetResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustGatewayProxyEndpointGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustGatewayProxyEndpointGetResponseEnvelopeMessages]
type zeroTrustGatewayProxyEndpointGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpointGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayProxyEndpointGetResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayProxyEndpointGetResponseEnvelopeSuccessTrue ZeroTrustGatewayProxyEndpointGetResponseEnvelopeSuccess = true
)

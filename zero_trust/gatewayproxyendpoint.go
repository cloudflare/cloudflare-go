// File generated from our OpenAPI spec by Stainless.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
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

// Creates a new Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) New(ctx context.Context, params GatewayProxyEndpointNewParams, opts ...option.RequestOption) (res *ZeroTrustGatewayProxyEndpoints, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) List(ctx context.Context, query GatewayProxyEndpointListParams, opts ...option.RequestOption) (res *[]ZeroTrustGatewayProxyEndpoints, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a configured Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) Delete(ctx context.Context, proxyEndpointID interface{}, body GatewayProxyEndpointDeleteParams, opts ...option.RequestOption) (res *GatewayProxyEndpointDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", body.AccountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) Edit(ctx context.Context, proxyEndpointID interface{}, params GatewayProxyEndpointEditParams, opts ...option.RequestOption) (res *ZeroTrustGatewayProxyEndpoints, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", params.AccountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all Zero Trust Gateway proxy endpoints for an account.
func (r *GatewayProxyEndpointService) Get(ctx context.Context, proxyEndpointID interface{}, query GatewayProxyEndpointGetParams, opts ...option.RequestOption) (res *ZeroTrustGatewayProxyEndpoints, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayProxyEndpointGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/proxy_endpoints/%v", query.AccountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustGatewayProxyEndpoints struct {
	ID        interface{} `json:"id"`
	CreatedAt time.Time   `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []string `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string                             `json:"subdomain"`
	UpdatedAt time.Time                          `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayProxyEndpointsJSON `json:"-"`
}

// zeroTrustGatewayProxyEndpointsJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayProxyEndpoints]
type zeroTrustGatewayProxyEndpointsJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayProxyEndpoints) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayProxyEndpointsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.GatewayProxyEndpointDeleteResponseUnknown] or
// [shared.UnionString].
type GatewayProxyEndpointDeleteResponse interface {
	ImplementsZeroTrustGatewayProxyEndpointDeleteResponse()
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

type GatewayProxyEndpointNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]string] `json:"ips,required"`
	// The name of the proxy endpoint.
	Name param.Field[string] `json:"name,required"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain param.Field[string] `json:"subdomain"`
}

func (r GatewayProxyEndpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayProxyEndpointNewResponseEnvelope struct {
	Errors   []GatewayProxyEndpointNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayProxyEndpointNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayProxyEndpoints                    `json:"result,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayProxyEndpointNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayProxyEndpointNewResponseEnvelope]
type gatewayProxyEndpointNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayProxyEndpointNewResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    gatewayProxyEndpointNewResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayProxyEndpointNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [GatewayProxyEndpointNewResponseEnvelopeErrors]
type gatewayProxyEndpointNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayProxyEndpointNewResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    gatewayProxyEndpointNewResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayProxyEndpointNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [GatewayProxyEndpointNewResponseEnvelopeMessages]
type gatewayProxyEndpointNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayProxyEndpointNewResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointNewResponseEnvelopeSuccessTrue GatewayProxyEndpointNewResponseEnvelopeSuccess = true
)

type GatewayProxyEndpointListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type GatewayProxyEndpointListResponseEnvelope struct {
	Errors   []GatewayProxyEndpointListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayProxyEndpointListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustGatewayProxyEndpoints                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    GatewayProxyEndpointListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo GatewayProxyEndpointListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       gatewayProxyEndpointListResponseEnvelopeJSON       `json:"-"`
}

// gatewayProxyEndpointListResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayProxyEndpointListResponseEnvelope]
type gatewayProxyEndpointListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r gatewayProxyEndpointListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayProxyEndpointListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayProxyEndpointListResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointListResponseEnvelopeSuccessTrue GatewayProxyEndpointListResponseEnvelopeSuccess = true
)

type GatewayProxyEndpointListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       gatewayProxyEndpointListResponseEnvelopeResultInfoJSON `json:"-"`
}

// gatewayProxyEndpointListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [GatewayProxyEndpointListResponseEnvelopeResultInfo]
type gatewayProxyEndpointListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type GatewayProxyEndpointDeleteParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

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

func (r gatewayProxyEndpointDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r gatewayProxyEndpointDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayProxyEndpointDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayProxyEndpointDeleteResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointDeleteResponseEnvelopeSuccessTrue GatewayProxyEndpointDeleteResponseEnvelopeSuccess = true
)

type GatewayProxyEndpointEditParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]string] `json:"ips"`
	// The name of the proxy endpoint.
	Name param.Field[string] `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain param.Field[string] `json:"subdomain"`
}

func (r GatewayProxyEndpointEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayProxyEndpointEditResponseEnvelope struct {
	Errors   []GatewayProxyEndpointEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayProxyEndpointEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayProxyEndpoints                     `json:"result,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayProxyEndpointEditResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayProxyEndpointEditResponseEnvelope]
type gatewayProxyEndpointEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type GatewayProxyEndpointEditResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    gatewayProxyEndpointEditResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayProxyEndpointEditResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [GatewayProxyEndpointEditResponseEnvelopeErrors]
type gatewayProxyEndpointEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type GatewayProxyEndpointEditResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    gatewayProxyEndpointEditResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayProxyEndpointEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [GatewayProxyEndpointEditResponseEnvelopeMessages]
type gatewayProxyEndpointEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayProxyEndpointEditResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointEditResponseEnvelopeSuccessTrue GatewayProxyEndpointEditResponseEnvelopeSuccess = true
)

type GatewayProxyEndpointGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type GatewayProxyEndpointGetResponseEnvelope struct {
	Errors   []GatewayProxyEndpointGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayProxyEndpointGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayProxyEndpoints                    `json:"result,required"`
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

func (r gatewayProxyEndpointGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r gatewayProxyEndpointGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r gatewayProxyEndpointGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayProxyEndpointGetResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointGetResponseEnvelopeSuccessTrue GatewayProxyEndpointGetResponseEnvelopeSuccess = true
)

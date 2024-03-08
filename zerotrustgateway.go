// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// ZeroTrustGatewayService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustGatewayService] method
// instead.
type ZeroTrustGatewayService struct {
	Options          []option.RequestOption
	AuditSSHSettings *ZeroTrustGatewayAuditSSHSettingService
	Categories       *ZeroTrustGatewayCategoryService
	AppTypes         *ZeroTrustGatewayAppTypeService
	Configurations   *ZeroTrustGatewayConfigurationService
	Lists            *ZeroTrustGatewayListService
	Locations        *ZeroTrustGatewayLocationService
	Loggings         *ZeroTrustGatewayLoggingService
	ProxyEndpoints   *ZeroTrustGatewayProxyEndpointService
	Rules            *ZeroTrustGatewayRuleService
}

// NewZeroTrustGatewayService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustGatewayService(opts ...option.RequestOption) (r *ZeroTrustGatewayService) {
	r = &ZeroTrustGatewayService{}
	r.Options = opts
	r.AuditSSHSettings = NewZeroTrustGatewayAuditSSHSettingService(opts...)
	r.Categories = NewZeroTrustGatewayCategoryService(opts...)
	r.AppTypes = NewZeroTrustGatewayAppTypeService(opts...)
	r.Configurations = NewZeroTrustGatewayConfigurationService(opts...)
	r.Lists = NewZeroTrustGatewayListService(opts...)
	r.Locations = NewZeroTrustGatewayLocationService(opts...)
	r.Loggings = NewZeroTrustGatewayLoggingService(opts...)
	r.ProxyEndpoints = NewZeroTrustGatewayProxyEndpointService(opts...)
	r.Rules = NewZeroTrustGatewayRuleService(opts...)
	return
}

// Creates a Zero Trust account with an existing Cloudflare account.
func (r *ZeroTrustGatewayService) New(ctx context.Context, body ZeroTrustGatewayNewParams, opts ...option.RequestOption) (res *ZeroTrustGatewayNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets information about the current Zero Trust account.
func (r *ZeroTrustGatewayService) List(ctx context.Context, query ZeroTrustGatewayListParams, opts ...option.RequestOption) (res *ZeroTrustGatewayListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustGatewayNewResponse struct {
	// Cloudflare account ID.
	ID string `json:"id"`
	// Gateway internal ID.
	GatewayTag string `json:"gateway_tag"`
	// The name of the provider. Usually Cloudflare.
	ProviderName string                          `json:"provider_name"`
	JSON         zeroTrustGatewayNewResponseJSON `json:"-"`
}

// zeroTrustGatewayNewResponseJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayNewResponse]
type zeroTrustGatewayNewResponseJSON struct {
	ID           apijson.Field
	GatewayTag   apijson.Field
	ProviderName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGatewayNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayListResponse struct {
	// Cloudflare account ID.
	ID string `json:"id"`
	// Gateway internal ID.
	GatewayTag string `json:"gateway_tag"`
	// The name of the provider. Usually Cloudflare.
	ProviderName string                           `json:"provider_name"`
	JSON         zeroTrustGatewayListResponseJSON `json:"-"`
}

// zeroTrustGatewayListResponseJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayListResponse]
type zeroTrustGatewayListResponseJSON struct {
	ID           apijson.Field
	GatewayTag   apijson.Field
	ProviderName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGatewayListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayNewParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayNewResponseEnvelope struct {
	Errors   []ZeroTrustGatewayNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayNewResponseEnvelope]
type zeroTrustGatewayNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zeroTrustGatewayNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayNewResponseEnvelopeErrors]
type zeroTrustGatewayNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zeroTrustGatewayNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayNewResponseEnvelopeMessages]
type zeroTrustGatewayNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustGatewayNewResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayNewResponseEnvelopeSuccessTrue ZeroTrustGatewayNewResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayListParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayListResponseEnvelope struct {
	Errors   []ZeroTrustGatewayListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayListResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayListResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayListResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayListResponseEnvelope]
type zeroTrustGatewayListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    zeroTrustGatewayListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayListResponseEnvelopeErrors]
type zeroTrustGatewayListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zeroTrustGatewayListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayListResponseEnvelopeMessages]
type zeroTrustGatewayListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustGatewayListResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayListResponseEnvelopeSuccessTrue ZeroTrustGatewayListResponseEnvelopeSuccess = true
)

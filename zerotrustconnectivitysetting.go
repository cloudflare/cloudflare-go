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

// ZeroTrustConnectivitySettingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustConnectivitySettingService] method instead.
type ZeroTrustConnectivitySettingService struct {
	Options []option.RequestOption
}

// NewZeroTrustConnectivitySettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustConnectivitySettingService(opts ...option.RequestOption) (r *ZeroTrustConnectivitySettingService) {
	r = &ZeroTrustConnectivitySettingService{}
	r.Options = opts
	return
}

// Updates the Zero Trust Connectivity Settings for the given account.
func (r *ZeroTrustConnectivitySettingService) Edit(ctx context.Context, params ZeroTrustConnectivitySettingEditParams, opts ...option.RequestOption) (res *ZeroTrustConnectivitySettingEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustConnectivitySettingEditResponseEnvelope
	path := fmt.Sprintf("accounts/%s/zerotrust/connectivity_settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the Zero Trust Connectivity Settings for the given account.
func (r *ZeroTrustConnectivitySettingService) Get(ctx context.Context, query ZeroTrustConnectivitySettingGetParams, opts ...option.RequestOption) (res *ZeroTrustConnectivitySettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustConnectivitySettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/zerotrust/connectivity_settings", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustConnectivitySettingEditResponse struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled bool `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWARPEnabled bool                                         `json:"offramp_warp_enabled"`
	JSON               zeroTrustConnectivitySettingEditResponseJSON `json:"-"`
}

// zeroTrustConnectivitySettingEditResponseJSON contains the JSON metadata for the
// struct [ZeroTrustConnectivitySettingEditResponse]
type zeroTrustConnectivitySettingEditResponseJSON struct {
	IcmpProxyEnabled   apijson.Field
	OfframpWARPEnabled apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustConnectivitySettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustConnectivitySettingEditResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustConnectivitySettingGetResponse struct {
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled bool `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWARPEnabled bool                                        `json:"offramp_warp_enabled"`
	JSON               zeroTrustConnectivitySettingGetResponseJSON `json:"-"`
}

// zeroTrustConnectivitySettingGetResponseJSON contains the JSON metadata for the
// struct [ZeroTrustConnectivitySettingGetResponse]
type zeroTrustConnectivitySettingGetResponseJSON struct {
	IcmpProxyEnabled   apijson.Field
	OfframpWARPEnabled apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustConnectivitySettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustConnectivitySettingGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustConnectivitySettingEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A flag to enable the ICMP proxy for the account network.
	IcmpProxyEnabled param.Field[bool] `json:"icmp_proxy_enabled"`
	// A flag to enable WARP to WARP traffic.
	OfframpWARPEnabled param.Field[bool] `json:"offramp_warp_enabled"`
}

func (r ZeroTrustConnectivitySettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustConnectivitySettingEditResponseEnvelope struct {
	Errors   []ZeroTrustConnectivitySettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustConnectivitySettingEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustConnectivitySettingEditResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustConnectivitySettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustConnectivitySettingEditResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustConnectivitySettingEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustConnectivitySettingEditResponseEnvelope]
type zeroTrustConnectivitySettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustConnectivitySettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustConnectivitySettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustConnectivitySettingEditResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustConnectivitySettingEditResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustConnectivitySettingEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustConnectivitySettingEditResponseEnvelopeErrors]
type zeroTrustConnectivitySettingEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustConnectivitySettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustConnectivitySettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustConnectivitySettingEditResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustConnectivitySettingEditResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustConnectivitySettingEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustConnectivitySettingEditResponseEnvelopeMessages]
type zeroTrustConnectivitySettingEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustConnectivitySettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustConnectivitySettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustConnectivitySettingEditResponseEnvelopeSuccess bool

const (
	ZeroTrustConnectivitySettingEditResponseEnvelopeSuccessTrue ZeroTrustConnectivitySettingEditResponseEnvelopeSuccess = true
)

type ZeroTrustConnectivitySettingGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustConnectivitySettingGetResponseEnvelope struct {
	Errors   []ZeroTrustConnectivitySettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustConnectivitySettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustConnectivitySettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustConnectivitySettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustConnectivitySettingGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustConnectivitySettingGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustConnectivitySettingGetResponseEnvelope]
type zeroTrustConnectivitySettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustConnectivitySettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustConnectivitySettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustConnectivitySettingGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustConnectivitySettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustConnectivitySettingGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustConnectivitySettingGetResponseEnvelopeErrors]
type zeroTrustConnectivitySettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustConnectivitySettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustConnectivitySettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustConnectivitySettingGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustConnectivitySettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustConnectivitySettingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustConnectivitySettingGetResponseEnvelopeMessages]
type zeroTrustConnectivitySettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustConnectivitySettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustConnectivitySettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustConnectivitySettingGetResponseEnvelopeSuccess bool

const (
	ZeroTrustConnectivitySettingGetResponseEnvelopeSuccessTrue ZeroTrustConnectivitySettingGetResponseEnvelopeSuccess = true
)

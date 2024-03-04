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

// ZeroTrustDevicePolicyDefaultPolicyService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustDevicePolicyDefaultPolicyService] method instead.
type ZeroTrustDevicePolicyDefaultPolicyService struct {
	Options []option.RequestOption
}

// NewZeroTrustDevicePolicyDefaultPolicyService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZeroTrustDevicePolicyDefaultPolicyService(opts ...option.RequestOption) (r *ZeroTrustDevicePolicyDefaultPolicyService) {
	r = &ZeroTrustDevicePolicyDefaultPolicyService{}
	r.Options = opts
	return
}

// Fetches the default device settings profile for an account.
func (r *ZeroTrustDevicePolicyDefaultPolicyService) Get(ctx context.Context, query ZeroTrustDevicePolicyDefaultPolicyGetParams, opts ...option.RequestOption) (res *[]ZeroTrustDevicePolicyDefaultPolicyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/devices/policy", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustDevicePolicyDefaultPolicyGetResponse = interface{}

type ZeroTrustDevicePolicyDefaultPolicyGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelope struct {
	Errors   []ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustDevicePolicyDefaultPolicyGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelope]
type zeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    zeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeErrors]
type zeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeMessages]
type zeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful.
type ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeSuccessTrue ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeSuccess = true
)

type ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                             `json:"total_count"`
	JSON       zeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeResultInfo]
type zeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDevicePolicyDefaultPolicyGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

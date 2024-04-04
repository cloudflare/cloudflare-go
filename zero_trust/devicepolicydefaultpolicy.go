// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DevicePolicyDefaultPolicyService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewDevicePolicyDefaultPolicyService] method instead.
type DevicePolicyDefaultPolicyService struct {
	Options []option.RequestOption
}

// NewDevicePolicyDefaultPolicyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyDefaultPolicyService(opts ...option.RequestOption) (r *DevicePolicyDefaultPolicyService) {
	r = &DevicePolicyDefaultPolicyService{}
	r.Options = opts
	return
}

// Fetches the default device settings profile for an account.
func (r *DevicePolicyDefaultPolicyService) Get(ctx context.Context, query DevicePolicyDefaultPolicyGetParams, opts ...option.RequestOption) (res *[]DevicePolicyDefaultPolicyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DevicePolicyDefaultPolicyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/devices/policy", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyDefaultPolicyGetResponse = interface{}

type DevicePolicyDefaultPolicyGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DevicePolicyDefaultPolicyGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172           `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172           `json:"messages,required"`
	Result   []DevicePolicyDefaultPolicyGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success    DevicePolicyDefaultPolicyGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DevicePolicyDefaultPolicyGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       devicePolicyDefaultPolicyGetResponseEnvelopeJSON       `json:"-"`
}

// devicePolicyDefaultPolicyGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DevicePolicyDefaultPolicyGetResponseEnvelope]
type devicePolicyDefaultPolicyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultPolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultPolicyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyDefaultPolicyGetResponseEnvelopeSuccess bool

const (
	DevicePolicyDefaultPolicyGetResponseEnvelopeSuccessTrue DevicePolicyDefaultPolicyGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyDefaultPolicyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyDefaultPolicyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyDefaultPolicyGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       devicePolicyDefaultPolicyGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// devicePolicyDefaultPolicyGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DevicePolicyDefaultPolicyGetResponseEnvelopeResultInfo]
type devicePolicyDefaultPolicyGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultPolicyGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultPolicyGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

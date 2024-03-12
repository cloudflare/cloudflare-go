// File generated from our OpenAPI spec by Stainless.

package alerting

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// V3DestinationEligibleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewV3DestinationEligibleService]
// method instead.
type V3DestinationEligibleService struct {
	Options []option.RequestOption
}

// NewV3DestinationEligibleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV3DestinationEligibleService(opts ...option.RequestOption) (r *V3DestinationEligibleService) {
	r = &V3DestinationEligibleService{}
	r.Options = opts
	return
}

// Get a list of all delivery mechanism types for which an account is eligible.
func (r *V3DestinationEligibleService) Get(ctx context.Context, query V3DestinationEligibleGetParams, opts ...option.RequestOption) (res *V3DestinationEligibleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3DestinationEligibleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/eligible", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [alerting.V3DestinationEligibleGetResponseUnknown],
// [alerting.V3DestinationEligibleGetResponseArray] or [shared.UnionString].
type V3DestinationEligibleGetResponse interface {
	ImplementsAlertingV3DestinationEligibleGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V3DestinationEligibleGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V3DestinationEligibleGetResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type V3DestinationEligibleGetResponseArray []interface{}

func (r V3DestinationEligibleGetResponseArray) ImplementsAlertingV3DestinationEligibleGetResponse() {}

type V3DestinationEligibleGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type V3DestinationEligibleGetResponseEnvelope struct {
	Errors   []V3DestinationEligibleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3DestinationEligibleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   V3DestinationEligibleGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    V3DestinationEligibleGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo V3DestinationEligibleGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       v3DestinationEligibleGetResponseEnvelopeJSON       `json:"-"`
}

// v3DestinationEligibleGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [V3DestinationEligibleGetResponseEnvelope]
type v3DestinationEligibleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationEligibleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationEligibleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3DestinationEligibleGetResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    v3DestinationEligibleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// v3DestinationEligibleGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [V3DestinationEligibleGetResponseEnvelopeErrors]
type v3DestinationEligibleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationEligibleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationEligibleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3DestinationEligibleGetResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    v3DestinationEligibleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// v3DestinationEligibleGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [V3DestinationEligibleGetResponseEnvelopeMessages]
type v3DestinationEligibleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationEligibleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationEligibleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3DestinationEligibleGetResponseEnvelopeSuccess bool

const (
	V3DestinationEligibleGetResponseEnvelopeSuccessTrue V3DestinationEligibleGetResponseEnvelopeSuccess = true
)

type V3DestinationEligibleGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       v3DestinationEligibleGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// v3DestinationEligibleGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [V3DestinationEligibleGetResponseEnvelopeResultInfo]
type v3DestinationEligibleGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationEligibleGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationEligibleGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

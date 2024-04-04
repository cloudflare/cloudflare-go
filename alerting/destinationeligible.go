// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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

// DestinationEligibleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDestinationEligibleService]
// method instead.
type DestinationEligibleService struct {
	Options []option.RequestOption
}

// NewDestinationEligibleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDestinationEligibleService(opts ...option.RequestOption) (r *DestinationEligibleService) {
	r = &DestinationEligibleService{}
	r.Options = opts
	return
}

// Get a list of all delivery mechanism types for which an account is eligible.
func (r *DestinationEligibleService) Get(ctx context.Context, query DestinationEligibleGetParams, opts ...option.RequestOption) (res *DestinationEligibleGetResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env DestinationEligibleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/eligible", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [alerting.DestinationEligibleGetResponseUnknown],
// [alerting.DestinationEligibleGetResponseArray] or [shared.UnionString].
type DestinationEligibleGetResponseUnion interface {
	ImplementsAlertingDestinationEligibleGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DestinationEligibleGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DestinationEligibleGetResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type DestinationEligibleGetResponseArray []interface{}

func (r DestinationEligibleGetResponseArray) ImplementsAlertingDestinationEligibleGetResponseUnion() {
}

type DestinationEligibleGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationEligibleGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   DestinationEligibleGetResponseUnion                       `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DestinationEligibleGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DestinationEligibleGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       destinationEligibleGetResponseEnvelopeJSON       `json:"-"`
}

// destinationEligibleGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationEligibleGetResponseEnvelope]
type destinationEligibleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationEligibleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationEligibleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationEligibleGetResponseEnvelopeSuccess bool

const (
	DestinationEligibleGetResponseEnvelopeSuccessTrue DestinationEligibleGetResponseEnvelopeSuccess = true
)

func (r DestinationEligibleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationEligibleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DestinationEligibleGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       destinationEligibleGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// destinationEligibleGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [DestinationEligibleGetResponseEnvelopeResultInfo]
type destinationEligibleGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationEligibleGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationEligibleGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

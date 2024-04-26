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
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// AvailableAlertService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAvailableAlertService] method
// instead.
type AvailableAlertService struct {
	Options []option.RequestOption
}

// NewAvailableAlertService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAvailableAlertService(opts ...option.RequestOption) (r *AvailableAlertService) {
	r = &AvailableAlertService{}
	r.Options = opts
	return
}

// Gets a list of all alert types for which an account is eligible.
func (r *AvailableAlertService) List(ctx context.Context, query AvailableAlertListParams, opts ...option.RequestOption) (res *AvailableAlertListResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env AvailableAlertListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/available_alerts", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [alerting.AvailableAlertListResponseUnknown],
// [alerting.AvailableAlertListResponseArray] or [shared.UnionString].
type AvailableAlertListResponseUnion interface {
	ImplementsAlertingAvailableAlertListResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AvailableAlertListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AvailableAlertListResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AvailableAlertListResponseArray []interface{}

func (r AvailableAlertListResponseArray) ImplementsAlertingAvailableAlertListResponseUnion() {}

type AvailableAlertListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type AvailableAlertListResponseEnvelope struct {
	Errors   []shared.ResponseInfo           `json:"errors,required"`
	Messages []shared.ResponseInfo           `json:"messages,required"`
	Result   AvailableAlertListResponseUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AvailableAlertListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AvailableAlertListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       availableAlertListResponseEnvelopeJSON       `json:"-"`
}

// availableAlertListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AvailableAlertListResponseEnvelope]
type availableAlertListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableAlertListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availableAlertListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AvailableAlertListResponseEnvelopeSuccess bool

const (
	AvailableAlertListResponseEnvelopeSuccessTrue AvailableAlertListResponseEnvelopeSuccess = true
)

func (r AvailableAlertListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AvailableAlertListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AvailableAlertListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       availableAlertListResponseEnvelopeResultInfoJSON `json:"-"`
}

// availableAlertListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [AvailableAlertListResponseEnvelopeResultInfo]
type availableAlertListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableAlertListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availableAlertListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

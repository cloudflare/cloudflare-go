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

// V3Service contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewV3Service] method instead.
type V3Service struct {
	Options      []option.RequestOption
	Destinations *V3DestinationService
	Histories    *V3HistoryService
	Policies     *V3PolicyService
}

// NewV3Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV3Service(opts ...option.RequestOption) (r *V3Service) {
	r = &V3Service{}
	r.Options = opts
	r.Destinations = NewV3DestinationService(opts...)
	r.Histories = NewV3HistoryService(opts...)
	r.Policies = NewV3PolicyService(opts...)
	return
}

// Gets a list of all alert types for which an account is eligible.
func (r *V3Service) List(ctx context.Context, query V3ListParams, opts ...option.RequestOption) (res *V3ListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3ListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/available_alerts", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [alerting.V3ListResponseUnknown],
// [alerting.V3ListResponseArray] or [shared.UnionString].
type V3ListResponse interface {
	ImplementsAlertingV3ListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V3ListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V3ListResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type V3ListResponseArray []interface{}

func (r V3ListResponseArray) ImplementsAlertingV3ListResponse() {}

type V3ListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type V3ListResponseEnvelope struct {
	Errors   []V3ListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3ListResponseEnvelopeMessages `json:"messages,required"`
	Result   V3ListResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    V3ListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo V3ListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       v3ListResponseEnvelopeJSON       `json:"-"`
}

// v3ListResponseEnvelopeJSON contains the JSON metadata for the struct
// [V3ListResponseEnvelope]
type v3ListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3ListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3ListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3ListResponseEnvelopeErrors struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    v3ListResponseEnvelopeErrorsJSON `json:"-"`
}

// v3ListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [V3ListResponseEnvelopeErrors]
type v3ListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3ListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3ListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3ListResponseEnvelopeMessages struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    v3ListResponseEnvelopeMessagesJSON `json:"-"`
}

// v3ListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [V3ListResponseEnvelopeMessages]
type v3ListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3ListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3ListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3ListResponseEnvelopeSuccess bool

const (
	V3ListResponseEnvelopeSuccessTrue V3ListResponseEnvelopeSuccess = true
)

type V3ListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                              `json:"total_count"`
	JSON       v3ListResponseEnvelopeResultInfoJSON `json:"-"`
}

// v3ListResponseEnvelopeResultInfoJSON contains the JSON metadata for the struct
// [V3ListResponseEnvelopeResultInfo]
type v3ListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3ListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3ListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

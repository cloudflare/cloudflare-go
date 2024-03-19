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

// V3DestinationPagerdutyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewV3DestinationPagerdutyService]
// method instead.
type V3DestinationPagerdutyService struct {
	Options []option.RequestOption
}

// NewV3DestinationPagerdutyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV3DestinationPagerdutyService(opts ...option.RequestOption) (r *V3DestinationPagerdutyService) {
	r = &V3DestinationPagerdutyService{}
	r.Options = opts
	return
}

// Creates a new token for integrating with PagerDuty.
func (r *V3DestinationPagerdutyService) New(ctx context.Context, body V3DestinationPagerdutyNewParams, opts ...option.RequestOption) (res *V3DestinationPagerdutyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3DestinationPagerdutyNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty/connect", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes all the PagerDuty Services connected to the account.
func (r *V3DestinationPagerdutyService) Delete(ctx context.Context, body V3DestinationPagerdutyDeleteParams, opts ...option.RequestOption) (res *V3DestinationPagerdutyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3DestinationPagerdutyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a list of all configured PagerDuty services.
func (r *V3DestinationPagerdutyService) Get(ctx context.Context, query V3DestinationPagerdutyGetParams, opts ...option.RequestOption) (res *[]AaaPagerduty, err error) {
	opts = append(r.Options[:], opts...)
	var env V3DestinationPagerdutyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Links PagerDuty with the account using the integration token.
func (r *V3DestinationPagerdutyService) Link(ctx context.Context, tokenID string, query V3DestinationPagerdutyLinkParams, opts ...option.RequestOption) (res *V3DestinationPagerdutyLinkResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3DestinationPagerdutyLinkResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty/connect/%s", query.AccountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AaaPagerduty struct {
	// UUID
	ID string `json:"id"`
	// The name of the pagerduty service.
	Name string           `json:"name"`
	JSON aaaPagerdutyJSON `json:"-"`
}

// aaaPagerdutyJSON contains the JSON metadata for the struct [AaaPagerduty]
type aaaPagerdutyJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AaaPagerduty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aaaPagerdutyJSON) RawJSON() string {
	return r.raw
}

type V3DestinationPagerdutyNewResponse struct {
	// token in form of UUID
	ID   string                                `json:"id"`
	JSON v3DestinationPagerdutyNewResponseJSON `json:"-"`
}

// v3DestinationPagerdutyNewResponseJSON contains the JSON metadata for the struct
// [V3DestinationPagerdutyNewResponse]
type v3DestinationPagerdutyNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyNewResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [alerting.V3DestinationPagerdutyDeleteResponseUnknown],
// [alerting.V3DestinationPagerdutyDeleteResponseArray] or [shared.UnionString].
type V3DestinationPagerdutyDeleteResponse interface {
	ImplementsAlertingV3DestinationPagerdutyDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V3DestinationPagerdutyDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V3DestinationPagerdutyDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type V3DestinationPagerdutyDeleteResponseArray []interface{}

func (r V3DestinationPagerdutyDeleteResponseArray) ImplementsAlertingV3DestinationPagerdutyDeleteResponse() {
}

type V3DestinationPagerdutyLinkResponse struct {
	// UUID
	ID   string                                 `json:"id"`
	JSON v3DestinationPagerdutyLinkResponseJSON `json:"-"`
}

// v3DestinationPagerdutyLinkResponseJSON contains the JSON metadata for the struct
// [V3DestinationPagerdutyLinkResponse]
type v3DestinationPagerdutyLinkResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyLinkResponseJSON) RawJSON() string {
	return r.raw
}

type V3DestinationPagerdutyNewParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type V3DestinationPagerdutyNewResponseEnvelope struct {
	Errors   []V3DestinationPagerdutyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3DestinationPagerdutyNewResponseEnvelopeMessages `json:"messages,required"`
	Result   V3DestinationPagerdutyNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V3DestinationPagerdutyNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    v3DestinationPagerdutyNewResponseEnvelopeJSON    `json:"-"`
}

// v3DestinationPagerdutyNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [V3DestinationPagerdutyNewResponseEnvelope]
type v3DestinationPagerdutyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3DestinationPagerdutyNewResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    v3DestinationPagerdutyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// v3DestinationPagerdutyNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [V3DestinationPagerdutyNewResponseEnvelopeErrors]
type v3DestinationPagerdutyNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3DestinationPagerdutyNewResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    v3DestinationPagerdutyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// v3DestinationPagerdutyNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [V3DestinationPagerdutyNewResponseEnvelopeMessages]
type v3DestinationPagerdutyNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3DestinationPagerdutyNewResponseEnvelopeSuccess bool

const (
	V3DestinationPagerdutyNewResponseEnvelopeSuccessTrue V3DestinationPagerdutyNewResponseEnvelopeSuccess = true
)

func (r V3DestinationPagerdutyNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V3DestinationPagerdutyNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V3DestinationPagerdutyDeleteParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type V3DestinationPagerdutyDeleteResponseEnvelope struct {
	Errors   []V3DestinationPagerdutyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3DestinationPagerdutyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   V3DestinationPagerdutyDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    V3DestinationPagerdutyDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo V3DestinationPagerdutyDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       v3DestinationPagerdutyDeleteResponseEnvelopeJSON       `json:"-"`
}

// v3DestinationPagerdutyDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [V3DestinationPagerdutyDeleteResponseEnvelope]
type v3DestinationPagerdutyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3DestinationPagerdutyDeleteResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    v3DestinationPagerdutyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// v3DestinationPagerdutyDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [V3DestinationPagerdutyDeleteResponseEnvelopeErrors]
type v3DestinationPagerdutyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3DestinationPagerdutyDeleteResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    v3DestinationPagerdutyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// v3DestinationPagerdutyDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [V3DestinationPagerdutyDeleteResponseEnvelopeMessages]
type v3DestinationPagerdutyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3DestinationPagerdutyDeleteResponseEnvelopeSuccess bool

const (
	V3DestinationPagerdutyDeleteResponseEnvelopeSuccessTrue V3DestinationPagerdutyDeleteResponseEnvelopeSuccess = true
)

func (r V3DestinationPagerdutyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V3DestinationPagerdutyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V3DestinationPagerdutyDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       v3DestinationPagerdutyDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// v3DestinationPagerdutyDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [V3DestinationPagerdutyDeleteResponseEnvelopeResultInfo]
type v3DestinationPagerdutyDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type V3DestinationPagerdutyGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type V3DestinationPagerdutyGetResponseEnvelope struct {
	Errors   []V3DestinationPagerdutyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3DestinationPagerdutyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []AaaPagerduty                                      `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    V3DestinationPagerdutyGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo V3DestinationPagerdutyGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       v3DestinationPagerdutyGetResponseEnvelopeJSON       `json:"-"`
}

// v3DestinationPagerdutyGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [V3DestinationPagerdutyGetResponseEnvelope]
type v3DestinationPagerdutyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3DestinationPagerdutyGetResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    v3DestinationPagerdutyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// v3DestinationPagerdutyGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [V3DestinationPagerdutyGetResponseEnvelopeErrors]
type v3DestinationPagerdutyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3DestinationPagerdutyGetResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    v3DestinationPagerdutyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// v3DestinationPagerdutyGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [V3DestinationPagerdutyGetResponseEnvelopeMessages]
type v3DestinationPagerdutyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3DestinationPagerdutyGetResponseEnvelopeSuccess bool

const (
	V3DestinationPagerdutyGetResponseEnvelopeSuccessTrue V3DestinationPagerdutyGetResponseEnvelopeSuccess = true
)

func (r V3DestinationPagerdutyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V3DestinationPagerdutyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type V3DestinationPagerdutyGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                 `json:"total_count"`
	JSON       v3DestinationPagerdutyGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// v3DestinationPagerdutyGetResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [V3DestinationPagerdutyGetResponseEnvelopeResultInfo]
type v3DestinationPagerdutyGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type V3DestinationPagerdutyLinkParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type V3DestinationPagerdutyLinkResponseEnvelope struct {
	Errors   []V3DestinationPagerdutyLinkResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3DestinationPagerdutyLinkResponseEnvelopeMessages `json:"messages,required"`
	Result   V3DestinationPagerdutyLinkResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V3DestinationPagerdutyLinkResponseEnvelopeSuccess `json:"success,required"`
	JSON    v3DestinationPagerdutyLinkResponseEnvelopeJSON    `json:"-"`
}

// v3DestinationPagerdutyLinkResponseEnvelopeJSON contains the JSON metadata for
// the struct [V3DestinationPagerdutyLinkResponseEnvelope]
type v3DestinationPagerdutyLinkResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyLinkResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyLinkResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3DestinationPagerdutyLinkResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    v3DestinationPagerdutyLinkResponseEnvelopeErrorsJSON `json:"-"`
}

// v3DestinationPagerdutyLinkResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [V3DestinationPagerdutyLinkResponseEnvelopeErrors]
type v3DestinationPagerdutyLinkResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyLinkResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyLinkResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3DestinationPagerdutyLinkResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    v3DestinationPagerdutyLinkResponseEnvelopeMessagesJSON `json:"-"`
}

// v3DestinationPagerdutyLinkResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [V3DestinationPagerdutyLinkResponseEnvelopeMessages]
type v3DestinationPagerdutyLinkResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationPagerdutyLinkResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationPagerdutyLinkResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3DestinationPagerdutyLinkResponseEnvelopeSuccess bool

const (
	V3DestinationPagerdutyLinkResponseEnvelopeSuccessTrue V3DestinationPagerdutyLinkResponseEnvelopeSuccess = true
)

func (r V3DestinationPagerdutyLinkResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case V3DestinationPagerdutyLinkResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

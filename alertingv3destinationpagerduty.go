// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AlertingV3DestinationPagerdutyService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAlertingV3DestinationPagerdutyService] method instead.
type AlertingV3DestinationPagerdutyService struct {
	Options []option.RequestOption
}

// NewAlertingV3DestinationPagerdutyService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAlertingV3DestinationPagerdutyService(opts ...option.RequestOption) (r *AlertingV3DestinationPagerdutyService) {
	r = &AlertingV3DestinationPagerdutyService{}
	r.Options = opts
	return
}

// Creates a new token for integrating with PagerDuty.
func (r *AlertingV3DestinationPagerdutyService) New(ctx context.Context, body AlertingV3DestinationPagerdutyNewParams, opts ...option.RequestOption) (res *AlertingV3DestinationPagerdutyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3DestinationPagerdutyNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty/connect", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes all the PagerDuty Services connected to the account.
func (r *AlertingV3DestinationPagerdutyService) Delete(ctx context.Context, body AlertingV3DestinationPagerdutyDeleteParams, opts ...option.RequestOption) (res *AlertingV3DestinationPagerdutyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3DestinationPagerdutyDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a list of all configured PagerDuty services.
func (r *AlertingV3DestinationPagerdutyService) Get(ctx context.Context, query AlertingV3DestinationPagerdutyGetParams, opts ...option.RequestOption) (res *[]AlertingV3DestinationPagerdutyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3DestinationPagerdutyGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Links PagerDuty with the account using the integration token.
func (r *AlertingV3DestinationPagerdutyService) Link(ctx context.Context, tokenID string, query AlertingV3DestinationPagerdutyLinkParams, opts ...option.RequestOption) (res *AlertingV3DestinationPagerdutyLinkResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3DestinationPagerdutyLinkResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty/connect/%s", query.AccountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AlertingV3DestinationPagerdutyNewResponse struct {
	// UUID
	ID   string                                        `json:"id"`
	JSON alertingV3DestinationPagerdutyNewResponseJSON `json:"-"`
}

// alertingV3DestinationPagerdutyNewResponseJSON contains the JSON metadata for the
// struct [AlertingV3DestinationPagerdutyNewResponse]
type alertingV3DestinationPagerdutyNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AlertingV3DestinationPagerdutyDeleteResponseUnknown],
// [AlertingV3DestinationPagerdutyDeleteResponseArray] or [shared.UnionString].
type AlertingV3DestinationPagerdutyDeleteResponse interface {
	ImplementsAlertingV3DestinationPagerdutyDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AlertingV3DestinationPagerdutyDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AlertingV3DestinationPagerdutyDeleteResponseArray []interface{}

func (r AlertingV3DestinationPagerdutyDeleteResponseArray) ImplementsAlertingV3DestinationPagerdutyDeleteResponse() {
}

type AlertingV3DestinationPagerdutyGetResponse struct {
	// UUID
	ID string `json:"id"`
	// The name of the pagerduty service.
	Name string                                        `json:"name"`
	JSON alertingV3DestinationPagerdutyGetResponseJSON `json:"-"`
}

// alertingV3DestinationPagerdutyGetResponseJSON contains the JSON metadata for the
// struct [AlertingV3DestinationPagerdutyGetResponse]
type alertingV3DestinationPagerdutyGetResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyLinkResponse struct {
	// UUID
	ID   string                                         `json:"id"`
	JSON alertingV3DestinationPagerdutyLinkResponseJSON `json:"-"`
}

// alertingV3DestinationPagerdutyLinkResponseJSON contains the JSON metadata for
// the struct [AlertingV3DestinationPagerdutyLinkResponse]
type alertingV3DestinationPagerdutyLinkResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyNewParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type AlertingV3DestinationPagerdutyNewResponseEnvelope struct {
	Errors   []AlertingV3DestinationPagerdutyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3DestinationPagerdutyNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3DestinationPagerdutyNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AlertingV3DestinationPagerdutyNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    alertingV3DestinationPagerdutyNewResponseEnvelopeJSON    `json:"-"`
}

// alertingV3DestinationPagerdutyNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [AlertingV3DestinationPagerdutyNewResponseEnvelope]
type alertingV3DestinationPagerdutyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyNewResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    alertingV3DestinationPagerdutyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3DestinationPagerdutyNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationPagerdutyNewResponseEnvelopeErrors]
type alertingV3DestinationPagerdutyNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyNewResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    alertingV3DestinationPagerdutyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3DestinationPagerdutyNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationPagerdutyNewResponseEnvelopeMessages]
type alertingV3DestinationPagerdutyNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3DestinationPagerdutyNewResponseEnvelopeSuccess bool

const (
	AlertingV3DestinationPagerdutyNewResponseEnvelopeSuccessTrue AlertingV3DestinationPagerdutyNewResponseEnvelopeSuccess = true
)

type AlertingV3DestinationPagerdutyDeleteParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type AlertingV3DestinationPagerdutyDeleteResponseEnvelope struct {
	Errors   []AlertingV3DestinationPagerdutyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3DestinationPagerdutyDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3DestinationPagerdutyDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AlertingV3DestinationPagerdutyDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AlertingV3DestinationPagerdutyDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       alertingV3DestinationPagerdutyDeleteResponseEnvelopeJSON       `json:"-"`
}

// alertingV3DestinationPagerdutyDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [AlertingV3DestinationPagerdutyDeleteResponseEnvelope]
type alertingV3DestinationPagerdutyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyDeleteResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    alertingV3DestinationPagerdutyDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3DestinationPagerdutyDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationPagerdutyDeleteResponseEnvelopeErrors]
type alertingV3DestinationPagerdutyDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyDeleteResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    alertingV3DestinationPagerdutyDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3DestinationPagerdutyDeleteResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [AlertingV3DestinationPagerdutyDeleteResponseEnvelopeMessages]
type alertingV3DestinationPagerdutyDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3DestinationPagerdutyDeleteResponseEnvelopeSuccess bool

const (
	AlertingV3DestinationPagerdutyDeleteResponseEnvelopeSuccessTrue AlertingV3DestinationPagerdutyDeleteResponseEnvelopeSuccess = true
)

type AlertingV3DestinationPagerdutyDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                            `json:"total_count"`
	JSON       alertingV3DestinationPagerdutyDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// alertingV3DestinationPagerdutyDeleteResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [AlertingV3DestinationPagerdutyDeleteResponseEnvelopeResultInfo]
type alertingV3DestinationPagerdutyDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type AlertingV3DestinationPagerdutyGetResponseEnvelope struct {
	Errors   []AlertingV3DestinationPagerdutyGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3DestinationPagerdutyGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []AlertingV3DestinationPagerdutyGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AlertingV3DestinationPagerdutyGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AlertingV3DestinationPagerdutyGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       alertingV3DestinationPagerdutyGetResponseEnvelopeJSON       `json:"-"`
}

// alertingV3DestinationPagerdutyGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [AlertingV3DestinationPagerdutyGetResponseEnvelope]
type alertingV3DestinationPagerdutyGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyGetResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    alertingV3DestinationPagerdutyGetResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3DestinationPagerdutyGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationPagerdutyGetResponseEnvelopeErrors]
type alertingV3DestinationPagerdutyGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyGetResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    alertingV3DestinationPagerdutyGetResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3DestinationPagerdutyGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationPagerdutyGetResponseEnvelopeMessages]
type alertingV3DestinationPagerdutyGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3DestinationPagerdutyGetResponseEnvelopeSuccess bool

const (
	AlertingV3DestinationPagerdutyGetResponseEnvelopeSuccessTrue AlertingV3DestinationPagerdutyGetResponseEnvelopeSuccess = true
)

type AlertingV3DestinationPagerdutyGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                         `json:"total_count"`
	JSON       alertingV3DestinationPagerdutyGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// alertingV3DestinationPagerdutyGetResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [AlertingV3DestinationPagerdutyGetResponseEnvelopeResultInfo]
type alertingV3DestinationPagerdutyGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyLinkParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type AlertingV3DestinationPagerdutyLinkResponseEnvelope struct {
	Errors   []AlertingV3DestinationPagerdutyLinkResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3DestinationPagerdutyLinkResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3DestinationPagerdutyLinkResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AlertingV3DestinationPagerdutyLinkResponseEnvelopeSuccess `json:"success,required"`
	JSON    alertingV3DestinationPagerdutyLinkResponseEnvelopeJSON    `json:"-"`
}

// alertingV3DestinationPagerdutyLinkResponseEnvelopeJSON contains the JSON
// metadata for the struct [AlertingV3DestinationPagerdutyLinkResponseEnvelope]
type alertingV3DestinationPagerdutyLinkResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyLinkResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyLinkResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    alertingV3DestinationPagerdutyLinkResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3DestinationPagerdutyLinkResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationPagerdutyLinkResponseEnvelopeErrors]
type alertingV3DestinationPagerdutyLinkResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyLinkResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationPagerdutyLinkResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    alertingV3DestinationPagerdutyLinkResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3DestinationPagerdutyLinkResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationPagerdutyLinkResponseEnvelopeMessages]
type alertingV3DestinationPagerdutyLinkResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationPagerdutyLinkResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3DestinationPagerdutyLinkResponseEnvelopeSuccess bool

const (
	AlertingV3DestinationPagerdutyLinkResponseEnvelopeSuccessTrue AlertingV3DestinationPagerdutyLinkResponseEnvelopeSuccess = true
)

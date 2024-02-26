// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AlertingV3DestinationWebhookService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAlertingV3DestinationWebhookService] method instead.
type AlertingV3DestinationWebhookService struct {
	Options []option.RequestOption
}

// NewAlertingV3DestinationWebhookService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAlertingV3DestinationWebhookService(opts ...option.RequestOption) (r *AlertingV3DestinationWebhookService) {
	r = &AlertingV3DestinationWebhookService{}
	r.Options = opts
	return
}

// Creates a new webhook destination.
func (r *AlertingV3DestinationWebhookService) New(ctx context.Context, params AlertingV3DestinationWebhookNewParams, opts ...option.RequestOption) (res *AlertingV3DestinationWebhookNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3DestinationWebhookNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a webhook destination.
func (r *AlertingV3DestinationWebhookService) Update(ctx context.Context, webhookID string, params AlertingV3DestinationWebhookUpdateParams, opts ...option.RequestOption) (res *AlertingV3DestinationWebhookUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3DestinationWebhookUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", params.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets a list of all configured webhook destinations.
func (r *AlertingV3DestinationWebhookService) List(ctx context.Context, query AlertingV3DestinationWebhookListParams, opts ...option.RequestOption) (res *[]AlertingV3DestinationWebhookListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3DestinationWebhookListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured webhook destination.
func (r *AlertingV3DestinationWebhookService) Delete(ctx context.Context, webhookID string, body AlertingV3DestinationWebhookDeleteParams, opts ...option.RequestOption) (res *AlertingV3DestinationWebhookDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3DestinationWebhookDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", body.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details for a single webhooks destination.
func (r *AlertingV3DestinationWebhookService) Get(ctx context.Context, webhookID string, query AlertingV3DestinationWebhookGetParams, opts ...option.RequestOption) (res *AlertingV3DestinationWebhookGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AlertingV3DestinationWebhookGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", query.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AlertingV3DestinationWebhookNewResponse struct {
	// UUID
	ID   string                                      `json:"id"`
	JSON alertingV3DestinationWebhookNewResponseJSON `json:"-"`
}

// alertingV3DestinationWebhookNewResponseJSON contains the JSON metadata for the
// struct [AlertingV3DestinationWebhookNewResponse]
type alertingV3DestinationWebhookNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookUpdateResponse struct {
	// UUID
	ID   string                                         `json:"id"`
	JSON alertingV3DestinationWebhookUpdateResponseJSON `json:"-"`
}

// alertingV3DestinationWebhookUpdateResponseJSON contains the JSON metadata for
// the struct [AlertingV3DestinationWebhookUpdateResponse]
type alertingV3DestinationWebhookUpdateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookListResponse struct {
	// The unique identifier of a webhook
	ID string `json:"id"`
	// Timestamp of when the webhook destination was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of the last time an attempt to dispatch a notification to this webhook
	// failed.
	LastFailure time.Time `json:"last_failure" format:"date-time"`
	// Timestamp of the last time Cloudflare was able to successfully dispatch a
	// notification using this webhook.
	LastSuccess time.Time `json:"last_success" format:"date-time"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name string `json:"name"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret string `json:"secret"`
	// Type of webhook endpoint.
	Type AlertingV3DestinationWebhookListResponseType `json:"type"`
	// The POST endpoint to call when dispatching a notification.
	URL  string                                       `json:"url"`
	JSON alertingV3DestinationWebhookListResponseJSON `json:"-"`
}

// alertingV3DestinationWebhookListResponseJSON contains the JSON metadata for the
// struct [AlertingV3DestinationWebhookListResponse]
type alertingV3DestinationWebhookListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	LastFailure apijson.Field
	LastSuccess apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of webhook endpoint.
type AlertingV3DestinationWebhookListResponseType string

const (
	AlertingV3DestinationWebhookListResponseTypeSlack   AlertingV3DestinationWebhookListResponseType = "slack"
	AlertingV3DestinationWebhookListResponseTypeGeneric AlertingV3DestinationWebhookListResponseType = "generic"
	AlertingV3DestinationWebhookListResponseTypeGchat   AlertingV3DestinationWebhookListResponseType = "gchat"
)

// Union satisfied by [AlertingV3DestinationWebhookDeleteResponseUnknown],
// [AlertingV3DestinationWebhookDeleteResponseArray] or [shared.UnionString].
type AlertingV3DestinationWebhookDeleteResponse interface {
	ImplementsAlertingV3DestinationWebhookDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AlertingV3DestinationWebhookDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AlertingV3DestinationWebhookDeleteResponseArray []interface{}

func (r AlertingV3DestinationWebhookDeleteResponseArray) ImplementsAlertingV3DestinationWebhookDeleteResponse() {
}

type AlertingV3DestinationWebhookGetResponse struct {
	// The unique identifier of a webhook
	ID string `json:"id"`
	// Timestamp of when the webhook destination was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of the last time an attempt to dispatch a notification to this webhook
	// failed.
	LastFailure time.Time `json:"last_failure" format:"date-time"`
	// Timestamp of the last time Cloudflare was able to successfully dispatch a
	// notification using this webhook.
	LastSuccess time.Time `json:"last_success" format:"date-time"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name string `json:"name"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret string `json:"secret"`
	// Type of webhook endpoint.
	Type AlertingV3DestinationWebhookGetResponseType `json:"type"`
	// The POST endpoint to call when dispatching a notification.
	URL  string                                      `json:"url"`
	JSON alertingV3DestinationWebhookGetResponseJSON `json:"-"`
}

// alertingV3DestinationWebhookGetResponseJSON contains the JSON metadata for the
// struct [AlertingV3DestinationWebhookGetResponse]
type alertingV3DestinationWebhookGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	LastFailure apijson.Field
	LastSuccess apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of webhook endpoint.
type AlertingV3DestinationWebhookGetResponseType string

const (
	AlertingV3DestinationWebhookGetResponseTypeSlack   AlertingV3DestinationWebhookGetResponseType = "slack"
	AlertingV3DestinationWebhookGetResponseTypeGeneric AlertingV3DestinationWebhookGetResponseType = "generic"
	AlertingV3DestinationWebhookGetResponseTypeGchat   AlertingV3DestinationWebhookGetResponseType = "gchat"
)

type AlertingV3DestinationWebhookNewParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name param.Field[string] `json:"name,required"`
	// The POST endpoint to call when dispatching a notification.
	URL param.Field[string] `json:"url,required"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret param.Field[string] `json:"secret"`
}

func (r AlertingV3DestinationWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertingV3DestinationWebhookNewResponseEnvelope struct {
	Errors   []AlertingV3DestinationWebhookNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3DestinationWebhookNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3DestinationWebhookNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AlertingV3DestinationWebhookNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    alertingV3DestinationWebhookNewResponseEnvelopeJSON    `json:"-"`
}

// alertingV3DestinationWebhookNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [AlertingV3DestinationWebhookNewResponseEnvelope]
type alertingV3DestinationWebhookNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookNewResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    alertingV3DestinationWebhookNewResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3DestinationWebhookNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AlertingV3DestinationWebhookNewResponseEnvelopeErrors]
type alertingV3DestinationWebhookNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookNewResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    alertingV3DestinationWebhookNewResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3DestinationWebhookNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationWebhookNewResponseEnvelopeMessages]
type alertingV3DestinationWebhookNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3DestinationWebhookNewResponseEnvelopeSuccess bool

const (
	AlertingV3DestinationWebhookNewResponseEnvelopeSuccessTrue AlertingV3DestinationWebhookNewResponseEnvelopeSuccess = true
)

type AlertingV3DestinationWebhookUpdateParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the webhook destination. This will be included in the request body
	// when you receive a webhook notification.
	Name param.Field[string] `json:"name,required"`
	// The POST endpoint to call when dispatching a notification.
	URL param.Field[string] `json:"url,required"`
	// Optional secret that will be passed in the `cf-webhook-auth` header when
	// dispatching generic webhook notifications or formatted for supported
	// destinations. Secrets are not returned in any API response body.
	Secret param.Field[string] `json:"secret"`
}

func (r AlertingV3DestinationWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AlertingV3DestinationWebhookUpdateResponseEnvelope struct {
	Errors   []AlertingV3DestinationWebhookUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3DestinationWebhookUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3DestinationWebhookUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AlertingV3DestinationWebhookUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    alertingV3DestinationWebhookUpdateResponseEnvelopeJSON    `json:"-"`
}

// alertingV3DestinationWebhookUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [AlertingV3DestinationWebhookUpdateResponseEnvelope]
type alertingV3DestinationWebhookUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookUpdateResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    alertingV3DestinationWebhookUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3DestinationWebhookUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationWebhookUpdateResponseEnvelopeErrors]
type alertingV3DestinationWebhookUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookUpdateResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    alertingV3DestinationWebhookUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3DestinationWebhookUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationWebhookUpdateResponseEnvelopeMessages]
type alertingV3DestinationWebhookUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3DestinationWebhookUpdateResponseEnvelopeSuccess bool

const (
	AlertingV3DestinationWebhookUpdateResponseEnvelopeSuccessTrue AlertingV3DestinationWebhookUpdateResponseEnvelopeSuccess = true
)

type AlertingV3DestinationWebhookListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type AlertingV3DestinationWebhookListResponseEnvelope struct {
	Errors   []AlertingV3DestinationWebhookListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3DestinationWebhookListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AlertingV3DestinationWebhookListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AlertingV3DestinationWebhookListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AlertingV3DestinationWebhookListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       alertingV3DestinationWebhookListResponseEnvelopeJSON       `json:"-"`
}

// alertingV3DestinationWebhookListResponseEnvelopeJSON contains the JSON metadata
// for the struct [AlertingV3DestinationWebhookListResponseEnvelope]
type alertingV3DestinationWebhookListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookListResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    alertingV3DestinationWebhookListResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3DestinationWebhookListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AlertingV3DestinationWebhookListResponseEnvelopeErrors]
type alertingV3DestinationWebhookListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookListResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    alertingV3DestinationWebhookListResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3DestinationWebhookListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationWebhookListResponseEnvelopeMessages]
type alertingV3DestinationWebhookListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3DestinationWebhookListResponseEnvelopeSuccess bool

const (
	AlertingV3DestinationWebhookListResponseEnvelopeSuccessTrue AlertingV3DestinationWebhookListResponseEnvelopeSuccess = true
)

type AlertingV3DestinationWebhookListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                        `json:"total_count"`
	JSON       alertingV3DestinationWebhookListResponseEnvelopeResultInfoJSON `json:"-"`
}

// alertingV3DestinationWebhookListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationWebhookListResponseEnvelopeResultInfo]
type alertingV3DestinationWebhookListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookDeleteParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type AlertingV3DestinationWebhookDeleteResponseEnvelope struct {
	Errors   []AlertingV3DestinationWebhookDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3DestinationWebhookDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3DestinationWebhookDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AlertingV3DestinationWebhookDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AlertingV3DestinationWebhookDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       alertingV3DestinationWebhookDeleteResponseEnvelopeJSON       `json:"-"`
}

// alertingV3DestinationWebhookDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [AlertingV3DestinationWebhookDeleteResponseEnvelope]
type alertingV3DestinationWebhookDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookDeleteResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    alertingV3DestinationWebhookDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3DestinationWebhookDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationWebhookDeleteResponseEnvelopeErrors]
type alertingV3DestinationWebhookDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookDeleteResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    alertingV3DestinationWebhookDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3DestinationWebhookDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationWebhookDeleteResponseEnvelopeMessages]
type alertingV3DestinationWebhookDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3DestinationWebhookDeleteResponseEnvelopeSuccess bool

const (
	AlertingV3DestinationWebhookDeleteResponseEnvelopeSuccessTrue AlertingV3DestinationWebhookDeleteResponseEnvelopeSuccess = true
)

type AlertingV3DestinationWebhookDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                          `json:"total_count"`
	JSON       alertingV3DestinationWebhookDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// alertingV3DestinationWebhookDeleteResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [AlertingV3DestinationWebhookDeleteResponseEnvelopeResultInfo]
type alertingV3DestinationWebhookDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type AlertingV3DestinationWebhookGetResponseEnvelope struct {
	Errors   []AlertingV3DestinationWebhookGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AlertingV3DestinationWebhookGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AlertingV3DestinationWebhookGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AlertingV3DestinationWebhookGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    alertingV3DestinationWebhookGetResponseEnvelopeJSON    `json:"-"`
}

// alertingV3DestinationWebhookGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [AlertingV3DestinationWebhookGetResponseEnvelope]
type alertingV3DestinationWebhookGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    alertingV3DestinationWebhookGetResponseEnvelopeErrorsJSON `json:"-"`
}

// alertingV3DestinationWebhookGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AlertingV3DestinationWebhookGetResponseEnvelopeErrors]
type alertingV3DestinationWebhookGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AlertingV3DestinationWebhookGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    alertingV3DestinationWebhookGetResponseEnvelopeMessagesJSON `json:"-"`
}

// alertingV3DestinationWebhookGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [AlertingV3DestinationWebhookGetResponseEnvelopeMessages]
type alertingV3DestinationWebhookGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AlertingV3DestinationWebhookGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AlertingV3DestinationWebhookGetResponseEnvelopeSuccess bool

const (
	AlertingV3DestinationWebhookGetResponseEnvelopeSuccessTrue AlertingV3DestinationWebhookGetResponseEnvelopeSuccess = true
)

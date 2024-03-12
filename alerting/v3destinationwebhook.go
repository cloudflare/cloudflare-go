// File generated from our OpenAPI spec by Stainless.

package alerting

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// V3DestinationWebhookService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewV3DestinationWebhookService]
// method instead.
type V3DestinationWebhookService struct {
	Options []option.RequestOption
}

// NewV3DestinationWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV3DestinationWebhookService(opts ...option.RequestOption) (r *V3DestinationWebhookService) {
	r = &V3DestinationWebhookService{}
	r.Options = opts
	return
}

// Creates a new webhook destination.
func (r *V3DestinationWebhookService) New(ctx context.Context, params V3DestinationWebhookNewParams, opts ...option.RequestOption) (res *V3DestinationWebhookNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3DestinationWebhookNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a webhook destination.
func (r *V3DestinationWebhookService) Update(ctx context.Context, webhookID string, params V3DestinationWebhookUpdateParams, opts ...option.RequestOption) (res *V3DestinationWebhookUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3DestinationWebhookUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", params.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets a list of all configured webhook destinations.
func (r *V3DestinationWebhookService) List(ctx context.Context, query V3DestinationWebhookListParams, opts ...option.RequestOption) (res *[]V3DestinationWebhookListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3DestinationWebhookListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a configured webhook destination.
func (r *V3DestinationWebhookService) Delete(ctx context.Context, webhookID string, body V3DestinationWebhookDeleteParams, opts ...option.RequestOption) (res *V3DestinationWebhookDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3DestinationWebhookDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", body.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details for a single webhooks destination.
func (r *V3DestinationWebhookService) Get(ctx context.Context, webhookID string, query V3DestinationWebhookGetParams, opts ...option.RequestOption) (res *V3DestinationWebhookGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env V3DestinationWebhookGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/webhooks/%s", query.AccountID, webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type V3DestinationWebhookNewResponse struct {
	// UUID
	ID   string                              `json:"id"`
	JSON v3DestinationWebhookNewResponseJSON `json:"-"`
}

// v3DestinationWebhookNewResponseJSON contains the JSON metadata for the struct
// [V3DestinationWebhookNewResponse]
type v3DestinationWebhookNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookNewResponseJSON) RawJSON() string {
	return r.raw
}

type V3DestinationWebhookUpdateResponse struct {
	// UUID
	ID   string                                 `json:"id"`
	JSON v3DestinationWebhookUpdateResponseJSON `json:"-"`
}

// v3DestinationWebhookUpdateResponseJSON contains the JSON metadata for the struct
// [V3DestinationWebhookUpdateResponse]
type v3DestinationWebhookUpdateResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type V3DestinationWebhookListResponse struct {
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
	Type V3DestinationWebhookListResponseType `json:"type"`
	// The POST endpoint to call when dispatching a notification.
	URL  string                               `json:"url"`
	JSON v3DestinationWebhookListResponseJSON `json:"-"`
}

// v3DestinationWebhookListResponseJSON contains the JSON metadata for the struct
// [V3DestinationWebhookListResponse]
type v3DestinationWebhookListResponseJSON struct {
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

func (r *V3DestinationWebhookListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookListResponseJSON) RawJSON() string {
	return r.raw
}

// Type of webhook endpoint.
type V3DestinationWebhookListResponseType string

const (
	V3DestinationWebhookListResponseTypeSlack   V3DestinationWebhookListResponseType = "slack"
	V3DestinationWebhookListResponseTypeGeneric V3DestinationWebhookListResponseType = "generic"
	V3DestinationWebhookListResponseTypeGchat   V3DestinationWebhookListResponseType = "gchat"
)

// Union satisfied by [alerting.V3DestinationWebhookDeleteResponseUnknown],
// [alerting.V3DestinationWebhookDeleteResponseArray] or [shared.UnionString].
type V3DestinationWebhookDeleteResponse interface {
	ImplementsAlertingV3DestinationWebhookDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*V3DestinationWebhookDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(V3DestinationWebhookDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type V3DestinationWebhookDeleteResponseArray []interface{}

func (r V3DestinationWebhookDeleteResponseArray) ImplementsAlertingV3DestinationWebhookDeleteResponse() {
}

type V3DestinationWebhookGetResponse struct {
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
	Type V3DestinationWebhookGetResponseType `json:"type"`
	// The POST endpoint to call when dispatching a notification.
	URL  string                              `json:"url"`
	JSON v3DestinationWebhookGetResponseJSON `json:"-"`
}

// v3DestinationWebhookGetResponseJSON contains the JSON metadata for the struct
// [V3DestinationWebhookGetResponse]
type v3DestinationWebhookGetResponseJSON struct {
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

func (r *V3DestinationWebhookGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookGetResponseJSON) RawJSON() string {
	return r.raw
}

// Type of webhook endpoint.
type V3DestinationWebhookGetResponseType string

const (
	V3DestinationWebhookGetResponseTypeSlack   V3DestinationWebhookGetResponseType = "slack"
	V3DestinationWebhookGetResponseTypeGeneric V3DestinationWebhookGetResponseType = "generic"
	V3DestinationWebhookGetResponseTypeGchat   V3DestinationWebhookGetResponseType = "gchat"
)

type V3DestinationWebhookNewParams struct {
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

func (r V3DestinationWebhookNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V3DestinationWebhookNewResponseEnvelope struct {
	Errors   []V3DestinationWebhookNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3DestinationWebhookNewResponseEnvelopeMessages `json:"messages,required"`
	Result   V3DestinationWebhookNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V3DestinationWebhookNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    v3DestinationWebhookNewResponseEnvelopeJSON    `json:"-"`
}

// v3DestinationWebhookNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [V3DestinationWebhookNewResponseEnvelope]
type v3DestinationWebhookNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3DestinationWebhookNewResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    v3DestinationWebhookNewResponseEnvelopeErrorsJSON `json:"-"`
}

// v3DestinationWebhookNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [V3DestinationWebhookNewResponseEnvelopeErrors]
type v3DestinationWebhookNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3DestinationWebhookNewResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    v3DestinationWebhookNewResponseEnvelopeMessagesJSON `json:"-"`
}

// v3DestinationWebhookNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [V3DestinationWebhookNewResponseEnvelopeMessages]
type v3DestinationWebhookNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3DestinationWebhookNewResponseEnvelopeSuccess bool

const (
	V3DestinationWebhookNewResponseEnvelopeSuccessTrue V3DestinationWebhookNewResponseEnvelopeSuccess = true
)

type V3DestinationWebhookUpdateParams struct {
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

func (r V3DestinationWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type V3DestinationWebhookUpdateResponseEnvelope struct {
	Errors   []V3DestinationWebhookUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3DestinationWebhookUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   V3DestinationWebhookUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V3DestinationWebhookUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    v3DestinationWebhookUpdateResponseEnvelopeJSON    `json:"-"`
}

// v3DestinationWebhookUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [V3DestinationWebhookUpdateResponseEnvelope]
type v3DestinationWebhookUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3DestinationWebhookUpdateResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    v3DestinationWebhookUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// v3DestinationWebhookUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [V3DestinationWebhookUpdateResponseEnvelopeErrors]
type v3DestinationWebhookUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3DestinationWebhookUpdateResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    v3DestinationWebhookUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// v3DestinationWebhookUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [V3DestinationWebhookUpdateResponseEnvelopeMessages]
type v3DestinationWebhookUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3DestinationWebhookUpdateResponseEnvelopeSuccess bool

const (
	V3DestinationWebhookUpdateResponseEnvelopeSuccessTrue V3DestinationWebhookUpdateResponseEnvelopeSuccess = true
)

type V3DestinationWebhookListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type V3DestinationWebhookListResponseEnvelope struct {
	Errors   []V3DestinationWebhookListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3DestinationWebhookListResponseEnvelopeMessages `json:"messages,required"`
	Result   []V3DestinationWebhookListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    V3DestinationWebhookListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo V3DestinationWebhookListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       v3DestinationWebhookListResponseEnvelopeJSON       `json:"-"`
}

// v3DestinationWebhookListResponseEnvelopeJSON contains the JSON metadata for the
// struct [V3DestinationWebhookListResponseEnvelope]
type v3DestinationWebhookListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3DestinationWebhookListResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    v3DestinationWebhookListResponseEnvelopeErrorsJSON `json:"-"`
}

// v3DestinationWebhookListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [V3DestinationWebhookListResponseEnvelopeErrors]
type v3DestinationWebhookListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3DestinationWebhookListResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    v3DestinationWebhookListResponseEnvelopeMessagesJSON `json:"-"`
}

// v3DestinationWebhookListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [V3DestinationWebhookListResponseEnvelopeMessages]
type v3DestinationWebhookListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3DestinationWebhookListResponseEnvelopeSuccess bool

const (
	V3DestinationWebhookListResponseEnvelopeSuccessTrue V3DestinationWebhookListResponseEnvelopeSuccess = true
)

type V3DestinationWebhookListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                `json:"total_count"`
	JSON       v3DestinationWebhookListResponseEnvelopeResultInfoJSON `json:"-"`
}

// v3DestinationWebhookListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [V3DestinationWebhookListResponseEnvelopeResultInfo]
type v3DestinationWebhookListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type V3DestinationWebhookDeleteParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type V3DestinationWebhookDeleteResponseEnvelope struct {
	Errors   []V3DestinationWebhookDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3DestinationWebhookDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   V3DestinationWebhookDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    V3DestinationWebhookDeleteResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo V3DestinationWebhookDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       v3DestinationWebhookDeleteResponseEnvelopeJSON       `json:"-"`
}

// v3DestinationWebhookDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [V3DestinationWebhookDeleteResponseEnvelope]
type v3DestinationWebhookDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3DestinationWebhookDeleteResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    v3DestinationWebhookDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// v3DestinationWebhookDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [V3DestinationWebhookDeleteResponseEnvelopeErrors]
type v3DestinationWebhookDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3DestinationWebhookDeleteResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    v3DestinationWebhookDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// v3DestinationWebhookDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [V3DestinationWebhookDeleteResponseEnvelopeMessages]
type v3DestinationWebhookDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3DestinationWebhookDeleteResponseEnvelopeSuccess bool

const (
	V3DestinationWebhookDeleteResponseEnvelopeSuccessTrue V3DestinationWebhookDeleteResponseEnvelopeSuccess = true
)

type V3DestinationWebhookDeleteResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                  `json:"total_count"`
	JSON       v3DestinationWebhookDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// v3DestinationWebhookDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [V3DestinationWebhookDeleteResponseEnvelopeResultInfo]
type v3DestinationWebhookDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type V3DestinationWebhookGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type V3DestinationWebhookGetResponseEnvelope struct {
	Errors   []V3DestinationWebhookGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []V3DestinationWebhookGetResponseEnvelopeMessages `json:"messages,required"`
	Result   V3DestinationWebhookGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success V3DestinationWebhookGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    v3DestinationWebhookGetResponseEnvelopeJSON    `json:"-"`
}

// v3DestinationWebhookGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [V3DestinationWebhookGetResponseEnvelope]
type v3DestinationWebhookGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type V3DestinationWebhookGetResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    v3DestinationWebhookGetResponseEnvelopeErrorsJSON `json:"-"`
}

// v3DestinationWebhookGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [V3DestinationWebhookGetResponseEnvelopeErrors]
type v3DestinationWebhookGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type V3DestinationWebhookGetResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    v3DestinationWebhookGetResponseEnvelopeMessagesJSON `json:"-"`
}

// v3DestinationWebhookGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [V3DestinationWebhookGetResponseEnvelopeMessages]
type v3DestinationWebhookGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *V3DestinationWebhookGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r v3DestinationWebhookGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type V3DestinationWebhookGetResponseEnvelopeSuccess bool

const (
	V3DestinationWebhookGetResponseEnvelopeSuccessTrue V3DestinationWebhookGetResponseEnvelopeSuccess = true
)

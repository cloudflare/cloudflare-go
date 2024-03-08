// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// StreamWebhookService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStreamWebhookService] method
// instead.
type StreamWebhookService struct {
	Options []option.RequestOption
}

// NewStreamWebhookService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStreamWebhookService(opts ...option.RequestOption) (r *StreamWebhookService) {
	r = &StreamWebhookService{}
	r.Options = opts
	return
}

// Creates a webhook notification.
func (r *StreamWebhookService) Update(ctx context.Context, params StreamWebhookUpdateParams, opts ...option.RequestOption) (res *StreamWebhookUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWebhookUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/webhook", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a webhook.
func (r *StreamWebhookService) Delete(ctx context.Context, body StreamWebhookDeleteParams, opts ...option.RequestOption) (res *StreamWebhookDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWebhookDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/webhook", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a list of webhooks.
func (r *StreamWebhookService) Get(ctx context.Context, query StreamWebhookGetParams, opts ...option.RequestOption) (res *StreamWebhookGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWebhookGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/webhook", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [StreamWebhookUpdateResponseUnknown] or [shared.UnionString].
type StreamWebhookUpdateResponse interface {
	ImplementsStreamWebhookUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamWebhookUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [StreamWebhookDeleteResponseUnknown] or [shared.UnionString].
type StreamWebhookDeleteResponse interface {
	ImplementsStreamWebhookDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamWebhookDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [StreamWebhookGetResponseUnknown] or [shared.UnionString].
type StreamWebhookGetResponse interface {
	ImplementsStreamWebhookGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamWebhookGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type StreamWebhookUpdateParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The URL where webhooks will be sent.
	NotificationURL param.Field[string] `json:"notificationUrl,required" format:"uri"`
}

func (r StreamWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamWebhookUpdateResponseEnvelope struct {
	Errors   []StreamWebhookUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamWebhookUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamWebhookUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamWebhookUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamWebhookUpdateResponseEnvelopeJSON    `json:"-"`
}

// streamWebhookUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamWebhookUpdateResponseEnvelope]
type streamWebhookUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWebhookUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamWebhookUpdateResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    streamWebhookUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// streamWebhookUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamWebhookUpdateResponseEnvelopeErrors]
type streamWebhookUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWebhookUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamWebhookUpdateResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    streamWebhookUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// streamWebhookUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamWebhookUpdateResponseEnvelopeMessages]
type streamWebhookUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWebhookUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamWebhookUpdateResponseEnvelopeSuccess bool

const (
	StreamWebhookUpdateResponseEnvelopeSuccessTrue StreamWebhookUpdateResponseEnvelopeSuccess = true
)

type StreamWebhookDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamWebhookDeleteResponseEnvelope struct {
	Errors   []StreamWebhookDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamWebhookDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamWebhookDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamWebhookDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamWebhookDeleteResponseEnvelopeJSON    `json:"-"`
}

// streamWebhookDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamWebhookDeleteResponseEnvelope]
type streamWebhookDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWebhookDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamWebhookDeleteResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    streamWebhookDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// streamWebhookDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamWebhookDeleteResponseEnvelopeErrors]
type streamWebhookDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWebhookDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamWebhookDeleteResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    streamWebhookDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// streamWebhookDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamWebhookDeleteResponseEnvelopeMessages]
type streamWebhookDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWebhookDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamWebhookDeleteResponseEnvelopeSuccess bool

const (
	StreamWebhookDeleteResponseEnvelopeSuccessTrue StreamWebhookDeleteResponseEnvelopeSuccess = true
)

type StreamWebhookGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type StreamWebhookGetResponseEnvelope struct {
	Errors   []StreamWebhookGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamWebhookGetResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamWebhookGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamWebhookGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamWebhookGetResponseEnvelopeJSON    `json:"-"`
}

// streamWebhookGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [StreamWebhookGetResponseEnvelope]
type streamWebhookGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWebhookGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type StreamWebhookGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    streamWebhookGetResponseEnvelopeErrorsJSON `json:"-"`
}

// streamWebhookGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [StreamWebhookGetResponseEnvelopeErrors]
type streamWebhookGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWebhookGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type StreamWebhookGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    streamWebhookGetResponseEnvelopeMessagesJSON `json:"-"`
}

// streamWebhookGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [StreamWebhookGetResponseEnvelopeMessages]
type streamWebhookGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r streamWebhookGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type StreamWebhookGetResponseEnvelopeSuccess bool

const (
	StreamWebhookGetResponseEnvelopeSuccessTrue StreamWebhookGetResponseEnvelopeSuccess = true
)

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

// Deletes a webhook.
func (r *StreamWebhookService) Delete(ctx context.Context, accountID string, opts ...option.RequestOption) (res *StreamWebhookDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWebhookDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a webhook notification.
func (r *StreamWebhookService) StreamWebhookNewWebhooks(ctx context.Context, accountID string, body StreamWebhookStreamWebhookNewWebhooksParams, opts ...option.RequestOption) (res *StreamWebhookStreamWebhookNewWebhooksResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWebhookStreamWebhookNewWebhooksResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a list of webhooks.
func (r *StreamWebhookService) StreamWebhookViewWebhooks(ctx context.Context, accountID string, opts ...option.RequestOption) (res *StreamWebhookStreamWebhookViewWebhooksResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWebhookStreamWebhookViewWebhooksResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

// Union satisfied by [StreamWebhookStreamWebhookNewWebhooksResponseUnknown] or
// [shared.UnionString].
type StreamWebhookStreamWebhookNewWebhooksResponse interface {
	ImplementsStreamWebhookStreamWebhookNewWebhooksResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamWebhookStreamWebhookNewWebhooksResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [StreamWebhookStreamWebhookViewWebhooksResponseUnknown] or
// [shared.UnionString].
type StreamWebhookStreamWebhookViewWebhooksResponse interface {
	ImplementsStreamWebhookStreamWebhookViewWebhooksResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamWebhookStreamWebhookViewWebhooksResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
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

// Whether the API call was successful
type StreamWebhookDeleteResponseEnvelopeSuccess bool

const (
	StreamWebhookDeleteResponseEnvelopeSuccessTrue StreamWebhookDeleteResponseEnvelopeSuccess = true
)

type StreamWebhookStreamWebhookNewWebhooksParams struct {
	// The URL where webhooks will be sent.
	NotificationURL param.Field[string] `json:"notificationUrl,required" format:"uri"`
}

func (r StreamWebhookStreamWebhookNewWebhooksParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamWebhookStreamWebhookNewWebhooksResponseEnvelope struct {
	Errors   []StreamWebhookStreamWebhookNewWebhooksResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamWebhookStreamWebhookNewWebhooksResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamWebhookStreamWebhookNewWebhooksResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamWebhookStreamWebhookNewWebhooksResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamWebhookStreamWebhookNewWebhooksResponseEnvelopeJSON    `json:"-"`
}

// streamWebhookStreamWebhookNewWebhooksResponseEnvelopeJSON contains the JSON
// metadata for the struct [StreamWebhookStreamWebhookNewWebhooksResponseEnvelope]
type streamWebhookStreamWebhookNewWebhooksResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookStreamWebhookNewWebhooksResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWebhookStreamWebhookNewWebhooksResponseEnvelopeErrors struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    streamWebhookStreamWebhookNewWebhooksResponseEnvelopeErrorsJSON `json:"-"`
}

// streamWebhookStreamWebhookNewWebhooksResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [StreamWebhookStreamWebhookNewWebhooksResponseEnvelopeErrors]
type streamWebhookStreamWebhookNewWebhooksResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookStreamWebhookNewWebhooksResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWebhookStreamWebhookNewWebhooksResponseEnvelopeMessages struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    streamWebhookStreamWebhookNewWebhooksResponseEnvelopeMessagesJSON `json:"-"`
}

// streamWebhookStreamWebhookNewWebhooksResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [StreamWebhookStreamWebhookNewWebhooksResponseEnvelopeMessages]
type streamWebhookStreamWebhookNewWebhooksResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookStreamWebhookNewWebhooksResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamWebhookStreamWebhookNewWebhooksResponseEnvelopeSuccess bool

const (
	StreamWebhookStreamWebhookNewWebhooksResponseEnvelopeSuccessTrue StreamWebhookStreamWebhookNewWebhooksResponseEnvelopeSuccess = true
)

type StreamWebhookStreamWebhookViewWebhooksResponseEnvelope struct {
	Errors   []StreamWebhookStreamWebhookViewWebhooksResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamWebhookStreamWebhookViewWebhooksResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamWebhookStreamWebhookViewWebhooksResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamWebhookStreamWebhookViewWebhooksResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamWebhookStreamWebhookViewWebhooksResponseEnvelopeJSON    `json:"-"`
}

// streamWebhookStreamWebhookViewWebhooksResponseEnvelopeJSON contains the JSON
// metadata for the struct [StreamWebhookStreamWebhookViewWebhooksResponseEnvelope]
type streamWebhookStreamWebhookViewWebhooksResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookStreamWebhookViewWebhooksResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWebhookStreamWebhookViewWebhooksResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    streamWebhookStreamWebhookViewWebhooksResponseEnvelopeErrorsJSON `json:"-"`
}

// streamWebhookStreamWebhookViewWebhooksResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [StreamWebhookStreamWebhookViewWebhooksResponseEnvelopeErrors]
type streamWebhookStreamWebhookViewWebhooksResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookStreamWebhookViewWebhooksResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWebhookStreamWebhookViewWebhooksResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    streamWebhookStreamWebhookViewWebhooksResponseEnvelopeMessagesJSON `json:"-"`
}

// streamWebhookStreamWebhookViewWebhooksResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [StreamWebhookStreamWebhookViewWebhooksResponseEnvelopeMessages]
type streamWebhookStreamWebhookViewWebhooksResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookStreamWebhookViewWebhooksResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamWebhookStreamWebhookViewWebhooksResponseEnvelopeSuccess bool

const (
	StreamWebhookStreamWebhookViewWebhooksResponseEnvelopeSuccessTrue StreamWebhookStreamWebhookViewWebhooksResponseEnvelopeSuccess = true
)

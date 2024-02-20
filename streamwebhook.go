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

// Retrieves a list of webhooks.
func (r *StreamWebhookService) Get(ctx context.Context, accountID string, opts ...option.RequestOption) (res *StreamWebhookGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWebhookGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a webhook notification.
func (r *StreamWebhookService) Replace(ctx context.Context, accountID string, body StreamWebhookReplaceParams, opts ...option.RequestOption) (res *StreamWebhookReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env StreamWebhookReplaceResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/webhook", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
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

// Union satisfied by [StreamWebhookReplaceResponseUnknown] or
// [shared.UnionString].
type StreamWebhookReplaceResponse interface {
	ImplementsStreamWebhookReplaceResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*StreamWebhookReplaceResponse)(nil)).Elem(),
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

// Whether the API call was successful
type StreamWebhookGetResponseEnvelopeSuccess bool

const (
	StreamWebhookGetResponseEnvelopeSuccessTrue StreamWebhookGetResponseEnvelopeSuccess = true
)

type StreamWebhookReplaceParams struct {
	// The URL where webhooks will be sent.
	NotificationURL param.Field[string] `json:"notificationUrl,required" format:"uri"`
}

func (r StreamWebhookReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type StreamWebhookReplaceResponseEnvelope struct {
	Errors   []StreamWebhookReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StreamWebhookReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   StreamWebhookReplaceResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success StreamWebhookReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    streamWebhookReplaceResponseEnvelopeJSON    `json:"-"`
}

// streamWebhookReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [StreamWebhookReplaceResponseEnvelope]
type streamWebhookReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWebhookReplaceResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    streamWebhookReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// streamWebhookReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StreamWebhookReplaceResponseEnvelopeErrors]
type streamWebhookReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StreamWebhookReplaceResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    streamWebhookReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// streamWebhookReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StreamWebhookReplaceResponseEnvelopeMessages]
type streamWebhookReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StreamWebhookReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StreamWebhookReplaceResponseEnvelopeSuccess bool

const (
	StreamWebhookReplaceResponseEnvelopeSuccessTrue StreamWebhookReplaceResponseEnvelopeSuccess = true
)

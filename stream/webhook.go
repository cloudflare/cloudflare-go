// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

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

// WebhookService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r *WebhookService) {
	r = &WebhookService{}
	r.Options = opts
	return
}

// Creates a webhook notification.
func (r *WebhookService) Update(ctx context.Context, params WebhookUpdateParams, opts ...option.RequestOption) (res *WebhookUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WebhookUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/webhook", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a webhook.
func (r *WebhookService) Delete(ctx context.Context, body WebhookDeleteParams, opts ...option.RequestOption) (res *WebhookDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WebhookDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/webhook", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a list of webhooks.
func (r *WebhookService) Get(ctx context.Context, query WebhookGetParams, opts ...option.RequestOption) (res *WebhookGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WebhookGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/stream/webhook", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [stream.WebhookUpdateResponseUnknown] or
// [shared.UnionString].
type WebhookUpdateResponse interface {
	ImplementsStreamWebhookUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WebhookUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [stream.WebhookDeleteResponseUnknown] or
// [shared.UnionString].
type WebhookDeleteResponse interface {
	ImplementsStreamWebhookDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WebhookDeleteResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [stream.WebhookGetResponseUnknown] or [shared.UnionString].
type WebhookGetResponse interface {
	ImplementsStreamWebhookGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WebhookGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type WebhookUpdateParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The URL where webhooks will be sent.
	NotificationURL param.Field[string] `json:"notificationUrl,required" format:"uri"`
}

func (r WebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WebhookUpdateResponseEnvelope struct {
	Errors   []WebhookUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WebhookUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WebhookUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WebhookUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    webhookUpdateResponseEnvelopeJSON    `json:"-"`
}

// webhookUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [WebhookUpdateResponseEnvelope]
type webhookUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WebhookUpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    webhookUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// webhookUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WebhookUpdateResponseEnvelopeErrors]
type webhookUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WebhookUpdateResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    webhookUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// webhookUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WebhookUpdateResponseEnvelopeMessages]
type webhookUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WebhookUpdateResponseEnvelopeSuccess bool

const (
	WebhookUpdateResponseEnvelopeSuccessTrue WebhookUpdateResponseEnvelopeSuccess = true
)

func (r WebhookUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WebhookUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WebhookDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type WebhookDeleteResponseEnvelope struct {
	Errors   []WebhookDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WebhookDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   WebhookDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WebhookDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    webhookDeleteResponseEnvelopeJSON    `json:"-"`
}

// webhookDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [WebhookDeleteResponseEnvelope]
type webhookDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WebhookDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    webhookDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// webhookDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WebhookDeleteResponseEnvelopeErrors]
type webhookDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WebhookDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    webhookDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// webhookDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WebhookDeleteResponseEnvelopeMessages]
type webhookDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WebhookDeleteResponseEnvelopeSuccess bool

const (
	WebhookDeleteResponseEnvelopeSuccessTrue WebhookDeleteResponseEnvelopeSuccess = true
)

func (r WebhookDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WebhookDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WebhookGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type WebhookGetResponseEnvelope struct {
	Errors   []WebhookGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WebhookGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WebhookGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WebhookGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    webhookGetResponseEnvelopeJSON    `json:"-"`
}

// webhookGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WebhookGetResponseEnvelope]
type webhookGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WebhookGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    webhookGetResponseEnvelopeErrorsJSON `json:"-"`
}

// webhookGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [WebhookGetResponseEnvelopeErrors]
type webhookGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WebhookGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    webhookGetResponseEnvelopeMessagesJSON `json:"-"`
}

// webhookGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [WebhookGetResponseEnvelopeMessages]
type webhookGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WebhookGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r webhookGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WebhookGetResponseEnvelopeSuccess bool

const (
	WebhookGetResponseEnvelopeSuccessTrue WebhookGetResponseEnvelopeSuccess = true
)

func (r WebhookGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WebhookGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

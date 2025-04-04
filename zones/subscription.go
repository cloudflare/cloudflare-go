// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// SubscriptionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubscriptionService] method instead.
type SubscriptionService struct {
	Options []option.RequestOption
}

// NewSubscriptionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubscriptionService(opts ...option.RequestOption) (r *SubscriptionService) {
	r = &SubscriptionService{}
	r.Options = opts
	return
}

// Create a zone subscription, either plan or add-ons.
func (r *SubscriptionService) New(ctx context.Context, identifier string, body SubscriptionNewParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env SubscriptionNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates zone subscriptions, either plan or add-ons.
func (r *SubscriptionService) Update(ctx context.Context, identifier string, body SubscriptionUpdateParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env SubscriptionUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists zone subscription details.
func (r *SubscriptionService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *interface{}, err error) {
	var env SubscriptionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/subscription", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SubscriptionNewParams struct {
	Subscription shared.SubscriptionParam `json:"subscription,required"`
}

func (r SubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Subscription)
}

type SubscriptionNewResponseEnvelope struct {
	Errors   []SubscriptionNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionNewResponseEnvelopeMessages `json:"messages,required"`
	Result   interface{}                               `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionNewResponseEnvelopeJSON    `json:"-"`
}

// subscriptionNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionNewResponseEnvelope]
type subscriptionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SubscriptionNewResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    subscriptionNewResponseEnvelopeErrorsJSON `json:"-"`
}

// subscriptionNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SubscriptionNewResponseEnvelopeErrors]
type subscriptionNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SubscriptionNewResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    subscriptionNewResponseEnvelopeMessagesJSON `json:"-"`
}

// subscriptionNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SubscriptionNewResponseEnvelopeMessages]
type subscriptionNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionNewResponseEnvelopeSuccess bool

const (
	SubscriptionNewResponseEnvelopeSuccessTrue SubscriptionNewResponseEnvelopeSuccess = true
)

func (r SubscriptionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubscriptionUpdateParams struct {
	Subscription shared.SubscriptionParam `json:"subscription,required"`
}

func (r SubscriptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Subscription)
}

type SubscriptionUpdateResponseEnvelope struct {
	Errors   []SubscriptionUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   interface{}                                  `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionUpdateResponseEnvelopeJSON    `json:"-"`
}

// subscriptionUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionUpdateResponseEnvelope]
type subscriptionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SubscriptionUpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    subscriptionUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// subscriptionUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SubscriptionUpdateResponseEnvelopeErrors]
type subscriptionUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SubscriptionUpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    subscriptionUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// subscriptionUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SubscriptionUpdateResponseEnvelopeMessages]
type subscriptionUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionUpdateResponseEnvelopeSuccess bool

const (
	SubscriptionUpdateResponseEnvelopeSuccessTrue SubscriptionUpdateResponseEnvelopeSuccess = true
)

func (r SubscriptionUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SubscriptionGetResponseEnvelope struct {
	Errors   []SubscriptionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SubscriptionGetResponseEnvelopeMessages `json:"messages,required"`
	Result   interface{}                               `json:"result,required"`
	// Whether the API call was successful
	Success SubscriptionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    subscriptionGetResponseEnvelopeJSON    `json:"-"`
}

// subscriptionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SubscriptionGetResponseEnvelope]
type subscriptionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SubscriptionGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    subscriptionGetResponseEnvelopeErrorsJSON `json:"-"`
}

// subscriptionGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SubscriptionGetResponseEnvelopeErrors]
type subscriptionGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SubscriptionGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    subscriptionGetResponseEnvelopeMessagesJSON `json:"-"`
}

// subscriptionGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SubscriptionGetResponseEnvelopeMessages]
type subscriptionGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubscriptionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SubscriptionGetResponseEnvelopeSuccess bool

const (
	SubscriptionGetResponseEnvelopeSuccessTrue SubscriptionGetResponseEnvelopeSuccess = true
)

func (r SubscriptionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SubscriptionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

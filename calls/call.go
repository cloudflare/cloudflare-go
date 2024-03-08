// File generated from our OpenAPI spec by Stainless.

package calls

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// CallService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewCallService] method instead.
type CallService struct {
	Options []option.RequestOption
}

// NewCallService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCallService(opts ...option.RequestOption) (r *CallService) {
	r = &CallService{}
	r.Options = opts
	return
}

// Creates a new Cloudflare calls app. An app is an unique enviroment where each
// Session can access all Tracks within the app.
func (r *CallService) New(ctx context.Context, params CallNewParams, opts ...option.RequestOption) (res *CallsAppWithSecret, err error) {
	opts = append(r.Options[:], opts...)
	var env CallNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/calls/apps", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit details for a single app.
func (r *CallService) Update(ctx context.Context, appID string, params CallUpdateParams, opts ...option.RequestOption) (res *CallsApp, err error) {
	opts = append(r.Options[:], opts...)
	var env CallUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/calls/apps/%s", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all apps in the Cloudflare account
func (r *CallService) List(ctx context.Context, query CallListParams, opts ...option.RequestOption) (res *[]CallsApp, err error) {
	opts = append(r.Options[:], opts...)
	var env CallListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/calls/apps", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an app from Cloudflare Calls
func (r *CallService) Delete(ctx context.Context, appID string, body CallDeleteParams, opts ...option.RequestOption) (res *CallsApp, err error) {
	opts = append(r.Options[:], opts...)
	var env CallDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/calls/apps/%s", body.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single Calls app.
func (r *CallService) Get(ctx context.Context, appID string, query CallGetParams, opts ...option.RequestOption) (res *CallsApp, err error) {
	opts = append(r.Options[:], opts...)
	var env CallGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/calls/apps/%s", query.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CallsApp struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	Uid  string       `json:"uid"`
	JSON callsAppJSON `json:"-"`
}

// callsAppJSON contains the JSON metadata for the struct [CallsApp]
type callsAppJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallsApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callsAppJSON) RawJSON() string {
	return r.raw
}

type CallsAppWithSecret struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// Bearer token to use the Calls API.
	Secret string `json:"secret"`
	// A Cloudflare-generated unique identifier for a item.
	Uid  string                 `json:"uid"`
	JSON callsAppWithSecretJSON `json:"-"`
}

// callsAppWithSecretJSON contains the JSON metadata for the struct
// [CallsAppWithSecret]
type callsAppWithSecretJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	Uid         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallsAppWithSecret) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callsAppWithSecretJSON) RawJSON() string {
	return r.raw
}

type CallNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A short description of Calls app, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r CallNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CallNewResponseEnvelope struct {
	Errors   []CallNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CallNewResponseEnvelopeMessages `json:"messages,required"`
	Result   CallsAppWithSecret                `json:"result,required"`
	// Whether the API call was successful
	Success CallNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    callNewResponseEnvelopeJSON    `json:"-"`
}

// callNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CallNewResponseEnvelope]
type callNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CallNewResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    callNewResponseEnvelopeErrorsJSON `json:"-"`
}

// callNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CallNewResponseEnvelopeErrors]
type callNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CallNewResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    callNewResponseEnvelopeMessagesJSON `json:"-"`
}

// callNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [CallNewResponseEnvelopeMessages]
type callNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CallNewResponseEnvelopeSuccess bool

const (
	CallNewResponseEnvelopeSuccessTrue CallNewResponseEnvelopeSuccess = true
)

type CallUpdateParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A short description of Calls app, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r CallUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CallUpdateResponseEnvelope struct {
	Errors   []CallUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CallUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   CallsApp                             `json:"result,required"`
	// Whether the API call was successful
	Success CallUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    callUpdateResponseEnvelopeJSON    `json:"-"`
}

// callUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [CallUpdateResponseEnvelope]
type callUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CallUpdateResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    callUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// callUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CallUpdateResponseEnvelopeErrors]
type callUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CallUpdateResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    callUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// callUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [CallUpdateResponseEnvelopeMessages]
type callUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CallUpdateResponseEnvelopeSuccess bool

const (
	CallUpdateResponseEnvelopeSuccessTrue CallUpdateResponseEnvelopeSuccess = true
)

type CallListParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type CallListResponseEnvelope struct {
	Errors   []CallListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CallListResponseEnvelopeMessages `json:"messages,required"`
	Result   []CallsApp                         `json:"result,required"`
	// Whether the API call was successful
	Success CallListResponseEnvelopeSuccess `json:"success,required"`
	JSON    callListResponseEnvelopeJSON    `json:"-"`
}

// callListResponseEnvelopeJSON contains the JSON metadata for the struct
// [CallListResponseEnvelope]
type callListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CallListResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    callListResponseEnvelopeErrorsJSON `json:"-"`
}

// callListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CallListResponseEnvelopeErrors]
type callListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CallListResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    callListResponseEnvelopeMessagesJSON `json:"-"`
}

// callListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [CallListResponseEnvelopeMessages]
type callListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CallListResponseEnvelopeSuccess bool

const (
	CallListResponseEnvelopeSuccessTrue CallListResponseEnvelopeSuccess = true
)

type CallDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type CallDeleteResponseEnvelope struct {
	Errors   []CallDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CallDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   CallsApp                             `json:"result,required"`
	// Whether the API call was successful
	Success CallDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    callDeleteResponseEnvelopeJSON    `json:"-"`
}

// callDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [CallDeleteResponseEnvelope]
type callDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CallDeleteResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    callDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// callDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CallDeleteResponseEnvelopeErrors]
type callDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CallDeleteResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    callDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// callDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [CallDeleteResponseEnvelopeMessages]
type callDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CallDeleteResponseEnvelopeSuccess bool

const (
	CallDeleteResponseEnvelopeSuccessTrue CallDeleteResponseEnvelopeSuccess = true
)

type CallGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type CallGetResponseEnvelope struct {
	Errors   []CallGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CallGetResponseEnvelopeMessages `json:"messages,required"`
	Result   CallsApp                          `json:"result,required"`
	// Whether the API call was successful
	Success CallGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    callGetResponseEnvelopeJSON    `json:"-"`
}

// callGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CallGetResponseEnvelope]
type callGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CallGetResponseEnvelopeErrors struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    callGetResponseEnvelopeErrorsJSON `json:"-"`
}

// callGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [CallGetResponseEnvelopeErrors]
type callGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CallGetResponseEnvelopeMessages struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    callGetResponseEnvelopeMessagesJSON `json:"-"`
}

// callGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [CallGetResponseEnvelopeMessages]
type callGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CallGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r callGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CallGetResponseEnvelopeSuccess bool

const (
	CallGetResponseEnvelopeSuccessTrue CallGetResponseEnvelopeSuccess = true
)

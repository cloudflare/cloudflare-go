// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *CallService) List(ctx context.Context, query CallListParams, opts ...option.RequestOption) (res *pagination.SinglePage[CallsApp], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/calls/apps", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all apps in the Cloudflare account
func (r *CallService) ListAutoPaging(ctx context.Context, query CallListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CallsApp] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes an app from Cloudflare Calls
func (r *CallService) Delete(ctx context.Context, appID string, body CallDeleteParams, opts ...option.RequestOption) (res *CallsApp, err error) {
	opts = append(r.Options[:], opts...)
	var env CallDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/calls/apps/%s", body.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
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
	UID  string       `json:"uid"`
	JSON callsAppJSON `json:"-"`
}

// callsAppJSON contains the JSON metadata for the struct [CallsApp]
type callsAppJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
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
	UID  string                 `json:"uid"`
	JSON callsAppWithSecretJSON `json:"-"`
}

// callsAppWithSecretJSON contains the JSON metadata for the struct
// [CallsAppWithSecret]
type callsAppWithSecretJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	UID         apijson.Field
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   CallsAppWithSecret    `json:"result,required"`
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

// Whether the API call was successful
type CallNewResponseEnvelopeSuccess bool

const (
	CallNewResponseEnvelopeSuccessTrue CallNewResponseEnvelopeSuccess = true
)

func (r CallNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CallNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   CallsApp              `json:"result,required"`
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

// Whether the API call was successful
type CallUpdateResponseEnvelopeSuccess bool

const (
	CallUpdateResponseEnvelopeSuccessTrue CallUpdateResponseEnvelopeSuccess = true
)

func (r CallUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CallUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CallListParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type CallDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type CallDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   CallsApp              `json:"result,required"`
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

// Whether the API call was successful
type CallDeleteResponseEnvelopeSuccess bool

const (
	CallDeleteResponseEnvelopeSuccessTrue CallDeleteResponseEnvelopeSuccess = true
)

func (r CallDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CallDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CallGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type CallGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   CallsApp              `json:"result,required"`
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

// Whether the API call was successful
type CallGetResponseEnvelopeSuccess bool

const (
	CallGetResponseEnvelopeSuccessTrue CallGetResponseEnvelopeSuccess = true
)

func (r CallGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CallGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

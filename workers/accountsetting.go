// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AccountSettingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountSettingService] method
// instead.
type AccountSettingService struct {
	Options []option.RequestOption
}

// NewAccountSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountSettingService(opts ...option.RequestOption) (r *AccountSettingService) {
	r = &AccountSettingService{}
	r.Options = opts
	return
}

// Creates Worker account settings for an account.
func (r *AccountSettingService) Update(ctx context.Context, params AccountSettingUpdateParams, opts ...option.RequestOption) (res *AccountSettingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountSettingUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/account-settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Worker account settings for an account.
func (r *AccountSettingService) Get(ctx context.Context, query AccountSettingGetParams, opts ...option.RequestOption) (res *AccountSettingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccountSettingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/workers/account-settings", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccountSettingUpdateResponse struct {
	DefaultUsageModel interface{}                      `json:"default_usage_model"`
	GreenCompute      interface{}                      `json:"green_compute"`
	JSON              accountSettingUpdateResponseJSON `json:"-"`
}

// accountSettingUpdateResponseJSON contains the JSON metadata for the struct
// [AccountSettingUpdateResponse]
type accountSettingUpdateResponseJSON struct {
	DefaultUsageModel apijson.Field
	GreenCompute      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSettingGetResponse struct {
	DefaultUsageModel interface{}                   `json:"default_usage_model"`
	GreenCompute      interface{}                   `json:"green_compute"`
	JSON              accountSettingGetResponseJSON `json:"-"`
}

// accountSettingGetResponseJSON contains the JSON metadata for the struct
// [AccountSettingGetResponse]
type accountSettingGetResponseJSON struct {
	DefaultUsageModel apijson.Field
	GreenCompute      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountSettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccountSettingUpdateParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r AccountSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccountSettingUpdateResponseEnvelope struct {
	Errors   []AccountSettingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccountSettingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountSettingUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccountSettingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accountSettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// accountSettingUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccountSettingUpdateResponseEnvelope]
type accountSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccountSettingUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountSettingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accountSettingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccountSettingUpdateResponseEnvelopeErrors]
type accountSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccountSettingUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountSettingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accountSettingUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccountSettingUpdateResponseEnvelopeMessages]
type accountSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountSettingUpdateResponseEnvelopeSuccess bool

const (
	AccountSettingUpdateResponseEnvelopeSuccessTrue AccountSettingUpdateResponseEnvelopeSuccess = true
)

type AccountSettingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccountSettingGetResponseEnvelope struct {
	Errors   []AccountSettingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccountSettingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccountSettingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccountSettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accountSettingGetResponseEnvelopeJSON    `json:"-"`
}

// accountSettingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccountSettingGetResponseEnvelope]
type accountSettingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccountSettingGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountSettingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accountSettingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccountSettingGetResponseEnvelopeErrors]
type accountSettingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSettingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccountSettingGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountSettingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accountSettingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccountSettingGetResponseEnvelopeMessages]
type accountSettingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSettingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accountSettingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccountSettingGetResponseEnvelopeSuccess bool

const (
	AccountSettingGetResponseEnvelopeSuccessTrue AccountSettingGetResponseEnvelopeSuccess = true
)

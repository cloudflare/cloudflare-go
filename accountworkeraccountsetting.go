// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountWorkerAccountSettingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountWorkerAccountSettingService] method instead.
type AccountWorkerAccountSettingService struct {
	Options []option.RequestOption
}

// NewAccountWorkerAccountSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerAccountSettingService(opts ...option.RequestOption) (r *AccountWorkerAccountSettingService) {
	r = &AccountWorkerAccountSettingService{}
	r.Options = opts
	return
}

// Creates Worker account settings for an account.
func (r *AccountWorkerAccountSettingService) WorkerAccountSettingsNewWorkerAccountSettings(ctx context.Context, accountIdentifier string, body AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsParams, opts ...option.RequestOption) (res *AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/account-settings", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches Worker account settings for an account.
func (r *AccountWorkerAccountSettingService) WorkerAccountSettingsFetchWorkerAccountSettings(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/account-settings", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponse struct {
	Errors   []AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseError   `json:"errors"`
	Messages []AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseMessage `json:"messages"`
	Result   AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseSuccess `json:"success"`
	JSON    accountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseJSON    `json:"-"`
}

// accountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseJSON
// contains the JSON metadata for the struct
// [AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponse]
type accountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseError struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    accountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseErrorJSON `json:"-"`
}

// accountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseError]
type accountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseMessage struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseMessageJSON `json:"-"`
}

// accountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseMessage]
type accountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseResult struct {
	DefaultUsageModel interface{}                                                                                `json:"default_usage_model"`
	GreenCompute      interface{}                                                                                `json:"green_compute"`
	JSON              accountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseResultJSON `json:"-"`
}

// accountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseResult]
type accountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseResultJSON struct {
	DefaultUsageModel apijson.Field
	GreenCompute      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseSuccess bool

const (
	AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseSuccessTrue AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsResponseSuccess = true
)

type AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponse struct {
	Errors   []AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseError   `json:"errors"`
	Messages []AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseMessage `json:"messages"`
	Result   AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseSuccess `json:"success"`
	JSON    accountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseJSON    `json:"-"`
}

// accountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseJSON
// contains the JSON metadata for the struct
// [AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponse]
type accountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseError struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    accountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseErrorJSON `json:"-"`
}

// accountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseError]
type accountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseMessage struct {
	Code    int64                                                                                         `json:"code,required"`
	Message string                                                                                        `json:"message,required"`
	JSON    accountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseMessageJSON `json:"-"`
}

// accountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseMessage]
type accountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseResult struct {
	DefaultUsageModel interface{}                                                                                  `json:"default_usage_model"`
	GreenCompute      interface{}                                                                                  `json:"green_compute"`
	JSON              accountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseResultJSON `json:"-"`
}

// accountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseResult]
type accountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseResultJSON struct {
	DefaultUsageModel apijson.Field
	GreenCompute      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseSuccess bool

const (
	AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseSuccessTrue AccountWorkerAccountSettingWorkerAccountSettingsFetchWorkerAccountSettingsResponseSuccess = true
)

type AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

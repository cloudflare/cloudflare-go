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

// Creates Worker Account Settings for an account.
func (r *AccountWorkerAccountSettingService) WorkerAccountSettingsNewWorkerAccountSettings(ctx context.Context, accountIdentifier string, body AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsParams, opts ...option.RequestOption) (res *AccountSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/account-settings", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches Worker Account Settings for an account.
func (r *AccountWorkerAccountSettingService) WorkerAccountSettingsFetchWorkerAccountSettings(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/account-settings", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountSettingsResponse struct {
	Errors   []AccountSettingsResponseError   `json:"errors"`
	Messages []AccountSettingsResponseMessage `json:"messages"`
	Result   AccountSettingsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountSettingsResponseSuccess `json:"success"`
	JSON    accountSettingsResponseJSON    `json:"-"`
}

// accountSettingsResponseJSON contains the JSON metadata for the struct
// [AccountSettingsResponse]
type accountSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSettingsResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    accountSettingsResponseErrorJSON `json:"-"`
}

// accountSettingsResponseErrorJSON contains the JSON metadata for the struct
// [AccountSettingsResponseError]
type accountSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSettingsResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    accountSettingsResponseMessageJSON `json:"-"`
}

// accountSettingsResponseMessageJSON contains the JSON metadata for the struct
// [AccountSettingsResponseMessage]
type accountSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountSettingsResponseResult struct {
	DefaultUsageModel interface{}                       `json:"default_usage_model"`
	GreenCompute      interface{}                       `json:"green_compute"`
	JSON              accountSettingsResponseResultJSON `json:"-"`
}

// accountSettingsResponseResultJSON contains the JSON metadata for the struct
// [AccountSettingsResponseResult]
type accountSettingsResponseResultJSON struct {
	DefaultUsageModel apijson.Field
	GreenCompute      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccountSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountSettingsResponseSuccess bool

const (
	AccountSettingsResponseSuccessTrue AccountSettingsResponseSuccess = true
)

type AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountWorkerAccountSettingWorkerAccountSettingsNewWorkerAccountSettingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

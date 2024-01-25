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

// AccountWorkerScriptScheduleService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountWorkerScriptScheduleService] method instead.
type AccountWorkerScriptScheduleService struct {
	Options []option.RequestOption
}

// NewAccountWorkerScriptScheduleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerScriptScheduleService(opts ...option.RequestOption) (r *AccountWorkerScriptScheduleService) {
	r = &AccountWorkerScriptScheduleService{}
	r.Options = opts
	return
}

// Fetches Cron Triggers for a Worker.
func (r *AccountWorkerScriptScheduleService) WorkerCronTriggerGetCronTriggers(ctx context.Context, accountIdentifier string, scriptName string, opts ...option.RequestOption) (res *AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates Cron Triggers for a Worker.
func (r *AccountWorkerScriptScheduleService) WorkerCronTriggerUpdateCronTriggers(ctx context.Context, accountIdentifier string, scriptName string, body AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersParams, opts ...option.RequestOption) (res *AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponse struct {
	Errors   []AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseError   `json:"errors"`
	Messages []AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseMessage `json:"messages"`
	Result   AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseSuccess `json:"success"`
	JSON    accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseJSON    `json:"-"`
}

// accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseJSON contains
// the JSON metadata for the struct
// [AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponse]
type accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseError struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseErrorJSON `json:"-"`
}

// accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseError]
type accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseMessage struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseMessageJSON `json:"-"`
}

// accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseMessage]
type accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseResult struct {
	Schedules []AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseResultSchedule `json:"schedules"`
	JSON      accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseResultJSON       `json:"-"`
}

// accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseResultJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseResult]
type accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseResultJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseResultSchedule struct {
	CreatedOn  interface{}                                                                           `json:"created_on"`
	Cron       interface{}                                                                           `json:"cron"`
	ModifiedOn interface{}                                                                           `json:"modified_on"`
	JSON       accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseResultScheduleJSON `json:"-"`
}

// accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseResultScheduleJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseResultSchedule]
type accountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseResultScheduleJSON struct {
	CreatedOn   apijson.Field
	Cron        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseResultSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseSuccess bool

const (
	AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseSuccessTrue AccountWorkerScriptScheduleWorkerCronTriggerGetCronTriggersResponseSuccess = true
)

type AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponse struct {
	Errors   []AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseError   `json:"errors"`
	Messages []AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseMessage `json:"messages"`
	Result   AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseSuccess `json:"success"`
	JSON    accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseJSON    `json:"-"`
}

// accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponse]
type accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseError struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseErrorJSON `json:"-"`
}

// accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseError]
type accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseMessage struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseMessageJSON `json:"-"`
}

// accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseMessage]
type accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseResult struct {
	Schedules []AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseResultSchedule `json:"schedules"`
	JSON      accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseResultJSON       `json:"-"`
}

// accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseResultJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseResult]
type accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseResultJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseResultSchedule struct {
	CreatedOn  interface{}                                                                              `json:"created_on"`
	Cron       interface{}                                                                              `json:"cron"`
	ModifiedOn interface{}                                                                              `json:"modified_on"`
	JSON       accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseResultScheduleJSON `json:"-"`
}

// accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseResultScheduleJSON
// contains the JSON metadata for the struct
// [AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseResultSchedule]
type accountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseResultScheduleJSON struct {
	CreatedOn   apijson.Field
	Cron        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseResultSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseSuccess bool

const (
	AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseSuccessTrue AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersResponseSuccess = true
)

type AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

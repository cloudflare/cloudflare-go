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
func (r *AccountWorkerScriptScheduleService) WorkerCronTriggerGetCronTriggers(ctx context.Context, accountIdentifier string, scriptName string, opts ...option.RequestOption) (res *CronTriggerResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates Cron Triggers for a Worker.
func (r *AccountWorkerScriptScheduleService) WorkerCronTriggerUpdateCronTriggers(ctx context.Context, accountIdentifier string, scriptName string, body AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersParams, opts ...option.RequestOption) (res *CronTriggerResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/schedules", accountIdentifier, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CronTriggerResponseCollection struct {
	Errors   []CronTriggerResponseCollectionError   `json:"errors"`
	Messages []CronTriggerResponseCollectionMessage `json:"messages"`
	Result   CronTriggerResponseCollectionResult    `json:"result"`
	// Whether the API call was successful
	Success CronTriggerResponseCollectionSuccess `json:"success"`
	JSON    cronTriggerResponseCollectionJSON    `json:"-"`
}

// cronTriggerResponseCollectionJSON contains the JSON metadata for the struct
// [CronTriggerResponseCollection]
type cronTriggerResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CronTriggerResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CronTriggerResponseCollectionError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    cronTriggerResponseCollectionErrorJSON `json:"-"`
}

// cronTriggerResponseCollectionErrorJSON contains the JSON metadata for the struct
// [CronTriggerResponseCollectionError]
type cronTriggerResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CronTriggerResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CronTriggerResponseCollectionMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    cronTriggerResponseCollectionMessageJSON `json:"-"`
}

// cronTriggerResponseCollectionMessageJSON contains the JSON metadata for the
// struct [CronTriggerResponseCollectionMessage]
type cronTriggerResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CronTriggerResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CronTriggerResponseCollectionResult struct {
	Schedules []CronTriggerResponseCollectionResultSchedule `json:"schedules"`
	JSON      cronTriggerResponseCollectionResultJSON       `json:"-"`
}

// cronTriggerResponseCollectionResultJSON contains the JSON metadata for the
// struct [CronTriggerResponseCollectionResult]
type cronTriggerResponseCollectionResultJSON struct {
	Schedules   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CronTriggerResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CronTriggerResponseCollectionResultSchedule struct {
	CreatedOn  interface{}                                     `json:"created_on"`
	Cron       interface{}                                     `json:"cron"`
	ModifiedOn interface{}                                     `json:"modified_on"`
	JSON       cronTriggerResponseCollectionResultScheduleJSON `json:"-"`
}

// cronTriggerResponseCollectionResultScheduleJSON contains the JSON metadata for
// the struct [CronTriggerResponseCollectionResultSchedule]
type cronTriggerResponseCollectionResultScheduleJSON struct {
	CreatedOn   apijson.Field
	Cron        apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CronTriggerResponseCollectionResultSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CronTriggerResponseCollectionSuccess bool

const (
	CronTriggerResponseCollectionSuccessTrue CronTriggerResponseCollectionSuccess = true
)

type AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountWorkerScriptScheduleWorkerCronTriggerUpdateCronTriggersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

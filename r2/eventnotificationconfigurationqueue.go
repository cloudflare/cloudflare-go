// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// EventNotificationConfigurationQueueService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventNotificationConfigurationQueueService] method instead.
type EventNotificationConfigurationQueueService struct {
	Options []option.RequestOption
}

// NewEventNotificationConfigurationQueueService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewEventNotificationConfigurationQueueService(opts ...option.RequestOption) (r *EventNotificationConfigurationQueueService) {
	r = &EventNotificationConfigurationQueueService{}
	r.Options = opts
	return
}

// Create event notification rule.
func (r *EventNotificationConfigurationQueueService) Update(ctx context.Context, bucketName string, queueID string, params EventNotificationConfigurationQueueUpdateParams, opts ...option.RequestOption) (res *EventNotificationConfigurationQueueUpdateResponse, err error) {
	var env EventNotificationConfigurationQueueUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration/queues/%s", params.AccountID, bucketName, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete an event notification rule. **If no body is provided, all rules for
// specified queue will be deleted**.
func (r *EventNotificationConfigurationQueueService) Delete(ctx context.Context, bucketName string, queueID string, params EventNotificationConfigurationQueueDeleteParams, opts ...option.RequestOption) (res *EventNotificationConfigurationQueueDeleteResponse, err error) {
	var env EventNotificationConfigurationQueueDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration/queues/%s", params.AccountID, bucketName, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EventNotificationConfigurationQueueUpdateResponse = interface{}

type EventNotificationConfigurationQueueDeleteResponse = interface{}

type EventNotificationConfigurationQueueUpdateParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Array of rules to drive notifications
	Rules param.Field[[]EventNotificationConfigurationQueueUpdateParamsRule] `json:"rules"`
	// The bucket jurisdiction
	CfR2Jurisdiction param.Field[EventNotificationConfigurationQueueUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r EventNotificationConfigurationQueueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventNotificationConfigurationQueueUpdateParamsRule struct {
	// Array of R2 object actions that will trigger notifications
	Actions param.Field[[]EventNotificationConfigurationQueueUpdateParamsRulesAction] `json:"actions,required"`
	// A description that can be used to identify the event notification rule after
	// creation
	Description param.Field[string] `json:"description"`
	// Notifications will be sent only for objects with this prefix
	Prefix param.Field[string] `json:"prefix"`
	// Notifications will be sent only for objects with this suffix
	Suffix param.Field[string] `json:"suffix"`
}

func (r EventNotificationConfigurationQueueUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EventNotificationConfigurationQueueUpdateParamsRulesAction string

const (
	EventNotificationConfigurationQueueUpdateParamsRulesActionPutObject               EventNotificationConfigurationQueueUpdateParamsRulesAction = "PutObject"
	EventNotificationConfigurationQueueUpdateParamsRulesActionCopyObject              EventNotificationConfigurationQueueUpdateParamsRulesAction = "CopyObject"
	EventNotificationConfigurationQueueUpdateParamsRulesActionDeleteObject            EventNotificationConfigurationQueueUpdateParamsRulesAction = "DeleteObject"
	EventNotificationConfigurationQueueUpdateParamsRulesActionCompleteMultipartUpload EventNotificationConfigurationQueueUpdateParamsRulesAction = "CompleteMultipartUpload"
	EventNotificationConfigurationQueueUpdateParamsRulesActionLifecycleDeletion       EventNotificationConfigurationQueueUpdateParamsRulesAction = "LifecycleDeletion"
)

func (r EventNotificationConfigurationQueueUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case EventNotificationConfigurationQueueUpdateParamsRulesActionPutObject, EventNotificationConfigurationQueueUpdateParamsRulesActionCopyObject, EventNotificationConfigurationQueueUpdateParamsRulesActionDeleteObject, EventNotificationConfigurationQueueUpdateParamsRulesActionCompleteMultipartUpload, EventNotificationConfigurationQueueUpdateParamsRulesActionLifecycleDeletion:
		return true
	}
	return false
}

// The bucket jurisdiction
type EventNotificationConfigurationQueueUpdateParamsCfR2Jurisdiction string

const (
	EventNotificationConfigurationQueueUpdateParamsCfR2JurisdictionDefault EventNotificationConfigurationQueueUpdateParamsCfR2Jurisdiction = "default"
	EventNotificationConfigurationQueueUpdateParamsCfR2JurisdictionEu      EventNotificationConfigurationQueueUpdateParamsCfR2Jurisdiction = "eu"
	EventNotificationConfigurationQueueUpdateParamsCfR2JurisdictionFedramp EventNotificationConfigurationQueueUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r EventNotificationConfigurationQueueUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case EventNotificationConfigurationQueueUpdateParamsCfR2JurisdictionDefault, EventNotificationConfigurationQueueUpdateParamsCfR2JurisdictionEu, EventNotificationConfigurationQueueUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type EventNotificationConfigurationQueueUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo                             `json:"errors,required"`
	Messages []string                                          `json:"messages,required"`
	Result   EventNotificationConfigurationQueueUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success EventNotificationConfigurationQueueUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    eventNotificationConfigurationQueueUpdateResponseEnvelopeJSON    `json:"-"`
}

// eventNotificationConfigurationQueueUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [EventNotificationConfigurationQueueUpdateResponseEnvelope]
type eventNotificationConfigurationQueueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventNotificationConfigurationQueueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventNotificationConfigurationQueueUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EventNotificationConfigurationQueueUpdateResponseEnvelopeSuccess bool

const (
	EventNotificationConfigurationQueueUpdateResponseEnvelopeSuccessTrue EventNotificationConfigurationQueueUpdateResponseEnvelopeSuccess = true
)

func (r EventNotificationConfigurationQueueUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EventNotificationConfigurationQueueUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type EventNotificationConfigurationQueueDeleteParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// The bucket jurisdiction
	CfR2Jurisdiction param.Field[EventNotificationConfigurationQueueDeleteParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction
type EventNotificationConfigurationQueueDeleteParamsCfR2Jurisdiction string

const (
	EventNotificationConfigurationQueueDeleteParamsCfR2JurisdictionDefault EventNotificationConfigurationQueueDeleteParamsCfR2Jurisdiction = "default"
	EventNotificationConfigurationQueueDeleteParamsCfR2JurisdictionEu      EventNotificationConfigurationQueueDeleteParamsCfR2Jurisdiction = "eu"
	EventNotificationConfigurationQueueDeleteParamsCfR2JurisdictionFedramp EventNotificationConfigurationQueueDeleteParamsCfR2Jurisdiction = "fedramp"
)

func (r EventNotificationConfigurationQueueDeleteParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case EventNotificationConfigurationQueueDeleteParamsCfR2JurisdictionDefault, EventNotificationConfigurationQueueDeleteParamsCfR2JurisdictionEu, EventNotificationConfigurationQueueDeleteParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type EventNotificationConfigurationQueueDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                             `json:"errors,required"`
	Messages []string                                          `json:"messages,required"`
	Result   EventNotificationConfigurationQueueDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success EventNotificationConfigurationQueueDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    eventNotificationConfigurationQueueDeleteResponseEnvelopeJSON    `json:"-"`
}

// eventNotificationConfigurationQueueDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [EventNotificationConfigurationQueueDeleteResponseEnvelope]
type eventNotificationConfigurationQueueDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventNotificationConfigurationQueueDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventNotificationConfigurationQueueDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EventNotificationConfigurationQueueDeleteResponseEnvelopeSuccess bool

const (
	EventNotificationConfigurationQueueDeleteResponseEnvelopeSuccessTrue EventNotificationConfigurationQueueDeleteResponseEnvelopeSuccess = true
)

func (r EventNotificationConfigurationQueueDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EventNotificationConfigurationQueueDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

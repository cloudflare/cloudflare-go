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

// EventNotificationConfigurationService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEventNotificationConfigurationService] method instead.
type EventNotificationConfigurationService struct {
	Options []option.RequestOption
	Queues  *EventNotificationConfigurationQueueService
}

// NewEventNotificationConfigurationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEventNotificationConfigurationService(opts ...option.RequestOption) (r *EventNotificationConfigurationService) {
	r = &EventNotificationConfigurationService{}
	r.Options = opts
	r.Queues = NewEventNotificationConfigurationQueueService(opts...)
	return
}

// List all event notification rules for a bucket.
func (r *EventNotificationConfigurationService) Get(ctx context.Context, bucketName string, params EventNotificationConfigurationGetParams, opts ...option.RequestOption) (res *EventNotificationConfigurationGetResponse, err error) {
	var env EventNotificationConfigurationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration", params.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EventNotificationConfigurationGetResponse struct {
	// Name of the bucket.
	BucketName string `json:"bucketName"`
	// List of queues associated with the bucket.
	Queues []EventNotificationConfigurationGetResponseQueue `json:"queues"`
	JSON   eventNotificationConfigurationGetResponseJSON    `json:"-"`
}

// eventNotificationConfigurationGetResponseJSON contains the JSON metadata for the
// struct [EventNotificationConfigurationGetResponse]
type eventNotificationConfigurationGetResponseJSON struct {
	BucketName  apijson.Field
	Queues      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventNotificationConfigurationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventNotificationConfigurationGetResponseJSON) RawJSON() string {
	return r.raw
}

type EventNotificationConfigurationGetResponseQueue struct {
	// Queue ID
	QueueID string `json:"queueId"`
	// Name of the queue
	QueueName string                                                `json:"queueName"`
	Rules     []EventNotificationConfigurationGetResponseQueuesRule `json:"rules"`
	JSON      eventNotificationConfigurationGetResponseQueueJSON    `json:"-"`
}

// eventNotificationConfigurationGetResponseQueueJSON contains the JSON metadata
// for the struct [EventNotificationConfigurationGetResponseQueue]
type eventNotificationConfigurationGetResponseQueueJSON struct {
	QueueID     apijson.Field
	QueueName   apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventNotificationConfigurationGetResponseQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventNotificationConfigurationGetResponseQueueJSON) RawJSON() string {
	return r.raw
}

type EventNotificationConfigurationGetResponseQueuesRule struct {
	// Array of R2 object actions that will trigger notifications
	Actions []EventNotificationConfigurationGetResponseQueuesRulesAction `json:"actions,required"`
	// Timestamp when the rule was created
	CreatedAt string `json:"createdAt"`
	// Notifications will be sent only for objects with this prefix
	Prefix string `json:"prefix"`
	// Rule ID
	RuleID string `json:"ruleId"`
	// Notifications will be sent only for objects with this suffix
	Suffix string                                                  `json:"suffix"`
	JSON   eventNotificationConfigurationGetResponseQueuesRuleJSON `json:"-"`
}

// eventNotificationConfigurationGetResponseQueuesRuleJSON contains the JSON
// metadata for the struct [EventNotificationConfigurationGetResponseQueuesRule]
type eventNotificationConfigurationGetResponseQueuesRuleJSON struct {
	Actions     apijson.Field
	CreatedAt   apijson.Field
	Prefix      apijson.Field
	RuleID      apijson.Field
	Suffix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventNotificationConfigurationGetResponseQueuesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventNotificationConfigurationGetResponseQueuesRuleJSON) RawJSON() string {
	return r.raw
}

type EventNotificationConfigurationGetResponseQueuesRulesAction string

const (
	EventNotificationConfigurationGetResponseQueuesRulesActionPutObject               EventNotificationConfigurationGetResponseQueuesRulesAction = "PutObject"
	EventNotificationConfigurationGetResponseQueuesRulesActionCopyObject              EventNotificationConfigurationGetResponseQueuesRulesAction = "CopyObject"
	EventNotificationConfigurationGetResponseQueuesRulesActionDeleteObject            EventNotificationConfigurationGetResponseQueuesRulesAction = "DeleteObject"
	EventNotificationConfigurationGetResponseQueuesRulesActionCompleteMultipartUpload EventNotificationConfigurationGetResponseQueuesRulesAction = "CompleteMultipartUpload"
	EventNotificationConfigurationGetResponseQueuesRulesActionLifecycleDeletion       EventNotificationConfigurationGetResponseQueuesRulesAction = "LifecycleDeletion"
)

func (r EventNotificationConfigurationGetResponseQueuesRulesAction) IsKnown() bool {
	switch r {
	case EventNotificationConfigurationGetResponseQueuesRulesActionPutObject, EventNotificationConfigurationGetResponseQueuesRulesActionCopyObject, EventNotificationConfigurationGetResponseQueuesRulesActionDeleteObject, EventNotificationConfigurationGetResponseQueuesRulesActionCompleteMultipartUpload, EventNotificationConfigurationGetResponseQueuesRulesActionLifecycleDeletion:
		return true
	}
	return false
}

type EventNotificationConfigurationGetParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// The bucket jurisdiction
	CfR2Jurisdiction param.Field[EventNotificationConfigurationGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction
type EventNotificationConfigurationGetParamsCfR2Jurisdiction string

const (
	EventNotificationConfigurationGetParamsCfR2JurisdictionDefault EventNotificationConfigurationGetParamsCfR2Jurisdiction = "default"
	EventNotificationConfigurationGetParamsCfR2JurisdictionEu      EventNotificationConfigurationGetParamsCfR2Jurisdiction = "eu"
	EventNotificationConfigurationGetParamsCfR2JurisdictionFedramp EventNotificationConfigurationGetParamsCfR2Jurisdiction = "fedramp"
)

func (r EventNotificationConfigurationGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case EventNotificationConfigurationGetParamsCfR2JurisdictionDefault, EventNotificationConfigurationGetParamsCfR2JurisdictionEu, EventNotificationConfigurationGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type EventNotificationConfigurationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                     `json:"errors,required"`
	Messages []string                                  `json:"messages,required"`
	Result   EventNotificationConfigurationGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success EventNotificationConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    eventNotificationConfigurationGetResponseEnvelopeJSON    `json:"-"`
}

// eventNotificationConfigurationGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [EventNotificationConfigurationGetResponseEnvelope]
type eventNotificationConfigurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EventNotificationConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r eventNotificationConfigurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EventNotificationConfigurationGetResponseEnvelopeSuccess bool

const (
	EventNotificationConfigurationGetResponseEnvelopeSuccessTrue EventNotificationConfigurationGetResponseEnvelopeSuccess = true
)

func (r EventNotificationConfigurationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EventNotificationConfigurationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

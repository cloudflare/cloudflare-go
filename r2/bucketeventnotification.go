// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// BucketEventNotificationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketEventNotificationService] method instead.
type BucketEventNotificationService struct {
	Options []option.RequestOption
}

// NewBucketEventNotificationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewBucketEventNotificationService(opts ...option.RequestOption) (r *BucketEventNotificationService) {
	r = &BucketEventNotificationService{}
	r.Options = opts
	return
}

// Create event notification rule.
func (r *BucketEventNotificationService) Update(ctx context.Context, bucketName string, queueID string, params BucketEventNotificationUpdateParams, opts ...option.RequestOption) (res *BucketEventNotificationUpdateResponse, err error) {
	var env BucketEventNotificationUpdateResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
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
func (r *BucketEventNotificationService) Delete(ctx context.Context, bucketName string, queueID string, params BucketEventNotificationDeleteParams, opts ...option.RequestOption) (res *BucketEventNotificationDeleteResponse, err error) {
	var env BucketEventNotificationDeleteResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
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

// List all event notification rules for a bucket.
func (r *BucketEventNotificationService) Get(ctx context.Context, bucketName string, params BucketEventNotificationGetParams, opts ...option.RequestOption) (res *BucketEventNotificationGetResponse, err error) {
	var env BucketEventNotificationGetResponseEnvelope
	if params.Jurisdiction.Present {
		opts = append(opts, option.WithHeader("cf-r2-jurisdiction", fmt.Sprintf("%s", params.Jurisdiction)))
	}
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

type BucketEventNotificationUpdateResponse = interface{}

type BucketEventNotificationDeleteResponse = interface{}

type BucketEventNotificationGetResponse struct {
	// Name of the bucket.
	BucketName string `json:"bucketName"`
	// List of queues associated with the bucket.
	Queues []BucketEventNotificationGetResponseQueue `json:"queues"`
	JSON   bucketEventNotificationGetResponseJSON    `json:"-"`
}

// bucketEventNotificationGetResponseJSON contains the JSON metadata for the struct
// [BucketEventNotificationGetResponse]
type bucketEventNotificationGetResponseJSON struct {
	BucketName  apijson.Field
	Queues      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationGetResponseQueue struct {
	// Queue ID.
	QueueID string `json:"queueId"`
	// Name of the queue.
	QueueName string                                         `json:"queueName"`
	Rules     []BucketEventNotificationGetResponseQueuesRule `json:"rules"`
	JSON      bucketEventNotificationGetResponseQueueJSON    `json:"-"`
}

// bucketEventNotificationGetResponseQueueJSON contains the JSON metadata for the
// struct [BucketEventNotificationGetResponseQueue]
type bucketEventNotificationGetResponseQueueJSON struct {
	QueueID     apijson.Field
	QueueName   apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponseQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseQueueJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationGetResponseQueuesRule struct {
	// Array of R2 object actions that will trigger notifications.
	Actions []BucketEventNotificationGetResponseQueuesRulesAction `json:"actions,required"`
	// Timestamp when the rule was created.
	CreatedAt string `json:"createdAt"`
	// A description that can be used to identify the event notification rule after
	// creation.
	Description string `json:"description"`
	// Notifications will be sent only for objects with this prefix.
	Prefix string `json:"prefix"`
	// Rule ID.
	RuleID string `json:"ruleId"`
	// Notifications will be sent only for objects with this suffix.
	Suffix string                                           `json:"suffix"`
	JSON   bucketEventNotificationGetResponseQueuesRuleJSON `json:"-"`
}

// bucketEventNotificationGetResponseQueuesRuleJSON contains the JSON metadata for
// the struct [BucketEventNotificationGetResponseQueuesRule]
type bucketEventNotificationGetResponseQueuesRuleJSON struct {
	Actions     apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Prefix      apijson.Field
	RuleID      apijson.Field
	Suffix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponseQueuesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseQueuesRuleJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationGetResponseQueuesRulesAction string

const (
	BucketEventNotificationGetResponseQueuesRulesActionPutObject               BucketEventNotificationGetResponseQueuesRulesAction = "PutObject"
	BucketEventNotificationGetResponseQueuesRulesActionCopyObject              BucketEventNotificationGetResponseQueuesRulesAction = "CopyObject"
	BucketEventNotificationGetResponseQueuesRulesActionDeleteObject            BucketEventNotificationGetResponseQueuesRulesAction = "DeleteObject"
	BucketEventNotificationGetResponseQueuesRulesActionCompleteMultipartUpload BucketEventNotificationGetResponseQueuesRulesAction = "CompleteMultipartUpload"
	BucketEventNotificationGetResponseQueuesRulesActionLifecycleDeletion       BucketEventNotificationGetResponseQueuesRulesAction = "LifecycleDeletion"
)

func (r BucketEventNotificationGetResponseQueuesRulesAction) IsKnown() bool {
	switch r {
	case BucketEventNotificationGetResponseQueuesRulesActionPutObject, BucketEventNotificationGetResponseQueuesRulesActionCopyObject, BucketEventNotificationGetResponseQueuesRulesActionDeleteObject, BucketEventNotificationGetResponseQueuesRulesActionCompleteMultipartUpload, BucketEventNotificationGetResponseQueuesRulesActionLifecycleDeletion:
		return true
	}
	return false
}

type BucketEventNotificationUpdateParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Array of rules to drive notifications.
	Rules param.Field[[]BucketEventNotificationUpdateParamsRule] `json:"rules"`
	// The bucket jurisdiction.
	Jurisdiction param.Field[BucketEventNotificationUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r BucketEventNotificationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BucketEventNotificationUpdateParamsRule struct {
	// Array of R2 object actions that will trigger notifications.
	Actions param.Field[[]BucketEventNotificationUpdateParamsRulesAction] `json:"actions,required"`
	// A description that can be used to identify the event notification rule after
	// creation.
	Description param.Field[string] `json:"description"`
	// Notifications will be sent only for objects with this prefix.
	Prefix param.Field[string] `json:"prefix"`
	// Notifications will be sent only for objects with this suffix.
	Suffix param.Field[string] `json:"suffix"`
}

func (r BucketEventNotificationUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BucketEventNotificationUpdateParamsRulesAction string

const (
	BucketEventNotificationUpdateParamsRulesActionPutObject               BucketEventNotificationUpdateParamsRulesAction = "PutObject"
	BucketEventNotificationUpdateParamsRulesActionCopyObject              BucketEventNotificationUpdateParamsRulesAction = "CopyObject"
	BucketEventNotificationUpdateParamsRulesActionDeleteObject            BucketEventNotificationUpdateParamsRulesAction = "DeleteObject"
	BucketEventNotificationUpdateParamsRulesActionCompleteMultipartUpload BucketEventNotificationUpdateParamsRulesAction = "CompleteMultipartUpload"
	BucketEventNotificationUpdateParamsRulesActionLifecycleDeletion       BucketEventNotificationUpdateParamsRulesAction = "LifecycleDeletion"
)

func (r BucketEventNotificationUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case BucketEventNotificationUpdateParamsRulesActionPutObject, BucketEventNotificationUpdateParamsRulesActionCopyObject, BucketEventNotificationUpdateParamsRulesActionDeleteObject, BucketEventNotificationUpdateParamsRulesActionCompleteMultipartUpload, BucketEventNotificationUpdateParamsRulesActionLifecycleDeletion:
		return true
	}
	return false
}

// The bucket jurisdiction.
type BucketEventNotificationUpdateParamsCfR2Jurisdiction string

const (
	BucketEventNotificationUpdateParamsCfR2JurisdictionDefault BucketEventNotificationUpdateParamsCfR2Jurisdiction = "default"
	BucketEventNotificationUpdateParamsCfR2JurisdictionEu      BucketEventNotificationUpdateParamsCfR2Jurisdiction = "eu"
	BucketEventNotificationUpdateParamsCfR2JurisdictionFedramp BucketEventNotificationUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketEventNotificationUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketEventNotificationUpdateParamsCfR2JurisdictionDefault, BucketEventNotificationUpdateParamsCfR2JurisdictionEu, BucketEventNotificationUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketEventNotificationUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo                 `json:"errors,required"`
	Messages []string                              `json:"messages,required"`
	Result   BucketEventNotificationUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketEventNotificationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketEventNotificationUpdateResponseEnvelopeJSON    `json:"-"`
}

// bucketEventNotificationUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [BucketEventNotificationUpdateResponseEnvelope]
type bucketEventNotificationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketEventNotificationUpdateResponseEnvelopeSuccess bool

const (
	BucketEventNotificationUpdateResponseEnvelopeSuccessTrue BucketEventNotificationUpdateResponseEnvelopeSuccess = true
)

func (r BucketEventNotificationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketEventNotificationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketEventNotificationDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// The bucket jurisdiction.
	Jurisdiction param.Field[BucketEventNotificationDeleteParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction.
type BucketEventNotificationDeleteParamsCfR2Jurisdiction string

const (
	BucketEventNotificationDeleteParamsCfR2JurisdictionDefault BucketEventNotificationDeleteParamsCfR2Jurisdiction = "default"
	BucketEventNotificationDeleteParamsCfR2JurisdictionEu      BucketEventNotificationDeleteParamsCfR2Jurisdiction = "eu"
	BucketEventNotificationDeleteParamsCfR2JurisdictionFedramp BucketEventNotificationDeleteParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketEventNotificationDeleteParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketEventNotificationDeleteParamsCfR2JurisdictionDefault, BucketEventNotificationDeleteParamsCfR2JurisdictionEu, BucketEventNotificationDeleteParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketEventNotificationDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                 `json:"errors,required"`
	Messages []string                              `json:"messages,required"`
	Result   BucketEventNotificationDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketEventNotificationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketEventNotificationDeleteResponseEnvelopeJSON    `json:"-"`
}

// bucketEventNotificationDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [BucketEventNotificationDeleteResponseEnvelope]
type bucketEventNotificationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketEventNotificationDeleteResponseEnvelopeSuccess bool

const (
	BucketEventNotificationDeleteResponseEnvelopeSuccessTrue BucketEventNotificationDeleteResponseEnvelopeSuccess = true
)

func (r BucketEventNotificationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketEventNotificationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketEventNotificationGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// The bucket jurisdiction.
	Jurisdiction param.Field[BucketEventNotificationGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction.
type BucketEventNotificationGetParamsCfR2Jurisdiction string

const (
	BucketEventNotificationGetParamsCfR2JurisdictionDefault BucketEventNotificationGetParamsCfR2Jurisdiction = "default"
	BucketEventNotificationGetParamsCfR2JurisdictionEu      BucketEventNotificationGetParamsCfR2Jurisdiction = "eu"
	BucketEventNotificationGetParamsCfR2JurisdictionFedramp BucketEventNotificationGetParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketEventNotificationGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketEventNotificationGetParamsCfR2JurisdictionDefault, BucketEventNotificationGetParamsCfR2JurisdictionEu, BucketEventNotificationGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketEventNotificationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo              `json:"errors,required"`
	Messages []string                           `json:"messages,required"`
	Result   BucketEventNotificationGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketEventNotificationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketEventNotificationGetResponseEnvelopeJSON    `json:"-"`
}

// bucketEventNotificationGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [BucketEventNotificationGetResponseEnvelope]
type bucketEventNotificationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketEventNotificationGetResponseEnvelopeSuccess bool

const (
	BucketEventNotificationGetResponseEnvelopeSuccessTrue BucketEventNotificationGetResponseEnvelopeSuccess = true
)

func (r BucketEventNotificationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketEventNotificationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

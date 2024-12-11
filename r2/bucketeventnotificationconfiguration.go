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

// BucketEventNotificationConfigurationService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketEventNotificationConfigurationService] method instead.
type BucketEventNotificationConfigurationService struct {
	Options []option.RequestOption
	Queues  *BucketEventNotificationConfigurationQueueService
}

// NewBucketEventNotificationConfigurationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewBucketEventNotificationConfigurationService(opts ...option.RequestOption) (r *BucketEventNotificationConfigurationService) {
	r = &BucketEventNotificationConfigurationService{}
	r.Options = opts
	r.Queues = NewBucketEventNotificationConfigurationQueueService(opts...)
	return
}

// List all event notification rules for a bucket.
func (r *BucketEventNotificationConfigurationService) Get(ctx context.Context, bucketName string, params BucketEventNotificationConfigurationGetParams, opts ...option.RequestOption) (res *BucketEventNotificationConfigurationGetResponse, err error) {
	var env BucketEventNotificationConfigurationGetResponseEnvelope
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

type BucketEventNotificationConfigurationGetResponse struct {
	// Name of the bucket.
	BucketName string `json:"bucketName"`
	// List of queues associated with the bucket.
	Queues []BucketEventNotificationConfigurationGetResponseQueue `json:"queues"`
	JSON   bucketEventNotificationConfigurationGetResponseJSON    `json:"-"`
}

// bucketEventNotificationConfigurationGetResponseJSON contains the JSON metadata
// for the struct [BucketEventNotificationConfigurationGetResponse]
type bucketEventNotificationConfigurationGetResponseJSON struct {
	BucketName  apijson.Field
	Queues      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationConfigurationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationConfigurationGetResponseJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationConfigurationGetResponseQueue struct {
	// Queue ID
	QueueID string `json:"queueId"`
	// Name of the queue
	QueueName string                                                      `json:"queueName"`
	Rules     []BucketEventNotificationConfigurationGetResponseQueuesRule `json:"rules"`
	JSON      bucketEventNotificationConfigurationGetResponseQueueJSON    `json:"-"`
}

// bucketEventNotificationConfigurationGetResponseQueueJSON contains the JSON
// metadata for the struct [BucketEventNotificationConfigurationGetResponseQueue]
type bucketEventNotificationConfigurationGetResponseQueueJSON struct {
	QueueID     apijson.Field
	QueueName   apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationConfigurationGetResponseQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationConfigurationGetResponseQueueJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationConfigurationGetResponseQueuesRule struct {
	// Array of R2 object actions that will trigger notifications
	Actions []BucketEventNotificationConfigurationGetResponseQueuesRulesAction `json:"actions,required"`
	// Timestamp when the rule was created
	CreatedAt string `json:"createdAt"`
	// A description that can be used to identify the event notification rule after
	// creation
	Description string `json:"description"`
	// Notifications will be sent only for objects with this prefix
	Prefix string `json:"prefix"`
	// Rule ID
	RuleID string `json:"ruleId"`
	// Notifications will be sent only for objects with this suffix
	Suffix string                                                        `json:"suffix"`
	JSON   bucketEventNotificationConfigurationGetResponseQueuesRuleJSON `json:"-"`
}

// bucketEventNotificationConfigurationGetResponseQueuesRuleJSON contains the JSON
// metadata for the struct
// [BucketEventNotificationConfigurationGetResponseQueuesRule]
type bucketEventNotificationConfigurationGetResponseQueuesRuleJSON struct {
	Actions     apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Prefix      apijson.Field
	RuleID      apijson.Field
	Suffix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationConfigurationGetResponseQueuesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationConfigurationGetResponseQueuesRuleJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationConfigurationGetResponseQueuesRulesAction string

const (
	BucketEventNotificationConfigurationGetResponseQueuesRulesActionPutObject               BucketEventNotificationConfigurationGetResponseQueuesRulesAction = "PutObject"
	BucketEventNotificationConfigurationGetResponseQueuesRulesActionCopyObject              BucketEventNotificationConfigurationGetResponseQueuesRulesAction = "CopyObject"
	BucketEventNotificationConfigurationGetResponseQueuesRulesActionDeleteObject            BucketEventNotificationConfigurationGetResponseQueuesRulesAction = "DeleteObject"
	BucketEventNotificationConfigurationGetResponseQueuesRulesActionCompleteMultipartUpload BucketEventNotificationConfigurationGetResponseQueuesRulesAction = "CompleteMultipartUpload"
	BucketEventNotificationConfigurationGetResponseQueuesRulesActionLifecycleDeletion       BucketEventNotificationConfigurationGetResponseQueuesRulesAction = "LifecycleDeletion"
)

func (r BucketEventNotificationConfigurationGetResponseQueuesRulesAction) IsKnown() bool {
	switch r {
	case BucketEventNotificationConfigurationGetResponseQueuesRulesActionPutObject, BucketEventNotificationConfigurationGetResponseQueuesRulesActionCopyObject, BucketEventNotificationConfigurationGetResponseQueuesRulesActionDeleteObject, BucketEventNotificationConfigurationGetResponseQueuesRulesActionCompleteMultipartUpload, BucketEventNotificationConfigurationGetResponseQueuesRulesActionLifecycleDeletion:
		return true
	}
	return false
}

type BucketEventNotificationConfigurationGetParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// The bucket jurisdiction
	CfR2Jurisdiction param.Field[BucketEventNotificationConfigurationGetParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction
type BucketEventNotificationConfigurationGetParamsCfR2Jurisdiction string

const (
	BucketEventNotificationConfigurationGetParamsCfR2JurisdictionDefault BucketEventNotificationConfigurationGetParamsCfR2Jurisdiction = "default"
	BucketEventNotificationConfigurationGetParamsCfR2JurisdictionEu      BucketEventNotificationConfigurationGetParamsCfR2Jurisdiction = "eu"
	BucketEventNotificationConfigurationGetParamsCfR2JurisdictionFedramp BucketEventNotificationConfigurationGetParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketEventNotificationConfigurationGetParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketEventNotificationConfigurationGetParamsCfR2JurisdictionDefault, BucketEventNotificationConfigurationGetParamsCfR2JurisdictionEu, BucketEventNotificationConfigurationGetParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketEventNotificationConfigurationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                           `json:"errors,required"`
	Messages []string                                        `json:"messages,required"`
	Result   BucketEventNotificationConfigurationGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success BucketEventNotificationConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketEventNotificationConfigurationGetResponseEnvelopeJSON    `json:"-"`
}

// bucketEventNotificationConfigurationGetResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [BucketEventNotificationConfigurationGetResponseEnvelope]
type bucketEventNotificationConfigurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationConfigurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BucketEventNotificationConfigurationGetResponseEnvelopeSuccess bool

const (
	BucketEventNotificationConfigurationGetResponseEnvelopeSuccessTrue BucketEventNotificationConfigurationGetResponseEnvelopeSuccess = true
)

func (r BucketEventNotificationConfigurationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketEventNotificationConfigurationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

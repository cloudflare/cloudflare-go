// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event_notifications

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

// R2ConfigurationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewR2ConfigurationService] method instead.
type R2ConfigurationService struct {
	Options []option.RequestOption
	Queues  *R2ConfigurationQueueService
}

// NewR2ConfigurationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewR2ConfigurationService(opts ...option.RequestOption) (r *R2ConfigurationService) {
	r = &R2ConfigurationService{}
	r.Options = opts
	r.Queues = NewR2ConfigurationQueueService(opts...)
	return
}

// List all event notification rules for a bucket.
func (r *R2ConfigurationService) Get(ctx context.Context, bucketName string, query R2ConfigurationGetParams, opts ...option.RequestOption) (res *R2ConfigurationGetResponse, err error) {
	var env R2ConfigurationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bucketName == "" {
		err = errors.New("missing required bucket_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration", query.AccountID, bucketName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type R2ConfigurationGetResponse struct {
	// Name of the bucket.
	BucketName string `json:"bucketName"`
	// List of queues associated with the bucket.
	Queues []R2ConfigurationGetResponseQueue `json:"queues"`
	JSON   r2ConfigurationGetResponseJSON    `json:"-"`
}

// r2ConfigurationGetResponseJSON contains the JSON metadata for the struct
// [R2ConfigurationGetResponse]
type r2ConfigurationGetResponseJSON struct {
	BucketName  apijson.Field
	Queues      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2ConfigurationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2ConfigurationGetResponseJSON) RawJSON() string {
	return r.raw
}

type R2ConfigurationGetResponseQueue struct {
	// Queue ID
	QueueID string `json:"queueId"`
	// Name of the queue
	QueueName string                                 `json:"queueName"`
	Rules     []R2ConfigurationGetResponseQueuesRule `json:"rules"`
	JSON      r2ConfigurationGetResponseQueueJSON    `json:"-"`
}

// r2ConfigurationGetResponseQueueJSON contains the JSON metadata for the struct
// [R2ConfigurationGetResponseQueue]
type r2ConfigurationGetResponseQueueJSON struct {
	QueueID     apijson.Field
	QueueName   apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2ConfigurationGetResponseQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2ConfigurationGetResponseQueueJSON) RawJSON() string {
	return r.raw
}

type R2ConfigurationGetResponseQueuesRule struct {
	// Array of R2 object actions that will trigger notifications
	Actions []R2ConfigurationGetResponseQueuesRulesAction `json:"actions,required"`
	// Timestamp when the rule was created
	CreatedAt string `json:"createdAt"`
	// Notifications will be sent only for objects with this prefix
	Prefix string `json:"prefix"`
	// Rule ID
	RuleID string `json:"ruleId"`
	// Notifications will be sent only for objects with this suffix
	Suffix string                                   `json:"suffix"`
	JSON   r2ConfigurationGetResponseQueuesRuleJSON `json:"-"`
}

// r2ConfigurationGetResponseQueuesRuleJSON contains the JSON metadata for the
// struct [R2ConfigurationGetResponseQueuesRule]
type r2ConfigurationGetResponseQueuesRuleJSON struct {
	Actions     apijson.Field
	CreatedAt   apijson.Field
	Prefix      apijson.Field
	RuleID      apijson.Field
	Suffix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2ConfigurationGetResponseQueuesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2ConfigurationGetResponseQueuesRuleJSON) RawJSON() string {
	return r.raw
}

type R2ConfigurationGetResponseQueuesRulesAction string

const (
	R2ConfigurationGetResponseQueuesRulesActionPutObject               R2ConfigurationGetResponseQueuesRulesAction = "PutObject"
	R2ConfigurationGetResponseQueuesRulesActionCopyObject              R2ConfigurationGetResponseQueuesRulesAction = "CopyObject"
	R2ConfigurationGetResponseQueuesRulesActionDeleteObject            R2ConfigurationGetResponseQueuesRulesAction = "DeleteObject"
	R2ConfigurationGetResponseQueuesRulesActionCompleteMultipartUpload R2ConfigurationGetResponseQueuesRulesAction = "CompleteMultipartUpload"
	R2ConfigurationGetResponseQueuesRulesActionLifecycleDeletion       R2ConfigurationGetResponseQueuesRulesAction = "LifecycleDeletion"
)

func (r R2ConfigurationGetResponseQueuesRulesAction) IsKnown() bool {
	switch r {
	case R2ConfigurationGetResponseQueuesRulesActionPutObject, R2ConfigurationGetResponseQueuesRulesActionCopyObject, R2ConfigurationGetResponseQueuesRulesActionDeleteObject, R2ConfigurationGetResponseQueuesRulesActionCompleteMultipartUpload, R2ConfigurationGetResponseQueuesRulesActionLifecycleDeletion:
		return true
	}
	return false
}

type R2ConfigurationGetParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type R2ConfigurationGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []string                   `json:"messages,required"`
	Result   R2ConfigurationGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success R2ConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    r2ConfigurationGetResponseEnvelopeJSON    `json:"-"`
}

// r2ConfigurationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [R2ConfigurationGetResponseEnvelope]
type r2ConfigurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2ConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2ConfigurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type R2ConfigurationGetResponseEnvelopeSuccess bool

const (
	R2ConfigurationGetResponseEnvelopeSuccessTrue R2ConfigurationGetResponseEnvelopeSuccess = true
)

func (r R2ConfigurationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case R2ConfigurationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

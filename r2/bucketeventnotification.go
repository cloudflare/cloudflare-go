// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/tidwall/gjson"
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

// List all event notification rules for a bucket.
func (r *BucketEventNotificationService) List(ctx context.Context, bucketName string, params BucketEventNotificationListParams, opts ...option.RequestOption) (res *BucketEventNotificationListResponse, err error) {
	var env BucketEventNotificationListResponseEnvelope
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

// Get a single event notification rule.
func (r *BucketEventNotificationService) Get(ctx context.Context, bucketName string, queueID string, params BucketEventNotificationGetParams, opts ...option.RequestOption) (res *BucketEventNotificationGetResponse, err error) {
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
	if queueID == "" {
		err = errors.New("missing required queue_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration/queues/%s", params.AccountID, bucketName, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BucketEventNotificationUpdateResponse = interface{}

type BucketEventNotificationListResponse struct {
	// Name of the bucket.
	BucketName string `json:"bucketName"`
	// List of queues associated with the bucket.
	Queues []BucketEventNotificationListResponseQueue `json:"queues"`
	JSON   bucketEventNotificationListResponseJSON    `json:"-"`
}

// bucketEventNotificationListResponseJSON contains the JSON metadata for the
// struct [BucketEventNotificationListResponse]
type bucketEventNotificationListResponseJSON struct {
	BucketName  apijson.Field
	Queues      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationListResponseJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationListResponseQueue struct {
	// Queue ID.
	QueueID string `json:"queueId"`
	// Name of the queue.
	QueueName string                                          `json:"queueName"`
	Rules     []BucketEventNotificationListResponseQueuesRule `json:"rules"`
	JSON      bucketEventNotificationListResponseQueueJSON    `json:"-"`
}

// bucketEventNotificationListResponseQueueJSON contains the JSON metadata for the
// struct [BucketEventNotificationListResponseQueue]
type bucketEventNotificationListResponseQueueJSON struct {
	QueueID     apijson.Field
	QueueName   apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationListResponseQueue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationListResponseQueueJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationListResponseQueuesRule struct {
	// Array of R2 object actions that will trigger notifications.
	Actions []BucketEventNotificationListResponseQueuesRulesAction `json:"actions,required"`
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
	Suffix string                                            `json:"suffix"`
	JSON   bucketEventNotificationListResponseQueuesRuleJSON `json:"-"`
}

// bucketEventNotificationListResponseQueuesRuleJSON contains the JSON metadata for
// the struct [BucketEventNotificationListResponseQueuesRule]
type bucketEventNotificationListResponseQueuesRuleJSON struct {
	Actions     apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Prefix      apijson.Field
	RuleID      apijson.Field
	Suffix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationListResponseQueuesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationListResponseQueuesRuleJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationListResponseQueuesRulesAction string

const (
	BucketEventNotificationListResponseQueuesRulesActionPutObject               BucketEventNotificationListResponseQueuesRulesAction = "PutObject"
	BucketEventNotificationListResponseQueuesRulesActionCopyObject              BucketEventNotificationListResponseQueuesRulesAction = "CopyObject"
	BucketEventNotificationListResponseQueuesRulesActionDeleteObject            BucketEventNotificationListResponseQueuesRulesAction = "DeleteObject"
	BucketEventNotificationListResponseQueuesRulesActionCompleteMultipartUpload BucketEventNotificationListResponseQueuesRulesAction = "CompleteMultipartUpload"
	BucketEventNotificationListResponseQueuesRulesActionLifecycleDeletion       BucketEventNotificationListResponseQueuesRulesAction = "LifecycleDeletion"
)

func (r BucketEventNotificationListResponseQueuesRulesAction) IsKnown() bool {
	switch r {
	case BucketEventNotificationListResponseQueuesRulesActionPutObject, BucketEventNotificationListResponseQueuesRulesActionCopyObject, BucketEventNotificationListResponseQueuesRulesActionDeleteObject, BucketEventNotificationListResponseQueuesRulesActionCompleteMultipartUpload, BucketEventNotificationListResponseQueuesRulesActionLifecycleDeletion:
		return true
	}
	return false
}

type BucketEventNotificationDeleteResponse = interface{}

type BucketEventNotificationGetResponse struct {
	// Unique identifier for this rule.
	ID string `json:"id,required"`
	// Conditions that apply to all transitions of this rule.
	Conditions BucketEventNotificationGetResponseConditions `json:"conditions,required"`
	// Whether or not this rule is in effect.
	Enabled bool `json:"enabled,required"`
	// Transition to abort ongoing multipart uploads.
	AbortMultipartUploadsTransition BucketEventNotificationGetResponseAbortMultipartUploadsTransition `json:"abortMultipartUploadsTransition"`
	// Transition to delete objects.
	DeleteObjectsTransition BucketEventNotificationGetResponseDeleteObjectsTransition `json:"deleteObjectsTransition"`
	// Transitions to change the storage class of objects.
	StorageClassTransitions []BucketEventNotificationGetResponseStorageClassTransition `json:"storageClassTransitions"`
	JSON                    bucketEventNotificationGetResponseJSON                     `json:"-"`
}

// bucketEventNotificationGetResponseJSON contains the JSON metadata for the struct
// [BucketEventNotificationGetResponse]
type bucketEventNotificationGetResponseJSON struct {
	ID                              apijson.Field
	Conditions                      apijson.Field
	Enabled                         apijson.Field
	AbortMultipartUploadsTransition apijson.Field
	DeleteObjectsTransition         apijson.Field
	StorageClassTransitions         apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseJSON) RawJSON() string {
	return r.raw
}

// Conditions that apply to all transitions of this rule.
type BucketEventNotificationGetResponseConditions struct {
	// Transitions will only apply to objects/uploads in the bucket that start with the
	// given prefix, an empty prefix can be provided to scope rule to all
	// objects/uploads.
	Prefix string                                           `json:"prefix,required"`
	JSON   bucketEventNotificationGetResponseConditionsJSON `json:"-"`
}

// bucketEventNotificationGetResponseConditionsJSON contains the JSON metadata for
// the struct [BucketEventNotificationGetResponseConditions]
type bucketEventNotificationGetResponseConditionsJSON struct {
	Prefix      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponseConditions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseConditionsJSON) RawJSON() string {
	return r.raw
}

// Transition to abort ongoing multipart uploads.
type BucketEventNotificationGetResponseAbortMultipartUploadsTransition struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition BucketEventNotificationGetResponseAbortMultipartUploadsTransitionCondition `json:"condition"`
	JSON      bucketEventNotificationGetResponseAbortMultipartUploadsTransitionJSON      `json:"-"`
}

// bucketEventNotificationGetResponseAbortMultipartUploadsTransitionJSON contains
// the JSON metadata for the struct
// [BucketEventNotificationGetResponseAbortMultipartUploadsTransition]
type bucketEventNotificationGetResponseAbortMultipartUploadsTransitionJSON struct {
	Condition   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponseAbortMultipartUploadsTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseAbortMultipartUploadsTransitionJSON) RawJSON() string {
	return r.raw
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketEventNotificationGetResponseAbortMultipartUploadsTransitionCondition struct {
	MaxAge int64                                                                          `json:"maxAge,required"`
	Type   BucketEventNotificationGetResponseAbortMultipartUploadsTransitionConditionType `json:"type,required"`
	JSON   bucketEventNotificationGetResponseAbortMultipartUploadsTransitionConditionJSON `json:"-"`
}

// bucketEventNotificationGetResponseAbortMultipartUploadsTransitionConditionJSON
// contains the JSON metadata for the struct
// [BucketEventNotificationGetResponseAbortMultipartUploadsTransitionCondition]
type bucketEventNotificationGetResponseAbortMultipartUploadsTransitionConditionJSON struct {
	MaxAge      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponseAbortMultipartUploadsTransitionCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseAbortMultipartUploadsTransitionConditionJSON) RawJSON() string {
	return r.raw
}

type BucketEventNotificationGetResponseAbortMultipartUploadsTransitionConditionType string

const (
	BucketEventNotificationGetResponseAbortMultipartUploadsTransitionConditionTypeAge BucketEventNotificationGetResponseAbortMultipartUploadsTransitionConditionType = "Age"
)

func (r BucketEventNotificationGetResponseAbortMultipartUploadsTransitionConditionType) IsKnown() bool {
	switch r {
	case BucketEventNotificationGetResponseAbortMultipartUploadsTransitionConditionTypeAge:
		return true
	}
	return false
}

// Transition to delete objects.
type BucketEventNotificationGetResponseDeleteObjectsTransition struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition BucketEventNotificationGetResponseDeleteObjectsTransitionCondition `json:"condition"`
	JSON      bucketEventNotificationGetResponseDeleteObjectsTransitionJSON      `json:"-"`
}

// bucketEventNotificationGetResponseDeleteObjectsTransitionJSON contains the JSON
// metadata for the struct
// [BucketEventNotificationGetResponseDeleteObjectsTransition]
type bucketEventNotificationGetResponseDeleteObjectsTransitionJSON struct {
	Condition   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponseDeleteObjectsTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseDeleteObjectsTransitionJSON) RawJSON() string {
	return r.raw
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketEventNotificationGetResponseDeleteObjectsTransitionCondition struct {
	Type   BucketEventNotificationGetResponseDeleteObjectsTransitionConditionType `json:"type,required"`
	Date   time.Time                                                              `json:"date" format:"date"`
	MaxAge int64                                                                  `json:"maxAge"`
	JSON   bucketEventNotificationGetResponseDeleteObjectsTransitionConditionJSON `json:"-"`
	union  BucketEventNotificationGetResponseDeleteObjectsTransitionConditionUnion
}

// bucketEventNotificationGetResponseDeleteObjectsTransitionConditionJSON contains
// the JSON metadata for the struct
// [BucketEventNotificationGetResponseDeleteObjectsTransitionCondition]
type bucketEventNotificationGetResponseDeleteObjectsTransitionConditionJSON struct {
	Type        apijson.Field
	Date        apijson.Field
	MaxAge      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r bucketEventNotificationGetResponseDeleteObjectsTransitionConditionJSON) RawJSON() string {
	return r.raw
}

func (r *BucketEventNotificationGetResponseDeleteObjectsTransitionCondition) UnmarshalJSON(data []byte) (err error) {
	*r = BucketEventNotificationGetResponseDeleteObjectsTransitionCondition{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [BucketEventNotificationGetResponseDeleteObjectsTransitionConditionUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeCondition],
// [BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateCondition].
func (r BucketEventNotificationGetResponseDeleteObjectsTransitionCondition) AsUnion() BucketEventNotificationGetResponseDeleteObjectsTransitionConditionUnion {
	return r.union
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
//
// Union satisfied by
// [BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeCondition]
// or
// [BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateCondition].
type BucketEventNotificationGetResponseDeleteObjectsTransitionConditionUnion interface {
	implementsBucketEventNotificationGetResponseDeleteObjectsTransitionCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BucketEventNotificationGetResponseDeleteObjectsTransitionConditionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateCondition{}),
		},
	)
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeCondition struct {
	MaxAge int64                                                                                         `json:"maxAge,required"`
	Type   BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeConditionType `json:"type,required"`
	JSON   bucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeConditionJSON `json:"-"`
}

// bucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeConditionJSON
// contains the JSON metadata for the struct
// [BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeCondition]
type bucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeConditionJSON struct {
	MaxAge      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeConditionJSON) RawJSON() string {
	return r.raw
}

func (r BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeCondition) implementsBucketEventNotificationGetResponseDeleteObjectsTransitionCondition() {
}

type BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeConditionType string

const (
	BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeConditionTypeAge BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeConditionType = "Age"
)

func (r BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeConditionType) IsKnown() bool {
	switch r {
	case BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleAgeConditionTypeAge:
		return true
	}
	return false
}

// Condition for lifecycle transitions to apply on a specific date.
type BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateCondition struct {
	Date time.Time                                                                                      `json:"date,required" format:"date"`
	Type BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateConditionType `json:"type,required"`
	JSON bucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateConditionJSON `json:"-"`
}

// bucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateConditionJSON
// contains the JSON metadata for the struct
// [BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateCondition]
type bucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateConditionJSON struct {
	Date        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateConditionJSON) RawJSON() string {
	return r.raw
}

func (r BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateCondition) implementsBucketEventNotificationGetResponseDeleteObjectsTransitionCondition() {
}

type BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateConditionType string

const (
	BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateConditionTypeDate BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateConditionType = "Date"
)

func (r BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateConditionType) IsKnown() bool {
	switch r {
	case BucketEventNotificationGetResponseDeleteObjectsTransitionConditionR2LifecycleDateConditionTypeDate:
		return true
	}
	return false
}

type BucketEventNotificationGetResponseDeleteObjectsTransitionConditionType string

const (
	BucketEventNotificationGetResponseDeleteObjectsTransitionConditionTypeAge  BucketEventNotificationGetResponseDeleteObjectsTransitionConditionType = "Age"
	BucketEventNotificationGetResponseDeleteObjectsTransitionConditionTypeDate BucketEventNotificationGetResponseDeleteObjectsTransitionConditionType = "Date"
)

func (r BucketEventNotificationGetResponseDeleteObjectsTransitionConditionType) IsKnown() bool {
	switch r {
	case BucketEventNotificationGetResponseDeleteObjectsTransitionConditionTypeAge, BucketEventNotificationGetResponseDeleteObjectsTransitionConditionTypeDate:
		return true
	}
	return false
}

type BucketEventNotificationGetResponseStorageClassTransition struct {
	// Condition for lifecycle transitions to apply after an object reaches an age in
	// seconds.
	Condition    BucketEventNotificationGetResponseStorageClassTransitionsCondition    `json:"condition,required"`
	StorageClass BucketEventNotificationGetResponseStorageClassTransitionsStorageClass `json:"storageClass,required"`
	JSON         bucketEventNotificationGetResponseStorageClassTransitionJSON          `json:"-"`
}

// bucketEventNotificationGetResponseStorageClassTransitionJSON contains the JSON
// metadata for the struct
// [BucketEventNotificationGetResponseStorageClassTransition]
type bucketEventNotificationGetResponseStorageClassTransitionJSON struct {
	Condition    apijson.Field
	StorageClass apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponseStorageClassTransition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseStorageClassTransitionJSON) RawJSON() string {
	return r.raw
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketEventNotificationGetResponseStorageClassTransitionsCondition struct {
	Type   BucketEventNotificationGetResponseStorageClassTransitionsConditionType `json:"type,required"`
	Date   time.Time                                                              `json:"date" format:"date"`
	MaxAge int64                                                                  `json:"maxAge"`
	JSON   bucketEventNotificationGetResponseStorageClassTransitionsConditionJSON `json:"-"`
	union  BucketEventNotificationGetResponseStorageClassTransitionsConditionUnion
}

// bucketEventNotificationGetResponseStorageClassTransitionsConditionJSON contains
// the JSON metadata for the struct
// [BucketEventNotificationGetResponseStorageClassTransitionsCondition]
type bucketEventNotificationGetResponseStorageClassTransitionsConditionJSON struct {
	Type        apijson.Field
	Date        apijson.Field
	MaxAge      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r bucketEventNotificationGetResponseStorageClassTransitionsConditionJSON) RawJSON() string {
	return r.raw
}

func (r *BucketEventNotificationGetResponseStorageClassTransitionsCondition) UnmarshalJSON(data []byte) (err error) {
	*r = BucketEventNotificationGetResponseStorageClassTransitionsCondition{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [BucketEventNotificationGetResponseStorageClassTransitionsConditionUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeCondition],
// [BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateCondition].
func (r BucketEventNotificationGetResponseStorageClassTransitionsCondition) AsUnion() BucketEventNotificationGetResponseStorageClassTransitionsConditionUnion {
	return r.union
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
//
// Union satisfied by
// [BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeCondition]
// or
// [BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateCondition].
type BucketEventNotificationGetResponseStorageClassTransitionsConditionUnion interface {
	implementsBucketEventNotificationGetResponseStorageClassTransitionsCondition()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*BucketEventNotificationGetResponseStorageClassTransitionsConditionUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeCondition{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateCondition{}),
		},
	)
}

// Condition for lifecycle transitions to apply after an object reaches an age in
// seconds.
type BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeCondition struct {
	MaxAge int64                                                                                         `json:"maxAge,required"`
	Type   BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeConditionType `json:"type,required"`
	JSON   bucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeConditionJSON `json:"-"`
}

// bucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeConditionJSON
// contains the JSON metadata for the struct
// [BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeCondition]
type bucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeConditionJSON struct {
	MaxAge      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeConditionJSON) RawJSON() string {
	return r.raw
}

func (r BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeCondition) implementsBucketEventNotificationGetResponseStorageClassTransitionsCondition() {
}

type BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeConditionType string

const (
	BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeConditionTypeAge BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeConditionType = "Age"
)

func (r BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeConditionType) IsKnown() bool {
	switch r {
	case BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleAgeConditionTypeAge:
		return true
	}
	return false
}

// Condition for lifecycle transitions to apply on a specific date.
type BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateCondition struct {
	Date time.Time                                                                                      `json:"date,required" format:"date"`
	Type BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateConditionType `json:"type,required"`
	JSON bucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateConditionJSON `json:"-"`
}

// bucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateConditionJSON
// contains the JSON metadata for the struct
// [BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateCondition]
type bucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateConditionJSON struct {
	Date        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateConditionJSON) RawJSON() string {
	return r.raw
}

func (r BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateCondition) implementsBucketEventNotificationGetResponseStorageClassTransitionsCondition() {
}

type BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateConditionType string

const (
	BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateConditionTypeDate BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateConditionType = "Date"
)

func (r BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateConditionType) IsKnown() bool {
	switch r {
	case BucketEventNotificationGetResponseStorageClassTransitionsConditionR2LifecycleDateConditionTypeDate:
		return true
	}
	return false
}

type BucketEventNotificationGetResponseStorageClassTransitionsConditionType string

const (
	BucketEventNotificationGetResponseStorageClassTransitionsConditionTypeAge  BucketEventNotificationGetResponseStorageClassTransitionsConditionType = "Age"
	BucketEventNotificationGetResponseStorageClassTransitionsConditionTypeDate BucketEventNotificationGetResponseStorageClassTransitionsConditionType = "Date"
)

func (r BucketEventNotificationGetResponseStorageClassTransitionsConditionType) IsKnown() bool {
	switch r {
	case BucketEventNotificationGetResponseStorageClassTransitionsConditionTypeAge, BucketEventNotificationGetResponseStorageClassTransitionsConditionTypeDate:
		return true
	}
	return false
}

type BucketEventNotificationGetResponseStorageClassTransitionsStorageClass string

const (
	BucketEventNotificationGetResponseStorageClassTransitionsStorageClassInfrequentAccess BucketEventNotificationGetResponseStorageClassTransitionsStorageClass = "InfrequentAccess"
)

func (r BucketEventNotificationGetResponseStorageClassTransitionsStorageClass) IsKnown() bool {
	switch r {
	case BucketEventNotificationGetResponseStorageClassTransitionsStorageClassInfrequentAccess:
		return true
	}
	return false
}

type BucketEventNotificationUpdateParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Array of rules to drive notifications.
	Rules param.Field[[]BucketEventNotificationUpdateParamsRule] `json:"rules"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
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

// Jurisdiction where objects in this bucket are guaranteed to be stored.
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

type BucketEventNotificationListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketEventNotificationListParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
type BucketEventNotificationListParamsCfR2Jurisdiction string

const (
	BucketEventNotificationListParamsCfR2JurisdictionDefault BucketEventNotificationListParamsCfR2Jurisdiction = "default"
	BucketEventNotificationListParamsCfR2JurisdictionEu      BucketEventNotificationListParamsCfR2Jurisdiction = "eu"
	BucketEventNotificationListParamsCfR2JurisdictionFedramp BucketEventNotificationListParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketEventNotificationListParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketEventNotificationListParamsCfR2JurisdictionDefault, BucketEventNotificationListParamsCfR2JurisdictionEu, BucketEventNotificationListParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketEventNotificationListResponseEnvelope struct {
	Errors   []shared.ResponseInfo               `json:"errors,required"`
	Messages []string                            `json:"messages,required"`
	Result   BucketEventNotificationListResponse `json:"result,required"`
	// Whether the API call was successful.
	Success BucketEventNotificationListResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketEventNotificationListResponseEnvelopeJSON    `json:"-"`
}

// bucketEventNotificationListResponseEnvelopeJSON contains the JSON metadata for
// the struct [BucketEventNotificationListResponseEnvelope]
type bucketEventNotificationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type BucketEventNotificationListResponseEnvelopeSuccess bool

const (
	BucketEventNotificationListResponseEnvelopeSuccessTrue BucketEventNotificationListResponseEnvelopeSuccess = true
)

func (r BucketEventNotificationListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketEventNotificationListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketEventNotificationDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Jurisdiction where objects in this bucket are guaranteed to be stored.
	Jurisdiction param.Field[BucketEventNotificationDeleteParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// Jurisdiction where objects in this bucket are guaranteed to be stored.
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

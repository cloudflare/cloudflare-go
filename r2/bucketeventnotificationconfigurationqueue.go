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

// BucketEventNotificationConfigurationQueueService contains methods and other
// services that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketEventNotificationConfigurationQueueService] method instead.
type BucketEventNotificationConfigurationQueueService struct {
	Options []option.RequestOption
}

// NewBucketEventNotificationConfigurationQueueService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewBucketEventNotificationConfigurationQueueService(opts ...option.RequestOption) (r *BucketEventNotificationConfigurationQueueService) {
	r = &BucketEventNotificationConfigurationQueueService{}
	r.Options = opts
	return
}

// Create event notification rule.
func (r *BucketEventNotificationConfigurationQueueService) Update(ctx context.Context, bucketName string, queueID string, params BucketEventNotificationConfigurationQueueUpdateParams, opts ...option.RequestOption) (res *BucketEventNotificationConfigurationQueueUpdateResponse, err error) {
	var env BucketEventNotificationConfigurationQueueUpdateResponseEnvelope
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
func (r *BucketEventNotificationConfigurationQueueService) Delete(ctx context.Context, bucketName string, queueID string, params BucketEventNotificationConfigurationQueueDeleteParams, opts ...option.RequestOption) (res *BucketEventNotificationConfigurationQueueDeleteResponse, err error) {
	var env BucketEventNotificationConfigurationQueueDeleteResponseEnvelope
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

type BucketEventNotificationConfigurationQueueUpdateResponse = interface{}

type BucketEventNotificationConfigurationQueueDeleteResponse = interface{}

type BucketEventNotificationConfigurationQueueUpdateParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// Array of rules to drive notifications
	Rules param.Field[[]BucketEventNotificationConfigurationQueueUpdateParamsRule] `json:"rules"`
	// The bucket jurisdiction
	CfR2Jurisdiction param.Field[BucketEventNotificationConfigurationQueueUpdateParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

func (r BucketEventNotificationConfigurationQueueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BucketEventNotificationConfigurationQueueUpdateParamsRule struct {
	// Array of R2 object actions that will trigger notifications
	Actions param.Field[[]BucketEventNotificationConfigurationQueueUpdateParamsRulesAction] `json:"actions,required"`
	// A description that can be used to identify the event notification rule after
	// creation
	Description param.Field[string] `json:"description"`
	// Notifications will be sent only for objects with this prefix
	Prefix param.Field[string] `json:"prefix"`
	// Notifications will be sent only for objects with this suffix
	Suffix param.Field[string] `json:"suffix"`
}

func (r BucketEventNotificationConfigurationQueueUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BucketEventNotificationConfigurationQueueUpdateParamsRulesAction string

const (
	BucketEventNotificationConfigurationQueueUpdateParamsRulesActionPutObject               BucketEventNotificationConfigurationQueueUpdateParamsRulesAction = "PutObject"
	BucketEventNotificationConfigurationQueueUpdateParamsRulesActionCopyObject              BucketEventNotificationConfigurationQueueUpdateParamsRulesAction = "CopyObject"
	BucketEventNotificationConfigurationQueueUpdateParamsRulesActionDeleteObject            BucketEventNotificationConfigurationQueueUpdateParamsRulesAction = "DeleteObject"
	BucketEventNotificationConfigurationQueueUpdateParamsRulesActionCompleteMultipartUpload BucketEventNotificationConfigurationQueueUpdateParamsRulesAction = "CompleteMultipartUpload"
	BucketEventNotificationConfigurationQueueUpdateParamsRulesActionLifecycleDeletion       BucketEventNotificationConfigurationQueueUpdateParamsRulesAction = "LifecycleDeletion"
)

func (r BucketEventNotificationConfigurationQueueUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case BucketEventNotificationConfigurationQueueUpdateParamsRulesActionPutObject, BucketEventNotificationConfigurationQueueUpdateParamsRulesActionCopyObject, BucketEventNotificationConfigurationQueueUpdateParamsRulesActionDeleteObject, BucketEventNotificationConfigurationQueueUpdateParamsRulesActionCompleteMultipartUpload, BucketEventNotificationConfigurationQueueUpdateParamsRulesActionLifecycleDeletion:
		return true
	}
	return false
}

// The bucket jurisdiction
type BucketEventNotificationConfigurationQueueUpdateParamsCfR2Jurisdiction string

const (
	BucketEventNotificationConfigurationQueueUpdateParamsCfR2JurisdictionDefault BucketEventNotificationConfigurationQueueUpdateParamsCfR2Jurisdiction = "default"
	BucketEventNotificationConfigurationQueueUpdateParamsCfR2JurisdictionEu      BucketEventNotificationConfigurationQueueUpdateParamsCfR2Jurisdiction = "eu"
	BucketEventNotificationConfigurationQueueUpdateParamsCfR2JurisdictionFedramp BucketEventNotificationConfigurationQueueUpdateParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketEventNotificationConfigurationQueueUpdateParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketEventNotificationConfigurationQueueUpdateParamsCfR2JurisdictionDefault, BucketEventNotificationConfigurationQueueUpdateParamsCfR2JurisdictionEu, BucketEventNotificationConfigurationQueueUpdateParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketEventNotificationConfigurationQueueUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo                                   `json:"errors,required"`
	Messages []string                                                `json:"messages,required"`
	Result   BucketEventNotificationConfigurationQueueUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success BucketEventNotificationConfigurationQueueUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketEventNotificationConfigurationQueueUpdateResponseEnvelopeJSON    `json:"-"`
}

// bucketEventNotificationConfigurationQueueUpdateResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [BucketEventNotificationConfigurationQueueUpdateResponseEnvelope]
type bucketEventNotificationConfigurationQueueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationConfigurationQueueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationConfigurationQueueUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BucketEventNotificationConfigurationQueueUpdateResponseEnvelopeSuccess bool

const (
	BucketEventNotificationConfigurationQueueUpdateResponseEnvelopeSuccessTrue BucketEventNotificationConfigurationQueueUpdateResponseEnvelopeSuccess = true
)

func (r BucketEventNotificationConfigurationQueueUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketEventNotificationConfigurationQueueUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type BucketEventNotificationConfigurationQueueDeleteParams struct {
	// Account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// The bucket jurisdiction
	CfR2Jurisdiction param.Field[BucketEventNotificationConfigurationQueueDeleteParamsCfR2Jurisdiction] `header:"cf-r2-jurisdiction"`
}

// The bucket jurisdiction
type BucketEventNotificationConfigurationQueueDeleteParamsCfR2Jurisdiction string

const (
	BucketEventNotificationConfigurationQueueDeleteParamsCfR2JurisdictionDefault BucketEventNotificationConfigurationQueueDeleteParamsCfR2Jurisdiction = "default"
	BucketEventNotificationConfigurationQueueDeleteParamsCfR2JurisdictionEu      BucketEventNotificationConfigurationQueueDeleteParamsCfR2Jurisdiction = "eu"
	BucketEventNotificationConfigurationQueueDeleteParamsCfR2JurisdictionFedramp BucketEventNotificationConfigurationQueueDeleteParamsCfR2Jurisdiction = "fedramp"
)

func (r BucketEventNotificationConfigurationQueueDeleteParamsCfR2Jurisdiction) IsKnown() bool {
	switch r {
	case BucketEventNotificationConfigurationQueueDeleteParamsCfR2JurisdictionDefault, BucketEventNotificationConfigurationQueueDeleteParamsCfR2JurisdictionEu, BucketEventNotificationConfigurationQueueDeleteParamsCfR2JurisdictionFedramp:
		return true
	}
	return false
}

type BucketEventNotificationConfigurationQueueDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                                   `json:"errors,required"`
	Messages []string                                                `json:"messages,required"`
	Result   BucketEventNotificationConfigurationQueueDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success BucketEventNotificationConfigurationQueueDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    bucketEventNotificationConfigurationQueueDeleteResponseEnvelopeJSON    `json:"-"`
}

// bucketEventNotificationConfigurationQueueDeleteResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [BucketEventNotificationConfigurationQueueDeleteResponseEnvelope]
type bucketEventNotificationConfigurationQueueDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BucketEventNotificationConfigurationQueueDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bucketEventNotificationConfigurationQueueDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BucketEventNotificationConfigurationQueueDeleteResponseEnvelopeSuccess bool

const (
	BucketEventNotificationConfigurationQueueDeleteResponseEnvelopeSuccessTrue BucketEventNotificationConfigurationQueueDeleteResponseEnvelopeSuccess = true
)

func (r BucketEventNotificationConfigurationQueueDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BucketEventNotificationConfigurationQueueDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

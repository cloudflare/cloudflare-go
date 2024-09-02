// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package event_notifications

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// R2ConfigurationQueueService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewR2ConfigurationQueueService] method instead.
type R2ConfigurationQueueService struct {
	Options []option.RequestOption
}

// NewR2ConfigurationQueueService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewR2ConfigurationQueueService(opts ...option.RequestOption) (r *R2ConfigurationQueueService) {
	r = &R2ConfigurationQueueService{}
	r.Options = opts
	return
}

// Define the rules for a given queue which will determine event notification
// production.
func (r *R2ConfigurationQueueService) Update(ctx context.Context, bucketName string, queueID string, params R2ConfigurationQueueUpdateParams, opts ...option.RequestOption) (res *R2ConfigurationQueueUpdateResponse, err error) {
	var env R2ConfigurationQueueUpdateResponseEnvelope
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

// Turn off all event notifications configured for delivery to a given queue. No
// further notifications will be produced for the queue once complete.
func (r *R2ConfigurationQueueService) Delete(ctx context.Context, bucketName string, queueID string, body R2ConfigurationQueueDeleteParams, opts ...option.RequestOption) (res *R2ConfigurationQueueDeleteResponseUnion, err error) {
	var env R2ConfigurationQueueDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
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
	path := fmt.Sprintf("accounts/%s/event_notifications/r2/%s/configuration/queues/%s", body.AccountID, bucketName, queueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type R2ConfigurationQueueUpdateResponse struct {
	EventNotificationDetailID string                                 `json:"event_notification_detail_id"`
	JSON                      r2ConfigurationQueueUpdateResponseJSON `json:"-"`
}

// r2ConfigurationQueueUpdateResponseJSON contains the JSON metadata for the struct
// [R2ConfigurationQueueUpdateResponse]
type r2ConfigurationQueueUpdateResponseJSON struct {
	EventNotificationDetailID apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *R2ConfigurationQueueUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2ConfigurationQueueUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by
// [event_notifications.R2ConfigurationQueueDeleteResponseUnknown],
// [event_notifications.R2ConfigurationQueueDeleteResponseArray] or
// [shared.UnionString].
type R2ConfigurationQueueDeleteResponseUnion interface {
	ImplementsEventNotificationsR2ConfigurationQueueDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*R2ConfigurationQueueDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(R2ConfigurationQueueDeleteResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type R2ConfigurationQueueDeleteResponseArray []interface{}

func (r R2ConfigurationQueueDeleteResponseArray) ImplementsEventNotificationsR2ConfigurationQueueDeleteResponseUnion() {
}

type R2ConfigurationQueueUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Array of rules to drive notifications
	Rules param.Field[[]R2ConfigurationQueueUpdateParamsRule] `json:"rules"`
}

func (r R2ConfigurationQueueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type R2ConfigurationQueueUpdateParamsRule struct {
	// Array of R2 object actions that will trigger notifications
	Actions param.Field[[]R2ConfigurationQueueUpdateParamsRulesAction] `json:"actions,required"`
	// Notifications will be sent only for objects with this prefix
	Prefix param.Field[string] `json:"prefix"`
	// Notifications will be sent only for objects with this suffix
	Suffix param.Field[string] `json:"suffix"`
}

func (r R2ConfigurationQueueUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type R2ConfigurationQueueUpdateParamsRulesAction string

const (
	R2ConfigurationQueueUpdateParamsRulesActionPutObject               R2ConfigurationQueueUpdateParamsRulesAction = "PutObject"
	R2ConfigurationQueueUpdateParamsRulesActionCopyObject              R2ConfigurationQueueUpdateParamsRulesAction = "CopyObject"
	R2ConfigurationQueueUpdateParamsRulesActionDeleteObject            R2ConfigurationQueueUpdateParamsRulesAction = "DeleteObject"
	R2ConfigurationQueueUpdateParamsRulesActionCompleteMultipartUpload R2ConfigurationQueueUpdateParamsRulesAction = "CompleteMultipartUpload"
	R2ConfigurationQueueUpdateParamsRulesActionAbortMultipartUpload    R2ConfigurationQueueUpdateParamsRulesAction = "AbortMultipartUpload"
)

func (r R2ConfigurationQueueUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case R2ConfigurationQueueUpdateParamsRulesActionPutObject, R2ConfigurationQueueUpdateParamsRulesActionCopyObject, R2ConfigurationQueueUpdateParamsRulesActionDeleteObject, R2ConfigurationQueueUpdateParamsRulesActionCompleteMultipartUpload, R2ConfigurationQueueUpdateParamsRulesActionAbortMultipartUpload:
		return true
	}
	return false
}

type R2ConfigurationQueueUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo              `json:"errors,required"`
	Messages []shared.ResponseInfo              `json:"messages,required"`
	Result   R2ConfigurationQueueUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success R2ConfigurationQueueUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    r2ConfigurationQueueUpdateResponseEnvelopeJSON    `json:"-"`
}

// r2ConfigurationQueueUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [R2ConfigurationQueueUpdateResponseEnvelope]
type r2ConfigurationQueueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2ConfigurationQueueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2ConfigurationQueueUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type R2ConfigurationQueueUpdateResponseEnvelopeSuccess bool

const (
	R2ConfigurationQueueUpdateResponseEnvelopeSuccessTrue R2ConfigurationQueueUpdateResponseEnvelopeSuccess = true
)

func (r R2ConfigurationQueueUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case R2ConfigurationQueueUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type R2ConfigurationQueueDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type R2ConfigurationQueueDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                   `json:"errors,required"`
	Messages []shared.ResponseInfo                   `json:"messages,required"`
	Result   R2ConfigurationQueueDeleteResponseUnion `json:"result,required"`
	// Whether the API call was successful.
	Success R2ConfigurationQueueDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    r2ConfigurationQueueDeleteResponseEnvelopeJSON    `json:"-"`
}

// r2ConfigurationQueueDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [R2ConfigurationQueueDeleteResponseEnvelope]
type r2ConfigurationQueueDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *R2ConfigurationQueueDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r r2ConfigurationQueueDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type R2ConfigurationQueueDeleteResponseEnvelopeSuccess bool

const (
	R2ConfigurationQueueDeleteResponseEnvelopeSuccessTrue R2ConfigurationQueueDeleteResponseEnvelopeSuccess = true
)

func (r R2ConfigurationQueueDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case R2ConfigurationQueueDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

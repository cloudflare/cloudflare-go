// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// CasbPostureWebhookJobService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbPostureWebhookJobService] method instead.
type CasbPostureWebhookJobService struct {
	Options []option.RequestOption
}

// NewCasbPostureWebhookJobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCasbPostureWebhookJobService(opts ...option.RequestOption) (r *CasbPostureWebhookJobService) {
	r = &CasbPostureWebhookJobService{}
	r.Options = opts
	return
}

// Creates webhook jobs to send a finding instance to one or more configured
// webhooks.
func (r *CasbPostureWebhookJobService) New(ctx context.Context, params CasbPostureWebhookJobNewParams, opts ...option.RequestOption) (res *CasbPostureWebhookJobNewResponse, err error) {
	var env CasbPostureWebhookJobNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/webhooks/jobs", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type CasbPostureWebhookJobNewResponse struct {
	// Successfully created webhook jobs.
	Created []CasbPostureWebhookJobNewResponseCreated `json:"created" api:"required"`
	// Failed webhook job creation attempts.
	Failed []CasbPostureWebhookJobNewResponseFailed `json:"failed" api:"required"`
	JSON   casbPostureWebhookJobNewResponseJSON     `json:"-"`
}

// casbPostureWebhookJobNewResponseJSON contains the JSON metadata for the struct
// [CasbPostureWebhookJobNewResponse]
type casbPostureWebhookJobNewResponseJSON struct {
	Created     apijson.Field
	Failed      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookJobNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookJobNewResponseJSON) RawJSON() string {
	return r.raw
}

// Information about a webhook job.
type CasbPostureWebhookJobNewResponseCreated struct {
	// Unique identifier for the webhook job.
	ID string `json:"id" api:"required" format:"uuid"`
	// Asset data associated with this webhook job.
	AssetData map[string]interface{} `json:"asset_data" api:"required"`
	// When the webhook job was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// ID of the integration.
	IntegrationID string `json:"integration_id" api:"required" format:"uuid"`
	// When the webhook job was last updated.
	LastUpdatedAt time.Time `json:"last_updated_at" api:"required" format:"date-time"`
	// Parameters for a webhook job.
	Parameters CasbPostureWebhookJobNewResponseCreatedParameters `json:"parameters" api:"required"`
	// Status of a webhook job.
	Status CasbPostureWebhookJobNewResponseCreatedStatus `json:"status" api:"required"`
	// Type of actor that triggered the webhook job.
	TriggeredByActor CasbPostureWebhookJobNewResponseCreatedTriggeredByActor `json:"triggered_by_actor" api:"required"`
	// ID of the actor that triggered the job.
	TriggeredByID string `json:"triggered_by_id" api:"required"`
	// ID of the webhook configuration.
	WebhookID string `json:"webhook_id" api:"required" format:"uuid"`
	// Additional details about the failure.
	FailureDetails map[string]interface{} `json:"failure_details"`
	// Reason for webhook job failure.
	FailureReason CasbPostureWebhookJobNewResponseCreatedFailureReason `json:"failure_reason"`
	JSON          casbPostureWebhookJobNewResponseCreatedJSON          `json:"-"`
}

// casbPostureWebhookJobNewResponseCreatedJSON contains the JSON metadata for the
// struct [CasbPostureWebhookJobNewResponseCreated]
type casbPostureWebhookJobNewResponseCreatedJSON struct {
	ID               apijson.Field
	AssetData        apijson.Field
	CreatedAt        apijson.Field
	IntegrationID    apijson.Field
	LastUpdatedAt    apijson.Field
	Parameters       apijson.Field
	Status           apijson.Field
	TriggeredByActor apijson.Field
	TriggeredByID    apijson.Field
	WebhookID        apijson.Field
	FailureDetails   apijson.Field
	FailureReason    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookJobNewResponseCreated) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookJobNewResponseCreatedJSON) RawJSON() string {
	return r.raw
}

// Parameters for a webhook job.
type CasbPostureWebhookJobNewResponseCreatedParameters struct {
	// ID of the finding instance.
	FindingInstanceID string                                                `json:"finding_instance_id" api:"required" format:"uuid"`
	JSON              casbPostureWebhookJobNewResponseCreatedParametersJSON `json:"-"`
}

// casbPostureWebhookJobNewResponseCreatedParametersJSON contains the JSON metadata
// for the struct [CasbPostureWebhookJobNewResponseCreatedParameters]
type casbPostureWebhookJobNewResponseCreatedParametersJSON struct {
	FindingInstanceID apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CasbPostureWebhookJobNewResponseCreatedParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookJobNewResponseCreatedParametersJSON) RawJSON() string {
	return r.raw
}

// Status of a webhook job.
type CasbPostureWebhookJobNewResponseCreatedStatus string

const (
	CasbPostureWebhookJobNewResponseCreatedStatusPending    CasbPostureWebhookJobNewResponseCreatedStatus = "pending"
	CasbPostureWebhookJobNewResponseCreatedStatusProcessing CasbPostureWebhookJobNewResponseCreatedStatus = "processing"
	CasbPostureWebhookJobNewResponseCreatedStatusCompleted  CasbPostureWebhookJobNewResponseCreatedStatus = "completed"
	CasbPostureWebhookJobNewResponseCreatedStatusFailed     CasbPostureWebhookJobNewResponseCreatedStatus = "failed"
)

func (r CasbPostureWebhookJobNewResponseCreatedStatus) IsKnown() bool {
	switch r {
	case CasbPostureWebhookJobNewResponseCreatedStatusPending, CasbPostureWebhookJobNewResponseCreatedStatusProcessing, CasbPostureWebhookJobNewResponseCreatedStatusCompleted, CasbPostureWebhookJobNewResponseCreatedStatusFailed:
		return true
	}
	return false
}

// Type of actor that triggered the webhook job.
type CasbPostureWebhookJobNewResponseCreatedTriggeredByActor string

const (
	CasbPostureWebhookJobNewResponseCreatedTriggeredByActorUser         CasbPostureWebhookJobNewResponseCreatedTriggeredByActor = "user"
	CasbPostureWebhookJobNewResponseCreatedTriggeredByActorAccountToken CasbPostureWebhookJobNewResponseCreatedTriggeredByActor = "account_token"
)

func (r CasbPostureWebhookJobNewResponseCreatedTriggeredByActor) IsKnown() bool {
	switch r {
	case CasbPostureWebhookJobNewResponseCreatedTriggeredByActorUser, CasbPostureWebhookJobNewResponseCreatedTriggeredByActorAccountToken:
		return true
	}
	return false
}

// Reason for webhook job failure.
type CasbPostureWebhookJobNewResponseCreatedFailureReason string

const (
	CasbPostureWebhookJobNewResponseCreatedFailureReasonPermissionDenied              CasbPostureWebhookJobNewResponseCreatedFailureReason = "Permission Denied"
	CasbPostureWebhookJobNewResponseCreatedFailureReasonIntegrationUnavailable        CasbPostureWebhookJobNewResponseCreatedFailureReason = "Integration Unavailable"
	CasbPostureWebhookJobNewResponseCreatedFailureReasonServiceTemporarilyUnavailable CasbPostureWebhookJobNewResponseCreatedFailureReason = "Service Temporarily Unavailable"
	CasbPostureWebhookJobNewResponseCreatedFailureReasonSystemError                   CasbPostureWebhookJobNewResponseCreatedFailureReason = "System Error"
)

func (r CasbPostureWebhookJobNewResponseCreatedFailureReason) IsKnown() bool {
	switch r {
	case CasbPostureWebhookJobNewResponseCreatedFailureReasonPermissionDenied, CasbPostureWebhookJobNewResponseCreatedFailureReasonIntegrationUnavailable, CasbPostureWebhookJobNewResponseCreatedFailureReasonServiceTemporarilyUnavailable, CasbPostureWebhookJobNewResponseCreatedFailureReasonSystemError:
		return true
	}
	return false
}

// Information about a failed webhook job creation.
type CasbPostureWebhookJobNewResponseFailed struct {
	// Error message describing the failure.
	Error string `json:"error" api:"required"`
	// ID of the finding instance that failed to create a webhook job.
	FindingInstanceID string `json:"finding_instance_id" api:"required" format:"uuid"`
	// ID of the webhook configuration.
	WebhookID string                                     `json:"webhook_id" api:"required" format:"uuid"`
	JSON      casbPostureWebhookJobNewResponseFailedJSON `json:"-"`
}

// casbPostureWebhookJobNewResponseFailedJSON contains the JSON metadata for the
// struct [CasbPostureWebhookJobNewResponseFailed]
type casbPostureWebhookJobNewResponseFailedJSON struct {
	Error             apijson.Field
	FindingInstanceID apijson.Field
	WebhookID         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CasbPostureWebhookJobNewResponseFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookJobNewResponseFailedJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookJobNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Array of finding instance IDs to send to the webhooks
	FindingInstanceIDs param.Field[[]string] `json:"finding_instance_ids" api:"required" format:"uuid"`
	// Array of webhook IDs to trigger jobs for
	WebhookIDs param.Field[[]string] `json:"webhook_ids" api:"required" format:"uuid"`
}

func (r CasbPostureWebhookJobNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Common response structure for all API endpoints.
type CasbPostureWebhookJobNewResponseEnvelope struct {
	Errors   []CasbPostureWebhookJobNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureWebhookJobNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   CasbPostureWebhookJobNewResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success bool                                         `json:"success" api:"required"`
	JSON    casbPostureWebhookJobNewResponseEnvelopeJSON `json:"-"`
}

// casbPostureWebhookJobNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CasbPostureWebhookJobNewResponseEnvelope]
type casbPostureWebhookJobNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookJobNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookJobNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookJobNewResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                               `json:"documentation_url" format:"uri"`
	Source           CasbPostureWebhookJobNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureWebhookJobNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureWebhookJobNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [CasbPostureWebhookJobNewResponseEnvelopeErrors]
type casbPostureWebhookJobNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookJobNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookJobNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookJobNewResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                   `json:"pointer"`
	JSON    casbPostureWebhookJobNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureWebhookJobNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [CasbPostureWebhookJobNewResponseEnvelopeErrorsSource]
type casbPostureWebhookJobNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookJobNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookJobNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookJobNewResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                 `json:"documentation_url" format:"uri"`
	Source           CasbPostureWebhookJobNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureWebhookJobNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureWebhookJobNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [CasbPostureWebhookJobNewResponseEnvelopeMessages]
type casbPostureWebhookJobNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureWebhookJobNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookJobNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureWebhookJobNewResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                     `json:"pointer"`
	JSON    casbPostureWebhookJobNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureWebhookJobNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CasbPostureWebhookJobNewResponseEnvelopeMessagesSource]
type casbPostureWebhookJobNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureWebhookJobNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureWebhookJobNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

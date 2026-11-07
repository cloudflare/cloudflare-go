// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ObservabilityIssueAutomationService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservabilityIssueAutomationService] method instead.
type ObservabilityIssueAutomationService struct {
	Options []option.RequestOption
}

// NewObservabilityIssueAutomationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObservabilityIssueAutomationService(opts ...option.RequestOption) (r *ObservabilityIssueAutomationService) {
	r = &ObservabilityIssueAutomationService{}
	r.Options = opts
	return
}

// Create an account or service issue automation.
func (r *ObservabilityIssueAutomationService) New(ctx context.Context, params ObservabilityIssueAutomationNewParams, opts ...option.RequestOption) (res *ObservabilityIssueAutomationNewResponse, err error) {
	var env ObservabilityIssueAutomationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues/automations", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replace an issue automation configuration.
func (r *ObservabilityIssueAutomationService) Update(ctx context.Context, automationID string, params ObservabilityIssueAutomationUpdateParams, opts ...option.RequestOption) (res *ObservabilityIssueAutomationUpdateResponse, err error) {
	var env ObservabilityIssueAutomationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if automationID == "" {
		err = errors.New("missing required automationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues/automations/%s", params.AccountID, automationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List account and service issue automations.
func (r *ObservabilityIssueAutomationService) List(ctx context.Context, params ObservabilityIssueAutomationListParams, opts ...option.RequestOption) (res *ObservabilityIssueAutomationListResponse, err error) {
	var env ObservabilityIssueAutomationListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues/automations", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Delete an issue automation while preserving its historical runs.
func (r *ObservabilityIssueAutomationService) Delete(ctx context.Context, automationID string, body ObservabilityIssueAutomationDeleteParams, opts ...option.RequestOption) (res *ObservabilityIssueAutomationDeleteResponse, err error) {
	var env ObservabilityIssueAutomationDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if automationID == "" {
		err = errors.New("missing required automationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues/automations/%s", body.AccountID, automationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve an issue automation.
func (r *ObservabilityIssueAutomationService) Get(ctx context.Context, automationID string, query ObservabilityIssueAutomationGetParams, opts ...option.RequestOption) (res *ObservabilityIssueAutomationGetResponse, err error) {
	var env ObservabilityIssueAutomationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if automationID == "" {
		err = errors.New("missing required automationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues/automations/%s", query.AccountID, automationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ObservabilityIssueAutomationNewResponse struct {
	Automation ObservabilityIssueAutomationNewResponseAutomation `json:"automation" api:"required"`
	JSON       observabilityIssueAutomationNewResponseJSON       `json:"-"`
}

// observabilityIssueAutomationNewResponseJSON contains the JSON metadata for the
// struct [ObservabilityIssueAutomationNewResponse]
type observabilityIssueAutomationNewResponseJSON struct {
	Automation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationNewResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationNewResponseAutomation struct {
	ID                string                                                       `json:"id" api:"required"`
	AnsPolicyID       string                                                       `json:"ansPolicyId" api:"required"`
	Created           float64                                                      `json:"created" api:"required"`
	CreatedByUserID   string                                                       `json:"createdByUserId" api:"required,nullable"`
	Enabled           bool                                                         `json:"enabled" api:"required"`
	InactivitySeconds float64                                                      `json:"inactivitySeconds" api:"required,nullable"`
	Name              string                                                       `json:"name" api:"required"`
	Revision          float64                                                      `json:"revision" api:"required"`
	Scope             ObservabilityIssueAutomationNewResponseAutomationScope       `json:"scope" api:"required"`
	Service           string                                                       `json:"service" api:"required,nullable"`
	ServiceType       string                                                       `json:"serviceType" api:"required,nullable"`
	Threshold         float64                                                      `json:"threshold" api:"required,nullable"`
	TriggerType       ObservabilityIssueAutomationNewResponseAutomationTriggerType `json:"triggerType" api:"required"`
	Updated           float64                                                      `json:"updated" api:"required"`
	UpdatedByUserID   string                                                       `json:"updatedByUserId" api:"required,nullable"`
	JSON              observabilityIssueAutomationNewResponseAutomationJSON        `json:"-"`
}

// observabilityIssueAutomationNewResponseAutomationJSON contains the JSON metadata
// for the struct [ObservabilityIssueAutomationNewResponseAutomation]
type observabilityIssueAutomationNewResponseAutomationJSON struct {
	ID                apijson.Field
	AnsPolicyID       apijson.Field
	Created           apijson.Field
	CreatedByUserID   apijson.Field
	Enabled           apijson.Field
	InactivitySeconds apijson.Field
	Name              apijson.Field
	Revision          apijson.Field
	Scope             apijson.Field
	Service           apijson.Field
	ServiceType       apijson.Field
	Threshold         apijson.Field
	TriggerType       apijson.Field
	Updated           apijson.Field
	UpdatedByUserID   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationNewResponseAutomation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationNewResponseAutomationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationNewResponseAutomationScope string

const (
	ObservabilityIssueAutomationNewResponseAutomationScopeAccount ObservabilityIssueAutomationNewResponseAutomationScope = "account"
	ObservabilityIssueAutomationNewResponseAutomationScopeService ObservabilityIssueAutomationNewResponseAutomationScope = "service"
)

func (r ObservabilityIssueAutomationNewResponseAutomationScope) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationNewResponseAutomationScopeAccount, ObservabilityIssueAutomationNewResponseAutomationScopeService:
		return true
	}
	return false
}

type ObservabilityIssueAutomationNewResponseAutomationTriggerType string

const (
	ObservabilityIssueAutomationNewResponseAutomationTriggerTypeOccurrenceThreshold       ObservabilityIssueAutomationNewResponseAutomationTriggerType = "occurrence_threshold"
	ObservabilityIssueAutomationNewResponseAutomationTriggerTypeRecurrenceAfterInactivity ObservabilityIssueAutomationNewResponseAutomationTriggerType = "recurrence_after_inactivity"
)

func (r ObservabilityIssueAutomationNewResponseAutomationTriggerType) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationNewResponseAutomationTriggerTypeOccurrenceThreshold, ObservabilityIssueAutomationNewResponseAutomationTriggerTypeRecurrenceAfterInactivity:
		return true
	}
	return false
}

type ObservabilityIssueAutomationUpdateResponse struct {
	Automation ObservabilityIssueAutomationUpdateResponseAutomation `json:"automation" api:"required"`
	JSON       observabilityIssueAutomationUpdateResponseJSON       `json:"-"`
}

// observabilityIssueAutomationUpdateResponseJSON contains the JSON metadata for
// the struct [ObservabilityIssueAutomationUpdateResponse]
type observabilityIssueAutomationUpdateResponseJSON struct {
	Automation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationUpdateResponseAutomation struct {
	ID                string                                                          `json:"id" api:"required"`
	AnsPolicyID       string                                                          `json:"ansPolicyId" api:"required"`
	Created           float64                                                         `json:"created" api:"required"`
	CreatedByUserID   string                                                          `json:"createdByUserId" api:"required,nullable"`
	Enabled           bool                                                            `json:"enabled" api:"required"`
	InactivitySeconds float64                                                         `json:"inactivitySeconds" api:"required,nullable"`
	Name              string                                                          `json:"name" api:"required"`
	Revision          float64                                                         `json:"revision" api:"required"`
	Scope             ObservabilityIssueAutomationUpdateResponseAutomationScope       `json:"scope" api:"required"`
	Service           string                                                          `json:"service" api:"required,nullable"`
	ServiceType       string                                                          `json:"serviceType" api:"required,nullable"`
	Threshold         float64                                                         `json:"threshold" api:"required,nullable"`
	TriggerType       ObservabilityIssueAutomationUpdateResponseAutomationTriggerType `json:"triggerType" api:"required"`
	Updated           float64                                                         `json:"updated" api:"required"`
	UpdatedByUserID   string                                                          `json:"updatedByUserId" api:"required,nullable"`
	JSON              observabilityIssueAutomationUpdateResponseAutomationJSON        `json:"-"`
}

// observabilityIssueAutomationUpdateResponseAutomationJSON contains the JSON
// metadata for the struct [ObservabilityIssueAutomationUpdateResponseAutomation]
type observabilityIssueAutomationUpdateResponseAutomationJSON struct {
	ID                apijson.Field
	AnsPolicyID       apijson.Field
	Created           apijson.Field
	CreatedByUserID   apijson.Field
	Enabled           apijson.Field
	InactivitySeconds apijson.Field
	Name              apijson.Field
	Revision          apijson.Field
	Scope             apijson.Field
	Service           apijson.Field
	ServiceType       apijson.Field
	Threshold         apijson.Field
	TriggerType       apijson.Field
	Updated           apijson.Field
	UpdatedByUserID   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationUpdateResponseAutomation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationUpdateResponseAutomationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationUpdateResponseAutomationScope string

const (
	ObservabilityIssueAutomationUpdateResponseAutomationScopeAccount ObservabilityIssueAutomationUpdateResponseAutomationScope = "account"
	ObservabilityIssueAutomationUpdateResponseAutomationScopeService ObservabilityIssueAutomationUpdateResponseAutomationScope = "service"
)

func (r ObservabilityIssueAutomationUpdateResponseAutomationScope) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationUpdateResponseAutomationScopeAccount, ObservabilityIssueAutomationUpdateResponseAutomationScopeService:
		return true
	}
	return false
}

type ObservabilityIssueAutomationUpdateResponseAutomationTriggerType string

const (
	ObservabilityIssueAutomationUpdateResponseAutomationTriggerTypeOccurrenceThreshold       ObservabilityIssueAutomationUpdateResponseAutomationTriggerType = "occurrence_threshold"
	ObservabilityIssueAutomationUpdateResponseAutomationTriggerTypeRecurrenceAfterInactivity ObservabilityIssueAutomationUpdateResponseAutomationTriggerType = "recurrence_after_inactivity"
)

func (r ObservabilityIssueAutomationUpdateResponseAutomationTriggerType) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationUpdateResponseAutomationTriggerTypeOccurrenceThreshold, ObservabilityIssueAutomationUpdateResponseAutomationTriggerTypeRecurrenceAfterInactivity:
		return true
	}
	return false
}

type ObservabilityIssueAutomationListResponse struct {
	Automations []ObservabilityIssueAutomationListResponseAutomation `json:"automations" api:"required"`
	JSON        observabilityIssueAutomationListResponseJSON         `json:"-"`
}

// observabilityIssueAutomationListResponseJSON contains the JSON metadata for the
// struct [ObservabilityIssueAutomationListResponse]
type observabilityIssueAutomationListResponseJSON struct {
	Automations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationListResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationListResponseAutomation struct {
	ID                string                                                         `json:"id" api:"required"`
	AnsPolicyID       string                                                         `json:"ansPolicyId" api:"required"`
	Created           float64                                                        `json:"created" api:"required"`
	CreatedByUserID   string                                                         `json:"createdByUserId" api:"required,nullable"`
	Enabled           bool                                                           `json:"enabled" api:"required"`
	InactivitySeconds float64                                                        `json:"inactivitySeconds" api:"required,nullable"`
	Name              string                                                         `json:"name" api:"required"`
	Revision          float64                                                        `json:"revision" api:"required"`
	Scope             ObservabilityIssueAutomationListResponseAutomationsScope       `json:"scope" api:"required"`
	Service           string                                                         `json:"service" api:"required,nullable"`
	ServiceType       string                                                         `json:"serviceType" api:"required,nullable"`
	Threshold         float64                                                        `json:"threshold" api:"required,nullable"`
	TriggerType       ObservabilityIssueAutomationListResponseAutomationsTriggerType `json:"triggerType" api:"required"`
	Updated           float64                                                        `json:"updated" api:"required"`
	UpdatedByUserID   string                                                         `json:"updatedByUserId" api:"required,nullable"`
	JSON              observabilityIssueAutomationListResponseAutomationJSON         `json:"-"`
}

// observabilityIssueAutomationListResponseAutomationJSON contains the JSON
// metadata for the struct [ObservabilityIssueAutomationListResponseAutomation]
type observabilityIssueAutomationListResponseAutomationJSON struct {
	ID                apijson.Field
	AnsPolicyID       apijson.Field
	Created           apijson.Field
	CreatedByUserID   apijson.Field
	Enabled           apijson.Field
	InactivitySeconds apijson.Field
	Name              apijson.Field
	Revision          apijson.Field
	Scope             apijson.Field
	Service           apijson.Field
	ServiceType       apijson.Field
	Threshold         apijson.Field
	TriggerType       apijson.Field
	Updated           apijson.Field
	UpdatedByUserID   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationListResponseAutomation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationListResponseAutomationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationListResponseAutomationsScope string

const (
	ObservabilityIssueAutomationListResponseAutomationsScopeAccount ObservabilityIssueAutomationListResponseAutomationsScope = "account"
	ObservabilityIssueAutomationListResponseAutomationsScopeService ObservabilityIssueAutomationListResponseAutomationsScope = "service"
)

func (r ObservabilityIssueAutomationListResponseAutomationsScope) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationListResponseAutomationsScopeAccount, ObservabilityIssueAutomationListResponseAutomationsScopeService:
		return true
	}
	return false
}

type ObservabilityIssueAutomationListResponseAutomationsTriggerType string

const (
	ObservabilityIssueAutomationListResponseAutomationsTriggerTypeOccurrenceThreshold       ObservabilityIssueAutomationListResponseAutomationsTriggerType = "occurrence_threshold"
	ObservabilityIssueAutomationListResponseAutomationsTriggerTypeRecurrenceAfterInactivity ObservabilityIssueAutomationListResponseAutomationsTriggerType = "recurrence_after_inactivity"
)

func (r ObservabilityIssueAutomationListResponseAutomationsTriggerType) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationListResponseAutomationsTriggerTypeOccurrenceThreshold, ObservabilityIssueAutomationListResponseAutomationsTriggerTypeRecurrenceAfterInactivity:
		return true
	}
	return false
}

type ObservabilityIssueAutomationDeleteResponse struct {
	Automation ObservabilityIssueAutomationDeleteResponseAutomation `json:"automation" api:"required"`
	JSON       observabilityIssueAutomationDeleteResponseJSON       `json:"-"`
}

// observabilityIssueAutomationDeleteResponseJSON contains the JSON metadata for
// the struct [ObservabilityIssueAutomationDeleteResponse]
type observabilityIssueAutomationDeleteResponseJSON struct {
	Automation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationDeleteResponseAutomation struct {
	ID                string                                                          `json:"id" api:"required"`
	AnsPolicyID       string                                                          `json:"ansPolicyId" api:"required"`
	Created           float64                                                         `json:"created" api:"required"`
	CreatedByUserID   string                                                          `json:"createdByUserId" api:"required,nullable"`
	Enabled           bool                                                            `json:"enabled" api:"required"`
	InactivitySeconds float64                                                         `json:"inactivitySeconds" api:"required,nullable"`
	Name              string                                                          `json:"name" api:"required"`
	Revision          float64                                                         `json:"revision" api:"required"`
	Scope             ObservabilityIssueAutomationDeleteResponseAutomationScope       `json:"scope" api:"required"`
	Service           string                                                          `json:"service" api:"required,nullable"`
	ServiceType       string                                                          `json:"serviceType" api:"required,nullable"`
	Threshold         float64                                                         `json:"threshold" api:"required,nullable"`
	TriggerType       ObservabilityIssueAutomationDeleteResponseAutomationTriggerType `json:"triggerType" api:"required"`
	Updated           float64                                                         `json:"updated" api:"required"`
	UpdatedByUserID   string                                                          `json:"updatedByUserId" api:"required,nullable"`
	JSON              observabilityIssueAutomationDeleteResponseAutomationJSON        `json:"-"`
}

// observabilityIssueAutomationDeleteResponseAutomationJSON contains the JSON
// metadata for the struct [ObservabilityIssueAutomationDeleteResponseAutomation]
type observabilityIssueAutomationDeleteResponseAutomationJSON struct {
	ID                apijson.Field
	AnsPolicyID       apijson.Field
	Created           apijson.Field
	CreatedByUserID   apijson.Field
	Enabled           apijson.Field
	InactivitySeconds apijson.Field
	Name              apijson.Field
	Revision          apijson.Field
	Scope             apijson.Field
	Service           apijson.Field
	ServiceType       apijson.Field
	Threshold         apijson.Field
	TriggerType       apijson.Field
	Updated           apijson.Field
	UpdatedByUserID   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationDeleteResponseAutomation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationDeleteResponseAutomationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationDeleteResponseAutomationScope string

const (
	ObservabilityIssueAutomationDeleteResponseAutomationScopeAccount ObservabilityIssueAutomationDeleteResponseAutomationScope = "account"
	ObservabilityIssueAutomationDeleteResponseAutomationScopeService ObservabilityIssueAutomationDeleteResponseAutomationScope = "service"
)

func (r ObservabilityIssueAutomationDeleteResponseAutomationScope) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationDeleteResponseAutomationScopeAccount, ObservabilityIssueAutomationDeleteResponseAutomationScopeService:
		return true
	}
	return false
}

type ObservabilityIssueAutomationDeleteResponseAutomationTriggerType string

const (
	ObservabilityIssueAutomationDeleteResponseAutomationTriggerTypeOccurrenceThreshold       ObservabilityIssueAutomationDeleteResponseAutomationTriggerType = "occurrence_threshold"
	ObservabilityIssueAutomationDeleteResponseAutomationTriggerTypeRecurrenceAfterInactivity ObservabilityIssueAutomationDeleteResponseAutomationTriggerType = "recurrence_after_inactivity"
)

func (r ObservabilityIssueAutomationDeleteResponseAutomationTriggerType) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationDeleteResponseAutomationTriggerTypeOccurrenceThreshold, ObservabilityIssueAutomationDeleteResponseAutomationTriggerTypeRecurrenceAfterInactivity:
		return true
	}
	return false
}

type ObservabilityIssueAutomationGetResponse struct {
	Automation ObservabilityIssueAutomationGetResponseAutomation `json:"automation" api:"required"`
	JSON       observabilityIssueAutomationGetResponseJSON       `json:"-"`
}

// observabilityIssueAutomationGetResponseJSON contains the JSON metadata for the
// struct [ObservabilityIssueAutomationGetResponse]
type observabilityIssueAutomationGetResponseJSON struct {
	Automation  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationGetResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationGetResponseAutomation struct {
	ID                string                                                       `json:"id" api:"required"`
	AnsPolicyID       string                                                       `json:"ansPolicyId" api:"required"`
	Created           float64                                                      `json:"created" api:"required"`
	CreatedByUserID   string                                                       `json:"createdByUserId" api:"required,nullable"`
	Enabled           bool                                                         `json:"enabled" api:"required"`
	InactivitySeconds float64                                                      `json:"inactivitySeconds" api:"required,nullable"`
	Name              string                                                       `json:"name" api:"required"`
	Revision          float64                                                      `json:"revision" api:"required"`
	Scope             ObservabilityIssueAutomationGetResponseAutomationScope       `json:"scope" api:"required"`
	Service           string                                                       `json:"service" api:"required,nullable"`
	ServiceType       string                                                       `json:"serviceType" api:"required,nullable"`
	Threshold         float64                                                      `json:"threshold" api:"required,nullable"`
	TriggerType       ObservabilityIssueAutomationGetResponseAutomationTriggerType `json:"triggerType" api:"required"`
	Updated           float64                                                      `json:"updated" api:"required"`
	UpdatedByUserID   string                                                       `json:"updatedByUserId" api:"required,nullable"`
	JSON              observabilityIssueAutomationGetResponseAutomationJSON        `json:"-"`
}

// observabilityIssueAutomationGetResponseAutomationJSON contains the JSON metadata
// for the struct [ObservabilityIssueAutomationGetResponseAutomation]
type observabilityIssueAutomationGetResponseAutomationJSON struct {
	ID                apijson.Field
	AnsPolicyID       apijson.Field
	Created           apijson.Field
	CreatedByUserID   apijson.Field
	Enabled           apijson.Field
	InactivitySeconds apijson.Field
	Name              apijson.Field
	Revision          apijson.Field
	Scope             apijson.Field
	Service           apijson.Field
	ServiceType       apijson.Field
	Threshold         apijson.Field
	TriggerType       apijson.Field
	Updated           apijson.Field
	UpdatedByUserID   apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationGetResponseAutomation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationGetResponseAutomationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationGetResponseAutomationScope string

const (
	ObservabilityIssueAutomationGetResponseAutomationScopeAccount ObservabilityIssueAutomationGetResponseAutomationScope = "account"
	ObservabilityIssueAutomationGetResponseAutomationScopeService ObservabilityIssueAutomationGetResponseAutomationScope = "service"
)

func (r ObservabilityIssueAutomationGetResponseAutomationScope) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationGetResponseAutomationScopeAccount, ObservabilityIssueAutomationGetResponseAutomationScopeService:
		return true
	}
	return false
}

type ObservabilityIssueAutomationGetResponseAutomationTriggerType string

const (
	ObservabilityIssueAutomationGetResponseAutomationTriggerTypeOccurrenceThreshold       ObservabilityIssueAutomationGetResponseAutomationTriggerType = "occurrence_threshold"
	ObservabilityIssueAutomationGetResponseAutomationTriggerTypeRecurrenceAfterInactivity ObservabilityIssueAutomationGetResponseAutomationTriggerType = "recurrence_after_inactivity"
)

func (r ObservabilityIssueAutomationGetResponseAutomationTriggerType) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationGetResponseAutomationTriggerTypeOccurrenceThreshold, ObservabilityIssueAutomationGetResponseAutomationTriggerTypeRecurrenceAfterInactivity:
		return true
	}
	return false
}

type ObservabilityIssueAutomationNewParams struct {
	AccountID              param.Field[string] `path:"account_id" api:"required"`
	PolicyID               param.Field[string] `json:"policyId" api:"required"`
	AfterInactivitySeconds param.Field[int64]  `json:"afterInactivitySeconds"`
	AfterOccurrences       param.Field[int64]  `json:"afterOccurrences"`
	Enabled                param.Field[bool]   `json:"enabled"`
	Name                   param.Field[string] `json:"name"`
	Service                param.Field[string] `json:"service"`
}

func (r ObservabilityIssueAutomationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityIssueAutomationNewResponseEnvelope struct {
	Errors   []ObservabilityIssueAutomationNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityIssueAutomationNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityIssueAutomationNewResponse                   `json:"result" api:"required"`
	Success  ObservabilityIssueAutomationNewResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityIssueAutomationNewResponseEnvelopeJSON       `json:"-"`
}

// observabilityIssueAutomationNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [ObservabilityIssueAutomationNewResponseEnvelope]
type observabilityIssueAutomationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationNewResponseEnvelopeErrors struct {
	Message string                                                    `json:"message" api:"required"`
	JSON    observabilityIssueAutomationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityIssueAutomationNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ObservabilityIssueAutomationNewResponseEnvelopeErrors]
type observabilityIssueAutomationNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationNewResponseEnvelopeMessages struct {
	Message ObservabilityIssueAutomationNewResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityIssueAutomationNewResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityIssueAutomationNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ObservabilityIssueAutomationNewResponseEnvelopeMessages]
type observabilityIssueAutomationNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationNewResponseEnvelopeMessagesMessage string

const (
	ObservabilityIssueAutomationNewResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityIssueAutomationNewResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityIssueAutomationNewResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationNewResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityIssueAutomationNewResponseEnvelopeSuccess bool

const (
	ObservabilityIssueAutomationNewResponseEnvelopeSuccessTrue ObservabilityIssueAutomationNewResponseEnvelopeSuccess = true
)

func (r ObservabilityIssueAutomationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityIssueAutomationUpdateParams struct {
	AccountID              param.Field[string] `path:"account_id" api:"required"`
	PolicyID               param.Field[string] `json:"policyId" api:"required"`
	AfterInactivitySeconds param.Field[int64]  `json:"afterInactivitySeconds"`
	AfterOccurrences       param.Field[int64]  `json:"afterOccurrences"`
	Enabled                param.Field[bool]   `json:"enabled"`
	Name                   param.Field[string] `json:"name"`
	Service                param.Field[string] `json:"service"`
}

func (r ObservabilityIssueAutomationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityIssueAutomationUpdateResponseEnvelope struct {
	Errors   []ObservabilityIssueAutomationUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityIssueAutomationUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityIssueAutomationUpdateResponse                   `json:"result" api:"required"`
	Success  ObservabilityIssueAutomationUpdateResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityIssueAutomationUpdateResponseEnvelopeJSON       `json:"-"`
}

// observabilityIssueAutomationUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [ObservabilityIssueAutomationUpdateResponseEnvelope]
type observabilityIssueAutomationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationUpdateResponseEnvelopeErrors struct {
	Message string                                                       `json:"message" api:"required"`
	JSON    observabilityIssueAutomationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityIssueAutomationUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ObservabilityIssueAutomationUpdateResponseEnvelopeErrors]
type observabilityIssueAutomationUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationUpdateResponseEnvelopeMessages struct {
	Message ObservabilityIssueAutomationUpdateResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityIssueAutomationUpdateResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityIssueAutomationUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ObservabilityIssueAutomationUpdateResponseEnvelopeMessages]
type observabilityIssueAutomationUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationUpdateResponseEnvelopeMessagesMessage string

const (
	ObservabilityIssueAutomationUpdateResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityIssueAutomationUpdateResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityIssueAutomationUpdateResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationUpdateResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityIssueAutomationUpdateResponseEnvelopeSuccess bool

const (
	ObservabilityIssueAutomationUpdateResponseEnvelopeSuccessTrue ObservabilityIssueAutomationUpdateResponseEnvelopeSuccess = true
)

func (r ObservabilityIssueAutomationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityIssueAutomationListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Service   param.Field[string] `query:"service"`
}

// URLQuery serializes [ObservabilityIssueAutomationListParams]'s query parameters
// as `url.Values`.
func (r ObservabilityIssueAutomationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ObservabilityIssueAutomationListResponseEnvelope struct {
	Errors   []ObservabilityIssueAutomationListResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityIssueAutomationListResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityIssueAutomationListResponse                   `json:"result" api:"required"`
	Success  ObservabilityIssueAutomationListResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityIssueAutomationListResponseEnvelopeJSON       `json:"-"`
}

// observabilityIssueAutomationListResponseEnvelopeJSON contains the JSON metadata
// for the struct [ObservabilityIssueAutomationListResponseEnvelope]
type observabilityIssueAutomationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationListResponseEnvelopeErrors struct {
	Message string                                                     `json:"message" api:"required"`
	JSON    observabilityIssueAutomationListResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityIssueAutomationListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ObservabilityIssueAutomationListResponseEnvelopeErrors]
type observabilityIssueAutomationListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationListResponseEnvelopeMessages struct {
	Message ObservabilityIssueAutomationListResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityIssueAutomationListResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityIssueAutomationListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ObservabilityIssueAutomationListResponseEnvelopeMessages]
type observabilityIssueAutomationListResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationListResponseEnvelopeMessagesMessage string

const (
	ObservabilityIssueAutomationListResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityIssueAutomationListResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityIssueAutomationListResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationListResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityIssueAutomationListResponseEnvelopeSuccess bool

const (
	ObservabilityIssueAutomationListResponseEnvelopeSuccessTrue ObservabilityIssueAutomationListResponseEnvelopeSuccess = true
)

func (r ObservabilityIssueAutomationListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityIssueAutomationDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ObservabilityIssueAutomationDeleteResponseEnvelope struct {
	Errors   []ObservabilityIssueAutomationDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityIssueAutomationDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityIssueAutomationDeleteResponse                   `json:"result" api:"required"`
	Success  ObservabilityIssueAutomationDeleteResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityIssueAutomationDeleteResponseEnvelopeJSON       `json:"-"`
}

// observabilityIssueAutomationDeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [ObservabilityIssueAutomationDeleteResponseEnvelope]
type observabilityIssueAutomationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationDeleteResponseEnvelopeErrors struct {
	Message string                                                       `json:"message" api:"required"`
	JSON    observabilityIssueAutomationDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityIssueAutomationDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ObservabilityIssueAutomationDeleteResponseEnvelopeErrors]
type observabilityIssueAutomationDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationDeleteResponseEnvelopeMessages struct {
	Message ObservabilityIssueAutomationDeleteResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityIssueAutomationDeleteResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityIssueAutomationDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ObservabilityIssueAutomationDeleteResponseEnvelopeMessages]
type observabilityIssueAutomationDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationDeleteResponseEnvelopeMessagesMessage string

const (
	ObservabilityIssueAutomationDeleteResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityIssueAutomationDeleteResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityIssueAutomationDeleteResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationDeleteResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityIssueAutomationDeleteResponseEnvelopeSuccess bool

const (
	ObservabilityIssueAutomationDeleteResponseEnvelopeSuccessTrue ObservabilityIssueAutomationDeleteResponseEnvelopeSuccess = true
)

func (r ObservabilityIssueAutomationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityIssueAutomationGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ObservabilityIssueAutomationGetResponseEnvelope struct {
	Errors   []ObservabilityIssueAutomationGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityIssueAutomationGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityIssueAutomationGetResponse                   `json:"result" api:"required"`
	Success  ObservabilityIssueAutomationGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityIssueAutomationGetResponseEnvelopeJSON       `json:"-"`
}

// observabilityIssueAutomationGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ObservabilityIssueAutomationGetResponseEnvelope]
type observabilityIssueAutomationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationGetResponseEnvelopeErrors struct {
	Message string                                                    `json:"message" api:"required"`
	JSON    observabilityIssueAutomationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityIssueAutomationGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ObservabilityIssueAutomationGetResponseEnvelopeErrors]
type observabilityIssueAutomationGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationGetResponseEnvelopeMessages struct {
	Message ObservabilityIssueAutomationGetResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityIssueAutomationGetResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityIssueAutomationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ObservabilityIssueAutomationGetResponseEnvelopeMessages]
type observabilityIssueAutomationGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAutomationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAutomationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAutomationGetResponseEnvelopeMessagesMessage string

const (
	ObservabilityIssueAutomationGetResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityIssueAutomationGetResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityIssueAutomationGetResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationGetResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityIssueAutomationGetResponseEnvelopeSuccess bool

const (
	ObservabilityIssueAutomationGetResponseEnvelopeSuccessTrue ObservabilityIssueAutomationGetResponseEnvelopeSuccess = true
)

func (r ObservabilityIssueAutomationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityIssueAutomationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

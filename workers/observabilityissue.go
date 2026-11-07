// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/tidwall/gjson"
)

// ObservabilityIssueService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservabilityIssueService] method instead.
type ObservabilityIssueService struct {
	Options     []option.RequestOption
	Automations *ObservabilityIssueAutomationService
}

// NewObservabilityIssueService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObservabilityIssueService(opts ...option.RequestOption) (r *ObservabilityIssueService) {
	r = &ObservabilityIssueService{}
	r.Options = opts
	r.Automations = NewObservabilityIssueAutomationService(opts...)
	return
}

// Modify an issue status or group title.
func (r *ObservabilityIssueService) Update(ctx context.Context, issueID string, params ObservabilityIssueUpdateParams, opts ...option.RequestOption) (res *ObservabilityIssueUpdateResponse, err error) {
	var env ObservabilityIssueUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if issueID == "" {
		err = errors.New("missing required issueId parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues/%s", params.AccountID, issueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List detected issues.
func (r *ObservabilityIssueService) List(ctx context.Context, params ObservabilityIssueListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ObservabilityIssueListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List detected issues.
func (r *ObservabilityIssueService) ListAutoPaging(ctx context.Context, params ObservabilityIssueListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ObservabilityIssueListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Add one or more ungrouped issues to an existing issue group.
func (r *ObservabilityIssueService) Add(ctx context.Context, groupID string, params ObservabilityIssueAddParams, opts ...option.RequestOption) (res *ObservabilityIssueAddResponse, err error) {
	var env ObservabilityIssueAddResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues/groups/%s/members", params.AccountID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve an issue.
func (r *ObservabilityIssueService) Get(ctx context.Context, issueID string, query ObservabilityIssueGetParams, opts ...option.RequestOption) (res *ObservabilityIssueGetResponse, err error) {
	var env ObservabilityIssueGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if issueID == "" {
		err = errors.New("missing required issueId parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues/%s", query.AccountID, issueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Group two or more issues under a new customer-defined identity.
func (r *ObservabilityIssueService) Group(ctx context.Context, params ObservabilityIssueGroupParams, opts ...option.RequestOption) (res *ObservabilityIssueGroupResponse, err error) {
	var env ObservabilityIssueGroupResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues/groups", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Send the issue to one ANS policy, or every matching Real-Time Issue policy when
// no policy ID is provided.
func (r *ObservabilityIssueService) Notify(ctx context.Context, issueID string, params ObservabilityIssueNotifyParams, opts ...option.RequestOption) (res *ObservabilityIssueNotifyResponse, err error) {
	var env ObservabilityIssueNotifyResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if issueID == "" {
		err = errors.New("missing required issueId parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues/%s/notification-runs", params.AccountID, issueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List the retained error occurrences grouped into an issue, newest first.
func (r *ObservabilityIssueService) Occurrences(ctx context.Context, issueID string, params ObservabilityIssueOccurrencesParams, opts ...option.RequestOption) (res *[]ObservabilityIssueOccurrencesResponse, err error) {
	var env ObservabilityIssueOccurrencesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if issueID == "" {
		err = errors.New("missing required issueId parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues/%s/occurrences", params.AccountID, issueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Remove an issue from a group, dissolving the group when fewer than two members
// remain.
func (r *ObservabilityIssueService) Remove(ctx context.Context, groupID string, issueID string, body ObservabilityIssueRemoveParams, opts ...option.RequestOption) (res *ObservabilityIssueRemoveResponse, err error) {
	var env ObservabilityIssueRemoveResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	if issueID == "" {
		err = errors.New("missing required issueId parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues/groups/%s/members/%s", body.AccountID, groupID, issueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve aggregate issue counts for the account.
func (r *ObservabilityIssueService) Summary(ctx context.Context, params ObservabilityIssueSummaryParams, opts ...option.RequestOption) (res *ObservabilityIssueSummaryResponse, err error) {
	var env ObservabilityIssueSummaryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues/summary", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Dissolve an issue group and restore its members as standalone issues.
func (r *ObservabilityIssueService) Ungroup(ctx context.Context, groupID string, body ObservabilityIssueUngroupParams, opts ...option.RequestOption) (res *ObservabilityIssueUngroupResponse, err error) {
	var env ObservabilityIssueUngroupResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if groupID == "" {
		err = errors.New("missing required groupId parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/issues/groups/%s", body.AccountID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ObservabilityIssueUpdateResponse struct {
	ID            string                                   `json:"id" api:"required"`
	Count         float64                                  `json:"count" api:"required"`
	Created       float64                                  `json:"created" api:"required"`
	FirstObserved float64                                  `json:"firstObserved" api:"required"`
	LastObserved  float64                                  `json:"lastObserved" api:"required"`
	Service       string                                   `json:"service" api:"required"`
	Status        ObservabilityIssueUpdateResponseStatus   `json:"status" api:"required"`
	StatusUpdated float64                                  `json:"statusUpdated" api:"required"`
	Title         string                                   `json:"title" api:"required"`
	Updated       float64                                  `json:"updated" api:"required"`
	Fingerprint   string                                   `json:"fingerprint"`
	Members       []ObservabilityIssueUpdateResponseMember `json:"members"`
	Type          string                                   `json:"type"`
	JSON          observabilityIssueUpdateResponseJSON     `json:"-"`
}

// observabilityIssueUpdateResponseJSON contains the JSON metadata for the struct
// [ObservabilityIssueUpdateResponse]
type observabilityIssueUpdateResponseJSON struct {
	ID            apijson.Field
	Count         apijson.Field
	Created       apijson.Field
	FirstObserved apijson.Field
	LastObserved  apijson.Field
	Service       apijson.Field
	Status        apijson.Field
	StatusUpdated apijson.Field
	Title         apijson.Field
	Updated       apijson.Field
	Fingerprint   apijson.Field
	Members       apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityIssueUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueUpdateResponseStatus string

const (
	ObservabilityIssueUpdateResponseStatusActive   ObservabilityIssueUpdateResponseStatus = "active"
	ObservabilityIssueUpdateResponseStatusResolved ObservabilityIssueUpdateResponseStatus = "resolved"
	ObservabilityIssueUpdateResponseStatusIgnored  ObservabilityIssueUpdateResponseStatus = "ignored"
)

func (r ObservabilityIssueUpdateResponseStatus) IsKnown() bool {
	switch r {
	case ObservabilityIssueUpdateResponseStatusActive, ObservabilityIssueUpdateResponseStatusResolved, ObservabilityIssueUpdateResponseStatusIgnored:
		return true
	}
	return false
}

type ObservabilityIssueUpdateResponseMember struct {
	Fingerprint string                                     `json:"fingerprint" api:"required"`
	IssueID     string                                     `json:"issueId" api:"required"`
	Title       string                                     `json:"title" api:"required"`
	JSON        observabilityIssueUpdateResponseMemberJSON `json:"-"`
}

// observabilityIssueUpdateResponseMemberJSON contains the JSON metadata for the
// struct [ObservabilityIssueUpdateResponseMember]
type observabilityIssueUpdateResponseMemberJSON struct {
	Fingerprint apijson.Field
	IssueID     apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueUpdateResponseMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueUpdateResponseMemberJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueListResponse struct {
	ID            string                                   `json:"id" api:"required"`
	Activity      []ObservabilityIssueListResponseActivity `json:"activity" api:"required"`
	Count         float64                                  `json:"count" api:"required"`
	Created       float64                                  `json:"created" api:"required"`
	FirstObserved float64                                  `json:"firstObserved" api:"required"`
	LastObserved  float64                                  `json:"lastObserved" api:"required"`
	Service       string                                   `json:"service" api:"required"`
	Status        ObservabilityIssueListResponseStatus     `json:"status" api:"required"`
	StatusUpdated float64                                  `json:"statusUpdated" api:"required"`
	Title         string                                   `json:"title" api:"required"`
	Updated       float64                                  `json:"updated" api:"required"`
	Fingerprint   string                                   `json:"fingerprint"`
	Members       []ObservabilityIssueListResponseMember   `json:"members"`
	Type          string                                   `json:"type"`
	JSON          observabilityIssueListResponseJSON       `json:"-"`
}

// observabilityIssueListResponseJSON contains the JSON metadata for the struct
// [ObservabilityIssueListResponse]
type observabilityIssueListResponseJSON struct {
	ID            apijson.Field
	Activity      apijson.Field
	Count         apijson.Field
	Created       apijson.Field
	FirstObserved apijson.Field
	LastObserved  apijson.Field
	Service       apijson.Field
	Status        apijson.Field
	StatusUpdated apijson.Field
	Title         apijson.Field
	Updated       apijson.Field
	Fingerprint   apijson.Field
	Members       apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityIssueListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueListResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueListResponseActivity struct {
	ID                 string                                            `json:"id" api:"required"`
	Created            float64                                           `json:"created" api:"required"`
	InactivitySeconds  float64                                           `json:"inactivitySeconds" api:"required,nullable"`
	InitiatedBy        ObservabilityIssueListResponseActivityInitiatedBy `json:"initiatedBy" api:"required"`
	IssueCount         float64                                           `json:"issueCount" api:"required"`
	IssueID            string                                            `json:"issueId" api:"required"`
	LastError          string                                            `json:"lastError" api:"required,nullable"`
	PolicyID           string                                            `json:"policyId" api:"required"`
	Status             ObservabilityIssueListResponseActivityStatus      `json:"status" api:"required"`
	TerminalAt         float64                                           `json:"terminalAt" api:"required,nullable"`
	Threshold          float64                                           `json:"threshold" api:"required,nullable"`
	Type               ObservabilityIssueListResponseActivityType        `json:"type" api:"required"`
	AutomationID       string                                            `json:"automationId"`
	AutomationName     string                                            `json:"automationName"`
	AutomationRevision float64                                           `json:"automationRevision"`
	Scope              ObservabilityIssueListResponseActivityScope       `json:"scope"`
	TriggerType        ObservabilityIssueListResponseActivityTriggerType `json:"triggerType"`
	JSON               observabilityIssueListResponseActivityJSON        `json:"-"`
}

// observabilityIssueListResponseActivityJSON contains the JSON metadata for the
// struct [ObservabilityIssueListResponseActivity]
type observabilityIssueListResponseActivityJSON struct {
	ID                 apijson.Field
	Created            apijson.Field
	InactivitySeconds  apijson.Field
	InitiatedBy        apijson.Field
	IssueCount         apijson.Field
	IssueID            apijson.Field
	LastError          apijson.Field
	PolicyID           apijson.Field
	Status             apijson.Field
	TerminalAt         apijson.Field
	Threshold          apijson.Field
	Type               apijson.Field
	AutomationID       apijson.Field
	AutomationName     apijson.Field
	AutomationRevision apijson.Field
	Scope              apijson.Field
	TriggerType        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ObservabilityIssueListResponseActivity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueListResponseActivityJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueListResponseActivityInitiatedBy string

const (
	ObservabilityIssueListResponseActivityInitiatedByAutomatic ObservabilityIssueListResponseActivityInitiatedBy = "automatic"
	ObservabilityIssueListResponseActivityInitiatedByManual    ObservabilityIssueListResponseActivityInitiatedBy = "manual"
)

func (r ObservabilityIssueListResponseActivityInitiatedBy) IsKnown() bool {
	switch r {
	case ObservabilityIssueListResponseActivityInitiatedByAutomatic, ObservabilityIssueListResponseActivityInitiatedByManual:
		return true
	}
	return false
}

type ObservabilityIssueListResponseActivityStatus string

const (
	ObservabilityIssueListResponseActivityStatusPending   ObservabilityIssueListResponseActivityStatus = "pending"
	ObservabilityIssueListResponseActivityStatusSucceeded ObservabilityIssueListResponseActivityStatus = "succeeded"
	ObservabilityIssueListResponseActivityStatusFailed    ObservabilityIssueListResponseActivityStatus = "failed"
)

func (r ObservabilityIssueListResponseActivityStatus) IsKnown() bool {
	switch r {
	case ObservabilityIssueListResponseActivityStatusPending, ObservabilityIssueListResponseActivityStatusSucceeded, ObservabilityIssueListResponseActivityStatusFailed:
		return true
	}
	return false
}

type ObservabilityIssueListResponseActivityType string

const (
	ObservabilityIssueListResponseActivityTypeNotification ObservabilityIssueListResponseActivityType = "notification"
)

func (r ObservabilityIssueListResponseActivityType) IsKnown() bool {
	switch r {
	case ObservabilityIssueListResponseActivityTypeNotification:
		return true
	}
	return false
}

type ObservabilityIssueListResponseActivityScope string

const (
	ObservabilityIssueListResponseActivityScopeAccount ObservabilityIssueListResponseActivityScope = "account"
	ObservabilityIssueListResponseActivityScopeService ObservabilityIssueListResponseActivityScope = "service"
)

func (r ObservabilityIssueListResponseActivityScope) IsKnown() bool {
	switch r {
	case ObservabilityIssueListResponseActivityScopeAccount, ObservabilityIssueListResponseActivityScopeService:
		return true
	}
	return false
}

type ObservabilityIssueListResponseActivityTriggerType string

const (
	ObservabilityIssueListResponseActivityTriggerTypeOccurrenceThreshold       ObservabilityIssueListResponseActivityTriggerType = "occurrence_threshold"
	ObservabilityIssueListResponseActivityTriggerTypeRecurrenceAfterInactivity ObservabilityIssueListResponseActivityTriggerType = "recurrence_after_inactivity"
)

func (r ObservabilityIssueListResponseActivityTriggerType) IsKnown() bool {
	switch r {
	case ObservabilityIssueListResponseActivityTriggerTypeOccurrenceThreshold, ObservabilityIssueListResponseActivityTriggerTypeRecurrenceAfterInactivity:
		return true
	}
	return false
}

type ObservabilityIssueListResponseStatus string

const (
	ObservabilityIssueListResponseStatusActive   ObservabilityIssueListResponseStatus = "active"
	ObservabilityIssueListResponseStatusResolved ObservabilityIssueListResponseStatus = "resolved"
	ObservabilityIssueListResponseStatusIgnored  ObservabilityIssueListResponseStatus = "ignored"
)

func (r ObservabilityIssueListResponseStatus) IsKnown() bool {
	switch r {
	case ObservabilityIssueListResponseStatusActive, ObservabilityIssueListResponseStatusResolved, ObservabilityIssueListResponseStatusIgnored:
		return true
	}
	return false
}

type ObservabilityIssueListResponseMember struct {
	Fingerprint string                                   `json:"fingerprint" api:"required"`
	IssueID     string                                   `json:"issueId" api:"required"`
	Title       string                                   `json:"title" api:"required"`
	JSON        observabilityIssueListResponseMemberJSON `json:"-"`
}

// observabilityIssueListResponseMemberJSON contains the JSON metadata for the
// struct [ObservabilityIssueListResponseMember]
type observabilityIssueListResponseMemberJSON struct {
	Fingerprint apijson.Field
	IssueID     apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueListResponseMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueListResponseMemberJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAddResponse struct {
	ID            string                                `json:"id" api:"required"`
	Count         float64                               `json:"count" api:"required"`
	Created       float64                               `json:"created" api:"required"`
	FirstObserved float64                               `json:"firstObserved" api:"required"`
	LastObserved  float64                               `json:"lastObserved" api:"required"`
	Service       string                                `json:"service" api:"required"`
	Status        ObservabilityIssueAddResponseStatus   `json:"status" api:"required"`
	StatusUpdated float64                               `json:"statusUpdated" api:"required"`
	Title         string                                `json:"title" api:"required"`
	Updated       float64                               `json:"updated" api:"required"`
	Fingerprint   string                                `json:"fingerprint"`
	Members       []ObservabilityIssueAddResponseMember `json:"members"`
	Type          string                                `json:"type"`
	JSON          observabilityIssueAddResponseJSON     `json:"-"`
}

// observabilityIssueAddResponseJSON contains the JSON metadata for the struct
// [ObservabilityIssueAddResponse]
type observabilityIssueAddResponseJSON struct {
	ID            apijson.Field
	Count         apijson.Field
	Created       apijson.Field
	FirstObserved apijson.Field
	LastObserved  apijson.Field
	Service       apijson.Field
	Status        apijson.Field
	StatusUpdated apijson.Field
	Title         apijson.Field
	Updated       apijson.Field
	Fingerprint   apijson.Field
	Members       apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityIssueAddResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAddResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAddResponseStatus string

const (
	ObservabilityIssueAddResponseStatusActive   ObservabilityIssueAddResponseStatus = "active"
	ObservabilityIssueAddResponseStatusResolved ObservabilityIssueAddResponseStatus = "resolved"
	ObservabilityIssueAddResponseStatusIgnored  ObservabilityIssueAddResponseStatus = "ignored"
)

func (r ObservabilityIssueAddResponseStatus) IsKnown() bool {
	switch r {
	case ObservabilityIssueAddResponseStatusActive, ObservabilityIssueAddResponseStatusResolved, ObservabilityIssueAddResponseStatusIgnored:
		return true
	}
	return false
}

type ObservabilityIssueAddResponseMember struct {
	Fingerprint string                                  `json:"fingerprint" api:"required"`
	IssueID     string                                  `json:"issueId" api:"required"`
	Title       string                                  `json:"title" api:"required"`
	JSON        observabilityIssueAddResponseMemberJSON `json:"-"`
}

// observabilityIssueAddResponseMemberJSON contains the JSON metadata for the
// struct [ObservabilityIssueAddResponseMember]
type observabilityIssueAddResponseMemberJSON struct {
	Fingerprint apijson.Field
	IssueID     apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAddResponseMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAddResponseMemberJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueGetResponse struct {
	Issue ObservabilityIssueGetResponseIssue `json:"issue" api:"required"`
	JSON  observabilityIssueGetResponseJSON  `json:"-"`
}

// observabilityIssueGetResponseJSON contains the JSON metadata for the struct
// [ObservabilityIssueGetResponse]
type observabilityIssueGetResponseJSON struct {
	Issue       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueGetResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueGetResponseIssue struct {
	ID            string                                       `json:"id" api:"required"`
	Activity      []ObservabilityIssueGetResponseIssueActivity `json:"activity" api:"required"`
	Count         float64                                      `json:"count" api:"required"`
	Created       float64                                      `json:"created" api:"required"`
	FirstObserved float64                                      `json:"firstObserved" api:"required"`
	LastObserved  float64                                      `json:"lastObserved" api:"required"`
	Service       string                                       `json:"service" api:"required"`
	Status        ObservabilityIssueGetResponseIssueStatus     `json:"status" api:"required"`
	StatusUpdated float64                                      `json:"statusUpdated" api:"required"`
	Title         string                                       `json:"title" api:"required"`
	Updated       float64                                      `json:"updated" api:"required"`
	Fingerprint   string                                       `json:"fingerprint"`
	Members       []ObservabilityIssueGetResponseIssueMember   `json:"members"`
	Type          string                                       `json:"type"`
	JSON          observabilityIssueGetResponseIssueJSON       `json:"-"`
}

// observabilityIssueGetResponseIssueJSON contains the JSON metadata for the struct
// [ObservabilityIssueGetResponseIssue]
type observabilityIssueGetResponseIssueJSON struct {
	ID            apijson.Field
	Activity      apijson.Field
	Count         apijson.Field
	Created       apijson.Field
	FirstObserved apijson.Field
	LastObserved  apijson.Field
	Service       apijson.Field
	Status        apijson.Field
	StatusUpdated apijson.Field
	Title         apijson.Field
	Updated       apijson.Field
	Fingerprint   apijson.Field
	Members       apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityIssueGetResponseIssue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueGetResponseIssueJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueGetResponseIssueActivity struct {
	ID                 string                                                `json:"id" api:"required"`
	Created            float64                                               `json:"created" api:"required"`
	InactivitySeconds  float64                                               `json:"inactivitySeconds" api:"required,nullable"`
	InitiatedBy        ObservabilityIssueGetResponseIssueActivityInitiatedBy `json:"initiatedBy" api:"required"`
	IssueCount         float64                                               `json:"issueCount" api:"required"`
	IssueID            string                                                `json:"issueId" api:"required"`
	LastError          string                                                `json:"lastError" api:"required,nullable"`
	PolicyID           string                                                `json:"policyId" api:"required"`
	Status             ObservabilityIssueGetResponseIssueActivityStatus      `json:"status" api:"required"`
	TerminalAt         float64                                               `json:"terminalAt" api:"required,nullable"`
	Threshold          float64                                               `json:"threshold" api:"required,nullable"`
	Type               ObservabilityIssueGetResponseIssueActivityType        `json:"type" api:"required"`
	AutomationID       string                                                `json:"automationId"`
	AutomationName     string                                                `json:"automationName"`
	AutomationRevision float64                                               `json:"automationRevision"`
	Scope              ObservabilityIssueGetResponseIssueActivityScope       `json:"scope"`
	TriggerType        ObservabilityIssueGetResponseIssueActivityTriggerType `json:"triggerType"`
	JSON               observabilityIssueGetResponseIssueActivityJSON        `json:"-"`
}

// observabilityIssueGetResponseIssueActivityJSON contains the JSON metadata for
// the struct [ObservabilityIssueGetResponseIssueActivity]
type observabilityIssueGetResponseIssueActivityJSON struct {
	ID                 apijson.Field
	Created            apijson.Field
	InactivitySeconds  apijson.Field
	InitiatedBy        apijson.Field
	IssueCount         apijson.Field
	IssueID            apijson.Field
	LastError          apijson.Field
	PolicyID           apijson.Field
	Status             apijson.Field
	TerminalAt         apijson.Field
	Threshold          apijson.Field
	Type               apijson.Field
	AutomationID       apijson.Field
	AutomationName     apijson.Field
	AutomationRevision apijson.Field
	Scope              apijson.Field
	TriggerType        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ObservabilityIssueGetResponseIssueActivity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueGetResponseIssueActivityJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueGetResponseIssueActivityInitiatedBy string

const (
	ObservabilityIssueGetResponseIssueActivityInitiatedByAutomatic ObservabilityIssueGetResponseIssueActivityInitiatedBy = "automatic"
	ObservabilityIssueGetResponseIssueActivityInitiatedByManual    ObservabilityIssueGetResponseIssueActivityInitiatedBy = "manual"
)

func (r ObservabilityIssueGetResponseIssueActivityInitiatedBy) IsKnown() bool {
	switch r {
	case ObservabilityIssueGetResponseIssueActivityInitiatedByAutomatic, ObservabilityIssueGetResponseIssueActivityInitiatedByManual:
		return true
	}
	return false
}

type ObservabilityIssueGetResponseIssueActivityStatus string

const (
	ObservabilityIssueGetResponseIssueActivityStatusPending   ObservabilityIssueGetResponseIssueActivityStatus = "pending"
	ObservabilityIssueGetResponseIssueActivityStatusSucceeded ObservabilityIssueGetResponseIssueActivityStatus = "succeeded"
	ObservabilityIssueGetResponseIssueActivityStatusFailed    ObservabilityIssueGetResponseIssueActivityStatus = "failed"
)

func (r ObservabilityIssueGetResponseIssueActivityStatus) IsKnown() bool {
	switch r {
	case ObservabilityIssueGetResponseIssueActivityStatusPending, ObservabilityIssueGetResponseIssueActivityStatusSucceeded, ObservabilityIssueGetResponseIssueActivityStatusFailed:
		return true
	}
	return false
}

type ObservabilityIssueGetResponseIssueActivityType string

const (
	ObservabilityIssueGetResponseIssueActivityTypeNotification ObservabilityIssueGetResponseIssueActivityType = "notification"
)

func (r ObservabilityIssueGetResponseIssueActivityType) IsKnown() bool {
	switch r {
	case ObservabilityIssueGetResponseIssueActivityTypeNotification:
		return true
	}
	return false
}

type ObservabilityIssueGetResponseIssueActivityScope string

const (
	ObservabilityIssueGetResponseIssueActivityScopeAccount ObservabilityIssueGetResponseIssueActivityScope = "account"
	ObservabilityIssueGetResponseIssueActivityScopeService ObservabilityIssueGetResponseIssueActivityScope = "service"
)

func (r ObservabilityIssueGetResponseIssueActivityScope) IsKnown() bool {
	switch r {
	case ObservabilityIssueGetResponseIssueActivityScopeAccount, ObservabilityIssueGetResponseIssueActivityScopeService:
		return true
	}
	return false
}

type ObservabilityIssueGetResponseIssueActivityTriggerType string

const (
	ObservabilityIssueGetResponseIssueActivityTriggerTypeOccurrenceThreshold       ObservabilityIssueGetResponseIssueActivityTriggerType = "occurrence_threshold"
	ObservabilityIssueGetResponseIssueActivityTriggerTypeRecurrenceAfterInactivity ObservabilityIssueGetResponseIssueActivityTriggerType = "recurrence_after_inactivity"
)

func (r ObservabilityIssueGetResponseIssueActivityTriggerType) IsKnown() bool {
	switch r {
	case ObservabilityIssueGetResponseIssueActivityTriggerTypeOccurrenceThreshold, ObservabilityIssueGetResponseIssueActivityTriggerTypeRecurrenceAfterInactivity:
		return true
	}
	return false
}

type ObservabilityIssueGetResponseIssueStatus string

const (
	ObservabilityIssueGetResponseIssueStatusActive   ObservabilityIssueGetResponseIssueStatus = "active"
	ObservabilityIssueGetResponseIssueStatusResolved ObservabilityIssueGetResponseIssueStatus = "resolved"
	ObservabilityIssueGetResponseIssueStatusIgnored  ObservabilityIssueGetResponseIssueStatus = "ignored"
)

func (r ObservabilityIssueGetResponseIssueStatus) IsKnown() bool {
	switch r {
	case ObservabilityIssueGetResponseIssueStatusActive, ObservabilityIssueGetResponseIssueStatusResolved, ObservabilityIssueGetResponseIssueStatusIgnored:
		return true
	}
	return false
}

type ObservabilityIssueGetResponseIssueMember struct {
	Fingerprint string                                       `json:"fingerprint" api:"required"`
	IssueID     string                                       `json:"issueId" api:"required"`
	Title       string                                       `json:"title" api:"required"`
	JSON        observabilityIssueGetResponseIssueMemberJSON `json:"-"`
}

// observabilityIssueGetResponseIssueMemberJSON contains the JSON metadata for the
// struct [ObservabilityIssueGetResponseIssueMember]
type observabilityIssueGetResponseIssueMemberJSON struct {
	Fingerprint apijson.Field
	IssueID     apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueGetResponseIssueMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueGetResponseIssueMemberJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueGroupResponse struct {
	ID            string                                  `json:"id" api:"required"`
	Count         float64                                 `json:"count" api:"required"`
	Created       float64                                 `json:"created" api:"required"`
	FirstObserved float64                                 `json:"firstObserved" api:"required"`
	LastObserved  float64                                 `json:"lastObserved" api:"required"`
	Service       string                                  `json:"service" api:"required"`
	Status        ObservabilityIssueGroupResponseStatus   `json:"status" api:"required"`
	StatusUpdated float64                                 `json:"statusUpdated" api:"required"`
	Title         string                                  `json:"title" api:"required"`
	Updated       float64                                 `json:"updated" api:"required"`
	Fingerprint   string                                  `json:"fingerprint"`
	Members       []ObservabilityIssueGroupResponseMember `json:"members"`
	Type          string                                  `json:"type"`
	JSON          observabilityIssueGroupResponseJSON     `json:"-"`
}

// observabilityIssueGroupResponseJSON contains the JSON metadata for the struct
// [ObservabilityIssueGroupResponse]
type observabilityIssueGroupResponseJSON struct {
	ID            apijson.Field
	Count         apijson.Field
	Created       apijson.Field
	FirstObserved apijson.Field
	LastObserved  apijson.Field
	Service       apijson.Field
	Status        apijson.Field
	StatusUpdated apijson.Field
	Title         apijson.Field
	Updated       apijson.Field
	Fingerprint   apijson.Field
	Members       apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityIssueGroupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueGroupResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueGroupResponseStatus string

const (
	ObservabilityIssueGroupResponseStatusActive   ObservabilityIssueGroupResponseStatus = "active"
	ObservabilityIssueGroupResponseStatusResolved ObservabilityIssueGroupResponseStatus = "resolved"
	ObservabilityIssueGroupResponseStatusIgnored  ObservabilityIssueGroupResponseStatus = "ignored"
)

func (r ObservabilityIssueGroupResponseStatus) IsKnown() bool {
	switch r {
	case ObservabilityIssueGroupResponseStatusActive, ObservabilityIssueGroupResponseStatusResolved, ObservabilityIssueGroupResponseStatusIgnored:
		return true
	}
	return false
}

type ObservabilityIssueGroupResponseMember struct {
	Fingerprint string                                    `json:"fingerprint" api:"required"`
	IssueID     string                                    `json:"issueId" api:"required"`
	Title       string                                    `json:"title" api:"required"`
	JSON        observabilityIssueGroupResponseMemberJSON `json:"-"`
}

// observabilityIssueGroupResponseMemberJSON contains the JSON metadata for the
// struct [ObservabilityIssueGroupResponseMember]
type observabilityIssueGroupResponseMemberJSON struct {
	Fingerprint apijson.Field
	IssueID     apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueGroupResponseMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueGroupResponseMemberJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueNotifyResponse struct {
	Run  ObservabilityIssueNotifyResponseRun  `json:"run" api:"required"`
	JSON observabilityIssueNotifyResponseJSON `json:"-"`
}

// observabilityIssueNotifyResponseJSON contains the JSON metadata for the struct
// [ObservabilityIssueNotifyResponse]
type observabilityIssueNotifyResponseJSON struct {
	Run         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueNotifyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueNotifyResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueNotifyResponseRun struct {
	ID                 string                                         `json:"id" api:"required"`
	Created            float64                                        `json:"created" api:"required"`
	InactivitySeconds  float64                                        `json:"inactivitySeconds" api:"required,nullable"`
	InitiatedBy        ObservabilityIssueNotifyResponseRunInitiatedBy `json:"initiatedBy" api:"required"`
	IssueCount         float64                                        `json:"issueCount" api:"required"`
	IssueID            string                                         `json:"issueId" api:"required"`
	LastError          string                                         `json:"lastError" api:"required,nullable"`
	PolicyID           string                                         `json:"policyId" api:"required"`
	Status             ObservabilityIssueNotifyResponseRunStatus      `json:"status" api:"required"`
	TerminalAt         float64                                        `json:"terminalAt" api:"required,nullable"`
	Threshold          float64                                        `json:"threshold" api:"required,nullable"`
	AutomationID       string                                         `json:"automationId"`
	AutomationName     string                                         `json:"automationName"`
	AutomationRevision float64                                        `json:"automationRevision"`
	Scope              ObservabilityIssueNotifyResponseRunScope       `json:"scope"`
	TriggerType        ObservabilityIssueNotifyResponseRunTriggerType `json:"triggerType"`
	JSON               observabilityIssueNotifyResponseRunJSON        `json:"-"`
}

// observabilityIssueNotifyResponseRunJSON contains the JSON metadata for the
// struct [ObservabilityIssueNotifyResponseRun]
type observabilityIssueNotifyResponseRunJSON struct {
	ID                 apijson.Field
	Created            apijson.Field
	InactivitySeconds  apijson.Field
	InitiatedBy        apijson.Field
	IssueCount         apijson.Field
	IssueID            apijson.Field
	LastError          apijson.Field
	PolicyID           apijson.Field
	Status             apijson.Field
	TerminalAt         apijson.Field
	Threshold          apijson.Field
	AutomationID       apijson.Field
	AutomationName     apijson.Field
	AutomationRevision apijson.Field
	Scope              apijson.Field
	TriggerType        apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ObservabilityIssueNotifyResponseRun) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueNotifyResponseRunJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueNotifyResponseRunInitiatedBy string

const (
	ObservabilityIssueNotifyResponseRunInitiatedByAutomatic ObservabilityIssueNotifyResponseRunInitiatedBy = "automatic"
	ObservabilityIssueNotifyResponseRunInitiatedByManual    ObservabilityIssueNotifyResponseRunInitiatedBy = "manual"
)

func (r ObservabilityIssueNotifyResponseRunInitiatedBy) IsKnown() bool {
	switch r {
	case ObservabilityIssueNotifyResponseRunInitiatedByAutomatic, ObservabilityIssueNotifyResponseRunInitiatedByManual:
		return true
	}
	return false
}

type ObservabilityIssueNotifyResponseRunStatus string

const (
	ObservabilityIssueNotifyResponseRunStatusPending   ObservabilityIssueNotifyResponseRunStatus = "pending"
	ObservabilityIssueNotifyResponseRunStatusSucceeded ObservabilityIssueNotifyResponseRunStatus = "succeeded"
	ObservabilityIssueNotifyResponseRunStatusFailed    ObservabilityIssueNotifyResponseRunStatus = "failed"
)

func (r ObservabilityIssueNotifyResponseRunStatus) IsKnown() bool {
	switch r {
	case ObservabilityIssueNotifyResponseRunStatusPending, ObservabilityIssueNotifyResponseRunStatusSucceeded, ObservabilityIssueNotifyResponseRunStatusFailed:
		return true
	}
	return false
}

type ObservabilityIssueNotifyResponseRunScope string

const (
	ObservabilityIssueNotifyResponseRunScopeAccount ObservabilityIssueNotifyResponseRunScope = "account"
	ObservabilityIssueNotifyResponseRunScopeService ObservabilityIssueNotifyResponseRunScope = "service"
)

func (r ObservabilityIssueNotifyResponseRunScope) IsKnown() bool {
	switch r {
	case ObservabilityIssueNotifyResponseRunScopeAccount, ObservabilityIssueNotifyResponseRunScopeService:
		return true
	}
	return false
}

type ObservabilityIssueNotifyResponseRunTriggerType string

const (
	ObservabilityIssueNotifyResponseRunTriggerTypeOccurrenceThreshold       ObservabilityIssueNotifyResponseRunTriggerType = "occurrence_threshold"
	ObservabilityIssueNotifyResponseRunTriggerTypeRecurrenceAfterInactivity ObservabilityIssueNotifyResponseRunTriggerType = "recurrence_after_inactivity"
)

func (r ObservabilityIssueNotifyResponseRunTriggerType) IsKnown() bool {
	switch r {
	case ObservabilityIssueNotifyResponseRunTriggerTypeOccurrenceThreshold, ObservabilityIssueNotifyResponseRunTriggerTypeRecurrenceAfterInactivity:
		return true
	}
	return false
}

type ObservabilityIssueOccurrencesResponse struct {
	// Unique occurrence ID.
	ID         string                                          `json:"id" api:"required"`
	Error      ObservabilityIssueOccurrencesResponseError      `json:"error" api:"required"`
	Invocation ObservabilityIssueOccurrencesResponseInvocation `json:"invocation" api:"required"`
	// Occurrence timestamp as Unix epoch milliseconds.
	Timestamp float64                                      `json:"timestamp" api:"required"`
	Trail     []ObservabilityIssueOccurrencesResponseTrail `json:"trail" api:"required"`
	Worker    ObservabilityIssueOccurrencesResponseWorker  `json:"worker" api:"required"`
	Context   ObservabilityIssueOccurrencesResponseContext `json:"context"`
	Request   ObservabilityIssueOccurrencesResponseRequest `json:"request"`
	JSON      observabilityIssueOccurrencesResponseJSON    `json:"-"`
}

// observabilityIssueOccurrencesResponseJSON contains the JSON metadata for the
// struct [ObservabilityIssueOccurrencesResponse]
type observabilityIssueOccurrencesResponseJSON struct {
	ID          apijson.Field
	Error       apijson.Field
	Invocation  apijson.Field
	Timestamp   apijson.Field
	Trail       apijson.Field
	Worker      apijson.Field
	Context     apijson.Field
	Request     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueOccurrencesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueOccurrencesResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueOccurrencesResponseError struct {
	Handled   bool                                                `json:"handled" api:"required"`
	Mechanism ObservabilityIssueOccurrencesResponseErrorMechanism `json:"mechanism" api:"required"`
	Message   string                                              `json:"message" api:"required"`
	Name      string                                              `json:"name" api:"required"`
	Signature string                                              `json:"signature" api:"required"`
	Stack     string                                              `json:"stack"`
	JSON      observabilityIssueOccurrencesResponseErrorJSON      `json:"-"`
}

// observabilityIssueOccurrencesResponseErrorJSON contains the JSON metadata for
// the struct [ObservabilityIssueOccurrencesResponseError]
type observabilityIssueOccurrencesResponseErrorJSON struct {
	Handled     apijson.Field
	Mechanism   apijson.Field
	Message     apijson.Field
	Name        apijson.Field
	Signature   apijson.Field
	Stack       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueOccurrencesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueOccurrencesResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueOccurrencesResponseErrorMechanism string

const (
	ObservabilityIssueOccurrencesResponseErrorMechanismStructuredLog ObservabilityIssueOccurrencesResponseErrorMechanism = "structured-log"
	ObservabilityIssueOccurrencesResponseErrorMechanismErrorLog      ObservabilityIssueOccurrencesResponseErrorMechanism = "error-log"
	ObservabilityIssueOccurrencesResponseErrorMechanismStackLog      ObservabilityIssueOccurrencesResponseErrorMechanism = "stack-log"
	ObservabilityIssueOccurrencesResponseErrorMechanismHTTPStatus    ObservabilityIssueOccurrencesResponseErrorMechanism = "http-status"
	ObservabilityIssueOccurrencesResponseErrorMechanismException     ObservabilityIssueOccurrencesResponseErrorMechanism = "exception"
	ObservabilityIssueOccurrencesResponseErrorMechanismOutcome       ObservabilityIssueOccurrencesResponseErrorMechanism = "outcome"
)

func (r ObservabilityIssueOccurrencesResponseErrorMechanism) IsKnown() bool {
	switch r {
	case ObservabilityIssueOccurrencesResponseErrorMechanismStructuredLog, ObservabilityIssueOccurrencesResponseErrorMechanismErrorLog, ObservabilityIssueOccurrencesResponseErrorMechanismStackLog, ObservabilityIssueOccurrencesResponseErrorMechanismHTTPStatus, ObservabilityIssueOccurrencesResponseErrorMechanismException, ObservabilityIssueOccurrencesResponseErrorMechanismOutcome:
		return true
	}
	return false
}

type ObservabilityIssueOccurrencesResponseInvocation struct {
	// Worker invocation request ID.
	ID            string  `json:"id" api:"required"`
	TraceID       string  `json:"traceId" api:"required"`
	Type          string  `json:"type" api:"required"`
	BatchSize     float64 `json:"batchSize"`
	Cron          string  `json:"cron"`
	JsrpcMethod   string  `json:"jsrpcMethod"`
	MailFrom      string  `json:"mailFrom"`
	Method        string  `json:"method"`
	Path          string  `json:"path"`
	Queue         string  `json:"queue"`
	RawSize       float64 `json:"rawSize"`
	RayID         string  `json:"rayId"`
	RcptTo        string  `json:"rcptTo"`
	ScheduledTime float64 `json:"scheduledTime"`
	SpanID        string  `json:"spanId"`
	StatusCode    int64   `json:"statusCode"`
	// Redacted request URL without query parameters.
	URL                string                                              `json:"url"`
	WebsocketCloseCode float64                                             `json:"websocketCloseCode"`
	WebsocketEvent     string                                              `json:"websocketEvent"`
	JSON               observabilityIssueOccurrencesResponseInvocationJSON `json:"-"`
}

// observabilityIssueOccurrencesResponseInvocationJSON contains the JSON metadata
// for the struct [ObservabilityIssueOccurrencesResponseInvocation]
type observabilityIssueOccurrencesResponseInvocationJSON struct {
	ID                 apijson.Field
	TraceID            apijson.Field
	Type               apijson.Field
	BatchSize          apijson.Field
	Cron               apijson.Field
	JsrpcMethod        apijson.Field
	MailFrom           apijson.Field
	Method             apijson.Field
	Path               apijson.Field
	Queue              apijson.Field
	RawSize            apijson.Field
	RayID              apijson.Field
	RcptTo             apijson.Field
	ScheduledTime      apijson.Field
	SpanID             apijson.Field
	StatusCode         apijson.Field
	URL                apijson.Field
	WebsocketCloseCode apijson.Field
	WebsocketEvent     apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ObservabilityIssueOccurrencesResponseInvocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueOccurrencesResponseInvocationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueOccurrencesResponseTrail struct {
	Timestamp    float64                                        `json:"timestamp" api:"required"`
	Type         ObservabilityIssueOccurrencesResponseTrailType `json:"type" api:"required"`
	Level        string                                         `json:"level"`
	Message      string                                         `json:"message"`
	Name         string                                         `json:"name"`
	Ongoing      bool                                           `json:"ongoing"`
	ParentSpanID string                                         `json:"parentSpanId"`
	SpanID       string                                         `json:"spanId"`
	JSON         observabilityIssueOccurrencesResponseTrailJSON `json:"-"`
	union        ObservabilityIssueOccurrencesResponseTrailUnion
}

// observabilityIssueOccurrencesResponseTrailJSON contains the JSON metadata for
// the struct [ObservabilityIssueOccurrencesResponseTrail]
type observabilityIssueOccurrencesResponseTrailJSON struct {
	Timestamp    apijson.Field
	Type         apijson.Field
	Level        apijson.Field
	Message      apijson.Field
	Name         apijson.Field
	Ongoing      apijson.Field
	ParentSpanID apijson.Field
	SpanID       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r observabilityIssueOccurrencesResponseTrailJSON) RawJSON() string {
	return r.raw
}

func (r *ObservabilityIssueOccurrencesResponseTrail) UnmarshalJSON(data []byte) (err error) {
	*r = ObservabilityIssueOccurrencesResponseTrail{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObservabilityIssueOccurrencesResponseTrailUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ObservabilityIssueOccurrencesResponseTrailObject],
// [ObservabilityIssueOccurrencesResponseTrailObject].
func (r ObservabilityIssueOccurrencesResponseTrail) AsUnion() ObservabilityIssueOccurrencesResponseTrailUnion {
	return r.union
}

// Union satisfied by [ObservabilityIssueOccurrencesResponseTrailObject] or
// [ObservabilityIssueOccurrencesResponseTrailObject].
type ObservabilityIssueOccurrencesResponseTrailUnion interface {
	implementsObservabilityIssueOccurrencesResponseTrail()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityIssueOccurrencesResponseTrailUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityIssueOccurrencesResponseTrailObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityIssueOccurrencesResponseTrailObject{}),
		},
	)
}

type ObservabilityIssueOccurrencesResponseTrailObject struct {
	Level     string                                               `json:"level" api:"required"`
	Message   string                                               `json:"message" api:"required"`
	Timestamp float64                                              `json:"timestamp" api:"required"`
	Type      ObservabilityIssueOccurrencesResponseTrailObjectType `json:"type" api:"required"`
	JSON      observabilityIssueOccurrencesResponseTrailObjectJSON `json:"-"`
}

// observabilityIssueOccurrencesResponseTrailObjectJSON contains the JSON metadata
// for the struct [ObservabilityIssueOccurrencesResponseTrailObject]
type observabilityIssueOccurrencesResponseTrailObjectJSON struct {
	Level       apijson.Field
	Message     apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueOccurrencesResponseTrailObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueOccurrencesResponseTrailObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilityIssueOccurrencesResponseTrailObject) implementsObservabilityIssueOccurrencesResponseTrail() {
}

type ObservabilityIssueOccurrencesResponseTrailObjectType string

const (
	ObservabilityIssueOccurrencesResponseTrailObjectTypeLog ObservabilityIssueOccurrencesResponseTrailObjectType = "log"
)

func (r ObservabilityIssueOccurrencesResponseTrailObjectType) IsKnown() bool {
	switch r {
	case ObservabilityIssueOccurrencesResponseTrailObjectTypeLog:
		return true
	}
	return false
}

type ObservabilityIssueOccurrencesResponseTrailType string

const (
	ObservabilityIssueOccurrencesResponseTrailTypeLog  ObservabilityIssueOccurrencesResponseTrailType = "log"
	ObservabilityIssueOccurrencesResponseTrailTypeSpan ObservabilityIssueOccurrencesResponseTrailType = "span"
)

func (r ObservabilityIssueOccurrencesResponseTrailType) IsKnown() bool {
	switch r {
	case ObservabilityIssueOccurrencesResponseTrailTypeLog, ObservabilityIssueOccurrencesResponseTrailTypeSpan:
		return true
	}
	return false
}

type ObservabilityIssueOccurrencesResponseWorker struct {
	ExecutionModel    string                                                   `json:"executionModel" api:"required"`
	ScriptName        string                                                   `json:"scriptName" api:"required"`
	ScriptTags        []string                                                 `json:"scriptTags" api:"required"`
	DispatchNamespace string                                                   `json:"dispatchNamespace"`
	Entrypoint        string                                                   `json:"entrypoint"`
	ScriptVersion     ObservabilityIssueOccurrencesResponseWorkerScriptVersion `json:"scriptVersion"`
	JSON              observabilityIssueOccurrencesResponseWorkerJSON          `json:"-"`
}

// observabilityIssueOccurrencesResponseWorkerJSON contains the JSON metadata for
// the struct [ObservabilityIssueOccurrencesResponseWorker]
type observabilityIssueOccurrencesResponseWorkerJSON struct {
	ExecutionModel    apijson.Field
	ScriptName        apijson.Field
	ScriptTags        apijson.Field
	DispatchNamespace apijson.Field
	Entrypoint        apijson.Field
	ScriptVersion     apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityIssueOccurrencesResponseWorker) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueOccurrencesResponseWorkerJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueOccurrencesResponseWorkerScriptVersion struct {
	ID   string                                                       `json:"id" api:"required"`
	Tag  string                                                       `json:"tag"`
	JSON observabilityIssueOccurrencesResponseWorkerScriptVersionJSON `json:"-"`
}

// observabilityIssueOccurrencesResponseWorkerScriptVersionJSON contains the JSON
// metadata for the struct
// [ObservabilityIssueOccurrencesResponseWorkerScriptVersion]
type observabilityIssueOccurrencesResponseWorkerScriptVersionJSON struct {
	ID          apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueOccurrencesResponseWorkerScriptVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueOccurrencesResponseWorkerScriptVersionJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueOccurrencesResponseContext struct {
	Account string                                           `json:"account"`
	Session string                                           `json:"session"`
	User    string                                           `json:"user"`
	JSON    observabilityIssueOccurrencesResponseContextJSON `json:"-"`
}

// observabilityIssueOccurrencesResponseContextJSON contains the JSON metadata for
// the struct [ObservabilityIssueOccurrencesResponseContext]
type observabilityIssueOccurrencesResponseContextJSON struct {
	Account     apijson.Field
	Session     apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueOccurrencesResponseContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueOccurrencesResponseContextJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueOccurrencesResponseRequest struct {
	Headers map[string]string                                `json:"headers" api:"required"`
	JSON    observabilityIssueOccurrencesResponseRequestJSON `json:"-"`
}

// observabilityIssueOccurrencesResponseRequestJSON contains the JSON metadata for
// the struct [ObservabilityIssueOccurrencesResponseRequest]
type observabilityIssueOccurrencesResponseRequestJSON struct {
	Headers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueOccurrencesResponseRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueOccurrencesResponseRequestJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueRemoveResponse struct {
	ID            string                                   `json:"id" api:"required"`
	Count         float64                                  `json:"count" api:"required"`
	Created       float64                                  `json:"created" api:"required"`
	FirstObserved float64                                  `json:"firstObserved" api:"required"`
	LastObserved  float64                                  `json:"lastObserved" api:"required"`
	Service       string                                   `json:"service" api:"required"`
	Status        ObservabilityIssueRemoveResponseStatus   `json:"status" api:"required"`
	StatusUpdated float64                                  `json:"statusUpdated" api:"required"`
	Title         string                                   `json:"title" api:"required"`
	Updated       float64                                  `json:"updated" api:"required"`
	Fingerprint   string                                   `json:"fingerprint"`
	Members       []ObservabilityIssueRemoveResponseMember `json:"members"`
	Type          string                                   `json:"type"`
	JSON          observabilityIssueRemoveResponseJSON     `json:"-"`
}

// observabilityIssueRemoveResponseJSON contains the JSON metadata for the struct
// [ObservabilityIssueRemoveResponse]
type observabilityIssueRemoveResponseJSON struct {
	ID            apijson.Field
	Count         apijson.Field
	Created       apijson.Field
	FirstObserved apijson.Field
	LastObserved  apijson.Field
	Service       apijson.Field
	Status        apijson.Field
	StatusUpdated apijson.Field
	Title         apijson.Field
	Updated       apijson.Field
	Fingerprint   apijson.Field
	Members       apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityIssueRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueRemoveResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueRemoveResponseStatus string

const (
	ObservabilityIssueRemoveResponseStatusActive   ObservabilityIssueRemoveResponseStatus = "active"
	ObservabilityIssueRemoveResponseStatusResolved ObservabilityIssueRemoveResponseStatus = "resolved"
	ObservabilityIssueRemoveResponseStatusIgnored  ObservabilityIssueRemoveResponseStatus = "ignored"
)

func (r ObservabilityIssueRemoveResponseStatus) IsKnown() bool {
	switch r {
	case ObservabilityIssueRemoveResponseStatusActive, ObservabilityIssueRemoveResponseStatusResolved, ObservabilityIssueRemoveResponseStatusIgnored:
		return true
	}
	return false
}

type ObservabilityIssueRemoveResponseMember struct {
	Fingerprint string                                     `json:"fingerprint" api:"required"`
	IssueID     string                                     `json:"issueId" api:"required"`
	Title       string                                     `json:"title" api:"required"`
	JSON        observabilityIssueRemoveResponseMemberJSON `json:"-"`
}

// observabilityIssueRemoveResponseMemberJSON contains the JSON metadata for the
// struct [ObservabilityIssueRemoveResponseMember]
type observabilityIssueRemoveResponseMemberJSON struct {
	Fingerprint apijson.Field
	IssueID     apijson.Field
	Title       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueRemoveResponseMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueRemoveResponseMemberJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueSummaryResponse struct {
	ActiveIssues      float64                               `json:"activeIssues" api:"required"`
	ActiveOccurrences float64                               `json:"activeOccurrences" api:"required"`
	LastIssue         float64                               `json:"lastIssue" api:"required,nullable"`
	ResolvedIssues    float64                               `json:"resolvedIssues" api:"required"`
	JSON              observabilityIssueSummaryResponseJSON `json:"-"`
}

// observabilityIssueSummaryResponseJSON contains the JSON metadata for the struct
// [ObservabilityIssueSummaryResponse]
type observabilityIssueSummaryResponseJSON struct {
	ActiveIssues      apijson.Field
	ActiveOccurrences apijson.Field
	LastIssue         apijson.Field
	ResolvedIssues    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityIssueSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueSummaryResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueUngroupResponse struct {
	Deleted ObservabilityIssueUngroupResponseDeleted `json:"deleted" api:"required"`
	JSON    observabilityIssueUngroupResponseJSON    `json:"-"`
}

// observabilityIssueUngroupResponseJSON contains the JSON metadata for the struct
// [ObservabilityIssueUngroupResponse]
type observabilityIssueUngroupResponseJSON struct {
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueUngroupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueUngroupResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueUngroupResponseDeleted bool

const (
	ObservabilityIssueUngroupResponseDeletedTrue ObservabilityIssueUngroupResponseDeleted = true
)

func (r ObservabilityIssueUngroupResponseDeleted) IsKnown() bool {
	switch r {
	case ObservabilityIssueUngroupResponseDeletedTrue:
		return true
	}
	return false
}

type ObservabilityIssueUpdateParams struct {
	AccountID param.Field[string]                               `path:"account_id" api:"required"`
	Status    param.Field[ObservabilityIssueUpdateParamsStatus] `json:"status"`
	Title     param.Field[string]                               `json:"title"`
}

func (r ObservabilityIssueUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityIssueUpdateParamsStatus string

const (
	ObservabilityIssueUpdateParamsStatusActive   ObservabilityIssueUpdateParamsStatus = "active"
	ObservabilityIssueUpdateParamsStatusResolved ObservabilityIssueUpdateParamsStatus = "resolved"
	ObservabilityIssueUpdateParamsStatusIgnored  ObservabilityIssueUpdateParamsStatus = "ignored"
)

func (r ObservabilityIssueUpdateParamsStatus) IsKnown() bool {
	switch r {
	case ObservabilityIssueUpdateParamsStatusActive, ObservabilityIssueUpdateParamsStatusResolved, ObservabilityIssueUpdateParamsStatusIgnored:
		return true
	}
	return false
}

type ObservabilityIssueUpdateResponseEnvelope struct {
	Errors   []ObservabilityIssueUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityIssueUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityIssueUpdateResponse                   `json:"result" api:"required"`
	Success  ObservabilityIssueUpdateResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityIssueUpdateResponseEnvelopeJSON       `json:"-"`
}

// observabilityIssueUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ObservabilityIssueUpdateResponseEnvelope]
type observabilityIssueUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueUpdateResponseEnvelopeErrors struct {
	Message string                                             `json:"message" api:"required"`
	JSON    observabilityIssueUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityIssueUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ObservabilityIssueUpdateResponseEnvelopeErrors]
type observabilityIssueUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueUpdateResponseEnvelopeMessages struct {
	Message ObservabilityIssueUpdateResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityIssueUpdateResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityIssueUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ObservabilityIssueUpdateResponseEnvelopeMessages]
type observabilityIssueUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueUpdateResponseEnvelopeMessagesMessage string

const (
	ObservabilityIssueUpdateResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityIssueUpdateResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityIssueUpdateResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityIssueUpdateResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityIssueUpdateResponseEnvelopeSuccess bool

const (
	ObservabilityIssueUpdateResponseEnvelopeSuccessTrue ObservabilityIssueUpdateResponseEnvelopeSuccess = true
)

func (r ObservabilityIssueUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityIssueUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityIssueListParams struct {
	AccountID param.Field[string]                              `path:"account_id" api:"required"`
	Order     param.Field[ObservabilityIssueListParamsOrder]   `query:"order"`
	OrderBy   param.Field[ObservabilityIssueListParamsOrderBy] `query:"orderBy"`
	Page      param.Field[int64]                               `query:"page"`
	PerPage   param.Field[int64]                               `query:"perPage"`
	Search    param.Field[string]                              `query:"search"`
	Service   param.Field[string]                              `query:"service"`
	Status    param.Field[ObservabilityIssueListParamsStatus]  `query:"status"`
	Type      param.Field[string]                              `query:"type"`
}

// URLQuery serializes [ObservabilityIssueListParams]'s query parameters as
// `url.Values`.
func (r ObservabilityIssueListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ObservabilityIssueListParamsOrder string

const (
	ObservabilityIssueListParamsOrderAsc  ObservabilityIssueListParamsOrder = "asc"
	ObservabilityIssueListParamsOrderDesc ObservabilityIssueListParamsOrder = "desc"
)

func (r ObservabilityIssueListParamsOrder) IsKnown() bool {
	switch r {
	case ObservabilityIssueListParamsOrderAsc, ObservabilityIssueListParamsOrderDesc:
		return true
	}
	return false
}

type ObservabilityIssueListParamsOrderBy string

const (
	ObservabilityIssueListParamsOrderByID            ObservabilityIssueListParamsOrderBy = "id"
	ObservabilityIssueListParamsOrderByService       ObservabilityIssueListParamsOrderBy = "service"
	ObservabilityIssueListParamsOrderByTitle         ObservabilityIssueListParamsOrderBy = "title"
	ObservabilityIssueListParamsOrderByType          ObservabilityIssueListParamsOrderBy = "type"
	ObservabilityIssueListParamsOrderByStatus        ObservabilityIssueListParamsOrderBy = "status"
	ObservabilityIssueListParamsOrderByStatusUpdated ObservabilityIssueListParamsOrderBy = "statusUpdated"
	ObservabilityIssueListParamsOrderByCount         ObservabilityIssueListParamsOrderBy = "count"
	ObservabilityIssueListParamsOrderByFirstObserved ObservabilityIssueListParamsOrderBy = "firstObserved"
	ObservabilityIssueListParamsOrderByLastObserved  ObservabilityIssueListParamsOrderBy = "lastObserved"
	ObservabilityIssueListParamsOrderByCreated       ObservabilityIssueListParamsOrderBy = "created"
	ObservabilityIssueListParamsOrderByUpdated       ObservabilityIssueListParamsOrderBy = "updated"
)

func (r ObservabilityIssueListParamsOrderBy) IsKnown() bool {
	switch r {
	case ObservabilityIssueListParamsOrderByID, ObservabilityIssueListParamsOrderByService, ObservabilityIssueListParamsOrderByTitle, ObservabilityIssueListParamsOrderByType, ObservabilityIssueListParamsOrderByStatus, ObservabilityIssueListParamsOrderByStatusUpdated, ObservabilityIssueListParamsOrderByCount, ObservabilityIssueListParamsOrderByFirstObserved, ObservabilityIssueListParamsOrderByLastObserved, ObservabilityIssueListParamsOrderByCreated, ObservabilityIssueListParamsOrderByUpdated:
		return true
	}
	return false
}

type ObservabilityIssueListParamsStatus string

const (
	ObservabilityIssueListParamsStatusActive   ObservabilityIssueListParamsStatus = "active"
	ObservabilityIssueListParamsStatusResolved ObservabilityIssueListParamsStatus = "resolved"
	ObservabilityIssueListParamsStatusIgnored  ObservabilityIssueListParamsStatus = "ignored"
)

func (r ObservabilityIssueListParamsStatus) IsKnown() bool {
	switch r {
	case ObservabilityIssueListParamsStatusActive, ObservabilityIssueListParamsStatusResolved, ObservabilityIssueListParamsStatusIgnored:
		return true
	}
	return false
}

type ObservabilityIssueAddParams struct {
	AccountID param.Field[string]   `path:"account_id" api:"required"`
	IssueIDs  param.Field[[]string] `json:"issueIds" api:"required"`
}

func (r ObservabilityIssueAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityIssueAddResponseEnvelope struct {
	Errors   []ObservabilityIssueAddResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityIssueAddResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityIssueAddResponse                   `json:"result" api:"required"`
	Success  ObservabilityIssueAddResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityIssueAddResponseEnvelopeJSON       `json:"-"`
}

// observabilityIssueAddResponseEnvelopeJSON contains the JSON metadata for the
// struct [ObservabilityIssueAddResponseEnvelope]
type observabilityIssueAddResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAddResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAddResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAddResponseEnvelopeErrors struct {
	Message string                                          `json:"message" api:"required"`
	JSON    observabilityIssueAddResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityIssueAddResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ObservabilityIssueAddResponseEnvelopeErrors]
type observabilityIssueAddResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAddResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAddResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAddResponseEnvelopeMessages struct {
	Message ObservabilityIssueAddResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityIssueAddResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityIssueAddResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ObservabilityIssueAddResponseEnvelopeMessages]
type observabilityIssueAddResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueAddResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueAddResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueAddResponseEnvelopeMessagesMessage string

const (
	ObservabilityIssueAddResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityIssueAddResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityIssueAddResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityIssueAddResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityIssueAddResponseEnvelopeSuccess bool

const (
	ObservabilityIssueAddResponseEnvelopeSuccessTrue ObservabilityIssueAddResponseEnvelopeSuccess = true
)

func (r ObservabilityIssueAddResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityIssueAddResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityIssueGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ObservabilityIssueGetResponseEnvelope struct {
	Errors   []ObservabilityIssueGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityIssueGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityIssueGetResponse                   `json:"result" api:"required"`
	Success  ObservabilityIssueGetResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityIssueGetResponseEnvelopeJSON       `json:"-"`
}

// observabilityIssueGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ObservabilityIssueGetResponseEnvelope]
type observabilityIssueGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueGetResponseEnvelopeErrors struct {
	Message string                                          `json:"message" api:"required"`
	JSON    observabilityIssueGetResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityIssueGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ObservabilityIssueGetResponseEnvelopeErrors]
type observabilityIssueGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueGetResponseEnvelopeMessages struct {
	Message ObservabilityIssueGetResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityIssueGetResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityIssueGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ObservabilityIssueGetResponseEnvelopeMessages]
type observabilityIssueGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueGetResponseEnvelopeMessagesMessage string

const (
	ObservabilityIssueGetResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityIssueGetResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityIssueGetResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityIssueGetResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityIssueGetResponseEnvelopeSuccess bool

const (
	ObservabilityIssueGetResponseEnvelopeSuccessTrue ObservabilityIssueGetResponseEnvelopeSuccess = true
)

func (r ObservabilityIssueGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityIssueGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityIssueGroupParams struct {
	AccountID param.Field[string]   `path:"account_id" api:"required"`
	IssueIDs  param.Field[[]string] `json:"issueIds" api:"required"`
	Title     param.Field[string]   `json:"title" api:"required"`
}

func (r ObservabilityIssueGroupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityIssueGroupResponseEnvelope struct {
	Errors   []ObservabilityIssueGroupResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityIssueGroupResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityIssueGroupResponse                   `json:"result" api:"required"`
	Success  ObservabilityIssueGroupResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityIssueGroupResponseEnvelopeJSON       `json:"-"`
}

// observabilityIssueGroupResponseEnvelopeJSON contains the JSON metadata for the
// struct [ObservabilityIssueGroupResponseEnvelope]
type observabilityIssueGroupResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueGroupResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueGroupResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueGroupResponseEnvelopeErrors struct {
	Message string                                            `json:"message" api:"required"`
	JSON    observabilityIssueGroupResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityIssueGroupResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ObservabilityIssueGroupResponseEnvelopeErrors]
type observabilityIssueGroupResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueGroupResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueGroupResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueGroupResponseEnvelopeMessages struct {
	Message ObservabilityIssueGroupResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityIssueGroupResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityIssueGroupResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ObservabilityIssueGroupResponseEnvelopeMessages]
type observabilityIssueGroupResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueGroupResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueGroupResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueGroupResponseEnvelopeMessagesMessage string

const (
	ObservabilityIssueGroupResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityIssueGroupResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityIssueGroupResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityIssueGroupResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityIssueGroupResponseEnvelopeSuccess bool

const (
	ObservabilityIssueGroupResponseEnvelopeSuccessTrue ObservabilityIssueGroupResponseEnvelopeSuccess = true
)

func (r ObservabilityIssueGroupResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityIssueGroupResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityIssueNotifyParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	PolicyID  param.Field[string] `json:"policyId"`
}

func (r ObservabilityIssueNotifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityIssueNotifyResponseEnvelope struct {
	Errors   []ObservabilityIssueNotifyResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityIssueNotifyResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityIssueNotifyResponse                   `json:"result" api:"required"`
	Success  ObservabilityIssueNotifyResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityIssueNotifyResponseEnvelopeJSON       `json:"-"`
}

// observabilityIssueNotifyResponseEnvelopeJSON contains the JSON metadata for the
// struct [ObservabilityIssueNotifyResponseEnvelope]
type observabilityIssueNotifyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueNotifyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueNotifyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueNotifyResponseEnvelopeErrors struct {
	Message string                                             `json:"message" api:"required"`
	JSON    observabilityIssueNotifyResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityIssueNotifyResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ObservabilityIssueNotifyResponseEnvelopeErrors]
type observabilityIssueNotifyResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueNotifyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueNotifyResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueNotifyResponseEnvelopeMessages struct {
	Message ObservabilityIssueNotifyResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityIssueNotifyResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityIssueNotifyResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ObservabilityIssueNotifyResponseEnvelopeMessages]
type observabilityIssueNotifyResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueNotifyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueNotifyResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueNotifyResponseEnvelopeMessagesMessage string

const (
	ObservabilityIssueNotifyResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityIssueNotifyResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityIssueNotifyResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityIssueNotifyResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityIssueNotifyResponseEnvelopeSuccess bool

const (
	ObservabilityIssueNotifyResponseEnvelopeSuccessTrue ObservabilityIssueNotifyResponseEnvelopeSuccess = true
)

func (r ObservabilityIssueNotifyResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityIssueNotifyResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityIssueOccurrencesParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Opaque cursor returned in result_info.cursors.after.
	Cursor  param.Field[string] `query:"cursor"`
	PerPage param.Field[int64]  `query:"per_page"`
}

// URLQuery serializes [ObservabilityIssueOccurrencesParams]'s query parameters as
// `url.Values`.
func (r ObservabilityIssueOccurrencesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ObservabilityIssueOccurrencesResponseEnvelope struct {
	Errors     []ObservabilityIssueOccurrencesResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages   []ObservabilityIssueOccurrencesResponseEnvelopeMessages `json:"messages" api:"required"`
	Result     []ObservabilityIssueOccurrencesResponse                 `json:"result" api:"required"`
	ResultInfo ObservabilityIssueOccurrencesResponseEnvelopeResultInfo `json:"result_info" api:"required"`
	Success    ObservabilityIssueOccurrencesResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON       observabilityIssueOccurrencesResponseEnvelopeJSON       `json:"-"`
}

// observabilityIssueOccurrencesResponseEnvelopeJSON contains the JSON metadata for
// the struct [ObservabilityIssueOccurrencesResponseEnvelope]
type observabilityIssueOccurrencesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueOccurrencesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueOccurrencesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueOccurrencesResponseEnvelopeErrors struct {
	Message string                                                  `json:"message" api:"required"`
	JSON    observabilityIssueOccurrencesResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityIssueOccurrencesResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ObservabilityIssueOccurrencesResponseEnvelopeErrors]
type observabilityIssueOccurrencesResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueOccurrencesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueOccurrencesResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueOccurrencesResponseEnvelopeMessages struct {
	Message ObservabilityIssueOccurrencesResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityIssueOccurrencesResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityIssueOccurrencesResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilityIssueOccurrencesResponseEnvelopeMessages]
type observabilityIssueOccurrencesResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueOccurrencesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueOccurrencesResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueOccurrencesResponseEnvelopeMessagesMessage string

const (
	ObservabilityIssueOccurrencesResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityIssueOccurrencesResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityIssueOccurrencesResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityIssueOccurrencesResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityIssueOccurrencesResponseEnvelopeResultInfo struct {
	Count   int64                                                          `json:"count" api:"required"`
	Cursors ObservabilityIssueOccurrencesResponseEnvelopeResultInfoCursors `json:"cursors" api:"required"`
	PerPage int64                                                          `json:"per_page" api:"required"`
	JSON    observabilityIssueOccurrencesResponseEnvelopeResultInfoJSON    `json:"-"`
}

// observabilityIssueOccurrencesResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [ObservabilityIssueOccurrencesResponseEnvelopeResultInfo]
type observabilityIssueOccurrencesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Cursors     apijson.Field
	PerPage     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueOccurrencesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueOccurrencesResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueOccurrencesResponseEnvelopeResultInfoCursors struct {
	After string                                                             `json:"after" api:"required,nullable"`
	JSON  observabilityIssueOccurrencesResponseEnvelopeResultInfoCursorsJSON `json:"-"`
}

// observabilityIssueOccurrencesResponseEnvelopeResultInfoCursorsJSON contains the
// JSON metadata for the struct
// [ObservabilityIssueOccurrencesResponseEnvelopeResultInfoCursors]
type observabilityIssueOccurrencesResponseEnvelopeResultInfoCursorsJSON struct {
	After       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueOccurrencesResponseEnvelopeResultInfoCursors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueOccurrencesResponseEnvelopeResultInfoCursorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueOccurrencesResponseEnvelopeSuccess bool

const (
	ObservabilityIssueOccurrencesResponseEnvelopeSuccessTrue ObservabilityIssueOccurrencesResponseEnvelopeSuccess = true
)

func (r ObservabilityIssueOccurrencesResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityIssueOccurrencesResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityIssueRemoveParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ObservabilityIssueRemoveResponseEnvelope struct {
	Errors   []ObservabilityIssueRemoveResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityIssueRemoveResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityIssueRemoveResponse                   `json:"result" api:"required,nullable"`
	Success  ObservabilityIssueRemoveResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityIssueRemoveResponseEnvelopeJSON       `json:"-"`
}

// observabilityIssueRemoveResponseEnvelopeJSON contains the JSON metadata for the
// struct [ObservabilityIssueRemoveResponseEnvelope]
type observabilityIssueRemoveResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueRemoveResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueRemoveResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueRemoveResponseEnvelopeErrors struct {
	Message string                                             `json:"message" api:"required"`
	JSON    observabilityIssueRemoveResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityIssueRemoveResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ObservabilityIssueRemoveResponseEnvelopeErrors]
type observabilityIssueRemoveResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueRemoveResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueRemoveResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueRemoveResponseEnvelopeMessages struct {
	Message ObservabilityIssueRemoveResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityIssueRemoveResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityIssueRemoveResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ObservabilityIssueRemoveResponseEnvelopeMessages]
type observabilityIssueRemoveResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueRemoveResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueRemoveResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueRemoveResponseEnvelopeMessagesMessage string

const (
	ObservabilityIssueRemoveResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityIssueRemoveResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityIssueRemoveResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityIssueRemoveResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityIssueRemoveResponseEnvelopeSuccess bool

const (
	ObservabilityIssueRemoveResponseEnvelopeSuccessTrue ObservabilityIssueRemoveResponseEnvelopeSuccess = true
)

func (r ObservabilityIssueRemoveResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityIssueRemoveResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityIssueSummaryParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Service   param.Field[string] `query:"service"`
}

// URLQuery serializes [ObservabilityIssueSummaryParams]'s query parameters as
// `url.Values`.
func (r ObservabilityIssueSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ObservabilityIssueSummaryResponseEnvelope struct {
	Errors   []ObservabilityIssueSummaryResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityIssueSummaryResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityIssueSummaryResponse                   `json:"result" api:"required"`
	Success  ObservabilityIssueSummaryResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityIssueSummaryResponseEnvelopeJSON       `json:"-"`
}

// observabilityIssueSummaryResponseEnvelopeJSON contains the JSON metadata for the
// struct [ObservabilityIssueSummaryResponseEnvelope]
type observabilityIssueSummaryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueSummaryResponseEnvelopeErrors struct {
	Message string                                              `json:"message" api:"required"`
	JSON    observabilityIssueSummaryResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityIssueSummaryResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ObservabilityIssueSummaryResponseEnvelopeErrors]
type observabilityIssueSummaryResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueSummaryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueSummaryResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueSummaryResponseEnvelopeMessages struct {
	Message ObservabilityIssueSummaryResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityIssueSummaryResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityIssueSummaryResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ObservabilityIssueSummaryResponseEnvelopeMessages]
type observabilityIssueSummaryResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueSummaryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueSummaryResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueSummaryResponseEnvelopeMessagesMessage string

const (
	ObservabilityIssueSummaryResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityIssueSummaryResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityIssueSummaryResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityIssueSummaryResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityIssueSummaryResponseEnvelopeSuccess bool

const (
	ObservabilityIssueSummaryResponseEnvelopeSuccessTrue ObservabilityIssueSummaryResponseEnvelopeSuccess = true
)

func (r ObservabilityIssueSummaryResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityIssueSummaryResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityIssueUngroupParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ObservabilityIssueUngroupResponseEnvelope struct {
	Errors   []ObservabilityIssueUngroupResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityIssueUngroupResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityIssueUngroupResponse                   `json:"result" api:"required"`
	Success  ObservabilityIssueUngroupResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityIssueUngroupResponseEnvelopeJSON       `json:"-"`
}

// observabilityIssueUngroupResponseEnvelopeJSON contains the JSON metadata for the
// struct [ObservabilityIssueUngroupResponseEnvelope]
type observabilityIssueUngroupResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueUngroupResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueUngroupResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueUngroupResponseEnvelopeErrors struct {
	Message string                                              `json:"message" api:"required"`
	JSON    observabilityIssueUngroupResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityIssueUngroupResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ObservabilityIssueUngroupResponseEnvelopeErrors]
type observabilityIssueUngroupResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueUngroupResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueUngroupResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueUngroupResponseEnvelopeMessages struct {
	Message ObservabilityIssueUngroupResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityIssueUngroupResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityIssueUngroupResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ObservabilityIssueUngroupResponseEnvelopeMessages]
type observabilityIssueUngroupResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityIssueUngroupResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityIssueUngroupResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityIssueUngroupResponseEnvelopeMessagesMessage string

const (
	ObservabilityIssueUngroupResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityIssueUngroupResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityIssueUngroupResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityIssueUngroupResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityIssueUngroupResponseEnvelopeSuccess bool

const (
	ObservabilityIssueUngroupResponseEnvelopeSuccessTrue ObservabilityIssueUngroupResponseEnvelopeSuccess = true
)

func (r ObservabilityIssueUngroupResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityIssueUngroupResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

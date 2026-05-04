// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_center

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// InsightAuditLogService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightAuditLogService] method instead.
type InsightAuditLogService struct {
	Options []option.RequestOption
}

// NewInsightAuditLogService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInsightAuditLogService(opts ...option.RequestOption) (r *InsightAuditLogService) {
	r = &InsightAuditLogService{}
	r.Options = opts
	return
}

// Lists audit log entries for all Security Center insights in the account or zone,
// showing changes to insight status and classification.
func (r *InsightAuditLogService) List(ctx context.Context, params InsightAuditLogListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[InsightAuditLogListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/security-center/insights/audit-log", accountOrZone, accountOrZoneID)
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

// Lists audit log entries for all Security Center insights in the account or zone,
// showing changes to insight status and classification.
func (r *InsightAuditLogService) ListAutoPaging(ctx context.Context, params InsightAuditLogListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[InsightAuditLogListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Lists audit log entries for a specific Security Center insight, showing changes
// to its status and classification over time.
func (r *InsightAuditLogService) ListByInsight(ctx context.Context, issueID string, params InsightAuditLogListByInsightParams, opts ...option.RequestOption) (res *pagination.CursorPagination[InsightAuditLogListByInsightResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if issueID == "" {
		err = errors.New("missing required issue_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/security-center/insights/%s/audit-log", accountOrZone, accountOrZoneID, issueID)
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

// Lists audit log entries for a specific Security Center insight, showing changes
// to its status and classification over time.
func (r *InsightAuditLogService) ListByInsightAutoPaging(ctx context.Context, issueID string, params InsightAuditLogListByInsightParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[InsightAuditLogListByInsightResponse] {
	return pagination.NewCursorPaginationAutoPager(r.ListByInsight(ctx, issueID, params, opts...))
}

type InsightAuditLogListResponse struct {
	// UUIDv7 identifier for the audit log entry, time-ordered.
	ID string `json:"id" format:"uuid"`
	// The timestamp when the change occurred.
	ChangedAt time.Time `json:"changed_at" format:"date-time"`
	// The actor that made the change. 'system' for automated changes, or a user
	// identifier.
	ChangedBy string `json:"changed_by"`
	// The value of the field after the change. Null if the field was cleared.
	CurrentValue string `json:"current_value" api:"nullable"`
	// The field that was changed.
	FieldChanged InsightAuditLogListResponseFieldChanged `json:"field_changed"`
	// The ID of the insight this audit log entry relates to.
	IssueID string `json:"issue_id"`
	// The value of the field before the change. Null if the field was not previously
	// set.
	PreviousValue string `json:"previous_value" api:"nullable"`
	// Optional rationale provided for the change.
	Rationale string `json:"rationale" api:"nullable"`
	// The zone ID associated with the insight. Only present for zone-level insights.
	ZoneID int64                           `json:"zone_id"`
	JSON   insightAuditLogListResponseJSON `json:"-"`
}

// insightAuditLogListResponseJSON contains the JSON metadata for the struct
// [InsightAuditLogListResponse]
type insightAuditLogListResponseJSON struct {
	ID            apijson.Field
	ChangedAt     apijson.Field
	ChangedBy     apijson.Field
	CurrentValue  apijson.Field
	FieldChanged  apijson.Field
	IssueID       apijson.Field
	PreviousValue apijson.Field
	Rationale     apijson.Field
	ZoneID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InsightAuditLogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightAuditLogListResponseJSON) RawJSON() string {
	return r.raw
}

// The field that was changed.
type InsightAuditLogListResponseFieldChanged string

const (
	InsightAuditLogListResponseFieldChangedStatus             InsightAuditLogListResponseFieldChanged = "status"
	InsightAuditLogListResponseFieldChangedUserClassification InsightAuditLogListResponseFieldChanged = "user_classification"
)

func (r InsightAuditLogListResponseFieldChanged) IsKnown() bool {
	switch r {
	case InsightAuditLogListResponseFieldChangedStatus, InsightAuditLogListResponseFieldChangedUserClassification:
		return true
	}
	return false
}

type InsightAuditLogListByInsightResponse struct {
	// UUIDv7 identifier for the audit log entry, time-ordered.
	ID string `json:"id" format:"uuid"`
	// The timestamp when the change occurred.
	ChangedAt time.Time `json:"changed_at" format:"date-time"`
	// The actor that made the change. 'system' for automated changes, or a user
	// identifier.
	ChangedBy string `json:"changed_by"`
	// The value of the field after the change. Null if the field was cleared.
	CurrentValue string `json:"current_value" api:"nullable"`
	// The field that was changed.
	FieldChanged InsightAuditLogListByInsightResponseFieldChanged `json:"field_changed"`
	// The ID of the insight this audit log entry relates to.
	IssueID string `json:"issue_id"`
	// The value of the field before the change. Null if the field was not previously
	// set.
	PreviousValue string `json:"previous_value" api:"nullable"`
	// Optional rationale provided for the change.
	Rationale string `json:"rationale" api:"nullable"`
	// The zone ID associated with the insight. Only present for zone-level insights.
	ZoneID int64                                    `json:"zone_id"`
	JSON   insightAuditLogListByInsightResponseJSON `json:"-"`
}

// insightAuditLogListByInsightResponseJSON contains the JSON metadata for the
// struct [InsightAuditLogListByInsightResponse]
type insightAuditLogListByInsightResponseJSON struct {
	ID            apijson.Field
	ChangedAt     apijson.Field
	ChangedBy     apijson.Field
	CurrentValue  apijson.Field
	FieldChanged  apijson.Field
	IssueID       apijson.Field
	PreviousValue apijson.Field
	Rationale     apijson.Field
	ZoneID        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *InsightAuditLogListByInsightResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightAuditLogListByInsightResponseJSON) RawJSON() string {
	return r.raw
}

// The field that was changed.
type InsightAuditLogListByInsightResponseFieldChanged string

const (
	InsightAuditLogListByInsightResponseFieldChangedStatus             InsightAuditLogListByInsightResponseFieldChanged = "status"
	InsightAuditLogListByInsightResponseFieldChangedUserClassification InsightAuditLogListByInsightResponseFieldChanged = "user_classification"
)

func (r InsightAuditLogListByInsightResponseFieldChanged) IsKnown() bool {
	switch r {
	case InsightAuditLogListByInsightResponseFieldChangedStatus, InsightAuditLogListByInsightResponseFieldChangedUserClassification:
		return true
	}
	return false
}

type InsightAuditLogListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Filter entries changed before this timestamp (RFC 3339).
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Filter by the actor that made the change.
	ChangedBy param.Field[string] `query:"changed_by"`
	// Opaque cursor for pagination. Use the cursor value from result_info of the
	// previous response.
	Cursor param.Field[string] `query:"cursor"`
	// Filter by the field that was changed.
	FieldChanged param.Field[InsightAuditLogListParamsFieldChanged] `query:"field_changed"`
	// Sort order for results. Use 'asc' for oldest first or 'desc' for newest first.
	Order param.Field[InsightAuditLogListParamsOrder] `query:"order"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter entries changed at or after this timestamp (RFC 3339).
	Since param.Field[time.Time] `query:"since" format:"date-time"`
}

// URLQuery serializes [InsightAuditLogListParams]'s query parameters as
// `url.Values`.
func (r InsightAuditLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by the field that was changed.
type InsightAuditLogListParamsFieldChanged string

const (
	InsightAuditLogListParamsFieldChangedStatus             InsightAuditLogListParamsFieldChanged = "status"
	InsightAuditLogListParamsFieldChangedUserClassification InsightAuditLogListParamsFieldChanged = "user_classification"
)

func (r InsightAuditLogListParamsFieldChanged) IsKnown() bool {
	switch r {
	case InsightAuditLogListParamsFieldChangedStatus, InsightAuditLogListParamsFieldChangedUserClassification:
		return true
	}
	return false
}

// Sort order for results. Use 'asc' for oldest first or 'desc' for newest first.
type InsightAuditLogListParamsOrder string

const (
	InsightAuditLogListParamsOrderAsc  InsightAuditLogListParamsOrder = "asc"
	InsightAuditLogListParamsOrderDesc InsightAuditLogListParamsOrder = "desc"
)

func (r InsightAuditLogListParamsOrder) IsKnown() bool {
	switch r {
	case InsightAuditLogListParamsOrderAsc, InsightAuditLogListParamsOrderDesc:
		return true
	}
	return false
}

type InsightAuditLogListByInsightParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Filter entries changed before this timestamp (RFC 3339).
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Filter by the actor that made the change.
	ChangedBy param.Field[string] `query:"changed_by"`
	// Opaque cursor for pagination. Use the cursor value from result_info of the
	// previous response.
	Cursor param.Field[string] `query:"cursor"`
	// Filter by the field that was changed.
	FieldChanged param.Field[InsightAuditLogListByInsightParamsFieldChanged] `query:"field_changed"`
	// Sort order for results. Use 'asc' for oldest first or 'desc' for newest first.
	Order param.Field[InsightAuditLogListByInsightParamsOrder] `query:"order"`
	// Number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter entries changed at or after this timestamp (RFC 3339).
	Since param.Field[time.Time] `query:"since" format:"date-time"`
}

// URLQuery serializes [InsightAuditLogListByInsightParams]'s query parameters as
// `url.Values`.
func (r InsightAuditLogListByInsightParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by the field that was changed.
type InsightAuditLogListByInsightParamsFieldChanged string

const (
	InsightAuditLogListByInsightParamsFieldChangedStatus             InsightAuditLogListByInsightParamsFieldChanged = "status"
	InsightAuditLogListByInsightParamsFieldChangedUserClassification InsightAuditLogListByInsightParamsFieldChanged = "user_classification"
)

func (r InsightAuditLogListByInsightParamsFieldChanged) IsKnown() bool {
	switch r {
	case InsightAuditLogListByInsightParamsFieldChangedStatus, InsightAuditLogListByInsightParamsFieldChangedUserClassification:
		return true
	}
	return false
}

// Sort order for results. Use 'asc' for oldest first or 'desc' for newest first.
type InsightAuditLogListByInsightParamsOrder string

const (
	InsightAuditLogListByInsightParamsOrderAsc  InsightAuditLogListByInsightParamsOrder = "asc"
	InsightAuditLogListByInsightParamsOrderDesc InsightAuditLogListByInsightParamsOrder = "desc"
)

func (r InsightAuditLogListByInsightParamsOrder) IsKnown() bool {
	switch r {
	case InsightAuditLogListByInsightParamsOrderAsc, InsightAuditLogListByInsightParamsOrderDesc:
		return true
	}
	return false
}

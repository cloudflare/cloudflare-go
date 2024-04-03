// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AuditLogService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAuditLogService] method instead.
type AuditLogService struct {
	Options []option.RequestOption
}

// NewAuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuditLogService(opts ...option.RequestOption) (r *AuditLogService) {
	r = &AuditLogService{}
	r.Options = opts
	return
}

// Gets a list of audit logs for a user account. Can be filtered by who made the
// change, on which zone, and the timeframe of the change.
func (r *AuditLogService) List(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AuditLogListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/audit_logs"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Gets a list of audit logs for a user account. Can be filtered by who made the
// change, on which zone, and the timeframe of the change.
func (r *AuditLogService) ListAutoPaging(ctx context.Context, query AuditLogListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AuditLogListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

type AuditLogListResponse struct {
	// A string that uniquely identifies the audit log.
	ID     string                     `json:"id"`
	Action AuditLogListResponseAction `json:"action"`
	Actor  AuditLogListResponseActor  `json:"actor"`
	// The source of the event.
	Interface string `json:"interface"`
	// An object which can lend more context to the action being logged. This is a
	// flexible value and varies between different actions.
	Metadata interface{} `json:"metadata"`
	// The new value of the resource that was modified.
	NewValue string `json:"newValue"`
	// The value of the resource before it was modified.
	OldValue string                       `json:"oldValue"`
	Owner    AuditLogListResponseOwner    `json:"owner"`
	Resource AuditLogListResponseResource `json:"resource"`
	// A UTC RFC3339 timestamp that specifies when the action being logged occured.
	When time.Time                `json:"when" format:"date-time"`
	JSON auditLogListResponseJSON `json:"-"`
}

// auditLogListResponseJSON contains the JSON metadata for the struct
// [AuditLogListResponse]
type auditLogListResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Actor       apijson.Field
	Interface   apijson.Field
	Metadata    apijson.Field
	NewValue    apijson.Field
	OldValue    apijson.Field
	Owner       apijson.Field
	Resource    apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogListResponseJSON) RawJSON() string {
	return r.raw
}

type AuditLogListResponseAction struct {
	// A boolean that indicates if the action attempted was successful.
	Result bool `json:"result"`
	// A short string that describes the action that was performed.
	Type string                         `json:"type"`
	JSON auditLogListResponseActionJSON `json:"-"`
}

// auditLogListResponseActionJSON contains the JSON metadata for the struct
// [AuditLogListResponseAction]
type auditLogListResponseActionJSON struct {
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogListResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogListResponseActionJSON) RawJSON() string {
	return r.raw
}

type AuditLogListResponseActor struct {
	// The ID of the actor that performed the action. If a user performed the action,
	// this will be their User ID.
	ID string `json:"id"`
	// The email of the user that performed the action.
	Email string `json:"email" format:"email"`
	// The IP address of the request that performed the action.
	IP string `json:"ip"`
	// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
	Type AuditLogListResponseActorType `json:"type"`
	JSON auditLogListResponseActorJSON `json:"-"`
}

// auditLogListResponseActorJSON contains the JSON metadata for the struct
// [AuditLogListResponseActor]
type auditLogListResponseActorJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	IP          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogListResponseActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogListResponseActorJSON) RawJSON() string {
	return r.raw
}

// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
type AuditLogListResponseActorType string

const (
	AuditLogListResponseActorTypeUser       AuditLogListResponseActorType = "user"
	AuditLogListResponseActorTypeAdmin      AuditLogListResponseActorType = "admin"
	AuditLogListResponseActorTypeCloudflare AuditLogListResponseActorType = "Cloudflare"
)

func (r AuditLogListResponseActorType) IsKnown() bool {
	switch r {
	case AuditLogListResponseActorTypeUser, AuditLogListResponseActorTypeAdmin, AuditLogListResponseActorTypeCloudflare:
		return true
	}
	return false
}

type AuditLogListResponseOwner struct {
	// Identifier
	ID   string                        `json:"id"`
	JSON auditLogListResponseOwnerJSON `json:"-"`
}

// auditLogListResponseOwnerJSON contains the JSON metadata for the struct
// [AuditLogListResponseOwner]
type auditLogListResponseOwnerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogListResponseOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogListResponseOwnerJSON) RawJSON() string {
	return r.raw
}

type AuditLogListResponseResource struct {
	// An identifier for the resource that was affected by the action.
	ID string `json:"id"`
	// A short string that describes the resource that was affected by the action.
	Type string                           `json:"type"`
	JSON auditLogListResponseResourceJSON `json:"-"`
}

// auditLogListResponseResourceJSON contains the JSON metadata for the struct
// [AuditLogListResponseResource]
type auditLogListResponseResourceJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogListResponseResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogListResponseResourceJSON) RawJSON() string {
	return r.raw
}

type AuditLogListParams struct {
	// Finds a specific log by its ID.
	ID     param.Field[string]                   `query:"id"`
	Action param.Field[AuditLogListParamsAction] `query:"action"`
	Actor  param.Field[AuditLogListParamsActor]  `query:"actor"`
	// Limits the returned results to logs older than the specified date. This can be a
	// date string `2019-04-30` or an absolute timestamp that conforms to RFC3339.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Changes the direction of the chronological sorting.
	Direction param.Field[AuditLogListParamsDirection] `query:"direction"`
	// Indicates that this request is an export of logs in CSV format.
	Export param.Field[bool] `query:"export"`
	// Indicates whether or not to hide user level audit logs.
	HideUserLogs param.Field[bool] `query:"hide_user_logs"`
	// Defines which page of results to return.
	Page param.Field[float64] `query:"page"`
	// Sets the number of results to return per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Limits the returned results to logs newer than the specified date. This can be a
	// date string `2019-04-30` or an absolute timestamp that conforms to RFC3339.
	Since param.Field[time.Time]              `query:"since" format:"date-time"`
	Zone  param.Field[AuditLogListParamsZone] `query:"zone"`
}

// URLQuery serializes [AuditLogListParams]'s query parameters as `url.Values`.
func (r AuditLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AuditLogListParamsAction struct {
	// Filters by the action type.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [AuditLogListParamsAction]'s query parameters as
// `url.Values`.
func (r AuditLogListParamsAction) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AuditLogListParamsActor struct {
	// Filters by the email address of the actor that made the change.
	Email param.Field[string] `query:"email" format:"email"`
	// Filters by the IP address of the request that made the change by specific IP
	// address or valid CIDR Range.
	IP param.Field[string] `query:"ip"`
}

// URLQuery serializes [AuditLogListParamsActor]'s query parameters as
// `url.Values`.
func (r AuditLogListParamsActor) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Changes the direction of the chronological sorting.
type AuditLogListParamsDirection string

const (
	AuditLogListParamsDirectionDesc AuditLogListParamsDirection = "desc"
	AuditLogListParamsDirectionAsc  AuditLogListParamsDirection = "asc"
)

func (r AuditLogListParamsDirection) IsKnown() bool {
	switch r {
	case AuditLogListParamsDirectionDesc, AuditLogListParamsDirectionAsc:
		return true
	}
	return false
}

type AuditLogListParamsZone struct {
	// Filters by the name of the zone associated to the change.
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [AuditLogListParamsZone]'s query parameters as `url.Values`.
func (r AuditLogListParamsZone) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

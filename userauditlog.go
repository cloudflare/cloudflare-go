// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserAuditLogService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserAuditLogService] method
// instead.
type UserAuditLogService struct {
	Options []option.RequestOption
}

// NewUserAuditLogService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserAuditLogService(opts ...option.RequestOption) (r *UserAuditLogService) {
	r = &UserAuditLogService{}
	r.Options = opts
	return
}

// Gets a list of audit logs for a user account. Can be filtered by who made the
// change, on which zone, and the timeframe of the change.
func (r *UserAuditLogService) List(ctx context.Context, query UserAuditLogListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[UserAuditLogListResponse], err error) {
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
func (r *UserAuditLogService) ListAutoPaging(ctx context.Context, query UserAuditLogListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[UserAuditLogListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

type UserAuditLogListResponse struct {
	// A string that uniquely identifies the audit log.
	ID     string                         `json:"id"`
	Action UserAuditLogListResponseAction `json:"action"`
	Actor  UserAuditLogListResponseActor  `json:"actor"`
	// The source of the event.
	Interface string `json:"interface"`
	// An object which can lend more context to the action being logged. This is a
	// flexible value and varies between different actions.
	Metadata interface{} `json:"metadata"`
	// The new value of the resource that was modified.
	NewValue string `json:"newValue"`
	// The value of the resource before it was modified.
	OldValue string                           `json:"oldValue"`
	Owner    UserAuditLogListResponseOwner    `json:"owner"`
	Resource UserAuditLogListResponseResource `json:"resource"`
	// A UTC RFC3339 timestamp that specifies when the action being logged occured.
	When time.Time                    `json:"when" format:"date-time"`
	JSON userAuditLogListResponseJSON `json:"-"`
}

// userAuditLogListResponseJSON contains the JSON metadata for the struct
// [UserAuditLogListResponse]
type userAuditLogListResponseJSON struct {
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

func (r *UserAuditLogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAuditLogListResponseJSON) RawJSON() string {
	return r.raw
}

type UserAuditLogListResponseAction struct {
	// A boolean that indicates if the action attempted was successful.
	Result bool `json:"result"`
	// A short string that describes the action that was performed.
	Type string                             `json:"type"`
	JSON userAuditLogListResponseActionJSON `json:"-"`
}

// userAuditLogListResponseActionJSON contains the JSON metadata for the struct
// [UserAuditLogListResponseAction]
type userAuditLogListResponseActionJSON struct {
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogListResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAuditLogListResponseActionJSON) RawJSON() string {
	return r.raw
}

type UserAuditLogListResponseActor struct {
	// The ID of the actor that performed the action. If a user performed the action,
	// this will be their User ID.
	ID string `json:"id"`
	// The email of the user that performed the action.
	Email string `json:"email" format:"email"`
	// The IP address of the request that performed the action.
	IP string `json:"ip"`
	// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
	Type UserAuditLogListResponseActorType `json:"type"`
	JSON userAuditLogListResponseActorJSON `json:"-"`
}

// userAuditLogListResponseActorJSON contains the JSON metadata for the struct
// [UserAuditLogListResponseActor]
type userAuditLogListResponseActorJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	IP          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogListResponseActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAuditLogListResponseActorJSON) RawJSON() string {
	return r.raw
}

// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
type UserAuditLogListResponseActorType string

const (
	UserAuditLogListResponseActorTypeUser       UserAuditLogListResponseActorType = "user"
	UserAuditLogListResponseActorTypeAdmin      UserAuditLogListResponseActorType = "admin"
	UserAuditLogListResponseActorTypeCloudflare UserAuditLogListResponseActorType = "Cloudflare"
)

type UserAuditLogListResponseOwner struct {
	// Identifier
	ID   string                            `json:"id"`
	JSON userAuditLogListResponseOwnerJSON `json:"-"`
}

// userAuditLogListResponseOwnerJSON contains the JSON metadata for the struct
// [UserAuditLogListResponseOwner]
type userAuditLogListResponseOwnerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogListResponseOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAuditLogListResponseOwnerJSON) RawJSON() string {
	return r.raw
}

type UserAuditLogListResponseResource struct {
	// An identifier for the resource that was affected by the action.
	ID string `json:"id"`
	// A short string that describes the resource that was affected by the action.
	Type string                               `json:"type"`
	JSON userAuditLogListResponseResourceJSON `json:"-"`
}

// userAuditLogListResponseResourceJSON contains the JSON metadata for the struct
// [UserAuditLogListResponseResource]
type userAuditLogListResponseResourceJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAuditLogListResponseResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAuditLogListResponseResourceJSON) RawJSON() string {
	return r.raw
}

type UserAuditLogListParams struct {
	// Finds a specific log by its ID.
	ID     param.Field[string]                       `query:"id"`
	Action param.Field[UserAuditLogListParamsAction] `query:"action"`
	Actor  param.Field[UserAuditLogListParamsActor]  `query:"actor"`
	// Limits the returned results to logs older than the specified date. This can be a
	// date string `2019-04-30` or an absolute timestamp that conforms to RFC3339.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Changes the direction of the chronological sorting.
	Direction param.Field[UserAuditLogListParamsDirection] `query:"direction"`
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
	Since param.Field[time.Time]                  `query:"since" format:"date-time"`
	Zone  param.Field[UserAuditLogListParamsZone] `query:"zone"`
}

// URLQuery serializes [UserAuditLogListParams]'s query parameters as `url.Values`.
func (r UserAuditLogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserAuditLogListParamsAction struct {
	// Filters by the action type.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [UserAuditLogListParamsAction]'s query parameters as
// `url.Values`.
func (r UserAuditLogListParamsAction) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserAuditLogListParamsActor struct {
	// Filters by the email address of the actor that made the change.
	Email param.Field[string] `query:"email" format:"email"`
	// Filters by the IP address of the request that made the change by specific IP
	// address or valid CIDR Range.
	IP param.Field[string] `query:"ip"`
}

// URLQuery serializes [UserAuditLogListParamsActor]'s query parameters as
// `url.Values`.
func (r UserAuditLogListParamsActor) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Changes the direction of the chronological sorting.
type UserAuditLogListParamsDirection string

const (
	UserAuditLogListParamsDirectionDesc UserAuditLogListParamsDirection = "desc"
	UserAuditLogListParamsDirectionAsc  UserAuditLogListParamsDirection = "asc"
)

type UserAuditLogListParamsZone struct {
	// Filters by the name of the zone associated to the change.
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [UserAuditLogListParamsZone]'s query parameters as
// `url.Values`.
func (r UserAuditLogListParamsZone) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

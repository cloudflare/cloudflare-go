// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
)

// LogAuditService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogAuditService] method instead.
type LogAuditService struct {
	Options []option.RequestOption
}

// NewLogAuditService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLogAuditService(opts ...option.RequestOption) (r *LogAuditService) {
	r = &LogAuditService{}
	r.Options = opts
	return
}

// Gets a list of audit logs for an account. <br /> <br /> This is the beta release
// of Audit Logs Version 2. Since this is a beta version, there may be gaps or
// missing entries in the available audit logs. Be aware of the following
// limitations. <br /> <ul> <li>Audit logs are available only for the past 30 days.
// <br /></li> <li>Error handling is not yet implemented. <br /> </li> </ul>
func (r *LogAuditService) List(ctx context.Context, params LogAuditListParams, opts ...option.RequestOption) (res *pagination.CursorLimitPagination[LogAuditListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logs/audit", params.AccountID)
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

// Gets a list of audit logs for an account. <br /> <br /> This is the beta release
// of Audit Logs Version 2. Since this is a beta version, there may be gaps or
// missing entries in the available audit logs. Be aware of the following
// limitations. <br /> <ul> <li>Audit logs are available only for the past 30 days.
// <br /></li> <li>Error handling is not yet implemented. <br /> </li> </ul>
func (r *LogAuditService) ListAutoPaging(ctx context.Context, params LogAuditListParams, opts ...option.RequestOption) *pagination.CursorLimitPaginationAutoPager[LogAuditListResponse] {
	return pagination.NewCursorLimitPaginationAutoPager(r.List(ctx, params, opts...))
}

type LogAuditListResponse struct {
	// A unique identifier for the audit log entry.
	ID string `json:"id"`
	// Contains account related information.
	Account LogAuditListResponseAccount `json:"account"`
	// Provides information about the action performed.
	Action LogAuditListResponseAction `json:"action"`
	// Provides details about the actor who performed the action.
	Actor LogAuditListResponseActor `json:"actor"`
	// Provides raw information about the request and response.
	Raw LogAuditListResponseRaw `json:"raw"`
	// Provides details about the affected resource.
	Resource LogAuditListResponseResource `json:"resource"`
	// Provides details about the zone affected by the action.
	Zone LogAuditListResponseZone `json:"zone"`
	JSON logAuditListResponseJSON `json:"-"`
}

// logAuditListResponseJSON contains the JSON metadata for the struct
// [LogAuditListResponse]
type logAuditListResponseJSON struct {
	ID          apijson.Field
	Account     apijson.Field
	Action      apijson.Field
	Actor       apijson.Field
	Raw         apijson.Field
	Resource    apijson.Field
	Zone        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditListResponseJSON) RawJSON() string {
	return r.raw
}

// Contains account related information.
type LogAuditListResponseAccount struct {
	// A unique identifier for the account.
	ID string `json:"id"`
	// A string that identifies the account name.
	Name string                          `json:"name"`
	JSON logAuditListResponseAccountJSON `json:"-"`
}

// logAuditListResponseAccountJSON contains the JSON metadata for the struct
// [LogAuditListResponseAccount]
type logAuditListResponseAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditListResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditListResponseAccountJSON) RawJSON() string {
	return r.raw
}

// Provides information about the action performed.
type LogAuditListResponseAction struct {
	// A short description of the action performed.
	Description string `json:"description"`
	// The result of the action, indicating success or failure.
	Result string `json:"result"`
	// A timestamp indicating when the action was logged.
	Time time.Time `json:"time" format:"date-time"`
	// A short string that describes the action that was performed.
	Type string                         `json:"type"`
	JSON logAuditListResponseActionJSON `json:"-"`
}

// logAuditListResponseActionJSON contains the JSON metadata for the struct
// [LogAuditListResponseAction]
type logAuditListResponseActionJSON struct {
	Description apijson.Field
	Result      apijson.Field
	Time        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditListResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditListResponseActionJSON) RawJSON() string {
	return r.raw
}

// Provides details about the actor who performed the action.
type LogAuditListResponseActor struct {
	// The ID of the actor who performed the action. If a user performed the action,
	// this will be their User ID.
	ID      string                           `json:"id"`
	Context LogAuditListResponseActorContext `json:"context"`
	// The email of the actor who performed the action.
	Email string `json:"email" format:"email"`
	// The IP address of the request that performed the action.
	IPAddress string `json:"ip_address" format:"ipv4 | ipv6"`
	// Filters by the API token ID when the actor context is an api_token.
	TokenID string `json:"token_id"`
	// Filters by the API token name when the actor context is an api_token.
	TokenName string `json:"token_name"`
	// The type of actor.
	Type LogAuditListResponseActorType `json:"type"`
	JSON logAuditListResponseActorJSON `json:"-"`
}

// logAuditListResponseActorJSON contains the JSON metadata for the struct
// [LogAuditListResponseActor]
type logAuditListResponseActorJSON struct {
	ID          apijson.Field
	Context     apijson.Field
	Email       apijson.Field
	IPAddress   apijson.Field
	TokenID     apijson.Field
	TokenName   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditListResponseActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditListResponseActorJSON) RawJSON() string {
	return r.raw
}

type LogAuditListResponseActorContext string

const (
	LogAuditListResponseActorContextAPIKey      LogAuditListResponseActorContext = "api_key"
	LogAuditListResponseActorContextAPIToken    LogAuditListResponseActorContext = "api_token"
	LogAuditListResponseActorContextDash        LogAuditListResponseActorContext = "dash"
	LogAuditListResponseActorContextOAuth       LogAuditListResponseActorContext = "oauth"
	LogAuditListResponseActorContextOriginCAKey LogAuditListResponseActorContext = "origin_ca_key"
)

func (r LogAuditListResponseActorContext) IsKnown() bool {
	switch r {
	case LogAuditListResponseActorContextAPIKey, LogAuditListResponseActorContextAPIToken, LogAuditListResponseActorContextDash, LogAuditListResponseActorContextOAuth, LogAuditListResponseActorContextOriginCAKey:
		return true
	}
	return false
}

// The type of actor.
type LogAuditListResponseActorType string

const (
	LogAuditListResponseActorTypeUser            LogAuditListResponseActorType = "user"
	LogAuditListResponseActorTypeAccount         LogAuditListResponseActorType = "account"
	LogAuditListResponseActorTypeCloudflareAdmin LogAuditListResponseActorType = "cloudflare-admin"
)

func (r LogAuditListResponseActorType) IsKnown() bool {
	switch r {
	case LogAuditListResponseActorTypeUser, LogAuditListResponseActorTypeAccount, LogAuditListResponseActorTypeCloudflareAdmin:
		return true
	}
	return false
}

// Provides raw information about the request and response.
type LogAuditListResponseRaw struct {
	// The Cloudflare Ray ID for the request.
	CfRayID string `json:"cf_ray_id"`
	// The HTTP method of the request.
	Method string `json:"method"`
	// The HTTP response status code returned by the API.
	StatusCode int64 `json:"status_code"`
	// The URI of the request.
	URI string `json:"uri"`
	// The client's user agent string sent with the request.
	UserAgent string                      `json:"user_agent"`
	JSON      logAuditListResponseRawJSON `json:"-"`
}

// logAuditListResponseRawJSON contains the JSON metadata for the struct
// [LogAuditListResponseRaw]
type logAuditListResponseRawJSON struct {
	CfRayID     apijson.Field
	Method      apijson.Field
	StatusCode  apijson.Field
	URI         apijson.Field
	UserAgent   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditListResponseRaw) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditListResponseRawJSON) RawJSON() string {
	return r.raw
}

// Provides details about the affected resource.
type LogAuditListResponseResource struct {
	// The unique identifier for the affected resource.
	ID string `json:"id"`
	// The Cloudflare product associated with the resource.
	Product  string      `json:"product"`
	Request  interface{} `json:"request"`
	Response interface{} `json:"response"`
	// The scope of the resource.
	Scope interface{} `json:"scope"`
	// The type of the resource.
	Type string                           `json:"type"`
	JSON logAuditListResponseResourceJSON `json:"-"`
}

// logAuditListResponseResourceJSON contains the JSON metadata for the struct
// [LogAuditListResponseResource]
type logAuditListResponseResourceJSON struct {
	ID          apijson.Field
	Product     apijson.Field
	Request     apijson.Field
	Response    apijson.Field
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditListResponseResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditListResponseResourceJSON) RawJSON() string {
	return r.raw
}

// Provides details about the zone affected by the action.
type LogAuditListResponseZone struct {
	// A string that identifies the zone id.
	ID string `json:"id"`
	// A string that identifies the zone name.
	Name string                       `json:"name"`
	JSON logAuditListResponseZoneJSON `json:"-"`
}

// logAuditListResponseZoneJSON contains the JSON metadata for the struct
// [LogAuditListResponseZone]
type logAuditListResponseZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditListResponseZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditListResponseZoneJSON) RawJSON() string {
	return r.raw
}

type LogAuditListParams struct {
	// The unique id that identifies the account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Filters actions based on a given timestamp in the format yyyy-mm-dd, returning
	// only logs that occurred on and before the specified date.
	Before param.Field[time.Time] `query:"before,required" format:"date"`
	// Filters actions based on a given timestamp in the format yyyy-mm-dd, returning
	// only logs that occurred on and after the specified date.
	Since param.Field[time.Time] `query:"since,required" format:"date"`
	// Filters by the account name.
	AccountName param.Field[string] `query:"account_name"`
	// Whether the action was successful or not.
	ActionResult param.Field[LogAuditListParamsActionResult] `query:"action_result"`
	// Filters by the action type.
	ActionType param.Field[LogAuditListParamsActionType] `query:"action_type"`
	// Filters by the actor context.
	ActorContext param.Field[LogAuditListParamsActorContext] `query:"actor_context"`
	// Filters by the actor's email address.
	ActorEmail param.Field[string] `query:"actor_email" format:"email"`
	// Filters by the actor ID. This can be either the Account ID or User ID.
	ActorID param.Field[string] `query:"actor_id"`
	// The IP address where the action was initiated.
	ActorIPAddress param.Field[string] `query:"actor_ip_address"`
	// Filters by the API token ID when the actor context is an api_token or oauth.
	ActorTokenID param.Field[string] `query:"actor_token_id"`
	// Filters by the API token name when the actor context is an api_token or oauth.
	ActorTokenName param.Field[string] `query:"actor_token_name"`
	// Filters by the actor type.
	ActorType param.Field[LogAuditListParamsActorType] `query:"actor_type"`
	// Finds a specific log by its ID.
	AuditLogID param.Field[string] `query:"audit_log_id"`
	// The cursor is an opaque token used to paginate through large sets of records. It
	// indicates the position from which to continue when requesting the next set of
	// records. A valid cursor value can be obtained from the cursor object in the
	// result_info structure of a previous response.
	Cursor param.Field[string] `query:"cursor"`
	// Sets sorting order.
	Direction param.Field[LogAuditListParamsDirection] `query:"direction"`
	// The number limits the objects to return. The cursor attribute may be used to
	// iterate over the next batch of objects if there are more than the limit.
	Limit param.Field[float64] `query:"limit"`
	// Filters by the response CF Ray ID.
	RawCfRayID param.Field[string] `query:"raw_cf_ray_id"`
	// The HTTP method for the API call.
	RawMethod param.Field[string] `query:"raw_method"`
	// The response status code that was returned.
	RawStatusCode param.Field[int64] `query:"raw_status_code"`
	// Filters by the request URI.
	RawURI param.Field[string] `query:"raw_uri"`
	// Filters by the resource ID.
	ResourceID param.Field[string] `query:"resource_id"`
	// Filters audit logs by the Cloudflare product associated with the changed
	// resource.
	ResourceProduct param.Field[string] `query:"resource_product"`
	// Filters by the resource scope, specifying whether the resource is associated
	// with an user, an account, or a zone.
	ResourceScope param.Field[LogAuditListParamsResourceScope] `query:"resource_scope"`
	// Filters audit logs based on the unique type of resource changed by the action.
	ResourceType param.Field[string] `query:"resource_type"`
	// Filters by the zone ID.
	ZoneID param.Field[string] `query:"zone_id"`
	// Filters by the zone name associated with the change.
	ZoneName param.Field[string] `query:"zone_name"`
}

// URLQuery serializes [LogAuditListParams]'s query parameters as `url.Values`.
func (r LogAuditListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Whether the action was successful or not.
type LogAuditListParamsActionResult string

const (
	LogAuditListParamsActionResultSuccess LogAuditListParamsActionResult = "success"
	LogAuditListParamsActionResultFailure LogAuditListParamsActionResult = "failure"
)

func (r LogAuditListParamsActionResult) IsKnown() bool {
	switch r {
	case LogAuditListParamsActionResultSuccess, LogAuditListParamsActionResultFailure:
		return true
	}
	return false
}

// Filters by the action type.
type LogAuditListParamsActionType string

const (
	LogAuditListParamsActionTypeCreate LogAuditListParamsActionType = "create"
	LogAuditListParamsActionTypeDelete LogAuditListParamsActionType = "delete"
	LogAuditListParamsActionTypeView   LogAuditListParamsActionType = "view"
	LogAuditListParamsActionTypeUpdate LogAuditListParamsActionType = "update"
)

func (r LogAuditListParamsActionType) IsKnown() bool {
	switch r {
	case LogAuditListParamsActionTypeCreate, LogAuditListParamsActionTypeDelete, LogAuditListParamsActionTypeView, LogAuditListParamsActionTypeUpdate:
		return true
	}
	return false
}

// Filters by the actor context.
type LogAuditListParamsActorContext string

const (
	LogAuditListParamsActorContextAPIKey      LogAuditListParamsActorContext = "api_key"
	LogAuditListParamsActorContextAPIToken    LogAuditListParamsActorContext = "api_token"
	LogAuditListParamsActorContextDash        LogAuditListParamsActorContext = "dash"
	LogAuditListParamsActorContextOAuth       LogAuditListParamsActorContext = "oauth"
	LogAuditListParamsActorContextOriginCAKey LogAuditListParamsActorContext = "origin_ca_key"
)

func (r LogAuditListParamsActorContext) IsKnown() bool {
	switch r {
	case LogAuditListParamsActorContextAPIKey, LogAuditListParamsActorContextAPIToken, LogAuditListParamsActorContextDash, LogAuditListParamsActorContextOAuth, LogAuditListParamsActorContextOriginCAKey:
		return true
	}
	return false
}

// Filters by the actor type.
type LogAuditListParamsActorType string

const (
	LogAuditListParamsActorTypeCloudflareAdmin LogAuditListParamsActorType = "cloudflare_admin"
	LogAuditListParamsActorTypeAccount         LogAuditListParamsActorType = "account"
	LogAuditListParamsActorTypeUser            LogAuditListParamsActorType = "user"
)

func (r LogAuditListParamsActorType) IsKnown() bool {
	switch r {
	case LogAuditListParamsActorTypeCloudflareAdmin, LogAuditListParamsActorTypeAccount, LogAuditListParamsActorTypeUser:
		return true
	}
	return false
}

// Sets sorting order.
type LogAuditListParamsDirection string

const (
	LogAuditListParamsDirectionDesc LogAuditListParamsDirection = "desc"
	LogAuditListParamsDirectionAsc  LogAuditListParamsDirection = "asc"
)

func (r LogAuditListParamsDirection) IsKnown() bool {
	switch r {
	case LogAuditListParamsDirectionDesc, LogAuditListParamsDirectionAsc:
		return true
	}
	return false
}

// Filters by the resource scope, specifying whether the resource is associated
// with an user, an account, or a zone.
type LogAuditListParamsResourceScope string

const (
	LogAuditListParamsResourceScopeAccounts LogAuditListParamsResourceScope = "accounts"
	LogAuditListParamsResourceScopeUser     LogAuditListParamsResourceScope = "user"
	LogAuditListParamsResourceScopeZones    LogAuditListParamsResourceScope = "zones"
)

func (r LogAuditListParamsResourceScope) IsKnown() bool {
	switch r {
	case LogAuditListParamsResourceScopeAccounts, LogAuditListParamsResourceScopeUser, LogAuditListParamsResourceScopeZones:
		return true
	}
	return false
}

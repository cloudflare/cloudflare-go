// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
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

// Gets a list of audit logs for an account.
func (r *LogAuditService) List(ctx context.Context, params LogAuditListParams, opts ...option.RequestOption) (res *pagination.CursorPaginationAfter[LogAuditListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
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

// Gets a list of audit logs for an account.
func (r *LogAuditService) ListAutoPaging(ctx context.Context, params LogAuditListParams, opts ...option.RequestOption) *pagination.CursorPaginationAfterAutoPager[LogAuditListResponse] {
	return pagination.NewCursorPaginationAfterAutoPager(r.List(ctx, params, opts...))
}

// Returns the chronological change history for the resource identified by the
// given audit log entry.
//
// The endpoint first locates the source audit log entry by `id` (using
// `action_time` to narrow the lookup window), derives identifying filters from
// that entry, and then returns matching audit logs within the `since`/`before`
// window.
//
// The `result_info.history_status` field indicates the quality of the resource
// identification used:
//
//   - `exact`: Resource was identified by the resource URI.
//   - `approximate`: Resource was identified without the resource URI.
//   - `unavailable`: The source audit log entry did not contain enough information
//     to identify the resource; an empty result is returned.
func (r *LogAuditService) History(ctx context.Context, id string, params LogAuditHistoryParams, opts ...option.RequestOption) (res *[]LogAuditHistoryResponse, err error) {
	var env LogAuditHistoryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/logs/audit/%s/history", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists the available audit log product categories and the resource products each
// one expands to. Use these values with the product_category filter on the account
// audit logs endpoint.
func (r *LogAuditService) ProductCategories(ctx context.Context, query LogAuditProductCategoriesParams, opts ...option.RequestOption) (res *pagination.SinglePage[LogAuditProductCategoriesResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/logs/audit/product_categories", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists the available audit log product categories and the resource products each
// one expands to. Use these values with the product_category filter on the account
// audit logs endpoint.
func (r *LogAuditService) ProductCategoriesAutoPaging(ctx context.Context, query LogAuditProductCategoriesParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[LogAuditProductCategoriesResponse] {
	return pagination.NewSinglePageAutoPager(r.ProductCategories(ctx, query, opts...))
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
	IPAddress string `json:"ip_address"`
	// The API token ID when the actor context is an api_token or oauth.
	TokenID string `json:"token_id"`
	// The API token name when the actor context is an api_token or oauth.
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
	LogAuditListResponseActorTypeAccount         LogAuditListResponseActorType = "account"
	LogAuditListResponseActorTypeCloudflareAdmin LogAuditListResponseActorType = "cloudflare_admin"
	LogAuditListResponseActorTypeSystem          LogAuditListResponseActorType = "system"
	LogAuditListResponseActorTypeUser            LogAuditListResponseActorType = "user"
)

func (r LogAuditListResponseActorType) IsKnown() bool {
	switch r {
	case LogAuditListResponseActorTypeAccount, LogAuditListResponseActorTypeCloudflareAdmin, LogAuditListResponseActorTypeSystem, LogAuditListResponseActorTypeUser:
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

type LogAuditHistoryResponse struct {
	// A unique identifier for the audit log entry.
	ID string `json:"id"`
	// Contains account related information.
	Account LogAuditHistoryResponseAccount `json:"account"`
	// Provides information about the action performed.
	Action LogAuditHistoryResponseAction `json:"action"`
	// Provides details about the actor who performed the action.
	Actor LogAuditHistoryResponseActor `json:"actor"`
	// Provides raw information about the request and response.
	Raw LogAuditHistoryResponseRaw `json:"raw"`
	// Provides details about the affected resource.
	Resource LogAuditHistoryResponseResource `json:"resource"`
	// Provides details about the zone affected by the action.
	Zone LogAuditHistoryResponseZone `json:"zone"`
	JSON logAuditHistoryResponseJSON `json:"-"`
}

// logAuditHistoryResponseJSON contains the JSON metadata for the struct
// [LogAuditHistoryResponse]
type logAuditHistoryResponseJSON struct {
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

func (r *LogAuditHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditHistoryResponseJSON) RawJSON() string {
	return r.raw
}

// Contains account related information.
type LogAuditHistoryResponseAccount struct {
	// A unique identifier for the account.
	ID string `json:"id"`
	// A string that identifies the account name.
	Name string                             `json:"name"`
	JSON logAuditHistoryResponseAccountJSON `json:"-"`
}

// logAuditHistoryResponseAccountJSON contains the JSON metadata for the struct
// [LogAuditHistoryResponseAccount]
type logAuditHistoryResponseAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditHistoryResponseAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditHistoryResponseAccountJSON) RawJSON() string {
	return r.raw
}

// Provides information about the action performed.
type LogAuditHistoryResponseAction struct {
	// A short description of the action performed.
	Description string `json:"description"`
	// The result of the action, indicating success or failure.
	Result string `json:"result"`
	// A timestamp indicating when the action was logged.
	Time time.Time `json:"time" format:"date-time"`
	// A short string that describes the action that was performed.
	Type string                            `json:"type"`
	JSON logAuditHistoryResponseActionJSON `json:"-"`
}

// logAuditHistoryResponseActionJSON contains the JSON metadata for the struct
// [LogAuditHistoryResponseAction]
type logAuditHistoryResponseActionJSON struct {
	Description apijson.Field
	Result      apijson.Field
	Time        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditHistoryResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditHistoryResponseActionJSON) RawJSON() string {
	return r.raw
}

// Provides details about the actor who performed the action.
type LogAuditHistoryResponseActor struct {
	// The ID of the actor who performed the action. If a user performed the action,
	// this will be their User ID.
	ID      string                              `json:"id"`
	Context LogAuditHistoryResponseActorContext `json:"context"`
	// The email of the actor who performed the action.
	Email string `json:"email" format:"email"`
	// The IP address of the request that performed the action.
	IPAddress string `json:"ip_address"`
	// The API token ID when the actor context is an api_token or oauth.
	TokenID string `json:"token_id"`
	// The API token name when the actor context is an api_token or oauth.
	TokenName string `json:"token_name"`
	// The type of actor.
	Type LogAuditHistoryResponseActorType `json:"type"`
	JSON logAuditHistoryResponseActorJSON `json:"-"`
}

// logAuditHistoryResponseActorJSON contains the JSON metadata for the struct
// [LogAuditHistoryResponseActor]
type logAuditHistoryResponseActorJSON struct {
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

func (r *LogAuditHistoryResponseActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditHistoryResponseActorJSON) RawJSON() string {
	return r.raw
}

type LogAuditHistoryResponseActorContext string

const (
	LogAuditHistoryResponseActorContextAPIKey      LogAuditHistoryResponseActorContext = "api_key"
	LogAuditHistoryResponseActorContextAPIToken    LogAuditHistoryResponseActorContext = "api_token"
	LogAuditHistoryResponseActorContextDash        LogAuditHistoryResponseActorContext = "dash"
	LogAuditHistoryResponseActorContextOAuth       LogAuditHistoryResponseActorContext = "oauth"
	LogAuditHistoryResponseActorContextOriginCAKey LogAuditHistoryResponseActorContext = "origin_ca_key"
)

func (r LogAuditHistoryResponseActorContext) IsKnown() bool {
	switch r {
	case LogAuditHistoryResponseActorContextAPIKey, LogAuditHistoryResponseActorContextAPIToken, LogAuditHistoryResponseActorContextDash, LogAuditHistoryResponseActorContextOAuth, LogAuditHistoryResponseActorContextOriginCAKey:
		return true
	}
	return false
}

// The type of actor.
type LogAuditHistoryResponseActorType string

const (
	LogAuditHistoryResponseActorTypeAccount         LogAuditHistoryResponseActorType = "account"
	LogAuditHistoryResponseActorTypeCloudflareAdmin LogAuditHistoryResponseActorType = "cloudflare_admin"
	LogAuditHistoryResponseActorTypeSystem          LogAuditHistoryResponseActorType = "system"
	LogAuditHistoryResponseActorTypeUser            LogAuditHistoryResponseActorType = "user"
)

func (r LogAuditHistoryResponseActorType) IsKnown() bool {
	switch r {
	case LogAuditHistoryResponseActorTypeAccount, LogAuditHistoryResponseActorTypeCloudflareAdmin, LogAuditHistoryResponseActorTypeSystem, LogAuditHistoryResponseActorTypeUser:
		return true
	}
	return false
}

// Provides raw information about the request and response.
type LogAuditHistoryResponseRaw struct {
	// The Cloudflare Ray ID for the request.
	CfRayID string `json:"cf_ray_id"`
	// The HTTP method of the request.
	Method string `json:"method"`
	// The HTTP response status code returned by the API.
	StatusCode int64 `json:"status_code"`
	// The URI of the request.
	URI string `json:"uri"`
	// The client's user agent string sent with the request.
	UserAgent string                         `json:"user_agent"`
	JSON      logAuditHistoryResponseRawJSON `json:"-"`
}

// logAuditHistoryResponseRawJSON contains the JSON metadata for the struct
// [LogAuditHistoryResponseRaw]
type logAuditHistoryResponseRawJSON struct {
	CfRayID     apijson.Field
	Method      apijson.Field
	StatusCode  apijson.Field
	URI         apijson.Field
	UserAgent   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditHistoryResponseRaw) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditHistoryResponseRawJSON) RawJSON() string {
	return r.raw
}

// Provides details about the affected resource.
type LogAuditHistoryResponseResource struct {
	// The unique identifier for the affected resource.
	ID string `json:"id"`
	// The Cloudflare product associated with the resource.
	Product  string      `json:"product"`
	Request  interface{} `json:"request"`
	Response interface{} `json:"response"`
	// The scope of the resource.
	Scope interface{} `json:"scope"`
	// The type of the resource.
	Type string                              `json:"type"`
	JSON logAuditHistoryResponseResourceJSON `json:"-"`
}

// logAuditHistoryResponseResourceJSON contains the JSON metadata for the struct
// [LogAuditHistoryResponseResource]
type logAuditHistoryResponseResourceJSON struct {
	ID          apijson.Field
	Product     apijson.Field
	Request     apijson.Field
	Response    apijson.Field
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditHistoryResponseResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditHistoryResponseResourceJSON) RawJSON() string {
	return r.raw
}

// Provides details about the zone affected by the action.
type LogAuditHistoryResponseZone struct {
	// A string that identifies the zone id.
	ID string `json:"id"`
	// A string that identifies the zone name.
	Name string                          `json:"name"`
	JSON logAuditHistoryResponseZoneJSON `json:"-"`
}

// logAuditHistoryResponseZoneJSON contains the JSON metadata for the struct
// [LogAuditHistoryResponseZone]
type logAuditHistoryResponseZoneJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditHistoryResponseZone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditHistoryResponseZoneJSON) RawJSON() string {
	return r.raw
}

// A predefined product category and the resource products it expands to.
type LogAuditProductCategoriesResponse struct {
	// A human-readable label for the product category.
	Label string `json:"label"`
	// The resource products that the product category expands to.
	Products []LogAuditProductCategoriesResponseProduct `json:"products"`
	// The product category identifier used with the product_category filter.
	Value string                                `json:"value"`
	JSON  logAuditProductCategoriesResponseJSON `json:"-"`
}

// logAuditProductCategoriesResponseJSON contains the JSON metadata for the struct
// [LogAuditProductCategoriesResponse]
type logAuditProductCategoriesResponseJSON struct {
	Label       apijson.Field
	Products    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditProductCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditProductCategoriesResponseJSON) RawJSON() string {
	return r.raw
}

// A resource product within a product category.
type LogAuditProductCategoriesResponseProduct struct {
	// A human-readable label for the product.
	Label string `json:"label"`
	// The resource_product value that the product category expands to.
	Value string                                       `json:"value"`
	JSON  logAuditProductCategoriesResponseProductJSON `json:"-"`
}

// logAuditProductCategoriesResponseProductJSON contains the JSON metadata for the
// struct [LogAuditProductCategoriesResponseProduct]
type logAuditProductCategoriesResponseProductJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditProductCategoriesResponseProduct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditProductCategoriesResponseProductJSON) RawJSON() string {
	return r.raw
}

type LogAuditListParams struct {
	// The unique id that identifies the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Limits the returned results to logs older than the specified date. This can be a
	// date string 2019-04-30 (interpreted in UTC) or an absolute timestamp that
	// conforms to RFC3339.
	Before param.Field[time.Time] `query:"before" api:"required" format:"date"`
	// Limits the returned results to logs newer than the specified date. This can be a
	// date string 2019-04-30 (interpreted in UTC) or an absolute timestamp that
	// conforms to RFC3339.
	Since          param.Field[time.Time]                        `query:"since" api:"required" format:"date"`
	ID             param.Field[LogAuditListParamsID]             `query:"id"`
	AccountName    param.Field[LogAuditListParamsAccountName]    `query:"account_name"`
	ActionResult   param.Field[LogAuditListParamsActionResult]   `query:"action_result"`
	ActionType     param.Field[LogAuditListParamsActionType]     `query:"action_type"`
	ActorContext   param.Field[LogAuditListParamsActorContext]   `query:"actor_context"`
	ActorEmail     param.Field[LogAuditListParamsActorEmail]     `query:"actor_email"`
	ActorID        param.Field[LogAuditListParamsActorID]        `query:"actor_id"`
	ActorIPAddress param.Field[LogAuditListParamsActorIPAddress] `query:"actor_ip_address"`
	ActorTokenID   param.Field[LogAuditListParamsActorTokenID]   `query:"actor_token_id"`
	ActorTokenName param.Field[LogAuditListParamsActorTokenName] `query:"actor_token_name"`
	ActorType      param.Field[LogAuditListParamsActorType]      `query:"actor_type"`
	AuditLogID     param.Field[LogAuditListParamsAuditLogID]     `query:"audit_log_id"`
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
	// Filters audit logs by one or more predefined product categories. Each product
	// category expands into a curated set of resource_product values and is unioned
	// with any explicit resource_product filter. Matched case-insensitively; unknown
	// product categories return 400. Repeatable. Use the audit log product categories
	// endpoint to discover the available values.
	ProductCategory param.Field[[]string]                          `query:"product_category"`
	RawCfRayID      param.Field[LogAuditListParamsRawCfRayID]      `query:"raw_cf_ray_id"`
	RawMethod       param.Field[LogAuditListParamsRawMethod]       `query:"raw_method"`
	RawStatusCode   param.Field[LogAuditListParamsRawStatusCode]   `query:"raw_status_code"`
	RawURI          param.Field[LogAuditListParamsRawURI]          `query:"raw_uri"`
	ResourceID      param.Field[LogAuditListParamsResourceID]      `query:"resource_id"`
	ResourceProduct param.Field[LogAuditListParamsResourceProduct] `query:"resource_product"`
	ResourceScope   param.Field[LogAuditListParamsResourceScope]   `query:"resource_scope"`
	ResourceType    param.Field[LogAuditListParamsResourceType]    `query:"resource_type"`
	ZoneID          param.Field[LogAuditListParamsZoneID]          `query:"zone_id"`
	ZoneName        param.Field[LogAuditListParamsZoneName]        `query:"zone_name"`
}

// URLQuery serializes [LogAuditListParams]'s query parameters as `url.Values`.
func (r LogAuditListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsID struct {
	// Filters out audit logs by their IDs.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsID]'s query parameters as `url.Values`.
func (r LogAuditListParamsID) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsAccountName struct {
	// Filters out audit logs by the account name.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsAccountName]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsAccountName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsActionResult struct {
	// Filters out audit logs by whether the action was successful or not.
	Not param.Field[[]LogAuditListParamsActionResultNot] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsActionResult]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsActionResult) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsActionResultNot string

const (
	LogAuditListParamsActionResultNotSuccess LogAuditListParamsActionResultNot = "success"
	LogAuditListParamsActionResultNotFailure LogAuditListParamsActionResultNot = "failure"
)

func (r LogAuditListParamsActionResultNot) IsKnown() bool {
	switch r {
	case LogAuditListParamsActionResultNotSuccess, LogAuditListParamsActionResultNotFailure:
		return true
	}
	return false
}

type LogAuditListParamsActionType struct {
	// Filters out audit logs by the action type.
	Not param.Field[[]LogAuditListParamsActionTypeNot] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsActionType]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsActionType) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsActionTypeNot string

const (
	LogAuditListParamsActionTypeNotCreate LogAuditListParamsActionTypeNot = "create"
	LogAuditListParamsActionTypeNotDelete LogAuditListParamsActionTypeNot = "delete"
	LogAuditListParamsActionTypeNotView   LogAuditListParamsActionTypeNot = "view"
	LogAuditListParamsActionTypeNotUpdate LogAuditListParamsActionTypeNot = "update"
)

func (r LogAuditListParamsActionTypeNot) IsKnown() bool {
	switch r {
	case LogAuditListParamsActionTypeNotCreate, LogAuditListParamsActionTypeNotDelete, LogAuditListParamsActionTypeNotView, LogAuditListParamsActionTypeNotUpdate:
		return true
	}
	return false
}

type LogAuditListParamsActorContext struct {
	// Filters out audit logs by the actor context.
	Not param.Field[[]LogAuditListParamsActorContextNot] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsActorContext]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsActorContext) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsActorContextNot string

const (
	LogAuditListParamsActorContextNotAPIKey      LogAuditListParamsActorContextNot = "api_key"
	LogAuditListParamsActorContextNotAPIToken    LogAuditListParamsActorContextNot = "api_token"
	LogAuditListParamsActorContextNotDash        LogAuditListParamsActorContextNot = "dash"
	LogAuditListParamsActorContextNotOAuth       LogAuditListParamsActorContextNot = "oauth"
	LogAuditListParamsActorContextNotOriginCAKey LogAuditListParamsActorContextNot = "origin_ca_key"
)

func (r LogAuditListParamsActorContextNot) IsKnown() bool {
	switch r {
	case LogAuditListParamsActorContextNotAPIKey, LogAuditListParamsActorContextNotAPIToken, LogAuditListParamsActorContextNotDash, LogAuditListParamsActorContextNotOAuth, LogAuditListParamsActorContextNotOriginCAKey:
		return true
	}
	return false
}

type LogAuditListParamsActorEmail struct {
	// Filters out audit logs by the actor's email address.
	Not param.Field[[]string] `query:"not" format:"email"`
}

// URLQuery serializes [LogAuditListParamsActorEmail]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsActorEmail) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsActorID struct {
	// Filters out audit logs by the actor ID. This can be either the Account ID or
	// User ID.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsActorID]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsActorID) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsActorIPAddress struct {
	// Filters out audit logs IP address where the action was initiated.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsActorIPAddress]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsActorIPAddress) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsActorTokenID struct {
	// Filters out audit logs by the API token ID when the actor context is an
	// api_token or oauth.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsActorTokenID]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsActorTokenID) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsActorTokenName struct {
	// Filters out audit logs by the API token name when the actor context is an
	// api_token or oauth.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsActorTokenName]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsActorTokenName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsActorType struct {
	// Filters out audit logs by the actor type.
	Not param.Field[[]LogAuditListParamsActorTypeNot] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsActorType]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsActorType) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsActorTypeNot string

const (
	LogAuditListParamsActorTypeNotAccount         LogAuditListParamsActorTypeNot = "account"
	LogAuditListParamsActorTypeNotCloudflareAdmin LogAuditListParamsActorTypeNot = "cloudflare_admin"
	LogAuditListParamsActorTypeNotSystem          LogAuditListParamsActorTypeNot = "system"
	LogAuditListParamsActorTypeNotUser            LogAuditListParamsActorTypeNot = "user"
)

func (r LogAuditListParamsActorTypeNot) IsKnown() bool {
	switch r {
	case LogAuditListParamsActorTypeNotAccount, LogAuditListParamsActorTypeNotCloudflareAdmin, LogAuditListParamsActorTypeNotSystem, LogAuditListParamsActorTypeNotUser:
		return true
	}
	return false
}

type LogAuditListParamsAuditLogID struct {
	// Filters out audit logs by their IDs.
	//
	// Deprecated: deprecated
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsAuditLogID]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsAuditLogID) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
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

type LogAuditListParamsRawCfRayID struct {
	// Filters out audit logs by the response CF Ray ID.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsRawCfRayID]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsRawCfRayID) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsRawMethod struct {
	// Filters out audit logs by the HTTP method for the API call.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsRawMethod]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsRawMethod) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsRawStatusCode struct {
	// Filters out audit logs by the response status code that was returned.
	Not param.Field[[]int64] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsRawStatusCode]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsRawStatusCode) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsRawURI struct {
	// Filters out audit logs by the request URI.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsRawURI]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsRawURI) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsResourceID struct {
	// Filters out audit logs by the resource ID.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsResourceID]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsResourceID) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsResourceProduct struct {
	// Filters out audit logs by the Cloudflare product associated with the changed
	// resource.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsResourceProduct]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsResourceProduct) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsResourceScope struct {
	// Filters out audit logs by the resource scope, specifying whether the resource is
	// associated with an user, an account, a zone, or a membership.
	Not param.Field[[]LogAuditListParamsResourceScopeNot] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsResourceScope]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsResourceScope) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsResourceScopeNot string

const (
	LogAuditListParamsResourceScopeNotAccounts    LogAuditListParamsResourceScopeNot = "accounts"
	LogAuditListParamsResourceScopeNotUser        LogAuditListParamsResourceScopeNot = "user"
	LogAuditListParamsResourceScopeNotZones       LogAuditListParamsResourceScopeNot = "zones"
	LogAuditListParamsResourceScopeNotMemberships LogAuditListParamsResourceScopeNot = "memberships"
)

func (r LogAuditListParamsResourceScopeNot) IsKnown() bool {
	switch r {
	case LogAuditListParamsResourceScopeNotAccounts, LogAuditListParamsResourceScopeNotUser, LogAuditListParamsResourceScopeNotZones, LogAuditListParamsResourceScopeNotMemberships:
		return true
	}
	return false
}

type LogAuditListParamsResourceType struct {
	// Filters out audit logs based on the unique type of resource changed by the
	// action.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsResourceType]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsResourceType) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsZoneID struct {
	// Filters out audit logs by the zone ID.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsZoneID]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsZoneID) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditListParamsZoneName struct {
	// Filters out audit logs by the zone name associated with the change.
	Not param.Field[[]string] `query:"not"`
}

// URLQuery serializes [LogAuditListParamsZoneName]'s query parameters as
// `url.Values`.
func (r LogAuditListParamsZoneName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogAuditHistoryParams struct {
	// The unique ID that identifies the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// RFC3339 timestamp of the source audit log entry's action time. Used to narrow
	// the source-entry lookup window. Provide the `action.time` value from the audit
	// log identified by `id`.
	ActionTime param.Field[time.Time] `query:"action_time" api:"required" format:"date-time"`
	// Limits the returned results to logs older than the specified date. This can be a
	// date string 2019-04-30 (interpreted in UTC) or an absolute timestamp that
	// conforms to RFC3339.
	Before param.Field[time.Time] `query:"before" api:"required" format:"date"`
	// Limits the returned results to logs newer than the specified date. This can be a
	// date string 2019-04-30 (interpreted in UTC) or an absolute timestamp that
	// conforms to RFC3339.
	Since param.Field[time.Time] `query:"since" api:"required" format:"date"`
	// The cursor is an opaque token used to paginate through large sets of records. It
	// indicates the position from which to continue when requesting the next set of
	// records. A valid cursor value can be obtained from the cursor object in the
	// result_info structure of a previous response.
	Cursor param.Field[string] `query:"cursor"`
	// Sets sorting order.
	Direction param.Field[LogAuditHistoryParamsDirection] `query:"direction"`
	// The number limits the objects to return. The cursor attribute may be used to
	// iterate over the next batch of objects if there are more than the limit.
	Limit param.Field[float64] `query:"limit"`
}

// URLQuery serializes [LogAuditHistoryParams]'s query parameters as `url.Values`.
func (r LogAuditHistoryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sets sorting order.
type LogAuditHistoryParamsDirection string

const (
	LogAuditHistoryParamsDirectionDesc LogAuditHistoryParamsDirection = "desc"
	LogAuditHistoryParamsDirectionAsc  LogAuditHistoryParamsDirection = "asc"
)

func (r LogAuditHistoryParamsDirection) IsKnown() bool {
	switch r {
	case LogAuditHistoryParamsDirectionDesc, LogAuditHistoryParamsDirectionAsc:
		return true
	}
	return false
}

type LogAuditHistoryResponseEnvelope struct {
	Errors []LogAuditHistoryResponseEnvelopeErrors `json:"errors" api:"required"`
	Result []LogAuditHistoryResponse               `json:"result" api:"required"`
	// Provides information about the result of the request, including count, cursor,
	// and identification quality.
	ResultInfo LogAuditHistoryResponseEnvelopeResultInfo `json:"result_info" api:"required"`
	// Indicates whether the API call was successful
	Success LogAuditHistoryResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    logAuditHistoryResponseEnvelopeJSON    `json:"-"`
}

// logAuditHistoryResponseEnvelopeJSON contains the JSON metadata for the struct
// [LogAuditHistoryResponseEnvelope]
type logAuditHistoryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditHistoryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditHistoryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LogAuditHistoryResponseEnvelopeErrors struct {
	Message string                                    `json:"message" api:"required"`
	JSON    logAuditHistoryResponseEnvelopeErrorsJSON `json:"-"`
}

// logAuditHistoryResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LogAuditHistoryResponseEnvelopeErrors]
type logAuditHistoryResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogAuditHistoryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditHistoryResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// Provides information about the result of the request, including count, cursor,
// and identification quality.
type LogAuditHistoryResponseEnvelopeResultInfo struct {
	// The number of records returned in the response.
	Count int64 `json:"count" api:"required"`
	// Indicates the quality of the resource identification used to derive the history.
	//
	//   - `exact`: Resource was identified by the resource URI.
	//   - `approximate`: Resource was identified without the resource URI.
	//   - `unavailable`: The source audit log entry did not contain enough information
	//     to identify the resource; result is empty.
	HistoryStatus LogAuditHistoryResponseEnvelopeResultInfoHistoryStatus `json:"history_status" api:"required"`
	// The cursor token used for pagination.
	Cursor string                                        `json:"cursor"`
	JSON   logAuditHistoryResponseEnvelopeResultInfoJSON `json:"-"`
}

// logAuditHistoryResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [LogAuditHistoryResponseEnvelopeResultInfo]
type logAuditHistoryResponseEnvelopeResultInfoJSON struct {
	Count         apijson.Field
	HistoryStatus apijson.Field
	Cursor        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LogAuditHistoryResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logAuditHistoryResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

// Indicates the quality of the resource identification used to derive the history.
//
//   - `exact`: Resource was identified by the resource URI.
//   - `approximate`: Resource was identified without the resource URI.
//   - `unavailable`: The source audit log entry did not contain enough information
//     to identify the resource; result is empty.
type LogAuditHistoryResponseEnvelopeResultInfoHistoryStatus string

const (
	LogAuditHistoryResponseEnvelopeResultInfoHistoryStatusExact       LogAuditHistoryResponseEnvelopeResultInfoHistoryStatus = "exact"
	LogAuditHistoryResponseEnvelopeResultInfoHistoryStatusApproximate LogAuditHistoryResponseEnvelopeResultInfoHistoryStatus = "approximate"
	LogAuditHistoryResponseEnvelopeResultInfoHistoryStatusUnavailable LogAuditHistoryResponseEnvelopeResultInfoHistoryStatus = "unavailable"
)

func (r LogAuditHistoryResponseEnvelopeResultInfoHistoryStatus) IsKnown() bool {
	switch r {
	case LogAuditHistoryResponseEnvelopeResultInfoHistoryStatusExact, LogAuditHistoryResponseEnvelopeResultInfoHistoryStatusApproximate, LogAuditHistoryResponseEnvelopeResultInfoHistoryStatusUnavailable:
		return true
	}
	return false
}

// Indicates whether the API call was successful
type LogAuditHistoryResponseEnvelopeSuccess bool

const (
	LogAuditHistoryResponseEnvelopeSuccessTrue LogAuditHistoryResponseEnvelopeSuccess = true
)

func (r LogAuditHistoryResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LogAuditHistoryResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LogAuditProductCategoriesParams struct {
	// The unique id that identifies the account.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

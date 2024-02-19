// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
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

// Gets a list of audit logs for an account. Can be filtered by who made the
// change, on which zone, and the timeframe of the change.
func (r *AuditLogService) AuditLogsGetAccountAuditLogs(ctx context.Context, accountIdentifier string, query AuditLogAuditLogsGetAccountAuditLogsParams, opts ...option.RequestOption) (res *AuditLogAuditLogsGetAccountAuditLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/audit_logs", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Union satisfied by [AuditLogAuditLogsGetAccountAuditLogsResponseObject] or
// [AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommon].
type AuditLogAuditLogsGetAccountAuditLogsResponse interface {
	implementsAuditLogAuditLogsGetAccountAuditLogsResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*AuditLogAuditLogsGetAccountAuditLogsResponse)(nil)).Elem(), "")
}

type AuditLogAuditLogsGetAccountAuditLogsResponseObject struct {
	Errors   interface{}                                                `json:"errors,nullable"`
	Messages []interface{}                                              `json:"messages"`
	Result   []AuditLogAuditLogsGetAccountAuditLogsResponseObjectResult `json:"result"`
	Success  bool                                                       `json:"success"`
	JSON     auditLogAuditLogsGetAccountAuditLogsResponseObjectJSON     `json:"-"`
}

// auditLogAuditLogsGetAccountAuditLogsResponseObjectJSON contains the JSON
// metadata for the struct [AuditLogAuditLogsGetAccountAuditLogsResponseObject]
type auditLogAuditLogsGetAccountAuditLogsResponseObjectJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogAuditLogsGetAccountAuditLogsResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AuditLogAuditLogsGetAccountAuditLogsResponseObject) implementsAuditLogAuditLogsGetAccountAuditLogsResponse() {
}

type AuditLogAuditLogsGetAccountAuditLogsResponseObjectResult struct {
	// A string that uniquely identifies the audit log.
	ID     string                                                         `json:"id"`
	Action AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultAction `json:"action"`
	Actor  AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultActor  `json:"actor"`
	// The source of the event.
	Interface string `json:"interface"`
	// An object which can lend more context to the action being logged. This is a
	// flexible value and varies between different actions.
	Metadata interface{} `json:"metadata"`
	// The new value of the resource that was modified.
	NewValue string `json:"newValue"`
	// The value of the resource before it was modified.
	OldValue string                                                           `json:"oldValue"`
	Owner    AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultOwner    `json:"owner"`
	Resource AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultResource `json:"resource"`
	// A UTC RFC3339 timestamp that specifies when the action being logged occured.
	When time.Time                                                    `json:"when" format:"date-time"`
	JSON auditLogAuditLogsGetAccountAuditLogsResponseObjectResultJSON `json:"-"`
}

// auditLogAuditLogsGetAccountAuditLogsResponseObjectResultJSON contains the JSON
// metadata for the struct
// [AuditLogAuditLogsGetAccountAuditLogsResponseObjectResult]
type auditLogAuditLogsGetAccountAuditLogsResponseObjectResultJSON struct {
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

func (r *AuditLogAuditLogsGetAccountAuditLogsResponseObjectResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultAction struct {
	// A boolean that indicates if the action attempted was successful.
	Result bool `json:"result"`
	// A short string that describes the action that was performed.
	Type string                                                             `json:"type"`
	JSON auditLogAuditLogsGetAccountAuditLogsResponseObjectResultActionJSON `json:"-"`
}

// auditLogAuditLogsGetAccountAuditLogsResponseObjectResultActionJSON contains the
// JSON metadata for the struct
// [AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultAction]
type auditLogAuditLogsGetAccountAuditLogsResponseObjectResultActionJSON struct {
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultActor struct {
	// The ID of the actor that performed the action. If a user performed the action,
	// this will be their User ID.
	ID string `json:"id"`
	// The email of the user that performed the action.
	Email string `json:"email" format:"email"`
	// The IP address of the request that performed the action.
	IP string `json:"ip"`
	// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
	Type AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultActorType `json:"type"`
	JSON auditLogAuditLogsGetAccountAuditLogsResponseObjectResultActorJSON `json:"-"`
}

// auditLogAuditLogsGetAccountAuditLogsResponseObjectResultActorJSON contains the
// JSON metadata for the struct
// [AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultActor]
type auditLogAuditLogsGetAccountAuditLogsResponseObjectResultActorJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	IP          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
type AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultActorType string

const (
	AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultActorTypeUser       AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultActorType = "user"
	AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultActorTypeAdmin      AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultActorType = "admin"
	AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultActorTypeCloudflare AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultActorType = "Cloudflare"
)

type AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultOwner struct {
	// Identifier
	ID   string                                                            `json:"id"`
	JSON auditLogAuditLogsGetAccountAuditLogsResponseObjectResultOwnerJSON `json:"-"`
}

// auditLogAuditLogsGetAccountAuditLogsResponseObjectResultOwnerJSON contains the
// JSON metadata for the struct
// [AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultOwner]
type auditLogAuditLogsGetAccountAuditLogsResponseObjectResultOwnerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultResource struct {
	// An identifier for the resource that was affected by the action.
	ID string `json:"id"`
	// A short string that describes the resource that was affected by the action.
	Type string                                                               `json:"type"`
	JSON auditLogAuditLogsGetAccountAuditLogsResponseObjectResultResourceJSON `json:"-"`
}

// auditLogAuditLogsGetAccountAuditLogsResponseObjectResultResourceJSON contains
// the JSON metadata for the struct
// [AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultResource]
type auditLogAuditLogsGetAccountAuditLogsResponseObjectResultResourceJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogAuditLogsGetAccountAuditLogsResponseObjectResultResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommon struct {
	Errors   []AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonError   `json:"errors,required"`
	Messages []AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonMessage `json:"messages,required"`
	Result   AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonResult    `json:"result,required"`
	// Whether the API call was successful
	Success AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonSuccess `json:"success,required"`
	JSON    auditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonJSON    `json:"-"`
}

// auditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonJSON
// contains the JSON metadata for the struct
// [AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommon]
type auditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommon) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommon) implementsAuditLogAuditLogsGetAccountAuditLogsResponse() {
}

type AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonError struct {
	Code    int64                                                                          `json:"code,required"`
	Message string                                                                         `json:"message,required"`
	JSON    auditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonErrorJSON `json:"-"`
}

// auditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonErrorJSON
// contains the JSON metadata for the struct
// [AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonError]
type auditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonMessage struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    auditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonMessageJSON `json:"-"`
}

// auditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonMessageJSON
// contains the JSON metadata for the struct
// [AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonMessage]
type auditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by
// [AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonResultUnknown],
// [AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonResultArray]
// or [shared.UnionString].
type AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonResult interface {
	ImplementsAuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonResult)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonResultArray []interface{}

func (r AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonResultArray) ImplementsAuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonResult() {
}

// Whether the API call was successful
type AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonSuccess bool

const (
	AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonSuccessTrue AuditLogAuditLogsGetAccountAuditLogsResponseVHtGuEw1APIResponseCommonSuccess = true
)

type AuditLogAuditLogsGetAccountAuditLogsParams struct {
	// Finds a specific log by its ID.
	ID     param.Field[string]                                           `query:"id"`
	Action param.Field[AuditLogAuditLogsGetAccountAuditLogsParamsAction] `query:"action"`
	Actor  param.Field[AuditLogAuditLogsGetAccountAuditLogsParamsActor]  `query:"actor"`
	// Limits the returned results to logs older than the specified date. This can be a
	// date string `2019-04-30` or an absolute timestamp that conforms to RFC3339.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Changes the direction of the chronological sorting.
	Direction param.Field[AuditLogAuditLogsGetAccountAuditLogsParamsDirection] `query:"direction"`
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
	Since param.Field[time.Time]                                      `query:"since" format:"date-time"`
	Zone  param.Field[AuditLogAuditLogsGetAccountAuditLogsParamsZone] `query:"zone"`
}

// URLQuery serializes [AuditLogAuditLogsGetAccountAuditLogsParams]'s query
// parameters as `url.Values`.
func (r AuditLogAuditLogsGetAccountAuditLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AuditLogAuditLogsGetAccountAuditLogsParamsAction struct {
	// Filters by the action type.
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [AuditLogAuditLogsGetAccountAuditLogsParamsAction]'s query
// parameters as `url.Values`.
func (r AuditLogAuditLogsGetAccountAuditLogsParamsAction) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AuditLogAuditLogsGetAccountAuditLogsParamsActor struct {
	// Filters by the email address of the actor that made the change.
	Email param.Field[string] `query:"email" format:"email"`
	// Filters by the IP address of the request that made the change by specific IP
	// address or valid CIDR Range.
	IP param.Field[string] `query:"ip"`
}

// URLQuery serializes [AuditLogAuditLogsGetAccountAuditLogsParamsActor]'s query
// parameters as `url.Values`.
func (r AuditLogAuditLogsGetAccountAuditLogsParamsActor) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Changes the direction of the chronological sorting.
type AuditLogAuditLogsGetAccountAuditLogsParamsDirection string

const (
	AuditLogAuditLogsGetAccountAuditLogsParamsDirectionDesc AuditLogAuditLogsGetAccountAuditLogsParamsDirection = "desc"
	AuditLogAuditLogsGetAccountAuditLogsParamsDirectionAsc  AuditLogAuditLogsGetAccountAuditLogsParamsDirection = "asc"
)

type AuditLogAuditLogsGetAccountAuditLogsParamsZone struct {
	// Filters by the name of the zone associated to the change.
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [AuditLogAuditLogsGetAccountAuditLogsParamsZone]'s query
// parameters as `url.Values`.
func (r AuditLogAuditLogsGetAccountAuditLogsParamsZone) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

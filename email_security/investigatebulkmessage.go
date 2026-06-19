// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/tidwall/gjson"
)

// InvestigateBulkMessageService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateBulkMessageService] method instead.
type InvestigateBulkMessageService struct {
	Options []option.RequestOption
}

// NewInvestigateBulkMessageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestigateBulkMessageService(opts ...option.RequestOption) (r *InvestigateBulkMessageService) {
	r = &InvestigateBulkMessageService{}
	r.Options = opts
	return
}

// List messages for a bulk action job
func (r *InvestigateBulkMessageService) List(ctx context.Context, jobID string, params InvestigateBulkMessageListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[InvestigateBulkMessageListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/bulk/%s/messages", params.AccountID, jobID)
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

// List messages for a bulk action job
func (r *InvestigateBulkMessageService) ListAutoPaging(ctx context.Context, jobID string, params InvestigateBulkMessageListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[InvestigateBulkMessageListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, jobID, params, opts...))
}

type InvestigateBulkMessageListResponse struct {
	ActionParams   InvestigateBulkMessageListResponseActionParams `json:"action_params" api:"required"`
	ActionType     InvestigateBulkMessageListResponseActionType   `json:"action_type" api:"required"`
	CreatedAt      time.Time                                      `json:"created_at" api:"required" format:"date-time"`
	MessageID      string                                         `json:"message_id" api:"required" format:"uuid"`
	PostfixID      string                                         `json:"postfix_id" api:"required"`
	RetryCount     int64                                          `json:"retry_count" api:"required"`
	Status         InvestigateBulkMessageListResponseStatus       `json:"status" api:"required"`
	AlertID        string                                         `json:"alert_id" api:"nullable"`
	EmailMessageID string                                         `json:"email_message_id" api:"nullable"`
	ProcessedAt    time.Time                                      `json:"processed_at" api:"nullable" format:"date-time"`
	// When to retry the action if it failed
	RetryAfter    time.Time                              `json:"retry_after" api:"nullable" format:"date-time"`
	StatusMessage string                                 `json:"status_message" api:"nullable"`
	JSON          investigateBulkMessageListResponseJSON `json:"-"`
}

// investigateBulkMessageListResponseJSON contains the JSON metadata for the struct
// [InvestigateBulkMessageListResponse]
type investigateBulkMessageListResponseJSON struct {
	ActionParams   apijson.Field
	ActionType     apijson.Field
	CreatedAt      apijson.Field
	MessageID      apijson.Field
	PostfixID      apijson.Field
	RetryCount     apijson.Field
	Status         apijson.Field
	AlertID        apijson.Field
	EmailMessageID apijson.Field
	ProcessedAt    apijson.Field
	RetryAfter     apijson.Field
	StatusMessage  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *InvestigateBulkMessageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkMessageListResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateBulkMessageListResponseActionParams struct {
	ClientRecipient     string                                                            `json:"client_recipient" api:"required"`
	Type                InvestigateBulkMessageListResponseActionParamsType                `json:"type" api:"required"`
	Destination         InvestigateBulkMessageListResponseActionParamsDestination         `json:"destination"`
	ExpectedDisposition InvestigateBulkMessageListResponseActionParamsExpectedDisposition `json:"expected_disposition"`
	JSON                investigateBulkMessageListResponseActionParamsJSON                `json:"-"`
	union               InvestigateBulkMessageListResponseActionParamsUnion
}

// investigateBulkMessageListResponseActionParamsJSON contains the JSON metadata
// for the struct [InvestigateBulkMessageListResponseActionParams]
type investigateBulkMessageListResponseActionParamsJSON struct {
	ClientRecipient     apijson.Field
	Type                apijson.Field
	Destination         apijson.Field
	ExpectedDisposition apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r investigateBulkMessageListResponseActionParamsJSON) RawJSON() string {
	return r.raw
}

func (r *InvestigateBulkMessageListResponseActionParams) UnmarshalJSON(data []byte) (err error) {
	*r = InvestigateBulkMessageListResponseActionParams{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InvestigateBulkMessageListResponseActionParamsUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [InvestigateBulkMessageListResponseActionParamsMove],
// [InvestigateBulkMessageListResponseActionParamsRelease].
func (r InvestigateBulkMessageListResponseActionParams) AsUnion() InvestigateBulkMessageListResponseActionParamsUnion {
	return r.union
}

// Union satisfied by [InvestigateBulkMessageListResponseActionParamsMove] or
// [InvestigateBulkMessageListResponseActionParamsRelease].
type InvestigateBulkMessageListResponseActionParamsUnion interface {
	implementsInvestigateBulkMessageListResponseActionParams()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InvestigateBulkMessageListResponseActionParamsUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvestigateBulkMessageListResponseActionParamsMove{}),
			DiscriminatorValue: "MOVE",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(InvestigateBulkMessageListResponseActionParamsRelease{}),
			DiscriminatorValue: "RELEASE",
		},
	)
}

type InvestigateBulkMessageListResponseActionParamsMove struct {
	ClientRecipient     string                                                                `json:"client_recipient" api:"required"`
	Destination         InvestigateBulkMessageListResponseActionParamsMoveDestination         `json:"destination" api:"required"`
	Type                InvestigateBulkMessageListResponseActionParamsMoveType                `json:"type" api:"required"`
	ExpectedDisposition InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition `json:"expected_disposition"`
	JSON                investigateBulkMessageListResponseActionParamsMoveJSON                `json:"-"`
}

// investigateBulkMessageListResponseActionParamsMoveJSON contains the JSON
// metadata for the struct [InvestigateBulkMessageListResponseActionParamsMove]
type investigateBulkMessageListResponseActionParamsMoveJSON struct {
	ClientRecipient     apijson.Field
	Destination         apijson.Field
	Type                apijson.Field
	ExpectedDisposition apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *InvestigateBulkMessageListResponseActionParamsMove) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkMessageListResponseActionParamsMoveJSON) RawJSON() string {
	return r.raw
}

func (r InvestigateBulkMessageListResponseActionParamsMove) implementsInvestigateBulkMessageListResponseActionParams() {
}

type InvestigateBulkMessageListResponseActionParamsMoveDestination string

const (
	InvestigateBulkMessageListResponseActionParamsMoveDestinationInbox                     InvestigateBulkMessageListResponseActionParamsMoveDestination = "Inbox"
	InvestigateBulkMessageListResponseActionParamsMoveDestinationJunkEmail                 InvestigateBulkMessageListResponseActionParamsMoveDestination = "JunkEmail"
	InvestigateBulkMessageListResponseActionParamsMoveDestinationDeletedItems              InvestigateBulkMessageListResponseActionParamsMoveDestination = "DeletedItems"
	InvestigateBulkMessageListResponseActionParamsMoveDestinationRecoverableItemsDeletions InvestigateBulkMessageListResponseActionParamsMoveDestination = "RecoverableItemsDeletions"
	InvestigateBulkMessageListResponseActionParamsMoveDestinationRecoverableItemsPurges    InvestigateBulkMessageListResponseActionParamsMoveDestination = "RecoverableItemsPurges"
)

func (r InvestigateBulkMessageListResponseActionParamsMoveDestination) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionParamsMoveDestinationInbox, InvestigateBulkMessageListResponseActionParamsMoveDestinationJunkEmail, InvestigateBulkMessageListResponseActionParamsMoveDestinationDeletedItems, InvestigateBulkMessageListResponseActionParamsMoveDestinationRecoverableItemsDeletions, InvestigateBulkMessageListResponseActionParamsMoveDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseActionParamsMoveType string

const (
	InvestigateBulkMessageListResponseActionParamsMoveTypeMove InvestigateBulkMessageListResponseActionParamsMoveType = "MOVE"
)

func (r InvestigateBulkMessageListResponseActionParamsMoveType) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionParamsMoveTypeMove:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition string

const (
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionMalicious    InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "MALICIOUS"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionMaliciousBec InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "MALICIOUS-BEC"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionSuspicious   InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "SUSPICIOUS"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionSpoof        InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "SPOOF"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionSpam         InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "SPAM"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionBulk         InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "BULK"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionEncrypted    InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "ENCRYPTED"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionExternal     InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "EXTERNAL"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionUnknown      InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "UNKNOWN"
	InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionNone         InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition = "NONE"
)

func (r InvestigateBulkMessageListResponseActionParamsMoveExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionMalicious, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionMaliciousBec, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionSuspicious, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionSpoof, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionSpam, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionBulk, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionEncrypted, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionExternal, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionUnknown, InvestigateBulkMessageListResponseActionParamsMoveExpectedDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseActionParamsRelease struct {
	ClientRecipient string                                                    `json:"client_recipient" api:"required"`
	Type            InvestigateBulkMessageListResponseActionParamsReleaseType `json:"type" api:"required"`
	JSON            investigateBulkMessageListResponseActionParamsReleaseJSON `json:"-"`
}

// investigateBulkMessageListResponseActionParamsReleaseJSON contains the JSON
// metadata for the struct [InvestigateBulkMessageListResponseActionParamsRelease]
type investigateBulkMessageListResponseActionParamsReleaseJSON struct {
	ClientRecipient apijson.Field
	Type            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *InvestigateBulkMessageListResponseActionParamsRelease) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateBulkMessageListResponseActionParamsReleaseJSON) RawJSON() string {
	return r.raw
}

func (r InvestigateBulkMessageListResponseActionParamsRelease) implementsInvestigateBulkMessageListResponseActionParams() {
}

type InvestigateBulkMessageListResponseActionParamsReleaseType string

const (
	InvestigateBulkMessageListResponseActionParamsReleaseTypeRelease InvestigateBulkMessageListResponseActionParamsReleaseType = "RELEASE"
)

func (r InvestigateBulkMessageListResponseActionParamsReleaseType) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionParamsReleaseTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseActionParamsType string

const (
	InvestigateBulkMessageListResponseActionParamsTypeMove    InvestigateBulkMessageListResponseActionParamsType = "MOVE"
	InvestigateBulkMessageListResponseActionParamsTypeRelease InvestigateBulkMessageListResponseActionParamsType = "RELEASE"
)

func (r InvestigateBulkMessageListResponseActionParamsType) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionParamsTypeMove, InvestigateBulkMessageListResponseActionParamsTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseActionParamsDestination string

const (
	InvestigateBulkMessageListResponseActionParamsDestinationInbox                     InvestigateBulkMessageListResponseActionParamsDestination = "Inbox"
	InvestigateBulkMessageListResponseActionParamsDestinationJunkEmail                 InvestigateBulkMessageListResponseActionParamsDestination = "JunkEmail"
	InvestigateBulkMessageListResponseActionParamsDestinationDeletedItems              InvestigateBulkMessageListResponseActionParamsDestination = "DeletedItems"
	InvestigateBulkMessageListResponseActionParamsDestinationRecoverableItemsDeletions InvestigateBulkMessageListResponseActionParamsDestination = "RecoverableItemsDeletions"
	InvestigateBulkMessageListResponseActionParamsDestinationRecoverableItemsPurges    InvestigateBulkMessageListResponseActionParamsDestination = "RecoverableItemsPurges"
)

func (r InvestigateBulkMessageListResponseActionParamsDestination) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionParamsDestinationInbox, InvestigateBulkMessageListResponseActionParamsDestinationJunkEmail, InvestigateBulkMessageListResponseActionParamsDestinationDeletedItems, InvestigateBulkMessageListResponseActionParamsDestinationRecoverableItemsDeletions, InvestigateBulkMessageListResponseActionParamsDestinationRecoverableItemsPurges:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseActionParamsExpectedDisposition string

const (
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionMalicious    InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "MALICIOUS"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionMaliciousBec InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "MALICIOUS-BEC"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionSuspicious   InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "SUSPICIOUS"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionSpoof        InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "SPOOF"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionSpam         InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "SPAM"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionBulk         InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "BULK"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionEncrypted    InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "ENCRYPTED"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionExternal     InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "EXTERNAL"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionUnknown      InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "UNKNOWN"
	InvestigateBulkMessageListResponseActionParamsExpectedDispositionNone         InvestigateBulkMessageListResponseActionParamsExpectedDisposition = "NONE"
)

func (r InvestigateBulkMessageListResponseActionParamsExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionParamsExpectedDispositionMalicious, InvestigateBulkMessageListResponseActionParamsExpectedDispositionMaliciousBec, InvestigateBulkMessageListResponseActionParamsExpectedDispositionSuspicious, InvestigateBulkMessageListResponseActionParamsExpectedDispositionSpoof, InvestigateBulkMessageListResponseActionParamsExpectedDispositionSpam, InvestigateBulkMessageListResponseActionParamsExpectedDispositionBulk, InvestigateBulkMessageListResponseActionParamsExpectedDispositionEncrypted, InvestigateBulkMessageListResponseActionParamsExpectedDispositionExternal, InvestigateBulkMessageListResponseActionParamsExpectedDispositionUnknown, InvestigateBulkMessageListResponseActionParamsExpectedDispositionNone:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseActionType string

const (
	InvestigateBulkMessageListResponseActionTypeMove    InvestigateBulkMessageListResponseActionType = "MOVE"
	InvestigateBulkMessageListResponseActionTypeRelease InvestigateBulkMessageListResponseActionType = "RELEASE"
)

func (r InvestigateBulkMessageListResponseActionType) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseActionTypeMove, InvestigateBulkMessageListResponseActionTypeRelease:
		return true
	}
	return false
}

type InvestigateBulkMessageListResponseStatus string

const (
	InvestigateBulkMessageListResponseStatusPending     InvestigateBulkMessageListResponseStatus = "PENDING"
	InvestigateBulkMessageListResponseStatusDiscovering InvestigateBulkMessageListResponseStatus = "DISCOVERING"
	InvestigateBulkMessageListResponseStatusProcessing  InvestigateBulkMessageListResponseStatus = "PROCESSING"
	InvestigateBulkMessageListResponseStatusCompleted   InvestigateBulkMessageListResponseStatus = "COMPLETED"
	InvestigateBulkMessageListResponseStatusFailed      InvestigateBulkMessageListResponseStatus = "FAILED"
	InvestigateBulkMessageListResponseStatusCancelled   InvestigateBulkMessageListResponseStatus = "CANCELLED"
	InvestigateBulkMessageListResponseStatusSkipped     InvestigateBulkMessageListResponseStatus = "SKIPPED"
)

func (r InvestigateBulkMessageListResponseStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListResponseStatusPending, InvestigateBulkMessageListResponseStatusDiscovering, InvestigateBulkMessageListResponseStatusProcessing, InvestigateBulkMessageListResponseStatusCompleted, InvestigateBulkMessageListResponseStatusFailed, InvestigateBulkMessageListResponseStatusCancelled, InvestigateBulkMessageListResponseStatusSkipped:
		return true
	}
	return false
}

type InvestigateBulkMessageListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Current page within paginated list of results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page. Maximum value is 1000.
	PerPage param.Field[int64]                                  `query:"per_page"`
	Status  param.Field[InvestigateBulkMessageListParamsStatus] `query:"status"`
}

// URLQuery serializes [InvestigateBulkMessageListParams]'s query parameters as
// `url.Values`.
func (r InvestigateBulkMessageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InvestigateBulkMessageListParamsStatus string

const (
	InvestigateBulkMessageListParamsStatusPending     InvestigateBulkMessageListParamsStatus = "PENDING"
	InvestigateBulkMessageListParamsStatusDiscovering InvestigateBulkMessageListParamsStatus = "DISCOVERING"
	InvestigateBulkMessageListParamsStatusProcessing  InvestigateBulkMessageListParamsStatus = "PROCESSING"
	InvestigateBulkMessageListParamsStatusCompleted   InvestigateBulkMessageListParamsStatus = "COMPLETED"
	InvestigateBulkMessageListParamsStatusFailed      InvestigateBulkMessageListParamsStatus = "FAILED"
	InvestigateBulkMessageListParamsStatusCancelled   InvestigateBulkMessageListParamsStatus = "CANCELLED"
	InvestigateBulkMessageListParamsStatusSkipped     InvestigateBulkMessageListParamsStatus = "SKIPPED"
)

func (r InvestigateBulkMessageListParamsStatus) IsKnown() bool {
	switch r {
	case InvestigateBulkMessageListParamsStatusPending, InvestigateBulkMessageListParamsStatusDiscovering, InvestigateBulkMessageListParamsStatusProcessing, InvestigateBulkMessageListParamsStatusCompleted, InvestigateBulkMessageListParamsStatusFailed, InvestigateBulkMessageListParamsStatusCancelled, InvestigateBulkMessageListParamsStatusSkipped:
		return true
	}
	return false
}

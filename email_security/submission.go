// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

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

// SubmissionService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubmissionService] method instead.
type SubmissionService struct {
	Options []option.RequestOption
}

// NewSubmissionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubmissionService(opts ...option.RequestOption) (r *SubmissionService) {
	r = &SubmissionService{}
	r.Options = opts
	return
}

// Returns information for submissions made to reclassify emails. Shows the status,
// outcome, and disposition changes for reclassification requests made by users or
// the security team. Useful for tracking false positive/negative reports.
func (r *SubmissionService) List(ctx context.Context, params SubmissionListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SubmissionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/submissions", params.AccountID)
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

// Returns information for submissions made to reclassify emails. Shows the status,
// outcome, and disposition changes for reclassification requests made by users or
// the security team. Useful for tracking false positive/negative reports.
func (r *SubmissionService) ListAutoPaging(ctx context.Context, params SubmissionListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SubmissionListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type SubmissionListResponse struct {
	// When the submission was requested (UTC).
	RequestedAt           time.Time                                 `json:"requested_at" api:"required" format:"date-time"`
	SubmissionID          string                                    `json:"submission_id" api:"required"`
	CustomerStatus        SubmissionListResponseCustomerStatus      `json:"customer_status" api:"nullable"`
	EscalatedAs           SubmissionListResponseEscalatedAs         `json:"escalated_as" api:"nullable"`
	EscalatedAt           time.Time                                 `json:"escalated_at" api:"nullable" format:"date-time"`
	EscalatedBy           string                                    `json:"escalated_by" api:"nullable"`
	EscalatedSubmissionID string                                    `json:"escalated_submission_id" api:"nullable"`
	OriginalDisposition   SubmissionListResponseOriginalDisposition `json:"original_disposition" api:"nullable"`
	OriginalEdfHash       string                                    `json:"original_edf_hash" api:"nullable"`
	// The postfix ID of the original message that was submitted
	OriginalPostfixID    string                                     `json:"original_postfix_id" api:"nullable"`
	Outcome              string                                     `json:"outcome" api:"nullable"`
	OutcomeDisposition   SubmissionListResponseOutcomeDisposition   `json:"outcome_disposition" api:"nullable"`
	RequestedBy          string                                     `json:"requested_by" api:"nullable"`
	RequestedDisposition SubmissionListResponseRequestedDisposition `json:"requested_disposition" api:"nullable"`
	// Deprecated, use `requested_at` instead
	//
	// Deprecated: deprecated
	RequestedTs string `json:"requested_ts"`
	Status      string `json:"status" api:"nullable"`
	Subject     string `json:"subject" api:"nullable"`
	// Whether the submission was created by a team member or an end user.
	Type SubmissionListResponseType `json:"type" api:"nullable"`
	JSON submissionListResponseJSON `json:"-"`
}

// submissionListResponseJSON contains the JSON metadata for the struct
// [SubmissionListResponse]
type submissionListResponseJSON struct {
	RequestedAt           apijson.Field
	SubmissionID          apijson.Field
	CustomerStatus        apijson.Field
	EscalatedAs           apijson.Field
	EscalatedAt           apijson.Field
	EscalatedBy           apijson.Field
	EscalatedSubmissionID apijson.Field
	OriginalDisposition   apijson.Field
	OriginalEdfHash       apijson.Field
	OriginalPostfixID     apijson.Field
	Outcome               apijson.Field
	OutcomeDisposition    apijson.Field
	RequestedBy           apijson.Field
	RequestedDisposition  apijson.Field
	RequestedTs           apijson.Field
	Status                apijson.Field
	Subject               apijson.Field
	Type                  apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SubmissionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submissionListResponseJSON) RawJSON() string {
	return r.raw
}

type SubmissionListResponseCustomerStatus string

const (
	SubmissionListResponseCustomerStatusEscalated  SubmissionListResponseCustomerStatus = "escalated"
	SubmissionListResponseCustomerStatusReviewed   SubmissionListResponseCustomerStatus = "reviewed"
	SubmissionListResponseCustomerStatusUnreviewed SubmissionListResponseCustomerStatus = "unreviewed"
)

func (r SubmissionListResponseCustomerStatus) IsKnown() bool {
	switch r {
	case SubmissionListResponseCustomerStatusEscalated, SubmissionListResponseCustomerStatusReviewed, SubmissionListResponseCustomerStatusUnreviewed:
		return true
	}
	return false
}

type SubmissionListResponseEscalatedAs string

const (
	SubmissionListResponseEscalatedAsMalicious  SubmissionListResponseEscalatedAs = "MALICIOUS"
	SubmissionListResponseEscalatedAsSuspicious SubmissionListResponseEscalatedAs = "SUSPICIOUS"
	SubmissionListResponseEscalatedAsSpoof      SubmissionListResponseEscalatedAs = "SPOOF"
	SubmissionListResponseEscalatedAsSpam       SubmissionListResponseEscalatedAs = "SPAM"
	SubmissionListResponseEscalatedAsBulk       SubmissionListResponseEscalatedAs = "BULK"
	SubmissionListResponseEscalatedAsNone       SubmissionListResponseEscalatedAs = "NONE"
)

func (r SubmissionListResponseEscalatedAs) IsKnown() bool {
	switch r {
	case SubmissionListResponseEscalatedAsMalicious, SubmissionListResponseEscalatedAsSuspicious, SubmissionListResponseEscalatedAsSpoof, SubmissionListResponseEscalatedAsSpam, SubmissionListResponseEscalatedAsBulk, SubmissionListResponseEscalatedAsNone:
		return true
	}
	return false
}

type SubmissionListResponseOriginalDisposition string

const (
	SubmissionListResponseOriginalDispositionMalicious  SubmissionListResponseOriginalDisposition = "MALICIOUS"
	SubmissionListResponseOriginalDispositionSuspicious SubmissionListResponseOriginalDisposition = "SUSPICIOUS"
	SubmissionListResponseOriginalDispositionSpoof      SubmissionListResponseOriginalDisposition = "SPOOF"
	SubmissionListResponseOriginalDispositionSpam       SubmissionListResponseOriginalDisposition = "SPAM"
	SubmissionListResponseOriginalDispositionBulk       SubmissionListResponseOriginalDisposition = "BULK"
	SubmissionListResponseOriginalDispositionNone       SubmissionListResponseOriginalDisposition = "NONE"
)

func (r SubmissionListResponseOriginalDisposition) IsKnown() bool {
	switch r {
	case SubmissionListResponseOriginalDispositionMalicious, SubmissionListResponseOriginalDispositionSuspicious, SubmissionListResponseOriginalDispositionSpoof, SubmissionListResponseOriginalDispositionSpam, SubmissionListResponseOriginalDispositionBulk, SubmissionListResponseOriginalDispositionNone:
		return true
	}
	return false
}

type SubmissionListResponseOutcomeDisposition string

const (
	SubmissionListResponseOutcomeDispositionMalicious  SubmissionListResponseOutcomeDisposition = "MALICIOUS"
	SubmissionListResponseOutcomeDispositionSuspicious SubmissionListResponseOutcomeDisposition = "SUSPICIOUS"
	SubmissionListResponseOutcomeDispositionSpoof      SubmissionListResponseOutcomeDisposition = "SPOOF"
	SubmissionListResponseOutcomeDispositionSpam       SubmissionListResponseOutcomeDisposition = "SPAM"
	SubmissionListResponseOutcomeDispositionBulk       SubmissionListResponseOutcomeDisposition = "BULK"
	SubmissionListResponseOutcomeDispositionNone       SubmissionListResponseOutcomeDisposition = "NONE"
)

func (r SubmissionListResponseOutcomeDisposition) IsKnown() bool {
	switch r {
	case SubmissionListResponseOutcomeDispositionMalicious, SubmissionListResponseOutcomeDispositionSuspicious, SubmissionListResponseOutcomeDispositionSpoof, SubmissionListResponseOutcomeDispositionSpam, SubmissionListResponseOutcomeDispositionBulk, SubmissionListResponseOutcomeDispositionNone:
		return true
	}
	return false
}

type SubmissionListResponseRequestedDisposition string

const (
	SubmissionListResponseRequestedDispositionMalicious  SubmissionListResponseRequestedDisposition = "MALICIOUS"
	SubmissionListResponseRequestedDispositionSuspicious SubmissionListResponseRequestedDisposition = "SUSPICIOUS"
	SubmissionListResponseRequestedDispositionSpoof      SubmissionListResponseRequestedDisposition = "SPOOF"
	SubmissionListResponseRequestedDispositionSpam       SubmissionListResponseRequestedDisposition = "SPAM"
	SubmissionListResponseRequestedDispositionBulk       SubmissionListResponseRequestedDisposition = "BULK"
	SubmissionListResponseRequestedDispositionNone       SubmissionListResponseRequestedDisposition = "NONE"
)

func (r SubmissionListResponseRequestedDisposition) IsKnown() bool {
	switch r {
	case SubmissionListResponseRequestedDispositionMalicious, SubmissionListResponseRequestedDispositionSuspicious, SubmissionListResponseRequestedDispositionSpoof, SubmissionListResponseRequestedDispositionSpam, SubmissionListResponseRequestedDispositionBulk, SubmissionListResponseRequestedDispositionNone:
		return true
	}
	return false
}

// Whether the submission was created by a team member or an end user.
type SubmissionListResponseType string

const (
	SubmissionListResponseTypeTeam SubmissionListResponseType = "Team"
	SubmissionListResponseTypeUser SubmissionListResponseType = "User"
)

func (r SubmissionListResponseType) IsKnown() bool {
	switch r {
	case SubmissionListResponseTypeTeam, SubmissionListResponseTypeUser:
		return true
	}
	return false
}

type SubmissionListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The end of the search date range. Defaults to `now`.
	End param.Field[time.Time] `query:"end" format:"date-time"`
	// When true, return only submissions that were escalated by an end user (vs. by
	// the security team). When false, return only submissions that were not escalated
	// by an end user. When omitted, no filter is applied.
	EscalatedFromUser   param.Field[bool]                                    `query:"escalated_from_user"`
	OriginalDisposition param.Field[SubmissionListParamsOriginalDisposition] `query:"original_disposition"`
	OutcomeDisposition  param.Field[SubmissionListParamsOutcomeDisposition]  `query:"outcome_disposition"`
	// Current page within paginated list of results.
	Page param.Field[int64] `query:"page"`
	// The number of results per page. Maximum value is 1000.
	PerPage              param.Field[int64]                                    `query:"per_page"`
	Query                param.Field[string]                                   `query:"query"`
	RequestedDisposition param.Field[SubmissionListParamsRequestedDisposition] `query:"requested_disposition"`
	// The beginning of the search date range. Defaults to `now - 30 days`.
	Start        param.Field[time.Time]                `query:"start" format:"date-time"`
	Status       param.Field[string]                   `query:"status"`
	SubmissionID param.Field[string]                   `query:"submission_id"`
	Type         param.Field[SubmissionListParamsType] `query:"type"`
}

// URLQuery serializes [SubmissionListParams]'s query parameters as `url.Values`.
func (r SubmissionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SubmissionListParamsOriginalDisposition string

const (
	SubmissionListParamsOriginalDispositionMalicious  SubmissionListParamsOriginalDisposition = "MALICIOUS"
	SubmissionListParamsOriginalDispositionSuspicious SubmissionListParamsOriginalDisposition = "SUSPICIOUS"
	SubmissionListParamsOriginalDispositionSpoof      SubmissionListParamsOriginalDisposition = "SPOOF"
	SubmissionListParamsOriginalDispositionSpam       SubmissionListParamsOriginalDisposition = "SPAM"
	SubmissionListParamsOriginalDispositionBulk       SubmissionListParamsOriginalDisposition = "BULK"
	SubmissionListParamsOriginalDispositionNone       SubmissionListParamsOriginalDisposition = "NONE"
)

func (r SubmissionListParamsOriginalDisposition) IsKnown() bool {
	switch r {
	case SubmissionListParamsOriginalDispositionMalicious, SubmissionListParamsOriginalDispositionSuspicious, SubmissionListParamsOriginalDispositionSpoof, SubmissionListParamsOriginalDispositionSpam, SubmissionListParamsOriginalDispositionBulk, SubmissionListParamsOriginalDispositionNone:
		return true
	}
	return false
}

type SubmissionListParamsOutcomeDisposition string

const (
	SubmissionListParamsOutcomeDispositionMalicious  SubmissionListParamsOutcomeDisposition = "MALICIOUS"
	SubmissionListParamsOutcomeDispositionSuspicious SubmissionListParamsOutcomeDisposition = "SUSPICIOUS"
	SubmissionListParamsOutcomeDispositionSpoof      SubmissionListParamsOutcomeDisposition = "SPOOF"
	SubmissionListParamsOutcomeDispositionSpam       SubmissionListParamsOutcomeDisposition = "SPAM"
	SubmissionListParamsOutcomeDispositionBulk       SubmissionListParamsOutcomeDisposition = "BULK"
	SubmissionListParamsOutcomeDispositionNone       SubmissionListParamsOutcomeDisposition = "NONE"
)

func (r SubmissionListParamsOutcomeDisposition) IsKnown() bool {
	switch r {
	case SubmissionListParamsOutcomeDispositionMalicious, SubmissionListParamsOutcomeDispositionSuspicious, SubmissionListParamsOutcomeDispositionSpoof, SubmissionListParamsOutcomeDispositionSpam, SubmissionListParamsOutcomeDispositionBulk, SubmissionListParamsOutcomeDispositionNone:
		return true
	}
	return false
}

type SubmissionListParamsRequestedDisposition string

const (
	SubmissionListParamsRequestedDispositionMalicious  SubmissionListParamsRequestedDisposition = "MALICIOUS"
	SubmissionListParamsRequestedDispositionSuspicious SubmissionListParamsRequestedDisposition = "SUSPICIOUS"
	SubmissionListParamsRequestedDispositionSpoof      SubmissionListParamsRequestedDisposition = "SPOOF"
	SubmissionListParamsRequestedDispositionSpam       SubmissionListParamsRequestedDisposition = "SPAM"
	SubmissionListParamsRequestedDispositionBulk       SubmissionListParamsRequestedDisposition = "BULK"
	SubmissionListParamsRequestedDispositionNone       SubmissionListParamsRequestedDisposition = "NONE"
)

func (r SubmissionListParamsRequestedDisposition) IsKnown() bool {
	switch r {
	case SubmissionListParamsRequestedDispositionMalicious, SubmissionListParamsRequestedDispositionSuspicious, SubmissionListParamsRequestedDispositionSpoof, SubmissionListParamsRequestedDispositionSpam, SubmissionListParamsRequestedDispositionBulk, SubmissionListParamsRequestedDispositionNone:
		return true
	}
	return false
}

type SubmissionListParamsType string

const (
	SubmissionListParamsTypeTeam SubmissionListParamsType = "TEAM"
	SubmissionListParamsTypeUser SubmissionListParamsType = "USER"
)

func (r SubmissionListParamsType) IsKnown() bool {
	switch r {
	case SubmissionListParamsTypeTeam, SubmissionListParamsTypeUser:
		return true
	}
	return false
}

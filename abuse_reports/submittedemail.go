// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package abuse_reports

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
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// SubmittedEmailService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSubmittedEmailService] method instead.
type SubmittedEmailService struct {
	Options []option.RequestOption
}

// NewSubmittedEmailService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSubmittedEmailService(opts ...option.RequestOption) (r *SubmittedEmailService) {
	r = &SubmittedEmailService{}
	r.Options = opts
	return
}

// List successful emails sent to the submitter of a report submitted by the
// account. Does not include emails sent to customers or hosts.
func (r *SubmittedEmailService) List(ctx context.Context, reportID string, params SubmittedEmailListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[SubmittedEmailListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if reportID == "" {
		err = errors.New("missing required report_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/abuse-reports/submitted/%s/emails", params.AccountID, reportID)
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

// List successful emails sent to the submitter of a report submitted by the
// account. Does not include emails sent to customers or hosts.
func (r *SubmittedEmailService) ListAutoPaging(ctx context.Context, reportID string, params SubmittedEmailListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[SubmittedEmailListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, reportID, params, opts...))
}

type SubmittedEmailListResponse struct {
	Emails []SubmittedEmailListResponseEmail `json:"emails" api:"required"`
	JSON   submittedEmailListResponseJSON    `json:"-"`
}

// submittedEmailListResponseJSON contains the JSON metadata for the struct
// [SubmittedEmailListResponse]
type submittedEmailListResponseJSON struct {
	Emails      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubmittedEmailListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submittedEmailListResponseJSON) RawJSON() string {
	return r.raw
}

// An email sent to the customer for an abuse report.
type SubmittedEmailListResponseEmail struct {
	// Unique identifier of the email.
	ID string `json:"id" api:"required"`
	// Body content of the email.
	Body string `json:"body" api:"required"`
	// Email address of the recipient.
	Recipient string `json:"recipient" api:"required"`
	// When the email was sent. Time in RFC 3339 format
	// (https://www.rfc-editor.org/rfc/rfc3339.html)
	SentAt string `json:"sent_at" api:"required"`
	// Subject line of the email.
	Subject string                              `json:"subject" api:"required"`
	JSON    submittedEmailListResponseEmailJSON `json:"-"`
}

// submittedEmailListResponseEmailJSON contains the JSON metadata for the struct
// [SubmittedEmailListResponseEmail]
type submittedEmailListResponseEmailJSON struct {
	ID          apijson.Field
	Body        apijson.Field
	Recipient   apijson.Field
	SentAt      apijson.Field
	Subject     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubmittedEmailListResponseEmail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r submittedEmailListResponseEmailJSON) RawJSON() string {
	return r.raw
}

type SubmittedEmailListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Page number to retrieve (default 1).
	Page param.Field[int64] `query:"page"`
	// Number of emails per page (default 20, max 100).
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [SubmittedEmailListParams]'s query parameters as
// `url.Values`.
func (r SubmittedEmailListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

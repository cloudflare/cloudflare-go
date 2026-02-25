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

// PhishguardReportService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhishguardReportService] method instead.
type PhishguardReportService struct {
	Options []option.RequestOption
}

// NewPhishguardReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPhishguardReportService(opts ...option.RequestOption) (r *PhishguardReportService) {
	r = &PhishguardReportService{}
	r.Options = opts
	return
}

// Get `PhishGuard` reports
func (r *PhishguardReportService) List(ctx context.Context, params PhishguardReportListParams, opts ...option.RequestOption) (res *pagination.SinglePage[PhishguardReportListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email-security/phishguard/reports", params.AccountID)
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

// Get `PhishGuard` reports
func (r *PhishguardReportService) ListAutoPaging(ctx context.Context, params PhishguardReportListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PhishguardReportListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

type PhishguardReportListResponse struct {
	ID          int64                                   `json:"id,required"`
	Content     string                                  `json:"content,required"`
	CreatedAt   time.Time                               `json:"created_at,required" format:"date-time"`
	Disposition PhishguardReportListResponseDisposition `json:"disposition,required"`
	Fields      PhishguardReportListResponseFields      `json:"fields,required"`
	Priority    string                                  `json:"priority,required"`
	Title       string                                  `json:"title,required"`
	Ts          time.Time                               `json:"ts,required" format:"date-time"`
	UpdatedAt   time.Time                               `json:"updated_at,required" format:"date-time"`
	Tags        []PhishguardReportListResponseTag       `json:"tags,nullable"`
	JSON        phishguardReportListResponseJSON        `json:"-"`
}

// phishguardReportListResponseJSON contains the JSON metadata for the struct
// [PhishguardReportListResponse]
type phishguardReportListResponseJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedAt   apijson.Field
	Disposition apijson.Field
	Fields      apijson.Field
	Priority    apijson.Field
	Title       apijson.Field
	Ts          apijson.Field
	UpdatedAt   apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhishguardReportListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phishguardReportListResponseJSON) RawJSON() string {
	return r.raw
}

type PhishguardReportListResponseDisposition string

const (
	PhishguardReportListResponseDispositionMalicious    PhishguardReportListResponseDisposition = "MALICIOUS"
	PhishguardReportListResponseDispositionMaliciousBec PhishguardReportListResponseDisposition = "MALICIOUS-BEC"
	PhishguardReportListResponseDispositionSuspicious   PhishguardReportListResponseDisposition = "SUSPICIOUS"
	PhishguardReportListResponseDispositionSpoof        PhishguardReportListResponseDisposition = "SPOOF"
	PhishguardReportListResponseDispositionSpam         PhishguardReportListResponseDisposition = "SPAM"
	PhishguardReportListResponseDispositionBulk         PhishguardReportListResponseDisposition = "BULK"
	PhishguardReportListResponseDispositionEncrypted    PhishguardReportListResponseDisposition = "ENCRYPTED"
	PhishguardReportListResponseDispositionExternal     PhishguardReportListResponseDisposition = "EXTERNAL"
	PhishguardReportListResponseDispositionUnknown      PhishguardReportListResponseDisposition = "UNKNOWN"
	PhishguardReportListResponseDispositionNone         PhishguardReportListResponseDisposition = "NONE"
)

func (r PhishguardReportListResponseDisposition) IsKnown() bool {
	switch r {
	case PhishguardReportListResponseDispositionMalicious, PhishguardReportListResponseDispositionMaliciousBec, PhishguardReportListResponseDispositionSuspicious, PhishguardReportListResponseDispositionSpoof, PhishguardReportListResponseDispositionSpam, PhishguardReportListResponseDispositionBulk, PhishguardReportListResponseDispositionEncrypted, PhishguardReportListResponseDispositionExternal, PhishguardReportListResponseDispositionUnknown, PhishguardReportListResponseDispositionNone:
		return true
	}
	return false
}

type PhishguardReportListResponseFields struct {
	To        []string                               `json:"to,required"`
	Ts        time.Time                              `json:"ts,required" format:"date-time"`
	From      string                                 `json:"from,nullable"`
	PostfixID string                                 `json:"postfix_id,nullable"`
	JSON      phishguardReportListResponseFieldsJSON `json:"-"`
}

// phishguardReportListResponseFieldsJSON contains the JSON metadata for the struct
// [PhishguardReportListResponseFields]
type phishguardReportListResponseFieldsJSON struct {
	To          apijson.Field
	Ts          apijson.Field
	From        apijson.Field
	PostfixID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhishguardReportListResponseFields) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phishguardReportListResponseFieldsJSON) RawJSON() string {
	return r.raw
}

type PhishguardReportListResponseTag struct {
	Category string                              `json:"category,required"`
	Value    string                              `json:"value,required"`
	JSON     phishguardReportListResponseTagJSON `json:"-"`
}

// phishguardReportListResponseTagJSON contains the JSON metadata for the struct
// [PhishguardReportListResponseTag]
type phishguardReportListResponseTagJSON struct {
	Category    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhishguardReportListResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phishguardReportListResponseTagJSON) RawJSON() string {
	return r.raw
}

type PhishguardReportListParams struct {
	// Account Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The end of the search date range (RFC3339 format).
	End      param.Field[time.Time] `query:"end" format:"date-time"`
	FromDate param.Field[time.Time] `query:"from_date" format:"date"`
	// The beginning of the search date range (RFC3339 format).
	Start  param.Field[time.Time] `query:"start" format:"date-time"`
	ToDate param.Field[time.Time] `query:"to_date" format:"date"`
}

// URLQuery serializes [PhishguardReportListParams]'s query parameters as
// `url.Values`.
func (r PhishguardReportListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

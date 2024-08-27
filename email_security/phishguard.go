// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
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

// PhishguardService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhishguardService] method instead.
type PhishguardService struct {
	Options []option.RequestOption
}

// NewPhishguardService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPhishguardService(opts ...option.RequestOption) (r *PhishguardService) {
	r = &PhishguardService{}
	r.Options = opts
	return
}

// Get PhishGuard reports
func (r *PhishguardService) List(ctx context.Context, params PhishguardListParams, opts ...option.RequestOption) (res *pagination.SinglePage[PhishguardListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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

// Get PhishGuard reports
func (r *PhishguardService) ListAutoPaging(ctx context.Context, params PhishguardListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PhishguardListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

type PhishguardListResponse struct {
	ID          int64                             `json:"id,required"`
	Content     string                            `json:"content,required"`
	Disposition PhishguardListResponseDisposition `json:"disposition,required"`
	Fields      PhishguardListResponseFields      `json:"fields,required"`
	Priority    string                            `json:"priority,required"`
	Title       string                            `json:"title,required"`
	Ts          time.Time                         `json:"ts,required" format:"date-time"`
	Tags        []PhishguardListResponseTag       `json:"tags,nullable"`
	JSON        phishguardListResponseJSON        `json:"-"`
}

// phishguardListResponseJSON contains the JSON metadata for the struct
// [PhishguardListResponse]
type phishguardListResponseJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	Disposition apijson.Field
	Fields      apijson.Field
	Priority    apijson.Field
	Title       apijson.Field
	Ts          apijson.Field
	Tags        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhishguardListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phishguardListResponseJSON) RawJSON() string {
	return r.raw
}

type PhishguardListResponseDisposition string

const (
	PhishguardListResponseDispositionMalicious    PhishguardListResponseDisposition = "MALICIOUS"
	PhishguardListResponseDispositionMaliciousBec PhishguardListResponseDisposition = "MALICIOUS-BEC"
	PhishguardListResponseDispositionSuspicious   PhishguardListResponseDisposition = "SUSPICIOUS"
	PhishguardListResponseDispositionSpoof        PhishguardListResponseDisposition = "SPOOF"
	PhishguardListResponseDispositionSpam         PhishguardListResponseDisposition = "SPAM"
	PhishguardListResponseDispositionBulk         PhishguardListResponseDisposition = "BULK"
	PhishguardListResponseDispositionEncrypted    PhishguardListResponseDisposition = "ENCRYPTED"
	PhishguardListResponseDispositionExternal     PhishguardListResponseDisposition = "EXTERNAL"
	PhishguardListResponseDispositionUnknown      PhishguardListResponseDisposition = "UNKNOWN"
	PhishguardListResponseDispositionNone         PhishguardListResponseDisposition = "NONE"
)

func (r PhishguardListResponseDisposition) IsKnown() bool {
	switch r {
	case PhishguardListResponseDispositionMalicious, PhishguardListResponseDispositionMaliciousBec, PhishguardListResponseDispositionSuspicious, PhishguardListResponseDispositionSpoof, PhishguardListResponseDispositionSpam, PhishguardListResponseDispositionBulk, PhishguardListResponseDispositionEncrypted, PhishguardListResponseDispositionExternal, PhishguardListResponseDispositionUnknown, PhishguardListResponseDispositionNone:
		return true
	}
	return false
}

type PhishguardListResponseFields struct {
	PostfixID string                           `json:"postfix_id,required"`
	To        []string                         `json:"to,required"`
	Ts        time.Time                        `json:"ts,required" format:"date-time"`
	From      string                           `json:"from,nullable"`
	JSON      phishguardListResponseFieldsJSON `json:"-"`
}

// phishguardListResponseFieldsJSON contains the JSON metadata for the struct
// [PhishguardListResponseFields]
type phishguardListResponseFieldsJSON struct {
	PostfixID   apijson.Field
	To          apijson.Field
	Ts          apijson.Field
	From        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhishguardListResponseFields) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phishguardListResponseFieldsJSON) RawJSON() string {
	return r.raw
}

type PhishguardListResponseTag struct {
	Category string                        `json:"category,required"`
	Value    string                        `json:"value,required"`
	JSON     phishguardListResponseTagJSON `json:"-"`
}

// phishguardListResponseTagJSON contains the JSON metadata for the struct
// [PhishguardListResponseTag]
type phishguardListResponseTagJSON struct {
	Category    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhishguardListResponseTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phishguardListResponseTagJSON) RawJSON() string {
	return r.raw
}

type PhishguardListParams struct {
	// Account Identifier
	AccountID param.Field[string]    `path:"account_id,required"`
	FromDate  param.Field[time.Time] `query:"from_date,required" format:"date"`
	ToDate    param.Field[time.Time] `query:"to_date,required" format:"date"`
}

// URLQuery serializes [PhishguardListParams]'s query parameters as `url.Values`.
func (r PhishguardListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

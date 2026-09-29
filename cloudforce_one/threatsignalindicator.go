// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

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
)

// ThreatSignalIndicatorService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatSignalIndicatorService] method instead.
type ThreatSignalIndicatorService struct {
	Options []option.RequestOption
}

// NewThreatSignalIndicatorService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatSignalIndicatorService(opts ...option.RequestOption) (r *ThreatSignalIndicatorService) {
	r = &ThreatSignalIndicatorService{}
	r.Options = opts
	return
}

// Lists indicators of compromise extracted from the account's Threat Signals
// articles.
func (r *ThreatSignalIndicatorService) List(ctx context.Context, params ThreatSignalIndicatorListParams, opts ...option.RequestOption) (res *ThreatSignalIndicatorListResponse, err error) {
	var env ThreatSignalIndicatorListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/indicators", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ThreatSignalIndicatorListResponse struct {
	Indicators []ThreatSignalIndicatorListResponseIndicator `json:"indicators" api:"required"`
	Pagination ThreatSignalIndicatorListResponsePagination  `json:"pagination" api:"required"`
	JSON       threatSignalIndicatorListResponseJSON        `json:"-"`
}

// threatSignalIndicatorListResponseJSON contains the JSON metadata for the struct
// [ThreatSignalIndicatorListResponse]
type threatSignalIndicatorListResponseJSON struct {
	Indicators  apijson.Field
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalIndicatorListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalIndicatorListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalIndicatorListResponseIndicator struct {
	ID           string `json:"id" api:"required" format:"uuid"`
	ArticleID    string `json:"article_id" api:"required" format:"uuid"`
	ArticleTitle string `json:"article_title" api:"required,nullable"`
	// Threat Events dataset identifier for navigating from this indicator. Null when
	// the account feeds dataset mapping is unavailable.
	DatasetID       string                                         `json:"dataset_id" api:"required,nullable"`
	FeedDisplayName string                                         `json:"feed_display_name" api:"required,nullable"`
	FeedID          string                                         `json:"feed_id" api:"required" format:"uuid"`
	Type            string                                         `json:"type" api:"required"`
	Value           string                                         `json:"value" api:"required"`
	JSON            threatSignalIndicatorListResponseIndicatorJSON `json:"-"`
}

// threatSignalIndicatorListResponseIndicatorJSON contains the JSON metadata for
// the struct [ThreatSignalIndicatorListResponseIndicator]
type threatSignalIndicatorListResponseIndicatorJSON struct {
	ID              apijson.Field
	ArticleID       apijson.Field
	ArticleTitle    apijson.Field
	DatasetID       apijson.Field
	FeedDisplayName apijson.Field
	FeedID          apijson.Field
	Type            apijson.Field
	Value           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ThreatSignalIndicatorListResponseIndicator) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalIndicatorListResponseIndicatorJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalIndicatorListResponsePagination struct {
	Count   int64  `json:"count" api:"required"`
	Cursor  string `json:"cursor" api:"required,nullable"`
	HasMore bool   `json:"has_more" api:"required"`
	// Ordinal of this cursor page; not a total-results offset.
	Page              int64                                           `json:"page" api:"required"`
	PerPage           int64                                           `json:"per_page" api:"required"`
	TotalCount        int64                                           `json:"total_count" api:"required,nullable"`
	TotalCountIsExact bool                                            `json:"total_count_is_exact" api:"required"`
	JSON              threatSignalIndicatorListResponsePaginationJSON `json:"-"`
}

// threatSignalIndicatorListResponsePaginationJSON contains the JSON metadata for
// the struct [ThreatSignalIndicatorListResponsePagination]
type threatSignalIndicatorListResponsePaginationJSON struct {
	Count             apijson.Field
	Cursor            apijson.Field
	HasMore           apijson.Field
	Page              apijson.Field
	PerPage           apijson.Field
	TotalCount        apijson.Field
	TotalCountIsExact apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ThreatSignalIndicatorListResponsePagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalIndicatorListResponsePaginationJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalIndicatorListParams struct {
	AccountID    param.Field[string] `path:"account_id" api:"required"`
	ArticleID    param.Field[string] `query:"article_id" format:"uuid"`
	Cursor       param.Field[string] `query:"cursor"`
	FeedID       param.Field[string] `query:"feed_id" format:"uuid"`
	IncludeTotal param.Field[bool]   `query:"include_total"`
	PerPage      param.Field[int64]  `query:"per_page"`
	// NFC-normalized and trimmed, case-insensitive literal substring search of
	// indicator values. Requires 3–500 Unicode code points; the upper code-point bound
	// is described here because OpenAPI string length cannot precisely express it
	// without imposing UTF-16 semantics.
	Search param.Field[string] `query:"search"`
	Sort   param.Field[string] `query:"sort"`
}

// URLQuery serializes [ThreatSignalIndicatorListParams]'s query parameters as
// `url.Values`.
func (r ThreatSignalIndicatorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatSignalIndicatorListResponseEnvelope struct {
	Errors  []ThreatSignalIndicatorListResponseEnvelopeErrors `json:"errors" api:"required"`
	Result  ThreatSignalIndicatorListResponse                 `json:"result" api:"required"`
	Success ThreatSignalIndicatorListResponseEnvelopeSuccess  `json:"success" api:"required"`
	JSON    threatSignalIndicatorListResponseEnvelopeJSON     `json:"-"`
}

// threatSignalIndicatorListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalIndicatorListResponseEnvelope]
type threatSignalIndicatorListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalIndicatorListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalIndicatorListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalIndicatorListResponseEnvelopeErrors struct {
	Message string                                              `json:"message" api:"required"`
	JSON    threatSignalIndicatorListResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalIndicatorListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ThreatSignalIndicatorListResponseEnvelopeErrors]
type threatSignalIndicatorListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalIndicatorListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalIndicatorListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalIndicatorListResponseEnvelopeSuccess bool

const (
	ThreatSignalIndicatorListResponseEnvelopeSuccessTrue ThreatSignalIndicatorListResponseEnvelopeSuccess = true
)

func (r ThreatSignalIndicatorListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalIndicatorListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

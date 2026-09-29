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

// ThreatSignalSearchService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatSignalSearchService] method instead.
type ThreatSignalSearchService struct {
	Options []option.RequestOption
}

// NewThreatSignalSearchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatSignalSearchService(opts ...option.RequestOption) (r *ThreatSignalSearchService) {
	r = &ThreatSignalSearchService{}
	r.Options = opts
	return
}

// Searches the account's Threat Signals articles using keyword and semantic
// retrieval.
func (r *ThreatSignalSearchService) Search(ctx context.Context, params ThreatSignalSearchSearchParams, opts ...option.RequestOption) (res *ThreatSignalSearchSearchResponse, err error) {
	var env ThreatSignalSearchSearchResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/search", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ThreatSignalSearchSearchResponse struct {
	// Number of unique article candidates returned in this response. Equal to
	// results.length.
	Count   int64                                    `json:"count" api:"required"`
	Results []ThreatSignalSearchSearchResponseResult `json:"results" api:"required"`
	JSON    threatSignalSearchSearchResponseJSON     `json:"-"`
}

// threatSignalSearchSearchResponseJSON contains the JSON metadata for the struct
// [ThreatSignalSearchSearchResponse]
type threatSignalSearchSearchResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSearchSearchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSearchSearchResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSearchSearchResponseResult struct {
	ArticleID string                                     `json:"article_id" api:"required" format:"uuid"`
	DatasetID string                                     `json:"dataset_id" api:"required,nullable" format:"uuid"`
	EventID   string                                     `json:"event_id" api:"required,nullable" format:"uuid"`
	FeedID    string                                     `json:"feed_id" api:"required" format:"uuid"`
	Score     float64                                    `json:"score" api:"required"`
	Text      string                                     `json:"text" api:"required"`
	JSON      threatSignalSearchSearchResponseResultJSON `json:"-"`
}

// threatSignalSearchSearchResponseResultJSON contains the JSON metadata for the
// struct [ThreatSignalSearchSearchResponseResult]
type threatSignalSearchSearchResponseResultJSON struct {
	ArticleID   apijson.Field
	DatasetID   apijson.Field
	EventID     apijson.Field
	FeedID      apijson.Field
	Score       apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSearchSearchResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSearchSearchResponseResultJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSearchSearchParams struct {
	AccountID  param.Field[string]                                   `path:"account_id" api:"required"`
	Query      param.Field[string]                                   `query:"query" api:"required"`
	FeedID     param.Field[string]                                   `query:"feed_id" format:"uuid"`
	MaxResults param.Field[ThreatSignalSearchSearchParamsMaxResults] `query:"max_results"`
}

// URLQuery serializes [ThreatSignalSearchSearchParams]'s query parameters as
// `url.Values`.
func (r ThreatSignalSearchSearchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatSignalSearchSearchParamsMaxResults string

const (
	ThreatSignalSearchSearchParamsMaxResultsEmpty ThreatSignalSearchSearchParamsMaxResults = ""
)

func (r ThreatSignalSearchSearchParamsMaxResults) IsKnown() bool {
	switch r {
	case ThreatSignalSearchSearchParamsMaxResultsEmpty:
		return true
	}
	return false
}

type ThreatSignalSearchSearchResponseEnvelope struct {
	Errors  []ThreatSignalSearchSearchResponseEnvelopeErrors `json:"errors" api:"required"`
	Result  ThreatSignalSearchSearchResponse                 `json:"result" api:"required"`
	Success ThreatSignalSearchSearchResponseEnvelopeSuccess  `json:"success" api:"required"`
	JSON    threatSignalSearchSearchResponseEnvelopeJSON     `json:"-"`
}

// threatSignalSearchSearchResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalSearchSearchResponseEnvelope]
type threatSignalSearchSearchResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSearchSearchResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSearchSearchResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSearchSearchResponseEnvelopeErrors struct {
	Message string                                             `json:"message" api:"required"`
	JSON    threatSignalSearchSearchResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalSearchSearchResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ThreatSignalSearchSearchResponseEnvelopeErrors]
type threatSignalSearchSearchResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSearchSearchResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSearchSearchResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSearchSearchResponseEnvelopeSuccess bool

const (
	ThreatSignalSearchSearchResponseEnvelopeSuccessTrue ThreatSignalSearchSearchResponseEnvelopeSuccess = true
)

func (r ThreatSignalSearchSearchResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalSearchSearchResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

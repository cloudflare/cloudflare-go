// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_tagging

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

// SummaryService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSummaryService] method instead.
type SummaryService struct {
	Options []option.RequestOption
}

// NewSummaryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSummaryService(opts ...option.RequestOption) (r *SummaryService) {
	r = &SummaryService{}
	r.Options = opts
	return
}

// Lists all distinct tag keys and their distinct values across resources in an
// account.
func (r *SummaryService) Get(ctx context.Context, params SummaryGetParams, opts ...option.RequestOption) (res *[]SummaryGetResponse, err error) {
	var env SummaryGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/tags/summary", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type SummaryGetResponse struct {
	// A tag key.
	Key string `json:"key" api:"required"`
	// All distinct values for this tag key.
	Values []string               `json:"values" api:"required"`
	JSON   summaryGetResponseJSON `json:"-"`
}

// summaryGetResponseJSON contains the JSON metadata for the struct
// [SummaryGetResponse]
type summaryGetResponseJSON struct {
	Key         apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SummaryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r summaryGetResponseJSON) RawJSON() string {
	return r.raw
}

type SummaryGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Cursor for pagination.
	Cursor param.Field[string] `query:"cursor"`
}

// URLQuery serializes [SummaryGetParams]'s query parameters as `url.Values`.
func (r SummaryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SummaryGetResponseEnvelope struct {
	Errors   []SummaryGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []SummaryGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success SummaryGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// Contains an array of tag keys with their distinct values.
	Result     []SummaryGetResponse                 `json:"result"`
	ResultInfo SummaryGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       summaryGetResponseEnvelopeJSON       `json:"-"`
}

// summaryGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SummaryGetResponseEnvelope]
type summaryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SummaryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r summaryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SummaryGetResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code" api:"required"`
	Message          string                                 `json:"message" api:"required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           SummaryGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             summaryGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// summaryGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SummaryGetResponseEnvelopeErrors]
type summaryGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SummaryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r summaryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SummaryGetResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    summaryGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// summaryGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [SummaryGetResponseEnvelopeErrorsSource]
type summaryGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SummaryGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r summaryGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SummaryGetResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code" api:"required"`
	Message          string                                   `json:"message" api:"required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           SummaryGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             summaryGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// summaryGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SummaryGetResponseEnvelopeMessages]
type summaryGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SummaryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r summaryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SummaryGetResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    summaryGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// summaryGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [SummaryGetResponseEnvelopeMessagesSource]
type summaryGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SummaryGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r summaryGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SummaryGetResponseEnvelopeSuccess bool

const (
	SummaryGetResponseEnvelopeSuccessTrue SummaryGetResponseEnvelopeSuccess = true
)

func (r SummaryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SummaryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SummaryGetResponseEnvelopeResultInfo struct {
	// Indicates the number of results returned in the current page.
	Count int64 `json:"count"`
	// Provides a cursor for the next page of results. Include this value in the next
	// request to continue pagination.
	Cursor string                                   `json:"cursor" api:"nullable"`
	JSON   summaryGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// summaryGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [SummaryGetResponseEnvelopeResultInfo]
type summaryGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SummaryGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r summaryGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

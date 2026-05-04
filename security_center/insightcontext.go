// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_center

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// InsightContextService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInsightContextService] method instead.
type InsightContextService struct {
	Options []option.RequestOption
}

// NewInsightContextService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInsightContextService(opts ...option.RequestOption) (r *InsightContextService) {
	r = &InsightContextService{}
	r.Options = opts
	return
}

// Returns the full context payload for an insight. This endpoint is used for
// insights with large payloads that are not included inline in the list response.
func (r *InsightContextService) Get(ctx context.Context, issueID string, query InsightContextGetParams, opts ...option.RequestOption) (res *InsightContextGetResponse, err error) {
	var env InsightContextGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if issueID == "" {
		err = errors.New("missing required issue_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/security-center/insights/%s/context", query.AccountID, issueID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type InsightContextGetResponse map[string]interface{}

type InsightContextGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type InsightContextGetResponseEnvelope struct {
	Errors   []InsightContextGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []InsightContextGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success InsightContextGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  InsightContextGetResponse                `json:"result"`
	JSON    insightContextGetResponseEnvelopeJSON    `json:"-"`
}

// insightContextGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InsightContextGetResponseEnvelope]
type insightContextGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightContextGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightContextGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InsightContextGetResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           InsightContextGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             insightContextGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// insightContextGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [InsightContextGetResponseEnvelopeErrors]
type insightContextGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InsightContextGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightContextGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InsightContextGetResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    insightContextGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// insightContextGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [InsightContextGetResponseEnvelopeErrorsSource]
type insightContextGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightContextGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightContextGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InsightContextGetResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           InsightContextGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             insightContextGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// insightContextGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [InsightContextGetResponseEnvelopeMessages]
type insightContextGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InsightContextGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightContextGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InsightContextGetResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    insightContextGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// insightContextGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [InsightContextGetResponseEnvelopeMessagesSource]
type insightContextGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InsightContextGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r insightContextGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InsightContextGetResponseEnvelopeSuccess bool

const (
	InsightContextGetResponseEnvelopeSuccessTrue InsightContextGetResponseEnvelopeSuccess = true
)

func (r InsightContextGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InsightContextGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

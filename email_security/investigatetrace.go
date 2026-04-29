// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// InvestigateTraceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateTraceService] method instead.
type InvestigateTraceService struct {
	Options []option.RequestOption
}

// NewInvestigateTraceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestigateTraceService(opts ...option.RequestOption) (r *InvestigateTraceService) {
	r = &InvestigateTraceService{}
	r.Options = opts
	return
}

// Retrieves delivery and processing trace information for an email message. Shows
// the delivery path, retraction history, and move operations performed on the
// message. Useful for debugging delivery issues.
func (r *InvestigateTraceService) Get(ctx context.Context, investigateID string, query InvestigateTraceGetParams, opts ...option.RequestOption) (res *InvestigateTraceGetResponse, err error) {
	var env InvestigateTraceGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if investigateID == "" {
		err = errors.New("missing required investigate_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/trace", query.AccountID, investigateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type InvestigateTraceGetResponse struct {
	Inbound  InvestigateTraceGetResponseInbound  `json:"inbound" api:"required"`
	Outbound InvestigateTraceGetResponseOutbound `json:"outbound" api:"required"`
	JSON     investigateTraceGetResponseJSON     `json:"-"`
}

// investigateTraceGetResponseJSON contains the JSON metadata for the struct
// [InvestigateTraceGetResponse]
type investigateTraceGetResponseJSON struct {
	Inbound     apijson.Field
	Outbound    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceGetResponseInbound struct {
	Lines   []InvestigateTraceGetResponseInboundLine `json:"lines" api:"nullable"`
	Pending bool                                     `json:"pending" api:"nullable"`
	JSON    investigateTraceGetResponseInboundJSON   `json:"-"`
}

// investigateTraceGetResponseInboundJSON contains the JSON metadata for the struct
// [InvestigateTraceGetResponseInbound]
type investigateTraceGetResponseInboundJSON struct {
	Lines       apijson.Field
	Pending     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceGetResponseInbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseInboundJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceGetResponseInboundLine struct {
	// Line number in the trace log
	Lineno   int64     `json:"lineno"`
	LoggedAt time.Time `json:"logged_at" api:"nullable" format:"date-time"`
	Message  string    `json:"message"`
	// Deprecated, use `logged_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	Ts   string                                     `json:"ts"`
	JSON investigateTraceGetResponseInboundLineJSON `json:"-"`
}

// investigateTraceGetResponseInboundLineJSON contains the JSON metadata for the
// struct [InvestigateTraceGetResponseInboundLine]
type investigateTraceGetResponseInboundLineJSON struct {
	Lineno      apijson.Field
	LoggedAt    apijson.Field
	Message     apijson.Field
	Ts          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceGetResponseInboundLine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseInboundLineJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceGetResponseOutbound struct {
	Lines   []InvestigateTraceGetResponseOutboundLine `json:"lines" api:"nullable"`
	Pending bool                                      `json:"pending" api:"nullable"`
	JSON    investigateTraceGetResponseOutboundJSON   `json:"-"`
}

// investigateTraceGetResponseOutboundJSON contains the JSON metadata for the
// struct [InvestigateTraceGetResponseOutbound]
type investigateTraceGetResponseOutboundJSON struct {
	Lines       apijson.Field
	Pending     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceGetResponseOutbound) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseOutboundJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceGetResponseOutboundLine struct {
	// Line number in the trace log
	Lineno   int64     `json:"lineno"`
	LoggedAt time.Time `json:"logged_at" api:"nullable" format:"date-time"`
	Message  string    `json:"message"`
	// Deprecated, use `logged_at` instead. End of life: November 1, 2026.
	//
	// Deprecated: deprecated
	Ts   string                                      `json:"ts"`
	JSON investigateTraceGetResponseOutboundLineJSON `json:"-"`
}

// investigateTraceGetResponseOutboundLineJSON contains the JSON metadata for the
// struct [InvestigateTraceGetResponseOutboundLine]
type investigateTraceGetResponseOutboundLineJSON struct {
	Lineno      apijson.Field
	LoggedAt    apijson.Field
	Message     apijson.Field
	Ts          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceGetResponseOutboundLine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseOutboundLineJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type InvestigateTraceGetResponseEnvelope struct {
	Errors   []InvestigateTraceGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []InvestigateTraceGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   InvestigateTraceGetResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success InvestigateTraceGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    investigateTraceGetResponseEnvelopeJSON    `json:"-"`
}

// investigateTraceGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [InvestigateTraceGetResponseEnvelope]
type investigateTraceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceGetResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           InvestigateTraceGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             investigateTraceGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// investigateTraceGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [InvestigateTraceGetResponseEnvelopeErrors]
type investigateTraceGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateTraceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceGetResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    investigateTraceGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// investigateTraceGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [InvestigateTraceGetResponseEnvelopeErrorsSource]
type investigateTraceGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceGetResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           InvestigateTraceGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             investigateTraceGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// investigateTraceGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [InvestigateTraceGetResponseEnvelopeMessages]
type investigateTraceGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateTraceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InvestigateTraceGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    investigateTraceGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// investigateTraceGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [InvestigateTraceGetResponseEnvelopeMessagesSource]
type investigateTraceGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateTraceGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateTraceGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InvestigateTraceGetResponseEnvelopeSuccess bool

const (
	InvestigateTraceGetResponseEnvelopeSuccessTrue InvestigateTraceGetResponseEnvelopeSuccess = true
)

func (r InvestigateTraceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InvestigateTraceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

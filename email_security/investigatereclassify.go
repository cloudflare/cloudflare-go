// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

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

// InvestigateReclassifyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateReclassifyService] method instead.
type InvestigateReclassifyService struct {
	Options []option.RequestOption
}

// NewInvestigateReclassifyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestigateReclassifyService(opts ...option.RequestOption) (r *InvestigateReclassifyService) {
	r = &InvestigateReclassifyService{}
	r.Options = opts
	return
}

// Submits a request to reclassify an email's disposition. Use for reporting false
// positives or false negatives. Optionally provide the raw EML content for
// reanalysis. The reclassification is processed asynchronously.
func (r *InvestigateReclassifyService) New(ctx context.Context, investigateID string, params InvestigateReclassifyNewParams, opts ...option.RequestOption) (res *InvestigateReclassifyNewResponse, err error) {
	var env InvestigateReclassifyNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if investigateID == "" {
		err = errors.New("missing required investigate_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/reclassify", params.AccountID, investigateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type InvestigateReclassifyNewResponse = interface{}

type InvestigateReclassifyNewParams struct {
	// Identifier.
	AccountID           param.Field[string]                                            `path:"account_id" api:"required"`
	ExpectedDisposition param.Field[InvestigateReclassifyNewParamsExpectedDisposition] `json:"expected_disposition" api:"required"`
	// Base64 encoded content of the EML file.
	EmlContent            param.Field[string] `json:"eml_content"`
	EscalatedSubmissionID param.Field[string] `json:"escalated_submission_id"`
}

func (r InvestigateReclassifyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestigateReclassifyNewParamsExpectedDisposition string

const (
	InvestigateReclassifyNewParamsExpectedDispositionNone       InvestigateReclassifyNewParamsExpectedDisposition = "NONE"
	InvestigateReclassifyNewParamsExpectedDispositionBulk       InvestigateReclassifyNewParamsExpectedDisposition = "BULK"
	InvestigateReclassifyNewParamsExpectedDispositionMalicious  InvestigateReclassifyNewParamsExpectedDisposition = "MALICIOUS"
	InvestigateReclassifyNewParamsExpectedDispositionSpam       InvestigateReclassifyNewParamsExpectedDisposition = "SPAM"
	InvestigateReclassifyNewParamsExpectedDispositionSpoof      InvestigateReclassifyNewParamsExpectedDisposition = "SPOOF"
	InvestigateReclassifyNewParamsExpectedDispositionSuspicious InvestigateReclassifyNewParamsExpectedDisposition = "SUSPICIOUS"
)

func (r InvestigateReclassifyNewParamsExpectedDisposition) IsKnown() bool {
	switch r {
	case InvestigateReclassifyNewParamsExpectedDispositionNone, InvestigateReclassifyNewParamsExpectedDispositionBulk, InvestigateReclassifyNewParamsExpectedDispositionMalicious, InvestigateReclassifyNewParamsExpectedDispositionSpam, InvestigateReclassifyNewParamsExpectedDispositionSpoof, InvestigateReclassifyNewParamsExpectedDispositionSuspicious:
		return true
	}
	return false
}

type InvestigateReclassifyNewResponseEnvelope struct {
	Errors   []InvestigateReclassifyNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []InvestigateReclassifyNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   InvestigateReclassifyNewResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success InvestigateReclassifyNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    investigateReclassifyNewResponseEnvelopeJSON    `json:"-"`
}

// investigateReclassifyNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [InvestigateReclassifyNewResponseEnvelope]
type investigateReclassifyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateReclassifyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateReclassifyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigateReclassifyNewResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           InvestigateReclassifyNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             investigateReclassifyNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// investigateReclassifyNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [InvestigateReclassifyNewResponseEnvelopeErrors]
type investigateReclassifyNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateReclassifyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateReclassifyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InvestigateReclassifyNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    investigateReclassifyNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// investigateReclassifyNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [InvestigateReclassifyNewResponseEnvelopeErrorsSource]
type investigateReclassifyNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateReclassifyNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateReclassifyNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InvestigateReclassifyNewResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code" api:"required"`
	Message          string                                                 `json:"message" api:"required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           InvestigateReclassifyNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             investigateReclassifyNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// investigateReclassifyNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [InvestigateReclassifyNewResponseEnvelopeMessages]
type investigateReclassifyNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateReclassifyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateReclassifyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InvestigateReclassifyNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    investigateReclassifyNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// investigateReclassifyNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [InvestigateReclassifyNewResponseEnvelopeMessagesSource]
type investigateReclassifyNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateReclassifyNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateReclassifyNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InvestigateReclassifyNewResponseEnvelopeSuccess bool

const (
	InvestigateReclassifyNewResponseEnvelopeSuccessTrue InvestigateReclassifyNewResponseEnvelopeSuccess = true
)

func (r InvestigateReclassifyNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InvestigateReclassifyNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

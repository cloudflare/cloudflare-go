// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_security

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
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

// Submits an email message for reclassification, updating its threat assessment
// based on new analysis.
func (r *InvestigateReclassifyService) New(ctx context.Context, postfixID string, params InvestigateReclassifyNewParams, opts ...option.RequestOption) (res *InvestigateReclassifyNewResponse, err error) {
	var env InvestigateReclassifyNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if postfixID == "" {
		err = errors.New("missing required postfix_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/reclassify", params.AccountID, postfixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type InvestigateReclassifyNewResponse = interface{}

type InvestigateReclassifyNewParams struct {
	// Account Identifier
	AccountID           param.Field[string]                                            `path:"account_id" api:"required"`
	ExpectedDisposition param.Field[InvestigateReclassifyNewParamsExpectedDisposition] `json:"expected_disposition" api:"required"`
	// When true, search the submissions datastore only. When false or omitted, search
	// the regular datastore only.
	Submission param.Field[bool] `query:"submission"`
	// Base64 encoded content of the EML file
	EmlContent            param.Field[string] `json:"eml_content"`
	EscalatedSubmissionID param.Field[string] `json:"escalated_submission_id"`
}

func (r InvestigateReclassifyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [InvestigateReclassifyNewParams]'s query parameters as
// `url.Values`.
func (r InvestigateReclassifyNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
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
	Errors   []shared.ResponseInfo                        `json:"errors" api:"required"`
	Messages []shared.ResponseInfo                        `json:"messages" api:"required"`
	Result   InvestigateReclassifyNewResponse             `json:"result" api:"required"`
	Success  bool                                         `json:"success" api:"required"`
	JSON     investigateReclassifyNewResponseEnvelopeJSON `json:"-"`
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

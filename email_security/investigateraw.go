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

// InvestigateRawService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigateRawService] method instead.
type InvestigateRawService struct {
	Options []option.RequestOption
}

// NewInvestigateRawService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInvestigateRawService(opts ...option.RequestOption) (r *InvestigateRawService) {
	r = &InvestigateRawService{}
	r.Options = opts
	return
}

// Returns the raw eml of any non-benign message.
func (r *InvestigateRawService) Get(ctx context.Context, investigateID string, query InvestigateRawGetParams, opts ...option.RequestOption) (res *InvestigateRawGetResponse, err error) {
	var env InvestigateRawGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if investigateID == "" {
		err = errors.New("missing required investigate_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/raw", query.AccountID, investigateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type InvestigateRawGetResponse struct {
	// A UTF-8 encoded eml file of the email.
	Raw  string                        `json:"raw" api:"required"`
	JSON investigateRawGetResponseJSON `json:"-"`
}

// investigateRawGetResponseJSON contains the JSON metadata for the struct
// [InvestigateRawGetResponse]
type investigateRawGetResponseJSON struct {
	Raw         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateRawGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateRawGetResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigateRawGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type InvestigateRawGetResponseEnvelope struct {
	Errors   []InvestigateRawGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []InvestigateRawGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   InvestigateRawGetResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success InvestigateRawGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    investigateRawGetResponseEnvelopeJSON    `json:"-"`
}

// investigateRawGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InvestigateRawGetResponseEnvelope]
type investigateRawGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateRawGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateRawGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigateRawGetResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           InvestigateRawGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             investigateRawGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// investigateRawGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [InvestigateRawGetResponseEnvelopeErrors]
type investigateRawGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateRawGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateRawGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InvestigateRawGetResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    investigateRawGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// investigateRawGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [InvestigateRawGetResponseEnvelopeErrorsSource]
type investigateRawGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateRawGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateRawGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InvestigateRawGetResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code" api:"required"`
	Message          string                                          `json:"message" api:"required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           InvestigateRawGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             investigateRawGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// investigateRawGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [InvestigateRawGetResponseEnvelopeMessages]
type investigateRawGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigateRawGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateRawGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InvestigateRawGetResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    investigateRawGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// investigateRawGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [InvestigateRawGetResponseEnvelopeMessagesSource]
type investigateRawGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigateRawGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigateRawGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InvestigateRawGetResponseEnvelopeSuccess bool

const (
	InvestigateRawGetResponseEnvelopeSuccessTrue InvestigateRawGetResponseEnvelopeSuccess = true
)

func (r InvestigateRawGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InvestigateRawGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

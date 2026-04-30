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

// InvestigatePreviewService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInvestigatePreviewService] method instead.
type InvestigatePreviewService struct {
	Options []option.RequestOption
}

// NewInvestigatePreviewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInvestigatePreviewService(opts ...option.RequestOption) (r *InvestigatePreviewService) {
	r = &InvestigatePreviewService{}
	r.Options = opts
	return
}

// Generates a preview image for a message that was not flagged as a detection.
// Useful for investigating benign messages. Returns a base64-encoded PNG
// screenshot of the email body.
func (r *InvestigatePreviewService) New(ctx context.Context, params InvestigatePreviewNewParams, opts ...option.RequestOption) (res *InvestigatePreviewNewResponse, err error) {
	var env InvestigatePreviewNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/preview", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns a preview of the message body as a base64 encoded PNG image for
// non-benign messages.
func (r *InvestigatePreviewService) Get(ctx context.Context, investigateID string, query InvestigatePreviewGetParams, opts ...option.RequestOption) (res *InvestigatePreviewGetResponse, err error) {
	var env InvestigatePreviewGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if investigateID == "" {
		err = errors.New("missing required investigate_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/email-security/investigate/%s/preview", query.AccountID, investigateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type InvestigatePreviewNewResponse struct {
	// A base64 encoded PNG image of the email.
	Screenshot string                            `json:"screenshot" api:"required"`
	JSON       investigatePreviewNewResponseJSON `json:"-"`
}

// investigatePreviewNewResponseJSON contains the JSON metadata for the struct
// [InvestigatePreviewNewResponse]
type investigatePreviewNewResponseJSON struct {
	Screenshot  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigatePreviewNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewNewResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigatePreviewGetResponse struct {
	// A base64 encoded PNG image of the email.
	Screenshot string                            `json:"screenshot" api:"required"`
	JSON       investigatePreviewGetResponseJSON `json:"-"`
}

// investigatePreviewGetResponseJSON contains the JSON metadata for the struct
// [InvestigatePreviewGetResponse]
type investigatePreviewGetResponseJSON struct {
	Screenshot  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigatePreviewGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewGetResponseJSON) RawJSON() string {
	return r.raw
}

type InvestigatePreviewNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The identifier of the message
	PostfixID param.Field[string] `json:"postfix_id" api:"required"`
}

func (r InvestigatePreviewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InvestigatePreviewNewResponseEnvelope struct {
	Errors   []InvestigatePreviewNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []InvestigatePreviewNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   InvestigatePreviewNewResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success InvestigatePreviewNewResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    investigatePreviewNewResponseEnvelopeJSON    `json:"-"`
}

// investigatePreviewNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [InvestigatePreviewNewResponseEnvelope]
type investigatePreviewNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigatePreviewNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigatePreviewNewResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           InvestigatePreviewNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             investigatePreviewNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// investigatePreviewNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [InvestigatePreviewNewResponseEnvelopeErrors]
type investigatePreviewNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigatePreviewNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InvestigatePreviewNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    investigatePreviewNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// investigatePreviewNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [InvestigatePreviewNewResponseEnvelopeErrorsSource]
type investigatePreviewNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigatePreviewNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InvestigatePreviewNewResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           InvestigatePreviewNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             investigatePreviewNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// investigatePreviewNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [InvestigatePreviewNewResponseEnvelopeMessages]
type investigatePreviewNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigatePreviewNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InvestigatePreviewNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    investigatePreviewNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// investigatePreviewNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [InvestigatePreviewNewResponseEnvelopeMessagesSource]
type investigatePreviewNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigatePreviewNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InvestigatePreviewNewResponseEnvelopeSuccess bool

const (
	InvestigatePreviewNewResponseEnvelopeSuccessTrue InvestigatePreviewNewResponseEnvelopeSuccess = true
)

func (r InvestigatePreviewNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InvestigatePreviewNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type InvestigatePreviewGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type InvestigatePreviewGetResponseEnvelope struct {
	Errors   []InvestigatePreviewGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []InvestigatePreviewGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   InvestigatePreviewGetResponse                   `json:"result" api:"required"`
	// Whether the API call was successful.
	Success InvestigatePreviewGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    investigatePreviewGetResponseEnvelopeJSON    `json:"-"`
}

// investigatePreviewGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [InvestigatePreviewGetResponseEnvelope]
type investigatePreviewGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigatePreviewGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InvestigatePreviewGetResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           InvestigatePreviewGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             investigatePreviewGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// investigatePreviewGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [InvestigatePreviewGetResponseEnvelopeErrors]
type investigatePreviewGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigatePreviewGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InvestigatePreviewGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    investigatePreviewGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// investigatePreviewGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [InvestigatePreviewGetResponseEnvelopeErrorsSource]
type investigatePreviewGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigatePreviewGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InvestigatePreviewGetResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           InvestigatePreviewGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             investigatePreviewGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// investigatePreviewGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [InvestigatePreviewGetResponseEnvelopeMessages]
type investigatePreviewGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InvestigatePreviewGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InvestigatePreviewGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    investigatePreviewGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// investigatePreviewGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [InvestigatePreviewGetResponseEnvelopeMessagesSource]
type investigatePreviewGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InvestigatePreviewGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r investigatePreviewGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InvestigatePreviewGetResponseEnvelopeSuccess bool

const (
	InvestigatePreviewGetResponseEnvelopeSuccessTrue InvestigatePreviewGetResponseEnvelopeSuccess = true
)

func (r InvestigatePreviewGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InvestigatePreviewGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

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

// TransformerContentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransformerContentService] method instead.
type TransformerContentService struct {
	Options []option.RequestOption
}

// NewTransformerContentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTransformerContentService(opts ...option.RequestOption) (r *TransformerContentService) {
	r = &TransformerContentService{}
	r.Options = opts
	return
}

// Returns the SQL query content for a transformer. Without query params, returns
// the latest version. With `version_id`, returns the specified version.
func (r *TransformerContentService) Get(ctx context.Context, transformerID int64, params TransformerContentGetParams, opts ...option.RequestOption) (res *TransformerContentGetResponse, err error) {
	var env TransformerContentGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/logpush/transformers/%v/content", params.AccountID, transformerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type TransformerContentGetResponse struct {
	// The SQL query content.
	Content string                            `json:"content"`
	JSON    transformerContentGetResponseJSON `json:"-"`
}

// transformerContentGetResponseJSON contains the JSON metadata for the struct
// [TransformerContentGetResponse]
type transformerContentGetResponseJSON struct {
	Content     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerContentGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerContentGetResponseJSON) RawJSON() string {
	return r.raw
}

type TransformerContentGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Specific version ID to retrieve. When omitted, the latest version is returned.
	VersionID param.Field[int64] `query:"version_id"`
}

// URLQuery serializes [TransformerContentGetParams]'s query parameters as
// `url.Values`.
func (r TransformerContentGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type TransformerContentGetResponseEnvelope struct {
	Errors   []TransformerContentGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []TransformerContentGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success TransformerContentGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  TransformerContentGetResponse                `json:"result"`
	JSON    transformerContentGetResponseEnvelopeJSON    `json:"-"`
}

// transformerContentGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [TransformerContentGetResponseEnvelope]
type transformerContentGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerContentGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerContentGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TransformerContentGetResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code" api:"required"`
	Message          string                                            `json:"message" api:"required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           TransformerContentGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             transformerContentGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// transformerContentGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [TransformerContentGetResponseEnvelopeErrors]
type transformerContentGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransformerContentGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerContentGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TransformerContentGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    transformerContentGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// transformerContentGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [TransformerContentGetResponseEnvelopeErrorsSource]
type transformerContentGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerContentGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerContentGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TransformerContentGetResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code" api:"required"`
	Message          string                                              `json:"message" api:"required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           TransformerContentGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             transformerContentGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// transformerContentGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [TransformerContentGetResponseEnvelopeMessages]
type transformerContentGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransformerContentGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerContentGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TransformerContentGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    transformerContentGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// transformerContentGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [TransformerContentGetResponseEnvelopeMessagesSource]
type transformerContentGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformerContentGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r transformerContentGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TransformerContentGetResponseEnvelopeSuccess bool

const (
	TransformerContentGetResponseEnvelopeSuccessTrue TransformerContentGetResponseEnvelopeSuccess = true
)

func (r TransformerContentGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TransformerContentGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

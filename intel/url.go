// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

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

// URLService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewURLService] method instead.
type URLService struct {
	Options []option.RequestOption
}

// NewURLService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewURLService(opts ...option.RequestOption) (r *URLService) {
	r = &URLService{}
	r.Options = opts
	return
}

// Gets security information about a URL, including content categories and risk
// types. The URL must be provided as a query parameter.
func (r *URLService) Get(ctx context.Context, params URLGetParams, opts ...option.RequestOption) (res *URL, err error) {
	var env URLGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/intel/url", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type URL struct {
	// Content categories associated with this URL.
	ContentCategories []URLContentCategory `json:"content_categories" api:"required"`
	// The full URL that was looked up.
	FullURL string `json:"full_url" api:"required"`
	// The hostname of the URL.
	Hostname string `json:"hostname" api:"required"`
	// Security risk types associated with this URL.
	RiskType []URLRiskType `json:"risk_type" api:"required"`
	// The path component of the URL.
	URLPath string  `json:"url_path" api:"required"`
	JSON    urlJSON `json:"-"`
}

// urlJSON contains the JSON metadata for the struct [URL]
type urlJSON struct {
	ContentCategories apijson.Field
	FullURL           apijson.Field
	Hostname          apijson.Field
	RiskType          apijson.Field
	URLPath           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *URL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlJSON) RawJSON() string {
	return r.raw
}

type URLContentCategory struct {
	ID              int64                  `json:"id"`
	Name            string                 `json:"name"`
	SourceID        int64                  `json:"source_id"`
	SuperCategoryID int64                  `json:"super_category_id"`
	JSON            urlContentCategoryJSON `json:"-"`
}

// urlContentCategoryJSON contains the JSON metadata for the struct
// [URLContentCategory]
type urlContentCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SourceID        apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLContentCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlContentCategoryJSON) RawJSON() string {
	return r.raw
}

type URLRiskType struct {
	ID              int64           `json:"id"`
	Name            string          `json:"name"`
	SourceID        int64           `json:"source_id"`
	SuperCategoryID int64           `json:"super_category_id"`
	JSON            urlRiskTypeJSON `json:"-"`
}

// urlRiskTypeJSON contains the JSON metadata for the struct [URLRiskType]
type urlRiskTypeJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SourceID        apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *URLRiskType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlRiskTypeJSON) RawJSON() string {
	return r.raw
}

type URLGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The URL to look up.
	URL param.Field[string] `query:"url" api:"required"`
}

// URLQuery serializes [URLGetParams]'s query parameters as `url.Values`.
func (r URLGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type URLGetResponseEnvelope struct {
	Errors   []URLGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []URLGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success URLGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  URL                           `json:"result"`
	JSON    urlGetResponseEnvelopeJSON    `json:"-"`
}

// urlGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [URLGetResponseEnvelope]
type urlGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type URLGetResponseEnvelopeErrors struct {
	Code             int64                              `json:"code" api:"required"`
	Message          string                             `json:"message" api:"required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           URLGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             urlGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// urlGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [URLGetResponseEnvelopeErrors]
type urlGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *URLGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type URLGetResponseEnvelopeErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    urlGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// urlGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the struct
// [URLGetResponseEnvelopeErrorsSource]
type urlGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type URLGetResponseEnvelopeMessages struct {
	Code             int64                                `json:"code" api:"required"`
	Message          string                               `json:"message" api:"required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           URLGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             urlGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// urlGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [URLGetResponseEnvelopeMessages]
type urlGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *URLGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type URLGetResponseEnvelopeMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    urlGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// urlGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [URLGetResponseEnvelopeMessagesSource]
type urlGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type URLGetResponseEnvelopeSuccess bool

const (
	URLGetResponseEnvelopeSuccessTrue URLGetResponseEnvelopeSuccess = true
)

func (r URLGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case URLGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

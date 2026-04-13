// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DLPLimitService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPLimitService] method instead.
type DLPLimitService struct {
	Options []option.RequestOption
}

// NewDLPLimitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPLimitService(opts ...option.RequestOption) (r *DLPLimitService) {
	r = &DLPLimitService{}
	r.Options = opts
	return
}

// Retrieves current DLP usage limits and quotas for the account, including maximum
// allowed counts and current usage for custom entries, dataset cells, and document
// fingerprints.
func (r *DLPLimitService) List(ctx context.Context, query DLPLimitListParams, opts ...option.RequestOption) (res *DLPLimitListResponse, err error) {
	var env DLPLimitListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/limits", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type DLPLimitListResponse struct {
	// Maximum number of custom regex entries allowed for the account.
	MaxCustomRegexEntries int64 `json:"max_custom_regex_entries" api:"required"`
	// Maximum number of dataset cells allowed for the account, across all EDM and CWL
	// datasets.
	MaxDatasetCells int64 `json:"max_dataset_cells" api:"required"`
	// Maximum number of document fingerprints allowed for the account.
	MaxDocumentFingerprints int64 `json:"max_document_fingerprints" api:"required"`
	// Number of custom regex entries currently configured for the account.
	UsedCustomRegexEntries int64 `json:"used_custom_regex_entries" api:"required"`
	// Number of dataset cells currently configured for the account, across all EDM and
	// CWL datasets. Document fingerprints do not count towards this limit.
	UsedDatasetCells int64 `json:"used_dataset_cells" api:"required"`
	// Number of document fingerprints currently configured for the account.
	UsedDocumentFingerprints int64                    `json:"used_document_fingerprints" api:"required"`
	JSON                     dlpLimitListResponseJSON `json:"-"`
}

// dlpLimitListResponseJSON contains the JSON metadata for the struct
// [DLPLimitListResponse]
type dlpLimitListResponseJSON struct {
	MaxCustomRegexEntries    apijson.Field
	MaxDatasetCells          apijson.Field
	MaxDocumentFingerprints  apijson.Field
	UsedCustomRegexEntries   apijson.Field
	UsedDatasetCells         apijson.Field
	UsedDocumentFingerprints apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *DLPLimitListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpLimitListResponseJSON) RawJSON() string {
	return r.raw
}

type DLPLimitListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPLimitListResponseEnvelope struct {
	Errors   []DLPLimitListResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPLimitListResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPLimitListResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPLimitListResponse                `json:"result"`
	JSON    dlpLimitListResponseEnvelopeJSON    `json:"-"`
}

// dlpLimitListResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPLimitListResponseEnvelope]
type dlpLimitListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPLimitListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpLimitListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPLimitListResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code" api:"required"`
	Message          string                                   `json:"message" api:"required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           DLPLimitListResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpLimitListResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpLimitListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DLPLimitListResponseEnvelopeErrors]
type dlpLimitListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPLimitListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpLimitListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPLimitListResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    dlpLimitListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpLimitListResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DLPLimitListResponseEnvelopeErrorsSource]
type dlpLimitListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPLimitListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpLimitListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPLimitListResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code" api:"required"`
	Message          string                                     `json:"message" api:"required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           DLPLimitListResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpLimitListResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpLimitListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPLimitListResponseEnvelopeMessages]
type dlpLimitListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPLimitListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpLimitListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPLimitListResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    dlpLimitListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpLimitListResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPLimitListResponseEnvelopeMessagesSource]
type dlpLimitListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPLimitListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpLimitListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DLPLimitListResponseEnvelopeSuccess bool

const (
	DLPLimitListResponseEnvelopeSuccessTrue DLPLimitListResponseEnvelopeSuccess = true
)

func (r DLPLimitListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPLimitListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

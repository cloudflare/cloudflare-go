// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DomainHistoryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDomainHistoryService] method
// instead.
type DomainHistoryService struct {
	Options []option.RequestOption
}

// NewDomainHistoryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDomainHistoryService(opts ...option.RequestOption) (r *DomainHistoryService) {
	r = &DomainHistoryService{}
	r.Options = opts
	return
}

// Get Domain History
func (r *DomainHistoryService) Get(ctx context.Context, params DomainHistoryGetParams, opts ...option.RequestOption) (res *[]IntelDomainHistory, err error) {
	opts = append(r.Options[:], opts...)
	var env DomainHistoryGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/domain-history", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelDomainHistory struct {
	Categorizations []IntelDomainHistoryCategorization `json:"categorizations"`
	Domain          string                             `json:"domain"`
	JSON            intelDomainHistoryJSON             `json:"-"`
}

// intelDomainHistoryJSON contains the JSON metadata for the struct
// [IntelDomainHistory]
type intelDomainHistoryJSON struct {
	Categorizations apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IntelDomainHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainHistoryJSON) RawJSON() string {
	return r.raw
}

type IntelDomainHistoryCategorization struct {
	Categories interface{}                          `json:"categories"`
	End        time.Time                            `json:"end" format:"date"`
	Start      time.Time                            `json:"start" format:"date"`
	JSON       intelDomainHistoryCategorizationJSON `json:"-"`
}

// intelDomainHistoryCategorizationJSON contains the JSON metadata for the struct
// [IntelDomainHistoryCategorization]
type intelDomainHistoryCategorizationJSON struct {
	Categories  apijson.Field
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainHistoryCategorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainHistoryCategorizationJSON) RawJSON() string {
	return r.raw
}

type DomainHistoryGetParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Domain    param.Field[interface{}] `query:"domain"`
}

// URLQuery serializes [DomainHistoryGetParams]'s query parameters as `url.Values`.
func (r DomainHistoryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DomainHistoryGetResponseEnvelope struct {
	Errors   []DomainHistoryGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DomainHistoryGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelDomainHistory                       `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    DomainHistoryGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo DomainHistoryGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       domainHistoryGetResponseEnvelopeJSON       `json:"-"`
}

// domainHistoryGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainHistoryGetResponseEnvelope]
type domainHistoryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainHistoryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainHistoryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DomainHistoryGetResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    domainHistoryGetResponseEnvelopeErrorsJSON `json:"-"`
}

// domainHistoryGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DomainHistoryGetResponseEnvelopeErrors]
type domainHistoryGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainHistoryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainHistoryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DomainHistoryGetResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    domainHistoryGetResponseEnvelopeMessagesJSON `json:"-"`
}

// domainHistoryGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DomainHistoryGetResponseEnvelopeMessages]
type domainHistoryGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainHistoryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainHistoryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DomainHistoryGetResponseEnvelopeSuccess bool

const (
	DomainHistoryGetResponseEnvelopeSuccessTrue DomainHistoryGetResponseEnvelopeSuccess = true
)

func (r DomainHistoryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainHistoryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DomainHistoryGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                        `json:"total_count"`
	JSON       domainHistoryGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// domainHistoryGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [DomainHistoryGetResponseEnvelopeResultInfo]
type domainHistoryGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainHistoryGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainHistoryGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

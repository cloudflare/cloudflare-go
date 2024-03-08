// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// IntelDomainHistoryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewIntelDomainHistoryService] method
// instead.
type IntelDomainHistoryService struct {
	Options []option.RequestOption
}

// NewIntelDomainHistoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIntelDomainHistoryService(opts ...option.RequestOption) (r *IntelDomainHistoryService) {
	r = &IntelDomainHistoryService{}
	r.Options = opts
	return
}

// Get Domain History
func (r *IntelDomainHistoryService) Get(ctx context.Context, params IntelDomainHistoryGetParams, opts ...option.RequestOption) (res *[]IntelDomainHistoryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelDomainHistoryGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/domain-history", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelDomainHistoryGetResponse struct {
	Categorizations []IntelDomainHistoryGetResponseCategorization `json:"categorizations"`
	Domain          string                                        `json:"domain"`
	JSON            intelDomainHistoryGetResponseJSON             `json:"-"`
}

// intelDomainHistoryGetResponseJSON contains the JSON metadata for the struct
// [IntelDomainHistoryGetResponse]
type intelDomainHistoryGetResponseJSON struct {
	Categorizations apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IntelDomainHistoryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainHistoryGetResponseJSON) RawJSON() string {
	return r.raw
}

type IntelDomainHistoryGetResponseCategorization struct {
	Categories interface{}                                     `json:"categories"`
	End        time.Time                                       `json:"end" format:"date"`
	Start      time.Time                                       `json:"start" format:"date"`
	JSON       intelDomainHistoryGetResponseCategorizationJSON `json:"-"`
}

// intelDomainHistoryGetResponseCategorizationJSON contains the JSON metadata for
// the struct [IntelDomainHistoryGetResponseCategorization]
type intelDomainHistoryGetResponseCategorizationJSON struct {
	Categories  apijson.Field
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainHistoryGetResponseCategorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainHistoryGetResponseCategorizationJSON) RawJSON() string {
	return r.raw
}

type IntelDomainHistoryGetParams struct {
	// Identifier
	AccountID param.Field[string]      `path:"account_id,required"`
	Domain    param.Field[interface{}] `query:"domain"`
}

// URLQuery serializes [IntelDomainHistoryGetParams]'s query parameters as
// `url.Values`.
func (r IntelDomainHistoryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntelDomainHistoryGetResponseEnvelope struct {
	Errors   []IntelDomainHistoryGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelDomainHistoryGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelDomainHistoryGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    IntelDomainHistoryGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo IntelDomainHistoryGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       intelDomainHistoryGetResponseEnvelopeJSON       `json:"-"`
}

// intelDomainHistoryGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [IntelDomainHistoryGetResponseEnvelope]
type intelDomainHistoryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainHistoryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainHistoryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IntelDomainHistoryGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    intelDomainHistoryGetResponseEnvelopeErrorsJSON `json:"-"`
}

// intelDomainHistoryGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [IntelDomainHistoryGetResponseEnvelopeErrors]
type intelDomainHistoryGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainHistoryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainHistoryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IntelDomainHistoryGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    intelDomainHistoryGetResponseEnvelopeMessagesJSON `json:"-"`
}

// intelDomainHistoryGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [IntelDomainHistoryGetResponseEnvelopeMessages]
type intelDomainHistoryGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainHistoryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainHistoryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IntelDomainHistoryGetResponseEnvelopeSuccess bool

const (
	IntelDomainHistoryGetResponseEnvelopeSuccessTrue IntelDomainHistoryGetResponseEnvelopeSuccess = true
)

type IntelDomainHistoryGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                             `json:"total_count"`
	JSON       intelDomainHistoryGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// intelDomainHistoryGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [IntelDomainHistoryGetResponseEnvelopeResultInfo]
type intelDomainHistoryGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainHistoryGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainHistoryGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

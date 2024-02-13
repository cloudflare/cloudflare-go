// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
func (r *IntelDomainHistoryService) List(ctx context.Context, accountID string, query IntelDomainHistoryListParams, opts ...option.RequestOption) (res *[]IntelDomainHistoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelDomainHistoryListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/domain-history", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelDomainHistoryListResponse struct {
	Categorizations []IntelDomainHistoryListResponseCategorization `json:"categorizations"`
	Domain          string                                         `json:"domain"`
	JSON            intelDomainHistoryListResponseJSON             `json:"-"`
}

// intelDomainHistoryListResponseJSON contains the JSON metadata for the struct
// [IntelDomainHistoryListResponse]
type intelDomainHistoryListResponseJSON struct {
	Categorizations apijson.Field
	Domain          apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IntelDomainHistoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainHistoryListResponseCategorization struct {
	Categories interface{}                                      `json:"categories"`
	End        time.Time                                        `json:"end" format:"date"`
	Start      time.Time                                        `json:"start" format:"date"`
	JSON       intelDomainHistoryListResponseCategorizationJSON `json:"-"`
}

// intelDomainHistoryListResponseCategorizationJSON contains the JSON metadata for
// the struct [IntelDomainHistoryListResponseCategorization]
type intelDomainHistoryListResponseCategorizationJSON struct {
	Categories  apijson.Field
	End         apijson.Field
	Start       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainHistoryListResponseCategorization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainHistoryListParams struct {
	Domain param.Field[interface{}] `query:"domain"`
}

// URLQuery serializes [IntelDomainHistoryListParams]'s query parameters as
// `url.Values`.
func (r IntelDomainHistoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntelDomainHistoryListResponseEnvelope struct {
	Errors   []IntelDomainHistoryListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelDomainHistoryListResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelDomainHistoryListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    IntelDomainHistoryListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo IntelDomainHistoryListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       intelDomainHistoryListResponseEnvelopeJSON       `json:"-"`
}

// intelDomainHistoryListResponseEnvelopeJSON contains the JSON metadata for the
// struct [IntelDomainHistoryListResponseEnvelope]
type intelDomainHistoryListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainHistoryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainHistoryListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    intelDomainHistoryListResponseEnvelopeErrorsJSON `json:"-"`
}

// intelDomainHistoryListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [IntelDomainHistoryListResponseEnvelopeErrors]
type intelDomainHistoryListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainHistoryListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainHistoryListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    intelDomainHistoryListResponseEnvelopeMessagesJSON `json:"-"`
}

// intelDomainHistoryListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [IntelDomainHistoryListResponseEnvelopeMessages]
type intelDomainHistoryListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainHistoryListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelDomainHistoryListResponseEnvelopeSuccess bool

const (
	IntelDomainHistoryListResponseEnvelopeSuccessTrue IntelDomainHistoryListResponseEnvelopeSuccess = true
)

type IntelDomainHistoryListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       intelDomainHistoryListResponseEnvelopeResultInfoJSON `json:"-"`
}

// intelDomainHistoryListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [IntelDomainHistoryListResponseEnvelopeResultInfo]
type intelDomainHistoryListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainHistoryListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

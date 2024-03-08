// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// IntelDomainBulkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewIntelDomainBulkService] method
// instead.
type IntelDomainBulkService struct {
	Options []option.RequestOption
}

// NewIntelDomainBulkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntelDomainBulkService(opts ...option.RequestOption) (r *IntelDomainBulkService) {
	r = &IntelDomainBulkService{}
	r.Options = opts
	return
}

// Get Multiple Domain Details
func (r *IntelDomainBulkService) Get(ctx context.Context, params IntelDomainBulkGetParams, opts ...option.RequestOption) (res *[]IntelDomainBulkGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelDomainBulkGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/domain/bulk", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelDomainBulkGetResponse struct {
	// Additional information related to the host name.
	AdditionalInformation IntelDomainBulkGetResponseAdditionalInformation `json:"additional_information"`
	// Application that the hostname belongs to.
	Application IntelDomainBulkGetResponseApplication `json:"application"`
	// Current content categories.
	ContentCategories          interface{}                                          `json:"content_categories"`
	Domain                     string                                               `json:"domain"`
	InheritedContentCategories []IntelDomainBulkGetResponseInheritedContentCategory `json:"inherited_content_categories"`
	// Domain from which `inherited_content_categories` and `inherited_risk_types` are
	// inherited, if applicable.
	InheritedFrom      string                                        `json:"inherited_from"`
	InheritedRiskTypes []IntelDomainBulkGetResponseInheritedRiskType `json:"inherited_risk_types"`
	// Global Cloudflare 100k ranking for the last 30 days, if available for the
	// hostname. The top ranked domain is 1, the lowest ranked domain is 100,000.
	PopularityRank int64 `json:"popularity_rank"`
	// Hostname risk score, which is a value between 0 (lowest risk) to 1 (highest
	// risk).
	RiskScore float64                        `json:"risk_score"`
	RiskTypes interface{}                    `json:"risk_types"`
	JSON      intelDomainBulkGetResponseJSON `json:"-"`
}

// intelDomainBulkGetResponseJSON contains the JSON metadata for the struct
// [IntelDomainBulkGetResponse]
type intelDomainBulkGetResponseJSON struct {
	AdditionalInformation      apijson.Field
	Application                apijson.Field
	ContentCategories          apijson.Field
	Domain                     apijson.Field
	InheritedContentCategories apijson.Field
	InheritedFrom              apijson.Field
	InheritedRiskTypes         apijson.Field
	PopularityRank             apijson.Field
	RiskScore                  apijson.Field
	RiskTypes                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *IntelDomainBulkGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainBulkGetResponseJSON) RawJSON() string {
	return r.raw
}

// Additional information related to the host name.
type IntelDomainBulkGetResponseAdditionalInformation struct {
	// Suspected DGA malware family.
	SuspectedMalwareFamily string                                              `json:"suspected_malware_family"`
	JSON                   intelDomainBulkGetResponseAdditionalInformationJSON `json:"-"`
}

// intelDomainBulkGetResponseAdditionalInformationJSON contains the JSON metadata
// for the struct [IntelDomainBulkGetResponseAdditionalInformation]
type intelDomainBulkGetResponseAdditionalInformationJSON struct {
	SuspectedMalwareFamily apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IntelDomainBulkGetResponseAdditionalInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainBulkGetResponseAdditionalInformationJSON) RawJSON() string {
	return r.raw
}

// Application that the hostname belongs to.
type IntelDomainBulkGetResponseApplication struct {
	ID   int64                                     `json:"id"`
	Name string                                    `json:"name"`
	JSON intelDomainBulkGetResponseApplicationJSON `json:"-"`
}

// intelDomainBulkGetResponseApplicationJSON contains the JSON metadata for the
// struct [IntelDomainBulkGetResponseApplication]
type intelDomainBulkGetResponseApplicationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainBulkGetResponseApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainBulkGetResponseApplicationJSON) RawJSON() string {
	return r.raw
}

type IntelDomainBulkGetResponseInheritedContentCategory struct {
	ID              int64                                                  `json:"id"`
	Name            string                                                 `json:"name"`
	SuperCategoryID int64                                                  `json:"super_category_id"`
	JSON            intelDomainBulkGetResponseInheritedContentCategoryJSON `json:"-"`
}

// intelDomainBulkGetResponseInheritedContentCategoryJSON contains the JSON
// metadata for the struct [IntelDomainBulkGetResponseInheritedContentCategory]
type intelDomainBulkGetResponseInheritedContentCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IntelDomainBulkGetResponseInheritedContentCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainBulkGetResponseInheritedContentCategoryJSON) RawJSON() string {
	return r.raw
}

type IntelDomainBulkGetResponseInheritedRiskType struct {
	ID              int64                                           `json:"id"`
	Name            string                                          `json:"name"`
	SuperCategoryID int64                                           `json:"super_category_id"`
	JSON            intelDomainBulkGetResponseInheritedRiskTypeJSON `json:"-"`
}

// intelDomainBulkGetResponseInheritedRiskTypeJSON contains the JSON metadata for
// the struct [IntelDomainBulkGetResponseInheritedRiskType]
type intelDomainBulkGetResponseInheritedRiskTypeJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IntelDomainBulkGetResponseInheritedRiskType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainBulkGetResponseInheritedRiskTypeJSON) RawJSON() string {
	return r.raw
}

type IntelDomainBulkGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Accepts multiple values, i.e. `?domain=cloudflare.com&domain=example.com`.
	Domain param.Field[interface{}] `query:"domain"`
}

// URLQuery serializes [IntelDomainBulkGetParams]'s query parameters as
// `url.Values`.
func (r IntelDomainBulkGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntelDomainBulkGetResponseEnvelope struct {
	Errors   []IntelDomainBulkGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelDomainBulkGetResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelDomainBulkGetResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    IntelDomainBulkGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo IntelDomainBulkGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       intelDomainBulkGetResponseEnvelopeJSON       `json:"-"`
}

// intelDomainBulkGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IntelDomainBulkGetResponseEnvelope]
type intelDomainBulkGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainBulkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainBulkGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IntelDomainBulkGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    intelDomainBulkGetResponseEnvelopeErrorsJSON `json:"-"`
}

// intelDomainBulkGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IntelDomainBulkGetResponseEnvelopeErrors]
type intelDomainBulkGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainBulkGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainBulkGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IntelDomainBulkGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    intelDomainBulkGetResponseEnvelopeMessagesJSON `json:"-"`
}

// intelDomainBulkGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [IntelDomainBulkGetResponseEnvelopeMessages]
type intelDomainBulkGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainBulkGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainBulkGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IntelDomainBulkGetResponseEnvelopeSuccess bool

const (
	IntelDomainBulkGetResponseEnvelopeSuccessTrue IntelDomainBulkGetResponseEnvelopeSuccess = true
)

type IntelDomainBulkGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                          `json:"total_count"`
	JSON       intelDomainBulkGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// intelDomainBulkGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [IntelDomainBulkGetResponseEnvelopeResultInfo]
type intelDomainBulkGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainBulkGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelDomainBulkGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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
func (r *IntelDomainBulkService) DomainIntelligenceGetMultipleDomainDetails(ctx context.Context, accountID string, query IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsParams, opts ...option.RequestOption) (res *[]IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/domain/bulk", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponse struct {
	// Additional information related to the host name.
	AdditionalInformation IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseAdditionalInformation `json:"additional_information"`
	// Application that the hostname belongs to.
	Application IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseApplication `json:"application"`
	// Current content categories.
	ContentCategories          interface{}                                                                                 `json:"content_categories"`
	Domain                     string                                                                                      `json:"domain"`
	InheritedContentCategories []IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseInheritedContentCategory `json:"inherited_content_categories"`
	// Domain from which `inherited_content_categories` and `inherited_risk_types` are
	// inherited, if applicable.
	InheritedFrom      string                                                                               `json:"inherited_from"`
	InheritedRiskTypes []IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseInheritedRiskType `json:"inherited_risk_types"`
	// Global Cloudflare 100k ranking for the last 30 days, if available for the
	// hostname. The top ranked domain is 1, the lowest ranked domain is 100,000.
	PopularityRank int64 `json:"popularity_rank"`
	// Hostname risk score, which is a value between 0 (lowest risk) to 1 (highest
	// risk).
	RiskScore float64                                                               `json:"risk_score"`
	RiskTypes interface{}                                                           `json:"risk_types"`
	JSON      intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseJSON `json:"-"`
}

// intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseJSON contains
// the JSON metadata for the struct
// [IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponse]
type intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseJSON struct {
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

func (r *IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information related to the host name.
type IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseAdditionalInformation struct {
	// Suspected DGA malware family.
	SuspectedMalwareFamily string                                                                                     `json:"suspected_malware_family"`
	JSON                   intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseAdditionalInformationJSON `json:"-"`
}

// intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseAdditionalInformationJSON
// contains the JSON metadata for the struct
// [IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseAdditionalInformation]
type intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseAdditionalInformationJSON struct {
	SuspectedMalwareFamily apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseAdditionalInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Application that the hostname belongs to.
type IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseApplication struct {
	ID   int64                                                                            `json:"id"`
	Name string                                                                           `json:"name"`
	JSON intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseApplicationJSON `json:"-"`
}

// intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseApplicationJSON
// contains the JSON metadata for the struct
// [IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseApplication]
type intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseApplicationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseInheritedContentCategory struct {
	ID              int64                                                                                         `json:"id"`
	Name            string                                                                                        `json:"name"`
	SuperCategoryID int64                                                                                         `json:"super_category_id"`
	JSON            intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseInheritedContentCategoryJSON `json:"-"`
}

// intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseInheritedContentCategoryJSON
// contains the JSON metadata for the struct
// [IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseInheritedContentCategory]
type intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseInheritedContentCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseInheritedContentCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseInheritedRiskType struct {
	ID              int64                                                                                  `json:"id"`
	Name            string                                                                                 `json:"name"`
	SuperCategoryID int64                                                                                  `json:"super_category_id"`
	JSON            intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseInheritedRiskTypeJSON `json:"-"`
}

// intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseInheritedRiskTypeJSON
// contains the JSON metadata for the struct
// [IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseInheritedRiskType]
type intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseInheritedRiskTypeJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseInheritedRiskType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsParams struct {
	// Accepts multiple values, i.e. `?domain=cloudflare.com&domain=example.com`.
	Domain param.Field[interface{}] `query:"domain"`
}

// URLQuery serializes
// [IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsParams]'s query
// parameters as `url.Values`.
func (r IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelope struct {
	Errors   []IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeMessages `json:"messages,required"`
	Result   []IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeJSON       `json:"-"`
}

// intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelope]
type intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeErrors struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeErrorsJSON `json:"-"`
}

// intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeErrors]
type intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeMessages struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeMessagesJSON `json:"-"`
}

// intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeMessages]
type intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeSuccess bool

const (
	IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeSuccessTrue IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeSuccess = true
)

type IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                 `json:"total_count"`
	JSON       intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeResultInfoJSON `json:"-"`
}

// intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeResultInfo]
type intelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

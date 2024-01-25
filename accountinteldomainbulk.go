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

// AccountIntelDomainBulkService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountIntelDomainBulkService]
// method instead.
type AccountIntelDomainBulkService struct {
	Options []option.RequestOption
}

// NewAccountIntelDomainBulkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountIntelDomainBulkService(opts ...option.RequestOption) (r *AccountIntelDomainBulkService) {
	r = &AccountIntelDomainBulkService{}
	r.Options = opts
	return
}

// Get Multiple Domain Details
func (r *AccountIntelDomainBulkService) DomainIntelligenceGetMultipleDomainDetails(ctx context.Context, accountIdentifier string, query AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsParams, opts ...option.RequestOption) (res *AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/domain/bulk", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponse struct {
	Errors     []AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseError    `json:"errors"`
	Messages   []AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseMessage  `json:"messages"`
	Result     []AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResult   `json:"result"`
	ResultInfo AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseSuccess `json:"success"`
	JSON    accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseJSON    `json:"-"`
}

// accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseJSON
// contains the JSON metadata for the struct
// [AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponse]
type accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseError struct {
	Code    int64                                                                             `json:"code,required"`
	Message string                                                                            `json:"message,required"`
	JSON    accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseErrorJSON `json:"-"`
}

// accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseError]
type accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseMessage struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseMessageJSON `json:"-"`
}

// accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseMessage]
type accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResult struct {
	// Additional information related to the host name.
	AdditionalInformation AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultAdditionalInformation `json:"additional_information"`
	// Application that the hostname belongs to.
	Application AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultApplication `json:"application"`
	// Current content categories.
	ContentCategories          interface{}                                                                                              `json:"content_categories"`
	Domain                     string                                                                                                   `json:"domain"`
	InheritedContentCategories []AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInheritedContentCategory `json:"inherited_content_categories"`
	// Domain from which `inherited_content_categories` and `inherited_risk_types` are
	// inherited, if applicable.
	InheritedFrom      string                                                                                            `json:"inherited_from"`
	InheritedRiskTypes []AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInheritedRiskType `json:"inherited_risk_types"`
	// Global Cloudflare 100k ranking for the last 30 days, if available for the
	// hostname. The top ranked domain is 1, the lowest ranked domain is 100,000.
	PopularityRank int64 `json:"popularity_rank"`
	// Hostname risk score, which is a value between 0 (lowest risk) to 1 (highest
	// risk).
	RiskScore float64                                                                            `json:"risk_score"`
	RiskTypes interface{}                                                                        `json:"risk_types"`
	JSON      accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultJSON `json:"-"`
}

// accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultJSON
// contains the JSON metadata for the struct
// [AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResult]
type accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultJSON struct {
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

func (r *AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information related to the host name.
type AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultAdditionalInformation struct {
	// Suspected DGA malware family.
	SuspectedMalwareFamily string                                                                                                  `json:"suspected_malware_family"`
	JSON                   accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultAdditionalInformationJSON `json:"-"`
}

// accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultAdditionalInformationJSON
// contains the JSON metadata for the struct
// [AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultAdditionalInformation]
type accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultAdditionalInformationJSON struct {
	SuspectedMalwareFamily apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultAdditionalInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Application that the hostname belongs to.
type AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultApplication struct {
	ID   int64                                                                                         `json:"id"`
	Name string                                                                                        `json:"name"`
	JSON accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultApplicationJSON `json:"-"`
}

// accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultApplicationJSON
// contains the JSON metadata for the struct
// [AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultApplication]
type accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultApplicationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInheritedContentCategory struct {
	ID              int64                                                                                                      `json:"id"`
	Name            string                                                                                                     `json:"name"`
	SuperCategoryID int64                                                                                                      `json:"super_category_id"`
	JSON            accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInheritedContentCategoryJSON `json:"-"`
}

// accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInheritedContentCategoryJSON
// contains the JSON metadata for the struct
// [AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInheritedContentCategory]
type accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInheritedContentCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInheritedContentCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInheritedRiskType struct {
	ID              int64                                                                                               `json:"id"`
	Name            string                                                                                              `json:"name"`
	SuperCategoryID int64                                                                                               `json:"super_category_id"`
	JSON            accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInheritedRiskTypeJSON `json:"-"`
}

// accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInheritedRiskTypeJSON
// contains the JSON metadata for the struct
// [AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInheritedRiskType]
type accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInheritedRiskTypeJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInheritedRiskType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                `json:"total_count"`
	JSON       accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInfoJSON `json:"-"`
}

// accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInfo]
type accountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseSuccess bool

const (
	AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseSuccessTrue AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsResponseSuccess = true
)

type AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsParams struct {
	// Accepts multiple values, i.e. `?domain=cloudflare.com&domain=example.com`.
	Domain param.Field[interface{}] `query:"domain"`
}

// URLQuery serializes
// [AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsParams]'s query
// parameters as `url.Values`.
func (r AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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

// AccountIntelDomainService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountIntelDomainService] method
// instead.
type AccountIntelDomainService struct {
	Options []option.RequestOption
	Bulks   *AccountIntelDomainBulkService
}

// NewAccountIntelDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountIntelDomainService(opts ...option.RequestOption) (r *AccountIntelDomainService) {
	r = &AccountIntelDomainService{}
	r.Options = opts
	r.Bulks = NewAccountIntelDomainBulkService(opts...)
	return
}

// Get Domain Details
func (r *AccountIntelDomainService) DomainIntelligenceGetDomainDetails(ctx context.Context, accountIdentifier string, query AccountIntelDomainDomainIntelligenceGetDomainDetailsParams, opts ...option.RequestOption) (res *AccountIntelDomainDomainIntelligenceGetDomainDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/domain", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountIntelDomainDomainIntelligenceGetDomainDetailsResponse struct {
	Errors   []AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseError   `json:"errors"`
	Messages []AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseMessage `json:"messages"`
	Result   AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseSuccess `json:"success"`
	JSON    accountIntelDomainDomainIntelligenceGetDomainDetailsResponseJSON    `json:"-"`
}

// accountIntelDomainDomainIntelligenceGetDomainDetailsResponseJSON contains the
// JSON metadata for the struct
// [AccountIntelDomainDomainIntelligenceGetDomainDetailsResponse]
type accountIntelDomainDomainIntelligenceGetDomainDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainDomainIntelligenceGetDomainDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseError struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accountIntelDomainDomainIntelligenceGetDomainDetailsResponseErrorJSON `json:"-"`
}

// accountIntelDomainDomainIntelligenceGetDomainDetailsResponseErrorJSON contains
// the JSON metadata for the struct
// [AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseError]
type accountIntelDomainDomainIntelligenceGetDomainDetailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseMessage struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    accountIntelDomainDomainIntelligenceGetDomainDetailsResponseMessageJSON `json:"-"`
}

// accountIntelDomainDomainIntelligenceGetDomainDetailsResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseMessage]
type accountIntelDomainDomainIntelligenceGetDomainDetailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResult struct {
	// Additional information related to the host name.
	AdditionalInformation AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultAdditionalInformation `json:"additional_information"`
	// Application that the hostname belongs to.
	Application AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultApplication `json:"application"`
	// Current content categories.
	ContentCategories          interface{}                                                                                  `json:"content_categories"`
	Domain                     string                                                                                       `json:"domain"`
	InheritedContentCategories []AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultInheritedContentCategory `json:"inherited_content_categories"`
	// Domain from which `inherited_content_categories` and `inherited_risk_types` are
	// inherited, if applicable.
	InheritedFrom      string                                                                                `json:"inherited_from"`
	InheritedRiskTypes []AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultInheritedRiskType `json:"inherited_risk_types"`
	// Global Cloudflare 100k ranking for the last 30 days, if available for the
	// hostname. The top ranked domain is 1, the lowest ranked domain is 100,000.
	PopularityRank int64 `json:"popularity_rank"`
	// Specifies a list of references to one or more IP addresses or domain names that
	// the domain name currently resolves to.
	ResolvesToRefs []AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultResolvesToRef `json:"resolves_to_refs"`
	// Hostname risk score, which is a value between 0 (lowest risk) to 1 (highest
	// risk).
	RiskScore float64                                                                `json:"risk_score"`
	RiskTypes interface{}                                                            `json:"risk_types"`
	JSON      accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultJSON `json:"-"`
}

// accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultJSON contains
// the JSON metadata for the struct
// [AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResult]
type accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultJSON struct {
	AdditionalInformation      apijson.Field
	Application                apijson.Field
	ContentCategories          apijson.Field
	Domain                     apijson.Field
	InheritedContentCategories apijson.Field
	InheritedFrom              apijson.Field
	InheritedRiskTypes         apijson.Field
	PopularityRank             apijson.Field
	ResolvesToRefs             apijson.Field
	RiskScore                  apijson.Field
	RiskTypes                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information related to the host name.
type AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultAdditionalInformation struct {
	// Suspected DGA malware family.
	SuspectedMalwareFamily string                                                                                      `json:"suspected_malware_family"`
	JSON                   accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultAdditionalInformationJSON `json:"-"`
}

// accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultAdditionalInformationJSON
// contains the JSON metadata for the struct
// [AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultAdditionalInformation]
type accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultAdditionalInformationJSON struct {
	SuspectedMalwareFamily apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultAdditionalInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Application that the hostname belongs to.
type AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultApplication struct {
	ID   int64                                                                             `json:"id"`
	Name string                                                                            `json:"name"`
	JSON accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultApplicationJSON `json:"-"`
}

// accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultApplicationJSON
// contains the JSON metadata for the struct
// [AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultApplication]
type accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultApplicationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultInheritedContentCategory struct {
	ID              int64                                                                                          `json:"id"`
	Name            string                                                                                         `json:"name"`
	SuperCategoryID int64                                                                                          `json:"super_category_id"`
	JSON            accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultInheritedContentCategoryJSON `json:"-"`
}

// accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultInheritedContentCategoryJSON
// contains the JSON metadata for the struct
// [AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultInheritedContentCategory]
type accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultInheritedContentCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultInheritedContentCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultInheritedRiskType struct {
	ID              int64                                                                                   `json:"id"`
	Name            string                                                                                  `json:"name"`
	SuperCategoryID int64                                                                                   `json:"super_category_id"`
	JSON            accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultInheritedRiskTypeJSON `json:"-"`
}

// accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultInheritedRiskTypeJSON
// contains the JSON metadata for the struct
// [AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultInheritedRiskType]
type accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultInheritedRiskTypeJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultInheritedRiskType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultResolvesToRef struct {
	// STIX 2.1 identifier:
	// https://docs.oasis-open.org/cti/stix/v2.1/cs02/stix-v2.1-cs02.html#_64yvzeku5a5c
	ID string `json:"id"`
	// IP address or domain name.
	Value string                                                                              `json:"value"`
	JSON  accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultResolvesToRefJSON `json:"-"`
}

// accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultResolvesToRefJSON
// contains the JSON metadata for the struct
// [AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultResolvesToRef]
type accountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultResolvesToRefJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseResultResolvesToRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseSuccess bool

const (
	AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseSuccessTrue AccountIntelDomainDomainIntelligenceGetDomainDetailsResponseSuccess = true
)

type AccountIntelDomainDomainIntelligenceGetDomainDetailsParams struct {
	Domain param.Field[string] `query:"domain"`
}

// URLQuery serializes
// [AccountIntelDomainDomainIntelligenceGetDomainDetailsParams]'s query parameters
// as `url.Values`.
func (r AccountIntelDomainDomainIntelligenceGetDomainDetailsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

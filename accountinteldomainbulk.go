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
func (r *AccountIntelDomainBulkService) DomainIntelligenceGetMultipleDomainDetails(ctx context.Context, accountIdentifier string, query AccountIntelDomainBulkDomainIntelligenceGetMultipleDomainDetailsParams, opts ...option.RequestOption) (res *SchemasCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/domain/bulk", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SchemasCollectionResponse struct {
	Errors     []SchemasCollectionResponseError    `json:"errors"`
	Messages   []SchemasCollectionResponseMessage  `json:"messages"`
	Result     []SchemasCollectionResponseResult   `json:"result"`
	ResultInfo SchemasCollectionResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success SchemasCollectionResponseSuccess `json:"success"`
	JSON    schemasCollectionResponseJSON    `json:"-"`
}

// schemasCollectionResponseJSON contains the JSON metadata for the struct
// [SchemasCollectionResponse]
type schemasCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasCollectionResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    schemasCollectionResponseErrorJSON `json:"-"`
}

// schemasCollectionResponseErrorJSON contains the JSON metadata for the struct
// [SchemasCollectionResponseError]
type schemasCollectionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCollectionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasCollectionResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    schemasCollectionResponseMessageJSON `json:"-"`
}

// schemasCollectionResponseMessageJSON contains the JSON metadata for the struct
// [SchemasCollectionResponseMessage]
type schemasCollectionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCollectionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasCollectionResponseResult struct {
	// Additional information related to the host name.
	AdditionalInformation SchemasCollectionResponseResultAdditionalInformation `json:"additional_information"`
	// Application that the hostname belongs to.
	Application SchemasCollectionResponseResultApplication `json:"application"`
	// Current content categories.
	ContentCategories interface{} `json:"content_categories"`
	Domain            string      `json:"domain"`
	// Global Cloudflare 100k ranking for the last 30 days, if available for the
	// hostname. The top ranked domain is 1, the lowest ranked domain is 100,000.
	PopularityRank int64 `json:"popularity_rank"`
	// Hostname risk score, which is a value between 0 (lowest risk) to 1 (highest
	// risk).
	RiskScore float64                             `json:"risk_score"`
	RiskTypes interface{}                         `json:"risk_types"`
	JSON      schemasCollectionResponseResultJSON `json:"-"`
}

// schemasCollectionResponseResultJSON contains the JSON metadata for the struct
// [SchemasCollectionResponseResult]
type schemasCollectionResponseResultJSON struct {
	AdditionalInformation apijson.Field
	Application           apijson.Field
	ContentCategories     apijson.Field
	Domain                apijson.Field
	PopularityRank        apijson.Field
	RiskScore             apijson.Field
	RiskTypes             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *SchemasCollectionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information related to the host name.
type SchemasCollectionResponseResultAdditionalInformation struct {
	// Suspected DGA malware family.
	SuspectedMalwareFamily string                                                   `json:"suspected_malware_family"`
	JSON                   schemasCollectionResponseResultAdditionalInformationJSON `json:"-"`
}

// schemasCollectionResponseResultAdditionalInformationJSON contains the JSON
// metadata for the struct [SchemasCollectionResponseResultAdditionalInformation]
type schemasCollectionResponseResultAdditionalInformationJSON struct {
	SuspectedMalwareFamily apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *SchemasCollectionResponseResultAdditionalInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Application that the hostname belongs to.
type SchemasCollectionResponseResultApplication struct {
	ID   int64                                          `json:"id"`
	Name string                                         `json:"name"`
	JSON schemasCollectionResponseResultApplicationJSON `json:"-"`
}

// schemasCollectionResponseResultApplicationJSON contains the JSON metadata for
// the struct [SchemasCollectionResponseResultApplication]
type schemasCollectionResponseResultApplicationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCollectionResponseResultApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasCollectionResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                 `json:"total_count"`
	JSON       schemasCollectionResponseResultInfoJSON `json:"-"`
}

// schemasCollectionResponseResultInfoJSON contains the JSON metadata for the
// struct [SchemasCollectionResponseResultInfo]
type schemasCollectionResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasCollectionResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasCollectionResponseSuccess bool

const (
	SchemasCollectionResponseSuccessTrue SchemasCollectionResponseSuccess = true
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

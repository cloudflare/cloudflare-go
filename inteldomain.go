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

// IntelDomainService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewIntelDomainService] method
// instead.
type IntelDomainService struct {
	Options []option.RequestOption
	Bulks   *IntelDomainBulkService
}

// NewIntelDomainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntelDomainService(opts ...option.RequestOption) (r *IntelDomainService) {
	r = &IntelDomainService{}
	r.Options = opts
	r.Bulks = NewIntelDomainBulkService(opts...)
	return
}

// Get Domain Details
func (r *IntelDomainService) DomainIntelligenceGetDomainDetails(ctx context.Context, accountID string, query IntelDomainDomainIntelligenceGetDomainDetailsParams, opts ...option.RequestOption) (res *IntelDomainDomainIntelligenceGetDomainDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/domain", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelDomainDomainIntelligenceGetDomainDetailsResponse struct {
	// Additional information related to the host name.
	AdditionalInformation IntelDomainDomainIntelligenceGetDomainDetailsResponseAdditionalInformation `json:"additional_information"`
	// Application that the hostname belongs to.
	Application IntelDomainDomainIntelligenceGetDomainDetailsResponseApplication `json:"application"`
	// Current content categories.
	ContentCategories          interface{}                                                                     `json:"content_categories"`
	Domain                     string                                                                          `json:"domain"`
	InheritedContentCategories []IntelDomainDomainIntelligenceGetDomainDetailsResponseInheritedContentCategory `json:"inherited_content_categories"`
	// Domain from which `inherited_content_categories` and `inherited_risk_types` are
	// inherited, if applicable.
	InheritedFrom      string                                                                   `json:"inherited_from"`
	InheritedRiskTypes []IntelDomainDomainIntelligenceGetDomainDetailsResponseInheritedRiskType `json:"inherited_risk_types"`
	// Global Cloudflare 100k ranking for the last 30 days, if available for the
	// hostname. The top ranked domain is 1, the lowest ranked domain is 100,000.
	PopularityRank int64 `json:"popularity_rank"`
	// Specifies a list of references to one or more IP addresses or domain names that
	// the domain name currently resolves to.
	ResolvesToRefs []IntelDomainDomainIntelligenceGetDomainDetailsResponseResolvesToRef `json:"resolves_to_refs"`
	// Hostname risk score, which is a value between 0 (lowest risk) to 1 (highest
	// risk).
	RiskScore float64                                                   `json:"risk_score"`
	RiskTypes interface{}                                               `json:"risk_types"`
	JSON      intelDomainDomainIntelligenceGetDomainDetailsResponseJSON `json:"-"`
}

// intelDomainDomainIntelligenceGetDomainDetailsResponseJSON contains the JSON
// metadata for the struct [IntelDomainDomainIntelligenceGetDomainDetailsResponse]
type intelDomainDomainIntelligenceGetDomainDetailsResponseJSON struct {
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

func (r *IntelDomainDomainIntelligenceGetDomainDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information related to the host name.
type IntelDomainDomainIntelligenceGetDomainDetailsResponseAdditionalInformation struct {
	// Suspected DGA malware family.
	SuspectedMalwareFamily string                                                                         `json:"suspected_malware_family"`
	JSON                   intelDomainDomainIntelligenceGetDomainDetailsResponseAdditionalInformationJSON `json:"-"`
}

// intelDomainDomainIntelligenceGetDomainDetailsResponseAdditionalInformationJSON
// contains the JSON metadata for the struct
// [IntelDomainDomainIntelligenceGetDomainDetailsResponseAdditionalInformation]
type intelDomainDomainIntelligenceGetDomainDetailsResponseAdditionalInformationJSON struct {
	SuspectedMalwareFamily apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IntelDomainDomainIntelligenceGetDomainDetailsResponseAdditionalInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Application that the hostname belongs to.
type IntelDomainDomainIntelligenceGetDomainDetailsResponseApplication struct {
	ID   int64                                                                `json:"id"`
	Name string                                                               `json:"name"`
	JSON intelDomainDomainIntelligenceGetDomainDetailsResponseApplicationJSON `json:"-"`
}

// intelDomainDomainIntelligenceGetDomainDetailsResponseApplicationJSON contains
// the JSON metadata for the struct
// [IntelDomainDomainIntelligenceGetDomainDetailsResponseApplication]
type intelDomainDomainIntelligenceGetDomainDetailsResponseApplicationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainDomainIntelligenceGetDomainDetailsResponseApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainDomainIntelligenceGetDomainDetailsResponseInheritedContentCategory struct {
	ID              int64                                                                             `json:"id"`
	Name            string                                                                            `json:"name"`
	SuperCategoryID int64                                                                             `json:"super_category_id"`
	JSON            intelDomainDomainIntelligenceGetDomainDetailsResponseInheritedContentCategoryJSON `json:"-"`
}

// intelDomainDomainIntelligenceGetDomainDetailsResponseInheritedContentCategoryJSON
// contains the JSON metadata for the struct
// [IntelDomainDomainIntelligenceGetDomainDetailsResponseInheritedContentCategory]
type intelDomainDomainIntelligenceGetDomainDetailsResponseInheritedContentCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IntelDomainDomainIntelligenceGetDomainDetailsResponseInheritedContentCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainDomainIntelligenceGetDomainDetailsResponseInheritedRiskType struct {
	ID              int64                                                                      `json:"id"`
	Name            string                                                                     `json:"name"`
	SuperCategoryID int64                                                                      `json:"super_category_id"`
	JSON            intelDomainDomainIntelligenceGetDomainDetailsResponseInheritedRiskTypeJSON `json:"-"`
}

// intelDomainDomainIntelligenceGetDomainDetailsResponseInheritedRiskTypeJSON
// contains the JSON metadata for the struct
// [IntelDomainDomainIntelligenceGetDomainDetailsResponseInheritedRiskType]
type intelDomainDomainIntelligenceGetDomainDetailsResponseInheritedRiskTypeJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IntelDomainDomainIntelligenceGetDomainDetailsResponseInheritedRiskType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainDomainIntelligenceGetDomainDetailsResponseResolvesToRef struct {
	// STIX 2.1 identifier:
	// https://docs.oasis-open.org/cti/stix/v2.1/cs02/stix-v2.1-cs02.html#_64yvzeku5a5c
	ID string `json:"id"`
	// IP address or domain name.
	Value string                                                                 `json:"value"`
	JSON  intelDomainDomainIntelligenceGetDomainDetailsResponseResolvesToRefJSON `json:"-"`
}

// intelDomainDomainIntelligenceGetDomainDetailsResponseResolvesToRefJSON contains
// the JSON metadata for the struct
// [IntelDomainDomainIntelligenceGetDomainDetailsResponseResolvesToRef]
type intelDomainDomainIntelligenceGetDomainDetailsResponseResolvesToRefJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainDomainIntelligenceGetDomainDetailsResponseResolvesToRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainDomainIntelligenceGetDomainDetailsParams struct {
	Domain param.Field[string] `query:"domain"`
}

// URLQuery serializes [IntelDomainDomainIntelligenceGetDomainDetailsParams]'s
// query parameters as `url.Values`.
func (r IntelDomainDomainIntelligenceGetDomainDetailsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelope struct {
	Errors   []IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelDomainDomainIntelligenceGetDomainDetailsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeJSON    `json:"-"`
}

// intelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelope]
type intelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeErrors struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    intelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeErrorsJSON `json:"-"`
}

// intelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeErrors]
type intelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeMessages struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    intelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeMessagesJSON `json:"-"`
}

// intelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeMessages]
type intelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeSuccess bool

const (
	IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeSuccessTrue IntelDomainDomainIntelligenceGetDomainDetailsResponseEnvelopeSuccess = true
)

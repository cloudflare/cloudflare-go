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
func (r *AccountIntelDomainService) DomainIntelligenceGetDomainDetails(ctx context.Context, accountIdentifier string, query AccountIntelDomainDomainIntelligenceGetDomainDetailsParams, opts ...option.RequestOption) (res *DomainSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/domain", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DomainSingleResponse struct {
	Errors   []DomainSingleResponseError   `json:"errors"`
	Messages []DomainSingleResponseMessage `json:"messages"`
	Result   DomainSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success DomainSingleResponseSuccess `json:"success"`
	JSON    domainSingleResponseJSON    `json:"-"`
}

// domainSingleResponseJSON contains the JSON metadata for the struct
// [DomainSingleResponse]
type domainSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainSingleResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    domainSingleResponseErrorJSON `json:"-"`
}

// domainSingleResponseErrorJSON contains the JSON metadata for the struct
// [DomainSingleResponseError]
type domainSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainSingleResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    domainSingleResponseMessageJSON `json:"-"`
}

// domainSingleResponseMessageJSON contains the JSON metadata for the struct
// [DomainSingleResponseMessage]
type domainSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainSingleResponseResult struct {
	// Additional information related to the host name.
	AdditionalInformation DomainSingleResponseResultAdditionalInformation `json:"additional_information"`
	// Application that the hostname belongs to.
	Application DomainSingleResponseResultApplication `json:"application"`
	// Current content categories.
	ContentCategories interface{} `json:"content_categories"`
	Domain            string      `json:"domain"`
	// Global Cloudflare 100k ranking for the last 30 days, if available for the
	// hostname. The top ranked domain is 1, the lowest ranked domain is 100,000.
	PopularityRank int64 `json:"popularity_rank"`
	// Specifies a list of references to one or more IP addresses or domain names that
	// the domain name currently resolves to.
	ResolvesToRefs []DomainSingleResponseResultResolvesToRef `json:"resolves_to_refs"`
	// Hostname risk score, which is a value between 0 (lowest risk) to 1 (highest
	// risk).
	RiskScore float64                        `json:"risk_score"`
	RiskTypes interface{}                    `json:"risk_types"`
	JSON      domainSingleResponseResultJSON `json:"-"`
}

// domainSingleResponseResultJSON contains the JSON metadata for the struct
// [DomainSingleResponseResult]
type domainSingleResponseResultJSON struct {
	AdditionalInformation apijson.Field
	Application           apijson.Field
	ContentCategories     apijson.Field
	Domain                apijson.Field
	PopularityRank        apijson.Field
	ResolvesToRefs        apijson.Field
	RiskScore             apijson.Field
	RiskTypes             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DomainSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information related to the host name.
type DomainSingleResponseResultAdditionalInformation struct {
	// Suspected DGA malware family.
	SuspectedMalwareFamily string                                              `json:"suspected_malware_family"`
	JSON                   domainSingleResponseResultAdditionalInformationJSON `json:"-"`
}

// domainSingleResponseResultAdditionalInformationJSON contains the JSON metadata
// for the struct [DomainSingleResponseResultAdditionalInformation]
type domainSingleResponseResultAdditionalInformationJSON struct {
	SuspectedMalwareFamily apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *DomainSingleResponseResultAdditionalInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Application that the hostname belongs to.
type DomainSingleResponseResultApplication struct {
	ID   int64                                     `json:"id"`
	Name string                                    `json:"name"`
	JSON domainSingleResponseResultApplicationJSON `json:"-"`
}

// domainSingleResponseResultApplicationJSON contains the JSON metadata for the
// struct [DomainSingleResponseResultApplication]
type domainSingleResponseResultApplicationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainSingleResponseResultApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainSingleResponseResultResolvesToRef struct {
	// STIX 2.1 identifier:
	// https://docs.oasis-open.org/cti/stix/v2.1/cs02/stix-v2.1-cs02.html#_64yvzeku5a5c
	ID string `json:"id"`
	// IP address or domain name.
	Value string                                      `json:"value"`
	JSON  domainSingleResponseResultResolvesToRefJSON `json:"-"`
}

// domainSingleResponseResultResolvesToRefJSON contains the JSON metadata for the
// struct [DomainSingleResponseResultResolvesToRef]
type domainSingleResponseResultResolvesToRefJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainSingleResponseResultResolvesToRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DomainSingleResponseSuccess bool

const (
	DomainSingleResponseSuccessTrue DomainSingleResponseSuccess = true
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

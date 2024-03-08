// File generated from our OpenAPI spec by Stainless.

package intel

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

// DomainService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDomainService] method instead.
type DomainService struct {
	Options []option.RequestOption
	Bulks   *DomainBulkService
}

// NewDomainService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomainService(opts ...option.RequestOption) (r *DomainService) {
	r = &DomainService{}
	r.Options = opts
	r.Bulks = NewDomainBulkService(opts...)
	return
}

// Get Domain Details
func (r *DomainService) Get(ctx context.Context, params DomainGetParams, opts ...option.RequestOption) (res *DomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DomainGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/domain", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DomainGetResponse struct {
	// Additional information related to the host name.
	AdditionalInformation DomainGetResponseAdditionalInformation `json:"additional_information"`
	// Application that the hostname belongs to.
	Application DomainGetResponseApplication `json:"application"`
	// Current content categories.
	ContentCategories          interface{}                                 `json:"content_categories"`
	Domain                     string                                      `json:"domain"`
	InheritedContentCategories []DomainGetResponseInheritedContentCategory `json:"inherited_content_categories"`
	// Domain from which `inherited_content_categories` and `inherited_risk_types` are
	// inherited, if applicable.
	InheritedFrom      string                               `json:"inherited_from"`
	InheritedRiskTypes []DomainGetResponseInheritedRiskType `json:"inherited_risk_types"`
	// Global Cloudflare 100k ranking for the last 30 days, if available for the
	// hostname. The top ranked domain is 1, the lowest ranked domain is 100,000.
	PopularityRank int64 `json:"popularity_rank"`
	// Specifies a list of references to one or more IP addresses or domain names that
	// the domain name currently resolves to.
	ResolvesToRefs []DomainGetResponseResolvesToRef `json:"resolves_to_refs"`
	// Hostname risk score, which is a value between 0 (lowest risk) to 1 (highest
	// risk).
	RiskScore float64               `json:"risk_score"`
	RiskTypes interface{}           `json:"risk_types"`
	JSON      domainGetResponseJSON `json:"-"`
}

// domainGetResponseJSON contains the JSON metadata for the struct
// [DomainGetResponse]
type domainGetResponseJSON struct {
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

func (r *DomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseJSON) RawJSON() string {
	return r.raw
}

// Additional information related to the host name.
type DomainGetResponseAdditionalInformation struct {
	// Suspected DGA malware family.
	SuspectedMalwareFamily string                                     `json:"suspected_malware_family"`
	JSON                   domainGetResponseAdditionalInformationJSON `json:"-"`
}

// domainGetResponseAdditionalInformationJSON contains the JSON metadata for the
// struct [DomainGetResponseAdditionalInformation]
type domainGetResponseAdditionalInformationJSON struct {
	SuspectedMalwareFamily apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *DomainGetResponseAdditionalInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseAdditionalInformationJSON) RawJSON() string {
	return r.raw
}

// Application that the hostname belongs to.
type DomainGetResponseApplication struct {
	ID   int64                            `json:"id"`
	Name string                           `json:"name"`
	JSON domainGetResponseApplicationJSON `json:"-"`
}

// domainGetResponseApplicationJSON contains the JSON metadata for the struct
// [DomainGetResponseApplication]
type domainGetResponseApplicationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseApplicationJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseInheritedContentCategory struct {
	ID              int64                                         `json:"id"`
	Name            string                                        `json:"name"`
	SuperCategoryID int64                                         `json:"super_category_id"`
	JSON            domainGetResponseInheritedContentCategoryJSON `json:"-"`
}

// domainGetResponseInheritedContentCategoryJSON contains the JSON metadata for the
// struct [DomainGetResponseInheritedContentCategory]
type domainGetResponseInheritedContentCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DomainGetResponseInheritedContentCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseInheritedContentCategoryJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseInheritedRiskType struct {
	ID              int64                                  `json:"id"`
	Name            string                                 `json:"name"`
	SuperCategoryID int64                                  `json:"super_category_id"`
	JSON            domainGetResponseInheritedRiskTypeJSON `json:"-"`
}

// domainGetResponseInheritedRiskTypeJSON contains the JSON metadata for the struct
// [DomainGetResponseInheritedRiskType]
type domainGetResponseInheritedRiskTypeJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DomainGetResponseInheritedRiskType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseInheritedRiskTypeJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseResolvesToRef struct {
	// STIX 2.1 identifier:
	// https://docs.oasis-open.org/cti/stix/v2.1/cs02/stix-v2.1-cs02.html#_64yvzeku5a5c
	ID string `json:"id"`
	// IP address or domain name.
	Value string                             `json:"value"`
	JSON  domainGetResponseResolvesToRefJSON `json:"-"`
}

// domainGetResponseResolvesToRefJSON contains the JSON metadata for the struct
// [DomainGetResponseResolvesToRef]
type domainGetResponseResolvesToRefJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseResolvesToRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseResolvesToRefJSON) RawJSON() string {
	return r.raw
}

type DomainGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Domain    param.Field[string] `query:"domain"`
}

// URLQuery serializes [DomainGetParams]'s query parameters as `url.Values`.
func (r DomainGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DomainGetResponseEnvelope struct {
	Errors   []DomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DomainGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DomainGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    domainGetResponseEnvelopeJSON    `json:"-"`
}

// domainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelope]
type domainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    domainGetResponseEnvelopeErrorsJSON `json:"-"`
}

// domainGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelopeErrors]
type domainGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DomainGetResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    domainGetResponseEnvelopeMessagesJSON `json:"-"`
}

// domainGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelopeMessages]
type domainGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DomainGetResponseEnvelopeSuccess bool

const (
	DomainGetResponseEnvelopeSuccessTrue DomainGetResponseEnvelopeSuccess = true
)

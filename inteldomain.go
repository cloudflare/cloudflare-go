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
func (r *IntelDomainService) Get(ctx context.Context, params IntelDomainGetParams, opts ...option.RequestOption) (res *IntelDomain, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelDomainGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/domain", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelDomain struct {
	// Additional information related to the host name.
	AdditionalInformation IntelDomainAdditionalInformation `json:"additional_information"`
	// Application that the hostname belongs to.
	Application IntelDomainApplication `json:"application"`
	// Current content categories.
	ContentCategories          interface{}                           `json:"content_categories"`
	Domain                     string                                `json:"domain"`
	InheritedContentCategories []IntelDomainInheritedContentCategory `json:"inherited_content_categories"`
	// Domain from which `inherited_content_categories` and `inherited_risk_types` are
	// inherited, if applicable.
	InheritedFrom      string                         `json:"inherited_from"`
	InheritedRiskTypes []IntelDomainInheritedRiskType `json:"inherited_risk_types"`
	// Global Cloudflare 100k ranking for the last 30 days, if available for the
	// hostname. The top ranked domain is 1, the lowest ranked domain is 100,000.
	PopularityRank int64 `json:"popularity_rank"`
	// Specifies a list of references to one or more IP addresses or domain names that
	// the domain name currently resolves to.
	ResolvesToRefs []IntelDomainResolvesToRef `json:"resolves_to_refs"`
	// Hostname risk score, which is a value between 0 (lowest risk) to 1 (highest
	// risk).
	RiskScore float64         `json:"risk_score"`
	RiskTypes interface{}     `json:"risk_types"`
	JSON      intelDomainJSON `json:"-"`
}

// intelDomainJSON contains the JSON metadata for the struct [IntelDomain]
type intelDomainJSON struct {
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

func (r *IntelDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Additional information related to the host name.
type IntelDomainAdditionalInformation struct {
	// Suspected DGA malware family.
	SuspectedMalwareFamily string                               `json:"suspected_malware_family"`
	JSON                   intelDomainAdditionalInformationJSON `json:"-"`
}

// intelDomainAdditionalInformationJSON contains the JSON metadata for the struct
// [IntelDomainAdditionalInformation]
type intelDomainAdditionalInformationJSON struct {
	SuspectedMalwareFamily apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *IntelDomainAdditionalInformation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Application that the hostname belongs to.
type IntelDomainApplication struct {
	ID   int64                      `json:"id"`
	Name string                     `json:"name"`
	JSON intelDomainApplicationJSON `json:"-"`
}

// intelDomainApplicationJSON contains the JSON metadata for the struct
// [IntelDomainApplication]
type intelDomainApplicationJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainApplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainInheritedContentCategory struct {
	ID              int64                                   `json:"id"`
	Name            string                                  `json:"name"`
	SuperCategoryID int64                                   `json:"super_category_id"`
	JSON            intelDomainInheritedContentCategoryJSON `json:"-"`
}

// intelDomainInheritedContentCategoryJSON contains the JSON metadata for the
// struct [IntelDomainInheritedContentCategory]
type intelDomainInheritedContentCategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IntelDomainInheritedContentCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainInheritedRiskType struct {
	ID              int64                            `json:"id"`
	Name            string                           `json:"name"`
	SuperCategoryID int64                            `json:"super_category_id"`
	JSON            intelDomainInheritedRiskTypeJSON `json:"-"`
}

// intelDomainInheritedRiskTypeJSON contains the JSON metadata for the struct
// [IntelDomainInheritedRiskType]
type intelDomainInheritedRiskTypeJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *IntelDomainInheritedRiskType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainResolvesToRef struct {
	// STIX 2.1 identifier:
	// https://docs.oasis-open.org/cti/stix/v2.1/cs02/stix-v2.1-cs02.html#_64yvzeku5a5c
	ID string `json:"id"`
	// IP address or domain name.
	Value string                       `json:"value"`
	JSON  intelDomainResolvesToRefJSON `json:"-"`
}

// intelDomainResolvesToRefJSON contains the JSON metadata for the struct
// [IntelDomainResolvesToRef]
type intelDomainResolvesToRefJSON struct {
	ID          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainResolvesToRef) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Domain    param.Field[string] `query:"domain"`
}

// URLQuery serializes [IntelDomainGetParams]'s query parameters as `url.Values`.
func (r IntelDomainGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntelDomainGetResponseEnvelope struct {
	Errors   []IntelDomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelDomainGetResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelDomain                              `json:"result,required"`
	// Whether the API call was successful
	Success IntelDomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelDomainGetResponseEnvelopeJSON    `json:"-"`
}

// intelDomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IntelDomainGetResponseEnvelope]
type intelDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    intelDomainGetResponseEnvelopeErrorsJSON `json:"-"`
}

// intelDomainGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IntelDomainGetResponseEnvelopeErrors]
type intelDomainGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelDomainGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    intelDomainGetResponseEnvelopeMessagesJSON `json:"-"`
}

// intelDomainGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IntelDomainGetResponseEnvelopeMessages]
type intelDomainGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelDomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelDomainGetResponseEnvelopeSuccess bool

const (
	IntelDomainGetResponseEnvelopeSuccessTrue IntelDomainGetResponseEnvelopeSuccess = true
)

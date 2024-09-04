// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
  "context"
  "errors"
  "fmt"
  "net/http"
  "net/url"

  "github.com/cloudflare/cloudflare-go/v2/internal/apijson"
  "github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
  "github.com/cloudflare/cloudflare-go/v2/internal/param"
  "github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
  "github.com/cloudflare/cloudflare-go/v2/option"
  "github.com/cloudflare/cloudflare-go/v2/shared"
)

// DomainService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainService] method instead.
type DomainService struct {
Options []option.RequestOption
Bulks *DomainBulkService
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
func (r *DomainService) Get(ctx context.Context, params DomainGetParams, opts ...option.RequestOption) (res *Domain, err error) {
  var env DomainGetResponseEnvelope
  opts = append(r.Options[:], opts...)
  if params.AccountID.Value == "" {
    err = errors.New("missing required account_id parameter")
    return
  }
  path := fmt.Sprintf("accounts/%s/intel/domain", params.AccountID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

type Domain struct {
// Additional information related to the host name.
AdditionalInformation DomainAdditionalInformation `json:"additional_information"`
// Application that the hostname belongs to.
Application DomainApplication `json:"application"`
// Current content categories.
ContentCategories []interface{} `json:"content_categories"`
Domain string `json:"domain"`
InheritedContentCategories []DomainInheritedContentCategory `json:"inherited_content_categories"`
// Domain from which `inherited_content_categories` and `inherited_risk_types` are
// inherited, if applicable.
InheritedFrom string `json:"inherited_from"`
InheritedRiskTypes []DomainInheritedRiskType `json:"inherited_risk_types"`
// Global Cloudflare 100k ranking for the last 30 days, if available for the
// hostname. The top ranked domain is 1, the lowest ranked domain is 100,000.
PopularityRank int64 `json:"popularity_rank"`
// Specifies a list of references to one or more IP addresses or domain names that
// the domain name currently resolves to.
ResolvesToRefs []DomainResolvesToRef `json:"resolves_to_refs"`
// Hostname risk score, which is a value between 0 (lowest risk) to 1 (highest
// risk).
RiskScore float64 `json:"risk_score"`
RiskTypes []interface{} `json:"risk_types"`
JSON domainJSON `json:"-"`
}

// domainJSON contains the JSON metadata for the struct [Domain]
type domainJSON struct {
AdditionalInformation apijson.Field
Application apijson.Field
ContentCategories apijson.Field
Domain apijson.Field
InheritedContentCategories apijson.Field
InheritedFrom apijson.Field
InheritedRiskTypes apijson.Field
PopularityRank apijson.Field
ResolvesToRefs apijson.Field
RiskScore apijson.Field
RiskTypes apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *Domain) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r domainJSON) RawJSON() (string) {
  return r.raw
}

// Additional information related to the host name.
type DomainAdditionalInformation struct {
// Suspected DGA malware family.
SuspectedMalwareFamily string `json:"suspected_malware_family"`
JSON domainAdditionalInformationJSON `json:"-"`
}

// domainAdditionalInformationJSON contains the JSON metadata for the struct
// [DomainAdditionalInformation]
type domainAdditionalInformationJSON struct {
SuspectedMalwareFamily apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DomainAdditionalInformation) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r domainAdditionalInformationJSON) RawJSON() (string) {
  return r.raw
}

// Application that the hostname belongs to.
type DomainApplication struct {
ID int64 `json:"id"`
Name string `json:"name"`
JSON domainApplicationJSON `json:"-"`
}

// domainApplicationJSON contains the JSON metadata for the struct
// [DomainApplication]
type domainApplicationJSON struct {
ID apijson.Field
Name apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DomainApplication) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r domainApplicationJSON) RawJSON() (string) {
  return r.raw
}

type DomainInheritedContentCategory struct {
ID int64 `json:"id"`
Name string `json:"name"`
SuperCategoryID int64 `json:"super_category_id"`
JSON domainInheritedContentCategoryJSON `json:"-"`
}

// domainInheritedContentCategoryJSON contains the JSON metadata for the struct
// [DomainInheritedContentCategory]
type domainInheritedContentCategoryJSON struct {
ID apijson.Field
Name apijson.Field
SuperCategoryID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DomainInheritedContentCategory) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r domainInheritedContentCategoryJSON) RawJSON() (string) {
  return r.raw
}

type DomainInheritedRiskType struct {
ID int64 `json:"id"`
Name string `json:"name"`
SuperCategoryID int64 `json:"super_category_id"`
JSON domainInheritedRiskTypeJSON `json:"-"`
}

// domainInheritedRiskTypeJSON contains the JSON metadata for the struct
// [DomainInheritedRiskType]
type domainInheritedRiskTypeJSON struct {
ID apijson.Field
Name apijson.Field
SuperCategoryID apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DomainInheritedRiskType) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r domainInheritedRiskTypeJSON) RawJSON() (string) {
  return r.raw
}

type DomainResolvesToRef struct {
// STIX 2.1 identifier:
// https://docs.oasis-open.org/cti/stix/v2.1/cs02/stix-v2.1-cs02.html#_64yvzeku5a5c
ID string `json:"id"`
// IP address or domain name.
Value string `json:"value"`
JSON domainResolvesToRefJSON `json:"-"`
}

// domainResolvesToRefJSON contains the JSON metadata for the struct
// [DomainResolvesToRef]
type domainResolvesToRefJSON struct {
ID apijson.Field
Value apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DomainResolvesToRef) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r domainResolvesToRefJSON) RawJSON() (string) {
  return r.raw
}

type DomainGetParams struct {
// Identifier
AccountID param.Field[string] `path:"account_id,required"`
Domain param.Field[string] `query:"domain"`
}

// URLQuery serializes [DomainGetParams]'s query parameters as `url.Values`.
func (r DomainGetParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

type DomainGetResponseEnvelope struct {
Errors []shared.ResponseInfo `json:"errors,required"`
Messages []shared.ResponseInfo `json:"messages,required"`
// Whether the API call was successful
Success DomainGetResponseEnvelopeSuccess `json:"success,required"`
Result Domain `json:"result"`
JSON domainGetResponseEnvelopeJSON `json:"-"`
}

// domainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelope]
type domainGetResponseEnvelopeJSON struct {
Errors apijson.Field
Messages apijson.Field
Success apijson.Field
Result apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type DomainGetResponseEnvelopeSuccess bool

const (
  DomainGetResponseEnvelopeSuccessTrue DomainGetResponseEnvelopeSuccess = true
)

func (r DomainGetResponseEnvelopeSuccess) IsKnown() (bool) {
  switch r {
  case DomainGetResponseEnvelopeSuccessTrue:
      return true
  }
  return false
}

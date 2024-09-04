// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
  "context"
  "net/http"
  "net/url"

  "github.com/cloudflare/cloudflare-go/v2/internal/apijson"
  "github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
  "github.com/cloudflare/cloudflare-go/v2/internal/param"
  "github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
  "github.com/cloudflare/cloudflare-go/v2/option"
)

// RobotsTXTService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRobotsTXTService] method instead.
type RobotsTXTService struct {
Options []option.RequestOption
Top *RobotsTXTTopService
}

// NewRobotsTXTService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRobotsTXTService(opts ...option.RequestOption) (r *RobotsTXTService) {
  r = &RobotsTXTService{}
  r.Options = opts
  r.Top = NewRobotsTXTTopService(opts...)
  return
}

// Get the top User-Agents on robots.txt files by domain.
func (r *RobotsTXTService) Domains(ctx context.Context, query RobotsTXTDomainsParams, opts ...option.RequestOption) (res *RobotsTXTDomainsResponse, err error) {
  var env RobotsTXTDomainsResponseEnvelope
  opts = append(r.Options[:], opts...)
  path := "radar/robots_txt/domains"
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

type RobotsTXTDomainsResponse struct {
Date string `json:"date,required"`
Domains []RobotsTXTDomainsResponseDomain `json:"domains,required"`
UserAgents []string `json:"userAgents,required"`
JSON robotsTXTDomainsResponseJSON `json:"-"`
}

// robotsTXTDomainsResponseJSON contains the JSON metadata for the struct
// [RobotsTXTDomainsResponse]
type robotsTXTDomainsResponseJSON struct {
Date apijson.Field
Domains apijson.Field
UserAgents apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTDomainsResponse) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTDomainsResponseJSON) RawJSON() (string) {
  return r.raw
}

type RobotsTXTDomainsResponseDomain struct {
string `json:"*,required"`
Amazonbot string `json:"amazonbot,required"`
CategoriesParent string `json:"categories_parent,required"`
CategoriesSub string `json:"categories_sub,required"`
Domain string `json:"domain,required"`
JSON robotsTXTDomainsResponseDomainJSON `json:"-"`
}

// robotsTXTDomainsResponseDomainJSON contains the JSON metadata for the struct
// [RobotsTXTDomainsResponseDomain]
type robotsTXTDomainsResponseDomainJSON struct {
apijson.Field
Amazonbot apijson.Field
CategoriesParent apijson.Field
CategoriesSub apijson.Field
Domain apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTDomainsResponseDomain) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTDomainsResponseDomainJSON) RawJSON() (string) {
  return r.raw
}

type RobotsTXTDomainsParams struct {
// Filter domains by category
DomainCategory param.Field[string] `query:"domainCategory"`
// Filter domains by name
DomainName param.Field[string] `query:"domainName"`
// Format results are returned in.
Format param.Field[RobotsTXTDomainsParamsFormat] `query:"format"`
// Limit the number of objects in the response.
Limit param.Field[int64] `query:"limit"`
// Number of objects to skip before grabbing results.
Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [RobotsTXTDomainsParams]'s query parameters as `url.Values`.
func (r RobotsTXTDomainsParams) URLQuery() (v url.Values) {
  return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
    ArrayFormat: apiquery.ArrayQueryFormatRepeat,
    NestedFormat: apiquery.NestedQueryFormatDots,
  })
}

// Format results are returned in.
type RobotsTXTDomainsParamsFormat string

const (
  RobotsTXTDomainsParamsFormatJson RobotsTXTDomainsParamsFormat = "JSON"
  RobotsTXTDomainsParamsFormatCsv RobotsTXTDomainsParamsFormat = "CSV"
)

func (r RobotsTXTDomainsParamsFormat) IsKnown() (bool) {
  switch r {
  case RobotsTXTDomainsParamsFormatJson, RobotsTXTDomainsParamsFormatCsv:
      return true
  }
  return false
}

type RobotsTXTDomainsResponseEnvelope struct {
Result RobotsTXTDomainsResponse `json:"result,required"`
ResultInfo RobotsTXTDomainsResponseEnvelopeResultInfo `json:"resultInfo,required"`
Success bool `json:"success,required"`
JSON robotsTXTDomainsResponseEnvelopeJSON `json:"-"`
}

// robotsTXTDomainsResponseEnvelopeJSON contains the JSON metadata for the struct
// [RobotsTXTDomainsResponseEnvelope]
type robotsTXTDomainsResponseEnvelopeJSON struct {
Result apijson.Field
ResultInfo apijson.Field
Success apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTDomainsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTDomainsResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

type RobotsTXTDomainsResponseEnvelopeResultInfo struct {
Limit int64 `json:"limit,required"`
Offset int64 `json:"offset,required"`
Total int64 `json:"total,required"`
JSON robotsTXTDomainsResponseEnvelopeResultInfoJSON `json:"-"`
}

// robotsTXTDomainsResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [RobotsTXTDomainsResponseEnvelopeResultInfo]
type robotsTXTDomainsResponseEnvelopeResultInfoJSON struct {
Limit apijson.Field
Offset apijson.Field
Total apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *RobotsTXTDomainsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r robotsTXTDomainsResponseEnvelopeResultInfoJSON) RawJSON() (string) {
  return r.raw
}

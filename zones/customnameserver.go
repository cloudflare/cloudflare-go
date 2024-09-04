// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
  "context"
  "errors"
  "fmt"
  "net/http"

  "github.com/cloudflare/cloudflare-go/v2/internal/apijson"
  "github.com/cloudflare/cloudflare-go/v2/internal/param"
  "github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
  "github.com/cloudflare/cloudflare-go/v2/option"
  "github.com/cloudflare/cloudflare-go/v2/shared"
)

// CustomNameserverService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomNameserverService] method instead.
type CustomNameserverService struct {
Options []option.RequestOption
}

// NewCustomNameserverService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomNameserverService(opts ...option.RequestOption) (r *CustomNameserverService) {
  r = &CustomNameserverService{}
  r.Options = opts
  return
}

// Set metadata for account-level custom nameservers on a zone.
//
// If you would like new zones in the account to use account custom nameservers by
// default, use PUT /accounts/:identifier to set the account setting
// use_account_custom_ns_by_default to true.
func (r *CustomNameserverService) Update(ctx context.Context, params CustomNameserverUpdateParams, opts ...option.RequestOption) (res *[]string, err error) {
  var env CustomNameserverUpdateResponseEnvelope
  opts = append(r.Options[:], opts...)
  if params.ZoneID.Value == "" {
    err = errors.New("missing required zone_id parameter")
    return
  }
  path := fmt.Sprintf("zones/%s/custom_ns", params.ZoneID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Get metadata for account-level custom nameservers on a zone.
func (r *CustomNameserverService) Get(ctx context.Context, query CustomNameserverGetParams, opts ...option.RequestOption) (res *[]CustomNameserverGetResponse, err error) {
  var env CustomNameserverGetResponseEnvelope
  opts = append(r.Options[:], opts...)
  if query.ZoneID.Value == "" {
    err = errors.New("missing required zone_id parameter")
    return
  }
  path := fmt.Sprintf("zones/%s/custom_ns", query.ZoneID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

type CustomNameserverGetResponse = interface{}

type CustomNameserverUpdateParams struct {
// Identifier
ZoneID param.Field[string] `path:"zone_id,required"`
// Whether zone uses account-level custom nameservers.
Enabled param.Field[bool] `json:"enabled"`
// The number of the name server set to assign to the zone.
NSSet param.Field[float64] `json:"ns_set"`
}

func (r CustomNameserverUpdateParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type CustomNameserverUpdateResponseEnvelope struct {
Errors []shared.ResponseInfo `json:"errors,required"`
Messages []shared.ResponseInfo `json:"messages,required"`
// Whether the API call was successful
Success CustomNameserverUpdateResponseEnvelopeSuccess `json:"success,required"`
Result []string `json:"result,nullable"`
ResultInfo CustomNameserverUpdateResponseEnvelopeResultInfo `json:"result_info"`
JSON customNameserverUpdateResponseEnvelopeJSON `json:"-"`
}

// customNameserverUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomNameserverUpdateResponseEnvelope]
type customNameserverUpdateResponseEnvelopeJSON struct {
Errors apijson.Field
Messages apijson.Field
Success apijson.Field
Result apijson.Field
ResultInfo apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverUpdateResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type CustomNameserverUpdateResponseEnvelopeSuccess bool

const (
  CustomNameserverUpdateResponseEnvelopeSuccessTrue CustomNameserverUpdateResponseEnvelopeSuccess = true
)

func (r CustomNameserverUpdateResponseEnvelopeSuccess) IsKnown() (bool) {
  switch r {
  case CustomNameserverUpdateResponseEnvelopeSuccessTrue:
      return true
  }
  return false
}

type CustomNameserverUpdateResponseEnvelopeResultInfo struct {
// Total number of results for the requested service
Count float64 `json:"count"`
// Current page within paginated list of results
Page float64 `json:"page"`
// Number of results per page of results
PerPage float64 `json:"per_page"`
// Total results available without any search parameters
TotalCount float64 `json:"total_count"`
JSON customNameserverUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// customNameserverUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [CustomNameserverUpdateResponseEnvelopeResultInfo]
type customNameserverUpdateResponseEnvelopeResultInfoJSON struct {
Count apijson.Field
Page apijson.Field
PerPage apijson.Field
TotalCount apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverUpdateResponseEnvelopeResultInfoJSON) RawJSON() (string) {
  return r.raw
}

type CustomNameserverGetParams struct {
// Identifier
ZoneID param.Field[string] `path:"zone_id,required"`
}

type CustomNameserverGetResponseEnvelope struct {
Errors []shared.ResponseInfo `json:"errors,required"`
Messages []shared.ResponseInfo `json:"messages,required"`
// Whether the API call was successful
Success CustomNameserverGetResponseEnvelopeSuccess `json:"success,required"`
// Whether zone uses account-level custom nameservers.
Enabled bool `json:"enabled"`
// The number of the name server set to assign to the zone.
NSSet float64 `json:"ns_set"`
Result []CustomNameserverGetResponse `json:"result,nullable"`
ResultInfo CustomNameserverGetResponseEnvelopeResultInfo `json:"result_info"`
JSON customNameserverGetResponseEnvelopeJSON `json:"-"`
}

// customNameserverGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomNameserverGetResponseEnvelope]
type customNameserverGetResponseEnvelopeJSON struct {
Errors apijson.Field
Messages apijson.Field
Success apijson.Field
Enabled apijson.Field
NSSet apijson.Field
Result apijson.Field
ResultInfo apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverGetResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type CustomNameserverGetResponseEnvelopeSuccess bool

const (
  CustomNameserverGetResponseEnvelopeSuccessTrue CustomNameserverGetResponseEnvelopeSuccess = true
)

func (r CustomNameserverGetResponseEnvelopeSuccess) IsKnown() (bool) {
  switch r {
  case CustomNameserverGetResponseEnvelopeSuccessTrue:
      return true
  }
  return false
}

type CustomNameserverGetResponseEnvelopeResultInfo struct {
// Total number of results for the requested service
Count float64 `json:"count"`
// Current page within paginated list of results
Page float64 `json:"page"`
// Number of results per page of results
PerPage float64 `json:"per_page"`
// Total results available without any search parameters
TotalCount float64 `json:"total_count"`
JSON customNameserverGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// customNameserverGetResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [CustomNameserverGetResponseEnvelopeResultInfo]
type customNameserverGetResponseEnvelopeResultInfoJSON struct {
Count apijson.Field
Page apijson.Field
PerPage apijson.Field
TotalCount apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverGetResponseEnvelopeResultInfoJSON) RawJSON() (string) {
  return r.raw
}

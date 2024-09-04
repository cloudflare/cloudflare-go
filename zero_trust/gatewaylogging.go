// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// GatewayLoggingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayLoggingService] method instead.
type GatewayLoggingService struct {
Options []option.RequestOption
}

// NewGatewayLoggingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayLoggingService(opts ...option.RequestOption) (r *GatewayLoggingService) {
  r = &GatewayLoggingService{}
  r.Options = opts
  return
}

// Updates logging settings for the current Zero Trust account.
func (r *GatewayLoggingService) Update(ctx context.Context, params GatewayLoggingUpdateParams, opts ...option.RequestOption) (res *LoggingSetting, err error) {
  var env GatewayLoggingUpdateResponseEnvelope
  opts = append(r.Options[:], opts...)
  if params.AccountID.Value == "" {
    err = errors.New("missing required account_id parameter")
    return
  }
  path := fmt.Sprintf("accounts/%s/gateway/logging", params.AccountID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

// Fetches the current logging settings for Zero Trust account.
func (r *GatewayLoggingService) Get(ctx context.Context, query GatewayLoggingGetParams, opts ...option.RequestOption) (res *LoggingSetting, err error) {
  var env GatewayLoggingGetResponseEnvelope
  opts = append(r.Options[:], opts...)
  if query.AccountID.Value == "" {
    err = errors.New("missing required account_id parameter")
    return
  }
  path := fmt.Sprintf("accounts/%s/gateway/logging", query.AccountID)
  err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
  if err != nil {
    return
  }
  res = &env.Result
  return
}

type LoggingSetting struct {
// Redact personally identifiable information from activity logging (PII fields
// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
RedactPii bool `json:"redact_pii"`
// Logging settings by rule type.
SettingsByRuleType LoggingSettingSettingsByRuleType `json:"settings_by_rule_type"`
JSON loggingSettingJSON `json:"-"`
}

// loggingSettingJSON contains the JSON metadata for the struct [LoggingSetting]
type loggingSettingJSON struct {
RedactPii apijson.Field
SettingsByRuleType apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *LoggingSetting) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r loggingSettingJSON) RawJSON() (string) {
  return r.raw
}

// Logging settings by rule type.
type LoggingSettingSettingsByRuleType struct {
// Logging settings for DNS firewall.
DNS interface{} `json:"dns"`
// Logging settings for HTTP/HTTPS firewall.
HTTP interface{} `json:"http"`
// Logging settings for Network firewall.
L4 interface{} `json:"l4"`
JSON loggingSettingSettingsByRuleTypeJSON `json:"-"`
}

// loggingSettingSettingsByRuleTypeJSON contains the JSON metadata for the struct
// [LoggingSettingSettingsByRuleType]
type loggingSettingSettingsByRuleTypeJSON struct {
DNS apijson.Field
HTTP apijson.Field
L4 apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *LoggingSettingSettingsByRuleType) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r loggingSettingSettingsByRuleTypeJSON) RawJSON() (string) {
  return r.raw
}

type LoggingSettingParam struct {
// Redact personally identifiable information from activity logging (PII fields
// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
RedactPii param.Field[bool] `json:"redact_pii"`
// Logging settings by rule type.
SettingsByRuleType param.Field[LoggingSettingSettingsByRuleTypeParam] `json:"settings_by_rule_type"`
}

func (r LoggingSettingParam) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

// Logging settings by rule type.
type LoggingSettingSettingsByRuleTypeParam struct {
// Logging settings for DNS firewall.
DNS param.Field[interface{}] `json:"dns"`
// Logging settings for HTTP/HTTPS firewall.
HTTP param.Field[interface{}] `json:"http"`
// Logging settings for Network firewall.
L4 param.Field[interface{}] `json:"l4"`
}

func (r LoggingSettingSettingsByRuleTypeParam) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r)
}

type GatewayLoggingUpdateParams struct {
AccountID param.Field[string] `path:"account_id,required"`
LoggingSetting LoggingSettingParam `json:"logging_setting,required"`
}

func (r GatewayLoggingUpdateParams) MarshalJSON() (data []byte, err error) {
  return apijson.MarshalRoot(r.LoggingSetting)
}

type GatewayLoggingUpdateResponseEnvelope struct {
Errors []shared.ResponseInfo `json:"errors,required"`
Messages []shared.ResponseInfo `json:"messages,required"`
// Whether the API call was successful
Success GatewayLoggingUpdateResponseEnvelopeSuccess `json:"success,required"`
Result LoggingSetting `json:"result"`
JSON gatewayLoggingUpdateResponseEnvelopeJSON `json:"-"`
}

// gatewayLoggingUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayLoggingUpdateResponseEnvelope]
type gatewayLoggingUpdateResponseEnvelopeJSON struct {
Errors apijson.Field
Messages apijson.Field
Success apijson.Field
Result apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLoggingUpdateResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type GatewayLoggingUpdateResponseEnvelopeSuccess bool

const (
  GatewayLoggingUpdateResponseEnvelopeSuccessTrue GatewayLoggingUpdateResponseEnvelopeSuccess = true
)

func (r GatewayLoggingUpdateResponseEnvelopeSuccess) IsKnown() (bool) {
  switch r {
  case GatewayLoggingUpdateResponseEnvelopeSuccessTrue:
      return true
  }
  return false
}

type GatewayLoggingGetParams struct {
AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayLoggingGetResponseEnvelope struct {
Errors []shared.ResponseInfo `json:"errors,required"`
Messages []shared.ResponseInfo `json:"messages,required"`
// Whether the API call was successful
Success GatewayLoggingGetResponseEnvelopeSuccess `json:"success,required"`
Result LoggingSetting `json:"result"`
JSON gatewayLoggingGetResponseEnvelopeJSON `json:"-"`
}

// gatewayLoggingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayLoggingGetResponseEnvelope]
type gatewayLoggingGetResponseEnvelopeJSON struct {
Errors apijson.Field
Messages apijson.Field
Success apijson.Field
Result apijson.Field
raw string
ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
  return apijson.UnmarshalRoot(data, r)
}

func (r gatewayLoggingGetResponseEnvelopeJSON) RawJSON() (string) {
  return r.raw
}

// Whether the API call was successful
type GatewayLoggingGetResponseEnvelopeSuccess bool

const (
  GatewayLoggingGetResponseEnvelopeSuccessTrue GatewayLoggingGetResponseEnvelopeSuccess = true
)

func (r GatewayLoggingGetResponseEnvelopeSuccess) IsKnown() (bool) {
  switch r {
  case GatewayLoggingGetResponseEnvelopeSuccessTrue:
      return true
  }
  return false
}

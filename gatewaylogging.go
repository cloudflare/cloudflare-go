// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// GatewayLoggingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayLoggingService] method
// instead.
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
func (r *GatewayLoggingService) Update(ctx context.Context, accountID interface{}, body GatewayLoggingUpdateParams, opts ...option.RequestOption) (res *GatewayLoggingUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLoggingUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/logging", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the current logging settings for Zero Trust account.
func (r *GatewayLoggingService) Get(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *GatewayLoggingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLoggingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/logging", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayLoggingUpdateResponse struct {
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii bool `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType GatewayLoggingUpdateResponseSettingsByRuleType `json:"settings_by_rule_type"`
	JSON               gatewayLoggingUpdateResponseJSON               `json:"-"`
}

// gatewayLoggingUpdateResponseJSON contains the JSON metadata for the struct
// [GatewayLoggingUpdateResponse]
type gatewayLoggingUpdateResponseJSON struct {
	RedactPii          apijson.Field
	SettingsByRuleType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *GatewayLoggingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Logging settings by rule type.
type GatewayLoggingUpdateResponseSettingsByRuleType struct {
	// Logging settings for DNS firewall.
	DNS interface{} `json:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	HTTP interface{} `json:"http"`
	// Logging settings for Network firewall.
	L4   interface{}                                        `json:"l4"`
	JSON gatewayLoggingUpdateResponseSettingsByRuleTypeJSON `json:"-"`
}

// gatewayLoggingUpdateResponseSettingsByRuleTypeJSON contains the JSON metadata
// for the struct [GatewayLoggingUpdateResponseSettingsByRuleType]
type gatewayLoggingUpdateResponseSettingsByRuleTypeJSON struct {
	DNS         apijson.Field
	HTTP        apijson.Field
	L4          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingUpdateResponseSettingsByRuleType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLoggingGetResponse struct {
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii bool `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType GatewayLoggingGetResponseSettingsByRuleType `json:"settings_by_rule_type"`
	JSON               gatewayLoggingGetResponseJSON               `json:"-"`
}

// gatewayLoggingGetResponseJSON contains the JSON metadata for the struct
// [GatewayLoggingGetResponse]
type gatewayLoggingGetResponseJSON struct {
	RedactPii          apijson.Field
	SettingsByRuleType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *GatewayLoggingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Logging settings by rule type.
type GatewayLoggingGetResponseSettingsByRuleType struct {
	// Logging settings for DNS firewall.
	DNS interface{} `json:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	HTTP interface{} `json:"http"`
	// Logging settings for Network firewall.
	L4   interface{}                                     `json:"l4"`
	JSON gatewayLoggingGetResponseSettingsByRuleTypeJSON `json:"-"`
}

// gatewayLoggingGetResponseSettingsByRuleTypeJSON contains the JSON metadata for
// the struct [GatewayLoggingGetResponseSettingsByRuleType]
type gatewayLoggingGetResponseSettingsByRuleTypeJSON struct {
	DNS         apijson.Field
	HTTP        apijson.Field
	L4          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingGetResponseSettingsByRuleType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLoggingUpdateParams struct {
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii param.Field[bool] `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType param.Field[GatewayLoggingUpdateParamsSettingsByRuleType] `json:"settings_by_rule_type"`
}

func (r GatewayLoggingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logging settings by rule type.
type GatewayLoggingUpdateParamsSettingsByRuleType struct {
	// Logging settings for DNS firewall.
	DNS param.Field[interface{}] `json:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	HTTP param.Field[interface{}] `json:"http"`
	// Logging settings for Network firewall.
	L4 param.Field[interface{}] `json:"l4"`
}

func (r GatewayLoggingUpdateParamsSettingsByRuleType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLoggingUpdateResponseEnvelope struct {
	Errors   []GatewayLoggingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayLoggingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayLoggingUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayLoggingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayLoggingUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayLoggingUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayLoggingUpdateResponseEnvelope]
type gatewayLoggingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLoggingUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    gatewayLoggingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayLoggingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [GatewayLoggingUpdateResponseEnvelopeErrors]
type gatewayLoggingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLoggingUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    gatewayLoggingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayLoggingUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [GatewayLoggingUpdateResponseEnvelopeMessages]
type gatewayLoggingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayLoggingUpdateResponseEnvelopeSuccess bool

const (
	GatewayLoggingUpdateResponseEnvelopeSuccessTrue GatewayLoggingUpdateResponseEnvelopeSuccess = true
)

type GatewayLoggingGetResponseEnvelope struct {
	Errors   []GatewayLoggingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayLoggingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayLoggingGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayLoggingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayLoggingGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayLoggingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayLoggingGetResponseEnvelope]
type gatewayLoggingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLoggingGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    gatewayLoggingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayLoggingGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [GatewayLoggingGetResponseEnvelopeErrors]
type gatewayLoggingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLoggingGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    gatewayLoggingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayLoggingGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [GatewayLoggingGetResponseEnvelopeMessages]
type gatewayLoggingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayLoggingGetResponseEnvelopeSuccess bool

const (
	GatewayLoggingGetResponseEnvelopeSuccessTrue GatewayLoggingGetResponseEnvelopeSuccess = true
)

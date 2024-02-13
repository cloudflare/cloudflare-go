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

// Fetches the current logging settings for Zero Trust account.
func (r *GatewayLoggingService) ZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccount(ctx context.Context, accountID interface{}, opts ...option.RequestOption) (res *GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/logging", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates logging settings for the current Zero Trust account.
func (r *GatewayLoggingService) ZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccount(ctx context.Context, accountID interface{}, body GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountParams, opts ...option.RequestOption) (res *GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/logging", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponse struct {
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii bool `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseSettingsByRuleType `json:"settings_by_rule_type"`
	JSON               gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseJSON               `json:"-"`
}

// gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseJSON
// contains the JSON metadata for the struct
// [GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponse]
type gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseJSON struct {
	RedactPii          apijson.Field
	SettingsByRuleType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Logging settings by rule type.
type GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseSettingsByRuleType struct {
	// Logging settings for DNS firewall.
	DNS interface{} `json:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	HTTP interface{} `json:"http"`
	// Logging settings for Network firewall.
	L4   interface{}                                                                                           `json:"l4"`
	JSON gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseSettingsByRuleTypeJSON `json:"-"`
}

// gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseSettingsByRuleTypeJSON
// contains the JSON metadata for the struct
// [GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseSettingsByRuleType]
type gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseSettingsByRuleTypeJSON struct {
	DNS         apijson.Field
	HTTP        apijson.Field
	L4          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseSettingsByRuleType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponse struct {
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii bool `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseSettingsByRuleType `json:"settings_by_rule_type"`
	JSON               gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseJSON               `json:"-"`
}

// gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseJSON
// contains the JSON metadata for the struct
// [GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponse]
type gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseJSON struct {
	RedactPii          apijson.Field
	SettingsByRuleType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Logging settings by rule type.
type GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseSettingsByRuleType struct {
	// Logging settings for DNS firewall.
	DNS interface{} `json:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	HTTP interface{} `json:"http"`
	// Logging settings for Network firewall.
	L4   interface{}                                                                                              `json:"l4"`
	JSON gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseSettingsByRuleTypeJSON `json:"-"`
}

// gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseSettingsByRuleTypeJSON
// contains the JSON metadata for the struct
// [GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseSettingsByRuleType]
type gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseSettingsByRuleTypeJSON struct {
	DNS         apijson.Field
	HTTP        apijson.Field
	L4          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseSettingsByRuleType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelope struct {
	Errors   []GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeJSON    `json:"-"`
}

// gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelope]
type gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeErrors struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeErrors]
type gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeMessages struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeMessages]
type gatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeSuccess bool

const (
	GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeSuccessTrue GatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseEnvelopeSuccess = true
)

type GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountParams struct {
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii param.Field[bool] `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType param.Field[GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountParamsSettingsByRuleType] `json:"settings_by_rule_type"`
}

func (r GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logging settings by rule type.
type GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountParamsSettingsByRuleType struct {
	// Logging settings for DNS firewall.
	DNS param.Field[interface{}] `json:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	HTTP param.Field[interface{}] `json:"http"`
	// Logging settings for Network firewall.
	L4 param.Field[interface{}] `json:"l4"`
}

func (r GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountParamsSettingsByRuleType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelope struct {
	Errors   []GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeErrors   `json:"errors,required"`
	Messages []GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeMessages `json:"messages,required"`
	Result   GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeJSON    `json:"-"`
}

// gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelope]
type gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeErrors struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeErrorsJSON `json:"-"`
}

// gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeErrors]
type gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeMessages struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeMessagesJSON `json:"-"`
}

// gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeMessages]
type gatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeSuccess bool

const (
	GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeSuccessTrue GatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseEnvelopeSuccess = true
)

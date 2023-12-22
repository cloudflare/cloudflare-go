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

// AccountGatewayLoggingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountGatewayLoggingService]
// method instead.
type AccountGatewayLoggingService struct {
	Options []option.RequestOption
}

// NewAccountGatewayLoggingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountGatewayLoggingService(opts ...option.RequestOption) (r *AccountGatewayLoggingService) {
	r = &AccountGatewayLoggingService{}
	r.Options = opts
	return
}

// Describes the current logging settings for Zero Trust account.
func (r *AccountGatewayLoggingService) ZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccount(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *GatewayAccountLoggingSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/logging", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the current logging settings for the Zero Trust accounty.
func (r *AccountGatewayLoggingService) ZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccount(ctx context.Context, identifier interface{}, body AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountParams, opts ...option.RequestOption) (res *GatewayAccountLoggingSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/logging", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type GatewayAccountLoggingSettingsResponse struct {
	Errors   []GatewayAccountLoggingSettingsResponseError   `json:"errors"`
	Messages []GatewayAccountLoggingSettingsResponseMessage `json:"messages"`
	Result   GatewayAccountLoggingSettingsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success GatewayAccountLoggingSettingsResponseSuccess `json:"success"`
	JSON    gatewayAccountLoggingSettingsResponseJSON    `json:"-"`
}

// gatewayAccountLoggingSettingsResponseJSON contains the JSON metadata for the
// struct [GatewayAccountLoggingSettingsResponse]
type gatewayAccountLoggingSettingsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountLoggingSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAccountLoggingSettingsResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    gatewayAccountLoggingSettingsResponseErrorJSON `json:"-"`
}

// gatewayAccountLoggingSettingsResponseErrorJSON contains the JSON metadata for
// the struct [GatewayAccountLoggingSettingsResponseError]
type gatewayAccountLoggingSettingsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountLoggingSettingsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAccountLoggingSettingsResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    gatewayAccountLoggingSettingsResponseMessageJSON `json:"-"`
}

// gatewayAccountLoggingSettingsResponseMessageJSON contains the JSON metadata for
// the struct [GatewayAccountLoggingSettingsResponseMessage]
type gatewayAccountLoggingSettingsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountLoggingSettingsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GatewayAccountLoggingSettingsResponseResult struct {
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii bool `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType GatewayAccountLoggingSettingsResponseResultSettingsByRuleType `json:"settings_by_rule_type"`
	JSON               gatewayAccountLoggingSettingsResponseResultJSON               `json:"-"`
}

// gatewayAccountLoggingSettingsResponseResultJSON contains the JSON metadata for
// the struct [GatewayAccountLoggingSettingsResponseResult]
type gatewayAccountLoggingSettingsResponseResultJSON struct {
	RedactPii          apijson.Field
	SettingsByRuleType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *GatewayAccountLoggingSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Logging settings by rule type.
type GatewayAccountLoggingSettingsResponseResultSettingsByRuleType struct {
	// Logging settings for DNS firewall.
	DNS interface{} `json:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	HTTP interface{} `json:"http"`
	// Logging settings for Network firewall.
	L4   interface{}                                                       `json:"l4"`
	JSON gatewayAccountLoggingSettingsResponseResultSettingsByRuleTypeJSON `json:"-"`
}

// gatewayAccountLoggingSettingsResponseResultSettingsByRuleTypeJSON contains the
// JSON metadata for the struct
// [GatewayAccountLoggingSettingsResponseResultSettingsByRuleType]
type gatewayAccountLoggingSettingsResponseResultSettingsByRuleTypeJSON struct {
	DNS         apijson.Field
	HTTP        apijson.Field
	L4          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayAccountLoggingSettingsResponseResultSettingsByRuleType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GatewayAccountLoggingSettingsResponseSuccess bool

const (
	GatewayAccountLoggingSettingsResponseSuccessTrue GatewayAccountLoggingSettingsResponseSuccess = true
)

type AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountParams struct {
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii param.Field[bool] `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType param.Field[AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountParamsSettingsByRuleType] `json:"settings_by_rule_type"`
}

func (r AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logging settings by rule type.
type AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountParamsSettingsByRuleType struct {
	// Logging settings for DNS firewall.
	DNS param.Field[interface{}] `json:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	HTTP param.Field[interface{}] `json:"http"`
	// Logging settings for Network firewall.
	L4 param.Field[interface{}] `json:"l4"`
}

func (r AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountParamsSettingsByRuleType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

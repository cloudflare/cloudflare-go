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

// Fetches the current logging settings for Zero Trust account.
func (r *AccountGatewayLoggingService) ZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccount(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/logging", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates logging settings for the current Zero Trust account.
func (r *AccountGatewayLoggingService) ZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccount(ctx context.Context, identifier interface{}, body AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountParams, opts ...option.RequestOption) (res *AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/logging", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponse struct {
	Errors   []AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseError   `json:"errors"`
	Messages []AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseMessage `json:"messages"`
	Result   AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseSuccess `json:"success"`
	JSON    accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseJSON    `json:"-"`
}

// accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseJSON
// contains the JSON metadata for the struct
// [AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponse]
type accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseError struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseErrorJSON `json:"-"`
}

// accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseError]
type accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseMessage struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseMessageJSON `json:"-"`
}

// accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseMessage]
type accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseResult struct {
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii bool `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseResultSettingsByRuleType `json:"settings_by_rule_type"`
	JSON               accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseResultJSON               `json:"-"`
}

// accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseResultJSON
// contains the JSON metadata for the struct
// [AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseResult]
type accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseResultJSON struct {
	RedactPii          apijson.Field
	SettingsByRuleType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Logging settings by rule type.
type AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseResultSettingsByRuleType struct {
	// Logging settings for DNS firewall.
	DNS interface{} `json:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	HTTP interface{} `json:"http"`
	// Logging settings for Network firewall.
	L4   interface{}                                                                                                        `json:"l4"`
	JSON accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseResultSettingsByRuleTypeJSON `json:"-"`
}

// accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseResultSettingsByRuleTypeJSON
// contains the JSON metadata for the struct
// [AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseResultSettingsByRuleType]
type accountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseResultSettingsByRuleTypeJSON struct {
	DNS         apijson.Field
	HTTP        apijson.Field
	L4          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseResultSettingsByRuleType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseSuccess bool

const (
	AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseSuccessTrue AccountGatewayLoggingZeroTrustAccountsGetLoggingSettingsForTheZeroTrustAccountResponseSuccess = true
)

type AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponse struct {
	Errors   []AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseError   `json:"errors"`
	Messages []AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseMessage `json:"messages"`
	Result   AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseSuccess `json:"success"`
	JSON    accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseJSON    `json:"-"`
}

// accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseJSON
// contains the JSON metadata for the struct
// [AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponse]
type accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseError struct {
	Code    int64                                                                                              `json:"code,required"`
	Message string                                                                                             `json:"message,required"`
	JSON    accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseErrorJSON `json:"-"`
}

// accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseError]
type accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseMessage struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseMessageJSON `json:"-"`
}

// accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseMessage]
type accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseResult struct {
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii bool `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseResultSettingsByRuleType `json:"settings_by_rule_type"`
	JSON               accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseResultJSON               `json:"-"`
}

// accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseResultJSON
// contains the JSON metadata for the struct
// [AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseResult]
type accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseResultJSON struct {
	RedactPii          apijson.Field
	SettingsByRuleType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Logging settings by rule type.
type AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseResultSettingsByRuleType struct {
	// Logging settings for DNS firewall.
	DNS interface{} `json:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	HTTP interface{} `json:"http"`
	// Logging settings for Network firewall.
	L4   interface{}                                                                                                           `json:"l4"`
	JSON accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseResultSettingsByRuleTypeJSON `json:"-"`
}

// accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseResultSettingsByRuleTypeJSON
// contains the JSON metadata for the struct
// [AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseResultSettingsByRuleType]
type accountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseResultSettingsByRuleTypeJSON struct {
	DNS         apijson.Field
	HTTP        apijson.Field
	L4          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseResultSettingsByRuleType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseSuccess bool

const (
	AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseSuccessTrue AccountGatewayLoggingZeroTrustAccountsUpdateLoggingSettingsForTheZeroTrustAccountResponseSuccess = true
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

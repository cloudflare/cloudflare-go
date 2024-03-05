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

// ZeroTrustGatewayLoggingService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustGatewayLoggingService] method instead.
type ZeroTrustGatewayLoggingService struct {
	Options []option.RequestOption
}

// NewZeroTrustGatewayLoggingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustGatewayLoggingService(opts ...option.RequestOption) (r *ZeroTrustGatewayLoggingService) {
	r = &ZeroTrustGatewayLoggingService{}
	r.Options = opts
	return
}

// Updates logging settings for the current Zero Trust account.
func (r *ZeroTrustGatewayLoggingService) Update(ctx context.Context, params ZeroTrustGatewayLoggingUpdateParams, opts ...option.RequestOption) (res *ZeroTrustGatewayGatewayAccountLoggingSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayLoggingUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/logging", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the current logging settings for Zero Trust account.
func (r *ZeroTrustGatewayLoggingService) Get(ctx context.Context, query ZeroTrustGatewayLoggingGetParams, opts ...option.RequestOption) (res *ZeroTrustGatewayGatewayAccountLoggingSettings, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustGatewayLoggingGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/gateway/logging", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustGatewayGatewayAccountLoggingSettings struct {
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii bool `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType ZeroTrustGatewayGatewayAccountLoggingSettingsSettingsByRuleType `json:"settings_by_rule_type"`
	JSON               zeroTrustGatewayGatewayAccountLoggingSettingsJSON               `json:"-"`
}

// zeroTrustGatewayGatewayAccountLoggingSettingsJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayGatewayAccountLoggingSettings]
type zeroTrustGatewayGatewayAccountLoggingSettingsJSON struct {
	RedactPii          apijson.Field
	SettingsByRuleType apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZeroTrustGatewayGatewayAccountLoggingSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Logging settings by rule type.
type ZeroTrustGatewayGatewayAccountLoggingSettingsSettingsByRuleType struct {
	// Logging settings for DNS firewall.
	DNS interface{} `json:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	HTTP interface{} `json:"http"`
	// Logging settings for Network firewall.
	L4   interface{}                                                         `json:"l4"`
	JSON zeroTrustGatewayGatewayAccountLoggingSettingsSettingsByRuleTypeJSON `json:"-"`
}

// zeroTrustGatewayGatewayAccountLoggingSettingsSettingsByRuleTypeJSON contains the
// JSON metadata for the struct
// [ZeroTrustGatewayGatewayAccountLoggingSettingsSettingsByRuleType]
type zeroTrustGatewayGatewayAccountLoggingSettingsSettingsByRuleTypeJSON struct {
	DNS         apijson.Field
	HTTP        apijson.Field
	L4          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayGatewayAccountLoggingSettingsSettingsByRuleType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLoggingUpdateParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
	// Redact personally identifiable information from activity logging (PII fields
	// are: source IP, user email, user ID, device ID, URL, referrer, user agent).
	RedactPii param.Field[bool] `json:"redact_pii"`
	// Logging settings by rule type.
	SettingsByRuleType param.Field[ZeroTrustGatewayLoggingUpdateParamsSettingsByRuleType] `json:"settings_by_rule_type"`
}

func (r ZeroTrustGatewayLoggingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Logging settings by rule type.
type ZeroTrustGatewayLoggingUpdateParamsSettingsByRuleType struct {
	// Logging settings for DNS firewall.
	DNS param.Field[interface{}] `json:"dns"`
	// Logging settings for HTTP/HTTPS firewall.
	HTTP param.Field[interface{}] `json:"http"`
	// Logging settings for Network firewall.
	L4 param.Field[interface{}] `json:"l4"`
}

func (r ZeroTrustGatewayLoggingUpdateParamsSettingsByRuleType) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustGatewayLoggingUpdateResponseEnvelope struct {
	Errors   []ZeroTrustGatewayLoggingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayLoggingUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayGatewayAccountLoggingSettings           `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayLoggingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayLoggingUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayLoggingUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayLoggingUpdateResponseEnvelope]
type zeroTrustGatewayLoggingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLoggingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLoggingUpdateResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustGatewayLoggingUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayLoggingUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayLoggingUpdateResponseEnvelopeErrors]
type zeroTrustGatewayLoggingUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLoggingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLoggingUpdateResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustGatewayLoggingUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayLoggingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayLoggingUpdateResponseEnvelopeMessages]
type zeroTrustGatewayLoggingUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLoggingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayLoggingUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayLoggingUpdateResponseEnvelopeSuccessTrue ZeroTrustGatewayLoggingUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustGatewayLoggingGetParams struct {
	AccountID param.Field[interface{}] `path:"account_id,required"`
}

type ZeroTrustGatewayLoggingGetResponseEnvelope struct {
	Errors   []ZeroTrustGatewayLoggingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustGatewayLoggingGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustGatewayGatewayAccountLoggingSettings        `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustGatewayLoggingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustGatewayLoggingGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustGatewayLoggingGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayLoggingGetResponseEnvelope]
type zeroTrustGatewayLoggingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLoggingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLoggingGetResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    zeroTrustGatewayLoggingGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustGatewayLoggingGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayLoggingGetResponseEnvelopeErrors]
type zeroTrustGatewayLoggingGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLoggingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustGatewayLoggingGetResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    zeroTrustGatewayLoggingGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustGatewayLoggingGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayLoggingGetResponseEnvelopeMessages]
type zeroTrustGatewayLoggingGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayLoggingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustGatewayLoggingGetResponseEnvelopeSuccess bool

const (
	ZeroTrustGatewayLoggingGetResponseEnvelopeSuccessTrue ZeroTrustGatewayLoggingGetResponseEnvelopeSuccess = true
)

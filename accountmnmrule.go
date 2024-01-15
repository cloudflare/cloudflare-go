// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountMnmRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountMnmRuleService] method
// instead.
type AccountMnmRuleService struct {
	Options        []option.RequestOption
	Advertisements *AccountMnmRuleAdvertisementService
}

// NewAccountMnmRuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountMnmRuleService(opts ...option.RequestOption) (r *AccountMnmRuleService) {
	r = &AccountMnmRuleService{}
	r.Options = opts
	r.Advertisements = NewAccountMnmRuleAdvertisementService(opts...)
	return
}

// List a single network monitoring rule for account.
func (r *AccountMnmRuleService) Get(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *AccountMnmRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a network monitoring rule for account.
func (r *AccountMnmRuleService) Update(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *AccountMnmRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Delete a network monitoring rule for account.
func (r *AccountMnmRuleService) Delete(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *AccountMnmRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create network monitoring rules for account. Currently only supports creating a
// single rule per API request.
func (r *AccountMnmRuleService) MagicNetworkMonitoringRulesNewRules(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Lists network monitoring rules for account.
func (r *AccountMnmRuleService) MagicNetworkMonitoringRulesListRules(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update network monitoring rules for account.
func (r *AccountMnmRuleService) MagicNetworkMonitoringRulesUpdateRules(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type AccountMnmRuleGetResponse struct {
	Errors   []AccountMnmRuleGetResponseError   `json:"errors"`
	Messages []AccountMnmRuleGetResponseMessage `json:"messages"`
	Result   AccountMnmRuleGetResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountMnmRuleGetResponseSuccess `json:"success"`
	JSON    accountMnmRuleGetResponseJSON    `json:"-"`
}

// accountMnmRuleGetResponseJSON contains the JSON metadata for the struct
// [AccountMnmRuleGetResponse]
type accountMnmRuleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleGetResponseError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    accountMnmRuleGetResponseErrorJSON `json:"-"`
}

// accountMnmRuleGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountMnmRuleGetResponseError]
type accountMnmRuleGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleGetResponseMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    accountMnmRuleGetResponseMessageJSON `json:"-"`
}

// accountMnmRuleGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountMnmRuleGetResponseMessage]
type accountMnmRuleGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleGetResponseResult struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool `json:"automatic_advertisement,required,nullable"`
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The final value must be equivalent to one of the following 8
	// values ["1m","5m","10m","15m","20m","30m","45m","60m"]. The format is
	// AhBmCsDmsEusFns where A, B, C, D, E and F durations are optional; however at
	// least one unit must be provided.
	Duration string `json:"duration,required"`
	// The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9,
	// underscore (\_), dash (-), period (.), and tilde (~). You can’t have a space in
	// the rule name. Max 256 characters.
	Name     string      `json:"name,required"`
	Prefixes []string    `json:"prefixes,required"`
	ID       interface{} `json:"id"`
	// The number of bits per second for the rule. When this value is exceeded for the
	// set duration, an alert notification is sent. Minimum of 1 and no maximum.
	BandwidthThreshold float64 `json:"bandwidth_threshold"`
	// The number of packets per second for the rule. When this value is exceeded for
	// the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	PacketThreshold float64                             `json:"packet_threshold"`
	JSON            accountMnmRuleGetResponseResultJSON `json:"-"`
}

// accountMnmRuleGetResponseResultJSON contains the JSON metadata for the struct
// [AccountMnmRuleGetResponseResult]
type accountMnmRuleGetResponseResultJSON struct {
	AutomaticAdvertisement apijson.Field
	Duration               apijson.Field
	Name                   apijson.Field
	Prefixes               apijson.Field
	ID                     apijson.Field
	BandwidthThreshold     apijson.Field
	PacketThreshold        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountMnmRuleGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMnmRuleGetResponseSuccess bool

const (
	AccountMnmRuleGetResponseSuccessTrue AccountMnmRuleGetResponseSuccess = true
)

type AccountMnmRuleUpdateResponse struct {
	Errors   []AccountMnmRuleUpdateResponseError   `json:"errors"`
	Messages []AccountMnmRuleUpdateResponseMessage `json:"messages"`
	Result   AccountMnmRuleUpdateResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountMnmRuleUpdateResponseSuccess `json:"success"`
	JSON    accountMnmRuleUpdateResponseJSON    `json:"-"`
}

// accountMnmRuleUpdateResponseJSON contains the JSON metadata for the struct
// [AccountMnmRuleUpdateResponse]
type accountMnmRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleUpdateResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountMnmRuleUpdateResponseErrorJSON `json:"-"`
}

// accountMnmRuleUpdateResponseErrorJSON contains the JSON metadata for the struct
// [AccountMnmRuleUpdateResponseError]
type accountMnmRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleUpdateResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountMnmRuleUpdateResponseMessageJSON `json:"-"`
}

// accountMnmRuleUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountMnmRuleUpdateResponseMessage]
type accountMnmRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleUpdateResponseResult struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool `json:"automatic_advertisement,required,nullable"`
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The final value must be equivalent to one of the following 8
	// values ["1m","5m","10m","15m","20m","30m","45m","60m"]. The format is
	// AhBmCsDmsEusFns where A, B, C, D, E and F durations are optional; however at
	// least one unit must be provided.
	Duration string `json:"duration,required"`
	// The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9,
	// underscore (\_), dash (-), period (.), and tilde (~). You can’t have a space in
	// the rule name. Max 256 characters.
	Name     string      `json:"name,required"`
	Prefixes []string    `json:"prefixes,required"`
	ID       interface{} `json:"id"`
	// The number of bits per second for the rule. When this value is exceeded for the
	// set duration, an alert notification is sent. Minimum of 1 and no maximum.
	BandwidthThreshold float64 `json:"bandwidth_threshold"`
	// The number of packets per second for the rule. When this value is exceeded for
	// the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	PacketThreshold float64                                `json:"packet_threshold"`
	JSON            accountMnmRuleUpdateResponseResultJSON `json:"-"`
}

// accountMnmRuleUpdateResponseResultJSON contains the JSON metadata for the struct
// [AccountMnmRuleUpdateResponseResult]
type accountMnmRuleUpdateResponseResultJSON struct {
	AutomaticAdvertisement apijson.Field
	Duration               apijson.Field
	Name                   apijson.Field
	Prefixes               apijson.Field
	ID                     apijson.Field
	BandwidthThreshold     apijson.Field
	PacketThreshold        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountMnmRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMnmRuleUpdateResponseSuccess bool

const (
	AccountMnmRuleUpdateResponseSuccessTrue AccountMnmRuleUpdateResponseSuccess = true
)

type AccountMnmRuleDeleteResponse struct {
	Errors   []AccountMnmRuleDeleteResponseError   `json:"errors"`
	Messages []AccountMnmRuleDeleteResponseMessage `json:"messages"`
	Result   AccountMnmRuleDeleteResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountMnmRuleDeleteResponseSuccess `json:"success"`
	JSON    accountMnmRuleDeleteResponseJSON    `json:"-"`
}

// accountMnmRuleDeleteResponseJSON contains the JSON metadata for the struct
// [AccountMnmRuleDeleteResponse]
type accountMnmRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleDeleteResponseError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountMnmRuleDeleteResponseErrorJSON `json:"-"`
}

// accountMnmRuleDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AccountMnmRuleDeleteResponseError]
type accountMnmRuleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleDeleteResponseMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountMnmRuleDeleteResponseMessageJSON `json:"-"`
}

// accountMnmRuleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountMnmRuleDeleteResponseMessage]
type accountMnmRuleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleDeleteResponseResult struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool `json:"automatic_advertisement,required,nullable"`
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The final value must be equivalent to one of the following 8
	// values ["1m","5m","10m","15m","20m","30m","45m","60m"]. The format is
	// AhBmCsDmsEusFns where A, B, C, D, E and F durations are optional; however at
	// least one unit must be provided.
	Duration string `json:"duration,required"`
	// The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9,
	// underscore (\_), dash (-), period (.), and tilde (~). You can’t have a space in
	// the rule name. Max 256 characters.
	Name     string      `json:"name,required"`
	Prefixes []string    `json:"prefixes,required"`
	ID       interface{} `json:"id"`
	// The number of bits per second for the rule. When this value is exceeded for the
	// set duration, an alert notification is sent. Minimum of 1 and no maximum.
	BandwidthThreshold float64 `json:"bandwidth_threshold"`
	// The number of packets per second for the rule. When this value is exceeded for
	// the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	PacketThreshold float64                                `json:"packet_threshold"`
	JSON            accountMnmRuleDeleteResponseResultJSON `json:"-"`
}

// accountMnmRuleDeleteResponseResultJSON contains the JSON metadata for the struct
// [AccountMnmRuleDeleteResponseResult]
type accountMnmRuleDeleteResponseResultJSON struct {
	AutomaticAdvertisement apijson.Field
	Duration               apijson.Field
	Name                   apijson.Field
	Prefixes               apijson.Field
	ID                     apijson.Field
	BandwidthThreshold     apijson.Field
	PacketThreshold        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountMnmRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMnmRuleDeleteResponseSuccess bool

const (
	AccountMnmRuleDeleteResponseSuccessTrue AccountMnmRuleDeleteResponseSuccess = true
)

type AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponse struct {
	Errors   []AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseError   `json:"errors"`
	Messages []AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseMessage `json:"messages"`
	Result   AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseSuccess `json:"success"`
	JSON    accountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseJSON    `json:"-"`
}

// accountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseJSON contains the JSON
// metadata for the struct
// [AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponse]
type accountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseError struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    accountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseErrorJSON `json:"-"`
}

// accountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseError]
type accountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseMessage struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    accountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseMessageJSON `json:"-"`
}

// accountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseMessage]
type accountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseResult struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool `json:"automatic_advertisement,required,nullable"`
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The final value must be equivalent to one of the following 8
	// values ["1m","5m","10m","15m","20m","30m","45m","60m"]. The format is
	// AhBmCsDmsEusFns where A, B, C, D, E and F durations are optional; however at
	// least one unit must be provided.
	Duration string `json:"duration,required"`
	// The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9,
	// underscore (\_), dash (-), period (.), and tilde (~). You can’t have a space in
	// the rule name. Max 256 characters.
	Name     string      `json:"name,required"`
	Prefixes []string    `json:"prefixes,required"`
	ID       interface{} `json:"id"`
	// The number of bits per second for the rule. When this value is exceeded for the
	// set duration, an alert notification is sent. Minimum of 1 and no maximum.
	BandwidthThreshold float64 `json:"bandwidth_threshold"`
	// The number of packets per second for the rule. When this value is exceeded for
	// the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	PacketThreshold float64                                                             `json:"packet_threshold"`
	JSON            accountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseResultJSON `json:"-"`
}

// accountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseResultJSON contains the
// JSON metadata for the struct
// [AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseResult]
type accountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseResultJSON struct {
	AutomaticAdvertisement apijson.Field
	Duration               apijson.Field
	Name                   apijson.Field
	Prefixes               apijson.Field
	ID                     apijson.Field
	BandwidthThreshold     apijson.Field
	PacketThreshold        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseSuccess bool

const (
	AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseSuccessTrue AccountMnmRuleMagicNetworkMonitoringRulesNewRulesResponseSuccess = true
)

type AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponse struct {
	Errors     []AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseError    `json:"errors"`
	Messages   []AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseMessage  `json:"messages"`
	Result     []AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseResult   `json:"result,nullable"`
	ResultInfo AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseSuccess `json:"success"`
	JSON    accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseJSON    `json:"-"`
}

// accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseJSON contains the JSON
// metadata for the struct
// [AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponse]
type accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseError struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseErrorJSON `json:"-"`
}

// accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseError]
type accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseMessage struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseMessageJSON `json:"-"`
}

// accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseMessage]
type accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseResult struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool `json:"automatic_advertisement,required,nullable"`
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The final value must be equivalent to one of the following 8
	// values ["1m","5m","10m","15m","20m","30m","45m","60m"]. The format is
	// AhBmCsDmsEusFns where A, B, C, D, E and F durations are optional; however at
	// least one unit must be provided.
	Duration string `json:"duration,required"`
	// The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9,
	// underscore (\_), dash (-), period (.), and tilde (~). You can’t have a space in
	// the rule name. Max 256 characters.
	Name     string      `json:"name,required"`
	Prefixes []string    `json:"prefixes,required"`
	ID       interface{} `json:"id"`
	// The number of bits per second for the rule. When this value is exceeded for the
	// set duration, an alert notification is sent. Minimum of 1 and no maximum.
	BandwidthThreshold float64 `json:"bandwidth_threshold"`
	// The number of packets per second for the rule. When this value is exceeded for
	// the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	PacketThreshold float64                                                              `json:"packet_threshold"`
	JSON            accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseResultJSON `json:"-"`
}

// accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseResultJSON contains
// the JSON metadata for the struct
// [AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseResult]
type accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseResultJSON struct {
	AutomaticAdvertisement apijson.Field
	Duration               apijson.Field
	Name                   apijson.Field
	Prefixes               apijson.Field
	ID                     apijson.Field
	BandwidthThreshold     apijson.Field
	PacketThreshold        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                  `json:"total_count"`
	JSON       accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseResultInfoJSON `json:"-"`
}

// accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseResultInfo]
type accountMnmRuleMagicNetworkMonitoringRulesListRulesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseSuccess bool

const (
	AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseSuccessTrue AccountMnmRuleMagicNetworkMonitoringRulesListRulesResponseSuccess = true
)

type AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponse struct {
	Errors   []AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseError   `json:"errors"`
	Messages []AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseMessage `json:"messages"`
	Result   AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseSuccess `json:"success"`
	JSON    accountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseJSON    `json:"-"`
}

// accountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseJSON contains the
// JSON metadata for the struct
// [AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponse]
type accountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseError struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseErrorJSON `json:"-"`
}

// accountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseErrorJSON contains
// the JSON metadata for the struct
// [AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseError]
type accountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseMessage struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    accountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseMessageJSON `json:"-"`
}

// accountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseMessage]
type accountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseResult struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool `json:"automatic_advertisement,required,nullable"`
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The final value must be equivalent to one of the following 8
	// values ["1m","5m","10m","15m","20m","30m","45m","60m"]. The format is
	// AhBmCsDmsEusFns where A, B, C, D, E and F durations are optional; however at
	// least one unit must be provided.
	Duration string `json:"duration,required"`
	// The name of the rule. Must be unique. Supports characters A-Z, a-z, 0-9,
	// underscore (\_), dash (-), period (.), and tilde (~). You can’t have a space in
	// the rule name. Max 256 characters.
	Name     string      `json:"name,required"`
	Prefixes []string    `json:"prefixes,required"`
	ID       interface{} `json:"id"`
	// The number of bits per second for the rule. When this value is exceeded for the
	// set duration, an alert notification is sent. Minimum of 1 and no maximum.
	BandwidthThreshold float64 `json:"bandwidth_threshold"`
	// The number of packets per second for the rule. When this value is exceeded for
	// the set duration, an alert notification is sent. Minimum of 1 and no maximum.
	PacketThreshold float64                                                                `json:"packet_threshold"`
	JSON            accountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseResultJSON `json:"-"`
}

// accountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseResultJSON contains
// the JSON metadata for the struct
// [AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseResult]
type accountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseResultJSON struct {
	AutomaticAdvertisement apijson.Field
	Duration               apijson.Field
	Name                   apijson.Field
	Prefixes               apijson.Field
	ID                     apijson.Field
	BandwidthThreshold     apijson.Field
	PacketThreshold        apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseSuccess bool

const (
	AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseSuccessTrue AccountMnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseSuccess = true
)

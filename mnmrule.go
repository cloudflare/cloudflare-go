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

// MnmRuleService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewMnmRuleService] method instead.
type MnmRuleService struct {
	Options        []option.RequestOption
	Advertisements *MnmRuleAdvertisementService
}

// NewMnmRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMnmRuleService(opts ...option.RequestOption) (r *MnmRuleService) {
	r = &MnmRuleService{}
	r.Options = opts
	r.Advertisements = NewMnmRuleAdvertisementService(opts...)
	return
}

// Update a network monitoring rule for account.
func (r *MnmRuleService) Update(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MnmRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmRuleUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a network monitoring rule for account.
func (r *MnmRuleService) Delete(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MnmRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmRuleDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List a single network monitoring rule for account.
func (r *MnmRuleService) Get(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MnmRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmRuleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create network monitoring rules for account. Currently only supports creating a
// single rule per API request.
func (r *MnmRuleService) MagicNetworkMonitoringRulesNewRules(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmRuleMagicNetworkMonitoringRulesNewRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists network monitoring rules for account.
func (r *MnmRuleService) MagicNetworkMonitoringRulesListRules(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *[]MnmRuleMagicNetworkMonitoringRulesListRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update network monitoring rules for account.
func (r *MnmRuleService) MagicNetworkMonitoringRulesUpdateRules(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MnmRuleUpdateResponse struct {
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
	PacketThreshold float64                   `json:"packet_threshold"`
	JSON            mnmRuleUpdateResponseJSON `json:"-"`
}

// mnmRuleUpdateResponseJSON contains the JSON metadata for the struct
// [MnmRuleUpdateResponse]
type mnmRuleUpdateResponseJSON struct {
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

func (r *MnmRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleDeleteResponse struct {
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
	PacketThreshold float64                   `json:"packet_threshold"`
	JSON            mnmRuleDeleteResponseJSON `json:"-"`
}

// mnmRuleDeleteResponseJSON contains the JSON metadata for the struct
// [MnmRuleDeleteResponse]
type mnmRuleDeleteResponseJSON struct {
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

func (r *MnmRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleGetResponse struct {
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
	PacketThreshold float64                `json:"packet_threshold"`
	JSON            mnmRuleGetResponseJSON `json:"-"`
}

// mnmRuleGetResponseJSON contains the JSON metadata for the struct
// [MnmRuleGetResponse]
type mnmRuleGetResponseJSON struct {
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

func (r *MnmRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleMagicNetworkMonitoringRulesNewRulesResponse struct {
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
	PacketThreshold float64                                                `json:"packet_threshold"`
	JSON            mnmRuleMagicNetworkMonitoringRulesNewRulesResponseJSON `json:"-"`
}

// mnmRuleMagicNetworkMonitoringRulesNewRulesResponseJSON contains the JSON
// metadata for the struct [MnmRuleMagicNetworkMonitoringRulesNewRulesResponse]
type mnmRuleMagicNetworkMonitoringRulesNewRulesResponseJSON struct {
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

func (r *MnmRuleMagicNetworkMonitoringRulesNewRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleMagicNetworkMonitoringRulesListRulesResponse struct {
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
	PacketThreshold float64                                                 `json:"packet_threshold"`
	JSON            mnmRuleMagicNetworkMonitoringRulesListRulesResponseJSON `json:"-"`
}

// mnmRuleMagicNetworkMonitoringRulesListRulesResponseJSON contains the JSON
// metadata for the struct [MnmRuleMagicNetworkMonitoringRulesListRulesResponse]
type mnmRuleMagicNetworkMonitoringRulesListRulesResponseJSON struct {
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

func (r *MnmRuleMagicNetworkMonitoringRulesListRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponse struct {
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
	PacketThreshold float64                                                   `json:"packet_threshold"`
	JSON            mnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseJSON `json:"-"`
}

// mnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseJSON contains the JSON
// metadata for the struct [MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponse]
type mnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseJSON struct {
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

func (r *MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleUpdateResponseEnvelope struct {
	Errors   []MnmRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmRuleUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MnmRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// mnmRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [MnmRuleUpdateResponseEnvelope]
type mnmRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    mnmRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MnmRuleUpdateResponseEnvelopeErrors]
type mnmRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    mnmRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MnmRuleUpdateResponseEnvelopeMessages]
type mnmRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmRuleUpdateResponseEnvelopeSuccess bool

const (
	MnmRuleUpdateResponseEnvelopeSuccessTrue MnmRuleUpdateResponseEnvelopeSuccess = true
)

type MnmRuleDeleteResponseEnvelope struct {
	Errors   []MnmRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmRuleDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MnmRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// mnmRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MnmRuleDeleteResponseEnvelope]
type mnmRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    mnmRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MnmRuleDeleteResponseEnvelopeErrors]
type mnmRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    mnmRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MnmRuleDeleteResponseEnvelopeMessages]
type mnmRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmRuleDeleteResponseEnvelopeSuccess bool

const (
	MnmRuleDeleteResponseEnvelopeSuccessTrue MnmRuleDeleteResponseEnvelopeSuccess = true
)

type MnmRuleGetResponseEnvelope struct {
	Errors   []MnmRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmRuleGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MnmRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmRuleGetResponseEnvelopeJSON    `json:"-"`
}

// mnmRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MnmRuleGetResponseEnvelope]
type mnmRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    mnmRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MnmRuleGetResponseEnvelopeErrors]
type mnmRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    mnmRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [MnmRuleGetResponseEnvelopeMessages]
type mnmRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmRuleGetResponseEnvelopeSuccess bool

const (
	MnmRuleGetResponseEnvelopeSuccessTrue MnmRuleGetResponseEnvelopeSuccess = true
)

type MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelope struct {
	Errors   []MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmRuleMagicNetworkMonitoringRulesNewRulesResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeJSON    `json:"-"`
}

// mnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelope]
type mnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeErrors struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    mnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeErrors]
type mnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeMessages struct {
	Code    int64                                                                  `json:"code,required"`
	Message string                                                                 `json:"message,required"`
	JSON    mnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeMessages]
type mnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeSuccess bool

const (
	MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeSuccessTrue MnmRuleMagicNetworkMonitoringRulesNewRulesResponseEnvelopeSuccess = true
)

type MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelope struct {
	Errors   []MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   []MnmRuleMagicNetworkMonitoringRulesListRulesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       mnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeJSON       `json:"-"`
}

// mnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelope]
type mnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeErrors struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    mnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeErrors]
type mnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeMessages struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    mnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeMessagesJSON contains
// the JSON metadata for the struct
// [MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeMessages]
type mnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeSuccess bool

const (
	MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeSuccessTrue MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeSuccess = true
)

type MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                   `json:"total_count"`
	JSON       mnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeResultInfoJSON `json:"-"`
}

// mnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeResultInfo]
type mnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleMagicNetworkMonitoringRulesListRulesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelope struct {
	Errors   []MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeJSON    `json:"-"`
}

// mnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelope]
type mnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeErrors struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    mnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeErrors]
type mnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeMessages struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    mnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeMessages]
type mnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeSuccess bool

const (
	MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeSuccessTrue MnmRuleMagicNetworkMonitoringRulesUpdateRulesResponseEnvelopeSuccess = true
)

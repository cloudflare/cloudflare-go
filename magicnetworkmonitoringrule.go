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

// MagicNetworkMonitoringRuleService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewMagicNetworkMonitoringRuleService] method instead.
type MagicNetworkMonitoringRuleService struct {
	Options        []option.RequestOption
	Advertisements *MagicNetworkMonitoringRuleAdvertisementService
}

// NewMagicNetworkMonitoringRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMagicNetworkMonitoringRuleService(opts ...option.RequestOption) (r *MagicNetworkMonitoringRuleService) {
	r = &MagicNetworkMonitoringRuleService{}
	r.Options = opts
	r.Advertisements = NewMagicNetworkMonitoringRuleAdvertisementService(opts...)
	return
}

// Create network monitoring rules for account. Currently only supports creating a
// single rule per API request.
func (r *MagicNetworkMonitoringRuleService) New(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MagicNetworkMonitoringRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicNetworkMonitoringRuleNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update network monitoring rules for account.
func (r *MagicNetworkMonitoringRuleService) Update(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MagicNetworkMonitoringRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicNetworkMonitoringRuleUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists network monitoring rules for account.
func (r *MagicNetworkMonitoringRuleService) List(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *[]MagicNetworkMonitoringRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicNetworkMonitoringRuleListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a network monitoring rule for account.
func (r *MagicNetworkMonitoringRuleService) Delete(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MagicNetworkMonitoringRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicNetworkMonitoringRuleDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a network monitoring rule for account.
func (r *MagicNetworkMonitoringRuleService) Edit(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MagicNetworkMonitoringRuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicNetworkMonitoringRuleEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List a single network monitoring rule for account.
func (r *MagicNetworkMonitoringRuleService) Get(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MagicNetworkMonitoringRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MagicNetworkMonitoringRuleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MagicNetworkMonitoringRuleNewResponse struct {
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
	PacketThreshold float64                                   `json:"packet_threshold"`
	JSON            magicNetworkMonitoringRuleNewResponseJSON `json:"-"`
}

// magicNetworkMonitoringRuleNewResponseJSON contains the JSON metadata for the
// struct [MagicNetworkMonitoringRuleNewResponse]
type magicNetworkMonitoringRuleNewResponseJSON struct {
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

func (r *MagicNetworkMonitoringRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleUpdateResponse struct {
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
	PacketThreshold float64                                      `json:"packet_threshold"`
	JSON            magicNetworkMonitoringRuleUpdateResponseJSON `json:"-"`
}

// magicNetworkMonitoringRuleUpdateResponseJSON contains the JSON metadata for the
// struct [MagicNetworkMonitoringRuleUpdateResponse]
type magicNetworkMonitoringRuleUpdateResponseJSON struct {
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

func (r *MagicNetworkMonitoringRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleListResponse struct {
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
	PacketThreshold float64                                    `json:"packet_threshold"`
	JSON            magicNetworkMonitoringRuleListResponseJSON `json:"-"`
}

// magicNetworkMonitoringRuleListResponseJSON contains the JSON metadata for the
// struct [MagicNetworkMonitoringRuleListResponse]
type magicNetworkMonitoringRuleListResponseJSON struct {
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

func (r *MagicNetworkMonitoringRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleDeleteResponse struct {
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
	PacketThreshold float64                                      `json:"packet_threshold"`
	JSON            magicNetworkMonitoringRuleDeleteResponseJSON `json:"-"`
}

// magicNetworkMonitoringRuleDeleteResponseJSON contains the JSON metadata for the
// struct [MagicNetworkMonitoringRuleDeleteResponse]
type magicNetworkMonitoringRuleDeleteResponseJSON struct {
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

func (r *MagicNetworkMonitoringRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleEditResponse struct {
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
	PacketThreshold float64                                    `json:"packet_threshold"`
	JSON            magicNetworkMonitoringRuleEditResponseJSON `json:"-"`
}

// magicNetworkMonitoringRuleEditResponseJSON contains the JSON metadata for the
// struct [MagicNetworkMonitoringRuleEditResponse]
type magicNetworkMonitoringRuleEditResponseJSON struct {
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

func (r *MagicNetworkMonitoringRuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleEditResponseJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleGetResponse struct {
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
	PacketThreshold float64                                   `json:"packet_threshold"`
	JSON            magicNetworkMonitoringRuleGetResponseJSON `json:"-"`
}

// magicNetworkMonitoringRuleGetResponseJSON contains the JSON metadata for the
// struct [MagicNetworkMonitoringRuleGetResponse]
type magicNetworkMonitoringRuleGetResponseJSON struct {
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

func (r *MagicNetworkMonitoringRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleGetResponseJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleNewResponseEnvelope struct {
	Errors   []MagicNetworkMonitoringRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicNetworkMonitoringRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringRuleNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MagicNetworkMonitoringRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicNetworkMonitoringRuleNewResponseEnvelopeJSON    `json:"-"`
}

// magicNetworkMonitoringRuleNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [MagicNetworkMonitoringRuleNewResponseEnvelope]
type magicNetworkMonitoringRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleNewResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    magicNetworkMonitoringRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// magicNetworkMonitoringRuleNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringRuleNewResponseEnvelopeErrors]
type magicNetworkMonitoringRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleNewResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    magicNetworkMonitoringRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// magicNetworkMonitoringRuleNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringRuleNewResponseEnvelopeMessages]
type magicNetworkMonitoringRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicNetworkMonitoringRuleNewResponseEnvelopeSuccess bool

const (
	MagicNetworkMonitoringRuleNewResponseEnvelopeSuccessTrue MagicNetworkMonitoringRuleNewResponseEnvelopeSuccess = true
)

type MagicNetworkMonitoringRuleUpdateResponseEnvelope struct {
	Errors   []MagicNetworkMonitoringRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicNetworkMonitoringRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringRuleUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MagicNetworkMonitoringRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicNetworkMonitoringRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// magicNetworkMonitoringRuleUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [MagicNetworkMonitoringRuleUpdateResponseEnvelope]
type magicNetworkMonitoringRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    magicNetworkMonitoringRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// magicNetworkMonitoringRuleUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringRuleUpdateResponseEnvelopeErrors]
type magicNetworkMonitoringRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    magicNetworkMonitoringRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// magicNetworkMonitoringRuleUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [MagicNetworkMonitoringRuleUpdateResponseEnvelopeMessages]
type magicNetworkMonitoringRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicNetworkMonitoringRuleUpdateResponseEnvelopeSuccess bool

const (
	MagicNetworkMonitoringRuleUpdateResponseEnvelopeSuccessTrue MagicNetworkMonitoringRuleUpdateResponseEnvelopeSuccess = true
)

type MagicNetworkMonitoringRuleListResponseEnvelope struct {
	Errors   []MagicNetworkMonitoringRuleListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicNetworkMonitoringRuleListResponseEnvelopeMessages `json:"messages,required"`
	Result   []MagicNetworkMonitoringRuleListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    MagicNetworkMonitoringRuleListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo MagicNetworkMonitoringRuleListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       magicNetworkMonitoringRuleListResponseEnvelopeJSON       `json:"-"`
}

// magicNetworkMonitoringRuleListResponseEnvelopeJSON contains the JSON metadata
// for the struct [MagicNetworkMonitoringRuleListResponseEnvelope]
type magicNetworkMonitoringRuleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleListResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    magicNetworkMonitoringRuleListResponseEnvelopeErrorsJSON `json:"-"`
}

// magicNetworkMonitoringRuleListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringRuleListResponseEnvelopeErrors]
type magicNetworkMonitoringRuleListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleListResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    magicNetworkMonitoringRuleListResponseEnvelopeMessagesJSON `json:"-"`
}

// magicNetworkMonitoringRuleListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringRuleListResponseEnvelopeMessages]
type magicNetworkMonitoringRuleListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicNetworkMonitoringRuleListResponseEnvelopeSuccess bool

const (
	MagicNetworkMonitoringRuleListResponseEnvelopeSuccessTrue MagicNetworkMonitoringRuleListResponseEnvelopeSuccess = true
)

type MagicNetworkMonitoringRuleListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       magicNetworkMonitoringRuleListResponseEnvelopeResultInfoJSON `json:"-"`
}

// magicNetworkMonitoringRuleListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [MagicNetworkMonitoringRuleListResponseEnvelopeResultInfo]
type magicNetworkMonitoringRuleListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleDeleteResponseEnvelope struct {
	Errors   []MagicNetworkMonitoringRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicNetworkMonitoringRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringRuleDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MagicNetworkMonitoringRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicNetworkMonitoringRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// magicNetworkMonitoringRuleDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [MagicNetworkMonitoringRuleDeleteResponseEnvelope]
type magicNetworkMonitoringRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    magicNetworkMonitoringRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// magicNetworkMonitoringRuleDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringRuleDeleteResponseEnvelopeErrors]
type magicNetworkMonitoringRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    magicNetworkMonitoringRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// magicNetworkMonitoringRuleDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [MagicNetworkMonitoringRuleDeleteResponseEnvelopeMessages]
type magicNetworkMonitoringRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicNetworkMonitoringRuleDeleteResponseEnvelopeSuccess bool

const (
	MagicNetworkMonitoringRuleDeleteResponseEnvelopeSuccessTrue MagicNetworkMonitoringRuleDeleteResponseEnvelopeSuccess = true
)

type MagicNetworkMonitoringRuleEditResponseEnvelope struct {
	Errors   []MagicNetworkMonitoringRuleEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicNetworkMonitoringRuleEditResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringRuleEditResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MagicNetworkMonitoringRuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicNetworkMonitoringRuleEditResponseEnvelopeJSON    `json:"-"`
}

// magicNetworkMonitoringRuleEditResponseEnvelopeJSON contains the JSON metadata
// for the struct [MagicNetworkMonitoringRuleEditResponseEnvelope]
type magicNetworkMonitoringRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleEditResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    magicNetworkMonitoringRuleEditResponseEnvelopeErrorsJSON `json:"-"`
}

// magicNetworkMonitoringRuleEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringRuleEditResponseEnvelopeErrors]
type magicNetworkMonitoringRuleEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleEditResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    magicNetworkMonitoringRuleEditResponseEnvelopeMessagesJSON `json:"-"`
}

// magicNetworkMonitoringRuleEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringRuleEditResponseEnvelopeMessages]
type magicNetworkMonitoringRuleEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicNetworkMonitoringRuleEditResponseEnvelopeSuccess bool

const (
	MagicNetworkMonitoringRuleEditResponseEnvelopeSuccessTrue MagicNetworkMonitoringRuleEditResponseEnvelopeSuccess = true
)

type MagicNetworkMonitoringRuleGetResponseEnvelope struct {
	Errors   []MagicNetworkMonitoringRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MagicNetworkMonitoringRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MagicNetworkMonitoringRuleGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MagicNetworkMonitoringRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    magicNetworkMonitoringRuleGetResponseEnvelopeJSON    `json:"-"`
}

// magicNetworkMonitoringRuleGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [MagicNetworkMonitoringRuleGetResponseEnvelope]
type magicNetworkMonitoringRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    magicNetworkMonitoringRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// magicNetworkMonitoringRuleGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringRuleGetResponseEnvelopeErrors]
type magicNetworkMonitoringRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MagicNetworkMonitoringRuleGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    magicNetworkMonitoringRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// magicNetworkMonitoringRuleGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [MagicNetworkMonitoringRuleGetResponseEnvelopeMessages]
type magicNetworkMonitoringRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MagicNetworkMonitoringRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r magicNetworkMonitoringRuleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type MagicNetworkMonitoringRuleGetResponseEnvelopeSuccess bool

const (
	MagicNetworkMonitoringRuleGetResponseEnvelopeSuccessTrue MagicNetworkMonitoringRuleGetResponseEnvelopeSuccess = true
)

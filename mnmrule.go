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

// MNMRuleService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewMNMRuleService] method instead.
type MNMRuleService struct {
	Options        []option.RequestOption
	Advertisements *MNMRuleAdvertisementService
}

// NewMNMRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMNMRuleService(opts ...option.RequestOption) (r *MNMRuleService) {
	r = &MNMRuleService{}
	r.Options = opts
	r.Advertisements = NewMNMRuleAdvertisementService(opts...)
	return
}

// Create network monitoring rules for account. Currently only supports creating a
// single rule per API request.
func (r *MNMRuleService) New(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MNMRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MNMRuleNewResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update network monitoring rules for account.
func (r *MNMRuleService) Update(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MNMRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MNMRuleUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists network monitoring rules for account.
func (r *MNMRuleService) List(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *[]MNMRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MNMRuleListResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete a network monitoring rule for account.
func (r *MNMRuleService) Delete(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MNMRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MNMRuleDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a network monitoring rule for account.
func (r *MNMRuleService) Edit(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MNMRuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MNMRuleEditResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List a single network monitoring rule for account.
func (r *MNMRuleService) Get(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MNMRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env MNMRuleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MNMRuleNewResponse struct {
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
	JSON            mnmRuleNewResponseJSON `json:"-"`
}

// mnmRuleNewResponseJSON contains the JSON metadata for the struct
// [MNMRuleNewResponse]
type mnmRuleNewResponseJSON struct {
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

func (r *MNMRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleUpdateResponse struct {
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
// [MNMRuleUpdateResponse]
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

func (r *MNMRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleListResponse struct {
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
	PacketThreshold float64                 `json:"packet_threshold"`
	JSON            mnmRuleListResponseJSON `json:"-"`
}

// mnmRuleListResponseJSON contains the JSON metadata for the struct
// [MNMRuleListResponse]
type mnmRuleListResponseJSON struct {
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

func (r *MNMRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleDeleteResponse struct {
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
// [MNMRuleDeleteResponse]
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

func (r *MNMRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleEditResponse struct {
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
	PacketThreshold float64                 `json:"packet_threshold"`
	JSON            mnmRuleEditResponseJSON `json:"-"`
}

// mnmRuleEditResponseJSON contains the JSON metadata for the struct
// [MNMRuleEditResponse]
type mnmRuleEditResponseJSON struct {
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

func (r *MNMRuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleGetResponse struct {
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
// [MNMRuleGetResponse]
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

func (r *MNMRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleNewResponseEnvelope struct {
	Errors   []MNMRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MNMRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   MNMRuleNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MNMRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmRuleNewResponseEnvelopeJSON    `json:"-"`
}

// mnmRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MNMRuleNewResponseEnvelope]
type mnmRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleNewResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    mnmRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MNMRuleNewResponseEnvelopeErrors]
type mnmRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleNewResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    mnmRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [MNMRuleNewResponseEnvelopeMessages]
type mnmRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MNMRuleNewResponseEnvelopeSuccess bool

const (
	MNMRuleNewResponseEnvelopeSuccessTrue MNMRuleNewResponseEnvelopeSuccess = true
)

type MNMRuleUpdateResponseEnvelope struct {
	Errors   []MNMRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MNMRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   MNMRuleUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MNMRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// mnmRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [MNMRuleUpdateResponseEnvelope]
type mnmRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    mnmRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MNMRuleUpdateResponseEnvelopeErrors]
type mnmRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    mnmRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MNMRuleUpdateResponseEnvelopeMessages]
type mnmRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MNMRuleUpdateResponseEnvelopeSuccess bool

const (
	MNMRuleUpdateResponseEnvelopeSuccessTrue MNMRuleUpdateResponseEnvelopeSuccess = true
)

type MNMRuleListResponseEnvelope struct {
	Errors   []MNMRuleListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MNMRuleListResponseEnvelopeMessages `json:"messages,required"`
	Result   []MNMRuleListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    MNMRuleListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo MNMRuleListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       mnmRuleListResponseEnvelopeJSON       `json:"-"`
}

// mnmRuleListResponseEnvelopeJSON contains the JSON metadata for the struct
// [MNMRuleListResponseEnvelope]
type mnmRuleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleListResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    mnmRuleListResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MNMRuleListResponseEnvelopeErrors]
type mnmRuleListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleListResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    mnmRuleListResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MNMRuleListResponseEnvelopeMessages]
type mnmRuleListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MNMRuleListResponseEnvelopeSuccess bool

const (
	MNMRuleListResponseEnvelopeSuccessTrue MNMRuleListResponseEnvelopeSuccess = true
)

type MNMRuleListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       mnmRuleListResponseEnvelopeResultInfoJSON `json:"-"`
}

// mnmRuleListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [MNMRuleListResponseEnvelopeResultInfo]
type mnmRuleListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleDeleteResponseEnvelope struct {
	Errors   []MNMRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MNMRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   MNMRuleDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MNMRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// mnmRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MNMRuleDeleteResponseEnvelope]
type mnmRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    mnmRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MNMRuleDeleteResponseEnvelopeErrors]
type mnmRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    mnmRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MNMRuleDeleteResponseEnvelopeMessages]
type mnmRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MNMRuleDeleteResponseEnvelopeSuccess bool

const (
	MNMRuleDeleteResponseEnvelopeSuccessTrue MNMRuleDeleteResponseEnvelopeSuccess = true
)

type MNMRuleEditResponseEnvelope struct {
	Errors   []MNMRuleEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MNMRuleEditResponseEnvelopeMessages `json:"messages,required"`
	Result   MNMRuleEditResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MNMRuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmRuleEditResponseEnvelopeJSON    `json:"-"`
}

// mnmRuleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [MNMRuleEditResponseEnvelope]
type mnmRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleEditResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    mnmRuleEditResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MNMRuleEditResponseEnvelopeErrors]
type mnmRuleEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleEditResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    mnmRuleEditResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MNMRuleEditResponseEnvelopeMessages]
type mnmRuleEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MNMRuleEditResponseEnvelopeSuccess bool

const (
	MNMRuleEditResponseEnvelopeSuccessTrue MNMRuleEditResponseEnvelopeSuccess = true
)

type MNMRuleGetResponseEnvelope struct {
	Errors   []MNMRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MNMRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   MNMRuleGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success MNMRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    mnmRuleGetResponseEnvelopeJSON    `json:"-"`
}

// mnmRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MNMRuleGetResponseEnvelope]
type mnmRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleGetResponseEnvelopeErrors struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    mnmRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// mnmRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MNMRuleGetResponseEnvelopeErrors]
type mnmRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MNMRuleGetResponseEnvelopeMessages struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    mnmRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// mnmRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [MNMRuleGetResponseEnvelopeMessages]
type mnmRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MNMRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MNMRuleGetResponseEnvelopeSuccess bool

const (
	MNMRuleGetResponseEnvelopeSuccessTrue MNMRuleGetResponseEnvelopeSuccess = true
)

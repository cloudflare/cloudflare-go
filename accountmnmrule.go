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
func (r *AccountMnmRuleService) Get(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MnmRulesSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a network monitoring rule for account.
func (r *AccountMnmRuleService) Update(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MnmRulesSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

// Delete a network monitoring rule for account.
func (r *AccountMnmRuleService) Delete(ctx context.Context, accountIdentifier interface{}, ruleIdentifier interface{}, opts ...option.RequestOption) (res *MnmRulesSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/rules/%v", accountIdentifier, ruleIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create network monitoring rules for account.
func (r *AccountMnmRuleService) MagicNetworkMonitoringRulesNewRules(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmRulesSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Lists network monitoring rules for account.
func (r *AccountMnmRuleService) MagicNetworkMonitoringRulesListRules(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmRulesCollectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update network monitoring rules for account.
func (r *AccountMnmRuleService) MagicNetworkMonitoringRulesUpdateRules(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *MnmRulesSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/mnm/rules", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type MnmRulesCollectionResponse struct {
	Errors     []MnmRulesCollectionResponseError    `json:"errors"`
	Messages   []MnmRulesCollectionResponseMessage  `json:"messages"`
	Result     []MnmRulesCollectionResponseResult   `json:"result,nullable"`
	ResultInfo MnmRulesCollectionResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success MnmRulesCollectionResponseSuccess `json:"success"`
	JSON    mnmRulesCollectionResponseJSON    `json:"-"`
}

// mnmRulesCollectionResponseJSON contains the JSON metadata for the struct
// [MnmRulesCollectionResponse]
type mnmRulesCollectionResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRulesCollectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRulesCollectionResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    mnmRulesCollectionResponseErrorJSON `json:"-"`
}

// mnmRulesCollectionResponseErrorJSON contains the JSON metadata for the struct
// [MnmRulesCollectionResponseError]
type mnmRulesCollectionResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRulesCollectionResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRulesCollectionResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    mnmRulesCollectionResponseMessageJSON `json:"-"`
}

// mnmRulesCollectionResponseMessageJSON contains the JSON metadata for the struct
// [MnmRulesCollectionResponseMessage]
type mnmRulesCollectionResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRulesCollectionResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRulesCollectionResponseResult struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool `json:"automatic_advertisement,required,nullable"`
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The minimum is 60 seconds and maximum is 21600 seconds. The format
	// is XhYmZs where X, Y, and Z durations are optional; however at least one unit
	// must be provided.
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
	PacketThreshold float64                              `json:"packet_threshold"`
	JSON            mnmRulesCollectionResponseResultJSON `json:"-"`
}

// mnmRulesCollectionResponseResultJSON contains the JSON metadata for the struct
// [MnmRulesCollectionResponseResult]
type mnmRulesCollectionResponseResultJSON struct {
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

func (r *MnmRulesCollectionResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRulesCollectionResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       mnmRulesCollectionResponseResultInfoJSON `json:"-"`
}

// mnmRulesCollectionResponseResultInfoJSON contains the JSON metadata for the
// struct [MnmRulesCollectionResponseResultInfo]
type mnmRulesCollectionResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRulesCollectionResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmRulesCollectionResponseSuccess bool

const (
	MnmRulesCollectionResponseSuccessTrue MnmRulesCollectionResponseSuccess = true
)

type MnmRulesSingleResponse struct {
	Errors   []MnmRulesSingleResponseError   `json:"errors"`
	Messages []MnmRulesSingleResponseMessage `json:"messages"`
	Result   MnmRulesSingleResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success MnmRulesSingleResponseSuccess `json:"success"`
	JSON    mnmRulesSingleResponseJSON    `json:"-"`
}

// mnmRulesSingleResponseJSON contains the JSON metadata for the struct
// [MnmRulesSingleResponse]
type mnmRulesSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRulesSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRulesSingleResponseError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    mnmRulesSingleResponseErrorJSON `json:"-"`
}

// mnmRulesSingleResponseErrorJSON contains the JSON metadata for the struct
// [MnmRulesSingleResponseError]
type mnmRulesSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRulesSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRulesSingleResponseMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    mnmRulesSingleResponseMessageJSON `json:"-"`
}

// mnmRulesSingleResponseMessageJSON contains the JSON metadata for the struct
// [MnmRulesSingleResponseMessage]
type mnmRulesSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MnmRulesSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type MnmRulesSingleResponseResult struct {
	// Toggle on if you would like Cloudflare to automatically advertise the IP
	// Prefixes within the rule via Magic Transit when the rule is triggered. Only
	// available for users of Magic Transit.
	AutomaticAdvertisement bool `json:"automatic_advertisement,required,nullable"`
	// The amount of time that the rule threshold must be exceeded to send an alert
	// notification. The minimum is 60 seconds and maximum is 21600 seconds. The format
	// is XhYmZs where X, Y, and Z durations are optional; however at least one unit
	// must be provided.
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
	PacketThreshold float64                          `json:"packet_threshold"`
	JSON            mnmRulesSingleResponseResultJSON `json:"-"`
}

// mnmRulesSingleResponseResultJSON contains the JSON metadata for the struct
// [MnmRulesSingleResponseResult]
type mnmRulesSingleResponseResultJSON struct {
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

func (r *MnmRulesSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type MnmRulesSingleResponseSuccess bool

const (
	MnmRulesSingleResponseSuccessTrue MnmRulesSingleResponseSuccess = true
)

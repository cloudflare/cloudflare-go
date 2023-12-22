// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneFirewallUaRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneFirewallUaRuleService] method
// instead.
type ZoneFirewallUaRuleService struct {
	Options []option.RequestOption
}

// NewZoneFirewallUaRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneFirewallUaRuleService(opts ...option.RequestOption) (r *ZoneFirewallUaRuleService) {
	r = &ZoneFirewallUaRuleService{}
	r.Options = opts
	return
}

// Fetches the details of a User Agent Blocking rule.
func (r *ZoneFirewallUaRuleService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FirewalluablockResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing User Agent Blocking rule.
func (r *ZoneFirewallUaRuleService) Update(ctx context.Context, zoneIdentifier string, id string, body ZoneFirewallUaRuleUpdateParams, opts ...option.RequestOption) (res *FirewalluablockResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an existing User Agent Blocking rule.
func (r *ZoneFirewallUaRuleService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *ZoneFirewallUaRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new User Agent Blocking rule in a zone.
func (r *ZoneFirewallUaRuleService) UserAgentBlockingRulesNewAUserAgentBlockingRule(ctx context.Context, zoneIdentifier string, body ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleParams, opts ...option.RequestOption) (res *FirewalluablockResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/ua_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches User Agent Blocking rules in a zone. You can filter the results using
// several optional parameters.
func (r *ZoneFirewallUaRuleService) UserAgentBlockingRulesListUserAgentBlockingRules(ctx context.Context, zoneIdentifier string, query ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesParams, opts ...option.RequestOption) (res *FirewalluablockResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/ua_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type FirewalluablockResponseCollection struct {
	Errors     []FirewalluablockResponseCollectionError    `json:"errors"`
	Messages   []FirewalluablockResponseCollectionMessage  `json:"messages"`
	Result     []FirewalluablockResponseCollectionResult   `json:"result"`
	ResultInfo FirewalluablockResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success FirewalluablockResponseCollectionSuccess `json:"success"`
	JSON    firewalluablockResponseCollectionJSON    `json:"-"`
}

// firewalluablockResponseCollectionJSON contains the JSON metadata for the struct
// [FirewalluablockResponseCollection]
type firewalluablockResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewalluablockResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewalluablockResponseCollectionError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    firewalluablockResponseCollectionErrorJSON `json:"-"`
}

// firewalluablockResponseCollectionErrorJSON contains the JSON metadata for the
// struct [FirewalluablockResponseCollectionError]
type firewalluablockResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewalluablockResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewalluablockResponseCollectionMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    firewalluablockResponseCollectionMessageJSON `json:"-"`
}

// firewalluablockResponseCollectionMessageJSON contains the JSON metadata for the
// struct [FirewalluablockResponseCollectionMessage]
type firewalluablockResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewalluablockResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewalluablockResponseCollectionResult struct {
	// The unique identifier of the User Agent Blocking rule.
	ID string `json:"id"`
	// The configuration object for the current rule.
	Configuration FirewalluablockResponseCollectionResultConfiguration `json:"configuration"`
	// An informative summary of the rule.
	Description string `json:"description"`
	// The action to apply to a matched request.
	Mode FirewalluablockResponseCollectionResultMode `json:"mode"`
	// When true, indicates that the rule is currently paused.
	Paused bool                                        `json:"paused"`
	JSON   firewalluablockResponseCollectionResultJSON `json:"-"`
}

// firewalluablockResponseCollectionResultJSON contains the JSON metadata for the
// struct [FirewalluablockResponseCollectionResult]
type firewalluablockResponseCollectionResultJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	Mode          apijson.Field
	Paused        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewalluablockResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object for the current rule.
type FirewalluablockResponseCollectionResultConfiguration struct {
	// The configuration target for this rule. You must set the target to `ua` for User
	// Agent Blocking rules.
	Target string `json:"target"`
	// The exact user agent string to match. This value will be compared to the
	// received `User-Agent` HTTP header value.
	Value string                                                   `json:"value"`
	JSON  firewalluablockResponseCollectionResultConfigurationJSON `json:"-"`
}

// firewalluablockResponseCollectionResultConfigurationJSON contains the JSON
// metadata for the struct [FirewalluablockResponseCollectionResultConfiguration]
type firewalluablockResponseCollectionResultConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewalluablockResponseCollectionResultConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type FirewalluablockResponseCollectionResultMode string

const (
	FirewalluablockResponseCollectionResultModeBlock            FirewalluablockResponseCollectionResultMode = "block"
	FirewalluablockResponseCollectionResultModeChallenge        FirewalluablockResponseCollectionResultMode = "challenge"
	FirewalluablockResponseCollectionResultModeJsChallenge      FirewalluablockResponseCollectionResultMode = "js_challenge"
	FirewalluablockResponseCollectionResultModeManagedChallenge FirewalluablockResponseCollectionResultMode = "managed_challenge"
)

type FirewalluablockResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                         `json:"total_count"`
	JSON       firewalluablockResponseCollectionResultInfoJSON `json:"-"`
}

// firewalluablockResponseCollectionResultInfoJSON contains the JSON metadata for
// the struct [FirewalluablockResponseCollectionResultInfo]
type firewalluablockResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewalluablockResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewalluablockResponseCollectionSuccess bool

const (
	FirewalluablockResponseCollectionSuccessTrue FirewalluablockResponseCollectionSuccess = true
)

type FirewalluablockResponseSingle struct {
	Errors   []FirewalluablockResponseSingleError   `json:"errors"`
	Messages []FirewalluablockResponseSingleMessage `json:"messages"`
	Result   interface{}                            `json:"result"`
	// Whether the API call was successful
	Success FirewalluablockResponseSingleSuccess `json:"success"`
	JSON    firewalluablockResponseSingleJSON    `json:"-"`
}

// firewalluablockResponseSingleJSON contains the JSON metadata for the struct
// [FirewalluablockResponseSingle]
type firewalluablockResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewalluablockResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewalluablockResponseSingleError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    firewalluablockResponseSingleErrorJSON `json:"-"`
}

// firewalluablockResponseSingleErrorJSON contains the JSON metadata for the struct
// [FirewalluablockResponseSingleError]
type firewalluablockResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewalluablockResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewalluablockResponseSingleMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    firewalluablockResponseSingleMessageJSON `json:"-"`
}

// firewalluablockResponseSingleMessageJSON contains the JSON metadata for the
// struct [FirewalluablockResponseSingleMessage]
type firewalluablockResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewalluablockResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewalluablockResponseSingleSuccess bool

const (
	FirewalluablockResponseSingleSuccessTrue FirewalluablockResponseSingleSuccess = true
)

type ZoneFirewallUaRuleDeleteResponse struct {
	Errors   []ZoneFirewallUaRuleDeleteResponseError   `json:"errors"`
	Messages []ZoneFirewallUaRuleDeleteResponseMessage `json:"messages"`
	Result   ZoneFirewallUaRuleDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallUaRuleDeleteResponseSuccess `json:"success"`
	JSON    zoneFirewallUaRuleDeleteResponseJSON    `json:"-"`
}

// zoneFirewallUaRuleDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallUaRuleDeleteResponse]
type zoneFirewallUaRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallUaRuleDeleteResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneFirewallUaRuleDeleteResponseErrorJSON `json:"-"`
}

// zoneFirewallUaRuleDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneFirewallUaRuleDeleteResponseError]
type zoneFirewallUaRuleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallUaRuleDeleteResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneFirewallUaRuleDeleteResponseMessageJSON `json:"-"`
}

// zoneFirewallUaRuleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneFirewallUaRuleDeleteResponseMessage]
type zoneFirewallUaRuleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallUaRuleDeleteResponseResult struct {
	// The unique identifier of the User Agent Blocking rule.
	ID   string                                     `json:"id"`
	JSON zoneFirewallUaRuleDeleteResponseResultJSON `json:"-"`
}

// zoneFirewallUaRuleDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneFirewallUaRuleDeleteResponseResult]
type zoneFirewallUaRuleDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFirewallUaRuleDeleteResponseSuccess bool

const (
	ZoneFirewallUaRuleDeleteResponseSuccessTrue ZoneFirewallUaRuleDeleteResponseSuccess = true
)

type ZoneFirewallUaRuleUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallUaRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesParams struct {
	// A string to search for in the description of existing rules.
	Description param.Field[string] `query:"description"`
	// A string to search for in the description of existing rules.
	DescriptionSearch param.Field[string] `query:"description_search"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
	// A string to search for in the user agent values of existing rules.
	UaSearch param.Field[string] `query:"ua_search"`
}

// URLQuery serializes
// [ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesParams]'s
// query parameters as `url.Values`.
func (r ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

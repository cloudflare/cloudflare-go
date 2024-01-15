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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
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
func (r *ZoneFirewallUaRuleService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *ZoneFirewallUaRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing User Agent Blocking rule.
func (r *ZoneFirewallUaRuleService) Update(ctx context.Context, zoneIdentifier string, id string, body ZoneFirewallUaRuleUpdateParams, opts ...option.RequestOption) (res *ZoneFirewallUaRuleUpdateResponse, err error) {
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
func (r *ZoneFirewallUaRuleService) UserAgentBlockingRulesNewAUserAgentBlockingRule(ctx context.Context, zoneIdentifier string, body ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleParams, opts ...option.RequestOption) (res *ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/ua_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches User Agent Blocking rules in a zone. You can filter the results using
// several optional parameters.
func (r *ZoneFirewallUaRuleService) UserAgentBlockingRulesListUserAgentBlockingRules(ctx context.Context, zoneIdentifier string, query ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesParams, opts ...option.RequestOption) (res *shared.Page[ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/ua_rules", zoneIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

type ZoneFirewallUaRuleGetResponse struct {
	Errors   []ZoneFirewallUaRuleGetResponseError   `json:"errors"`
	Messages []ZoneFirewallUaRuleGetResponseMessage `json:"messages"`
	Result   interface{}                            `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallUaRuleGetResponseSuccess `json:"success"`
	JSON    zoneFirewallUaRuleGetResponseJSON    `json:"-"`
}

// zoneFirewallUaRuleGetResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallUaRuleGetResponse]
type zoneFirewallUaRuleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallUaRuleGetResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneFirewallUaRuleGetResponseErrorJSON `json:"-"`
}

// zoneFirewallUaRuleGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneFirewallUaRuleGetResponseError]
type zoneFirewallUaRuleGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallUaRuleGetResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneFirewallUaRuleGetResponseMessageJSON `json:"-"`
}

// zoneFirewallUaRuleGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneFirewallUaRuleGetResponseMessage]
type zoneFirewallUaRuleGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFirewallUaRuleGetResponseSuccess bool

const (
	ZoneFirewallUaRuleGetResponseSuccessTrue ZoneFirewallUaRuleGetResponseSuccess = true
)

type ZoneFirewallUaRuleUpdateResponse struct {
	Errors   []ZoneFirewallUaRuleUpdateResponseError   `json:"errors"`
	Messages []ZoneFirewallUaRuleUpdateResponseMessage `json:"messages"`
	Result   interface{}                               `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallUaRuleUpdateResponseSuccess `json:"success"`
	JSON    zoneFirewallUaRuleUpdateResponseJSON    `json:"-"`
}

// zoneFirewallUaRuleUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallUaRuleUpdateResponse]
type zoneFirewallUaRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallUaRuleUpdateResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneFirewallUaRuleUpdateResponseErrorJSON `json:"-"`
}

// zoneFirewallUaRuleUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneFirewallUaRuleUpdateResponseError]
type zoneFirewallUaRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallUaRuleUpdateResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneFirewallUaRuleUpdateResponseMessageJSON `json:"-"`
}

// zoneFirewallUaRuleUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneFirewallUaRuleUpdateResponseMessage]
type zoneFirewallUaRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFirewallUaRuleUpdateResponseSuccess bool

const (
	ZoneFirewallUaRuleUpdateResponseSuccessTrue ZoneFirewallUaRuleUpdateResponseSuccess = true
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

type ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponse struct {
	Errors   []ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseError   `json:"errors"`
	Messages []ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseMessage `json:"messages"`
	Result   interface{}                                                                        `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseSuccess `json:"success"`
	JSON    zoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseJSON    `json:"-"`
}

// zoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseJSON
// contains the JSON metadata for the struct
// [ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponse]
type zoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseError struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    zoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseErrorJSON `json:"-"`
}

// zoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseError]
type zoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseMessage struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    zoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseMessageJSON `json:"-"`
}

// zoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseMessage]
type zoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseSuccess bool

const (
	ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseSuccessTrue ZoneFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseSuccess = true
)

type ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponse struct {
	// The unique identifier of the User Agent Blocking rule.
	ID string `json:"id"`
	// The configuration object for the current rule.
	Configuration ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseConfiguration `json:"configuration"`
	// An informative summary of the rule.
	Description string `json:"description"`
	// The action to apply to a matched request.
	Mode ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseMode `json:"mode"`
	// When true, indicates that the rule is currently paused.
	Paused bool                                                                           `json:"paused"`
	JSON   zoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseJSON `json:"-"`
}

// zoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseJSON
// contains the JSON metadata for the struct
// [ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponse]
type zoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	Mode          apijson.Field
	Paused        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object for the current rule.
type ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseConfiguration struct {
	// The configuration target for this rule. You must set the target to `ua` for User
	// Agent Blocking rules.
	Target string `json:"target"`
	// The exact user agent string to match. This value will be compared to the
	// received `User-Agent` HTTP header value.
	Value string                                                                                      `json:"value"`
	JSON  zoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseConfigurationJSON `json:"-"`
}

// zoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseConfiguration]
type zoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseMode string

const (
	ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseModeBlock            ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseMode = "block"
	ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseModeChallenge        ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseMode = "challenge"
	ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseModeJsChallenge      ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseMode = "js_challenge"
	ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseModeManagedChallenge ZoneFirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseMode = "managed_challenge"
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

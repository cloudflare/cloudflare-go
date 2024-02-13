// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// FirewallUaRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallUaRuleService] method
// instead.
type FirewallUaRuleService struct {
	Options []option.RequestOption
}

// NewFirewallUaRuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFirewallUaRuleService(opts ...option.RequestOption) (r *FirewallUaRuleService) {
	r = &FirewallUaRuleService{}
	r.Options = opts
	return
}

// Updates an existing User Agent Blocking rule.
func (r *FirewallUaRuleService) Update(ctx context.Context, zoneIdentifier string, id string, body FirewallUaRuleUpdateParams, opts ...option.RequestOption) (res *FirewallUaRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallUaRuleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing User Agent Blocking rule.
func (r *FirewallUaRuleService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FirewallUaRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallUaRuleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a User Agent Blocking rule.
func (r *FirewallUaRuleService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FirewallUaRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallUaRuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Creates a new User Agent Blocking rule in a zone.
func (r *FirewallUaRuleService) UserAgentBlockingRulesNewAUserAgentBlockingRule(ctx context.Context, zoneIdentifier string, body FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleParams, opts ...option.RequestOption) (res *FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/ua_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches User Agent Blocking rules in a zone. You can filter the results using
// several optional parameters.
func (r *FirewallUaRuleService) UserAgentBlockingRulesListUserAgentBlockingRules(ctx context.Context, zoneIdentifier string, query FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesParams, opts ...option.RequestOption) (res *[]FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/ua_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [FirewallUaRuleUpdateResponseUnknown] or
// [shared.UnionString].
type FirewallUaRuleUpdateResponse interface {
	ImplementsFirewallUaRuleUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallUaRuleUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallUaRuleDeleteResponse struct {
	// The unique identifier of the User Agent Blocking rule.
	ID   string                           `json:"id"`
	JSON firewallUaRuleDeleteResponseJSON `json:"-"`
}

// firewallUaRuleDeleteResponseJSON contains the JSON metadata for the struct
// [FirewallUaRuleDeleteResponse]
type firewallUaRuleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [FirewallUaRuleGetResponseUnknown] or [shared.UnionString].
type FirewallUaRuleGetResponse interface {
	ImplementsFirewallUaRuleGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallUaRuleGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by
// [FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseUnknown]
// or [shared.UnionString].
type FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponse interface {
	ImplementsFirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponse struct {
	// The unique identifier of the User Agent Blocking rule.
	ID string `json:"id"`
	// The configuration object for the current rule.
	Configuration FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseConfiguration `json:"configuration"`
	// An informative summary of the rule.
	Description string `json:"description"`
	// The action to apply to a matched request.
	Mode FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseMode `json:"mode"`
	// When true, indicates that the rule is currently paused.
	Paused bool                                                                       `json:"paused"`
	JSON   firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseJSON `json:"-"`
}

// firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseJSON
// contains the JSON metadata for the struct
// [FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponse]
type firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	Mode          apijson.Field
	Paused        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object for the current rule.
type FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseConfiguration struct {
	// The configuration target for this rule. You must set the target to `ua` for User
	// Agent Blocking rules.
	Target string `json:"target"`
	// The exact user agent string to match. This value will be compared to the
	// received `User-Agent` HTTP header value.
	Value string                                                                                  `json:"value"`
	JSON  firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseConfigurationJSON `json:"-"`
}

// firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseConfiguration]
type firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseMode string

const (
	FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseModeBlock            FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseMode = "block"
	FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseModeChallenge        FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseMode = "challenge"
	FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseModeJsChallenge      FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseMode = "js_challenge"
	FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseModeManagedChallenge FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseMode = "managed_challenge"
)

type FirewallUaRuleUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallUaRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallUaRuleUpdateResponseEnvelope struct {
	Errors   []FirewallUaRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallUaRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallUaRuleUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallUaRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallUaRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// firewallUaRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallUaRuleUpdateResponseEnvelope]
type firewallUaRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallUaRuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    firewallUaRuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallUaRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallUaRuleUpdateResponseEnvelopeErrors]
type firewallUaRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallUaRuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    firewallUaRuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallUaRuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallUaRuleUpdateResponseEnvelopeMessages]
type firewallUaRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallUaRuleUpdateResponseEnvelopeSuccess bool

const (
	FirewallUaRuleUpdateResponseEnvelopeSuccessTrue FirewallUaRuleUpdateResponseEnvelopeSuccess = true
)

type FirewallUaRuleDeleteResponseEnvelope struct {
	Errors   []FirewallUaRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallUaRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallUaRuleDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallUaRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallUaRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// firewallUaRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallUaRuleDeleteResponseEnvelope]
type firewallUaRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallUaRuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    firewallUaRuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallUaRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallUaRuleDeleteResponseEnvelopeErrors]
type firewallUaRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallUaRuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    firewallUaRuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallUaRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallUaRuleDeleteResponseEnvelopeMessages]
type firewallUaRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallUaRuleDeleteResponseEnvelopeSuccess bool

const (
	FirewallUaRuleDeleteResponseEnvelopeSuccessTrue FirewallUaRuleDeleteResponseEnvelopeSuccess = true
)

type FirewallUaRuleGetResponseEnvelope struct {
	Errors   []FirewallUaRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallUaRuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallUaRuleGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallUaRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallUaRuleGetResponseEnvelopeJSON    `json:"-"`
}

// firewallUaRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallUaRuleGetResponseEnvelope]
type firewallUaRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallUaRuleGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    firewallUaRuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallUaRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallUaRuleGetResponseEnvelopeErrors]
type firewallUaRuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallUaRuleGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    firewallUaRuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallUaRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FirewallUaRuleGetResponseEnvelopeMessages]
type firewallUaRuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallUaRuleGetResponseEnvelopeSuccess bool

const (
	FirewallUaRuleGetResponseEnvelopeSuccessTrue FirewallUaRuleGetResponseEnvelopeSuccess = true
)

type FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelope struct {
	Errors   []FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeJSON    `json:"-"`
}

// firewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelope]
type firewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeErrors struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    firewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeErrors]
type firewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeMessages struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    firewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeMessages]
type firewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeSuccess bool

const (
	FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeSuccessTrue FirewallUaRuleUserAgentBlockingRulesNewAUserAgentBlockingRuleResponseEnvelopeSuccess = true
)

type FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesParams struct {
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
// [FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesParams]'s query
// parameters as `url.Values`.
func (r FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelope struct {
	Errors   []FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   []FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeJSON       `json:"-"`
}

// firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelope]
type firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeErrors struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeErrors]
type firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeMessages struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeMessages]
type firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeSuccess bool

const (
	FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeSuccessTrue FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeSuccess = true
)

type FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                      `json:"total_count"`
	JSON       firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeResultInfoJSON `json:"-"`
}

// firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeResultInfo]
type firewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleUserAgentBlockingRulesListUserAgentBlockingRulesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

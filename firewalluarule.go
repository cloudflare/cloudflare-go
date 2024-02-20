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

// Creates a new User Agent Blocking rule in a zone.
func (r *FirewallUaRuleService) New(ctx context.Context, zoneIdentifier string, body FirewallUaRuleNewParams, opts ...option.RequestOption) (res *FirewallUaRuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallUaRuleNewResponseEnvelope
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
func (r *FirewallUaRuleService) List(ctx context.Context, zoneIdentifier string, query FirewallUaRuleListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[FirewallUaRuleListResponse], err error) {
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

// Fetches User Agent Blocking rules in a zone. You can filter the results using
// several optional parameters.
func (r *FirewallUaRuleService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query FirewallUaRuleListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[FirewallUaRuleListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
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

// Updates an existing User Agent Blocking rule.
func (r *FirewallUaRuleService) Replace(ctx context.Context, zoneIdentifier string, id string, body FirewallUaRuleReplaceParams, opts ...option.RequestOption) (res *FirewallUaRuleReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallUaRuleReplaceResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [FirewallUaRuleNewResponseUnknown] or [shared.UnionString].
type FirewallUaRuleNewResponse interface {
	ImplementsFirewallUaRuleNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallUaRuleNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallUaRuleListResponse struct {
	// The unique identifier of the User Agent Blocking rule.
	ID string `json:"id"`
	// The configuration object for the current rule.
	Configuration FirewallUaRuleListResponseConfiguration `json:"configuration"`
	// An informative summary of the rule.
	Description string `json:"description"`
	// The action to apply to a matched request.
	Mode FirewallUaRuleListResponseMode `json:"mode"`
	// When true, indicates that the rule is currently paused.
	Paused bool                           `json:"paused"`
	JSON   firewallUaRuleListResponseJSON `json:"-"`
}

// firewallUaRuleListResponseJSON contains the JSON metadata for the struct
// [FirewallUaRuleListResponse]
type firewallUaRuleListResponseJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	Mode          apijson.Field
	Paused        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallUaRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration object for the current rule.
type FirewallUaRuleListResponseConfiguration struct {
	// The configuration target for this rule. You must set the target to `ua` for User
	// Agent Blocking rules.
	Target string `json:"target"`
	// The exact user agent string to match. This value will be compared to the
	// received `User-Agent` HTTP header value.
	Value string                                      `json:"value"`
	JSON  firewallUaRuleListResponseConfigurationJSON `json:"-"`
}

// firewallUaRuleListResponseConfigurationJSON contains the JSON metadata for the
// struct [FirewallUaRuleListResponseConfiguration]
type firewallUaRuleListResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleListResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request.
type FirewallUaRuleListResponseMode string

const (
	FirewallUaRuleListResponseModeBlock            FirewallUaRuleListResponseMode = "block"
	FirewallUaRuleListResponseModeChallenge        FirewallUaRuleListResponseMode = "challenge"
	FirewallUaRuleListResponseModeJsChallenge      FirewallUaRuleListResponseMode = "js_challenge"
	FirewallUaRuleListResponseModeManagedChallenge FirewallUaRuleListResponseMode = "managed_challenge"
)

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

// Union satisfied by [FirewallUaRuleReplaceResponseUnknown] or
// [shared.UnionString].
type FirewallUaRuleReplaceResponse interface {
	ImplementsFirewallUaRuleReplaceResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallUaRuleReplaceResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallUaRuleNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallUaRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallUaRuleNewResponseEnvelope struct {
	Errors   []FirewallUaRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallUaRuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallUaRuleNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallUaRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallUaRuleNewResponseEnvelopeJSON    `json:"-"`
}

// firewallUaRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallUaRuleNewResponseEnvelope]
type firewallUaRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallUaRuleNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    firewallUaRuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallUaRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallUaRuleNewResponseEnvelopeErrors]
type firewallUaRuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallUaRuleNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    firewallUaRuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallUaRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FirewallUaRuleNewResponseEnvelopeMessages]
type firewallUaRuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallUaRuleNewResponseEnvelopeSuccess bool

const (
	FirewallUaRuleNewResponseEnvelopeSuccessTrue FirewallUaRuleNewResponseEnvelopeSuccess = true
)

type FirewallUaRuleListParams struct {
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

// URLQuery serializes [FirewallUaRuleListParams]'s query parameters as
// `url.Values`.
func (r FirewallUaRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

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

type FirewallUaRuleReplaceParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallUaRuleReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallUaRuleReplaceResponseEnvelope struct {
	Errors   []FirewallUaRuleReplaceResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallUaRuleReplaceResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallUaRuleReplaceResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallUaRuleReplaceResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallUaRuleReplaceResponseEnvelopeJSON    `json:"-"`
}

// firewallUaRuleReplaceResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallUaRuleReplaceResponseEnvelope]
type firewallUaRuleReplaceResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleReplaceResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallUaRuleReplaceResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    firewallUaRuleReplaceResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallUaRuleReplaceResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallUaRuleReplaceResponseEnvelopeErrors]
type firewallUaRuleReplaceResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleReplaceResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallUaRuleReplaceResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    firewallUaRuleReplaceResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallUaRuleReplaceResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallUaRuleReplaceResponseEnvelopeMessages]
type firewallUaRuleReplaceResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUaRuleReplaceResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallUaRuleReplaceResponseEnvelopeSuccess bool

const (
	FirewallUaRuleReplaceResponseEnvelopeSuccessTrue FirewallUaRuleReplaceResponseEnvelopeSuccess = true
)

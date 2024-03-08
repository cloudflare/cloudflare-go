// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// FirewallUARuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallUARuleService] method
// instead.
type FirewallUARuleService struct {
	Options []option.RequestOption
}

// NewFirewallUARuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFirewallUARuleService(opts ...option.RequestOption) (r *FirewallUARuleService) {
	r = &FirewallUARuleService{}
	r.Options = opts
	return
}

// Creates a new User Agent Blocking rule in a zone.
func (r *FirewallUARuleService) New(ctx context.Context, zoneIdentifier string, body FirewallUARuleNewParams, opts ...option.RequestOption) (res *FirewallUARuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallUARuleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/ua_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing User Agent Blocking rule.
func (r *FirewallUARuleService) Update(ctx context.Context, zoneIdentifier string, id string, body FirewallUARuleUpdateParams, opts ...option.RequestOption) (res *FirewallUARuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallUARuleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches User Agent Blocking rules in a zone. You can filter the results using
// several optional parameters.
func (r *FirewallUARuleService) List(ctx context.Context, zoneIdentifier string, query FirewallUARuleListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[FirewallUARuleListResponse], err error) {
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
func (r *FirewallUARuleService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query FirewallUARuleListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[FirewallUARuleListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing User Agent Blocking rule.
func (r *FirewallUARuleService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FirewallUARuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallUARuleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a User Agent Blocking rule.
func (r *FirewallUARuleService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FirewallUARuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallUARuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [FirewallUARuleNewResponseUnknown] or [shared.UnionString].
type FirewallUARuleNewResponse interface {
	ImplementsFirewallUARuleNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallUARuleNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [FirewallUARuleUpdateResponseUnknown] or
// [shared.UnionString].
type FirewallUARuleUpdateResponse interface {
	ImplementsFirewallUARuleUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallUARuleUpdateResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallUARuleListResponse struct {
	// The unique identifier of the User Agent Blocking rule.
	ID string `json:"id"`
	// The configuration object for the current rule.
	Configuration FirewallUARuleListResponseConfiguration `json:"configuration"`
	// An informative summary of the rule.
	Description string `json:"description"`
	// The action to apply to a matched request.
	Mode FirewallUARuleListResponseMode `json:"mode"`
	// When true, indicates that the rule is currently paused.
	Paused bool                           `json:"paused"`
	JSON   firewallUARuleListResponseJSON `json:"-"`
}

// firewallUARuleListResponseJSON contains the JSON metadata for the struct
// [FirewallUARuleListResponse]
type firewallUARuleListResponseJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	Mode          apijson.Field
	Paused        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *FirewallUARuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleListResponseJSON) RawJSON() string {
	return r.raw
}

// The configuration object for the current rule.
type FirewallUARuleListResponseConfiguration struct {
	// The configuration target for this rule. You must set the target to `ua` for User
	// Agent Blocking rules.
	Target string `json:"target"`
	// The exact user agent string to match. This value will be compared to the
	// received `User-Agent` HTTP header value.
	Value string                                      `json:"value"`
	JSON  firewallUARuleListResponseConfigurationJSON `json:"-"`
}

// firewallUARuleListResponseConfigurationJSON contains the JSON metadata for the
// struct [FirewallUARuleListResponseConfiguration]
type firewallUARuleListResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUARuleListResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleListResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type FirewallUARuleListResponseMode string

const (
	FirewallUARuleListResponseModeBlock            FirewallUARuleListResponseMode = "block"
	FirewallUARuleListResponseModeChallenge        FirewallUARuleListResponseMode = "challenge"
	FirewallUARuleListResponseModeJsChallenge      FirewallUARuleListResponseMode = "js_challenge"
	FirewallUARuleListResponseModeManagedChallenge FirewallUARuleListResponseMode = "managed_challenge"
)

type FirewallUARuleDeleteResponse struct {
	// The unique identifier of the User Agent Blocking rule.
	ID   string                           `json:"id"`
	JSON firewallUARuleDeleteResponseJSON `json:"-"`
}

// firewallUARuleDeleteResponseJSON contains the JSON metadata for the struct
// [FirewallUARuleDeleteResponse]
type firewallUARuleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUARuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [FirewallUARuleGetResponseUnknown] or [shared.UnionString].
type FirewallUARuleGetResponse interface {
	ImplementsFirewallUARuleGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallUARuleGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type FirewallUARuleNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallUARuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallUARuleNewResponseEnvelope struct {
	Errors   []FirewallUARuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallUARuleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallUARuleNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallUARuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallUARuleNewResponseEnvelopeJSON    `json:"-"`
}

// firewallUARuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallUARuleNewResponseEnvelope]
type firewallUARuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUARuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallUARuleNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    firewallUARuleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallUARuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallUARuleNewResponseEnvelopeErrors]
type firewallUARuleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUARuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallUARuleNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    firewallUARuleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallUARuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FirewallUARuleNewResponseEnvelopeMessages]
type firewallUARuleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUARuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallUARuleNewResponseEnvelopeSuccess bool

const (
	FirewallUARuleNewResponseEnvelopeSuccessTrue FirewallUARuleNewResponseEnvelopeSuccess = true
)

type FirewallUARuleUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallUARuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallUARuleUpdateResponseEnvelope struct {
	Errors   []FirewallUARuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallUARuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallUARuleUpdateResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallUARuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallUARuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// firewallUARuleUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallUARuleUpdateResponseEnvelope]
type firewallUARuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUARuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallUARuleUpdateResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    firewallUARuleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallUARuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallUARuleUpdateResponseEnvelopeErrors]
type firewallUARuleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUARuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallUARuleUpdateResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    firewallUARuleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallUARuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallUARuleUpdateResponseEnvelopeMessages]
type firewallUARuleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUARuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallUARuleUpdateResponseEnvelopeSuccess bool

const (
	FirewallUARuleUpdateResponseEnvelopeSuccessTrue FirewallUARuleUpdateResponseEnvelopeSuccess = true
)

type FirewallUARuleListParams struct {
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
	UASearch param.Field[string] `query:"ua_search"`
}

// URLQuery serializes [FirewallUARuleListParams]'s query parameters as
// `url.Values`.
func (r FirewallUARuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallUARuleDeleteResponseEnvelope struct {
	Errors   []FirewallUARuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallUARuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallUARuleDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallUARuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallUARuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// firewallUARuleDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallUARuleDeleteResponseEnvelope]
type firewallUARuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUARuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallUARuleDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    firewallUARuleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallUARuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallUARuleDeleteResponseEnvelopeErrors]
type firewallUARuleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUARuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallUARuleDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    firewallUARuleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallUARuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallUARuleDeleteResponseEnvelopeMessages]
type firewallUARuleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUARuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallUARuleDeleteResponseEnvelopeSuccess bool

const (
	FirewallUARuleDeleteResponseEnvelopeSuccessTrue FirewallUARuleDeleteResponseEnvelopeSuccess = true
)

type FirewallUARuleGetResponseEnvelope struct {
	Errors   []FirewallUARuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallUARuleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallUARuleGetResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallUARuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallUARuleGetResponseEnvelopeJSON    `json:"-"`
}

// firewallUARuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [FirewallUARuleGetResponseEnvelope]
type firewallUARuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUARuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallUARuleGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    firewallUARuleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallUARuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallUARuleGetResponseEnvelopeErrors]
type firewallUARuleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUARuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallUARuleGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    firewallUARuleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallUARuleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FirewallUARuleGetResponseEnvelopeMessages]
type firewallUARuleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUARuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallUARuleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallUARuleGetResponseEnvelopeSuccess bool

const (
	FirewallUARuleGetResponseEnvelopeSuccessTrue FirewallUARuleGetResponseEnvelopeSuccess = true
)

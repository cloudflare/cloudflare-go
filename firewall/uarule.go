// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// UARuleService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewUARuleService] method instead.
type UARuleService struct {
	Options []option.RequestOption
}

// NewUARuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUARuleService(opts ...option.RequestOption) (r *UARuleService) {
	r = &UARuleService{}
	r.Options = opts
	return
}

// Creates a new User Agent Blocking rule in a zone.
func (r *UARuleService) New(ctx context.Context, zoneIdentifier string, body UARuleNewParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env UARuleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/ua_rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing User Agent Blocking rule.
func (r *UARuleService) Update(ctx context.Context, zoneIdentifier string, id string, body UARuleUpdateParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env UARuleUpdateResponseEnvelope
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
func (r *UARuleService) List(ctx context.Context, zoneIdentifier string, query UARuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[UserAgentRule], err error) {
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
func (r *UARuleService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query UARuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[UserAgentRule] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing User Agent Blocking rule.
func (r *UARuleService) Delete(ctx context.Context, zoneIdentifier string, id string, body UARuleDeleteParams, opts ...option.RequestOption) (res *UARuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UARuleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a User Agent Blocking rule.
func (r *UARuleService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env UARuleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/ua_rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserAgentRule struct {
	// The unique identifier of the User Agent Blocking rule.
	ID string `json:"id"`
	// The configuration object for the current rule.
	Configuration UserAgentRuleConfiguration `json:"configuration"`
	// An informative summary of the rule.
	Description string `json:"description"`
	// The action to apply to a matched request.
	Mode UserAgentRuleMode `json:"mode"`
	// When true, indicates that the rule is currently paused.
	Paused bool              `json:"paused"`
	JSON   userAgentRuleJSON `json:"-"`
}

// userAgentRuleJSON contains the JSON metadata for the struct [UserAgentRule]
type userAgentRuleJSON struct {
	ID            apijson.Field
	Configuration apijson.Field
	Description   apijson.Field
	Mode          apijson.Field
	Paused        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserAgentRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAgentRuleJSON) RawJSON() string {
	return r.raw
}

// The configuration object for the current rule.
type UserAgentRuleConfiguration struct {
	// The configuration target for this rule. You must set the target to `ua` for User
	// Agent Blocking rules.
	Target string `json:"target"`
	// The exact user agent string to match. This value will be compared to the
	// received `User-Agent` HTTP header value.
	Value string                         `json:"value"`
	JSON  userAgentRuleConfigurationJSON `json:"-"`
}

// userAgentRuleConfigurationJSON contains the JSON metadata for the struct
// [UserAgentRuleConfiguration]
type userAgentRuleConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserAgentRuleConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userAgentRuleConfigurationJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type UserAgentRuleMode string

const (
	UserAgentRuleModeBlock            UserAgentRuleMode = "block"
	UserAgentRuleModeChallenge        UserAgentRuleMode = "challenge"
	UserAgentRuleModeJsChallenge      UserAgentRuleMode = "js_challenge"
	UserAgentRuleModeManagedChallenge UserAgentRuleMode = "managed_challenge"
)

func (r UserAgentRuleMode) IsKnown() bool {
	switch r {
	case UserAgentRuleModeBlock, UserAgentRuleModeChallenge, UserAgentRuleModeJsChallenge, UserAgentRuleModeManagedChallenge:
		return true
	}
	return false
}

type UARuleDeleteResponse struct {
	// The unique identifier of the User Agent Blocking rule.
	ID   string                   `json:"id"`
	JSON uaRuleDeleteResponseJSON `json:"-"`
}

// uaRuleDeleteResponseJSON contains the JSON metadata for the struct
// [UARuleDeleteResponse]
type uaRuleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UARuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type UARuleNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r UARuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type UARuleNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"messages,required"`
	Result   shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion `json:"result,required"`
	// Whether the API call was successful
	Success UARuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    uaRuleNewResponseEnvelopeJSON    `json:"-"`
}

// uaRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [UARuleNewResponseEnvelope]
type uaRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UARuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UARuleNewResponseEnvelopeSuccess bool

const (
	UARuleNewResponseEnvelopeSuccessTrue UARuleNewResponseEnvelopeSuccess = true
)

func (r UARuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UARuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UARuleUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r UARuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type UARuleUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"messages,required"`
	Result   shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion `json:"result,required"`
	// Whether the API call was successful
	Success UARuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    uaRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// uaRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [UARuleUpdateResponseEnvelope]
type uaRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UARuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UARuleUpdateResponseEnvelopeSuccess bool

const (
	UARuleUpdateResponseEnvelopeSuccessTrue UARuleUpdateResponseEnvelopeSuccess = true
)

func (r UARuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UARuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UARuleListParams struct {
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

// URLQuery serializes [UARuleListParams]'s query parameters as `url.Values`.
func (r UARuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UARuleDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r UARuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type UARuleDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   UARuleDeleteResponse                                      `json:"result,required,nullable"`
	// Whether the API call was successful
	Success UARuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    uaRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// uaRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [UARuleDeleteResponseEnvelope]
type uaRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UARuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UARuleDeleteResponseEnvelopeSuccess bool

const (
	UARuleDeleteResponseEnvelopeSuccessTrue UARuleDeleteResponseEnvelopeSuccess = true
)

func (r UARuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UARuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UARuleGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"messages,required"`
	Result   shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion `json:"result,required"`
	// Whether the API call was successful
	Success UARuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    uaRuleGetResponseEnvelopeJSON    `json:"-"`
}

// uaRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UARuleGetResponseEnvelope]
type uaRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UARuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r uaRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type UARuleGetResponseEnvelopeSuccess bool

const (
	UARuleGetResponseEnvelopeSuccessTrue UARuleGetResponseEnvelopeSuccess = true
)

func (r UARuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UARuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

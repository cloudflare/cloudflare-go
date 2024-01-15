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
)

// ZoneFirewallRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneFirewallRuleService] method
// instead.
type ZoneFirewallRuleService struct {
	Options []option.RequestOption
}

// NewZoneFirewallRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneFirewallRuleService(opts ...option.RequestOption) (r *ZoneFirewallRuleService) {
	r = &ZoneFirewallRuleService{}
	r.Options = opts
	return
}

// Fetches the details of a firewall rule.
func (r *ZoneFirewallRuleService) Get(ctx context.Context, zoneIdentifier string, id string, query ZoneFirewallRuleGetParams, opts ...option.RequestOption) (res *ZoneFirewallRuleGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates an existing firewall rule.
func (r *ZoneFirewallRuleService) Update(ctx context.Context, zoneIdentifier string, id string, body ZoneFirewallRuleUpdateParams, opts ...option.RequestOption) (res *ZoneFirewallRuleUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an existing firewall rule.
func (r *ZoneFirewallRuleService) Delete(ctx context.Context, zoneIdentifier string, id string, body ZoneFirewallRuleDeleteParams, opts ...option.RequestOption) (res *ZoneFirewallRuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Create one or more firewall rules.
func (r *ZoneFirewallRuleService) FirewallRulesNewFirewallRules(ctx context.Context, zoneIdentifier string, body ZoneFirewallRuleFirewallRulesNewFirewallRulesParams, opts ...option.RequestOption) (res *ZoneFirewallRuleFirewallRulesNewFirewallRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches firewall rules in a zone. You can filter the results using several
// optional parameters.
func (r *ZoneFirewallRuleService) FirewallRulesListFirewallRules(ctx context.Context, zoneIdentifier string, query ZoneFirewallRuleFirewallRulesListFirewallRulesParams, opts ...option.RequestOption) (res *shared.Page[ZoneFirewallRuleFirewallRulesListFirewallRulesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
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

// Updates one or more existing firewall rules.
func (r *ZoneFirewallRuleService) FirewallRulesUpdateFirewallRules(ctx context.Context, zoneIdentifier string, body ZoneFirewallRuleFirewallRulesUpdateFirewallRulesParams, opts ...option.RequestOption) (res *ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Updates the priority of existing firewall rules.
func (r *ZoneFirewallRuleService) FirewallRulesUpdatePriorityOfFirewallRules(ctx context.Context, zoneIdentifier string, body ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesParams, opts ...option.RequestOption) (res *ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/rules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type ZoneFirewallRuleGetResponse struct {
	Errors   []ZoneFirewallRuleGetResponseError   `json:"errors"`
	Messages []ZoneFirewallRuleGetResponseMessage `json:"messages"`
	Result   ZoneFirewallRuleGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallRuleGetResponseSuccess `json:"success"`
	JSON    zoneFirewallRuleGetResponseJSON    `json:"-"`
}

// zoneFirewallRuleGetResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallRuleGetResponse]
type zoneFirewallRuleGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleGetResponseError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    zoneFirewallRuleGetResponseErrorJSON `json:"-"`
}

// zoneFirewallRuleGetResponseErrorJSON contains the JSON metadata for the struct
// [ZoneFirewallRuleGetResponseError]
type zoneFirewallRuleGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleGetResponseMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    zoneFirewallRuleGetResponseMessageJSON `json:"-"`
}

// zoneFirewallRuleGetResponseMessageJSON contains the JSON metadata for the struct
// [ZoneFirewallRuleGetResponseMessage]
type zoneFirewallRuleGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleGetResponseResult struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action ZoneFirewallRuleGetResponseResultAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                                  `json:"description"`
	Filter      ZoneFirewallRuleGetResponseResultFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                    `json:"priority"`
	Products []ZoneFirewallRuleGetResponseResultProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                                `json:"ref"`
	JSON zoneFirewallRuleGetResponseResultJSON `json:"-"`
}

// zoneFirewallRuleGetResponseResultJSON contains the JSON metadata for the struct
// [ZoneFirewallRuleGetResponseResult]
type zoneFirewallRuleGetResponseResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type ZoneFirewallRuleGetResponseResultAction string

const (
	ZoneFirewallRuleGetResponseResultActionBlock            ZoneFirewallRuleGetResponseResultAction = "block"
	ZoneFirewallRuleGetResponseResultActionChallenge        ZoneFirewallRuleGetResponseResultAction = "challenge"
	ZoneFirewallRuleGetResponseResultActionJsChallenge      ZoneFirewallRuleGetResponseResultAction = "js_challenge"
	ZoneFirewallRuleGetResponseResultActionManagedChallenge ZoneFirewallRuleGetResponseResultAction = "managed_challenge"
	ZoneFirewallRuleGetResponseResultActionAllow            ZoneFirewallRuleGetResponseResultAction = "allow"
	ZoneFirewallRuleGetResponseResultActionLog              ZoneFirewallRuleGetResponseResultAction = "log"
	ZoneFirewallRuleGetResponseResultActionBypass           ZoneFirewallRuleGetResponseResultAction = "bypass"
)

// Union satisfied by [ZoneFirewallRuleGetResponseResultFilterDZw70ubTFilter] or
// [ZoneFirewallRuleGetResponseResultFilterDZw70ubTDeletedFilter].
type ZoneFirewallRuleGetResponseResultFilter interface {
	implementsZoneFirewallRuleGetResponseResultFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallRuleGetResponseResultFilter)(nil)).Elem(), "")
}

type ZoneFirewallRuleGetResponseResultFilterDZw70ubTFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                                                    `json:"ref"`
	JSON zoneFirewallRuleGetResponseResultFilterDZw70ubTFilterJSON `json:"-"`
}

// zoneFirewallRuleGetResponseResultFilterDZw70ubTFilterJSON contains the JSON
// metadata for the struct [ZoneFirewallRuleGetResponseResultFilterDZw70ubTFilter]
type zoneFirewallRuleGetResponseResultFilterDZw70ubTFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleGetResponseResultFilterDZw70ubTFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallRuleGetResponseResultFilterDZw70ubTFilter) implementsZoneFirewallRuleGetResponseResultFilter() {
}

type ZoneFirewallRuleGetResponseResultFilterDZw70ubTDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                             `json:"deleted,required"`
	JSON    zoneFirewallRuleGetResponseResultFilterDZw70ubTDeletedFilterJSON `json:"-"`
}

// zoneFirewallRuleGetResponseResultFilterDZw70ubTDeletedFilterJSON contains the
// JSON metadata for the struct
// [ZoneFirewallRuleGetResponseResultFilterDZw70ubTDeletedFilter]
type zoneFirewallRuleGetResponseResultFilterDZw70ubTDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleGetResponseResultFilterDZw70ubTDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallRuleGetResponseResultFilterDZw70ubTDeletedFilter) implementsZoneFirewallRuleGetResponseResultFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type ZoneFirewallRuleGetResponseResultProduct string

const (
	ZoneFirewallRuleGetResponseResultProductZoneLockdown  ZoneFirewallRuleGetResponseResultProduct = "zoneLockdown"
	ZoneFirewallRuleGetResponseResultProductUaBlock       ZoneFirewallRuleGetResponseResultProduct = "uaBlock"
	ZoneFirewallRuleGetResponseResultProductBic           ZoneFirewallRuleGetResponseResultProduct = "bic"
	ZoneFirewallRuleGetResponseResultProductHot           ZoneFirewallRuleGetResponseResultProduct = "hot"
	ZoneFirewallRuleGetResponseResultProductSecurityLevel ZoneFirewallRuleGetResponseResultProduct = "securityLevel"
	ZoneFirewallRuleGetResponseResultProductRateLimit     ZoneFirewallRuleGetResponseResultProduct = "rateLimit"
	ZoneFirewallRuleGetResponseResultProductWaf           ZoneFirewallRuleGetResponseResultProduct = "waf"
)

// Whether the API call was successful
type ZoneFirewallRuleGetResponseSuccess bool

const (
	ZoneFirewallRuleGetResponseSuccessTrue ZoneFirewallRuleGetResponseSuccess = true
)

type ZoneFirewallRuleUpdateResponse struct {
	Errors   []ZoneFirewallRuleUpdateResponseError   `json:"errors"`
	Messages []ZoneFirewallRuleUpdateResponseMessage `json:"messages"`
	Result   ZoneFirewallRuleUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallRuleUpdateResponseSuccess `json:"success"`
	JSON    zoneFirewallRuleUpdateResponseJSON    `json:"-"`
}

// zoneFirewallRuleUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallRuleUpdateResponse]
type zoneFirewallRuleUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleUpdateResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneFirewallRuleUpdateResponseErrorJSON `json:"-"`
}

// zoneFirewallRuleUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneFirewallRuleUpdateResponseError]
type zoneFirewallRuleUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleUpdateResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneFirewallRuleUpdateResponseMessageJSON `json:"-"`
}

// zoneFirewallRuleUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneFirewallRuleUpdateResponseMessage]
type zoneFirewallRuleUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleUpdateResponseResult struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action ZoneFirewallRuleUpdateResponseResultAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                                     `json:"description"`
	Filter      ZoneFirewallRuleUpdateResponseResultFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                       `json:"priority"`
	Products []ZoneFirewallRuleUpdateResponseResultProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                                   `json:"ref"`
	JSON zoneFirewallRuleUpdateResponseResultJSON `json:"-"`
}

// zoneFirewallRuleUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneFirewallRuleUpdateResponseResult]
type zoneFirewallRuleUpdateResponseResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type ZoneFirewallRuleUpdateResponseResultAction string

const (
	ZoneFirewallRuleUpdateResponseResultActionBlock            ZoneFirewallRuleUpdateResponseResultAction = "block"
	ZoneFirewallRuleUpdateResponseResultActionChallenge        ZoneFirewallRuleUpdateResponseResultAction = "challenge"
	ZoneFirewallRuleUpdateResponseResultActionJsChallenge      ZoneFirewallRuleUpdateResponseResultAction = "js_challenge"
	ZoneFirewallRuleUpdateResponseResultActionManagedChallenge ZoneFirewallRuleUpdateResponseResultAction = "managed_challenge"
	ZoneFirewallRuleUpdateResponseResultActionAllow            ZoneFirewallRuleUpdateResponseResultAction = "allow"
	ZoneFirewallRuleUpdateResponseResultActionLog              ZoneFirewallRuleUpdateResponseResultAction = "log"
	ZoneFirewallRuleUpdateResponseResultActionBypass           ZoneFirewallRuleUpdateResponseResultAction = "bypass"
)

// Union satisfied by [ZoneFirewallRuleUpdateResponseResultFilterDZw70ubTFilter] or
// [ZoneFirewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilter].
type ZoneFirewallRuleUpdateResponseResultFilter interface {
	implementsZoneFirewallRuleUpdateResponseResultFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallRuleUpdateResponseResultFilter)(nil)).Elem(), "")
}

type ZoneFirewallRuleUpdateResponseResultFilterDZw70ubTFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                                                       `json:"ref"`
	JSON zoneFirewallRuleUpdateResponseResultFilterDZw70ubTFilterJSON `json:"-"`
}

// zoneFirewallRuleUpdateResponseResultFilterDZw70ubTFilterJSON contains the JSON
// metadata for the struct
// [ZoneFirewallRuleUpdateResponseResultFilterDZw70ubTFilter]
type zoneFirewallRuleUpdateResponseResultFilterDZw70ubTFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleUpdateResponseResultFilterDZw70ubTFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallRuleUpdateResponseResultFilterDZw70ubTFilter) implementsZoneFirewallRuleUpdateResponseResultFilter() {
}

type ZoneFirewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                                `json:"deleted,required"`
	JSON    zoneFirewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilterJSON `json:"-"`
}

// zoneFirewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilterJSON contains the
// JSON metadata for the struct
// [ZoneFirewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilter]
type zoneFirewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallRuleUpdateResponseResultFilterDZw70ubTDeletedFilter) implementsZoneFirewallRuleUpdateResponseResultFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type ZoneFirewallRuleUpdateResponseResultProduct string

const (
	ZoneFirewallRuleUpdateResponseResultProductZoneLockdown  ZoneFirewallRuleUpdateResponseResultProduct = "zoneLockdown"
	ZoneFirewallRuleUpdateResponseResultProductUaBlock       ZoneFirewallRuleUpdateResponseResultProduct = "uaBlock"
	ZoneFirewallRuleUpdateResponseResultProductBic           ZoneFirewallRuleUpdateResponseResultProduct = "bic"
	ZoneFirewallRuleUpdateResponseResultProductHot           ZoneFirewallRuleUpdateResponseResultProduct = "hot"
	ZoneFirewallRuleUpdateResponseResultProductSecurityLevel ZoneFirewallRuleUpdateResponseResultProduct = "securityLevel"
	ZoneFirewallRuleUpdateResponseResultProductRateLimit     ZoneFirewallRuleUpdateResponseResultProduct = "rateLimit"
	ZoneFirewallRuleUpdateResponseResultProductWaf           ZoneFirewallRuleUpdateResponseResultProduct = "waf"
)

// Whether the API call was successful
type ZoneFirewallRuleUpdateResponseSuccess bool

const (
	ZoneFirewallRuleUpdateResponseSuccessTrue ZoneFirewallRuleUpdateResponseSuccess = true
)

type ZoneFirewallRuleDeleteResponse struct {
	Errors   []ZoneFirewallRuleDeleteResponseError   `json:"errors"`
	Messages []ZoneFirewallRuleDeleteResponseMessage `json:"messages"`
	Result   ZoneFirewallRuleDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallRuleDeleteResponseSuccess `json:"success"`
	JSON    zoneFirewallRuleDeleteResponseJSON    `json:"-"`
}

// zoneFirewallRuleDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallRuleDeleteResponse]
type zoneFirewallRuleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleDeleteResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zoneFirewallRuleDeleteResponseErrorJSON `json:"-"`
}

// zoneFirewallRuleDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneFirewallRuleDeleteResponseError]
type zoneFirewallRuleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleDeleteResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zoneFirewallRuleDeleteResponseMessageJSON `json:"-"`
}

// zoneFirewallRuleDeleteResponseMessageJSON contains the JSON metadata for the
// struct [ZoneFirewallRuleDeleteResponseMessage]
type zoneFirewallRuleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleDeleteResponseResult struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action ZoneFirewallRuleDeleteResponseResultAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                                     `json:"description"`
	Filter      ZoneFirewallRuleDeleteResponseResultFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                       `json:"priority"`
	Products []ZoneFirewallRuleDeleteResponseResultProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                                   `json:"ref"`
	JSON zoneFirewallRuleDeleteResponseResultJSON `json:"-"`
}

// zoneFirewallRuleDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneFirewallRuleDeleteResponseResult]
type zoneFirewallRuleDeleteResponseResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type ZoneFirewallRuleDeleteResponseResultAction string

const (
	ZoneFirewallRuleDeleteResponseResultActionBlock            ZoneFirewallRuleDeleteResponseResultAction = "block"
	ZoneFirewallRuleDeleteResponseResultActionChallenge        ZoneFirewallRuleDeleteResponseResultAction = "challenge"
	ZoneFirewallRuleDeleteResponseResultActionJsChallenge      ZoneFirewallRuleDeleteResponseResultAction = "js_challenge"
	ZoneFirewallRuleDeleteResponseResultActionManagedChallenge ZoneFirewallRuleDeleteResponseResultAction = "managed_challenge"
	ZoneFirewallRuleDeleteResponseResultActionAllow            ZoneFirewallRuleDeleteResponseResultAction = "allow"
	ZoneFirewallRuleDeleteResponseResultActionLog              ZoneFirewallRuleDeleteResponseResultAction = "log"
	ZoneFirewallRuleDeleteResponseResultActionBypass           ZoneFirewallRuleDeleteResponseResultAction = "bypass"
)

// Union satisfied by [ZoneFirewallRuleDeleteResponseResultFilterDZw70ubTFilter] or
// [ZoneFirewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilter].
type ZoneFirewallRuleDeleteResponseResultFilter interface {
	implementsZoneFirewallRuleDeleteResponseResultFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallRuleDeleteResponseResultFilter)(nil)).Elem(), "")
}

type ZoneFirewallRuleDeleteResponseResultFilterDZw70ubTFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                                                       `json:"ref"`
	JSON zoneFirewallRuleDeleteResponseResultFilterDZw70ubTFilterJSON `json:"-"`
}

// zoneFirewallRuleDeleteResponseResultFilterDZw70ubTFilterJSON contains the JSON
// metadata for the struct
// [ZoneFirewallRuleDeleteResponseResultFilterDZw70ubTFilter]
type zoneFirewallRuleDeleteResponseResultFilterDZw70ubTFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleDeleteResponseResultFilterDZw70ubTFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallRuleDeleteResponseResultFilterDZw70ubTFilter) implementsZoneFirewallRuleDeleteResponseResultFilter() {
}

type ZoneFirewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                                `json:"deleted,required"`
	JSON    zoneFirewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilterJSON `json:"-"`
}

// zoneFirewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilterJSON contains the
// JSON metadata for the struct
// [ZoneFirewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilter]
type zoneFirewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallRuleDeleteResponseResultFilterDZw70ubTDeletedFilter) implementsZoneFirewallRuleDeleteResponseResultFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type ZoneFirewallRuleDeleteResponseResultProduct string

const (
	ZoneFirewallRuleDeleteResponseResultProductZoneLockdown  ZoneFirewallRuleDeleteResponseResultProduct = "zoneLockdown"
	ZoneFirewallRuleDeleteResponseResultProductUaBlock       ZoneFirewallRuleDeleteResponseResultProduct = "uaBlock"
	ZoneFirewallRuleDeleteResponseResultProductBic           ZoneFirewallRuleDeleteResponseResultProduct = "bic"
	ZoneFirewallRuleDeleteResponseResultProductHot           ZoneFirewallRuleDeleteResponseResultProduct = "hot"
	ZoneFirewallRuleDeleteResponseResultProductSecurityLevel ZoneFirewallRuleDeleteResponseResultProduct = "securityLevel"
	ZoneFirewallRuleDeleteResponseResultProductRateLimit     ZoneFirewallRuleDeleteResponseResultProduct = "rateLimit"
	ZoneFirewallRuleDeleteResponseResultProductWaf           ZoneFirewallRuleDeleteResponseResultProduct = "waf"
)

// Whether the API call was successful
type ZoneFirewallRuleDeleteResponseSuccess bool

const (
	ZoneFirewallRuleDeleteResponseSuccessTrue ZoneFirewallRuleDeleteResponseSuccess = true
)

type ZoneFirewallRuleFirewallRulesNewFirewallRulesResponse struct {
	Errors     []ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseError    `json:"errors"`
	Messages   []ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseMessage  `json:"messages"`
	Result     []ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResult   `json:"result"`
	ResultInfo ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseSuccess `json:"success"`
	JSON    zoneFirewallRuleFirewallRulesNewFirewallRulesResponseJSON    `json:"-"`
}

// zoneFirewallRuleFirewallRulesNewFirewallRulesResponseJSON contains the JSON
// metadata for the struct [ZoneFirewallRuleFirewallRulesNewFirewallRulesResponse]
type zoneFirewallRuleFirewallRulesNewFirewallRulesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesNewFirewallRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseError struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zoneFirewallRuleFirewallRulesNewFirewallRulesResponseErrorJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesNewFirewallRulesResponseErrorJSON contains the JSON
// metadata for the struct
// [ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseError]
type zoneFirewallRuleFirewallRulesNewFirewallRulesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseMessage struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    zoneFirewallRuleFirewallRulesNewFirewallRulesResponseMessageJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesNewFirewallRulesResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseMessage]
type zoneFirewallRuleFirewallRulesNewFirewallRulesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResult struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                                                            `json:"description"`
	Filter      ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                                              `json:"priority"`
	Products []ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                                                          `json:"ref"`
	JSON zoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResult]
type zoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultAction string

const (
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultActionBlock            ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultAction = "block"
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultActionChallenge        ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultAction = "challenge"
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultActionJsChallenge      ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultAction = "js_challenge"
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultActionManagedChallenge ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultAction = "managed_challenge"
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultActionAllow            ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultAction = "allow"
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultActionLog              ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultAction = "log"
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultActionBypass           ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultAction = "bypass"
)

// Union satisfied by
// [ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTFilter]
// or
// [ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTDeletedFilter].
type ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilter interface {
	implementsZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilter)(nil)).Elem(), "")
}

type ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                                                                              `json:"ref"`
	JSON zoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTFilterJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTFilterJSON
// contains the JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTFilter]
type zoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTFilter) implementsZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilter() {
}

type ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                                                       `json:"deleted,required"`
	JSON    zoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTDeletedFilterJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTDeletedFilterJSON
// contains the JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTDeletedFilter]
type zoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilterDZw70ubTDeletedFilter) implementsZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProduct string

const (
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProductZoneLockdown  ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProduct = "zoneLockdown"
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProductUaBlock       ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProduct = "uaBlock"
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProductBic           ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProduct = "bic"
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProductHot           ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProduct = "hot"
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProductSecurityLevel ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProduct = "securityLevel"
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProductRateLimit     ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProduct = "rateLimit"
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProductWaf           ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultProduct = "waf"
)

type ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                             `json:"total_count"`
	JSON       zoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultInfoJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultInfoJSON contains the
// JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultInfo]
type zoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseSuccess bool

const (
	ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseSuccessTrue ZoneFirewallRuleFirewallRulesNewFirewallRulesResponseSuccess = true
)

type ZoneFirewallRuleFirewallRulesListFirewallRulesResponse struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action ZoneFirewallRuleFirewallRulesListFirewallRulesResponseAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                                                       `json:"description"`
	Filter      ZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                                         `json:"priority"`
	Products []ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                                                     `json:"ref"`
	JSON zoneFirewallRuleFirewallRulesListFirewallRulesResponseJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesListFirewallRulesResponseJSON contains the JSON
// metadata for the struct [ZoneFirewallRuleFirewallRulesListFirewallRulesResponse]
type zoneFirewallRuleFirewallRulesListFirewallRulesResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesListFirewallRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type ZoneFirewallRuleFirewallRulesListFirewallRulesResponseAction string

const (
	ZoneFirewallRuleFirewallRulesListFirewallRulesResponseActionBlock            ZoneFirewallRuleFirewallRulesListFirewallRulesResponseAction = "block"
	ZoneFirewallRuleFirewallRulesListFirewallRulesResponseActionChallenge        ZoneFirewallRuleFirewallRulesListFirewallRulesResponseAction = "challenge"
	ZoneFirewallRuleFirewallRulesListFirewallRulesResponseActionJsChallenge      ZoneFirewallRuleFirewallRulesListFirewallRulesResponseAction = "js_challenge"
	ZoneFirewallRuleFirewallRulesListFirewallRulesResponseActionManagedChallenge ZoneFirewallRuleFirewallRulesListFirewallRulesResponseAction = "managed_challenge"
	ZoneFirewallRuleFirewallRulesListFirewallRulesResponseActionAllow            ZoneFirewallRuleFirewallRulesListFirewallRulesResponseAction = "allow"
	ZoneFirewallRuleFirewallRulesListFirewallRulesResponseActionLog              ZoneFirewallRuleFirewallRulesListFirewallRulesResponseAction = "log"
	ZoneFirewallRuleFirewallRulesListFirewallRulesResponseActionBypass           ZoneFirewallRuleFirewallRulesListFirewallRulesResponseAction = "bypass"
)

// Union satisfied by
// [ZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTFilter] or
// [ZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTDeletedFilter].
type ZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilter interface {
	implementsZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilter)(nil)).Elem(), "")
}

type ZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                                                                         `json:"ref"`
	JSON zoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTFilterJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTFilterJSON
// contains the JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTFilter]
type zoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTFilter) implementsZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilter() {
}

type ZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                                                  `json:"deleted,required"`
	JSON    zoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTDeletedFilterJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTDeletedFilterJSON
// contains the JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTDeletedFilter]
type zoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilterDZw70ubTDeletedFilter) implementsZoneFirewallRuleFirewallRulesListFirewallRulesResponseFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProduct string

const (
	ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProductZoneLockdown  ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProduct = "zoneLockdown"
	ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProductUaBlock       ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProduct = "uaBlock"
	ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProductBic           ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProduct = "bic"
	ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProductHot           ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProduct = "hot"
	ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProductSecurityLevel ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProduct = "securityLevel"
	ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProductRateLimit     ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProduct = "rateLimit"
	ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProductWaf           ZoneFirewallRuleFirewallRulesListFirewallRulesResponseProduct = "waf"
)

type ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponse struct {
	Errors     []ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseError    `json:"errors"`
	Messages   []ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseMessage  `json:"messages"`
	Result     []ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResult   `json:"result"`
	ResultInfo ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseSuccess `json:"success"`
	JSON    zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseJSON    `json:"-"`
}

// zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseJSON contains the JSON
// metadata for the struct
// [ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponse]
type zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseErrorJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseErrorJSON contains the
// JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseError]
type zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseMessageJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseMessageJSON contains the
// JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseMessage]
type zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResult struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                                                               `json:"description"`
	Filter      ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                                                 `json:"priority"`
	Products []ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                                                             `json:"ref"`
	JSON zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResult]
type zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultAction string

const (
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultActionBlock            ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultAction = "block"
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultActionChallenge        ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultAction = "challenge"
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultActionJsChallenge      ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultAction = "js_challenge"
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultActionManagedChallenge ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultAction = "managed_challenge"
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultActionAllow            ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultAction = "allow"
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultActionLog              ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultAction = "log"
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultActionBypass           ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultAction = "bypass"
)

// Union satisfied by
// [ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTFilter]
// or
// [ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTDeletedFilter].
type ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilter interface {
	implementsZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilter)(nil)).Elem(), "")
}

type ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                                                                                 `json:"ref"`
	JSON zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTFilterJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTFilterJSON
// contains the JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTFilter]
type zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTFilter) implementsZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilter() {
}

type ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                                                          `json:"deleted,required"`
	JSON    zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTDeletedFilterJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTDeletedFilterJSON
// contains the JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTDeletedFilter]
type zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilterDZw70ubTDeletedFilter) implementsZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProduct string

const (
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProductZoneLockdown  ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProduct = "zoneLockdown"
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProductUaBlock       ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProduct = "uaBlock"
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProductBic           ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProduct = "bic"
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProductHot           ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProduct = "hot"
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProductSecurityLevel ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProduct = "securityLevel"
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProductRateLimit     ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProduct = "rateLimit"
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProductWaf           ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultProduct = "waf"
)

type ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                `json:"total_count"`
	JSON       zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultInfoJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultInfoJSON contains
// the JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultInfo]
type zoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseSuccess bool

const (
	ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseSuccessTrue ZoneFirewallRuleFirewallRulesUpdateFirewallRulesResponseSuccess = true
)

type ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponse struct {
	Errors     []ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseError    `json:"errors"`
	Messages   []ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseMessage  `json:"messages"`
	Result     []ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResult   `json:"result"`
	ResultInfo ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseSuccess `json:"success"`
	JSON    zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseJSON    `json:"-"`
}

// zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseJSON contains
// the JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponse]
type zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseError struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseErrorJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseError]
type zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseMessage struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseMessageJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseMessage]
type zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResult struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultAction `json:"action"`
	// An informative summary of the firewall rule.
	Description string                                                                         `json:"description"`
	Filter      ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64                                                                           `json:"priority"`
	Products []ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProduct `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string                                                                       `json:"ref"`
	JSON zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResult]
type zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultAction string

const (
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultActionBlock            ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultAction = "block"
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultActionChallenge        ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultAction = "challenge"
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultActionJsChallenge      ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultAction = "js_challenge"
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultActionManagedChallenge ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultAction = "managed_challenge"
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultActionAllow            ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultAction = "allow"
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultActionLog              ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultAction = "log"
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultActionBypass           ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultAction = "bypass"
)

// Union satisfied by
// [ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTFilter]
// or
// [ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTDeletedFilter].
type ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilter interface {
	implementsZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilter()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilter)(nil)).Elem(), "")
}

type ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref  string                                                                                           `json:"ref"`
	JSON zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTFilterJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTFilterJSON
// contains the JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTFilter]
type zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTFilterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTFilter) implementsZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilter() {
}

type ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTDeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool                                                                                                    `json:"deleted,required"`
	JSON    zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTDeletedFilterJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTDeletedFilterJSON
// contains the JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTDeletedFilter]
type zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTDeletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTDeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilterDZw70ubTDeletedFilter) implementsZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultFilter() {
}

// A list of products to bypass for a request when using the `bypass` action.
type ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProduct string

const (
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProductZoneLockdown  ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProduct = "zoneLockdown"
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProductUaBlock       ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProduct = "uaBlock"
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProductBic           ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProduct = "bic"
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProductHot           ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProduct = "hot"
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProductSecurityLevel ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProduct = "securityLevel"
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProductRateLimit     ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProduct = "rateLimit"
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProductWaf           ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultProduct = "waf"
)

type ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                          `json:"total_count"`
	JSON       zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultInfoJSON `json:"-"`
}

// zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultInfo]
type zoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseSuccess bool

const (
	ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseSuccessTrue ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesResponseSuccess = true
)

type ZoneFirewallRuleGetParams struct {
}

type ZoneFirewallRuleUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFirewallRuleDeleteParams struct {
	// When true, indicates that Cloudflare should also delete the associated filter if
	// there are no other firewall rules referencing the filter.
	DeleteFilterIfUnused param.Field[bool] `json:"delete_filter_if_unused"`
}

func (r ZoneFirewallRuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneFirewallRuleFirewallRulesNewFirewallRulesParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallRuleFirewallRulesNewFirewallRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFirewallRuleFirewallRulesListFirewallRulesParams struct {
	// The action to search for. Must be an exact match.
	Action param.Field[string] `query:"action"`
	// A case-insensitive string to find in the description.
	Description param.Field[string] `query:"description"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// When true, indicates that the firewall rule is currently paused.
	Paused param.Field[bool] `query:"paused"`
	// Number of firewall rules per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [ZoneFirewallRuleFirewallRulesListFirewallRulesParams]'s
// query parameters as `url.Values`.
func (r ZoneFirewallRuleFirewallRulesListFirewallRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneFirewallRuleFirewallRulesUpdateFirewallRulesParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallRuleFirewallRulesUpdateFirewallRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallRuleFirewallRulesUpdatePriorityOfFirewallRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

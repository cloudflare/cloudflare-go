// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// LockdownService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewLockdownService] method instead.
type LockdownService struct {
	Options []option.RequestOption
}

// NewLockdownService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLockdownService(opts ...option.RequestOption) (r *LockdownService) {
	r = &LockdownService{}
	r.Options = opts
	return
}

// Creates a new Zone Lockdown rule.
func (r *LockdownService) New(ctx context.Context, zoneIdentifier string, body LockdownNewParams, opts ...option.RequestOption) (res *FirewallZoneLockdown, err error) {
	opts = append(r.Options[:], opts...)
	var env LockdownNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Zone Lockdown rule.
func (r *LockdownService) Update(ctx context.Context, zoneIdentifier string, id string, body LockdownUpdateParams, opts ...option.RequestOption) (res *FirewallZoneLockdown, err error) {
	opts = append(r.Options[:], opts...)
	var env LockdownUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Zone Lockdown rules. You can filter the results using several optional
// parameters.
func (r *LockdownService) List(ctx context.Context, zoneIdentifier string, query LockdownListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[FirewallZoneLockdown], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneIdentifier)
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

// Fetches Zone Lockdown rules. You can filter the results using several optional
// parameters.
func (r *LockdownService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query LockdownListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[FirewallZoneLockdown] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing Zone Lockdown rule.
func (r *LockdownService) Delete(ctx context.Context, zoneIdentifier string, id string, body LockdownDeleteParams, opts ...option.RequestOption) (res *LockdownDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LockdownDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a Zone Lockdown rule.
func (r *LockdownService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FirewallZoneLockdown, err error) {
	opts = append(r.Options[:], opts...)
	var env LockdownGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallZoneLockdown struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations FirewallZoneLockdownConfigurations `json:"configurations,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// An informative summary of the rule.
	Description string `json:"description,required"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// When true, indicates that the rule is currently paused.
	Paused bool `json:"paused,required"`
	// The URLs to include in the rule definition. You can use wildcards. Each entered
	// URL will be escaped before use, which means you can only use simple wildcard
	// patterns.
	URLs []string                 `json:"urls,required"`
	JSON firewallZoneLockdownJSON `json:"-"`
}

// firewallZoneLockdownJSON contains the JSON metadata for the struct
// [FirewallZoneLockdown]
type firewallZoneLockdownJSON struct {
	ID             apijson.Field
	Configurations apijson.Field
	CreatedOn      apijson.Field
	Description    apijson.Field
	ModifiedOn     apijson.Field
	Paused         apijson.Field
	URLs           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *FirewallZoneLockdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallZoneLockdownJSON) RawJSON() string {
	return r.raw
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
type FirewallZoneLockdownConfigurations struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target FirewallZoneLockdownConfigurationsTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                 `json:"value"`
	JSON  firewallZoneLockdownConfigurationsJSON `json:"-"`
	union FirewallZoneLockdownConfigurationsUnion
}

// firewallZoneLockdownConfigurationsJSON contains the JSON metadata for the struct
// [FirewallZoneLockdownConfigurations]
type firewallZoneLockdownConfigurationsJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r firewallZoneLockdownConfigurationsJSON) RawJSON() string {
	return r.raw
}

func (r *FirewallZoneLockdownConfigurations) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r FirewallZoneLockdownConfigurations) AsUnion() FirewallZoneLockdownConfigurationsUnion {
	return r.union
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [firewall.FirewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfiguration] or
// [firewall.FirewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfiguration].
type FirewallZoneLockdownConfigurationsUnion interface {
	implementsFirewallFirewallZoneLockdownConfigurations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallZoneLockdownConfigurationsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(FirewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfiguration{}),
		},
	)
}

type FirewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target FirewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                `json:"value"`
	JSON  firewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfigurationJSON `json:"-"`
}

// firewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfigurationJSON contains
// the JSON metadata for the struct
// [FirewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfiguration]
type firewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfiguration) implementsFirewallFirewallZoneLockdownConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type FirewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfigurationTarget string

const (
	FirewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfigurationTargetIP FirewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfigurationTarget = "ip"
)

func (r FirewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallZoneLockdownConfigurationsLegacyJhsSchemasIPConfigurationTargetIP:
		return true
	}
	return false
}

type FirewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target FirewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                  `json:"value"`
	JSON  firewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfigurationJSON `json:"-"`
}

// firewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfigurationJSON contains
// the JSON metadata for the struct
// [FirewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfiguration]
type firewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r FirewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfiguration) implementsFirewallFirewallZoneLockdownConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type FirewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfigurationTarget string

const (
	FirewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfigurationTargetIPRange FirewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfigurationTarget = "ip_range"
)

func (r FirewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case FirewallZoneLockdownConfigurationsLegacyJhsSchemasCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type FirewallZoneLockdownConfigurationsTarget string

const (
	FirewallZoneLockdownConfigurationsTargetIP      FirewallZoneLockdownConfigurationsTarget = "ip"
	FirewallZoneLockdownConfigurationsTargetIPRange FirewallZoneLockdownConfigurationsTarget = "ip_range"
)

func (r FirewallZoneLockdownConfigurationsTarget) IsKnown() bool {
	switch r {
	case FirewallZoneLockdownConfigurationsTargetIP, FirewallZoneLockdownConfigurationsTargetIPRange:
		return true
	}
	return false
}

type LockdownDeleteResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID   string                     `json:"id"`
	JSON lockdownDeleteResponseJSON `json:"-"`
}

// lockdownDeleteResponseJSON contains the JSON metadata for the struct
// [LockdownDeleteResponse]
type lockdownDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type LockdownNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r LockdownNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LockdownNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   FirewallZoneLockdown                                      `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LockdownNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    lockdownNewResponseEnvelopeJSON    `json:"-"`
}

// lockdownNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownNewResponseEnvelope]
type lockdownNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LockdownNewResponseEnvelopeSuccess bool

const (
	LockdownNewResponseEnvelopeSuccessTrue LockdownNewResponseEnvelopeSuccess = true
)

func (r LockdownNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LockdownNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LockdownUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r LockdownUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LockdownUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   FirewallZoneLockdown                                      `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LockdownUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    lockdownUpdateResponseEnvelopeJSON    `json:"-"`
}

// lockdownUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownUpdateResponseEnvelope]
type lockdownUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LockdownUpdateResponseEnvelopeSuccess bool

const (
	LockdownUpdateResponseEnvelopeSuccessTrue LockdownUpdateResponseEnvelopeSuccess = true
)

func (r LockdownUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LockdownUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LockdownListParams struct {
	// The timestamp of when the rule was created.
	CreatedOn param.Field[time.Time] `query:"created_on" format:"date-time"`
	// A string to search for in the description of existing rules.
	Description param.Field[string] `query:"description"`
	// A string to search for in the description of existing rules.
	DescriptionSearch param.Field[string] `query:"description_search"`
	// A single IP address to search for in existing rules.
	IP param.Field[string] `query:"ip"`
	// A single IP address range to search for in existing rules.
	IPRangeSearch param.Field[string] `query:"ip_range_search"`
	// A single IP address to search for in existing rules.
	IPSearch param.Field[string] `query:"ip_search"`
	// The timestamp of when the rule was last modified.
	ModifiedOn param.Field[time.Time] `query:"modified_on" format:"date-time"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The maximum number of results per page. You can only set the value to `1` or to
	// a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
	// The priority of the rule to control the processing order. A lower number
	// indicates higher priority. If not provided, any rules with a configured priority
	// will be processed before rules without a priority.
	Priority param.Field[float64] `query:"priority"`
	// A single URI to search for in the list of URLs of existing rules.
	URISearch param.Field[string] `query:"uri_search"`
}

// URLQuery serializes [LockdownListParams]'s query parameters as `url.Values`.
func (r LockdownListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LockdownDeleteParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r LockdownDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type LockdownDeleteResponseEnvelope struct {
	Result LockdownDeleteResponse             `json:"result"`
	JSON   lockdownDeleteResponseEnvelopeJSON `json:"-"`
}

// lockdownDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownDeleteResponseEnvelope]
type lockdownDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LockdownGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   FirewallZoneLockdown                                      `json:"result,required,nullable"`
	// Whether the API call was successful
	Success LockdownGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    lockdownGetResponseEnvelopeJSON    `json:"-"`
}

// lockdownGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [LockdownGetResponseEnvelope]
type lockdownGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LockdownGetResponseEnvelopeSuccess bool

const (
	LockdownGetResponseEnvelopeSuccessTrue LockdownGetResponseEnvelopeSuccess = true
)

func (r LockdownGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LockdownGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

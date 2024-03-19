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
func (r *LockdownService) New(ctx context.Context, zoneIdentifier string, body LockdownNewParams, opts ...option.RequestOption) (res *LegacyJhsZonelockdown, err error) {
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
func (r *LockdownService) Update(ctx context.Context, zoneIdentifier string, id string, body LockdownUpdateParams, opts ...option.RequestOption) (res *LegacyJhsZonelockdown, err error) {
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
func (r *LockdownService) List(ctx context.Context, zoneIdentifier string, query LockdownListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[LegacyJhsZonelockdown], err error) {
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
func (r *LockdownService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query LockdownListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[LegacyJhsZonelockdown] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing Zone Lockdown rule.
func (r *LockdownService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *LockdownDeleteResponse, err error) {
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
func (r *LockdownService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *LegacyJhsZonelockdown, err error) {
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

type LegacyJhsZonelockdown struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations LegacyJhsZonelockdownConfigurations `json:"configurations,required"`
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
	URLs []string                  `json:"urls,required"`
	JSON legacyJhsZonelockdownJSON `json:"-"`
}

// legacyJhsZonelockdownJSON contains the JSON metadata for the struct
// [LegacyJhsZonelockdown]
type legacyJhsZonelockdownJSON struct {
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

func (r *LegacyJhsZonelockdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legacyJhsZonelockdownJSON) RawJSON() string {
	return r.raw
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [firewall.LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfiguration] or
// [firewall.LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfiguration].
type LegacyJhsZonelockdownConfigurations interface {
	implementsFirewallLegacyJhsZonelockdownConfigurations()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LegacyJhsZonelockdownConfigurations)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfiguration{}),
		},
	)
}

type LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                 `json:"value"`
	JSON  legacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfigurationJSON `json:"-"`
}

// legacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfigurationJSON contains
// the JSON metadata for the struct
// [LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfiguration]
type legacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfiguration) implementsFirewallLegacyJhsZonelockdownConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfigurationTarget string

const (
	LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfigurationTargetIP LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfigurationTarget = "ip"
)

func (r LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfigurationTarget) IsKnown() bool {
	switch r {
	case LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfigurationTargetIP:
		return true
	}
	return false
}

type LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                   `json:"value"`
	JSON  legacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfigurationJSON `json:"-"`
}

// legacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfigurationJSON
// contains the JSON metadata for the struct
// [LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfiguration]
type legacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r legacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfiguration) implementsFirewallLegacyJhsZonelockdownConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfigurationTarget string

const (
	LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfigurationTargetIPRange LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfigurationTarget = "ip_range"
)

func (r LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCIDRConfigurationTargetIPRange:
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
	Errors   []LockdownNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LockdownNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsZonelockdown                 `json:"result,required,nullable"`
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

type LockdownNewResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    lockdownNewResponseEnvelopeErrorsJSON `json:"-"`
}

// lockdownNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [LockdownNewResponseEnvelopeErrors]
type lockdownNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LockdownNewResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    lockdownNewResponseEnvelopeMessagesJSON `json:"-"`
}

// lockdownNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LockdownNewResponseEnvelopeMessages]
type lockdownNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownNewResponseEnvelopeMessagesJSON) RawJSON() string {
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
	Errors   []LockdownUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LockdownUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsZonelockdown                    `json:"result,required,nullable"`
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

type LockdownUpdateResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    lockdownUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// lockdownUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [LockdownUpdateResponseEnvelopeErrors]
type lockdownUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LockdownUpdateResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    lockdownUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// lockdownUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LockdownUpdateResponseEnvelopeMessages]
type lockdownUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
	Errors   []LockdownGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LockdownGetResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsZonelockdown                 `json:"result,required,nullable"`
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

type LockdownGetResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    lockdownGetResponseEnvelopeErrorsJSON `json:"-"`
}

// lockdownGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [LockdownGetResponseEnvelopeErrors]
type lockdownGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LockdownGetResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    lockdownGetResponseEnvelopeMessagesJSON `json:"-"`
}

// lockdownGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [LockdownGetResponseEnvelopeMessages]
type lockdownGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LockdownGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lockdownGetResponseEnvelopeMessagesJSON) RawJSON() string {
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

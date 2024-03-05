// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// FirewallLockdownService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallLockdownService] method
// instead.
type FirewallLockdownService struct {
	Options []option.RequestOption
}

// NewFirewallLockdownService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFirewallLockdownService(opts ...option.RequestOption) (r *FirewallLockdownService) {
	r = &FirewallLockdownService{}
	r.Options = opts
	return
}

// Creates a new Zone Lockdown rule.
func (r *FirewallLockdownService) New(ctx context.Context, zoneIdentifier string, body FirewallLockdownNewParams, opts ...option.RequestOption) (res *LegacyJhsZonelockdown, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallLockdownNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Zone Lockdown rule.
func (r *FirewallLockdownService) Update(ctx context.Context, zoneIdentifier string, id string, body FirewallLockdownUpdateParams, opts ...option.RequestOption) (res *LegacyJhsZonelockdown, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallLockdownUpdateResponseEnvelope
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
func (r *FirewallLockdownService) List(ctx context.Context, zoneIdentifier string, query FirewallLockdownListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[LegacyJhsZonelockdown], err error) {
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
func (r *FirewallLockdownService) ListAutoPaging(ctx context.Context, zoneIdentifier string, query FirewallLockdownListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[LegacyJhsZonelockdown] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, zoneIdentifier, query, opts...))
}

// Deletes an existing Zone Lockdown rule.
func (r *FirewallLockdownService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FirewallLockdownDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallLockdownDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a Zone Lockdown rule.
func (r *FirewallLockdownService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *LegacyJhsZonelockdown, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallLockdownGetResponseEnvelope
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

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfiguration] or
// [LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCidrConfiguration].
type LegacyJhsZonelockdownConfigurations interface {
	implementsLegacyJhsZonelockdownConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*LegacyJhsZonelockdownConfigurations)(nil)).Elem(), "")
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

func (r LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfiguration) implementsLegacyJhsZonelockdownConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfigurationTarget string

const (
	LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfigurationTargetIP LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasIPConfigurationTarget = "ip"
)

type LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                   `json:"value"`
	JSON  legacyJhsZonelockdownConfigurationsLegacyJhsSchemasCidrConfigurationJSON `json:"-"`
}

// legacyJhsZonelockdownConfigurationsLegacyJhsSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCidrConfiguration]
type legacyJhsZonelockdownConfigurationsLegacyJhsSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCidrConfiguration) implementsLegacyJhsZonelockdownConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCidrConfigurationTarget string

const (
	LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCidrConfigurationTargetIPRange LegacyJhsZonelockdownConfigurationsLegacyJhsSchemasCidrConfigurationTarget = "ip_range"
)

type FirewallLockdownDeleteResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID   string                             `json:"id"`
	JSON firewallLockdownDeleteResponseJSON `json:"-"`
}

// firewallLockdownDeleteResponseJSON contains the JSON metadata for the struct
// [FirewallLockdownDeleteResponse]
type firewallLockdownDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownNewParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallLockdownNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallLockdownNewResponseEnvelope struct {
	Errors   []FirewallLockdownNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallLockdownNewResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsZonelockdown                         `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallLockdownNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallLockdownNewResponseEnvelopeJSON    `json:"-"`
}

// firewallLockdownNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallLockdownNewResponseEnvelope]
type firewallLockdownNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownNewResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    firewallLockdownNewResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallLockdownNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallLockdownNewResponseEnvelopeErrors]
type firewallLockdownNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownNewResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    firewallLockdownNewResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallLockdownNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallLockdownNewResponseEnvelopeMessages]
type firewallLockdownNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallLockdownNewResponseEnvelopeSuccess bool

const (
	FirewallLockdownNewResponseEnvelopeSuccessTrue FirewallLockdownNewResponseEnvelopeSuccess = true
)

type FirewallLockdownUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallLockdownUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallLockdownUpdateResponseEnvelope struct {
	Errors   []FirewallLockdownUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallLockdownUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsZonelockdown                            `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallLockdownUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallLockdownUpdateResponseEnvelopeJSON    `json:"-"`
}

// firewallLockdownUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallLockdownUpdateResponseEnvelope]
type firewallLockdownUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownUpdateResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    firewallLockdownUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallLockdownUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [FirewallLockdownUpdateResponseEnvelopeErrors]
type firewallLockdownUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownUpdateResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    firewallLockdownUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallLockdownUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [FirewallLockdownUpdateResponseEnvelopeMessages]
type firewallLockdownUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallLockdownUpdateResponseEnvelopeSuccess bool

const (
	FirewallLockdownUpdateResponseEnvelopeSuccessTrue FirewallLockdownUpdateResponseEnvelopeSuccess = true
)

type FirewallLockdownListParams struct {
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
	UriSearch param.Field[string] `query:"uri_search"`
}

// URLQuery serializes [FirewallLockdownListParams]'s query parameters as
// `url.Values`.
func (r FirewallLockdownListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallLockdownDeleteResponseEnvelope struct {
	Result FirewallLockdownDeleteResponse             `json:"result"`
	JSON   firewallLockdownDeleteResponseEnvelopeJSON `json:"-"`
}

// firewallLockdownDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallLockdownDeleteResponseEnvelope]
type firewallLockdownDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownGetResponseEnvelope struct {
	Errors   []FirewallLockdownGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallLockdownGetResponseEnvelopeMessages `json:"messages,required"`
	Result   LegacyJhsZonelockdown                         `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallLockdownGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallLockdownGetResponseEnvelopeJSON    `json:"-"`
}

// firewallLockdownGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [FirewallLockdownGetResponseEnvelope]
type firewallLockdownGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownGetResponseEnvelopeErrors struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    firewallLockdownGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallLockdownGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [FirewallLockdownGetResponseEnvelopeErrors]
type firewallLockdownGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownGetResponseEnvelopeMessages struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    firewallLockdownGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallLockdownGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [FirewallLockdownGetResponseEnvelopeMessages]
type firewallLockdownGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallLockdownGetResponseEnvelopeSuccess bool

const (
	FirewallLockdownGetResponseEnvelopeSuccessTrue FirewallLockdownGetResponseEnvelopeSuccess = true
)

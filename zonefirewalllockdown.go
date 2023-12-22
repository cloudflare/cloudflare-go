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
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneFirewallLockdownService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneFirewallLockdownService]
// method instead.
type ZoneFirewallLockdownService struct {
	Options []option.RequestOption
}

// NewZoneFirewallLockdownService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneFirewallLockdownService(opts ...option.RequestOption) (r *ZoneFirewallLockdownService) {
	r = &ZoneFirewallLockdownService{}
	r.Options = opts
	return
}

// Fetches the details of a Zone Lockdown rule.
func (r *ZoneFirewallLockdownService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *ZonelockdownResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing Zone Lockdown rule.
func (r *ZoneFirewallLockdownService) Update(ctx context.Context, zoneIdentifier string, id string, body ZoneFirewallLockdownUpdateParams, opts ...option.RequestOption) (res *ZonelockdownResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an existing Zone Lockdown rule.
func (r *ZoneFirewallLockdownService) Delete(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *ZoneFirewallLockdownDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new Zone Lockdown rule.
func (r *ZoneFirewallLockdownService) ZoneLockdownNewAZoneLockdownRule(ctx context.Context, zoneIdentifier string, body ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleParams, opts ...option.RequestOption) (res *ZonelockdownResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches Zone Lockdown rules. You can filter the results using several optional
// parameters.
func (r *ZoneFirewallLockdownService) ZoneLockdownListZoneLockdownRules(ctx context.Context, zoneIdentifier string, query ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesParams, opts ...option.RequestOption) (res *ZonelockdownResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZonelockdownResponseCollection struct {
	Errors     []ZonelockdownResponseCollectionError    `json:"errors"`
	Messages   []ZonelockdownResponseCollectionMessage  `json:"messages"`
	Result     []ZonelockdownResponseCollectionResult   `json:"result"`
	ResultInfo ZonelockdownResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZonelockdownResponseCollectionSuccess `json:"success"`
	JSON    zonelockdownResponseCollectionJSON    `json:"-"`
}

// zonelockdownResponseCollectionJSON contains the JSON metadata for the struct
// [ZonelockdownResponseCollection]
type zonelockdownResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonelockdownResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonelockdownResponseCollectionError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    zonelockdownResponseCollectionErrorJSON `json:"-"`
}

// zonelockdownResponseCollectionErrorJSON contains the JSON metadata for the
// struct [ZonelockdownResponseCollectionError]
type zonelockdownResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonelockdownResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonelockdownResponseCollectionMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    zonelockdownResponseCollectionMessageJSON `json:"-"`
}

// zonelockdownResponseCollectionMessageJSON contains the JSON metadata for the
// struct [ZonelockdownResponseCollectionMessage]
type zonelockdownResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonelockdownResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonelockdownResponseCollectionResult struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations ZonelockdownResponseCollectionResultConfigurations `json:"configurations,required"`
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
	URLs []string                                 `json:"urls,required"`
	JSON zonelockdownResponseCollectionResultJSON `json:"-"`
}

// zonelockdownResponseCollectionResultJSON contains the JSON metadata for the
// struct [ZonelockdownResponseCollectionResult]
type zonelockdownResponseCollectionResultJSON struct {
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

func (r *ZonelockdownResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [ZonelockdownResponseCollectionResultConfigurationsSchemasIPConfiguration] or
// [ZonelockdownResponseCollectionResultConfigurationsSchemasCidrConfiguration].
type ZonelockdownResponseCollectionResultConfigurations interface {
	implementsZonelockdownResponseCollectionResultConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZonelockdownResponseCollectionResultConfigurations)(nil)).Elem(), "")
}

type ZonelockdownResponseCollectionResultConfigurationsSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target ZonelockdownResponseCollectionResultConfigurationsSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                       `json:"value"`
	JSON  zonelockdownResponseCollectionResultConfigurationsSchemasIPConfigurationJSON `json:"-"`
}

// zonelockdownResponseCollectionResultConfigurationsSchemasIPConfigurationJSON
// contains the JSON metadata for the struct
// [ZonelockdownResponseCollectionResultConfigurationsSchemasIPConfiguration]
type zonelockdownResponseCollectionResultConfigurationsSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonelockdownResponseCollectionResultConfigurationsSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonelockdownResponseCollectionResultConfigurationsSchemasIPConfiguration) implementsZonelockdownResponseCollectionResultConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type ZonelockdownResponseCollectionResultConfigurationsSchemasIPConfigurationTarget string

const (
	ZonelockdownResponseCollectionResultConfigurationsSchemasIPConfigurationTargetIP ZonelockdownResponseCollectionResultConfigurationsSchemasIPConfigurationTarget = "ip"
)

type ZonelockdownResponseCollectionResultConfigurationsSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target ZonelockdownResponseCollectionResultConfigurationsSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                         `json:"value"`
	JSON  zonelockdownResponseCollectionResultConfigurationsSchemasCidrConfigurationJSON `json:"-"`
}

// zonelockdownResponseCollectionResultConfigurationsSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [ZonelockdownResponseCollectionResultConfigurationsSchemasCidrConfiguration]
type zonelockdownResponseCollectionResultConfigurationsSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonelockdownResponseCollectionResultConfigurationsSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonelockdownResponseCollectionResultConfigurationsSchemasCidrConfiguration) implementsZonelockdownResponseCollectionResultConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type ZonelockdownResponseCollectionResultConfigurationsSchemasCidrConfigurationTarget string

const (
	ZonelockdownResponseCollectionResultConfigurationsSchemasCidrConfigurationTargetIPRange ZonelockdownResponseCollectionResultConfigurationsSchemasCidrConfigurationTarget = "ip_range"
)

type ZonelockdownResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       zonelockdownResponseCollectionResultInfoJSON `json:"-"`
}

// zonelockdownResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [ZonelockdownResponseCollectionResultInfo]
type zonelockdownResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonelockdownResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZonelockdownResponseCollectionSuccess bool

const (
	ZonelockdownResponseCollectionSuccessTrue ZonelockdownResponseCollectionSuccess = true
)

type ZonelockdownResponseSingle struct {
	Errors   []ZonelockdownResponseSingleError   `json:"errors"`
	Messages []ZonelockdownResponseSingleMessage `json:"messages"`
	Result   ZonelockdownResponseSingleResult    `json:"result"`
	// Whether the API call was successful
	Success ZonelockdownResponseSingleSuccess `json:"success"`
	JSON    zonelockdownResponseSingleJSON    `json:"-"`
}

// zonelockdownResponseSingleJSON contains the JSON metadata for the struct
// [ZonelockdownResponseSingle]
type zonelockdownResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonelockdownResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonelockdownResponseSingleError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zonelockdownResponseSingleErrorJSON `json:"-"`
}

// zonelockdownResponseSingleErrorJSON contains the JSON metadata for the struct
// [ZonelockdownResponseSingleError]
type zonelockdownResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonelockdownResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonelockdownResponseSingleMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zonelockdownResponseSingleMessageJSON `json:"-"`
}

// zonelockdownResponseSingleMessageJSON contains the JSON metadata for the struct
// [ZonelockdownResponseSingleMessage]
type zonelockdownResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonelockdownResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonelockdownResponseSingleResult struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations ZonelockdownResponseSingleResultConfigurations `json:"configurations,required"`
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
	URLs []string                             `json:"urls,required"`
	JSON zonelockdownResponseSingleResultJSON `json:"-"`
}

// zonelockdownResponseSingleResultJSON contains the JSON metadata for the struct
// [ZonelockdownResponseSingleResult]
type zonelockdownResponseSingleResultJSON struct {
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

func (r *ZonelockdownResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [ZonelockdownResponseSingleResultConfigurationsSchemasIPConfiguration] or
// [ZonelockdownResponseSingleResultConfigurationsSchemasCidrConfiguration].
type ZonelockdownResponseSingleResultConfigurations interface {
	implementsZonelockdownResponseSingleResultConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZonelockdownResponseSingleResultConfigurations)(nil)).Elem(), "")
}

type ZonelockdownResponseSingleResultConfigurationsSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target ZonelockdownResponseSingleResultConfigurationsSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                   `json:"value"`
	JSON  zonelockdownResponseSingleResultConfigurationsSchemasIPConfigurationJSON `json:"-"`
}

// zonelockdownResponseSingleResultConfigurationsSchemasIPConfigurationJSON
// contains the JSON metadata for the struct
// [ZonelockdownResponseSingleResultConfigurationsSchemasIPConfiguration]
type zonelockdownResponseSingleResultConfigurationsSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonelockdownResponseSingleResultConfigurationsSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonelockdownResponseSingleResultConfigurationsSchemasIPConfiguration) implementsZonelockdownResponseSingleResultConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type ZonelockdownResponseSingleResultConfigurationsSchemasIPConfigurationTarget string

const (
	ZonelockdownResponseSingleResultConfigurationsSchemasIPConfigurationTargetIP ZonelockdownResponseSingleResultConfigurationsSchemasIPConfigurationTarget = "ip"
)

type ZonelockdownResponseSingleResultConfigurationsSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target ZonelockdownResponseSingleResultConfigurationsSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                     `json:"value"`
	JSON  zonelockdownResponseSingleResultConfigurationsSchemasCidrConfigurationJSON `json:"-"`
}

// zonelockdownResponseSingleResultConfigurationsSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [ZonelockdownResponseSingleResultConfigurationsSchemasCidrConfiguration]
type zonelockdownResponseSingleResultConfigurationsSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonelockdownResponseSingleResultConfigurationsSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZonelockdownResponseSingleResultConfigurationsSchemasCidrConfiguration) implementsZonelockdownResponseSingleResultConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type ZonelockdownResponseSingleResultConfigurationsSchemasCidrConfigurationTarget string

const (
	ZonelockdownResponseSingleResultConfigurationsSchemasCidrConfigurationTargetIPRange ZonelockdownResponseSingleResultConfigurationsSchemasCidrConfigurationTarget = "ip_range"
)

// Whether the API call was successful
type ZonelockdownResponseSingleSuccess bool

const (
	ZonelockdownResponseSingleSuccessTrue ZonelockdownResponseSingleSuccess = true
)

type ZoneFirewallLockdownDeleteResponse struct {
	Result ZoneFirewallLockdownDeleteResponseResult `json:"result"`
	JSON   zoneFirewallLockdownDeleteResponseJSON   `json:"-"`
}

// zoneFirewallLockdownDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallLockdownDeleteResponse]
type zoneFirewallLockdownDeleteResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallLockdownDeleteResponseResult struct {
	// The unique identifier of the Zone Lockdown rule.
	ID   string                                       `json:"id"`
	JSON zoneFirewallLockdownDeleteResponseResultJSON `json:"-"`
}

// zoneFirewallLockdownDeleteResponseResultJSON contains the JSON metadata for the
// struct [ZoneFirewallLockdownDeleteResponseResult]
type zoneFirewallLockdownDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallLockdownUpdateParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallLockdownUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesParams struct {
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

// URLQuery serializes
// [ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesParams]'s query parameters
// as `url.Values`.
func (r ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

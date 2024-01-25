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
func (r *ZoneFirewallLockdownService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *ZoneFirewallLockdownGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/lockdowns/%s", zoneIdentifier, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing Zone Lockdown rule.
func (r *ZoneFirewallLockdownService) Update(ctx context.Context, zoneIdentifier string, id string, body ZoneFirewallLockdownUpdateParams, opts ...option.RequestOption) (res *ZoneFirewallLockdownUpdateResponse, err error) {
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
func (r *ZoneFirewallLockdownService) ZoneLockdownNewAZoneLockdownRule(ctx context.Context, zoneIdentifier string, body ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleParams, opts ...option.RequestOption) (res *ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches Zone Lockdown rules. You can filter the results using several optional
// parameters.
func (r *ZoneFirewallLockdownService) ZoneLockdownListZoneLockdownRules(ctx context.Context, zoneIdentifier string, query ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesParams, opts ...option.RequestOption) (res *shared.Page[ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponse], err error) {
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

type ZoneFirewallLockdownGetResponse struct {
	Errors   []ZoneFirewallLockdownGetResponseError   `json:"errors"`
	Messages []ZoneFirewallLockdownGetResponseMessage `json:"messages"`
	Result   ZoneFirewallLockdownGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallLockdownGetResponseSuccess `json:"success"`
	JSON    zoneFirewallLockdownGetResponseJSON    `json:"-"`
}

// zoneFirewallLockdownGetResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallLockdownGetResponse]
type zoneFirewallLockdownGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallLockdownGetResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    zoneFirewallLockdownGetResponseErrorJSON `json:"-"`
}

// zoneFirewallLockdownGetResponseErrorJSON contains the JSON metadata for the
// struct [ZoneFirewallLockdownGetResponseError]
type zoneFirewallLockdownGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallLockdownGetResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneFirewallLockdownGetResponseMessageJSON `json:"-"`
}

// zoneFirewallLockdownGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneFirewallLockdownGetResponseMessage]
type zoneFirewallLockdownGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallLockdownGetResponseResult struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations ZoneFirewallLockdownGetResponseResultConfigurations `json:"configurations,required"`
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
	URLs []string                                  `json:"urls,required"`
	JSON zoneFirewallLockdownGetResponseResultJSON `json:"-"`
}

// zoneFirewallLockdownGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneFirewallLockdownGetResponseResult]
type zoneFirewallLockdownGetResponseResultJSON struct {
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

func (r *ZoneFirewallLockdownGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasIPConfiguration]
// or
// [ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration].
type ZoneFirewallLockdownGetResponseResultConfigurations interface {
	implementsZoneFirewallLockdownGetResponseResultConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallLockdownGetResponseResultConfigurations)(nil)).Elem(), "")
}

type ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                                `json:"value"`
	JSON  zoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasIPConfigurationJSON `json:"-"`
}

// zoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasIPConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasIPConfiguration]
type zoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasIPConfiguration) implementsZoneFirewallLockdownGetResponseResultConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasIPConfigurationTarget string

const (
	ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasIPConfigurationTargetIP ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasIPConfigurationTarget = "ip"
)

type ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                                  `json:"value"`
	JSON  zoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationJSON `json:"-"`
}

// zoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration]
type zoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration) implementsZoneFirewallLockdownGetResponseResultConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationTarget string

const (
	ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationTargetIPRange ZoneFirewallLockdownGetResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationTarget = "ip_range"
)

// Whether the API call was successful
type ZoneFirewallLockdownGetResponseSuccess bool

const (
	ZoneFirewallLockdownGetResponseSuccessTrue ZoneFirewallLockdownGetResponseSuccess = true
)

type ZoneFirewallLockdownUpdateResponse struct {
	Errors   []ZoneFirewallLockdownUpdateResponseError   `json:"errors"`
	Messages []ZoneFirewallLockdownUpdateResponseMessage `json:"messages"`
	Result   ZoneFirewallLockdownUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallLockdownUpdateResponseSuccess `json:"success"`
	JSON    zoneFirewallLockdownUpdateResponseJSON    `json:"-"`
}

// zoneFirewallLockdownUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneFirewallLockdownUpdateResponse]
type zoneFirewallLockdownUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallLockdownUpdateResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    zoneFirewallLockdownUpdateResponseErrorJSON `json:"-"`
}

// zoneFirewallLockdownUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneFirewallLockdownUpdateResponseError]
type zoneFirewallLockdownUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallLockdownUpdateResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneFirewallLockdownUpdateResponseMessageJSON `json:"-"`
}

// zoneFirewallLockdownUpdateResponseMessageJSON contains the JSON metadata for the
// struct [ZoneFirewallLockdownUpdateResponseMessage]
type zoneFirewallLockdownUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallLockdownUpdateResponseResult struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations ZoneFirewallLockdownUpdateResponseResultConfigurations `json:"configurations,required"`
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
	URLs []string                                     `json:"urls,required"`
	JSON zoneFirewallLockdownUpdateResponseResultJSON `json:"-"`
}

// zoneFirewallLockdownUpdateResponseResultJSON contains the JSON metadata for the
// struct [ZoneFirewallLockdownUpdateResponseResult]
type zoneFirewallLockdownUpdateResponseResultJSON struct {
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

func (r *ZoneFirewallLockdownUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasIPConfiguration]
// or
// [ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration].
type ZoneFirewallLockdownUpdateResponseResultConfigurations interface {
	implementsZoneFirewallLockdownUpdateResponseResultConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallLockdownUpdateResponseResultConfigurations)(nil)).Elem(), "")
}

type ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                                   `json:"value"`
	JSON  zoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasIPConfigurationJSON `json:"-"`
}

// zoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasIPConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasIPConfiguration]
type zoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasIPConfiguration) implementsZoneFirewallLockdownUpdateResponseResultConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasIPConfigurationTarget string

const (
	ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasIPConfigurationTargetIP ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasIPConfigurationTarget = "ip"
)

type ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                                     `json:"value"`
	JSON  zoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationJSON `json:"-"`
}

// zoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration]
type zoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration) implementsZoneFirewallLockdownUpdateResponseResultConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationTarget string

const (
	ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationTargetIPRange ZoneFirewallLockdownUpdateResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationTarget = "ip_range"
)

// Whether the API call was successful
type ZoneFirewallLockdownUpdateResponseSuccess bool

const (
	ZoneFirewallLockdownUpdateResponseSuccessTrue ZoneFirewallLockdownUpdateResponseSuccess = true
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

type ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponse struct {
	Errors   []ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseError   `json:"errors"`
	Messages []ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseMessage `json:"messages"`
	Result   ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseSuccess `json:"success"`
	JSON    zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseJSON    `json:"-"`
}

// zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseJSON contains the
// JSON metadata for the struct
// [ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponse]
type zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseError struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseErrorJSON `json:"-"`
}

// zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseErrorJSON contains
// the JSON metadata for the struct
// [ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseError]
type zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseMessage struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseMessageJSON `json:"-"`
}

// zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseMessageJSON contains
// the JSON metadata for the struct
// [ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseMessage]
type zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResult struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurations `json:"configurations,required"`
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
	URLs []string                                                               `json:"urls,required"`
	JSON zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultJSON `json:"-"`
}

// zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultJSON contains
// the JSON metadata for the struct
// [ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResult]
type zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultJSON struct {
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

func (r *ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasIPConfiguration]
// or
// [ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration].
type ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurations interface {
	implementsZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurations)(nil)).Elem(), "")
}

type ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                                                             `json:"value"`
	JSON  zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasIPConfigurationJSON `json:"-"`
}

// zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasIPConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasIPConfiguration]
type zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasIPConfiguration) implementsZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasIPConfigurationTarget string

const (
	ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasIPConfigurationTargetIP ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasIPConfigurationTarget = "ip"
)

type ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                                                               `json:"value"`
	JSON  zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationJSON `json:"-"`
}

// zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration]
type zoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasCidrConfiguration) implementsZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationTarget string

const (
	ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationTargetIPRange ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseResultConfigurationsDZw70ubTSchemasCidrConfigurationTarget = "ip_range"
)

// Whether the API call was successful
type ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseSuccess bool

const (
	ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseSuccessTrue ZoneFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseSuccess = true
)

type ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurations `json:"configurations,required"`
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
	URLs []string                                                          `json:"urls,required"`
	JSON zoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseJSON `json:"-"`
}

// zoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseJSON contains the
// JSON metadata for the struct
// [ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponse]
type zoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseJSON struct {
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

func (r *ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasIPConfiguration]
// or
// [ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasCidrConfiguration].
type ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurations interface {
	implementsZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurations)(nil)).Elem(), "")
}

type ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                                                        `json:"value"`
	JSON  zoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasIPConfigurationJSON `json:"-"`
}

// zoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasIPConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasIPConfiguration]
type zoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasIPConfiguration) implementsZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasIPConfigurationTarget string

const (
	ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasIPConfigurationTargetIP ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasIPConfigurationTarget = "ip"
)

type ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                                                          `json:"value"`
	JSON  zoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasCidrConfigurationJSON `json:"-"`
}

// zoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasCidrConfiguration]
type zoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasCidrConfiguration) implementsZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasCidrConfigurationTarget string

const (
	ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasCidrConfigurationTargetIPRange ZoneFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsDZw70ubTSchemasCidrConfigurationTarget = "ip_range"
)

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

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

// Updates an existing Zone Lockdown rule.
func (r *FirewallLockdownService) Update(ctx context.Context, zoneIdentifier string, id string, body FirewallLockdownUpdateParams, opts ...option.RequestOption) (res *FirewallLockdownUpdateResponse, err error) {
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
func (r *FirewallLockdownService) Get(ctx context.Context, zoneIdentifier string, id string, opts ...option.RequestOption) (res *FirewallLockdownGetResponse, err error) {
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

// Creates a new Zone Lockdown rule.
func (r *FirewallLockdownService) ZoneLockdownNewAZoneLockdownRule(ctx context.Context, zoneIdentifier string, body FirewallLockdownZoneLockdownNewAZoneLockdownRuleParams, opts ...option.RequestOption) (res *FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Zone Lockdown rules. You can filter the results using several optional
// parameters.
func (r *FirewallLockdownService) ZoneLockdownListZoneLockdownRules(ctx context.Context, zoneIdentifier string, query FirewallLockdownZoneLockdownListZoneLockdownRulesParams, opts ...option.RequestOption) (res *[]FirewallLockdownZoneLockdownListZoneLockdownRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelope
	path := fmt.Sprintf("zones/%s/firewall/lockdowns", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallLockdownUpdateResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations FirewallLockdownUpdateResponseConfigurations `json:"configurations,required"`
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
	URLs []string                           `json:"urls,required"`
	JSON firewallLockdownUpdateResponseJSON `json:"-"`
}

// firewallLockdownUpdateResponseJSON contains the JSON metadata for the struct
// [FirewallLockdownUpdateResponse]
type firewallLockdownUpdateResponseJSON struct {
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

func (r *FirewallLockdownUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasIPConfiguration] or
// [FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasCidrConfiguration].
type FirewallLockdownUpdateResponseConfigurations interface {
	implementsFirewallLockdownUpdateResponseConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallLockdownUpdateResponseConfigurations)(nil)).Elem(), "")
}

type FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                          `json:"value"`
	JSON  firewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON `json:"-"`
}

// firewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasIPConfiguration]
type firewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasIPConfiguration) implementsFirewallLockdownUpdateResponseConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget string

const (
	FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasIPConfigurationTargetIP FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget = "ip"
)

type FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                            `json:"value"`
	JSON  firewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON `json:"-"`
}

// firewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasCidrConfiguration]
type firewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasCidrConfiguration) implementsFirewallLockdownUpdateResponseConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget string

const (
	FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasCidrConfigurationTargetIPRange FirewallLockdownUpdateResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget = "ip_range"
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

type FirewallLockdownGetResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations FirewallLockdownGetResponseConfigurations `json:"configurations,required"`
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
	URLs []string                        `json:"urls,required"`
	JSON firewallLockdownGetResponseJSON `json:"-"`
}

// firewallLockdownGetResponseJSON contains the JSON metadata for the struct
// [FirewallLockdownGetResponse]
type firewallLockdownGetResponseJSON struct {
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

func (r *FirewallLockdownGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfiguration] or
// [FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfiguration].
type FirewallLockdownGetResponseConfigurations interface {
	implementsFirewallLockdownGetResponseConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallLockdownGetResponseConfigurations)(nil)).Elem(), "")
}

type FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                       `json:"value"`
	JSON  firewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON `json:"-"`
}

// firewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfiguration]
type firewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfiguration) implementsFirewallLockdownGetResponseConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget string

const (
	FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfigurationTargetIP FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget = "ip"
)

type FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                         `json:"value"`
	JSON  firewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON `json:"-"`
}

// firewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfiguration]
type firewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfiguration) implementsFirewallLockdownGetResponseConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget string

const (
	FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfigurationTargetIPRange FirewallLockdownGetResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget = "ip_range"
)

type FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurations `json:"configurations,required"`
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
	URLs []string                                                     `json:"urls,required"`
	JSON firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseJSON `json:"-"`
}

// firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseJSON contains the JSON
// metadata for the struct
// [FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponse]
type firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseJSON struct {
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

func (r *FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasIPConfiguration]
// or
// [FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasCidrConfiguration].
type FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurations interface {
	implementsFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurations)(nil)).Elem(), "")
}

type FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                                                    `json:"value"`
	JSON  firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON `json:"-"`
}

// firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasIPConfiguration]
type firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasIPConfiguration) implementsFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget string

const (
	FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasIPConfigurationTargetIP FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget = "ip"
)

type FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                                                      `json:"value"`
	JSON  firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON `json:"-"`
}

// firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasCidrConfiguration]
type firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasCidrConfiguration) implementsFirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget string

const (
	FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasCidrConfigurationTargetIPRange FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget = "ip_range"
)

type FirewallLockdownZoneLockdownListZoneLockdownRulesResponse struct {
	// The unique identifier of the Zone Lockdown rule.
	ID string `json:"id,required"`
	// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
	// specified in the Zone Lockdown rule. You can include any number of `ip` or
	// `ip_range` configurations.
	Configurations FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurations `json:"configurations,required"`
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
	URLs []string                                                      `json:"urls,required"`
	JSON firewallLockdownZoneLockdownListZoneLockdownRulesResponseJSON `json:"-"`
}

// firewallLockdownZoneLockdownListZoneLockdownRulesResponseJSON contains the JSON
// metadata for the struct
// [FirewallLockdownZoneLockdownListZoneLockdownRulesResponse]
type firewallLockdownZoneLockdownListZoneLockdownRulesResponseJSON struct {
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

func (r *FirewallLockdownZoneLockdownListZoneLockdownRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A list of IP addresses or CIDR ranges that will be allowed to access the URLs
// specified in the Zone Lockdown rule. You can include any number of `ip` or
// `ip_range` configurations.
//
// Union satisfied by
// [FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasIPConfiguration]
// or
// [FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasCidrConfiguration].
type FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurations interface {
	implementsFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurations()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurations)(nil)).Elem(), "")
}

type FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the Zone Lockdown rule.
	Target FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                                                                                     `json:"value"`
	JSON  firewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON `json:"-"`
}

// firewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasIPConfiguration]
type firewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasIPConfiguration) implementsFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurations() {
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the Zone Lockdown rule.
type FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget string

const (
	FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasIPConfigurationTargetIP FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasIPConfigurationTarget = "ip"
)

type FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasCidrConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the Zone Lockdown rule.
	Target FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`.
	Value string                                                                                                       `json:"value"`
	JSON  firewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON `json:"-"`
}

// firewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON
// contains the JSON metadata for the struct
// [FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasCidrConfiguration]
type firewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasCidrConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasCidrConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasCidrConfiguration) implementsFirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurations() {
}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the Zone Lockdown rule.
type FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget string

const (
	FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasCidrConfigurationTargetIPRange FirewallLockdownZoneLockdownListZoneLockdownRulesResponseConfigurationsLegacyJhsSchemasCidrConfigurationTarget = "ip_range"
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
	Result   FirewallLockdownUpdateResponse                   `json:"result,required,nullable"`
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
	Result   FirewallLockdownGetResponse                   `json:"result,required,nullable"`
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

type FirewallLockdownZoneLockdownNewAZoneLockdownRuleParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r FirewallLockdownZoneLockdownNewAZoneLockdownRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelope struct {
	Errors   []FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeJSON    `json:"-"`
}

// firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelope]
type firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeErrors struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeErrors]
type firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeMessages struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeMessages]
type firewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeSuccess bool

const (
	FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeSuccessTrue FirewallLockdownZoneLockdownNewAZoneLockdownRuleResponseEnvelopeSuccess = true
)

type FirewallLockdownZoneLockdownListZoneLockdownRulesParams struct {
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

// URLQuery serializes [FirewallLockdownZoneLockdownListZoneLockdownRulesParams]'s
// query parameters as `url.Values`.
func (r FirewallLockdownZoneLockdownListZoneLockdownRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelope struct {
	Errors   []FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeMessages `json:"messages,required"`
	Result   []FirewallLockdownZoneLockdownListZoneLockdownRulesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       firewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeJSON       `json:"-"`
}

// firewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelope]
type firewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    firewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeErrors]
type firewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    firewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeMessages]
type firewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeSuccess bool

const (
	FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeSuccessTrue FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeSuccess = true
)

type FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                         `json:"total_count"`
	JSON       firewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeResultInfoJSON `json:"-"`
}

// firewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeResultInfo]
type firewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallLockdownZoneLockdownListZoneLockdownRulesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/shared"
)

// GatewayRuleService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayRuleService] method instead.
type GatewayRuleService struct {
	Options []option.RequestOption
}

// NewGatewayRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayRuleService(opts ...option.RequestOption) (r *GatewayRuleService) {
	r = &GatewayRuleService{}
	r.Options = opts
	return
}

// Create a new Zero Trust Gateway rule.
func (r *GatewayRuleService) New(ctx context.Context, params GatewayRuleNewParams, opts ...option.RequestOption) (res *GatewayRule, err error) {
	var env GatewayRuleNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a configured Zero Trust Gateway rule.
func (r *GatewayRuleService) Update(ctx context.Context, ruleID string, params GatewayRuleUpdateParams, opts ...option.RequestOption) (res *GatewayRule, err error) {
	var env GatewayRuleUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Zero Trust Gateway rules for an account.
func (r *GatewayRuleService) List(ctx context.Context, query GatewayRuleListParams, opts ...option.RequestOption) (res *pagination.SinglePage[GatewayRule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// List Zero Trust Gateway rules for an account.
func (r *GatewayRuleService) ListAutoPaging(ctx context.Context, query GatewayRuleListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[GatewayRule] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Zero Trust Gateway rule.
func (r *GatewayRuleService) Delete(ctx context.Context, ruleID string, body GatewayRuleDeleteParams, opts ...option.RequestOption) (res *GatewayRuleDeleteResponse, err error) {
	var env GatewayRuleDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", body.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a single Zero Trust Gateway rule.
func (r *GatewayRuleService) Get(ctx context.Context, ruleID string, query GatewayRuleGetParams, opts ...option.RequestOption) (res *GatewayRule, err error) {
	var env GatewayRuleGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Zero Trust Gateway rules for the parent account of an account in the MSP
// configuration.
func (r *GatewayRuleService) ListTenant(ctx context.Context, query GatewayRuleListTenantParams, opts ...option.RequestOption) (res *pagination.SinglePage[GatewayRule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules/tenant", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// List Zero Trust Gateway rules for the parent account of an account in the MSP
// configuration.
func (r *GatewayRuleService) ListTenantAutoPaging(ctx context.Context, query GatewayRuleListTenantParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[GatewayRule] {
	return pagination.NewSinglePageAutoPager(r.ListTenant(ctx, query, opts...))
}

// Resets the expiration of a Zero Trust Gateway Rule if its duration elapsed and
// it has a default duration. The Zero Trust Gateway Rule must have values for both
// `expiration.expires_at` and `expiration.duration`.
func (r *GatewayRuleService) ResetExpiration(ctx context.Context, ruleID string, body GatewayRuleResetExpirationParams, opts ...option.RequestOption) (res *GatewayRule, err error) {
	var env GatewayRuleResetExpirationResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s/reset_expiration", body.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSResolverSettingsV4 struct {
	// Specify the IPv4 address of the upstream resolver.
	IP string `json:"ip,required"`
	// Specify a port number to use for the upstream resolver. Defaults to 53 if
	// unspecified.
	Port int64 `json:"port"`
	// Indicate whether to connect to this resolver over a private network. Must set
	// when vnet_id set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Specify an optional virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                    `json:"vnet_id"`
	JSON   dnsResolverSettingsV4JSON `json:"-"`
}

// dnsResolverSettingsV4JSON contains the JSON metadata for the struct
// [DNSResolverSettingsV4]
type dnsResolverSettingsV4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *DNSResolverSettingsV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsResolverSettingsV4JSON) RawJSON() string {
	return r.raw
}

type DNSResolverSettingsV4Param struct {
	// Specify the IPv4 address of the upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// Specify a port number to use for the upstream resolver. Defaults to 53 if
	// unspecified.
	Port param.Field[int64] `json:"port"`
	// Indicate whether to connect to this resolver over a private network. Must set
	// when vnet_id set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Specify an optional virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r DNSResolverSettingsV4Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSResolverSettingsV6 struct {
	// Specify the IPv6 address of the upstream resolver.
	IP string `json:"ip,required"`
	// Specify a port number to use for the upstream resolver. Defaults to 53 if
	// unspecified.
	Port int64 `json:"port"`
	// Indicate whether to connect to this resolver over a private network. Must set
	// when vnet_id set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Specify an optional virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                    `json:"vnet_id"`
	JSON   dnsResolverSettingsV6JSON `json:"-"`
}

// dnsResolverSettingsV6JSON contains the JSON metadata for the struct
// [DNSResolverSettingsV6]
type dnsResolverSettingsV6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *DNSResolverSettingsV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsResolverSettingsV6JSON) RawJSON() string {
	return r.raw
}

type DNSResolverSettingsV6Param struct {
	// Specify the IPv6 address of the upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// Specify a port number to use for the upstream resolver. Defaults to 53 if
	// unspecified.
	Port param.Field[int64] `json:"port"`
	// Indicate whether to connect to this resolver over a private network. Must set
	// when vnet_id set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Specify an optional virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r DNSResolverSettingsV6Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify the protocol or layer to use.
type GatewayFilter string

const (
	GatewayFilterHTTP        GatewayFilter = "http"
	GatewayFilterDNS         GatewayFilter = "dns"
	GatewayFilterL4          GatewayFilter = "l4"
	GatewayFilterEgress      GatewayFilter = "egress"
	GatewayFilterDNSResolver GatewayFilter = "dns_resolver"
)

func (r GatewayFilter) IsKnown() bool {
	switch r {
	case GatewayFilterHTTP, GatewayFilterDNS, GatewayFilterL4, GatewayFilterEgress, GatewayFilterDNSResolver:
		return true
	}
	return false
}

type GatewayRule struct {
	// Specify the action to perform when the associated traffic, identity, and device
	// posture expressions either absent or evaluate to `true`.
	Action GatewayRuleAction `json:"action,required"`
	// Specify whether the rule is enabled.
	Enabled bool `json:"enabled,required"`
	// Specify the protocol or layer to evaluate the traffic, identity, and device
	// posture expressions. Can only contain a single value.
	Filters []GatewayFilter `json:"filters,required"`
	// Specify the rule name.
	Name string `json:"name,required"`
	// Set the order of your rules. Lower values indicate higher precedence. At each
	// processing phase, evaluate applicable rules in ascending order of this value.
	// Refer to
	// [Order of enforcement](http://developers.cloudflare.com/learning-paths/secure-internet-traffic/understand-policies/order-of-enforcement/#manage-precedence-with-terraform)
	// to manage precedence via Terraform.
	Precedence int64 `json:"precedence,required"`
	// Specify the wirefilter expression used for traffic matching. The API
	// automatically formats and sanitizes expressions before storing them. To prevent
	// Terraform state drift, use the formatted expression returned in the API
	// response.
	Traffic string `json:"traffic,required"`
	// Identify the API resource with a UUID.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Indicate the date of deletion, if any.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// Specify the rule description.
	Description string `json:"description"`
	// Specify the wirefilter expression used for device posture check. The API
	// automatically formats and sanitizes expressions before storing them. To prevent
	// Terraform state drift, use the formatted expression returned in the API
	// response.
	DevicePosture string `json:"device_posture"`
	// Defines the expiration time stamp and default duration of a DNS policy. Takes
	// precedence over the policy's `schedule` configuration, if any. This does not
	// apply to HTTP or network policies. Settable only for `dns` rules.
	Expiration GatewayRuleExpiration `json:"expiration,nullable"`
	// Specify the wirefilter expression used for identity matching. The API
	// automatically formats and sanitizes expressions before storing them. To prevent
	// Terraform state drift, use the formatted expression returned in the API
	// response.
	Identity string `json:"identity"`
	// Indicate that this rule is shared via the Orgs API and read only.
	ReadOnly bool `json:"read_only"`
	// Defines settings for this rule. Settings apply only to specific rule types and
	// must use compatible selectors. If Terraform detects drift, confirm the setting
	// supports your rule type and check whether the API modifies the value. Use
	// API-returned values in your configuration to prevent drift.
	RuleSettings RuleSetting `json:"rule_settings"`
	// Defines the schedule for activating DNS policies. Settable only for `dns` and
	// `dns_resolver` rules.
	Schedule Schedule `json:"schedule,nullable"`
	// Indicate that this rule is sharable via the Orgs API.
	Sharable bool `json:"sharable"`
	// Provide the account tag of the account that created the rule.
	SourceAccount string    `json:"source_account"`
	UpdatedAt     time.Time `json:"updated_at" format:"date-time"`
	// Indicate the version number of the rule(read-only).
	Version int64 `json:"version"`
	// Indicate a warning for a misconfigured rule, if any.
	WarningStatus string          `json:"warning_status,nullable"`
	JSON          gatewayRuleJSON `json:"-"`
}

// gatewayRuleJSON contains the JSON metadata for the struct [GatewayRule]
type gatewayRuleJSON struct {
	Action        apijson.Field
	Enabled       apijson.Field
	Filters       apijson.Field
	Name          apijson.Field
	Precedence    apijson.Field
	Traffic       apijson.Field
	ID            apijson.Field
	CreatedAt     apijson.Field
	DeletedAt     apijson.Field
	Description   apijson.Field
	DevicePosture apijson.Field
	Expiration    apijson.Field
	Identity      apijson.Field
	ReadOnly      apijson.Field
	RuleSettings  apijson.Field
	Schedule      apijson.Field
	Sharable      apijson.Field
	SourceAccount apijson.Field
	UpdatedAt     apijson.Field
	Version       apijson.Field
	WarningStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleJSON) RawJSON() string {
	return r.raw
}

// Specify the action to perform when the associated traffic, identity, and device
// posture expressions either absent or evaluate to `true`.
type GatewayRuleAction string

const (
	GatewayRuleActionOn           GatewayRuleAction = "on"
	GatewayRuleActionOff          GatewayRuleAction = "off"
	GatewayRuleActionAllow        GatewayRuleAction = "allow"
	GatewayRuleActionBlock        GatewayRuleAction = "block"
	GatewayRuleActionScan         GatewayRuleAction = "scan"
	GatewayRuleActionNoscan       GatewayRuleAction = "noscan"
	GatewayRuleActionSafesearch   GatewayRuleAction = "safesearch"
	GatewayRuleActionYtrestricted GatewayRuleAction = "ytrestricted"
	GatewayRuleActionIsolate      GatewayRuleAction = "isolate"
	GatewayRuleActionNoisolate    GatewayRuleAction = "noisolate"
	GatewayRuleActionOverride     GatewayRuleAction = "override"
	GatewayRuleActionL4Override   GatewayRuleAction = "l4_override"
	GatewayRuleActionEgress       GatewayRuleAction = "egress"
	GatewayRuleActionResolve      GatewayRuleAction = "resolve"
	GatewayRuleActionQuarantine   GatewayRuleAction = "quarantine"
	GatewayRuleActionRedirect     GatewayRuleAction = "redirect"
)

func (r GatewayRuleAction) IsKnown() bool {
	switch r {
	case GatewayRuleActionOn, GatewayRuleActionOff, GatewayRuleActionAllow, GatewayRuleActionBlock, GatewayRuleActionScan, GatewayRuleActionNoscan, GatewayRuleActionSafesearch, GatewayRuleActionYtrestricted, GatewayRuleActionIsolate, GatewayRuleActionNoisolate, GatewayRuleActionOverride, GatewayRuleActionL4Override, GatewayRuleActionEgress, GatewayRuleActionResolve, GatewayRuleActionQuarantine, GatewayRuleActionRedirect:
		return true
	}
	return false
}

// Defines the expiration time stamp and default duration of a DNS policy. Takes
// precedence over the policy's `schedule` configuration, if any. This does not
// apply to HTTP or network policies. Settable only for `dns` rules.
type GatewayRuleExpiration struct {
	// Show the timestamp when the policy expires and stops applying. The value must
	// follow RFC 3339 and include a UTC offset. The system accepts non-zero offsets
	// but converts them to the equivalent UTC+00:00 value and returns timestamps with
	// a trailing Z. Expiration policies ignore client timezones and expire globally at
	// the specified expires_at time.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// Defines the default duration a policy active in minutes. Must set in order to
	// use the `reset_expiration` endpoint on this rule.
	Duration int64 `json:"duration"`
	// Indicates whether the policy is expired.
	Expired bool                      `json:"expired"`
	JSON    gatewayRuleExpirationJSON `json:"-"`
}

// gatewayRuleExpirationJSON contains the JSON metadata for the struct
// [GatewayRuleExpiration]
type gatewayRuleExpirationJSON struct {
	ExpiresAt   apijson.Field
	Duration    apijson.Field
	Expired     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleExpiration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleExpirationJSON) RawJSON() string {
	return r.raw
}

// Defines settings for this rule. Settings apply only to specific rule types and
// must use compatible selectors. If Terraform detects drift, confirm the setting
// supports your rule type and check whether the API modifies the value. Use
// API-returned values in your configuration to prevent drift.
type RuleSetting struct {
	// Add custom headers to allowed requests as key-value pairs. Use header names as
	// keys that map to arrays of header values. Settable only for `http` rules with
	// the action set to `allow`.
	AddHeaders map[string][]string `json:"add_headers,nullable"`
	// Set to enable MSP children to bypass this rule. Only parent MSP accounts can set
	// this. this rule. Settable for all types of rules.
	AllowChildBypass bool `json:"allow_child_bypass,nullable"`
	// Define the settings for the Audit SSH action. Settable only for `l4` rules with
	// `audit_ssh` action.
	AuditSSH RuleSettingAuditSSH `json:"audit_ssh,nullable"`
	// Configure browser isolation behavior. Settable only for `http` rules with the
	// action set to `isolate`.
	BISOAdminControls RuleSettingBISOAdminControls `json:"biso_admin_controls"`
	// Configure custom block page settings. If missing or null, use the account
	// settings. Settable only for `http` rules with the action set to `block`.
	BlockPage RuleSettingBlockPage `json:"block_page,nullable"`
	// Enable the custom block page. Settable only for `dns` rules with action `block`.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// Explain why the rule blocks the request. The custom block page shows this text
	// (if enabled). Settable only for `dns`, `l4`, and `http` rules when the action
	// set to `block`.
	BlockReason string `json:"block_reason,nullable"`
	// Set to enable MSP accounts to bypass their parent's rules. Only MSP child
	// accounts can set this. Settable for all types of rules.
	BypassParentRule bool `json:"bypass_parent_rule,nullable"`
	// Configure session check behavior. Settable only for `l4` and `http` rules with
	// the action set to `allow`.
	CheckSession RuleSettingCheckSession `json:"check_session,nullable"`
	// Configure custom resolvers to route queries that match the resolver policy.
	// Unused with 'resolve_dns_through_cloudflare' or 'resolve_dns_internally'
	// settings. DNS queries get routed to the address closest to their origin. Only
	// valid when a rule's action set to 'resolve'. Settable only for `dns_resolver`
	// rules.
	DNSResolvers RuleSettingDNSResolvers `json:"dns_resolvers,nullable"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs. Settable only for `egress` rules.
	Egress RuleSettingEgress `json:"egress,nullable"`
	// Ignore category matches at CNAME domains in a response. When off, evaluate
	// categories in this rule against all CNAME domain categories in the response.
	// Settable only for `dns` and `dns_resolver` rules.
	IgnoreCNAMECategoryMatches bool `json:"ignore_cname_category_matches"`
	// Specify whether to disable DNSSEC validation (for Allow actions) [INSECURE].
	// Settable only for `dns` rules.
	InsecureDisableDNSSECValidation bool `json:"insecure_disable_dnssec_validation"`
	// Enable IPs in DNS resolver category blocks. The system blocks only domain name
	// categories unless you enable this setting. Settable only for `dns` and
	// `dns_resolver` rules.
	IPCategories bool `json:"ip_categories"`
	// Indicates whether to include IPs in DNS resolver indicator feed blocks. Default,
	// indicator feeds block only domain names. Settable only for `dns` and
	// `dns_resolver` rules.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port. Settable
	// only for `l4` rules with the action set to `l4_override`.
	L4override RuleSettingL4override `json:"l4override,nullable"`
	// Configure a notification to display on the user's device when this rule matched.
	// Settable for all types of rules with the action set to `block`.
	NotificationSettings RuleSettingNotificationSettings `json:"notification_settings,nullable"`
	// Defines a hostname for override, for the matching DNS queries. Settable only for
	// `dns` rules with the action set to `override`.
	OverrideHost string `json:"override_host"`
	// Defines a an IP or set of IPs for overriding matched DNS queries. Settable only
	// for `dns` rules with the action set to `override`.
	OverrideIPs []string `json:"override_ips,nullable"`
	// Configure DLP payload logging. Settable only for `http` rules.
	PayloadLog RuleSettingPayloadLog `json:"payload_log,nullable"`
	// Configure settings that apply to quarantine rules. Settable only for `http`
	// rules.
	Quarantine RuleSettingQuarantine `json:"quarantine,nullable"`
	// Apply settings to redirect rules. Settable only for `http` rules with the action
	// set to `redirect`.
	Redirect RuleSettingRedirect `json:"redirect,nullable"`
	// Configure to forward the query to the internal DNS service, passing the
	// specified 'view_id' as input. Not used when 'dns_resolvers' is specified or
	// 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action set to
	// 'resolve'. Settable only for `dns_resolver` rules.
	ResolveDNSInternally RuleSettingResolveDNSInternally `json:"resolve_dns_internally,nullable"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot set when 'dns_resolvers' specified or 'resolve_dns_internally'
	// is set. Only valid when a rule's action set to 'resolve'. Settable only for
	// `dns_resolver` rules.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare,nullable"`
	// Configure behavior when an upstream certificate is invalid or an SSL error
	// occurs. Settable only for `http` rules with the action set to `allow`.
	UntrustedCERT RuleSettingUntrustedCERT `json:"untrusted_cert,nullable"`
	JSON          ruleSettingJSON          `json:"-"`
}

// ruleSettingJSON contains the JSON metadata for the struct [RuleSetting]
type ruleSettingJSON struct {
	AddHeaders                      apijson.Field
	AllowChildBypass                apijson.Field
	AuditSSH                        apijson.Field
	BISOAdminControls               apijson.Field
	BlockPage                       apijson.Field
	BlockPageEnabled                apijson.Field
	BlockReason                     apijson.Field
	BypassParentRule                apijson.Field
	CheckSession                    apijson.Field
	DNSResolvers                    apijson.Field
	Egress                          apijson.Field
	IgnoreCNAMECategoryMatches      apijson.Field
	InsecureDisableDNSSECValidation apijson.Field
	IPCategories                    apijson.Field
	IPIndicatorFeeds                apijson.Field
	L4override                      apijson.Field
	NotificationSettings            apijson.Field
	OverrideHost                    apijson.Field
	OverrideIPs                     apijson.Field
	PayloadLog                      apijson.Field
	Quarantine                      apijson.Field
	Redirect                        apijson.Field
	ResolveDNSInternally            apijson.Field
	ResolveDNSThroughCloudflare     apijson.Field
	UntrustedCERT                   apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *RuleSetting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingJSON) RawJSON() string {
	return r.raw
}

// Define the settings for the Audit SSH action. Settable only for `l4` rules with
// `audit_ssh` action.
type RuleSettingAuditSSH struct {
	// Enable SSH command logging.
	CommandLogging bool                    `json:"command_logging"`
	JSON           ruleSettingAuditSSHJSON `json:"-"`
}

// ruleSettingAuditSSHJSON contains the JSON metadata for the struct
// [RuleSettingAuditSSH]
type ruleSettingAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuleSettingAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingAuditSSHJSON) RawJSON() string {
	return r.raw
}

// Configure browser isolation behavior. Settable only for `http` rules with the
// action set to `isolate`.
type RuleSettingBISOAdminControls struct {
	// Configure copy behavior. If set to remote_only, users cannot copy isolated
	// content from the remote browser to the local clipboard. If this field is absent,
	// copying remains enabled. Applies only when version == "v2".
	Copy RuleSettingBISOAdminControlsCopy `json:"copy"`
	// Set to false to enable copy-pasting. Only applies when `version == "v1"`.
	DCP bool `json:"dcp"`
	// Set to false to enable downloading. Only applies when `version == "v1"`.
	DD bool `json:"dd"`
	// Set to false to enable keyboard usage. Only applies when `version == "v1"`.
	DK bool `json:"dk"`
	// Configure download behavior. When set to remote_only, users can view downloads
	// but cannot save them. Applies only when version == "v2".
	Download RuleSettingBISOAdminControlsDownload `json:"download"`
	// Set to false to enable printing. Only applies when `version == "v1"`.
	DP bool `json:"dp"`
	// Set to false to enable uploading. Only applies when `version == "v1"`.
	DU bool `json:"du"`
	// Configure keyboard usage behavior. If this field is absent, keyboard usage
	// remains enabled. Applies only when version == "v2".
	Keyboard RuleSettingBISOAdminControlsKeyboard `json:"keyboard"`
	// Configure paste behavior. If set to remote_only, users cannot paste content from
	// the local clipboard into isolated pages. If this field is absent, pasting
	// remains enabled. Applies only when version == "v2".
	Paste RuleSettingBISOAdminControlsPaste `json:"paste"`
	// Configure print behavior. Default, Printing is enabled. Applies only when
	// version == "v2".
	Printing RuleSettingBISOAdminControlsPrinting `json:"printing"`
	// Configure upload behavior. If this field is absent, uploading remains enabled.
	// Applies only when version == "v2".
	Upload RuleSettingBISOAdminControlsUpload `json:"upload"`
	// Indicate which version of the browser isolation controls should apply.
	Version RuleSettingBISOAdminControlsVersion `json:"version"`
	JSON    ruleSettingBISOAdminControlsJSON    `json:"-"`
}

// ruleSettingBISOAdminControlsJSON contains the JSON metadata for the struct
// [RuleSettingBISOAdminControls]
type ruleSettingBISOAdminControlsJSON struct {
	Copy        apijson.Field
	DCP         apijson.Field
	DD          apijson.Field
	DK          apijson.Field
	Download    apijson.Field
	DP          apijson.Field
	DU          apijson.Field
	Keyboard    apijson.Field
	Paste       apijson.Field
	Printing    apijson.Field
	Upload      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingBISOAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingBISOAdminControlsJSON) RawJSON() string {
	return r.raw
}

// Configure copy behavior. If set to remote_only, users cannot copy isolated
// content from the remote browser to the local clipboard. If this field is absent,
// copying remains enabled. Applies only when version == "v2".
type RuleSettingBISOAdminControlsCopy string

const (
	RuleSettingBISOAdminControlsCopyEnabled    RuleSettingBISOAdminControlsCopy = "enabled"
	RuleSettingBISOAdminControlsCopyDisabled   RuleSettingBISOAdminControlsCopy = "disabled"
	RuleSettingBISOAdminControlsCopyRemoteOnly RuleSettingBISOAdminControlsCopy = "remote_only"
)

func (r RuleSettingBISOAdminControlsCopy) IsKnown() bool {
	switch r {
	case RuleSettingBISOAdminControlsCopyEnabled, RuleSettingBISOAdminControlsCopyDisabled, RuleSettingBISOAdminControlsCopyRemoteOnly:
		return true
	}
	return false
}

// Configure download behavior. When set to remote_only, users can view downloads
// but cannot save them. Applies only when version == "v2".
type RuleSettingBISOAdminControlsDownload string

const (
	RuleSettingBISOAdminControlsDownloadEnabled    RuleSettingBISOAdminControlsDownload = "enabled"
	RuleSettingBISOAdminControlsDownloadDisabled   RuleSettingBISOAdminControlsDownload = "disabled"
	RuleSettingBISOAdminControlsDownloadRemoteOnly RuleSettingBISOAdminControlsDownload = "remote_only"
)

func (r RuleSettingBISOAdminControlsDownload) IsKnown() bool {
	switch r {
	case RuleSettingBISOAdminControlsDownloadEnabled, RuleSettingBISOAdminControlsDownloadDisabled, RuleSettingBISOAdminControlsDownloadRemoteOnly:
		return true
	}
	return false
}

// Configure keyboard usage behavior. If this field is absent, keyboard usage
// remains enabled. Applies only when version == "v2".
type RuleSettingBISOAdminControlsKeyboard string

const (
	RuleSettingBISOAdminControlsKeyboardEnabled  RuleSettingBISOAdminControlsKeyboard = "enabled"
	RuleSettingBISOAdminControlsKeyboardDisabled RuleSettingBISOAdminControlsKeyboard = "disabled"
)

func (r RuleSettingBISOAdminControlsKeyboard) IsKnown() bool {
	switch r {
	case RuleSettingBISOAdminControlsKeyboardEnabled, RuleSettingBISOAdminControlsKeyboardDisabled:
		return true
	}
	return false
}

// Configure paste behavior. If set to remote_only, users cannot paste content from
// the local clipboard into isolated pages. If this field is absent, pasting
// remains enabled. Applies only when version == "v2".
type RuleSettingBISOAdminControlsPaste string

const (
	RuleSettingBISOAdminControlsPasteEnabled    RuleSettingBISOAdminControlsPaste = "enabled"
	RuleSettingBISOAdminControlsPasteDisabled   RuleSettingBISOAdminControlsPaste = "disabled"
	RuleSettingBISOAdminControlsPasteRemoteOnly RuleSettingBISOAdminControlsPaste = "remote_only"
)

func (r RuleSettingBISOAdminControlsPaste) IsKnown() bool {
	switch r {
	case RuleSettingBISOAdminControlsPasteEnabled, RuleSettingBISOAdminControlsPasteDisabled, RuleSettingBISOAdminControlsPasteRemoteOnly:
		return true
	}
	return false
}

// Configure print behavior. Default, Printing is enabled. Applies only when
// version == "v2".
type RuleSettingBISOAdminControlsPrinting string

const (
	RuleSettingBISOAdminControlsPrintingEnabled  RuleSettingBISOAdminControlsPrinting = "enabled"
	RuleSettingBISOAdminControlsPrintingDisabled RuleSettingBISOAdminControlsPrinting = "disabled"
)

func (r RuleSettingBISOAdminControlsPrinting) IsKnown() bool {
	switch r {
	case RuleSettingBISOAdminControlsPrintingEnabled, RuleSettingBISOAdminControlsPrintingDisabled:
		return true
	}
	return false
}

// Configure upload behavior. If this field is absent, uploading remains enabled.
// Applies only when version == "v2".
type RuleSettingBISOAdminControlsUpload string

const (
	RuleSettingBISOAdminControlsUploadEnabled  RuleSettingBISOAdminControlsUpload = "enabled"
	RuleSettingBISOAdminControlsUploadDisabled RuleSettingBISOAdminControlsUpload = "disabled"
)

func (r RuleSettingBISOAdminControlsUpload) IsKnown() bool {
	switch r {
	case RuleSettingBISOAdminControlsUploadEnabled, RuleSettingBISOAdminControlsUploadDisabled:
		return true
	}
	return false
}

// Indicate which version of the browser isolation controls should apply.
type RuleSettingBISOAdminControlsVersion string

const (
	RuleSettingBISOAdminControlsVersionV1 RuleSettingBISOAdminControlsVersion = "v1"
	RuleSettingBISOAdminControlsVersionV2 RuleSettingBISOAdminControlsVersion = "v2"
)

func (r RuleSettingBISOAdminControlsVersion) IsKnown() bool {
	switch r {
	case RuleSettingBISOAdminControlsVersionV1, RuleSettingBISOAdminControlsVersionV2:
		return true
	}
	return false
}

// Configure custom block page settings. If missing or null, use the account
// settings. Settable only for `http` rules with the action set to `block`.
type RuleSettingBlockPage struct {
	// Specify the URI to which the user is redirected.
	TargetURI string `json:"target_uri,required" format:"uri"`
	// Specify whether to pass the context information as query parameters.
	IncludeContext bool                     `json:"include_context"`
	JSON           ruleSettingBlockPageJSON `json:"-"`
}

// ruleSettingBlockPageJSON contains the JSON metadata for the struct
// [RuleSettingBlockPage]
type ruleSettingBlockPageJSON struct {
	TargetURI      apijson.Field
	IncludeContext apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuleSettingBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingBlockPageJSON) RawJSON() string {
	return r.raw
}

// Configure session check behavior. Settable only for `l4` and `http` rules with
// the action set to `allow`.
type RuleSettingCheckSession struct {
	// Sets the required session freshness threshold. The API returns a normalized
	// version of this value.
	Duration string `json:"duration"`
	// Enable session enforcement.
	Enforce bool                        `json:"enforce"`
	JSON    ruleSettingCheckSessionJSON `json:"-"`
}

// ruleSettingCheckSessionJSON contains the JSON metadata for the struct
// [RuleSettingCheckSession]
type ruleSettingCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingCheckSessionJSON) RawJSON() string {
	return r.raw
}

// Configure custom resolvers to route queries that match the resolver policy.
// Unused with 'resolve_dns_through_cloudflare' or 'resolve_dns_internally'
// settings. DNS queries get routed to the address closest to their origin. Only
// valid when a rule's action set to 'resolve'. Settable only for `dns_resolver`
// rules.
type RuleSettingDNSResolvers struct {
	IPV4 []DNSResolverSettingsV4     `json:"ipv4"`
	IPV6 []DNSResolverSettingsV6     `json:"ipv6"`
	JSON ruleSettingDNSResolversJSON `json:"-"`
}

// ruleSettingDNSResolversJSON contains the JSON metadata for the struct
// [RuleSettingDNSResolvers]
type ruleSettingDNSResolversJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingDNSResolversJSON) RawJSON() string {
	return r.raw
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs. Settable only for `egress` rules.
type RuleSettingEgress struct {
	// Specify the IPv4 address to use for egress.
	IPV4 string `json:"ipv4"`
	// Specify the fallback IPv4 address to use for egress when the primary IPv4 fails.
	// Set '0.0.0.0' to indicate local egress via WARP IPs.
	IPV4Fallback string `json:"ipv4_fallback"`
	// Specify the IPv6 range to use for egress.
	IPV6 string                `json:"ipv6"`
	JSON ruleSettingEgressJSON `json:"-"`
}

// ruleSettingEgressJSON contains the JSON metadata for the struct
// [RuleSettingEgress]
type ruleSettingEgressJSON struct {
	IPV4         apijson.Field
	IPV4Fallback apijson.Field
	IPV6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RuleSettingEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingEgressJSON) RawJSON() string {
	return r.raw
}

// Send matching traffic to the supplied destination IP address and port. Settable
// only for `l4` rules with the action set to `l4_override`.
type RuleSettingL4override struct {
	// Defines the IPv4 or IPv6 address.
	IP string `json:"ip"`
	// Defines a port number to use for TCP/UDP overrides.
	Port int64                     `json:"port"`
	JSON ruleSettingL4overrideJSON `json:"-"`
}

// ruleSettingL4overrideJSON contains the JSON metadata for the struct
// [RuleSettingL4override]
type ruleSettingL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingL4overrideJSON) RawJSON() string {
	return r.raw
}

// Configure a notification to display on the user's device when this rule matched.
// Settable for all types of rules with the action set to `block`.
type RuleSettingNotificationSettings struct {
	// Enable notification.
	Enabled bool `json:"enabled"`
	// Indicates whether to pass the context information as query parameters.
	IncludeContext bool `json:"include_context"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Defines an optional URL to direct users to additional information. If unset, the
	// notification opens a block page.
	SupportURL string                              `json:"support_url"`
	JSON       ruleSettingNotificationSettingsJSON `json:"-"`
}

// ruleSettingNotificationSettingsJSON contains the JSON metadata for the struct
// [RuleSettingNotificationSettings]
type ruleSettingNotificationSettingsJSON struct {
	Enabled        apijson.Field
	IncludeContext apijson.Field
	Msg            apijson.Field
	SupportURL     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuleSettingNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Configure DLP payload logging. Settable only for `http` rules.
type RuleSettingPayloadLog struct {
	// Enable DLP payload logging for this rule.
	Enabled bool                      `json:"enabled"`
	JSON    ruleSettingPayloadLogJSON `json:"-"`
}

// ruleSettingPayloadLogJSON contains the JSON metadata for the struct
// [RuleSettingPayloadLog]
type ruleSettingPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingPayloadLogJSON) RawJSON() string {
	return r.raw
}

// Configure settings that apply to quarantine rules. Settable only for `http`
// rules.
type RuleSettingQuarantine struct {
	// Specify the types of files to sandbox.
	FileTypes []RuleSettingQuarantineFileType `json:"file_types"`
	JSON      ruleSettingQuarantineJSON       `json:"-"`
}

// ruleSettingQuarantineJSON contains the JSON metadata for the struct
// [RuleSettingQuarantine]
type ruleSettingQuarantineJSON struct {
	FileTypes   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingQuarantine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingQuarantineJSON) RawJSON() string {
	return r.raw
}

type RuleSettingQuarantineFileType string

const (
	RuleSettingQuarantineFileTypeExe  RuleSettingQuarantineFileType = "exe"
	RuleSettingQuarantineFileTypePDF  RuleSettingQuarantineFileType = "pdf"
	RuleSettingQuarantineFileTypeDoc  RuleSettingQuarantineFileType = "doc"
	RuleSettingQuarantineFileTypeDocm RuleSettingQuarantineFileType = "docm"
	RuleSettingQuarantineFileTypeDocx RuleSettingQuarantineFileType = "docx"
	RuleSettingQuarantineFileTypeRtf  RuleSettingQuarantineFileType = "rtf"
	RuleSettingQuarantineFileTypePpt  RuleSettingQuarantineFileType = "ppt"
	RuleSettingQuarantineFileTypePptx RuleSettingQuarantineFileType = "pptx"
	RuleSettingQuarantineFileTypeXls  RuleSettingQuarantineFileType = "xls"
	RuleSettingQuarantineFileTypeXlsm RuleSettingQuarantineFileType = "xlsm"
	RuleSettingQuarantineFileTypeXlsx RuleSettingQuarantineFileType = "xlsx"
	RuleSettingQuarantineFileTypeZip  RuleSettingQuarantineFileType = "zip"
	RuleSettingQuarantineFileTypeRar  RuleSettingQuarantineFileType = "rar"
)

func (r RuleSettingQuarantineFileType) IsKnown() bool {
	switch r {
	case RuleSettingQuarantineFileTypeExe, RuleSettingQuarantineFileTypePDF, RuleSettingQuarantineFileTypeDoc, RuleSettingQuarantineFileTypeDocm, RuleSettingQuarantineFileTypeDocx, RuleSettingQuarantineFileTypeRtf, RuleSettingQuarantineFileTypePpt, RuleSettingQuarantineFileTypePptx, RuleSettingQuarantineFileTypeXls, RuleSettingQuarantineFileTypeXlsm, RuleSettingQuarantineFileTypeXlsx, RuleSettingQuarantineFileTypeZip, RuleSettingQuarantineFileTypeRar:
		return true
	}
	return false
}

// Apply settings to redirect rules. Settable only for `http` rules with the action
// set to `redirect`.
type RuleSettingRedirect struct {
	// Specify the URI to which the user is redirected.
	TargetURI string `json:"target_uri,required" format:"uri"`
	// Specify whether to pass the context information as query parameters.
	IncludeContext bool `json:"include_context"`
	// Specify whether to append the path and query parameters from the original
	// request to target_uri.
	PreservePathAndQuery bool                    `json:"preserve_path_and_query"`
	JSON                 ruleSettingRedirectJSON `json:"-"`
}

// ruleSettingRedirectJSON contains the JSON metadata for the struct
// [RuleSettingRedirect]
type ruleSettingRedirectJSON struct {
	TargetURI            apijson.Field
	IncludeContext       apijson.Field
	PreservePathAndQuery apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RuleSettingRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingRedirectJSON) RawJSON() string {
	return r.raw
}

// Configure to forward the query to the internal DNS service, passing the
// specified 'view_id' as input. Not used when 'dns_resolvers' is specified or
// 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action set to
// 'resolve'. Settable only for `dns_resolver` rules.
type RuleSettingResolveDNSInternally struct {
	// Specify the fallback behavior to apply when the internal DNS response code
	// differs from 'NOERROR' or when the response data contains only CNAME records for
	// 'A' or 'AAAA' queries.
	Fallback RuleSettingResolveDNSInternallyFallback `json:"fallback"`
	// Specify the internal DNS view identifier to pass to the internal DNS service.
	ViewID string                              `json:"view_id"`
	JSON   ruleSettingResolveDNSInternallyJSON `json:"-"`
}

// ruleSettingResolveDNSInternallyJSON contains the JSON metadata for the struct
// [RuleSettingResolveDNSInternally]
type ruleSettingResolveDNSInternallyJSON struct {
	Fallback    apijson.Field
	ViewID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingResolveDNSInternally) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingResolveDNSInternallyJSON) RawJSON() string {
	return r.raw
}

// Specify the fallback behavior to apply when the internal DNS response code
// differs from 'NOERROR' or when the response data contains only CNAME records for
// 'A' or 'AAAA' queries.
type RuleSettingResolveDNSInternallyFallback string

const (
	RuleSettingResolveDNSInternallyFallbackNone      RuleSettingResolveDNSInternallyFallback = "none"
	RuleSettingResolveDNSInternallyFallbackPublicDNS RuleSettingResolveDNSInternallyFallback = "public_dns"
)

func (r RuleSettingResolveDNSInternallyFallback) IsKnown() bool {
	switch r {
	case RuleSettingResolveDNSInternallyFallbackNone, RuleSettingResolveDNSInternallyFallbackPublicDNS:
		return true
	}
	return false
}

// Configure behavior when an upstream certificate is invalid or an SSL error
// occurs. Settable only for `http` rules with the action set to `allow`.
type RuleSettingUntrustedCERT struct {
	// Defines the action performed when an untrusted certificate seen. The default
	// action an error with HTTP code 526.
	Action RuleSettingUntrustedCERTAction `json:"action"`
	JSON   ruleSettingUntrustedCERTJSON   `json:"-"`
}

// ruleSettingUntrustedCERTJSON contains the JSON metadata for the struct
// [RuleSettingUntrustedCERT]
type ruleSettingUntrustedCERTJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingUntrustedCERT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingUntrustedCERTJSON) RawJSON() string {
	return r.raw
}

// Defines the action performed when an untrusted certificate seen. The default
// action an error with HTTP code 526.
type RuleSettingUntrustedCERTAction string

const (
	RuleSettingUntrustedCERTActionPassThrough RuleSettingUntrustedCERTAction = "pass_through"
	RuleSettingUntrustedCERTActionBlock       RuleSettingUntrustedCERTAction = "block"
	RuleSettingUntrustedCERTActionError       RuleSettingUntrustedCERTAction = "error"
)

func (r RuleSettingUntrustedCERTAction) IsKnown() bool {
	switch r {
	case RuleSettingUntrustedCERTActionPassThrough, RuleSettingUntrustedCERTActionBlock, RuleSettingUntrustedCERTActionError:
		return true
	}
	return false
}

// Defines settings for this rule. Settings apply only to specific rule types and
// must use compatible selectors. If Terraform detects drift, confirm the setting
// supports your rule type and check whether the API modifies the value. Use
// API-returned values in your configuration to prevent drift.
type RuleSettingParam struct {
	// Add custom headers to allowed requests as key-value pairs. Use header names as
	// keys that map to arrays of header values. Settable only for `http` rules with
	// the action set to `allow`.
	AddHeaders param.Field[map[string][]string] `json:"add_headers"`
	// Set to enable MSP children to bypass this rule. Only parent MSP accounts can set
	// this. this rule. Settable for all types of rules.
	AllowChildBypass param.Field[bool] `json:"allow_child_bypass"`
	// Define the settings for the Audit SSH action. Settable only for `l4` rules with
	// `audit_ssh` action.
	AuditSSH param.Field[RuleSettingAuditSSHParam] `json:"audit_ssh"`
	// Configure browser isolation behavior. Settable only for `http` rules with the
	// action set to `isolate`.
	BISOAdminControls param.Field[RuleSettingBISOAdminControlsParam] `json:"biso_admin_controls"`
	// Configure custom block page settings. If missing or null, use the account
	// settings. Settable only for `http` rules with the action set to `block`.
	BlockPage param.Field[RuleSettingBlockPageParam] `json:"block_page"`
	// Enable the custom block page. Settable only for `dns` rules with action `block`.
	BlockPageEnabled param.Field[bool] `json:"block_page_enabled"`
	// Explain why the rule blocks the request. The custom block page shows this text
	// (if enabled). Settable only for `dns`, `l4`, and `http` rules when the action
	// set to `block`.
	BlockReason param.Field[string] `json:"block_reason"`
	// Set to enable MSP accounts to bypass their parent's rules. Only MSP child
	// accounts can set this. Settable for all types of rules.
	BypassParentRule param.Field[bool] `json:"bypass_parent_rule"`
	// Configure session check behavior. Settable only for `l4` and `http` rules with
	// the action set to `allow`.
	CheckSession param.Field[RuleSettingCheckSessionParam] `json:"check_session"`
	// Configure custom resolvers to route queries that match the resolver policy.
	// Unused with 'resolve_dns_through_cloudflare' or 'resolve_dns_internally'
	// settings. DNS queries get routed to the address closest to their origin. Only
	// valid when a rule's action set to 'resolve'. Settable only for `dns_resolver`
	// rules.
	DNSResolvers param.Field[RuleSettingDNSResolversParam] `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs. Settable only for `egress` rules.
	Egress param.Field[RuleSettingEgressParam] `json:"egress"`
	// Ignore category matches at CNAME domains in a response. When off, evaluate
	// categories in this rule against all CNAME domain categories in the response.
	// Settable only for `dns` and `dns_resolver` rules.
	IgnoreCNAMECategoryMatches param.Field[bool] `json:"ignore_cname_category_matches"`
	// Specify whether to disable DNSSEC validation (for Allow actions) [INSECURE].
	// Settable only for `dns` rules.
	InsecureDisableDNSSECValidation param.Field[bool] `json:"insecure_disable_dnssec_validation"`
	// Enable IPs in DNS resolver category blocks. The system blocks only domain name
	// categories unless you enable this setting. Settable only for `dns` and
	// `dns_resolver` rules.
	IPCategories param.Field[bool] `json:"ip_categories"`
	// Indicates whether to include IPs in DNS resolver indicator feed blocks. Default,
	// indicator feeds block only domain names. Settable only for `dns` and
	// `dns_resolver` rules.
	IPIndicatorFeeds param.Field[bool] `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port. Settable
	// only for `l4` rules with the action set to `l4_override`.
	L4override param.Field[RuleSettingL4overrideParam] `json:"l4override"`
	// Configure a notification to display on the user's device when this rule matched.
	// Settable for all types of rules with the action set to `block`.
	NotificationSettings param.Field[RuleSettingNotificationSettingsParam] `json:"notification_settings"`
	// Defines a hostname for override, for the matching DNS queries. Settable only for
	// `dns` rules with the action set to `override`.
	OverrideHost param.Field[string] `json:"override_host"`
	// Defines a an IP or set of IPs for overriding matched DNS queries. Settable only
	// for `dns` rules with the action set to `override`.
	OverrideIPs param.Field[[]string] `json:"override_ips"`
	// Configure DLP payload logging. Settable only for `http` rules.
	PayloadLog param.Field[RuleSettingPayloadLogParam] `json:"payload_log"`
	// Configure settings that apply to quarantine rules. Settable only for `http`
	// rules.
	Quarantine param.Field[RuleSettingQuarantineParam] `json:"quarantine"`
	// Apply settings to redirect rules. Settable only for `http` rules with the action
	// set to `redirect`.
	Redirect param.Field[RuleSettingRedirectParam] `json:"redirect"`
	// Configure to forward the query to the internal DNS service, passing the
	// specified 'view_id' as input. Not used when 'dns_resolvers' is specified or
	// 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action set to
	// 'resolve'. Settable only for `dns_resolver` rules.
	ResolveDNSInternally param.Field[RuleSettingResolveDNSInternallyParam] `json:"resolve_dns_internally"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot set when 'dns_resolvers' specified or 'resolve_dns_internally'
	// is set. Only valid when a rule's action set to 'resolve'. Settable only for
	// `dns_resolver` rules.
	ResolveDNSThroughCloudflare param.Field[bool] `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream certificate is invalid or an SSL error
	// occurs. Settable only for `http` rules with the action set to `allow`.
	UntrustedCERT param.Field[RuleSettingUntrustedCERTParam] `json:"untrusted_cert"`
}

func (r RuleSettingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the settings for the Audit SSH action. Settable only for `l4` rules with
// `audit_ssh` action.
type RuleSettingAuditSSHParam struct {
	// Enable SSH command logging.
	CommandLogging param.Field[bool] `json:"command_logging"`
}

func (r RuleSettingAuditSSHParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure browser isolation behavior. Settable only for `http` rules with the
// action set to `isolate`.
type RuleSettingBISOAdminControlsParam struct {
	// Configure copy behavior. If set to remote_only, users cannot copy isolated
	// content from the remote browser to the local clipboard. If this field is absent,
	// copying remains enabled. Applies only when version == "v2".
	Copy param.Field[RuleSettingBISOAdminControlsCopy] `json:"copy"`
	// Set to false to enable copy-pasting. Only applies when `version == "v1"`.
	DCP param.Field[bool] `json:"dcp"`
	// Set to false to enable downloading. Only applies when `version == "v1"`.
	DD param.Field[bool] `json:"dd"`
	// Set to false to enable keyboard usage. Only applies when `version == "v1"`.
	DK param.Field[bool] `json:"dk"`
	// Configure download behavior. When set to remote_only, users can view downloads
	// but cannot save them. Applies only when version == "v2".
	Download param.Field[RuleSettingBISOAdminControlsDownload] `json:"download"`
	// Set to false to enable printing. Only applies when `version == "v1"`.
	DP param.Field[bool] `json:"dp"`
	// Set to false to enable uploading. Only applies when `version == "v1"`.
	DU param.Field[bool] `json:"du"`
	// Configure keyboard usage behavior. If this field is absent, keyboard usage
	// remains enabled. Applies only when version == "v2".
	Keyboard param.Field[RuleSettingBISOAdminControlsKeyboard] `json:"keyboard"`
	// Configure paste behavior. If set to remote_only, users cannot paste content from
	// the local clipboard into isolated pages. If this field is absent, pasting
	// remains enabled. Applies only when version == "v2".
	Paste param.Field[RuleSettingBISOAdminControlsPaste] `json:"paste"`
	// Configure print behavior. Default, Printing is enabled. Applies only when
	// version == "v2".
	Printing param.Field[RuleSettingBISOAdminControlsPrinting] `json:"printing"`
	// Configure upload behavior. If this field is absent, uploading remains enabled.
	// Applies only when version == "v2".
	Upload param.Field[RuleSettingBISOAdminControlsUpload] `json:"upload"`
	// Indicate which version of the browser isolation controls should apply.
	Version param.Field[RuleSettingBISOAdminControlsVersion] `json:"version"`
}

func (r RuleSettingBISOAdminControlsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure custom block page settings. If missing or null, use the account
// settings. Settable only for `http` rules with the action set to `block`.
type RuleSettingBlockPageParam struct {
	// Specify the URI to which the user is redirected.
	TargetURI param.Field[string] `json:"target_uri,required" format:"uri"`
	// Specify whether to pass the context information as query parameters.
	IncludeContext param.Field[bool] `json:"include_context"`
}

func (r RuleSettingBlockPageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure session check behavior. Settable only for `l4` and `http` rules with
// the action set to `allow`.
type RuleSettingCheckSessionParam struct {
	// Sets the required session freshness threshold. The API returns a normalized
	// version of this value.
	Duration param.Field[string] `json:"duration"`
	// Enable session enforcement.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r RuleSettingCheckSessionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure custom resolvers to route queries that match the resolver policy.
// Unused with 'resolve_dns_through_cloudflare' or 'resolve_dns_internally'
// settings. DNS queries get routed to the address closest to their origin. Only
// valid when a rule's action set to 'resolve'. Settable only for `dns_resolver`
// rules.
type RuleSettingDNSResolversParam struct {
	IPV4 param.Field[[]DNSResolverSettingsV4Param] `json:"ipv4"`
	IPV6 param.Field[[]DNSResolverSettingsV6Param] `json:"ipv6"`
}

func (r RuleSettingDNSResolversParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs. Settable only for `egress` rules.
type RuleSettingEgressParam struct {
	// Specify the IPv4 address to use for egress.
	IPV4 param.Field[string] `json:"ipv4"`
	// Specify the fallback IPv4 address to use for egress when the primary IPv4 fails.
	// Set '0.0.0.0' to indicate local egress via WARP IPs.
	IPV4Fallback param.Field[string] `json:"ipv4_fallback"`
	// Specify the IPv6 range to use for egress.
	IPV6 param.Field[string] `json:"ipv6"`
}

func (r RuleSettingEgressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Send matching traffic to the supplied destination IP address and port. Settable
// only for `l4` rules with the action set to `l4_override`.
type RuleSettingL4overrideParam struct {
	// Defines the IPv4 or IPv6 address.
	IP param.Field[string] `json:"ip"`
	// Defines a port number to use for TCP/UDP overrides.
	Port param.Field[int64] `json:"port"`
}

func (r RuleSettingL4overrideParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a notification to display on the user's device when this rule matched.
// Settable for all types of rules with the action set to `block`.
type RuleSettingNotificationSettingsParam struct {
	// Enable notification.
	Enabled param.Field[bool] `json:"enabled"`
	// Indicates whether to pass the context information as query parameters.
	IncludeContext param.Field[bool] `json:"include_context"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Defines an optional URL to direct users to additional information. If unset, the
	// notification opens a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r RuleSettingNotificationSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure DLP payload logging. Settable only for `http` rules.
type RuleSettingPayloadLogParam struct {
	// Enable DLP payload logging for this rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r RuleSettingPayloadLogParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure settings that apply to quarantine rules. Settable only for `http`
// rules.
type RuleSettingQuarantineParam struct {
	// Specify the types of files to sandbox.
	FileTypes param.Field[[]RuleSettingQuarantineFileType] `json:"file_types"`
}

func (r RuleSettingQuarantineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Apply settings to redirect rules. Settable only for `http` rules with the action
// set to `redirect`.
type RuleSettingRedirectParam struct {
	// Specify the URI to which the user is redirected.
	TargetURI param.Field[string] `json:"target_uri,required" format:"uri"`
	// Specify whether to pass the context information as query parameters.
	IncludeContext param.Field[bool] `json:"include_context"`
	// Specify whether to append the path and query parameters from the original
	// request to target_uri.
	PreservePathAndQuery param.Field[bool] `json:"preserve_path_and_query"`
}

func (r RuleSettingRedirectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure to forward the query to the internal DNS service, passing the
// specified 'view_id' as input. Not used when 'dns_resolvers' is specified or
// 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action set to
// 'resolve'. Settable only for `dns_resolver` rules.
type RuleSettingResolveDNSInternallyParam struct {
	// Specify the fallback behavior to apply when the internal DNS response code
	// differs from 'NOERROR' or when the response data contains only CNAME records for
	// 'A' or 'AAAA' queries.
	Fallback param.Field[RuleSettingResolveDNSInternallyFallback] `json:"fallback"`
	// Specify the internal DNS view identifier to pass to the internal DNS service.
	ViewID param.Field[string] `json:"view_id"`
}

func (r RuleSettingResolveDNSInternallyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure behavior when an upstream certificate is invalid or an SSL error
// occurs. Settable only for `http` rules with the action set to `allow`.
type RuleSettingUntrustedCERTParam struct {
	// Defines the action performed when an untrusted certificate seen. The default
	// action an error with HTTP code 526.
	Action param.Field[RuleSettingUntrustedCERTAction] `json:"action"`
}

func (r RuleSettingUntrustedCERTParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines the schedule for activating DNS policies. Settable only for `dns` and
// `dns_resolver` rules.
type Schedule struct {
	// Specify the time intervals when the rule is active on Fridays, in the increasing
	// order from 00:00-24:00. If this parameter omitted, the rule is deactivated on
	// Fridays. API returns a formatted version of this string, which may cause
	// Terraform drift if a unformatted value is used.
	Fri string `json:"fri"`
	// Specify the time intervals when the rule is active on Mondays, in the increasing
	// order from 00:00-24:00(capped at maximum of 6 time splits). If this parameter
	// omitted, the rule is deactivated on Mondays. API returns a formatted version of
	// this string, which may cause Terraform drift if a unformatted value is used.
	Mon string `json:"mon"`
	// Specify the time intervals when the rule is active on Saturdays, in the
	// increasing order from 00:00-24:00. If this parameter omitted, the rule is
	// deactivated on Saturdays. API returns a formatted version of this string, which
	// may cause Terraform drift if a unformatted value is used.
	Sat string `json:"sat"`
	// Specify the time intervals when the rule is active on Sundays, in the increasing
	// order from 00:00-24:00. If this parameter omitted, the rule is deactivated on
	// Sundays. API returns a formatted version of this string, which may cause
	// Terraform drift if a unformatted value is used.
	Sun string `json:"sun"`
	// Specify the time intervals when the rule is active on Thursdays, in the
	// increasing order from 00:00-24:00. If this parameter omitted, the rule is
	// deactivated on Thursdays. API returns a formatted version of this string, which
	// may cause Terraform drift if a unformatted value is used.
	Thu string `json:"thu"`
	// Specify the time zone for rule evaluation. When a
	// [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)
	// is provided, Gateway always uses the current time for that time zone. When this
	// parameter is omitted, Gateway uses the time zone determined from the user's IP
	// address. Colo time zone is used when the user's IP address does not resolve to a
	// location.
	TimeZone string `json:"time_zone"`
	// Specify the time intervals when the rule is active on Tuesdays, in the
	// increasing order from 00:00-24:00. If this parameter omitted, the rule is
	// deactivated on Tuesdays. API returns a formatted version of this string, which
	// may cause Terraform drift if a unformatted value is used.
	Tue string `json:"tue"`
	// Specify the time intervals when the rule is active on Wednesdays, in the
	// increasing order from 00:00-24:00. If this parameter omitted, the rule is
	// deactivated on Wednesdays. API returns a formatted version of this string, which
	// may cause Terraform drift if a unformatted value is used.
	Wed  string       `json:"wed"`
	JSON scheduleJSON `json:"-"`
}

// scheduleJSON contains the JSON metadata for the struct [Schedule]
type scheduleJSON struct {
	Fri         apijson.Field
	Mon         apijson.Field
	Sat         apijson.Field
	Sun         apijson.Field
	Thu         apijson.Field
	TimeZone    apijson.Field
	Tue         apijson.Field
	Wed         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Schedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleJSON) RawJSON() string {
	return r.raw
}

// Defines the schedule for activating DNS policies. Settable only for `dns` and
// `dns_resolver` rules.
type ScheduleParam struct {
	// Specify the time intervals when the rule is active on Fridays, in the increasing
	// order from 00:00-24:00. If this parameter omitted, the rule is deactivated on
	// Fridays. API returns a formatted version of this string, which may cause
	// Terraform drift if a unformatted value is used.
	Fri param.Field[string] `json:"fri"`
	// Specify the time intervals when the rule is active on Mondays, in the increasing
	// order from 00:00-24:00(capped at maximum of 6 time splits). If this parameter
	// omitted, the rule is deactivated on Mondays. API returns a formatted version of
	// this string, which may cause Terraform drift if a unformatted value is used.
	Mon param.Field[string] `json:"mon"`
	// Specify the time intervals when the rule is active on Saturdays, in the
	// increasing order from 00:00-24:00. If this parameter omitted, the rule is
	// deactivated on Saturdays. API returns a formatted version of this string, which
	// may cause Terraform drift if a unformatted value is used.
	Sat param.Field[string] `json:"sat"`
	// Specify the time intervals when the rule is active on Sundays, in the increasing
	// order from 00:00-24:00. If this parameter omitted, the rule is deactivated on
	// Sundays. API returns a formatted version of this string, which may cause
	// Terraform drift if a unformatted value is used.
	Sun param.Field[string] `json:"sun"`
	// Specify the time intervals when the rule is active on Thursdays, in the
	// increasing order from 00:00-24:00. If this parameter omitted, the rule is
	// deactivated on Thursdays. API returns a formatted version of this string, which
	// may cause Terraform drift if a unformatted value is used.
	Thu param.Field[string] `json:"thu"`
	// Specify the time zone for rule evaluation. When a
	// [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)
	// is provided, Gateway always uses the current time for that time zone. When this
	// parameter is omitted, Gateway uses the time zone determined from the user's IP
	// address. Colo time zone is used when the user's IP address does not resolve to a
	// location.
	TimeZone param.Field[string] `json:"time_zone"`
	// Specify the time intervals when the rule is active on Tuesdays, in the
	// increasing order from 00:00-24:00. If this parameter omitted, the rule is
	// deactivated on Tuesdays. API returns a formatted version of this string, which
	// may cause Terraform drift if a unformatted value is used.
	Tue param.Field[string] `json:"tue"`
	// Specify the time intervals when the rule is active on Wednesdays, in the
	// increasing order from 00:00-24:00. If this parameter omitted, the rule is
	// deactivated on Wednesdays. API returns a formatted version of this string, which
	// may cause Terraform drift if a unformatted value is used.
	Wed param.Field[string] `json:"wed"`
}

func (r ScheduleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleDeleteResponse = interface{}

type GatewayRuleNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Specify the action to perform when the associated traffic, identity, and device
	// posture expressions either absent or evaluate to `true`.
	Action param.Field[GatewayRuleNewParamsAction] `json:"action,required"`
	// Specify the rule name.
	Name param.Field[string] `json:"name,required"`
	// Specify the rule description.
	Description param.Field[string] `json:"description"`
	// Specify the wirefilter expression used for device posture check. The API
	// automatically formats and sanitizes expressions before storing them. To prevent
	// Terraform state drift, use the formatted expression returned in the API
	// response.
	DevicePosture param.Field[string] `json:"device_posture"`
	// Specify whether the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Defines the expiration time stamp and default duration of a DNS policy. Takes
	// precedence over the policy's `schedule` configuration, if any. This does not
	// apply to HTTP or network policies. Settable only for `dns` rules.
	Expiration param.Field[GatewayRuleNewParamsExpiration] `json:"expiration"`
	// Specify the protocol or layer to evaluate the traffic, identity, and device
	// posture expressions. Can only contain a single value.
	Filters param.Field[[]GatewayFilter] `json:"filters"`
	// Specify the wirefilter expression used for identity matching. The API
	// automatically formats and sanitizes expressions before storing them. To prevent
	// Terraform state drift, use the formatted expression returned in the API
	// response.
	Identity param.Field[string] `json:"identity"`
	// Set the order of your rules. Lower values indicate higher precedence. At each
	// processing phase, evaluate applicable rules in ascending order of this value.
	// Refer to
	// [Order of enforcement](http://developers.cloudflare.com/learning-paths/secure-internet-traffic/understand-policies/order-of-enforcement/#manage-precedence-with-terraform)
	// to manage precedence via Terraform.
	Precedence param.Field[int64] `json:"precedence"`
	// Defines settings for this rule. Settings apply only to specific rule types and
	// must use compatible selectors. If Terraform detects drift, confirm the setting
	// supports your rule type and check whether the API modifies the value. Use
	// API-returned values in your configuration to prevent drift.
	RuleSettings param.Field[RuleSettingParam] `json:"rule_settings"`
	// Defines the schedule for activating DNS policies. Settable only for `dns` and
	// `dns_resolver` rules.
	Schedule param.Field[ScheduleParam] `json:"schedule"`
	// Specify the wirefilter expression used for traffic matching. The API
	// automatically formats and sanitizes expressions before storing them. To prevent
	// Terraform state drift, use the formatted expression returned in the API
	// response.
	Traffic param.Field[string] `json:"traffic"`
}

func (r GatewayRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify the action to perform when the associated traffic, identity, and device
// posture expressions either absent or evaluate to `true`.
type GatewayRuleNewParamsAction string

const (
	GatewayRuleNewParamsActionOn           GatewayRuleNewParamsAction = "on"
	GatewayRuleNewParamsActionOff          GatewayRuleNewParamsAction = "off"
	GatewayRuleNewParamsActionAllow        GatewayRuleNewParamsAction = "allow"
	GatewayRuleNewParamsActionBlock        GatewayRuleNewParamsAction = "block"
	GatewayRuleNewParamsActionScan         GatewayRuleNewParamsAction = "scan"
	GatewayRuleNewParamsActionNoscan       GatewayRuleNewParamsAction = "noscan"
	GatewayRuleNewParamsActionSafesearch   GatewayRuleNewParamsAction = "safesearch"
	GatewayRuleNewParamsActionYtrestricted GatewayRuleNewParamsAction = "ytrestricted"
	GatewayRuleNewParamsActionIsolate      GatewayRuleNewParamsAction = "isolate"
	GatewayRuleNewParamsActionNoisolate    GatewayRuleNewParamsAction = "noisolate"
	GatewayRuleNewParamsActionOverride     GatewayRuleNewParamsAction = "override"
	GatewayRuleNewParamsActionL4Override   GatewayRuleNewParamsAction = "l4_override"
	GatewayRuleNewParamsActionEgress       GatewayRuleNewParamsAction = "egress"
	GatewayRuleNewParamsActionResolve      GatewayRuleNewParamsAction = "resolve"
	GatewayRuleNewParamsActionQuarantine   GatewayRuleNewParamsAction = "quarantine"
	GatewayRuleNewParamsActionRedirect     GatewayRuleNewParamsAction = "redirect"
)

func (r GatewayRuleNewParamsAction) IsKnown() bool {
	switch r {
	case GatewayRuleNewParamsActionOn, GatewayRuleNewParamsActionOff, GatewayRuleNewParamsActionAllow, GatewayRuleNewParamsActionBlock, GatewayRuleNewParamsActionScan, GatewayRuleNewParamsActionNoscan, GatewayRuleNewParamsActionSafesearch, GatewayRuleNewParamsActionYtrestricted, GatewayRuleNewParamsActionIsolate, GatewayRuleNewParamsActionNoisolate, GatewayRuleNewParamsActionOverride, GatewayRuleNewParamsActionL4Override, GatewayRuleNewParamsActionEgress, GatewayRuleNewParamsActionResolve, GatewayRuleNewParamsActionQuarantine, GatewayRuleNewParamsActionRedirect:
		return true
	}
	return false
}

// Defines the expiration time stamp and default duration of a DNS policy. Takes
// precedence over the policy's `schedule` configuration, if any. This does not
// apply to HTTP or network policies. Settable only for `dns` rules.
type GatewayRuleNewParamsExpiration struct {
	// Show the timestamp when the policy expires and stops applying. The value must
	// follow RFC 3339 and include a UTC offset. The system accepts non-zero offsets
	// but converts them to the equivalent UTC+00:00 value and returns timestamps with
	// a trailing Z. Expiration policies ignore client timezones and expire globally at
	// the specified expires_at time.
	ExpiresAt param.Field[time.Time] `json:"expires_at,required" format:"date-time"`
	// Defines the default duration a policy active in minutes. Must set in order to
	// use the `reset_expiration` endpoint on this rule.
	Duration param.Field[int64] `json:"duration"`
}

func (r GatewayRuleNewParamsExpiration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayRuleNewResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayRule                           `json:"result"`
	JSON    gatewayRuleNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleNewResponseEnvelope]
type gatewayRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
type GatewayRuleNewResponseEnvelopeSuccess bool

const (
	GatewayRuleNewResponseEnvelopeSuccessTrue GatewayRuleNewResponseEnvelopeSuccess = true
)

func (r GatewayRuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayRuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayRuleUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Specify the action to perform when the associated traffic, identity, and device
	// posture expressions either absent or evaluate to `true`.
	Action param.Field[GatewayRuleUpdateParamsAction] `json:"action,required"`
	// Specify the rule name.
	Name param.Field[string] `json:"name,required"`
	// Specify the rule description.
	Description param.Field[string] `json:"description"`
	// Specify the wirefilter expression used for device posture check. The API
	// automatically formats and sanitizes expressions before storing them. To prevent
	// Terraform state drift, use the formatted expression returned in the API
	// response.
	DevicePosture param.Field[string] `json:"device_posture"`
	// Specify whether the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// Defines the expiration time stamp and default duration of a DNS policy. Takes
	// precedence over the policy's `schedule` configuration, if any. This does not
	// apply to HTTP or network policies. Settable only for `dns` rules.
	Expiration param.Field[GatewayRuleUpdateParamsExpiration] `json:"expiration"`
	// Specify the protocol or layer to evaluate the traffic, identity, and device
	// posture expressions. Can only contain a single value.
	Filters param.Field[[]GatewayFilter] `json:"filters"`
	// Specify the wirefilter expression used for identity matching. The API
	// automatically formats and sanitizes expressions before storing them. To prevent
	// Terraform state drift, use the formatted expression returned in the API
	// response.
	Identity param.Field[string] `json:"identity"`
	// Set the order of your rules. Lower values indicate higher precedence. At each
	// processing phase, evaluate applicable rules in ascending order of this value.
	// Refer to
	// [Order of enforcement](http://developers.cloudflare.com/learning-paths/secure-internet-traffic/understand-policies/order-of-enforcement/#manage-precedence-with-terraform)
	// to manage precedence via Terraform.
	Precedence param.Field[int64] `json:"precedence"`
	// Defines settings for this rule. Settings apply only to specific rule types and
	// must use compatible selectors. If Terraform detects drift, confirm the setting
	// supports your rule type and check whether the API modifies the value. Use
	// API-returned values in your configuration to prevent drift.
	RuleSettings param.Field[RuleSettingParam] `json:"rule_settings"`
	// Defines the schedule for activating DNS policies. Settable only for `dns` and
	// `dns_resolver` rules.
	Schedule param.Field[ScheduleParam] `json:"schedule"`
	// Specify the wirefilter expression used for traffic matching. The API
	// automatically formats and sanitizes expressions before storing them. To prevent
	// Terraform state drift, use the formatted expression returned in the API
	// response.
	Traffic param.Field[string] `json:"traffic"`
}

func (r GatewayRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify the action to perform when the associated traffic, identity, and device
// posture expressions either absent or evaluate to `true`.
type GatewayRuleUpdateParamsAction string

const (
	GatewayRuleUpdateParamsActionOn           GatewayRuleUpdateParamsAction = "on"
	GatewayRuleUpdateParamsActionOff          GatewayRuleUpdateParamsAction = "off"
	GatewayRuleUpdateParamsActionAllow        GatewayRuleUpdateParamsAction = "allow"
	GatewayRuleUpdateParamsActionBlock        GatewayRuleUpdateParamsAction = "block"
	GatewayRuleUpdateParamsActionScan         GatewayRuleUpdateParamsAction = "scan"
	GatewayRuleUpdateParamsActionNoscan       GatewayRuleUpdateParamsAction = "noscan"
	GatewayRuleUpdateParamsActionSafesearch   GatewayRuleUpdateParamsAction = "safesearch"
	GatewayRuleUpdateParamsActionYtrestricted GatewayRuleUpdateParamsAction = "ytrestricted"
	GatewayRuleUpdateParamsActionIsolate      GatewayRuleUpdateParamsAction = "isolate"
	GatewayRuleUpdateParamsActionNoisolate    GatewayRuleUpdateParamsAction = "noisolate"
	GatewayRuleUpdateParamsActionOverride     GatewayRuleUpdateParamsAction = "override"
	GatewayRuleUpdateParamsActionL4Override   GatewayRuleUpdateParamsAction = "l4_override"
	GatewayRuleUpdateParamsActionEgress       GatewayRuleUpdateParamsAction = "egress"
	GatewayRuleUpdateParamsActionResolve      GatewayRuleUpdateParamsAction = "resolve"
	GatewayRuleUpdateParamsActionQuarantine   GatewayRuleUpdateParamsAction = "quarantine"
	GatewayRuleUpdateParamsActionRedirect     GatewayRuleUpdateParamsAction = "redirect"
)

func (r GatewayRuleUpdateParamsAction) IsKnown() bool {
	switch r {
	case GatewayRuleUpdateParamsActionOn, GatewayRuleUpdateParamsActionOff, GatewayRuleUpdateParamsActionAllow, GatewayRuleUpdateParamsActionBlock, GatewayRuleUpdateParamsActionScan, GatewayRuleUpdateParamsActionNoscan, GatewayRuleUpdateParamsActionSafesearch, GatewayRuleUpdateParamsActionYtrestricted, GatewayRuleUpdateParamsActionIsolate, GatewayRuleUpdateParamsActionNoisolate, GatewayRuleUpdateParamsActionOverride, GatewayRuleUpdateParamsActionL4Override, GatewayRuleUpdateParamsActionEgress, GatewayRuleUpdateParamsActionResolve, GatewayRuleUpdateParamsActionQuarantine, GatewayRuleUpdateParamsActionRedirect:
		return true
	}
	return false
}

// Defines the expiration time stamp and default duration of a DNS policy. Takes
// precedence over the policy's `schedule` configuration, if any. This does not
// apply to HTTP or network policies. Settable only for `dns` rules.
type GatewayRuleUpdateParamsExpiration struct {
	// Show the timestamp when the policy expires and stops applying. The value must
	// follow RFC 3339 and include a UTC offset. The system accepts non-zero offsets
	// but converts them to the equivalent UTC+00:00 value and returns timestamps with
	// a trailing Z. Expiration policies ignore client timezones and expire globally at
	// the specified expires_at time.
	ExpiresAt param.Field[time.Time] `json:"expires_at,required" format:"date-time"`
	// Defines the default duration a policy active in minutes. Must set in order to
	// use the `reset_expiration` endpoint on this rule.
	Duration param.Field[int64] `json:"duration"`
}

func (r GatewayRuleUpdateParamsExpiration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayRule                              `json:"result"`
	JSON    gatewayRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleUpdateResponseEnvelope]
type gatewayRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
type GatewayRuleUpdateResponseEnvelopeSuccess bool

const (
	GatewayRuleUpdateResponseEnvelopeSuccessTrue GatewayRuleUpdateResponseEnvelopeSuccess = true
)

func (r GatewayRuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayRuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayRuleListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayRuleDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayRuleDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayRuleDeleteResponse                `json:"result"`
	JSON    gatewayRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleDeleteResponseEnvelope]
type gatewayRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
type GatewayRuleDeleteResponseEnvelopeSuccess bool

const (
	GatewayRuleDeleteResponseEnvelopeSuccessTrue GatewayRuleDeleteResponseEnvelopeSuccess = true
)

func (r GatewayRuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayRuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayRuleGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayRuleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayRuleGetResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayRule                           `json:"result"`
	JSON    gatewayRuleGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleGetResponseEnvelope]
type gatewayRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
type GatewayRuleGetResponseEnvelopeSuccess bool

const (
	GatewayRuleGetResponseEnvelopeSuccessTrue GatewayRuleGetResponseEnvelopeSuccess = true
)

func (r GatewayRuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayRuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayRuleListTenantParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayRuleResetExpirationParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayRuleResetExpirationResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicate whether the API call was successful.
	Success GatewayRuleResetExpirationResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayRule                                       `json:"result"`
	JSON    gatewayRuleResetExpirationResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleResetExpirationResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayRuleResetExpirationResponseEnvelope]
type gatewayRuleResetExpirationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleResetExpirationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleResetExpirationResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Indicate whether the API call was successful.
type GatewayRuleResetExpirationResponseEnvelopeSuccess bool

const (
	GatewayRuleResetExpirationResponseEnvelopeSuccessTrue GatewayRuleResetExpirationResponseEnvelopeSuccess = true
)

func (r GatewayRuleResetExpirationResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayRuleResetExpirationResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

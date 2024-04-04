// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// GatewayRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewGatewayRuleService] method
// instead.
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

// Creates a new Zero Trust Gateway rule.
func (r *GatewayRuleService) New(ctx context.Context, params GatewayRuleNewParams, opts ...option.RequestOption) (res *ZeroTrustGatewayRules, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/rules", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust Gateway rule.
func (r *GatewayRuleService) Update(ctx context.Context, ruleID string, params GatewayRuleUpdateParams, opts ...option.RequestOption) (res *ZeroTrustGatewayRules, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleUpdateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the Zero Trust Gateway rules for an account.
func (r *GatewayRuleService) List(ctx context.Context, query GatewayRuleListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ZeroTrustGatewayRules], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/gateway/rules", query.AccountID)
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

// Fetches the Zero Trust Gateway rules for an account.
func (r *GatewayRuleService) ListAutoPaging(ctx context.Context, query GatewayRuleListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ZeroTrustGatewayRules] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a Zero Trust Gateway rule.
func (r *GatewayRuleService) Delete(ctx context.Context, ruleID string, params GatewayRuleDeleteParams, opts ...option.RequestOption) (res *GatewayRuleDeleteResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust Gateway rule.
func (r *GatewayRuleService) Get(ctx context.Context, ruleID string, query GatewayRuleGetParams, opts ...option.RequestOption) (res *ZeroTrustGatewayRules, err error) {
	opts = append(r.Options[:], opts...)
	var env GatewayRuleGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustGatewayRules struct {
	// The API resource UUID.
	ID string `json:"id"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action    ZeroTrustGatewayRulesAction `json:"action"`
	CreatedAt time.Time                   `json:"created_at" format:"date-time"`
	// Date of deletion, if any.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// The description of the rule.
	Description string `json:"description"`
	// The wirefilter expression used for device posture check matching.
	DevicePosture string `json:"device_posture"`
	// True if the rule is enabled.
	Enabled bool `json:"enabled"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters []ZeroTrustGatewayRulesFilter `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity string `json:"identity"`
	// The name of the rule.
	Name string `json:"name"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence int64 `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings ZeroTrustGatewayRulesRuleSettings `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule ZeroTrustGatewayRulesSchedule `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic   string                    `json:"traffic"`
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	JSON      zeroTrustGatewayRulesJSON `json:"-"`
}

// zeroTrustGatewayRulesJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayRules]
type zeroTrustGatewayRulesJSON struct {
	ID            apijson.Field
	Action        apijson.Field
	CreatedAt     apijson.Field
	DeletedAt     apijson.Field
	Description   apijson.Field
	DevicePosture apijson.Field
	Enabled       apijson.Field
	Filters       apijson.Field
	Identity      apijson.Field
	Name          apijson.Field
	Precedence    apijson.Field
	RuleSettings  apijson.Field
	Schedule      apijson.Field
	Traffic       apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ZeroTrustGatewayRules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesJSON) RawJSON() string {
	return r.raw
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type ZeroTrustGatewayRulesAction string

const (
	ZeroTrustGatewayRulesActionOn           ZeroTrustGatewayRulesAction = "on"
	ZeroTrustGatewayRulesActionOff          ZeroTrustGatewayRulesAction = "off"
	ZeroTrustGatewayRulesActionAllow        ZeroTrustGatewayRulesAction = "allow"
	ZeroTrustGatewayRulesActionBlock        ZeroTrustGatewayRulesAction = "block"
	ZeroTrustGatewayRulesActionScan         ZeroTrustGatewayRulesAction = "scan"
	ZeroTrustGatewayRulesActionNoscan       ZeroTrustGatewayRulesAction = "noscan"
	ZeroTrustGatewayRulesActionSafesearch   ZeroTrustGatewayRulesAction = "safesearch"
	ZeroTrustGatewayRulesActionYtrestricted ZeroTrustGatewayRulesAction = "ytrestricted"
	ZeroTrustGatewayRulesActionIsolate      ZeroTrustGatewayRulesAction = "isolate"
	ZeroTrustGatewayRulesActionNoisolate    ZeroTrustGatewayRulesAction = "noisolate"
	ZeroTrustGatewayRulesActionOverride     ZeroTrustGatewayRulesAction = "override"
	ZeroTrustGatewayRulesActionL4Override   ZeroTrustGatewayRulesAction = "l4_override"
	ZeroTrustGatewayRulesActionEgress       ZeroTrustGatewayRulesAction = "egress"
	ZeroTrustGatewayRulesActionAuditSSH     ZeroTrustGatewayRulesAction = "audit_ssh"
	ZeroTrustGatewayRulesActionResolve      ZeroTrustGatewayRulesAction = "resolve"
)

func (r ZeroTrustGatewayRulesAction) IsKnown() bool {
	switch r {
	case ZeroTrustGatewayRulesActionOn, ZeroTrustGatewayRulesActionOff, ZeroTrustGatewayRulesActionAllow, ZeroTrustGatewayRulesActionBlock, ZeroTrustGatewayRulesActionScan, ZeroTrustGatewayRulesActionNoscan, ZeroTrustGatewayRulesActionSafesearch, ZeroTrustGatewayRulesActionYtrestricted, ZeroTrustGatewayRulesActionIsolate, ZeroTrustGatewayRulesActionNoisolate, ZeroTrustGatewayRulesActionOverride, ZeroTrustGatewayRulesActionL4Override, ZeroTrustGatewayRulesActionEgress, ZeroTrustGatewayRulesActionAuditSSH, ZeroTrustGatewayRulesActionResolve:
		return true
	}
	return false
}

// The protocol or layer to use.
type ZeroTrustGatewayRulesFilter string

const (
	ZeroTrustGatewayRulesFilterHTTP   ZeroTrustGatewayRulesFilter = "http"
	ZeroTrustGatewayRulesFilterDNS    ZeroTrustGatewayRulesFilter = "dns"
	ZeroTrustGatewayRulesFilterL4     ZeroTrustGatewayRulesFilter = "l4"
	ZeroTrustGatewayRulesFilterEgress ZeroTrustGatewayRulesFilter = "egress"
)

func (r ZeroTrustGatewayRulesFilter) IsKnown() bool {
	switch r {
	case ZeroTrustGatewayRulesFilterHTTP, ZeroTrustGatewayRulesFilterDNS, ZeroTrustGatewayRulesFilterL4, ZeroTrustGatewayRulesFilterEgress:
		return true
	}
	return false
}

// Additional settings that modify the rule's action.
type ZeroTrustGatewayRulesRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders interface{} `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH ZeroTrustGatewayRulesRuleSettingsAuditSSH `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls ZeroTrustGatewayRulesRuleSettingsBisoAdminControls `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason string `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession ZeroTrustGatewayRulesRuleSettingsCheckSession `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin. Only valid when a rule's action is
	// set to 'resolve'.
	DNSResolvers ZeroTrustGatewayRulesRuleSettingsDNSResolvers `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress ZeroTrustGatewayRulesRuleSettingsEgress `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation bool `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories bool `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override ZeroTrustGatewayRulesRuleSettingsL4override `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings ZeroTrustGatewayRulesRuleSettingsNotificationSettings `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs []string `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog ZeroTrustGatewayRulesRuleSettingsPayloadLog `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified. Only valid when a
	// rule's action is set to 'resolve'.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCERT ZeroTrustGatewayRulesRuleSettingsUntrustedCERT `json:"untrusted_cert"`
	JSON          zeroTrustGatewayRulesRuleSettingsJSON          `json:"-"`
}

// zeroTrustGatewayRulesRuleSettingsJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayRulesRuleSettings]
type zeroTrustGatewayRulesRuleSettingsJSON struct {
	AddHeaders                      apijson.Field
	AllowChildBypass                apijson.Field
	AuditSSH                        apijson.Field
	BisoAdminControls               apijson.Field
	BlockPageEnabled                apijson.Field
	BlockReason                     apijson.Field
	BypassParentRule                apijson.Field
	CheckSession                    apijson.Field
	DNSResolvers                    apijson.Field
	Egress                          apijson.Field
	InsecureDisableDNSSECValidation apijson.Field
	IPCategories                    apijson.Field
	IPIndicatorFeeds                apijson.Field
	L4override                      apijson.Field
	NotificationSettings            apijson.Field
	OverrideHost                    apijson.Field
	OverrideIPs                     apijson.Field
	PayloadLog                      apijson.Field
	ResolveDNSThroughCloudflare     apijson.Field
	UntrustedCERT                   apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *ZeroTrustGatewayRulesRuleSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesRuleSettingsJSON) RawJSON() string {
	return r.raw
}

// Settings for the Audit SSH action.
type ZeroTrustGatewayRulesRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging bool                                          `json:"command_logging"`
	JSON           zeroTrustGatewayRulesRuleSettingsAuditSSHJSON `json:"-"`
}

// zeroTrustGatewayRulesRuleSettingsAuditSSHJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayRulesRuleSettingsAuditSSH]
type zeroTrustGatewayRulesRuleSettingsAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ZeroTrustGatewayRulesRuleSettingsAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesRuleSettingsAuditSSHJSON) RawJSON() string {
	return r.raw
}

// Configure how browser isolation behaves.
type ZeroTrustGatewayRulesRuleSettingsBisoAdminControls struct {
	// Set to true to enable copy-pasting.
	Dcp bool `json:"dcp"`
	// Set to true to enable downloading.
	Dd bool `json:"dd"`
	// Set to true to enable keyboard usage.
	Dk bool `json:"dk"`
	// Set to true to enable printing.
	Dp bool `json:"dp"`
	// Set to true to enable uploading.
	Du   bool                                                   `json:"du"`
	JSON zeroTrustGatewayRulesRuleSettingsBisoAdminControlsJSON `json:"-"`
}

// zeroTrustGatewayRulesRuleSettingsBisoAdminControlsJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayRulesRuleSettingsBisoAdminControls]
type zeroTrustGatewayRulesRuleSettingsBisoAdminControlsJSON struct {
	Dcp         apijson.Field
	Dd          apijson.Field
	Dk          apijson.Field
	Dp          apijson.Field
	Du          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRulesRuleSettingsBisoAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesRuleSettingsBisoAdminControlsJSON) RawJSON() string {
	return r.raw
}

// Configure how session check behaves.
type ZeroTrustGatewayRulesRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration string `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce bool                                              `json:"enforce"`
	JSON    zeroTrustGatewayRulesRuleSettingsCheckSessionJSON `json:"-"`
}

// zeroTrustGatewayRulesRuleSettingsCheckSessionJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayRulesRuleSettingsCheckSession]
type zeroTrustGatewayRulesRuleSettingsCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRulesRuleSettingsCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesRuleSettingsCheckSessionJSON) RawJSON() string {
	return r.raw
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin. Only valid when a rule's action is
// set to 'resolve'.
type ZeroTrustGatewayRulesRuleSettingsDNSResolvers struct {
	IPV4 []ZeroTrustGatewayRulesRuleSettingsDNSResolversIPV4 `json:"ipv4"`
	IPV6 []ZeroTrustGatewayRulesRuleSettingsDNSResolversIPV6 `json:"ipv6"`
	JSON zeroTrustGatewayRulesRuleSettingsDNSResolversJSON   `json:"-"`
}

// zeroTrustGatewayRulesRuleSettingsDNSResolversJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayRulesRuleSettingsDNSResolvers]
type zeroTrustGatewayRulesRuleSettingsDNSResolversJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRulesRuleSettingsDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesRuleSettingsDNSResolversJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayRulesRuleSettingsDNSResolversIPV4 struct {
	// IPv4 address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                `json:"vnet_id"`
	JSON   zeroTrustGatewayRulesRuleSettingsDNSResolversIPV4JSON `json:"-"`
}

// zeroTrustGatewayRulesRuleSettingsDNSResolversIPV4JSON contains the JSON metadata
// for the struct [ZeroTrustGatewayRulesRuleSettingsDNSResolversIPV4]
type zeroTrustGatewayRulesRuleSettingsDNSResolversIPV4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ZeroTrustGatewayRulesRuleSettingsDNSResolversIPV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesRuleSettingsDNSResolversIPV4JSON) RawJSON() string {
	return r.raw
}

type ZeroTrustGatewayRulesRuleSettingsDNSResolversIPV6 struct {
	// IPv6 address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                                                `json:"vnet_id"`
	JSON   zeroTrustGatewayRulesRuleSettingsDNSResolversIPV6JSON `json:"-"`
}

// zeroTrustGatewayRulesRuleSettingsDNSResolversIPV6JSON contains the JSON metadata
// for the struct [ZeroTrustGatewayRulesRuleSettingsDNSResolversIPV6]
type zeroTrustGatewayRulesRuleSettingsDNSResolversIPV6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *ZeroTrustGatewayRulesRuleSettingsDNSResolversIPV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesRuleSettingsDNSResolversIPV6JSON) RawJSON() string {
	return r.raw
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type ZeroTrustGatewayRulesRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	IPV4 string `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 string                                      `json:"ipv6"`
	JSON zeroTrustGatewayRulesRuleSettingsEgressJSON `json:"-"`
}

// zeroTrustGatewayRulesRuleSettingsEgressJSON contains the JSON metadata for the
// struct [ZeroTrustGatewayRulesRuleSettingsEgress]
type zeroTrustGatewayRulesRuleSettingsEgressJSON struct {
	IPV4         apijson.Field
	IPV4Fallback apijson.Field
	IPV6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustGatewayRulesRuleSettingsEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesRuleSettingsEgressJSON) RawJSON() string {
	return r.raw
}

// Send matching traffic to the supplied destination IP address and port.
type ZeroTrustGatewayRulesRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                                           `json:"port"`
	JSON zeroTrustGatewayRulesRuleSettingsL4overrideJSON `json:"-"`
}

// zeroTrustGatewayRulesRuleSettingsL4overrideJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayRulesRuleSettingsL4override]
type zeroTrustGatewayRulesRuleSettingsL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRulesRuleSettingsL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesRuleSettingsL4overrideJSON) RawJSON() string {
	return r.raw
}

// Configure a notification to display on the user's device when this rule is
// matched.
type ZeroTrustGatewayRulesRuleSettingsNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                                                    `json:"support_url"`
	JSON       zeroTrustGatewayRulesRuleSettingsNotificationSettingsJSON `json:"-"`
}

// zeroTrustGatewayRulesRuleSettingsNotificationSettingsJSON contains the JSON
// metadata for the struct [ZeroTrustGatewayRulesRuleSettingsNotificationSettings]
type zeroTrustGatewayRulesRuleSettingsNotificationSettingsJSON struct {
	Enabled     apijson.Field
	Msg         apijson.Field
	SupportURL  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRulesRuleSettingsNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesRuleSettingsNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Configure DLP payload logging.
type ZeroTrustGatewayRulesRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled bool                                            `json:"enabled"`
	JSON    zeroTrustGatewayRulesRuleSettingsPayloadLogJSON `json:"-"`
}

// zeroTrustGatewayRulesRuleSettingsPayloadLogJSON contains the JSON metadata for
// the struct [ZeroTrustGatewayRulesRuleSettingsPayloadLog]
type zeroTrustGatewayRulesRuleSettingsPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRulesRuleSettingsPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesRuleSettingsPayloadLogJSON) RawJSON() string {
	return r.raw
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type ZeroTrustGatewayRulesRuleSettingsUntrustedCERT struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action ZeroTrustGatewayRulesRuleSettingsUntrustedCERTAction `json:"action"`
	JSON   zeroTrustGatewayRulesRuleSettingsUntrustedCERTJSON   `json:"-"`
}

// zeroTrustGatewayRulesRuleSettingsUntrustedCERTJSON contains the JSON metadata
// for the struct [ZeroTrustGatewayRulesRuleSettingsUntrustedCERT]
type zeroTrustGatewayRulesRuleSettingsUntrustedCERTJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustGatewayRulesRuleSettingsUntrustedCERT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesRuleSettingsUntrustedCERTJSON) RawJSON() string {
	return r.raw
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type ZeroTrustGatewayRulesRuleSettingsUntrustedCERTAction string

const (
	ZeroTrustGatewayRulesRuleSettingsUntrustedCERTActionPassThrough ZeroTrustGatewayRulesRuleSettingsUntrustedCERTAction = "pass_through"
	ZeroTrustGatewayRulesRuleSettingsUntrustedCERTActionBlock       ZeroTrustGatewayRulesRuleSettingsUntrustedCERTAction = "block"
	ZeroTrustGatewayRulesRuleSettingsUntrustedCERTActionError       ZeroTrustGatewayRulesRuleSettingsUntrustedCERTAction = "error"
)

func (r ZeroTrustGatewayRulesRuleSettingsUntrustedCERTAction) IsKnown() bool {
	switch r {
	case ZeroTrustGatewayRulesRuleSettingsUntrustedCERTActionPassThrough, ZeroTrustGatewayRulesRuleSettingsUntrustedCERTActionBlock, ZeroTrustGatewayRulesRuleSettingsUntrustedCERTActionError:
		return true
	}
	return false
}

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type ZeroTrustGatewayRulesSchedule struct {
	// The time intervals when the rule will be active on Fridays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Fridays.
	Fri string `json:"fri"`
	// The time intervals when the rule will be active on Mondays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Mondays.
	Mon string `json:"mon"`
	// The time intervals when the rule will be active on Saturdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Saturdays.
	Sat string `json:"sat"`
	// The time intervals when the rule will be active on Sundays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Sundays.
	Sun string `json:"sun"`
	// The time intervals when the rule will be active on Thursdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Thursdays.
	Thu string `json:"thu"`
	// The time zone the rule will be evaluated against. If a
	// [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)
	// is provided, Gateway will always use the current time at that time zone. If this
	// parameter is omitted, then Gateway will use the time zone inferred from the
	// user's source IP to evaluate the rule. If Gateway cannot determine the time zone
	// from the IP, we will fall back to the time zone of the user's connected data
	// center.
	TimeZone string `json:"time_zone"`
	// The time intervals when the rule will be active on Tuesdays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Tuesdays.
	Tue string `json:"tue"`
	// The time intervals when the rule will be active on Wednesdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Wednesdays.
	Wed  string                            `json:"wed"`
	JSON zeroTrustGatewayRulesScheduleJSON `json:"-"`
}

// zeroTrustGatewayRulesScheduleJSON contains the JSON metadata for the struct
// [ZeroTrustGatewayRulesSchedule]
type zeroTrustGatewayRulesScheduleJSON struct {
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

func (r *ZeroTrustGatewayRulesSchedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustGatewayRulesScheduleJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.GatewayRuleDeleteResponseUnknown] or
// [shared.UnionString].
type GatewayRuleDeleteResponseUnion interface {
	ImplementsZeroTrustGatewayRuleDeleteResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GatewayRuleDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type GatewayRuleNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action param.Field[GatewayRuleNewParamsAction] `json:"action,required"`
	// The name of the rule.
	Name param.Field[string] `json:"name,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// The wirefilter expression used for device posture check matching.
	DevicePosture param.Field[string] `json:"device_posture"`
	// True if the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters param.Field[[]GatewayRuleNewParamsFilter] `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[GatewayRuleNewParamsRuleSettings] `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule param.Field[GatewayRuleNewParamsSchedule] `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic param.Field[string] `json:"traffic"`
}

func (r GatewayRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
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
	GatewayRuleNewParamsActionAuditSSH     GatewayRuleNewParamsAction = "audit_ssh"
	GatewayRuleNewParamsActionResolve      GatewayRuleNewParamsAction = "resolve"
)

func (r GatewayRuleNewParamsAction) IsKnown() bool {
	switch r {
	case GatewayRuleNewParamsActionOn, GatewayRuleNewParamsActionOff, GatewayRuleNewParamsActionAllow, GatewayRuleNewParamsActionBlock, GatewayRuleNewParamsActionScan, GatewayRuleNewParamsActionNoscan, GatewayRuleNewParamsActionSafesearch, GatewayRuleNewParamsActionYtrestricted, GatewayRuleNewParamsActionIsolate, GatewayRuleNewParamsActionNoisolate, GatewayRuleNewParamsActionOverride, GatewayRuleNewParamsActionL4Override, GatewayRuleNewParamsActionEgress, GatewayRuleNewParamsActionAuditSSH, GatewayRuleNewParamsActionResolve:
		return true
	}
	return false
}

// The protocol or layer to use.
type GatewayRuleNewParamsFilter string

const (
	GatewayRuleNewParamsFilterHTTP   GatewayRuleNewParamsFilter = "http"
	GatewayRuleNewParamsFilterDNS    GatewayRuleNewParamsFilter = "dns"
	GatewayRuleNewParamsFilterL4     GatewayRuleNewParamsFilter = "l4"
	GatewayRuleNewParamsFilterEgress GatewayRuleNewParamsFilter = "egress"
)

func (r GatewayRuleNewParamsFilter) IsKnown() bool {
	switch r {
	case GatewayRuleNewParamsFilterHTTP, GatewayRuleNewParamsFilterDNS, GatewayRuleNewParamsFilterL4, GatewayRuleNewParamsFilterEgress:
		return true
	}
	return false
}

// Additional settings that modify the rule's action.
type GatewayRuleNewParamsRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders param.Field[interface{}] `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass param.Field[bool] `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH param.Field[GatewayRuleNewParamsRuleSettingsAuditSSH] `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls param.Field[GatewayRuleNewParamsRuleSettingsBisoAdminControls] `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled param.Field[bool] `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason param.Field[string] `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule param.Field[bool] `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession param.Field[GatewayRuleNewParamsRuleSettingsCheckSession] `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin. Only valid when a rule's action is
	// set to 'resolve'.
	DNSResolvers param.Field[GatewayRuleNewParamsRuleSettingsDNSResolvers] `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress param.Field[GatewayRuleNewParamsRuleSettingsEgress] `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation param.Field[bool] `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories param.Field[bool] `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds param.Field[bool] `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override param.Field[GatewayRuleNewParamsRuleSettingsL4override] `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings param.Field[GatewayRuleNewParamsRuleSettingsNotificationSettings] `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost param.Field[string] `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs param.Field[[]string] `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog param.Field[GatewayRuleNewParamsRuleSettingsPayloadLog] `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified. Only valid when a
	// rule's action is set to 'resolve'.
	ResolveDNSThroughCloudflare param.Field[bool] `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCERT param.Field[GatewayRuleNewParamsRuleSettingsUntrustedCERT] `json:"untrusted_cert"`
}

func (r GatewayRuleNewParamsRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for the Audit SSH action.
type GatewayRuleNewParamsRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging param.Field[bool] `json:"command_logging"`
}

func (r GatewayRuleNewParamsRuleSettingsAuditSSH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how browser isolation behaves.
type GatewayRuleNewParamsRuleSettingsBisoAdminControls struct {
	// Set to true to enable copy-pasting.
	Dcp param.Field[bool] `json:"dcp"`
	// Set to true to enable downloading.
	Dd param.Field[bool] `json:"dd"`
	// Set to true to enable keyboard usage.
	Dk param.Field[bool] `json:"dk"`
	// Set to true to enable printing.
	Dp param.Field[bool] `json:"dp"`
	// Set to true to enable uploading.
	Du param.Field[bool] `json:"du"`
}

func (r GatewayRuleNewParamsRuleSettingsBisoAdminControls) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how session check behaves.
type GatewayRuleNewParamsRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration param.Field[string] `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r GatewayRuleNewParamsRuleSettingsCheckSession) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin. Only valid when a rule's action is
// set to 'resolve'.
type GatewayRuleNewParamsRuleSettingsDNSResolvers struct {
	IPV4 param.Field[[]GatewayRuleNewParamsRuleSettingsDNSResolversIPV4] `json:"ipv4"`
	IPV6 param.Field[[]GatewayRuleNewParamsRuleSettingsDNSResolversIPV6] `json:"ipv6"`
}

func (r GatewayRuleNewParamsRuleSettingsDNSResolvers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleNewParamsRuleSettingsDNSResolversIPV4 struct {
	// IPv4 address of upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port param.Field[int64] `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r GatewayRuleNewParamsRuleSettingsDNSResolversIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleNewParamsRuleSettingsDNSResolversIPV6 struct {
	// IPv6 address of upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port param.Field[int64] `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r GatewayRuleNewParamsRuleSettingsDNSResolversIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type GatewayRuleNewParamsRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	IPV4 param.Field[string] `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback param.Field[string] `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 param.Field[string] `json:"ipv6"`
}

func (r GatewayRuleNewParamsRuleSettingsEgress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Send matching traffic to the supplied destination IP address and port.
type GatewayRuleNewParamsRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP param.Field[string] `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port param.Field[int64] `json:"port"`
}

func (r GatewayRuleNewParamsRuleSettingsL4override) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type GatewayRuleNewParamsRuleSettingsNotificationSettings struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r GatewayRuleNewParamsRuleSettingsNotificationSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure DLP payload logging.
type GatewayRuleNewParamsRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayRuleNewParamsRuleSettingsPayloadLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type GatewayRuleNewParamsRuleSettingsUntrustedCERT struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action param.Field[GatewayRuleNewParamsRuleSettingsUntrustedCERTAction] `json:"action"`
}

func (r GatewayRuleNewParamsRuleSettingsUntrustedCERT) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type GatewayRuleNewParamsRuleSettingsUntrustedCERTAction string

const (
	GatewayRuleNewParamsRuleSettingsUntrustedCERTActionPassThrough GatewayRuleNewParamsRuleSettingsUntrustedCERTAction = "pass_through"
	GatewayRuleNewParamsRuleSettingsUntrustedCERTActionBlock       GatewayRuleNewParamsRuleSettingsUntrustedCERTAction = "block"
	GatewayRuleNewParamsRuleSettingsUntrustedCERTActionError       GatewayRuleNewParamsRuleSettingsUntrustedCERTAction = "error"
)

func (r GatewayRuleNewParamsRuleSettingsUntrustedCERTAction) IsKnown() bool {
	switch r {
	case GatewayRuleNewParamsRuleSettingsUntrustedCERTActionPassThrough, GatewayRuleNewParamsRuleSettingsUntrustedCERTActionBlock, GatewayRuleNewParamsRuleSettingsUntrustedCERTActionError:
		return true
	}
	return false
}

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type GatewayRuleNewParamsSchedule struct {
	// The time intervals when the rule will be active on Fridays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Fridays.
	Fri param.Field[string] `json:"fri"`
	// The time intervals when the rule will be active on Mondays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Mondays.
	Mon param.Field[string] `json:"mon"`
	// The time intervals when the rule will be active on Saturdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Saturdays.
	Sat param.Field[string] `json:"sat"`
	// The time intervals when the rule will be active on Sundays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Sundays.
	Sun param.Field[string] `json:"sun"`
	// The time intervals when the rule will be active on Thursdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Thursdays.
	Thu param.Field[string] `json:"thu"`
	// The time zone the rule will be evaluated against. If a
	// [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)
	// is provided, Gateway will always use the current time at that time zone. If this
	// parameter is omitted, then Gateway will use the time zone inferred from the
	// user's source IP to evaluate the rule. If Gateway cannot determine the time zone
	// from the IP, we will fall back to the time zone of the user's connected data
	// center.
	TimeZone param.Field[string] `json:"time_zone"`
	// The time intervals when the rule will be active on Tuesdays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Tuesdays.
	Tue param.Field[string] `json:"tue"`
	// The time intervals when the rule will be active on Wednesdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Wednesdays.
	Wed param.Field[string] `json:"wed"`
}

func (r GatewayRuleNewParamsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ZeroTrustGatewayRules `json:"result,required"`
	// Whether the API call was successful
	Success GatewayRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayRuleNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleNewResponseEnvelope]
type gatewayRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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
	// The action to preform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action param.Field[GatewayRuleUpdateParamsAction] `json:"action,required"`
	// The name of the rule.
	Name param.Field[string] `json:"name,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// The wirefilter expression used for device posture check matching.
	DevicePosture param.Field[string] `json:"device_posture"`
	// True if the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters param.Field[[]GatewayRuleUpdateParamsFilter] `json:"filters"`
	// The wirefilter expression used for identity matching.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[GatewayRuleUpdateParamsRuleSettings] `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule param.Field[GatewayRuleUpdateParamsSchedule] `json:"schedule"`
	// The wirefilter expression used for traffic matching.
	Traffic param.Field[string] `json:"traffic"`
}

func (r GatewayRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to preform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
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
	GatewayRuleUpdateParamsActionAuditSSH     GatewayRuleUpdateParamsAction = "audit_ssh"
	GatewayRuleUpdateParamsActionResolve      GatewayRuleUpdateParamsAction = "resolve"
)

func (r GatewayRuleUpdateParamsAction) IsKnown() bool {
	switch r {
	case GatewayRuleUpdateParamsActionOn, GatewayRuleUpdateParamsActionOff, GatewayRuleUpdateParamsActionAllow, GatewayRuleUpdateParamsActionBlock, GatewayRuleUpdateParamsActionScan, GatewayRuleUpdateParamsActionNoscan, GatewayRuleUpdateParamsActionSafesearch, GatewayRuleUpdateParamsActionYtrestricted, GatewayRuleUpdateParamsActionIsolate, GatewayRuleUpdateParamsActionNoisolate, GatewayRuleUpdateParamsActionOverride, GatewayRuleUpdateParamsActionL4Override, GatewayRuleUpdateParamsActionEgress, GatewayRuleUpdateParamsActionAuditSSH, GatewayRuleUpdateParamsActionResolve:
		return true
	}
	return false
}

// The protocol or layer to use.
type GatewayRuleUpdateParamsFilter string

const (
	GatewayRuleUpdateParamsFilterHTTP   GatewayRuleUpdateParamsFilter = "http"
	GatewayRuleUpdateParamsFilterDNS    GatewayRuleUpdateParamsFilter = "dns"
	GatewayRuleUpdateParamsFilterL4     GatewayRuleUpdateParamsFilter = "l4"
	GatewayRuleUpdateParamsFilterEgress GatewayRuleUpdateParamsFilter = "egress"
)

func (r GatewayRuleUpdateParamsFilter) IsKnown() bool {
	switch r {
	case GatewayRuleUpdateParamsFilterHTTP, GatewayRuleUpdateParamsFilterDNS, GatewayRuleUpdateParamsFilterL4, GatewayRuleUpdateParamsFilterEgress:
		return true
	}
	return false
}

// Additional settings that modify the rule's action.
type GatewayRuleUpdateParamsRuleSettings struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders param.Field[interface{}] `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass param.Field[bool] `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH param.Field[GatewayRuleUpdateParamsRuleSettingsAuditSSH] `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BisoAdminControls param.Field[GatewayRuleUpdateParamsRuleSettingsBisoAdminControls] `json:"biso_admin_controls"`
	// Enable the custom block page.
	BlockPageEnabled param.Field[bool] `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason param.Field[string] `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule param.Field[bool] `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession param.Field[GatewayRuleUpdateParamsRuleSettingsCheckSession] `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
	// route to the address closest to their origin. Only valid when a rule's action is
	// set to 'resolve'.
	DNSResolvers param.Field[GatewayRuleUpdateParamsRuleSettingsDNSResolvers] `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress param.Field[GatewayRuleUpdateParamsRuleSettingsEgress] `json:"egress"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation param.Field[bool] `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories param.Field[bool] `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds param.Field[bool] `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override param.Field[GatewayRuleUpdateParamsRuleSettingsL4override] `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings param.Field[GatewayRuleUpdateParamsRuleSettingsNotificationSettings] `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost param.Field[string] `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs param.Field[[]string] `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog param.Field[GatewayRuleUpdateParamsRuleSettingsPayloadLog] `json:"payload_log"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when dns_resolvers are specified. Only valid when a
	// rule's action is set to 'resolve'.
	ResolveDNSThroughCloudflare param.Field[bool] `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCERT param.Field[GatewayRuleUpdateParamsRuleSettingsUntrustedCERT] `json:"untrusted_cert"`
}

func (r GatewayRuleUpdateParamsRuleSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for the Audit SSH action.
type GatewayRuleUpdateParamsRuleSettingsAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging param.Field[bool] `json:"command_logging"`
}

func (r GatewayRuleUpdateParamsRuleSettingsAuditSSH) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how browser isolation behaves.
type GatewayRuleUpdateParamsRuleSettingsBisoAdminControls struct {
	// Set to true to enable copy-pasting.
	Dcp param.Field[bool] `json:"dcp"`
	// Set to true to enable downloading.
	Dd param.Field[bool] `json:"dd"`
	// Set to true to enable keyboard usage.
	Dk param.Field[bool] `json:"dk"`
	// Set to true to enable printing.
	Dp param.Field[bool] `json:"dp"`
	// Set to true to enable uploading.
	Du param.Field[bool] `json:"du"`
}

func (r GatewayRuleUpdateParamsRuleSettingsBisoAdminControls) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how session check behaves.
type GatewayRuleUpdateParamsRuleSettingsCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid.
	Duration param.Field[string] `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r GatewayRuleUpdateParamsRuleSettingsCheckSession) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when resolve_dns_through_cloudflare is set. DNS queries will
// route to the address closest to their origin. Only valid when a rule's action is
// set to 'resolve'.
type GatewayRuleUpdateParamsRuleSettingsDNSResolvers struct {
	IPV4 param.Field[[]GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV4] `json:"ipv4"`
	IPV6 param.Field[[]GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV6] `json:"ipv6"`
}

func (r GatewayRuleUpdateParamsRuleSettingsDNSResolvers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV4 struct {
	// IPv4 address of upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port param.Field[int64] `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV4) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV6 struct {
	// IPv6 address of upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port param.Field[int64] `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r GatewayRuleUpdateParamsRuleSettingsDNSResolversIPV6) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type GatewayRuleUpdateParamsRuleSettingsEgress struct {
	// The IPv4 address to be used for egress.
	IPV4 param.Field[string] `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback param.Field[string] `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 param.Field[string] `json:"ipv6"`
}

func (r GatewayRuleUpdateParamsRuleSettingsEgress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Send matching traffic to the supplied destination IP address and port.
type GatewayRuleUpdateParamsRuleSettingsL4override struct {
	// IPv4 or IPv6 address.
	IP param.Field[string] `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port param.Field[int64] `json:"port"`
}

func (r GatewayRuleUpdateParamsRuleSettingsL4override) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type GatewayRuleUpdateParamsRuleSettingsNotificationSettings struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r GatewayRuleUpdateParamsRuleSettingsNotificationSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure DLP payload logging.
type GatewayRuleUpdateParamsRuleSettingsPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r GatewayRuleUpdateParamsRuleSettingsPayloadLog) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type GatewayRuleUpdateParamsRuleSettingsUntrustedCERT struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action param.Field[GatewayRuleUpdateParamsRuleSettingsUntrustedCERTAction] `json:"action"`
}

func (r GatewayRuleUpdateParamsRuleSettingsUntrustedCERT) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type GatewayRuleUpdateParamsRuleSettingsUntrustedCERTAction string

const (
	GatewayRuleUpdateParamsRuleSettingsUntrustedCERTActionPassThrough GatewayRuleUpdateParamsRuleSettingsUntrustedCERTAction = "pass_through"
	GatewayRuleUpdateParamsRuleSettingsUntrustedCERTActionBlock       GatewayRuleUpdateParamsRuleSettingsUntrustedCERTAction = "block"
	GatewayRuleUpdateParamsRuleSettingsUntrustedCERTActionError       GatewayRuleUpdateParamsRuleSettingsUntrustedCERTAction = "error"
)

func (r GatewayRuleUpdateParamsRuleSettingsUntrustedCERTAction) IsKnown() bool {
	switch r {
	case GatewayRuleUpdateParamsRuleSettingsUntrustedCERTActionPassThrough, GatewayRuleUpdateParamsRuleSettingsUntrustedCERTActionBlock, GatewayRuleUpdateParamsRuleSettingsUntrustedCERTActionError:
		return true
	}
	return false
}

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type GatewayRuleUpdateParamsSchedule struct {
	// The time intervals when the rule will be active on Fridays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Fridays.
	Fri param.Field[string] `json:"fri"`
	// The time intervals when the rule will be active on Mondays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Mondays.
	Mon param.Field[string] `json:"mon"`
	// The time intervals when the rule will be active on Saturdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Saturdays.
	Sat param.Field[string] `json:"sat"`
	// The time intervals when the rule will be active on Sundays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Sundays.
	Sun param.Field[string] `json:"sun"`
	// The time intervals when the rule will be active on Thursdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Thursdays.
	Thu param.Field[string] `json:"thu"`
	// The time zone the rule will be evaluated against. If a
	// [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)
	// is provided, Gateway will always use the current time at that time zone. If this
	// parameter is omitted, then Gateway will use the time zone inferred from the
	// user's source IP to evaluate the rule. If Gateway cannot determine the time zone
	// from the IP, we will fall back to the time zone of the user's connected data
	// center.
	TimeZone param.Field[string] `json:"time_zone"`
	// The time intervals when the rule will be active on Tuesdays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Tuesdays.
	Tue param.Field[string] `json:"tue"`
	// The time intervals when the rule will be active on Wednesdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Wednesdays.
	Wed param.Field[string] `json:"wed"`
}

func (r GatewayRuleUpdateParamsSchedule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ZeroTrustGatewayRules `json:"result,required"`
	// Whether the API call was successful
	Success GatewayRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleUpdateResponseEnvelope]
type gatewayRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      param.Field[interface{}] `json:"body,required"`
}

func (r GatewayRuleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type GatewayRuleDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   GatewayRuleDeleteResponseUnion `json:"result,required"`
	// Whether the API call was successful
	Success GatewayRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleDeleteResponseEnvelope]
type gatewayRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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
	Result   ZeroTrustGatewayRules `json:"result,required"`
	// Whether the API call was successful
	Success GatewayRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    gatewayRuleGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleGetResponseEnvelope]
type gatewayRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
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
